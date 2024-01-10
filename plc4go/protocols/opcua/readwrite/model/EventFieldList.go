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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// EventFieldList is the corresponding interface of EventFieldList
type EventFieldList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetClientHandle returns ClientHandle (property field)
	GetClientHandle() uint32
	// GetNoOfEventFields returns NoOfEventFields (property field)
	GetNoOfEventFields() int32
	// GetEventFields returns EventFields (property field)
	GetEventFields() []Variant
}

// EventFieldListExactly can be used when we want exactly this type and not a type which fulfills EventFieldList.
// This is useful for switch cases.
type EventFieldListExactly interface {
	EventFieldList
	isEventFieldList() bool
}

// _EventFieldList is the data-structure of this message
type _EventFieldList struct {
	*_ExtensionObjectDefinition
	ClientHandle    uint32
	NoOfEventFields int32
	EventFields     []Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EventFieldList) GetIdentifier() string {
	return "919"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EventFieldList) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_EventFieldList) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EventFieldList) GetClientHandle() uint32 {
	return m.ClientHandle
}

func (m *_EventFieldList) GetNoOfEventFields() int32 {
	return m.NoOfEventFields
}

func (m *_EventFieldList) GetEventFields() []Variant {
	return m.EventFields
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEventFieldList factory function for _EventFieldList
func NewEventFieldList(clientHandle uint32, noOfEventFields int32, eventFields []Variant) *_EventFieldList {
	_result := &_EventFieldList{
		ClientHandle:               clientHandle,
		NoOfEventFields:            noOfEventFields,
		EventFields:                eventFields,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEventFieldList(structType any) EventFieldList {
	if casted, ok := structType.(EventFieldList); ok {
		return casted
	}
	if casted, ok := structType.(*EventFieldList); ok {
		return *casted
	}
	return nil
}

func (m *_EventFieldList) GetTypeName() string {
	return "EventFieldList"
}

func (m *_EventFieldList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (clientHandle)
	lengthInBits += 32

	// Simple field (noOfEventFields)
	lengthInBits += 32

	// Array field
	if len(m.EventFields) > 0 {
		for _curItem, element := range m.EventFields {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EventFields), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_EventFieldList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EventFieldListParse(ctx context.Context, theBytes []byte, identifier string) (EventFieldList, error) {
	return EventFieldListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func EventFieldListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (EventFieldList, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("EventFieldList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EventFieldList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (clientHandle)
	_clientHandle, _clientHandleErr := readBuffer.ReadUint32("clientHandle", 32)
	if _clientHandleErr != nil {
		return nil, errors.Wrap(_clientHandleErr, "Error parsing 'clientHandle' field of EventFieldList")
	}
	clientHandle := _clientHandle

	// Simple Field (noOfEventFields)
	_noOfEventFields, _noOfEventFieldsErr := readBuffer.ReadInt32("noOfEventFields", 32)
	if _noOfEventFieldsErr != nil {
		return nil, errors.Wrap(_noOfEventFieldsErr, "Error parsing 'noOfEventFields' field of EventFieldList")
	}
	noOfEventFields := _noOfEventFields

	// Array field (eventFields)
	if pullErr := readBuffer.PullContext("eventFields", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventFields")
	}
	// Count array
	eventFields := make([]Variant, utils.Max(noOfEventFields, 0))
	// This happens when the size is set conditional to 0
	if len(eventFields) == 0 {
		eventFields = nil
	}
	{
		_numItems := uint16(utils.Max(noOfEventFields, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := VariantParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'eventFields' field of EventFieldList")
			}
			eventFields[_curItem] = _item.(Variant)
		}
	}
	if closeErr := readBuffer.CloseContext("eventFields", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventFields")
	}

	if closeErr := readBuffer.CloseContext("EventFieldList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EventFieldList")
	}

	// Create a partially initialized instance
	_child := &_EventFieldList{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ClientHandle:               clientHandle,
		NoOfEventFields:            noOfEventFields,
		EventFields:                eventFields,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_EventFieldList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EventFieldList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EventFieldList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EventFieldList")
		}

		// Simple Field (clientHandle)
		clientHandle := uint32(m.GetClientHandle())
		_clientHandleErr := writeBuffer.WriteUint32("clientHandle", 32, uint32((clientHandle)))
		if _clientHandleErr != nil {
			return errors.Wrap(_clientHandleErr, "Error serializing 'clientHandle' field")
		}

		// Simple Field (noOfEventFields)
		noOfEventFields := int32(m.GetNoOfEventFields())
		_noOfEventFieldsErr := writeBuffer.WriteInt32("noOfEventFields", 32, int32((noOfEventFields)))
		if _noOfEventFieldsErr != nil {
			return errors.Wrap(_noOfEventFieldsErr, "Error serializing 'noOfEventFields' field")
		}

		// Array Field (eventFields)
		if pushErr := writeBuffer.PushContext("eventFields", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventFields")
		}
		for _curItem, _element := range m.GetEventFields() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetEventFields()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'eventFields' field")
			}
		}
		if popErr := writeBuffer.PopContext("eventFields", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventFields")
		}

		if popErr := writeBuffer.PopContext("EventFieldList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EventFieldList")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EventFieldList) isEventFieldList() bool {
	return true
}

func (m *_EventFieldList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
