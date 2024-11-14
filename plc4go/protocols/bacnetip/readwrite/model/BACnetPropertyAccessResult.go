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

// BACnetPropertyAccessResult is the corresponding interface of BACnetPropertyAccessResult
type BACnetPropertyAccessResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetAccessResult returns AccessResult (property field)
	GetAccessResult() BACnetPropertyAccessResultAccessResult
	// IsBACnetPropertyAccessResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyAccessResult()
	// CreateBuilder creates a BACnetPropertyAccessResultBuilder
	CreateBACnetPropertyAccessResultBuilder() BACnetPropertyAccessResultBuilder
}

// _BACnetPropertyAccessResult is the data-structure of this message
type _BACnetPropertyAccessResult struct {
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	PropertyArrayIndex BACnetContextTagUnsignedInteger
	DeviceIdentifier   BACnetContextTagObjectIdentifier
	AccessResult       BACnetPropertyAccessResultAccessResult
}

var _ BACnetPropertyAccessResult = (*_BACnetPropertyAccessResult)(nil)

// NewBACnetPropertyAccessResult factory function for _BACnetPropertyAccessResult
func NewBACnetPropertyAccessResult(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, deviceIdentifier BACnetContextTagObjectIdentifier, accessResult BACnetPropertyAccessResultAccessResult) *_BACnetPropertyAccessResult {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetPropertyAccessResult must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetPropertyAccessResult must not be nil")
	}
	if accessResult == nil {
		panic("accessResult of type BACnetPropertyAccessResultAccessResult for BACnetPropertyAccessResult must not be nil")
	}
	return &_BACnetPropertyAccessResult{ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, PropertyArrayIndex: propertyArrayIndex, DeviceIdentifier: deviceIdentifier, AccessResult: accessResult}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyAccessResultBuilder is a builder for BACnetPropertyAccessResult
type BACnetPropertyAccessResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, accessResult BACnetPropertyAccessResultAccessResult) BACnetPropertyAccessResultBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetPropertyAccessResultBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetPropertyAccessResultBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetPropertyAccessResultBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyAccessResultBuilder
	// WithPropertyArrayIndex adds PropertyArrayIndex (property field)
	WithOptionalPropertyArrayIndex(BACnetContextTagUnsignedInteger) BACnetPropertyAccessResultBuilder
	// WithOptionalPropertyArrayIndexBuilder adds PropertyArrayIndex (property field) which is build by the builder
	WithOptionalPropertyArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyAccessResultBuilder
	// WithDeviceIdentifier adds DeviceIdentifier (property field)
	WithOptionalDeviceIdentifier(BACnetContextTagObjectIdentifier) BACnetPropertyAccessResultBuilder
	// WithOptionalDeviceIdentifierBuilder adds DeviceIdentifier (property field) which is build by the builder
	WithOptionalDeviceIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetPropertyAccessResultBuilder
	// WithAccessResult adds AccessResult (property field)
	WithAccessResult(BACnetPropertyAccessResultAccessResult) BACnetPropertyAccessResultBuilder
	// WithAccessResultBuilder adds AccessResult (property field) which is build by the builder
	WithAccessResultBuilder(func(BACnetPropertyAccessResultAccessResultBuilder) BACnetPropertyAccessResultAccessResultBuilder) BACnetPropertyAccessResultBuilder
	// Build builds the BACnetPropertyAccessResult or returns an error if something is wrong
	Build() (BACnetPropertyAccessResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyAccessResult
}

// NewBACnetPropertyAccessResultBuilder() creates a BACnetPropertyAccessResultBuilder
func NewBACnetPropertyAccessResultBuilder() BACnetPropertyAccessResultBuilder {
	return &_BACnetPropertyAccessResultBuilder{_BACnetPropertyAccessResult: new(_BACnetPropertyAccessResult)}
}

type _BACnetPropertyAccessResultBuilder struct {
	*_BACnetPropertyAccessResult

	err *utils.MultiError
}

var _ (BACnetPropertyAccessResultBuilder) = (*_BACnetPropertyAccessResultBuilder)(nil)

func (b *_BACnetPropertyAccessResultBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, accessResult BACnetPropertyAccessResultAccessResult) BACnetPropertyAccessResultBuilder {
	return b.WithObjectIdentifier(objectIdentifier).WithPropertyIdentifier(propertyIdentifier).WithAccessResult(accessResult)
}

func (b *_BACnetPropertyAccessResultBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetPropertyAccessResultBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetPropertyAccessResultBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetPropertyAccessResultBuilder {
	b.PropertyIdentifier = propertyIdentifier
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyAccessResultBuilder {
	builder := builderSupplier(b.PropertyIdentifier.CreateBACnetPropertyIdentifierTaggedBuilder())
	var err error
	b.PropertyIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyIdentifierTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithOptionalPropertyArrayIndex(propertyArrayIndex BACnetContextTagUnsignedInteger) BACnetPropertyAccessResultBuilder {
	b.PropertyArrayIndex = propertyArrayIndex
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithOptionalPropertyArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyAccessResultBuilder {
	builder := builderSupplier(b.PropertyArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.PropertyArrayIndex, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithOptionalDeviceIdentifier(deviceIdentifier BACnetContextTagObjectIdentifier) BACnetPropertyAccessResultBuilder {
	b.DeviceIdentifier = deviceIdentifier
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithOptionalDeviceIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetPropertyAccessResultBuilder {
	builder := builderSupplier(b.DeviceIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.DeviceIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithAccessResult(accessResult BACnetPropertyAccessResultAccessResult) BACnetPropertyAccessResultBuilder {
	b.AccessResult = accessResult
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) WithAccessResultBuilder(builderSupplier func(BACnetPropertyAccessResultAccessResultBuilder) BACnetPropertyAccessResultAccessResultBuilder) BACnetPropertyAccessResultBuilder {
	builder := builderSupplier(b.AccessResult.CreateBACnetPropertyAccessResultAccessResultBuilder())
	var err error
	b.AccessResult, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyAccessResultAccessResultBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyAccessResultBuilder) Build() (BACnetPropertyAccessResult, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.PropertyIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if b.AccessResult == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'accessResult' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyAccessResult.deepCopy(), nil
}

func (b *_BACnetPropertyAccessResultBuilder) MustBuild() BACnetPropertyAccessResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyAccessResultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyAccessResultBuilder().(*_BACnetPropertyAccessResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyAccessResultBuilder creates a BACnetPropertyAccessResultBuilder
func (b *_BACnetPropertyAccessResult) CreateBACnetPropertyAccessResultBuilder() BACnetPropertyAccessResultBuilder {
	if b == nil {
		return NewBACnetPropertyAccessResultBuilder()
	}
	return &_BACnetPropertyAccessResultBuilder{_BACnetPropertyAccessResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResult) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetPropertyAccessResult) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetPropertyAccessResult) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetPropertyAccessResult) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetPropertyAccessResult) GetAccessResult() BACnetPropertyAccessResultAccessResult {
	return m.AccessResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResult(structType any) BACnetPropertyAccessResult {
	if casted, ok := structType.(BACnetPropertyAccessResult); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResult); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResult) GetTypeName() string {
	return "BACnetPropertyAccessResult"
}

func (m *_BACnetPropertyAccessResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)
	}

	// Simple field (accessResult)
	lengthInBits += m.AccessResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyAccessResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyAccessResultParse(ctx context.Context, theBytes []byte) (BACnetPropertyAccessResult, error) {
	return BACnetPropertyAccessResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPropertyAccessResultParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyAccessResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyAccessResult, error) {
		return BACnetPropertyAccessResultParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetPropertyAccessResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyAccessResult, error) {
	v, err := (&_BACnetPropertyAccessResult{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPropertyAccessResult) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetPropertyAccessResult BACnetPropertyAccessResult, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var propertyArrayIndex BACnetContextTagUnsignedInteger
	_propertyArrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyArrayIndex' field"))
	}
	if _propertyArrayIndex != nil {
		propertyArrayIndex = *_propertyArrayIndex
		m.PropertyArrayIndex = propertyArrayIndex
	}

	var deviceIdentifier BACnetContextTagObjectIdentifier
	_deviceIdentifier, err := ReadOptionalField[BACnetContextTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}
	if _deviceIdentifier != nil {
		deviceIdentifier = *_deviceIdentifier
		m.DeviceIdentifier = deviceIdentifier
	}

	accessResult, err := ReadSimpleField[BACnetPropertyAccessResultAccessResult](ctx, "accessResult", ReadComplex[BACnetPropertyAccessResultAccessResult](BACnetPropertyAccessResultAccessResultParseWithBufferProducer[BACnetPropertyAccessResultAccessResult]((BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((propertyArrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessResult' field"))
	}
	m.AccessResult = accessResult

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResult")
	}

	return m, nil
}

func (m *_BACnetPropertyAccessResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyAccessResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResult")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", GetRef(m.GetPropertyArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyArrayIndex' field")
	}

	if err := WriteOptionalField[BACnetContextTagObjectIdentifier](ctx, "deviceIdentifier", GetRef(m.GetDeviceIdentifier()), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'deviceIdentifier' field")
	}

	if err := WriteSimpleField[BACnetPropertyAccessResultAccessResult](ctx, "accessResult", m.GetAccessResult(), WriteComplex[BACnetPropertyAccessResultAccessResult](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'accessResult' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyAccessResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResult")
	}
	return nil
}

func (m *_BACnetPropertyAccessResult) IsBACnetPropertyAccessResult() {}

func (m *_BACnetPropertyAccessResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyAccessResult) deepCopy() *_BACnetPropertyAccessResult {
	if m == nil {
		return nil
	}
	_BACnetPropertyAccessResultCopy := &_BACnetPropertyAccessResult{
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.ObjectIdentifier),
		utils.DeepCopy[BACnetPropertyIdentifierTagged](m.PropertyIdentifier),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.PropertyArrayIndex),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.DeviceIdentifier),
		utils.DeepCopy[BACnetPropertyAccessResultAccessResult](m.AccessResult),
	}
	return _BACnetPropertyAccessResultCopy
}

func (m *_BACnetPropertyAccessResult) String() string {
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
