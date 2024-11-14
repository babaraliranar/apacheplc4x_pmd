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

// FilterOperand is the corresponding interface of FilterOperand
type FilterOperand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsFilterOperand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFilterOperand()
	// CreateBuilder creates a FilterOperandBuilder
	CreateFilterOperandBuilder() FilterOperandBuilder
}

// _FilterOperand is the data-structure of this message
type _FilterOperand struct {
	ExtensionObjectDefinitionContract
}

var _ FilterOperand = (*_FilterOperand)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FilterOperand)(nil)

// NewFilterOperand factory function for _FilterOperand
func NewFilterOperand() *_FilterOperand {
	_result := &_FilterOperand{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FilterOperandBuilder is a builder for FilterOperand
type FilterOperandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() FilterOperandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the FilterOperand or returns an error if something is wrong
	Build() (FilterOperand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FilterOperand
}

// NewFilterOperandBuilder() creates a FilterOperandBuilder
func NewFilterOperandBuilder() FilterOperandBuilder {
	return &_FilterOperandBuilder{_FilterOperand: new(_FilterOperand)}
}

type _FilterOperandBuilder struct {
	*_FilterOperand

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FilterOperandBuilder) = (*_FilterOperandBuilder)(nil)

func (b *_FilterOperandBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._FilterOperand
}

func (b *_FilterOperandBuilder) WithMandatoryFields() FilterOperandBuilder {
	return b
}

func (b *_FilterOperandBuilder) Build() (FilterOperand, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FilterOperand.deepCopy(), nil
}

func (b *_FilterOperandBuilder) MustBuild() FilterOperand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FilterOperandBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_FilterOperandBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FilterOperandBuilder) DeepCopy() any {
	_copy := b.CreateFilterOperandBuilder().(*_FilterOperandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFilterOperandBuilder creates a FilterOperandBuilder
func (b *_FilterOperand) CreateFilterOperandBuilder() FilterOperandBuilder {
	if b == nil {
		return NewFilterOperandBuilder()
	}
	return &_FilterOperandBuilder{_FilterOperand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FilterOperand) GetExtensionId() int32 {
	return int32(591)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FilterOperand) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastFilterOperand(structType any) FilterOperand {
	if casted, ok := structType.(FilterOperand); ok {
		return casted
	}
	if casted, ok := structType.(*FilterOperand); ok {
		return *casted
	}
	return nil
}

func (m *_FilterOperand) GetTypeName() string {
	return "FilterOperand"
}

func (m *_FilterOperand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_FilterOperand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FilterOperand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__filterOperand FilterOperand, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FilterOperand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FilterOperand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("FilterOperand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FilterOperand")
	}

	return m, nil
}

func (m *_FilterOperand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FilterOperand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FilterOperand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FilterOperand")
		}

		if popErr := writeBuffer.PopContext("FilterOperand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FilterOperand")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FilterOperand) IsFilterOperand() {}

func (m *_FilterOperand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FilterOperand) deepCopy() *_FilterOperand {
	if m == nil {
		return nil
	}
	_FilterOperandCopy := &_FilterOperand{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	_FilterOperandCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FilterOperandCopy
}

func (m *_FilterOperand) String() string {
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
