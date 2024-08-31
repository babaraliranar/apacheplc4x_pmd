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

// BACnetConstructedDataMachineRoomID is the corresponding interface of BACnetConstructedDataMachineRoomID
type BACnetConstructedDataMachineRoomID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMachineRoomId returns MachineRoomId (property field)
	GetMachineRoomId() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
}

// BACnetConstructedDataMachineRoomIDExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMachineRoomID.
// This is useful for switch cases.
type BACnetConstructedDataMachineRoomIDExactly interface {
	BACnetConstructedDataMachineRoomID
	isBACnetConstructedDataMachineRoomID() bool
}

// _BACnetConstructedDataMachineRoomID is the data-structure of this message
type _BACnetConstructedDataMachineRoomID struct {
	*_BACnetConstructedData
	MachineRoomId BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMachineRoomID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MACHINE_ROOM_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMachineRoomID) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMachineRoomID) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetMachineRoomId() BACnetApplicationTagObjectIdentifier {
	return m.MachineRoomId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetMachineRoomId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMachineRoomID factory function for _BACnetConstructedDataMachineRoomID
func NewBACnetConstructedDataMachineRoomID(machineRoomId BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMachineRoomID {
	_result := &_BACnetConstructedDataMachineRoomID{
		MachineRoomId:          machineRoomId,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMachineRoomID(structType any) BACnetConstructedDataMachineRoomID {
	if casted, ok := structType.(BACnetConstructedDataMachineRoomID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMachineRoomID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMachineRoomID) GetTypeName() string {
	return "BACnetConstructedDataMachineRoomID"
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (machineRoomId)
	lengthInBits += m.MachineRoomId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataMachineRoomIDParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMachineRoomID, error) {
	return BACnetConstructedDataMachineRoomIDParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataMachineRoomIDParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataMachineRoomID, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataMachineRoomID, error) {
		return BACnetConstructedDataMachineRoomIDParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataMachineRoomIDParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMachineRoomID, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMachineRoomID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMachineRoomID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	machineRoomId, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "machineRoomId", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'machineRoomId' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagObjectIdentifier](ctx, "actualValue", (*BACnetApplicationTagObjectIdentifier)(nil), machineRoomId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMachineRoomID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMachineRoomID")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMachineRoomID{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MachineRoomId: machineRoomId,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMachineRoomID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMachineRoomID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMachineRoomID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMachineRoomID")
		}

		// Simple Field (machineRoomId)
		if pushErr := writeBuffer.PushContext("machineRoomId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for machineRoomId")
		}
		_machineRoomIdErr := writeBuffer.WriteSerializable(ctx, m.GetMachineRoomId())
		if popErr := writeBuffer.PopContext("machineRoomId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for machineRoomId")
		}
		if _machineRoomIdErr != nil {
			return errors.Wrap(_machineRoomIdErr, "Error serializing 'machineRoomId' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMachineRoomID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMachineRoomID")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMachineRoomID) isBACnetConstructedDataMachineRoomID() bool {
	return true
}

func (m *_BACnetConstructedDataMachineRoomID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
