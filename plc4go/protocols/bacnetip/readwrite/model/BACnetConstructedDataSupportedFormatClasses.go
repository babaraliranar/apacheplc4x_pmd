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

// BACnetConstructedDataSupportedFormatClasses is the corresponding interface of BACnetConstructedDataSupportedFormatClasses
type BACnetConstructedDataSupportedFormatClasses interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSupportedFormats returns SupportedFormats (property field)
	GetSupportedFormats() []BACnetApplicationTagUnsignedInteger
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataSupportedFormatClasses is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSupportedFormatClasses()
	// CreateBuilder creates a BACnetConstructedDataSupportedFormatClassesBuilder
	CreateBACnetConstructedDataSupportedFormatClassesBuilder() BACnetConstructedDataSupportedFormatClassesBuilder
}

// _BACnetConstructedDataSupportedFormatClasses is the data-structure of this message
type _BACnetConstructedDataSupportedFormatClasses struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	SupportedFormats     []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataSupportedFormatClasses = (*_BACnetConstructedDataSupportedFormatClasses)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSupportedFormatClasses)(nil)

// NewBACnetConstructedDataSupportedFormatClasses factory function for _BACnetConstructedDataSupportedFormatClasses
func NewBACnetConstructedDataSupportedFormatClasses(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, supportedFormats []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSupportedFormatClasses {
	_result := &_BACnetConstructedDataSupportedFormatClasses{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		SupportedFormats:              supportedFormats,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSupportedFormatClassesBuilder is a builder for BACnetConstructedDataSupportedFormatClasses
type BACnetConstructedDataSupportedFormatClassesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(supportedFormats []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedFormatClassesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedFormatClassesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSupportedFormatClassesBuilder
	// WithSupportedFormats adds SupportedFormats (property field)
	WithSupportedFormats(...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedFormatClassesBuilder
	// Build builds the BACnetConstructedDataSupportedFormatClasses or returns an error if something is wrong
	Build() (BACnetConstructedDataSupportedFormatClasses, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSupportedFormatClasses
}

// NewBACnetConstructedDataSupportedFormatClassesBuilder() creates a BACnetConstructedDataSupportedFormatClassesBuilder
func NewBACnetConstructedDataSupportedFormatClassesBuilder() BACnetConstructedDataSupportedFormatClassesBuilder {
	return &_BACnetConstructedDataSupportedFormatClassesBuilder{_BACnetConstructedDataSupportedFormatClasses: new(_BACnetConstructedDataSupportedFormatClasses)}
}

type _BACnetConstructedDataSupportedFormatClassesBuilder struct {
	*_BACnetConstructedDataSupportedFormatClasses

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSupportedFormatClassesBuilder) = (*_BACnetConstructedDataSupportedFormatClassesBuilder)(nil)

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) WithMandatoryFields(supportedFormats []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedFormatClassesBuilder {
	return b.WithSupportedFormats(supportedFormats...)
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedFormatClassesBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSupportedFormatClassesBuilder {
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

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) WithSupportedFormats(supportedFormats ...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedFormatClassesBuilder {
	b.SupportedFormats = supportedFormats
	return b
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) Build() (BACnetConstructedDataSupportedFormatClasses, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSupportedFormatClasses.deepCopy(), nil
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) MustBuild() BACnetConstructedDataSupportedFormatClasses {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSupportedFormatClassesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSupportedFormatClassesBuilder().(*_BACnetConstructedDataSupportedFormatClassesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSupportedFormatClassesBuilder creates a BACnetConstructedDataSupportedFormatClassesBuilder
func (b *_BACnetConstructedDataSupportedFormatClasses) CreateBACnetConstructedDataSupportedFormatClassesBuilder() BACnetConstructedDataSupportedFormatClassesBuilder {
	if b == nil {
		return NewBACnetConstructedDataSupportedFormatClassesBuilder()
	}
	return &_BACnetConstructedDataSupportedFormatClassesBuilder{_BACnetConstructedDataSupportedFormatClasses: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSupportedFormatClasses) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSupportedFormatClasses) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUPPORTED_FORMAT_CLASSES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSupportedFormatClasses) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSupportedFormatClasses) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSupportedFormatClasses) GetSupportedFormats() []BACnetApplicationTagUnsignedInteger {
	return m.SupportedFormats
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSupportedFormatClasses) GetZero() uint64 {
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
func CastBACnetConstructedDataSupportedFormatClasses(structType any) BACnetConstructedDataSupportedFormatClasses {
	if casted, ok := structType.(BACnetConstructedDataSupportedFormatClasses); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSupportedFormatClasses); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSupportedFormatClasses) GetTypeName() string {
	return "BACnetConstructedDataSupportedFormatClasses"
}

func (m *_BACnetConstructedDataSupportedFormatClasses) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.SupportedFormats) > 0 {
		for _, element := range m.SupportedFormats {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSupportedFormatClasses) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSupportedFormatClasses) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSupportedFormatClasses BACnetConstructedDataSupportedFormatClasses, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSupportedFormatClasses"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSupportedFormatClasses")
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

	supportedFormats, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "supportedFormats", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'supportedFormats' field"))
	}
	m.SupportedFormats = supportedFormats

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSupportedFormatClasses"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSupportedFormatClasses")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSupportedFormatClasses) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSupportedFormatClasses) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSupportedFormatClasses"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSupportedFormatClasses")
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

		if err := WriteComplexTypeArrayField(ctx, "supportedFormats", m.GetSupportedFormats(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'supportedFormats' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSupportedFormatClasses"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSupportedFormatClasses")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSupportedFormatClasses) IsBACnetConstructedDataSupportedFormatClasses() {
}

func (m *_BACnetConstructedDataSupportedFormatClasses) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSupportedFormatClasses) deepCopy() *_BACnetConstructedDataSupportedFormatClasses {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSupportedFormatClassesCopy := &_BACnetConstructedDataSupportedFormatClasses{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.SupportedFormats),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSupportedFormatClassesCopy
}

func (m *_BACnetConstructedDataSupportedFormatClasses) String() string {
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
