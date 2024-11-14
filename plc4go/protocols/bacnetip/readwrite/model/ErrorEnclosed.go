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

// ErrorEnclosed is the corresponding interface of ErrorEnclosed
type ErrorEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetError returns Error (property field)
	GetError() Error
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsErrorEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorEnclosed()
	// CreateBuilder creates a ErrorEnclosedBuilder
	CreateErrorEnclosedBuilder() ErrorEnclosedBuilder
}

// _ErrorEnclosed is the data-structure of this message
type _ErrorEnclosed struct {
	OpeningTag BACnetOpeningTag
	Error      Error
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ ErrorEnclosed = (*_ErrorEnclosed)(nil)

// NewErrorEnclosed factory function for _ErrorEnclosed
func NewErrorEnclosed(openingTag BACnetOpeningTag, error Error, closingTag BACnetClosingTag, tagNumber uint8) *_ErrorEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for ErrorEnclosed must not be nil")
	}
	if error == nil {
		panic("error of type Error for ErrorEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for ErrorEnclosed must not be nil")
	}
	return &_ErrorEnclosed{OpeningTag: openingTag, Error: error, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorEnclosedBuilder is a builder for ErrorEnclosed
type ErrorEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, error Error, closingTag BACnetClosingTag) ErrorEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) ErrorEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) ErrorEnclosedBuilder
	// WithError adds Error (property field)
	WithError(Error) ErrorEnclosedBuilder
	// WithErrorBuilder adds Error (property field) which is build by the builder
	WithErrorBuilder(func(ErrorBuilder) ErrorBuilder) ErrorEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) ErrorEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) ErrorEnclosedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) ErrorEnclosedBuilder
	// Build builds the ErrorEnclosed or returns an error if something is wrong
	Build() (ErrorEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorEnclosed
}

// NewErrorEnclosedBuilder() creates a ErrorEnclosedBuilder
func NewErrorEnclosedBuilder() ErrorEnclosedBuilder {
	return &_ErrorEnclosedBuilder{_ErrorEnclosed: new(_ErrorEnclosed)}
}

type _ErrorEnclosedBuilder struct {
	*_ErrorEnclosed

	err *utils.MultiError
}

var _ (ErrorEnclosedBuilder) = (*_ErrorEnclosedBuilder)(nil)

func (b *_ErrorEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, error Error, closingTag BACnetClosingTag) ErrorEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithError(error).WithClosingTag(closingTag)
}

func (b *_ErrorEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) ErrorEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_ErrorEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) ErrorEnclosedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_ErrorEnclosedBuilder) WithError(error Error) ErrorEnclosedBuilder {
	b.Error = error
	return b
}

func (b *_ErrorEnclosedBuilder) WithErrorBuilder(builderSupplier func(ErrorBuilder) ErrorBuilder) ErrorEnclosedBuilder {
	builder := builderSupplier(b.Error.CreateErrorBuilder())
	var err error
	b.Error, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorBuilder failed"))
	}
	return b
}

func (b *_ErrorEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) ErrorEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_ErrorEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) ErrorEnclosedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_ErrorEnclosedBuilder) WithArgTagNumber(tagNumber uint8) ErrorEnclosedBuilder {
	b.TagNumber = tagNumber
	return b
}

func (b *_ErrorEnclosedBuilder) Build() (ErrorEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.Error == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'error' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorEnclosed.deepCopy(), nil
}

func (b *_ErrorEnclosedBuilder) MustBuild() ErrorEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateErrorEnclosedBuilder().(*_ErrorEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorEnclosedBuilder creates a ErrorEnclosedBuilder
func (b *_ErrorEnclosed) CreateErrorEnclosedBuilder() ErrorEnclosedBuilder {
	if b == nil {
		return NewErrorEnclosedBuilder()
	}
	return &_ErrorEnclosedBuilder{_ErrorEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ErrorEnclosed) GetError() Error {
	return m.Error
}

func (m *_ErrorEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorEnclosed(structType any) ErrorEnclosed {
	if casted, ok := structType.(ErrorEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorEnclosed) GetTypeName() string {
	return "ErrorEnclosed"
}

func (m *_ErrorEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ErrorEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (ErrorEnclosed, error) {
	return ErrorEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func ErrorEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorEnclosed, error) {
		return ErrorEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func ErrorEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (ErrorEnclosed, error) {
	v, err := (&_ErrorEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ErrorEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__errorEnclosed ErrorEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	error, err := ReadSimpleField[Error](ctx, "error", ReadComplex[Error](ErrorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("ErrorEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorEnclosed")
	}

	return m, nil
}

func (m *_ErrorEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[Error](ctx, "error", m.GetError(), WriteComplex[Error](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'error' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("ErrorEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_ErrorEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_ErrorEnclosed) IsErrorEnclosed() {}

func (m *_ErrorEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorEnclosed) deepCopy() *_ErrorEnclosed {
	if m == nil {
		return nil
	}
	_ErrorEnclosedCopy := &_ErrorEnclosed{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[Error](m.Error),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _ErrorEnclosedCopy
}

func (m *_ErrorEnclosed) String() string {
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
