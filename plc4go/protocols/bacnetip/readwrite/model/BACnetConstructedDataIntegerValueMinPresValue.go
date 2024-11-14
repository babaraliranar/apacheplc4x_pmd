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

// BACnetConstructedDataIntegerValueMinPresValue is the corresponding interface of BACnetConstructedDataIntegerValueMinPresValue
type BACnetConstructedDataIntegerValueMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataIntegerValueMinPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueMinPresValue()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueMinPresValueBuilder
	CreateBACnetConstructedDataIntegerValueMinPresValueBuilder() BACnetConstructedDataIntegerValueMinPresValueBuilder
}

// _BACnetConstructedDataIntegerValueMinPresValue is the data-structure of this message
type _BACnetConstructedDataIntegerValueMinPresValue struct {
	BACnetConstructedDataContract
	MinPresValue BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataIntegerValueMinPresValue = (*_BACnetConstructedDataIntegerValueMinPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueMinPresValue)(nil)

// NewBACnetConstructedDataIntegerValueMinPresValue factory function for _BACnetConstructedDataIntegerValueMinPresValue
func NewBACnetConstructedDataIntegerValueMinPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minPresValue BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueMinPresValue {
	if minPresValue == nil {
		panic("minPresValue of type BACnetApplicationTagSignedInteger for BACnetConstructedDataIntegerValueMinPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueMinPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinPresValue:                  minPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueMinPresValueBuilder is a builder for BACnetConstructedDataIntegerValueMinPresValue
type BACnetConstructedDataIntegerValueMinPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minPresValue BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueMinPresValueBuilder
	// WithMinPresValue adds MinPresValue (property field)
	WithMinPresValue(BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueMinPresValueBuilder
	// WithMinPresValueBuilder adds MinPresValue (property field) which is build by the builder
	WithMinPresValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueMinPresValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIntegerValueMinPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueMinPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueMinPresValue
}

// NewBACnetConstructedDataIntegerValueMinPresValueBuilder() creates a BACnetConstructedDataIntegerValueMinPresValueBuilder
func NewBACnetConstructedDataIntegerValueMinPresValueBuilder() BACnetConstructedDataIntegerValueMinPresValueBuilder {
	return &_BACnetConstructedDataIntegerValueMinPresValueBuilder{_BACnetConstructedDataIntegerValueMinPresValue: new(_BACnetConstructedDataIntegerValueMinPresValue)}
}

type _BACnetConstructedDataIntegerValueMinPresValueBuilder struct {
	*_BACnetConstructedDataIntegerValueMinPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueMinPresValueBuilder) = (*_BACnetConstructedDataIntegerValueMinPresValueBuilder)(nil)

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIntegerValueMinPresValue
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) WithMandatoryFields(minPresValue BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueMinPresValueBuilder {
	return b.WithMinPresValue(minPresValue)
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) WithMinPresValue(minPresValue BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueMinPresValueBuilder {
	b.MinPresValue = minPresValue
	return b
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) WithMinPresValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueMinPresValueBuilder {
	builder := builderSupplier(b.MinPresValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.MinPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) Build() (BACnetConstructedDataIntegerValueMinPresValue, error) {
	if b.MinPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'minPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegerValueMinPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) MustBuild() BACnetConstructedDataIntegerValueMinPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegerValueMinPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegerValueMinPresValueBuilder().(*_BACnetConstructedDataIntegerValueMinPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegerValueMinPresValueBuilder creates a BACnetConstructedDataIntegerValueMinPresValueBuilder
func (b *_BACnetConstructedDataIntegerValueMinPresValue) CreateBACnetConstructedDataIntegerValueMinPresValueBuilder() BACnetConstructedDataIntegerValueMinPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegerValueMinPresValueBuilder()
	}
	return &_BACnetConstructedDataIntegerValueMinPresValueBuilder{_BACnetConstructedDataIntegerValueMinPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetMinPresValue() BACnetApplicationTagSignedInteger {
	return m.MinPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueMinPresValue(structType any) BACnetConstructedDataIntegerValueMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueMinPresValue"
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueMinPresValue BACnetConstructedDataIntegerValueMinPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minPresValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "minPresValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPresValue' field"))
	}
	m.MinPresValue = minPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), minPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueMinPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueMinPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "minPresValue", m.GetMinPresValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueMinPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) IsBACnetConstructedDataIntegerValueMinPresValue() {
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) deepCopy() *_BACnetConstructedDataIntegerValueMinPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueMinPresValueCopy := &_BACnetConstructedDataIntegerValueMinPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.MinPresValue),
	}
	_BACnetConstructedDataIntegerValueMinPresValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueMinPresValueCopy
}

func (m *_BACnetConstructedDataIntegerValueMinPresValue) String() string {
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
