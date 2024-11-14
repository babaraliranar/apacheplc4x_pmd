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

// CBusCommand is the corresponding interface of CBusCommand
type CBusCommand interface {
	CBusCommandContract
	CBusCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCBusCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusCommand()
	// CreateBuilder creates a CBusCommandBuilder
	CreateCBusCommandBuilder() CBusCommandBuilder
}

// CBusCommandContract provides a set of functions which can be overwritten by a sub struct
type CBusCommandContract interface {
	// GetHeader returns Header (property field)
	GetHeader() CBusHeader
	// GetIsDeviceManagement returns IsDeviceManagement (virtual field)
	GetIsDeviceManagement() bool
	// GetDestinationAddressType returns DestinationAddressType (virtual field)
	GetDestinationAddressType() DestinationAddressType
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// IsCBusCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusCommand()
	// CreateBuilder creates a CBusCommandBuilder
	CreateCBusCommandBuilder() CBusCommandBuilder
}

// CBusCommandRequirements provides a set of functions which need to be implemented by a sub struct
type CBusCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetDestinationAddressType returns DestinationAddressType (discriminator field)
	GetDestinationAddressType() DestinationAddressType
	// GetIsDeviceManagement returns IsDeviceManagement (discriminator field)
	GetIsDeviceManagement() bool
}

// _CBusCommand is the data-structure of this message
type _CBusCommand struct {
	_SubType interface {
		CBusCommandContract
		CBusCommandRequirements
	}
	Header CBusHeader

	// Arguments.
	CBusOptions CBusOptions
}

var _ CBusCommandContract = (*_CBusCommand)(nil)

// NewCBusCommand factory function for _CBusCommand
func NewCBusCommand(header CBusHeader, cBusOptions CBusOptions) *_CBusCommand {
	if header == nil {
		panic("header of type CBusHeader for CBusCommand must not be nil")
	}
	return &_CBusCommand{Header: header, CBusOptions: cBusOptions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusCommandBuilder is a builder for CBusCommand
type CBusCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header CBusHeader) CBusCommandBuilder
	// WithHeader adds Header (property field)
	WithHeader(CBusHeader) CBusCommandBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(CBusHeaderBuilder) CBusHeaderBuilder) CBusCommandBuilder
	// WithArgCBusOptions sets a parser argument
	WithArgCBusOptions(CBusOptions) CBusCommandBuilder
	// AsCBusCommandDeviceManagement converts this build to a subType of CBusCommand. It is always possible to return to current builder using Done()
	AsCBusCommandDeviceManagement() CBusCommandDeviceManagementBuilder
	// AsCBusCommandPointToPointToMultiPoint converts this build to a subType of CBusCommand. It is always possible to return to current builder using Done()
	AsCBusCommandPointToPointToMultiPoint() CBusCommandPointToPointToMultiPointBuilder
	// AsCBusCommandPointToMultiPoint converts this build to a subType of CBusCommand. It is always possible to return to current builder using Done()
	AsCBusCommandPointToMultiPoint() CBusCommandPointToMultiPointBuilder
	// AsCBusCommandPointToPoint converts this build to a subType of CBusCommand. It is always possible to return to current builder using Done()
	AsCBusCommandPointToPoint() CBusCommandPointToPointBuilder
	// Build builds the CBusCommand or returns an error if something is wrong
	PartialBuild() (CBusCommandContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() CBusCommandContract
	// Build builds the CBusCommand or returns an error if something is wrong
	Build() (CBusCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusCommand
}

// NewCBusCommandBuilder() creates a CBusCommandBuilder
func NewCBusCommandBuilder() CBusCommandBuilder {
	return &_CBusCommandBuilder{_CBusCommand: new(_CBusCommand)}
}

type _CBusCommandChildBuilder interface {
	utils.Copyable
	setParent(CBusCommandContract)
	buildForCBusCommand() (CBusCommand, error)
}

type _CBusCommandBuilder struct {
	*_CBusCommand

	childBuilder _CBusCommandChildBuilder

	err *utils.MultiError
}

var _ (CBusCommandBuilder) = (*_CBusCommandBuilder)(nil)

func (b *_CBusCommandBuilder) WithMandatoryFields(header CBusHeader) CBusCommandBuilder {
	return b.WithHeader(header)
}

func (b *_CBusCommandBuilder) WithHeader(header CBusHeader) CBusCommandBuilder {
	b.Header = header
	return b
}

func (b *_CBusCommandBuilder) WithHeaderBuilder(builderSupplier func(CBusHeaderBuilder) CBusHeaderBuilder) CBusCommandBuilder {
	builder := builderSupplier(b.Header.CreateCBusHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CBusHeaderBuilder failed"))
	}
	return b
}

func (b *_CBusCommandBuilder) WithArgCBusOptions(cBusOptions CBusOptions) CBusCommandBuilder {
	b.CBusOptions = cBusOptions
	return b
}

func (b *_CBusCommandBuilder) PartialBuild() (CBusCommandContract, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusCommand.deepCopy(), nil
}

func (b *_CBusCommandBuilder) PartialMustBuild() CBusCommandContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusCommandBuilder) AsCBusCommandDeviceManagement() CBusCommandDeviceManagementBuilder {
	if cb, ok := b.childBuilder.(CBusCommandDeviceManagementBuilder); ok {
		return cb
	}
	cb := NewCBusCommandDeviceManagementBuilder().(*_CBusCommandDeviceManagementBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusCommandBuilder) AsCBusCommandPointToPointToMultiPoint() CBusCommandPointToPointToMultiPointBuilder {
	if cb, ok := b.childBuilder.(CBusCommandPointToPointToMultiPointBuilder); ok {
		return cb
	}
	cb := NewCBusCommandPointToPointToMultiPointBuilder().(*_CBusCommandPointToPointToMultiPointBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusCommandBuilder) AsCBusCommandPointToMultiPoint() CBusCommandPointToMultiPointBuilder {
	if cb, ok := b.childBuilder.(CBusCommandPointToMultiPointBuilder); ok {
		return cb
	}
	cb := NewCBusCommandPointToMultiPointBuilder().(*_CBusCommandPointToMultiPointBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusCommandBuilder) AsCBusCommandPointToPoint() CBusCommandPointToPointBuilder {
	if cb, ok := b.childBuilder.(CBusCommandPointToPointBuilder); ok {
		return cb
	}
	cb := NewCBusCommandPointToPointBuilder().(*_CBusCommandPointToPointBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusCommandBuilder) Build() (CBusCommand, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCBusCommand()
}

func (b *_CBusCommandBuilder) MustBuild() CBusCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusCommandBuilder) DeepCopy() any {
	_copy := b.CreateCBusCommandBuilder().(*_CBusCommandBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_CBusCommandChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusCommandBuilder creates a CBusCommandBuilder
func (b *_CBusCommand) CreateCBusCommandBuilder() CBusCommandBuilder {
	if b == nil {
		return NewCBusCommandBuilder()
	}
	return &_CBusCommandBuilder{_CBusCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommand) GetHeader() CBusHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_CBusCommand) GetIsDeviceManagement() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(m.GetHeader().GetDp())
}

func (pm *_CBusCommand) GetDestinationAddressType() DestinationAddressType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastDestinationAddressType(m.GetHeader().GetDestinationAddressType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusCommand(structType any) CBusCommand {
	if casted, ok := structType.(CBusCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommand) GetTypeName() string {
	return "CBusCommand"
}

func (m *_CBusCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_CBusCommand) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_CBusCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CBusCommandParse[T CBusCommand](ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (T, error) {
	return CBusCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusCommandParseWithBufferProducer[T CBusCommand](cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CBusCommandParseWithBuffer[T](ctx, readBuffer, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CBusCommandParseWithBuffer[T CBusCommand](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (T, error) {
	v, err := (&_CBusCommand{CBusOptions: cBusOptions}).parse(ctx, readBuffer, cBusOptions)
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

func (m *_CBusCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (__cBusCommand CBusCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[CBusHeader](ctx, "header", ReadComplex[CBusHeader](CBusHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	isDeviceManagement, err := ReadVirtualField[bool](ctx, "isDeviceManagement", (*bool)(nil), header.GetDp())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDeviceManagement' field"))
	}
	_ = isDeviceManagement

	destinationAddressType, err := ReadVirtualField[DestinationAddressType](ctx, "destinationAddressType", (*DestinationAddressType)(nil), header.GetDestinationAddressType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationAddressType' field"))
	}
	_ = destinationAddressType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CBusCommand
	switch {
	case 0 == 0 && isDeviceManagement == bool(true): // CBusCommandDeviceManagement
		if _child, err = new(_CBusCommandDeviceManagement).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusCommandDeviceManagement for type-switch of CBusCommand")
		}
	case destinationAddressType == DestinationAddressType_PointToPointToMultiPoint: // CBusCommandPointToPointToMultiPoint
		if _child, err = new(_CBusCommandPointToPointToMultiPoint).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusCommandPointToPointToMultiPoint for type-switch of CBusCommand")
		}
	case destinationAddressType == DestinationAddressType_PointToMultiPoint: // CBusCommandPointToMultiPoint
		if _child, err = new(_CBusCommandPointToMultiPoint).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusCommandPointToMultiPoint for type-switch of CBusCommand")
		}
	case destinationAddressType == DestinationAddressType_PointToPoint: // CBusCommandPointToPoint
		if _child, err = new(_CBusCommandPointToPoint).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusCommandPointToPoint for type-switch of CBusCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [destinationAddressType=%v, isDeviceManagement=%v]", destinationAddressType, isDeviceManagement)
	}

	if closeErr := readBuffer.CloseContext("CBusCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommand")
	}

	return _child, nil
}

func (pm *_CBusCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusCommand")
	}

	if err := WriteSimpleField[CBusHeader](ctx, "header", m.GetHeader(), WriteComplex[CBusHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}
	// Virtual field
	isDeviceManagement := m.GetIsDeviceManagement()
	_ = isDeviceManagement
	if _isDeviceManagementErr := writeBuffer.WriteVirtual(ctx, "isDeviceManagement", m.GetIsDeviceManagement()); _isDeviceManagementErr != nil {
		return errors.Wrap(_isDeviceManagementErr, "Error serializing 'isDeviceManagement' field")
	}
	// Virtual field
	destinationAddressType := m.GetDestinationAddressType()
	_ = destinationAddressType
	if _destinationAddressTypeErr := writeBuffer.WriteVirtual(ctx, "destinationAddressType", m.GetDestinationAddressType()); _destinationAddressTypeErr != nil {
		return errors.Wrap(_destinationAddressTypeErr, "Error serializing 'destinationAddressType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusCommand) IsCBusCommand() {}

func (m *_CBusCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusCommand) deepCopy() *_CBusCommand {
	if m == nil {
		return nil
	}
	_CBusCommandCopy := &_CBusCommand{
		nil, // will be set by child
		utils.DeepCopy[CBusHeader](m.Header),
		m.CBusOptions,
	}
	return _CBusCommandCopy
}
