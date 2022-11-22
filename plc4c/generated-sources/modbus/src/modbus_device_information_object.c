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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/driver_modbus_static.h>

#include "modbus_device_information_object.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_device_information_object_parse(plc4c_spi_read_buffer* readBuffer, plc4c_modbus_read_write_modbus_device_information_object** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_modbus_read_write_modbus_device_information_object));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (objectId)
  uint8_t objectId = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &objectId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->object_id = objectId;

  // Implicit Field (objectLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t objectLength = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &objectLength);
  if(_res != OK) {
    return _res;
  }

  // Array field (data)
  plc4c_list* data = NULL;
  plc4c_utils_list_create(&data);
  if(data == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) objectLength;
    for(int curItem = 0; curItem < itemCount; curItem++) {
      
      char* _value = malloc(sizeof(char));
      _res = plc4c_spi_read_char(readBuffer, (char*) _value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(data, _value);
    }
  }
  (*_message)->data = data;

  return OK;
}

plc4c_return_code plc4c_modbus_read_write_modbus_device_information_object_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_modbus_read_write_modbus_device_information_object* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (objectId)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, _message->object_id);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (objectLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_spi_evaluation_helper_count(_message->data));
  if(_res != OK) {
    return _res;
  }

  // Array field (data)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->data);
    for(int curItem = 0; curItem < itemCount; curItem++) {

      char* _value = (char*) plc4c_utils_list_get_value(_message->data, curItem);
      plc4c_spi_write_char(writeBuffer, *_value);
    }
  }

  return OK;
}

uint16_t plc4c_modbus_read_write_modbus_device_information_object_length_in_bytes(plc4c_modbus_read_write_modbus_device_information_object* _message) {
  return plc4c_modbus_read_write_modbus_device_information_object_length_in_bits(_message) / 8;
}

uint16_t plc4c_modbus_read_write_modbus_device_information_object_length_in_bits(plc4c_modbus_read_write_modbus_device_information_object* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (objectId)
  lengthInBits += 8;

  // Implicit Field (objectLength)
  lengthInBits += 8;

  // Array field
  lengthInBits += 8 * plc4c_utils_list_size(_message->data);

  return lengthInBits;
}

