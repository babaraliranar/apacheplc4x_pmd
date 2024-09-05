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

// BACnetEventSummary is the corresponding interface of BACnetEventSummary
type BACnetEventSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetAcknowledgedTransitions returns AcknowledgedTransitions (property field)
	GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged
	// GetEventTimestamps returns EventTimestamps (property field)
	GetEventTimestamps() BACnetEventTimestampsEnclosed
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetEventEnable returns EventEnable (property field)
	GetEventEnable() BACnetEventTransitionBitsTagged
	// GetEventPriorities returns EventPriorities (property field)
	GetEventPriorities() BACnetEventPriorities
	// IsBACnetEventSummary is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventSummary()
}

// _BACnetEventSummary is the data-structure of this message
type _BACnetEventSummary struct {
	ObjectIdentifier        BACnetContextTagObjectIdentifier
	EventState              BACnetEventStateTagged
	AcknowledgedTransitions BACnetEventTransitionBitsTagged
	EventTimestamps         BACnetEventTimestampsEnclosed
	NotifyType              BACnetNotifyTypeTagged
	EventEnable             BACnetEventTransitionBitsTagged
	EventPriorities         BACnetEventPriorities
}

var _ BACnetEventSummary = (*_BACnetEventSummary)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventSummary) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetEventSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetEventSummary) GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged {
	return m.AcknowledgedTransitions
}

func (m *_BACnetEventSummary) GetEventTimestamps() BACnetEventTimestampsEnclosed {
	return m.EventTimestamps
}

func (m *_BACnetEventSummary) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *_BACnetEventSummary) GetEventEnable() BACnetEventTransitionBitsTagged {
	return m.EventEnable
}

func (m *_BACnetEventSummary) GetEventPriorities() BACnetEventPriorities {
	return m.EventPriorities
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventSummary factory function for _BACnetEventSummary
func NewBACnetEventSummary(objectIdentifier BACnetContextTagObjectIdentifier, eventState BACnetEventStateTagged, acknowledgedTransitions BACnetEventTransitionBitsTagged, eventTimestamps BACnetEventTimestampsEnclosed, notifyType BACnetNotifyTypeTagged, eventEnable BACnetEventTransitionBitsTagged, eventPriorities BACnetEventPriorities) *_BACnetEventSummary {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetEventSummary must not be nil")
	}
	if eventState == nil {
		panic("eventState of type BACnetEventStateTagged for BACnetEventSummary must not be nil")
	}
	if acknowledgedTransitions == nil {
		panic("acknowledgedTransitions of type BACnetEventTransitionBitsTagged for BACnetEventSummary must not be nil")
	}
	if eventTimestamps == nil {
		panic("eventTimestamps of type BACnetEventTimestampsEnclosed for BACnetEventSummary must not be nil")
	}
	if notifyType == nil {
		panic("notifyType of type BACnetNotifyTypeTagged for BACnetEventSummary must not be nil")
	}
	if eventEnable == nil {
		panic("eventEnable of type BACnetEventTransitionBitsTagged for BACnetEventSummary must not be nil")
	}
	if eventPriorities == nil {
		panic("eventPriorities of type BACnetEventPriorities for BACnetEventSummary must not be nil")
	}
	return &_BACnetEventSummary{ObjectIdentifier: objectIdentifier, EventState: eventState, AcknowledgedTransitions: acknowledgedTransitions, EventTimestamps: eventTimestamps, NotifyType: notifyType, EventEnable: eventEnable, EventPriorities: eventPriorities}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventSummary(structType any) BACnetEventSummary {
	if casted, ok := structType.(BACnetEventSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventSummary) GetTypeName() string {
	return "BACnetEventSummary"
}

func (m *_BACnetEventSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (acknowledgedTransitions)
	lengthInBits += m.AcknowledgedTransitions.GetLengthInBits(ctx)

	// Simple field (eventTimestamps)
	lengthInBits += m.EventTimestamps.GetLengthInBits(ctx)

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits(ctx)

	// Simple field (eventEnable)
	lengthInBits += m.EventEnable.GetLengthInBits(ctx)

	// Simple field (eventPriorities)
	lengthInBits += m.EventPriorities.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventSummaryParse(ctx context.Context, theBytes []byte) (BACnetEventSummary, error) {
	return BACnetEventSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventSummaryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventSummary, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventSummary, error) {
		return BACnetEventSummaryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventSummary, error) {
	v, err := (&_BACnetEventSummary{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetEventSummary BACnetEventSummary, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	eventState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	acknowledgedTransitions, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "acknowledgedTransitions", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgedTransitions' field"))
	}
	m.AcknowledgedTransitions = acknowledgedTransitions

	eventTimestamps, err := ReadSimpleField[BACnetEventTimestampsEnclosed](ctx, "eventTimestamps", ReadComplex[BACnetEventTimestampsEnclosed](BACnetEventTimestampsEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventTimestamps' field"))
	}
	m.EventTimestamps = eventTimestamps

	notifyType, err := ReadSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", ReadComplex[BACnetNotifyTypeTagged](BACnetNotifyTypeTaggedParseWithBufferProducer((uint8)(uint8(4)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notifyType' field"))
	}
	m.NotifyType = notifyType

	eventEnable, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "eventEnable", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(5)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventEnable' field"))
	}
	m.EventEnable = eventEnable

	eventPriorities, err := ReadSimpleField[BACnetEventPriorities](ctx, "eventPriorities", ReadComplex[BACnetEventPriorities](BACnetEventPrioritiesParseWithBufferProducer((uint8)(uint8(6))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventPriorities' field"))
	}
	m.EventPriorities = eventPriorities

	if closeErr := readBuffer.CloseContext("BACnetEventSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventSummary")
	}

	return m, nil
}

func (m *_BACnetEventSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventSummary"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventSummary")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetEventStateTagged](ctx, "eventState", m.GetEventState(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventState' field")
	}

	if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "acknowledgedTransitions", m.GetAcknowledgedTransitions(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'acknowledgedTransitions' field")
	}

	if err := WriteSimpleField[BACnetEventTimestampsEnclosed](ctx, "eventTimestamps", m.GetEventTimestamps(), WriteComplex[BACnetEventTimestampsEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventTimestamps' field")
	}

	if err := WriteSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", m.GetNotifyType(), WriteComplex[BACnetNotifyTypeTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'notifyType' field")
	}

	if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "eventEnable", m.GetEventEnable(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventEnable' field")
	}

	if err := WriteSimpleField[BACnetEventPriorities](ctx, "eventPriorities", m.GetEventPriorities(), WriteComplex[BACnetEventPriorities](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventPriorities' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventSummary"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventSummary")
	}
	return nil
}

func (m *_BACnetEventSummary) IsBACnetEventSummary() {}

func (m *_BACnetEventSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
