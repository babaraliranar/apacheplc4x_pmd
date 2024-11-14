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

// BACnetConstructedDataNetworkType is the corresponding interface of BACnetConstructedDataNetworkType
type BACnetConstructedDataNetworkType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNetworkType returns NetworkType (property field)
	GetNetworkType() BACnetNetworkTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetNetworkTypeTagged
	// IsBACnetConstructedDataNetworkType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNetworkType()
	// CreateBuilder creates a BACnetConstructedDataNetworkTypeBuilder
	CreateBACnetConstructedDataNetworkTypeBuilder() BACnetConstructedDataNetworkTypeBuilder
}

// _BACnetConstructedDataNetworkType is the data-structure of this message
type _BACnetConstructedDataNetworkType struct {
	BACnetConstructedDataContract
	NetworkType BACnetNetworkTypeTagged
}

var _ BACnetConstructedDataNetworkType = (*_BACnetConstructedDataNetworkType)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNetworkType)(nil)

// NewBACnetConstructedDataNetworkType factory function for _BACnetConstructedDataNetworkType
func NewBACnetConstructedDataNetworkType(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, networkType BACnetNetworkTypeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkType {
	if networkType == nil {
		panic("networkType of type BACnetNetworkTypeTagged for BACnetConstructedDataNetworkType must not be nil")
	}
	_result := &_BACnetConstructedDataNetworkType{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NetworkType:                   networkType,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNetworkTypeBuilder is a builder for BACnetConstructedDataNetworkType
type BACnetConstructedDataNetworkTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkType BACnetNetworkTypeTagged) BACnetConstructedDataNetworkTypeBuilder
	// WithNetworkType adds NetworkType (property field)
	WithNetworkType(BACnetNetworkTypeTagged) BACnetConstructedDataNetworkTypeBuilder
	// WithNetworkTypeBuilder adds NetworkType (property field) which is build by the builder
	WithNetworkTypeBuilder(func(BACnetNetworkTypeTaggedBuilder) BACnetNetworkTypeTaggedBuilder) BACnetConstructedDataNetworkTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataNetworkType or returns an error if something is wrong
	Build() (BACnetConstructedDataNetworkType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNetworkType
}

// NewBACnetConstructedDataNetworkTypeBuilder() creates a BACnetConstructedDataNetworkTypeBuilder
func NewBACnetConstructedDataNetworkTypeBuilder() BACnetConstructedDataNetworkTypeBuilder {
	return &_BACnetConstructedDataNetworkTypeBuilder{_BACnetConstructedDataNetworkType: new(_BACnetConstructedDataNetworkType)}
}

type _BACnetConstructedDataNetworkTypeBuilder struct {
	*_BACnetConstructedDataNetworkType

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNetworkTypeBuilder) = (*_BACnetConstructedDataNetworkTypeBuilder)(nil)

func (b *_BACnetConstructedDataNetworkTypeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataNetworkType
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) WithMandatoryFields(networkType BACnetNetworkTypeTagged) BACnetConstructedDataNetworkTypeBuilder {
	return b.WithNetworkType(networkType)
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) WithNetworkType(networkType BACnetNetworkTypeTagged) BACnetConstructedDataNetworkTypeBuilder {
	b.NetworkType = networkType
	return b
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) WithNetworkTypeBuilder(builderSupplier func(BACnetNetworkTypeTaggedBuilder) BACnetNetworkTypeTaggedBuilder) BACnetConstructedDataNetworkTypeBuilder {
	builder := builderSupplier(b.NetworkType.CreateBACnetNetworkTypeTaggedBuilder())
	var err error
	b.NetworkType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNetworkTypeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) Build() (BACnetConstructedDataNetworkType, error) {
	if b.NetworkType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'networkType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNetworkType.deepCopy(), nil
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) MustBuild() BACnetConstructedDataNetworkType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNetworkTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNetworkTypeBuilder().(*_BACnetConstructedDataNetworkTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNetworkTypeBuilder creates a BACnetConstructedDataNetworkTypeBuilder
func (b *_BACnetConstructedDataNetworkType) CreateBACnetConstructedDataNetworkTypeBuilder() BACnetConstructedDataNetworkTypeBuilder {
	if b == nil {
		return NewBACnetConstructedDataNetworkTypeBuilder()
	}
	return &_BACnetConstructedDataNetworkTypeBuilder{_BACnetConstructedDataNetworkType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNetworkType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkType) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkType) GetNetworkType() BACnetNetworkTypeTagged {
	return m.NetworkType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkType) GetActualValue() BACnetNetworkTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetNetworkTypeTagged(m.GetNetworkType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkType(structType any) BACnetConstructedDataNetworkType {
	if casted, ok := structType.(BACnetConstructedDataNetworkType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkType) GetTypeName() string {
	return "BACnetConstructedDataNetworkType"
}

func (m *_BACnetConstructedDataNetworkType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (networkType)
	lengthInBits += m.NetworkType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNetworkType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNetworkType BACnetConstructedDataNetworkType, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkType, err := ReadSimpleField[BACnetNetworkTypeTagged](ctx, "networkType", ReadComplex[BACnetNetworkTypeTagged](BACnetNetworkTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkType' field"))
	}
	m.NetworkType = networkType

	actualValue, err := ReadVirtualField[BACnetNetworkTypeTagged](ctx, "actualValue", (*BACnetNetworkTypeTagged)(nil), networkType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkType")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNetworkType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkType")
		}

		if err := WriteSimpleField[BACnetNetworkTypeTagged](ctx, "networkType", m.GetNetworkType(), WriteComplex[BACnetNetworkTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkType")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkType) IsBACnetConstructedDataNetworkType() {}

func (m *_BACnetConstructedDataNetworkType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNetworkType) deepCopy() *_BACnetConstructedDataNetworkType {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNetworkTypeCopy := &_BACnetConstructedDataNetworkType{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetNetworkTypeTagged](m.NetworkType),
	}
	_BACnetConstructedDataNetworkTypeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNetworkTypeCopy
}

func (m *_BACnetConstructedDataNetworkType) String() string {
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
