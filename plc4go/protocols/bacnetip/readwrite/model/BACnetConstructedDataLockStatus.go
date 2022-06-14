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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataLockStatus is the data-structure of this message
type BACnetConstructedDataLockStatus struct {
	*BACnetConstructedData
	LockStatus *BACnetLockStatusTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLockStatus is the corresponding interface of BACnetConstructedDataLockStatus
type IBACnetConstructedDataLockStatus interface {
	IBACnetConstructedData
	// GetLockStatus returns LockStatus (property field)
	GetLockStatus() *BACnetLockStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetLockStatusTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataLockStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLockStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCK_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLockStatus) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLockStatus) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLockStatus) GetLockStatus() *BACnetLockStatusTagged {
	return m.LockStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLockStatus) GetActualValue() *BACnetLockStatusTagged {
	return CastBACnetLockStatusTagged(m.GetLockStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLockStatus factory function for BACnetConstructedDataLockStatus
func NewBACnetConstructedDataLockStatus(lockStatus *BACnetLockStatusTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLockStatus {
	_result := &BACnetConstructedDataLockStatus{
		LockStatus:            lockStatus,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLockStatus(structType interface{}) *BACnetConstructedDataLockStatus {
	if casted, ok := structType.(BACnetConstructedDataLockStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLockStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLockStatus(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLockStatus) GetTypeName() string {
	return "BACnetConstructedDataLockStatus"
}

func (m *BACnetConstructedDataLockStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLockStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lockStatus)
	lengthInBits += m.LockStatus.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLockStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLockStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLockStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lockStatus)
	if pullErr := readBuffer.PullContext("lockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lockStatus")
	}
	_lockStatus, _lockStatusErr := BACnetLockStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _lockStatusErr != nil {
		return nil, errors.Wrap(_lockStatusErr, "Error parsing 'lockStatus' field")
	}
	lockStatus := CastBACnetLockStatusTagged(_lockStatus)
	if closeErr := readBuffer.CloseContext("lockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lockStatus")
	}

	// Virtual field
	_actualValue := lockStatus
	actualValue := CastBACnetLockStatusTagged(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockStatus")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLockStatus{
		LockStatus:            CastBACnetLockStatusTagged(lockStatus),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLockStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockStatus")
		}

		// Simple Field (lockStatus)
		if pushErr := writeBuffer.PushContext("lockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lockStatus")
		}
		_lockStatusErr := writeBuffer.WriteSerializable(m.LockStatus)
		if popErr := writeBuffer.PopContext("lockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lockStatus")
		}
		if _lockStatusErr != nil {
			return errors.Wrap(_lockStatusErr, "Error serializing 'lockStatus' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
