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

// BACnetConstructedDataAlarmValue is the corresponding interface of BACnetConstructedDataAlarmValue
type BACnetConstructedDataAlarmValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBinaryPv returns BinaryPv (property field)
	GetBinaryPv() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
}

// BACnetConstructedDataAlarmValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAlarmValue.
// This is useful for switch cases.
type BACnetConstructedDataAlarmValueExactly interface {
	BACnetConstructedDataAlarmValue
	isBACnetConstructedDataAlarmValue() bool
}

// _BACnetConstructedDataAlarmValue is the data-structure of this message
type _BACnetConstructedDataAlarmValue struct {
	*_BACnetConstructedData
	BinaryPv BACnetBinaryPVTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlarmValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAlarmValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlarmValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAlarmValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAlarmValue) GetBinaryPv() BACnetBinaryPVTagged {
	return m.BinaryPv
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAlarmValue) GetActualValue() BACnetBinaryPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryPVTagged(m.GetBinaryPv())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAlarmValue factory function for _BACnetConstructedDataAlarmValue
func NewBACnetConstructedDataAlarmValue(binaryPv BACnetBinaryPVTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAlarmValue {
	_result := &_BACnetConstructedDataAlarmValue{
		BinaryPv:               binaryPv,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlarmValue(structType any) BACnetConstructedDataAlarmValue {
	if casted, ok := structType.(BACnetConstructedDataAlarmValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlarmValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlarmValue) GetTypeName() string {
	return "BACnetConstructedDataAlarmValue"
}

func (m *_BACnetConstructedDataAlarmValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (binaryPv)
	lengthInBits += m.BinaryPv.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAlarmValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAlarmValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAlarmValue, error) {
	return BACnetConstructedDataAlarmValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAlarmValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAlarmValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlarmValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlarmValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryPv)
	if pullErr := readBuffer.PullContext("binaryPv"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for binaryPv")
	}
	_binaryPv, _binaryPvErr := BACnetBinaryPVTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _binaryPvErr != nil {
		return nil, errors.Wrap(_binaryPvErr, "Error parsing 'binaryPv' field of BACnetConstructedDataAlarmValue")
	}
	binaryPv := _binaryPv.(BACnetBinaryPVTagged)
	if closeErr := readBuffer.CloseContext("binaryPv"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for binaryPv")
	}

	// Virtual field
	_actualValue := binaryPv
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlarmValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlarmValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAlarmValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BinaryPv: binaryPv,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAlarmValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAlarmValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlarmValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlarmValue")
		}

		// Simple Field (binaryPv)
		if pushErr := writeBuffer.PushContext("binaryPv"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for binaryPv")
		}
		_binaryPvErr := writeBuffer.WriteSerializable(ctx, m.GetBinaryPv())
		if popErr := writeBuffer.PopContext("binaryPv"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for binaryPv")
		}
		if _binaryPvErr != nil {
			return errors.Wrap(_binaryPvErr, "Error serializing 'binaryPv' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlarmValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlarmValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAlarmValue) isBACnetConstructedDataAlarmValue() bool {
	return true
}

func (m *_BACnetConstructedDataAlarmValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
