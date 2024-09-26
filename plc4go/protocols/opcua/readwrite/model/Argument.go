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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Argument is the corresponding interface of Argument
type Argument interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetDataType returns DataType (property field)
	GetDataType() NodeId
	// GetValueRank returns ValueRank (property field)
	GetValueRank() int32
	// GetNoOfArrayDimensions returns NoOfArrayDimensions (property field)
	GetNoOfArrayDimensions() int32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() []uint32
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// IsArgument is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsArgument()
}

// _Argument is the data-structure of this message
type _Argument struct {
	ExtensionObjectDefinitionContract
	Name                PascalString
	DataType            NodeId
	ValueRank           int32
	NoOfArrayDimensions int32
	ArrayDimensions     []uint32
	Description         LocalizedText
}

var _ Argument = (*_Argument)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_Argument)(nil)

// NewArgument factory function for _Argument
func NewArgument(name PascalString, dataType NodeId, valueRank int32, noOfArrayDimensions int32, arrayDimensions []uint32, description LocalizedText) *_Argument {
	if name == nil {
		panic("name of type PascalString for Argument must not be nil")
	}
	if dataType == nil {
		panic("dataType of type NodeId for Argument must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for Argument must not be nil")
	}
	_result := &_Argument{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		DataType:                          dataType,
		ValueRank:                         valueRank,
		NoOfArrayDimensions:               noOfArrayDimensions,
		ArrayDimensions:                   arrayDimensions,
		Description:                       description,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Argument) GetIdentifier() string {
	return "298"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Argument) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Argument) GetName() PascalString {
	return m.Name
}

func (m *_Argument) GetDataType() NodeId {
	return m.DataType
}

func (m *_Argument) GetValueRank() int32 {
	return m.ValueRank
}

func (m *_Argument) GetNoOfArrayDimensions() int32 {
	return m.NoOfArrayDimensions
}

func (m *_Argument) GetArrayDimensions() []uint32 {
	return m.ArrayDimensions
}

func (m *_Argument) GetDescription() LocalizedText {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastArgument(structType any) Argument {
	if casted, ok := structType.(Argument); ok {
		return casted
	}
	if casted, ok := structType.(*Argument); ok {
		return *casted
	}
	return nil
}

func (m *_Argument) GetTypeName() string {
	return "Argument"
}

func (m *_Argument) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (dataType)
	lengthInBits += m.DataType.GetLengthInBits(ctx)

	// Simple field (valueRank)
	lengthInBits += 32

	// Simple field (noOfArrayDimensions)
	lengthInBits += 32

	// Array field
	if len(m.ArrayDimensions) > 0 {
		lengthInBits += 32 * uint16(len(m.ArrayDimensions))
	}

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Argument) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_Argument) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__argument Argument, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Argument"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Argument")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	dataType, err := ReadSimpleField[NodeId](ctx, "dataType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	valueRank, err := ReadSimpleField(ctx, "valueRank", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueRank' field"))
	}
	m.ValueRank = valueRank

	noOfArrayDimensions, err := ReadSimpleField(ctx, "noOfArrayDimensions", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfArrayDimensions' field"))
	}
	m.NoOfArrayDimensions = noOfArrayDimensions

	arrayDimensions, err := ReadCountArrayField[uint32](ctx, "arrayDimensions", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfArrayDimensions))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensions' field"))
	}
	m.ArrayDimensions = arrayDimensions

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	if closeErr := readBuffer.CloseContext("Argument"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Argument")
	}

	return m, nil
}

func (m *_Argument) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Argument) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Argument"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Argument")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "dataType", m.GetDataType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataType' field")
		}

		if err := WriteSimpleField[int32](ctx, "valueRank", m.GetValueRank(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'valueRank' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfArrayDimensions", m.GetNoOfArrayDimensions(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfArrayDimensions' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "arrayDimensions", m.GetArrayDimensions(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayDimensions' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("Argument"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Argument")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Argument) IsArgument() {}

func (m *_Argument) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Argument) deepCopy() *_Argument {
	if m == nil {
		return nil
	}
	_ArgumentCopy := &_Argument{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Name.DeepCopy().(PascalString),
		m.DataType.DeepCopy().(NodeId),
		m.ValueRank,
		m.NoOfArrayDimensions,
		utils.DeepCopySlice[uint32, uint32](m.ArrayDimensions),
		m.Description.DeepCopy().(LocalizedText),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ArgumentCopy
}

func (m *_Argument) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
