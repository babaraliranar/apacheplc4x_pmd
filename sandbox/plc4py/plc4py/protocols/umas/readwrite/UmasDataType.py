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
from aenum import AutoNumberEnum


class UmasDataType(AutoNumberEnum):
    _init_ = "value, data_type_size"
    BOOL = (1, int(2))
    BYTE = (2, int(2))
    WORD = (3, int(2))
    DWORD = (4, int(4))
    LWORD = (5, int(8))
    SINT = (6, int(2))
    INT = (7, int(2))
    DINT = (8, int(4))
    LINT = (9, int(8))
    USINT = (10, int(2))
    UINT = (11, int(2))
    UDINT = (12, int(4))
    ULINT = (13, int(8))
    REAL = (14, int(4))
    LREAL = (15, int(8))
    TIME = (16, int(8))
    LTIME = (17, int(8))
    DATE = (18, int(8))
    LDATE = (19, int(8))
    TIME_OF_DAY = (20, int(8))
    LTIME_OF_DAY = (21, int(8))
    DATE_AND_TIME = (22, int(8))
    LDATE_AND_TIME = (23, int(8))
    CHAR = (24, int(1))
    WCHAR = (25, int(2))
    STRING = (26, int(1))
    WSTRING = (27, int(2))
