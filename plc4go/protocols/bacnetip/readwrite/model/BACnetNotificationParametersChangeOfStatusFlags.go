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

// BACnetNotificationParametersChangeOfStatusFlags is the corresponding interface of BACnetNotificationParametersChangeOfStatusFlags
type BACnetNotificationParametersChangeOfStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetConstructedData
	// GetReferencedFlags returns ReferencedFlags (property field)
	GetReferencedFlags() BACnetStatusFlagsTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersChangeOfStatusFlags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfStatusFlags()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfStatusFlagsBuilder
	CreateBACnetNotificationParametersChangeOfStatusFlagsBuilder() BACnetNotificationParametersChangeOfStatusFlagsBuilder
}

// _BACnetNotificationParametersChangeOfStatusFlags is the data-structure of this message
type _BACnetNotificationParametersChangeOfStatusFlags struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	PresentValue    BACnetConstructedData
	ReferencedFlags BACnetStatusFlagsTagged
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersChangeOfStatusFlags = (*_BACnetNotificationParametersChangeOfStatusFlags)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersChangeOfStatusFlags)(nil)

// NewBACnetNotificationParametersChangeOfStatusFlags factory function for _BACnetNotificationParametersChangeOfStatusFlags
func NewBACnetNotificationParametersChangeOfStatusFlags(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, presentValue BACnetConstructedData, referencedFlags BACnetStatusFlagsTagged, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfStatusFlags {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfStatusFlags must not be nil")
	}
	if presentValue == nil {
		panic("presentValue of type BACnetConstructedData for BACnetNotificationParametersChangeOfStatusFlags must not be nil")
	}
	if referencedFlags == nil {
		panic("referencedFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersChangeOfStatusFlags must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfStatusFlags must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfStatusFlags{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		PresentValue:                         presentValue,
		ReferencedFlags:                      referencedFlags,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfStatusFlagsBuilder is a builder for BACnetNotificationParametersChangeOfStatusFlags
type BACnetNotificationParametersChangeOfStatusFlagsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, presentValue BACnetConstructedData, referencedFlags BACnetStatusFlagsTagged, innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetConstructedData) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithReferencedFlags adds ReferencedFlags (property field)
	WithReferencedFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithReferencedFlagsBuilder adds ReferencedFlags (property field) which is build by the builder
	WithReferencedFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder
	// Build builds the BACnetNotificationParametersChangeOfStatusFlags or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfStatusFlags, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfStatusFlags
}

// NewBACnetNotificationParametersChangeOfStatusFlagsBuilder() creates a BACnetNotificationParametersChangeOfStatusFlagsBuilder
func NewBACnetNotificationParametersChangeOfStatusFlagsBuilder() BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	return &_BACnetNotificationParametersChangeOfStatusFlagsBuilder{_BACnetNotificationParametersChangeOfStatusFlags: new(_BACnetNotificationParametersChangeOfStatusFlags)}
}

type _BACnetNotificationParametersChangeOfStatusFlagsBuilder struct {
	*_BACnetNotificationParametersChangeOfStatusFlags

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfStatusFlagsBuilder) = (*_BACnetNotificationParametersChangeOfStatusFlagsBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, presentValue BACnetConstructedData, referencedFlags BACnetStatusFlagsTagged, innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithPresentValue(presentValue).WithReferencedFlags(referencedFlags).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithPresentValue(presentValue BACnetConstructedData) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithPresentValueBuilder(builderSupplier func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetConstructedDataBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConstructedDataBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithReferencedFlags(referencedFlags BACnetStatusFlagsTagged) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	b.ReferencedFlags = referencedFlags
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithReferencedFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	builder := builderSupplier(b.ReferencedFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.ReferencedFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) Build() (BACnetNotificationParametersChangeOfStatusFlags, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.ReferencedFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referencedFlags' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfStatusFlags.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) MustBuild() BACnetNotificationParametersChangeOfStatusFlags {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) Done() BACnetNotificationParametersBuilder {
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfStatusFlagsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfStatusFlagsBuilder().(*_BACnetNotificationParametersChangeOfStatusFlagsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfStatusFlagsBuilder creates a BACnetNotificationParametersChangeOfStatusFlagsBuilder
func (b *_BACnetNotificationParametersChangeOfStatusFlags) CreateBACnetNotificationParametersChangeOfStatusFlagsBuilder() BACnetNotificationParametersChangeOfStatusFlagsBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfStatusFlagsBuilder()
	}
	return &_BACnetNotificationParametersChangeOfStatusFlagsBuilder{_BACnetNotificationParametersChangeOfStatusFlags: b.deepCopy()}
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

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetPresentValue() BACnetConstructedData {
	return m.PresentValue
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetReferencedFlags() BACnetStatusFlagsTagged {
	return m.ReferencedFlags
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfStatusFlags(structType any) BACnetNotificationParametersChangeOfStatusFlags {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfStatusFlags"
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// Simple field (referencedFlags)
	lengthInBits += m.ReferencedFlags.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersChangeOfStatusFlags BACnetNotificationParametersChangeOfStatusFlags, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	presentValue, err := ReadSimpleField[BACnetConstructedData](ctx, "presentValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(0)), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	referencedFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "referencedFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencedFlags' field"))
	}
	m.ReferencedFlags = referencedFlags

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfStatusFlags")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfStatusFlags")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "referencedFlags", m.GetReferencedFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referencedFlags' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfStatusFlags")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) IsBACnetNotificationParametersChangeOfStatusFlags() {
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) deepCopy() *_BACnetNotificationParametersChangeOfStatusFlags {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfStatusFlagsCopy := &_BACnetNotificationParametersChangeOfStatusFlags{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.InnerOpeningTag.DeepCopy().(BACnetOpeningTag),
		m.PresentValue.DeepCopy().(BACnetConstructedData),
		m.ReferencedFlags.DeepCopy().(BACnetStatusFlagsTagged),
		m.InnerClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersChangeOfStatusFlagsCopy
}

func (m *_BACnetNotificationParametersChangeOfStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
