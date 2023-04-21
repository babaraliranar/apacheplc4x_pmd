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

// FirmataMessageSubscribeAnalogPinValue is the corresponding interface of FirmataMessageSubscribeAnalogPinValue
type FirmataMessageSubscribeAnalogPinValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	FirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetEnable returns Enable (property field)
	GetEnable() bool
}

// FirmataMessageSubscribeAnalogPinValueExactly can be used when we want exactly this type and not a type which fulfills FirmataMessageSubscribeAnalogPinValue.
// This is useful for switch cases.
type FirmataMessageSubscribeAnalogPinValueExactly interface {
	FirmataMessageSubscribeAnalogPinValue
	isFirmataMessageSubscribeAnalogPinValue() bool
}

// _FirmataMessageSubscribeAnalogPinValue is the data-structure of this message
type _FirmataMessageSubscribeAnalogPinValue struct {
	*_FirmataMessage
	Pin    uint8
	Enable bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageSubscribeAnalogPinValue) GetMessageType() uint8 {
	return 0xC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageSubscribeAnalogPinValue) InitializeParent(parent FirmataMessage) {}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetParent() FirmataMessage {
	return m._FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageSubscribeAnalogPinValue) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetEnable() bool {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageSubscribeAnalogPinValue factory function for _FirmataMessageSubscribeAnalogPinValue
func NewFirmataMessageSubscribeAnalogPinValue(pin uint8, enable bool, response bool) *_FirmataMessageSubscribeAnalogPinValue {
	_result := &_FirmataMessageSubscribeAnalogPinValue{
		Pin:             pin,
		Enable:          enable,
		_FirmataMessage: NewFirmataMessage(response),
	}
	_result._FirmataMessage._FirmataMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataMessageSubscribeAnalogPinValue(structType any) FirmataMessageSubscribeAnalogPinValue {
	if casted, ok := structType.(FirmataMessageSubscribeAnalogPinValue); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageSubscribeAnalogPinValue); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetTypeName() string {
	return "FirmataMessageSubscribeAnalogPinValue"
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 4

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enable)
	lengthInBits += 1

	return lengthInBits
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataMessageSubscribeAnalogPinValueParse(theBytes []byte, response bool) (FirmataMessageSubscribeAnalogPinValue, error) {
	return FirmataMessageSubscribeAnalogPinValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response)
}

func FirmataMessageSubscribeAnalogPinValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (FirmataMessageSubscribeAnalogPinValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageSubscribeAnalogPinValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageSubscribeAnalogPinValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 4)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of FirmataMessageSubscribeAnalogPinValue")
	}
	pin := _pin

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of FirmataMessageSubscribeAnalogPinValue")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (enable)
	_enable, _enableErr := readBuffer.ReadBit("enable")
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field of FirmataMessageSubscribeAnalogPinValue")
	}
	enable := _enable

	if closeErr := readBuffer.CloseContext("FirmataMessageSubscribeAnalogPinValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageSubscribeAnalogPinValue")
	}

	// Create a partially initialized instance
	_child := &_FirmataMessageSubscribeAnalogPinValue{
		_FirmataMessage: &_FirmataMessage{
			Response: response,
		},
		Pin:            pin,
		Enable:         enable,
		reservedField0: reservedField0,
	}
	_child._FirmataMessage._FirmataMessageChildRequirements = _child
	return _child, nil
}

func (m *_FirmataMessageSubscribeAnalogPinValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageSubscribeAnalogPinValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageSubscribeAnalogPinValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageSubscribeAnalogPinValue")
		}

		// Simple Field (pin)
		pin := uint8(m.GetPin())
		_pinErr := writeBuffer.WriteUint8("pin", 4, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (enable)
		enable := bool(m.GetEnable())
		_enableErr := writeBuffer.WriteBit("enable", (enable))
		if _enableErr != nil {
			return errors.Wrap(_enableErr, "Error serializing 'enable' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageSubscribeAnalogPinValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageSubscribeAnalogPinValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageSubscribeAnalogPinValue) isFirmataMessageSubscribeAnalogPinValue() bool {
	return true
}

func (m *_FirmataMessageSubscribeAnalogPinValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
