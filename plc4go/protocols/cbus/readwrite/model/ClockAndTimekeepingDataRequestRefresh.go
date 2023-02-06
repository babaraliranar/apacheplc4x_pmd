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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ClockAndTimekeepingDataRequestRefresh is the corresponding interface of ClockAndTimekeepingDataRequestRefresh
type ClockAndTimekeepingDataRequestRefresh interface {
	utils.LengthAware
	utils.Serializable
	ClockAndTimekeepingData
}

// ClockAndTimekeepingDataRequestRefreshExactly can be used when we want exactly this type and not a type which fulfills ClockAndTimekeepingDataRequestRefresh.
// This is useful for switch cases.
type ClockAndTimekeepingDataRequestRefreshExactly interface {
	ClockAndTimekeepingDataRequestRefresh
	isClockAndTimekeepingDataRequestRefresh() bool
}

// _ClockAndTimekeepingDataRequestRefresh is the data-structure of this message
type _ClockAndTimekeepingDataRequestRefresh struct {
	*_ClockAndTimekeepingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClockAndTimekeepingDataRequestRefresh) InitializeParent(parent ClockAndTimekeepingData, commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetParent() ClockAndTimekeepingData {
	return m._ClockAndTimekeepingData
}

// NewClockAndTimekeepingDataRequestRefresh factory function for _ClockAndTimekeepingDataRequestRefresh
func NewClockAndTimekeepingDataRequestRefresh(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) *_ClockAndTimekeepingDataRequestRefresh {
	_result := &_ClockAndTimekeepingDataRequestRefresh{
		_ClockAndTimekeepingData: NewClockAndTimekeepingData(commandTypeContainer, argument),
	}
	_result._ClockAndTimekeepingData._ClockAndTimekeepingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingDataRequestRefresh(structType interface{}) ClockAndTimekeepingDataRequestRefresh {
	if casted, ok := structType.(ClockAndTimekeepingDataRequestRefresh); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingDataRequestRefresh); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetTypeName() string {
	return "ClockAndTimekeepingDataRequestRefresh"
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClockAndTimekeepingDataRequestRefreshParse(theBytes []byte) (ClockAndTimekeepingDataRequestRefresh, error) {
	return ClockAndTimekeepingDataRequestRefreshParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ClockAndTimekeepingDataRequestRefreshParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClockAndTimekeepingDataRequestRefresh, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingDataRequestRefresh"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingDataRequestRefresh")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingDataRequestRefresh"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingDataRequestRefresh")
	}

	// Create a partially initialized instance
	_child := &_ClockAndTimekeepingDataRequestRefresh{
		_ClockAndTimekeepingData: &_ClockAndTimekeepingData{},
	}
	_child._ClockAndTimekeepingData._ClockAndTimekeepingDataChildRequirements = _child
	return _child, nil
}

func (m *_ClockAndTimekeepingDataRequestRefresh) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClockAndTimekeepingDataRequestRefresh) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClockAndTimekeepingDataRequestRefresh"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingDataRequestRefresh")
		}

		if popErr := writeBuffer.PopContext("ClockAndTimekeepingDataRequestRefresh"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingDataRequestRefresh")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClockAndTimekeepingDataRequestRefresh) isClockAndTimekeepingDataRequestRefresh() bool {
	return true
}

func (m *_ClockAndTimekeepingDataRequestRefresh) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
