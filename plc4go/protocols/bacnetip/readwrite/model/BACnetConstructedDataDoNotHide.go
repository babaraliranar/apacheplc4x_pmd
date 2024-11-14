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

// BACnetConstructedDataDoNotHide is the corresponding interface of BACnetConstructedDataDoNotHide
type BACnetConstructedDataDoNotHide interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDoNotHide returns DoNotHide (property field)
	GetDoNotHide() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataDoNotHide is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDoNotHide()
	// CreateBuilder creates a BACnetConstructedDataDoNotHideBuilder
	CreateBACnetConstructedDataDoNotHideBuilder() BACnetConstructedDataDoNotHideBuilder
}

// _BACnetConstructedDataDoNotHide is the data-structure of this message
type _BACnetConstructedDataDoNotHide struct {
	BACnetConstructedDataContract
	DoNotHide BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataDoNotHide = (*_BACnetConstructedDataDoNotHide)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDoNotHide)(nil)

// NewBACnetConstructedDataDoNotHide factory function for _BACnetConstructedDataDoNotHide
func NewBACnetConstructedDataDoNotHide(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, doNotHide BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoNotHide {
	if doNotHide == nil {
		panic("doNotHide of type BACnetApplicationTagBoolean for BACnetConstructedDataDoNotHide must not be nil")
	}
	_result := &_BACnetConstructedDataDoNotHide{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DoNotHide:                     doNotHide,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDoNotHideBuilder is a builder for BACnetConstructedDataDoNotHide
type BACnetConstructedDataDoNotHideBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doNotHide BACnetApplicationTagBoolean) BACnetConstructedDataDoNotHideBuilder
	// WithDoNotHide adds DoNotHide (property field)
	WithDoNotHide(BACnetApplicationTagBoolean) BACnetConstructedDataDoNotHideBuilder
	// WithDoNotHideBuilder adds DoNotHide (property field) which is build by the builder
	WithDoNotHideBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataDoNotHideBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDoNotHide or returns an error if something is wrong
	Build() (BACnetConstructedDataDoNotHide, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDoNotHide
}

// NewBACnetConstructedDataDoNotHideBuilder() creates a BACnetConstructedDataDoNotHideBuilder
func NewBACnetConstructedDataDoNotHideBuilder() BACnetConstructedDataDoNotHideBuilder {
	return &_BACnetConstructedDataDoNotHideBuilder{_BACnetConstructedDataDoNotHide: new(_BACnetConstructedDataDoNotHide)}
}

type _BACnetConstructedDataDoNotHideBuilder struct {
	*_BACnetConstructedDataDoNotHide

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDoNotHideBuilder) = (*_BACnetConstructedDataDoNotHideBuilder)(nil)

func (b *_BACnetConstructedDataDoNotHideBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDoNotHide
}

func (b *_BACnetConstructedDataDoNotHideBuilder) WithMandatoryFields(doNotHide BACnetApplicationTagBoolean) BACnetConstructedDataDoNotHideBuilder {
	return b.WithDoNotHide(doNotHide)
}

func (b *_BACnetConstructedDataDoNotHideBuilder) WithDoNotHide(doNotHide BACnetApplicationTagBoolean) BACnetConstructedDataDoNotHideBuilder {
	b.DoNotHide = doNotHide
	return b
}

func (b *_BACnetConstructedDataDoNotHideBuilder) WithDoNotHideBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataDoNotHideBuilder {
	builder := builderSupplier(b.DoNotHide.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.DoNotHide, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDoNotHideBuilder) Build() (BACnetConstructedDataDoNotHide, error) {
	if b.DoNotHide == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doNotHide' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDoNotHide.deepCopy(), nil
}

func (b *_BACnetConstructedDataDoNotHideBuilder) MustBuild() BACnetConstructedDataDoNotHide {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDoNotHideBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDoNotHideBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDoNotHideBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDoNotHideBuilder().(*_BACnetConstructedDataDoNotHideBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDoNotHideBuilder creates a BACnetConstructedDataDoNotHideBuilder
func (b *_BACnetConstructedDataDoNotHide) CreateBACnetConstructedDataDoNotHideBuilder() BACnetConstructedDataDoNotHideBuilder {
	if b == nil {
		return NewBACnetConstructedDataDoNotHideBuilder()
	}
	return &_BACnetConstructedDataDoNotHideBuilder{_BACnetConstructedDataDoNotHide: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoNotHide) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DO_NOT_HIDE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetDoNotHide() BACnetApplicationTagBoolean {
	return m.DoNotHide
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetDoNotHide())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoNotHide(structType any) BACnetConstructedDataDoNotHide {
	if casted, ok := structType.(BACnetConstructedDataDoNotHide); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoNotHide); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoNotHide) GetTypeName() string {
	return "BACnetConstructedDataDoNotHide"
}

func (m *_BACnetConstructedDataDoNotHide) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (doNotHide)
	lengthInBits += m.DoNotHide.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoNotHide) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDoNotHide) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDoNotHide BACnetConstructedDataDoNotHide, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoNotHide"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoNotHide")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doNotHide, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "doNotHide", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doNotHide' field"))
	}
	m.DoNotHide = doNotHide

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), doNotHide)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoNotHide"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoNotHide")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDoNotHide) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoNotHide) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoNotHide"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoNotHide")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "doNotHide", m.GetDoNotHide(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doNotHide' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoNotHide"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoNotHide")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoNotHide) IsBACnetConstructedDataDoNotHide() {}

func (m *_BACnetConstructedDataDoNotHide) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDoNotHide) deepCopy() *_BACnetConstructedDataDoNotHide {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDoNotHideCopy := &_BACnetConstructedDataDoNotHide{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.DoNotHide),
	}
	_BACnetConstructedDataDoNotHideCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDoNotHideCopy
}

func (m *_BACnetConstructedDataDoNotHide) String() string {
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
