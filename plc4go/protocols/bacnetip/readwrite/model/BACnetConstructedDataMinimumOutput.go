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

// BACnetConstructedDataMinimumOutput is the corresponding interface of BACnetConstructedDataMinimumOutput
type BACnetConstructedDataMinimumOutput interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinimumOutput returns MinimumOutput (property field)
	GetMinimumOutput() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataMinimumOutputExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMinimumOutput.
// This is useful for switch cases.
type BACnetConstructedDataMinimumOutputExactly interface {
	BACnetConstructedDataMinimumOutput
	isBACnetConstructedDataMinimumOutput() bool
}

// _BACnetConstructedDataMinimumOutput is the data-structure of this message
type _BACnetConstructedDataMinimumOutput struct {
	*_BACnetConstructedData
	MinimumOutput BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumOutput) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_OUTPUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumOutput) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMinimumOutput) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetMinimumOutput() BACnetApplicationTagReal {
	return m.MinimumOutput
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetMinimumOutput())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumOutput factory function for _BACnetConstructedDataMinimumOutput
func NewBACnetConstructedDataMinimumOutput(minimumOutput BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumOutput {
	_result := &_BACnetConstructedDataMinimumOutput{
		MinimumOutput:          minimumOutput,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumOutput(structType interface{}) BACnetConstructedDataMinimumOutput {
	if casted, ok := structType.(BACnetConstructedDataMinimumOutput); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOutput); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumOutput) GetTypeName() string {
	return "BACnetConstructedDataMinimumOutput"
}

func (m *_BACnetConstructedDataMinimumOutput) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMinimumOutput) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumOutput)
	lengthInBits += m.MinimumOutput.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumOutput) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumOutputParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMinimumOutput, error) {
	return BACnetConstructedDataMinimumOutputParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataMinimumOutputParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMinimumOutput, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOutput")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumOutput)
	if pullErr := readBuffer.PullContext("minimumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumOutput")
	}
	_minimumOutput, _minimumOutputErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _minimumOutputErr != nil {
		return nil, errors.Wrap(_minimumOutputErr, "Error parsing 'minimumOutput' field of BACnetConstructedDataMinimumOutput")
	}
	minimumOutput := _minimumOutput.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("minimumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumOutput")
	}

	// Virtual field
	_actualValue := minimumOutput
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOutput")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMinimumOutput{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MinimumOutput: minimumOutput,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMinimumOutput) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumOutput) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOutput")
		}

		// Simple Field (minimumOutput)
		if pushErr := writeBuffer.PushContext("minimumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumOutput")
		}
		_minimumOutputErr := writeBuffer.WriteSerializable(m.GetMinimumOutput())
		if popErr := writeBuffer.PopContext("minimumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumOutput")
		}
		if _minimumOutputErr != nil {
			return errors.Wrap(_minimumOutputErr, "Error serializing 'minimumOutput' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOutput")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumOutput) isBACnetConstructedDataMinimumOutput() bool {
	return true
}

func (m *_BACnetConstructedDataMinimumOutput) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
