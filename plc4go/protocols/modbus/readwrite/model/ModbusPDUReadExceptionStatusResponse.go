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

package model

import (
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadExceptionStatusResponse is the corresponding interface of ModbusPDUReadExceptionStatusResponse
type ModbusPDUReadExceptionStatusResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() uint8
}

// ModbusPDUReadExceptionStatusResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadExceptionStatusResponse.
// This is useful for switch cases.
type ModbusPDUReadExceptionStatusResponseExactly interface {
	ModbusPDUReadExceptionStatusResponse
	isModbusPDUReadExceptionStatusResponse() bool
}

// _ModbusPDUReadExceptionStatusResponse is the data-structure of this message
type _ModbusPDUReadExceptionStatusResponse struct {
	*_ModbusPDU
	Value uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadExceptionStatusResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadExceptionStatusResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadExceptionStatusResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadExceptionStatusResponse) GetValue() uint8 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadExceptionStatusResponse factory function for _ModbusPDUReadExceptionStatusResponse
func NewModbusPDUReadExceptionStatusResponse(value uint8) *_ModbusPDUReadExceptionStatusResponse {
	_result := &_ModbusPDUReadExceptionStatusResponse{
		Value:      value,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadExceptionStatusResponse(structType interface{}) ModbusPDUReadExceptionStatusResponse {
	if casted, ok := structType.(ModbusPDUReadExceptionStatusResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadExceptionStatusResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusResponse"
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadExceptionStatusResponseParse(theBytes []byte, response bool) (ModbusPDUReadExceptionStatusResponse, error) {
	return ModbusPDUReadExceptionStatusResponseParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response) // TODO: get endianness from mspec
}

func ModbusPDUReadExceptionStatusResponseParseWithBuffer(readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadExceptionStatusResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadExceptionStatusResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ModbusPDUReadExceptionStatusResponse")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadExceptionStatusResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadExceptionStatusResponse{
		_ModbusPDU: &_ModbusPDU{},
		Value:      value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadExceptionStatusResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadExceptionStatusResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadExceptionStatusResponse")
		}

		// Simple Field (value)
		value := uint8(m.GetValue())
		_valueErr := writeBuffer.WriteUint8("value", 8, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadExceptionStatusResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUReadExceptionStatusResponse) isModbusPDUReadExceptionStatusResponse() bool {
	return true
}

func (m *_ModbusPDUReadExceptionStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
