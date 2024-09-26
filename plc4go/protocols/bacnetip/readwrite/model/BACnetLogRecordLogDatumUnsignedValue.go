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

// BACnetLogRecordLogDatumUnsignedValue is the corresponding interface of BACnetLogRecordLogDatumUnsignedValue
type BACnetLogRecordLogDatumUnsignedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetContextTagUnsignedInteger
	// IsBACnetLogRecordLogDatumUnsignedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumUnsignedValue()
}

// _BACnetLogRecordLogDatumUnsignedValue is the data-structure of this message
type _BACnetLogRecordLogDatumUnsignedValue struct {
	BACnetLogRecordLogDatumContract
	UnsignedValue BACnetContextTagUnsignedInteger
}

var _ BACnetLogRecordLogDatumUnsignedValue = (*_BACnetLogRecordLogDatumUnsignedValue)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumUnsignedValue)(nil)

// NewBACnetLogRecordLogDatumUnsignedValue factory function for _BACnetLogRecordLogDatumUnsignedValue
func NewBACnetLogRecordLogDatumUnsignedValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, unsignedValue BACnetContextTagUnsignedInteger, tagNumber uint8) *_BACnetLogRecordLogDatumUnsignedValue {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetContextTagUnsignedInteger for BACnetLogRecordLogDatumUnsignedValue must not be nil")
	}
	_result := &_BACnetLogRecordLogDatumUnsignedValue{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		UnsignedValue:                   unsignedValue,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetUnsignedValue() BACnetContextTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumUnsignedValue(structType any) BACnetLogRecordLogDatumUnsignedValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumUnsignedValue"
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumUnsignedValue BACnetLogRecordLogDatumUnsignedValue, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumUnsignedValue")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumUnsignedValue")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumUnsignedValue")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) IsBACnetLogRecordLogDatumUnsignedValue() {}

func (m *_BACnetLogRecordLogDatumUnsignedValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) deepCopy() *_BACnetLogRecordLogDatumUnsignedValue {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumUnsignedValueCopy := &_BACnetLogRecordLogDatumUnsignedValue{
		m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).deepCopy(),
		m.UnsignedValue.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumUnsignedValueCopy
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
