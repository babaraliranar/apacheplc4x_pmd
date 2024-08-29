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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAccessZoneAlarmValues is the corresponding interface of BACnetConstructedDataAccessZoneAlarmValues
type BACnetConstructedDataAccessZoneAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetAccessZoneOccupancyStateTagged
}

// BACnetConstructedDataAccessZoneAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessZoneAlarmValues.
// This is useful for switch cases.
type BACnetConstructedDataAccessZoneAlarmValuesExactly interface {
	BACnetConstructedDataAccessZoneAlarmValues
	isBACnetConstructedDataAccessZoneAlarmValues() bool
}

// _BACnetConstructedDataAccessZoneAlarmValues is the data-structure of this message
type _BACnetConstructedDataAccessZoneAlarmValues struct {
	*_BACnetConstructedData
	AlarmValues []BACnetAccessZoneOccupancyStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_ZONE
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessZoneAlarmValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetAlarmValues() []BACnetAccessZoneOccupancyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessZoneAlarmValues factory function for _BACnetConstructedDataAccessZoneAlarmValues
func NewBACnetConstructedDataAccessZoneAlarmValues(alarmValues []BACnetAccessZoneOccupancyStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessZoneAlarmValues {
	_result := &_BACnetConstructedDataAccessZoneAlarmValues{
		AlarmValues:            alarmValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessZoneAlarmValues(structType any) BACnetConstructedDataAccessZoneAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataAccessZoneAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessZoneAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataAccessZoneAlarmValues"
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccessZoneAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessZoneAlarmValues, error) {
	return BACnetConstructedDataAccessZoneAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccessZoneAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessZoneAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessZoneAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessZoneAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (alarmValues)
	if pullErr := readBuffer.PullContext("alarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValues")
	}
	// Terminated array
	var alarmValues []BACnetAccessZoneOccupancyStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetAccessZoneOccupancyStateTaggedParseWithBuffer(ctx, readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'alarmValues' field of BACnetConstructedDataAccessZoneAlarmValues")
			}
			alarmValues = append(alarmValues, _item.(BACnetAccessZoneOccupancyStateTagged))
		}
	}
	if closeErr := readBuffer.CloseContext("alarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessZoneAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessZoneAlarmValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessZoneAlarmValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AlarmValues: alarmValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessZoneAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessZoneAlarmValues")
		}

		// Array Field (alarmValues)
		if pushErr := writeBuffer.PushContext("alarmValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmValues")
		}
		for _curItem, _element := range m.GetAlarmValues() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAlarmValues()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'alarmValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("alarmValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessZoneAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessZoneAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) isBACnetConstructedDataAccessZoneAlarmValues() bool {
	return true
}

func (m *_BACnetConstructedDataAccessZoneAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
