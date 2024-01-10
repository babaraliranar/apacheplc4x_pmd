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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// AlarmMessagePushType is the corresponding interface of AlarmMessagePushType
type AlarmMessagePushType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectPushType
}

// AlarmMessagePushTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessagePushType.
// This is useful for switch cases.
type AlarmMessagePushTypeExactly interface {
	AlarmMessagePushType
	isAlarmMessagePushType() bool
}

// _AlarmMessagePushType is the data-structure of this message
type _AlarmMessagePushType struct {
	TimeStamp       DateAndTime
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []AlarmMessageObjectPushType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessagePushType) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

func (m *_AlarmMessagePushType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessagePushType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessagePushType) GetMessageObjects() []AlarmMessageObjectPushType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessagePushType factory function for _AlarmMessagePushType
func NewAlarmMessagePushType(TimeStamp DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectPushType) *_AlarmMessagePushType {
	return &_AlarmMessagePushType{TimeStamp: TimeStamp, FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessagePushType(structType any) AlarmMessagePushType {
	if casted, ok := structType.(AlarmMessagePushType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessagePushType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessagePushType) GetTypeName() string {
	return "AlarmMessagePushType"
}

func (m *_AlarmMessagePushType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (TimeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessagePushType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessagePushTypeParse(ctx context.Context, theBytes []byte) (AlarmMessagePushType, error) {
	return AlarmMessagePushTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessagePushTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessagePushType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AlarmMessagePushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessagePushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (TimeStamp)
	if pullErr := readBuffer.PullContext("TimeStamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TimeStamp")
	}
	_TimeStamp, _TimeStampErr := DateAndTimeParseWithBuffer(ctx, readBuffer)
	if _TimeStampErr != nil {
		return nil, errors.Wrap(_TimeStampErr, "Error parsing 'TimeStamp' field of AlarmMessagePushType")
	}
	TimeStamp := _TimeStamp.(DateAndTime)
	if closeErr := readBuffer.CloseContext("TimeStamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TimeStamp")
	}

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field of AlarmMessagePushType")
	}
	functionId := _functionId

	// Simple Field (numberOfObjects)
	_numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field of AlarmMessagePushType")
	}
	numberOfObjects := _numberOfObjects

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageObjects")
	}
	// Count array
	messageObjects := make([]AlarmMessageObjectPushType, utils.Max(numberOfObjects, 0))
	// This happens when the size is set conditional to 0
	if len(messageObjects) == 0 {
		messageObjects = nil
	}
	{
		_numItems := uint16(utils.Max(numberOfObjects, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := AlarmMessageObjectPushTypeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field of AlarmMessagePushType")
			}
			messageObjects[_curItem] = _item.(AlarmMessageObjectPushType)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageObjects")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessagePushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessagePushType")
	}

	// Create the instance
	return &_AlarmMessagePushType{
		TimeStamp:       TimeStamp,
		FunctionId:      functionId,
		NumberOfObjects: numberOfObjects,
		MessageObjects:  messageObjects,
	}, nil
}

func (m *_AlarmMessagePushType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessagePushType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessagePushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessagePushType")
	}

	// Simple Field (TimeStamp)
	if pushErr := writeBuffer.PushContext("TimeStamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TimeStamp")
	}
	_TimeStampErr := writeBuffer.WriteSerializable(ctx, m.GetTimeStamp())
	if popErr := writeBuffer.PopContext("TimeStamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TimeStamp")
	}
	if _TimeStampErr != nil {
		return errors.Wrap(_TimeStampErr, "Error serializing 'TimeStamp' field")
	}

	// Simple Field (functionId)
	functionId := uint8(m.GetFunctionId())
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, uint8((functionId)))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.GetNumberOfObjects())
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, uint8((numberOfObjects)))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for messageObjects")
	}
	for _curItem, _element := range m.GetMessageObjects() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMessageObjects()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
		}
	}
	if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for messageObjects")
	}

	if popErr := writeBuffer.PopContext("AlarmMessagePushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessagePushType")
	}
	return nil
}

func (m *_AlarmMessagePushType) isAlarmMessagePushType() bool {
	return true
}

func (m *_AlarmMessagePushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
