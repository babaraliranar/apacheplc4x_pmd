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
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DeviceStatus is the corresponding interface of DeviceStatus
type DeviceStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetProgramMode returns ProgramMode (property field)
	GetProgramMode() bool
}

// DeviceStatusExactly can be used when we want exactly this type and not a type which fulfills DeviceStatus.
// This is useful for switch cases.
type DeviceStatusExactly interface {
	DeviceStatus
	isDeviceStatus() bool
}

// _DeviceStatus is the data-structure of this message
type _DeviceStatus struct {
	ProgramMode bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceStatus) GetProgramMode() bool {
	return m.ProgramMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceStatus factory function for _DeviceStatus
func NewDeviceStatus(programMode bool) *_DeviceStatus {
	return &_DeviceStatus{ProgramMode: programMode}
}

// Deprecated: use the interface for direct cast
func CastDeviceStatus(structType any) DeviceStatus {
	if casted, ok := structType.(DeviceStatus); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceStatus); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceStatus) GetTypeName() string {
	return "DeviceStatus"
}

func (m *_DeviceStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (programMode)
	lengthInBits += 1

	return lengthInBits
}

func (m *_DeviceStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceStatusParse(ctx context.Context, theBytes []byte) (DeviceStatus, error) {
	return DeviceStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeviceStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeviceStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DeviceStatus")
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

	// Simple Field (programMode)
	_programMode, _programModeErr := readBuffer.ReadBit("programMode")
	if _programModeErr != nil {
		return nil, errors.Wrap(_programModeErr, "Error parsing 'programMode' field of DeviceStatus")
	}
	programMode := _programMode

	if closeErr := readBuffer.CloseContext("DeviceStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceStatus")
	}

	// Create the instance
	return &_DeviceStatus{
		ProgramMode:    programMode,
		reservedField0: reservedField0,
	}, nil
}

func (m *_DeviceStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DeviceStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceStatus")
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
		_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (programMode)
	programMode := bool(m.GetProgramMode())
	_programModeErr := writeBuffer.WriteBit("programMode", (programMode))
	if _programModeErr != nil {
		return errors.Wrap(_programModeErr, "Error serializing 'programMode' field")
	}

	if popErr := writeBuffer.PopContext("DeviceStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceStatus")
	}
	return nil
}

func (m *_DeviceStatus) isDeviceStatus() bool {
	return true
}

func (m *_DeviceStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
