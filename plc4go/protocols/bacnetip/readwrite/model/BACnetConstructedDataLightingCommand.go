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

// BACnetConstructedDataLightingCommand is the data-structure of this message
type BACnetConstructedDataLightingCommand struct {
	*BACnetConstructedData
	LightingCommand *BACnetLightingCommand

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLightingCommand is the corresponding interface of BACnetConstructedDataLightingCommand
type IBACnetConstructedDataLightingCommand interface {
	IBACnetConstructedData
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() *BACnetLightingCommand
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetLightingCommand
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

func (m *BACnetConstructedDataLightingCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLightingCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLightingCommand) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLightingCommand) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLightingCommand) GetLightingCommand() *BACnetLightingCommand {
	return m.LightingCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLightingCommand) GetActualValue() *BACnetLightingCommand {
	return CastBACnetLightingCommand(m.GetLightingCommand())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingCommand factory function for BACnetConstructedDataLightingCommand
func NewBACnetConstructedDataLightingCommand(lightingCommand *BACnetLightingCommand, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLightingCommand {
	_result := &BACnetConstructedDataLightingCommand{
		LightingCommand:       lightingCommand,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLightingCommand(structType interface{}) *BACnetConstructedDataLightingCommand {
	if casted, ok := structType.(BACnetConstructedDataLightingCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingCommand(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingCommand(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLightingCommand) GetTypeName() string {
	return "BACnetConstructedDataLightingCommand"
}

func (m *BACnetConstructedDataLightingCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLightingCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLightingCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLightingCommandParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLightingCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightingCommand)
	if pullErr := readBuffer.PullContext("lightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightingCommand")
	}
	_lightingCommand, _lightingCommandErr := BACnetLightingCommandParse(readBuffer)
	if _lightingCommandErr != nil {
		return nil, errors.Wrap(_lightingCommandErr, "Error parsing 'lightingCommand' field")
	}
	lightingCommand := CastBACnetLightingCommand(_lightingCommand)
	if closeErr := readBuffer.CloseContext("lightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightingCommand")
	}

	// Virtual field
	_actualValue := lightingCommand
	actualValue := CastBACnetLightingCommand(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommand")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLightingCommand{
		LightingCommand:       CastBACnetLightingCommand(lightingCommand),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLightingCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommand")
		}

		// Simple Field (lightingCommand)
		if pushErr := writeBuffer.PushContext("lightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightingCommand")
		}
		_lightingCommandErr := writeBuffer.WriteSerializable(m.LightingCommand)
		if popErr := writeBuffer.PopContext("lightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightingCommand")
		}
		if _lightingCommandErr != nil {
			return errors.Wrap(_lightingCommandErr, "Error serializing 'lightingCommand' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
