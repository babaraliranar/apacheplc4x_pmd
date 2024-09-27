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

// ErrorCodeTagged is the corresponding interface of ErrorCodeTagged
type ErrorCodeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() ErrorCode
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsErrorCodeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorCodeTagged()
	// CreateBuilder creates a ErrorCodeTaggedBuilder
	CreateErrorCodeTaggedBuilder() ErrorCodeTaggedBuilder
}

// _ErrorCodeTagged is the data-structure of this message
type _ErrorCodeTagged struct {
	Header           BACnetTagHeader
	Value            ErrorCode
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ ErrorCodeTagged = (*_ErrorCodeTagged)(nil)

// NewErrorCodeTagged factory function for _ErrorCodeTagged
func NewErrorCodeTagged(header BACnetTagHeader, value ErrorCode, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_ErrorCodeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for ErrorCodeTagged must not be nil")
	}
	return &_ErrorCodeTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorCodeTaggedBuilder is a builder for ErrorCodeTagged
type ErrorCodeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value ErrorCode, proprietaryValue uint32) ErrorCodeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) ErrorCodeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) ErrorCodeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(ErrorCode) ErrorCodeTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) ErrorCodeTaggedBuilder
	// Build builds the ErrorCodeTagged or returns an error if something is wrong
	Build() (ErrorCodeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorCodeTagged
}

// NewErrorCodeTaggedBuilder() creates a ErrorCodeTaggedBuilder
func NewErrorCodeTaggedBuilder() ErrorCodeTaggedBuilder {
	return &_ErrorCodeTaggedBuilder{_ErrorCodeTagged: new(_ErrorCodeTagged)}
}

type _ErrorCodeTaggedBuilder struct {
	*_ErrorCodeTagged

	err *utils.MultiError
}

var _ (ErrorCodeTaggedBuilder) = (*_ErrorCodeTaggedBuilder)(nil)

func (m *_ErrorCodeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value ErrorCode, proprietaryValue uint32) ErrorCodeTaggedBuilder {
	return m.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (m *_ErrorCodeTaggedBuilder) WithHeader(header BACnetTagHeader) ErrorCodeTaggedBuilder {
	m.Header = header
	return m
}

func (m *_ErrorCodeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) ErrorCodeTaggedBuilder {
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

func (m *_ErrorCodeTaggedBuilder) WithValue(value ErrorCode) ErrorCodeTaggedBuilder {
	m.Value = value
	return m
}

func (m *_ErrorCodeTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) ErrorCodeTaggedBuilder {
	m.ProprietaryValue = proprietaryValue
	return m
}

func (m *_ErrorCodeTaggedBuilder) Build() (ErrorCodeTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ErrorCodeTagged.deepCopy(), nil
}

func (m *_ErrorCodeTaggedBuilder) MustBuild() ErrorCodeTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ErrorCodeTaggedBuilder) DeepCopy() any {
	return m.CreateErrorCodeTaggedBuilder()
}

// CreateErrorCodeTaggedBuilder creates a ErrorCodeTaggedBuilder
func (m *_ErrorCodeTagged) CreateErrorCodeTaggedBuilder() ErrorCodeTaggedBuilder {
	if m == nil {
		return NewErrorCodeTaggedBuilder()
	}
	return &_ErrorCodeTaggedBuilder{_ErrorCodeTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorCodeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_ErrorCodeTagged) GetValue() ErrorCode {
	return m.Value
}

func (m *_ErrorCodeTagged) GetProprietaryValue() uint32 {
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

func (m *_ErrorCodeTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (ErrorCode_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorCodeTagged(structType any) ErrorCodeTagged {
	if casted, ok := structType.(ErrorCodeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorCodeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorCodeTagged) GetTypeName() string {
	return "ErrorCodeTagged"
}

func (m *_ErrorCodeTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_ErrorCodeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorCodeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (ErrorCodeTagged, error) {
	return ErrorCodeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func ErrorCodeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorCodeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorCodeTagged, error) {
		return ErrorCodeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func ErrorCodeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (ErrorCodeTagged, error) {
	v, err := (&_ErrorCodeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ErrorCodeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__errorCodeTagged ErrorCodeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorCodeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorCodeTagged")
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

	value, err := ReadManualField[ErrorCode](ctx, "value", readBuffer, EnsureType[ErrorCode](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), ErrorCode_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (ErrorCode_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("ErrorCodeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorCodeTagged")
	}

	return m, nil
}

func (m *_ErrorCodeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorCodeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorCodeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorCodeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[ErrorCode](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("ErrorCodeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorCodeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_ErrorCodeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_ErrorCodeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_ErrorCodeTagged) IsErrorCodeTagged() {}

func (m *_ErrorCodeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorCodeTagged) deepCopy() *_ErrorCodeTagged {
	if m == nil {
		return nil
	}
	_ErrorCodeTaggedCopy := &_ErrorCodeTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _ErrorCodeTaggedCopy
}

func (m *_ErrorCodeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
