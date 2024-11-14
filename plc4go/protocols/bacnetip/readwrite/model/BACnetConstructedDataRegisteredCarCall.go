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

// BACnetConstructedDataRegisteredCarCall is the corresponding interface of BACnetConstructedDataRegisteredCarCall
type BACnetConstructedDataRegisteredCarCall interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetRegisteredCarCall returns RegisteredCarCall (property field)
	GetRegisteredCarCall() []BACnetLiftCarCallList
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataRegisteredCarCall is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRegisteredCarCall()
	// CreateBuilder creates a BACnetConstructedDataRegisteredCarCallBuilder
	CreateBACnetConstructedDataRegisteredCarCallBuilder() BACnetConstructedDataRegisteredCarCallBuilder
}

// _BACnetConstructedDataRegisteredCarCall is the data-structure of this message
type _BACnetConstructedDataRegisteredCarCall struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	RegisteredCarCall    []BACnetLiftCarCallList
}

var _ BACnetConstructedDataRegisteredCarCall = (*_BACnetConstructedDataRegisteredCarCall)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRegisteredCarCall)(nil)

// NewBACnetConstructedDataRegisteredCarCall factory function for _BACnetConstructedDataRegisteredCarCall
func NewBACnetConstructedDataRegisteredCarCall(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, registeredCarCall []BACnetLiftCarCallList, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRegisteredCarCall {
	_result := &_BACnetConstructedDataRegisteredCarCall{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		RegisteredCarCall:             registeredCarCall,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRegisteredCarCallBuilder is a builder for BACnetConstructedDataRegisteredCarCall
type BACnetConstructedDataRegisteredCarCallBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(registeredCarCall []BACnetLiftCarCallList) BACnetConstructedDataRegisteredCarCallBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRegisteredCarCallBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRegisteredCarCallBuilder
	// WithRegisteredCarCall adds RegisteredCarCall (property field)
	WithRegisteredCarCall(...BACnetLiftCarCallList) BACnetConstructedDataRegisteredCarCallBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataRegisteredCarCall or returns an error if something is wrong
	Build() (BACnetConstructedDataRegisteredCarCall, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRegisteredCarCall
}

// NewBACnetConstructedDataRegisteredCarCallBuilder() creates a BACnetConstructedDataRegisteredCarCallBuilder
func NewBACnetConstructedDataRegisteredCarCallBuilder() BACnetConstructedDataRegisteredCarCallBuilder {
	return &_BACnetConstructedDataRegisteredCarCallBuilder{_BACnetConstructedDataRegisteredCarCall: new(_BACnetConstructedDataRegisteredCarCall)}
}

type _BACnetConstructedDataRegisteredCarCallBuilder struct {
	*_BACnetConstructedDataRegisteredCarCall

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataRegisteredCarCallBuilder) = (*_BACnetConstructedDataRegisteredCarCallBuilder)(nil)

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataRegisteredCarCall
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) WithMandatoryFields(registeredCarCall []BACnetLiftCarCallList) BACnetConstructedDataRegisteredCarCallBuilder {
	return b.WithRegisteredCarCall(registeredCarCall...)
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRegisteredCarCallBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRegisteredCarCallBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) WithRegisteredCarCall(registeredCarCall ...BACnetLiftCarCallList) BACnetConstructedDataRegisteredCarCallBuilder {
	b.RegisteredCarCall = registeredCarCall
	return b
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) Build() (BACnetConstructedDataRegisteredCarCall, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataRegisteredCarCall.deepCopy(), nil
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) MustBuild() BACnetConstructedDataRegisteredCarCall {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataRegisteredCarCallBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataRegisteredCarCallBuilder().(*_BACnetConstructedDataRegisteredCarCallBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataRegisteredCarCallBuilder creates a BACnetConstructedDataRegisteredCarCallBuilder
func (b *_BACnetConstructedDataRegisteredCarCall) CreateBACnetConstructedDataRegisteredCarCallBuilder() BACnetConstructedDataRegisteredCarCallBuilder {
	if b == nil {
		return NewBACnetConstructedDataRegisteredCarCallBuilder()
	}
	return &_BACnetConstructedDataRegisteredCarCallBuilder{_BACnetConstructedDataRegisteredCarCall: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REGISTERED_CAR_CALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetRegisteredCarCall() []BACnetLiftCarCallList {
	return m.RegisteredCarCall
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRegisteredCarCall(structType any) BACnetConstructedDataRegisteredCarCall {
	if casted, ok := structType.(BACnetConstructedDataRegisteredCarCall); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRegisteredCarCall); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetTypeName() string {
	return "BACnetConstructedDataRegisteredCarCall"
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.RegisteredCarCall) > 0 {
		for _, element := range m.RegisteredCarCall {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRegisteredCarCall) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRegisteredCarCall BACnetConstructedDataRegisteredCarCall, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRegisteredCarCall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRegisteredCarCall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	registeredCarCall, err := ReadTerminatedArrayField[BACnetLiftCarCallList](ctx, "registeredCarCall", ReadComplex[BACnetLiftCarCallList](BACnetLiftCarCallListParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'registeredCarCall' field"))
	}
	m.RegisteredCarCall = registeredCarCall

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRegisteredCarCall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRegisteredCarCall")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRegisteredCarCall) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRegisteredCarCall) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRegisteredCarCall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRegisteredCarCall")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "registeredCarCall", m.GetRegisteredCarCall(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'registeredCarCall' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRegisteredCarCall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRegisteredCarCall")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRegisteredCarCall) IsBACnetConstructedDataRegisteredCarCall() {}

func (m *_BACnetConstructedDataRegisteredCarCall) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRegisteredCarCall) deepCopy() *_BACnetConstructedDataRegisteredCarCall {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRegisteredCarCallCopy := &_BACnetConstructedDataRegisteredCarCall{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetLiftCarCallList, BACnetLiftCarCallList](m.RegisteredCarCall),
	}
	_BACnetConstructedDataRegisteredCarCallCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRegisteredCarCallCopy
}

func (m *_BACnetConstructedDataRegisteredCarCall) String() string {
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
