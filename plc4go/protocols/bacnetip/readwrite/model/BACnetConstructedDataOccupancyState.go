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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataOccupancyState is the corresponding interface of BACnetConstructedDataOccupancyState
type BACnetConstructedDataOccupancyState interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOccupancyState returns OccupancyState (property field)
	GetOccupancyState() BACnetAccessZoneOccupancyStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessZoneOccupancyStateTagged
}

// BACnetConstructedDataOccupancyStateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOccupancyState.
// This is useful for switch cases.
type BACnetConstructedDataOccupancyStateExactly interface {
	BACnetConstructedDataOccupancyState
	isBACnetConstructedDataOccupancyState() bool
}

// _BACnetConstructedDataOccupancyState is the data-structure of this message
type _BACnetConstructedDataOccupancyState struct {
	*_BACnetConstructedData
	OccupancyState BACnetAccessZoneOccupancyStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyState) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOccupancyState) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetOccupancyState() BACnetAccessZoneOccupancyStateTagged {
	return m.OccupancyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetActualValue() BACnetAccessZoneOccupancyStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessZoneOccupancyStateTagged(m.GetOccupancyState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyState factory function for _BACnetConstructedDataOccupancyState
func NewBACnetConstructedDataOccupancyState(occupancyState BACnetAccessZoneOccupancyStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyState {
	_result := &_BACnetConstructedDataOccupancyState{
		OccupancyState:         occupancyState,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyState(structType interface{}) BACnetConstructedDataOccupancyState {
	if casted, ok := structType.(BACnetConstructedDataOccupancyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyState) GetTypeName() string {
	return "BACnetConstructedDataOccupancyState"
}

func (m *_BACnetConstructedDataOccupancyState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (occupancyState)
	lengthInBits += m.OccupancyState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataOccupancyStateParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyState, error) {
	return BACnetConstructedDataOccupancyStateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataOccupancyStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyState)
	if pullErr := readBuffer.PullContext("occupancyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for occupancyState")
	}
	_occupancyState, _occupancyStateErr := BACnetAccessZoneOccupancyStateTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _occupancyStateErr != nil {
		return nil, errors.Wrap(_occupancyStateErr, "Error parsing 'occupancyState' field of BACnetConstructedDataOccupancyState")
	}
	occupancyState := _occupancyState.(BACnetAccessZoneOccupancyStateTagged)
	if closeErr := readBuffer.CloseContext("occupancyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for occupancyState")
	}

	// Virtual field
	_actualValue := occupancyState
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyState")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOccupancyState{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		OccupancyState: occupancyState,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOccupancyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyState")
		}

		// Simple Field (occupancyState)
		if pushErr := writeBuffer.PushContext("occupancyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for occupancyState")
		}
		_occupancyStateErr := writeBuffer.WriteSerializable(ctx, m.GetOccupancyState())
		if popErr := writeBuffer.PopContext("occupancyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for occupancyState")
		}
		if _occupancyStateErr != nil {
			return errors.Wrap(_occupancyStateErr, "Error serializing 'occupancyState' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyState")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyState) isBACnetConstructedDataOccupancyState() bool {
	return true
}

func (m *_BACnetConstructedDataOccupancyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
