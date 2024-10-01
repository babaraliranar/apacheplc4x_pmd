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

// BACnetConstructedDataCOVUPeriod is the corresponding interface of BACnetConstructedDataCOVUPeriod
type BACnetConstructedDataCOVUPeriod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCovuPeriod returns CovuPeriod (property field)
	GetCovuPeriod() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataCOVUPeriod is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCOVUPeriod()
	// CreateBuilder creates a BACnetConstructedDataCOVUPeriodBuilder
	CreateBACnetConstructedDataCOVUPeriodBuilder() BACnetConstructedDataCOVUPeriodBuilder
}

// _BACnetConstructedDataCOVUPeriod is the data-structure of this message
type _BACnetConstructedDataCOVUPeriod struct {
	BACnetConstructedDataContract
	CovuPeriod BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataCOVUPeriod = (*_BACnetConstructedDataCOVUPeriod)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCOVUPeriod)(nil)

// NewBACnetConstructedDataCOVUPeriod factory function for _BACnetConstructedDataCOVUPeriod
func NewBACnetConstructedDataCOVUPeriod(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, covuPeriod BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVUPeriod {
	if covuPeriod == nil {
		panic("covuPeriod of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataCOVUPeriod must not be nil")
	}
	_result := &_BACnetConstructedDataCOVUPeriod{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovuPeriod:                    covuPeriod,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCOVUPeriodBuilder is a builder for BACnetConstructedDataCOVUPeriod
type BACnetConstructedDataCOVUPeriodBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(covuPeriod BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVUPeriodBuilder
	// WithCovuPeriod adds CovuPeriod (property field)
	WithCovuPeriod(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVUPeriodBuilder
	// WithCovuPeriodBuilder adds CovuPeriod (property field) which is build by the builder
	WithCovuPeriodBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCOVUPeriodBuilder
	// Build builds the BACnetConstructedDataCOVUPeriod or returns an error if something is wrong
	Build() (BACnetConstructedDataCOVUPeriod, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCOVUPeriod
}

// NewBACnetConstructedDataCOVUPeriodBuilder() creates a BACnetConstructedDataCOVUPeriodBuilder
func NewBACnetConstructedDataCOVUPeriodBuilder() BACnetConstructedDataCOVUPeriodBuilder {
	return &_BACnetConstructedDataCOVUPeriodBuilder{_BACnetConstructedDataCOVUPeriod: new(_BACnetConstructedDataCOVUPeriod)}
}

type _BACnetConstructedDataCOVUPeriodBuilder struct {
	*_BACnetConstructedDataCOVUPeriod

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCOVUPeriodBuilder) = (*_BACnetConstructedDataCOVUPeriodBuilder)(nil)

func (b *_BACnetConstructedDataCOVUPeriodBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) WithMandatoryFields(covuPeriod BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVUPeriodBuilder {
	return b.WithCovuPeriod(covuPeriod)
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) WithCovuPeriod(covuPeriod BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVUPeriodBuilder {
	b.CovuPeriod = covuPeriod
	return b
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) WithCovuPeriodBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCOVUPeriodBuilder {
	builder := builderSupplier(b.CovuPeriod.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.CovuPeriod, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) Build() (BACnetConstructedDataCOVUPeriod, error) {
	if b.CovuPeriod == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'covuPeriod' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCOVUPeriod.deepCopy(), nil
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) MustBuild() BACnetConstructedDataCOVUPeriod {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataCOVUPeriodBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCOVUPeriodBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCOVUPeriodBuilder().(*_BACnetConstructedDataCOVUPeriodBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCOVUPeriodBuilder creates a BACnetConstructedDataCOVUPeriodBuilder
func (b *_BACnetConstructedDataCOVUPeriod) CreateBACnetConstructedDataCOVUPeriodBuilder() BACnetConstructedDataCOVUPeriodBuilder {
	if b == nil {
		return NewBACnetConstructedDataCOVUPeriodBuilder()
	}
	return &_BACnetConstructedDataCOVUPeriodBuilder{_BACnetConstructedDataCOVUPeriod: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVUPeriod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COVU_PERIOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetCovuPeriod() BACnetApplicationTagUnsignedInteger {
	return m.CovuPeriod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovuPeriod())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVUPeriod(structType any) BACnetConstructedDataCOVUPeriod {
	if casted, ok := structType.(BACnetConstructedDataCOVUPeriod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVUPeriod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVUPeriod) GetTypeName() string {
	return "BACnetConstructedDataCOVUPeriod"
}

func (m *_BACnetConstructedDataCOVUPeriod) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covuPeriod)
	lengthInBits += m.CovuPeriod.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVUPeriod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCOVUPeriod) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCOVUPeriod BACnetConstructedDataCOVUPeriod, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVUPeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVUPeriod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covuPeriod, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covuPeriod", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covuPeriod' field"))
	}
	m.CovuPeriod = covuPeriod

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), covuPeriod)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVUPeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVUPeriod")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCOVUPeriod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVUPeriod) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVUPeriod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVUPeriod")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covuPeriod", m.GetCovuPeriod(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covuPeriod' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVUPeriod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVUPeriod")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVUPeriod) IsBACnetConstructedDataCOVUPeriod() {}

func (m *_BACnetConstructedDataCOVUPeriod) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCOVUPeriod) deepCopy() *_BACnetConstructedDataCOVUPeriod {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCOVUPeriodCopy := &_BACnetConstructedDataCOVUPeriod{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.CovuPeriod.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCOVUPeriodCopy
}

func (m *_BACnetConstructedDataCOVUPeriod) String() string {
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
