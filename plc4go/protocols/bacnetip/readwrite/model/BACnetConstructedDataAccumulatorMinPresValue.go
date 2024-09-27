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

// BACnetConstructedDataAccumulatorMinPresValue is the corresponding interface of BACnetConstructedDataAccumulatorMinPresValue
type BACnetConstructedDataAccumulatorMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAccumulatorMinPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccumulatorMinPresValue()
	// CreateBuilder creates a BACnetConstructedDataAccumulatorMinPresValueBuilder
	CreateBACnetConstructedDataAccumulatorMinPresValueBuilder() BACnetConstructedDataAccumulatorMinPresValueBuilder
}

// _BACnetConstructedDataAccumulatorMinPresValue is the data-structure of this message
type _BACnetConstructedDataAccumulatorMinPresValue struct {
	BACnetConstructedDataContract
	MinPresValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAccumulatorMinPresValue = (*_BACnetConstructedDataAccumulatorMinPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccumulatorMinPresValue)(nil)

// NewBACnetConstructedDataAccumulatorMinPresValue factory function for _BACnetConstructedDataAccumulatorMinPresValue
func NewBACnetConstructedDataAccumulatorMinPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minPresValue BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorMinPresValue {
	if minPresValue == nil {
		panic("minPresValue of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAccumulatorMinPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataAccumulatorMinPresValue{
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

// BACnetConstructedDataAccumulatorMinPresValueBuilder is a builder for BACnetConstructedDataAccumulatorMinPresValue
type BACnetConstructedDataAccumulatorMinPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMinPresValueBuilder
	// WithMinPresValue adds MinPresValue (property field)
	WithMinPresValue(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMinPresValueBuilder
	// WithMinPresValueBuilder adds MinPresValue (property field) which is build by the builder
	WithMinPresValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorMinPresValueBuilder
	// Build builds the BACnetConstructedDataAccumulatorMinPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAccumulatorMinPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccumulatorMinPresValue
}

// NewBACnetConstructedDataAccumulatorMinPresValueBuilder() creates a BACnetConstructedDataAccumulatorMinPresValueBuilder
func NewBACnetConstructedDataAccumulatorMinPresValueBuilder() BACnetConstructedDataAccumulatorMinPresValueBuilder {
	return &_BACnetConstructedDataAccumulatorMinPresValueBuilder{_BACnetConstructedDataAccumulatorMinPresValue: new(_BACnetConstructedDataAccumulatorMinPresValue)}
}

type _BACnetConstructedDataAccumulatorMinPresValueBuilder struct {
	*_BACnetConstructedDataAccumulatorMinPresValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccumulatorMinPresValueBuilder) = (*_BACnetConstructedDataAccumulatorMinPresValueBuilder)(nil)

func (m *_BACnetConstructedDataAccumulatorMinPresValueBuilder) WithMandatoryFields(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMinPresValueBuilder {
	return m.WithMinPresValue(minPresValue)
}

func (m *_BACnetConstructedDataAccumulatorMinPresValueBuilder) WithMinPresValue(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMinPresValueBuilder {
	m.MinPresValue = minPresValue
	return m
}

func (m *_BACnetConstructedDataAccumulatorMinPresValueBuilder) WithMinPresValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorMinPresValueBuilder {
	builder := builderSupplier(m.MinPresValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.MinPresValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataAccumulatorMinPresValueBuilder) Build() (BACnetConstructedDataAccumulatorMinPresValue, error) {
	if m.MinPresValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'minPresValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataAccumulatorMinPresValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValueBuilder) MustBuild() BACnetConstructedDataAccumulatorMinPresValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataAccumulatorMinPresValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataAccumulatorMinPresValueBuilder()
}

// CreateBACnetConstructedDataAccumulatorMinPresValueBuilder creates a BACnetConstructedDataAccumulatorMinPresValueBuilder
func (m *_BACnetConstructedDataAccumulatorMinPresValue) CreateBACnetConstructedDataAccumulatorMinPresValueBuilder() BACnetConstructedDataAccumulatorMinPresValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataAccumulatorMinPresValueBuilder()
	}
	return &_BACnetConstructedDataAccumulatorMinPresValueBuilder{_BACnetConstructedDataAccumulatorMinPresValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetMinPresValue() BACnetApplicationTagUnsignedInteger {
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

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorMinPresValue(structType any) BACnetConstructedDataAccumulatorMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorMinPresValue"
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccumulatorMinPresValue BACnetConstructedDataAccumulatorMinPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minPresValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minPresValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPresValue' field"))
	}
	m.MinPresValue = minPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), minPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorMinPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorMinPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minPresValue", m.GetMinPresValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorMinPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) IsBACnetConstructedDataAccumulatorMinPresValue() {
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) deepCopy() *_BACnetConstructedDataAccumulatorMinPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccumulatorMinPresValueCopy := &_BACnetConstructedDataAccumulatorMinPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MinPresValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccumulatorMinPresValueCopy
}

func (m *_BACnetConstructedDataAccumulatorMinPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
