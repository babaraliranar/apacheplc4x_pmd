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


// ModbusPDUGetComEventLogResponse is the corresponding interface of ModbusPDUGetComEventLogResponse
type ModbusPDUGetComEventLogResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
	// GetMessageCount returns MessageCount (property field)
	GetMessageCount() uint16
	// GetEvents returns Events (property field)
	GetEvents() []byte
}

// ModbusPDUGetComEventLogResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUGetComEventLogResponse.
// This is useful for switch cases.
type ModbusPDUGetComEventLogResponseExactly interface {
	ModbusPDUGetComEventLogResponse
	isModbusPDUGetComEventLogResponse() bool
}

// _ModbusPDUGetComEventLogResponse is the data-structure of this message
type _ModbusPDUGetComEventLogResponse struct {
	*_ModbusPDU
        Status uint16
        EventCount uint16
        MessageCount uint16
        Events []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventLogResponse)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUGetComEventLogResponse)  GetFunctionFlag() uint8 {
return 0x0C}

func (m *_ModbusPDUGetComEventLogResponse)  GetResponse() bool {
return bool(true)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventLogResponse) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUGetComEventLogResponse)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUGetComEventLogResponse) GetStatus() uint16 {
	return m.Status
}

func (m *_ModbusPDUGetComEventLogResponse) GetEventCount() uint16 {
	return m.EventCount
}

func (m *_ModbusPDUGetComEventLogResponse) GetMessageCount() uint16 {
	return m.MessageCount
}

func (m *_ModbusPDUGetComEventLogResponse) GetEvents() []byte {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUGetComEventLogResponse factory function for _ModbusPDUGetComEventLogResponse
func NewModbusPDUGetComEventLogResponse( status uint16 , eventCount uint16 , messageCount uint16 , events []byte ) *_ModbusPDUGetComEventLogResponse {
	_result := &_ModbusPDUGetComEventLogResponse{
		Status: status,
		EventCount: eventCount,
		MessageCount: messageCount,
		Events: events,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventLogResponse(structType interface{}) ModbusPDUGetComEventLogResponse {
    if casted, ok := structType.(ModbusPDUGetComEventLogResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventLogResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventLogResponse) GetTypeName() string {
	return "ModbusPDUGetComEventLogResponse"
}

func (m *_ModbusPDUGetComEventLogResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16;

	// Simple field (eventCount)
	lengthInBits += 16;

	// Simple field (messageCount)
	lengthInBits += 16;

	// Array field
	if len(m.Events) > 0 {
		lengthInBits += 8 * uint16(len(m.Events))
	}

	return lengthInBits
}


func (m *_ModbusPDUGetComEventLogResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUGetComEventLogResponseParse(theBytes []byte, response bool) (ModbusPDUGetComEventLogResponse, error) {
	return ModbusPDUGetComEventLogResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUGetComEventLogResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUGetComEventLogResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventLogResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventLogResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUGetComEventLogResponse")
	}

	// Simple Field (status)
_status, _statusErr := readBuffer.ReadUint16("status", 16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of ModbusPDUGetComEventLogResponse")
	}
	status := _status

	// Simple Field (eventCount)
_eventCount, _eventCountErr := readBuffer.ReadUint16("eventCount", 16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field of ModbusPDUGetComEventLogResponse")
	}
	eventCount := _eventCount

	// Simple Field (messageCount)
_messageCount, _messageCountErr := readBuffer.ReadUint16("messageCount", 16)
	if _messageCountErr != nil {
		return nil, errors.Wrap(_messageCountErr, "Error parsing 'messageCount' field of ModbusPDUGetComEventLogResponse")
	}
	messageCount := _messageCount
	// Byte Array field (events)
	numberOfBytesevents := int(uint16(byteCount) - uint16(uint16(6)))
	events, _readArrayErr := readBuffer.ReadByteArray("events", numberOfBytesevents)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'events' field of ModbusPDUGetComEventLogResponse")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventLogResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventLogResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUGetComEventLogResponse{
		_ModbusPDU: &_ModbusPDU{
		},
		Status: status,
		EventCount: eventCount,
		MessageCount: messageCount,
		Events: events,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUGetComEventLogResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUGetComEventLogResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventLogResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventLogResponse")
		}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount := uint8(uint8(uint8(len(m.GetEvents()))) + uint8(uint8(6)))
	_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
	if _byteCountErr != nil {
		return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
	}

	// Simple Field (status)
	status := uint16(m.GetStatus())
	_statusErr := writeBuffer.WriteUint16("status", 16, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Simple Field (eventCount)
	eventCount := uint16(m.GetEventCount())
	_eventCountErr := writeBuffer.WriteUint16("eventCount", 16, (eventCount))
	if _eventCountErr != nil {
		return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
	}

	// Simple Field (messageCount)
	messageCount := uint16(m.GetMessageCount())
	_messageCountErr := writeBuffer.WriteUint16("messageCount", 16, (messageCount))
	if _messageCountErr != nil {
		return errors.Wrap(_messageCountErr, "Error serializing 'messageCount' field")
	}

	// Array Field (events)
	// Byte Array field (events)
	if err := writeBuffer.WriteByteArray("events", m.GetEvents()); err != nil {
		return errors.Wrap(err, "Error serializing 'events' field")
	}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventLogResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventLogResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ModbusPDUGetComEventLogResponse) isModbusPDUGetComEventLogResponse() bool {
	return true
}

func (m *_ModbusPDUGetComEventLogResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



