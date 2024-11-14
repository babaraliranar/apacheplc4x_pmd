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

// BACnetConstructedDataAnalogValueFaultLowLimit is the corresponding interface of BACnetConstructedDataAnalogValueFaultLowLimit
type BACnetConstructedDataAnalogValueFaultLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataAnalogValueFaultLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogValueFaultLowLimit()
	// CreateBuilder creates a BACnetConstructedDataAnalogValueFaultLowLimitBuilder
	CreateBACnetConstructedDataAnalogValueFaultLowLimitBuilder() BACnetConstructedDataAnalogValueFaultLowLimitBuilder
}

// _BACnetConstructedDataAnalogValueFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataAnalogValueFaultLowLimit struct {
	BACnetConstructedDataContract
	FaultLowLimit BACnetApplicationTagReal
}

var _ BACnetConstructedDataAnalogValueFaultLowLimit = (*_BACnetConstructedDataAnalogValueFaultLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogValueFaultLowLimit)(nil)

// NewBACnetConstructedDataAnalogValueFaultLowLimit factory function for _BACnetConstructedDataAnalogValueFaultLowLimit
func NewBACnetConstructedDataAnalogValueFaultLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultLowLimit BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogValueFaultLowLimit {
	if faultLowLimit == nil {
		panic("faultLowLimit of type BACnetApplicationTagReal for BACnetConstructedDataAnalogValueFaultLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogValueFaultLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultLowLimit:                 faultLowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogValueFaultLowLimitBuilder is a builder for BACnetConstructedDataAnalogValueFaultLowLimit
type BACnetConstructedDataAnalogValueFaultLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultLowLimit BACnetApplicationTagReal) BACnetConstructedDataAnalogValueFaultLowLimitBuilder
	// WithFaultLowLimit adds FaultLowLimit (property field)
	WithFaultLowLimit(BACnetApplicationTagReal) BACnetConstructedDataAnalogValueFaultLowLimitBuilder
	// WithFaultLowLimitBuilder adds FaultLowLimit (property field) which is build by the builder
	WithFaultLowLimitBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogValueFaultLowLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAnalogValueFaultLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogValueFaultLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogValueFaultLowLimit
}

// NewBACnetConstructedDataAnalogValueFaultLowLimitBuilder() creates a BACnetConstructedDataAnalogValueFaultLowLimitBuilder
func NewBACnetConstructedDataAnalogValueFaultLowLimitBuilder() BACnetConstructedDataAnalogValueFaultLowLimitBuilder {
	return &_BACnetConstructedDataAnalogValueFaultLowLimitBuilder{_BACnetConstructedDataAnalogValueFaultLowLimit: new(_BACnetConstructedDataAnalogValueFaultLowLimit)}
}

type _BACnetConstructedDataAnalogValueFaultLowLimitBuilder struct {
	*_BACnetConstructedDataAnalogValueFaultLowLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogValueFaultLowLimitBuilder) = (*_BACnetConstructedDataAnalogValueFaultLowLimitBuilder)(nil)

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAnalogValueFaultLowLimit
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) WithMandatoryFields(faultLowLimit BACnetApplicationTagReal) BACnetConstructedDataAnalogValueFaultLowLimitBuilder {
	return b.WithFaultLowLimit(faultLowLimit)
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) WithFaultLowLimit(faultLowLimit BACnetApplicationTagReal) BACnetConstructedDataAnalogValueFaultLowLimitBuilder {
	b.FaultLowLimit = faultLowLimit
	return b
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) WithFaultLowLimitBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogValueFaultLowLimitBuilder {
	builder := builderSupplier(b.FaultLowLimit.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.FaultLowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) Build() (BACnetConstructedDataAnalogValueFaultLowLimit, error) {
	if b.FaultLowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'faultLowLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogValueFaultLowLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) MustBuild() BACnetConstructedDataAnalogValueFaultLowLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogValueFaultLowLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogValueFaultLowLimitBuilder().(*_BACnetConstructedDataAnalogValueFaultLowLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogValueFaultLowLimitBuilder creates a BACnetConstructedDataAnalogValueFaultLowLimitBuilder
func (b *_BACnetConstructedDataAnalogValueFaultLowLimit) CreateBACnetConstructedDataAnalogValueFaultLowLimitBuilder() BACnetConstructedDataAnalogValueFaultLowLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogValueFaultLowLimitBuilder()
	}
	return &_BACnetConstructedDataAnalogValueFaultLowLimitBuilder{_BACnetConstructedDataAnalogValueFaultLowLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagReal {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogValueFaultLowLimit(structType any) BACnetConstructedDataAnalogValueFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAnalogValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValueFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAnalogValueFaultLowLimit"
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogValueFaultLowLimit BACnetConstructedDataAnalogValueFaultLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValueFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValueFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultLowLimit, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "faultLowLimit", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultLowLimit' field"))
	}
	m.FaultLowLimit = faultLowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), faultLowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValueFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValueFaultLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValueFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValueFaultLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "faultLowLimit", m.GetFaultLowLimit(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValueFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValueFaultLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) IsBACnetConstructedDataAnalogValueFaultLowLimit() {
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) deepCopy() *_BACnetConstructedDataAnalogValueFaultLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogValueFaultLowLimitCopy := &_BACnetConstructedDataAnalogValueFaultLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.FaultLowLimit),
	}
	_BACnetConstructedDataAnalogValueFaultLowLimitCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogValueFaultLowLimitCopy
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) String() string {
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
