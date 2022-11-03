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

// MeteringDataGasConsumption is the corresponding interface of MeteringDataGasConsumption
type MeteringDataGasConsumption interface {
	utils.LengthAware
	utils.Serializable
	MeteringData
	// GetMJ returns MJ (property field)
	GetMJ() uint32
}

// MeteringDataGasConsumptionExactly can be used when we want exactly this type and not a type which fulfills MeteringDataGasConsumption.
// This is useful for switch cases.
type MeteringDataGasConsumptionExactly interface {
	MeteringDataGasConsumption
	isMeteringDataGasConsumption() bool
}

// _MeteringDataGasConsumption is the data-structure of this message
type _MeteringDataGasConsumption struct {
	*_MeteringData
	MJ uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataGasConsumption) InitializeParent(parent MeteringData, commandTypeContainer MeteringCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataGasConsumption) GetParent() MeteringData {
	return m._MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataGasConsumption) GetMJ() uint32 {
	return m.MJ
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeteringDataGasConsumption factory function for _MeteringDataGasConsumption
func NewMeteringDataGasConsumption(mJ uint32, commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataGasConsumption {
	_result := &_MeteringDataGasConsumption{
		MJ:            mJ,
		_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataGasConsumption(structType interface{}) MeteringDataGasConsumption {
	if casted, ok := structType.(MeteringDataGasConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataGasConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataGasConsumption) GetTypeName() string {
	return "MeteringDataGasConsumption"
}

func (m *_MeteringDataGasConsumption) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MeteringDataGasConsumption) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (mJ)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MeteringDataGasConsumption) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeteringDataGasConsumptionParse(theBytes []byte) (MeteringDataGasConsumption, error) {
	return MeteringDataGasConsumptionParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func MeteringDataGasConsumptionParseWithBuffer(readBuffer utils.ReadBuffer) (MeteringDataGasConsumption, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataGasConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataGasConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (mJ)
	_mJ, _mJErr := readBuffer.ReadUint32("mJ", 32)
	if _mJErr != nil {
		return nil, errors.Wrap(_mJErr, "Error parsing 'mJ' field of MeteringDataGasConsumption")
	}
	mJ := _mJ

	if closeErr := readBuffer.CloseContext("MeteringDataGasConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataGasConsumption")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataGasConsumption{
		_MeteringData: &_MeteringData{},
		MJ:            mJ,
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataGasConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataGasConsumption) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataGasConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataGasConsumption")
		}

		// Simple Field (mJ)
		mJ := uint32(m.GetMJ())
		_mJErr := writeBuffer.WriteUint32("mJ", 32, (mJ))
		if _mJErr != nil {
			return errors.Wrap(_mJErr, "Error serializing 'mJ' field")
		}

		if popErr := writeBuffer.PopContext("MeteringDataGasConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataGasConsumption")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MeteringDataGasConsumption) isMeteringDataGasConsumption() bool {
	return true
}

func (m *_MeteringDataGasConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
