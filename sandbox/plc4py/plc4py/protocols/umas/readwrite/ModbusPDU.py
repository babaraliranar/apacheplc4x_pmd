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

from abc import ABC
from abc import abstractmethod
from plc4py.api.exceptions.exceptions import ParseException
from plc4py.api.exceptions.exceptions import PlcRuntimeException
from plc4py.api.exceptions.exceptions import SerializationException
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.spi.generation.ReadBuffer import ReadBuffer
from plc4py.spi.generation.WriteBuffer import WriteBuffer
import math


@dataclass
class ModbusPDU(ABC, PlcMessage):
    # Arguments.
    response: bool
    length: int

    # Abstract accessors for discriminator values.
    @property
    @abstractmethod
    def error_flag(self) -> bool:
        pass

    @property
    @abstractmethod
    def function_flag(self) -> int:
        pass

    @abstractmethod
    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer) -> None:
        pass

    def serialize(self, write_buffer: WriteBuffer):
        write_buffer.push_context("ModbusPDU")

        # Discriminator Field (errorFlag) (Used as input to a switch field)
        write_buffer.write_bit(
            self.error_flag,
            logical_name="errorFlag",
            bit_length=1,
        )

        # Discriminator Field (functionFlag) (Used as input to a switch field)
        write_buffer.write_unsigned_byte(
            self.function_flag,
            logical_name="functionFlag",
            bit_length=7,
        )

        # Switch field (Serialize the sub-type)
        self.serialize_modbus_pdu_child(write_buffer)

        write_buffer.pop_context("ModbusPDU")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.length_in_bits() / 8.0)))

    def length_in_bits(self) -> int:
        length_in_bits: int = 0
        _value: ModbusPDU = self

        # Discriminator Field (errorFlag)
        length_in_bits += 1

        # Discriminator Field (functionFlag)
        length_in_bits += 7

        # Length of subtype elements will be added by sub-type...

        return length_in_bits

    @staticmethod
    def static_parse(read_buffer: ReadBuffer, **kwargs):
        if kwargs is None:
            raise PlcRuntimeException(
                "Wrong number of arguments, expected 2, but got None"
            )

        response: bool = False
        if isinstance(kwargs.get("response"), bool):
            response = bool(kwargs.get("response"))
        elif isinstance(kwargs.get("response"), str):
            response = bool(str(kwargs.get("response")))
        else:
            raise PlcRuntimeException(
                "Argument 0 expected to be of type bool or a string which is parseable but was "
                + kwargs.get("response").getClass().getName()
            )

        length: int = 0
        if isinstance(kwargs.get("length"), int):
            length = int(kwargs.get("length"))
        elif isinstance(kwargs.get("length"), str):
            length = int(str(kwargs.get("length")))
        else:
            raise PlcRuntimeException(
                "Argument 1 expected to be of type int or a string which is parseable but was "
                + kwargs.get("length").getClass().getName()
            )

        return ModbusPDU.static_parse_context(read_buffer, response, length)

    @staticmethod
    def static_parse_context(read_buffer: ReadBuffer, response: bool, length: int):
        read_buffer.push_context("ModbusPDU")
        error_flag: bool = read_buffer.read_bit(
            logical_name="errorFlag", bit_length=1, response=response, length=length
        )

        function_flag: int = read_buffer.read_unsigned_byte(
            logical_name="functionFlag", bit_length=7, response=response, length=length
        )

        # Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
        builder: ModbusPDUBuilder = None
        from plc4py.protocols.umas.readwrite.ModbusPDUError import ModbusPDUError

        if error_flag == bool(True):
            builder = ModbusPDUError.static_parse_builder(read_buffer, response, length)
        from plc4py.protocols.umas.readwrite.UmasPDU import UmasPDU

        if error_flag == bool(False) and function_flag == int(0x5A):
            builder = UmasPDU.static_parse_builder(read_buffer, response, length)
        if builder is None:
            raise ParseException(
                "Unsupported case for discriminated type"
                + " parameters ["
                + "errorFlag="
                + str(error_flag)
                + " "
                + "functionFlag="
                + str(function_flag)
                + "]"
            )

        read_buffer.pop_context("ModbusPDU")
        # Create the instance
        _modbus_pdu: ModbusPDU = builder.build(response, length)
        return _modbus_pdu

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDU):
            return False

        that: ModbusPDU = ModbusPDU(o)
        return True

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        pass
        # write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        # try:
        #    write_buffer_box_based.writeSerializable(self)
        # except SerializationException as e:
        #    raise PlcRuntimeException(e)

        # return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class ModbusPDUBuilder:
    def build(self, response: bool, length: int) -> ModbusPDU:
        pass
