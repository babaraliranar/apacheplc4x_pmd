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

// KnxAddress is the corresponding interface of KnxAddress
type KnxAddress interface {
	utils.LengthAware
	utils.Serializable
	// GetMainGroup returns MainGroup (property field)
	GetMainGroup() uint8
	// GetMiddleGroup returns MiddleGroup (property field)
	GetMiddleGroup() uint8
	// GetSubGroup returns SubGroup (property field)
	GetSubGroup() uint8
}

// KnxAddressExactly can be used when we want exactly this type and not a type which fulfills KnxAddress.
// This is useful for switch cases.
type KnxAddressExactly interface {
	KnxAddress
	isKnxAddress() bool
}

// _KnxAddress is the data-structure of this message
type _KnxAddress struct {
	MainGroup   uint8
	MiddleGroup uint8
	SubGroup    uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxAddress) GetMainGroup() uint8 {
	return m.MainGroup
}

func (m *_KnxAddress) GetMiddleGroup() uint8 {
	return m.MiddleGroup
}

func (m *_KnxAddress) GetSubGroup() uint8 {
	return m.SubGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxAddress factory function for _KnxAddress
func NewKnxAddress(mainGroup uint8, middleGroup uint8, subGroup uint8) *_KnxAddress {
	return &_KnxAddress{MainGroup: mainGroup, MiddleGroup: middleGroup, SubGroup: subGroup}
}

// Deprecated: use the interface for direct cast
func CastKnxAddress(structType interface{}) KnxAddress {
	if casted, ok := structType.(KnxAddress); ok {
		return casted
	}
	if casted, ok := structType.(*KnxAddress); ok {
		return *casted
	}
	return nil
}

func (m *_KnxAddress) GetTypeName() string {
	return "KnxAddress"
}

func (m *_KnxAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (mainGroup)
	lengthInBits += 4

	// Simple field (middleGroup)
	lengthInBits += 4

	// Simple field (subGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxAddressParse(theBytes []byte) (KnxAddress, error) {
	return KnxAddressParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func KnxAddressParseWithBuffer(readBuffer utils.ReadBuffer) (KnxAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (mainGroup)
	_mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 4)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field of KnxAddress")
	}
	mainGroup := _mainGroup

	// Simple Field (middleGroup)
	_middleGroup, _middleGroupErr := readBuffer.ReadUint8("middleGroup", 4)
	if _middleGroupErr != nil {
		return nil, errors.Wrap(_middleGroupErr, "Error parsing 'middleGroup' field of KnxAddress")
	}
	middleGroup := _middleGroup

	// Simple Field (subGroup)
	_subGroup, _subGroupErr := readBuffer.ReadUint8("subGroup", 8)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field of KnxAddress")
	}
	subGroup := _subGroup

	if closeErr := readBuffer.CloseContext("KnxAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxAddress")
	}

	// Create the instance
	return &_KnxAddress{
		MainGroup:   mainGroup,
		MiddleGroup: middleGroup,
		SubGroup:    subGroup,
	}, nil
}

func (m *_KnxAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxAddress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("KnxAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for KnxAddress")
	}

	// Simple Field (mainGroup)
	mainGroup := uint8(m.GetMainGroup())
	_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 4, (mainGroup))
	if _mainGroupErr != nil {
		return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
	}

	// Simple Field (middleGroup)
	middleGroup := uint8(m.GetMiddleGroup())
	_middleGroupErr := writeBuffer.WriteUint8("middleGroup", 4, (middleGroup))
	if _middleGroupErr != nil {
		return errors.Wrap(_middleGroupErr, "Error serializing 'middleGroup' field")
	}

	// Simple Field (subGroup)
	subGroup := uint8(m.GetSubGroup())
	_subGroupErr := writeBuffer.WriteUint8("subGroup", 8, (subGroup))
	if _subGroupErr != nil {
		return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
	}

	if popErr := writeBuffer.PopContext("KnxAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for KnxAddress")
	}
	return nil
}

func (m *_KnxAddress) isKnxAddress() bool {
	return true
}

func (m *_KnxAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
