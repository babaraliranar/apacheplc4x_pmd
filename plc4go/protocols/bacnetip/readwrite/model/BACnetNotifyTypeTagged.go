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

// BACnetNotifyTypeTagged is the corresponding interface of BACnetNotifyTypeTagged
type BACnetNotifyTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNotifyType
	// IsBACnetNotifyTypeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotifyTypeTagged()
	// CreateBuilder creates a BACnetNotifyTypeTaggedBuilder
	CreateBACnetNotifyTypeTaggedBuilder() BACnetNotifyTypeTaggedBuilder
}

// _BACnetNotifyTypeTagged is the data-structure of this message
type _BACnetNotifyTypeTagged struct {
	Header BACnetTagHeader
	Value  BACnetNotifyType

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetNotifyTypeTagged = (*_BACnetNotifyTypeTagged)(nil)

// NewBACnetNotifyTypeTagged factory function for _BACnetNotifyTypeTagged
func NewBACnetNotifyTypeTagged(header BACnetTagHeader, value BACnetNotifyType, tagNumber uint8, tagClass TagClass) *_BACnetNotifyTypeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetNotifyTypeTagged must not be nil")
	}
	return &_BACnetNotifyTypeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotifyTypeTaggedBuilder is a builder for BACnetNotifyTypeTagged
type BACnetNotifyTypeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetNotifyType) BACnetNotifyTypeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetNotifyTypeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNotifyTypeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetNotifyType) BACnetNotifyTypeTaggedBuilder
	// Build builds the BACnetNotifyTypeTagged or returns an error if something is wrong
	Build() (BACnetNotifyTypeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotifyTypeTagged
}

// NewBACnetNotifyTypeTaggedBuilder() creates a BACnetNotifyTypeTaggedBuilder
func NewBACnetNotifyTypeTaggedBuilder() BACnetNotifyTypeTaggedBuilder {
	return &_BACnetNotifyTypeTaggedBuilder{_BACnetNotifyTypeTagged: new(_BACnetNotifyTypeTagged)}
}

type _BACnetNotifyTypeTaggedBuilder struct {
	*_BACnetNotifyTypeTagged

	err *utils.MultiError
}

var _ (BACnetNotifyTypeTaggedBuilder) = (*_BACnetNotifyTypeTaggedBuilder)(nil)

func (m *_BACnetNotifyTypeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetNotifyType) BACnetNotifyTypeTaggedBuilder {
	return m.WithHeader(header).WithValue(value)
}

func (m *_BACnetNotifyTypeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetNotifyTypeTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetNotifyTypeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNotifyTypeTaggedBuilder {
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

func (m *_BACnetNotifyTypeTaggedBuilder) WithValue(value BACnetNotifyType) BACnetNotifyTypeTaggedBuilder {
	m.Value = value
	return m
}

func (m *_BACnetNotifyTypeTaggedBuilder) Build() (BACnetNotifyTypeTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetNotifyTypeTagged.deepCopy(), nil
}

func (m *_BACnetNotifyTypeTaggedBuilder) MustBuild() BACnetNotifyTypeTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetNotifyTypeTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetNotifyTypeTaggedBuilder()
}

// CreateBACnetNotifyTypeTaggedBuilder creates a BACnetNotifyTypeTaggedBuilder
func (m *_BACnetNotifyTypeTagged) CreateBACnetNotifyTypeTaggedBuilder() BACnetNotifyTypeTaggedBuilder {
	if m == nil {
		return NewBACnetNotifyTypeTaggedBuilder()
	}
	return &_BACnetNotifyTypeTaggedBuilder{_BACnetNotifyTypeTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotifyTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNotifyTypeTagged) GetValue() BACnetNotifyType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotifyTypeTagged(structType any) BACnetNotifyTypeTagged {
	if casted, ok := structType.(BACnetNotifyTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotifyTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotifyTypeTagged) GetTypeName() string {
	return "BACnetNotifyTypeTagged"
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotifyTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetNotifyTypeTagged, error) {
	return BACnetNotifyTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetNotifyTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNotifyTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNotifyTypeTagged, error) {
		return BACnetNotifyTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetNotifyTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNotifyTypeTagged, error) {
	v, err := (&_BACnetNotifyTypeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetNotifyTypeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetNotifyTypeTagged BACnetNotifyTypeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotifyTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotifyTypeTagged")
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

	value, err := ReadManualField[BACnetNotifyType](ctx, "value", readBuffer, EnsureType[BACnetNotifyType](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetNotifyType_ALARM)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetNotifyTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotifyTypeTagged")
	}

	return m, nil
}

func (m *_BACnetNotifyTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotifyTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotifyTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotifyTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetNotifyType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotifyTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotifyTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotifyTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNotifyTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetNotifyTypeTagged) IsBACnetNotifyTypeTagged() {}

func (m *_BACnetNotifyTypeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotifyTypeTagged) deepCopy() *_BACnetNotifyTypeTagged {
	if m == nil {
		return nil
	}
	_BACnetNotifyTypeTaggedCopy := &_BACnetNotifyTypeTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetNotifyTypeTaggedCopy
}

func (m *_BACnetNotifyTypeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
