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
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// IdentifyReplyCommandOutputUnitSummary is the corresponding interface of IdentifyReplyCommandOutputUnitSummary
type IdentifyReplyCommandOutputUnitSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetUnitFlags returns UnitFlags (property field)
	GetUnitFlags() IdentifyReplyCommandUnitSummary
	// GetGavStoreEnabledByte1 returns GavStoreEnabledByte1 (property field)
	GetGavStoreEnabledByte1() *byte
	// GetGavStoreEnabledByte2 returns GavStoreEnabledByte2 (property field)
	GetGavStoreEnabledByte2() *byte
	// GetTimeFromLastRecoverOfMainsInSeconds returns TimeFromLastRecoverOfMainsInSeconds (property field)
	GetTimeFromLastRecoverOfMainsInSeconds() uint8
}

// IdentifyReplyCommandOutputUnitSummaryExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandOutputUnitSummary.
// This is useful for switch cases.
type IdentifyReplyCommandOutputUnitSummaryExactly interface {
	IdentifyReplyCommandOutputUnitSummary
	isIdentifyReplyCommandOutputUnitSummary() bool
}

// _IdentifyReplyCommandOutputUnitSummary is the data-structure of this message
type _IdentifyReplyCommandOutputUnitSummary struct {
	*_IdentifyReplyCommand
	UnitFlags                           IdentifyReplyCommandUnitSummary
	GavStoreEnabledByte1                *byte
	GavStoreEnabledByte2                *byte
	TimeFromLastRecoverOfMainsInSeconds uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandOutputUnitSummary) GetAttribute() Attribute {
	return Attribute_OutputUnitSummary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandOutputUnitSummary) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandOutputUnitSummary) GetUnitFlags() IdentifyReplyCommandUnitSummary {
	return m.UnitFlags
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetGavStoreEnabledByte1() *byte {
	return m.GavStoreEnabledByte1
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetGavStoreEnabledByte2() *byte {
	return m.GavStoreEnabledByte2
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetTimeFromLastRecoverOfMainsInSeconds() uint8 {
	return m.TimeFromLastRecoverOfMainsInSeconds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandOutputUnitSummary factory function for _IdentifyReplyCommandOutputUnitSummary
func NewIdentifyReplyCommandOutputUnitSummary(unitFlags IdentifyReplyCommandUnitSummary, gavStoreEnabledByte1 *byte, gavStoreEnabledByte2 *byte, timeFromLastRecoverOfMainsInSeconds uint8, numBytes uint8) *_IdentifyReplyCommandOutputUnitSummary {
	_result := &_IdentifyReplyCommandOutputUnitSummary{
		UnitFlags:                           unitFlags,
		GavStoreEnabledByte1:                gavStoreEnabledByte1,
		GavStoreEnabledByte2:                gavStoreEnabledByte2,
		TimeFromLastRecoverOfMainsInSeconds: timeFromLastRecoverOfMainsInSeconds,
		_IdentifyReplyCommand:               NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandOutputUnitSummary(structType any) IdentifyReplyCommandOutputUnitSummary {
	if casted, ok := structType.(IdentifyReplyCommandOutputUnitSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandOutputUnitSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetTypeName() string {
	return "IdentifyReplyCommandOutputUnitSummary"
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unitFlags)
	lengthInBits += m.UnitFlags.GetLengthInBits(ctx)

	// Optional Field (gavStoreEnabledByte1)
	if m.GavStoreEnabledByte1 != nil {
		lengthInBits += 8
	}

	// Optional Field (gavStoreEnabledByte2)
	if m.GavStoreEnabledByte2 != nil {
		lengthInBits += 8
	}

	// Simple field (timeFromLastRecoverOfMainsInSeconds)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandOutputUnitSummaryParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandOutputUnitSummary, error) {
	return IdentifyReplyCommandOutputUnitSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandOutputUnitSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandOutputUnitSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandOutputUnitSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandOutputUnitSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unitFlags)
	if pullErr := readBuffer.PullContext("unitFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unitFlags")
	}
	_unitFlags, _unitFlagsErr := IdentifyReplyCommandUnitSummaryParseWithBuffer(ctx, readBuffer)
	if _unitFlagsErr != nil {
		return nil, errors.Wrap(_unitFlagsErr, "Error parsing 'unitFlags' field of IdentifyReplyCommandOutputUnitSummary")
	}
	unitFlags := _unitFlags.(IdentifyReplyCommandUnitSummary)
	if closeErr := readBuffer.CloseContext("unitFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unitFlags")
	}

	// Optional Field (gavStoreEnabledByte1) (Can be skipped, if a given expression evaluates to false)
	var gavStoreEnabledByte1 *byte = nil
	if bool((numBytes) > (1)) {
		currentPos = positionAware.GetPos()
		_val, _err := readBuffer.ReadByte("gavStoreEnabledByte1")
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'gavStoreEnabledByte1' field of IdentifyReplyCommandOutputUnitSummary")
		default:
			gavStoreEnabledByte1 = &_val
		}
	}

	// Optional Field (gavStoreEnabledByte2) (Can be skipped, if a given expression evaluates to false)
	var gavStoreEnabledByte2 *byte = nil
	if bool((numBytes) > (2)) {
		currentPos = positionAware.GetPos()
		_val, _err := readBuffer.ReadByte("gavStoreEnabledByte2")
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'gavStoreEnabledByte2' field of IdentifyReplyCommandOutputUnitSummary")
		default:
			gavStoreEnabledByte2 = &_val
		}
	}

	// Simple Field (timeFromLastRecoverOfMainsInSeconds)
	_timeFromLastRecoverOfMainsInSeconds, _timeFromLastRecoverOfMainsInSecondsErr := readBuffer.ReadUint8("timeFromLastRecoverOfMainsInSeconds", 8)
	if _timeFromLastRecoverOfMainsInSecondsErr != nil {
		return nil, errors.Wrap(_timeFromLastRecoverOfMainsInSecondsErr, "Error parsing 'timeFromLastRecoverOfMainsInSeconds' field of IdentifyReplyCommandOutputUnitSummary")
	}
	timeFromLastRecoverOfMainsInSeconds := _timeFromLastRecoverOfMainsInSeconds

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandOutputUnitSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandOutputUnitSummary")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandOutputUnitSummary{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		UnitFlags:                           unitFlags,
		GavStoreEnabledByte1:                gavStoreEnabledByte1,
		GavStoreEnabledByte2:                gavStoreEnabledByte2,
		TimeFromLastRecoverOfMainsInSeconds: timeFromLastRecoverOfMainsInSeconds,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandOutputUnitSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandOutputUnitSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandOutputUnitSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandOutputUnitSummary")
		}

		// Simple Field (unitFlags)
		if pushErr := writeBuffer.PushContext("unitFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unitFlags")
		}
		_unitFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetUnitFlags())
		if popErr := writeBuffer.PopContext("unitFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unitFlags")
		}
		if _unitFlagsErr != nil {
			return errors.Wrap(_unitFlagsErr, "Error serializing 'unitFlags' field")
		}

		// Optional Field (gavStoreEnabledByte1) (Can be skipped, if the value is null)
		var gavStoreEnabledByte1 *byte = nil
		if m.GetGavStoreEnabledByte1() != nil {
			gavStoreEnabledByte1 = m.GetGavStoreEnabledByte1()
			_gavStoreEnabledByte1Err := writeBuffer.WriteByte("gavStoreEnabledByte1", *(gavStoreEnabledByte1))
			if _gavStoreEnabledByte1Err != nil {
				return errors.Wrap(_gavStoreEnabledByte1Err, "Error serializing 'gavStoreEnabledByte1' field")
			}
		}

		// Optional Field (gavStoreEnabledByte2) (Can be skipped, if the value is null)
		var gavStoreEnabledByte2 *byte = nil
		if m.GetGavStoreEnabledByte2() != nil {
			gavStoreEnabledByte2 = m.GetGavStoreEnabledByte2()
			_gavStoreEnabledByte2Err := writeBuffer.WriteByte("gavStoreEnabledByte2", *(gavStoreEnabledByte2))
			if _gavStoreEnabledByte2Err != nil {
				return errors.Wrap(_gavStoreEnabledByte2Err, "Error serializing 'gavStoreEnabledByte2' field")
			}
		}

		// Simple Field (timeFromLastRecoverOfMainsInSeconds)
		timeFromLastRecoverOfMainsInSeconds := uint8(m.GetTimeFromLastRecoverOfMainsInSeconds())
		_timeFromLastRecoverOfMainsInSecondsErr := writeBuffer.WriteUint8("timeFromLastRecoverOfMainsInSeconds", 8, uint8((timeFromLastRecoverOfMainsInSeconds)))
		if _timeFromLastRecoverOfMainsInSecondsErr != nil {
			return errors.Wrap(_timeFromLastRecoverOfMainsInSecondsErr, "Error serializing 'timeFromLastRecoverOfMainsInSeconds' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandOutputUnitSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandOutputUnitSummary")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandOutputUnitSummary) isIdentifyReplyCommandOutputUnitSummary() bool {
	return true
}

func (m *_IdentifyReplyCommandOutputUnitSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
