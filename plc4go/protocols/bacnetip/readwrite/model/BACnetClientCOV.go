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

// BACnetClientCOV is the corresponding interface of BACnetClientCOV
type BACnetClientCOV interface {
	BACnetClientCOVContract
	BACnetClientCOVRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetClientCOV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetClientCOV()
	// CreateBuilder creates a BACnetClientCOVBuilder
	CreateBACnetClientCOVBuilder() BACnetClientCOVBuilder
}

// BACnetClientCOVContract provides a set of functions which can be overwritten by a sub struct
type BACnetClientCOVContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetClientCOV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetClientCOV()
	// CreateBuilder creates a BACnetClientCOVBuilder
	CreateBACnetClientCOVBuilder() BACnetClientCOVBuilder
}

// BACnetClientCOVRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetClientCOVRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetClientCOV is the data-structure of this message
type _BACnetClientCOV struct {
	_SubType interface {
		BACnetClientCOVContract
		BACnetClientCOVRequirements
	}
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetClientCOVContract = (*_BACnetClientCOV)(nil)

// NewBACnetClientCOV factory function for _BACnetClientCOV
func NewBACnetClientCOV(peekedTagHeader BACnetTagHeader) *_BACnetClientCOV {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetClientCOV must not be nil")
	}
	return &_BACnetClientCOV{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetClientCOVBuilder is a builder for BACnetClientCOV
type BACnetClientCOVBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetClientCOVBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetClientCOVBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetClientCOVBuilder
	// AsBACnetClientCOVObject converts this build to a subType of BACnetClientCOV. It is always possible to return to current builder using Done()
	AsBACnetClientCOVObject() interface {
		BACnetClientCOVObjectBuilder
		Done() BACnetClientCOVBuilder
	}
	// AsBACnetClientCOVNone converts this build to a subType of BACnetClientCOV. It is always possible to return to current builder using Done()
	AsBACnetClientCOVNone() interface {
		BACnetClientCOVNoneBuilder
		Done() BACnetClientCOVBuilder
	}
	// Build builds the BACnetClientCOV or returns an error if something is wrong
	PartialBuild() (BACnetClientCOVContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetClientCOVContract
	// Build builds the BACnetClientCOV or returns an error if something is wrong
	Build() (BACnetClientCOV, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetClientCOV
}

// NewBACnetClientCOVBuilder() creates a BACnetClientCOVBuilder
func NewBACnetClientCOVBuilder() BACnetClientCOVBuilder {
	return &_BACnetClientCOVBuilder{_BACnetClientCOV: new(_BACnetClientCOV)}
}

type _BACnetClientCOVChildBuilder interface {
	utils.Copyable
	setParent(BACnetClientCOVContract)
	buildForBACnetClientCOV() (BACnetClientCOV, error)
}

type _BACnetClientCOVBuilder struct {
	*_BACnetClientCOV

	childBuilder _BACnetClientCOVChildBuilder

	err *utils.MultiError
}

var _ (BACnetClientCOVBuilder) = (*_BACnetClientCOVBuilder)(nil)

func (b *_BACnetClientCOVBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetClientCOVBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader)
}

func (b *_BACnetClientCOVBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetClientCOVBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetClientCOVBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetClientCOVBuilder {
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

func (b *_BACnetClientCOVBuilder) PartialBuild() (BACnetClientCOVContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetClientCOV.deepCopy(), nil
}

func (b *_BACnetClientCOVBuilder) PartialMustBuild() BACnetClientCOVContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetClientCOVBuilder) AsBACnetClientCOVObject() interface {
	BACnetClientCOVObjectBuilder
	Done() BACnetClientCOVBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetClientCOVObjectBuilder
		Done() BACnetClientCOVBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetClientCOVObjectBuilder().(*_BACnetClientCOVObjectBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetClientCOVBuilder) AsBACnetClientCOVNone() interface {
	BACnetClientCOVNoneBuilder
	Done() BACnetClientCOVBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetClientCOVNoneBuilder
		Done() BACnetClientCOVBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetClientCOVNoneBuilder().(*_BACnetClientCOVNoneBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetClientCOVBuilder) Build() (BACnetClientCOV, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetClientCOV()
}

func (b *_BACnetClientCOVBuilder) MustBuild() BACnetClientCOV {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetClientCOVBuilder) DeepCopy() any {
	_copy := b.CreateBACnetClientCOVBuilder().(*_BACnetClientCOVBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetClientCOVChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetClientCOVBuilder creates a BACnetClientCOVBuilder
func (b *_BACnetClientCOV) CreateBACnetClientCOVBuilder() BACnetClientCOVBuilder {
	if b == nil {
		return NewBACnetClientCOVBuilder()
	}
	return &_BACnetClientCOVBuilder{_BACnetClientCOV: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetClientCOV) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetClientCOV) GetPeekedTagNumber() uint8 {
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
func CastBACnetClientCOV(structType any) BACnetClientCOV {
	if casted, ok := structType.(BACnetClientCOV); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetClientCOV); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetClientCOV) GetTypeName() string {
	return "BACnetClientCOV"
}

func (m *_BACnetClientCOV) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetClientCOV) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetClientCOVParse[T BACnetClientCOV](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetClientCOVParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetClientCOVParseWithBufferProducer[T BACnetClientCOV]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetClientCOVParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetClientCOVParseWithBuffer[T BACnetClientCOV](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetClientCOV{}).parse(ctx, readBuffer)
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

func (m *_BACnetClientCOV) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetClientCOV BACnetClientCOV, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClientCOV"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetClientCOV")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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
	var _child BACnetClientCOV
	switch {
	case peekedTagNumber == 0x4: // BACnetClientCOVObject
		if _child, err = new(_BACnetClientCOVObject).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetClientCOVObject for type-switch of BACnetClientCOV")
		}
	case peekedTagNumber == 0x0: // BACnetClientCOVNone
		if _child, err = new(_BACnetClientCOVNone).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetClientCOVNone for type-switch of BACnetClientCOV")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetClientCOV"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetClientCOV")
	}

	return _child, nil
}

func (pm *_BACnetClientCOV) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetClientCOV, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetClientCOV"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetClientCOV")
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

	if popErr := writeBuffer.PopContext("BACnetClientCOV"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetClientCOV")
	}
	return nil
}

func (m *_BACnetClientCOV) IsBACnetClientCOV() {}

func (m *_BACnetClientCOV) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetClientCOV) deepCopy() *_BACnetClientCOV {
	if m == nil {
		return nil
	}
	_BACnetClientCOVCopy := &_BACnetClientCOV{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetClientCOVCopy
}
