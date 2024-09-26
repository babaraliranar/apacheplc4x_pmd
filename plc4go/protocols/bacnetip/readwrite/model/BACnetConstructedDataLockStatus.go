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

// BACnetConstructedDataLockStatus is the corresponding interface of BACnetConstructedDataLockStatus
type BACnetConstructedDataLockStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLockStatus returns LockStatus (property field)
	GetLockStatus() BACnetLockStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLockStatusTagged
	// IsBACnetConstructedDataLockStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLockStatus()
}

// _BACnetConstructedDataLockStatus is the data-structure of this message
type _BACnetConstructedDataLockStatus struct {
	BACnetConstructedDataContract
	LockStatus BACnetLockStatusTagged
}

var _ BACnetConstructedDataLockStatus = (*_BACnetConstructedDataLockStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLockStatus)(nil)

// NewBACnetConstructedDataLockStatus factory function for _BACnetConstructedDataLockStatus
func NewBACnetConstructedDataLockStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lockStatus BACnetLockStatusTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLockStatus {
	if lockStatus == nil {
		panic("lockStatus of type BACnetLockStatusTagged for BACnetConstructedDataLockStatus must not be nil")
	}
	_result := &_BACnetConstructedDataLockStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LockStatus:                    lockStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLockStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLockStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCK_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLockStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLockStatus) GetLockStatus() BACnetLockStatusTagged {
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

func (m *_BACnetConstructedDataLockStatus) GetActualValue() BACnetLockStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLockStatusTagged(m.GetLockStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLockStatus(structType any) BACnetConstructedDataLockStatus {
	if casted, ok := structType.(BACnetConstructedDataLockStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLockStatus) GetTypeName() string {
	return "BACnetConstructedDataLockStatus"
}

func (m *_BACnetConstructedDataLockStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lockStatus)
	lengthInBits += m.LockStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLockStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLockStatus BACnetConstructedDataLockStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lockStatus, err := ReadSimpleField[BACnetLockStatusTagged](ctx, "lockStatus", ReadComplex[BACnetLockStatusTagged](BACnetLockStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lockStatus' field"))
	}
	m.LockStatus = lockStatus

	actualValue, err := ReadVirtualField[BACnetLockStatusTagged](ctx, "actualValue", (*BACnetLockStatusTagged)(nil), lockStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLockStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLockStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockStatus")
		}

		if err := WriteSimpleField[BACnetLockStatusTagged](ctx, "lockStatus", m.GetLockStatus(), WriteComplex[BACnetLockStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lockStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLockStatus) IsBACnetConstructedDataLockStatus() {}

func (m *_BACnetConstructedDataLockStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLockStatus) deepCopy() *_BACnetConstructedDataLockStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLockStatusCopy := &_BACnetConstructedDataLockStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LockStatus.DeepCopy().(BACnetLockStatusTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLockStatusCopy
}

func (m *_BACnetConstructedDataLockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
