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

// BACnetEventLogRecordLogDatum is the corresponding interface of BACnetEventLogRecordLogDatum
type BACnetEventLogRecordLogDatum interface {
	BACnetEventLogRecordLogDatumContract
	BACnetEventLogRecordLogDatumRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetEventLogRecordLogDatum is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventLogRecordLogDatum()
	// CreateBuilder creates a BACnetEventLogRecordLogDatumBuilder
	CreateBACnetEventLogRecordLogDatumBuilder() BACnetEventLogRecordLogDatumBuilder
}

// BACnetEventLogRecordLogDatumContract provides a set of functions which can be overwritten by a sub struct
type BACnetEventLogRecordLogDatumContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetEventLogRecordLogDatum is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventLogRecordLogDatum()
	// CreateBuilder creates a BACnetEventLogRecordLogDatumBuilder
	CreateBACnetEventLogRecordLogDatumBuilder() BACnetEventLogRecordLogDatumBuilder
}

// BACnetEventLogRecordLogDatumRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetEventLogRecordLogDatumRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetEventLogRecordLogDatum is the data-structure of this message
type _BACnetEventLogRecordLogDatum struct {
	_SubType interface {
		BACnetEventLogRecordLogDatumContract
		BACnetEventLogRecordLogDatumRequirements
	}
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventLogRecordLogDatumContract = (*_BACnetEventLogRecordLogDatum)(nil)

// NewBACnetEventLogRecordLogDatum factory function for _BACnetEventLogRecordLogDatum
func NewBACnetEventLogRecordLogDatum(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventLogRecordLogDatum {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventLogRecordLogDatum must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetEventLogRecordLogDatum must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventLogRecordLogDatum must not be nil")
	}
	return &_BACnetEventLogRecordLogDatum{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventLogRecordLogDatumBuilder is a builder for BACnetEventLogRecordLogDatum
type BACnetEventLogRecordLogDatumBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetEventLogRecordLogDatumBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventLogRecordLogDatumBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventLogRecordLogDatumBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetEventLogRecordLogDatumBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetEventLogRecordLogDatumBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventLogRecordLogDatumBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventLogRecordLogDatumBuilder
	// AsBACnetEventLogRecordLogDatumLogStatus converts this build to a subType of BACnetEventLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetEventLogRecordLogDatumLogStatus() interface {
		BACnetEventLogRecordLogDatumLogStatusBuilder
		Done() BACnetEventLogRecordLogDatumBuilder
	}
	// AsBACnetEventLogRecordLogDatumNotification converts this build to a subType of BACnetEventLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetEventLogRecordLogDatumNotification() interface {
		BACnetEventLogRecordLogDatumNotificationBuilder
		Done() BACnetEventLogRecordLogDatumBuilder
	}
	// AsBACnetEventLogRecordLogDatumTimeChange converts this build to a subType of BACnetEventLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetEventLogRecordLogDatumTimeChange() interface {
		BACnetEventLogRecordLogDatumTimeChangeBuilder
		Done() BACnetEventLogRecordLogDatumBuilder
	}
	// Build builds the BACnetEventLogRecordLogDatum or returns an error if something is wrong
	PartialBuild() (BACnetEventLogRecordLogDatumContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetEventLogRecordLogDatumContract
	// Build builds the BACnetEventLogRecordLogDatum or returns an error if something is wrong
	Build() (BACnetEventLogRecordLogDatum, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventLogRecordLogDatum
}

// NewBACnetEventLogRecordLogDatumBuilder() creates a BACnetEventLogRecordLogDatumBuilder
func NewBACnetEventLogRecordLogDatumBuilder() BACnetEventLogRecordLogDatumBuilder {
	return &_BACnetEventLogRecordLogDatumBuilder{_BACnetEventLogRecordLogDatum: new(_BACnetEventLogRecordLogDatum)}
}

type _BACnetEventLogRecordLogDatumChildBuilder interface {
	utils.Copyable
	setParent(BACnetEventLogRecordLogDatumContract)
	buildForBACnetEventLogRecordLogDatum() (BACnetEventLogRecordLogDatum, error)
}

type _BACnetEventLogRecordLogDatumBuilder struct {
	*_BACnetEventLogRecordLogDatum

	childBuilder _BACnetEventLogRecordLogDatumChildBuilder

	err *utils.MultiError
}

var _ (BACnetEventLogRecordLogDatumBuilder) = (*_BACnetEventLogRecordLogDatumBuilder)(nil)

func (b *_BACnetEventLogRecordLogDatumBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetEventLogRecordLogDatumBuilder {
	return b.WithOpeningTag(openingTag).WithPeekedTagHeader(peekedTagHeader).WithClosingTag(closingTag)
}

func (b *_BACnetEventLogRecordLogDatumBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventLogRecordLogDatumBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventLogRecordLogDatumBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventLogRecordLogDatumBuilder {
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

func (b *_BACnetEventLogRecordLogDatumBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetEventLogRecordLogDatumBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetEventLogRecordLogDatumBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetEventLogRecordLogDatumBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetEventLogRecordLogDatumBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventLogRecordLogDatumBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventLogRecordLogDatumBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventLogRecordLogDatumBuilder {
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

func (b *_BACnetEventLogRecordLogDatumBuilder) PartialBuild() (BACnetEventLogRecordLogDatumContract, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
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
	return b._BACnetEventLogRecordLogDatum.deepCopy(), nil
}

func (b *_BACnetEventLogRecordLogDatumBuilder) PartialMustBuild() BACnetEventLogRecordLogDatumContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventLogRecordLogDatumBuilder) AsBACnetEventLogRecordLogDatumLogStatus() interface {
	BACnetEventLogRecordLogDatumLogStatusBuilder
	Done() BACnetEventLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetEventLogRecordLogDatumLogStatusBuilder
		Done() BACnetEventLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetEventLogRecordLogDatumLogStatusBuilder().(*_BACnetEventLogRecordLogDatumLogStatusBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetEventLogRecordLogDatumBuilder) AsBACnetEventLogRecordLogDatumNotification() interface {
	BACnetEventLogRecordLogDatumNotificationBuilder
	Done() BACnetEventLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetEventLogRecordLogDatumNotificationBuilder
		Done() BACnetEventLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetEventLogRecordLogDatumNotificationBuilder().(*_BACnetEventLogRecordLogDatumNotificationBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetEventLogRecordLogDatumBuilder) AsBACnetEventLogRecordLogDatumTimeChange() interface {
	BACnetEventLogRecordLogDatumTimeChangeBuilder
	Done() BACnetEventLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetEventLogRecordLogDatumTimeChangeBuilder
		Done() BACnetEventLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetEventLogRecordLogDatumTimeChangeBuilder().(*_BACnetEventLogRecordLogDatumTimeChangeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetEventLogRecordLogDatumBuilder) Build() (BACnetEventLogRecordLogDatum, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetEventLogRecordLogDatum()
}

func (b *_BACnetEventLogRecordLogDatumBuilder) MustBuild() BACnetEventLogRecordLogDatum {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventLogRecordLogDatumBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventLogRecordLogDatumBuilder().(*_BACnetEventLogRecordLogDatumBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetEventLogRecordLogDatumChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventLogRecordLogDatumBuilder creates a BACnetEventLogRecordLogDatumBuilder
func (b *_BACnetEventLogRecordLogDatum) CreateBACnetEventLogRecordLogDatumBuilder() BACnetEventLogRecordLogDatumBuilder {
	if b == nil {
		return NewBACnetEventLogRecordLogDatumBuilder()
	}
	return &_BACnetEventLogRecordLogDatumBuilder{_BACnetEventLogRecordLogDatum: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatum) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventLogRecordLogDatum) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetEventLogRecordLogDatum) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetEventLogRecordLogDatum) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatum(structType any) BACnetEventLogRecordLogDatum {
	if casted, ok := structType.(BACnetEventLogRecordLogDatum); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatum); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatum) GetTypeName() string {
	return "BACnetEventLogRecordLogDatum"
}

func (m *_BACnetEventLogRecordLogDatum) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatum) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetEventLogRecordLogDatum) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetEventLogRecordLogDatumParse[T BACnetEventLogRecordLogDatum](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetEventLogRecordLogDatumParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventLogRecordLogDatumParseWithBufferProducer[T BACnetEventLogRecordLogDatum](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetEventLogRecordLogDatumParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetEventLogRecordLogDatumParseWithBuffer[T BACnetEventLogRecordLogDatum](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetEventLogRecordLogDatum{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetEventLogRecordLogDatum) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventLogRecordLogDatum BACnetEventLogRecordLogDatum, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetEventLogRecordLogDatum
	switch {
	case peekedTagNumber == uint8(0): // BACnetEventLogRecordLogDatumLogStatus
		if _child, err = new(_BACnetEventLogRecordLogDatumLogStatus).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventLogRecordLogDatumLogStatus for type-switch of BACnetEventLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(1): // BACnetEventLogRecordLogDatumNotification
		if _child, err = new(_BACnetEventLogRecordLogDatumNotification).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventLogRecordLogDatumNotification for type-switch of BACnetEventLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(2): // BACnetEventLogRecordLogDatumTimeChange
		if _child, err = new(_BACnetEventLogRecordLogDatumTimeChange).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventLogRecordLogDatumTimeChange for type-switch of BACnetEventLogRecordLogDatum")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatum")
	}

	return _child, nil
}

func (pm *_BACnetEventLogRecordLogDatum) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventLogRecordLogDatum, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatum")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatum")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventLogRecordLogDatum) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventLogRecordLogDatum) IsBACnetEventLogRecordLogDatum() {}

func (m *_BACnetEventLogRecordLogDatum) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventLogRecordLogDatum) deepCopy() *_BACnetEventLogRecordLogDatum {
	if m == nil {
		return nil
	}
	_BACnetEventLogRecordLogDatumCopy := &_BACnetEventLogRecordLogDatum{
		nil, // will be set by child
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetEventLogRecordLogDatumCopy
}
