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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// Constant values.
const AlarmMessageObjectAckType_VARIABLESPEC uint8 = 0x12
const AlarmMessageObjectAckType_LENGTH uint8 = 0x08

// AlarmMessageObjectAckType is the corresponding interface of AlarmMessageObjectAckType
type AlarmMessageObjectAckType interface {
	utils.LengthAware
	utils.Serializable
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
}

// AlarmMessageObjectAckTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessageObjectAckType.
// This is useful for switch cases.
type AlarmMessageObjectAckTypeExactly interface {
	AlarmMessageObjectAckType
	isAlarmMessageObjectAckType() bool
}

// _AlarmMessageObjectAckType is the data-structure of this message
type _AlarmMessageObjectAckType struct {
        SyntaxId SyntaxIdType
        NumberOfValues uint8
        EventId uint32
        AckStateGoing State
        AckStateComing State
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageObjectAckType) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_AlarmMessageObjectAckType) GetNumberOfValues() uint8 {
	return m.NumberOfValues
}

func (m *_AlarmMessageObjectAckType) GetEventId() uint32 {
	return m.EventId
}

func (m *_AlarmMessageObjectAckType) GetAckStateGoing() State {
	return m.AckStateGoing
}

func (m *_AlarmMessageObjectAckType) GetAckStateComing() State {
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

func (m *_AlarmMessageObjectAckType) GetVariableSpec() uint8 {
	return AlarmMessageObjectAckType_VARIABLESPEC
}

func (m *_AlarmMessageObjectAckType) GetLength() uint8 {
	return AlarmMessageObjectAckType_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAlarmMessageObjectAckType factory function for _AlarmMessageObjectAckType
func NewAlarmMessageObjectAckType( syntaxId SyntaxIdType , numberOfValues uint8 , eventId uint32 , ackStateGoing State , ackStateComing State ) *_AlarmMessageObjectAckType {
return &_AlarmMessageObjectAckType{ SyntaxId: syntaxId , NumberOfValues: numberOfValues , EventId: eventId , AckStateGoing: ackStateGoing , AckStateComing: ackStateComing }
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageObjectAckType(structType interface{}) AlarmMessageObjectAckType {
    if casted, ok := structType.(AlarmMessageObjectAckType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageObjectAckType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageObjectAckType) GetTypeName() string {
	return "AlarmMessageObjectAckType"
}

func (m *_AlarmMessageObjectAckType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (variableSpec)
	lengthInBits += 8

	// Const Field (length)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Simple field (numberOfValues)
	lengthInBits += 8;

	// Simple field (eventId)
	lengthInBits += 32;

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits(ctx)

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_AlarmMessageObjectAckType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageObjectAckTypeParse(theBytes []byte) (AlarmMessageObjectAckType, error) {
	return AlarmMessageObjectAckTypeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageObjectAckTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectAckType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageObjectAckType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageObjectAckType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field of AlarmMessageObjectAckType")
	}
	if variableSpec != AlarmMessageObjectAckType_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageObjectAckType_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Const Field (length)
	length, _lengthErr := readBuffer.ReadUint8("length", 8)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AlarmMessageObjectAckType")
	}
	if length != AlarmMessageObjectAckType_LENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageObjectAckType_LENGTH) + " but got " + fmt.Sprintf("%d", length))
	}

	// Simple Field (syntaxId)
	if pullErr := readBuffer.PullContext("syntaxId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for syntaxId")
	}
_syntaxId, _syntaxIdErr := SyntaxIdTypeParseWithBuffer(ctx, readBuffer)
	if _syntaxIdErr != nil {
		return nil, errors.Wrap(_syntaxIdErr, "Error parsing 'syntaxId' field of AlarmMessageObjectAckType")
	}
	syntaxId := _syntaxId
	if closeErr := readBuffer.CloseContext("syntaxId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for syntaxId")
	}

	// Simple Field (numberOfValues)
_numberOfValues, _numberOfValuesErr := readBuffer.ReadUint8("numberOfValues", 8)
	if _numberOfValuesErr != nil {
		return nil, errors.Wrap(_numberOfValuesErr, "Error parsing 'numberOfValues' field of AlarmMessageObjectAckType")
	}
	numberOfValues := _numberOfValues

	// Simple Field (eventId)
_eventId, _eventIdErr := readBuffer.ReadUint32("eventId", 32)
	if _eventIdErr != nil {
		return nil, errors.Wrap(_eventIdErr, "Error parsing 'eventId' field of AlarmMessageObjectAckType")
	}
	eventId := _eventId

	// Simple Field (ackStateGoing)
	if pullErr := readBuffer.PullContext("ackStateGoing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateGoing")
	}
_ackStateGoing, _ackStateGoingErr := StateParseWithBuffer(ctx, readBuffer)
	if _ackStateGoingErr != nil {
		return nil, errors.Wrap(_ackStateGoingErr, "Error parsing 'ackStateGoing' field of AlarmMessageObjectAckType")
	}
	ackStateGoing := _ackStateGoing.(State)
	if closeErr := readBuffer.CloseContext("ackStateGoing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateGoing")
	}

	// Simple Field (ackStateComing)
	if pullErr := readBuffer.PullContext("ackStateComing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateComing")
	}
_ackStateComing, _ackStateComingErr := StateParseWithBuffer(ctx, readBuffer)
	if _ackStateComingErr != nil {
		return nil, errors.Wrap(_ackStateComingErr, "Error parsing 'ackStateComing' field of AlarmMessageObjectAckType")
	}
	ackStateComing := _ackStateComing.(State)
	if closeErr := readBuffer.CloseContext("ackStateComing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateComing")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectAckType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageObjectAckType")
	}

	// Create the instance
	return &_AlarmMessageObjectAckType{
			SyntaxId: syntaxId,
			NumberOfValues: numberOfValues,
			EventId: eventId,
			AckStateGoing: ackStateGoing,
			AckStateComing: ackStateComing,
		}, nil
}

func (m *_AlarmMessageObjectAckType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageObjectAckType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("AlarmMessageObjectAckType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageObjectAckType")
	}

	// Const Field (variableSpec)
	_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, 0x12)
	if _variableSpecErr != nil {
		return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
	}

	// Const Field (length)
	_lengthErr := writeBuffer.WriteUint8("length", 8, 0x08)
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (syntaxId)
	if pushErr := writeBuffer.PushContext("syntaxId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for syntaxId")
	}
	_syntaxIdErr := writeBuffer.WriteSerializable(ctx, m.GetSyntaxId())
	if popErr := writeBuffer.PopContext("syntaxId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for syntaxId")
	}
	if _syntaxIdErr != nil {
		return errors.Wrap(_syntaxIdErr, "Error serializing 'syntaxId' field")
	}

	// Simple Field (numberOfValues)
	numberOfValues := uint8(m.GetNumberOfValues())
	_numberOfValuesErr := writeBuffer.WriteUint8("numberOfValues", 8, (numberOfValues))
	if _numberOfValuesErr != nil {
		return errors.Wrap(_numberOfValuesErr, "Error serializing 'numberOfValues' field")
	}

	// Simple Field (eventId)
	eventId := uint32(m.GetEventId())
	_eventIdErr := writeBuffer.WriteUint32("eventId", 32, (eventId))
	if _eventIdErr != nil {
		return errors.Wrap(_eventIdErr, "Error serializing 'eventId' field")
	}

	// Simple Field (ackStateGoing)
	if pushErr := writeBuffer.PushContext("ackStateGoing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateGoing")
	}
	_ackStateGoingErr := writeBuffer.WriteSerializable(ctx, m.GetAckStateGoing())
	if popErr := writeBuffer.PopContext("ackStateGoing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateGoing")
	}
	if _ackStateGoingErr != nil {
		return errors.Wrap(_ackStateGoingErr, "Error serializing 'ackStateGoing' field")
	}

	// Simple Field (ackStateComing)
	if pushErr := writeBuffer.PushContext("ackStateComing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateComing")
	}
	_ackStateComingErr := writeBuffer.WriteSerializable(ctx, m.GetAckStateComing())
	if popErr := writeBuffer.PopContext("ackStateComing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateComing")
	}
	if _ackStateComingErr != nil {
		return errors.Wrap(_ackStateComingErr, "Error serializing 'ackStateComing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectAckType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageObjectAckType")
	}
	return nil
}


func (m *_AlarmMessageObjectAckType) isAlarmMessageObjectAckType() bool {
	return true
}

func (m *_AlarmMessageObjectAckType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



