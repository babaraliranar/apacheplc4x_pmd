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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataCarMovingDirection is the corresponding interface of BACnetConstructedDataCarMovingDirection
type BACnetConstructedDataCarMovingDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCarMovingDirection returns CarMovingDirection (property field)
	GetCarMovingDirection() BACnetLiftCarDirectionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftCarDirectionTagged
}

// BACnetConstructedDataCarMovingDirectionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarMovingDirection.
// This is useful for switch cases.
type BACnetConstructedDataCarMovingDirectionExactly interface {
	BACnetConstructedDataCarMovingDirection
	isBACnetConstructedDataCarMovingDirection() bool
}

// _BACnetConstructedDataCarMovingDirection is the data-structure of this message
type _BACnetConstructedDataCarMovingDirection struct {
	*_BACnetConstructedData
	CarMovingDirection BACnetLiftCarDirectionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarMovingDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_MOVING_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarMovingDirection) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetCarMovingDirection() BACnetLiftCarDirectionTagged {
	return m.CarMovingDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetActualValue() BACnetLiftCarDirectionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLiftCarDirectionTagged(m.GetCarMovingDirection())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarMovingDirection factory function for _BACnetConstructedDataCarMovingDirection
func NewBACnetConstructedDataCarMovingDirection(carMovingDirection BACnetLiftCarDirectionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarMovingDirection {
	_result := &_BACnetConstructedDataCarMovingDirection{
		CarMovingDirection:     carMovingDirection,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarMovingDirection(structType any) BACnetConstructedDataCarMovingDirection {
	if casted, ok := structType.(BACnetConstructedDataCarMovingDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarMovingDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarMovingDirection) GetTypeName() string {
	return "BACnetConstructedDataCarMovingDirection"
}

func (m *_BACnetConstructedDataCarMovingDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (carMovingDirection)
	lengthInBits += m.CarMovingDirection.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarMovingDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCarMovingDirectionParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarMovingDirection, error) {
	return BACnetConstructedDataCarMovingDirectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCarMovingDirectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarMovingDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarMovingDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarMovingDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carMovingDirection)
	if pullErr := readBuffer.PullContext("carMovingDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carMovingDirection")
	}
	_carMovingDirection, _carMovingDirectionErr := BACnetLiftCarDirectionTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _carMovingDirectionErr != nil {
		return nil, errors.Wrap(_carMovingDirectionErr, "Error parsing 'carMovingDirection' field of BACnetConstructedDataCarMovingDirection")
	}
	carMovingDirection := _carMovingDirection.(BACnetLiftCarDirectionTagged)
	if closeErr := readBuffer.CloseContext("carMovingDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carMovingDirection")
	}

	// Virtual field
	_actualValue := carMovingDirection
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarMovingDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarMovingDirection")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarMovingDirection{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CarMovingDirection: carMovingDirection,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarMovingDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarMovingDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarMovingDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarMovingDirection")
		}

		// Simple Field (carMovingDirection)
		if pushErr := writeBuffer.PushContext("carMovingDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for carMovingDirection")
		}
		_carMovingDirectionErr := writeBuffer.WriteSerializable(ctx, m.GetCarMovingDirection())
		if popErr := writeBuffer.PopContext("carMovingDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for carMovingDirection")
		}
		if _carMovingDirectionErr != nil {
			return errors.Wrap(_carMovingDirectionErr, "Error serializing 'carMovingDirection' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarMovingDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarMovingDirection")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarMovingDirection) isBACnetConstructedDataCarMovingDirection() bool {
	return true
}

func (m *_BACnetConstructedDataCarMovingDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
