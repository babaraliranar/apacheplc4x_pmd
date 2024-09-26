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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7ParameterWriteVarResponse is the corresponding interface of S7ParameterWriteVarResponse
type S7ParameterWriteVarResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Parameter
	// GetNumItems returns NumItems (property field)
	GetNumItems() uint8
	// IsS7ParameterWriteVarResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterWriteVarResponse()
}

// _S7ParameterWriteVarResponse is the data-structure of this message
type _S7ParameterWriteVarResponse struct {
	S7ParameterContract
	NumItems uint8
}

var _ S7ParameterWriteVarResponse = (*_S7ParameterWriteVarResponse)(nil)
var _ S7ParameterRequirements = (*_S7ParameterWriteVarResponse)(nil)

// NewS7ParameterWriteVarResponse factory function for _S7ParameterWriteVarResponse
func NewS7ParameterWriteVarResponse(numItems uint8) *_S7ParameterWriteVarResponse {
	_result := &_S7ParameterWriteVarResponse{
		S7ParameterContract: NewS7Parameter(),
		NumItems:            numItems,
	}
	_result.S7ParameterContract.(*_S7Parameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterWriteVarResponse) GetParameterType() uint8 {
	return 0x05
}

func (m *_S7ParameterWriteVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterWriteVarResponse) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterWriteVarResponse) GetNumItems() uint8 {
	return m.NumItems
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterWriteVarResponse(structType any) S7ParameterWriteVarResponse {
	if casted, ok := structType.(S7ParameterWriteVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterWriteVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterWriteVarResponse) GetTypeName() string {
	return "S7ParameterWriteVarResponse"
}

func (m *_S7ParameterWriteVarResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Simple field (numItems)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterWriteVarResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterWriteVarResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterWriteVarResponse S7ParameterWriteVarResponse, err error) {
	m.S7ParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterWriteVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numItems, err := ReadSimpleField(ctx, "numItems", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numItems' field"))
	}
	m.NumItems = numItems

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterWriteVarResponse")
	}

	return m, nil
}

func (m *_S7ParameterWriteVarResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterWriteVarResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterWriteVarResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "numItems", m.GetNumItems(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numItems' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterWriteVarResponse")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterWriteVarResponse) IsS7ParameterWriteVarResponse() {}

func (m *_S7ParameterWriteVarResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterWriteVarResponse) deepCopy() *_S7ParameterWriteVarResponse {
	if m == nil {
		return nil
	}
	_S7ParameterWriteVarResponseCopy := &_S7ParameterWriteVarResponse{
		m.S7ParameterContract.(*_S7Parameter).deepCopy(),
		m.NumItems,
	}
	m.S7ParameterContract.(*_S7Parameter)._SubType = m
	return _S7ParameterWriteVarResponseCopy
}

func (m *_S7ParameterWriteVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
