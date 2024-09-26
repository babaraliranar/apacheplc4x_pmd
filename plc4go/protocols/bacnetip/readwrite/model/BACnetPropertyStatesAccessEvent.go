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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesAccessEvent is the corresponding interface of BACnetPropertyStatesAccessEvent
type BACnetPropertyStatesAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetAccessEvent returns AccessEvent (property field)
	GetAccessEvent() BACnetAccessEventTagged
	// IsBACnetPropertyStatesAccessEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesAccessEvent()
}

// _BACnetPropertyStatesAccessEvent is the data-structure of this message
type _BACnetPropertyStatesAccessEvent struct {
	BACnetPropertyStatesContract
	AccessEvent BACnetAccessEventTagged
}

var _ BACnetPropertyStatesAccessEvent = (*_BACnetPropertyStatesAccessEvent)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesAccessEvent)(nil)

// NewBACnetPropertyStatesAccessEvent factory function for _BACnetPropertyStatesAccessEvent
func NewBACnetPropertyStatesAccessEvent(peekedTagHeader BACnetTagHeader, accessEvent BACnetAccessEventTagged) *_BACnetPropertyStatesAccessEvent {
	if accessEvent == nil {
		panic("accessEvent of type BACnetAccessEventTagged for BACnetPropertyStatesAccessEvent must not be nil")
	}
	_result := &_BACnetPropertyStatesAccessEvent{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		AccessEvent:                  accessEvent,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesAccessEvent) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAccessEvent) GetAccessEvent() BACnetAccessEventTagged {
	return m.AccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAccessEvent(structType any) BACnetPropertyStatesAccessEvent {
	if casted, ok := structType.(BACnetPropertyStatesAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAccessEvent) GetTypeName() string {
	return "BACnetPropertyStatesAccessEvent"
}

func (m *_BACnetPropertyStatesAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (accessEvent)
	lengthInBits += m.AccessEvent.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesAccessEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesAccessEvent BACnetPropertyStatesAccessEvent, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessEvent, err := ReadSimpleField[BACnetAccessEventTagged](ctx, "accessEvent", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEvent' field"))
	}
	m.AccessEvent = accessEvent

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAccessEvent")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAccessEvent")
		}

		if err := WriteSimpleField[BACnetAccessEventTagged](ctx, "accessEvent", m.GetAccessEvent(), WriteComplex[BACnetAccessEventTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEvent' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAccessEvent")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAccessEvent) IsBACnetPropertyStatesAccessEvent() {}

func (m *_BACnetPropertyStatesAccessEvent) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesAccessEvent) deepCopy() *_BACnetPropertyStatesAccessEvent {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesAccessEventCopy := &_BACnetPropertyStatesAccessEvent{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.AccessEvent.DeepCopy().(BACnetAccessEventTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesAccessEventCopy
}

func (m *_BACnetPropertyStatesAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
