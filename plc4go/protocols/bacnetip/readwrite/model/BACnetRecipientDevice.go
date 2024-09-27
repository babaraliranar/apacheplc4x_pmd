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

// BACnetRecipientDevice is the corresponding interface of BACnetRecipientDevice
type BACnetRecipientDevice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetRecipient
	// GetDeviceValue returns DeviceValue (property field)
	GetDeviceValue() BACnetContextTagObjectIdentifier
	// IsBACnetRecipientDevice is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipientDevice()
	// CreateBuilder creates a BACnetRecipientDeviceBuilder
	CreateBACnetRecipientDeviceBuilder() BACnetRecipientDeviceBuilder
}

// _BACnetRecipientDevice is the data-structure of this message
type _BACnetRecipientDevice struct {
	BACnetRecipientContract
	DeviceValue BACnetContextTagObjectIdentifier
}

var _ BACnetRecipientDevice = (*_BACnetRecipientDevice)(nil)
var _ BACnetRecipientRequirements = (*_BACnetRecipientDevice)(nil)

// NewBACnetRecipientDevice factory function for _BACnetRecipientDevice
func NewBACnetRecipientDevice(peekedTagHeader BACnetTagHeader, deviceValue BACnetContextTagObjectIdentifier) *_BACnetRecipientDevice {
	if deviceValue == nil {
		panic("deviceValue of type BACnetContextTagObjectIdentifier for BACnetRecipientDevice must not be nil")
	}
	_result := &_BACnetRecipientDevice{
		BACnetRecipientContract: NewBACnetRecipient(peekedTagHeader),
		DeviceValue:             deviceValue,
	}
	_result.BACnetRecipientContract.(*_BACnetRecipient)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRecipientDeviceBuilder is a builder for BACnetRecipientDevice
type BACnetRecipientDeviceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deviceValue BACnetContextTagObjectIdentifier) BACnetRecipientDeviceBuilder
	// WithDeviceValue adds DeviceValue (property field)
	WithDeviceValue(BACnetContextTagObjectIdentifier) BACnetRecipientDeviceBuilder
	// WithDeviceValueBuilder adds DeviceValue (property field) which is build by the builder
	WithDeviceValueBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetRecipientDeviceBuilder
	// Build builds the BACnetRecipientDevice or returns an error if something is wrong
	Build() (BACnetRecipientDevice, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRecipientDevice
}

// NewBACnetRecipientDeviceBuilder() creates a BACnetRecipientDeviceBuilder
func NewBACnetRecipientDeviceBuilder() BACnetRecipientDeviceBuilder {
	return &_BACnetRecipientDeviceBuilder{_BACnetRecipientDevice: new(_BACnetRecipientDevice)}
}

type _BACnetRecipientDeviceBuilder struct {
	*_BACnetRecipientDevice

	err *utils.MultiError
}

var _ (BACnetRecipientDeviceBuilder) = (*_BACnetRecipientDeviceBuilder)(nil)

func (m *_BACnetRecipientDeviceBuilder) WithMandatoryFields(deviceValue BACnetContextTagObjectIdentifier) BACnetRecipientDeviceBuilder {
	return m.WithDeviceValue(deviceValue)
}

func (m *_BACnetRecipientDeviceBuilder) WithDeviceValue(deviceValue BACnetContextTagObjectIdentifier) BACnetRecipientDeviceBuilder {
	m.DeviceValue = deviceValue
	return m
}

func (m *_BACnetRecipientDeviceBuilder) WithDeviceValueBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetRecipientDeviceBuilder {
	builder := builderSupplier(m.DeviceValue.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.DeviceValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetRecipientDeviceBuilder) Build() (BACnetRecipientDevice, error) {
	if m.DeviceValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'deviceValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetRecipientDevice.deepCopy(), nil
}

func (m *_BACnetRecipientDeviceBuilder) MustBuild() BACnetRecipientDevice {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetRecipientDeviceBuilder) DeepCopy() any {
	return m.CreateBACnetRecipientDeviceBuilder()
}

// CreateBACnetRecipientDeviceBuilder creates a BACnetRecipientDeviceBuilder
func (m *_BACnetRecipientDevice) CreateBACnetRecipientDeviceBuilder() BACnetRecipientDeviceBuilder {
	if m == nil {
		return NewBACnetRecipientDeviceBuilder()
	}
	return &_BACnetRecipientDeviceBuilder{_BACnetRecipientDevice: m.deepCopy()}
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

func (m *_BACnetRecipientDevice) GetParent() BACnetRecipientContract {
	return m.BACnetRecipientContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientDevice) GetDeviceValue() BACnetContextTagObjectIdentifier {
	return m.DeviceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetRecipientDevice(structType any) BACnetRecipientDevice {
	if casted, ok := structType.(BACnetRecipientDevice); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientDevice); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientDevice) GetTypeName() string {
	return "BACnetRecipientDevice"
}

func (m *_BACnetRecipientDevice) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetRecipientContract.(*_BACnetRecipient).getLengthInBits(ctx))

	// Simple field (deviceValue)
	lengthInBits += m.DeviceValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientDevice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetRecipientDevice) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetRecipient) (__bACnetRecipientDevice BACnetRecipientDevice, err error) {
	m.BACnetRecipientContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceValue, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "deviceValue", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceValue' field"))
	}
	m.DeviceValue = deviceValue

	if closeErr := readBuffer.CloseContext("BACnetRecipientDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientDevice")
	}

	return m, nil
}

func (m *_BACnetRecipientDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientDevice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetRecipientDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetRecipientDevice")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "deviceValue", m.GetDeviceValue(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetRecipientDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetRecipientDevice")
		}
		return nil
	}
	return m.BACnetRecipientContract.(*_BACnetRecipient).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetRecipientDevice) IsBACnetRecipientDevice() {}

func (m *_BACnetRecipientDevice) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRecipientDevice) deepCopy() *_BACnetRecipientDevice {
	if m == nil {
		return nil
	}
	_BACnetRecipientDeviceCopy := &_BACnetRecipientDevice{
		m.BACnetRecipientContract.(*_BACnetRecipient).deepCopy(),
		m.DeviceValue.DeepCopy().(BACnetContextTagObjectIdentifier),
	}
	m.BACnetRecipientContract.(*_BACnetRecipient)._SubType = m
	return _BACnetRecipientDeviceCopy
}

func (m *_BACnetRecipientDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
