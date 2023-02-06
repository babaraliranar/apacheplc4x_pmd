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

// BACnetServiceAckAuthenticate is the corresponding interface of BACnetServiceAckAuthenticate
type BACnetServiceAckAuthenticate interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
}

// BACnetServiceAckAuthenticateExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAuthenticate.
// This is useful for switch cases.
type BACnetServiceAckAuthenticateExactly interface {
	BACnetServiceAckAuthenticate
	isBACnetServiceAckAuthenticate() bool
}

// _BACnetServiceAckAuthenticate is the data-structure of this message
type _BACnetServiceAckAuthenticate struct {
	*_BACnetServiceAck
	BytesOfRemovedService []byte

	// Arguments.
	ServiceAckPayloadLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckAuthenticate) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_AUTHENTICATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAuthenticate) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckAuthenticate) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAuthenticate) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAuthenticate factory function for _BACnetServiceAckAuthenticate
func NewBACnetServiceAckAuthenticate(bytesOfRemovedService []byte, serviceAckPayloadLength uint32, serviceAckLength uint32) *_BACnetServiceAckAuthenticate {
	_result := &_BACnetServiceAckAuthenticate{
		BytesOfRemovedService: bytesOfRemovedService,
		_BACnetServiceAck:     NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAuthenticate(structType interface{}) BACnetServiceAckAuthenticate {
	if casted, ok := structType.(BACnetServiceAckAuthenticate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAuthenticate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAuthenticate) GetTypeName() string {
	return "BACnetServiceAckAuthenticate"
}

func (m *_BACnetServiceAckAuthenticate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetServiceAckAuthenticate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckAuthenticateParse(theBytes []byte, serviceAckPayloadLength uint32, serviceAckLength uint32) (BACnetServiceAckAuthenticate, error) {
	return BACnetServiceAckAuthenticateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceAckPayloadLength, serviceAckLength)
}

func BACnetServiceAckAuthenticateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckPayloadLength uint32, serviceAckLength uint32) (BACnetServiceAckAuthenticate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAuthenticate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAuthenticate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (bytesOfRemovedService)
	numberOfBytesbytesOfRemovedService := int(serviceAckPayloadLength)
	bytesOfRemovedService, _readArrayErr := readBuffer.ReadByteArray("bytesOfRemovedService", numberOfBytesbytesOfRemovedService)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bytesOfRemovedService' field of BACnetServiceAckAuthenticate")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAuthenticate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAuthenticate")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckAuthenticate{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		BytesOfRemovedService: bytesOfRemovedService,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckAuthenticate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAuthenticate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAuthenticate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAuthenticate")
		}

		// Array Field (bytesOfRemovedService)
		// Byte Array field (bytesOfRemovedService)
		if err := writeBuffer.WriteByteArray("bytesOfRemovedService", m.GetBytesOfRemovedService()); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAuthenticate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAuthenticate")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetServiceAckAuthenticate) GetServiceAckPayloadLength() uint32 {
	return m.ServiceAckPayloadLength
}

//
////

func (m *_BACnetServiceAckAuthenticate) isBACnetServiceAckAuthenticate() bool {
	return true
}

func (m *_BACnetServiceAckAuthenticate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
