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

// BACnetLogDataLogDataEntryAnyValue is the corresponding interface of BACnetLogDataLogDataEntryAnyValue
type BACnetLogDataLogDataEntryAnyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogDataLogDataEntry
	// GetAnyValue returns AnyValue (property field)
	GetAnyValue() BACnetConstructedData
	// IsBACnetLogDataLogDataEntryAnyValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryAnyValue()
}

// _BACnetLogDataLogDataEntryAnyValue is the data-structure of this message
type _BACnetLogDataLogDataEntryAnyValue struct {
	BACnetLogDataLogDataEntryContract
	AnyValue BACnetConstructedData
}

var _ BACnetLogDataLogDataEntryAnyValue = (*_BACnetLogDataLogDataEntryAnyValue)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryAnyValue)(nil)

// NewBACnetLogDataLogDataEntryAnyValue factory function for _BACnetLogDataLogDataEntryAnyValue
func NewBACnetLogDataLogDataEntryAnyValue(peekedTagHeader BACnetTagHeader, anyValue BACnetConstructedData) *_BACnetLogDataLogDataEntryAnyValue {
	_result := &_BACnetLogDataLogDataEntryAnyValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		AnyValue:                          anyValue,
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

func (m *_BACnetLogDataLogDataEntryAnyValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryAnyValue) GetAnyValue() BACnetConstructedData {
	return m.AnyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryAnyValue(structType any) BACnetLogDataLogDataEntryAnyValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryAnyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryAnyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryAnyValue"
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Optional Field (anyValue)
	if m.AnyValue != nil {
		lengthInBits += m.AnyValue.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryAnyValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryAnyValue BACnetLogDataLogDataEntryAnyValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryAnyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryAnyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var anyValue BACnetConstructedData
	_anyValue, err := ReadOptionalField[BACnetConstructedData](ctx, "anyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(8)), (BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'anyValue' field"))
	}
	if _anyValue != nil {
		anyValue = *_anyValue
		m.AnyValue = anyValue
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryAnyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryAnyValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryAnyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryAnyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryAnyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryAnyValue")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "anyValue", GetRef(m.GetAnyValue()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'anyValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryAnyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryAnyValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryAnyValue) IsBACnetLogDataLogDataEntryAnyValue() {}

func (m *_BACnetLogDataLogDataEntryAnyValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataEntryAnyValue) deepCopy() *_BACnetLogDataLogDataEntryAnyValue {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataEntryAnyValueCopy := &_BACnetLogDataLogDataEntryAnyValue{
		m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).deepCopy(),
		m.AnyValue.DeepCopy().(BACnetConstructedData),
	}
	m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = m
	return _BACnetLogDataLogDataEntryAnyValueCopy
}

func (m *_BACnetLogDataLogDataEntryAnyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
