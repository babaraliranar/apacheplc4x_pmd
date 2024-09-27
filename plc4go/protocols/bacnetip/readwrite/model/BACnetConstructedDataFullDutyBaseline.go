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

// BACnetConstructedDataFullDutyBaseline is the corresponding interface of BACnetConstructedDataFullDutyBaseline
type BACnetConstructedDataFullDutyBaseline interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFullDutyBaseLine returns FullDutyBaseLine (property field)
	GetFullDutyBaseLine() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataFullDutyBaseline is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFullDutyBaseline()
	// CreateBuilder creates a BACnetConstructedDataFullDutyBaselineBuilder
	CreateBACnetConstructedDataFullDutyBaselineBuilder() BACnetConstructedDataFullDutyBaselineBuilder
}

// _BACnetConstructedDataFullDutyBaseline is the data-structure of this message
type _BACnetConstructedDataFullDutyBaseline struct {
	BACnetConstructedDataContract
	FullDutyBaseLine BACnetApplicationTagReal
}

var _ BACnetConstructedDataFullDutyBaseline = (*_BACnetConstructedDataFullDutyBaseline)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFullDutyBaseline)(nil)

// NewBACnetConstructedDataFullDutyBaseline factory function for _BACnetConstructedDataFullDutyBaseline
func NewBACnetConstructedDataFullDutyBaseline(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, fullDutyBaseLine BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFullDutyBaseline {
	if fullDutyBaseLine == nil {
		panic("fullDutyBaseLine of type BACnetApplicationTagReal for BACnetConstructedDataFullDutyBaseline must not be nil")
	}
	_result := &_BACnetConstructedDataFullDutyBaseline{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FullDutyBaseLine:              fullDutyBaseLine,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFullDutyBaselineBuilder is a builder for BACnetConstructedDataFullDutyBaseline
type BACnetConstructedDataFullDutyBaselineBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fullDutyBaseLine BACnetApplicationTagReal) BACnetConstructedDataFullDutyBaselineBuilder
	// WithFullDutyBaseLine adds FullDutyBaseLine (property field)
	WithFullDutyBaseLine(BACnetApplicationTagReal) BACnetConstructedDataFullDutyBaselineBuilder
	// WithFullDutyBaseLineBuilder adds FullDutyBaseLine (property field) which is build by the builder
	WithFullDutyBaseLineBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataFullDutyBaselineBuilder
	// Build builds the BACnetConstructedDataFullDutyBaseline or returns an error if something is wrong
	Build() (BACnetConstructedDataFullDutyBaseline, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFullDutyBaseline
}

// NewBACnetConstructedDataFullDutyBaselineBuilder() creates a BACnetConstructedDataFullDutyBaselineBuilder
func NewBACnetConstructedDataFullDutyBaselineBuilder() BACnetConstructedDataFullDutyBaselineBuilder {
	return &_BACnetConstructedDataFullDutyBaselineBuilder{_BACnetConstructedDataFullDutyBaseline: new(_BACnetConstructedDataFullDutyBaseline)}
}

type _BACnetConstructedDataFullDutyBaselineBuilder struct {
	*_BACnetConstructedDataFullDutyBaseline

	err *utils.MultiError
}

var _ (BACnetConstructedDataFullDutyBaselineBuilder) = (*_BACnetConstructedDataFullDutyBaselineBuilder)(nil)

func (m *_BACnetConstructedDataFullDutyBaselineBuilder) WithMandatoryFields(fullDutyBaseLine BACnetApplicationTagReal) BACnetConstructedDataFullDutyBaselineBuilder {
	return m.WithFullDutyBaseLine(fullDutyBaseLine)
}

func (m *_BACnetConstructedDataFullDutyBaselineBuilder) WithFullDutyBaseLine(fullDutyBaseLine BACnetApplicationTagReal) BACnetConstructedDataFullDutyBaselineBuilder {
	m.FullDutyBaseLine = fullDutyBaseLine
	return m
}

func (m *_BACnetConstructedDataFullDutyBaselineBuilder) WithFullDutyBaseLineBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataFullDutyBaselineBuilder {
	builder := builderSupplier(m.FullDutyBaseLine.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.FullDutyBaseLine, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataFullDutyBaselineBuilder) Build() (BACnetConstructedDataFullDutyBaseline, error) {
	if m.FullDutyBaseLine == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'fullDutyBaseLine' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataFullDutyBaseline.deepCopy(), nil
}

func (m *_BACnetConstructedDataFullDutyBaselineBuilder) MustBuild() BACnetConstructedDataFullDutyBaseline {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataFullDutyBaselineBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataFullDutyBaselineBuilder()
}

// CreateBACnetConstructedDataFullDutyBaselineBuilder creates a BACnetConstructedDataFullDutyBaselineBuilder
func (m *_BACnetConstructedDataFullDutyBaseline) CreateBACnetConstructedDataFullDutyBaselineBuilder() BACnetConstructedDataFullDutyBaselineBuilder {
	if m == nil {
		return NewBACnetConstructedDataFullDutyBaselineBuilder()
	}
	return &_BACnetConstructedDataFullDutyBaselineBuilder{_BACnetConstructedDataFullDutyBaseline: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FULL_DUTY_BASELINE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetFullDutyBaseLine() BACnetApplicationTagReal {
	return m.FullDutyBaseLine
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFullDutyBaseLine())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFullDutyBaseline(structType any) BACnetConstructedDataFullDutyBaseline {
	if casted, ok := structType.(BACnetConstructedDataFullDutyBaseline); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFullDutyBaseline); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetTypeName() string {
	return "BACnetConstructedDataFullDutyBaseline"
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (fullDutyBaseLine)
	lengthInBits += m.FullDutyBaseLine.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFullDutyBaseline) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFullDutyBaseline BACnetConstructedDataFullDutyBaseline, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFullDutyBaseline"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFullDutyBaseline")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fullDutyBaseLine, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "fullDutyBaseLine", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fullDutyBaseLine' field"))
	}
	m.FullDutyBaseLine = fullDutyBaseLine

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), fullDutyBaseLine)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFullDutyBaseline"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFullDutyBaseline")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFullDutyBaseline) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFullDutyBaseline) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFullDutyBaseline"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFullDutyBaseline")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "fullDutyBaseLine", m.GetFullDutyBaseLine(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fullDutyBaseLine' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFullDutyBaseline"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFullDutyBaseline")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFullDutyBaseline) IsBACnetConstructedDataFullDutyBaseline() {}

func (m *_BACnetConstructedDataFullDutyBaseline) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFullDutyBaseline) deepCopy() *_BACnetConstructedDataFullDutyBaseline {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFullDutyBaselineCopy := &_BACnetConstructedDataFullDutyBaseline{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.FullDutyBaseLine.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFullDutyBaselineCopy
}

func (m *_BACnetConstructedDataFullDutyBaseline) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
