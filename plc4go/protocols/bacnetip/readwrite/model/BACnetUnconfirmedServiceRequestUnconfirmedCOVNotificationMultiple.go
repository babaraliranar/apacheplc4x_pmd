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

// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetListOfCovNotifications returns ListOfCovNotifications (property field)
	GetListOfCovNotifications() ListOfCovNotificationsList
	// IsBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	CreateBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder() BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
}

// _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple struct {
	BACnetUnconfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	TimeRemaining               BACnetContextTagUnsignedInteger
	Timestamp                   BACnetTimeStampEnclosed
	ListOfCovNotifications      ListOfCovNotificationsList
}

var _ BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple = (*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple)(nil)

// NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple factory function for _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, timestamp BACnetTimeStampEnclosed, listOfCovNotifications ListOfCovNotificationsList, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple must not be nil")
	}
	if initiatingDeviceIdentifier == nil {
		panic("initiatingDeviceIdentifier of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple must not be nil")
	}
	if timeRemaining == nil {
		panic("timeRemaining of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple must not be nil")
	}
	if listOfCovNotifications == nil {
		panic("listOfCovNotifications of type ListOfCovNotificationsList for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:             subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:              initiatingDeviceIdentifier,
		TimeRemaining:                           timeRemaining,
		Timestamp:                               timestamp,
		ListOfCovNotifications:                  listOfCovNotifications,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder is a builder for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, listOfCovNotifications ListOfCovNotificationsList) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithSubscriberProcessIdentifier adds SubscriberProcessIdentifier (property field)
	WithSubscriberProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithSubscriberProcessIdentifierBuilder adds SubscriberProcessIdentifier (property field) which is build by the builder
	WithSubscriberProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithInitiatingDeviceIdentifier adds InitiatingDeviceIdentifier (property field)
	WithInitiatingDeviceIdentifier(BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithInitiatingDeviceIdentifierBuilder adds InitiatingDeviceIdentifier (property field) which is build by the builder
	WithInitiatingDeviceIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithTimeRemaining adds TimeRemaining (property field)
	WithTimeRemaining(BACnetContextTagUnsignedInteger) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithTimeRemainingBuilder adds TimeRemaining (property field) which is build by the builder
	WithTimeRemainingBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithTimestamp adds Timestamp (property field)
	WithOptionalTimestamp(BACnetTimeStampEnclosed) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithOptionalTimestampBuilder adds Timestamp (property field) which is build by the builder
	WithOptionalTimestampBuilder(func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithListOfCovNotifications adds ListOfCovNotifications (property field)
	WithListOfCovNotifications(ListOfCovNotificationsList) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// WithListOfCovNotificationsBuilder adds ListOfCovNotifications (property field) which is build by the builder
	WithListOfCovNotificationsBuilder(func(ListOfCovNotificationsListBuilder) ListOfCovNotificationsListBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetUnconfirmedServiceRequestBuilder
	// Build builds the BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
}

// NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder() creates a BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder() BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	return &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder{_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple: new(_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple)}
}

type _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder struct {
	*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple

	parentBuilder *_BACnetUnconfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) = (*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder)(nil)

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) setParent(contract BACnetUnconfirmedServiceRequestContract) {
	b.BACnetUnconfirmedServiceRequestContract = contract
	contract.(*_BACnetUnconfirmedServiceRequest)._SubType = b._BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, listOfCovNotifications ListOfCovNotificationsList) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	return b.WithSubscriberProcessIdentifier(subscriberProcessIdentifier).WithInitiatingDeviceIdentifier(initiatingDeviceIdentifier).WithTimeRemaining(timeRemaining).WithListOfCovNotifications(listOfCovNotifications)
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithSubscriberProcessIdentifier(subscriberProcessIdentifier BACnetContextTagUnsignedInteger) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	b.SubscriberProcessIdentifier = subscriberProcessIdentifier
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithSubscriberProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.SubscriberProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.SubscriberProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithInitiatingDeviceIdentifier(initiatingDeviceIdentifier BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	b.InitiatingDeviceIdentifier = initiatingDeviceIdentifier
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithInitiatingDeviceIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.InitiatingDeviceIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.InitiatingDeviceIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithTimeRemaining(timeRemaining BACnetContextTagUnsignedInteger) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	b.TimeRemaining = timeRemaining
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithTimeRemainingBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.TimeRemaining.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeRemaining, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithOptionalTimestamp(timestamp BACnetTimeStampEnclosed) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	b.Timestamp = timestamp
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithOptionalTimestampBuilder(builderSupplier func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.Timestamp.CreateBACnetTimeStampEnclosedBuilder())
	var err error
	b.Timestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithListOfCovNotifications(listOfCovNotifications ListOfCovNotificationsList) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	b.ListOfCovNotifications = listOfCovNotifications
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) WithListOfCovNotificationsBuilder(builderSupplier func(ListOfCovNotificationsListBuilder) ListOfCovNotificationsListBuilder) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.ListOfCovNotifications.CreateListOfCovNotificationsListBuilder())
	var err error
	b.ListOfCovNotifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ListOfCovNotificationsListBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) Build() (BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple, error) {
	if b.SubscriberProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'subscriberProcessIdentifier' not set"))
	}
	if b.InitiatingDeviceIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'initiatingDeviceIdentifier' not set"))
	}
	if b.TimeRemaining == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeRemaining' not set"))
	}
	if b.ListOfCovNotifications == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfCovNotifications' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple.deepCopy(), nil
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) MustBuild() BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) Done() BACnetUnconfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetUnconfirmedServiceRequestBuilder().(*_BACnetUnconfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) buildForBACnetUnconfirmedServiceRequest() (BACnetUnconfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder().(*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder creates a BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder
func (b *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) CreateBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder() BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder {
	if b == nil {
		return NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleBuilder{_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetListOfCovNotifications() ListOfCovNotificationsList {
	return m.ListOfCovNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(structType any) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (timestamp)
	if m.Timestamp != nil {
		lengthInBits += m.Timestamp.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovNotifications)
	lengthInBits += m.ListOfCovNotifications.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}
	m.InitiatingDeviceIdentifier = initiatingDeviceIdentifier

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}
	m.TimeRemaining = timeRemaining

	var timestamp BACnetTimeStampEnclosed
	_timestamp, err := ReadOptionalField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	if _timestamp != nil {
		timestamp = *_timestamp
		m.Timestamp = timestamp
	}

	listOfCovNotifications, err := ReadSimpleField[ListOfCovNotificationsList](ctx, "listOfCovNotifications", ReadComplex[ListOfCovNotificationsList](ListOfCovNotificationsListParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovNotifications' field"))
	}
	m.ListOfCovNotifications = listOfCovNotifications

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeRemaining' field")
		}

		if err := WriteOptionalField[BACnetTimeStampEnclosed](ctx, "timestamp", GetRef(m.GetTimestamp()), WriteComplex[BACnetTimeStampEnclosed](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[ListOfCovNotificationsList](ctx, "listOfCovNotifications", m.GetListOfCovNotifications(), WriteComplex[ListOfCovNotificationsList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfCovNotifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) IsBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple() {
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) deepCopy() *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleCopy := &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.SubscriberProcessIdentifier),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.InitiatingDeviceIdentifier),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeRemaining),
		utils.DeepCopy[BACnetTimeStampEnclosed](m.Timestamp),
		utils.DeepCopy[ListOfCovNotificationsList](m.ListOfCovNotifications),
	}
	_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleCopy.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleCopy
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
