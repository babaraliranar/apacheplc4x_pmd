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

// BACnetConstructedDataEgressTime is the corresponding interface of BACnetConstructedDataEgressTime
type BACnetConstructedDataEgressTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEgressTime returns EgressTime (property field)
	GetEgressTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataEgressTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEgressTime()
	// CreateBuilder creates a BACnetConstructedDataEgressTimeBuilder
	CreateBACnetConstructedDataEgressTimeBuilder() BACnetConstructedDataEgressTimeBuilder
}

// _BACnetConstructedDataEgressTime is the data-structure of this message
type _BACnetConstructedDataEgressTime struct {
	BACnetConstructedDataContract
	EgressTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataEgressTime = (*_BACnetConstructedDataEgressTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEgressTime)(nil)

// NewBACnetConstructedDataEgressTime factory function for _BACnetConstructedDataEgressTime
func NewBACnetConstructedDataEgressTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, egressTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEgressTime {
	if egressTime == nil {
		panic("egressTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataEgressTime must not be nil")
	}
	_result := &_BACnetConstructedDataEgressTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EgressTime:                    egressTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEgressTimeBuilder is a builder for BACnetConstructedDataEgressTime
type BACnetConstructedDataEgressTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(egressTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataEgressTimeBuilder
	// WithEgressTime adds EgressTime (property field)
	WithEgressTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataEgressTimeBuilder
	// WithEgressTimeBuilder adds EgressTime (property field) which is build by the builder
	WithEgressTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataEgressTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEgressTime or returns an error if something is wrong
	Build() (BACnetConstructedDataEgressTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEgressTime
}

// NewBACnetConstructedDataEgressTimeBuilder() creates a BACnetConstructedDataEgressTimeBuilder
func NewBACnetConstructedDataEgressTimeBuilder() BACnetConstructedDataEgressTimeBuilder {
	return &_BACnetConstructedDataEgressTimeBuilder{_BACnetConstructedDataEgressTime: new(_BACnetConstructedDataEgressTime)}
}

type _BACnetConstructedDataEgressTimeBuilder struct {
	*_BACnetConstructedDataEgressTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEgressTimeBuilder) = (*_BACnetConstructedDataEgressTimeBuilder)(nil)

func (b *_BACnetConstructedDataEgressTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEgressTime
}

func (b *_BACnetConstructedDataEgressTimeBuilder) WithMandatoryFields(egressTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataEgressTimeBuilder {
	return b.WithEgressTime(egressTime)
}

func (b *_BACnetConstructedDataEgressTimeBuilder) WithEgressTime(egressTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataEgressTimeBuilder {
	b.EgressTime = egressTime
	return b
}

func (b *_BACnetConstructedDataEgressTimeBuilder) WithEgressTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataEgressTimeBuilder {
	builder := builderSupplier(b.EgressTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.EgressTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEgressTimeBuilder) Build() (BACnetConstructedDataEgressTime, error) {
	if b.EgressTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'egressTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEgressTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataEgressTimeBuilder) MustBuild() BACnetConstructedDataEgressTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEgressTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEgressTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEgressTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEgressTimeBuilder().(*_BACnetConstructedDataEgressTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEgressTimeBuilder creates a BACnetConstructedDataEgressTimeBuilder
func (b *_BACnetConstructedDataEgressTime) CreateBACnetConstructedDataEgressTimeBuilder() BACnetConstructedDataEgressTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataEgressTimeBuilder()
	}
	return &_BACnetConstructedDataEgressTimeBuilder{_BACnetConstructedDataEgressTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEgressTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEgressTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EGRESS_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEgressTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEgressTime) GetEgressTime() BACnetApplicationTagUnsignedInteger {
	return m.EgressTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEgressTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetEgressTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEgressTime(structType any) BACnetConstructedDataEgressTime {
	if casted, ok := structType.(BACnetConstructedDataEgressTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEgressTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEgressTime) GetTypeName() string {
	return "BACnetConstructedDataEgressTime"
}

func (m *_BACnetConstructedDataEgressTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (egressTime)
	lengthInBits += m.EgressTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEgressTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEgressTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEgressTime BACnetConstructedDataEgressTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEgressTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEgressTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	egressTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "egressTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'egressTime' field"))
	}
	m.EgressTime = egressTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), egressTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEgressTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEgressTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEgressTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEgressTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEgressTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEgressTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "egressTime", m.GetEgressTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'egressTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEgressTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEgressTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEgressTime) IsBACnetConstructedDataEgressTime() {}

func (m *_BACnetConstructedDataEgressTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEgressTime) deepCopy() *_BACnetConstructedDataEgressTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEgressTimeCopy := &_BACnetConstructedDataEgressTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.EgressTime),
	}
	_BACnetConstructedDataEgressTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEgressTimeCopy
}

func (m *_BACnetConstructedDataEgressTime) String() string {
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
