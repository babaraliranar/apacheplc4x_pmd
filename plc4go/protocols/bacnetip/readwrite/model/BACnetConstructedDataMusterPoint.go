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

// BACnetConstructedDataMusterPoint is the corresponding interface of BACnetConstructedDataMusterPoint
type BACnetConstructedDataMusterPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMusterPoint returns MusterPoint (property field)
	GetMusterPoint() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataMusterPoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMusterPoint()
	// CreateBuilder creates a BACnetConstructedDataMusterPointBuilder
	CreateBACnetConstructedDataMusterPointBuilder() BACnetConstructedDataMusterPointBuilder
}

// _BACnetConstructedDataMusterPoint is the data-structure of this message
type _BACnetConstructedDataMusterPoint struct {
	BACnetConstructedDataContract
	MusterPoint BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataMusterPoint = (*_BACnetConstructedDataMusterPoint)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMusterPoint)(nil)

// NewBACnetConstructedDataMusterPoint factory function for _BACnetConstructedDataMusterPoint
func NewBACnetConstructedDataMusterPoint(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, musterPoint BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMusterPoint {
	if musterPoint == nil {
		panic("musterPoint of type BACnetApplicationTagBoolean for BACnetConstructedDataMusterPoint must not be nil")
	}
	_result := &_BACnetConstructedDataMusterPoint{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MusterPoint:                   musterPoint,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMusterPointBuilder is a builder for BACnetConstructedDataMusterPoint
type BACnetConstructedDataMusterPointBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(musterPoint BACnetApplicationTagBoolean) BACnetConstructedDataMusterPointBuilder
	// WithMusterPoint adds MusterPoint (property field)
	WithMusterPoint(BACnetApplicationTagBoolean) BACnetConstructedDataMusterPointBuilder
	// WithMusterPointBuilder adds MusterPoint (property field) which is build by the builder
	WithMusterPointBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataMusterPointBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMusterPoint or returns an error if something is wrong
	Build() (BACnetConstructedDataMusterPoint, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMusterPoint
}

// NewBACnetConstructedDataMusterPointBuilder() creates a BACnetConstructedDataMusterPointBuilder
func NewBACnetConstructedDataMusterPointBuilder() BACnetConstructedDataMusterPointBuilder {
	return &_BACnetConstructedDataMusterPointBuilder{_BACnetConstructedDataMusterPoint: new(_BACnetConstructedDataMusterPoint)}
}

type _BACnetConstructedDataMusterPointBuilder struct {
	*_BACnetConstructedDataMusterPoint

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMusterPointBuilder) = (*_BACnetConstructedDataMusterPointBuilder)(nil)

func (b *_BACnetConstructedDataMusterPointBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMusterPoint
}

func (b *_BACnetConstructedDataMusterPointBuilder) WithMandatoryFields(musterPoint BACnetApplicationTagBoolean) BACnetConstructedDataMusterPointBuilder {
	return b.WithMusterPoint(musterPoint)
}

func (b *_BACnetConstructedDataMusterPointBuilder) WithMusterPoint(musterPoint BACnetApplicationTagBoolean) BACnetConstructedDataMusterPointBuilder {
	b.MusterPoint = musterPoint
	return b
}

func (b *_BACnetConstructedDataMusterPointBuilder) WithMusterPointBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataMusterPointBuilder {
	builder := builderSupplier(b.MusterPoint.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.MusterPoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMusterPointBuilder) Build() (BACnetConstructedDataMusterPoint, error) {
	if b.MusterPoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'musterPoint' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMusterPoint.deepCopy(), nil
}

func (b *_BACnetConstructedDataMusterPointBuilder) MustBuild() BACnetConstructedDataMusterPoint {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMusterPointBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMusterPointBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMusterPointBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMusterPointBuilder().(*_BACnetConstructedDataMusterPointBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMusterPointBuilder creates a BACnetConstructedDataMusterPointBuilder
func (b *_BACnetConstructedDataMusterPoint) CreateBACnetConstructedDataMusterPointBuilder() BACnetConstructedDataMusterPointBuilder {
	if b == nil {
		return NewBACnetConstructedDataMusterPointBuilder()
	}
	return &_BACnetConstructedDataMusterPointBuilder{_BACnetConstructedDataMusterPoint: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMusterPoint) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MUSTER_POINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetMusterPoint() BACnetApplicationTagBoolean {
	return m.MusterPoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetMusterPoint())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMusterPoint(structType any) BACnetConstructedDataMusterPoint {
	if casted, ok := structType.(BACnetConstructedDataMusterPoint); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMusterPoint); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMusterPoint) GetTypeName() string {
	return "BACnetConstructedDataMusterPoint"
}

func (m *_BACnetConstructedDataMusterPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (musterPoint)
	lengthInBits += m.MusterPoint.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMusterPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMusterPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMusterPoint BACnetConstructedDataMusterPoint, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMusterPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMusterPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	musterPoint, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "musterPoint", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'musterPoint' field"))
	}
	m.MusterPoint = musterPoint

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), musterPoint)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMusterPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMusterPoint")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMusterPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMusterPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMusterPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMusterPoint")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "musterPoint", m.GetMusterPoint(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'musterPoint' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMusterPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMusterPoint")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMusterPoint) IsBACnetConstructedDataMusterPoint() {}

func (m *_BACnetConstructedDataMusterPoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMusterPoint) deepCopy() *_BACnetConstructedDataMusterPoint {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMusterPointCopy := &_BACnetConstructedDataMusterPoint{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.MusterPoint),
	}
	_BACnetConstructedDataMusterPointCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMusterPointCopy
}

func (m *_BACnetConstructedDataMusterPoint) String() string {
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
