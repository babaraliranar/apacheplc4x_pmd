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

// Constant values.
const AlarmMessageAckObjectPushType_VARIABLESPEC uint8 = 0x12

// AlarmMessageAckObjectPushType is the corresponding interface of AlarmMessageAckObjectPushType
type AlarmMessageAckObjectPushType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetLengthSpec returns LengthSpec (property field)
	GetLengthSpec() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetNumberOfValues returns NumberOfValues (property field)
	GetNumberOfValues() uint8
	// GetEventId returns EventId (property field)
	GetEventId() uint32
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() State
	// IsAlarmMessageAckObjectPushType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarmMessageAckObjectPushType()
	// CreateBuilder creates a AlarmMessageAckObjectPushTypeBuilder
	CreateAlarmMessageAckObjectPushTypeBuilder() AlarmMessageAckObjectPushTypeBuilder
}

// _AlarmMessageAckObjectPushType is the data-structure of this message
type _AlarmMessageAckObjectPushType struct {
	LengthSpec     uint8
	SyntaxId       SyntaxIdType
	NumberOfValues uint8
	EventId        uint32
	AckStateGoing  State
	AckStateComing State
}

var _ AlarmMessageAckObjectPushType = (*_AlarmMessageAckObjectPushType)(nil)

// NewAlarmMessageAckObjectPushType factory function for _AlarmMessageAckObjectPushType
func NewAlarmMessageAckObjectPushType(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, ackStateGoing State, ackStateComing State) *_AlarmMessageAckObjectPushType {
	if ackStateGoing == nil {
		panic("ackStateGoing of type State for AlarmMessageAckObjectPushType must not be nil")
	}
	if ackStateComing == nil {
		panic("ackStateComing of type State for AlarmMessageAckObjectPushType must not be nil")
	}
	return &_AlarmMessageAckObjectPushType{LengthSpec: lengthSpec, SyntaxId: syntaxId, NumberOfValues: numberOfValues, EventId: eventId, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AlarmMessageAckObjectPushTypeBuilder is a builder for AlarmMessageAckObjectPushType
type AlarmMessageAckObjectPushTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, ackStateGoing State, ackStateComing State) AlarmMessageAckObjectPushTypeBuilder
	// WithLengthSpec adds LengthSpec (property field)
	WithLengthSpec(uint8) AlarmMessageAckObjectPushTypeBuilder
	// WithSyntaxId adds SyntaxId (property field)
	WithSyntaxId(SyntaxIdType) AlarmMessageAckObjectPushTypeBuilder
	// WithNumberOfValues adds NumberOfValues (property field)
	WithNumberOfValues(uint8) AlarmMessageAckObjectPushTypeBuilder
	// WithEventId adds EventId (property field)
	WithEventId(uint32) AlarmMessageAckObjectPushTypeBuilder
	// WithAckStateGoing adds AckStateGoing (property field)
	WithAckStateGoing(State) AlarmMessageAckObjectPushTypeBuilder
	// WithAckStateGoingBuilder adds AckStateGoing (property field) which is build by the builder
	WithAckStateGoingBuilder(func(StateBuilder) StateBuilder) AlarmMessageAckObjectPushTypeBuilder
	// WithAckStateComing adds AckStateComing (property field)
	WithAckStateComing(State) AlarmMessageAckObjectPushTypeBuilder
	// WithAckStateComingBuilder adds AckStateComing (property field) which is build by the builder
	WithAckStateComingBuilder(func(StateBuilder) StateBuilder) AlarmMessageAckObjectPushTypeBuilder
	// Build builds the AlarmMessageAckObjectPushType or returns an error if something is wrong
	Build() (AlarmMessageAckObjectPushType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AlarmMessageAckObjectPushType
}

// NewAlarmMessageAckObjectPushTypeBuilder() creates a AlarmMessageAckObjectPushTypeBuilder
func NewAlarmMessageAckObjectPushTypeBuilder() AlarmMessageAckObjectPushTypeBuilder {
	return &_AlarmMessageAckObjectPushTypeBuilder{_AlarmMessageAckObjectPushType: new(_AlarmMessageAckObjectPushType)}
}

type _AlarmMessageAckObjectPushTypeBuilder struct {
	*_AlarmMessageAckObjectPushType

	err *utils.MultiError
}

var _ (AlarmMessageAckObjectPushTypeBuilder) = (*_AlarmMessageAckObjectPushTypeBuilder)(nil)

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithMandatoryFields(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, ackStateGoing State, ackStateComing State) AlarmMessageAckObjectPushTypeBuilder {
	return m.WithLengthSpec(lengthSpec).WithSyntaxId(syntaxId).WithNumberOfValues(numberOfValues).WithEventId(eventId).WithAckStateGoing(ackStateGoing).WithAckStateComing(ackStateComing)
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithLengthSpec(lengthSpec uint8) AlarmMessageAckObjectPushTypeBuilder {
	m.LengthSpec = lengthSpec
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithSyntaxId(syntaxId SyntaxIdType) AlarmMessageAckObjectPushTypeBuilder {
	m.SyntaxId = syntaxId
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithNumberOfValues(numberOfValues uint8) AlarmMessageAckObjectPushTypeBuilder {
	m.NumberOfValues = numberOfValues
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithEventId(eventId uint32) AlarmMessageAckObjectPushTypeBuilder {
	m.EventId = eventId
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithAckStateGoing(ackStateGoing State) AlarmMessageAckObjectPushTypeBuilder {
	m.AckStateGoing = ackStateGoing
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithAckStateGoingBuilder(builderSupplier func(StateBuilder) StateBuilder) AlarmMessageAckObjectPushTypeBuilder {
	builder := builderSupplier(m.AckStateGoing.CreateStateBuilder())
	var err error
	m.AckStateGoing, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "StateBuilder failed"))
	}
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithAckStateComing(ackStateComing State) AlarmMessageAckObjectPushTypeBuilder {
	m.AckStateComing = ackStateComing
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) WithAckStateComingBuilder(builderSupplier func(StateBuilder) StateBuilder) AlarmMessageAckObjectPushTypeBuilder {
	builder := builderSupplier(m.AckStateComing.CreateStateBuilder())
	var err error
	m.AckStateComing, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "StateBuilder failed"))
	}
	return m
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) Build() (AlarmMessageAckObjectPushType, error) {
	if m.AckStateGoing == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'ackStateGoing' not set"))
	}
	if m.AckStateComing == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'ackStateComing' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AlarmMessageAckObjectPushType.deepCopy(), nil
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) MustBuild() AlarmMessageAckObjectPushType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AlarmMessageAckObjectPushTypeBuilder) DeepCopy() any {
	return m.CreateAlarmMessageAckObjectPushTypeBuilder()
}

// CreateAlarmMessageAckObjectPushTypeBuilder creates a AlarmMessageAckObjectPushTypeBuilder
func (m *_AlarmMessageAckObjectPushType) CreateAlarmMessageAckObjectPushTypeBuilder() AlarmMessageAckObjectPushTypeBuilder {
	if m == nil {
		return NewAlarmMessageAckObjectPushTypeBuilder()
	}
	return &_AlarmMessageAckObjectPushTypeBuilder{_AlarmMessageAckObjectPushType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageAckObjectPushType) GetLengthSpec() uint8 {
	return m.LengthSpec
}

func (m *_AlarmMessageAckObjectPushType) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_AlarmMessageAckObjectPushType) GetNumberOfValues() uint8 {
	return m.NumberOfValues
}

func (m *_AlarmMessageAckObjectPushType) GetEventId() uint32 {
	return m.EventId
}

func (m *_AlarmMessageAckObjectPushType) GetAckStateGoing() State {
	return m.AckStateGoing
}

func (m *_AlarmMessageAckObjectPushType) GetAckStateComing() State {
	return m.AckStateComing
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AlarmMessageAckObjectPushType) GetVariableSpec() uint8 {
	return AlarmMessageAckObjectPushType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAlarmMessageAckObjectPushType(structType any) AlarmMessageAckObjectPushType {
	if casted, ok := structType.(AlarmMessageAckObjectPushType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageAckObjectPushType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageAckObjectPushType) GetTypeName() string {
	return "AlarmMessageAckObjectPushType"
}

func (m *_AlarmMessageAckObjectPushType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (lengthSpec)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Simple field (numberOfValues)
	lengthInBits += 8

	// Simple field (eventId)
	lengthInBits += 32

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits(ctx)

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AlarmMessageAckObjectPushType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageAckObjectPushTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageAckObjectPushType, error) {
	return AlarmMessageAckObjectPushTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageAckObjectPushTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckObjectPushType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckObjectPushType, error) {
		return AlarmMessageAckObjectPushTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageAckObjectPushTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckObjectPushType, error) {
	v, err := (&_AlarmMessageAckObjectPushType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AlarmMessageAckObjectPushType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessageAckObjectPushType AlarmMessageAckObjectPushType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageAckObjectPushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckObjectPushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), AlarmMessageAckObjectPushType_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	lengthSpec, err := ReadSimpleField(ctx, "lengthSpec", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lengthSpec' field"))
	}
	m.LengthSpec = lengthSpec

	syntaxId, err := ReadEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", ReadEnum(SyntaxIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}
	m.SyntaxId = syntaxId

	numberOfValues, err := ReadSimpleField(ctx, "numberOfValues", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfValues' field"))
	}
	m.NumberOfValues = numberOfValues

	eventId, err := ReadSimpleField(ctx, "eventId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventId' field"))
	}
	m.EventId = eventId

	ackStateGoing, err := ReadSimpleField[State](ctx, "ackStateGoing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateGoing' field"))
	}
	m.AckStateGoing = ackStateGoing

	ackStateComing, err := ReadSimpleField[State](ctx, "ackStateComing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateComing' field"))
	}
	m.AckStateComing = ackStateComing

	if closeErr := readBuffer.CloseContext("AlarmMessageAckObjectPushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckObjectPushType")
	}

	return m, nil
}

func (m *_AlarmMessageAckObjectPushType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageAckObjectPushType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageAckObjectPushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckObjectPushType")
	}

	if err := WriteConstField(ctx, "variableSpec", AlarmMessageAckObjectPushType_VARIABLESPEC, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'variableSpec' field")
	}

	if err := WriteSimpleField[uint8](ctx, "lengthSpec", m.GetLengthSpec(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'lengthSpec' field")
	}

	if err := WriteSimpleEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", m.GetSyntaxId(), WriteEnum[SyntaxIdType, uint8](SyntaxIdType.GetValue, SyntaxIdType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'syntaxId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "numberOfValues", m.GetNumberOfValues(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfValues' field")
	}

	if err := WriteSimpleField[uint32](ctx, "eventId", m.GetEventId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventId' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateGoing", m.GetAckStateGoing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateGoing' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateComing", m.GetAckStateComing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateComing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckObjectPushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckObjectPushType")
	}
	return nil
}

func (m *_AlarmMessageAckObjectPushType) IsAlarmMessageAckObjectPushType() {}

func (m *_AlarmMessageAckObjectPushType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AlarmMessageAckObjectPushType) deepCopy() *_AlarmMessageAckObjectPushType {
	if m == nil {
		return nil
	}
	_AlarmMessageAckObjectPushTypeCopy := &_AlarmMessageAckObjectPushType{
		m.LengthSpec,
		m.SyntaxId,
		m.NumberOfValues,
		m.EventId,
		m.AckStateGoing.DeepCopy().(State),
		m.AckStateComing.DeepCopy().(State),
	}
	return _AlarmMessageAckObjectPushTypeCopy
}

func (m *_AlarmMessageAckObjectPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
