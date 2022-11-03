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

// BACnetPriorityValueNull is the corresponding interface of BACnetPriorityValueNull
type BACnetPriorityValueNull interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetPriorityValueNullExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueNull.
// This is useful for switch cases.
type BACnetPriorityValueNullExactly interface {
	BACnetPriorityValueNull
	isBACnetPriorityValueNull() bool
}

// _BACnetPriorityValueNull is the data-structure of this message
type _BACnetPriorityValueNull struct {
	*_BACnetPriorityValue
	NullValue BACnetApplicationTagNull
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueNull) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueNull) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueNull factory function for _BACnetPriorityValueNull
func NewBACnetPriorityValueNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueNull {
	_result := &_BACnetPriorityValueNull{
		NullValue:            nullValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueNull(structType interface{}) BACnetPriorityValueNull {
	if casted, ok := structType.(BACnetPriorityValueNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueNull) GetTypeName() string {
	return "BACnetPriorityValueNull"
}

func (m *_BACnetPriorityValueNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueNullParse(theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPriorityValueNull, error) {
	return BACnetPriorityValueNullParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), objectTypeArgument) // TODO: get endianness from mspec
}

func BACnetPriorityValueNullParseWithBuffer(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field of BACnetPriorityValueNull")
	}
	nullValue := _nullValue.(BACnetApplicationTagNull)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueNull{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		NullValue: nullValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueNull) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueNull")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := writeBuffer.WriteSerializable(m.GetNullValue())
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueNull) isBACnetPriorityValueNull() bool {
	return true
}

func (m *_BACnetPriorityValueNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
