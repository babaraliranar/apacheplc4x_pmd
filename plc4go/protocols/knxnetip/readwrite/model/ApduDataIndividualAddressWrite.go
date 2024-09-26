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

// ApduDataIndividualAddressWrite is the corresponding interface of ApduDataIndividualAddressWrite
type ApduDataIndividualAddressWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// IsApduDataIndividualAddressWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataIndividualAddressWrite()
}

// _ApduDataIndividualAddressWrite is the data-structure of this message
type _ApduDataIndividualAddressWrite struct {
	ApduDataContract
}

var _ ApduDataIndividualAddressWrite = (*_ApduDataIndividualAddressWrite)(nil)
var _ ApduDataRequirements = (*_ApduDataIndividualAddressWrite)(nil)

// NewApduDataIndividualAddressWrite factory function for _ApduDataIndividualAddressWrite
func NewApduDataIndividualAddressWrite(dataLength uint8) *_ApduDataIndividualAddressWrite {
	_result := &_ApduDataIndividualAddressWrite{
		ApduDataContract: NewApduData(dataLength),
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataIndividualAddressWrite) GetApciType() uint8 {
	return 0x3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataIndividualAddressWrite) GetParent() ApduDataContract {
	return m.ApduDataContract
}

// Deprecated: use the interface for direct cast
func CastApduDataIndividualAddressWrite(structType any) ApduDataIndividualAddressWrite {
	if casted, ok := structType.(ApduDataIndividualAddressWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataIndividualAddressWrite) GetTypeName() string {
	return "ApduDataIndividualAddressWrite"
}

func (m *_ApduDataIndividualAddressWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataIndividualAddressWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataIndividualAddressWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataIndividualAddressWrite ApduDataIndividualAddressWrite, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataIndividualAddressWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataIndividualAddressWrite")
	}

	return m, nil
}

func (m *_ApduDataIndividualAddressWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataIndividualAddressWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataIndividualAddressWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataIndividualAddressWrite")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataIndividualAddressWrite) IsApduDataIndividualAddressWrite() {}

func (m *_ApduDataIndividualAddressWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataIndividualAddressWrite) deepCopy() *_ApduDataIndividualAddressWrite {
	if m == nil {
		return nil
	}
	_ApduDataIndividualAddressWriteCopy := &_ApduDataIndividualAddressWrite{
		m.ApduDataContract.(*_ApduData).deepCopy(),
	}
	m.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataIndividualAddressWriteCopy
}

func (m *_ApduDataIndividualAddressWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
