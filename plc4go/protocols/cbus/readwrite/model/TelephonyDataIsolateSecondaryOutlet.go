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

// TelephonyDataIsolateSecondaryOutlet is the corresponding interface of TelephonyDataIsolateSecondaryOutlet
type TelephonyDataIsolateSecondaryOutlet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetIsolateStatus returns IsolateStatus (property field)
	GetIsolateStatus() byte
	// GetIsBehaveNormal returns IsBehaveNormal (virtual field)
	GetIsBehaveNormal() bool
	// GetIsToBeIsolated returns IsToBeIsolated (virtual field)
	GetIsToBeIsolated() bool
}

// TelephonyDataIsolateSecondaryOutletExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataIsolateSecondaryOutlet.
// This is useful for switch cases.
type TelephonyDataIsolateSecondaryOutletExactly interface {
	TelephonyDataIsolateSecondaryOutlet
	isTelephonyDataIsolateSecondaryOutlet() bool
}

// _TelephonyDataIsolateSecondaryOutlet is the data-structure of this message
type _TelephonyDataIsolateSecondaryOutlet struct {
	*_TelephonyData
	IsolateStatus byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataIsolateSecondaryOutlet) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetParent() TelephonyData {
	return m._TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataIsolateSecondaryOutlet) GetIsolateStatus() byte {
	return m.IsolateStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyDataIsolateSecondaryOutlet) GetIsBehaveNormal() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetIsolateStatus()) == (0x00)))
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetIsToBeIsolated() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetIsolateStatus()) == (0x01)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataIsolateSecondaryOutlet factory function for _TelephonyDataIsolateSecondaryOutlet
func NewTelephonyDataIsolateSecondaryOutlet(isolateStatus byte, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataIsolateSecondaryOutlet {
	_result := &_TelephonyDataIsolateSecondaryOutlet{
		IsolateStatus:  isolateStatus,
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataIsolateSecondaryOutlet(structType any) TelephonyDataIsolateSecondaryOutlet {
	if casted, ok := structType.(TelephonyDataIsolateSecondaryOutlet); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataIsolateSecondaryOutlet); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetTypeName() string {
	return "TelephonyDataIsolateSecondaryOutlet"
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (isolateStatus)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataIsolateSecondaryOutletParse(theBytes []byte) (TelephonyDataIsolateSecondaryOutlet, error) {
	return TelephonyDataIsolateSecondaryOutletParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func TelephonyDataIsolateSecondaryOutletParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyDataIsolateSecondaryOutlet, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataIsolateSecondaryOutlet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataIsolateSecondaryOutlet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (isolateStatus)
	_isolateStatus, _isolateStatusErr := readBuffer.ReadByte("isolateStatus")
	if _isolateStatusErr != nil {
		return nil, errors.Wrap(_isolateStatusErr, "Error parsing 'isolateStatus' field of TelephonyDataIsolateSecondaryOutlet")
	}
	isolateStatus := _isolateStatus

	// Virtual field
	_isBehaveNormal := bool((isolateStatus) == (0x00))
	isBehaveNormal := bool(_isBehaveNormal)
	_ = isBehaveNormal

	// Virtual field
	_isToBeIsolated := bool((isolateStatus) == (0x01))
	isToBeIsolated := bool(_isToBeIsolated)
	_ = isToBeIsolated

	if closeErr := readBuffer.CloseContext("TelephonyDataIsolateSecondaryOutlet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataIsolateSecondaryOutlet")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataIsolateSecondaryOutlet{
		_TelephonyData: &_TelephonyData{},
		IsolateStatus:  isolateStatus,
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataIsolateSecondaryOutlet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataIsolateSecondaryOutlet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataIsolateSecondaryOutlet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataIsolateSecondaryOutlet")
		}

		// Simple Field (isolateStatus)
		isolateStatus := byte(m.GetIsolateStatus())
		_isolateStatusErr := writeBuffer.WriteByte("isolateStatus", (isolateStatus))
		if _isolateStatusErr != nil {
			return errors.Wrap(_isolateStatusErr, "Error serializing 'isolateStatus' field")
		}
		// Virtual field
		if _isBehaveNormalErr := writeBuffer.WriteVirtual(ctx, "isBehaveNormal", m.GetIsBehaveNormal()); _isBehaveNormalErr != nil {
			return errors.Wrap(_isBehaveNormalErr, "Error serializing 'isBehaveNormal' field")
		}
		// Virtual field
		if _isToBeIsolatedErr := writeBuffer.WriteVirtual(ctx, "isToBeIsolated", m.GetIsToBeIsolated()); _isToBeIsolatedErr != nil {
			return errors.Wrap(_isToBeIsolatedErr, "Error serializing 'isToBeIsolated' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataIsolateSecondaryOutlet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataIsolateSecondaryOutlet")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataIsolateSecondaryOutlet) isTelephonyDataIsolateSecondaryOutlet() bool {
	return true
}

func (m *_TelephonyDataIsolateSecondaryOutlet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
