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

// TamperStatus is the corresponding interface of TamperStatus
type TamperStatus interface {
	utils.LengthAware
	utils.Serializable
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetIsNoTamper returns IsNoTamper (virtual field)
	GetIsNoTamper() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsTamperActive returns IsTamperActive (virtual field)
	GetIsTamperActive() bool
}

// TamperStatusExactly can be used when we want exactly this type and not a type which fulfills TamperStatus.
// This is useful for switch cases.
type TamperStatusExactly interface {
	TamperStatus
	isTamperStatus() bool
}

// _TamperStatus is the data-structure of this message
type _TamperStatus struct {
	Status uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TamperStatus) GetStatus() uint8 {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TamperStatus) GetIsNoTamper() bool {
	return bool(bool((m.GetStatus()) == (0x00)))
}

func (m *_TamperStatus) GetIsReserved() bool {
	return bool(bool(bool((m.GetStatus()) >= (0x01))) && bool(bool((m.GetStatus()) <= (0xFE))))
}

func (m *_TamperStatus) GetIsTamperActive() bool {
	return bool(bool((m.GetStatus()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTamperStatus factory function for _TamperStatus
func NewTamperStatus(status uint8) *_TamperStatus {
	return &_TamperStatus{Status: status}
}

// Deprecated: use the interface for direct cast
func CastTamperStatus(structType interface{}) TamperStatus {
	if casted, ok := structType.(TamperStatus); ok {
		return casted
	}
	if casted, ok := structType.(*TamperStatus); ok {
		return *casted
	}
	return nil
}

func (m *_TamperStatus) GetTypeName() string {
	return "TamperStatus"
}

func (m *_TamperStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TamperStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (status)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TamperStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TamperStatusParse(theBytes []byte) (TamperStatus, error) {
	return TamperStatusParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func TamperStatusParseWithBuffer(readBuffer utils.ReadBuffer) (TamperStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TamperStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TamperStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of TamperStatus")
	}
	status := _status

	// Virtual field
	_isNoTamper := bool((status) == (0x00))
	isNoTamper := bool(_isNoTamper)
	_ = isNoTamper

	// Virtual field
	_isReserved := bool(bool((status) >= (0x01))) && bool(bool((status) <= (0xFE)))
	isReserved := bool(_isReserved)
	_ = isReserved

	// Virtual field
	_isTamperActive := bool((status) > (0xFE))
	isTamperActive := bool(_isTamperActive)
	_ = isTamperActive

	if closeErr := readBuffer.CloseContext("TamperStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TamperStatus")
	}

	// Create the instance
	return &_TamperStatus{
		Status: status,
	}, nil
}

func (m *_TamperStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TamperStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("TamperStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TamperStatus")
	}

	// Simple Field (status)
	status := uint8(m.GetStatus())
	_statusErr := writeBuffer.WriteUint8("status", 8, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}
	// Virtual field
	if _isNoTamperErr := writeBuffer.WriteVirtual("isNoTamper", m.GetIsNoTamper()); _isNoTamperErr != nil {
		return errors.Wrap(_isNoTamperErr, "Error serializing 'isNoTamper' field")
	}
	// Virtual field
	if _isReservedErr := writeBuffer.WriteVirtual("isReserved", m.GetIsReserved()); _isReservedErr != nil {
		return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
	}
	// Virtual field
	if _isTamperActiveErr := writeBuffer.WriteVirtual("isTamperActive", m.GetIsTamperActive()); _isTamperActiveErr != nil {
		return errors.Wrap(_isTamperActiveErr, "Error serializing 'isTamperActive' field")
	}

	if popErr := writeBuffer.PopContext("TamperStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TamperStatus")
	}
	return nil
}

func (m *_TamperStatus) isTamperStatus() bool {
	return true
}

func (m *_TamperStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
