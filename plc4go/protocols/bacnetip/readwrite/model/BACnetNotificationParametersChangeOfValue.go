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

// BACnetNotificationParametersChangeOfValue is the corresponding interface of BACnetNotificationParametersChangeOfValue
type BACnetNotificationParametersChangeOfValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewValue returns NewValue (property field)
	GetNewValue() BACnetNotificationParametersChangeOfValueNewValue
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersChangeOfValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfValue()
}

// _BACnetNotificationParametersChangeOfValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfValue struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	NewValue        BACnetNotificationParametersChangeOfValueNewValue
	StatusFlags     BACnetStatusFlagsTagged
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersChangeOfValue = (*_BACnetNotificationParametersChangeOfValue)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersChangeOfValue)(nil)

// NewBACnetNotificationParametersChangeOfValue factory function for _BACnetNotificationParametersChangeOfValue
func NewBACnetNotificationParametersChangeOfValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, newValue BACnetNotificationParametersChangeOfValueNewValue, statusFlags BACnetStatusFlagsTagged, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfValue {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfValue must not be nil")
	}
	if newValue == nil {
		panic("newValue of type BACnetNotificationParametersChangeOfValueNewValue for BACnetNotificationParametersChangeOfValue must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersChangeOfValue must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfValue must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfValue{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		NewValue:                             newValue,
		StatusFlags:                          statusFlags,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
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

func (m *_BACnetNotificationParametersChangeOfValue) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValue) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfValue) GetNewValue() BACnetNotificationParametersChangeOfValueNewValue {
	return m.NewValue
}

func (m *_BACnetNotificationParametersChangeOfValue) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfValue) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValue(structType any) BACnetNotificationParametersChangeOfValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValue"
}

func (m *_BACnetNotificationParametersChangeOfValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (newValue)
	lengthInBits += m.NewValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersChangeOfValue BACnetNotificationParametersChangeOfValue, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	newValue, err := ReadSimpleField[BACnetNotificationParametersChangeOfValueNewValue](ctx, "newValue", ReadComplex[BACnetNotificationParametersChangeOfValueNewValue](BACnetNotificationParametersChangeOfValueNewValueParseWithBufferProducer[BACnetNotificationParametersChangeOfValueNewValue]((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'newValue' field"))
	}
	m.NewValue = newValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValue")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValue")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetNotificationParametersChangeOfValueNewValue](ctx, "newValue", m.GetNewValue(), WriteComplex[BACnetNotificationParametersChangeOfValueNewValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'newValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValue")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfValue) IsBACnetNotificationParametersChangeOfValue() {}

func (m *_BACnetNotificationParametersChangeOfValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfValue) deepCopy() *_BACnetNotificationParametersChangeOfValue {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfValueCopy := &_BACnetNotificationParametersChangeOfValue{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.InnerOpeningTag.DeepCopy().(BACnetOpeningTag),
		m.NewValue.DeepCopy().(BACnetNotificationParametersChangeOfValueNewValue),
		m.StatusFlags.DeepCopy().(BACnetStatusFlagsTagged),
		m.InnerClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersChangeOfValueCopy
}

func (m *_BACnetNotificationParametersChangeOfValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
