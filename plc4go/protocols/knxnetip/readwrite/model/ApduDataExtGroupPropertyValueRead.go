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

// ApduDataExtGroupPropertyValueRead is the corresponding interface of ApduDataExtGroupPropertyValueRead
type ApduDataExtGroupPropertyValueRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtGroupPropertyValueRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtGroupPropertyValueRead()
}

// _ApduDataExtGroupPropertyValueRead is the data-structure of this message
type _ApduDataExtGroupPropertyValueRead struct {
	ApduDataExtContract
}

var _ ApduDataExtGroupPropertyValueRead = (*_ApduDataExtGroupPropertyValueRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtGroupPropertyValueRead)(nil)

// NewApduDataExtGroupPropertyValueRead factory function for _ApduDataExtGroupPropertyValueRead
func NewApduDataExtGroupPropertyValueRead(length uint8) *_ApduDataExtGroupPropertyValueRead {
	_result := &_ApduDataExtGroupPropertyValueRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueRead) GetExtApciType() uint8 {
	return 0x28
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueRead(structType any) ApduDataExtGroupPropertyValueRead {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueRead) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueRead"
}

func (m *_ApduDataExtGroupPropertyValueRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtGroupPropertyValueRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtGroupPropertyValueRead ApduDataExtGroupPropertyValueRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueRead")
	}

	return m, nil
}

func (m *_ApduDataExtGroupPropertyValueRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtGroupPropertyValueRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueRead) IsApduDataExtGroupPropertyValueRead() {}

func (m *_ApduDataExtGroupPropertyValueRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtGroupPropertyValueRead) deepCopy() *_ApduDataExtGroupPropertyValueRead {
	if m == nil {
		return nil
	}
	_ApduDataExtGroupPropertyValueReadCopy := &_ApduDataExtGroupPropertyValueRead{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtGroupPropertyValueReadCopy
}

func (m *_ApduDataExtGroupPropertyValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
