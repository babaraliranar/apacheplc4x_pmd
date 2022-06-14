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

// BACnetConstructedDataPositiveIntegerValueHighLimit is the data-structure of this message
type BACnetConstructedDataPositiveIntegerValueHighLimit struct {
	*BACnetConstructedData
	HighLimit *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPositiveIntegerValueHighLimit is the corresponding interface of BACnetConstructedDataPositiveIntegerValueHighLimit
type IBACnetConstructedDataPositiveIntegerValueHighLimit interface {
	IBACnetConstructedData
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetHighLimit() *BACnetApplicationTagUnsignedInteger {
	return m.HighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPositiveIntegerValueHighLimit factory function for BACnetConstructedDataPositiveIntegerValueHighLimit
func NewBACnetConstructedDataPositiveIntegerValueHighLimit(highLimit *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPositiveIntegerValueHighLimit {
	_result := &BACnetConstructedDataPositiveIntegerValueHighLimit{
		HighLimit:             highLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPositiveIntegerValueHighLimit(structType interface{}) *BACnetConstructedDataPositiveIntegerValueHighLimit {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueHighLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPositiveIntegerValueHighLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPositiveIntegerValueHighLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueHighLimit"
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPositiveIntegerValueHighLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPositiveIntegerValueHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (highLimit)
	if pullErr := readBuffer.PullContext("highLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highLimit")
	}
	_highLimit, _highLimitErr := BACnetApplicationTagParse(readBuffer)
	if _highLimitErr != nil {
		return nil, errors.Wrap(_highLimitErr, "Error parsing 'highLimit' field")
	}
	highLimit := CastBACnetApplicationTagUnsignedInteger(_highLimit)
	if closeErr := readBuffer.CloseContext("highLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highLimit")
	}

	// Virtual field
	_actualValue := highLimit
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueHighLimit")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPositiveIntegerValueHighLimit{
		HighLimit:             CastBACnetApplicationTagUnsignedInteger(highLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueHighLimit")
		}

		// Simple Field (highLimit)
		if pushErr := writeBuffer.PushContext("highLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highLimit")
		}
		_highLimitErr := writeBuffer.WriteSerializable(m.HighLimit)
		if popErr := writeBuffer.PopContext("highLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highLimit")
		}
		if _highLimitErr != nil {
			return errors.Wrap(_highLimitErr, "Error serializing 'highLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueHighLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPositiveIntegerValueHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
