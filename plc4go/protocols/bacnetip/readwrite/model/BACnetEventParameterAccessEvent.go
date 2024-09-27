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

// BACnetEventParameterAccessEvent is the corresponding interface of BACnetEventParameterAccessEvent
type BACnetEventParameterAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAccessEvents returns ListOfAccessEvents (property field)
	GetListOfAccessEvents() BACnetEventParameterAccessEventListOfAccessEvents
	// GetAccessEventTimeReference returns AccessEventTimeReference (property field)
	GetAccessEventTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterAccessEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterAccessEvent()
	// CreateBuilder creates a BACnetEventParameterAccessEventBuilder
	CreateBACnetEventParameterAccessEventBuilder() BACnetEventParameterAccessEventBuilder
}

// _BACnetEventParameterAccessEvent is the data-structure of this message
type _BACnetEventParameterAccessEvent struct {
	BACnetEventParameterContract
	OpeningTag               BACnetOpeningTag
	ListOfAccessEvents       BACnetEventParameterAccessEventListOfAccessEvents
	AccessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag               BACnetClosingTag
}

var _ BACnetEventParameterAccessEvent = (*_BACnetEventParameterAccessEvent)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterAccessEvent)(nil)

// NewBACnetEventParameterAccessEvent factory function for _BACnetEventParameterAccessEvent
func NewBACnetEventParameterAccessEvent(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, listOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents, accessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetEventParameterAccessEvent {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterAccessEvent must not be nil")
	}
	if listOfAccessEvents == nil {
		panic("listOfAccessEvents of type BACnetEventParameterAccessEventListOfAccessEvents for BACnetEventParameterAccessEvent must not be nil")
	}
	if accessEventTimeReference == nil {
		panic("accessEventTimeReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetEventParameterAccessEvent must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterAccessEvent must not be nil")
	}
	_result := &_BACnetEventParameterAccessEvent{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		ListOfAccessEvents:           listOfAccessEvents,
		AccessEventTimeReference:     accessEventTimeReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterAccessEventBuilder is a builder for BACnetEventParameterAccessEvent
type BACnetEventParameterAccessEventBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents, accessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterAccessEventBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterAccessEventBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterAccessEventBuilder
	// WithListOfAccessEvents adds ListOfAccessEvents (property field)
	WithListOfAccessEvents(BACnetEventParameterAccessEventListOfAccessEvents) BACnetEventParameterAccessEventBuilder
	// WithListOfAccessEventsBuilder adds ListOfAccessEvents (property field) which is build by the builder
	WithListOfAccessEventsBuilder(func(BACnetEventParameterAccessEventListOfAccessEventsBuilder) BACnetEventParameterAccessEventListOfAccessEventsBuilder) BACnetEventParameterAccessEventBuilder
	// WithAccessEventTimeReference adds AccessEventTimeReference (property field)
	WithAccessEventTimeReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterAccessEventBuilder
	// WithAccessEventTimeReferenceBuilder adds AccessEventTimeReference (property field) which is build by the builder
	WithAccessEventTimeReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterAccessEventBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterAccessEventBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterAccessEventBuilder
	// Build builds the BACnetEventParameterAccessEvent or returns an error if something is wrong
	Build() (BACnetEventParameterAccessEvent, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterAccessEvent
}

// NewBACnetEventParameterAccessEventBuilder() creates a BACnetEventParameterAccessEventBuilder
func NewBACnetEventParameterAccessEventBuilder() BACnetEventParameterAccessEventBuilder {
	return &_BACnetEventParameterAccessEventBuilder{_BACnetEventParameterAccessEvent: new(_BACnetEventParameterAccessEvent)}
}

type _BACnetEventParameterAccessEventBuilder struct {
	*_BACnetEventParameterAccessEvent

	err *utils.MultiError
}

var _ (BACnetEventParameterAccessEventBuilder) = (*_BACnetEventParameterAccessEventBuilder)(nil)

func (m *_BACnetEventParameterAccessEventBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents, accessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterAccessEventBuilder {
	return m.WithOpeningTag(openingTag).WithListOfAccessEvents(listOfAccessEvents).WithAccessEventTimeReference(accessEventTimeReference).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterAccessEventBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterAccessEventBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterAccessEventBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithListOfAccessEvents(listOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents) BACnetEventParameterAccessEventBuilder {
	m.ListOfAccessEvents = listOfAccessEvents
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithListOfAccessEventsBuilder(builderSupplier func(BACnetEventParameterAccessEventListOfAccessEventsBuilder) BACnetEventParameterAccessEventListOfAccessEventsBuilder) BACnetEventParameterAccessEventBuilder {
	builder := builderSupplier(m.ListOfAccessEvents.CreateBACnetEventParameterAccessEventListOfAccessEventsBuilder())
	var err error
	m.ListOfAccessEvents, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEventParameterAccessEventListOfAccessEventsBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithAccessEventTimeReference(accessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterAccessEventBuilder {
	m.AccessEventTimeReference = accessEventTimeReference
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithAccessEventTimeReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterAccessEventBuilder {
	builder := builderSupplier(m.AccessEventTimeReference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	m.AccessEventTimeReference, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterAccessEventBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterAccessEventBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterAccessEventBuilder) Build() (BACnetEventParameterAccessEvent, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.ListOfAccessEvents == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'listOfAccessEvents' not set"))
	}
	if m.AccessEventTimeReference == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'accessEventTimeReference' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetEventParameterAccessEvent.deepCopy(), nil
}

func (m *_BACnetEventParameterAccessEventBuilder) MustBuild() BACnetEventParameterAccessEvent {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterAccessEventBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterAccessEventBuilder()
}

// CreateBACnetEventParameterAccessEventBuilder creates a BACnetEventParameterAccessEventBuilder
func (m *_BACnetEventParameterAccessEvent) CreateBACnetEventParameterAccessEventBuilder() BACnetEventParameterAccessEventBuilder {
	if m == nil {
		return NewBACnetEventParameterAccessEventBuilder()
	}
	return &_BACnetEventParameterAccessEventBuilder{_BACnetEventParameterAccessEvent: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterAccessEvent) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterAccessEvent) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterAccessEvent) GetListOfAccessEvents() BACnetEventParameterAccessEventListOfAccessEvents {
	return m.ListOfAccessEvents
}

func (m *_BACnetEventParameterAccessEvent) GetAccessEventTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.AccessEventTimeReference
}

func (m *_BACnetEventParameterAccessEvent) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterAccessEvent(structType any) BACnetEventParameterAccessEvent {
	if casted, ok := structType.(BACnetEventParameterAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterAccessEvent) GetTypeName() string {
	return "BACnetEventParameterAccessEvent"
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfAccessEvents)
	lengthInBits += m.ListOfAccessEvents.GetLengthInBits(ctx)

	// Simple field (accessEventTimeReference)
	lengthInBits += m.AccessEventTimeReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterAccessEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterAccessEvent BACnetEventParameterAccessEvent, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(13))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfAccessEvents, err := ReadSimpleField[BACnetEventParameterAccessEventListOfAccessEvents](ctx, "listOfAccessEvents", ReadComplex[BACnetEventParameterAccessEventListOfAccessEvents](BACnetEventParameterAccessEventListOfAccessEventsParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAccessEvents' field"))
	}
	m.ListOfAccessEvents = listOfAccessEvents

	accessEventTimeReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "accessEventTimeReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEventTimeReference' field"))
	}
	m.AccessEventTimeReference = accessEventTimeReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(13))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterAccessEvent")
	}

	return m, nil
}

func (m *_BACnetEventParameterAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterAccessEvent")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetEventParameterAccessEventListOfAccessEvents](ctx, "listOfAccessEvents", m.GetListOfAccessEvents(), WriteComplex[BACnetEventParameterAccessEventListOfAccessEvents](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfAccessEvents' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "accessEventTimeReference", m.GetAccessEventTimeReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEventTimeReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterAccessEvent")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterAccessEvent) IsBACnetEventParameterAccessEvent() {}

func (m *_BACnetEventParameterAccessEvent) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterAccessEvent) deepCopy() *_BACnetEventParameterAccessEvent {
	if m == nil {
		return nil
	}
	_BACnetEventParameterAccessEventCopy := &_BACnetEventParameterAccessEvent{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ListOfAccessEvents.DeepCopy().(BACnetEventParameterAccessEventListOfAccessEvents),
		m.AccessEventTimeReference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterAccessEventCopy
}

func (m *_BACnetEventParameterAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
