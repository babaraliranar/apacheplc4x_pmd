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

// ErrorReportingData is the corresponding interface of ErrorReportingData
type ErrorReportingData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() ErrorReportingCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() ErrorReportingCommandType
}

// ErrorReportingDataExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingData.
// This is useful for switch cases.
type ErrorReportingDataExactly interface {
	ErrorReportingData
	isErrorReportingData() bool
}

// _ErrorReportingData is the data-structure of this message
type _ErrorReportingData struct {
	_ErrorReportingDataChildRequirements
	CommandTypeContainer ErrorReportingCommandTypeContainer
}

type _ErrorReportingDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() ErrorReportingCommandType
}

type ErrorReportingDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ErrorReportingData, serializeChildFunction func() error) error
	GetTypeName() string
}

type ErrorReportingDataChild interface {
	utils.Serializable
	InitializeParent(parent ErrorReportingData, commandTypeContainer ErrorReportingCommandTypeContainer)
	GetParent() *ErrorReportingData

	GetTypeName() string
	ErrorReportingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingData) GetCommandTypeContainer() ErrorReportingCommandTypeContainer {
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

func (m *_ErrorReportingData) GetCommandType() ErrorReportingCommandType {
	ctx := context.Background()
	_ = ctx
	return CastErrorReportingCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingData factory function for _ErrorReportingData
func NewErrorReportingData(commandTypeContainer ErrorReportingCommandTypeContainer) *_ErrorReportingData {
	return &_ErrorReportingData{CommandTypeContainer: commandTypeContainer}
}

// Deprecated: use the interface for direct cast
func CastErrorReportingData(structType any) ErrorReportingData {
	if casted, ok := structType.(ErrorReportingData); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingData); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingData) GetTypeName() string {
	return "ErrorReportingData"
}

func (m *_ErrorReportingData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ErrorReportingData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingDataParse(ctx context.Context, theBytes []byte) (ErrorReportingData, error) {
	return ErrorReportingDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingDataParseWithBufferProducer[T ErrorReportingData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := ErrorReportingDataParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func ErrorReportingDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ErrorReportingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsErrorReportingCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[ErrorReportingCommandTypeContainer](ctx, "commandTypeContainer", "ErrorReportingCommandTypeContainer", ReadEnum(ErrorReportingCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}

	commandType, err := ReadVirtualField[ErrorReportingCommandType](ctx, "commandType", (*ErrorReportingCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ErrorReportingDataChildSerializeRequirement interface {
		ErrorReportingData
		InitializeParent(ErrorReportingData, ErrorReportingCommandTypeContainer)
		GetParent() ErrorReportingData
	}
	var _childTemp any
	var _child ErrorReportingDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case 0 == 0: // ErrorReportingDataGeneric
		_childTemp, typeSwitchError = ErrorReportingDataGenericParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ErrorReportingData")
	}
	_child = _childTemp.(ErrorReportingDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ErrorReportingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer)
	return _child, nil
}

func (pm *_ErrorReportingData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ErrorReportingData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorReportingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorReportingData")
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

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ErrorReportingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorReportingData")
	}
	return nil
}

func (m *_ErrorReportingData) isErrorReportingData() bool {
	return true
}

func (m *_ErrorReportingData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
