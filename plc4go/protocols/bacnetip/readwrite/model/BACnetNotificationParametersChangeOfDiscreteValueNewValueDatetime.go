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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
type BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeExactly interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
	DateTimeValue BACnetDateTimeEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime(dateTimeValue BACnetDateTimeEnclosed, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime{
		DateTimeValue: dateTimeValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeParse(theBytes []byte, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime, error) {
	return BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field of BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime")
	}
	dateTimeValue := _dateTimeValue.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime{
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{
			TagNumber: tagNumber,
		},
		DateTimeValue: dateTimeValue,
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime")
		}

		// Simple Field (dateTimeValue)
		if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateTimeValue")
		}
		_dateTimeValueErr := writeBuffer.WriteSerializable(ctx, m.GetDateTimeValue())
		if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateTimeValue")
		}
		if _dateTimeValueErr != nil {
			return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) isBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
