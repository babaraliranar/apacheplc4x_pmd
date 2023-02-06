/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#include <string.h>
#include <plc4c/driver_plc4x_static.h>

#include "plc4x_value_type.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_plc4x_read_write_plc4x_value_type plc4c_plc4x_read_write_plc4x_value_type_null_const;

plc4c_plc4x_read_write_plc4x_value_type plc4c_plc4x_read_write_plc4x_value_type_null() {
  return plc4c_plc4x_read_write_plc4x_value_type_null_const;
}

// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_value_type_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_value_type* _message) {
    plc4c_return_code _res = OK;

    uint8_t value;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
    *_message = plc4c_plc4x_read_write_plc4x_value_type_for_value(value);

    return _res;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_value_type_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_value_type* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_plc4x_read_write_plc4x_value_type plc4c_plc4x_read_write_plc4x_value_type_for_value(uint8_t value) {
    for(int i = 0; i < plc4c_plc4x_read_write_plc4x_value_type_num_values(); i++) {
        if(plc4c_plc4x_read_write_plc4x_value_type_value_for_index(i) == value) {
            return plc4c_plc4x_read_write_plc4x_value_type_value_for_index(i);
        }
    }
    return -1;
}

plc4c_plc4x_read_write_plc4x_value_type plc4c_plc4x_read_write_plc4x_value_type_value_of(char* value_string) {
    if(strcmp(value_string, "NULL") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_NULL;
    }
    if(strcmp(value_string, "BOOL") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_BOOL;
    }
    if(strcmp(value_string, "BYTE") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_BYTE;
    }
    if(strcmp(value_string, "WORD") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_WORD;
    }
    if(strcmp(value_string, "DWORD") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_DWORD;
    }
    if(strcmp(value_string, "LWORD") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LWORD;
    }
    if(strcmp(value_string, "USINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_USINT;
    }
    if(strcmp(value_string, "UINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_UINT;
    }
    if(strcmp(value_string, "UDINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_UDINT;
    }
    if(strcmp(value_string, "ULINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_ULINT;
    }
    if(strcmp(value_string, "SINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_SINT;
    }
    if(strcmp(value_string, "INT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_INT;
    }
    if(strcmp(value_string, "DINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_DINT;
    }
    if(strcmp(value_string, "LINT") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LINT;
    }
    if(strcmp(value_string, "REAL") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_REAL;
    }
    if(strcmp(value_string, "LREAL") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LREAL;
    }
    if(strcmp(value_string, "CHAR") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_CHAR;
    }
    if(strcmp(value_string, "WCHAR") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_WCHAR;
    }
    if(strcmp(value_string, "STRING") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_STRING;
    }
    if(strcmp(value_string, "WSTRING") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_WSTRING;
    }
    if(strcmp(value_string, "TIME") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_TIME;
    }
    if(strcmp(value_string, "LTIME") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LTIME;
    }
    if(strcmp(value_string, "DATE") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_DATE;
    }
    if(strcmp(value_string, "LDATE") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LDATE;
    }
    if(strcmp(value_string, "TIME_OF_DAY") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_TIME_OF_DAY;
    }
    if(strcmp(value_string, "LTIME_OF_DAY") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LTIME_OF_DAY;
    }
    if(strcmp(value_string, "DATE_AND_TIME") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_TIME;
    }
    if(strcmp(value_string, "LDATE_AND_TIME") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_LDATE_AND_TIME;
    }
    if(strcmp(value_string, "Struct") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_Struct;
    }
    if(strcmp(value_string, "List") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_List;
    }
    if(strcmp(value_string, "RAW_BYTE_ARRAY") == 0) {
        return plc4c_plc4x_read_write_plc4x_value_type_RAW_BYTE_ARRAY;
    }
    return -1;
}

int plc4c_plc4x_read_write_plc4x_value_type_num_values() {
  return 31;
}

plc4c_plc4x_read_write_plc4x_value_type plc4c_plc4x_read_write_plc4x_value_type_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_plc4x_read_write_plc4x_value_type_NULL;
      }
      case 1: {
        return plc4c_plc4x_read_write_plc4x_value_type_BOOL;
      }
      case 2: {
        return plc4c_plc4x_read_write_plc4x_value_type_BYTE;
      }
      case 3: {
        return plc4c_plc4x_read_write_plc4x_value_type_WORD;
      }
      case 4: {
        return plc4c_plc4x_read_write_plc4x_value_type_DWORD;
      }
      case 5: {
        return plc4c_plc4x_read_write_plc4x_value_type_LWORD;
      }
      case 6: {
        return plc4c_plc4x_read_write_plc4x_value_type_USINT;
      }
      case 7: {
        return plc4c_plc4x_read_write_plc4x_value_type_UINT;
      }
      case 8: {
        return plc4c_plc4x_read_write_plc4x_value_type_UDINT;
      }
      case 9: {
        return plc4c_plc4x_read_write_plc4x_value_type_ULINT;
      }
      case 10: {
        return plc4c_plc4x_read_write_plc4x_value_type_SINT;
      }
      case 11: {
        return plc4c_plc4x_read_write_plc4x_value_type_INT;
      }
      case 12: {
        return plc4c_plc4x_read_write_plc4x_value_type_DINT;
      }
      case 13: {
        return plc4c_plc4x_read_write_plc4x_value_type_LINT;
      }
      case 14: {
        return plc4c_plc4x_read_write_plc4x_value_type_REAL;
      }
      case 15: {
        return plc4c_plc4x_read_write_plc4x_value_type_LREAL;
      }
      case 16: {
        return plc4c_plc4x_read_write_plc4x_value_type_CHAR;
      }
      case 17: {
        return plc4c_plc4x_read_write_plc4x_value_type_WCHAR;
      }
      case 18: {
        return plc4c_plc4x_read_write_plc4x_value_type_STRING;
      }
      case 19: {
        return plc4c_plc4x_read_write_plc4x_value_type_WSTRING;
      }
      case 20: {
        return plc4c_plc4x_read_write_plc4x_value_type_TIME;
      }
      case 21: {
        return plc4c_plc4x_read_write_plc4x_value_type_LTIME;
      }
      case 22: {
        return plc4c_plc4x_read_write_plc4x_value_type_DATE;
      }
      case 23: {
        return plc4c_plc4x_read_write_plc4x_value_type_LDATE;
      }
      case 24: {
        return plc4c_plc4x_read_write_plc4x_value_type_TIME_OF_DAY;
      }
      case 25: {
        return plc4c_plc4x_read_write_plc4x_value_type_LTIME_OF_DAY;
      }
      case 26: {
        return plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_TIME;
      }
      case 27: {
        return plc4c_plc4x_read_write_plc4x_value_type_LDATE_AND_TIME;
      }
      case 28: {
        return plc4c_plc4x_read_write_plc4x_value_type_Struct;
      }
      case 29: {
        return plc4c_plc4x_read_write_plc4x_value_type_List;
      }
      case 30: {
        return plc4c_plc4x_read_write_plc4x_value_type_RAW_BYTE_ARRAY;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_plc4x_read_write_plc4x_value_type_length_in_bytes(plc4x_spi_context ctx, plc4c_plc4x_read_write_plc4x_value_type* _message) {
    return plc4c_plc4x_read_write_plc4x_value_type_length_in_bits(ctx, _message) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_value_type_length_in_bits(plc4x_spi_context ctx, plc4c_plc4x_read_write_plc4x_value_type* _message) {
    return 8;
}
