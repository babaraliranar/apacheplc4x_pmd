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

// BACnetConstructedDataLastCredentialRemovedTime is the corresponding interface of BACnetConstructedDataLastCredentialRemovedTime
type BACnetConstructedDataLastCredentialRemovedTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastCredentialRemovedTime returns LastCredentialRemovedTime (property field)
	GetLastCredentialRemovedTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataLastCredentialRemovedTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastCredentialRemovedTime.
// This is useful for switch cases.
type BACnetConstructedDataLastCredentialRemovedTimeExactly interface {
	BACnetConstructedDataLastCredentialRemovedTime
	isBACnetConstructedDataLastCredentialRemovedTime() bool
}

// _BACnetConstructedDataLastCredentialRemovedTime is the data-structure of this message
type _BACnetConstructedDataLastCredentialRemovedTime struct {
	*_BACnetConstructedData
	LastCredentialRemovedTime BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_REMOVED_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLastCredentialRemovedTime() BACnetDateTime {
	return m.LastCredentialRemovedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetLastCredentialRemovedTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastCredentialRemovedTime factory function for _BACnetConstructedDataLastCredentialRemovedTime
func NewBACnetConstructedDataLastCredentialRemovedTime(lastCredentialRemovedTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCredentialRemovedTime {
	_result := &_BACnetConstructedDataLastCredentialRemovedTime{
		LastCredentialRemovedTime: lastCredentialRemovedTime,
		_BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCredentialRemovedTime(structType interface{}) BACnetConstructedDataLastCredentialRemovedTime {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialRemovedTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialRemovedTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialRemovedTime"
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastCredentialRemovedTime)
	lengthInBits += m.LastCredentialRemovedTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastCredentialRemovedTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCredentialRemovedTime, error) {
	return BACnetConstructedDataLastCredentialRemovedTimeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataLastCredentialRemovedTimeParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCredentialRemovedTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialRemovedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCredentialRemovedTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCredentialRemovedTime)
	if pullErr := readBuffer.PullContext("lastCredentialRemovedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastCredentialRemovedTime")
	}
	_lastCredentialRemovedTime, _lastCredentialRemovedTimeErr := BACnetDateTimeParseWithBuffer(readBuffer)
	if _lastCredentialRemovedTimeErr != nil {
		return nil, errors.Wrap(_lastCredentialRemovedTimeErr, "Error parsing 'lastCredentialRemovedTime' field of BACnetConstructedDataLastCredentialRemovedTime")
	}
	lastCredentialRemovedTime := _lastCredentialRemovedTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("lastCredentialRemovedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastCredentialRemovedTime")
	}

	// Virtual field
	_actualValue := lastCredentialRemovedTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialRemovedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCredentialRemovedTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastCredentialRemovedTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastCredentialRemovedTime: lastCredentialRemovedTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialRemovedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCredentialRemovedTime")
		}

		// Simple Field (lastCredentialRemovedTime)
		if pushErr := writeBuffer.PushContext("lastCredentialRemovedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastCredentialRemovedTime")
		}
		_lastCredentialRemovedTimeErr := writeBuffer.WriteSerializable(m.GetLastCredentialRemovedTime())
		if popErr := writeBuffer.PopContext("lastCredentialRemovedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastCredentialRemovedTime")
		}
		if _lastCredentialRemovedTimeErr != nil {
			return errors.Wrap(_lastCredentialRemovedTimeErr, "Error serializing 'lastCredentialRemovedTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialRemovedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCredentialRemovedTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) isBACnetConstructedDataLastCredentialRemovedTime() bool {
	return true
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
