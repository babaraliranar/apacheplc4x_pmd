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

// BACnetConstructedDataLargeAnalogValuePresentValue is the corresponding interface of BACnetConstructedDataLargeAnalogValuePresentValue
type BACnetConstructedDataLargeAnalogValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValuePresentValueBuilder
	CreateBACnetConstructedDataLargeAnalogValuePresentValueBuilder() BACnetConstructedDataLargeAnalogValuePresentValueBuilder
}

// _BACnetConstructedDataLargeAnalogValuePresentValue is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValuePresentValue = (*_BACnetConstructedDataLargeAnalogValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValuePresentValue)(nil)

// NewBACnetConstructedDataLargeAnalogValuePresentValue factory function for _BACnetConstructedDataLargeAnalogValuePresentValue
func NewBACnetConstructedDataLargeAnalogValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValuePresentValueBuilder is a builder for BACnetConstructedDataLargeAnalogValuePresentValue
type BACnetConstructedDataLargeAnalogValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValuePresentValueBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValuePresentValue
}

// NewBACnetConstructedDataLargeAnalogValuePresentValueBuilder() creates a BACnetConstructedDataLargeAnalogValuePresentValueBuilder
func NewBACnetConstructedDataLargeAnalogValuePresentValueBuilder() BACnetConstructedDataLargeAnalogValuePresentValueBuilder {
	return &_BACnetConstructedDataLargeAnalogValuePresentValueBuilder{_BACnetConstructedDataLargeAnalogValuePresentValue: new(_BACnetConstructedDataLargeAnalogValuePresentValue)}
}

type _BACnetConstructedDataLargeAnalogValuePresentValueBuilder struct {
	*_BACnetConstructedDataLargeAnalogValuePresentValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValuePresentValueBuilder) = (*_BACnetConstructedDataLargeAnalogValuePresentValueBuilder)(nil)

func (m *_BACnetConstructedDataLargeAnalogValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValuePresentValueBuilder {
	return m.WithPresentValue(presentValue)
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValuePresentValueBuilder {
	m.PresentValue = presentValue
	return m
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValuePresentValueBuilder {
	builder := builderSupplier(m.PresentValue.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	m.PresentValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValueBuilder) Build() (BACnetConstructedDataLargeAnalogValuePresentValue, error) {
	if m.PresentValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLargeAnalogValuePresentValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValueBuilder) MustBuild() BACnetConstructedDataLargeAnalogValuePresentValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLargeAnalogValuePresentValueBuilder()
}

// CreateBACnetConstructedDataLargeAnalogValuePresentValueBuilder creates a BACnetConstructedDataLargeAnalogValuePresentValueBuilder
func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) CreateBACnetConstructedDataLargeAnalogValuePresentValueBuilder() BACnetConstructedDataLargeAnalogValuePresentValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataLargeAnalogValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValuePresentValueBuilder{_BACnetConstructedDataLargeAnalogValuePresentValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetPresentValue() BACnetApplicationTagDouble {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValuePresentValue(structType any) BACnetConstructedDataLargeAnalogValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValuePresentValue"
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValuePresentValue BACnetConstructedDataLargeAnalogValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "presentValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) IsBACnetConstructedDataLargeAnalogValuePresentValue() {
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) deepCopy() *_BACnetConstructedDataLargeAnalogValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValuePresentValueCopy := &_BACnetConstructedDataLargeAnalogValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValuePresentValueCopy
}

func (m *_BACnetConstructedDataLargeAnalogValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
