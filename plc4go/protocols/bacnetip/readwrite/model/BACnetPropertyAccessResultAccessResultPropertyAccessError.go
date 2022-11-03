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

// BACnetPropertyAccessResultAccessResultPropertyAccessError is the corresponding interface of BACnetPropertyAccessResultAccessResultPropertyAccessError
type BACnetPropertyAccessResultAccessResultPropertyAccessError interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyAccessResultAccessResult
	// GetPropertyAccessError returns PropertyAccessError (property field)
	GetPropertyAccessError() ErrorEnclosed
}

// BACnetPropertyAccessResultAccessResultPropertyAccessErrorExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyAccessResultAccessResultPropertyAccessError.
// This is useful for switch cases.
type BACnetPropertyAccessResultAccessResultPropertyAccessErrorExactly interface {
	BACnetPropertyAccessResultAccessResultPropertyAccessError
	isBACnetPropertyAccessResultAccessResultPropertyAccessError() bool
}

// _BACnetPropertyAccessResultAccessResultPropertyAccessError is the data-structure of this message
type _BACnetPropertyAccessResultAccessResultPropertyAccessError struct {
	*_BACnetPropertyAccessResultAccessResult
	PropertyAccessError ErrorEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) InitializeParent(parent BACnetPropertyAccessResultAccessResult, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetParent() BACnetPropertyAccessResultAccessResult {
	return m._BACnetPropertyAccessResultAccessResult
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetPropertyAccessError() ErrorEnclosed {
	return m.PropertyAccessError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyAccessResultAccessResultPropertyAccessError factory function for _BACnetPropertyAccessResultAccessResultPropertyAccessError
func NewBACnetPropertyAccessResultAccessResultPropertyAccessError(propertyAccessError ErrorEnclosed, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPropertyAccessResultAccessResultPropertyAccessError {
	_result := &_BACnetPropertyAccessResultAccessResultPropertyAccessError{
		PropertyAccessError:                     propertyAccessError,
		_BACnetPropertyAccessResultAccessResult: NewBACnetPropertyAccessResultAccessResult(peekedTagHeader, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument),
	}
	_result._BACnetPropertyAccessResultAccessResult._BACnetPropertyAccessResultAccessResultChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResultAccessResultPropertyAccessError(structType interface{}) BACnetPropertyAccessResultAccessResultPropertyAccessError {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResultPropertyAccessError); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResultPropertyAccessError); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResultPropertyAccessError"
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (propertyAccessError)
	lengthInBits += m.PropertyAccessError.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyAccessResultAccessResultPropertyAccessErrorParse(theBytes []byte, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPropertyAccessResultAccessResultPropertyAccessError, error) {
	return BACnetPropertyAccessResultAccessResultPropertyAccessErrorParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetPropertyAccessResultAccessResultPropertyAccessErrorParseWithBuffer(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPropertyAccessResultAccessResultPropertyAccessError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResultPropertyAccessError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyAccessError)
	if pullErr := readBuffer.PullContext("propertyAccessError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyAccessError")
	}
	_propertyAccessError, _propertyAccessErrorErr := ErrorEnclosedParseWithBuffer(readBuffer, uint8(uint8(5)))
	if _propertyAccessErrorErr != nil {
		return nil, errors.Wrap(_propertyAccessErrorErr, "Error parsing 'propertyAccessError' field of BACnetPropertyAccessResultAccessResultPropertyAccessError")
	}
	propertyAccessError := _propertyAccessError.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("propertyAccessError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyAccessError")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResultPropertyAccessError")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyAccessResultAccessResultPropertyAccessError{
		_BACnetPropertyAccessResultAccessResult: &_BACnetPropertyAccessResultAccessResult{
			ObjectTypeArgument:         objectTypeArgument,
			PropertyIdentifierArgument: propertyIdentifierArgument,
			PropertyArrayIndexArgument: propertyArrayIndexArgument,
		},
		PropertyAccessError: propertyAccessError,
	}
	_child._BACnetPropertyAccessResultAccessResult._BACnetPropertyAccessResultAccessResultChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResultPropertyAccessError")
		}

		// Simple Field (propertyAccessError)
		if pushErr := writeBuffer.PushContext("propertyAccessError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyAccessError")
		}
		_propertyAccessErrorErr := writeBuffer.WriteSerializable(m.GetPropertyAccessError())
		if popErr := writeBuffer.PopContext("propertyAccessError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyAccessError")
		}
		if _propertyAccessErrorErr != nil {
			return errors.Wrap(_propertyAccessErrorErr, "Error serializing 'propertyAccessError' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResultPropertyAccessError")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) isBACnetPropertyAccessResultAccessResultPropertyAccessError() bool {
	return true
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
