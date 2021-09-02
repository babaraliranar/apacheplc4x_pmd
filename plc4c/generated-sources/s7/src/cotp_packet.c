/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
#include "cotp_packet.h"

// Code generated by code-generation. DO NOT EDIT.

// Array of discriminator values that match the enum type constants.
// (The order is identical to the enum constants so we can use the
// enum constant to directly access a given types discriminator values)
const plc4c_s7_read_write_cotp_packet_discriminator plc4c_s7_read_write_cotp_packet_discriminators[] = {
  {/* plc4c_s7_read_write_cotp_packet_data */
   .tpduCode = 0xF0},
  {/* plc4c_s7_read_write_cotp_packet_connection_request */
   .tpduCode = 0xE0},
  {/* plc4c_s7_read_write_cotp_packet_connection_response */
   .tpduCode = 0xD0},
  {/* plc4c_s7_read_write_cotp_packet_disconnect_request */
   .tpduCode = 0x80},
  {/* plc4c_s7_read_write_cotp_packet_disconnect_response */
   .tpduCode = 0xC0},
  {/* plc4c_s7_read_write_cotp_packet_tpdu_error */
   .tpduCode = 0x70}

};

// Function returning the discriminator values for a given type constant.
plc4c_s7_read_write_cotp_packet_discriminator plc4c_s7_read_write_cotp_packet_get_discriminator(plc4c_s7_read_write_cotp_packet_type type) {
  return plc4c_s7_read_write_cotp_packet_discriminators[type];
}

// Create an empty NULL-struct
static const plc4c_s7_read_write_cotp_packet plc4c_s7_read_write_cotp_packet_null_const;

plc4c_s7_read_write_cotp_packet plc4c_s7_read_write_cotp_packet_null() {
  return plc4c_s7_read_write_cotp_packet_null_const;
}


// Parse function.
plc4c_return_code plc4c_s7_read_write_cotp_packet_parse(plc4c_spi_read_buffer* readBuffer, uint16_t cotpLen, plc4c_s7_read_write_cotp_packet** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  uint16_t curPos;
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_cotp_packet));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  uint8_t headerLength = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &headerLength);
  if(_res != OK) {
    return _res;
  }
        // Discriminator Field (tpduCode)

  // Discriminator Field (tpduCode) (Used as input to a switch field)
  uint8_t tpduCode = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &tpduCode);
  if(_res != OK) {
    return _res;
  }

  // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
  if(tpduCode == 0xF0) { /* COTPPacketData */
    (*_message)->_type = plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_data;
                    
    // Simple Field (eot)
    bool eot = false;
    _res = plc4c_spi_read_bit(readBuffer, (bool*) &eot);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_data_eot = eot;


                    
    // Simple Field (tpduRef)
    uint8_t tpduRef = 0;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 7, (uint8_t*) &tpduRef);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_data_tpdu_ref = tpduRef;

  } else 
  if(tpduCode == 0xE0) { /* COTPPacketConnectionRequest */
    (*_message)->_type = plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_connection_request;
                    
    // Simple Field (destinationReference)
    uint16_t destinationReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &destinationReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_connection_request_destination_reference = destinationReference;


                    
    // Simple Field (sourceReference)
    uint16_t sourceReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &sourceReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_connection_request_source_reference = sourceReference;


                    
    // Simple Field (protocolClass)
    plc4c_s7_read_write_cotp_protocol_class* protocolClass;
    _res = plc4c_s7_read_write_cotp_protocol_class_parse(readBuffer, (void*) &protocolClass);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_connection_request_protocol_class = *protocolClass;

  } else 
  if(tpduCode == 0xD0) { /* COTPPacketConnectionResponse */
    (*_message)->_type = plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_connection_response;
                    
    // Simple Field (destinationReference)
    uint16_t destinationReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &destinationReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_connection_response_destination_reference = destinationReference;


                    
    // Simple Field (sourceReference)
    uint16_t sourceReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &sourceReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_connection_response_source_reference = sourceReference;


                    
    // Simple Field (protocolClass)
    plc4c_s7_read_write_cotp_protocol_class* protocolClass;
    _res = plc4c_s7_read_write_cotp_protocol_class_parse(readBuffer, (void*) &protocolClass);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_connection_response_protocol_class = *protocolClass;

  } else 
  if(tpduCode == 0x80) { /* COTPPacketDisconnectRequest */
    (*_message)->_type = plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_disconnect_request;
                    
    // Simple Field (destinationReference)
    uint16_t destinationReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &destinationReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_disconnect_request_destination_reference = destinationReference;


                    
    // Simple Field (sourceReference)
    uint16_t sourceReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &sourceReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_disconnect_request_source_reference = sourceReference;


                    
    // Simple Field (protocolClass)
    plc4c_s7_read_write_cotp_protocol_class* protocolClass;
    _res = plc4c_s7_read_write_cotp_protocol_class_parse(readBuffer, (void*) &protocolClass);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_disconnect_request_protocol_class = *protocolClass;

  } else 
  if(tpduCode == 0xC0) { /* COTPPacketDisconnectResponse */
    (*_message)->_type = plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_disconnect_response;
                    
    // Simple Field (destinationReference)
    uint16_t destinationReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &destinationReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_disconnect_response_destination_reference = destinationReference;


                    
    // Simple Field (sourceReference)
    uint16_t sourceReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &sourceReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_disconnect_response_source_reference = sourceReference;

  } else 
  if(tpduCode == 0x70) { /* COTPPacketTpduError */
    (*_message)->_type = plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_tpdu_error;
                    
    // Simple Field (destinationReference)
    uint16_t destinationReference = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &destinationReference);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_tpdu_error_destination_reference = destinationReference;


                    
    // Simple Field (rejectCause)
    uint8_t rejectCause = 0;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &rejectCause);
    if(_res != OK) {
      return _res;
    }
    (*_message)->cotp_packet_tpdu_error_reject_cause = rejectCause;

  }

  // Array field (parameters)
  curPos = plc4c_spi_read_get_pos(readBuffer) - startPos;
  plc4c_list* parameters = NULL;
  plc4c_utils_list_create(&parameters);
  if(parameters == NULL) {
    return NO_MEMORY;
  }
  {
    // Length array
    uint8_t _parametersLength = (((headerLength) + (1))) - (curPos);
    uint8_t parametersEndPos = plc4c_spi_read_get_pos(readBuffer) + _parametersLength;
    while(plc4c_spi_read_get_pos(readBuffer) < parametersEndPos) {
      plc4c_s7_read_write_cotp_parameter* _value = NULL;
      _res = plc4c_s7_read_write_cotp_parameter_parse(readBuffer, (((headerLength) + (1))) - (curPos), (void*) &_value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(parameters, _value);
      curPos = plc4c_spi_read_get_pos(readBuffer) - startPos;
    }
  }
  (*_message)->parameters = parameters;

  // Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
  curPos = plc4c_spi_read_get_pos(readBuffer) - startPos;
  plc4c_s7_read_write_s7_message* payload = NULL;
  if((curPos) < (cotpLen)) {
    _res = plc4c_s7_read_write_s7_message_parse(readBuffer, &payload);
    if(_res != OK) {
      return _res;
    }
    (*_message)->payload = payload;
  } else {
    (*_message)->payload = NULL;
  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_cotp_packet_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_cotp_packet* _message) {
  plc4c_return_code _res = OK;

  // Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (plc4c_s7_read_write_cotp_packet_length_in_bytes(_message)) - ((((((((_message->payload) != (NULL))) ? plc4c_s7_read_write_s7_message_length_in_bytes(_message->payload) : 0))) + (1))));
  if(_res != OK) {
    return _res;
  }

  // Discriminator Field (tpduCode)
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_s7_read_write_cotp_packet_get_discriminator(_message->_type).tpduCode);

  // Switch Field (Depending of the current type, serialize the sub-type elements)
  switch(_message->_type) {
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_data: {

      // Simple Field (eot)
      _res = plc4c_spi_write_bit(writeBuffer, _message->cotp_packet_data_eot);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (tpduRef)
      _res = plc4c_spi_write_unsigned_byte(writeBuffer, 7, _message->cotp_packet_data_tpdu_ref);
      if(_res != OK) {
        return _res;
      }

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_connection_request: {

      // Simple Field (destinationReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_connection_request_destination_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (sourceReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_connection_request_source_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (protocolClass)
      _res = plc4c_s7_read_write_cotp_protocol_class_serialize(writeBuffer, &_message->cotp_packet_connection_request_protocol_class);
      if(_res != OK) {
        return _res;
      }

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_connection_response: {

      // Simple Field (destinationReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_connection_response_destination_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (sourceReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_connection_response_source_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (protocolClass)
      _res = plc4c_s7_read_write_cotp_protocol_class_serialize(writeBuffer, &_message->cotp_packet_connection_response_protocol_class);
      if(_res != OK) {
        return _res;
      }

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_disconnect_request: {

      // Simple Field (destinationReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_disconnect_request_destination_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (sourceReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_disconnect_request_source_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (protocolClass)
      _res = plc4c_s7_read_write_cotp_protocol_class_serialize(writeBuffer, &_message->cotp_packet_disconnect_request_protocol_class);
      if(_res != OK) {
        return _res;
      }

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_disconnect_response: {

      // Simple Field (destinationReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_disconnect_response_destination_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (sourceReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_disconnect_response_source_reference);
      if(_res != OK) {
        return _res;
      }

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_tpdu_error: {

      // Simple Field (destinationReference)
      _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->cotp_packet_tpdu_error_destination_reference);
      if(_res != OK) {
        return _res;
      }

      // Simple Field (rejectCause)
      _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, _message->cotp_packet_tpdu_error_reject_cause);
      if(_res != OK) {
        return _res;
      }

      break;
    }
  }

  // Array field (parameters)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->parameters);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_s7_read_write_cotp_parameter* _value = (plc4c_s7_read_write_cotp_parameter*) plc4c_utils_list_get_value(_message->parameters, curItem);
      _res = plc4c_s7_read_write_cotp_parameter_serialize(writeBuffer, (void*) _value);
      if(_res != OK) {
        return _res;
      }
    }
  }

  // Optional Field (payload)
  if(_message->payload != NULL) {
    _res = plc4c_s7_read_write_s7_message_serialize(writeBuffer, _message->payload);
    if(_res != OK) {
      return _res;
    }
  }

  return OK;
}

uint16_t plc4c_s7_read_write_cotp_packet_length_in_bytes(plc4c_s7_read_write_cotp_packet* _message) {
  return plc4c_s7_read_write_cotp_packet_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_cotp_packet_length_in_bits(plc4c_s7_read_write_cotp_packet* _message) {
  uint16_t lengthInBits = 0;

  // Implicit Field (headerLength)
  lengthInBits += 8;

  // Discriminator Field (tpduCode)
  lengthInBits += 8;

  // Depending of the current type, add the length of sub-type elements ...
  switch(_message->_type) {
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_data: {

      // Simple field (eot)
      lengthInBits += 1;


      // Simple field (tpduRef)
      lengthInBits += 7;

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_connection_request: {

      // Simple field (destinationReference)
      lengthInBits += 16;


      // Simple field (sourceReference)
      lengthInBits += 16;


      // Simple field (protocolClass)
      lengthInBits += plc4c_s7_read_write_cotp_protocol_class_length_in_bits(&_message->cotp_packet_connection_request_protocol_class);

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_connection_response: {

      // Simple field (destinationReference)
      lengthInBits += 16;


      // Simple field (sourceReference)
      lengthInBits += 16;


      // Simple field (protocolClass)
      lengthInBits += plc4c_s7_read_write_cotp_protocol_class_length_in_bits(&_message->cotp_packet_connection_response_protocol_class);

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_disconnect_request: {

      // Simple field (destinationReference)
      lengthInBits += 16;


      // Simple field (sourceReference)
      lengthInBits += 16;


      // Simple field (protocolClass)
      lengthInBits += plc4c_s7_read_write_cotp_protocol_class_length_in_bits(&_message->cotp_packet_disconnect_request_protocol_class);

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_disconnect_response: {

      // Simple field (destinationReference)
      lengthInBits += 16;


      // Simple field (sourceReference)
      lengthInBits += 16;

      break;
    }
    case plc4c_s7_read_write_cotp_packet_type_plc4c_s7_read_write_cotp_packet_tpdu_error: {

      // Simple field (destinationReference)
      lengthInBits += 16;


      // Simple field (rejectCause)
      lengthInBits += 8;

      break;
    }
  }

  // Array field
  if(_message->parameters != NULL) {
    plc4c_list_element* curElement = _message->parameters->tail;
    while (curElement != NULL) {
      lengthInBits += plc4c_s7_read_write_cotp_parameter_length_in_bits((plc4c_s7_read_write_cotp_parameter*) curElement->value);
      curElement = curElement->next;
    }
  }

  // Optional Field (payload)
  if(_message->payload != NULL) {
    lengthInBits += plc4c_s7_read_write_s7_message_length_in_bits(_message->payload);
  }

  return lengthInBits;
}

