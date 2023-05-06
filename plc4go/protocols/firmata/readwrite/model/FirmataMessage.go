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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessage is the corresponding interface of FirmataMessage
type FirmataMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
}

// FirmataMessageExactly can be used when we want exactly this type and not a type which fulfills FirmataMessage.
// This is useful for switch cases.
type FirmataMessageExactly interface {
	FirmataMessage
	isFirmataMessage() bool
}

// _FirmataMessage is the data-structure of this message
type _FirmataMessage struct {
	_FirmataMessageChildRequirements

	// Arguments.
	Response bool
}

type _FirmataMessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() uint8
}

type FirmataMessageParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child FirmataMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type FirmataMessageChild interface {
	utils.Serializable
	InitializeParent(parent FirmataMessage)
	GetParent() *FirmataMessage

	GetTypeName() string
	FirmataMessage
}

// NewFirmataMessage factory function for _FirmataMessage
func NewFirmataMessage(response bool) *_FirmataMessage {
	return &_FirmataMessage{Response: response}
}

// Deprecated: use the interface for direct cast
func CastFirmataMessage(structType any) FirmataMessage {
	if casted, ok := structType.(FirmataMessage); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessage); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessage) GetTypeName() string {
	return "FirmataMessage"
}

func (m *_FirmataMessage) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_FirmataMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataMessageParse(theBytes []byte, response bool) (FirmataMessage, error) {
	return FirmataMessageParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response)
}

func FirmataMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (FirmataMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 4)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field of FirmataMessage")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type FirmataMessageChildSerializeRequirement interface {
		FirmataMessage
		InitializeParent(FirmataMessage)
		GetParent() FirmataMessage
	}
	var _childTemp any
	var _child FirmataMessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case messageType == 0xE: // FirmataMessageAnalogIO
		_childTemp, typeSwitchError = FirmataMessageAnalogIOParseWithBuffer(ctx, readBuffer, response)
	case messageType == 0x9: // FirmataMessageDigitalIO
		_childTemp, typeSwitchError = FirmataMessageDigitalIOParseWithBuffer(ctx, readBuffer, response)
	case messageType == 0xC: // FirmataMessageSubscribeAnalogPinValue
		_childTemp, typeSwitchError = FirmataMessageSubscribeAnalogPinValueParseWithBuffer(ctx, readBuffer, response)
	case messageType == 0xD: // FirmataMessageSubscribeDigitalPinValue
		_childTemp, typeSwitchError = FirmataMessageSubscribeDigitalPinValueParseWithBuffer(ctx, readBuffer, response)
	case messageType == 0xF: // FirmataMessageCommand
		_childTemp, typeSwitchError = FirmataMessageCommandParseWithBuffer(ctx, readBuffer, response)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageType=%v]", messageType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of FirmataMessage")
	}
	_child = _childTemp.(FirmataMessageChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("FirmataMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessage")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_FirmataMessage) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child FirmataMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("FirmataMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for FirmataMessage")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 4, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("FirmataMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for FirmataMessage")
	}
	return nil
}

////
// Arguments Getter

func (m *_FirmataMessage) GetResponse() bool {
	return m.Response
}

//
////

func (m *_FirmataMessage) isFirmataMessage() bool {
	return true
}

func (m *_FirmataMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
