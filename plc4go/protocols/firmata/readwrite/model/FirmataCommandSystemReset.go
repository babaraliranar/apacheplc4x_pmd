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

// FirmataCommandSystemReset is the corresponding interface of FirmataCommandSystemReset
type FirmataCommandSystemReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataCommand
	// IsFirmataCommandSystemReset is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataCommandSystemReset()
}

// _FirmataCommandSystemReset is the data-structure of this message
type _FirmataCommandSystemReset struct {
	FirmataCommandContract
}

var _ FirmataCommandSystemReset = (*_FirmataCommandSystemReset)(nil)
var _ FirmataCommandRequirements = (*_FirmataCommandSystemReset)(nil)

// NewFirmataCommandSystemReset factory function for _FirmataCommandSystemReset
func NewFirmataCommandSystemReset(response bool) *_FirmataCommandSystemReset {
	_result := &_FirmataCommandSystemReset{
		FirmataCommandContract: NewFirmataCommand(response),
	}
	_result.FirmataCommandContract.(*_FirmataCommand)._SubType = _result
	return _result
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

func (m *_FirmataCommandSystemReset) GetParent() FirmataCommandContract {
	return m.FirmataCommandContract
}

// Deprecated: use the interface for direct cast
func CastFirmataCommandSystemReset(structType any) FirmataCommandSystemReset {
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

func (m *_FirmataCommandSystemReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataCommandContract.(*_FirmataCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_FirmataCommandSystemReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataCommandSystemReset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataCommand, response bool) (__firmataCommandSystemReset FirmataCommandSystemReset, err error) {
	m.FirmataCommandContract = parent
	parent._SubType = m
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

	return m, nil
}

func (m *_FirmataCommandSystemReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSystemReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSystemReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSystemReset")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSystemReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSystemReset")
		}
		return nil
	}
	return m.FirmataCommandContract.(*_FirmataCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataCommandSystemReset) IsFirmataCommandSystemReset() {}

func (m *_FirmataCommandSystemReset) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataCommandSystemReset) deepCopy() *_FirmataCommandSystemReset {
	if m == nil {
		return nil
	}
	_FirmataCommandSystemResetCopy := &_FirmataCommandSystemReset{
		m.FirmataCommandContract.(*_FirmataCommand).deepCopy(),
	}
	m.FirmataCommandContract.(*_FirmataCommand)._SubType = m
	return _FirmataCommandSystemResetCopy
}

func (m *_FirmataCommandSystemReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
