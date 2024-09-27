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

// BACnetConstructedDataOccupancyCount is the corresponding interface of BACnetConstructedDataOccupancyCount
type BACnetConstructedDataOccupancyCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOccupancyCount returns OccupancyCount (property field)
	GetOccupancyCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataOccupancyCount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyCount()
	// CreateBuilder creates a BACnetConstructedDataOccupancyCountBuilder
	CreateBACnetConstructedDataOccupancyCountBuilder() BACnetConstructedDataOccupancyCountBuilder
}

// _BACnetConstructedDataOccupancyCount is the data-structure of this message
type _BACnetConstructedDataOccupancyCount struct {
	BACnetConstructedDataContract
	OccupancyCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataOccupancyCount = (*_BACnetConstructedDataOccupancyCount)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyCount)(nil)

// NewBACnetConstructedDataOccupancyCount factory function for _BACnetConstructedDataOccupancyCount
func NewBACnetConstructedDataOccupancyCount(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyCount BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyCount {
	if occupancyCount == nil {
		panic("occupancyCount of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataOccupancyCount must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyCount{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyCount:                occupancyCount,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOccupancyCountBuilder is a builder for BACnetConstructedDataOccupancyCount
type BACnetConstructedDataOccupancyCountBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(occupancyCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyCountBuilder
	// WithOccupancyCount adds OccupancyCount (property field)
	WithOccupancyCount(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyCountBuilder
	// WithOccupancyCountBuilder adds OccupancyCount (property field) which is build by the builder
	WithOccupancyCountBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataOccupancyCountBuilder
	// Build builds the BACnetConstructedDataOccupancyCount or returns an error if something is wrong
	Build() (BACnetConstructedDataOccupancyCount, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOccupancyCount
}

// NewBACnetConstructedDataOccupancyCountBuilder() creates a BACnetConstructedDataOccupancyCountBuilder
func NewBACnetConstructedDataOccupancyCountBuilder() BACnetConstructedDataOccupancyCountBuilder {
	return &_BACnetConstructedDataOccupancyCountBuilder{_BACnetConstructedDataOccupancyCount: new(_BACnetConstructedDataOccupancyCount)}
}

type _BACnetConstructedDataOccupancyCountBuilder struct {
	*_BACnetConstructedDataOccupancyCount

	err *utils.MultiError
}

var _ (BACnetConstructedDataOccupancyCountBuilder) = (*_BACnetConstructedDataOccupancyCountBuilder)(nil)

func (m *_BACnetConstructedDataOccupancyCountBuilder) WithMandatoryFields(occupancyCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyCountBuilder {
	return m.WithOccupancyCount(occupancyCount)
}

func (m *_BACnetConstructedDataOccupancyCountBuilder) WithOccupancyCount(occupancyCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyCountBuilder {
	m.OccupancyCount = occupancyCount
	return m
}

func (m *_BACnetConstructedDataOccupancyCountBuilder) WithOccupancyCountBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataOccupancyCountBuilder {
	builder := builderSupplier(m.OccupancyCount.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.OccupancyCount, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataOccupancyCountBuilder) Build() (BACnetConstructedDataOccupancyCount, error) {
	if m.OccupancyCount == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'occupancyCount' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataOccupancyCount.deepCopy(), nil
}

func (m *_BACnetConstructedDataOccupancyCountBuilder) MustBuild() BACnetConstructedDataOccupancyCount {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataOccupancyCountBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataOccupancyCountBuilder()
}

// CreateBACnetConstructedDataOccupancyCountBuilder creates a BACnetConstructedDataOccupancyCountBuilder
func (m *_BACnetConstructedDataOccupancyCount) CreateBACnetConstructedDataOccupancyCountBuilder() BACnetConstructedDataOccupancyCountBuilder {
	if m == nil {
		return NewBACnetConstructedDataOccupancyCountBuilder()
	}
	return &_BACnetConstructedDataOccupancyCountBuilder{_BACnetConstructedDataOccupancyCount: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyCount) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCount) GetOccupancyCount() BACnetApplicationTagUnsignedInteger {
	return m.OccupancyCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetOccupancyCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyCount(structType any) BACnetConstructedDataOccupancyCount {
	if casted, ok := structType.(BACnetConstructedDataOccupancyCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyCount) GetTypeName() string {
	return "BACnetConstructedDataOccupancyCount"
}

func (m *_BACnetConstructedDataOccupancyCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyCount)
	lengthInBits += m.OccupancyCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyCount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyCount BACnetConstructedDataOccupancyCount, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "occupancyCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyCount' field"))
	}
	m.OccupancyCount = occupancyCount

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), occupancyCount)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyCount")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyCount")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "occupancyCount", m.GetOccupancyCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyCount' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyCount")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyCount) IsBACnetConstructedDataOccupancyCount() {}

func (m *_BACnetConstructedDataOccupancyCount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOccupancyCount) deepCopy() *_BACnetConstructedDataOccupancyCount {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOccupancyCountCopy := &_BACnetConstructedDataOccupancyCount{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.OccupancyCount.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOccupancyCountCopy
}

func (m *_BACnetConstructedDataOccupancyCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
