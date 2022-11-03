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

// BACnetConstructedDataEventDetectionEnable is the corresponding interface of BACnetConstructedDataEventDetectionEnable
type BACnetConstructedDataEventDetectionEnable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventDetectionEnable returns EventDetectionEnable (property field)
	GetEventDetectionEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataEventDetectionEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventDetectionEnable.
// This is useful for switch cases.
type BACnetConstructedDataEventDetectionEnableExactly interface {
	BACnetConstructedDataEventDetectionEnable
	isBACnetConstructedDataEventDetectionEnable() bool
}

// _BACnetConstructedDataEventDetectionEnable is the data-structure of this message
type _BACnetConstructedDataEventDetectionEnable struct {
	*_BACnetConstructedData
	EventDetectionEnable BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_DETECTION_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetEventDetectionEnable() BACnetApplicationTagBoolean {
	return m.EventDetectionEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventDetectionEnable) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetEventDetectionEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventDetectionEnable factory function for _BACnetConstructedDataEventDetectionEnable
func NewBACnetConstructedDataEventDetectionEnable(eventDetectionEnable BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventDetectionEnable {
	_result := &_BACnetConstructedDataEventDetectionEnable{
		EventDetectionEnable:   eventDetectionEnable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventDetectionEnable(structType interface{}) BACnetConstructedDataEventDetectionEnable {
	if casted, ok := structType.(BACnetConstructedDataEventDetectionEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventDetectionEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetTypeName() string {
	return "BACnetConstructedDataEventDetectionEnable"
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventDetectionEnable)
	lengthInBits += m.EventDetectionEnable.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventDetectionEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventDetectionEnableParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventDetectionEnable, error) {
	return BACnetConstructedDataEventDetectionEnableParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataEventDetectionEnableParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventDetectionEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventDetectionEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventDetectionEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventDetectionEnable)
	if pullErr := readBuffer.PullContext("eventDetectionEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventDetectionEnable")
	}
	_eventDetectionEnable, _eventDetectionEnableErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _eventDetectionEnableErr != nil {
		return nil, errors.Wrap(_eventDetectionEnableErr, "Error parsing 'eventDetectionEnable' field of BACnetConstructedDataEventDetectionEnable")
	}
	eventDetectionEnable := _eventDetectionEnable.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("eventDetectionEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventDetectionEnable")
	}

	// Virtual field
	_actualValue := eventDetectionEnable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventDetectionEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventDetectionEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventDetectionEnable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EventDetectionEnable: eventDetectionEnable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventDetectionEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventDetectionEnable) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventDetectionEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventDetectionEnable")
		}

		// Simple Field (eventDetectionEnable)
		if pushErr := writeBuffer.PushContext("eventDetectionEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventDetectionEnable")
		}
		_eventDetectionEnableErr := writeBuffer.WriteSerializable(m.GetEventDetectionEnable())
		if popErr := writeBuffer.PopContext("eventDetectionEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventDetectionEnable")
		}
		if _eventDetectionEnableErr != nil {
			return errors.Wrap(_eventDetectionEnableErr, "Error serializing 'eventDetectionEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventDetectionEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventDetectionEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventDetectionEnable) isBACnetConstructedDataEventDetectionEnable() bool {
	return true
}

func (m *_BACnetConstructedDataEventDetectionEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
