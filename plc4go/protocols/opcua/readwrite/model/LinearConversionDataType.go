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

// LinearConversionDataType is the corresponding interface of LinearConversionDataType
type LinearConversionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetInitialAddend returns InitialAddend (property field)
	GetInitialAddend() float32
	// GetMultiplicand returns Multiplicand (property field)
	GetMultiplicand() float32
	// GetDivisor returns Divisor (property field)
	GetDivisor() float32
	// GetFinalAddend returns FinalAddend (property field)
	GetFinalAddend() float32
	// IsLinearConversionDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLinearConversionDataType()
	// CreateBuilder creates a LinearConversionDataTypeBuilder
	CreateLinearConversionDataTypeBuilder() LinearConversionDataTypeBuilder
}

// _LinearConversionDataType is the data-structure of this message
type _LinearConversionDataType struct {
	ExtensionObjectDefinitionContract
	InitialAddend float32
	Multiplicand  float32
	Divisor       float32
	FinalAddend   float32
}

var _ LinearConversionDataType = (*_LinearConversionDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_LinearConversionDataType)(nil)

// NewLinearConversionDataType factory function for _LinearConversionDataType
func NewLinearConversionDataType(initialAddend float32, multiplicand float32, divisor float32, finalAddend float32) *_LinearConversionDataType {
	_result := &_LinearConversionDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		InitialAddend:                     initialAddend,
		Multiplicand:                      multiplicand,
		Divisor:                           divisor,
		FinalAddend:                       finalAddend,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LinearConversionDataTypeBuilder is a builder for LinearConversionDataType
type LinearConversionDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(initialAddend float32, multiplicand float32, divisor float32, finalAddend float32) LinearConversionDataTypeBuilder
	// WithInitialAddend adds InitialAddend (property field)
	WithInitialAddend(float32) LinearConversionDataTypeBuilder
	// WithMultiplicand adds Multiplicand (property field)
	WithMultiplicand(float32) LinearConversionDataTypeBuilder
	// WithDivisor adds Divisor (property field)
	WithDivisor(float32) LinearConversionDataTypeBuilder
	// WithFinalAddend adds FinalAddend (property field)
	WithFinalAddend(float32) LinearConversionDataTypeBuilder
	// Build builds the LinearConversionDataType or returns an error if something is wrong
	Build() (LinearConversionDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LinearConversionDataType
}

// NewLinearConversionDataTypeBuilder() creates a LinearConversionDataTypeBuilder
func NewLinearConversionDataTypeBuilder() LinearConversionDataTypeBuilder {
	return &_LinearConversionDataTypeBuilder{_LinearConversionDataType: new(_LinearConversionDataType)}
}

type _LinearConversionDataTypeBuilder struct {
	*_LinearConversionDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (LinearConversionDataTypeBuilder) = (*_LinearConversionDataTypeBuilder)(nil)

func (b *_LinearConversionDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_LinearConversionDataTypeBuilder) WithMandatoryFields(initialAddend float32, multiplicand float32, divisor float32, finalAddend float32) LinearConversionDataTypeBuilder {
	return b.WithInitialAddend(initialAddend).WithMultiplicand(multiplicand).WithDivisor(divisor).WithFinalAddend(finalAddend)
}

func (b *_LinearConversionDataTypeBuilder) WithInitialAddend(initialAddend float32) LinearConversionDataTypeBuilder {
	b.InitialAddend = initialAddend
	return b
}

func (b *_LinearConversionDataTypeBuilder) WithMultiplicand(multiplicand float32) LinearConversionDataTypeBuilder {
	b.Multiplicand = multiplicand
	return b
}

func (b *_LinearConversionDataTypeBuilder) WithDivisor(divisor float32) LinearConversionDataTypeBuilder {
	b.Divisor = divisor
	return b
}

func (b *_LinearConversionDataTypeBuilder) WithFinalAddend(finalAddend float32) LinearConversionDataTypeBuilder {
	b.FinalAddend = finalAddend
	return b
}

func (b *_LinearConversionDataTypeBuilder) Build() (LinearConversionDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LinearConversionDataType.deepCopy(), nil
}

func (b *_LinearConversionDataTypeBuilder) MustBuild() LinearConversionDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_LinearConversionDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_LinearConversionDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_LinearConversionDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateLinearConversionDataTypeBuilder().(*_LinearConversionDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLinearConversionDataTypeBuilder creates a LinearConversionDataTypeBuilder
func (b *_LinearConversionDataType) CreateLinearConversionDataTypeBuilder() LinearConversionDataTypeBuilder {
	if b == nil {
		return NewLinearConversionDataTypeBuilder()
	}
	return &_LinearConversionDataTypeBuilder{_LinearConversionDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LinearConversionDataType) GetExtensionId() int32 {
	return int32(32437)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LinearConversionDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LinearConversionDataType) GetInitialAddend() float32 {
	return m.InitialAddend
}

func (m *_LinearConversionDataType) GetMultiplicand() float32 {
	return m.Multiplicand
}

func (m *_LinearConversionDataType) GetDivisor() float32 {
	return m.Divisor
}

func (m *_LinearConversionDataType) GetFinalAddend() float32 {
	return m.FinalAddend
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLinearConversionDataType(structType any) LinearConversionDataType {
	if casted, ok := structType.(LinearConversionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*LinearConversionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_LinearConversionDataType) GetTypeName() string {
	return "LinearConversionDataType"
}

func (m *_LinearConversionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (initialAddend)
	lengthInBits += 32

	// Simple field (multiplicand)
	lengthInBits += 32

	// Simple field (divisor)
	lengthInBits += 32

	// Simple field (finalAddend)
	lengthInBits += 32

	return lengthInBits
}

func (m *_LinearConversionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LinearConversionDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__linearConversionDataType LinearConversionDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LinearConversionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LinearConversionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	initialAddend, err := ReadSimpleField(ctx, "initialAddend", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initialAddend' field"))
	}
	m.InitialAddend = initialAddend

	multiplicand, err := ReadSimpleField(ctx, "multiplicand", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'multiplicand' field"))
	}
	m.Multiplicand = multiplicand

	divisor, err := ReadSimpleField(ctx, "divisor", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'divisor' field"))
	}
	m.Divisor = divisor

	finalAddend, err := ReadSimpleField(ctx, "finalAddend", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'finalAddend' field"))
	}
	m.FinalAddend = finalAddend

	if closeErr := readBuffer.CloseContext("LinearConversionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LinearConversionDataType")
	}

	return m, nil
}

func (m *_LinearConversionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LinearConversionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LinearConversionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LinearConversionDataType")
		}

		if err := WriteSimpleField[float32](ctx, "initialAddend", m.GetInitialAddend(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'initialAddend' field")
		}

		if err := WriteSimpleField[float32](ctx, "multiplicand", m.GetMultiplicand(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'multiplicand' field")
		}

		if err := WriteSimpleField[float32](ctx, "divisor", m.GetDivisor(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'divisor' field")
		}

		if err := WriteSimpleField[float32](ctx, "finalAddend", m.GetFinalAddend(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'finalAddend' field")
		}

		if popErr := writeBuffer.PopContext("LinearConversionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LinearConversionDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LinearConversionDataType) IsLinearConversionDataType() {}

func (m *_LinearConversionDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LinearConversionDataType) deepCopy() *_LinearConversionDataType {
	if m == nil {
		return nil
	}
	_LinearConversionDataTypeCopy := &_LinearConversionDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.InitialAddend,
		m.Multiplicand,
		m.Divisor,
		m.FinalAddend,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _LinearConversionDataTypeCopy
}

func (m *_LinearConversionDataType) String() string {
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
