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

// BACnetEscalatorFaultTagged is the corresponding interface of BACnetEscalatorFaultTagged
type BACnetEscalatorFaultTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetEscalatorFault
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetEscalatorFaultTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEscalatorFaultTagged()
	// CreateBuilder creates a BACnetEscalatorFaultTaggedBuilder
	CreateBACnetEscalatorFaultTaggedBuilder() BACnetEscalatorFaultTaggedBuilder
}

// _BACnetEscalatorFaultTagged is the data-structure of this message
type _BACnetEscalatorFaultTagged struct {
	Header           BACnetTagHeader
	Value            BACnetEscalatorFault
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetEscalatorFaultTagged = (*_BACnetEscalatorFaultTagged)(nil)

// NewBACnetEscalatorFaultTagged factory function for _BACnetEscalatorFaultTagged
func NewBACnetEscalatorFaultTagged(header BACnetTagHeader, value BACnetEscalatorFault, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetEscalatorFaultTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetEscalatorFaultTagged must not be nil")
	}
	return &_BACnetEscalatorFaultTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEscalatorFaultTaggedBuilder is a builder for BACnetEscalatorFaultTagged
type BACnetEscalatorFaultTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetEscalatorFault, proprietaryValue uint32) BACnetEscalatorFaultTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetEscalatorFaultTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetEscalatorFaultTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetEscalatorFault) BACnetEscalatorFaultTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetEscalatorFaultTaggedBuilder
	// Build builds the BACnetEscalatorFaultTagged or returns an error if something is wrong
	Build() (BACnetEscalatorFaultTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEscalatorFaultTagged
}

// NewBACnetEscalatorFaultTaggedBuilder() creates a BACnetEscalatorFaultTaggedBuilder
func NewBACnetEscalatorFaultTaggedBuilder() BACnetEscalatorFaultTaggedBuilder {
	return &_BACnetEscalatorFaultTaggedBuilder{_BACnetEscalatorFaultTagged: new(_BACnetEscalatorFaultTagged)}
}

type _BACnetEscalatorFaultTaggedBuilder struct {
	*_BACnetEscalatorFaultTagged

	err *utils.MultiError
}

var _ (BACnetEscalatorFaultTaggedBuilder) = (*_BACnetEscalatorFaultTaggedBuilder)(nil)

func (m *_BACnetEscalatorFaultTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetEscalatorFault, proprietaryValue uint32) BACnetEscalatorFaultTaggedBuilder {
	return m.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (m *_BACnetEscalatorFaultTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetEscalatorFaultTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetEscalatorFaultTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetEscalatorFaultTaggedBuilder {
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

func (m *_BACnetEscalatorFaultTaggedBuilder) WithValue(value BACnetEscalatorFault) BACnetEscalatorFaultTaggedBuilder {
	m.Value = value
	return m
}

func (m *_BACnetEscalatorFaultTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetEscalatorFaultTaggedBuilder {
	m.ProprietaryValue = proprietaryValue
	return m
}

func (m *_BACnetEscalatorFaultTaggedBuilder) Build() (BACnetEscalatorFaultTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetEscalatorFaultTagged.deepCopy(), nil
}

func (m *_BACnetEscalatorFaultTaggedBuilder) MustBuild() BACnetEscalatorFaultTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEscalatorFaultTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetEscalatorFaultTaggedBuilder()
}

// CreateBACnetEscalatorFaultTaggedBuilder creates a BACnetEscalatorFaultTaggedBuilder
func (m *_BACnetEscalatorFaultTagged) CreateBACnetEscalatorFaultTaggedBuilder() BACnetEscalatorFaultTaggedBuilder {
	if m == nil {
		return NewBACnetEscalatorFaultTaggedBuilder()
	}
	return &_BACnetEscalatorFaultTaggedBuilder{_BACnetEscalatorFaultTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEscalatorFaultTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetEscalatorFaultTagged) GetValue() BACnetEscalatorFault {
	return m.Value
}

func (m *_BACnetEscalatorFaultTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetEscalatorFaultTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEscalatorFaultTagged(structType any) BACnetEscalatorFaultTagged {
	if casted, ok := structType.(BACnetEscalatorFaultTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEscalatorFaultTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEscalatorFaultTagged) GetTypeName() string {
	return "BACnetEscalatorFaultTagged"
}

func (m *_BACnetEscalatorFaultTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetEscalatorFaultTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEscalatorFaultTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetEscalatorFaultTagged, error) {
	return BACnetEscalatorFaultTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetEscalatorFaultTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEscalatorFaultTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEscalatorFaultTagged, error) {
		return BACnetEscalatorFaultTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetEscalatorFaultTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetEscalatorFaultTagged, error) {
	v, err := (&_BACnetEscalatorFaultTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEscalatorFaultTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetEscalatorFaultTagged BACnetEscalatorFaultTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEscalatorFaultTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEscalatorFaultTagged")
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

	value, err := ReadManualField[BACnetEscalatorFault](ctx, "value", readBuffer, EnsureType[BACnetEscalatorFault](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetEscalatorFault_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetEscalatorFaultTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEscalatorFaultTagged")
	}

	return m, nil
}

func (m *_BACnetEscalatorFaultTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEscalatorFaultTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEscalatorFaultTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEscalatorFaultTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetEscalatorFault](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetEscalatorFaultTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEscalatorFaultTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEscalatorFaultTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetEscalatorFaultTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetEscalatorFaultTagged) IsBACnetEscalatorFaultTagged() {}

func (m *_BACnetEscalatorFaultTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEscalatorFaultTagged) deepCopy() *_BACnetEscalatorFaultTagged {
	if m == nil {
		return nil
	}
	_BACnetEscalatorFaultTaggedCopy := &_BACnetEscalatorFaultTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetEscalatorFaultTaggedCopy
}

func (m *_BACnetEscalatorFaultTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
