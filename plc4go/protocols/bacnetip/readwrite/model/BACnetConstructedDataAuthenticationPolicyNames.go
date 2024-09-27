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

// BACnetConstructedDataAuthenticationPolicyNames is the corresponding interface of BACnetConstructedDataAuthenticationPolicyNames
type BACnetConstructedDataAuthenticationPolicyNames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationPolicyNames returns AuthenticationPolicyNames (property field)
	GetAuthenticationPolicyNames() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataAuthenticationPolicyNames is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAuthenticationPolicyNames()
	// CreateBuilder creates a BACnetConstructedDataAuthenticationPolicyNamesBuilder
	CreateBACnetConstructedDataAuthenticationPolicyNamesBuilder() BACnetConstructedDataAuthenticationPolicyNamesBuilder
}

// _BACnetConstructedDataAuthenticationPolicyNames is the data-structure of this message
type _BACnetConstructedDataAuthenticationPolicyNames struct {
	BACnetConstructedDataContract
	NumberOfDataElements      BACnetApplicationTagUnsignedInteger
	AuthenticationPolicyNames []BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataAuthenticationPolicyNames = (*_BACnetConstructedDataAuthenticationPolicyNames)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAuthenticationPolicyNames)(nil)

// NewBACnetConstructedDataAuthenticationPolicyNames factory function for _BACnetConstructedDataAuthenticationPolicyNames
func NewBACnetConstructedDataAuthenticationPolicyNames(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, authenticationPolicyNames []BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationPolicyNames {
	_result := &_BACnetConstructedDataAuthenticationPolicyNames{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AuthenticationPolicyNames:     authenticationPolicyNames,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAuthenticationPolicyNamesBuilder is a builder for BACnetConstructedDataAuthenticationPolicyNames
type BACnetConstructedDataAuthenticationPolicyNamesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(authenticationPolicyNames []BACnetApplicationTagCharacterString) BACnetConstructedDataAuthenticationPolicyNamesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAuthenticationPolicyNamesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAuthenticationPolicyNamesBuilder
	// WithAuthenticationPolicyNames adds AuthenticationPolicyNames (property field)
	WithAuthenticationPolicyNames(...BACnetApplicationTagCharacterString) BACnetConstructedDataAuthenticationPolicyNamesBuilder
	// Build builds the BACnetConstructedDataAuthenticationPolicyNames or returns an error if something is wrong
	Build() (BACnetConstructedDataAuthenticationPolicyNames, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAuthenticationPolicyNames
}

// NewBACnetConstructedDataAuthenticationPolicyNamesBuilder() creates a BACnetConstructedDataAuthenticationPolicyNamesBuilder
func NewBACnetConstructedDataAuthenticationPolicyNamesBuilder() BACnetConstructedDataAuthenticationPolicyNamesBuilder {
	return &_BACnetConstructedDataAuthenticationPolicyNamesBuilder{_BACnetConstructedDataAuthenticationPolicyNames: new(_BACnetConstructedDataAuthenticationPolicyNames)}
}

type _BACnetConstructedDataAuthenticationPolicyNamesBuilder struct {
	*_BACnetConstructedDataAuthenticationPolicyNames

	err *utils.MultiError
}

var _ (BACnetConstructedDataAuthenticationPolicyNamesBuilder) = (*_BACnetConstructedDataAuthenticationPolicyNamesBuilder)(nil)

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) WithMandatoryFields(authenticationPolicyNames []BACnetApplicationTagCharacterString) BACnetConstructedDataAuthenticationPolicyNamesBuilder {
	return m.WithAuthenticationPolicyNames(authenticationPolicyNames...)
}

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAuthenticationPolicyNamesBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAuthenticationPolicyNamesBuilder {
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

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) WithAuthenticationPolicyNames(authenticationPolicyNames ...BACnetApplicationTagCharacterString) BACnetConstructedDataAuthenticationPolicyNamesBuilder {
	m.AuthenticationPolicyNames = authenticationPolicyNames
	return m
}

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) Build() (BACnetConstructedDataAuthenticationPolicyNames, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataAuthenticationPolicyNames.deepCopy(), nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) MustBuild() BACnetConstructedDataAuthenticationPolicyNames {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataAuthenticationPolicyNamesBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataAuthenticationPolicyNamesBuilder()
}

// CreateBACnetConstructedDataAuthenticationPolicyNamesBuilder creates a BACnetConstructedDataAuthenticationPolicyNamesBuilder
func (m *_BACnetConstructedDataAuthenticationPolicyNames) CreateBACnetConstructedDataAuthenticationPolicyNamesBuilder() BACnetConstructedDataAuthenticationPolicyNamesBuilder {
	if m == nil {
		return NewBACnetConstructedDataAuthenticationPolicyNamesBuilder()
	}
	return &_BACnetConstructedDataAuthenticationPolicyNamesBuilder{_BACnetConstructedDataAuthenticationPolicyNames: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_POLICY_NAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetAuthenticationPolicyNames() []BACnetApplicationTagCharacterString {
	return m.AuthenticationPolicyNames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetZero() uint64 {
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
func CastBACnetConstructedDataAuthenticationPolicyNames(structType any) BACnetConstructedDataAuthenticationPolicyNames {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationPolicyNames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationPolicyNames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationPolicyNames"
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AuthenticationPolicyNames) > 0 {
		for _, element := range m.AuthenticationPolicyNames {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAuthenticationPolicyNames BACnetConstructedDataAuthenticationPolicyNames, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationPolicyNames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationPolicyNames")
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

	authenticationPolicyNames, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "authenticationPolicyNames", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationPolicyNames' field"))
	}
	m.AuthenticationPolicyNames = authenticationPolicyNames

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationPolicyNames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationPolicyNames")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationPolicyNames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationPolicyNames")
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

		if err := WriteComplexTypeArrayField(ctx, "authenticationPolicyNames", m.GetAuthenticationPolicyNames(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationPolicyNames' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationPolicyNames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationPolicyNames")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) IsBACnetConstructedDataAuthenticationPolicyNames() {
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) deepCopy() *_BACnetConstructedDataAuthenticationPolicyNames {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAuthenticationPolicyNamesCopy := &_BACnetConstructedDataAuthenticationPolicyNames{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagCharacterString, BACnetApplicationTagCharacterString](m.AuthenticationPolicyNames),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAuthenticationPolicyNamesCopy
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
