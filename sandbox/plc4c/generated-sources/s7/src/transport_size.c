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

#include "transport_size.h"
#include <string.h>

#include "data_transport_size.h"
#include "transport_size.h"

// Create an empty NULL-struct
static const plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_null_const;

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_null() {
  return plc4c_s7_read_write_transport_size_null_const;
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_value_of(char* value_string) {
    if(strcmp(value_string, "BOOL") == 0) {
        return plc4c_s7_read_write_transport_size_BOOL;
    }
    if(strcmp(value_string, "BYTE") == 0) {
        return plc4c_s7_read_write_transport_size_BYTE;
    }
    if(strcmp(value_string, "WORD") == 0) {
        return plc4c_s7_read_write_transport_size_WORD;
    }
    if(strcmp(value_string, "DWORD") == 0) {
        return plc4c_s7_read_write_transport_size_DWORD;
    }
    if(strcmp(value_string, "LWORD") == 0) {
        return plc4c_s7_read_write_transport_size_LWORD;
    }
    if(strcmp(value_string, "INT") == 0) {
        return plc4c_s7_read_write_transport_size_INT;
    }
    if(strcmp(value_string, "UINT") == 0) {
        return plc4c_s7_read_write_transport_size_UINT;
    }
    if(strcmp(value_string, "SINT") == 0) {
        return plc4c_s7_read_write_transport_size_SINT;
    }
    if(strcmp(value_string, "USINT") == 0) {
        return plc4c_s7_read_write_transport_size_USINT;
    }
    if(strcmp(value_string, "DINT") == 0) {
        return plc4c_s7_read_write_transport_size_DINT;
    }
    if(strcmp(value_string, "UDINT") == 0) {
        return plc4c_s7_read_write_transport_size_UDINT;
    }
    if(strcmp(value_string, "LINT") == 0) {
        return plc4c_s7_read_write_transport_size_LINT;
    }
    if(strcmp(value_string, "ULINT") == 0) {
        return plc4c_s7_read_write_transport_size_ULINT;
    }
    if(strcmp(value_string, "REAL") == 0) {
        return plc4c_s7_read_write_transport_size_REAL;
    }
    if(strcmp(value_string, "LREAL") == 0) {
        return plc4c_s7_read_write_transport_size_LREAL;
    }
    if(strcmp(value_string, "CHAR") == 0) {
        return plc4c_s7_read_write_transport_size_CHAR;
    }
    if(strcmp(value_string, "WCHAR") == 0) {
        return plc4c_s7_read_write_transport_size_WCHAR;
    }
    if(strcmp(value_string, "STRING") == 0) {
        return plc4c_s7_read_write_transport_size_STRING;
    }
    if(strcmp(value_string, "WSTRING") == 0) {
        return plc4c_s7_read_write_transport_size_WSTRING;
    }
    if(strcmp(value_string, "TIME") == 0) {
        return plc4c_s7_read_write_transport_size_TIME;
    }
    if(strcmp(value_string, "LTIME") == 0) {
        return plc4c_s7_read_write_transport_size_LTIME;
    }
    if(strcmp(value_string, "DATE") == 0) {
        return plc4c_s7_read_write_transport_size_DATE;
    }
    if(strcmp(value_string, "TIME_OF_DAY") == 0) {
        return plc4c_s7_read_write_transport_size_TIME_OF_DAY;
    }
    if(strcmp(value_string, "TOD") == 0) {
        return plc4c_s7_read_write_transport_size_TOD;
    }
    if(strcmp(value_string, "DATE_AND_TIME") == 0) {
        return plc4c_s7_read_write_transport_size_DATE_AND_TIME;
    }
    if(strcmp(value_string, "DT") == 0) {
        return plc4c_s7_read_write_transport_size_DT;
    }
    return -1;
}

int plc4c_s7_read_write_transport_size_num_values() {
  return 26;
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_s7_read_write_transport_size_BOOL;
      }
      case 1: {
        return plc4c_s7_read_write_transport_size_BYTE;
      }
      case 2: {
        return plc4c_s7_read_write_transport_size_WORD;
      }
      case 3: {
        return plc4c_s7_read_write_transport_size_DWORD;
      }
      case 4: {
        return plc4c_s7_read_write_transport_size_LWORD;
      }
      case 5: {
        return plc4c_s7_read_write_transport_size_INT;
      }
      case 6: {
        return plc4c_s7_read_write_transport_size_UINT;
      }
      case 7: {
        return plc4c_s7_read_write_transport_size_SINT;
      }
      case 8: {
        return plc4c_s7_read_write_transport_size_USINT;
      }
      case 9: {
        return plc4c_s7_read_write_transport_size_DINT;
      }
      case 10: {
        return plc4c_s7_read_write_transport_size_UDINT;
      }
      case 11: {
        return plc4c_s7_read_write_transport_size_LINT;
      }
      case 12: {
        return plc4c_s7_read_write_transport_size_ULINT;
      }
      case 13: {
        return plc4c_s7_read_write_transport_size_REAL;
      }
      case 14: {
        return plc4c_s7_read_write_transport_size_LREAL;
      }
      case 15: {
        return plc4c_s7_read_write_transport_size_CHAR;
      }
      case 16: {
        return plc4c_s7_read_write_transport_size_WCHAR;
      }
      case 17: {
        return plc4c_s7_read_write_transport_size_STRING;
      }
      case 18: {
        return plc4c_s7_read_write_transport_size_WSTRING;
      }
      case 19: {
        return plc4c_s7_read_write_transport_size_TIME;
      }
      case 20: {
        return plc4c_s7_read_write_transport_size_LTIME;
      }
      case 21: {
        return plc4c_s7_read_write_transport_size_DATE;
      }
      case 22: {
        return plc4c_s7_read_write_transport_size_TIME_OF_DAY;
      }
      case 23: {
        return plc4c_s7_read_write_transport_size_TOD;
      }
      case 24: {
        return plc4c_s7_read_write_transport_size_DATE_AND_TIME;
      }
      case 25: {
        return plc4c_s7_read_write_transport_size_DT;
      }
      default: {
        return -1;
      }
    }
}

bool plc4c_s7_read_write_transport_size_get_supported__s7_300(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return true;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_300(bool value) {
    switch(value) {
        case false: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        case true: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

bool plc4c_s7_read_write_transport_size_get_supported__logo(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return false;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__logo(bool value) {
    switch(value) {
        case false: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        case true: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

uint8_t plc4c_s7_read_write_transport_size_get_code(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return 6;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return 0;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return 5;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return 5;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return 7;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return 7;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return 0;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return 0;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return 8;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return 48;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return 3;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return 19;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return 3;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return 0;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return 11;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return 0;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return 9;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return 6;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return 6;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return 15;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return 15;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_code(uint8_t value) {
    switch(value) {
        case 0: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        case 1: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        case 2: {
            return plc4c_s7_read_write_transport_size_BYTE;
        }
        case 3: {
            return plc4c_s7_read_write_transport_size_CHAR;
        }
        case 4: {
            return plc4c_s7_read_write_transport_size_WORD;
        }
        case 5: {
            return plc4c_s7_read_write_transport_size_INT;
        }
        case 6: {
            return plc4c_s7_read_write_transport_size_DWORD;
        }
        case 7: {
            return plc4c_s7_read_write_transport_size_DINT;
        }
        case 8: {
            return plc4c_s7_read_write_transport_size_REAL;
        }
        case 9: {
            return plc4c_s7_read_write_transport_size_DATE;
        }
        case 11: {
            return plc4c_s7_read_write_transport_size_TIME;
        }
        case 15: {
            return plc4c_s7_read_write_transport_size_DATE_AND_TIME;
        }
        case 19: {
            return plc4c_s7_read_write_transport_size_WCHAR;
        }
        case 48: {
            return plc4c_s7_read_write_transport_size_LREAL;
        }
        default: {
            return -1;
        }
    }
}

uint8_t plc4c_s7_read_write_transport_size_get_size_in_bytes(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return 8;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return 8;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return 16;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return 8;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return 1;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return 8;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return 2;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return 4;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return 12;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return 12;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_size_in_bytes(uint8_t value) {
    switch(value) {
        case 1: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        case 12: {
            return plc4c_s7_read_write_transport_size_DATE_AND_TIME;
        }
        case 16: {
            return plc4c_s7_read_write_transport_size_ULINT;
        }
        case 2: {
            return plc4c_s7_read_write_transport_size_WORD;
        }
        case 4: {
            return plc4c_s7_read_write_transport_size_DWORD;
        }
        case 8: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        default: {
            return -1;
        }
    }
}

bool plc4c_s7_read_write_transport_size_get_supported__s7_400(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return true;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_400(bool value) {
    switch(value) {
        case false: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        case true: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

bool plc4c_s7_read_write_transport_size_get_supported__s7_1200(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return false;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return false;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_1200(bool value) {
    switch(value) {
        case false: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        case true: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

uint8_t plc4c_s7_read_write_transport_size_get_short_name(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return 'B';
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return 'W';
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return 'D';
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return 'W';
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return 'W';
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return 'B';
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return 'B';
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return 'D';
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return 'D';
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return 'D';
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return 'B';
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return 'X';
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return 'X';
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_short_name(uint8_t value) {
    switch(value) {
        case 'B': {
            return plc4c_s7_read_write_transport_size_BYTE;
        }
        case 'D': {
            return plc4c_s7_read_write_transport_size_DWORD;
        }
        case 'W': {
            return plc4c_s7_read_write_transport_size_WORD;
        }
        case 'X': {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

bool plc4c_s7_read_write_transport_size_get_supported__s7_1500(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return true;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return true;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_1500(bool value) {
    switch(value) {
        case true: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

plc4c_s7_read_write_data_transport_size plc4c_s7_read_write_transport_size_get_data_transport_size(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return plc4c_s7_read_write_data_transport_size_BIT;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return plc4c_s7_read_write_data_transport_size_INTEGER;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return plc4c_s7_read_write_data_transport_size_INTEGER;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return plc4c_s7_read_write_data_transport_size_INTEGER;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return plc4c_s7_read_write_data_transport_size_INTEGER;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return -1;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_data_transport_size(plc4c_s7_read_write_data_transport_size value) {
    switch(value) {
        case plc4c_s7_read_write_data_transport_size_BIT: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        case plc4c_s7_read_write_data_transport_size_BYTE_WORD_DWORD: {
            return plc4c_s7_read_write_transport_size_BYTE;
        }
        case plc4c_s7_read_write_data_transport_size_INTEGER: {
            return plc4c_s7_read_write_transport_size_INT;
        }
        case -1: {
            return plc4c_s7_read_write_transport_size_LWORD;
        }
        default: {
            return -1;
        }
    }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_base_type(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return plc4c_s7_read_write_transport_size_WORD;
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return plc4c_s7_read_write_transport_size_INT;
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return plc4c_s7_read_write_transport_size_REAL;
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return plc4c_s7_read_write_transport_size_TIME;
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return -1;
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return -1;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_base_type(plc4c_s7_read_write_transport_size value) {
    switch(value) {
        case plc4c_s7_read_write_transport_size_INT: {
            return plc4c_s7_read_write_transport_size_UINT;
        }
        case plc4c_s7_read_write_transport_size_REAL: {
            return plc4c_s7_read_write_transport_size_LREAL;
        }
        case plc4c_s7_read_write_transport_size_TIME: {
            return plc4c_s7_read_write_transport_size_LTIME;
        }
        case plc4c_s7_read_write_transport_size_WORD: {
            return plc4c_s7_read_write_transport_size_DWORD;
        }
        case -1: {
            return plc4c_s7_read_write_transport_size_BOOL;
        }
        default: {
            return -1;
        }
    }
}

char* plc4c_s7_read_write_transport_size_get_data_protocol_id(plc4c_s7_read_write_transport_size value) {
  switch(value) {
    case plc4c_s7_read_write_transport_size_BOOL: { /* '0x01' */
      return "IEC61131_BOOL";
    }
    case plc4c_s7_read_write_transport_size_BYTE: { /* '0x02' */
      return "IEC61131_BYTE";
    }
    case plc4c_s7_read_write_transport_size_WORD: { /* '0x03' */
      return "IEC61131_WORD";
    }
    case plc4c_s7_read_write_transport_size_DWORD: { /* '0x04' */
      return "IEC61131_DWORD";
    }
    case plc4c_s7_read_write_transport_size_LWORD: { /* '0x05' */
      return "IEC61131_LWORD";
    }
    case plc4c_s7_read_write_transport_size_INT: { /* '0x06' */
      return "IEC61131_INT";
    }
    case plc4c_s7_read_write_transport_size_UINT: { /* '0x07' */
      return "IEC61131_UINT";
    }
    case plc4c_s7_read_write_transport_size_SINT: { /* '0x08' */
      return "IEC61131_SINT";
    }
    case plc4c_s7_read_write_transport_size_USINT: { /* '0x09' */
      return "IEC61131_USINT";
    }
    case plc4c_s7_read_write_transport_size_DINT: { /* '0x0A' */
      return "IEC61131_DINT";
    }
    case plc4c_s7_read_write_transport_size_UDINT: { /* '0x0B' */
      return "IEC61131_UDINT";
    }
    case plc4c_s7_read_write_transport_size_LINT: { /* '0x0C' */
      return "IEC61131_LINT";
    }
    case plc4c_s7_read_write_transport_size_ULINT: { /* '0x0D' */
      return "IEC61131_ULINT";
    }
    case plc4c_s7_read_write_transport_size_REAL: { /* '0x0E' */
      return "IEC61131_REAL";
    }
    case plc4c_s7_read_write_transport_size_LREAL: { /* '0x0F' */
      return "IEC61131_LREAL";
    }
    case plc4c_s7_read_write_transport_size_CHAR: { /* '0x10' */
      return "IEC61131_CHAR";
    }
    case plc4c_s7_read_write_transport_size_WCHAR: { /* '0x11' */
      return "IEC61131_WCHAR";
    }
    case plc4c_s7_read_write_transport_size_STRING: { /* '0x12' */
      return "IEC61131_STRING";
    }
    case plc4c_s7_read_write_transport_size_WSTRING: { /* '0x13' */
      return "IEC61131_WSTRING";
    }
    case plc4c_s7_read_write_transport_size_TIME: { /* '0x14' */
      return "IEC61131_TIME";
    }
    case plc4c_s7_read_write_transport_size_LTIME: { /* '0x16' */
      return "IEC61131_LTIME";
    }
    case plc4c_s7_read_write_transport_size_DATE: { /* '0x17' */
      return "IEC61131_DATE";
    }
    case plc4c_s7_read_write_transport_size_TIME_OF_DAY: { /* '0x18' */
      return "IEC61131_TIME_OF_DAY";
    }
    case plc4c_s7_read_write_transport_size_TOD: { /* '0x19' */
      return "IEC61131_TIME_OF_DAY";
    }
    case plc4c_s7_read_write_transport_size_DATE_AND_TIME: { /* '0x1A' */
      return "IEC61131_DATE_AND_TIME";
    }
    case plc4c_s7_read_write_transport_size_DT: { /* '0x1B' */
      return "IEC61131_DATE_AND_TIME";
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_data_protocol_id(char* value) {
    if (strcmp(value, "IEC61131_BOOL") == 0) {
        return plc4c_s7_read_write_transport_size_BOOL;
    }
    if (strcmp(value, "IEC61131_BYTE") == 0) {
        return plc4c_s7_read_write_transport_size_BYTE;
    }
    if (strcmp(value, "IEC61131_CHAR") == 0) {
        return plc4c_s7_read_write_transport_size_CHAR;
    }
    if (strcmp(value, "IEC61131_DATE") == 0) {
        return plc4c_s7_read_write_transport_size_DATE;
    }
    if (strcmp(value, "IEC61131_DATE_AND_TIME") == 0) {
        return plc4c_s7_read_write_transport_size_DATE_AND_TIME;
    }
    if (strcmp(value, "IEC61131_DINT") == 0) {
        return plc4c_s7_read_write_transport_size_DINT;
    }
    if (strcmp(value, "IEC61131_DWORD") == 0) {
        return plc4c_s7_read_write_transport_size_DWORD;
    }
    if (strcmp(value, "IEC61131_INT") == 0) {
        return plc4c_s7_read_write_transport_size_INT;
    }
    if (strcmp(value, "IEC61131_LINT") == 0) {
        return plc4c_s7_read_write_transport_size_LINT;
    }
    if (strcmp(value, "IEC61131_LREAL") == 0) {
        return plc4c_s7_read_write_transport_size_LREAL;
    }
    if (strcmp(value, "IEC61131_LTIME") == 0) {
        return plc4c_s7_read_write_transport_size_LTIME;
    }
    if (strcmp(value, "IEC61131_LWORD") == 0) {
        return plc4c_s7_read_write_transport_size_LWORD;
    }
    if (strcmp(value, "IEC61131_REAL") == 0) {
        return plc4c_s7_read_write_transport_size_REAL;
    }
    if (strcmp(value, "IEC61131_SINT") == 0) {
        return plc4c_s7_read_write_transport_size_SINT;
    }
    if (strcmp(value, "IEC61131_STRING") == 0) {
        return plc4c_s7_read_write_transport_size_STRING;
    }
    if (strcmp(value, "IEC61131_TIME") == 0) {
        return plc4c_s7_read_write_transport_size_TIME;
    }
    if (strcmp(value, "IEC61131_TIME_OF_DAY") == 0) {
        return plc4c_s7_read_write_transport_size_TIME_OF_DAY;
    }
    if (strcmp(value, "IEC61131_UDINT") == 0) {
        return plc4c_s7_read_write_transport_size_UDINT;
    }
    if (strcmp(value, "IEC61131_UINT") == 0) {
        return plc4c_s7_read_write_transport_size_UINT;
    }
    if (strcmp(value, "IEC61131_ULINT") == 0) {
        return plc4c_s7_read_write_transport_size_ULINT;
    }
    if (strcmp(value, "IEC61131_USINT") == 0) {
        return plc4c_s7_read_write_transport_size_USINT;
    }
    if (strcmp(value, "IEC61131_WCHAR") == 0) {
        return plc4c_s7_read_write_transport_size_WCHAR;
    }
    if (strcmp(value, "IEC61131_WORD") == 0) {
        return plc4c_s7_read_write_transport_size_WORD;
    }
    if (strcmp(value, "IEC61131_WSTRING") == 0) {
        return plc4c_s7_read_write_transport_size_WSTRING;
    }
}
