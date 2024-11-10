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

// LightingData is the corresponding interface of LightingData
type LightingData interface {
	LightingDataContract
	LightingDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsLightingData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLightingData()
	// CreateBuilder creates a LightingDataBuilder
	CreateLightingDataBuilder() LightingDataBuilder
}

// LightingDataContract provides a set of functions which can be overwritten by a sub struct
type LightingDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() LightingCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() LightingCommandType
	// IsLightingData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLightingData()
	// CreateBuilder creates a LightingDataBuilder
	CreateLightingDataBuilder() LightingDataBuilder
}

// LightingDataRequirements provides a set of functions which need to be implemented by a sub struct
type LightingDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() LightingCommandType
}

// _LightingData is the data-structure of this message
type _LightingData struct {
	_SubType interface {
		LightingDataContract
		LightingDataRequirements
	}
	CommandTypeContainer LightingCommandTypeContainer
}

var _ LightingDataContract = (*_LightingData)(nil)

// NewLightingData factory function for _LightingData
func NewLightingData(commandTypeContainer LightingCommandTypeContainer) *_LightingData {
	return &_LightingData{CommandTypeContainer: commandTypeContainer}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LightingDataBuilder is a builder for LightingData
type LightingDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(commandTypeContainer LightingCommandTypeContainer) LightingDataBuilder
	// WithCommandTypeContainer adds CommandTypeContainer (property field)
	WithCommandTypeContainer(LightingCommandTypeContainer) LightingDataBuilder
	// AsLightingDataOff converts this build to a subType of LightingData. It is always possible to return to current builder using Done()
	AsLightingDataOff() interface {
		LightingDataOffBuilder
		Done() LightingDataBuilder
	}
	// AsLightingDataOn converts this build to a subType of LightingData. It is always possible to return to current builder using Done()
	AsLightingDataOn() interface {
		LightingDataOnBuilder
		Done() LightingDataBuilder
	}
	// AsLightingDataRampToLevel converts this build to a subType of LightingData. It is always possible to return to current builder using Done()
	AsLightingDataRampToLevel() interface {
		LightingDataRampToLevelBuilder
		Done() LightingDataBuilder
	}
	// AsLightingDataTerminateRamp converts this build to a subType of LightingData. It is always possible to return to current builder using Done()
	AsLightingDataTerminateRamp() interface {
		LightingDataTerminateRampBuilder
		Done() LightingDataBuilder
	}
	// AsLightingDataLabel converts this build to a subType of LightingData. It is always possible to return to current builder using Done()
	AsLightingDataLabel() interface {
		LightingDataLabelBuilder
		Done() LightingDataBuilder
	}
	// Build builds the LightingData or returns an error if something is wrong
	PartialBuild() (LightingDataContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() LightingDataContract
	// Build builds the LightingData or returns an error if something is wrong
	Build() (LightingData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LightingData
}

// NewLightingDataBuilder() creates a LightingDataBuilder
func NewLightingDataBuilder() LightingDataBuilder {
	return &_LightingDataBuilder{_LightingData: new(_LightingData)}
}

type _LightingDataChildBuilder interface {
	utils.Copyable
	setParent(LightingDataContract)
	buildForLightingData() (LightingData, error)
}

type _LightingDataBuilder struct {
	*_LightingData

	childBuilder _LightingDataChildBuilder

	err *utils.MultiError
}

var _ (LightingDataBuilder) = (*_LightingDataBuilder)(nil)

func (b *_LightingDataBuilder) WithMandatoryFields(commandTypeContainer LightingCommandTypeContainer) LightingDataBuilder {
	return b.WithCommandTypeContainer(commandTypeContainer)
}

func (b *_LightingDataBuilder) WithCommandTypeContainer(commandTypeContainer LightingCommandTypeContainer) LightingDataBuilder {
	b.CommandTypeContainer = commandTypeContainer
	return b
}

func (b *_LightingDataBuilder) PartialBuild() (LightingDataContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LightingData.deepCopy(), nil
}

func (b *_LightingDataBuilder) PartialMustBuild() LightingDataContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LightingDataBuilder) AsLightingDataOff() interface {
	LightingDataOffBuilder
	Done() LightingDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LightingDataOffBuilder
		Done() LightingDataBuilder
	}); ok {
		return cb
	}
	cb := NewLightingDataOffBuilder().(*_LightingDataOffBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LightingDataBuilder) AsLightingDataOn() interface {
	LightingDataOnBuilder
	Done() LightingDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LightingDataOnBuilder
		Done() LightingDataBuilder
	}); ok {
		return cb
	}
	cb := NewLightingDataOnBuilder().(*_LightingDataOnBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LightingDataBuilder) AsLightingDataRampToLevel() interface {
	LightingDataRampToLevelBuilder
	Done() LightingDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LightingDataRampToLevelBuilder
		Done() LightingDataBuilder
	}); ok {
		return cb
	}
	cb := NewLightingDataRampToLevelBuilder().(*_LightingDataRampToLevelBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LightingDataBuilder) AsLightingDataTerminateRamp() interface {
	LightingDataTerminateRampBuilder
	Done() LightingDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LightingDataTerminateRampBuilder
		Done() LightingDataBuilder
	}); ok {
		return cb
	}
	cb := NewLightingDataTerminateRampBuilder().(*_LightingDataTerminateRampBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LightingDataBuilder) AsLightingDataLabel() interface {
	LightingDataLabelBuilder
	Done() LightingDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LightingDataLabelBuilder
		Done() LightingDataBuilder
	}); ok {
		return cb
	}
	cb := NewLightingDataLabelBuilder().(*_LightingDataLabelBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LightingDataBuilder) Build() (LightingData, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForLightingData()
}

func (b *_LightingDataBuilder) MustBuild() LightingData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LightingDataBuilder) DeepCopy() any {
	_copy := b.CreateLightingDataBuilder().(*_LightingDataBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_LightingDataChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLightingDataBuilder creates a LightingDataBuilder
func (b *_LightingData) CreateLightingDataBuilder() LightingDataBuilder {
	if b == nil {
		return NewLightingDataBuilder()
	}
	return &_LightingDataBuilder{_LightingData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingData) GetCommandTypeContainer() LightingCommandTypeContainer {
	return m.CommandTypeContainer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_LightingData) GetCommandType() LightingCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastLightingCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLightingData(structType any) LightingData {
	if casted, ok := structType.(LightingData); ok {
		return casted
	}
	if casted, ok := structType.(*LightingData); ok {
		return *casted
	}
	return nil
}

func (m *_LightingData) GetTypeName() string {
	return "LightingData"
}

func (m *_LightingData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_LightingData) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_LightingData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func LightingDataParse[T LightingData](ctx context.Context, theBytes []byte) (T, error) {
	return LightingDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func LightingDataParseWithBufferProducer[T LightingData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := LightingDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func LightingDataParseWithBuffer[T LightingData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_LightingData{}).parse(ctx, readBuffer)
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

func (m *_LightingData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__lightingData LightingData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsLightingCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[LightingCommandTypeContainer](ctx, "commandTypeContainer", "LightingCommandTypeContainer", ReadEnum(LightingCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[LightingCommandType](ctx, "commandType", (*LightingCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child LightingData
	switch {
	case commandType == LightingCommandType_OFF: // LightingDataOff
		if _child, err = new(_LightingDataOff).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LightingDataOff for type-switch of LightingData")
		}
	case commandType == LightingCommandType_ON: // LightingDataOn
		if _child, err = new(_LightingDataOn).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LightingDataOn for type-switch of LightingData")
		}
	case commandType == LightingCommandType_RAMP_TO_LEVEL: // LightingDataRampToLevel
		if _child, err = new(_LightingDataRampToLevel).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LightingDataRampToLevel for type-switch of LightingData")
		}
	case commandType == LightingCommandType_TERMINATE_RAMP: // LightingDataTerminateRamp
		if _child, err = new(_LightingDataTerminateRamp).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LightingDataTerminateRamp for type-switch of LightingData")
		}
	case commandType == LightingCommandType_LABEL: // LightingDataLabel
		if _child, err = new(_LightingDataLabel).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LightingDataLabel for type-switch of LightingData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("LightingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingData")
	}

	return _child, nil
}

func (pm *_LightingData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LightingData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LightingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LightingData")
	}

	if err := WriteSimpleEnumField[LightingCommandTypeContainer](ctx, "commandTypeContainer", "LightingCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[LightingCommandTypeContainer, uint8](LightingCommandTypeContainer.GetValue, LightingCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LightingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LightingData")
	}
	return nil
}

func (m *_LightingData) IsLightingData() {}

func (m *_LightingData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LightingData) deepCopy() *_LightingData {
	if m == nil {
		return nil
	}
	_LightingDataCopy := &_LightingData{
		nil, // will be set by child
		m.CommandTypeContainer,
	}
	return _LightingDataCopy
}
