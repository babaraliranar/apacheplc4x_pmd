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

// VariantQualifiedName is the corresponding interface of VariantQualifiedName
type VariantQualifiedName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []QualifiedName
	// IsVariantQualifiedName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantQualifiedName()
	// CreateBuilder creates a VariantQualifiedNameBuilder
	CreateVariantQualifiedNameBuilder() VariantQualifiedNameBuilder
}

// _VariantQualifiedName is the data-structure of this message
type _VariantQualifiedName struct {
	VariantContract
	ArrayLength *int32
	Value       []QualifiedName
}

var _ VariantQualifiedName = (*_VariantQualifiedName)(nil)
var _ VariantRequirements = (*_VariantQualifiedName)(nil)

// NewVariantQualifiedName factory function for _VariantQualifiedName
func NewVariantQualifiedName(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []QualifiedName) *_VariantQualifiedName {
	_result := &_VariantQualifiedName{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantQualifiedNameBuilder is a builder for VariantQualifiedName
type VariantQualifiedNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []QualifiedName) VariantQualifiedNameBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantQualifiedNameBuilder
	// WithValue adds Value (property field)
	WithValue(...QualifiedName) VariantQualifiedNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantQualifiedName or returns an error if something is wrong
	Build() (VariantQualifiedName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantQualifiedName
}

// NewVariantQualifiedNameBuilder() creates a VariantQualifiedNameBuilder
func NewVariantQualifiedNameBuilder() VariantQualifiedNameBuilder {
	return &_VariantQualifiedNameBuilder{_VariantQualifiedName: new(_VariantQualifiedName)}
}

type _VariantQualifiedNameBuilder struct {
	*_VariantQualifiedName

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantQualifiedNameBuilder) = (*_VariantQualifiedNameBuilder)(nil)

func (b *_VariantQualifiedNameBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
	contract.(*_Variant)._SubType = b._VariantQualifiedName
}

func (b *_VariantQualifiedNameBuilder) WithMandatoryFields(value []QualifiedName) VariantQualifiedNameBuilder {
	return b.WithValue(value...)
}

func (b *_VariantQualifiedNameBuilder) WithOptionalArrayLength(arrayLength int32) VariantQualifiedNameBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantQualifiedNameBuilder) WithValue(value ...QualifiedName) VariantQualifiedNameBuilder {
	b.Value = value
	return b
}

func (b *_VariantQualifiedNameBuilder) Build() (VariantQualifiedName, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantQualifiedName.deepCopy(), nil
}

func (b *_VariantQualifiedNameBuilder) MustBuild() VariantQualifiedName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantQualifiedNameBuilder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantQualifiedNameBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantQualifiedNameBuilder) DeepCopy() any {
	_copy := b.CreateVariantQualifiedNameBuilder().(*_VariantQualifiedNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantQualifiedNameBuilder creates a VariantQualifiedNameBuilder
func (b *_VariantQualifiedName) CreateVariantQualifiedNameBuilder() VariantQualifiedNameBuilder {
	if b == nil {
		return NewVariantQualifiedNameBuilder()
	}
	return &_VariantQualifiedNameBuilder{_VariantQualifiedName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantQualifiedName) GetVariantType() uint8 {
	return uint8(20)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantQualifiedName) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantQualifiedName) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantQualifiedName) GetValue() []QualifiedName {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantQualifiedName(structType any) VariantQualifiedName {
	if casted, ok := structType.(VariantQualifiedName); ok {
		return casted
	}
	if casted, ok := structType.(*VariantQualifiedName); ok {
		return *casted
	}
	return nil
}

func (m *_VariantQualifiedName) GetTypeName() string {
	return "VariantQualifiedName"
}

func (m *_VariantQualifiedName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_VariantQualifiedName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantQualifiedName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantQualifiedName VariantQualifiedName, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantQualifiedName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantQualifiedName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[QualifiedName](ctx, "value", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantQualifiedName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantQualifiedName")
	}

	return m, nil
}

func (m *_VariantQualifiedName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantQualifiedName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantQualifiedName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantQualifiedName")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "value", m.GetValue(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantQualifiedName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantQualifiedName")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantQualifiedName) IsVariantQualifiedName() {}

func (m *_VariantQualifiedName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantQualifiedName) deepCopy() *_VariantQualifiedName {
	if m == nil {
		return nil
	}
	_VariantQualifiedNameCopy := &_VariantQualifiedName{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[QualifiedName, QualifiedName](m.Value),
	}
	_VariantQualifiedNameCopy.VariantContract.(*_Variant)._SubType = m
	return _VariantQualifiedNameCopy
}

func (m *_VariantQualifiedName) String() string {
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
