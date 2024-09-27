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

// BACnetRestartReasonTagged is the corresponding interface of BACnetRestartReasonTagged
type BACnetRestartReasonTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetRestartReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetRestartReasonTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRestartReasonTagged()
	// CreateBuilder creates a BACnetRestartReasonTaggedBuilder
	CreateBACnetRestartReasonTaggedBuilder() BACnetRestartReasonTaggedBuilder
}

// _BACnetRestartReasonTagged is the data-structure of this message
type _BACnetRestartReasonTagged struct {
	Header           BACnetTagHeader
	Value            BACnetRestartReason
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetRestartReasonTagged = (*_BACnetRestartReasonTagged)(nil)

// NewBACnetRestartReasonTagged factory function for _BACnetRestartReasonTagged
func NewBACnetRestartReasonTagged(header BACnetTagHeader, value BACnetRestartReason, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetRestartReasonTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetRestartReasonTagged must not be nil")
	}
	return &_BACnetRestartReasonTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRestartReasonTaggedBuilder is a builder for BACnetRestartReasonTagged
type BACnetRestartReasonTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetRestartReason, proprietaryValue uint32) BACnetRestartReasonTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetRestartReasonTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetRestartReasonTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetRestartReason) BACnetRestartReasonTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetRestartReasonTaggedBuilder
	// Build builds the BACnetRestartReasonTagged or returns an error if something is wrong
	Build() (BACnetRestartReasonTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRestartReasonTagged
}

// NewBACnetRestartReasonTaggedBuilder() creates a BACnetRestartReasonTaggedBuilder
func NewBACnetRestartReasonTaggedBuilder() BACnetRestartReasonTaggedBuilder {
	return &_BACnetRestartReasonTaggedBuilder{_BACnetRestartReasonTagged: new(_BACnetRestartReasonTagged)}
}

type _BACnetRestartReasonTaggedBuilder struct {
	*_BACnetRestartReasonTagged

	err *utils.MultiError
}

var _ (BACnetRestartReasonTaggedBuilder) = (*_BACnetRestartReasonTaggedBuilder)(nil)

func (m *_BACnetRestartReasonTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetRestartReason, proprietaryValue uint32) BACnetRestartReasonTaggedBuilder {
	return m.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (m *_BACnetRestartReasonTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetRestartReasonTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetRestartReasonTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetRestartReasonTaggedBuilder {
	builder := builderSupplier(m.Header.CreateBACnetTagHeaderBuilder())
	var err error
	m.Header, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetRestartReasonTaggedBuilder) WithValue(value BACnetRestartReason) BACnetRestartReasonTaggedBuilder {
	m.Value = value
	return m
}

func (m *_BACnetRestartReasonTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetRestartReasonTaggedBuilder {
	m.ProprietaryValue = proprietaryValue
	return m
}

func (m *_BACnetRestartReasonTaggedBuilder) Build() (BACnetRestartReasonTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetRestartReasonTagged.deepCopy(), nil
}

func (m *_BACnetRestartReasonTaggedBuilder) MustBuild() BACnetRestartReasonTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetRestartReasonTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetRestartReasonTaggedBuilder()
}

// CreateBACnetRestartReasonTaggedBuilder creates a BACnetRestartReasonTaggedBuilder
func (m *_BACnetRestartReasonTagged) CreateBACnetRestartReasonTaggedBuilder() BACnetRestartReasonTaggedBuilder {
	if m == nil {
		return NewBACnetRestartReasonTaggedBuilder()
	}
	return &_BACnetRestartReasonTaggedBuilder{_BACnetRestartReasonTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRestartReasonTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetRestartReasonTagged) GetValue() BACnetRestartReason {
	return m.Value
}

func (m *_BACnetRestartReasonTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetRestartReasonTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetRestartReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetRestartReasonTagged(structType any) BACnetRestartReasonTagged {
	if casted, ok := structType.(BACnetRestartReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRestartReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRestartReasonTagged) GetTypeName() string {
	return "BACnetRestartReasonTagged"
}

func (m *_BACnetRestartReasonTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetRestartReasonTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRestartReasonTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetRestartReasonTagged, error) {
	return BACnetRestartReasonTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetRestartReasonTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRestartReasonTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRestartReasonTagged, error) {
		return BACnetRestartReasonTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetRestartReasonTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetRestartReasonTagged, error) {
	v, err := (&_BACnetRestartReasonTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetRestartReasonTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetRestartReasonTagged BACnetRestartReasonTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRestartReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRestartReasonTagged")
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

	value, err := ReadManualField[BACnetRestartReason](ctx, "value", readBuffer, EnsureType[BACnetRestartReason](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetRestartReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetRestartReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetRestartReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRestartReasonTagged")
	}

	return m, nil
}

func (m *_BACnetRestartReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRestartReasonTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRestartReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRestartReasonTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetRestartReason](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetRestartReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRestartReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRestartReasonTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetRestartReasonTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetRestartReasonTagged) IsBACnetRestartReasonTagged() {}

func (m *_BACnetRestartReasonTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRestartReasonTagged) deepCopy() *_BACnetRestartReasonTagged {
	if m == nil {
		return nil
	}
	_BACnetRestartReasonTaggedCopy := &_BACnetRestartReasonTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetRestartReasonTaggedCopy
}

func (m *_BACnetRestartReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
