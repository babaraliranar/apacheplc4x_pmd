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

// VariantUInt64 is the corresponding interface of VariantUInt64
type VariantUInt64 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []uint64
	// IsVariantUInt64 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantUInt64()
}

// _VariantUInt64 is the data-structure of this message
type _VariantUInt64 struct {
	VariantContract
	ArrayLength *int32
	Value       []uint64
}

var _ VariantUInt64 = (*_VariantUInt64)(nil)
var _ VariantRequirements = (*_VariantUInt64)(nil)

// NewVariantUInt64 factory function for _VariantUInt64
func NewVariantUInt64(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []uint64) *_VariantUInt64 {
	_result := &_VariantUInt64{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantUInt64) GetVariantType() uint8 {
	return uint8(9)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantUInt64) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantUInt64) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantUInt64) GetValue() []uint64 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantUInt64(structType any) VariantUInt64 {
	if casted, ok := structType.(VariantUInt64); ok {
		return casted
	}
	if casted, ok := structType.(*VariantUInt64); ok {
		return *casted
	}
	return nil
}

func (m *_VariantUInt64) GetTypeName() string {
	return "VariantUInt64"
}

func (m *_VariantUInt64) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 64 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantUInt64) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantUInt64) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantUInt64 VariantUInt64, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantUInt64"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantUInt64")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[uint64](ctx, "value", ReadUnsignedLong(readBuffer, uint8(64)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantUInt64"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantUInt64")
	}

	return m, nil
}

func (m *_VariantUInt64) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantUInt64) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantUInt64"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantUInt64")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteUnsignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantUInt64"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantUInt64")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantUInt64) IsVariantUInt64() {}

func (m *_VariantUInt64) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantUInt64) deepCopy() *_VariantUInt64 {
	if m == nil {
		return nil
	}
	_VariantUInt64Copy := &_VariantUInt64{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[uint64, uint64](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantUInt64Copy
}

func (m *_VariantUInt64) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
