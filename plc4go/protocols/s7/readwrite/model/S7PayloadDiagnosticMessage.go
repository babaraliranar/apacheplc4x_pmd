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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// S7PayloadDiagnosticMessage is the corresponding interface of S7PayloadDiagnosticMessage
type S7PayloadDiagnosticMessage interface {
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetEventId returns EventId (property field)
	GetEventId() uint16
	// GetPriorityClass returns PriorityClass (property field)
	GetPriorityClass() uint8
	// GetObNumber returns ObNumber (property field)
	GetObNumber() uint8
	// GetDatId returns DatId (property field)
	GetDatId() uint16
	// GetInfo1 returns Info1 (property field)
	GetInfo1() uint16
	// GetInfo2 returns Info2 (property field)
	GetInfo2() uint32
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
}

// S7PayloadDiagnosticMessageExactly can be used when we want exactly this type and not a type which fulfills S7PayloadDiagnosticMessage.
// This is useful for switch cases.
type S7PayloadDiagnosticMessageExactly interface {
	S7PayloadDiagnosticMessage
	isS7PayloadDiagnosticMessage() bool
}

// _S7PayloadDiagnosticMessage is the data-structure of this message
type _S7PayloadDiagnosticMessage struct {
	*_S7PayloadUserDataItem
        EventId uint16
        PriorityClass uint8
        ObNumber uint8
        DatId uint16
        Info1 uint16
        Info2 uint32
        TimeStamp DateAndTime
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadDiagnosticMessage)  GetCpuFunctionType() uint8 {
return 0x00}

func (m *_S7PayloadDiagnosticMessage)  GetCpuSubfunction() uint8 {
return 0x03}

func (m *_S7PayloadDiagnosticMessage)  GetDataLength() uint16 {
return 0}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadDiagnosticMessage) InitializeParent(parent S7PayloadUserDataItem , returnCode DataTransportErrorCode , transportSize DataTransportSize ) {	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadDiagnosticMessage)  GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadDiagnosticMessage) GetEventId() uint16 {
	return m.EventId
}

func (m *_S7PayloadDiagnosticMessage) GetPriorityClass() uint8 {
	return m.PriorityClass
}

func (m *_S7PayloadDiagnosticMessage) GetObNumber() uint8 {
	return m.ObNumber
}

func (m *_S7PayloadDiagnosticMessage) GetDatId() uint16 {
	return m.DatId
}

func (m *_S7PayloadDiagnosticMessage) GetInfo1() uint16 {
	return m.Info1
}

func (m *_S7PayloadDiagnosticMessage) GetInfo2() uint32 {
	return m.Info2
}

func (m *_S7PayloadDiagnosticMessage) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewS7PayloadDiagnosticMessage factory function for _S7PayloadDiagnosticMessage
func NewS7PayloadDiagnosticMessage( EventId uint16 , PriorityClass uint8 , ObNumber uint8 , DatId uint16 , Info1 uint16 , Info2 uint32 , TimeStamp DateAndTime , returnCode DataTransportErrorCode , transportSize DataTransportSize ) *_S7PayloadDiagnosticMessage {
	_result := &_S7PayloadDiagnosticMessage{
		EventId: EventId,
		PriorityClass: PriorityClass,
		ObNumber: ObNumber,
		DatId: DatId,
		Info1: Info1,
		Info2: Info2,
		TimeStamp: TimeStamp,
    	_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadDiagnosticMessage(structType interface{}) S7PayloadDiagnosticMessage {
    if casted, ok := structType.(S7PayloadDiagnosticMessage); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadDiagnosticMessage); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadDiagnosticMessage) GetTypeName() string {
	return "S7PayloadDiagnosticMessage"
}

func (m *_S7PayloadDiagnosticMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (EventId)
	lengthInBits += 16;

	// Simple field (PriorityClass)
	lengthInBits += 8;

	// Simple field (ObNumber)
	lengthInBits += 8;

	// Simple field (DatId)
	lengthInBits += 16;

	// Simple field (Info1)
	lengthInBits += 16;

	// Simple field (Info2)
	lengthInBits += 32;

	// Simple field (TimeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_S7PayloadDiagnosticMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadDiagnosticMessageParse(theBytes []byte, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadDiagnosticMessage, error) {
	return S7PayloadDiagnosticMessageParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cpuFunctionType, cpuSubfunction)
}

func S7PayloadDiagnosticMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadDiagnosticMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadDiagnosticMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadDiagnosticMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (EventId)
_EventId, _EventIdErr := readBuffer.ReadUint16("EventId", 16)
	if _EventIdErr != nil {
		return nil, errors.Wrap(_EventIdErr, "Error parsing 'EventId' field of S7PayloadDiagnosticMessage")
	}
	EventId := _EventId

	// Simple Field (PriorityClass)
_PriorityClass, _PriorityClassErr := readBuffer.ReadUint8("PriorityClass", 8)
	if _PriorityClassErr != nil {
		return nil, errors.Wrap(_PriorityClassErr, "Error parsing 'PriorityClass' field of S7PayloadDiagnosticMessage")
	}
	PriorityClass := _PriorityClass

	// Simple Field (ObNumber)
_ObNumber, _ObNumberErr := readBuffer.ReadUint8("ObNumber", 8)
	if _ObNumberErr != nil {
		return nil, errors.Wrap(_ObNumberErr, "Error parsing 'ObNumber' field of S7PayloadDiagnosticMessage")
	}
	ObNumber := _ObNumber

	// Simple Field (DatId)
_DatId, _DatIdErr := readBuffer.ReadUint16("DatId", 16)
	if _DatIdErr != nil {
		return nil, errors.Wrap(_DatIdErr, "Error parsing 'DatId' field of S7PayloadDiagnosticMessage")
	}
	DatId := _DatId

	// Simple Field (Info1)
_Info1, _Info1Err := readBuffer.ReadUint16("Info1", 16)
	if _Info1Err != nil {
		return nil, errors.Wrap(_Info1Err, "Error parsing 'Info1' field of S7PayloadDiagnosticMessage")
	}
	Info1 := _Info1

	// Simple Field (Info2)
_Info2, _Info2Err := readBuffer.ReadUint32("Info2", 32)
	if _Info2Err != nil {
		return nil, errors.Wrap(_Info2Err, "Error parsing 'Info2' field of S7PayloadDiagnosticMessage")
	}
	Info2 := _Info2

	// Simple Field (TimeStamp)
	if pullErr := readBuffer.PullContext("TimeStamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TimeStamp")
	}
_TimeStamp, _TimeStampErr := DateAndTimeParseWithBuffer(ctx, readBuffer)
	if _TimeStampErr != nil {
		return nil, errors.Wrap(_TimeStampErr, "Error parsing 'TimeStamp' field of S7PayloadDiagnosticMessage")
	}
	TimeStamp := _TimeStamp.(DateAndTime)
	if closeErr := readBuffer.CloseContext("TimeStamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TimeStamp")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadDiagnosticMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadDiagnosticMessage")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadDiagnosticMessage{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{
		},
		EventId: EventId,
		PriorityClass: PriorityClass,
		ObNumber: ObNumber,
		DatId: DatId,
		Info1: Info1,
		Info2: Info2,
		TimeStamp: TimeStamp,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadDiagnosticMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadDiagnosticMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadDiagnosticMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadDiagnosticMessage")
		}

	// Simple Field (EventId)
	EventId := uint16(m.GetEventId())
	_EventIdErr := writeBuffer.WriteUint16("EventId", 16, (EventId))
	if _EventIdErr != nil {
		return errors.Wrap(_EventIdErr, "Error serializing 'EventId' field")
	}

	// Simple Field (PriorityClass)
	PriorityClass := uint8(m.GetPriorityClass())
	_PriorityClassErr := writeBuffer.WriteUint8("PriorityClass", 8, (PriorityClass))
	if _PriorityClassErr != nil {
		return errors.Wrap(_PriorityClassErr, "Error serializing 'PriorityClass' field")
	}

	// Simple Field (ObNumber)
	ObNumber := uint8(m.GetObNumber())
	_ObNumberErr := writeBuffer.WriteUint8("ObNumber", 8, (ObNumber))
	if _ObNumberErr != nil {
		return errors.Wrap(_ObNumberErr, "Error serializing 'ObNumber' field")
	}

	// Simple Field (DatId)
	DatId := uint16(m.GetDatId())
	_DatIdErr := writeBuffer.WriteUint16("DatId", 16, (DatId))
	if _DatIdErr != nil {
		return errors.Wrap(_DatIdErr, "Error serializing 'DatId' field")
	}

	// Simple Field (Info1)
	Info1 := uint16(m.GetInfo1())
	_Info1Err := writeBuffer.WriteUint16("Info1", 16, (Info1))
	if _Info1Err != nil {
		return errors.Wrap(_Info1Err, "Error serializing 'Info1' field")
	}

	// Simple Field (Info2)
	Info2 := uint32(m.GetInfo2())
	_Info2Err := writeBuffer.WriteUint32("Info2", 32, (Info2))
	if _Info2Err != nil {
		return errors.Wrap(_Info2Err, "Error serializing 'Info2' field")
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

		if popErr := writeBuffer.PopContext("S7PayloadDiagnosticMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadDiagnosticMessage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_S7PayloadDiagnosticMessage) isS7PayloadDiagnosticMessage() bool {
	return true
}

func (m *_S7PayloadDiagnosticMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



