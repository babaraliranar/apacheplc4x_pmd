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

// BACnetConstructedDataLinkSpeed is the corresponding interface of BACnetConstructedDataLinkSpeed
type BACnetConstructedDataLinkSpeed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLinkSpeed returns LinkSpeed (property field)
	GetLinkSpeed() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataLinkSpeed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLinkSpeed()
	// CreateBuilder creates a BACnetConstructedDataLinkSpeedBuilder
	CreateBACnetConstructedDataLinkSpeedBuilder() BACnetConstructedDataLinkSpeedBuilder
}

// _BACnetConstructedDataLinkSpeed is the data-structure of this message
type _BACnetConstructedDataLinkSpeed struct {
	BACnetConstructedDataContract
	LinkSpeed BACnetApplicationTagReal
}

var _ BACnetConstructedDataLinkSpeed = (*_BACnetConstructedDataLinkSpeed)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLinkSpeed)(nil)

// NewBACnetConstructedDataLinkSpeed factory function for _BACnetConstructedDataLinkSpeed
func NewBACnetConstructedDataLinkSpeed(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, linkSpeed BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLinkSpeed {
	if linkSpeed == nil {
		panic("linkSpeed of type BACnetApplicationTagReal for BACnetConstructedDataLinkSpeed must not be nil")
	}
	_result := &_BACnetConstructedDataLinkSpeed{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LinkSpeed:                     linkSpeed,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLinkSpeedBuilder is a builder for BACnetConstructedDataLinkSpeed
type BACnetConstructedDataLinkSpeedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(linkSpeed BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedBuilder
	// WithLinkSpeed adds LinkSpeed (property field)
	WithLinkSpeed(BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedBuilder
	// WithLinkSpeedBuilder adds LinkSpeed (property field) which is build by the builder
	WithLinkSpeedBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataLinkSpeedBuilder
	// Build builds the BACnetConstructedDataLinkSpeed or returns an error if something is wrong
	Build() (BACnetConstructedDataLinkSpeed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLinkSpeed
}

// NewBACnetConstructedDataLinkSpeedBuilder() creates a BACnetConstructedDataLinkSpeedBuilder
func NewBACnetConstructedDataLinkSpeedBuilder() BACnetConstructedDataLinkSpeedBuilder {
	return &_BACnetConstructedDataLinkSpeedBuilder{_BACnetConstructedDataLinkSpeed: new(_BACnetConstructedDataLinkSpeed)}
}

type _BACnetConstructedDataLinkSpeedBuilder struct {
	*_BACnetConstructedDataLinkSpeed

	err *utils.MultiError
}

var _ (BACnetConstructedDataLinkSpeedBuilder) = (*_BACnetConstructedDataLinkSpeedBuilder)(nil)

func (m *_BACnetConstructedDataLinkSpeedBuilder) WithMandatoryFields(linkSpeed BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedBuilder {
	return m.WithLinkSpeed(linkSpeed)
}

func (m *_BACnetConstructedDataLinkSpeedBuilder) WithLinkSpeed(linkSpeed BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedBuilder {
	m.LinkSpeed = linkSpeed
	return m
}

func (m *_BACnetConstructedDataLinkSpeedBuilder) WithLinkSpeedBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataLinkSpeedBuilder {
	builder := builderSupplier(m.LinkSpeed.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.LinkSpeed, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLinkSpeedBuilder) Build() (BACnetConstructedDataLinkSpeed, error) {
	if m.LinkSpeed == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'linkSpeed' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLinkSpeed.deepCopy(), nil
}

func (m *_BACnetConstructedDataLinkSpeedBuilder) MustBuild() BACnetConstructedDataLinkSpeed {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLinkSpeedBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLinkSpeedBuilder()
}

// CreateBACnetConstructedDataLinkSpeedBuilder creates a BACnetConstructedDataLinkSpeedBuilder
func (m *_BACnetConstructedDataLinkSpeed) CreateBACnetConstructedDataLinkSpeedBuilder() BACnetConstructedDataLinkSpeedBuilder {
	if m == nil {
		return NewBACnetConstructedDataLinkSpeedBuilder()
	}
	return &_BACnetConstructedDataLinkSpeedBuilder{_BACnetConstructedDataLinkSpeed: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeed) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLinkSpeed) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LINK_SPEED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLinkSpeed) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeed) GetLinkSpeed() BACnetApplicationTagReal {
	return m.LinkSpeed
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeed) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetLinkSpeed())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLinkSpeed(structType any) BACnetConstructedDataLinkSpeed {
	if casted, ok := structType.(BACnetConstructedDataLinkSpeed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLinkSpeed) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeed"
}

func (m *_BACnetConstructedDataLinkSpeed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (linkSpeed)
	lengthInBits += m.LinkSpeed.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLinkSpeed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLinkSpeed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLinkSpeed BACnetConstructedDataLinkSpeed, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLinkSpeed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	linkSpeed, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "linkSpeed", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkSpeed' field"))
	}
	m.LinkSpeed = linkSpeed

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), linkSpeed)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLinkSpeed")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLinkSpeed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLinkSpeed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLinkSpeed")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "linkSpeed", m.GetLinkSpeed(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkSpeed' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLinkSpeed")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLinkSpeed) IsBACnetConstructedDataLinkSpeed() {}

func (m *_BACnetConstructedDataLinkSpeed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLinkSpeed) deepCopy() *_BACnetConstructedDataLinkSpeed {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLinkSpeedCopy := &_BACnetConstructedDataLinkSpeed{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LinkSpeed.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLinkSpeedCopy
}

func (m *_BACnetConstructedDataLinkSpeed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
