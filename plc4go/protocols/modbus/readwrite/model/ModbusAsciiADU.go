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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusAsciiADU is the corresponding interface of ModbusAsciiADU
type ModbusAsciiADU interface {
	utils.LengthAware
	utils.Serializable
	ModbusADU
	// GetAddress returns Address (property field)
	GetAddress() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
}

// ModbusAsciiADUExactly can be used when we want exactly this type and not a type which fulfills ModbusAsciiADU.
// This is useful for switch cases.
type ModbusAsciiADUExactly interface {
	ModbusAsciiADU
	isModbusAsciiADU() bool
}

// _ModbusAsciiADU is the data-structure of this message
type _ModbusAsciiADU struct {
	*_ModbusADU
	Address uint8
	Pdu     ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusAsciiADU) GetDriverType() DriverType {
	return DriverType_MODBUS_ASCII
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusAsciiADU) InitializeParent(parent ModbusADU) {}

func (m *_ModbusAsciiADU) GetParent() ModbusADU {
	return m._ModbusADU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusAsciiADU) GetAddress() uint8 {
	return m.Address
}

func (m *_ModbusAsciiADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusAsciiADU factory function for _ModbusAsciiADU
func NewModbusAsciiADU(address uint8, pdu ModbusPDU, response bool) *_ModbusAsciiADU {
	_result := &_ModbusAsciiADU{
		Address:    address,
		Pdu:        pdu,
		_ModbusADU: NewModbusADU(response),
	}
	_result._ModbusADU._ModbusADUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusAsciiADU(structType interface{}) ModbusAsciiADU {
	if casted, ok := structType.(ModbusAsciiADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusAsciiADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusAsciiADU) GetTypeName() string {
	return "ModbusAsciiADU"
}

func (m *_ModbusAsciiADU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusAsciiADU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits()

	// Checksum Field (checksum)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusAsciiADU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusAsciiADUParse(theBytes []byte, driverType DriverType, response bool) (ModbusAsciiADU, error) {
	return ModbusAsciiADUParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), driverType, response) // TODO: get endianness from mspec
}

func ModbusAsciiADUParseWithBuffer(readBuffer utils.ReadBuffer, driverType DriverType, response bool) (ModbusAsciiADU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusAsciiADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusAsciiADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint8("address", 8)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ModbusAsciiADU")
	}
	address := _address

	// Simple Field (pdu)
	if pullErr := readBuffer.PullContext("pdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pdu")
	}
	_pdu, _pduErr := ModbusPDUParseWithBuffer(readBuffer, bool(response))
	if _pduErr != nil {
		return nil, errors.Wrap(_pduErr, "Error parsing 'pdu' field of ModbusAsciiADU")
	}
	pdu := _pdu.(ModbusPDU)
	if closeErr := readBuffer.CloseContext("pdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pdu")
	}

	// Checksum Field (checksum)
	{
		checksumRef, _checksumRefErr := readBuffer.ReadUint8("checksum", 8)
		if _checksumRefErr != nil {
			return nil, errors.Wrap(_checksumRefErr, "Error parsing 'checksum' field of ModbusAsciiADU")
		}
		checksum, _checksumErr := AsciiLrcCheck(address, pdu)
		if _checksumErr != nil {
			return nil, errors.Wrap(_checksumErr, "Checksum verification failed")
		}
		if checksum != checksumRef {
			return nil, errors.Errorf("Checksum verification failed. Expected %x but got %x", checksumRef, checksum)
		}
	}

	if closeErr := readBuffer.CloseContext("ModbusAsciiADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusAsciiADU")
	}

	// Create a partially initialized instance
	_child := &_ModbusAsciiADU{
		_ModbusADU: &_ModbusADU{
			Response: response,
		},
		Address: address,
		Pdu:     pdu,
	}
	_child._ModbusADU._ModbusADUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusAsciiADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusAsciiADU) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusAsciiADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusAsciiADU")
		}

		// Simple Field (address)
		address := uint8(m.GetAddress())
		_addressErr := writeBuffer.WriteUint8("address", 8, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (pdu)
		if pushErr := writeBuffer.PushContext("pdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pdu")
		}
		_pduErr := writeBuffer.WriteSerializable(m.GetPdu())
		if popErr := writeBuffer.PopContext("pdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pdu")
		}
		if _pduErr != nil {
			return errors.Wrap(_pduErr, "Error serializing 'pdu' field")
		}

		// Checksum Field (checksum) (Calculated)
		{
			_checksum, _checksumErr := AsciiLrcCheck(m.GetAddress(), m.GetPdu())
			if _checksumErr != nil {
				return errors.Wrap(_checksumErr, "Checksum calculation failed")
			}
			_checksumWriteErr := writeBuffer.WriteUint8("checksum", 8, (_checksum))
			if _checksumWriteErr != nil {
				return errors.Wrap(_checksumWriteErr, "Error serializing 'checksum' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusAsciiADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusAsciiADU")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusAsciiADU) isModbusAsciiADU() bool {
	return true
}

func (m *_ModbusAsciiADU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
