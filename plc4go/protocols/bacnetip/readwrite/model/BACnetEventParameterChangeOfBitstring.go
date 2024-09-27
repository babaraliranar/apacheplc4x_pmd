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

// BACnetEventParameterChangeOfBitstring is the corresponding interface of BACnetEventParameterChangeOfBitstring
type BACnetEventParameterChangeOfBitstring interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetBitmask returns Bitmask (property field)
	GetBitmask() BACnetContextTagBitString
	// GetListOfBitstringValues returns ListOfBitstringValues (property field)
	GetListOfBitstringValues() BACnetEventParameterChangeOfBitstringListOfBitstringValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfBitstring is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfBitstring()
	// CreateBuilder creates a BACnetEventParameterChangeOfBitstringBuilder
	CreateBACnetEventParameterChangeOfBitstringBuilder() BACnetEventParameterChangeOfBitstringBuilder
}

// _BACnetEventParameterChangeOfBitstring is the data-structure of this message
type _BACnetEventParameterChangeOfBitstring struct {
	BACnetEventParameterContract
	OpeningTag            BACnetOpeningTag
	TimeDelay             BACnetContextTagUnsignedInteger
	Bitmask               BACnetContextTagBitString
	ListOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues
	ClosingTag            BACnetClosingTag
}

var _ BACnetEventParameterChangeOfBitstring = (*_BACnetEventParameterChangeOfBitstring)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfBitstring)(nil)

// NewBACnetEventParameterChangeOfBitstring factory function for _BACnetEventParameterChangeOfBitstring
func NewBACnetEventParameterChangeOfBitstring(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, bitmask BACnetContextTagBitString, listOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues, closingTag BACnetClosingTag) *_BACnetEventParameterChangeOfBitstring {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfBitstring must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterChangeOfBitstring must not be nil")
	}
	if bitmask == nil {
		panic("bitmask of type BACnetContextTagBitString for BACnetEventParameterChangeOfBitstring must not be nil")
	}
	if listOfBitstringValues == nil {
		panic("listOfBitstringValues of type BACnetEventParameterChangeOfBitstringListOfBitstringValues for BACnetEventParameterChangeOfBitstring must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfBitstring must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfBitstring{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		Bitmask:                      bitmask,
		ListOfBitstringValues:        listOfBitstringValues,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfBitstringBuilder is a builder for BACnetEventParameterChangeOfBitstring
type BACnetEventParameterChangeOfBitstringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, bitmask BACnetContextTagBitString, listOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues, closingTag BACnetClosingTag) BACnetEventParameterChangeOfBitstringBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfBitstringBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfBitstringBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfBitstringBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfBitstringBuilder
	// WithBitmask adds Bitmask (property field)
	WithBitmask(BACnetContextTagBitString) BACnetEventParameterChangeOfBitstringBuilder
	// WithBitmaskBuilder adds Bitmask (property field) which is build by the builder
	WithBitmaskBuilder(func(BACnetContextTagBitStringBuilder) BACnetContextTagBitStringBuilder) BACnetEventParameterChangeOfBitstringBuilder
	// WithListOfBitstringValues adds ListOfBitstringValues (property field)
	WithListOfBitstringValues(BACnetEventParameterChangeOfBitstringListOfBitstringValues) BACnetEventParameterChangeOfBitstringBuilder
	// WithListOfBitstringValuesBuilder adds ListOfBitstringValues (property field) which is build by the builder
	WithListOfBitstringValuesBuilder(func(BACnetEventParameterChangeOfBitstringListOfBitstringValuesBuilder) BACnetEventParameterChangeOfBitstringListOfBitstringValuesBuilder) BACnetEventParameterChangeOfBitstringBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfBitstringBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfBitstringBuilder
	// Build builds the BACnetEventParameterChangeOfBitstring or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfBitstring, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfBitstring
}

// NewBACnetEventParameterChangeOfBitstringBuilder() creates a BACnetEventParameterChangeOfBitstringBuilder
func NewBACnetEventParameterChangeOfBitstringBuilder() BACnetEventParameterChangeOfBitstringBuilder {
	return &_BACnetEventParameterChangeOfBitstringBuilder{_BACnetEventParameterChangeOfBitstring: new(_BACnetEventParameterChangeOfBitstring)}
}

type _BACnetEventParameterChangeOfBitstringBuilder struct {
	*_BACnetEventParameterChangeOfBitstring

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfBitstringBuilder) = (*_BACnetEventParameterChangeOfBitstringBuilder)(nil)

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, bitmask BACnetContextTagBitString, listOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues, closingTag BACnetClosingTag) BACnetEventParameterChangeOfBitstringBuilder {
	return m.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithBitmask(bitmask).WithListOfBitstringValues(listOfBitstringValues).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfBitstringBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfBitstringBuilder {
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

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfBitstringBuilder {
	m.TimeDelay = timeDelay
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfBitstringBuilder {
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

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithBitmask(bitmask BACnetContextTagBitString) BACnetEventParameterChangeOfBitstringBuilder {
	m.Bitmask = bitmask
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithBitmaskBuilder(builderSupplier func(BACnetContextTagBitStringBuilder) BACnetContextTagBitStringBuilder) BACnetEventParameterChangeOfBitstringBuilder {
	builder := builderSupplier(m.Bitmask.CreateBACnetContextTagBitStringBuilder())
	var err error
	m.Bitmask, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagBitStringBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithListOfBitstringValues(listOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues) BACnetEventParameterChangeOfBitstringBuilder {
	m.ListOfBitstringValues = listOfBitstringValues
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithListOfBitstringValuesBuilder(builderSupplier func(BACnetEventParameterChangeOfBitstringListOfBitstringValuesBuilder) BACnetEventParameterChangeOfBitstringListOfBitstringValuesBuilder) BACnetEventParameterChangeOfBitstringBuilder {
	builder := builderSupplier(m.ListOfBitstringValues.CreateBACnetEventParameterChangeOfBitstringListOfBitstringValuesBuilder())
	var err error
	m.ListOfBitstringValues, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEventParameterChangeOfBitstringListOfBitstringValuesBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfBitstringBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfBitstringBuilder {
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

func (m *_BACnetEventParameterChangeOfBitstringBuilder) Build() (BACnetEventParameterChangeOfBitstring, error) {
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
	if m.Bitmask == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'bitmask' not set"))
	}
	if m.ListOfBitstringValues == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'listOfBitstringValues' not set"))
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
	return m._BACnetEventParameterChangeOfBitstring.deepCopy(), nil
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) MustBuild() BACnetEventParameterChangeOfBitstring {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterChangeOfBitstringBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterChangeOfBitstringBuilder()
}

// CreateBACnetEventParameterChangeOfBitstringBuilder creates a BACnetEventParameterChangeOfBitstringBuilder
func (m *_BACnetEventParameterChangeOfBitstring) CreateBACnetEventParameterChangeOfBitstringBuilder() BACnetEventParameterChangeOfBitstringBuilder {
	if m == nil {
		return NewBACnetEventParameterChangeOfBitstringBuilder()
	}
	return &_BACnetEventParameterChangeOfBitstringBuilder{_BACnetEventParameterChangeOfBitstring: m.deepCopy()}
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

func (m *_BACnetEventParameterChangeOfBitstring) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfBitstring) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfBitstring) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfBitstring) GetBitmask() BACnetContextTagBitString {
	return m.Bitmask
}

func (m *_BACnetEventParameterChangeOfBitstring) GetListOfBitstringValues() BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	return m.ListOfBitstringValues
}

func (m *_BACnetEventParameterChangeOfBitstring) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfBitstring(structType any) BACnetEventParameterChangeOfBitstring {
	if casted, ok := structType.(BACnetEventParameterChangeOfBitstring); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfBitstring); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfBitstring) GetTypeName() string {
	return "BACnetEventParameterChangeOfBitstring"
}

func (m *_BACnetEventParameterChangeOfBitstring) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (bitmask)
	lengthInBits += m.Bitmask.GetLengthInBits(ctx)

	// Simple field (listOfBitstringValues)
	lengthInBits += m.ListOfBitstringValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfBitstring) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfBitstring) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfBitstring BACnetEventParameterChangeOfBitstring, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfBitstring"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfBitstring")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	bitmask, err := ReadSimpleField[BACnetContextTagBitString](ctx, "bitmask", ReadComplex[BACnetContextTagBitString](BACnetContextTagParseWithBufferProducer[BACnetContextTagBitString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BIT_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitmask' field"))
	}
	m.Bitmask = bitmask

	listOfBitstringValues, err := ReadSimpleField[BACnetEventParameterChangeOfBitstringListOfBitstringValues](ctx, "listOfBitstringValues", ReadComplex[BACnetEventParameterChangeOfBitstringListOfBitstringValues](BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfBitstringValues' field"))
	}
	m.ListOfBitstringValues = listOfBitstringValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfBitstring"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfBitstring")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfBitstring) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfBitstring) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfBitstring"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfBitstring")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagBitString](ctx, "bitmask", m.GetBitmask(), WriteComplex[BACnetContextTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitmask' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfBitstringListOfBitstringValues](ctx, "listOfBitstringValues", m.GetListOfBitstringValues(), WriteComplex[BACnetEventParameterChangeOfBitstringListOfBitstringValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfBitstringValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfBitstring"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfBitstring")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfBitstring) IsBACnetEventParameterChangeOfBitstring() {}

func (m *_BACnetEventParameterChangeOfBitstring) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfBitstring) deepCopy() *_BACnetEventParameterChangeOfBitstring {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfBitstringCopy := &_BACnetEventParameterChangeOfBitstring{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.TimeDelay.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Bitmask.DeepCopy().(BACnetContextTagBitString),
		m.ListOfBitstringValues.DeepCopy().(BACnetEventParameterChangeOfBitstringListOfBitstringValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterChangeOfBitstringCopy
}

func (m *_BACnetEventParameterChangeOfBitstring) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
