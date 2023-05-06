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
#include <plc4c/spi/context.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/driver_s7_static.h>

#include "associated_value_type.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_s7_read_write_associated_value_type_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_associated_value_type** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_associated_value_type));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (returnCode)
  plc4c_s7_read_write_data_transport_error_code returnCode;
  _res = plc4c_s7_read_write_data_transport_error_code_parse(ctx, readBuffer, (void*) &returnCode);
  if(_res != OK) {
    return _res;
  }
  (*_message)->return_code = returnCode;

  // Simple Field (transportSize)
  plc4c_s7_read_write_data_transport_size transportSize;
  _res = plc4c_s7_read_write_data_transport_size_parse(ctx, readBuffer, (void*) &transportSize);
  if(_res != OK) {
    return _res;
  }
  (*_message)->transport_size = transportSize;

  // Manual Field (valueLength)
  uint16_t valueLength = (uint16_t) (plc4c_s7_read_write_right_shift3(readBuffer, transportSize));
  (*_message)->value_length = valueLength;

  // Array field (data)
  plc4c_list* data = NULL;
  plc4c_utils_list_create(&data);
  if(data == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) plc4c_s7_read_write_event_item_length(readBuffer, valueLength);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      uint8_t* _value = malloc(sizeof(uint8_t));
      _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) _value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(data, _value);
    }
  }
  (*_message)->data = data;

  return OK;
}

plc4c_return_code plc4c_s7_read_write_associated_value_type_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_associated_value_type* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (returnCode)
  _res = plc4c_s7_read_write_data_transport_error_code_serialize(ctx, writeBuffer, &_message->return_code);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (transportSize)
  _res = plc4c_s7_read_write_data_transport_size_serialize(ctx, writeBuffer, &_message->transport_size);
  if(_res != OK) {
    return _res;
  }

  // Manual Field (valueLength)  {
  _res = plc4c_s7_read_write_left_shift3(writeBuffer, _message->value_length);
  if(_res != OK) {
    return _res;
  }

  // Array field (data)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->data);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      uint8_t* _value = (uint8_t*) plc4c_utils_list_get_value(_message->data, curItem);
      plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_value);
    }
  }

  return OK;
}

uint16_t plc4c_s7_read_write_associated_value_type_length_in_bytes(plc4x_spi_context ctx, plc4c_s7_read_write_associated_value_type* _message) {
  return plc4c_s7_read_write_associated_value_type_length_in_bits(ctx, _message) / 8;
}

uint16_t plc4c_s7_read_write_associated_value_type_length_in_bits(plc4x_spi_context ctx, plc4c_s7_read_write_associated_value_type* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (returnCode)
  lengthInBits += plc4c_s7_read_write_data_transport_error_code_length_in_bits(ctx, &_message->return_code);

  // Simple field (transportSize)
  lengthInBits += plc4c_s7_read_write_data_transport_size_length_in_bits(ctx, &_message->transport_size);

  // Manual Field (valueLength)
  lengthInBits += 2;

  // Array field
  lengthInBits += 8 * plc4c_utils_list_size(_message->data);

  return lengthInBits;
}

