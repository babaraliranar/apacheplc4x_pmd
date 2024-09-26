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

// ModbusPDUGetComEventLogRequest is the corresponding interface of ModbusPDUGetComEventLogRequest
type ModbusPDUGetComEventLogRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// IsModbusPDUGetComEventLogRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUGetComEventLogRequest()
}

// _ModbusPDUGetComEventLogRequest is the data-structure of this message
type _ModbusPDUGetComEventLogRequest struct {
	ModbusPDUContract
}

var _ ModbusPDUGetComEventLogRequest = (*_ModbusPDUGetComEventLogRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUGetComEventLogRequest)(nil)

// NewModbusPDUGetComEventLogRequest factory function for _ModbusPDUGetComEventLogRequest
func NewModbusPDUGetComEventLogRequest() *_ModbusPDUGetComEventLogRequest {
	_result := &_ModbusPDUGetComEventLogRequest{
		ModbusPDUContract: NewModbusPDU(),
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventLogRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUGetComEventLogRequest) GetFunctionFlag() uint8 {
	return 0x0C
}

func (m *_ModbusPDUGetComEventLogRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventLogRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventLogRequest(structType any) ModbusPDUGetComEventLogRequest {
	if casted, ok := structType.(ModbusPDUGetComEventLogRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventLogRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventLogRequest) GetTypeName() string {
	return "ModbusPDUGetComEventLogRequest"
}

func (m *_ModbusPDUGetComEventLogRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ModbusPDUGetComEventLogRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUGetComEventLogRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUGetComEventLogRequest ModbusPDUGetComEventLogRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventLogRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventLogRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventLogRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventLogRequest")
	}

	return m, nil
}

func (m *_ModbusPDUGetComEventLogRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUGetComEventLogRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventLogRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventLogRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventLogRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventLogRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUGetComEventLogRequest) IsModbusPDUGetComEventLogRequest() {}

func (m *_ModbusPDUGetComEventLogRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUGetComEventLogRequest) deepCopy() *_ModbusPDUGetComEventLogRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUGetComEventLogRequestCopy := &_ModbusPDUGetComEventLogRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUGetComEventLogRequestCopy
}

func (m *_ModbusPDUGetComEventLogRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
