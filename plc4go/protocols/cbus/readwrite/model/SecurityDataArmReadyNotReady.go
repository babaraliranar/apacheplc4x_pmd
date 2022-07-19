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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataArmReadyNotReady is the corresponding interface of SecurityDataArmReadyNotReady
type SecurityDataArmReadyNotReady interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataArmReadyNotReadyExactly can be used when we want exactly this type and not a type which fulfills SecurityDataArmReadyNotReady.
// This is useful for switch cases.
type SecurityDataArmReadyNotReadyExactly interface {
	SecurityDataArmReadyNotReady
	isSecurityDataArmReadyNotReady() bool
}

// _SecurityDataArmReadyNotReady is the data-structure of this message
type _SecurityDataArmReadyNotReady struct {
	*_SecurityData
	ZoneNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataArmReadyNotReady) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataArmReadyNotReady) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataArmReadyNotReady) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataArmReadyNotReady factory function for _SecurityDataArmReadyNotReady
func NewSecurityDataArmReadyNotReady(zoneNumber uint8, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataArmReadyNotReady {
	_result := &_SecurityDataArmReadyNotReady{
		ZoneNumber:    zoneNumber,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataArmReadyNotReady(structType interface{}) SecurityDataArmReadyNotReady {
	if casted, ok := structType.(SecurityDataArmReadyNotReady); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataArmReadyNotReady); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataArmReadyNotReady) GetTypeName() string {
	return "SecurityDataArmReadyNotReady"
}

func (m *_SecurityDataArmReadyNotReady) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataArmReadyNotReady) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataArmReadyNotReady) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataArmReadyNotReadyParse(readBuffer utils.ReadBuffer) (SecurityDataArmReadyNotReady, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataArmReadyNotReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataArmReadyNotReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
	_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataArmReadyNotReady")
	}
	zoneNumber := _zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataArmReadyNotReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataArmReadyNotReady")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataArmReadyNotReady{
		ZoneNumber:    zoneNumber,
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataArmReadyNotReady) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataArmReadyNotReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataArmReadyNotReady")
		}

		// Simple Field (zoneNumber)
		zoneNumber := uint8(m.GetZoneNumber())
		_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, (zoneNumber))
		if _zoneNumberErr != nil {
			return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataArmReadyNotReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataArmReadyNotReady")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataArmReadyNotReady) isSecurityDataArmReadyNotReady() bool {
	return true
}

func (m *_SecurityDataArmReadyNotReady) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
