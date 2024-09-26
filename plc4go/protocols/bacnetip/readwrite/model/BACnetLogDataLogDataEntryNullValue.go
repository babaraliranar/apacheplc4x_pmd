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

// BACnetLogDataLogDataEntryNullValue is the corresponding interface of BACnetLogDataLogDataEntryNullValue
type BACnetLogDataLogDataEntryNullValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogDataLogDataEntry
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetContextTagNull
	// IsBACnetLogDataLogDataEntryNullValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryNullValue()
}

// _BACnetLogDataLogDataEntryNullValue is the data-structure of this message
type _BACnetLogDataLogDataEntryNullValue struct {
	BACnetLogDataLogDataEntryContract
	NullValue BACnetContextTagNull
}

var _ BACnetLogDataLogDataEntryNullValue = (*_BACnetLogDataLogDataEntryNullValue)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryNullValue)(nil)

// NewBACnetLogDataLogDataEntryNullValue factory function for _BACnetLogDataLogDataEntryNullValue
func NewBACnetLogDataLogDataEntryNullValue(peekedTagHeader BACnetTagHeader, nullValue BACnetContextTagNull) *_BACnetLogDataLogDataEntryNullValue {
	if nullValue == nil {
		panic("nullValue of type BACnetContextTagNull for BACnetLogDataLogDataEntryNullValue must not be nil")
	}
	_result := &_BACnetLogDataLogDataEntryNullValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		NullValue:                         nullValue,
	}
	_result.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = _result
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

func (m *_BACnetLogDataLogDataEntryNullValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryNullValue) GetNullValue() BACnetContextTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryNullValue(structType any) BACnetLogDataLogDataEntryNullValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryNullValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryNullValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryNullValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryNullValue"
}

func (m *_BACnetLogDataLogDataEntryNullValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryNullValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryNullValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryNullValue BACnetLogDataLogDataEntryNullValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryNullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryNullValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetContextTagNull](ctx, "nullValue", ReadComplex[BACnetContextTagNull](BACnetContextTagParseWithBufferProducer[BACnetContextTagNull]((uint8)(uint8(6)), (BACnetDataType)(BACnetDataType_NULL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryNullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryNullValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryNullValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryNullValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryNullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryNullValue")
		}

		if err := WriteSimpleField[BACnetContextTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetContextTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryNullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryNullValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryNullValue) IsBACnetLogDataLogDataEntryNullValue() {}

func (m *_BACnetLogDataLogDataEntryNullValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataEntryNullValue) deepCopy() *_BACnetLogDataLogDataEntryNullValue {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataEntryNullValueCopy := &_BACnetLogDataLogDataEntryNullValue{
		m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).deepCopy(),
		m.NullValue.DeepCopy().(BACnetContextTagNull),
	}
	m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = m
	return _BACnetLogDataLogDataEntryNullValueCopy
}

func (m *_BACnetLogDataLogDataEntryNullValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
