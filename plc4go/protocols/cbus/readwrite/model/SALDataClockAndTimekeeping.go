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

// SALDataClockAndTimekeeping is the corresponding interface of SALDataClockAndTimekeeping
type SALDataClockAndTimekeeping interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetClockAndTimekeepingData returns ClockAndTimekeepingData (property field)
	GetClockAndTimekeepingData() ClockAndTimekeepingData
}

// SALDataClockAndTimekeepingExactly can be used when we want exactly this type and not a type which fulfills SALDataClockAndTimekeeping.
// This is useful for switch cases.
type SALDataClockAndTimekeepingExactly interface {
	SALDataClockAndTimekeeping
	isSALDataClockAndTimekeeping() bool
}

// _SALDataClockAndTimekeeping is the data-structure of this message
type _SALDataClockAndTimekeeping struct {
	*_SALData
	ClockAndTimekeepingData ClockAndTimekeepingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataClockAndTimekeeping) GetApplicationId() ApplicationId {
	return ApplicationId_CLOCK_AND_TIMEKEEPING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataClockAndTimekeeping) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataClockAndTimekeeping) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataClockAndTimekeeping) GetClockAndTimekeepingData() ClockAndTimekeepingData {
	return m.ClockAndTimekeepingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataClockAndTimekeeping factory function for _SALDataClockAndTimekeeping
func NewSALDataClockAndTimekeeping(clockAndTimekeepingData ClockAndTimekeepingData, salData SALData) *_SALDataClockAndTimekeeping {
	_result := &_SALDataClockAndTimekeeping{
		ClockAndTimekeepingData: clockAndTimekeepingData,
		_SALData:                NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataClockAndTimekeeping(structType interface{}) SALDataClockAndTimekeeping {
	if casted, ok := structType.(SALDataClockAndTimekeeping); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataClockAndTimekeeping); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataClockAndTimekeeping) GetTypeName() string {
	return "SALDataClockAndTimekeeping"
}

func (m *_SALDataClockAndTimekeeping) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataClockAndTimekeeping) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (clockAndTimekeepingData)
	lengthInBits += m.ClockAndTimekeepingData.GetLengthInBits()

	return lengthInBits
}

func (m *_SALDataClockAndTimekeeping) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataClockAndTimekeepingParse(theBytes []byte, applicationId ApplicationId) (SALDataClockAndTimekeeping, error) {
	return SALDataClockAndTimekeepingParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), applicationId) // TODO: get endianness from mspec
}

func SALDataClockAndTimekeepingParseWithBuffer(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataClockAndTimekeeping, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataClockAndTimekeeping"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataClockAndTimekeeping")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (clockAndTimekeepingData)
	if pullErr := readBuffer.PullContext("clockAndTimekeepingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clockAndTimekeepingData")
	}
	_clockAndTimekeepingData, _clockAndTimekeepingDataErr := ClockAndTimekeepingDataParseWithBuffer(readBuffer)
	if _clockAndTimekeepingDataErr != nil {
		return nil, errors.Wrap(_clockAndTimekeepingDataErr, "Error parsing 'clockAndTimekeepingData' field of SALDataClockAndTimekeeping")
	}
	clockAndTimekeepingData := _clockAndTimekeepingData.(ClockAndTimekeepingData)
	if closeErr := readBuffer.CloseContext("clockAndTimekeepingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clockAndTimekeepingData")
	}

	if closeErr := readBuffer.CloseContext("SALDataClockAndTimekeeping"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataClockAndTimekeeping")
	}

	// Create a partially initialized instance
	_child := &_SALDataClockAndTimekeeping{
		_SALData:                &_SALData{},
		ClockAndTimekeepingData: clockAndTimekeepingData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataClockAndTimekeeping) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataClockAndTimekeeping) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataClockAndTimekeeping"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataClockAndTimekeeping")
		}

		// Simple Field (clockAndTimekeepingData)
		if pushErr := writeBuffer.PushContext("clockAndTimekeepingData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clockAndTimekeepingData")
		}
		_clockAndTimekeepingDataErr := writeBuffer.WriteSerializable(m.GetClockAndTimekeepingData())
		if popErr := writeBuffer.PopContext("clockAndTimekeepingData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clockAndTimekeepingData")
		}
		if _clockAndTimekeepingDataErr != nil {
			return errors.Wrap(_clockAndTimekeepingDataErr, "Error serializing 'clockAndTimekeepingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataClockAndTimekeeping"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataClockAndTimekeeping")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataClockAndTimekeeping) isSALDataClockAndTimekeeping() bool {
	return true
}

func (m *_SALDataClockAndTimekeeping) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
