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


// PowerUpReply is the corresponding interface of PowerUpReply
type PowerUpReply interface {
	utils.LengthAware
	utils.Serializable
	Reply
	// GetPowerUpIndicator returns PowerUpIndicator (property field)
	GetPowerUpIndicator() PowerUp
}

// PowerUpReplyExactly can be used when we want exactly this type and not a type which fulfills PowerUpReply.
// This is useful for switch cases.
type PowerUpReplyExactly interface {
	PowerUpReply
	isPowerUpReply() bool
}

// _PowerUpReply is the data-structure of this message
type _PowerUpReply struct {
	*_Reply
        PowerUpIndicator PowerUp
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PowerUpReply) InitializeParent(parent Reply , peekedByte byte ) {	m.PeekedByte = peekedByte
}

func (m *_PowerUpReply)  GetParent() Reply {
	return m._Reply
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PowerUpReply) GetPowerUpIndicator() PowerUp {
	return m.PowerUpIndicator
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewPowerUpReply factory function for _PowerUpReply
func NewPowerUpReply( powerUpIndicator PowerUp , peekedByte byte , cBusOptions CBusOptions , requestContext RequestContext ) *_PowerUpReply {
	_result := &_PowerUpReply{
		PowerUpIndicator: powerUpIndicator,
    	_Reply: NewReply(peekedByte, cBusOptions, requestContext),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPowerUpReply(structType interface{}) PowerUpReply {
    if casted, ok := structType.(PowerUpReply); ok {
		return casted
	}
	if casted, ok := structType.(*PowerUpReply); ok {
		return *casted
	}
	return nil
}

func (m *_PowerUpReply) GetTypeName() string {
	return "PowerUpReply"
}

func (m *_PowerUpReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (powerUpIndicator)
	lengthInBits += m.PowerUpIndicator.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_PowerUpReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PowerUpReplyParse(theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (PowerUpReply, error) {
	return PowerUpReplyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func PowerUpReplyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (PowerUpReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PowerUpReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PowerUpReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (powerUpIndicator)
	if pullErr := readBuffer.PullContext("powerUpIndicator"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for powerUpIndicator")
	}
_powerUpIndicator, _powerUpIndicatorErr := PowerUpParseWithBuffer(ctx, readBuffer)
	if _powerUpIndicatorErr != nil {
		return nil, errors.Wrap(_powerUpIndicatorErr, "Error parsing 'powerUpIndicator' field of PowerUpReply")
	}
	powerUpIndicator := _powerUpIndicator.(PowerUp)
	if closeErr := readBuffer.CloseContext("powerUpIndicator"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for powerUpIndicator")
	}

	if closeErr := readBuffer.CloseContext("PowerUpReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PowerUpReply")
	}

	// Create a partially initialized instance
	_child := &_PowerUpReply{
		_Reply: &_Reply{
			CBusOptions: cBusOptions,
			RequestContext: requestContext,
		},
		PowerUpIndicator: powerUpIndicator,
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_PowerUpReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PowerUpReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PowerUpReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PowerUpReply")
		}

	// Simple Field (powerUpIndicator)
	if pushErr := writeBuffer.PushContext("powerUpIndicator"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for powerUpIndicator")
	}
	_powerUpIndicatorErr := writeBuffer.WriteSerializable(ctx, m.GetPowerUpIndicator())
	if popErr := writeBuffer.PopContext("powerUpIndicator"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for powerUpIndicator")
	}
	if _powerUpIndicatorErr != nil {
		return errors.Wrap(_powerUpIndicatorErr, "Error serializing 'powerUpIndicator' field")
	}

		if popErr := writeBuffer.PopContext("PowerUpReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PowerUpReply")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_PowerUpReply) isPowerUpReply() bool {
	return true
}

func (m *_PowerUpReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



