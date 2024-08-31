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

// TelephonyData is the corresponding interface of TelephonyData
type TelephonyData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() TelephonyCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() TelephonyCommandType
}

// TelephonyDataExactly can be used when we want exactly this type and not a type which fulfills TelephonyData.
// This is useful for switch cases.
type TelephonyDataExactly interface {
	TelephonyData
	isTelephonyData() bool
}

// _TelephonyData is the data-structure of this message
type _TelephonyData struct {
	_TelephonyDataChildRequirements
	CommandTypeContainer TelephonyCommandTypeContainer
	Argument             byte
}

type _TelephonyDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetArgument() byte
	GetCommandType() TelephonyCommandType
}

type TelephonyDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TelephonyData, serializeChildFunction func() error) error
	GetTypeName() string
}

type TelephonyDataChild interface {
	utils.Serializable
	InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte)
	GetParent() *TelephonyData

	GetTypeName() string
	TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyData) GetCommandTypeContainer() TelephonyCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_TelephonyData) GetArgument() byte {
	return m.Argument
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyData) GetCommandType() TelephonyCommandType {
	ctx := context.Background()
	_ = ctx
	return CastTelephonyCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyData factory function for _TelephonyData
func NewTelephonyData(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyData {
	return &_TelephonyData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

// Deprecated: use the interface for direct cast
func CastTelephonyData(structType any) TelephonyData {
	if casted, ok := structType.(TelephonyData); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyData); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyData) GetTypeName() string {
	return "TelephonyData"
}

func (m *_TelephonyData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataParse(ctx context.Context, theBytes []byte) (TelephonyData, error) {
	return TelephonyDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TelephonyDataParseWithBufferProducer[T TelephonyData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := TelephonyDataParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func TelephonyDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TelephonyData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsTelephonyCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[TelephonyCommandTypeContainer](ctx, "commandTypeContainer", "TelephonyCommandTypeContainer", ReadEnum(TelephonyCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}

	commandType, err := ReadVirtualField[TelephonyCommandType](ctx, "commandType", (*TelephonyCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	argument, err := ReadSimpleField(ctx, "argument", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'argument' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type TelephonyDataChildSerializeRequirement interface {
		TelephonyData
		InitializeParent(TelephonyData, TelephonyCommandTypeContainer, byte)
		GetParent() TelephonyData
	}
	var _childTemp any
	var _child TelephonyDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == TelephonyCommandType_EVENT && argument == 0x01: // TelephonyDataLineOnHook
		_childTemp, typeSwitchError = TelephonyDataLineOnHookParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x02: // TelephonyDataLineOffHook
		_childTemp, typeSwitchError = TelephonyDataLineOffHookParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x03: // TelephonyDataDialOutFailure
		_childTemp, typeSwitchError = TelephonyDataDialOutFailureParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x04: // TelephonyDataDialInFailure
		_childTemp, typeSwitchError = TelephonyDataDialInFailureParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x05: // TelephonyDataRinging
		_childTemp, typeSwitchError = TelephonyDataRingingParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x06: // TelephonyDataRecallLastNumber
		_childTemp, typeSwitchError = TelephonyDataRecallLastNumberParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x07: // TelephonyDataInternetConnectionRequestMade
		_childTemp, typeSwitchError = TelephonyDataInternetConnectionRequestMadeParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x80: // TelephonyDataIsolateSecondaryOutlet
		_childTemp, typeSwitchError = TelephonyDataIsolateSecondaryOutletParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x81: // TelephonyDataRecallLastNumberRequest
		_childTemp, typeSwitchError = TelephonyDataRecallLastNumberRequestParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x82: // TelephonyDataRejectIncomingCall
		_childTemp, typeSwitchError = TelephonyDataRejectIncomingCallParseWithBuffer(ctx, readBuffer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x83: // TelephonyDataDivert
		_childTemp, typeSwitchError = TelephonyDataDivertParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == TelephonyCommandType_EVENT && argument == 0x84: // TelephonyDataClearDiversion
		_childTemp, typeSwitchError = TelephonyDataClearDiversionParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of TelephonyData")
	}
	_child = _childTemp.(TelephonyDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("TelephonyData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, argument)
	return _child, nil
}

func (pm *_TelephonyData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TelephonyData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TelephonyData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TelephonyData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (argument)
	argument := byte(m.GetArgument())
	_argumentErr := /*TODO: migrate me*/ writeBuffer.WriteByte("argument", (argument))
	if _argumentErr != nil {
		return errors.Wrap(_argumentErr, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("TelephonyData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TelephonyData")
	}
	return nil
}

func (m *_TelephonyData) isTelephonyData() bool {
	return true
}

func (m *_TelephonyData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
