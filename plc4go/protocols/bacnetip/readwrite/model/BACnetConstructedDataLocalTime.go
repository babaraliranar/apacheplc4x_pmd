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

// BACnetConstructedDataLocalTime is the corresponding interface of BACnetConstructedDataLocalTime
type BACnetConstructedDataLocalTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLocalTime returns LocalTime (property field)
	GetLocalTime() BACnetApplicationTagTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagTime
	// IsBACnetConstructedDataLocalTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLocalTime()
	// CreateBuilder creates a BACnetConstructedDataLocalTimeBuilder
	CreateBACnetConstructedDataLocalTimeBuilder() BACnetConstructedDataLocalTimeBuilder
}

// _BACnetConstructedDataLocalTime is the data-structure of this message
type _BACnetConstructedDataLocalTime struct {
	BACnetConstructedDataContract
	LocalTime BACnetApplicationTagTime
}

var _ BACnetConstructedDataLocalTime = (*_BACnetConstructedDataLocalTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLocalTime)(nil)

// NewBACnetConstructedDataLocalTime factory function for _BACnetConstructedDataLocalTime
func NewBACnetConstructedDataLocalTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, localTime BACnetApplicationTagTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLocalTime {
	if localTime == nil {
		panic("localTime of type BACnetApplicationTagTime for BACnetConstructedDataLocalTime must not be nil")
	}
	_result := &_BACnetConstructedDataLocalTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LocalTime:                     localTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLocalTimeBuilder is a builder for BACnetConstructedDataLocalTime
type BACnetConstructedDataLocalTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(localTime BACnetApplicationTagTime) BACnetConstructedDataLocalTimeBuilder
	// WithLocalTime adds LocalTime (property field)
	WithLocalTime(BACnetApplicationTagTime) BACnetConstructedDataLocalTimeBuilder
	// WithLocalTimeBuilder adds LocalTime (property field) which is build by the builder
	WithLocalTimeBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataLocalTimeBuilder
	// Build builds the BACnetConstructedDataLocalTime or returns an error if something is wrong
	Build() (BACnetConstructedDataLocalTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLocalTime
}

// NewBACnetConstructedDataLocalTimeBuilder() creates a BACnetConstructedDataLocalTimeBuilder
func NewBACnetConstructedDataLocalTimeBuilder() BACnetConstructedDataLocalTimeBuilder {
	return &_BACnetConstructedDataLocalTimeBuilder{_BACnetConstructedDataLocalTime: new(_BACnetConstructedDataLocalTime)}
}

type _BACnetConstructedDataLocalTimeBuilder struct {
	*_BACnetConstructedDataLocalTime

	err *utils.MultiError
}

var _ (BACnetConstructedDataLocalTimeBuilder) = (*_BACnetConstructedDataLocalTimeBuilder)(nil)

func (m *_BACnetConstructedDataLocalTimeBuilder) WithMandatoryFields(localTime BACnetApplicationTagTime) BACnetConstructedDataLocalTimeBuilder {
	return m.WithLocalTime(localTime)
}

func (m *_BACnetConstructedDataLocalTimeBuilder) WithLocalTime(localTime BACnetApplicationTagTime) BACnetConstructedDataLocalTimeBuilder {
	m.LocalTime = localTime
	return m
}

func (m *_BACnetConstructedDataLocalTimeBuilder) WithLocalTimeBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataLocalTimeBuilder {
	builder := builderSupplier(m.LocalTime.CreateBACnetApplicationTagTimeBuilder())
	var err error
	m.LocalTime, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLocalTimeBuilder) Build() (BACnetConstructedDataLocalTime, error) {
	if m.LocalTime == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'localTime' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLocalTime.deepCopy(), nil
}

func (m *_BACnetConstructedDataLocalTimeBuilder) MustBuild() BACnetConstructedDataLocalTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLocalTimeBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLocalTimeBuilder()
}

// CreateBACnetConstructedDataLocalTimeBuilder creates a BACnetConstructedDataLocalTimeBuilder
func (m *_BACnetConstructedDataLocalTime) CreateBACnetConstructedDataLocalTimeBuilder() BACnetConstructedDataLocalTimeBuilder {
	if m == nil {
		return NewBACnetConstructedDataLocalTimeBuilder()
	}
	return &_BACnetConstructedDataLocalTimeBuilder{_BACnetConstructedDataLocalTime: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLocalTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLocalTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCAL_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLocalTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLocalTime) GetLocalTime() BACnetApplicationTagTime {
	return m.LocalTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLocalTime) GetActualValue() BACnetApplicationTagTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagTime(m.GetLocalTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLocalTime(structType any) BACnetConstructedDataLocalTime {
	if casted, ok := structType.(BACnetConstructedDataLocalTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLocalTime) GetTypeName() string {
	return "BACnetConstructedDataLocalTime"
}

func (m *_BACnetConstructedDataLocalTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (localTime)
	lengthInBits += m.LocalTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLocalTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLocalTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLocalTime BACnetConstructedDataLocalTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	localTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "localTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localTime' field"))
	}
	m.LocalTime = localTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagTime](ctx, "actualValue", (*BACnetApplicationTagTime)(nil), localTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLocalTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLocalTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "localTime", m.GetLocalTime(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'localTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLocalTime) IsBACnetConstructedDataLocalTime() {}

func (m *_BACnetConstructedDataLocalTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLocalTime) deepCopy() *_BACnetConstructedDataLocalTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLocalTimeCopy := &_BACnetConstructedDataLocalTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LocalTime.DeepCopy().(BACnetApplicationTagTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLocalTimeCopy
}

func (m *_BACnetConstructedDataLocalTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
