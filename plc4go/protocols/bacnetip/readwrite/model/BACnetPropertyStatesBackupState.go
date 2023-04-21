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

// BACnetPropertyStatesBackupState is the corresponding interface of BACnetPropertyStatesBackupState
type BACnetPropertyStatesBackupState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetBackupState returns BackupState (property field)
	GetBackupState() BACnetBackupStateTagged
}

// BACnetPropertyStatesBackupStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesBackupState.
// This is useful for switch cases.
type BACnetPropertyStatesBackupStateExactly interface {
	BACnetPropertyStatesBackupState
	isBACnetPropertyStatesBackupState() bool
}

// _BACnetPropertyStatesBackupState is the data-structure of this message
type _BACnetPropertyStatesBackupState struct {
	*_BACnetPropertyStates
	BackupState BACnetBackupStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesBackupState) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesBackupState) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBackupState) GetBackupState() BACnetBackupStateTagged {
	return m.BackupState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBackupState factory function for _BACnetPropertyStatesBackupState
func NewBACnetPropertyStatesBackupState(backupState BACnetBackupStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesBackupState {
	_result := &_BACnetPropertyStatesBackupState{
		BackupState:           backupState,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBackupState(structType any) BACnetPropertyStatesBackupState {
	if casted, ok := structType.(BACnetPropertyStatesBackupState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBackupState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBackupState) GetTypeName() string {
	return "BACnetPropertyStatesBackupState"
}

func (m *_BACnetPropertyStatesBackupState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (backupState)
	lengthInBits += m.BackupState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBackupState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesBackupStateParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesBackupState, error) {
	return BACnetPropertyStatesBackupStateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesBackupStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesBackupState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBackupState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBackupState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (backupState)
	if pullErr := readBuffer.PullContext("backupState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for backupState")
	}
	_backupState, _backupStateErr := BACnetBackupStateTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _backupStateErr != nil {
		return nil, errors.Wrap(_backupStateErr, "Error parsing 'backupState' field of BACnetPropertyStatesBackupState")
	}
	backupState := _backupState.(BACnetBackupStateTagged)
	if closeErr := readBuffer.CloseContext("backupState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for backupState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBackupState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBackupState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesBackupState{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		BackupState:           backupState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesBackupState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBackupState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBackupState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBackupState")
		}

		// Simple Field (backupState)
		if pushErr := writeBuffer.PushContext("backupState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for backupState")
		}
		_backupStateErr := writeBuffer.WriteSerializable(ctx, m.GetBackupState())
		if popErr := writeBuffer.PopContext("backupState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for backupState")
		}
		if _backupStateErr != nil {
			return errors.Wrap(_backupStateErr, "Error serializing 'backupState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBackupState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBackupState")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBackupState) isBACnetPropertyStatesBackupState() bool {
	return true
}

func (m *_BACnetPropertyStatesBackupState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
