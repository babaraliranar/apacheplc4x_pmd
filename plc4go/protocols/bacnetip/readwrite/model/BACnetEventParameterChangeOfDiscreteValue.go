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

// BACnetEventParameterChangeOfDiscreteValue is the corresponding interface of BACnetEventParameterChangeOfDiscreteValue
type BACnetEventParameterChangeOfDiscreteValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfDiscreteValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfDiscreteValue()
	// CreateBuilder creates a BACnetEventParameterChangeOfDiscreteValueBuilder
	CreateBACnetEventParameterChangeOfDiscreteValueBuilder() BACnetEventParameterChangeOfDiscreteValueBuilder
}

// _BACnetEventParameterChangeOfDiscreteValue is the data-structure of this message
type _BACnetEventParameterChangeOfDiscreteValue struct {
	BACnetEventParameterContract
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	ClosingTag BACnetClosingTag
}

var _ BACnetEventParameterChangeOfDiscreteValue = (*_BACnetEventParameterChangeOfDiscreteValue)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfDiscreteValue)(nil)

// NewBACnetEventParameterChangeOfDiscreteValue factory function for _BACnetEventParameterChangeOfDiscreteValue
func NewBACnetEventParameterChangeOfDiscreteValue(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) *_BACnetEventParameterChangeOfDiscreteValue {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfDiscreteValue must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterChangeOfDiscreteValue must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfDiscreteValue must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfDiscreteValue{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfDiscreteValueBuilder is a builder for BACnetEventParameterChangeOfDiscreteValue
type BACnetEventParameterChangeOfDiscreteValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterChangeOfDiscreteValueBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfDiscreteValueBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfDiscreteValueBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfDiscreteValueBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfDiscreteValueBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfDiscreteValueBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfDiscreteValueBuilder
	// Build builds the BACnetEventParameterChangeOfDiscreteValue or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfDiscreteValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfDiscreteValue
}

// NewBACnetEventParameterChangeOfDiscreteValueBuilder() creates a BACnetEventParameterChangeOfDiscreteValueBuilder
func NewBACnetEventParameterChangeOfDiscreteValueBuilder() BACnetEventParameterChangeOfDiscreteValueBuilder {
	return &_BACnetEventParameterChangeOfDiscreteValueBuilder{_BACnetEventParameterChangeOfDiscreteValue: new(_BACnetEventParameterChangeOfDiscreteValue)}
}

type _BACnetEventParameterChangeOfDiscreteValueBuilder struct {
	*_BACnetEventParameterChangeOfDiscreteValue

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfDiscreteValueBuilder) = (*_BACnetEventParameterChangeOfDiscreteValueBuilder)(nil)

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterChangeOfDiscreteValueBuilder {
	return m.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfDiscreteValueBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfDiscreteValueBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfDiscreteValueBuilder {
	m.TimeDelay = timeDelay
	return m
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfDiscreteValueBuilder {
	builder := builderSupplier(m.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.TimeDelay, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfDiscreteValueBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfDiscreteValueBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) Build() (BACnetEventParameterChangeOfDiscreteValue, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.TimeDelay == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetEventParameterChangeOfDiscreteValue.deepCopy(), nil
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) MustBuild() BACnetEventParameterChangeOfDiscreteValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterChangeOfDiscreteValueBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterChangeOfDiscreteValueBuilder()
}

// CreateBACnetEventParameterChangeOfDiscreteValueBuilder creates a BACnetEventParameterChangeOfDiscreteValueBuilder
func (m *_BACnetEventParameterChangeOfDiscreteValue) CreateBACnetEventParameterChangeOfDiscreteValueBuilder() BACnetEventParameterChangeOfDiscreteValueBuilder {
	if m == nil {
		return NewBACnetEventParameterChangeOfDiscreteValueBuilder()
	}
	return &_BACnetEventParameterChangeOfDiscreteValueBuilder{_BACnetEventParameterChangeOfDiscreteValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfDiscreteValue(structType any) BACnetEventParameterChangeOfDiscreteValue {
	if casted, ok := structType.(BACnetEventParameterChangeOfDiscreteValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfDiscreteValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetTypeName() string {
	return "BACnetEventParameterChangeOfDiscreteValue"
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfDiscreteValue BACnetEventParameterChangeOfDiscreteValue, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfDiscreteValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfDiscreteValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(21))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(21))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfDiscreteValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfDiscreteValue")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfDiscreteValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfDiscreteValue")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfDiscreteValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfDiscreteValue")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) IsBACnetEventParameterChangeOfDiscreteValue() {}

func (m *_BACnetEventParameterChangeOfDiscreteValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) deepCopy() *_BACnetEventParameterChangeOfDiscreteValue {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfDiscreteValueCopy := &_BACnetEventParameterChangeOfDiscreteValue{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.TimeDelay.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterChangeOfDiscreteValueCopy
}

func (m *_BACnetEventParameterChangeOfDiscreteValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
