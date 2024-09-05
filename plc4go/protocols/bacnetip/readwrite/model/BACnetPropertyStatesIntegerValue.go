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

// BACnetPropertyStatesIntegerValue is the corresponding interface of BACnetPropertyStatesIntegerValue
type BACnetPropertyStatesIntegerValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetContextTagSignedInteger
	// IsBACnetPropertyStatesIntegerValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesIntegerValue()
}

// _BACnetPropertyStatesIntegerValue is the data-structure of this message
type _BACnetPropertyStatesIntegerValue struct {
	BACnetPropertyStatesContract
	IntegerValue BACnetContextTagSignedInteger
}

var _ BACnetPropertyStatesIntegerValue = (*_BACnetPropertyStatesIntegerValue)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesIntegerValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesIntegerValue) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesIntegerValue) GetIntegerValue() BACnetContextTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesIntegerValue factory function for _BACnetPropertyStatesIntegerValue
func NewBACnetPropertyStatesIntegerValue(integerValue BACnetContextTagSignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesIntegerValue {
	if integerValue == nil {
		panic("integerValue of type BACnetContextTagSignedInteger for BACnetPropertyStatesIntegerValue must not be nil")
	}
	_result := &_BACnetPropertyStatesIntegerValue{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		IntegerValue:                 integerValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesIntegerValue(structType any) BACnetPropertyStatesIntegerValue {
	if casted, ok := structType.(BACnetPropertyStatesIntegerValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesIntegerValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesIntegerValue) GetTypeName() string {
	return "BACnetPropertyStatesIntegerValue"
}

func (m *_BACnetPropertyStatesIntegerValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesIntegerValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesIntegerValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesIntegerValue BACnetPropertyStatesIntegerValue, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesIntegerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesIntegerValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(peekedTagNumber), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesIntegerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesIntegerValue")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesIntegerValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesIntegerValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesIntegerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesIntegerValue")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesIntegerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesIntegerValue")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesIntegerValue) IsBACnetPropertyStatesIntegerValue() {}

func (m *_BACnetPropertyStatesIntegerValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
