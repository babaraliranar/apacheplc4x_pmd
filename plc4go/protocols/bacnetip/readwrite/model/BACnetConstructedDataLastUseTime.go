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

// BACnetConstructedDataLastUseTime is the corresponding interface of BACnetConstructedDataLastUseTime
type BACnetConstructedDataLastUseTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastUseTime returns LastUseTime (property field)
	GetLastUseTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataLastUseTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastUseTime()
	// CreateBuilder creates a BACnetConstructedDataLastUseTimeBuilder
	CreateBACnetConstructedDataLastUseTimeBuilder() BACnetConstructedDataLastUseTimeBuilder
}

// _BACnetConstructedDataLastUseTime is the data-structure of this message
type _BACnetConstructedDataLastUseTime struct {
	BACnetConstructedDataContract
	LastUseTime BACnetDateTime
}

var _ BACnetConstructedDataLastUseTime = (*_BACnetConstructedDataLastUseTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastUseTime)(nil)

// NewBACnetConstructedDataLastUseTime factory function for _BACnetConstructedDataLastUseTime
func NewBACnetConstructedDataLastUseTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastUseTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastUseTime {
	if lastUseTime == nil {
		panic("lastUseTime of type BACnetDateTime for BACnetConstructedDataLastUseTime must not be nil")
	}
	_result := &_BACnetConstructedDataLastUseTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastUseTime:                   lastUseTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastUseTimeBuilder is a builder for BACnetConstructedDataLastUseTime
type BACnetConstructedDataLastUseTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastUseTime BACnetDateTime) BACnetConstructedDataLastUseTimeBuilder
	// WithLastUseTime adds LastUseTime (property field)
	WithLastUseTime(BACnetDateTime) BACnetConstructedDataLastUseTimeBuilder
	// WithLastUseTimeBuilder adds LastUseTime (property field) which is build by the builder
	WithLastUseTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataLastUseTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastUseTime or returns an error if something is wrong
	Build() (BACnetConstructedDataLastUseTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastUseTime
}

// NewBACnetConstructedDataLastUseTimeBuilder() creates a BACnetConstructedDataLastUseTimeBuilder
func NewBACnetConstructedDataLastUseTimeBuilder() BACnetConstructedDataLastUseTimeBuilder {
	return &_BACnetConstructedDataLastUseTimeBuilder{_BACnetConstructedDataLastUseTime: new(_BACnetConstructedDataLastUseTime)}
}

type _BACnetConstructedDataLastUseTimeBuilder struct {
	*_BACnetConstructedDataLastUseTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastUseTimeBuilder) = (*_BACnetConstructedDataLastUseTimeBuilder)(nil)

func (b *_BACnetConstructedDataLastUseTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLastUseTime
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) WithMandatoryFields(lastUseTime BACnetDateTime) BACnetConstructedDataLastUseTimeBuilder {
	return b.WithLastUseTime(lastUseTime)
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) WithLastUseTime(lastUseTime BACnetDateTime) BACnetConstructedDataLastUseTimeBuilder {
	b.LastUseTime = lastUseTime
	return b
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) WithLastUseTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataLastUseTimeBuilder {
	builder := builderSupplier(b.LastUseTime.CreateBACnetDateTimeBuilder())
	var err error
	b.LastUseTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) Build() (BACnetConstructedDataLastUseTime, error) {
	if b.LastUseTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastUseTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastUseTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) MustBuild() BACnetConstructedDataLastUseTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastUseTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastUseTimeBuilder().(*_BACnetConstructedDataLastUseTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastUseTimeBuilder creates a BACnetConstructedDataLastUseTimeBuilder
func (b *_BACnetConstructedDataLastUseTime) CreateBACnetConstructedDataLastUseTimeBuilder() BACnetConstructedDataLastUseTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastUseTimeBuilder()
	}
	return &_BACnetConstructedDataLastUseTimeBuilder{_BACnetConstructedDataLastUseTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastUseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastUseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_USE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastUseTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastUseTime) GetLastUseTime() BACnetDateTime {
	return m.LastUseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastUseTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetLastUseTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastUseTime(structType any) BACnetConstructedDataLastUseTime {
	if casted, ok := structType.(BACnetConstructedDataLastUseTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastUseTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastUseTime) GetTypeName() string {
	return "BACnetConstructedDataLastUseTime"
}

func (m *_BACnetConstructedDataLastUseTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastUseTime)
	lengthInBits += m.LastUseTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastUseTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastUseTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastUseTime BACnetConstructedDataLastUseTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastUseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastUseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastUseTime, err := ReadSimpleField[BACnetDateTime](ctx, "lastUseTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastUseTime' field"))
	}
	m.LastUseTime = lastUseTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), lastUseTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastUseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastUseTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastUseTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastUseTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastUseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastUseTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "lastUseTime", m.GetLastUseTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastUseTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastUseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastUseTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastUseTime) IsBACnetConstructedDataLastUseTime() {}

func (m *_BACnetConstructedDataLastUseTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastUseTime) deepCopy() *_BACnetConstructedDataLastUseTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastUseTimeCopy := &_BACnetConstructedDataLastUseTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.LastUseTime),
	}
	_BACnetConstructedDataLastUseTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastUseTimeCopy
}

func (m *_BACnetConstructedDataLastUseTime) String() string {
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
