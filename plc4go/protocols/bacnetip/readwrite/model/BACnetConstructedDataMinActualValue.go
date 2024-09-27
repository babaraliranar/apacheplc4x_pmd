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

// BACnetConstructedDataMinActualValue is the corresponding interface of BACnetConstructedDataMinActualValue
type BACnetConstructedDataMinActualValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinActualValue returns MinActualValue (property field)
	GetMinActualValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataMinActualValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMinActualValue()
	// CreateBuilder creates a BACnetConstructedDataMinActualValueBuilder
	CreateBACnetConstructedDataMinActualValueBuilder() BACnetConstructedDataMinActualValueBuilder
}

// _BACnetConstructedDataMinActualValue is the data-structure of this message
type _BACnetConstructedDataMinActualValue struct {
	BACnetConstructedDataContract
	MinActualValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataMinActualValue = (*_BACnetConstructedDataMinActualValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMinActualValue)(nil)

// NewBACnetConstructedDataMinActualValue factory function for _BACnetConstructedDataMinActualValue
func NewBACnetConstructedDataMinActualValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minActualValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinActualValue {
	if minActualValue == nil {
		panic("minActualValue of type BACnetApplicationTagReal for BACnetConstructedDataMinActualValue must not be nil")
	}
	_result := &_BACnetConstructedDataMinActualValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinActualValue:                minActualValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMinActualValueBuilder is a builder for BACnetConstructedDataMinActualValue
type BACnetConstructedDataMinActualValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minActualValue BACnetApplicationTagReal) BACnetConstructedDataMinActualValueBuilder
	// WithMinActualValue adds MinActualValue (property field)
	WithMinActualValue(BACnetApplicationTagReal) BACnetConstructedDataMinActualValueBuilder
	// WithMinActualValueBuilder adds MinActualValue (property field) which is build by the builder
	WithMinActualValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataMinActualValueBuilder
	// Build builds the BACnetConstructedDataMinActualValue or returns an error if something is wrong
	Build() (BACnetConstructedDataMinActualValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMinActualValue
}

// NewBACnetConstructedDataMinActualValueBuilder() creates a BACnetConstructedDataMinActualValueBuilder
func NewBACnetConstructedDataMinActualValueBuilder() BACnetConstructedDataMinActualValueBuilder {
	return &_BACnetConstructedDataMinActualValueBuilder{_BACnetConstructedDataMinActualValue: new(_BACnetConstructedDataMinActualValue)}
}

type _BACnetConstructedDataMinActualValueBuilder struct {
	*_BACnetConstructedDataMinActualValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataMinActualValueBuilder) = (*_BACnetConstructedDataMinActualValueBuilder)(nil)

func (m *_BACnetConstructedDataMinActualValueBuilder) WithMandatoryFields(minActualValue BACnetApplicationTagReal) BACnetConstructedDataMinActualValueBuilder {
	return m.WithMinActualValue(minActualValue)
}

func (m *_BACnetConstructedDataMinActualValueBuilder) WithMinActualValue(minActualValue BACnetApplicationTagReal) BACnetConstructedDataMinActualValueBuilder {
	m.MinActualValue = minActualValue
	return m
}

func (m *_BACnetConstructedDataMinActualValueBuilder) WithMinActualValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataMinActualValueBuilder {
	builder := builderSupplier(m.MinActualValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.MinActualValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataMinActualValueBuilder) Build() (BACnetConstructedDataMinActualValue, error) {
	if m.MinActualValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'minActualValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataMinActualValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataMinActualValueBuilder) MustBuild() BACnetConstructedDataMinActualValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataMinActualValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataMinActualValueBuilder()
}

// CreateBACnetConstructedDataMinActualValueBuilder creates a BACnetConstructedDataMinActualValueBuilder
func (m *_BACnetConstructedDataMinActualValue) CreateBACnetConstructedDataMinActualValueBuilder() BACnetConstructedDataMinActualValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataMinActualValueBuilder()
	}
	return &_BACnetConstructedDataMinActualValueBuilder{_BACnetConstructedDataMinActualValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinActualValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinActualValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_ACTUAL_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinActualValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinActualValue) GetMinActualValue() BACnetApplicationTagReal {
	return m.MinActualValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinActualValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMinActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinActualValue(structType any) BACnetConstructedDataMinActualValue {
	if casted, ok := structType.(BACnetConstructedDataMinActualValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinActualValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinActualValue) GetTypeName() string {
	return "BACnetConstructedDataMinActualValue"
}

func (m *_BACnetConstructedDataMinActualValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minActualValue)
	lengthInBits += m.MinActualValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinActualValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMinActualValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMinActualValue BACnetConstructedDataMinActualValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinActualValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minActualValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "minActualValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minActualValue' field"))
	}
	m.MinActualValue = minActualValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), minActualValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinActualValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMinActualValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinActualValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinActualValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "minActualValue", m.GetMinActualValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minActualValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinActualValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinActualValue) IsBACnetConstructedDataMinActualValue() {}

func (m *_BACnetConstructedDataMinActualValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMinActualValue) deepCopy() *_BACnetConstructedDataMinActualValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMinActualValueCopy := &_BACnetConstructedDataMinActualValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MinActualValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMinActualValueCopy
}

func (m *_BACnetConstructedDataMinActualValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
