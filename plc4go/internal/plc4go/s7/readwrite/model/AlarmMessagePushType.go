/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AlarmMessagePushType struct {
	TimeStamp       *DateAndTime
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []*AlarmMessageObjectPushType
}

// The corresponding interface
type IAlarmMessagePushType interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewAlarmMessagePushType(TimeStamp *DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []*AlarmMessageObjectPushType) *AlarmMessagePushType {
	return &AlarmMessagePushType{TimeStamp: TimeStamp, FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

func CastAlarmMessagePushType(structType interface{}) *AlarmMessagePushType {
	castFunc := func(typ interface{}) *AlarmMessagePushType {
		if casted, ok := typ.(AlarmMessagePushType); ok {
			return &casted
		}
		if casted, ok := typ.(*AlarmMessagePushType); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AlarmMessagePushType) GetTypeName() string {
	return "AlarmMessagePushType"
}

func (m *AlarmMessagePushType) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AlarmMessagePushType) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (TimeStamp)
	lengthInBits += m.TimeStamp.LengthInBits()

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *AlarmMessagePushType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AlarmMessagePushTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessagePushType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessagePushType"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (TimeStamp)
	if pullErr := readBuffer.PullContext("TimeStamp"); pullErr != nil {
		return nil, pullErr
	}
	TimeStamp, _TimeStampErr := DateAndTimeParse(readBuffer)
	if _TimeStampErr != nil {
		return nil, errors.Wrap(_TimeStampErr, "Error parsing 'TimeStamp' field")
	}
	if closeErr := readBuffer.CloseContext("TimeStamp"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (functionId)
	functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field")
	}

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	messageObjects := make([]*AlarmMessageObjectPushType, numberOfObjects)
	for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
		_item, _err := AlarmMessageObjectPushTypeParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
		}
		messageObjects[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessagePushType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessagePushType(TimeStamp, functionId, numberOfObjects, messageObjects), nil
}

func (m *AlarmMessagePushType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessagePushType"); pushErr != nil {
		return pushErr
	}

	// Simple Field (TimeStamp)
	if pushErr := writeBuffer.PushContext("TimeStamp"); pushErr != nil {
		return pushErr
	}
	_TimeStampErr := m.TimeStamp.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("TimeStamp"); popErr != nil {
		return popErr
	}
	if _TimeStampErr != nil {
		return errors.Wrap(_TimeStampErr, "Error serializing 'TimeStamp' field")
	}

	// Simple Field (functionId)
	functionId := uint8(m.FunctionId)
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.NumberOfObjects)
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if m.MessageObjects != nil {
		if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.MessageObjects {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
			}
		}
		if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AlarmMessagePushType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessagePushType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
