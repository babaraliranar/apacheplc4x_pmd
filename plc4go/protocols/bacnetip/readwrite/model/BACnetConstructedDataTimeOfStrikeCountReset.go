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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTimeOfStrikeCountReset is the corresponding interface of BACnetConstructedDataTimeOfStrikeCountReset
type BACnetConstructedDataTimeOfStrikeCountReset interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeOfStrikeCountReset returns TimeOfStrikeCountReset (property field)
	GetTimeOfStrikeCountReset() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataTimeOfStrikeCountResetExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeOfStrikeCountReset.
// This is useful for switch cases.
type BACnetConstructedDataTimeOfStrikeCountResetExactly interface {
	BACnetConstructedDataTimeOfStrikeCountReset
	isBACnetConstructedDataTimeOfStrikeCountReset() bool
}

// _BACnetConstructedDataTimeOfStrikeCountReset is the data-structure of this message
type _BACnetConstructedDataTimeOfStrikeCountReset struct {
	*_BACnetConstructedData
	TimeOfStrikeCountReset BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_STRIKE_COUNT_RESET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetTimeOfStrikeCountReset() BACnetDateTime {
	return m.TimeOfStrikeCountReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetTimeOfStrikeCountReset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfStrikeCountReset factory function for _BACnetConstructedDataTimeOfStrikeCountReset
func NewBACnetConstructedDataTimeOfStrikeCountReset(timeOfStrikeCountReset BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfStrikeCountReset {
	_result := &_BACnetConstructedDataTimeOfStrikeCountReset{
		TimeOfStrikeCountReset: timeOfStrikeCountReset,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfStrikeCountReset(structType interface{}) BACnetConstructedDataTimeOfStrikeCountReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfStrikeCountReset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfStrikeCountReset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfStrikeCountReset"
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeOfStrikeCountReset)
	lengthInBits += m.TimeOfStrikeCountReset.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeOfStrikeCountResetParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeOfStrikeCountReset, error) {
	return BACnetConstructedDataTimeOfStrikeCountResetParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataTimeOfStrikeCountResetParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeOfStrikeCountReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfStrikeCountReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfStrikeCountReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeOfStrikeCountReset)
	if pullErr := readBuffer.PullContext("timeOfStrikeCountReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeOfStrikeCountReset")
	}
	_timeOfStrikeCountReset, _timeOfStrikeCountResetErr := BACnetDateTimeParseWithBuffer(readBuffer)
	if _timeOfStrikeCountResetErr != nil {
		return nil, errors.Wrap(_timeOfStrikeCountResetErr, "Error parsing 'timeOfStrikeCountReset' field of BACnetConstructedDataTimeOfStrikeCountReset")
	}
	timeOfStrikeCountReset := _timeOfStrikeCountReset.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("timeOfStrikeCountReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeOfStrikeCountReset")
	}

	// Virtual field
	_actualValue := timeOfStrikeCountReset
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfStrikeCountReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfStrikeCountReset")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeOfStrikeCountReset{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimeOfStrikeCountReset: timeOfStrikeCountReset,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfStrikeCountReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfStrikeCountReset")
		}

		// Simple Field (timeOfStrikeCountReset)
		if pushErr := writeBuffer.PushContext("timeOfStrikeCountReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeOfStrikeCountReset")
		}
		_timeOfStrikeCountResetErr := writeBuffer.WriteSerializable(m.GetTimeOfStrikeCountReset())
		if popErr := writeBuffer.PopContext("timeOfStrikeCountReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeOfStrikeCountReset")
		}
		if _timeOfStrikeCountResetErr != nil {
			return errors.Wrap(_timeOfStrikeCountResetErr, "Error serializing 'timeOfStrikeCountReset' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfStrikeCountReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfStrikeCountReset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) isBACnetConstructedDataTimeOfStrikeCountReset() bool {
	return true
}

func (m *_BACnetConstructedDataTimeOfStrikeCountReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
