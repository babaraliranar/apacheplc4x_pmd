/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#ifndef PLC4C_MODBUS_READ_WRITE_MODBUS_ERROR_CODE_H_
#define PLC4C_MODBUS_READ_WRITE_MODBUS_ERROR_CODE_H_

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

enum plc4c_modbus_read_write_modbus_error_code {
  plc4c_modbus_read_write_modbus_error_code_ILLEGAL_FUNCTION = 1,
  plc4c_modbus_read_write_modbus_error_code_ILLEGAL_DATA_ADDRESS = 2,
  plc4c_modbus_read_write_modbus_error_code_ILLEGAL_DATA_VALUE = 3,
  plc4c_modbus_read_write_modbus_error_code_SLAVE_DEVICE_FAILURE = 4,
  plc4c_modbus_read_write_modbus_error_code_ACKNOWLEDGE = 5,
  plc4c_modbus_read_write_modbus_error_code_SLAVE_DEVICE_BUSY = 6,
  plc4c_modbus_read_write_modbus_error_code_NEGATIVE_ACKNOWLEDGE = 7,
  plc4c_modbus_read_write_modbus_error_code_MEMORY_PARITY_ERROR = 8,
  plc4c_modbus_read_write_modbus_error_code_GATEWAY_PATH_UNAVAILABLE = 10,
  plc4c_modbus_read_write_modbus_error_code_GATEWAY_TARGET_DEVICE_FAILED_TO_RESPOND = 11
};
typedef enum plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code;

// Get an empty NULL-struct
plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_null();

plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_value_of(char* value_string);

int plc4c_modbus_read_write_modbus_error_code_num_values();

plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_value_for_index(int index);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_MODBUS_READ_WRITE_MODBUS_ERROR_CODE_H_
