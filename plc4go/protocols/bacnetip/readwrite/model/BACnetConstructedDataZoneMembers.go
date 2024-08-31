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

// BACnetConstructedDataZoneMembers is the corresponding interface of BACnetConstructedDataZoneMembers
type BACnetConstructedDataZoneMembers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMembers returns Members (property field)
	GetMembers() []BACnetDeviceObjectReference
}

// BACnetConstructedDataZoneMembersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataZoneMembers.
// This is useful for switch cases.
type BACnetConstructedDataZoneMembersExactly interface {
	BACnetConstructedDataZoneMembers
	isBACnetConstructedDataZoneMembers() bool
}

// _BACnetConstructedDataZoneMembers is the data-structure of this message
type _BACnetConstructedDataZoneMembers struct {
	*_BACnetConstructedData
	Members []BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataZoneMembers) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataZoneMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataZoneMembers) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataZoneMembers) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataZoneMembers) GetMembers() []BACnetDeviceObjectReference {
	return m.Members
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataZoneMembers factory function for _BACnetConstructedDataZoneMembers
func NewBACnetConstructedDataZoneMembers(members []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataZoneMembers {
	_result := &_BACnetConstructedDataZoneMembers{
		Members:                members,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataZoneMembers(structType any) BACnetConstructedDataZoneMembers {
	if casted, ok := structType.(BACnetConstructedDataZoneMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataZoneMembers) GetTypeName() string {
	return "BACnetConstructedDataZoneMembers"
}

func (m *_BACnetConstructedDataZoneMembers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Members) > 0 {
		for _, element := range m.Members {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataZoneMembers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataZoneMembersParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataZoneMembers, error) {
	return BACnetConstructedDataZoneMembersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataZoneMembersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataZoneMembers, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	members, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "members", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'members' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneMembers")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataZoneMembers{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Members: members,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataZoneMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataZoneMembers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneMembers")
		}

		// Array Field (members)
		if pushErr := writeBuffer.PushContext("members", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for members")
		}
		for _curItem, _element := range m.GetMembers() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMembers()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'members' field")
			}
		}
		if popErr := writeBuffer.PopContext("members", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for members")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneMembers")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataZoneMembers) isBACnetConstructedDataZoneMembers() bool {
	return true
}

func (m *_BACnetConstructedDataZoneMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
