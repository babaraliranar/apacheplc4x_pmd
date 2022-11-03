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

// BACnetConstructedDataPulseConverterAdjustValue is the corresponding interface of BACnetConstructedDataPulseConverterAdjustValue
type BACnetConstructedDataPulseConverterAdjustValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataPulseConverterAdjustValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPulseConverterAdjustValue.
// This is useful for switch cases.
type BACnetConstructedDataPulseConverterAdjustValueExactly interface {
	BACnetConstructedDataPulseConverterAdjustValue
	isBACnetConstructedDataPulseConverterAdjustValue() bool
}

// _BACnetConstructedDataPulseConverterAdjustValue is the data-structure of this message
type _BACnetConstructedDataPulseConverterAdjustValue struct {
	*_BACnetConstructedData
	AdjustValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_PULSE_CONVERTER
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetAdjustValue() BACnetApplicationTagReal {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPulseConverterAdjustValue factory function for _BACnetConstructedDataPulseConverterAdjustValue
func NewBACnetConstructedDataPulseConverterAdjustValue(adjustValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPulseConverterAdjustValue {
	_result := &_BACnetConstructedDataPulseConverterAdjustValue{
		AdjustValue:            adjustValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPulseConverterAdjustValue(structType interface{}) BACnetConstructedDataPulseConverterAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataPulseConverterAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPulseConverterAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataPulseConverterAdjustValue"
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPulseConverterAdjustValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPulseConverterAdjustValue, error) {
	return BACnetConstructedDataPulseConverterAdjustValueParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataPulseConverterAdjustValueParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPulseConverterAdjustValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPulseConverterAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPulseConverterAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (adjustValue)
	if pullErr := readBuffer.PullContext("adjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for adjustValue")
	}
	_adjustValue, _adjustValueErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _adjustValueErr != nil {
		return nil, errors.Wrap(_adjustValueErr, "Error parsing 'adjustValue' field of BACnetConstructedDataPulseConverterAdjustValue")
	}
	adjustValue := _adjustValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("adjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for adjustValue")
	}

	// Virtual field
	_actualValue := adjustValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPulseConverterAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPulseConverterAdjustValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPulseConverterAdjustValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AdjustValue: adjustValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPulseConverterAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPulseConverterAdjustValue")
		}

		// Simple Field (adjustValue)
		if pushErr := writeBuffer.PushContext("adjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for adjustValue")
		}
		_adjustValueErr := writeBuffer.WriteSerializable(m.GetAdjustValue())
		if popErr := writeBuffer.PopContext("adjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for adjustValue")
		}
		if _adjustValueErr != nil {
			return errors.Wrap(_adjustValueErr, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPulseConverterAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPulseConverterAdjustValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) isBACnetConstructedDataPulseConverterAdjustValue() bool {
	return true
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
