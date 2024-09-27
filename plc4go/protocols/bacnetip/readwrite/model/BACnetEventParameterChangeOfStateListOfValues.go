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

// BACnetEventParameterChangeOfStateListOfValues is the corresponding interface of BACnetEventParameterChangeOfStateListOfValues
type BACnetEventParameterChangeOfStateListOfValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() []BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfStateListOfValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfStateListOfValues()
	// CreateBuilder creates a BACnetEventParameterChangeOfStateListOfValuesBuilder
	CreateBACnetEventParameterChangeOfStateListOfValuesBuilder() BACnetEventParameterChangeOfStateListOfValuesBuilder
}

// _BACnetEventParameterChangeOfStateListOfValues is the data-structure of this message
type _BACnetEventParameterChangeOfStateListOfValues struct {
	OpeningTag   BACnetOpeningTag
	ListOfValues []BACnetPropertyStates
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfStateListOfValues = (*_BACnetEventParameterChangeOfStateListOfValues)(nil)

// NewBACnetEventParameterChangeOfStateListOfValues factory function for _BACnetEventParameterChangeOfStateListOfValues
func NewBACnetEventParameterChangeOfStateListOfValues(openingTag BACnetOpeningTag, listOfValues []BACnetPropertyStates, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfStateListOfValues {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfStateListOfValues must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfStateListOfValues must not be nil")
	}
	return &_BACnetEventParameterChangeOfStateListOfValues{OpeningTag: openingTag, ListOfValues: listOfValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfStateListOfValuesBuilder is a builder for BACnetEventParameterChangeOfStateListOfValues
type BACnetEventParameterChangeOfStateListOfValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfValues []BACnetPropertyStates, closingTag BACnetClosingTag) BACnetEventParameterChangeOfStateListOfValuesBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfStateListOfValuesBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfStateListOfValuesBuilder
	// WithListOfValues adds ListOfValues (property field)
	WithListOfValues(...BACnetPropertyStates) BACnetEventParameterChangeOfStateListOfValuesBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfStateListOfValuesBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfStateListOfValuesBuilder
	// Build builds the BACnetEventParameterChangeOfStateListOfValues or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfStateListOfValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfStateListOfValues
}

// NewBACnetEventParameterChangeOfStateListOfValuesBuilder() creates a BACnetEventParameterChangeOfStateListOfValuesBuilder
func NewBACnetEventParameterChangeOfStateListOfValuesBuilder() BACnetEventParameterChangeOfStateListOfValuesBuilder {
	return &_BACnetEventParameterChangeOfStateListOfValuesBuilder{_BACnetEventParameterChangeOfStateListOfValues: new(_BACnetEventParameterChangeOfStateListOfValues)}
}

type _BACnetEventParameterChangeOfStateListOfValuesBuilder struct {
	*_BACnetEventParameterChangeOfStateListOfValues

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfStateListOfValuesBuilder) = (*_BACnetEventParameterChangeOfStateListOfValuesBuilder)(nil)

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfValues []BACnetPropertyStates, closingTag BACnetClosingTag) BACnetEventParameterChangeOfStateListOfValuesBuilder {
	return m.WithOpeningTag(openingTag).WithListOfValues(listOfValues...).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfStateListOfValuesBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfStateListOfValuesBuilder {
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

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) WithListOfValues(listOfValues ...BACnetPropertyStates) BACnetEventParameterChangeOfStateListOfValuesBuilder {
	m.ListOfValues = listOfValues
	return m
}

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfStateListOfValuesBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfStateListOfValuesBuilder {
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

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) Build() (BACnetEventParameterChangeOfStateListOfValues, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
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
	return m._BACnetEventParameterChangeOfStateListOfValues.deepCopy(), nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) MustBuild() BACnetEventParameterChangeOfStateListOfValues {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterChangeOfStateListOfValuesBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterChangeOfStateListOfValuesBuilder()
}

// CreateBACnetEventParameterChangeOfStateListOfValuesBuilder creates a BACnetEventParameterChangeOfStateListOfValuesBuilder
func (m *_BACnetEventParameterChangeOfStateListOfValues) CreateBACnetEventParameterChangeOfStateListOfValuesBuilder() BACnetEventParameterChangeOfStateListOfValuesBuilder {
	if m == nil {
		return NewBACnetEventParameterChangeOfStateListOfValuesBuilder()
	}
	return &_BACnetEventParameterChangeOfStateListOfValuesBuilder{_BACnetEventParameterChangeOfStateListOfValues: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetListOfValues() []BACnetPropertyStates {
	return m.ListOfValues
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfStateListOfValues(structType any) BACnetEventParameterChangeOfStateListOfValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfStateListOfValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfStateListOfValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfStateListOfValues"
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfValues) > 0 {
		for _, element := range m.ListOfValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfStateListOfValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfStateListOfValues, error) {
	return BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfStateListOfValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfStateListOfValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfStateListOfValues, error) {
		return BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfStateListOfValues, error) {
	v, err := (&_BACnetEventParameterChangeOfStateListOfValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterChangeOfStateListOfValues BACnetEventParameterChangeOfStateListOfValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfStateListOfValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfStateListOfValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfValues, err := ReadTerminatedArrayField[BACnetPropertyStates](ctx, "listOfValues", ReadComplex[BACnetPropertyStates](BACnetPropertyStatesParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}
	m.ListOfValues = listOfValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfStateListOfValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfStateListOfValues")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfStateListOfValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfStateListOfValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfValues", m.GetListOfValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfStateListOfValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfStateListOfValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfStateListOfValues) IsBACnetEventParameterChangeOfStateListOfValues() {
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) deepCopy() *_BACnetEventParameterChangeOfStateListOfValues {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfStateListOfValuesCopy := &_BACnetEventParameterChangeOfStateListOfValues{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetPropertyStates, BACnetPropertyStates](m.ListOfValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetEventParameterChangeOfStateListOfValuesCopy
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
