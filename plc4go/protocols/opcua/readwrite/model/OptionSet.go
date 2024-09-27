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

// OptionSet is the corresponding interface of OptionSet
type OptionSet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetValue returns Value (property field)
	GetValue() PascalByteString
	// GetValidBits returns ValidBits (property field)
	GetValidBits() PascalByteString
	// IsOptionSet is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOptionSet()
	// CreateBuilder creates a OptionSetBuilder
	CreateOptionSetBuilder() OptionSetBuilder
}

// _OptionSet is the data-structure of this message
type _OptionSet struct {
	ExtensionObjectDefinitionContract
	Value     PascalByteString
	ValidBits PascalByteString
}

var _ OptionSet = (*_OptionSet)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_OptionSet)(nil)

// NewOptionSet factory function for _OptionSet
func NewOptionSet(value PascalByteString, validBits PascalByteString) *_OptionSet {
	if value == nil {
		panic("value of type PascalByteString for OptionSet must not be nil")
	}
	if validBits == nil {
		panic("validBits of type PascalByteString for OptionSet must not be nil")
	}
	_result := &_OptionSet{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Value:                             value,
		ValidBits:                         validBits,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OptionSetBuilder is a builder for OptionSet
type OptionSetBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value PascalByteString, validBits PascalByteString) OptionSetBuilder
	// WithValue adds Value (property field)
	WithValue(PascalByteString) OptionSetBuilder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) OptionSetBuilder
	// WithValidBits adds ValidBits (property field)
	WithValidBits(PascalByteString) OptionSetBuilder
	// WithValidBitsBuilder adds ValidBits (property field) which is build by the builder
	WithValidBitsBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) OptionSetBuilder
	// Build builds the OptionSet or returns an error if something is wrong
	Build() (OptionSet, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OptionSet
}

// NewOptionSetBuilder() creates a OptionSetBuilder
func NewOptionSetBuilder() OptionSetBuilder {
	return &_OptionSetBuilder{_OptionSet: new(_OptionSet)}
}

type _OptionSetBuilder struct {
	*_OptionSet

	err *utils.MultiError
}

var _ (OptionSetBuilder) = (*_OptionSetBuilder)(nil)

func (m *_OptionSetBuilder) WithMandatoryFields(value PascalByteString, validBits PascalByteString) OptionSetBuilder {
	return m.WithValue(value).WithValidBits(validBits)
}

func (m *_OptionSetBuilder) WithValue(value PascalByteString) OptionSetBuilder {
	m.Value = value
	return m
}

func (m *_OptionSetBuilder) WithValueBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) OptionSetBuilder {
	builder := builderSupplier(m.Value.CreatePascalByteStringBuilder())
	var err error
	m.Value, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return m
}

func (m *_OptionSetBuilder) WithValidBits(validBits PascalByteString) OptionSetBuilder {
	m.ValidBits = validBits
	return m
}

func (m *_OptionSetBuilder) WithValidBitsBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) OptionSetBuilder {
	builder := builderSupplier(m.ValidBits.CreatePascalByteStringBuilder())
	var err error
	m.ValidBits, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return m
}

func (m *_OptionSetBuilder) Build() (OptionSet, error) {
	if m.Value == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if m.ValidBits == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'validBits' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._OptionSet.deepCopy(), nil
}

func (m *_OptionSetBuilder) MustBuild() OptionSet {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_OptionSetBuilder) DeepCopy() any {
	return m.CreateOptionSetBuilder()
}

// CreateOptionSetBuilder creates a OptionSetBuilder
func (m *_OptionSet) CreateOptionSetBuilder() OptionSetBuilder {
	if m == nil {
		return NewOptionSetBuilder()
	}
	return &_OptionSetBuilder{_OptionSet: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OptionSet) GetIdentifier() string {
	return "12757"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OptionSet) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OptionSet) GetValue() PascalByteString {
	return m.Value
}

func (m *_OptionSet) GetValidBits() PascalByteString {
	return m.ValidBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOptionSet(structType any) OptionSet {
	if casted, ok := structType.(OptionSet); ok {
		return casted
	}
	if casted, ok := structType.(*OptionSet); ok {
		return *casted
	}
	return nil
}

func (m *_OptionSet) GetTypeName() string {
	return "OptionSet"
}

func (m *_OptionSet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Simple field (validBits)
	lengthInBits += m.ValidBits.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OptionSet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OptionSet) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__optionSet OptionSet, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OptionSet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OptionSet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField[PascalByteString](ctx, "value", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	validBits, err := ReadSimpleField[PascalByteString](ctx, "validBits", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'validBits' field"))
	}
	m.ValidBits = validBits

	if closeErr := readBuffer.CloseContext("OptionSet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OptionSet")
	}

	return m, nil
}

func (m *_OptionSet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OptionSet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OptionSet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OptionSet")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "value", m.GetValue(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "validBits", m.GetValidBits(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'validBits' field")
		}

		if popErr := writeBuffer.PopContext("OptionSet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OptionSet")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OptionSet) IsOptionSet() {}

func (m *_OptionSet) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OptionSet) deepCopy() *_OptionSet {
	if m == nil {
		return nil
	}
	_OptionSetCopy := &_OptionSet{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Value.DeepCopy().(PascalByteString),
		m.ValidBits.DeepCopy().(PascalByteString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _OptionSetCopy
}

func (m *_OptionSet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
