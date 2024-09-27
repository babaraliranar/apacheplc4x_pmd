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

// AlarmMessageAckResponseType is the corresponding interface of AlarmMessageAckResponseType
type AlarmMessageAckResponseType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []uint8
	// IsAlarmMessageAckResponseType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarmMessageAckResponseType()
	// CreateBuilder creates a AlarmMessageAckResponseTypeBuilder
	CreateAlarmMessageAckResponseTypeBuilder() AlarmMessageAckResponseTypeBuilder
}

// _AlarmMessageAckResponseType is the data-structure of this message
type _AlarmMessageAckResponseType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []uint8
}

var _ AlarmMessageAckResponseType = (*_AlarmMessageAckResponseType)(nil)

// NewAlarmMessageAckResponseType factory function for _AlarmMessageAckResponseType
func NewAlarmMessageAckResponseType(functionId uint8, numberOfObjects uint8, messageObjects []uint8) *_AlarmMessageAckResponseType {
	return &_AlarmMessageAckResponseType{FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AlarmMessageAckResponseTypeBuilder is a builder for AlarmMessageAckResponseType
type AlarmMessageAckResponseTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(functionId uint8, numberOfObjects uint8, messageObjects []uint8) AlarmMessageAckResponseTypeBuilder
	// WithFunctionId adds FunctionId (property field)
	WithFunctionId(uint8) AlarmMessageAckResponseTypeBuilder
	// WithNumberOfObjects adds NumberOfObjects (property field)
	WithNumberOfObjects(uint8) AlarmMessageAckResponseTypeBuilder
	// WithMessageObjects adds MessageObjects (property field)
	WithMessageObjects(...uint8) AlarmMessageAckResponseTypeBuilder
	// Build builds the AlarmMessageAckResponseType or returns an error if something is wrong
	Build() (AlarmMessageAckResponseType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AlarmMessageAckResponseType
}

// NewAlarmMessageAckResponseTypeBuilder() creates a AlarmMessageAckResponseTypeBuilder
func NewAlarmMessageAckResponseTypeBuilder() AlarmMessageAckResponseTypeBuilder {
	return &_AlarmMessageAckResponseTypeBuilder{_AlarmMessageAckResponseType: new(_AlarmMessageAckResponseType)}
}

type _AlarmMessageAckResponseTypeBuilder struct {
	*_AlarmMessageAckResponseType

	err *utils.MultiError
}

var _ (AlarmMessageAckResponseTypeBuilder) = (*_AlarmMessageAckResponseTypeBuilder)(nil)

func (m *_AlarmMessageAckResponseTypeBuilder) WithMandatoryFields(functionId uint8, numberOfObjects uint8, messageObjects []uint8) AlarmMessageAckResponseTypeBuilder {
	return m.WithFunctionId(functionId).WithNumberOfObjects(numberOfObjects).WithMessageObjects(messageObjects...)
}

func (m *_AlarmMessageAckResponseTypeBuilder) WithFunctionId(functionId uint8) AlarmMessageAckResponseTypeBuilder {
	m.FunctionId = functionId
	return m
}

func (m *_AlarmMessageAckResponseTypeBuilder) WithNumberOfObjects(numberOfObjects uint8) AlarmMessageAckResponseTypeBuilder {
	m.NumberOfObjects = numberOfObjects
	return m
}

func (m *_AlarmMessageAckResponseTypeBuilder) WithMessageObjects(messageObjects ...uint8) AlarmMessageAckResponseTypeBuilder {
	m.MessageObjects = messageObjects
	return m
}

func (m *_AlarmMessageAckResponseTypeBuilder) Build() (AlarmMessageAckResponseType, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AlarmMessageAckResponseType.deepCopy(), nil
}

func (m *_AlarmMessageAckResponseTypeBuilder) MustBuild() AlarmMessageAckResponseType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AlarmMessageAckResponseTypeBuilder) DeepCopy() any {
	return m.CreateAlarmMessageAckResponseTypeBuilder()
}

// CreateAlarmMessageAckResponseTypeBuilder creates a AlarmMessageAckResponseTypeBuilder
func (m *_AlarmMessageAckResponseType) CreateAlarmMessageAckResponseTypeBuilder() AlarmMessageAckResponseTypeBuilder {
	if m == nil {
		return NewAlarmMessageAckResponseTypeBuilder()
	}
	return &_AlarmMessageAckResponseTypeBuilder{_AlarmMessageAckResponseType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageAckResponseType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessageAckResponseType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessageAckResponseType) GetMessageObjects() []uint8 {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAlarmMessageAckResponseType(structType any) AlarmMessageAckResponseType {
	if casted, ok := structType.(AlarmMessageAckResponseType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageAckResponseType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageAckResponseType) GetTypeName() string {
	return "AlarmMessageAckResponseType"
}

func (m *_AlarmMessageAckResponseType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		lengthInBits += 8 * uint16(len(m.MessageObjects))
	}

	return lengthInBits
}

func (m *_AlarmMessageAckResponseType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageAckResponseTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageAckResponseType, error) {
	return AlarmMessageAckResponseTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageAckResponseTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckResponseType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckResponseType, error) {
		return AlarmMessageAckResponseTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageAckResponseTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckResponseType, error) {
	v, err := (&_AlarmMessageAckResponseType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AlarmMessageAckResponseType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessageAckResponseType AlarmMessageAckResponseType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageAckResponseType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckResponseType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadSimpleField(ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	m.FunctionId = functionId

	numberOfObjects, err := ReadSimpleField(ctx, "numberOfObjects", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfObjects' field"))
	}
	m.NumberOfObjects = numberOfObjects

	messageObjects, err := ReadCountArrayField[uint8](ctx, "messageObjects", ReadUnsignedByte(readBuffer, uint8(8)), uint64(numberOfObjects))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageObjects' field"))
	}
	m.MessageObjects = messageObjects

	if closeErr := readBuffer.CloseContext("AlarmMessageAckResponseType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckResponseType")
	}

	return m, nil
}

func (m *_AlarmMessageAckResponseType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageAckResponseType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageAckResponseType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckResponseType")
	}

	if err := WriteSimpleField[uint8](ctx, "functionId", m.GetFunctionId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "numberOfObjects", m.GetNumberOfObjects(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfObjects' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "messageObjects", m.GetMessageObjects(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageObjects' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckResponseType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckResponseType")
	}
	return nil
}

func (m *_AlarmMessageAckResponseType) IsAlarmMessageAckResponseType() {}

func (m *_AlarmMessageAckResponseType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AlarmMessageAckResponseType) deepCopy() *_AlarmMessageAckResponseType {
	if m == nil {
		return nil
	}
	_AlarmMessageAckResponseTypeCopy := &_AlarmMessageAckResponseType{
		m.FunctionId,
		m.NumberOfObjects,
		utils.DeepCopySlice[uint8, uint8](m.MessageObjects),
	}
	return _AlarmMessageAckResponseTypeCopy
}

func (m *_AlarmMessageAckResponseType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
