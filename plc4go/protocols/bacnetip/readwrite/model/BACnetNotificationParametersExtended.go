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

// BACnetNotificationParametersExtended is the corresponding interface of BACnetNotificationParametersExtended
type BACnetNotificationParametersExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedEventType returns ExtendedEventType (property field)
	GetExtendedEventType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetNotificationParametersExtendedParameters
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersExtended is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersExtended()
	// CreateBuilder creates a BACnetNotificationParametersExtendedBuilder
	CreateBACnetNotificationParametersExtendedBuilder() BACnetNotificationParametersExtendedBuilder
}

// _BACnetNotificationParametersExtended is the data-structure of this message
type _BACnetNotificationParametersExtended struct {
	BACnetNotificationParametersContract
	InnerOpeningTag   BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedEventType BACnetContextTagUnsignedInteger
	Parameters        BACnetNotificationParametersExtendedParameters
	InnerClosingTag   BACnetClosingTag
}

var _ BACnetNotificationParametersExtended = (*_BACnetNotificationParametersExtended)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersExtended)(nil)

// NewBACnetNotificationParametersExtended factory function for _BACnetNotificationParametersExtended
func NewBACnetNotificationParametersExtended(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetNotificationParametersExtendedParameters, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersExtended {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersExtended must not be nil")
	}
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for BACnetNotificationParametersExtended must not be nil")
	}
	if extendedEventType == nil {
		panic("extendedEventType of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersExtended must not be nil")
	}
	if parameters == nil {
		panic("parameters of type BACnetNotificationParametersExtendedParameters for BACnetNotificationParametersExtended must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersExtended must not be nil")
	}
	_result := &_BACnetNotificationParametersExtended{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		VendorId:                             vendorId,
		ExtendedEventType:                    extendedEventType,
		Parameters:                           parameters,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersExtendedBuilder is a builder for BACnetNotificationParametersExtended
type BACnetNotificationParametersExtendedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetNotificationParametersExtendedParameters, innerClosingTag BACnetClosingTag) BACnetNotificationParametersExtendedBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersExtendedBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersExtendedBuilder
	// WithVendorId adds VendorId (property field)
	WithVendorId(BACnetVendorIdTagged) BACnetNotificationParametersExtendedBuilder
	// WithVendorIdBuilder adds VendorId (property field) which is build by the builder
	WithVendorIdBuilder(func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetNotificationParametersExtendedBuilder
	// WithExtendedEventType adds ExtendedEventType (property field)
	WithExtendedEventType(BACnetContextTagUnsignedInteger) BACnetNotificationParametersExtendedBuilder
	// WithExtendedEventTypeBuilder adds ExtendedEventType (property field) which is build by the builder
	WithExtendedEventTypeBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersExtendedBuilder
	// WithParameters adds Parameters (property field)
	WithParameters(BACnetNotificationParametersExtendedParameters) BACnetNotificationParametersExtendedBuilder
	// WithParametersBuilder adds Parameters (property field) which is build by the builder
	WithParametersBuilder(func(BACnetNotificationParametersExtendedParametersBuilder) BACnetNotificationParametersExtendedParametersBuilder) BACnetNotificationParametersExtendedBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersExtendedBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersExtendedBuilder
	// Build builds the BACnetNotificationParametersExtended or returns an error if something is wrong
	Build() (BACnetNotificationParametersExtended, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersExtended
}

// NewBACnetNotificationParametersExtendedBuilder() creates a BACnetNotificationParametersExtendedBuilder
func NewBACnetNotificationParametersExtendedBuilder() BACnetNotificationParametersExtendedBuilder {
	return &_BACnetNotificationParametersExtendedBuilder{_BACnetNotificationParametersExtended: new(_BACnetNotificationParametersExtended)}
}

type _BACnetNotificationParametersExtendedBuilder struct {
	*_BACnetNotificationParametersExtended

	err *utils.MultiError
}

var _ (BACnetNotificationParametersExtendedBuilder) = (*_BACnetNotificationParametersExtendedBuilder)(nil)

func (m *_BACnetNotificationParametersExtendedBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetNotificationParametersExtendedParameters, innerClosingTag BACnetClosingTag) BACnetNotificationParametersExtendedBuilder {
	return m.WithInnerOpeningTag(innerOpeningTag).WithVendorId(vendorId).WithExtendedEventType(extendedEventType).WithParameters(parameters).WithInnerClosingTag(innerClosingTag)
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersExtendedBuilder {
	m.InnerOpeningTag = innerOpeningTag
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersExtendedBuilder {
	builder := builderSupplier(m.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithVendorId(vendorId BACnetVendorIdTagged) BACnetNotificationParametersExtendedBuilder {
	m.VendorId = vendorId
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithVendorIdBuilder(builderSupplier func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetNotificationParametersExtendedBuilder {
	builder := builderSupplier(m.VendorId.CreateBACnetVendorIdTaggedBuilder())
	var err error
	m.VendorId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetVendorIdTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithExtendedEventType(extendedEventType BACnetContextTagUnsignedInteger) BACnetNotificationParametersExtendedBuilder {
	m.ExtendedEventType = extendedEventType
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithExtendedEventTypeBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersExtendedBuilder {
	builder := builderSupplier(m.ExtendedEventType.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ExtendedEventType, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithParameters(parameters BACnetNotificationParametersExtendedParameters) BACnetNotificationParametersExtendedBuilder {
	m.Parameters = parameters
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithParametersBuilder(builderSupplier func(BACnetNotificationParametersExtendedParametersBuilder) BACnetNotificationParametersExtendedParametersBuilder) BACnetNotificationParametersExtendedBuilder {
	builder := builderSupplier(m.Parameters.CreateBACnetNotificationParametersExtendedParametersBuilder())
	var err error
	m.Parameters, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetNotificationParametersExtendedParametersBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersExtendedBuilder {
	m.InnerClosingTag = innerClosingTag
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersExtendedBuilder {
	builder := builderSupplier(m.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.InnerClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersExtendedBuilder) Build() (BACnetNotificationParametersExtended, error) {
	if m.InnerOpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if m.VendorId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'vendorId' not set"))
	}
	if m.ExtendedEventType == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'extendedEventType' not set"))
	}
	if m.Parameters == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'parameters' not set"))
	}
	if m.InnerClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetNotificationParametersExtended.deepCopy(), nil
}

func (m *_BACnetNotificationParametersExtendedBuilder) MustBuild() BACnetNotificationParametersExtended {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetNotificationParametersExtendedBuilder) DeepCopy() any {
	return m.CreateBACnetNotificationParametersExtendedBuilder()
}

// CreateBACnetNotificationParametersExtendedBuilder creates a BACnetNotificationParametersExtendedBuilder
func (m *_BACnetNotificationParametersExtended) CreateBACnetNotificationParametersExtendedBuilder() BACnetNotificationParametersExtendedBuilder {
	if m == nil {
		return NewBACnetNotificationParametersExtendedBuilder()
	}
	return &_BACnetNotificationParametersExtendedBuilder{_BACnetNotificationParametersExtended: m.deepCopy()}
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

func (m *_BACnetNotificationParametersExtended) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersExtended) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetNotificationParametersExtended) GetExtendedEventType() BACnetContextTagUnsignedInteger {
	return m.ExtendedEventType
}

func (m *_BACnetNotificationParametersExtended) GetParameters() BACnetNotificationParametersExtendedParameters {
	return m.Parameters
}

func (m *_BACnetNotificationParametersExtended) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersExtended(structType any) BACnetNotificationParametersExtended {
	if casted, ok := structType.(BACnetNotificationParametersExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersExtended) GetTypeName() string {
	return "BACnetNotificationParametersExtended"
}

func (m *_BACnetNotificationParametersExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedEventType)
	lengthInBits += m.ExtendedEventType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersExtended BACnetNotificationParametersExtended, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	extendedEventType, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedEventType", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedEventType' field"))
	}
	m.ExtendedEventType = extendedEventType

	parameters, err := ReadSimpleField[BACnetNotificationParametersExtendedParameters](ctx, "parameters", ReadComplex[BACnetNotificationParametersExtendedParameters](BACnetNotificationParametersExtendedParametersParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}
	m.Parameters = parameters

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersExtended")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersExtended")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedEventType", m.GetExtendedEventType(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedEventType' field")
		}

		if err := WriteSimpleField[BACnetNotificationParametersExtendedParameters](ctx, "parameters", m.GetParameters(), WriteComplex[BACnetNotificationParametersExtendedParameters](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameters' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersExtended")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersExtended) IsBACnetNotificationParametersExtended() {}

func (m *_BACnetNotificationParametersExtended) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersExtended) deepCopy() *_BACnetNotificationParametersExtended {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersExtendedCopy := &_BACnetNotificationParametersExtended{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.InnerOpeningTag.DeepCopy().(BACnetOpeningTag),
		m.VendorId.DeepCopy().(BACnetVendorIdTagged),
		m.ExtendedEventType.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Parameters.DeepCopy().(BACnetNotificationParametersExtendedParameters),
		m.InnerClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersExtendedCopy
}

func (m *_BACnetNotificationParametersExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
