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

// BACnetConstructedDataFaultParameters is the corresponding interface of BACnetConstructedDataFaultParameters
type BACnetConstructedDataFaultParameters interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultParameters returns FaultParameters (property field)
	GetFaultParameters() BACnetFaultParameter
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetFaultParameter
}

// BACnetConstructedDataFaultParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFaultParameters.
// This is useful for switch cases.
type BACnetConstructedDataFaultParametersExactly interface {
	BACnetConstructedDataFaultParameters
	isBACnetConstructedDataFaultParameters() bool
}

// _BACnetConstructedDataFaultParameters is the data-structure of this message
type _BACnetConstructedDataFaultParameters struct {
	*_BACnetConstructedData
	FaultParameters BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFaultParameters) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_PARAMETERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFaultParameters) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFaultParameters) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetFaultParameters() BACnetFaultParameter {
	return m.FaultParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetActualValue() BACnetFaultParameter {
	return CastBACnetFaultParameter(m.GetFaultParameters())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFaultParameters factory function for _BACnetConstructedDataFaultParameters
func NewBACnetConstructedDataFaultParameters(faultParameters BACnetFaultParameter, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFaultParameters {
	_result := &_BACnetConstructedDataFaultParameters{
		FaultParameters:        faultParameters,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFaultParameters(structType interface{}) BACnetConstructedDataFaultParameters {
	if casted, ok := structType.(BACnetConstructedDataFaultParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFaultParameters) GetTypeName() string {
	return "BACnetConstructedDataFaultParameters"
}

func (m *_BACnetConstructedDataFaultParameters) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataFaultParameters) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultParameters)
	lengthInBits += m.FaultParameters.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFaultParameters) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFaultParametersParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFaultParameters, error) {
	return BACnetConstructedDataFaultParametersParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataFaultParametersParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFaultParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFaultParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultParameters)
	if pullErr := readBuffer.PullContext("faultParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultParameters")
	}
	_faultParameters, _faultParametersErr := BACnetFaultParameterParseWithBuffer(readBuffer)
	if _faultParametersErr != nil {
		return nil, errors.Wrap(_faultParametersErr, "Error parsing 'faultParameters' field of BACnetConstructedDataFaultParameters")
	}
	faultParameters := _faultParameters.(BACnetFaultParameter)
	if closeErr := readBuffer.CloseContext("faultParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultParameters")
	}

	// Virtual field
	_actualValue := faultParameters
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFaultParameters")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFaultParameters{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultParameters: faultParameters,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFaultParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFaultParameters) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFaultParameters")
		}

		// Simple Field (faultParameters)
		if pushErr := writeBuffer.PushContext("faultParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultParameters")
		}
		_faultParametersErr := writeBuffer.WriteSerializable(m.GetFaultParameters())
		if popErr := writeBuffer.PopContext("faultParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultParameters")
		}
		if _faultParametersErr != nil {
			return errors.Wrap(_faultParametersErr, "Error serializing 'faultParameters' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFaultParameters")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFaultParameters) isBACnetConstructedDataFaultParameters() bool {
	return true
}

func (m *_BACnetConstructedDataFaultParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
