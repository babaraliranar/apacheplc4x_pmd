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

// BACnetUnconfirmedServiceRequestIHave is the corresponding interface of BACnetUnconfirmedServiceRequestIHave
type BACnetUnconfirmedServiceRequestIHave interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetObjectName returns ObjectName (property field)
	GetObjectName() BACnetApplicationTagCharacterString
	// IsBACnetUnconfirmedServiceRequestIHave is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestIHave()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestIHaveBuilder
	CreateBACnetUnconfirmedServiceRequestIHaveBuilder() BACnetUnconfirmedServiceRequestIHaveBuilder
}

// _BACnetUnconfirmedServiceRequestIHave is the data-structure of this message
type _BACnetUnconfirmedServiceRequestIHave struct {
	BACnetUnconfirmedServiceRequestContract
	DeviceIdentifier BACnetApplicationTagObjectIdentifier
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
	ObjectName       BACnetApplicationTagCharacterString
}

var _ BACnetUnconfirmedServiceRequestIHave = (*_BACnetUnconfirmedServiceRequestIHave)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestIHave)(nil)

// NewBACnetUnconfirmedServiceRequestIHave factory function for _BACnetUnconfirmedServiceRequestIHave
func NewBACnetUnconfirmedServiceRequestIHave(deviceIdentifier BACnetApplicationTagObjectIdentifier, objectIdentifier BACnetApplicationTagObjectIdentifier, objectName BACnetApplicationTagCharacterString, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestIHave {
	if deviceIdentifier == nil {
		panic("deviceIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetUnconfirmedServiceRequestIHave must not be nil")
	}
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetUnconfirmedServiceRequestIHave must not be nil")
	}
	if objectName == nil {
		panic("objectName of type BACnetApplicationTagCharacterString for BACnetUnconfirmedServiceRequestIHave must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestIHave{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		DeviceIdentifier:                        deviceIdentifier,
		ObjectIdentifier:                        objectIdentifier,
		ObjectName:                              objectName,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestIHaveBuilder is a builder for BACnetUnconfirmedServiceRequestIHave
type BACnetUnconfirmedServiceRequestIHaveBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deviceIdentifier BACnetApplicationTagObjectIdentifier, objectIdentifier BACnetApplicationTagObjectIdentifier, objectName BACnetApplicationTagCharacterString) BACnetUnconfirmedServiceRequestIHaveBuilder
	// WithDeviceIdentifier adds DeviceIdentifier (property field)
	WithDeviceIdentifier(BACnetApplicationTagObjectIdentifier) BACnetUnconfirmedServiceRequestIHaveBuilder
	// WithDeviceIdentifierBuilder adds DeviceIdentifier (property field) which is build by the builder
	WithDeviceIdentifierBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestIHaveBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetApplicationTagObjectIdentifier) BACnetUnconfirmedServiceRequestIHaveBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestIHaveBuilder
	// WithObjectName adds ObjectName (property field)
	WithObjectName(BACnetApplicationTagCharacterString) BACnetUnconfirmedServiceRequestIHaveBuilder
	// WithObjectNameBuilder adds ObjectName (property field) which is build by the builder
	WithObjectNameBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetUnconfirmedServiceRequestIHaveBuilder
	// Build builds the BACnetUnconfirmedServiceRequestIHave or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestIHave, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestIHave
}

// NewBACnetUnconfirmedServiceRequestIHaveBuilder() creates a BACnetUnconfirmedServiceRequestIHaveBuilder
func NewBACnetUnconfirmedServiceRequestIHaveBuilder() BACnetUnconfirmedServiceRequestIHaveBuilder {
	return &_BACnetUnconfirmedServiceRequestIHaveBuilder{_BACnetUnconfirmedServiceRequestIHave: new(_BACnetUnconfirmedServiceRequestIHave)}
}

type _BACnetUnconfirmedServiceRequestIHaveBuilder struct {
	*_BACnetUnconfirmedServiceRequestIHave

	parentBuilder *_BACnetUnconfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestIHaveBuilder) = (*_BACnetUnconfirmedServiceRequestIHaveBuilder)(nil)

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) setParent(contract BACnetUnconfirmedServiceRequestContract) {
	b.BACnetUnconfirmedServiceRequestContract = contract
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithMandatoryFields(deviceIdentifier BACnetApplicationTagObjectIdentifier, objectIdentifier BACnetApplicationTagObjectIdentifier, objectName BACnetApplicationTagCharacterString) BACnetUnconfirmedServiceRequestIHaveBuilder {
	return b.WithDeviceIdentifier(deviceIdentifier).WithObjectIdentifier(objectIdentifier).WithObjectName(objectName)
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithDeviceIdentifier(deviceIdentifier BACnetApplicationTagObjectIdentifier) BACnetUnconfirmedServiceRequestIHaveBuilder {
	b.DeviceIdentifier = deviceIdentifier
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithDeviceIdentifierBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestIHaveBuilder {
	builder := builderSupplier(b.DeviceIdentifier.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.DeviceIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithObjectIdentifier(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetUnconfirmedServiceRequestIHaveBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestIHaveBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithObjectName(objectName BACnetApplicationTagCharacterString) BACnetUnconfirmedServiceRequestIHaveBuilder {
	b.ObjectName = objectName
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) WithObjectNameBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetUnconfirmedServiceRequestIHaveBuilder {
	builder := builderSupplier(b.ObjectName.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.ObjectName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) Build() (BACnetUnconfirmedServiceRequestIHave, error) {
	if b.DeviceIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deviceIdentifier' not set"))
	}
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.ObjectName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetUnconfirmedServiceRequestIHave.deepCopy(), nil
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) MustBuild() BACnetUnconfirmedServiceRequestIHave {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) Done() BACnetUnconfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) buildForBACnetUnconfirmedServiceRequest() (BACnetUnconfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetUnconfirmedServiceRequestIHaveBuilder) DeepCopy() any {
	_copy := b.CreateBACnetUnconfirmedServiceRequestIHaveBuilder().(*_BACnetUnconfirmedServiceRequestIHaveBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetUnconfirmedServiceRequestIHaveBuilder creates a BACnetUnconfirmedServiceRequestIHaveBuilder
func (b *_BACnetUnconfirmedServiceRequestIHave) CreateBACnetUnconfirmedServiceRequestIHaveBuilder() BACnetUnconfirmedServiceRequestIHaveBuilder {
	if b == nil {
		return NewBACnetUnconfirmedServiceRequestIHaveBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestIHaveBuilder{_BACnetUnconfirmedServiceRequestIHave: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIHave) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_I_HAVE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestIHave) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIHave) GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetObjectName() BACnetApplicationTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestIHave(structType any) BACnetUnconfirmedServiceRequestIHave {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestIHave); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestIHave); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIHave"
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestIHave) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestIHave BACnetUnconfirmedServiceRequestIHave, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIHave"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestIHave")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}
	m.DeviceIdentifier = deviceIdentifier

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	objectName, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "objectName", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectName' field"))
	}
	m.ObjectName = objectName

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIHave"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestIHave")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestIHave) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestIHave) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIHave"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestIHave")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", m.GetDeviceIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "objectName", m.GetObjectName(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIHave"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestIHave")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestIHave) IsBACnetUnconfirmedServiceRequestIHave() {}

func (m *_BACnetUnconfirmedServiceRequestIHave) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestIHave) deepCopy() *_BACnetUnconfirmedServiceRequestIHave {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestIHaveCopy := &_BACnetUnconfirmedServiceRequestIHave{
		m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).deepCopy(),
		m.DeviceIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
		m.ObjectIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
		m.ObjectName.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestIHaveCopy
}

func (m *_BACnetUnconfirmedServiceRequestIHave) String() string {
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
