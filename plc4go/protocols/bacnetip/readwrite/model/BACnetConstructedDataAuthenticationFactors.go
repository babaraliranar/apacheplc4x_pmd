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

// BACnetConstructedDataAuthenticationFactors is the corresponding interface of BACnetConstructedDataAuthenticationFactors
type BACnetConstructedDataAuthenticationFactors interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationFactors returns AuthenticationFactors (property field)
	GetAuthenticationFactors() []BACnetCredentialAuthenticationFactor
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataAuthenticationFactors is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAuthenticationFactors()
	// CreateBuilder creates a BACnetConstructedDataAuthenticationFactorsBuilder
	CreateBACnetConstructedDataAuthenticationFactorsBuilder() BACnetConstructedDataAuthenticationFactorsBuilder
}

// _BACnetConstructedDataAuthenticationFactors is the data-structure of this message
type _BACnetConstructedDataAuthenticationFactors struct {
	BACnetConstructedDataContract
	NumberOfDataElements  BACnetApplicationTagUnsignedInteger
	AuthenticationFactors []BACnetCredentialAuthenticationFactor
}

var _ BACnetConstructedDataAuthenticationFactors = (*_BACnetConstructedDataAuthenticationFactors)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAuthenticationFactors)(nil)

// NewBACnetConstructedDataAuthenticationFactors factory function for _BACnetConstructedDataAuthenticationFactors
func NewBACnetConstructedDataAuthenticationFactors(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, authenticationFactors []BACnetCredentialAuthenticationFactor, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationFactors {
	_result := &_BACnetConstructedDataAuthenticationFactors{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AuthenticationFactors:         authenticationFactors,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAuthenticationFactorsBuilder is a builder for BACnetConstructedDataAuthenticationFactors
type BACnetConstructedDataAuthenticationFactorsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(authenticationFactors []BACnetCredentialAuthenticationFactor) BACnetConstructedDataAuthenticationFactorsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAuthenticationFactorsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAuthenticationFactorsBuilder
	// WithAuthenticationFactors adds AuthenticationFactors (property field)
	WithAuthenticationFactors(...BACnetCredentialAuthenticationFactor) BACnetConstructedDataAuthenticationFactorsBuilder
	// Build builds the BACnetConstructedDataAuthenticationFactors or returns an error if something is wrong
	Build() (BACnetConstructedDataAuthenticationFactors, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAuthenticationFactors
}

// NewBACnetConstructedDataAuthenticationFactorsBuilder() creates a BACnetConstructedDataAuthenticationFactorsBuilder
func NewBACnetConstructedDataAuthenticationFactorsBuilder() BACnetConstructedDataAuthenticationFactorsBuilder {
	return &_BACnetConstructedDataAuthenticationFactorsBuilder{_BACnetConstructedDataAuthenticationFactors: new(_BACnetConstructedDataAuthenticationFactors)}
}

type _BACnetConstructedDataAuthenticationFactorsBuilder struct {
	*_BACnetConstructedDataAuthenticationFactors

	err *utils.MultiError
}

var _ (BACnetConstructedDataAuthenticationFactorsBuilder) = (*_BACnetConstructedDataAuthenticationFactorsBuilder)(nil)

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) WithMandatoryFields(authenticationFactors []BACnetCredentialAuthenticationFactor) BACnetConstructedDataAuthenticationFactorsBuilder {
	return m.WithAuthenticationFactors(authenticationFactors...)
}

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAuthenticationFactorsBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAuthenticationFactorsBuilder {
	builder := builderSupplier(m.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) WithAuthenticationFactors(authenticationFactors ...BACnetCredentialAuthenticationFactor) BACnetConstructedDataAuthenticationFactorsBuilder {
	m.AuthenticationFactors = authenticationFactors
	return m
}

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) Build() (BACnetConstructedDataAuthenticationFactors, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataAuthenticationFactors.deepCopy(), nil
}

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) MustBuild() BACnetConstructedDataAuthenticationFactors {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataAuthenticationFactorsBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataAuthenticationFactorsBuilder()
}

// CreateBACnetConstructedDataAuthenticationFactorsBuilder creates a BACnetConstructedDataAuthenticationFactorsBuilder
func (m *_BACnetConstructedDataAuthenticationFactors) CreateBACnetConstructedDataAuthenticationFactorsBuilder() BACnetConstructedDataAuthenticationFactorsBuilder {
	if m == nil {
		return NewBACnetConstructedDataAuthenticationFactorsBuilder()
	}
	return &_BACnetConstructedDataAuthenticationFactorsBuilder{_BACnetConstructedDataAuthenticationFactors: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_FACTORS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetAuthenticationFactors() []BACnetCredentialAuthenticationFactor {
	return m.AuthenticationFactors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetZero() uint64 {
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
func CastBACnetConstructedDataAuthenticationFactors(structType any) BACnetConstructedDataAuthenticationFactors {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationFactors); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationFactors); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationFactors"
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AuthenticationFactors) > 0 {
		for _, element := range m.AuthenticationFactors {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAuthenticationFactors) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAuthenticationFactors BACnetConstructedDataAuthenticationFactors, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationFactors"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationFactors")
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

	authenticationFactors, err := ReadTerminatedArrayField[BACnetCredentialAuthenticationFactor](ctx, "authenticationFactors", ReadComplex[BACnetCredentialAuthenticationFactor](BACnetCredentialAuthenticationFactorParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationFactors' field"))
	}
	m.AuthenticationFactors = authenticationFactors

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationFactors"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationFactors")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAuthenticationFactors) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthenticationFactors) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationFactors"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationFactors")
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

		if err := WriteComplexTypeArrayField(ctx, "authenticationFactors", m.GetAuthenticationFactors(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationFactors' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationFactors"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationFactors")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationFactors) IsBACnetConstructedDataAuthenticationFactors() {
}

func (m *_BACnetConstructedDataAuthenticationFactors) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAuthenticationFactors) deepCopy() *_BACnetConstructedDataAuthenticationFactors {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAuthenticationFactorsCopy := &_BACnetConstructedDataAuthenticationFactors{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetCredentialAuthenticationFactor, BACnetCredentialAuthenticationFactor](m.AuthenticationFactors),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAuthenticationFactorsCopy
}

func (m *_BACnetConstructedDataAuthenticationFactors) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
