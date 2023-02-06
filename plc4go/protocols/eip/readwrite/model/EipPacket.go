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
	"encoding/binary"
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// EipPacket is the corresponding interface of EipPacket
type EipPacket interface {
	utils.LengthAware
	utils.Serializable
	// GetCommand returns Command (discriminator field)
	GetCommand() uint16
	// GetSessionHandle returns SessionHandle (property field)
	GetSessionHandle() uint32
	// GetStatus returns Status (property field)
	GetStatus() uint32
	// GetSenderContext returns SenderContext (property field)
	GetSenderContext() []uint8
	// GetOptions returns Options (property field)
	GetOptions() uint32
}

// EipPacketExactly can be used when we want exactly this type and not a type which fulfills EipPacket.
// This is useful for switch cases.
type EipPacketExactly interface {
	EipPacket
	isEipPacket() bool
}

// _EipPacket is the data-structure of this message
type _EipPacket struct {
	_EipPacketChildRequirements
	SessionHandle uint32
	Status        uint32
	SenderContext []uint8
	Options       uint32
}

type _EipPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommand() uint16
}

type EipPacketParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child EipPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type EipPacketChild interface {
	utils.Serializable
	InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32)
	GetParent() *EipPacket

	GetTypeName() string
	EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EipPacket) GetSessionHandle() uint32 {
	return m.SessionHandle
}

func (m *_EipPacket) GetStatus() uint32 {
	return m.Status
}

func (m *_EipPacket) GetSenderContext() []uint8 {
	return m.SenderContext
}

func (m *_EipPacket) GetOptions() uint32 {
	return m.Options
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEipPacket factory function for _EipPacket
func NewEipPacket(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_EipPacket {
	return &_EipPacket{SessionHandle: sessionHandle, Status: status, SenderContext: senderContext, Options: options}
}

// Deprecated: use the interface for direct cast
func CastEipPacket(structType interface{}) EipPacket {
	if casted, ok := structType.(EipPacket); ok {
		return casted
	}
	if casted, ok := structType.(*EipPacket); ok {
		return *casted
	}
	return nil
}

func (m *_EipPacket) GetTypeName() string {
	return "EipPacket"
}

func (m *_EipPacket) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (command)
	lengthInBits += 16

	// Implicit Field (packetLength)
	lengthInBits += 16

	// Simple field (sessionHandle)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += 32

	// Array field
	if len(m.SenderContext) > 0 {
		lengthInBits += 8 * uint16(len(m.SenderContext))
	}

	// Simple field (options)
	lengthInBits += 32

	return lengthInBits
}

func (m *_EipPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EipPacketParse(theBytes []byte) (EipPacket, error) {
	return EipPacketParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func EipPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EipPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (command) (Used as input to a switch field)
	command, _commandErr := readBuffer.ReadUint16("command", 16)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of EipPacket")
	}

	// Implicit Field (packetLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	packetLength, _packetLengthErr := readBuffer.ReadUint16("packetLength", 16)
	_ = packetLength
	if _packetLengthErr != nil {
		return nil, errors.Wrap(_packetLengthErr, "Error parsing 'packetLength' field of EipPacket")
	}

	// Simple Field (sessionHandle)
	_sessionHandle, _sessionHandleErr := readBuffer.ReadUint32("sessionHandle", 32)
	if _sessionHandleErr != nil {
		return nil, errors.Wrap(_sessionHandleErr, "Error parsing 'sessionHandle' field of EipPacket")
	}
	sessionHandle := _sessionHandle

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint32("status", 32)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of EipPacket")
	}
	status := _status

	// Array field (senderContext)
	if pullErr := readBuffer.PullContext("senderContext", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for senderContext")
	}
	// Count array
	senderContext := make([]uint8, uint16(8))
	// This happens when the size is set conditional to 0
	if len(senderContext) == 0 {
		senderContext = nil
	}
	{
		_numItems := uint16(uint16(8))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'senderContext' field of EipPacket")
			}
			senderContext[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("senderContext", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for senderContext")
	}

	// Simple Field (options)
	_options, _optionsErr := readBuffer.ReadUint32("options", 32)
	if _optionsErr != nil {
		return nil, errors.Wrap(_optionsErr, "Error parsing 'options' field of EipPacket")
	}
	options := _options

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type EipPacketChildSerializeRequirement interface {
		EipPacket
		InitializeParent(EipPacket, uint32, uint32, []uint8, uint32)
		GetParent() EipPacket
	}
	var _childTemp interface{}
	var _child EipPacketChildSerializeRequirement
	var typeSwitchError error
	switch {
	case command == 0x0065: // EipConnectionRequest
		_childTemp, typeSwitchError = EipConnectionRequestParseWithBuffer(ctx, readBuffer)
	case command == 0x0066: // EipDisconnectRequest
		_childTemp, typeSwitchError = EipDisconnectRequestParseWithBuffer(ctx, readBuffer)
	case command == 0x006F: // CipRRData
		_childTemp, typeSwitchError = CipRRDataParseWithBuffer(ctx, readBuffer, packetLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [command=%v]", command)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of EipPacket")
	}
	_child = _childTemp.(EipPacketChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("EipPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipPacket")
	}

	// Finish initializing
	_child.InitializeParent(_child, sessionHandle, status, senderContext, options)
	return _child, nil
}

func (pm *_EipPacket) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child EipPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("EipPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EipPacket")
	}

	// Discriminator Field (command) (Used as input to a switch field)
	command := uint16(child.GetCommand())
	_commandErr := writeBuffer.WriteUint16("command", 16, (command))

	if _commandErr != nil {
		return errors.Wrap(_commandErr, "Error serializing 'command' field")
	}

	// Implicit Field (packetLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	packetLength := uint16(uint16(uint16(m.GetLengthInBytes(ctx))) - uint16(uint16(24)))
	_packetLengthErr := writeBuffer.WriteUint16("packetLength", 16, (packetLength))
	if _packetLengthErr != nil {
		return errors.Wrap(_packetLengthErr, "Error serializing 'packetLength' field")
	}

	// Simple Field (sessionHandle)
	sessionHandle := uint32(m.GetSessionHandle())
	_sessionHandleErr := writeBuffer.WriteUint32("sessionHandle", 32, (sessionHandle))
	if _sessionHandleErr != nil {
		return errors.Wrap(_sessionHandleErr, "Error serializing 'sessionHandle' field")
	}

	// Simple Field (status)
	status := uint32(m.GetStatus())
	_statusErr := writeBuffer.WriteUint32("status", 32, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Array Field (senderContext)
	if pushErr := writeBuffer.PushContext("senderContext", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for senderContext")
	}
	for _curItem, _element := range m.GetSenderContext() {
		_ = _curItem
		_elementErr := writeBuffer.WriteUint8("", 8, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'senderContext' field")
		}
	}
	if popErr := writeBuffer.PopContext("senderContext", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for senderContext")
	}

	// Simple Field (options)
	options := uint32(m.GetOptions())
	_optionsErr := writeBuffer.WriteUint32("options", 32, (options))
	if _optionsErr != nil {
		return errors.Wrap(_optionsErr, "Error serializing 'options' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("EipPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EipPacket")
	}
	return nil
}

func (m *_EipPacket) isEipPacket() bool {
	return true
}

func (m *_EipPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
