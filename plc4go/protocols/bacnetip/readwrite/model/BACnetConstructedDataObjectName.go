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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataObjectName is the data-structure of this message
type BACnetConstructedDataObjectName struct {
	*BACnetConstructedData
	ObjectName *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataObjectName is the corresponding interface of BACnetConstructedDataObjectName
type IBACnetConstructedDataObjectName interface {
	IBACnetConstructedData
	// GetObjectName returns ObjectName (property field)
	GetObjectName() *BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataObjectName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataObjectName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataObjectName) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataObjectName) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataObjectName) GetObjectName() *BACnetApplicationTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataObjectName) GetActualValue() *BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetObjectName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataObjectName factory function for BACnetConstructedDataObjectName
func NewBACnetConstructedDataObjectName(objectName *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataObjectName {
	_result := &BACnetConstructedDataObjectName{
		ObjectName:            objectName,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataObjectName(structType interface{}) *BACnetConstructedDataObjectName {
	if casted, ok := structType.(BACnetConstructedDataObjectName); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectName); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataObjectName(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataObjectName(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataObjectName) GetTypeName() string {
	return "BACnetConstructedDataObjectName"
}

func (m *BACnetConstructedDataObjectName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataObjectName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataObjectName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataObjectNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataObjectName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectName)
	if pullErr := readBuffer.PullContext("objectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectName")
	}
	_objectName, _objectNameErr := BACnetApplicationTagParse(readBuffer)
	if _objectNameErr != nil {
		return nil, errors.Wrap(_objectNameErr, "Error parsing 'objectName' field")
	}
	objectName := CastBACnetApplicationTagCharacterString(_objectName)
	if closeErr := readBuffer.CloseContext("objectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectName")
	}

	// Virtual field
	_actualValue := objectName
	actualValue := CastBACnetApplicationTagCharacterString(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectName")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataObjectName{
		ObjectName:            CastBACnetApplicationTagCharacterString(objectName),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataObjectName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectName")
		}

		// Simple Field (objectName)
		if pushErr := writeBuffer.PushContext("objectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectName")
		}
		_objectNameErr := writeBuffer.WriteSerializable(m.ObjectName)
		if popErr := writeBuffer.PopContext("objectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectName")
		}
		if _objectNameErr != nil {
			return errors.Wrap(_objectNameErr, "Error serializing 'objectName' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataObjectName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
