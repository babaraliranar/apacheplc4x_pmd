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

// BACnetConstructedDataValueSource is the corresponding interface of BACnetConstructedDataValueSource
type BACnetConstructedDataValueSource interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetValueSource returns ValueSource (property field)
	GetValueSource() BACnetValueSource
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetValueSource
	// IsBACnetConstructedDataValueSource is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataValueSource()
	// CreateBuilder creates a BACnetConstructedDataValueSourceBuilder
	CreateBACnetConstructedDataValueSourceBuilder() BACnetConstructedDataValueSourceBuilder
}

// _BACnetConstructedDataValueSource is the data-structure of this message
type _BACnetConstructedDataValueSource struct {
	BACnetConstructedDataContract
	ValueSource BACnetValueSource
}

var _ BACnetConstructedDataValueSource = (*_BACnetConstructedDataValueSource)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataValueSource)(nil)

// NewBACnetConstructedDataValueSource factory function for _BACnetConstructedDataValueSource
func NewBACnetConstructedDataValueSource(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, valueSource BACnetValueSource, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValueSource {
	if valueSource == nil {
		panic("valueSource of type BACnetValueSource for BACnetConstructedDataValueSource must not be nil")
	}
	_result := &_BACnetConstructedDataValueSource{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ValueSource:                   valueSource,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataValueSourceBuilder is a builder for BACnetConstructedDataValueSource
type BACnetConstructedDataValueSourceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(valueSource BACnetValueSource) BACnetConstructedDataValueSourceBuilder
	// WithValueSource adds ValueSource (property field)
	WithValueSource(BACnetValueSource) BACnetConstructedDataValueSourceBuilder
	// Build builds the BACnetConstructedDataValueSource or returns an error if something is wrong
	Build() (BACnetConstructedDataValueSource, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataValueSource
}

// NewBACnetConstructedDataValueSourceBuilder() creates a BACnetConstructedDataValueSourceBuilder
func NewBACnetConstructedDataValueSourceBuilder() BACnetConstructedDataValueSourceBuilder {
	return &_BACnetConstructedDataValueSourceBuilder{_BACnetConstructedDataValueSource: new(_BACnetConstructedDataValueSource)}
}

type _BACnetConstructedDataValueSourceBuilder struct {
	*_BACnetConstructedDataValueSource

	err *utils.MultiError
}

var _ (BACnetConstructedDataValueSourceBuilder) = (*_BACnetConstructedDataValueSourceBuilder)(nil)

func (m *_BACnetConstructedDataValueSourceBuilder) WithMandatoryFields(valueSource BACnetValueSource) BACnetConstructedDataValueSourceBuilder {
	return m.WithValueSource(valueSource)
}

func (m *_BACnetConstructedDataValueSourceBuilder) WithValueSource(valueSource BACnetValueSource) BACnetConstructedDataValueSourceBuilder {
	m.ValueSource = valueSource
	return m
}

func (m *_BACnetConstructedDataValueSourceBuilder) Build() (BACnetConstructedDataValueSource, error) {
	if m.ValueSource == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'valueSource' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataValueSource.deepCopy(), nil
}

func (m *_BACnetConstructedDataValueSourceBuilder) MustBuild() BACnetConstructedDataValueSource {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataValueSourceBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataValueSourceBuilder()
}

// CreateBACnetConstructedDataValueSourceBuilder creates a BACnetConstructedDataValueSourceBuilder
func (m *_BACnetConstructedDataValueSource) CreateBACnetConstructedDataValueSourceBuilder() BACnetConstructedDataValueSourceBuilder {
	if m == nil {
		return NewBACnetConstructedDataValueSourceBuilder()
	}
	return &_BACnetConstructedDataValueSourceBuilder{_BACnetConstructedDataValueSource: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValueSource) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValueSource) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_SOURCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValueSource) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValueSource) GetValueSource() BACnetValueSource {
	return m.ValueSource
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValueSource) GetActualValue() BACnetValueSource {
	ctx := context.Background()
	_ = ctx
	return CastBACnetValueSource(m.GetValueSource())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValueSource(structType any) BACnetConstructedDataValueSource {
	if casted, ok := structType.(BACnetConstructedDataValueSource); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueSource); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValueSource) GetTypeName() string {
	return "BACnetConstructedDataValueSource"
}

func (m *_BACnetConstructedDataValueSource) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (valueSource)
	lengthInBits += m.ValueSource.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValueSource) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataValueSource) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataValueSource BACnetConstructedDataValueSource, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValueSource")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	valueSource, err := ReadSimpleField[BACnetValueSource](ctx, "valueSource", ReadComplex[BACnetValueSource](BACnetValueSourceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueSource' field"))
	}
	m.ValueSource = valueSource

	actualValue, err := ReadVirtualField[BACnetValueSource](ctx, "actualValue", (*BACnetValueSource)(nil), valueSource)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValueSource")
	}

	return m, nil
}

func (m *_BACnetConstructedDataValueSource) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataValueSource) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueSource"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValueSource")
		}

		if err := WriteSimpleField[BACnetValueSource](ctx, "valueSource", m.GetValueSource(), WriteComplex[BACnetValueSource](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'valueSource' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueSource"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValueSource")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValueSource) IsBACnetConstructedDataValueSource() {}

func (m *_BACnetConstructedDataValueSource) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataValueSource) deepCopy() *_BACnetConstructedDataValueSource {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataValueSourceCopy := &_BACnetConstructedDataValueSource{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ValueSource.DeepCopy().(BACnetValueSource),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataValueSourceCopy
}

func (m *_BACnetConstructedDataValueSource) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
