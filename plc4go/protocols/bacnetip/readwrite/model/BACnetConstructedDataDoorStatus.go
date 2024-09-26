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

// BACnetConstructedDataDoorStatus is the corresponding interface of BACnetConstructedDataDoorStatus
type BACnetConstructedDataDoorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() BACnetDoorStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorStatusTagged
	// IsBACnetConstructedDataDoorStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDoorStatus()
}

// _BACnetConstructedDataDoorStatus is the data-structure of this message
type _BACnetConstructedDataDoorStatus struct {
	BACnetConstructedDataContract
	DoorStatus BACnetDoorStatusTagged
}

var _ BACnetConstructedDataDoorStatus = (*_BACnetConstructedDataDoorStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDoorStatus)(nil)

// NewBACnetConstructedDataDoorStatus factory function for _BACnetConstructedDataDoorStatus
func NewBACnetConstructedDataDoorStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, doorStatus BACnetDoorStatusTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorStatus {
	if doorStatus == nil {
		panic("doorStatus of type BACnetDoorStatusTagged for BACnetConstructedDataDoorStatus must not be nil")
	}
	_result := &_BACnetConstructedDataDoorStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DoorStatus:                    doorStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetDoorStatus() BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetActualValue() BACnetDoorStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorStatusTagged(m.GetDoorStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorStatus(structType any) BACnetConstructedDataDoorStatus {
	if casted, ok := structType.(BACnetConstructedDataDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorStatus) GetTypeName() string {
	return "BACnetConstructedDataDoorStatus"
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDoorStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDoorStatus BACnetConstructedDataDoorStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorStatus, err := ReadSimpleField[BACnetDoorStatusTagged](ctx, "doorStatus", ReadComplex[BACnetDoorStatusTagged](BACnetDoorStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorStatus' field"))
	}
	m.DoorStatus = doorStatus

	actualValue, err := ReadVirtualField[BACnetDoorStatusTagged](ctx, "actualValue", (*BACnetDoorStatusTagged)(nil), doorStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDoorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorStatus")
		}

		if err := WriteSimpleField[BACnetDoorStatusTagged](ctx, "doorStatus", m.GetDoorStatus(), WriteComplex[BACnetDoorStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorStatus) IsBACnetConstructedDataDoorStatus() {}

func (m *_BACnetConstructedDataDoorStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDoorStatus) deepCopy() *_BACnetConstructedDataDoorStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDoorStatusCopy := &_BACnetConstructedDataDoorStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DoorStatus.DeepCopy().(BACnetDoorStatusTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDoorStatusCopy
}

func (m *_BACnetConstructedDataDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
