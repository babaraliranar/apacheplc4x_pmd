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

// BACnetConstructedDataLargeAnalogValueRelinquishDefault is the corresponding interface of BACnetConstructedDataLargeAnalogValueRelinquishDefault
type BACnetConstructedDataLargeAnalogValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
	CreateBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder() BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
}

// _BACnetConstructedDataLargeAnalogValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueRelinquishDefault = (*_BACnetConstructedDataLargeAnalogValueRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueRelinquishDefault)(nil)

// NewBACnetConstructedDataLargeAnalogValueRelinquishDefault factory function for _BACnetConstructedDataLargeAnalogValueRelinquishDefault
func NewBACnetConstructedDataLargeAnalogValueRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder is a builder for BACnetConstructedDataLargeAnalogValueRelinquishDefault
type BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueRelinquishDefault
}

// NewBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder() creates a BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
func NewBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder() BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder {
	return &_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder{_BACnetConstructedDataLargeAnalogValueRelinquishDefault: new(_BACnetConstructedDataLargeAnalogValueRelinquishDefault)}
}

type _BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueRelinquishDefault

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) = (*_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder)(nil)

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder {
	return m.WithRelinquishDefault(relinquishDefault)
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder {
	m.RelinquishDefault = relinquishDefault
	return m
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder {
	builder := builderSupplier(m.RelinquishDefault.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	m.RelinquishDefault, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) Build() (BACnetConstructedDataLargeAnalogValueRelinquishDefault, error) {
	if m.RelinquishDefault == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLargeAnalogValueRelinquishDefault.deepCopy(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueRelinquishDefault {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder()
}

// CreateBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder creates a BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder
func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) CreateBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder() BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder {
	if m == nil {
		return NewBACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueRelinquishDefaultBuilder{_BACnetConstructedDataLargeAnalogValueRelinquishDefault: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagDouble {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueRelinquishDefault(structType any) BACnetConstructedDataLargeAnalogValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueRelinquishDefault"
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueRelinquishDefault BACnetConstructedDataLargeAnalogValueRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "relinquishDefault", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) IsBACnetConstructedDataLargeAnalogValueRelinquishDefault() {
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) deepCopy() *_BACnetConstructedDataLargeAnalogValueRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueRelinquishDefaultCopy := &_BACnetConstructedDataLargeAnalogValueRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.RelinquishDefault.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
