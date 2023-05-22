#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from ctypes import c_uint8
from enum import IntEnum


class ModbusErrorCode(IntEnum):
    ILLEGAL_FUNCTION: c_uint8 = 1
    ILLEGAL_DATA_ADDRESS: c_uint8 = 2
    ILLEGAL_DATA_VALUE: c_uint8 = 3
    SLAVE_DEVICE_FAILURE: c_uint8 = 4
    ACKNOWLEDGE: c_uint8 = 5
    SLAVE_DEVICE_BUSY: c_uint8 = 6
    NEGATIVE_ACKNOWLEDGE: c_uint8 = 7
    MEMORY_PARITY_ERROR: c_uint8 = 8
    GATEWAY_PATH_UNAVAILABLE: c_uint8 = 10
    GATEWAY_TARGET_DEVICE_FAILED_TO_RESPOND: c_uint8 = 11
