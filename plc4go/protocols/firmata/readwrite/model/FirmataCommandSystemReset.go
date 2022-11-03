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

// FirmataCommandSystemReset is the corresponding interface of FirmataCommandSystemReset
type FirmataCommandSystemReset interface {
	utils.LengthAware
	utils.Serializable
	FirmataCommand
}

// FirmataCommandSystemResetExactly can be used when we want exactly this type and not a type which fulfills FirmataCommandSystemReset.
// This is useful for switch cases.
type FirmataCommandSystemResetExactly interface {
	FirmataCommandSystemReset
	isFirmataCommandSystemReset() bool
}

// _FirmataCommandSystemReset is the data-structure of this message
type _FirmataCommandSystemReset struct {
	*_FirmataCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandSystemReset) GetCommandCode() uint8 {
	return 0xF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandSystemReset) InitializeParent(parent FirmataCommand) {}

func (m *_FirmataCommandSystemReset) GetParent() FirmataCommand {
	return m._FirmataCommand
}

// NewFirmataCommandSystemReset factory function for _FirmataCommandSystemReset
func NewFirmataCommandSystemReset(response bool) *_FirmataCommandSystemReset {
	_result := &_FirmataCommandSystemReset{
		_FirmataCommand: NewFirmataCommand(response),
	}
	_result._FirmataCommand._FirmataCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataCommandSystemReset(structType interface{}) FirmataCommandSystemReset {
	if casted, ok := structType.(FirmataCommandSystemReset); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandSystemReset); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandSystemReset) GetTypeName() string {
	return "FirmataCommandSystemReset"
}

func (m *_FirmataCommandSystemReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_FirmataCommandSystemReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_FirmataCommandSystemReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSystemResetParse(theBytes []byte, response bool) (FirmataCommandSystemReset, error) {
	return FirmataCommandSystemResetParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response) // TODO: get endianness from mspec
}

func FirmataCommandSystemResetParseWithBuffer(readBuffer utils.ReadBuffer, response bool) (FirmataCommandSystemReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandSystemReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandSystemReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("FirmataCommandSystemReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandSystemReset")
	}

	// Create a partially initialized instance
	_child := &_FirmataCommandSystemReset{
		_FirmataCommand: &_FirmataCommand{
			Response: response,
		},
	}
	_child._FirmataCommand._FirmataCommandChildRequirements = _child
	return _child, nil
}

func (m *_FirmataCommandSystemReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSystemReset) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSystemReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSystemReset")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSystemReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSystemReset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_FirmataCommandSystemReset) isFirmataCommandSystemReset() bool {
	return true
}

func (m *_FirmataCommandSystemReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
