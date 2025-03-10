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

// BACnetPriorityValueConstructedValue is the corresponding interface of BACnetPriorityValueConstructedValue
type BACnetPriorityValueConstructedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetConstructedValue returns ConstructedValue (property field)
	GetConstructedValue() BACnetConstructedData
	// IsBACnetPriorityValueConstructedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueConstructedValue()
	// CreateBuilder creates a BACnetPriorityValueConstructedValueBuilder
	CreateBACnetPriorityValueConstructedValueBuilder() BACnetPriorityValueConstructedValueBuilder
}

// _BACnetPriorityValueConstructedValue is the data-structure of this message
type _BACnetPriorityValueConstructedValue struct {
	BACnetPriorityValueContract
	ConstructedValue BACnetConstructedData
}

var _ BACnetPriorityValueConstructedValue = (*_BACnetPriorityValueConstructedValue)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueConstructedValue)(nil)

// NewBACnetPriorityValueConstructedValue factory function for _BACnetPriorityValueConstructedValue
func NewBACnetPriorityValueConstructedValue(peekedTagHeader BACnetTagHeader, constructedValue BACnetConstructedData, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueConstructedValue {
	if constructedValue == nil {
		panic("constructedValue of type BACnetConstructedData for BACnetPriorityValueConstructedValue must not be nil")
	}
	_result := &_BACnetPriorityValueConstructedValue{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		ConstructedValue:            constructedValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueConstructedValueBuilder is a builder for BACnetPriorityValueConstructedValue
type BACnetPriorityValueConstructedValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(constructedValue BACnetConstructedData) BACnetPriorityValueConstructedValueBuilder
	// WithConstructedValue adds ConstructedValue (property field)
	WithConstructedValue(BACnetConstructedData) BACnetPriorityValueConstructedValueBuilder
	// WithConstructedValueBuilder adds ConstructedValue (property field) which is build by the builder
	WithConstructedValueBuilder(func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetPriorityValueConstructedValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPriorityValueBuilder
	// Build builds the BACnetPriorityValueConstructedValue or returns an error if something is wrong
	Build() (BACnetPriorityValueConstructedValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueConstructedValue
}

// NewBACnetPriorityValueConstructedValueBuilder() creates a BACnetPriorityValueConstructedValueBuilder
func NewBACnetPriorityValueConstructedValueBuilder() BACnetPriorityValueConstructedValueBuilder {
	return &_BACnetPriorityValueConstructedValueBuilder{_BACnetPriorityValueConstructedValue: new(_BACnetPriorityValueConstructedValue)}
}

type _BACnetPriorityValueConstructedValueBuilder struct {
	*_BACnetPriorityValueConstructedValue

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueConstructedValueBuilder) = (*_BACnetPriorityValueConstructedValueBuilder)(nil)

func (b *_BACnetPriorityValueConstructedValueBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
	contract.(*_BACnetPriorityValue)._SubType = b._BACnetPriorityValueConstructedValue
}

func (b *_BACnetPriorityValueConstructedValueBuilder) WithMandatoryFields(constructedValue BACnetConstructedData) BACnetPriorityValueConstructedValueBuilder {
	return b.WithConstructedValue(constructedValue)
}

func (b *_BACnetPriorityValueConstructedValueBuilder) WithConstructedValue(constructedValue BACnetConstructedData) BACnetPriorityValueConstructedValueBuilder {
	b.ConstructedValue = constructedValue
	return b
}

func (b *_BACnetPriorityValueConstructedValueBuilder) WithConstructedValueBuilder(builderSupplier func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetPriorityValueConstructedValueBuilder {
	builder := builderSupplier(b.ConstructedValue.CreateBACnetConstructedDataBuilder())
	var err error
	b.ConstructedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConstructedDataBuilder failed"))
	}
	return b
}

func (b *_BACnetPriorityValueConstructedValueBuilder) Build() (BACnetPriorityValueConstructedValue, error) {
	if b.ConstructedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'constructedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueConstructedValue.deepCopy(), nil
}

func (b *_BACnetPriorityValueConstructedValueBuilder) MustBuild() BACnetPriorityValueConstructedValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueConstructedValueBuilder) Done() BACnetPriorityValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPriorityValueBuilder().(*_BACnetPriorityValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPriorityValueConstructedValueBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueConstructedValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueConstructedValueBuilder().(*_BACnetPriorityValueConstructedValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueConstructedValueBuilder creates a BACnetPriorityValueConstructedValueBuilder
func (b *_BACnetPriorityValueConstructedValue) CreateBACnetPriorityValueConstructedValueBuilder() BACnetPriorityValueConstructedValueBuilder {
	if b == nil {
		return NewBACnetPriorityValueConstructedValueBuilder()
	}
	return &_BACnetPriorityValueConstructedValueBuilder{_BACnetPriorityValueConstructedValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueConstructedValue) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueConstructedValue) GetConstructedValue() BACnetConstructedData {
	return m.ConstructedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueConstructedValue(structType any) BACnetPriorityValueConstructedValue {
	if casted, ok := structType.(BACnetPriorityValueConstructedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueConstructedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueConstructedValue) GetTypeName() string {
	return "BACnetPriorityValueConstructedValue"
}

func (m *_BACnetPriorityValueConstructedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (constructedValue)
	lengthInBits += m.ConstructedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueConstructedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueConstructedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueConstructedValue BACnetPriorityValueConstructedValue, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueConstructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueConstructedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	constructedValue, err := ReadSimpleField[BACnetConstructedData](ctx, "constructedValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(0)), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'constructedValue' field"))
	}
	m.ConstructedValue = constructedValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueConstructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueConstructedValue")
	}

	return m, nil
}

func (m *_BACnetPriorityValueConstructedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueConstructedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueConstructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueConstructedValue")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "constructedValue", m.GetConstructedValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'constructedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueConstructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueConstructedValue")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueConstructedValue) IsBACnetPriorityValueConstructedValue() {}

func (m *_BACnetPriorityValueConstructedValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueConstructedValue) deepCopy() *_BACnetPriorityValueConstructedValue {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueConstructedValueCopy := &_BACnetPriorityValueConstructedValue{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetConstructedData](m.ConstructedValue),
	}
	_BACnetPriorityValueConstructedValueCopy.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueConstructedValueCopy
}

func (m *_BACnetPriorityValueConstructedValue) String() string {
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
