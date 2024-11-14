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

// BACnetConstructedDataFDBBMDAddress is the corresponding interface of BACnetConstructedDataFDBBMDAddress
type BACnetConstructedDataFDBBMDAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFDBBMDAddress returns FDBBMDAddress (property field)
	GetFDBBMDAddress() BACnetHostNPort
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetHostNPort
	// IsBACnetConstructedDataFDBBMDAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFDBBMDAddress()
	// CreateBuilder creates a BACnetConstructedDataFDBBMDAddressBuilder
	CreateBACnetConstructedDataFDBBMDAddressBuilder() BACnetConstructedDataFDBBMDAddressBuilder
}

// _BACnetConstructedDataFDBBMDAddress is the data-structure of this message
type _BACnetConstructedDataFDBBMDAddress struct {
	BACnetConstructedDataContract
	FDBBMDAddress BACnetHostNPort
}

var _ BACnetConstructedDataFDBBMDAddress = (*_BACnetConstructedDataFDBBMDAddress)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFDBBMDAddress)(nil)

// NewBACnetConstructedDataFDBBMDAddress factory function for _BACnetConstructedDataFDBBMDAddress
func NewBACnetConstructedDataFDBBMDAddress(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, fDBBMDAddress BACnetHostNPort, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFDBBMDAddress {
	if fDBBMDAddress == nil {
		panic("fDBBMDAddress of type BACnetHostNPort for BACnetConstructedDataFDBBMDAddress must not be nil")
	}
	_result := &_BACnetConstructedDataFDBBMDAddress{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FDBBMDAddress:                 fDBBMDAddress,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFDBBMDAddressBuilder is a builder for BACnetConstructedDataFDBBMDAddress
type BACnetConstructedDataFDBBMDAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fDBBMDAddress BACnetHostNPort) BACnetConstructedDataFDBBMDAddressBuilder
	// WithFDBBMDAddress adds FDBBMDAddress (property field)
	WithFDBBMDAddress(BACnetHostNPort) BACnetConstructedDataFDBBMDAddressBuilder
	// WithFDBBMDAddressBuilder adds FDBBMDAddress (property field) which is build by the builder
	WithFDBBMDAddressBuilder(func(BACnetHostNPortBuilder) BACnetHostNPortBuilder) BACnetConstructedDataFDBBMDAddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataFDBBMDAddress or returns an error if something is wrong
	Build() (BACnetConstructedDataFDBBMDAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFDBBMDAddress
}

// NewBACnetConstructedDataFDBBMDAddressBuilder() creates a BACnetConstructedDataFDBBMDAddressBuilder
func NewBACnetConstructedDataFDBBMDAddressBuilder() BACnetConstructedDataFDBBMDAddressBuilder {
	return &_BACnetConstructedDataFDBBMDAddressBuilder{_BACnetConstructedDataFDBBMDAddress: new(_BACnetConstructedDataFDBBMDAddress)}
}

type _BACnetConstructedDataFDBBMDAddressBuilder struct {
	*_BACnetConstructedDataFDBBMDAddress

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFDBBMDAddressBuilder) = (*_BACnetConstructedDataFDBBMDAddressBuilder)(nil)

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataFDBBMDAddress
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) WithMandatoryFields(fDBBMDAddress BACnetHostNPort) BACnetConstructedDataFDBBMDAddressBuilder {
	return b.WithFDBBMDAddress(fDBBMDAddress)
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) WithFDBBMDAddress(fDBBMDAddress BACnetHostNPort) BACnetConstructedDataFDBBMDAddressBuilder {
	b.FDBBMDAddress = fDBBMDAddress
	return b
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) WithFDBBMDAddressBuilder(builderSupplier func(BACnetHostNPortBuilder) BACnetHostNPortBuilder) BACnetConstructedDataFDBBMDAddressBuilder {
	builder := builderSupplier(b.FDBBMDAddress.CreateBACnetHostNPortBuilder())
	var err error
	b.FDBBMDAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetHostNPortBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) Build() (BACnetConstructedDataFDBBMDAddress, error) {
	if b.FDBBMDAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fDBBMDAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFDBBMDAddress.deepCopy(), nil
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) MustBuild() BACnetConstructedDataFDBBMDAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFDBBMDAddressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFDBBMDAddressBuilder().(*_BACnetConstructedDataFDBBMDAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFDBBMDAddressBuilder creates a BACnetConstructedDataFDBBMDAddressBuilder
func (b *_BACnetConstructedDataFDBBMDAddress) CreateBACnetConstructedDataFDBBMDAddressBuilder() BACnetConstructedDataFDBBMDAddressBuilder {
	if b == nil {
		return NewBACnetConstructedDataFDBBMDAddressBuilder()
	}
	return &_BACnetConstructedDataFDBBMDAddressBuilder{_BACnetConstructedDataFDBBMDAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FD_BBMD_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetFDBBMDAddress() BACnetHostNPort {
	return m.FDBBMDAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetActualValue() BACnetHostNPort {
	ctx := context.Background()
	_ = ctx
	return CastBACnetHostNPort(m.GetFDBBMDAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFDBBMDAddress(structType any) BACnetConstructedDataFDBBMDAddress {
	if casted, ok := structType.(BACnetConstructedDataFDBBMDAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFDBBMDAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetTypeName() string {
	return "BACnetConstructedDataFDBBMDAddress"
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (fDBBMDAddress)
	lengthInBits += m.FDBBMDAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFDBBMDAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFDBBMDAddress BACnetConstructedDataFDBBMDAddress, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFDBBMDAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFDBBMDAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fDBBMDAddress, err := ReadSimpleField[BACnetHostNPort](ctx, "fDBBMDAddress", ReadComplex[BACnetHostNPort](BACnetHostNPortParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fDBBMDAddress' field"))
	}
	m.FDBBMDAddress = fDBBMDAddress

	actualValue, err := ReadVirtualField[BACnetHostNPort](ctx, "actualValue", (*BACnetHostNPort)(nil), fDBBMDAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFDBBMDAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFDBBMDAddress")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFDBBMDAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFDBBMDAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFDBBMDAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFDBBMDAddress")
		}

		if err := WriteSimpleField[BACnetHostNPort](ctx, "fDBBMDAddress", m.GetFDBBMDAddress(), WriteComplex[BACnetHostNPort](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fDBBMDAddress' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFDBBMDAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFDBBMDAddress")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFDBBMDAddress) IsBACnetConstructedDataFDBBMDAddress() {}

func (m *_BACnetConstructedDataFDBBMDAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFDBBMDAddress) deepCopy() *_BACnetConstructedDataFDBBMDAddress {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFDBBMDAddressCopy := &_BACnetConstructedDataFDBBMDAddress{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetHostNPort](m.FDBBMDAddress),
	}
	_BACnetConstructedDataFDBBMDAddressCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFDBBMDAddressCopy
}

func (m *_BACnetConstructedDataFDBBMDAddress) String() string {
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
