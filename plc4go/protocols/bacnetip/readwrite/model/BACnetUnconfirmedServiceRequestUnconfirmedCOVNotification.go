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

// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() BACnetContextTagUnsignedInteger
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
	// IsBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification()
}

// _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification struct {
	BACnetUnconfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	LifetimeInSeconds           BACnetContextTagUnsignedInteger
	ListOfValues                BACnetPropertyValues
}

var _ BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification = (*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLifetimeInSeconds() BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification factory function for _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, lifetimeInSeconds BACnetContextTagUnsignedInteger, listOfValues BACnetPropertyValues, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification must not be nil")
	}
	if initiatingDeviceIdentifier == nil {
		panic("initiatingDeviceIdentifier of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification must not be nil")
	}
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification must not be nil")
	}
	if lifetimeInSeconds == nil {
		panic("lifetimeInSeconds of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification must not be nil")
	}
	if listOfValues == nil {
		panic("listOfValues of type BACnetPropertyValues for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:             subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:              initiatingDeviceIdentifier,
		MonitoredObjectIdentifier:               monitoredObjectIdentifier,
		LifetimeInSeconds:                       lifetimeInSeconds,
		ListOfValues:                            listOfValues,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification(structType any) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.GetLengthInBits(ctx)

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestUnconfirmedCOVNotification BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification")
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

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	lifetimeInSeconds, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetimeInSeconds' field"))
	}
	m.LifetimeInSeconds = lifetimeInSeconds

	listOfValues, err := ReadSimpleField[BACnetPropertyValues](ctx, "listOfValues", ReadComplex[BACnetPropertyValues](BACnetPropertyValuesParseWithBufferProducer((uint8)(uint8(4)), (BACnetObjectType)(monitoredObjectIdentifier.GetObjectType())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}
	m.ListOfValues = listOfValues

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", m.GetLifetimeInSeconds(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifetimeInSeconds' field")
		}

		if err := WriteSimpleField[BACnetPropertyValues](ctx, "listOfValues", m.GetListOfValues(), WriteComplex[BACnetPropertyValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) IsBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification() {
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
