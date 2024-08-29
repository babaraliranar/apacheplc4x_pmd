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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BitFieldDefinition is the corresponding interface of BitFieldDefinition
type BitFieldDefinition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetStartingBitPosition returns StartingBitPosition (property field)
	GetStartingBitPosition() uint32
	// GetEndingBitPosition returns EndingBitPosition (property field)
	GetEndingBitPosition() uint32
}

// BitFieldDefinitionExactly can be used when we want exactly this type and not a type which fulfills BitFieldDefinition.
// This is useful for switch cases.
type BitFieldDefinitionExactly interface {
	BitFieldDefinition
	isBitFieldDefinition() bool
}

// _BitFieldDefinition is the data-structure of this message
type _BitFieldDefinition struct {
	*_ExtensionObjectDefinition
	Name                PascalString
	Description         LocalizedText
	StartingBitPosition uint32
	EndingBitPosition   uint32
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BitFieldDefinition) GetIdentifier() string {
	return "32423"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BitFieldDefinition) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BitFieldDefinition) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BitFieldDefinition) GetName() PascalString {
	return m.Name
}

func (m *_BitFieldDefinition) GetDescription() LocalizedText {
	return m.Description
}

func (m *_BitFieldDefinition) GetStartingBitPosition() uint32 {
	return m.StartingBitPosition
}

func (m *_BitFieldDefinition) GetEndingBitPosition() uint32 {
	return m.EndingBitPosition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBitFieldDefinition factory function for _BitFieldDefinition
func NewBitFieldDefinition(name PascalString, description LocalizedText, startingBitPosition uint32, endingBitPosition uint32) *_BitFieldDefinition {
	_result := &_BitFieldDefinition{
		Name:                       name,
		Description:                description,
		StartingBitPosition:        startingBitPosition,
		EndingBitPosition:          endingBitPosition,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBitFieldDefinition(structType any) BitFieldDefinition {
	if casted, ok := structType.(BitFieldDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*BitFieldDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_BitFieldDefinition) GetTypeName() string {
	return "BitFieldDefinition"
}

func (m *_BitFieldDefinition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (startingBitPosition)
	lengthInBits += 32

	// Simple field (endingBitPosition)
	lengthInBits += 32

	return lengthInBits
}

func (m *_BitFieldDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BitFieldDefinitionParse(ctx context.Context, theBytes []byte, identifier string) (BitFieldDefinition, error) {
	return BitFieldDefinitionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BitFieldDefinitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BitFieldDefinition, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BitFieldDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BitFieldDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of BitFieldDefinition")
	}
	name := _name.(PascalString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	// Simple Field (description)
	if pullErr := readBuffer.PullContext("description"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for description")
	}
	_description, _descriptionErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _descriptionErr != nil {
		return nil, errors.Wrap(_descriptionErr, "Error parsing 'description' field of BitFieldDefinition")
	}
	description := _description.(LocalizedText)
	if closeErr := readBuffer.CloseContext("description"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for description")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of BitFieldDefinition")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of BitFieldDefinition")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (startingBitPosition)
	_startingBitPosition, _startingBitPositionErr := readBuffer.ReadUint32("startingBitPosition", 32)
	if _startingBitPositionErr != nil {
		return nil, errors.Wrap(_startingBitPositionErr, "Error parsing 'startingBitPosition' field of BitFieldDefinition")
	}
	startingBitPosition := _startingBitPosition

	// Simple Field (endingBitPosition)
	_endingBitPosition, _endingBitPositionErr := readBuffer.ReadUint32("endingBitPosition", 32)
	if _endingBitPositionErr != nil {
		return nil, errors.Wrap(_endingBitPositionErr, "Error parsing 'endingBitPosition' field of BitFieldDefinition")
	}
	endingBitPosition := _endingBitPosition

	if closeErr := readBuffer.CloseContext("BitFieldDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BitFieldDefinition")
	}

	// Create a partially initialized instance
	_child := &_BitFieldDefinition{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Name:                       name,
		Description:                description,
		StartingBitPosition:        startingBitPosition,
		EndingBitPosition:          endingBitPosition,
		reservedField0:             reservedField0,
		reservedField1:             reservedField1,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BitFieldDefinition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BitFieldDefinition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BitFieldDefinition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BitFieldDefinition")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		// Simple Field (description)
		if pushErr := writeBuffer.PushContext("description"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for description")
		}
		_descriptionErr := writeBuffer.WriteSerializable(ctx, m.GetDescription())
		if popErr := writeBuffer.PopContext("description"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for description")
		}
		if _descriptionErr != nil {
			return errors.Wrap(_descriptionErr, "Error serializing 'description' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved bool = bool(false)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": bool(false),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteBit("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (startingBitPosition)
		startingBitPosition := uint32(m.GetStartingBitPosition())
		_startingBitPositionErr := writeBuffer.WriteUint32("startingBitPosition", 32, uint32((startingBitPosition)))
		if _startingBitPositionErr != nil {
			return errors.Wrap(_startingBitPositionErr, "Error serializing 'startingBitPosition' field")
		}

		// Simple Field (endingBitPosition)
		endingBitPosition := uint32(m.GetEndingBitPosition())
		_endingBitPositionErr := writeBuffer.WriteUint32("endingBitPosition", 32, uint32((endingBitPosition)))
		if _endingBitPositionErr != nil {
			return errors.Wrap(_endingBitPositionErr, "Error serializing 'endingBitPosition' field")
		}

		if popErr := writeBuffer.PopContext("BitFieldDefinition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BitFieldDefinition")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BitFieldDefinition) isBitFieldDefinition() bool {
	return true
}

func (m *_BitFieldDefinition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
