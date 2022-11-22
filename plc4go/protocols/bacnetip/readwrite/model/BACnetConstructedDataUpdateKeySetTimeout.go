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

// BACnetConstructedDataUpdateKeySetTimeout is the corresponding interface of BACnetConstructedDataUpdateKeySetTimeout
type BACnetConstructedDataUpdateKeySetTimeout interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUpdateKeySetTimeout returns UpdateKeySetTimeout (property field)
	GetUpdateKeySetTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataUpdateKeySetTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUpdateKeySetTimeout.
// This is useful for switch cases.
type BACnetConstructedDataUpdateKeySetTimeoutExactly interface {
	BACnetConstructedDataUpdateKeySetTimeout
	isBACnetConstructedDataUpdateKeySetTimeout() bool
}

// _BACnetConstructedDataUpdateKeySetTimeout is the data-structure of this message
type _BACnetConstructedDataUpdateKeySetTimeout struct {
	*_BACnetConstructedData
	UpdateKeySetTimeout BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_KEY_SET_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetUpdateKeySetTimeout() BACnetApplicationTagUnsignedInteger {
	return m.UpdateKeySetTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetUpdateKeySetTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUpdateKeySetTimeout factory function for _BACnetConstructedDataUpdateKeySetTimeout
func NewBACnetConstructedDataUpdateKeySetTimeout(updateKeySetTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUpdateKeySetTimeout {
	_result := &_BACnetConstructedDataUpdateKeySetTimeout{
		UpdateKeySetTimeout:    updateKeySetTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUpdateKeySetTimeout(structType interface{}) BACnetConstructedDataUpdateKeySetTimeout {
	if casted, ok := structType.(BACnetConstructedDataUpdateKeySetTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUpdateKeySetTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetTypeName() string {
	return "BACnetConstructedDataUpdateKeySetTimeout"
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (updateKeySetTimeout)
	lengthInBits += m.UpdateKeySetTimeout.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUpdateKeySetTimeoutParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUpdateKeySetTimeout, error) {
	return BACnetConstructedDataUpdateKeySetTimeoutParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataUpdateKeySetTimeoutParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUpdateKeySetTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUpdateKeySetTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUpdateKeySetTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (updateKeySetTimeout)
	if pullErr := readBuffer.PullContext("updateKeySetTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateKeySetTimeout")
	}
	_updateKeySetTimeout, _updateKeySetTimeoutErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _updateKeySetTimeoutErr != nil {
		return nil, errors.Wrap(_updateKeySetTimeoutErr, "Error parsing 'updateKeySetTimeout' field of BACnetConstructedDataUpdateKeySetTimeout")
	}
	updateKeySetTimeout := _updateKeySetTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("updateKeySetTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateKeySetTimeout")
	}

	// Virtual field
	_actualValue := updateKeySetTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUpdateKeySetTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUpdateKeySetTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUpdateKeySetTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UpdateKeySetTimeout: updateKeySetTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUpdateKeySetTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUpdateKeySetTimeout")
		}

		// Simple Field (updateKeySetTimeout)
		if pushErr := writeBuffer.PushContext("updateKeySetTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for updateKeySetTimeout")
		}
		_updateKeySetTimeoutErr := writeBuffer.WriteSerializable(m.GetUpdateKeySetTimeout())
		if popErr := writeBuffer.PopContext("updateKeySetTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for updateKeySetTimeout")
		}
		if _updateKeySetTimeoutErr != nil {
			return errors.Wrap(_updateKeySetTimeoutErr, "Error serializing 'updateKeySetTimeout' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUpdateKeySetTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUpdateKeySetTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) isBACnetConstructedDataUpdateKeySetTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
