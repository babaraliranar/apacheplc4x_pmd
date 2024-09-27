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

// BACnetContextTag is the corresponding interface of BACnetContextTag
type BACnetContextTag interface {
	BACnetContextTagContract
	BACnetContextTagRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetContextTag is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTag()
	// CreateBuilder creates a BACnetContextTagBuilder
	CreateBACnetContextTagBuilder() BACnetContextTagBuilder
}

// BACnetContextTagContract provides a set of functions which can be overwritten by a sub struct
type BACnetContextTagContract interface {
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetTagNumber returns TagNumber (virtual field)
	GetTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
	// GetTagNumberArgument() returns a parser argument
	GetTagNumberArgument() uint8
	// IsBACnetContextTag is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTag()
	// CreateBuilder creates a BACnetContextTagBuilder
	CreateBACnetContextTagBuilder() BACnetContextTagBuilder
}

// BACnetContextTagRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetContextTagRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetDataType returns DataType (discriminator field)
	GetDataType() BACnetDataType
}

// _BACnetContextTag is the data-structure of this message
type _BACnetContextTag struct {
	_SubType BACnetContextTag
	Header   BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

var _ BACnetContextTagContract = (*_BACnetContextTag)(nil)

// NewBACnetContextTag factory function for _BACnetContextTag
func NewBACnetContextTag(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTag {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetContextTag must not be nil")
	}
	return &_BACnetContextTag{Header: header, TagNumberArgument: tagNumberArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagBuilder is a builder for BACnetContextTag
type BACnetContextTagBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader) BACnetContextTagBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetContextTagBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetContextTagBuilder
	// Build builds the BACnetContextTag or returns an error if something is wrong
	Build() (BACnetContextTagContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagContract
}

// NewBACnetContextTagBuilder() creates a BACnetContextTagBuilder
func NewBACnetContextTagBuilder() BACnetContextTagBuilder {
	return &_BACnetContextTagBuilder{_BACnetContextTag: new(_BACnetContextTag)}
}

type _BACnetContextTagBuilder struct {
	*_BACnetContextTag

	err *utils.MultiError
}

var _ (BACnetContextTagBuilder) = (*_BACnetContextTagBuilder)(nil)

func (m *_BACnetContextTagBuilder) WithMandatoryFields(header BACnetTagHeader) BACnetContextTagBuilder {
	return m.WithHeader(header)
}

func (m *_BACnetContextTagBuilder) WithHeader(header BACnetTagHeader) BACnetContextTagBuilder {
	m.Header = header
	return m
}

func (m *_BACnetContextTagBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetContextTagBuilder {
	builder := builderSupplier(m.Header.CreateBACnetTagHeaderBuilder())
	var err error
	m.Header, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetContextTagBuilder) Build() (BACnetContextTagContract, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetContextTag.deepCopy(), nil
}

func (m *_BACnetContextTagBuilder) MustBuild() BACnetContextTagContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetContextTagBuilder) DeepCopy() any {
	return m.CreateBACnetContextTagBuilder()
}

// CreateBACnetContextTagBuilder creates a BACnetContextTagBuilder
func (m *_BACnetContextTag) CreateBACnetContextTagBuilder() BACnetContextTagBuilder {
	if m == nil {
		return NewBACnetContextTagBuilder()
	}
	return &_BACnetContextTagBuilder{_BACnetContextTag: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetContextTag) GetTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetHeader().GetTagNumber())
}

func (pm *_BACnetContextTag) GetActualLength() uint32 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetHeader().GetActualLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTag(structType any) BACnetContextTag {
	if casted, ok := structType.(BACnetContextTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTag) GetTypeName() string {
	return "BACnetContextTag"
}

func (m *_BACnetContextTag) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTag) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetContextTagParse[T BACnetContextTag](ctx context.Context, theBytes []byte, tagNumberArgument uint8, dataType BACnetDataType) (T, error) {
	return BACnetContextTagParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumberArgument, dataType)
}

func BACnetContextTagParseWithBufferProducer[T BACnetContextTag](tagNumberArgument uint8, dataType BACnetDataType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetContextTagParseWithBuffer[T](ctx, readBuffer, tagNumberArgument, dataType)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetContextTagParseWithBuffer[T BACnetContextTag](ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (T, error) {
	v, err := (&_BACnetContextTag{TagNumberArgument: tagNumberArgument}).parse(ctx, readBuffer, tagNumberArgument, dataType)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetContextTag) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTag BACnetContextTag, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "should be a context tag"})
	}

	tagNumber, err := ReadVirtualField[uint8](ctx, "tagNumber", (*uint8)(nil), header.GetTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tagNumber' field"))
	}
	_ = tagNumber

	actualLength, err := ReadVirtualField[uint32](ctx, "actualLength", (*uint32)(nil), header.GetActualLength())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualLength' field"))
	}
	_ = actualLength

	// Validation
	if !(bool(bool((header.GetLengthValueType()) != (6))) && bool(bool((header.GetLengthValueType()) != (7)))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "length 6 and 7 reserved for opening and closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetContextTag
	switch {
	case dataType == BACnetDataType_NULL: // BACnetContextTagNull
		if _child, err = new(_BACnetContextTagNull).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagNull for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_BOOLEAN: // BACnetContextTagBoolean
		if _child, err = new(_BACnetContextTagBoolean).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagBoolean for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetContextTagUnsignedInteger
		if _child, err = new(_BACnetContextTagUnsignedInteger).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagUnsignedInteger for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetContextTagSignedInteger
		if _child, err = new(_BACnetContextTagSignedInteger).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagSignedInteger for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_REAL: // BACnetContextTagReal
		if _child, err = new(_BACnetContextTagReal).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagReal for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_DOUBLE: // BACnetContextTagDouble
		if _child, err = new(_BACnetContextTagDouble).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagDouble for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_OCTET_STRING: // BACnetContextTagOctetString
		if _child, err = new(_BACnetContextTagOctetString).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagOctetString for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetContextTagCharacterString
		if _child, err = new(_BACnetContextTagCharacterString).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagCharacterString for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_BIT_STRING: // BACnetContextTagBitString
		if _child, err = new(_BACnetContextTagBitString).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagBitString for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_ENUMERATED: // BACnetContextTagEnumerated
		if _child, err = new(_BACnetContextTagEnumerated).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagEnumerated for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_DATE: // BACnetContextTagDate
		if _child, err = new(_BACnetContextTagDate).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagDate for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_TIME: // BACnetContextTagTime
		if _child, err = new(_BACnetContextTagTime).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagTime for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetContextTagObjectIdentifier
		if _child, err = new(_BACnetContextTagObjectIdentifier).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagObjectIdentifier for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_UNKNOWN: // BACnetContextTagUnknown
		if _child, err = new(_BACnetContextTagUnknown).parse(ctx, readBuffer, m, actualLength, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagUnknown for type-switch of BACnetContextTag")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [dataType=%v]", dataType)
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTag")
	}

	return _child, nil
}

func (pm *_BACnetContextTag) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetContextTag, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetContextTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetContextTag")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}
	// Virtual field
	tagNumber := m.GetTagNumber()
	_ = tagNumber
	if _tagNumberErr := writeBuffer.WriteVirtual(ctx, "tagNumber", m.GetTagNumber()); _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}
	// Virtual field
	actualLength := m.GetActualLength()
	_ = actualLength
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetContextTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetContextTag")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetContextTag) GetTagNumberArgument() uint8 {
	return m.TagNumberArgument
}

//
////

func (m *_BACnetContextTag) IsBACnetContextTag() {}

func (m *_BACnetContextTag) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTag) deepCopy() *_BACnetContextTag {
	if m == nil {
		return nil
	}
	_BACnetContextTagCopy := &_BACnetContextTag{
		nil, // will be set by child
		m.Header.DeepCopy().(BACnetTagHeader),
		m.TagNumberArgument,
	}
	return _BACnetContextTagCopy
}
