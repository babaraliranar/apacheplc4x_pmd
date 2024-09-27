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

// BACnetAuthenticationPolicyList is the corresponding interface of BACnetAuthenticationPolicyList
type BACnetAuthenticationPolicyList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetEntries returns Entries (property field)
	GetEntries() []BACnetAuthenticationPolicyListEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetAuthenticationPolicyList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAuthenticationPolicyList()
	// CreateBuilder creates a BACnetAuthenticationPolicyListBuilder
	CreateBACnetAuthenticationPolicyListBuilder() BACnetAuthenticationPolicyListBuilder
}

// _BACnetAuthenticationPolicyList is the data-structure of this message
type _BACnetAuthenticationPolicyList struct {
	OpeningTag BACnetOpeningTag
	Entries    []BACnetAuthenticationPolicyListEntry
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetAuthenticationPolicyList = (*_BACnetAuthenticationPolicyList)(nil)

// NewBACnetAuthenticationPolicyList factory function for _BACnetAuthenticationPolicyList
func NewBACnetAuthenticationPolicyList(openingTag BACnetOpeningTag, entries []BACnetAuthenticationPolicyListEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAuthenticationPolicyList {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetAuthenticationPolicyList must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetAuthenticationPolicyList must not be nil")
	}
	return &_BACnetAuthenticationPolicyList{OpeningTag: openingTag, Entries: entries, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAuthenticationPolicyListBuilder is a builder for BACnetAuthenticationPolicyList
type BACnetAuthenticationPolicyListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, entries []BACnetAuthenticationPolicyListEntry, closingTag BACnetClosingTag) BACnetAuthenticationPolicyListBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetAuthenticationPolicyListBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAuthenticationPolicyListBuilder
	// WithEntries adds Entries (property field)
	WithEntries(...BACnetAuthenticationPolicyListEntry) BACnetAuthenticationPolicyListBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetAuthenticationPolicyListBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAuthenticationPolicyListBuilder
	// Build builds the BACnetAuthenticationPolicyList or returns an error if something is wrong
	Build() (BACnetAuthenticationPolicyList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAuthenticationPolicyList
}

// NewBACnetAuthenticationPolicyListBuilder() creates a BACnetAuthenticationPolicyListBuilder
func NewBACnetAuthenticationPolicyListBuilder() BACnetAuthenticationPolicyListBuilder {
	return &_BACnetAuthenticationPolicyListBuilder{_BACnetAuthenticationPolicyList: new(_BACnetAuthenticationPolicyList)}
}

type _BACnetAuthenticationPolicyListBuilder struct {
	*_BACnetAuthenticationPolicyList

	err *utils.MultiError
}

var _ (BACnetAuthenticationPolicyListBuilder) = (*_BACnetAuthenticationPolicyListBuilder)(nil)

func (m *_BACnetAuthenticationPolicyListBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, entries []BACnetAuthenticationPolicyListEntry, closingTag BACnetClosingTag) BACnetAuthenticationPolicyListBuilder {
	return m.WithOpeningTag(openingTag).WithEntries(entries...).WithClosingTag(closingTag)
}

func (m *_BACnetAuthenticationPolicyListBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetAuthenticationPolicyListBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetAuthenticationPolicyListBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAuthenticationPolicyListBuilder {
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

func (m *_BACnetAuthenticationPolicyListBuilder) WithEntries(entries ...BACnetAuthenticationPolicyListEntry) BACnetAuthenticationPolicyListBuilder {
	m.Entries = entries
	return m
}

func (m *_BACnetAuthenticationPolicyListBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetAuthenticationPolicyListBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetAuthenticationPolicyListBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAuthenticationPolicyListBuilder {
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

func (m *_BACnetAuthenticationPolicyListBuilder) Build() (BACnetAuthenticationPolicyList, error) {
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
	return m._BACnetAuthenticationPolicyList.deepCopy(), nil
}

func (m *_BACnetAuthenticationPolicyListBuilder) MustBuild() BACnetAuthenticationPolicyList {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetAuthenticationPolicyListBuilder) DeepCopy() any {
	return m.CreateBACnetAuthenticationPolicyListBuilder()
}

// CreateBACnetAuthenticationPolicyListBuilder creates a BACnetAuthenticationPolicyListBuilder
func (m *_BACnetAuthenticationPolicyList) CreateBACnetAuthenticationPolicyListBuilder() BACnetAuthenticationPolicyListBuilder {
	if m == nil {
		return NewBACnetAuthenticationPolicyListBuilder()
	}
	return &_BACnetAuthenticationPolicyListBuilder{_BACnetAuthenticationPolicyList: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationPolicyList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAuthenticationPolicyList) GetEntries() []BACnetAuthenticationPolicyListEntry {
	return m.Entries
}

func (m *_BACnetAuthenticationPolicyList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationPolicyList(structType any) BACnetAuthenticationPolicyList {
	if casted, ok := structType.(BACnetAuthenticationPolicyList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicyList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationPolicyList) GetTypeName() string {
	return "BACnetAuthenticationPolicyList"
}

func (m *_BACnetAuthenticationPolicyList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Entries) > 0 {
		for _, element := range m.Entries {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationPolicyList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationPolicyListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAuthenticationPolicyList, error) {
	return BACnetAuthenticationPolicyListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAuthenticationPolicyListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyList, error) {
		return BACnetAuthenticationPolicyListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetAuthenticationPolicyListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAuthenticationPolicyList, error) {
	v, err := (&_BACnetAuthenticationPolicyList{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAuthenticationPolicyList) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetAuthenticationPolicyList BACnetAuthenticationPolicyList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicyList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicyList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	entries, err := ReadTerminatedArrayField[BACnetAuthenticationPolicyListEntry](ctx, "entries", ReadComplex[BACnetAuthenticationPolicyListEntry](BACnetAuthenticationPolicyListEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entries' field"))
	}
	m.Entries = entries

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicyList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicyList")
	}

	return m, nil
}

func (m *_BACnetAuthenticationPolicyList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationPolicyList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicyList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicyList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "entries", m.GetEntries(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'entries' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicyList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationPolicyList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthenticationPolicyList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAuthenticationPolicyList) IsBACnetAuthenticationPolicyList() {}

func (m *_BACnetAuthenticationPolicyList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAuthenticationPolicyList) deepCopy() *_BACnetAuthenticationPolicyList {
	if m == nil {
		return nil
	}
	_BACnetAuthenticationPolicyListCopy := &_BACnetAuthenticationPolicyList{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetAuthenticationPolicyListEntry, BACnetAuthenticationPolicyListEntry](m.Entries),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetAuthenticationPolicyListCopy
}

func (m *_BACnetAuthenticationPolicyList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
