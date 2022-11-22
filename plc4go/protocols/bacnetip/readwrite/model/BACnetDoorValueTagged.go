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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDoorValueTagged is the corresponding interface of BACnetDoorValueTagged
type BACnetDoorValueTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetDoorValue
}

// BACnetDoorValueTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetDoorValueTagged.
// This is useful for switch cases.
type BACnetDoorValueTaggedExactly interface {
	BACnetDoorValueTagged
	isBACnetDoorValueTagged() bool
}

// _BACnetDoorValueTagged is the data-structure of this message
type _BACnetDoorValueTagged struct {
	Header BACnetTagHeader
	Value  BACnetDoorValue

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDoorValueTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetDoorValueTagged) GetValue() BACnetDoorValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDoorValueTagged factory function for _BACnetDoorValueTagged
func NewBACnetDoorValueTagged(header BACnetTagHeader, value BACnetDoorValue, tagNumber uint8, tagClass TagClass) *_BACnetDoorValueTagged {
	return &_BACnetDoorValueTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetDoorValueTagged(structType interface{}) BACnetDoorValueTagged {
	if casted, ok := structType.(BACnetDoorValueTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDoorValueTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDoorValueTagged) GetTypeName() string {
	return "BACnetDoorValueTagged"
}

func (m *_BACnetDoorValueTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetDoorValueTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetDoorValueTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDoorValueTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetDoorValueTagged, error) {
	return BACnetDoorValueTaggedParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetDoorValueTaggedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetDoorValueTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDoorValueTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDoorValueTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetDoorValueTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetDoorValue_LOCK)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetDoorValueTagged")
	}
	var value BACnetDoorValue
	if _value != nil {
		value = _value.(BACnetDoorValue)
	}

	if closeErr := readBuffer.CloseContext("BACnetDoorValueTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDoorValueTagged")
	}

	// Create the instance
	return &_BACnetDoorValueTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetDoorValueTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDoorValueTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDoorValueTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDoorValueTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDoorValueTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDoorValueTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDoorValueTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetDoorValueTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetDoorValueTagged) isBACnetDoorValueTagged() bool {
	return true
}

func (m *_BACnetDoorValueTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
