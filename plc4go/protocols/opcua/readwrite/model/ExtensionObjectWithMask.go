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

// ExtensionObjectWithMask is the corresponding interface of ExtensionObjectWithMask
type ExtensionObjectWithMask interface {
	ExtensionObjectWithMaskContract
	ExtensionObjectWithMaskRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsExtensionObjectWithMask is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensionObjectWithMask()
	// CreateBuilder creates a ExtensionObjectWithMaskBuilder
	CreateExtensionObjectWithMaskBuilder() ExtensionObjectWithMaskBuilder
}

// ExtensionObjectWithMaskContract provides a set of functions which can be overwritten by a sub struct
type ExtensionObjectWithMaskContract interface {
	ExtensionObjectContract
	// GetEncodingMask returns EncodingMask (property field)
	GetEncodingMask() ExtensionObjectEncodingMask
	// GetExtensionId() returns a parser argument
	GetExtensionId() int32
	// GetIncludeEncodingMask returns IncludeEncodingMask (discriminator field)
	GetIncludeEncodingMask() bool
	// IsExtensionObjectWithMask is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensionObjectWithMask()
	// CreateBuilder creates a ExtensionObjectWithMaskBuilder
	CreateExtensionObjectWithMaskBuilder() ExtensionObjectWithMaskBuilder
}

// ExtensionObjectWithMaskRequirements provides a set of functions which need to be implemented by a sub struct
type ExtensionObjectWithMaskRequirements interface {
	ExtensionObjectRequirements
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
}

// _ExtensionObjectWithMask is the data-structure of this message
type _ExtensionObjectWithMask struct {
	ExtensionObjectContract
	_SubType interface {
		ExtensionObjectWithMaskContract
		ExtensionObjectWithMaskRequirements
	}
	EncodingMask ExtensionObjectEncodingMask

	// Arguments.
	ExtensionId int32
}

var _ ExtensionObjectWithMaskContract = (*_ExtensionObjectWithMask)(nil)

// NewExtensionObjectWithMask factory function for _ExtensionObjectWithMask
func NewExtensionObjectWithMask(typeId ExpandedNodeId, encodingMask ExtensionObjectEncodingMask, extensionId int32) *_ExtensionObjectWithMask {
	if encodingMask == nil {
		panic("encodingMask of type ExtensionObjectEncodingMask for ExtensionObjectWithMask must not be nil")
	}
	_result := &_ExtensionObjectWithMask{
		ExtensionObjectContract: NewExtensionObject(typeId),
		EncodingMask:            encodingMask,
	}
	_result.ExtensionObjectContract.(*_ExtensionObject)._SubType = _result._SubType
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ExtensionObjectWithMaskBuilder is a builder for ExtensionObjectWithMask
type ExtensionObjectWithMaskBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(encodingMask ExtensionObjectEncodingMask) ExtensionObjectWithMaskBuilder
	// WithEncodingMask adds EncodingMask (property field)
	WithEncodingMask(ExtensionObjectEncodingMask) ExtensionObjectWithMaskBuilder
	// WithEncodingMaskBuilder adds EncodingMask (property field) which is build by the builder
	WithEncodingMaskBuilder(func(ExtensionObjectEncodingMaskBuilder) ExtensionObjectEncodingMaskBuilder) ExtensionObjectWithMaskBuilder
	// AsBinaryExtensionObjectWithMask converts this build to a subType of ExtensionObjectWithMask. It is always possible to return to current builder using Done()
	AsBinaryExtensionObjectWithMask() interface {
		BinaryExtensionObjectWithMaskBuilder
		Done() ExtensionObjectWithMaskBuilder
	}
	// AsNullExtensionObjectWithMask converts this build to a subType of ExtensionObjectWithMask. It is always possible to return to current builder using Done()
	AsNullExtensionObjectWithMask() interface {
		NullExtensionObjectWithMaskBuilder
		Done() ExtensionObjectWithMaskBuilder
	}
	// Build builds the ExtensionObjectWithMask or returns an error if something is wrong
	PartialBuild() (ExtensionObjectWithMaskContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ExtensionObjectWithMaskContract
	// Build builds the ExtensionObjectWithMask or returns an error if something is wrong
	Build() (ExtensionObjectWithMask, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ExtensionObjectWithMask
}

// NewExtensionObjectWithMaskBuilder() creates a ExtensionObjectWithMaskBuilder
func NewExtensionObjectWithMaskBuilder() ExtensionObjectWithMaskBuilder {
	return &_ExtensionObjectWithMaskBuilder{_ExtensionObjectWithMask: new(_ExtensionObjectWithMask)}
}

type _ExtensionObjectWithMaskChildBuilder interface {
	utils.Copyable
	setParent(ExtensionObjectWithMaskContract)
	buildForExtensionObjectWithMask() (ExtensionObjectWithMask, error)
}

type _ExtensionObjectWithMaskBuilder struct {
	*_ExtensionObjectWithMask

	parentBuilder *_ExtensionObjectBuilder

	childBuilder _ExtensionObjectWithMaskChildBuilder

	err *utils.MultiError
}

var _ (ExtensionObjectWithMaskBuilder) = (*_ExtensionObjectWithMaskBuilder)(nil)

func (b *_ExtensionObjectWithMaskBuilder) setParent(contract ExtensionObjectContract) {
	b.ExtensionObjectContract = contract
}

func (b *_ExtensionObjectWithMaskBuilder) WithMandatoryFields(encodingMask ExtensionObjectEncodingMask) ExtensionObjectWithMaskBuilder {
	return b.WithEncodingMask(encodingMask)
}

func (b *_ExtensionObjectWithMaskBuilder) WithEncodingMask(encodingMask ExtensionObjectEncodingMask) ExtensionObjectWithMaskBuilder {
	b.EncodingMask = encodingMask
	return b
}

func (b *_ExtensionObjectWithMaskBuilder) WithEncodingMaskBuilder(builderSupplier func(ExtensionObjectEncodingMaskBuilder) ExtensionObjectEncodingMaskBuilder) ExtensionObjectWithMaskBuilder {
	builder := builderSupplier(b.EncodingMask.CreateExtensionObjectEncodingMaskBuilder())
	var err error
	b.EncodingMask, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectEncodingMaskBuilder failed"))
	}
	return b
}

func (b *_ExtensionObjectWithMaskBuilder) PartialBuild() (ExtensionObjectWithMaskContract, error) {
	if b.EncodingMask == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'encodingMask' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ExtensionObjectWithMask.deepCopy(), nil
}

func (b *_ExtensionObjectWithMaskBuilder) PartialMustBuild() ExtensionObjectWithMaskContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ExtensionObjectWithMaskBuilder) AsBinaryExtensionObjectWithMask() interface {
	BinaryExtensionObjectWithMaskBuilder
	Done() ExtensionObjectWithMaskBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BinaryExtensionObjectWithMaskBuilder
		Done() ExtensionObjectWithMaskBuilder
	}); ok {
		return cb
	}
	cb := NewBinaryExtensionObjectWithMaskBuilder().(*_BinaryExtensionObjectWithMaskBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ExtensionObjectWithMaskBuilder) AsNullExtensionObjectWithMask() interface {
	NullExtensionObjectWithMaskBuilder
	Done() ExtensionObjectWithMaskBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NullExtensionObjectWithMaskBuilder
		Done() ExtensionObjectWithMaskBuilder
	}); ok {
		return cb
	}
	cb := NewNullExtensionObjectWithMaskBuilder().(*_NullExtensionObjectWithMaskBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ExtensionObjectWithMaskBuilder) Done() ExtensionObjectBuilder {
	return b.parentBuilder
}

func (b *_ExtensionObjectWithMaskBuilder) buildForExtensionObject() (ExtensionObject, error) {
	return b.Build()
}

func (b *_ExtensionObjectWithMaskBuilder) Build() (ExtensionObjectWithMask, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForExtensionObjectWithMask()
}

func (b *_ExtensionObjectWithMaskBuilder) MustBuild() ExtensionObjectWithMask {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ExtensionObjectWithMaskBuilder) DeepCopy() any {
	_copy := b.CreateExtensionObjectWithMaskBuilder().(*_ExtensionObjectWithMaskBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ExtensionObjectWithMaskChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateExtensionObjectWithMaskBuilder creates a ExtensionObjectWithMaskBuilder
func (b *_ExtensionObjectWithMask) CreateExtensionObjectWithMaskBuilder() ExtensionObjectWithMaskBuilder {
	if b == nil {
		return NewExtensionObjectWithMaskBuilder()
	}
	return &_ExtensionObjectWithMaskBuilder{_ExtensionObjectWithMask: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ExtensionObjectWithMask) GetIncludeEncodingMask() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensionObjectWithMask) GetEncodingMask() ExtensionObjectEncodingMask {
	return m.EncodingMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastExtensionObjectWithMask(structType any) ExtensionObjectWithMask {
	if casted, ok := structType.(ExtensionObjectWithMask); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensionObjectWithMask); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensionObjectWithMask) GetTypeName() string {
	return "ExtensionObjectWithMask"
}

func (m *_ExtensionObjectWithMask) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (encodingMask)
	lengthInBits += m.EncodingMask.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ExtensionObjectWithMask) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_ExtensionObjectWithMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func (m *_ExtensionObjectWithMask) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObject, extensionId int32, includeEncodingMask bool) (__extensionObjectWithMask ExtensionObjectWithMask, err error) {
	m.ExtensionObjectContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtensionObjectWithMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensionObjectWithMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	encodingMask, err := ReadSimpleField[ExtensionObjectEncodingMask](ctx, "encodingMask", ReadComplex[ExtensionObjectEncodingMask](ExtensionObjectEncodingMaskParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encodingMask' field"))
	}
	m.EncodingMask = encodingMask

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ExtensionObjectWithMask
	switch {
	case encodingMask.GetXmlBody() == (false) && encodingMask.GetBinaryBody() == (true): // BinaryExtensionObjectWithMask
		if _child, err = new(_BinaryExtensionObjectWithMask).parse(ctx, readBuffer, m, extensionId, includeEncodingMask); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BinaryExtensionObjectWithMask for type-switch of ExtensionObjectWithMask")
		}
	case encodingMask.GetXmlBody() == (false) && encodingMask.GetBinaryBody() == (false): // NullExtensionObjectWithMask
		if _child, err = new(_NullExtensionObjectWithMask).parse(ctx, readBuffer, m, extensionId, includeEncodingMask); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NullExtensionObjectWithMask for type-switch of ExtensionObjectWithMask")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [encodingMaskxmlBody=%v, encodingMaskbinaryBody=%v]", encodingMask.GetXmlBody(), encodingMask.GetBinaryBody())
	}

	if closeErr := readBuffer.CloseContext("ExtensionObjectWithMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensionObjectWithMask")
	}

	parent._SubType = _child
	return _child, nil
}

func (pm *_ExtensionObjectWithMask) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ExtensionObjectWithMask, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ExtensionObjectWithMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ExtensionObjectWithMask")
		}

		if err := WriteSimpleField[ExtensionObjectEncodingMask](ctx, "encodingMask", m.GetEncodingMask(), WriteComplex[ExtensionObjectEncodingMask](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'encodingMask' field")
		}

		// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
		if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
			return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
		}

		if popErr := writeBuffer.PopContext("ExtensionObjectWithMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ExtensionObjectWithMask")
		}
		return nil
	}
	return pm.ExtensionObjectContract.(*_ExtensionObject).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_ExtensionObjectWithMask) GetExtensionId() int32 {
	return m.ExtensionId
}

//
////

func (m *_ExtensionObjectWithMask) IsExtensionObjectWithMask() {}

func (m *_ExtensionObjectWithMask) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ExtensionObjectWithMask) deepCopy() *_ExtensionObjectWithMask {
	if m == nil {
		return nil
	}
	_ExtensionObjectWithMaskCopy := &_ExtensionObjectWithMask{
		m.ExtensionObjectContract.(*_ExtensionObject).deepCopy(),
		nil, // will be set by child
		m.EncodingMask.DeepCopy().(ExtensionObjectEncodingMask),
		m.ExtensionId,
	}
	m.ExtensionObjectContract.(*_ExtensionObject)._SubType = m
	return _ExtensionObjectWithMaskCopy
}
