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

// BACnetConstructedDataFloorText is the corresponding interface of BACnetConstructedDataFloorText
type BACnetConstructedDataFloorText interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataFloorText is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFloorText()
	// CreateBuilder creates a BACnetConstructedDataFloorTextBuilder
	CreateBACnetConstructedDataFloorTextBuilder() BACnetConstructedDataFloorTextBuilder
}

// _BACnetConstructedDataFloorText is the data-structure of this message
type _BACnetConstructedDataFloorText struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	FloorText            []BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataFloorText = (*_BACnetConstructedDataFloorText)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFloorText)(nil)

// NewBACnetConstructedDataFloorText factory function for _BACnetConstructedDataFloorText
func NewBACnetConstructedDataFloorText(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, floorText []BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFloorText {
	_result := &_BACnetConstructedDataFloorText{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		FloorText:                     floorText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFloorTextBuilder is a builder for BACnetConstructedDataFloorText
type BACnetConstructedDataFloorTextBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(floorText []BACnetApplicationTagCharacterString) BACnetConstructedDataFloorTextBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFloorTextBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFloorTextBuilder
	// WithFloorText adds FloorText (property field)
	WithFloorText(...BACnetApplicationTagCharacterString) BACnetConstructedDataFloorTextBuilder
	// Build builds the BACnetConstructedDataFloorText or returns an error if something is wrong
	Build() (BACnetConstructedDataFloorText, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFloorText
}

// NewBACnetConstructedDataFloorTextBuilder() creates a BACnetConstructedDataFloorTextBuilder
func NewBACnetConstructedDataFloorTextBuilder() BACnetConstructedDataFloorTextBuilder {
	return &_BACnetConstructedDataFloorTextBuilder{_BACnetConstructedDataFloorText: new(_BACnetConstructedDataFloorText)}
}

type _BACnetConstructedDataFloorTextBuilder struct {
	*_BACnetConstructedDataFloorText

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFloorTextBuilder) = (*_BACnetConstructedDataFloorTextBuilder)(nil)

func (b *_BACnetConstructedDataFloorTextBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataFloorTextBuilder) WithMandatoryFields(floorText []BACnetApplicationTagCharacterString) BACnetConstructedDataFloorTextBuilder {
	return b.WithFloorText(floorText...)
}

func (b *_BACnetConstructedDataFloorTextBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFloorTextBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataFloorTextBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFloorTextBuilder {
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

func (b *_BACnetConstructedDataFloorTextBuilder) WithFloorText(floorText ...BACnetApplicationTagCharacterString) BACnetConstructedDataFloorTextBuilder {
	b.FloorText = floorText
	return b
}

func (b *_BACnetConstructedDataFloorTextBuilder) Build() (BACnetConstructedDataFloorText, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFloorText.deepCopy(), nil
}

func (b *_BACnetConstructedDataFloorTextBuilder) MustBuild() BACnetConstructedDataFloorText {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataFloorTextBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFloorTextBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFloorTextBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFloorTextBuilder().(*_BACnetConstructedDataFloorTextBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFloorTextBuilder creates a BACnetConstructedDataFloorTextBuilder
func (b *_BACnetConstructedDataFloorText) CreateBACnetConstructedDataFloorTextBuilder() BACnetConstructedDataFloorTextBuilder {
	if b == nil {
		return NewBACnetConstructedDataFloorTextBuilder()
	}
	return &_BACnetConstructedDataFloorTextBuilder{_BACnetConstructedDataFloorText: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFloorText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFloorText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FLOOR_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFloorText) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFloorText) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataFloorText) GetFloorText() []BACnetApplicationTagCharacterString {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFloorText) GetZero() uint64 {
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
func CastBACnetConstructedDataFloorText(structType any) BACnetConstructedDataFloorText {
	if casted, ok := structType.(BACnetConstructedDataFloorText); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFloorText); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFloorText) GetTypeName() string {
	return "BACnetConstructedDataFloorText"
}

func (m *_BACnetConstructedDataFloorText) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataFloorText) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFloorText) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFloorText BACnetConstructedDataFloorText, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFloorText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFloorText")
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

	floorText, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "floorText", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorText' field"))
	}
	m.FloorText = floorText

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFloorText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFloorText")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFloorText) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFloorText) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFloorText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFloorText")
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

		if err := WriteComplexTypeArrayField(ctx, "floorText", m.GetFloorText(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'floorText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFloorText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFloorText")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFloorText) IsBACnetConstructedDataFloorText() {}

func (m *_BACnetConstructedDataFloorText) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFloorText) deepCopy() *_BACnetConstructedDataFloorText {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFloorTextCopy := &_BACnetConstructedDataFloorText{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagCharacterString, BACnetApplicationTagCharacterString](m.FloorText),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFloorTextCopy
}

func (m *_BACnetConstructedDataFloorText) String() string {
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
