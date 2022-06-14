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

// BACnetConstructedDataValidSamples is the data-structure of this message
type BACnetConstructedDataValidSamples struct {
	*BACnetConstructedData
	ValidSamples *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataValidSamples is the corresponding interface of BACnetConstructedDataValidSamples
type IBACnetConstructedDataValidSamples interface {
	IBACnetConstructedData
	// GetValidSamples returns ValidSamples (property field)
	GetValidSamples() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataValidSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataValidSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALID_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataValidSamples) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataValidSamples) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataValidSamples) GetValidSamples() *BACnetApplicationTagUnsignedInteger {
	return m.ValidSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataValidSamples) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetValidSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValidSamples factory function for BACnetConstructedDataValidSamples
func NewBACnetConstructedDataValidSamples(validSamples *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataValidSamples {
	_result := &BACnetConstructedDataValidSamples{
		ValidSamples:          validSamples,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataValidSamples(structType interface{}) *BACnetConstructedDataValidSamples {
	if casted, ok := structType.(BACnetConstructedDataValidSamples); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValidSamples); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataValidSamples(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataValidSamples(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataValidSamples) GetTypeName() string {
	return "BACnetConstructedDataValidSamples"
}

func (m *BACnetConstructedDataValidSamples) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataValidSamples) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (validSamples)
	lengthInBits += m.ValidSamples.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataValidSamples) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataValidSamplesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataValidSamples, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValidSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValidSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (validSamples)
	if pullErr := readBuffer.PullContext("validSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for validSamples")
	}
	_validSamples, _validSamplesErr := BACnetApplicationTagParse(readBuffer)
	if _validSamplesErr != nil {
		return nil, errors.Wrap(_validSamplesErr, "Error parsing 'validSamples' field")
	}
	validSamples := CastBACnetApplicationTagUnsignedInteger(_validSamples)
	if closeErr := readBuffer.CloseContext("validSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for validSamples")
	}

	// Virtual field
	_actualValue := validSamples
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValidSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValidSamples")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataValidSamples{
		ValidSamples:          CastBACnetApplicationTagUnsignedInteger(validSamples),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataValidSamples) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValidSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValidSamples")
		}

		// Simple Field (validSamples)
		if pushErr := writeBuffer.PushContext("validSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for validSamples")
		}
		_validSamplesErr := writeBuffer.WriteSerializable(m.ValidSamples)
		if popErr := writeBuffer.PopContext("validSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for validSamples")
		}
		if _validSamplesErr != nil {
			return errors.Wrap(_validSamplesErr, "Error serializing 'validSamples' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValidSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValidSamples")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataValidSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
