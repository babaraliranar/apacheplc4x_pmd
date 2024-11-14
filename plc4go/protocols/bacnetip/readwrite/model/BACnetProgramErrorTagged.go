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

// BACnetProgramErrorTagged is the corresponding interface of BACnetProgramErrorTagged
type BACnetProgramErrorTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProgramError
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetProgramErrorTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProgramErrorTagged()
	// CreateBuilder creates a BACnetProgramErrorTaggedBuilder
	CreateBACnetProgramErrorTaggedBuilder() BACnetProgramErrorTaggedBuilder
}

// _BACnetProgramErrorTagged is the data-structure of this message
type _BACnetProgramErrorTagged struct {
	Header           BACnetTagHeader
	Value            BACnetProgramError
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetProgramErrorTagged = (*_BACnetProgramErrorTagged)(nil)

// NewBACnetProgramErrorTagged factory function for _BACnetProgramErrorTagged
func NewBACnetProgramErrorTagged(header BACnetTagHeader, value BACnetProgramError, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetProgramErrorTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetProgramErrorTagged must not be nil")
	}
	return &_BACnetProgramErrorTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetProgramErrorTaggedBuilder is a builder for BACnetProgramErrorTagged
type BACnetProgramErrorTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetProgramError, proprietaryValue uint32) BACnetProgramErrorTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetProgramErrorTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProgramErrorTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetProgramError) BACnetProgramErrorTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetProgramErrorTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetProgramErrorTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetProgramErrorTaggedBuilder
	// Build builds the BACnetProgramErrorTagged or returns an error if something is wrong
	Build() (BACnetProgramErrorTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetProgramErrorTagged
}

// NewBACnetProgramErrorTaggedBuilder() creates a BACnetProgramErrorTaggedBuilder
func NewBACnetProgramErrorTaggedBuilder() BACnetProgramErrorTaggedBuilder {
	return &_BACnetProgramErrorTaggedBuilder{_BACnetProgramErrorTagged: new(_BACnetProgramErrorTagged)}
}

type _BACnetProgramErrorTaggedBuilder struct {
	*_BACnetProgramErrorTagged

	err *utils.MultiError
}

var _ (BACnetProgramErrorTaggedBuilder) = (*_BACnetProgramErrorTaggedBuilder)(nil)

func (b *_BACnetProgramErrorTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetProgramError, proprietaryValue uint32) BACnetProgramErrorTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetProgramErrorTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetProgramErrorTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetProgramErrorTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProgramErrorTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetProgramErrorTaggedBuilder) WithValue(value BACnetProgramError) BACnetProgramErrorTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetProgramErrorTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetProgramErrorTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetProgramErrorTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetProgramErrorTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetProgramErrorTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetProgramErrorTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetProgramErrorTaggedBuilder) Build() (BACnetProgramErrorTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetProgramErrorTagged.deepCopy(), nil
}

func (b *_BACnetProgramErrorTaggedBuilder) MustBuild() BACnetProgramErrorTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetProgramErrorTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetProgramErrorTaggedBuilder().(*_BACnetProgramErrorTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetProgramErrorTaggedBuilder creates a BACnetProgramErrorTaggedBuilder
func (b *_BACnetProgramErrorTagged) CreateBACnetProgramErrorTaggedBuilder() BACnetProgramErrorTaggedBuilder {
	if b == nil {
		return NewBACnetProgramErrorTaggedBuilder()
	}
	return &_BACnetProgramErrorTaggedBuilder{_BACnetProgramErrorTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProgramErrorTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetProgramErrorTagged) GetValue() BACnetProgramError {
	return m.Value
}

func (m *_BACnetProgramErrorTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetProgramErrorTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetProgramError_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetProgramErrorTagged(structType any) BACnetProgramErrorTagged {
	if casted, ok := structType.(BACnetProgramErrorTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProgramErrorTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProgramErrorTagged) GetTypeName() string {
	return "BACnetProgramErrorTagged"
}

func (m *_BACnetProgramErrorTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetProgramErrorTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProgramErrorTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetProgramErrorTagged, error) {
	return BACnetProgramErrorTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetProgramErrorTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramErrorTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramErrorTagged, error) {
		return BACnetProgramErrorTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetProgramErrorTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetProgramErrorTagged, error) {
	v, err := (&_BACnetProgramErrorTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetProgramErrorTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetProgramErrorTagged BACnetProgramErrorTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProgramErrorTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProgramErrorTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetProgramError](ctx, "value", readBuffer, EnsureType[BACnetProgramError](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetProgramError_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetProgramError_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetProgramErrorTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProgramErrorTagged")
	}

	return m, nil
}

func (m *_BACnetProgramErrorTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProgramErrorTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProgramErrorTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProgramErrorTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetProgramError](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProgramErrorTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProgramErrorTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetProgramErrorTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetProgramErrorTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetProgramErrorTagged) IsBACnetProgramErrorTagged() {}

func (m *_BACnetProgramErrorTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetProgramErrorTagged) deepCopy() *_BACnetProgramErrorTagged {
	if m == nil {
		return nil
	}
	_BACnetProgramErrorTaggedCopy := &_BACnetProgramErrorTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetProgramErrorTaggedCopy
}

func (m *_BACnetProgramErrorTagged) String() string {
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
