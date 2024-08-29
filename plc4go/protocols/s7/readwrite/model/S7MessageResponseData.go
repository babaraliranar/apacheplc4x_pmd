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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7MessageResponseData is the corresponding interface of S7MessageResponseData
type S7MessageResponseData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Message
	// GetErrorClass returns ErrorClass (property field)
	GetErrorClass() uint8
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() uint8
}

// S7MessageResponseDataExactly can be used when we want exactly this type and not a type which fulfills S7MessageResponseData.
// This is useful for switch cases.
type S7MessageResponseDataExactly interface {
	S7MessageResponseData
	isS7MessageResponseData() bool
}

// _S7MessageResponseData is the data-structure of this message
type _S7MessageResponseData struct {
	*_S7Message
	ErrorClass uint8
	ErrorCode  uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageResponseData) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageResponseData) InitializeParent(parent S7Message, tpduReference uint16, parameter S7Parameter, payload S7Payload) {
	m.TpduReference = tpduReference
	m.Parameter = parameter
	m.Payload = payload
}

func (m *_S7MessageResponseData) GetParent() S7Message {
	return m._S7Message
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7MessageResponseData) GetErrorClass() uint8 {
	return m.ErrorClass
}

func (m *_S7MessageResponseData) GetErrorCode() uint8 {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7MessageResponseData factory function for _S7MessageResponseData
func NewS7MessageResponseData(errorClass uint8, errorCode uint8, tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7MessageResponseData {
	_result := &_S7MessageResponseData{
		ErrorClass: errorClass,
		ErrorCode:  errorCode,
		_S7Message: NewS7Message(tpduReference, parameter, payload),
	}
	_result._S7Message._S7MessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageResponseData(structType any) S7MessageResponseData {
	if casted, ok := structType.(S7MessageResponseData); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageResponseData); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageResponseData) GetTypeName() string {
	return "S7MessageResponseData"
}

func (m *_S7MessageResponseData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (errorClass)
	lengthInBits += 8

	// Simple field (errorCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7MessageResponseData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7MessageResponseDataParse(ctx context.Context, theBytes []byte) (S7MessageResponseData, error) {
	return S7MessageResponseDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7MessageResponseDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7MessageResponseData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7MessageResponseData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageResponseData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorClass)
	_errorClass, _errorClassErr := readBuffer.ReadUint8("errorClass", 8)
	if _errorClassErr != nil {
		return nil, errors.Wrap(_errorClassErr, "Error parsing 'errorClass' field of S7MessageResponseData")
	}
	errorClass := _errorClass

	// Simple Field (errorCode)
	_errorCode, _errorCodeErr := readBuffer.ReadUint8("errorCode", 8)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field of S7MessageResponseData")
	}
	errorCode := _errorCode

	if closeErr := readBuffer.CloseContext("S7MessageResponseData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageResponseData")
	}

	// Create a partially initialized instance
	_child := &_S7MessageResponseData{
		_S7Message: &_S7Message{},
		ErrorClass: errorClass,
		ErrorCode:  errorCode,
	}
	_child._S7Message._S7MessageChildRequirements = _child
	return _child, nil
}

func (m *_S7MessageResponseData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7MessageResponseData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageResponseData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageResponseData")
		}

		// Simple Field (errorClass)
		errorClass := uint8(m.GetErrorClass())
		_errorClassErr := writeBuffer.WriteUint8("errorClass", 8, uint8((errorClass)))
		if _errorClassErr != nil {
			return errors.Wrap(_errorClassErr, "Error serializing 'errorClass' field")
		}

		// Simple Field (errorCode)
		errorCode := uint8(m.GetErrorCode())
		_errorCodeErr := writeBuffer.WriteUint8("errorCode", 8, uint8((errorCode)))
		if _errorCodeErr != nil {
			return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
		}

		if popErr := writeBuffer.PopContext("S7MessageResponseData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageResponseData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7MessageResponseData) isS7MessageResponseData() bool {
	return true
}

func (m *_S7MessageResponseData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
