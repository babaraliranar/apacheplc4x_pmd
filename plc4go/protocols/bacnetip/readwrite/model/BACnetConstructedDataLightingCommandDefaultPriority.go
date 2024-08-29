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

// BACnetConstructedDataLightingCommandDefaultPriority is the corresponding interface of BACnetConstructedDataLightingCommandDefaultPriority
type BACnetConstructedDataLightingCommandDefaultPriority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLightingCommandDefaultPriority returns LightingCommandDefaultPriority (property field)
	GetLightingCommandDefaultPriority() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataLightingCommandDefaultPriorityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLightingCommandDefaultPriority.
// This is useful for switch cases.
type BACnetConstructedDataLightingCommandDefaultPriorityExactly interface {
	BACnetConstructedDataLightingCommandDefaultPriority
	isBACnetConstructedDataLightingCommandDefaultPriority() bool
}

// _BACnetConstructedDataLightingCommandDefaultPriority is the data-structure of this message
type _BACnetConstructedDataLightingCommandDefaultPriority struct {
	*_BACnetConstructedData
	LightingCommandDefaultPriority BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND_DEFAULT_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLightingCommandDefaultPriority() BACnetApplicationTagUnsignedInteger {
	return m.LightingCommandDefaultPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLightingCommandDefaultPriority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingCommandDefaultPriority factory function for _BACnetConstructedDataLightingCommandDefaultPriority
func NewBACnetConstructedDataLightingCommandDefaultPriority(lightingCommandDefaultPriority BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingCommandDefaultPriority {
	_result := &_BACnetConstructedDataLightingCommandDefaultPriority{
		LightingCommandDefaultPriority: lightingCommandDefaultPriority,
		_BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingCommandDefaultPriority(structType any) BACnetConstructedDataLightingCommandDefaultPriority {
	if casted, ok := structType.(BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetTypeName() string {
	return "BACnetConstructedDataLightingCommandDefaultPriority"
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lightingCommandDefaultPriority)
	lengthInBits += m.LightingCommandDefaultPriority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLightingCommandDefaultPriorityParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingCommandDefaultPriority, error) {
	return BACnetConstructedDataLightingCommandDefaultPriorityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLightingCommandDefaultPriorityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingCommandDefaultPriority, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommandDefaultPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightingCommandDefaultPriority)
	if pullErr := readBuffer.PullContext("lightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightingCommandDefaultPriority")
	}
	_lightingCommandDefaultPriority, _lightingCommandDefaultPriorityErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _lightingCommandDefaultPriorityErr != nil {
		return nil, errors.Wrap(_lightingCommandDefaultPriorityErr, "Error parsing 'lightingCommandDefaultPriority' field of BACnetConstructedDataLightingCommandDefaultPriority")
	}
	lightingCommandDefaultPriority := _lightingCommandDefaultPriority.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightingCommandDefaultPriority")
	}

	// Virtual field
	_actualValue := lightingCommandDefaultPriority
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommandDefaultPriority")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLightingCommandDefaultPriority{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LightingCommandDefaultPriority: lightingCommandDefaultPriority,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommandDefaultPriority")
		}

		// Simple Field (lightingCommandDefaultPriority)
		if pushErr := writeBuffer.PushContext("lightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightingCommandDefaultPriority")
		}
		_lightingCommandDefaultPriorityErr := writeBuffer.WriteSerializable(ctx, m.GetLightingCommandDefaultPriority())
		if popErr := writeBuffer.PopContext("lightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightingCommandDefaultPriority")
		}
		if _lightingCommandDefaultPriorityErr != nil {
			return errors.Wrap(_lightingCommandDefaultPriorityErr, "Error serializing 'lightingCommandDefaultPriority' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommandDefaultPriority")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) isBACnetConstructedDataLightingCommandDefaultPriority() bool {
	return true
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
