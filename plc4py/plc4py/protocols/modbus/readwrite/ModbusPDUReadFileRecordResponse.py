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

from dataclasses import dataclass

from plc4py.api.exceptions.exceptions import PlcRuntimeException
from plc4py.api.exceptions.exceptions import SerializationException
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDUReadFileRecordResponseItem import (
    ModbusPDUReadFileRecordResponseItem,
)
from plc4py.spi.generation.ReadBuffer import ReadBuffer
from plc4py.spi.generation.WriteBuffer import WriteBuffer
from plc4py.spi.values.Common import get_size_of_array
from plc4py.utils.ConnectionStringHandling import strtobool
from typing import Any
from typing import ClassVar
from typing import List
import math


@dataclass
class ModbusPDUReadFileRecordResponse(ModbusPDU):
    items: List[ModbusPDUReadFileRecordResponseItem]
    # Accessors for discriminator values.
    error_flag: ClassVar[bool] = False
    function_flag: ClassVar[int] = 0x14
    response: ClassVar[bool] = True

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        write_buffer.push_context("ModbusPDUReadFileRecordResponse")

        # Implicit Field (byte_count) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        byte_count: int = int(get_size_of_array(self.items))
        write_buffer.write_unsigned_byte(byte_count, logical_name="byte_count")

        # Array Field (items)
        write_buffer.write_complex_array(self.items, logical_name="items")

        write_buffer.pop_context("ModbusPDUReadFileRecordResponse")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.length_in_bits() / 8.0)))

    def length_in_bits(self) -> int:
        length_in_bits: int = super().length_in_bits()
        _value: ModbusPDUReadFileRecordResponse = self

        # Implicit Field (byteCount)
        length_in_bits += 8

        # Array field
        if self.items is not None:
            for element in self.items:
                length_in_bits += element.length_in_bits()

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: bool):
        read_buffer.push_context("ModbusPDUReadFileRecordResponse")

        if isinstance(response, str):
            response = bool(strtobool(response))

        byte_count: int = read_buffer.read_unsigned_byte(
            logical_name="byte_count", response=response
        )

        items: List[Any] = read_buffer.read_array_field(
            logical_name="items",
            read_function=ModbusPDUReadFileRecordResponseItem.static_parse,
            length=byte_count,
            response=response,
        )

        read_buffer.pop_context("ModbusPDUReadFileRecordResponse")
        # Create the instance
        return ModbusPDUReadFileRecordResponseBuilder(items)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadFileRecordResponse):
            return False

        that: ModbusPDUReadFileRecordResponse = ModbusPDUReadFileRecordResponse(o)
        return (self.items == that.items) and super().equals(that) and True

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        # TODO:- Implement a generic python object to probably json convertor or something.
        return ""


@dataclass
class ModbusPDUReadFileRecordResponseBuilder:
    items: List[ModbusPDUReadFileRecordResponseItem]

    def build(
        self,
    ) -> ModbusPDUReadFileRecordResponse:
        modbus_pduread_file_record_response: ModbusPDUReadFileRecordResponse = (
            ModbusPDUReadFileRecordResponse(self.items)
        )
        return modbus_pduread_file_record_response
