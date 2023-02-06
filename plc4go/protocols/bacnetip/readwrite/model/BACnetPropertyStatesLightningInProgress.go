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

// BACnetPropertyStatesLightningInProgress is the corresponding interface of BACnetPropertyStatesLightningInProgress
type BACnetPropertyStatesLightningInProgress interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLightningInProgress returns LightningInProgress (property field)
	GetLightningInProgress() BACnetLightingInProgressTagged
}

// BACnetPropertyStatesLightningInProgressExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLightningInProgress.
// This is useful for switch cases.
type BACnetPropertyStatesLightningInProgressExactly interface {
	BACnetPropertyStatesLightningInProgress
	isBACnetPropertyStatesLightningInProgress() bool
}

// _BACnetPropertyStatesLightningInProgress is the data-structure of this message
type _BACnetPropertyStatesLightningInProgress struct {
	*_BACnetPropertyStates
	LightningInProgress BACnetLightingInProgressTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLightningInProgress) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLightningInProgress) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningInProgress) GetLightningInProgress() BACnetLightingInProgressTagged {
	return m.LightningInProgress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLightningInProgress factory function for _BACnetPropertyStatesLightningInProgress
func NewBACnetPropertyStatesLightningInProgress(lightningInProgress BACnetLightingInProgressTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLightningInProgress {
	_result := &_BACnetPropertyStatesLightningInProgress{
		LightningInProgress:   lightningInProgress,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningInProgress(structType interface{}) BACnetPropertyStatesLightningInProgress {
	if casted, ok := structType.(BACnetPropertyStatesLightningInProgress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningInProgress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningInProgress) GetTypeName() string {
	return "BACnetPropertyStatesLightningInProgress"
}

func (m *_BACnetPropertyStatesLightningInProgress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lightningInProgress)
	lengthInBits += m.LightningInProgress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningInProgress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesLightningInProgressParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesLightningInProgress, error) {
	return BACnetPropertyStatesLightningInProgressParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesLightningInProgressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLightningInProgress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningInProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningInProgress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightningInProgress)
	if pullErr := readBuffer.PullContext("lightningInProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightningInProgress")
	}
	_lightningInProgress, _lightningInProgressErr := BACnetLightingInProgressTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lightningInProgressErr != nil {
		return nil, errors.Wrap(_lightningInProgressErr, "Error parsing 'lightningInProgress' field of BACnetPropertyStatesLightningInProgress")
	}
	lightningInProgress := _lightningInProgress.(BACnetLightingInProgressTagged)
	if closeErr := readBuffer.CloseContext("lightningInProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightningInProgress")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningInProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningInProgress")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLightningInProgress{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LightningInProgress:   lightningInProgress,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLightningInProgress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningInProgress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningInProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningInProgress")
		}

		// Simple Field (lightningInProgress)
		if pushErr := writeBuffer.PushContext("lightningInProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightningInProgress")
		}
		_lightningInProgressErr := writeBuffer.WriteSerializable(ctx, m.GetLightningInProgress())
		if popErr := writeBuffer.PopContext("lightningInProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightningInProgress")
		}
		if _lightningInProgressErr != nil {
			return errors.Wrap(_lightningInProgressErr, "Error serializing 'lightningInProgress' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningInProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningInProgress")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningInProgress) isBACnetPropertyStatesLightningInProgress() bool {
	return true
}

func (m *_BACnetPropertyStatesLightningInProgress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
