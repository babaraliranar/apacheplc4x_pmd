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

// BACnetTagPayloadObjectIdentifier is the corresponding interface of BACnetTagPayloadObjectIdentifier
type BACnetTagPayloadObjectIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetObjectType returns ObjectType (property field)
	GetObjectType() BACnetObjectType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint16
	// GetInstanceNumber returns InstanceNumber (property field)
	GetInstanceNumber() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetTagPayloadObjectIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadObjectIdentifier()
}

// _BACnetTagPayloadObjectIdentifier is the data-structure of this message
type _BACnetTagPayloadObjectIdentifier struct {
	ObjectType       BACnetObjectType
	ProprietaryValue uint16
	InstanceNumber   uint32
}

var _ BACnetTagPayloadObjectIdentifier = (*_BACnetTagPayloadObjectIdentifier)(nil)

// NewBACnetTagPayloadObjectIdentifier factory function for _BACnetTagPayloadObjectIdentifier
func NewBACnetTagPayloadObjectIdentifier(objectType BACnetObjectType, proprietaryValue uint16, instanceNumber uint32) *_BACnetTagPayloadObjectIdentifier {
	return &_BACnetTagPayloadObjectIdentifier{ObjectType: objectType, ProprietaryValue: proprietaryValue, InstanceNumber: instanceNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadObjectIdentifier) GetObjectType() BACnetObjectType {
	return m.ObjectType
}

func (m *_BACnetTagPayloadObjectIdentifier) GetProprietaryValue() uint16 {
	return m.ProprietaryValue
}

func (m *_BACnetTagPayloadObjectIdentifier) GetInstanceNumber() uint32 {
	return m.InstanceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadObjectIdentifier) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetObjectType()) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadObjectIdentifier(structType any) BACnetTagPayloadObjectIdentifier {
	if casted, ok := structType.(BACnetTagPayloadObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadObjectIdentifier) GetTypeName() string {
	return "BACnetTagPayloadObjectIdentifier"
}

func (m *_BACnetTagPayloadObjectIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (objectType)
	lengthInBits += uint16(int32(10))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	// Simple field (instanceNumber)
	lengthInBits += 22

	return lengthInBits
}

func (m *_BACnetTagPayloadObjectIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadObjectIdentifierParse(ctx context.Context, theBytes []byte) (BACnetTagPayloadObjectIdentifier, error) {
	return BACnetTagPayloadObjectIdentifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadObjectIdentifierParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadObjectIdentifier, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadObjectIdentifier, error) {
		return BACnetTagPayloadObjectIdentifierParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTagPayloadObjectIdentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadObjectIdentifier, error) {
	v, err := (&_BACnetTagPayloadObjectIdentifier{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadObjectIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTagPayloadObjectIdentifier BACnetTagPayloadObjectIdentifier, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectType, err := ReadManualField[BACnetObjectType](ctx, "objectType", readBuffer, EnsureType[BACnetObjectType](ReadObjectType(ctx, readBuffer)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectType' field"))
	}
	m.ObjectType = objectType

	proprietaryValue, err := ReadManualField[uint16](ctx, "proprietaryValue", readBuffer, EnsureType[uint16](ReadProprietaryObjectType(ctx, readBuffer, objectType)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((objectType) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	instanceNumber, err := ReadSimpleField(ctx, "instanceNumber", ReadUnsignedInt(readBuffer, uint8(22)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceNumber' field"))
	}
	m.InstanceNumber = instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadObjectIdentifier")
	}

	return m, nil
}

func (m *_BACnetTagPayloadObjectIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadObjectIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadObjectIdentifier")
	}

	if err := WriteManualField[BACnetObjectType](ctx, "objectType", func(ctx context.Context) error { return WriteObjectType(ctx, writeBuffer, m.GetObjectType()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'objectType' field")
	}

	if err := WriteManualField[uint16](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryObjectType(ctx, writeBuffer, m.GetObjectType(), m.GetProprietaryValue())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteSimpleField[uint32](ctx, "instanceNumber", m.GetInstanceNumber(), WriteUnsignedInt(writeBuffer, 22)); err != nil {
		return errors.Wrap(err, "Error serializing 'instanceNumber' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadObjectIdentifier")
	}
	return nil
}

func (m *_BACnetTagPayloadObjectIdentifier) IsBACnetTagPayloadObjectIdentifier() {}

func (m *_BACnetTagPayloadObjectIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadObjectIdentifier) deepCopy() *_BACnetTagPayloadObjectIdentifier {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadObjectIdentifierCopy := &_BACnetTagPayloadObjectIdentifier{
		m.ObjectType,
		m.ProprietaryValue,
		m.InstanceNumber,
	}
	return _BACnetTagPayloadObjectIdentifierCopy
}

func (m *_BACnetTagPayloadObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
