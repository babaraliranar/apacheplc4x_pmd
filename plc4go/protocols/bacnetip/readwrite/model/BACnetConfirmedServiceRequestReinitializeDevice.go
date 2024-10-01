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

// BACnetConfirmedServiceRequestReinitializeDevice is the corresponding interface of BACnetConfirmedServiceRequestReinitializeDevice
type BACnetConfirmedServiceRequestReinitializeDevice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetReinitializedStateOfDevice returns ReinitializedStateOfDevice (property field)
	GetReinitializedStateOfDevice() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	// GetPassword returns Password (property field)
	GetPassword() BACnetContextTagCharacterString
	// IsBACnetConfirmedServiceRequestReinitializeDevice is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReinitializeDevice()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReinitializeDeviceBuilder
	CreateBACnetConfirmedServiceRequestReinitializeDeviceBuilder() BACnetConfirmedServiceRequestReinitializeDeviceBuilder
}

// _BACnetConfirmedServiceRequestReinitializeDevice is the data-structure of this message
type _BACnetConfirmedServiceRequestReinitializeDevice struct {
	BACnetConfirmedServiceRequestContract
	ReinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	Password                   BACnetContextTagCharacterString
}

var _ BACnetConfirmedServiceRequestReinitializeDevice = (*_BACnetConfirmedServiceRequestReinitializeDevice)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestReinitializeDevice)(nil)

// NewBACnetConfirmedServiceRequestReinitializeDevice factory function for _BACnetConfirmedServiceRequestReinitializeDevice
func NewBACnetConfirmedServiceRequestReinitializeDevice(reinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, password BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReinitializeDevice {
	if reinitializedStateOfDevice == nil {
		panic("reinitializedStateOfDevice of type BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged for BACnetConfirmedServiceRequestReinitializeDevice must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestReinitializeDevice{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ReinitializedStateOfDevice:            reinitializedStateOfDevice,
		Password:                              password,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReinitializeDeviceBuilder is a builder for BACnetConfirmedServiceRequestReinitializeDevice
type BACnetConfirmedServiceRequestReinitializeDeviceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) BACnetConfirmedServiceRequestReinitializeDeviceBuilder
	// WithReinitializedStateOfDevice adds ReinitializedStateOfDevice (property field)
	WithReinitializedStateOfDevice(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) BACnetConfirmedServiceRequestReinitializeDeviceBuilder
	// WithReinitializedStateOfDeviceBuilder adds ReinitializedStateOfDevice (property field) which is build by the builder
	WithReinitializedStateOfDeviceBuilder(func(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) BACnetConfirmedServiceRequestReinitializeDeviceBuilder
	// WithPassword adds Password (property field)
	WithOptionalPassword(BACnetContextTagCharacterString) BACnetConfirmedServiceRequestReinitializeDeviceBuilder
	// WithOptionalPasswordBuilder adds Password (property field) which is build by the builder
	WithOptionalPasswordBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestReinitializeDeviceBuilder
	// Build builds the BACnetConfirmedServiceRequestReinitializeDevice or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReinitializeDevice, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReinitializeDevice
}

// NewBACnetConfirmedServiceRequestReinitializeDeviceBuilder() creates a BACnetConfirmedServiceRequestReinitializeDeviceBuilder
func NewBACnetConfirmedServiceRequestReinitializeDeviceBuilder() BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	return &_BACnetConfirmedServiceRequestReinitializeDeviceBuilder{_BACnetConfirmedServiceRequestReinitializeDevice: new(_BACnetConfirmedServiceRequestReinitializeDevice)}
}

type _BACnetConfirmedServiceRequestReinitializeDeviceBuilder struct {
	*_BACnetConfirmedServiceRequestReinitializeDevice

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReinitializeDeviceBuilder) = (*_BACnetConfirmedServiceRequestReinitializeDeviceBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) WithMandatoryFields(reinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	return b.WithReinitializedStateOfDevice(reinitializedStateOfDevice)
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) WithReinitializedStateOfDevice(reinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	b.ReinitializedStateOfDevice = reinitializedStateOfDevice
	return b
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) WithReinitializedStateOfDeviceBuilder(builderSupplier func(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	builder := builderSupplier(b.ReinitializedStateOfDevice.CreateBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder())
	var err error
	b.ReinitializedStateOfDevice, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) WithOptionalPassword(password BACnetContextTagCharacterString) BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	b.Password = password
	return b
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) WithOptionalPasswordBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	builder := builderSupplier(b.Password.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	b.Password, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) Build() (BACnetConfirmedServiceRequestReinitializeDevice, error) {
	if b.ReinitializedStateOfDevice == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'reinitializedStateOfDevice' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestReinitializeDevice.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) MustBuild() BACnetConfirmedServiceRequestReinitializeDevice {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestReinitializeDeviceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestReinitializeDeviceBuilder().(*_BACnetConfirmedServiceRequestReinitializeDeviceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestReinitializeDeviceBuilder creates a BACnetConfirmedServiceRequestReinitializeDeviceBuilder
func (b *_BACnetConfirmedServiceRequestReinitializeDevice) CreateBACnetConfirmedServiceRequestReinitializeDeviceBuilder() BACnetConfirmedServiceRequestReinitializeDeviceBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestReinitializeDeviceBuilder()
	}
	return &_BACnetConfirmedServiceRequestReinitializeDeviceBuilder{_BACnetConfirmedServiceRequestReinitializeDevice: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetReinitializedStateOfDevice() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	return m.ReinitializedStateOfDevice
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetPassword() BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReinitializeDevice(structType any) BACnetConfirmedServiceRequestReinitializeDevice {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDevice"
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (reinitializedStateOfDevice)
	lengthInBits += m.ReinitializedStateOfDevice.GetLengthInBits(ctx)

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += m.Password.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestReinitializeDevice BACnetConfirmedServiceRequestReinitializeDevice, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReinitializeDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reinitializedStateOfDevice, err := ReadSimpleField[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](ctx, "reinitializedStateOfDevice", ReadComplex[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reinitializedStateOfDevice' field"))
	}
	m.ReinitializedStateOfDevice = reinitializedStateOfDevice

	var password BACnetContextTagCharacterString
	_password, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "password", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'password' field"))
	}
	if _password != nil {
		password = *_password
		m.Password = password
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReinitializeDevice")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReinitializeDevice")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](ctx, "reinitializedStateOfDevice", m.GetReinitializedStateOfDevice(), WriteComplex[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reinitializedStateOfDevice' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "password", GetRef(m.GetPassword()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'password' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReinitializeDevice")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) IsBACnetConfirmedServiceRequestReinitializeDevice() {
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) deepCopy() *_BACnetConfirmedServiceRequestReinitializeDevice {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReinitializeDeviceCopy := &_BACnetConfirmedServiceRequestReinitializeDevice{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.ReinitializedStateOfDevice.DeepCopy().(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged),
		m.Password.DeepCopy().(BACnetContextTagCharacterString),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestReinitializeDeviceCopy
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) String() string {
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
