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

// BACnetConstructedDataValidSamples is the corresponding interface of BACnetConstructedDataValidSamples
type BACnetConstructedDataValidSamples interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetValidSamples returns ValidSamples (property field)
	GetValidSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataValidSamples is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataValidSamples()
	// CreateBuilder creates a BACnetConstructedDataValidSamplesBuilder
	CreateBACnetConstructedDataValidSamplesBuilder() BACnetConstructedDataValidSamplesBuilder
}

// _BACnetConstructedDataValidSamples is the data-structure of this message
type _BACnetConstructedDataValidSamples struct {
	BACnetConstructedDataContract
	ValidSamples BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataValidSamples = (*_BACnetConstructedDataValidSamples)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataValidSamples)(nil)

// NewBACnetConstructedDataValidSamples factory function for _BACnetConstructedDataValidSamples
func NewBACnetConstructedDataValidSamples(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, validSamples BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValidSamples {
	if validSamples == nil {
		panic("validSamples of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataValidSamples must not be nil")
	}
	_result := &_BACnetConstructedDataValidSamples{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ValidSamples:                  validSamples,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataValidSamplesBuilder is a builder for BACnetConstructedDataValidSamples
type BACnetConstructedDataValidSamplesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(validSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValidSamplesBuilder
	// WithValidSamples adds ValidSamples (property field)
	WithValidSamples(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValidSamplesBuilder
	// WithValidSamplesBuilder adds ValidSamples (property field) which is build by the builder
	WithValidSamplesBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataValidSamplesBuilder
	// Build builds the BACnetConstructedDataValidSamples or returns an error if something is wrong
	Build() (BACnetConstructedDataValidSamples, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataValidSamples
}

// NewBACnetConstructedDataValidSamplesBuilder() creates a BACnetConstructedDataValidSamplesBuilder
func NewBACnetConstructedDataValidSamplesBuilder() BACnetConstructedDataValidSamplesBuilder {
	return &_BACnetConstructedDataValidSamplesBuilder{_BACnetConstructedDataValidSamples: new(_BACnetConstructedDataValidSamples)}
}

type _BACnetConstructedDataValidSamplesBuilder struct {
	*_BACnetConstructedDataValidSamples

	err *utils.MultiError
}

var _ (BACnetConstructedDataValidSamplesBuilder) = (*_BACnetConstructedDataValidSamplesBuilder)(nil)

func (m *_BACnetConstructedDataValidSamplesBuilder) WithMandatoryFields(validSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValidSamplesBuilder {
	return m.WithValidSamples(validSamples)
}

func (m *_BACnetConstructedDataValidSamplesBuilder) WithValidSamples(validSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataValidSamplesBuilder {
	m.ValidSamples = validSamples
	return m
}

func (m *_BACnetConstructedDataValidSamplesBuilder) WithValidSamplesBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataValidSamplesBuilder {
	builder := builderSupplier(m.ValidSamples.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.ValidSamples, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataValidSamplesBuilder) Build() (BACnetConstructedDataValidSamples, error) {
	if m.ValidSamples == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'validSamples' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataValidSamples.deepCopy(), nil
}

func (m *_BACnetConstructedDataValidSamplesBuilder) MustBuild() BACnetConstructedDataValidSamples {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataValidSamplesBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataValidSamplesBuilder()
}

// CreateBACnetConstructedDataValidSamplesBuilder creates a BACnetConstructedDataValidSamplesBuilder
func (m *_BACnetConstructedDataValidSamples) CreateBACnetConstructedDataValidSamplesBuilder() BACnetConstructedDataValidSamplesBuilder {
	if m == nil {
		return NewBACnetConstructedDataValidSamplesBuilder()
	}
	return &_BACnetConstructedDataValidSamplesBuilder{_BACnetConstructedDataValidSamples: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValidSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValidSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALID_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValidSamples) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValidSamples) GetValidSamples() BACnetApplicationTagUnsignedInteger {
	return m.ValidSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValidSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetValidSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValidSamples(structType any) BACnetConstructedDataValidSamples {
	if casted, ok := structType.(BACnetConstructedDataValidSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValidSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValidSamples) GetTypeName() string {
	return "BACnetConstructedDataValidSamples"
}

func (m *_BACnetConstructedDataValidSamples) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (validSamples)
	lengthInBits += m.ValidSamples.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValidSamples) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataValidSamples) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataValidSamples BACnetConstructedDataValidSamples, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValidSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValidSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	validSamples, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "validSamples", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'validSamples' field"))
	}
	m.ValidSamples = validSamples

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), validSamples)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValidSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValidSamples")
	}

	return m, nil
}

func (m *_BACnetConstructedDataValidSamples) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataValidSamples) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValidSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValidSamples")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "validSamples", m.GetValidSamples(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'validSamples' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValidSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValidSamples")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValidSamples) IsBACnetConstructedDataValidSamples() {}

func (m *_BACnetConstructedDataValidSamples) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataValidSamples) deepCopy() *_BACnetConstructedDataValidSamples {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataValidSamplesCopy := &_BACnetConstructedDataValidSamples{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ValidSamples.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataValidSamplesCopy
}

func (m *_BACnetConstructedDataValidSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
