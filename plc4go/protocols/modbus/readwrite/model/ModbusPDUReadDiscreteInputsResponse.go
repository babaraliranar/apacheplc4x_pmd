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

// ModbusPDUReadDiscreteInputsResponse is the corresponding interface of ModbusPDUReadDiscreteInputsResponse
type ModbusPDUReadDiscreteInputsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsModbusPDUReadDiscreteInputsResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadDiscreteInputsResponse()
}

// _ModbusPDUReadDiscreteInputsResponse is the data-structure of this message
type _ModbusPDUReadDiscreteInputsResponse struct {
	ModbusPDUContract
	Value []byte
}

var _ ModbusPDUReadDiscreteInputsResponse = (*_ModbusPDUReadDiscreteInputsResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadDiscreteInputsResponse)(nil)

// NewModbusPDUReadDiscreteInputsResponse factory function for _ModbusPDUReadDiscreteInputsResponse
func NewModbusPDUReadDiscreteInputsResponse(value []byte) *_ModbusPDUReadDiscreteInputsResponse {
	_result := &_ModbusPDUReadDiscreteInputsResponse{
		ModbusPDUContract: NewModbusPDU(),
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadDiscreteInputsResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetFunctionFlag() uint8 {
	return 0x02
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadDiscreteInputsResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadDiscreteInputsResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadDiscreteInputsResponse(structType any) ModbusPDUReadDiscreteInputsResponse {
	if casted, ok := structType.(ModbusPDUReadDiscreteInputsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadDiscreteInputsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetTypeName() string {
	return "ModbusPDUReadDiscreteInputsResponse"
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadDiscreteInputsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadDiscreteInputsResponse ModbusPDUReadDiscreteInputsResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadDiscreteInputsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadDiscreteInputsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDiscreteInputsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadDiscreteInputsResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReadDiscreteInputsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadDiscreteInputsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDiscreteInputsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadDiscreteInputsResponse")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDiscreteInputsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadDiscreteInputsResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadDiscreteInputsResponse) IsModbusPDUReadDiscreteInputsResponse() {}

func (m *_ModbusPDUReadDiscreteInputsResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadDiscreteInputsResponse) deepCopy() *_ModbusPDUReadDiscreteInputsResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUReadDiscreteInputsResponseCopy := &_ModbusPDUReadDiscreteInputsResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Value),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadDiscreteInputsResponseCopy
}

func (m *_ModbusPDUReadDiscreteInputsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
