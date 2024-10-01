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

// ParameterValueCustomTypes is the corresponding interface of ParameterValueCustomTypes
type ParameterValueCustomTypes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() CustomTypes
	// IsParameterValueCustomTypes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueCustomTypes()
	// CreateBuilder creates a ParameterValueCustomTypesBuilder
	CreateParameterValueCustomTypesBuilder() ParameterValueCustomTypesBuilder
}

// _ParameterValueCustomTypes is the data-structure of this message
type _ParameterValueCustomTypes struct {
	ParameterValueContract
	Value CustomTypes
}

var _ ParameterValueCustomTypes = (*_ParameterValueCustomTypes)(nil)
var _ ParameterValueRequirements = (*_ParameterValueCustomTypes)(nil)

// NewParameterValueCustomTypes factory function for _ParameterValueCustomTypes
func NewParameterValueCustomTypes(value CustomTypes, numBytes uint8) *_ParameterValueCustomTypes {
	if value == nil {
		panic("value of type CustomTypes for ParameterValueCustomTypes must not be nil")
	}
	_result := &_ParameterValueCustomTypes{
		ParameterValueContract: NewParameterValue(numBytes),
		Value:                  value,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueCustomTypesBuilder is a builder for ParameterValueCustomTypes
type ParameterValueCustomTypesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value CustomTypes) ParameterValueCustomTypesBuilder
	// WithValue adds Value (property field)
	WithValue(CustomTypes) ParameterValueCustomTypesBuilder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(CustomTypesBuilder) CustomTypesBuilder) ParameterValueCustomTypesBuilder
	// Build builds the ParameterValueCustomTypes or returns an error if something is wrong
	Build() (ParameterValueCustomTypes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueCustomTypes
}

// NewParameterValueCustomTypesBuilder() creates a ParameterValueCustomTypesBuilder
func NewParameterValueCustomTypesBuilder() ParameterValueCustomTypesBuilder {
	return &_ParameterValueCustomTypesBuilder{_ParameterValueCustomTypes: new(_ParameterValueCustomTypes)}
}

type _ParameterValueCustomTypesBuilder struct {
	*_ParameterValueCustomTypes

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueCustomTypesBuilder) = (*_ParameterValueCustomTypesBuilder)(nil)

func (b *_ParameterValueCustomTypesBuilder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
}

func (b *_ParameterValueCustomTypesBuilder) WithMandatoryFields(value CustomTypes) ParameterValueCustomTypesBuilder {
	return b.WithValue(value)
}

func (b *_ParameterValueCustomTypesBuilder) WithValue(value CustomTypes) ParameterValueCustomTypesBuilder {
	b.Value = value
	return b
}

func (b *_ParameterValueCustomTypesBuilder) WithValueBuilder(builderSupplier func(CustomTypesBuilder) CustomTypesBuilder) ParameterValueCustomTypesBuilder {
	builder := builderSupplier(b.Value.CreateCustomTypesBuilder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CustomTypesBuilder failed"))
	}
	return b
}

func (b *_ParameterValueCustomTypesBuilder) Build() (ParameterValueCustomTypes, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueCustomTypes.deepCopy(), nil
}

func (b *_ParameterValueCustomTypesBuilder) MustBuild() ParameterValueCustomTypes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ParameterValueCustomTypesBuilder) Done() ParameterValueBuilder {
	return b.parentBuilder
}

func (b *_ParameterValueCustomTypesBuilder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueCustomTypesBuilder) DeepCopy() any {
	_copy := b.CreateParameterValueCustomTypesBuilder().(*_ParameterValueCustomTypesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueCustomTypesBuilder creates a ParameterValueCustomTypesBuilder
func (b *_ParameterValueCustomTypes) CreateParameterValueCustomTypesBuilder() ParameterValueCustomTypesBuilder {
	if b == nil {
		return NewParameterValueCustomTypesBuilder()
	}
	return &_ParameterValueCustomTypesBuilder{_ParameterValueCustomTypes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueCustomTypes) GetParameterType() ParameterType {
	return ParameterType_CUSTOM_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueCustomTypes) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueCustomTypes) GetValue() CustomTypes {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueCustomTypes(structType any) ParameterValueCustomTypes {
	if casted, ok := structType.(ParameterValueCustomTypes); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueCustomTypes); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueCustomTypes) GetTypeName() string {
	return "ParameterValueCustomTypes"
}

func (m *_ParameterValueCustomTypes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ParameterValueCustomTypes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueCustomTypes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueCustomTypes ParameterValueCustomTypes, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueCustomTypes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueCustomTypes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField[CustomTypes](ctx, "value", ReadComplex[CustomTypes](CustomTypesParseWithBufferProducer((uint8)(numBytes)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ParameterValueCustomTypes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueCustomTypes")
	}

	return m, nil
}

func (m *_ParameterValueCustomTypes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueCustomTypes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueCustomTypes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueCustomTypes")
		}

		if err := WriteSimpleField[CustomTypes](ctx, "value", m.GetValue(), WriteComplex[CustomTypes](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueCustomTypes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueCustomTypes")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueCustomTypes) IsParameterValueCustomTypes() {}

func (m *_ParameterValueCustomTypes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueCustomTypes) deepCopy() *_ParameterValueCustomTypes {
	if m == nil {
		return nil
	}
	_ParameterValueCustomTypesCopy := &_ParameterValueCustomTypes{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		m.Value.DeepCopy().(CustomTypes),
	}
	m.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueCustomTypesCopy
}

func (m *_ParameterValueCustomTypes) String() string {
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
