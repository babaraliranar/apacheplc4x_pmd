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

// BACnetConstructedDataChannelNumber is the corresponding interface of BACnetConstructedDataChannelNumber
type BACnetConstructedDataChannelNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetChannelNumber returns ChannelNumber (property field)
	GetChannelNumber() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataChannelNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataChannelNumber()
	// CreateBuilder creates a BACnetConstructedDataChannelNumberBuilder
	CreateBACnetConstructedDataChannelNumberBuilder() BACnetConstructedDataChannelNumberBuilder
}

// _BACnetConstructedDataChannelNumber is the data-structure of this message
type _BACnetConstructedDataChannelNumber struct {
	BACnetConstructedDataContract
	ChannelNumber BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataChannelNumber = (*_BACnetConstructedDataChannelNumber)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataChannelNumber)(nil)

// NewBACnetConstructedDataChannelNumber factory function for _BACnetConstructedDataChannelNumber
func NewBACnetConstructedDataChannelNumber(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, channelNumber BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChannelNumber {
	if channelNumber == nil {
		panic("channelNumber of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataChannelNumber must not be nil")
	}
	_result := &_BACnetConstructedDataChannelNumber{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ChannelNumber:                 channelNumber,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataChannelNumberBuilder is a builder for BACnetConstructedDataChannelNumber
type BACnetConstructedDataChannelNumberBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(channelNumber BACnetApplicationTagUnsignedInteger) BACnetConstructedDataChannelNumberBuilder
	// WithChannelNumber adds ChannelNumber (property field)
	WithChannelNumber(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataChannelNumberBuilder
	// WithChannelNumberBuilder adds ChannelNumber (property field) which is build by the builder
	WithChannelNumberBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataChannelNumberBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataChannelNumber or returns an error if something is wrong
	Build() (BACnetConstructedDataChannelNumber, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataChannelNumber
}

// NewBACnetConstructedDataChannelNumberBuilder() creates a BACnetConstructedDataChannelNumberBuilder
func NewBACnetConstructedDataChannelNumberBuilder() BACnetConstructedDataChannelNumberBuilder {
	return &_BACnetConstructedDataChannelNumberBuilder{_BACnetConstructedDataChannelNumber: new(_BACnetConstructedDataChannelNumber)}
}

type _BACnetConstructedDataChannelNumberBuilder struct {
	*_BACnetConstructedDataChannelNumber

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataChannelNumberBuilder) = (*_BACnetConstructedDataChannelNumberBuilder)(nil)

func (b *_BACnetConstructedDataChannelNumberBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataChannelNumber
}

func (b *_BACnetConstructedDataChannelNumberBuilder) WithMandatoryFields(channelNumber BACnetApplicationTagUnsignedInteger) BACnetConstructedDataChannelNumberBuilder {
	return b.WithChannelNumber(channelNumber)
}

func (b *_BACnetConstructedDataChannelNumberBuilder) WithChannelNumber(channelNumber BACnetApplicationTagUnsignedInteger) BACnetConstructedDataChannelNumberBuilder {
	b.ChannelNumber = channelNumber
	return b
}

func (b *_BACnetConstructedDataChannelNumberBuilder) WithChannelNumberBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataChannelNumberBuilder {
	builder := builderSupplier(b.ChannelNumber.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ChannelNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataChannelNumberBuilder) Build() (BACnetConstructedDataChannelNumber, error) {
	if b.ChannelNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'channelNumber' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataChannelNumber.deepCopy(), nil
}

func (b *_BACnetConstructedDataChannelNumberBuilder) MustBuild() BACnetConstructedDataChannelNumber {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataChannelNumberBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataChannelNumberBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataChannelNumberBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataChannelNumberBuilder().(*_BACnetConstructedDataChannelNumberBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataChannelNumberBuilder creates a BACnetConstructedDataChannelNumberBuilder
func (b *_BACnetConstructedDataChannelNumber) CreateBACnetConstructedDataChannelNumberBuilder() BACnetConstructedDataChannelNumberBuilder {
	if b == nil {
		return NewBACnetConstructedDataChannelNumberBuilder()
	}
	return &_BACnetConstructedDataChannelNumberBuilder{_BACnetConstructedDataChannelNumber: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChannelNumber) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANNEL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetChannelNumber() BACnetApplicationTagUnsignedInteger {
	return m.ChannelNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetChannelNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChannelNumber(structType any) BACnetConstructedDataChannelNumber {
	if casted, ok := structType.(BACnetConstructedDataChannelNumber); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChannelNumber); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChannelNumber) GetTypeName() string {
	return "BACnetConstructedDataChannelNumber"
}

func (m *_BACnetConstructedDataChannelNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (channelNumber)
	lengthInBits += m.ChannelNumber.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChannelNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataChannelNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataChannelNumber BACnetConstructedDataChannelNumber, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChannelNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChannelNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	channelNumber, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "channelNumber", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelNumber' field"))
	}
	m.ChannelNumber = channelNumber

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), channelNumber)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChannelNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChannelNumber")
	}

	return m, nil
}

func (m *_BACnetConstructedDataChannelNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChannelNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChannelNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChannelNumber")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "channelNumber", m.GetChannelNumber(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'channelNumber' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChannelNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChannelNumber")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChannelNumber) IsBACnetConstructedDataChannelNumber() {}

func (m *_BACnetConstructedDataChannelNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataChannelNumber) deepCopy() *_BACnetConstructedDataChannelNumber {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataChannelNumberCopy := &_BACnetConstructedDataChannelNumber{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ChannelNumber),
	}
	_BACnetConstructedDataChannelNumberCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataChannelNumberCopy
}

func (m *_BACnetConstructedDataChannelNumber) String() string {
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
