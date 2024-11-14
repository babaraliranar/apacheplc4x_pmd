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

// BACnetConstructedDataBinaryLightingOutputRelinquishDefault is the corresponding interface of BACnetConstructedDataBinaryLightingOutputRelinquishDefault
type BACnetConstructedDataBinaryLightingOutputRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetBinaryLightingPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryLightingPVTagged
	// IsBACnetConstructedDataBinaryLightingOutputRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryLightingOutputRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
	CreateBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder() BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
}

// _BACnetConstructedDataBinaryLightingOutputRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataBinaryLightingOutputRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetBinaryLightingPVTagged
}

var _ BACnetConstructedDataBinaryLightingOutputRelinquishDefault = (*_BACnetConstructedDataBinaryLightingOutputRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryLightingOutputRelinquishDefault)(nil)

// NewBACnetConstructedDataBinaryLightingOutputRelinquishDefault factory function for _BACnetConstructedDataBinaryLightingOutputRelinquishDefault
func NewBACnetConstructedDataBinaryLightingOutputRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetBinaryLightingPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetBinaryLightingPVTagged for BACnetConstructedDataBinaryLightingOutputRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryLightingOutputRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder is a builder for BACnetConstructedDataBinaryLightingOutputRelinquishDefault
type BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBinaryLightingOutputRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryLightingOutputRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryLightingOutputRelinquishDefault
}

// NewBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder() creates a BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
func NewBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder() BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder {
	return &_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder{_BACnetConstructedDataBinaryLightingOutputRelinquishDefault: new(_BACnetConstructedDataBinaryLightingOutputRelinquishDefault)}
}

type _BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataBinaryLightingOutputRelinquishDefault

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) = (*_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder)(nil)

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBinaryLightingOutputRelinquishDefault
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder {
	return b.WithRelinquishDefault(relinquishDefault)
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder {
	b.RelinquishDefault = relinquishDefault
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder {
	builder := builderSupplier(b.RelinquishDefault.CreateBACnetBinaryLightingPVTaggedBuilder())
	var err error
	b.RelinquishDefault, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetBinaryLightingPVTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) Build() (BACnetConstructedDataBinaryLightingOutputRelinquishDefault, error) {
	if b.RelinquishDefault == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryLightingOutputRelinquishDefault.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataBinaryLightingOutputRelinquishDefault {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder().(*_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder creates a BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder
func (b *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) CreateBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder() BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultBuilder{_BACnetConstructedDataBinaryLightingOutputRelinquishDefault: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetRelinquishDefault() BACnetBinaryLightingPVTagged {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetActualValue() BACnetBinaryLightingPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryLightingPVTagged(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryLightingOutputRelinquishDefault(structType any) BACnetConstructedDataBinaryLightingOutputRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataBinaryLightingOutputRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryLightingOutputRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataBinaryLightingOutputRelinquishDefault"
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryLightingOutputRelinquishDefault BACnetConstructedDataBinaryLightingOutputRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryLightingOutputRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryLightingOutputRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetBinaryLightingPVTagged](ctx, "relinquishDefault", ReadComplex[BACnetBinaryLightingPVTagged](BACnetBinaryLightingPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetBinaryLightingPVTagged](ctx, "actualValue", (*BACnetBinaryLightingPVTagged)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryLightingOutputRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryLightingOutputRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryLightingOutputRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryLightingOutputRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetBinaryLightingPVTagged](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetBinaryLightingPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryLightingOutputRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryLightingOutputRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) IsBACnetConstructedDataBinaryLightingOutputRelinquishDefault() {
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) deepCopy() *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultCopy := &_BACnetConstructedDataBinaryLightingOutputRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetBinaryLightingPVTagged](m.RelinquishDefault),
	}
	_BACnetConstructedDataBinaryLightingOutputRelinquishDefaultCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryLightingOutputRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataBinaryLightingOutputRelinquishDefault) String() string {
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
