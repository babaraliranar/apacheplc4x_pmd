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

// ModbusPDUWriteMultipleHoldingRegistersResponse is the corresponding interface of ModbusPDUWriteMultipleHoldingRegistersResponse
type ModbusPDUWriteMultipleHoldingRegistersResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
}

// ModbusPDUWriteMultipleHoldingRegistersResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteMultipleHoldingRegistersResponse.
// This is useful for switch cases.
type ModbusPDUWriteMultipleHoldingRegistersResponseExactly interface {
	ModbusPDUWriteMultipleHoldingRegistersResponse
	isModbusPDUWriteMultipleHoldingRegistersResponse() bool
}

// _ModbusPDUWriteMultipleHoldingRegistersResponse is the data-structure of this message
type _ModbusPDUWriteMultipleHoldingRegistersResponse struct {
	*_ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetFunctionFlag() uint8 {
	return 0x10
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteMultipleHoldingRegistersResponse factory function for _ModbusPDUWriteMultipleHoldingRegistersResponse
func NewModbusPDUWriteMultipleHoldingRegistersResponse(startingAddress uint16, quantity uint16) *_ModbusPDUWriteMultipleHoldingRegistersResponse {
	_result := &_ModbusPDUWriteMultipleHoldingRegistersResponse{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		_ModbusPDU:      NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteMultipleHoldingRegistersResponse(structType interface{}) ModbusPDUWriteMultipleHoldingRegistersResponse {
	if casted, ok := structType.(ModbusPDUWriteMultipleHoldingRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteMultipleHoldingRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUWriteMultipleHoldingRegistersResponse"
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteMultipleHoldingRegistersResponseParse(theBytes []byte, response bool) (ModbusPDUWriteMultipleHoldingRegistersResponse, error) {
	return ModbusPDUWriteMultipleHoldingRegistersResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUWriteMultipleHoldingRegistersResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteMultipleHoldingRegistersResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteMultipleHoldingRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field of ModbusPDUWriteMultipleHoldingRegistersResponse")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field of ModbusPDUWriteMultipleHoldingRegistersResponse")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteMultipleHoldingRegistersResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteMultipleHoldingRegistersResponse{
		_ModbusPDU:      &_ModbusPDU{},
		StartingAddress: startingAddress,
		Quantity:        quantity,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteMultipleHoldingRegistersResponse")
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.GetStartingAddress())
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.GetQuantity())
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleHoldingRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteMultipleHoldingRegistersResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) isModbusPDUWriteMultipleHoldingRegistersResponse() bool {
	return true
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
