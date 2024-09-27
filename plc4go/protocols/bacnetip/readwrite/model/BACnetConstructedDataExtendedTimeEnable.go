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

// BACnetConstructedDataExtendedTimeEnable is the corresponding interface of BACnetConstructedDataExtendedTimeEnable
type BACnetConstructedDataExtendedTimeEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetExtendedTimeEnable returns ExtendedTimeEnable (property field)
	GetExtendedTimeEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataExtendedTimeEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataExtendedTimeEnable()
	// CreateBuilder creates a BACnetConstructedDataExtendedTimeEnableBuilder
	CreateBACnetConstructedDataExtendedTimeEnableBuilder() BACnetConstructedDataExtendedTimeEnableBuilder
}

// _BACnetConstructedDataExtendedTimeEnable is the data-structure of this message
type _BACnetConstructedDataExtendedTimeEnable struct {
	BACnetConstructedDataContract
	ExtendedTimeEnable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataExtendedTimeEnable = (*_BACnetConstructedDataExtendedTimeEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataExtendedTimeEnable)(nil)

// NewBACnetConstructedDataExtendedTimeEnable factory function for _BACnetConstructedDataExtendedTimeEnable
func NewBACnetConstructedDataExtendedTimeEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, extendedTimeEnable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExtendedTimeEnable {
	if extendedTimeEnable == nil {
		panic("extendedTimeEnable of type BACnetApplicationTagBoolean for BACnetConstructedDataExtendedTimeEnable must not be nil")
	}
	_result := &_BACnetConstructedDataExtendedTimeEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ExtendedTimeEnable:            extendedTimeEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataExtendedTimeEnableBuilder is a builder for BACnetConstructedDataExtendedTimeEnable
type BACnetConstructedDataExtendedTimeEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(extendedTimeEnable BACnetApplicationTagBoolean) BACnetConstructedDataExtendedTimeEnableBuilder
	// WithExtendedTimeEnable adds ExtendedTimeEnable (property field)
	WithExtendedTimeEnable(BACnetApplicationTagBoolean) BACnetConstructedDataExtendedTimeEnableBuilder
	// WithExtendedTimeEnableBuilder adds ExtendedTimeEnable (property field) which is build by the builder
	WithExtendedTimeEnableBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataExtendedTimeEnableBuilder
	// Build builds the BACnetConstructedDataExtendedTimeEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataExtendedTimeEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataExtendedTimeEnable
}

// NewBACnetConstructedDataExtendedTimeEnableBuilder() creates a BACnetConstructedDataExtendedTimeEnableBuilder
func NewBACnetConstructedDataExtendedTimeEnableBuilder() BACnetConstructedDataExtendedTimeEnableBuilder {
	return &_BACnetConstructedDataExtendedTimeEnableBuilder{_BACnetConstructedDataExtendedTimeEnable: new(_BACnetConstructedDataExtendedTimeEnable)}
}

type _BACnetConstructedDataExtendedTimeEnableBuilder struct {
	*_BACnetConstructedDataExtendedTimeEnable

	err *utils.MultiError
}

var _ (BACnetConstructedDataExtendedTimeEnableBuilder) = (*_BACnetConstructedDataExtendedTimeEnableBuilder)(nil)

func (m *_BACnetConstructedDataExtendedTimeEnableBuilder) WithMandatoryFields(extendedTimeEnable BACnetApplicationTagBoolean) BACnetConstructedDataExtendedTimeEnableBuilder {
	return m.WithExtendedTimeEnable(extendedTimeEnable)
}

func (m *_BACnetConstructedDataExtendedTimeEnableBuilder) WithExtendedTimeEnable(extendedTimeEnable BACnetApplicationTagBoolean) BACnetConstructedDataExtendedTimeEnableBuilder {
	m.ExtendedTimeEnable = extendedTimeEnable
	return m
}

func (m *_BACnetConstructedDataExtendedTimeEnableBuilder) WithExtendedTimeEnableBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataExtendedTimeEnableBuilder {
	builder := builderSupplier(m.ExtendedTimeEnable.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.ExtendedTimeEnable, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataExtendedTimeEnableBuilder) Build() (BACnetConstructedDataExtendedTimeEnable, error) {
	if m.ExtendedTimeEnable == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'extendedTimeEnable' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataExtendedTimeEnable.deepCopy(), nil
}

func (m *_BACnetConstructedDataExtendedTimeEnableBuilder) MustBuild() BACnetConstructedDataExtendedTimeEnable {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataExtendedTimeEnableBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataExtendedTimeEnableBuilder()
}

// CreateBACnetConstructedDataExtendedTimeEnableBuilder creates a BACnetConstructedDataExtendedTimeEnableBuilder
func (m *_BACnetConstructedDataExtendedTimeEnable) CreateBACnetConstructedDataExtendedTimeEnableBuilder() BACnetConstructedDataExtendedTimeEnableBuilder {
	if m == nil {
		return NewBACnetConstructedDataExtendedTimeEnableBuilder()
	}
	return &_BACnetConstructedDataExtendedTimeEnableBuilder{_BACnetConstructedDataExtendedTimeEnable: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXTENDED_TIME_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetExtendedTimeEnable() BACnetApplicationTagBoolean {
	return m.ExtendedTimeEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExtendedTimeEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetExtendedTimeEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExtendedTimeEnable(structType any) BACnetConstructedDataExtendedTimeEnable {
	if casted, ok := structType.(BACnetConstructedDataExtendedTimeEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExtendedTimeEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetTypeName() string {
	return "BACnetConstructedDataExtendedTimeEnable"
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (extendedTimeEnable)
	lengthInBits += m.ExtendedTimeEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataExtendedTimeEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataExtendedTimeEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataExtendedTimeEnable BACnetConstructedDataExtendedTimeEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExtendedTimeEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExtendedTimeEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	extendedTimeEnable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "extendedTimeEnable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedTimeEnable' field"))
	}
	m.ExtendedTimeEnable = extendedTimeEnable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), extendedTimeEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExtendedTimeEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExtendedTimeEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataExtendedTimeEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataExtendedTimeEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExtendedTimeEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExtendedTimeEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "extendedTimeEnable", m.GetExtendedTimeEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedTimeEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExtendedTimeEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExtendedTimeEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExtendedTimeEnable) IsBACnetConstructedDataExtendedTimeEnable() {}

func (m *_BACnetConstructedDataExtendedTimeEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataExtendedTimeEnable) deepCopy() *_BACnetConstructedDataExtendedTimeEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataExtendedTimeEnableCopy := &_BACnetConstructedDataExtendedTimeEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ExtendedTimeEnable.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataExtendedTimeEnableCopy
}

func (m *_BACnetConstructedDataExtendedTimeEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
