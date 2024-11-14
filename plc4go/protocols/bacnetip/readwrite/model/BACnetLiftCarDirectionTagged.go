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

// BACnetLiftCarDirectionTagged is the corresponding interface of BACnetLiftCarDirectionTagged
type BACnetLiftCarDirectionTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLiftCarDirection
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetLiftCarDirectionTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLiftCarDirectionTagged()
	// CreateBuilder creates a BACnetLiftCarDirectionTaggedBuilder
	CreateBACnetLiftCarDirectionTaggedBuilder() BACnetLiftCarDirectionTaggedBuilder
}

// _BACnetLiftCarDirectionTagged is the data-structure of this message
type _BACnetLiftCarDirectionTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLiftCarDirection
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLiftCarDirectionTagged = (*_BACnetLiftCarDirectionTagged)(nil)

// NewBACnetLiftCarDirectionTagged factory function for _BACnetLiftCarDirectionTagged
func NewBACnetLiftCarDirectionTagged(header BACnetTagHeader, value BACnetLiftCarDirection, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLiftCarDirectionTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLiftCarDirectionTagged must not be nil")
	}
	return &_BACnetLiftCarDirectionTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLiftCarDirectionTaggedBuilder is a builder for BACnetLiftCarDirectionTagged
type BACnetLiftCarDirectionTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetLiftCarDirection, proprietaryValue uint32) BACnetLiftCarDirectionTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetLiftCarDirectionTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLiftCarDirectionTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetLiftCarDirection) BACnetLiftCarDirectionTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetLiftCarDirectionTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetLiftCarDirectionTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetLiftCarDirectionTaggedBuilder
	// Build builds the BACnetLiftCarDirectionTagged or returns an error if something is wrong
	Build() (BACnetLiftCarDirectionTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLiftCarDirectionTagged
}

// NewBACnetLiftCarDirectionTaggedBuilder() creates a BACnetLiftCarDirectionTaggedBuilder
func NewBACnetLiftCarDirectionTaggedBuilder() BACnetLiftCarDirectionTaggedBuilder {
	return &_BACnetLiftCarDirectionTaggedBuilder{_BACnetLiftCarDirectionTagged: new(_BACnetLiftCarDirectionTagged)}
}

type _BACnetLiftCarDirectionTaggedBuilder struct {
	*_BACnetLiftCarDirectionTagged

	err *utils.MultiError
}

var _ (BACnetLiftCarDirectionTaggedBuilder) = (*_BACnetLiftCarDirectionTaggedBuilder)(nil)

func (b *_BACnetLiftCarDirectionTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetLiftCarDirection, proprietaryValue uint32) BACnetLiftCarDirectionTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetLiftCarDirectionTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLiftCarDirectionTaggedBuilder {
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

func (b *_BACnetLiftCarDirectionTaggedBuilder) WithValue(value BACnetLiftCarDirection) BACnetLiftCarDirectionTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetLiftCarDirectionTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetLiftCarDirectionTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetLiftCarDirectionTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetLiftCarDirectionTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) Build() (BACnetLiftCarDirectionTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLiftCarDirectionTagged.deepCopy(), nil
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) MustBuild() BACnetLiftCarDirectionTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLiftCarDirectionTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLiftCarDirectionTaggedBuilder().(*_BACnetLiftCarDirectionTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLiftCarDirectionTaggedBuilder creates a BACnetLiftCarDirectionTaggedBuilder
func (b *_BACnetLiftCarDirectionTagged) CreateBACnetLiftCarDirectionTaggedBuilder() BACnetLiftCarDirectionTaggedBuilder {
	if b == nil {
		return NewBACnetLiftCarDirectionTaggedBuilder()
	}
	return &_BACnetLiftCarDirectionTaggedBuilder{_BACnetLiftCarDirectionTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarDirectionTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLiftCarDirectionTagged) GetValue() BACnetLiftCarDirection {
	return m.Value
}

func (m *_BACnetLiftCarDirectionTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetLiftCarDirectionTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarDirectionTagged(structType any) BACnetLiftCarDirectionTagged {
	if casted, ok := structType.(BACnetLiftCarDirectionTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarDirectionTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarDirectionTagged) GetTypeName() string {
	return "BACnetLiftCarDirectionTagged"
}

func (m *_BACnetLiftCarDirectionTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetLiftCarDirectionTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarDirectionTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLiftCarDirectionTagged, error) {
	return BACnetLiftCarDirectionTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLiftCarDirectionTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarDirectionTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarDirectionTagged, error) {
		return BACnetLiftCarDirectionTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLiftCarDirectionTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLiftCarDirectionTagged, error) {
	v, err := (&_BACnetLiftCarDirectionTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLiftCarDirectionTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLiftCarDirectionTagged BACnetLiftCarDirectionTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftCarDirectionTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarDirectionTagged")
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

	value, err := ReadManualField[BACnetLiftCarDirection](ctx, "value", readBuffer, EnsureType[BACnetLiftCarDirection](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetLiftCarDirectionTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarDirectionTagged")
	}

	return m, nil
}

func (m *_BACnetLiftCarDirectionTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarDirectionTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLiftCarDirectionTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarDirectionTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLiftCarDirection](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetLiftCarDirectionTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarDirectionTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLiftCarDirectionTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLiftCarDirectionTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLiftCarDirectionTagged) IsBACnetLiftCarDirectionTagged() {}

func (m *_BACnetLiftCarDirectionTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLiftCarDirectionTagged) deepCopy() *_BACnetLiftCarDirectionTagged {
	if m == nil {
		return nil
	}
	_BACnetLiftCarDirectionTaggedCopy := &_BACnetLiftCarDirectionTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetLiftCarDirectionTaggedCopy
}

func (m *_BACnetLiftCarDirectionTagged) String() string {
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
