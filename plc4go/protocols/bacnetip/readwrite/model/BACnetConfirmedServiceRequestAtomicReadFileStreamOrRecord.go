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

// BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord interface {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	CreateBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
}

// BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract provides a set of functions which can be overwritten by a sub struct
type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	CreateBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
}

// BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord struct {
	_SubType        BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	PeekedTagHeader BACnetTagHeader
	OpeningTag      BACnetOpeningTag
	ClosingTag      BACnetClosingTag
}

var _ BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = (*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)(nil)

// NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord factory function for _BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
func NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord must not be nil")
	}
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord must not be nil")
	}
	return &_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder is a builder for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// Build builds the BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
}

// NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder() creates a BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
func NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	return &_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder{_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord: new(_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)}
}

type _BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder struct {
	*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) = (*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	return m.WithPeekedTagHeader(peekedTagHeader).WithOpeningTag(openingTag).WithClosingTag(closingTag)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	m.PeekedTagHeader = peekedTagHeader
	return m
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	builder := builderSupplier(m.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	m.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) Build() (BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract, error) {
	if m.PeekedTagHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) MustBuild() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder()
}

// CreateBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) CreateBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder()
	}
	return &_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder{_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) GetClosingTag() BACnetClosingTag {
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

func (pm *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) GetPeekedTagNumber() uint8 {
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
func CastBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(structType any) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParse[T BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParseWithBufferProducer[T BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParseWithBuffer[T BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord{}).parse(ctx, readBuffer)
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

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	switch {
	case peekedTagNumber == 0x0: // BACnetConfirmedServiceRequestAtomicReadFileStream
		if _child, err = new(_BACnetConfirmedServiceRequestAtomicReadFileStream).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAtomicReadFileStream for type-switch of BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord")
		}
	case peekedTagNumber == 0x1: // BACnetConfirmedServiceRequestAtomicReadFileRecord
		if _child, err = new(_BACnetConfirmedServiceRequestAtomicReadFileRecord).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAtomicReadFileRecord for type-switch of BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord")
	}

	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord")
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

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord")
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) IsBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord() {
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) deepCopy() *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordCopy := &_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	return _BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordCopy
}
