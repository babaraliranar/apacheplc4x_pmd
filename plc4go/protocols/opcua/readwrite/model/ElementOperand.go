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

// ElementOperand is the corresponding interface of ElementOperand
type ElementOperand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetIndex returns Index (property field)
	GetIndex() uint32
	// IsElementOperand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsElementOperand()
	// CreateBuilder creates a ElementOperandBuilder
	CreateElementOperandBuilder() ElementOperandBuilder
}

// _ElementOperand is the data-structure of this message
type _ElementOperand struct {
	ExtensionObjectDefinitionContract
	Index uint32
}

var _ ElementOperand = (*_ElementOperand)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ElementOperand)(nil)

// NewElementOperand factory function for _ElementOperand
func NewElementOperand(index uint32) *_ElementOperand {
	_result := &_ElementOperand{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Index:                             index,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ElementOperandBuilder is a builder for ElementOperand
type ElementOperandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(index uint32) ElementOperandBuilder
	// WithIndex adds Index (property field)
	WithIndex(uint32) ElementOperandBuilder
	// Build builds the ElementOperand or returns an error if something is wrong
	Build() (ElementOperand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ElementOperand
}

// NewElementOperandBuilder() creates a ElementOperandBuilder
func NewElementOperandBuilder() ElementOperandBuilder {
	return &_ElementOperandBuilder{_ElementOperand: new(_ElementOperand)}
}

type _ElementOperandBuilder struct {
	*_ElementOperand

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ElementOperandBuilder) = (*_ElementOperandBuilder)(nil)

func (b *_ElementOperandBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ElementOperandBuilder) WithMandatoryFields(index uint32) ElementOperandBuilder {
	return b.WithIndex(index)
}

func (b *_ElementOperandBuilder) WithIndex(index uint32) ElementOperandBuilder {
	b.Index = index
	return b
}

func (b *_ElementOperandBuilder) Build() (ElementOperand, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ElementOperand.deepCopy(), nil
}

func (b *_ElementOperandBuilder) MustBuild() ElementOperand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ElementOperandBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ElementOperandBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ElementOperandBuilder) DeepCopy() any {
	_copy := b.CreateElementOperandBuilder().(*_ElementOperandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateElementOperandBuilder creates a ElementOperandBuilder
func (b *_ElementOperand) CreateElementOperandBuilder() ElementOperandBuilder {
	if b == nil {
		return NewElementOperandBuilder()
	}
	return &_ElementOperandBuilder{_ElementOperand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ElementOperand) GetExtensionId() int32 {
	return int32(594)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ElementOperand) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ElementOperand) GetIndex() uint32 {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastElementOperand(structType any) ElementOperand {
	if casted, ok := structType.(ElementOperand); ok {
		return casted
	}
	if casted, ok := structType.(*ElementOperand); ok {
		return *casted
	}
	return nil
}

func (m *_ElementOperand) GetTypeName() string {
	return "ElementOperand"
}

func (m *_ElementOperand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (index)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ElementOperand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ElementOperand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__elementOperand ElementOperand, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ElementOperand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ElementOperand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	index, err := ReadSimpleField(ctx, "index", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}
	m.Index = index

	if closeErr := readBuffer.CloseContext("ElementOperand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ElementOperand")
	}

	return m, nil
}

func (m *_ElementOperand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ElementOperand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ElementOperand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ElementOperand")
		}

		if err := WriteSimpleField[uint32](ctx, "index", m.GetIndex(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'index' field")
		}

		if popErr := writeBuffer.PopContext("ElementOperand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ElementOperand")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ElementOperand) IsElementOperand() {}

func (m *_ElementOperand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ElementOperand) deepCopy() *_ElementOperand {
	if m == nil {
		return nil
	}
	_ElementOperandCopy := &_ElementOperand{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Index,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ElementOperandCopy
}

func (m *_ElementOperand) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
