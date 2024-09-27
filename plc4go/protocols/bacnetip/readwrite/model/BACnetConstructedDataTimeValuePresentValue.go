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

// BACnetConstructedDataTimeValuePresentValue is the corresponding interface of BACnetConstructedDataTimeValuePresentValue
type BACnetConstructedDataTimeValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagTime
	// IsBACnetConstructedDataTimeValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataTimeValuePresentValueBuilder
	CreateBACnetConstructedDataTimeValuePresentValueBuilder() BACnetConstructedDataTimeValuePresentValueBuilder
}

// _BACnetConstructedDataTimeValuePresentValue is the data-structure of this message
type _BACnetConstructedDataTimeValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagTime
}

var _ BACnetConstructedDataTimeValuePresentValue = (*_BACnetConstructedDataTimeValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeValuePresentValue)(nil)

// NewBACnetConstructedDataTimeValuePresentValue factory function for _BACnetConstructedDataTimeValuePresentValue
func NewBACnetConstructedDataTimeValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagTime for BACnetConstructedDataTimeValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataTimeValuePresentValue{
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

// BACnetConstructedDataTimeValuePresentValueBuilder is a builder for BACnetConstructedDataTimeValuePresentValue
type BACnetConstructedDataTimeValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagTime) BACnetConstructedDataTimeValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagTime) BACnetConstructedDataTimeValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataTimeValuePresentValueBuilder
	// Build builds the BACnetConstructedDataTimeValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataTimeValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimeValuePresentValue
}

// NewBACnetConstructedDataTimeValuePresentValueBuilder() creates a BACnetConstructedDataTimeValuePresentValueBuilder
func NewBACnetConstructedDataTimeValuePresentValueBuilder() BACnetConstructedDataTimeValuePresentValueBuilder {
	return &_BACnetConstructedDataTimeValuePresentValueBuilder{_BACnetConstructedDataTimeValuePresentValue: new(_BACnetConstructedDataTimeValuePresentValue)}
}

type _BACnetConstructedDataTimeValuePresentValueBuilder struct {
	*_BACnetConstructedDataTimeValuePresentValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimeValuePresentValueBuilder) = (*_BACnetConstructedDataTimeValuePresentValueBuilder)(nil)

func (m *_BACnetConstructedDataTimeValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagTime) BACnetConstructedDataTimeValuePresentValueBuilder {
	return m.WithPresentValue(presentValue)
}

func (m *_BACnetConstructedDataTimeValuePresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagTime) BACnetConstructedDataTimeValuePresentValueBuilder {
	m.PresentValue = presentValue
	return m
}

func (m *_BACnetConstructedDataTimeValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataTimeValuePresentValueBuilder {
	builder := builderSupplier(m.PresentValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	m.PresentValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataTimeValuePresentValueBuilder) Build() (BACnetConstructedDataTimeValuePresentValue, error) {
	if m.PresentValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataTimeValuePresentValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataTimeValuePresentValueBuilder) MustBuild() BACnetConstructedDataTimeValuePresentValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataTimeValuePresentValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataTimeValuePresentValueBuilder()
}

// CreateBACnetConstructedDataTimeValuePresentValueBuilder creates a BACnetConstructedDataTimeValuePresentValueBuilder
func (m *_BACnetConstructedDataTimeValuePresentValue) CreateBACnetConstructedDataTimeValuePresentValueBuilder() BACnetConstructedDataTimeValuePresentValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataTimeValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataTimeValuePresentValueBuilder{_BACnetConstructedDataTimeValuePresentValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIME_VALUE
}

func (m *_BACnetConstructedDataTimeValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeValuePresentValue) GetPresentValue() BACnetApplicationTagTime {
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

func (m *_BACnetConstructedDataTimeValuePresentValue) GetActualValue() BACnetApplicationTagTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagTime(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeValuePresentValue(structType any) BACnetConstructedDataTimeValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataTimeValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataTimeValuePresentValue"
}

func (m *_BACnetConstructedDataTimeValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeValuePresentValue BACnetConstructedDataTimeValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "presentValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagTime](ctx, "actualValue", (*BACnetApplicationTagTime)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeValuePresentValue) IsBACnetConstructedDataTimeValuePresentValue() {
}

func (m *_BACnetConstructedDataTimeValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeValuePresentValue) deepCopy() *_BACnetConstructedDataTimeValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeValuePresentValueCopy := &_BACnetConstructedDataTimeValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetApplicationTagTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeValuePresentValueCopy
}

func (m *_BACnetConstructedDataTimeValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
