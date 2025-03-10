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

// DoubleComplexNumberType is the corresponding interface of DoubleComplexNumberType
type DoubleComplexNumberType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetReal returns Real (property field)
	GetReal() float64
	// GetImaginary returns Imaginary (property field)
	GetImaginary() float64
	// IsDoubleComplexNumberType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDoubleComplexNumberType()
	// CreateBuilder creates a DoubleComplexNumberTypeBuilder
	CreateDoubleComplexNumberTypeBuilder() DoubleComplexNumberTypeBuilder
}

// _DoubleComplexNumberType is the data-structure of this message
type _DoubleComplexNumberType struct {
	ExtensionObjectDefinitionContract
	Real      float64
	Imaginary float64
}

var _ DoubleComplexNumberType = (*_DoubleComplexNumberType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DoubleComplexNumberType)(nil)

// NewDoubleComplexNumberType factory function for _DoubleComplexNumberType
func NewDoubleComplexNumberType(real float64, imaginary float64) *_DoubleComplexNumberType {
	_result := &_DoubleComplexNumberType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Real:                              real,
		Imaginary:                         imaginary,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DoubleComplexNumberTypeBuilder is a builder for DoubleComplexNumberType
type DoubleComplexNumberTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(real float64, imaginary float64) DoubleComplexNumberTypeBuilder
	// WithReal adds Real (property field)
	WithReal(float64) DoubleComplexNumberTypeBuilder
	// WithImaginary adds Imaginary (property field)
	WithImaginary(float64) DoubleComplexNumberTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DoubleComplexNumberType or returns an error if something is wrong
	Build() (DoubleComplexNumberType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DoubleComplexNumberType
}

// NewDoubleComplexNumberTypeBuilder() creates a DoubleComplexNumberTypeBuilder
func NewDoubleComplexNumberTypeBuilder() DoubleComplexNumberTypeBuilder {
	return &_DoubleComplexNumberTypeBuilder{_DoubleComplexNumberType: new(_DoubleComplexNumberType)}
}

type _DoubleComplexNumberTypeBuilder struct {
	*_DoubleComplexNumberType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DoubleComplexNumberTypeBuilder) = (*_DoubleComplexNumberTypeBuilder)(nil)

func (b *_DoubleComplexNumberTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._DoubleComplexNumberType
}

func (b *_DoubleComplexNumberTypeBuilder) WithMandatoryFields(real float64, imaginary float64) DoubleComplexNumberTypeBuilder {
	return b.WithReal(real).WithImaginary(imaginary)
}

func (b *_DoubleComplexNumberTypeBuilder) WithReal(real float64) DoubleComplexNumberTypeBuilder {
	b.Real = real
	return b
}

func (b *_DoubleComplexNumberTypeBuilder) WithImaginary(imaginary float64) DoubleComplexNumberTypeBuilder {
	b.Imaginary = imaginary
	return b
}

func (b *_DoubleComplexNumberTypeBuilder) Build() (DoubleComplexNumberType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DoubleComplexNumberType.deepCopy(), nil
}

func (b *_DoubleComplexNumberTypeBuilder) MustBuild() DoubleComplexNumberType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DoubleComplexNumberTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DoubleComplexNumberTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DoubleComplexNumberTypeBuilder) DeepCopy() any {
	_copy := b.CreateDoubleComplexNumberTypeBuilder().(*_DoubleComplexNumberTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDoubleComplexNumberTypeBuilder creates a DoubleComplexNumberTypeBuilder
func (b *_DoubleComplexNumberType) CreateDoubleComplexNumberTypeBuilder() DoubleComplexNumberTypeBuilder {
	if b == nil {
		return NewDoubleComplexNumberTypeBuilder()
	}
	return &_DoubleComplexNumberTypeBuilder{_DoubleComplexNumberType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DoubleComplexNumberType) GetExtensionId() int32 {
	return int32(12174)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DoubleComplexNumberType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DoubleComplexNumberType) GetReal() float64 {
	return m.Real
}

func (m *_DoubleComplexNumberType) GetImaginary() float64 {
	return m.Imaginary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDoubleComplexNumberType(structType any) DoubleComplexNumberType {
	if casted, ok := structType.(DoubleComplexNumberType); ok {
		return casted
	}
	if casted, ok := structType.(*DoubleComplexNumberType); ok {
		return *casted
	}
	return nil
}

func (m *_DoubleComplexNumberType) GetTypeName() string {
	return "DoubleComplexNumberType"
}

func (m *_DoubleComplexNumberType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (real)
	lengthInBits += 64

	// Simple field (imaginary)
	lengthInBits += 64

	return lengthInBits
}

func (m *_DoubleComplexNumberType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DoubleComplexNumberType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__doubleComplexNumberType DoubleComplexNumberType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DoubleComplexNumberType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DoubleComplexNumberType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	real, err := ReadSimpleField(ctx, "real", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'real' field"))
	}
	m.Real = real

	imaginary, err := ReadSimpleField(ctx, "imaginary", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'imaginary' field"))
	}
	m.Imaginary = imaginary

	if closeErr := readBuffer.CloseContext("DoubleComplexNumberType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DoubleComplexNumberType")
	}

	return m, nil
}

func (m *_DoubleComplexNumberType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DoubleComplexNumberType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DoubleComplexNumberType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DoubleComplexNumberType")
		}

		if err := WriteSimpleField[float64](ctx, "real", m.GetReal(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'real' field")
		}

		if err := WriteSimpleField[float64](ctx, "imaginary", m.GetImaginary(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'imaginary' field")
		}

		if popErr := writeBuffer.PopContext("DoubleComplexNumberType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DoubleComplexNumberType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DoubleComplexNumberType) IsDoubleComplexNumberType() {}

func (m *_DoubleComplexNumberType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DoubleComplexNumberType) deepCopy() *_DoubleComplexNumberType {
	if m == nil {
		return nil
	}
	_DoubleComplexNumberTypeCopy := &_DoubleComplexNumberType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Real,
		m.Imaginary,
	}
	_DoubleComplexNumberTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DoubleComplexNumberTypeCopy
}

func (m *_DoubleComplexNumberType) String() string {
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
