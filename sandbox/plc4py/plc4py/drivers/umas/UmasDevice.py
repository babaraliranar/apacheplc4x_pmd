#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#
import asyncio
import logging
from asyncio import Transport, AbstractEventLoop
from dataclasses import dataclass, field
from typing import Dict, List

from plc4py.drivers.umas.UmasTag import UmasTag

from plc4py.protocols.umas.readwrite.UmasPDUReadVariableResponse import (
    UmasPDUReadVariableResponse,
)

from plc4py.protocols.umas.readwrite.UmasDataType import UmasDataType

from plc4py.protocols.umas.readwrite.DataItem import DataItem
from plc4py.protocols.umas.readwrite.VariableRequestReference import (
    VariableRequestReference,
)

from plc4py.protocols.umas.readwrite.UmasPDUReadVariableRequest import (
    UmasPDUReadVariableRequestBuilder,
)

from plc4py.protocols.umas.readwrite import UmasPDUReadUnlocatedVariableResponse
from plc4py.protocols.umas.readwrite.UmasPDUReadDatatypeNamesResponse import (
    UmasPDUReadDatatypeNamesResponse,
)
from plc4py.protocols.umas.readwrite.UmasPDUReadUnlocatedVariableNamesResponse import (
    UmasPDUReadUnlocatedVariableNamesResponse,
)

from plc4py.protocols.umas.readwrite.UmasPDUReadUnlocatedVariableNamesRequest import (
    UmasPDUReadUnlocatedVariableNamesRequestBuilder,
)

from plc4py.protocols.umas.readwrite.UmasPDUReadUnlocatedVariableNames import (
    UmasPDUReadUnlocatedVariableNamesBuilder,
)

from plc4py.protocols.umas.readwrite.UmasMemoryBlockBasicInfo import (
    UmasMemoryBlockBasicInfo,
)
from plc4py.spi.generation.ReadBuffer import ReadBufferByteBased

from plc4py.protocols.umas.readwrite.UmasPDUReadMemoryBlockResponse import (
    UmasPDUReadMemoryBlockResponse,
)

from plc4py.protocols.umas.readwrite.PlcMemoryBlockIdent import PlcMemoryBlockIdent
from plc4py.protocols.umas.readwrite.UmasPDUPlcIdentResponse import (
    UmasPDUPlcIdentResponse,
)

from plc4py.protocols.umas.readwrite.UmasPDUReadMemoryBlockRequest import (
    UmasPDUReadMemoryBlockRequestBuilder,
)

from plc4py.protocols.umas.readwrite.UmasInitCommsResponse import UmasInitCommsResponse

from plc4py.protocols.umas.readwrite.UmasPDUPlcStatusResponse import (
    UmasPDUPlcStatusResponse,
)

from plc4py.protocols.umas.readwrite.UmasPDUPlcStatusRequest import (
    UmasPDUPlcStatusRequestBuilder,
)

from plc4py.protocols.umas.readwrite.UmasPDUProjectInfoRequest import (
    UmasPDUProjectInfoRequest,
    UmasPDUProjectInfoRequestBuilder,
)

from plc4py.protocols.umas.readwrite.UmasInitCommsRequest import (
    UmasInitCommsRequestBuilder,
)

from plc4py.protocols.umas.readwrite.UmasPDUPlcIdentRequest import (
    UmasPDUPlcIdentRequestBuilder,
)
from plc4py.spi.generation.WriteBuffer import WriteBufferByteBased

from plc4py.api.messages.PlcRequest import PlcReadRequest
from plc4py.api.messages.PlcResponse import PlcReadResponse
from plc4py.api.value.PlcValue import PlcValue, PlcResponseCode
from plc4py.drivers.umas.UmasConfiguration import UmasConfiguration
from plc4py.spi.messages.utils.ResponseItem import ResponseItem
from plc4py.utils.GenericTypes import ByteOrder


@dataclass
class UmasDevice:
    _configuration: UmasConfiguration
    tags: Dict[str, PlcValue] = field(default_factory=lambda: {})
    project_crc: int = -1
    max_frame_size: int = -1
    memory_blocks: List[PlcMemoryBlockIdent] = field(default_factory=lambda: [])
    hardware_id: int = -1
    index: int = -1

    async def connect(self, transport: Transport):
        # Create future to be returned when a value is returned
        loop = asyncio.get_running_loop()
        await self._send_plc_ident(transport, loop)
        await self._send_init_comms(transport, loop)
        await self._update_plc_project_info(transport, loop)

    async def _update_plc_project_info(self, transport, loop):
        await self._send_project_info(transport, loop)
        await self._send_read_memory_block(transport, loop)
        await self._send_unlocated_variable_datatype_request(transport, loop)
        await self._send_unlocated_variable_request(transport, loop)

    async def _send_plc_ident(self, transport: Transport, loop: AbstractEventLoop):
        message_future = loop.create_future()

        request_pdu = UmasPDUPlcIdentRequestBuilder().build(0, -1)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        ident_result: UmasPDUPlcIdentResponse = message_future.result()
        self.memory_blocks = ident_result.memory_idents

    async def _send_init_comms(self, transport: Transport, loop: AbstractEventLoop):
        message_future = loop.create_future()

        request_pdu = UmasInitCommsRequestBuilder(0).build(0, -1)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        init_result: UmasInitCommsResponse = message_future.result()
        self.max_frame_size = init_result.max_frame_size

    async def _send_project_info(self, transport: Transport, loop: AbstractEventLoop):
        message_future = loop.create_future()

        request_pdu = UmasPDUPlcStatusRequestBuilder().build(0, -1)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        project_info_result: UmasPDUPlcStatusResponse = message_future.result()
        if project_info_result.number_of_blocks > 3:
            self.project_crc = (
                project_info_result.blocks[3] + project_info_result.blocks[4]
            )

    async def _send_read_memory_block(
        self, transport: Transport, loop: AbstractEventLoop
    ):
        message_future = loop.create_future()

        request_pdu = UmasPDUReadMemoryBlockRequestBuilder(
            range=0,
            block_number=0x30,
            offset=0,
            number_of_bytes=0x21,
            unknown_object1=0,
        ).build(pairing_key=0, byte_length=0)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        memory_block_result: UmasPDUReadMemoryBlockResponse = message_future.result()
        read_buffer = ReadBufferByteBased(
            bytearray(memory_block_result.block),
            ByteOrder.LITTLE_ENDIAN,
        )
        basic_info = UmasMemoryBlockBasicInfo.static_parse_builder(
            read_buffer, 0x30, 0x00
        ).build()
        self.hardware_id = basic_info.hardware_id
        self.index = basic_info.index
        pass

    async def _send_unlocated_variable_datatype_request(
        self, transport: Transport, loop: AbstractEventLoop
    ):
        message_future = loop.create_future()

        request_pdu = UmasPDUReadUnlocatedVariableNamesRequestBuilder(
            record_type=0xDD03,
            block_no=0x0000,
            index=self.index,
            hardware_id=self.hardware_id,
        ).build(0, 0)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        data_type_response: UmasPDUReadUnlocatedVariableResponse = (
            message_future.result()
        )
        read_buffer = ReadBufferByteBased(
            bytearray(data_type_response.block),
            ByteOrder.BIG_ENDIAN,
            ByteOrder.BIG_ENDIAN,
        )
        basic_info = UmasPDUReadDatatypeNamesResponse.static_parse_builder(
            read_buffer, 0xDD03
        ).build()

    async def _send_unlocated_variable_request(
        self, transport: Transport, loop: AbstractEventLoop
    ):
        message_future = loop.create_future()

        request_pdu = UmasPDUReadUnlocatedVariableNamesRequestBuilder(
            record_type=0xDD02,
            block_no=0xFFFF,
            index=self.index,
            hardware_id=self.hardware_id,
        ).build(0, 0)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        variable_name_response: UmasPDUReadUnlocatedVariableResponse = (
            message_future.result()
        )
        read_buffer = ReadBufferByteBased(
            bytearray(variable_name_response.block), ByteOrder.LITTLE_ENDIAN
        )
        self.variables = UmasPDUReadUnlocatedVariableNamesResponse.static_parse_builder(
            read_buffer, 0xDD03
        ).build()
        pass

    async def _send_read_variable_request(
        self, transport: Transport, loop: AbstractEventLoop, request: PlcReadRequest
    ):
        message_future = loop.create_future()

        tag_list: List[VariableRequestReference] = []
        for kea, tag in request.tags.items():
            found_record = None
            for record in self.variables.records:
                if tag.tag_name.lower() == record.value.lower():
                    found_record = record
                    break

            if found_record is not None:
                tag_list.append(
                    VariableRequestReference(
                        data_type=UmasDataType(record.data_type).request_size,
                        block=found_record.block,
                        base_offset=0x0000,
                        offset=found_record.offset,
                    )
                )
        request_pdu = UmasPDUReadVariableRequestBuilder(
            crc=self.project_crc,
            variable_count=len(tag_list),
            variables=tag_list,
        ).build(0, 0)

        protocol = transport.protocol
        protocol.write_wait_for_response(
            request_pdu,
            transport,
            message_future,
        )

        await message_future
        variable_name_response: UmasPDUReadVariableResponse = message_future.result()
        read_buffer = ReadBufferByteBased(
            bytearray(variable_name_response.block), ByteOrder.LITTLE_ENDIAN
        )
        values: Dict[str, List[ResponseItem[PlcValue]]] = {}
        for key, tag in request.tags.items():
            response_items = [
                ResponseItem(
                    PlcResponseCode.OK,
                    DataItem.static_parse(read_buffer, tag.data_type, tag.quantity),
                )
            ]
            values[key] = response_items

        response = PlcReadResponse(PlcResponseCode.OK, values)
        return response

    async def read(
        self, request: PlcReadRequest, transport: Transport
    ) -> PlcReadResponse:
        """
        Reads one field from the Umas Device
        """
        loop = asyncio.get_running_loop()
        response = await self._send_read_variable_request(transport, loop, request)
        return response
