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

// BACnetNotificationParametersBufferReady is the corresponding interface of BACnetNotificationParametersBufferReady
type BACnetNotificationParametersBufferReady interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetBufferProperty returns BufferProperty (property field)
	GetBufferProperty() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetPreviousNotification returns PreviousNotification (property field)
	GetPreviousNotification() BACnetContextTagUnsignedInteger
	// GetCurrentNotification returns CurrentNotification (property field)
	GetCurrentNotification() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersBufferReady is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersBufferReady()
	// CreateBuilder creates a BACnetNotificationParametersBufferReadyBuilder
	CreateBACnetNotificationParametersBufferReadyBuilder() BACnetNotificationParametersBufferReadyBuilder
}

// _BACnetNotificationParametersBufferReady is the data-structure of this message
type _BACnetNotificationParametersBufferReady struct {
	BACnetNotificationParametersContract
	InnerOpeningTag      BACnetOpeningTag
	BufferProperty       BACnetDeviceObjectPropertyReferenceEnclosed
	PreviousNotification BACnetContextTagUnsignedInteger
	CurrentNotification  BACnetContextTagUnsignedInteger
	InnerClosingTag      BACnetClosingTag
}

var _ BACnetNotificationParametersBufferReady = (*_BACnetNotificationParametersBufferReady)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersBufferReady)(nil)

// NewBACnetNotificationParametersBufferReady factory function for _BACnetNotificationParametersBufferReady
func NewBACnetNotificationParametersBufferReady(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, bufferProperty BACnetDeviceObjectPropertyReferenceEnclosed, previousNotification BACnetContextTagUnsignedInteger, currentNotification BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersBufferReady {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersBufferReady must not be nil")
	}
	if bufferProperty == nil {
		panic("bufferProperty of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetNotificationParametersBufferReady must not be nil")
	}
	if previousNotification == nil {
		panic("previousNotification of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersBufferReady must not be nil")
	}
	if currentNotification == nil {
		panic("currentNotification of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersBufferReady must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersBufferReady must not be nil")
	}
	_result := &_BACnetNotificationParametersBufferReady{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		BufferProperty:                       bufferProperty,
		PreviousNotification:                 previousNotification,
		CurrentNotification:                  currentNotification,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersBufferReadyBuilder is a builder for BACnetNotificationParametersBufferReady
type BACnetNotificationParametersBufferReadyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, bufferProperty BACnetDeviceObjectPropertyReferenceEnclosed, previousNotification BACnetContextTagUnsignedInteger, currentNotification BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersBufferReadyBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersBufferReadyBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersBufferReadyBuilder
	// WithBufferProperty adds BufferProperty (property field)
	WithBufferProperty(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetNotificationParametersBufferReadyBuilder
	// WithBufferPropertyBuilder adds BufferProperty (property field) which is build by the builder
	WithBufferPropertyBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetNotificationParametersBufferReadyBuilder
	// WithPreviousNotification adds PreviousNotification (property field)
	WithPreviousNotification(BACnetContextTagUnsignedInteger) BACnetNotificationParametersBufferReadyBuilder
	// WithPreviousNotificationBuilder adds PreviousNotification (property field) which is build by the builder
	WithPreviousNotificationBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersBufferReadyBuilder
	// WithCurrentNotification adds CurrentNotification (property field)
	WithCurrentNotification(BACnetContextTagUnsignedInteger) BACnetNotificationParametersBufferReadyBuilder
	// WithCurrentNotificationBuilder adds CurrentNotification (property field) which is build by the builder
	WithCurrentNotificationBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersBufferReadyBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersBufferReadyBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersBufferReadyBuilder
	// Build builds the BACnetNotificationParametersBufferReady or returns an error if something is wrong
	Build() (BACnetNotificationParametersBufferReady, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersBufferReady
}

// NewBACnetNotificationParametersBufferReadyBuilder() creates a BACnetNotificationParametersBufferReadyBuilder
func NewBACnetNotificationParametersBufferReadyBuilder() BACnetNotificationParametersBufferReadyBuilder {
	return &_BACnetNotificationParametersBufferReadyBuilder{_BACnetNotificationParametersBufferReady: new(_BACnetNotificationParametersBufferReady)}
}

type _BACnetNotificationParametersBufferReadyBuilder struct {
	*_BACnetNotificationParametersBufferReady

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersBufferReadyBuilder) = (*_BACnetNotificationParametersBufferReadyBuilder)(nil)

func (b *_BACnetNotificationParametersBufferReadyBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, bufferProperty BACnetDeviceObjectPropertyReferenceEnclosed, previousNotification BACnetContextTagUnsignedInteger, currentNotification BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersBufferReadyBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithBufferProperty(bufferProperty).WithPreviousNotification(previousNotification).WithCurrentNotification(currentNotification).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersBufferReadyBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersBufferReadyBuilder {
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

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithBufferProperty(bufferProperty BACnetDeviceObjectPropertyReferenceEnclosed) BACnetNotificationParametersBufferReadyBuilder {
	b.BufferProperty = bufferProperty
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithBufferPropertyBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetNotificationParametersBufferReadyBuilder {
	builder := builderSupplier(b.BufferProperty.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.BufferProperty, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithPreviousNotification(previousNotification BACnetContextTagUnsignedInteger) BACnetNotificationParametersBufferReadyBuilder {
	b.PreviousNotification = previousNotification
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithPreviousNotificationBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersBufferReadyBuilder {
	builder := builderSupplier(b.PreviousNotification.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.PreviousNotification, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithCurrentNotification(currentNotification BACnetContextTagUnsignedInteger) BACnetNotificationParametersBufferReadyBuilder {
	b.CurrentNotification = currentNotification
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithCurrentNotificationBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersBufferReadyBuilder {
	builder := builderSupplier(b.CurrentNotification.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.CurrentNotification, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersBufferReadyBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersBufferReadyBuilder {
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

func (b *_BACnetNotificationParametersBufferReadyBuilder) Build() (BACnetNotificationParametersBufferReady, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.BufferProperty == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bufferProperty' not set"))
	}
	if b.PreviousNotification == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'previousNotification' not set"))
	}
	if b.CurrentNotification == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'currentNotification' not set"))
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
	return b._BACnetNotificationParametersBufferReady.deepCopy(), nil
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) MustBuild() BACnetNotificationParametersBufferReady {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetNotificationParametersBufferReadyBuilder) Done() BACnetNotificationParametersBuilder {
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersBufferReadyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersBufferReadyBuilder().(*_BACnetNotificationParametersBufferReadyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersBufferReadyBuilder creates a BACnetNotificationParametersBufferReadyBuilder
func (b *_BACnetNotificationParametersBufferReady) CreateBACnetNotificationParametersBufferReadyBuilder() BACnetNotificationParametersBufferReadyBuilder {
	if b == nil {
		return NewBACnetNotificationParametersBufferReadyBuilder()
	}
	return &_BACnetNotificationParametersBufferReadyBuilder{_BACnetNotificationParametersBufferReady: b.deepCopy()}
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

func (m *_BACnetNotificationParametersBufferReady) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersBufferReady) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersBufferReady) GetBufferProperty() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.BufferProperty
}

func (m *_BACnetNotificationParametersBufferReady) GetPreviousNotification() BACnetContextTagUnsignedInteger {
	return m.PreviousNotification
}

func (m *_BACnetNotificationParametersBufferReady) GetCurrentNotification() BACnetContextTagUnsignedInteger {
	return m.CurrentNotification
}

func (m *_BACnetNotificationParametersBufferReady) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersBufferReady(structType any) BACnetNotificationParametersBufferReady {
	if casted, ok := structType.(BACnetNotificationParametersBufferReady); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersBufferReady); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersBufferReady) GetTypeName() string {
	return "BACnetNotificationParametersBufferReady"
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (bufferProperty)
	lengthInBits += m.BufferProperty.GetLengthInBits(ctx)

	// Simple field (previousNotification)
	lengthInBits += m.PreviousNotification.GetLengthInBits(ctx)

	// Simple field (currentNotification)
	lengthInBits += m.CurrentNotification.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersBufferReady) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersBufferReady BACnetNotificationParametersBufferReady, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersBufferReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersBufferReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	bufferProperty, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "bufferProperty", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bufferProperty' field"))
	}
	m.BufferProperty = bufferProperty

	previousNotification, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "previousNotification", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'previousNotification' field"))
	}
	m.PreviousNotification = previousNotification

	currentNotification, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "currentNotification", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentNotification' field"))
	}
	m.CurrentNotification = currentNotification

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersBufferReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersBufferReady")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersBufferReady) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersBufferReady) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersBufferReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersBufferReady")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "bufferProperty", m.GetBufferProperty(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bufferProperty' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "previousNotification", m.GetPreviousNotification(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'previousNotification' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "currentNotification", m.GetCurrentNotification(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentNotification' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersBufferReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersBufferReady")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersBufferReady) IsBACnetNotificationParametersBufferReady() {}

func (m *_BACnetNotificationParametersBufferReady) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersBufferReady) deepCopy() *_BACnetNotificationParametersBufferReady {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersBufferReadyCopy := &_BACnetNotificationParametersBufferReady{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.InnerOpeningTag.DeepCopy().(BACnetOpeningTag),
		m.BufferProperty.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.PreviousNotification.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.CurrentNotification.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.InnerClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersBufferReadyCopy
}

func (m *_BACnetNotificationParametersBufferReady) String() string {
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
