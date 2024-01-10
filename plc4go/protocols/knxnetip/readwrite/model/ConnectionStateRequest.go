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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionStateRequest is the corresponding interface of ConnectionStateRequest
type ConnectionStateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
}

// ConnectionStateRequestExactly can be used when we want exactly this type and not a type which fulfills ConnectionStateRequest.
// This is useful for switch cases.
type ConnectionStateRequestExactly interface {
	ConnectionStateRequest
	isConnectionStateRequest() bool
}

// _ConnectionStateRequest is the data-structure of this message
type _ConnectionStateRequest struct {
	*_KnxNetIpMessage
	CommunicationChannelId uint8
	HpaiControlEndpoint    HPAIControlEndpoint
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionStateRequest) GetMsgType() uint16 {
	return 0x0207
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionStateRequest) InitializeParent(parent KnxNetIpMessage) {}

func (m *_ConnectionStateRequest) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionStateRequest) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_ConnectionStateRequest) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionStateRequest factory function for _ConnectionStateRequest
func NewConnectionStateRequest(communicationChannelId uint8, hpaiControlEndpoint HPAIControlEndpoint) *_ConnectionStateRequest {
	_result := &_ConnectionStateRequest{
		CommunicationChannelId: communicationChannelId,
		HpaiControlEndpoint:    hpaiControlEndpoint,
		_KnxNetIpMessage:       NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionStateRequest(structType any) ConnectionStateRequest {
	if casted, ok := structType.(ConnectionStateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionStateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionStateRequest) GetTypeName() string {
	return "ConnectionStateRequest"
}

func (m *_ConnectionStateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ConnectionStateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectionStateRequestParse(ctx context.Context, theBytes []byte) (ConnectionStateRequest, error) {
	return ConnectionStateRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func ConnectionStateRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectionStateRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ConnectionStateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionStateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (communicationChannelId)
	_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of ConnectionStateRequest")
	}
	communicationChannelId := _communicationChannelId

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ConnectionStateRequest")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (hpaiControlEndpoint)
	if pullErr := readBuffer.PullContext("hpaiControlEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiControlEndpoint")
	}
	_hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParseWithBuffer(ctx, readBuffer)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field of ConnectionStateRequest")
	}
	hpaiControlEndpoint := _hpaiControlEndpoint.(HPAIControlEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiControlEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiControlEndpoint")
	}

	if closeErr := readBuffer.CloseContext("ConnectionStateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionStateRequest")
	}

	// Create a partially initialized instance
	_child := &_ConnectionStateRequest{
		_KnxNetIpMessage:       &_KnxNetIpMessage{},
		CommunicationChannelId: communicationChannelId,
		HpaiControlEndpoint:    hpaiControlEndpoint,
		reservedField0:         reservedField0,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionStateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionStateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionStateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionStateRequest")
		}

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.GetCommunicationChannelId())
		_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, uint8((communicationChannelId)))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (hpaiControlEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiControlEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiControlEndpoint")
		}
		_hpaiControlEndpointErr := writeBuffer.WriteSerializable(ctx, m.GetHpaiControlEndpoint())
		if popErr := writeBuffer.PopContext("hpaiControlEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiControlEndpoint")
		}
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionStateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionStateRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionStateRequest) isConnectionStateRequest() bool {
	return true
}

func (m *_ConnectionStateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
