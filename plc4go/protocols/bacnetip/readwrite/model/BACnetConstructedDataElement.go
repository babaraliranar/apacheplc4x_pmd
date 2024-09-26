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

// BACnetConstructedDataElement is the corresponding interface of BACnetConstructedDataElement
type BACnetConstructedDataElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetApplicationTag returns ApplicationTag (property field)
	GetApplicationTag() BACnetApplicationTag
	// GetContextTag returns ContextTag (property field)
	GetContextTag() BACnetContextTag
	// GetConstructedData returns ConstructedData (property field)
	GetConstructedData() BACnetConstructedData
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetIsApplicationTag returns IsApplicationTag (virtual field)
	GetIsApplicationTag() bool
	// GetIsConstructedData returns IsConstructedData (virtual field)
	GetIsConstructedData() bool
	// GetIsContextTag returns IsContextTag (virtual field)
	GetIsContextTag() bool
	// IsBACnetConstructedDataElement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataElement()
}

// _BACnetConstructedDataElement is the data-structure of this message
type _BACnetConstructedDataElement struct {
	PeekedTagHeader BACnetTagHeader
	ApplicationTag  BACnetApplicationTag
	ContextTag      BACnetContextTag
	ConstructedData BACnetConstructedData

	// Arguments.
	ObjectTypeArgument         BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
	ArrayIndexArgument         BACnetTagPayloadUnsignedInteger
}

var _ BACnetConstructedDataElement = (*_BACnetConstructedDataElement)(nil)

// NewBACnetConstructedDataElement factory function for _BACnetConstructedDataElement
func NewBACnetConstructedDataElement(peekedTagHeader BACnetTagHeader, applicationTag BACnetApplicationTag, contextTag BACnetContextTag, constructedData BACnetConstructedData, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataElement {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetConstructedDataElement must not be nil")
	}
	return &_BACnetConstructedDataElement{PeekedTagHeader: peekedTagHeader, ApplicationTag: applicationTag, ContextTag: contextTag, ConstructedData: constructedData, ObjectTypeArgument: objectTypeArgument, PropertyIdentifierArgument: propertyIdentifierArgument, ArrayIndexArgument: arrayIndexArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataElement) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConstructedDataElement) GetApplicationTag() BACnetApplicationTag {
	return m.ApplicationTag
}

func (m *_BACnetConstructedDataElement) GetContextTag() BACnetContextTag {
	return m.ContextTag
}

func (m *_BACnetConstructedDataElement) GetConstructedData() BACnetConstructedData {
	return m.ConstructedData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataElement) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.GetApplicationTag()
	_ = applicationTag
	contextTag := m.GetContextTag()
	_ = contextTag
	constructedData := m.GetConstructedData()
	_ = constructedData
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetConstructedDataElement) GetIsApplicationTag() bool {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.GetApplicationTag()
	_ = applicationTag
	contextTag := m.GetContextTag()
	_ = contextTag
	constructedData := m.GetConstructedData()
	_ = constructedData
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_APPLICATION_TAGS)))
}

func (m *_BACnetConstructedDataElement) GetIsConstructedData() bool {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.GetApplicationTag()
	_ = applicationTag
	contextTag := m.GetContextTag()
	_ = contextTag
	constructedData := m.GetConstructedData()
	_ = constructedData
	return bool(bool(!(m.GetIsApplicationTag())) && bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x6))))
}

func (m *_BACnetConstructedDataElement) GetIsContextTag() bool {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.GetApplicationTag()
	_ = applicationTag
	contextTag := m.GetContextTag()
	_ = contextTag
	constructedData := m.GetConstructedData()
	_ = constructedData
	return bool(bool(!(m.GetIsConstructedData())) && bool(!(m.GetIsApplicationTag())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataElement(structType any) BACnetConstructedDataElement {
	if casted, ok := structType.(BACnetConstructedDataElement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataElement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataElement) GetTypeName() string {
	return "BACnetConstructedDataElement"
}

func (m *_BACnetConstructedDataElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (applicationTag)
	if m.ApplicationTag != nil {
		lengthInBits += m.ApplicationTag.GetLengthInBits(ctx)
	}

	// Optional Field (contextTag)
	if m.ContextTag != nil {
		lengthInBits += m.ContextTag.GetLengthInBits(ctx)
	}

	// Optional Field (constructedData)
	if m.ConstructedData != nil {
		lengthInBits += m.ConstructedData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataElementParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElement, error) {
	return BACnetConstructedDataElementParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataElementParseWithBufferProducer(objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataElement, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataElement, error) {
		return BACnetConstructedDataElementParseWithBuffer(ctx, readBuffer, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataElementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElement, error) {
	v, err := (&_BACnetConstructedDataElement{ObjectTypeArgument: objectTypeArgument, PropertyIdentifierArgument: propertyIdentifierArgument, ArrayIndexArgument: arrayIndexArgument}).parse(ctx, readBuffer, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetConstructedDataElement) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataElement BACnetConstructedDataElement, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	isApplicationTag, err := ReadVirtualField[bool](ctx, "isApplicationTag", (*bool)(nil), bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isApplicationTag' field"))
	}
	_ = isApplicationTag

	isConstructedData, err := ReadVirtualField[bool](ctx, "isConstructedData", (*bool)(nil), bool(!(isApplicationTag)) && bool(bool((peekedTagHeader.GetLengthValueType()) == (0x6))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isConstructedData' field"))
	}
	_ = isConstructedData

	isContextTag, err := ReadVirtualField[bool](ctx, "isContextTag", (*bool)(nil), bool(!(isConstructedData)) && bool(!(isApplicationTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isContextTag' field"))
	}
	_ = isContextTag

	// Validation
	if !(bool(!(isContextTag)) || bool((bool(isContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unexpected closing tag"})
	}

	var applicationTag BACnetApplicationTag
	_applicationTag, err := ReadOptionalField[BACnetApplicationTag](ctx, "applicationTag", ReadComplex[BACnetApplicationTag](BACnetApplicationTagParseWithBuffer, readBuffer), isApplicationTag)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationTag' field"))
	}
	if _applicationTag != nil {
		applicationTag = *_applicationTag
		m.ApplicationTag = applicationTag
	}

	var contextTag BACnetContextTag
	_contextTag, err := ReadOptionalField[BACnetContextTag](ctx, "contextTag", ReadComplex[BACnetContextTag](BACnetContextTagParseWithBufferProducer[BACnetContextTag]((uint8)(peekedTagNumber), (BACnetDataType)(BACnetDataType_UNKNOWN)), readBuffer), isContextTag)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'contextTag' field"))
	}
	if _contextTag != nil {
		contextTag = *_contextTag
		m.ContextTag = contextTag
	}

	var constructedData BACnetConstructedData
	_constructedData, err := ReadOptionalField[BACnetConstructedData](ctx, "constructedData", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(peekedTagNumber), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(propertyIdentifierArgument), (BACnetTagPayloadUnsignedInteger)(arrayIndexArgument)), readBuffer), isConstructedData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'constructedData' field"))
	}
	if _constructedData != nil {
		constructedData = *_constructedData
		m.ConstructedData = constructedData
	}

	// Validation
	if !(bool(bool((bool(isApplicationTag) && bool(bool((applicationTag) != (nil))))) || bool((bool(isContextTag) && bool(bool((contextTag) != (nil)))))) || bool((bool(isConstructedData) && bool(bool((constructedData) != (nil)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "BACnetConstructedDataElement could not parse anything"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataElement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConstructedDataElement"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataElement")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	isApplicationTag := m.GetIsApplicationTag()
	_ = isApplicationTag
	if _isApplicationTagErr := writeBuffer.WriteVirtual(ctx, "isApplicationTag", m.GetIsApplicationTag()); _isApplicationTagErr != nil {
		return errors.Wrap(_isApplicationTagErr, "Error serializing 'isApplicationTag' field")
	}
	// Virtual field
	isConstructedData := m.GetIsConstructedData()
	_ = isConstructedData
	if _isConstructedDataErr := writeBuffer.WriteVirtual(ctx, "isConstructedData", m.GetIsConstructedData()); _isConstructedDataErr != nil {
		return errors.Wrap(_isConstructedDataErr, "Error serializing 'isConstructedData' field")
	}
	// Virtual field
	isContextTag := m.GetIsContextTag()
	_ = isContextTag
	if _isContextTagErr := writeBuffer.WriteVirtual(ctx, "isContextTag", m.GetIsContextTag()); _isContextTagErr != nil {
		return errors.Wrap(_isContextTagErr, "Error serializing 'isContextTag' field")
	}

	if err := WriteOptionalField[BACnetApplicationTag](ctx, "applicationTag", GetRef(m.GetApplicationTag()), WriteComplex[BACnetApplicationTag](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'applicationTag' field")
	}

	if err := WriteOptionalField[BACnetContextTag](ctx, "contextTag", GetRef(m.GetContextTag()), WriteComplex[BACnetContextTag](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'contextTag' field")
	}

	if err := WriteOptionalField[BACnetConstructedData](ctx, "constructedData", GetRef(m.GetConstructedData()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'constructedData' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConstructedDataElement"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConstructedDataElement")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConstructedDataElement) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}
func (m *_BACnetConstructedDataElement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return m.PropertyIdentifierArgument
}
func (m *_BACnetConstructedDataElement) GetArrayIndexArgument() BACnetTagPayloadUnsignedInteger {
	return m.ArrayIndexArgument
}

//
////

func (m *_BACnetConstructedDataElement) IsBACnetConstructedDataElement() {}

func (m *_BACnetConstructedDataElement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataElement) deepCopy() *_BACnetConstructedDataElement {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataElementCopy := &_BACnetConstructedDataElement{
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.ApplicationTag.DeepCopy().(BACnetApplicationTag),
		m.ContextTag.DeepCopy().(BACnetContextTag),
		m.ConstructedData.DeepCopy().(BACnetConstructedData),
		m.ObjectTypeArgument,
		m.PropertyIdentifierArgument,
		m.ArrayIndexArgument,
	}
	return _BACnetConstructedDataElementCopy
}

func (m *_BACnetConstructedDataElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
