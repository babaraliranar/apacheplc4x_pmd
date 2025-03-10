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

// BACnetHostAddressEnclosed is the corresponding interface of BACnetHostAddressEnclosed
type BACnetHostAddressEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetHostAddress returns HostAddress (property field)
	GetHostAddress() BACnetHostAddress
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetHostAddressEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetHostAddressEnclosed()
	// CreateBuilder creates a BACnetHostAddressEnclosedBuilder
	CreateBACnetHostAddressEnclosedBuilder() BACnetHostAddressEnclosedBuilder
}

// _BACnetHostAddressEnclosed is the data-structure of this message
type _BACnetHostAddressEnclosed struct {
	OpeningTag  BACnetOpeningTag
	HostAddress BACnetHostAddress
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetHostAddressEnclosed = (*_BACnetHostAddressEnclosed)(nil)

// NewBACnetHostAddressEnclosed factory function for _BACnetHostAddressEnclosed
func NewBACnetHostAddressEnclosed(openingTag BACnetOpeningTag, hostAddress BACnetHostAddress, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetHostAddressEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetHostAddressEnclosed must not be nil")
	}
	if hostAddress == nil {
		panic("hostAddress of type BACnetHostAddress for BACnetHostAddressEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetHostAddressEnclosed must not be nil")
	}
	return &_BACnetHostAddressEnclosed{OpeningTag: openingTag, HostAddress: hostAddress, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetHostAddressEnclosedBuilder is a builder for BACnetHostAddressEnclosed
type BACnetHostAddressEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, hostAddress BACnetHostAddress, closingTag BACnetClosingTag) BACnetHostAddressEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetHostAddressEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetHostAddressEnclosedBuilder
	// WithHostAddress adds HostAddress (property field)
	WithHostAddress(BACnetHostAddress) BACnetHostAddressEnclosedBuilder
	// WithHostAddressBuilder adds HostAddress (property field) which is build by the builder
	WithHostAddressBuilder(func(BACnetHostAddressBuilder) BACnetHostAddressBuilder) BACnetHostAddressEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetHostAddressEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetHostAddressEnclosedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetHostAddressEnclosedBuilder
	// Build builds the BACnetHostAddressEnclosed or returns an error if something is wrong
	Build() (BACnetHostAddressEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetHostAddressEnclosed
}

// NewBACnetHostAddressEnclosedBuilder() creates a BACnetHostAddressEnclosedBuilder
func NewBACnetHostAddressEnclosedBuilder() BACnetHostAddressEnclosedBuilder {
	return &_BACnetHostAddressEnclosedBuilder{_BACnetHostAddressEnclosed: new(_BACnetHostAddressEnclosed)}
}

type _BACnetHostAddressEnclosedBuilder struct {
	*_BACnetHostAddressEnclosed

	err *utils.MultiError
}

var _ (BACnetHostAddressEnclosedBuilder) = (*_BACnetHostAddressEnclosedBuilder)(nil)

func (b *_BACnetHostAddressEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, hostAddress BACnetHostAddress, closingTag BACnetClosingTag) BACnetHostAddressEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithHostAddress(hostAddress).WithClosingTag(closingTag)
}

func (b *_BACnetHostAddressEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetHostAddressEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetHostAddressEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetHostAddressEnclosedBuilder {
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

func (b *_BACnetHostAddressEnclosedBuilder) WithHostAddress(hostAddress BACnetHostAddress) BACnetHostAddressEnclosedBuilder {
	b.HostAddress = hostAddress
	return b
}

func (b *_BACnetHostAddressEnclosedBuilder) WithHostAddressBuilder(builderSupplier func(BACnetHostAddressBuilder) BACnetHostAddressBuilder) BACnetHostAddressEnclosedBuilder {
	builder := builderSupplier(b.HostAddress.CreateBACnetHostAddressBuilder())
	var err error
	b.HostAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetHostAddressBuilder failed"))
	}
	return b
}

func (b *_BACnetHostAddressEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetHostAddressEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetHostAddressEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetHostAddressEnclosedBuilder {
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

func (b *_BACnetHostAddressEnclosedBuilder) WithArgTagNumber(tagNumber uint8) BACnetHostAddressEnclosedBuilder {
	b.TagNumber = tagNumber
	return b
}

func (b *_BACnetHostAddressEnclosedBuilder) Build() (BACnetHostAddressEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.HostAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hostAddress' not set"))
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
	return b._BACnetHostAddressEnclosed.deepCopy(), nil
}

func (b *_BACnetHostAddressEnclosedBuilder) MustBuild() BACnetHostAddressEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetHostAddressEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetHostAddressEnclosedBuilder().(*_BACnetHostAddressEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetHostAddressEnclosedBuilder creates a BACnetHostAddressEnclosedBuilder
func (b *_BACnetHostAddressEnclosed) CreateBACnetHostAddressEnclosedBuilder() BACnetHostAddressEnclosedBuilder {
	if b == nil {
		return NewBACnetHostAddressEnclosedBuilder()
	}
	return &_BACnetHostAddressEnclosedBuilder{_BACnetHostAddressEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetHostAddressEnclosed) GetHostAddress() BACnetHostAddress {
	return m.HostAddress
}

func (m *_BACnetHostAddressEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressEnclosed(structType any) BACnetHostAddressEnclosed {
	if casted, ok := structType.(BACnetHostAddressEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressEnclosed) GetTypeName() string {
	return "BACnetHostAddressEnclosed"
}

func (m *_BACnetHostAddressEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (hostAddress)
	lengthInBits += m.HostAddress.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostAddressEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostAddressEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetHostAddressEnclosed, error) {
	return BACnetHostAddressEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetHostAddressEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostAddressEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostAddressEnclosed, error) {
		return BACnetHostAddressEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetHostAddressEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetHostAddressEnclosed, error) {
	v, err := (&_BACnetHostAddressEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetHostAddressEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetHostAddressEnclosed BACnetHostAddressEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	hostAddress, err := ReadSimpleField[BACnetHostAddress](ctx, "hostAddress", ReadComplex[BACnetHostAddress](BACnetHostAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hostAddress' field"))
	}
	m.HostAddress = hostAddress

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetHostAddressEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressEnclosed")
	}

	return m, nil
}

func (m *_BACnetHostAddressEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostAddressEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetHostAddressEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetHostAddress](ctx, "hostAddress", m.GetHostAddress(), WriteComplex[BACnetHostAddress](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'hostAddress' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostAddressEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostAddressEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetHostAddressEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetHostAddressEnclosed) IsBACnetHostAddressEnclosed() {}

func (m *_BACnetHostAddressEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetHostAddressEnclosed) deepCopy() *_BACnetHostAddressEnclosed {
	if m == nil {
		return nil
	}
	_BACnetHostAddressEnclosedCopy := &_BACnetHostAddressEnclosed{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetHostAddress](m.HostAddress),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetHostAddressEnclosedCopy
}

func (m *_BACnetHostAddressEnclosed) String() string {
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
