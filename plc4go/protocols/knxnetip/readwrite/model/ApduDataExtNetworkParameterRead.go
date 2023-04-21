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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtNetworkParameterRead is the corresponding interface of ApduDataExtNetworkParameterRead
type ApduDataExtNetworkParameterRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtNetworkParameterReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtNetworkParameterRead.
// This is useful for switch cases.
type ApduDataExtNetworkParameterReadExactly interface {
	ApduDataExtNetworkParameterRead
	isApduDataExtNetworkParameterRead() bool
}

// _ApduDataExtNetworkParameterRead is the data-structure of this message
type _ApduDataExtNetworkParameterRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtNetworkParameterRead) GetExtApciType() uint8 {
	return 0x1A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtNetworkParameterRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtNetworkParameterRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtNetworkParameterRead factory function for _ApduDataExtNetworkParameterRead
func NewApduDataExtNetworkParameterRead(length uint8) *_ApduDataExtNetworkParameterRead {
	_result := &_ApduDataExtNetworkParameterRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtNetworkParameterRead(structType any) ApduDataExtNetworkParameterRead {
	if casted, ok := structType.(ApduDataExtNetworkParameterRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtNetworkParameterRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtNetworkParameterRead) GetTypeName() string {
	return "ApduDataExtNetworkParameterRead"
}

func (m *_ApduDataExtNetworkParameterRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtNetworkParameterRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtNetworkParameterReadParse(theBytes []byte, length uint8) (ApduDataExtNetworkParameterRead, error) {
	return ApduDataExtNetworkParameterReadParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtNetworkParameterReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtNetworkParameterRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtNetworkParameterRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtNetworkParameterRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtNetworkParameterRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtNetworkParameterRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtNetworkParameterRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtNetworkParameterRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtNetworkParameterRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtNetworkParameterRead) isApduDataExtNetworkParameterRead() bool {
	return true
}

func (m *_ApduDataExtNetworkParameterRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
