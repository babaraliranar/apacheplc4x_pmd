//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const KnxNetIpMessage_PROTOCOLVERSION uint8 = 0x10

// The data-structure of this message
type KnxNetIpMessage struct {
	Child IKnxNetIpMessageChild
}

// The corresponding interface
type IKnxNetIpMessage interface {
	MsgType() uint16
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IKnxNetIpMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IKnxNetIpMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type IKnxNetIpMessageChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *KnxNetIpMessage)
	GetTypeName() string
	IKnxNetIpMessage
}

func NewKnxNetIpMessage() *KnxNetIpMessage {
	return &KnxNetIpMessage{}
}

func CastKnxNetIpMessage(structType interface{}) *KnxNetIpMessage {
	castFunc := func(typ interface{}) *KnxNetIpMessage {
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxNetIpMessage) GetTypeName() string {
	return "KnxNetIpMessage"
}

func (m *KnxNetIpMessage) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxNetIpMessage) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *KnxNetIpMessage) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8

	// Const Field (protocolVersion)
	lengthInBits += 8
	// Discriminator Field (msgType)
	lengthInBits += 16

	// Implicit Field (totalLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *KnxNetIpMessage) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetIpMessageParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("KnxNetIpMessage"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength, _headerLengthErr := readBuffer.ReadUint8("headerLength", 8)
	_ = headerLength
	if _headerLengthErr != nil {
		return nil, errors.Wrap(_headerLengthErr, "Error parsing 'headerLength' field")
	}

	// Const Field (protocolVersion)
	protocolVersion, _protocolVersionErr := readBuffer.ReadUint8("protocolVersion", 8)
	if _protocolVersionErr != nil {
		return nil, errors.Wrap(_protocolVersionErr, "Error parsing 'protocolVersion' field")
	}
	if protocolVersion != KnxNetIpMessage_PROTOCOLVERSION {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", KnxNetIpMessage_PROTOCOLVERSION) + " but got " + fmt.Sprintf("%d", protocolVersion))
	}

	// Discriminator Field (msgType) (Used as input to a switch field)
	msgType, _msgTypeErr := readBuffer.ReadUint16("msgType", 16)
	if _msgTypeErr != nil {
		return nil, errors.Wrap(_msgTypeErr, "Error parsing 'msgType' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	totalLength, _totalLengthErr := readBuffer.ReadUint16("totalLength", 16)
	_ = totalLength
	if _totalLengthErr != nil {
		return nil, errors.Wrap(_totalLengthErr, "Error parsing 'totalLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *KnxNetIpMessage
	var typeSwitchError error
	switch {
	case msgType == 0x0201: // SearchRequest
		_parent, typeSwitchError = SearchRequestParse(readBuffer)
	case msgType == 0x0202: // SearchResponse
		_parent, typeSwitchError = SearchResponseParse(readBuffer)
	case msgType == 0x0203: // DescriptionRequest
		_parent, typeSwitchError = DescriptionRequestParse(readBuffer)
	case msgType == 0x0204: // DescriptionResponse
		_parent, typeSwitchError = DescriptionResponseParse(readBuffer)
	case msgType == 0x0205: // ConnectionRequest
		_parent, typeSwitchError = ConnectionRequestParse(readBuffer)
	case msgType == 0x0206: // ConnectionResponse
		_parent, typeSwitchError = ConnectionResponseParse(readBuffer)
	case msgType == 0x0207: // ConnectionStateRequest
		_parent, typeSwitchError = ConnectionStateRequestParse(readBuffer)
	case msgType == 0x0208: // ConnectionStateResponse
		_parent, typeSwitchError = ConnectionStateResponseParse(readBuffer)
	case msgType == 0x0209: // DisconnectRequest
		_parent, typeSwitchError = DisconnectRequestParse(readBuffer)
	case msgType == 0x020A: // DisconnectResponse
		_parent, typeSwitchError = DisconnectResponseParse(readBuffer)
	case msgType == 0x020B: // UnknownMessage
		_parent, typeSwitchError = UnknownMessageParse(readBuffer, totalLength)
	case msgType == 0x0310: // DeviceConfigurationRequest
		_parent, typeSwitchError = DeviceConfigurationRequestParse(readBuffer, totalLength)
	case msgType == 0x0311: // DeviceConfigurationAck
		_parent, typeSwitchError = DeviceConfigurationAckParse(readBuffer)
	case msgType == 0x0420: // TunnelingRequest
		_parent, typeSwitchError = TunnelingRequestParse(readBuffer, totalLength)
	case msgType == 0x0421: // TunnelingResponse
		_parent, typeSwitchError = TunnelingResponseParse(readBuffer)
	case msgType == 0x0530: // RoutingIndication
		_parent, typeSwitchError = RoutingIndicationParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("KnxNetIpMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *KnxNetIpMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *KnxNetIpMessage) SerializeParent(writeBuffer utils.WriteBuffer, child IKnxNetIpMessage, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("KnxNetIpMessage"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(6))
	_headerLengthErr := writeBuffer.WriteUint8("headerLength", 8, (headerLength))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	// Const Field (protocolVersion)
	_protocolVersionErr := writeBuffer.WriteUint8("protocolVersion", 8, 0x10)
	if _protocolVersionErr != nil {
		return errors.Wrap(_protocolVersionErr, "Error serializing 'protocolVersion' field")
	}

	// Discriminator Field (msgType) (Used as input to a switch field)
	msgType := uint16(child.MsgType())
	_msgTypeErr := writeBuffer.WriteUint16("msgType", 16, (msgType))

	if _msgTypeErr != nil {
		return errors.Wrap(_msgTypeErr, "Error serializing 'msgType' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	totalLength := uint16(uint16(m.LengthInBytes()))
	_totalLengthErr := writeBuffer.WriteUint16("totalLength", 16, (totalLength))
	if _totalLengthErr != nil {
		return errors.Wrap(_totalLengthErr, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("KnxNetIpMessage"); popErr != nil {
		return popErr
	}
	return nil
}
