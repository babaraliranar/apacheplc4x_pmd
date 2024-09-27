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

// BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues is the corresponding interface of BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
type BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() []BACnetLifeSafetyStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues()
	// CreateBuilder creates a BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	CreateBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
}

// _BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues struct {
	OpeningTag        BACnetOpeningTag
	ListOfAlarmValues []BACnetLifeSafetyStateTagged
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues = (*_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues)(nil)

// NewBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues factory function for _BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
func NewBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues must not be nil")
	}
	return &_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues{OpeningTag: openingTag, ListOfAlarmValues: listOfAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder is a builder for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
type BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	// WithListOfAlarmValues adds ListOfAlarmValues (property field)
	WithListOfAlarmValues(...BACnetLifeSafetyStateTagged) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
	// Build builds the BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
}

// NewBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder() creates a BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
func NewBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
	return &_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder{_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues: new(_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues)}
}

type _BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder struct {
	*_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) = (*_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder)(nil)

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
	return m.WithOpeningTag(openingTag).WithListOfAlarmValues(listOfAlarmValues...).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
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

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) WithListOfAlarmValues(listOfAlarmValues ...BACnetLifeSafetyStateTagged) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
	m.ListOfAlarmValues = listOfAlarmValues
	return m
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
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

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) Build() (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, error) {
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
	return m._BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues.deepCopy(), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) MustBuild() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder()
}

// CreateBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder creates a BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder
func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) CreateBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder {
	if m == nil {
		return NewBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder()
	}
	return &_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder{_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetListOfAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues(structType any) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues"
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfAlarmValues) > 0 {
		for _, element := range m.ListOfAlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, error) {
	return BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, error) {
		return BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, error) {
	v, err := (&_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterChangeOfLifeSavetyListOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfAlarmValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "listOfAlarmValues", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAlarmValues' field"))
	}
	m.ListOfAlarmValues = listOfAlarmValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfAlarmValues", m.GetListOfAlarmValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfAlarmValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) IsBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues() {
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) deepCopy() *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesCopy := &_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetLifeSafetyStateTagged, BACnetLifeSafetyStateTagged](m.ListOfAlarmValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesCopy
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
