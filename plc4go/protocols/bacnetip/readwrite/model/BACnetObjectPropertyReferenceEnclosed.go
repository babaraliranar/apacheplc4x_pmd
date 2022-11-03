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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetObjectPropertyReferenceEnclosed is the corresponding interface of BACnetObjectPropertyReferenceEnclosed
type BACnetObjectPropertyReferenceEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetObjectPropertyReference returns ObjectPropertyReference (property field)
	GetObjectPropertyReference() BACnetObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetObjectPropertyReferenceEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetObjectPropertyReferenceEnclosed.
// This is useful for switch cases.
type BACnetObjectPropertyReferenceEnclosedExactly interface {
	BACnetObjectPropertyReferenceEnclosed
	isBACnetObjectPropertyReferenceEnclosed() bool
}

// _BACnetObjectPropertyReferenceEnclosed is the data-structure of this message
type _BACnetObjectPropertyReferenceEnclosed struct {
	OpeningTag              BACnetOpeningTag
	ObjectPropertyReference BACnetObjectPropertyReference
	ClosingTag              BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetObjectPropertyReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetObjectPropertyReference() BACnetObjectPropertyReference {
	return m.ObjectPropertyReference
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetObjectPropertyReferenceEnclosed factory function for _BACnetObjectPropertyReferenceEnclosed
func NewBACnetObjectPropertyReferenceEnclosed(openingTag BACnetOpeningTag, objectPropertyReference BACnetObjectPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetObjectPropertyReferenceEnclosed {
	return &_BACnetObjectPropertyReferenceEnclosed{OpeningTag: openingTag, ObjectPropertyReference: objectPropertyReference, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetObjectPropertyReferenceEnclosed(structType interface{}) BACnetObjectPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetObjectPropertyReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetObjectPropertyReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetObjectPropertyReferenceEnclosed"
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (objectPropertyReference)
	lengthInBits += m.ObjectPropertyReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetObjectPropertyReferenceEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetObjectPropertyReferenceEnclosed, error) {
	return BACnetObjectPropertyReferenceEnclosedParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber) // TODO: get endianness from mspec
}

func BACnetObjectPropertyReferenceEnclosedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetObjectPropertyReferenceEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetObjectPropertyReferenceEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (objectPropertyReference)
	if pullErr := readBuffer.PullContext("objectPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectPropertyReference")
	}
	_objectPropertyReference, _objectPropertyReferenceErr := BACnetObjectPropertyReferenceParseWithBuffer(readBuffer)
	if _objectPropertyReferenceErr != nil {
		return nil, errors.Wrap(_objectPropertyReferenceErr, "Error parsing 'objectPropertyReference' field of BACnetObjectPropertyReferenceEnclosed")
	}
	objectPropertyReference := _objectPropertyReference.(BACnetObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("objectPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectPropertyReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetObjectPropertyReferenceEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetObjectPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectPropertyReferenceEnclosed")
	}

	// Create the instance
	return &_BACnetObjectPropertyReferenceEnclosed{
		TagNumber:               tagNumber,
		OpeningTag:              openingTag,
		ObjectPropertyReference: objectPropertyReference,
		ClosingTag:              closingTag,
	}, nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetObjectPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectPropertyReferenceEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (objectPropertyReference)
	if pushErr := writeBuffer.PushContext("objectPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectPropertyReference")
	}
	_objectPropertyReferenceErr := writeBuffer.WriteSerializable(m.GetObjectPropertyReference())
	if popErr := writeBuffer.PopContext("objectPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectPropertyReference")
	}
	if _objectPropertyReferenceErr != nil {
		return errors.Wrap(_objectPropertyReferenceErr, "Error serializing 'objectPropertyReference' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetObjectPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectPropertyReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetObjectPropertyReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetObjectPropertyReferenceEnclosed) isBACnetObjectPropertyReferenceEnclosed() bool {
	return true
}

func (m *_BACnetObjectPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
