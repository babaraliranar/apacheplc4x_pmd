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

// BACnetConstructedDataPulseConverterAdjustValue is the corresponding interface of BACnetConstructedDataPulseConverterAdjustValue
type BACnetConstructedDataPulseConverterAdjustValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataPulseConverterAdjustValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPulseConverterAdjustValue.
// This is useful for switch cases.
type BACnetConstructedDataPulseConverterAdjustValueExactly interface {
	BACnetConstructedDataPulseConverterAdjustValue
	isBACnetConstructedDataPulseConverterAdjustValue() bool
}

// _BACnetConstructedDataPulseConverterAdjustValue is the data-structure of this message
type _BACnetConstructedDataPulseConverterAdjustValue struct {
	*_BACnetConstructedData
	AdjustValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_PULSE_CONVERTER
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetAdjustValue() BACnetApplicationTagReal {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPulseConverterAdjustValue factory function for _BACnetConstructedDataPulseConverterAdjustValue
func NewBACnetConstructedDataPulseConverterAdjustValue(adjustValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPulseConverterAdjustValue {
	_result := &_BACnetConstructedDataPulseConverterAdjustValue{
		AdjustValue:            adjustValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPulseConverterAdjustValue(structType any) BACnetConstructedDataPulseConverterAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataPulseConverterAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPulseConverterAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataPulseConverterAdjustValue"
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPulseConverterAdjustValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPulseConverterAdjustValue, error) {
	return BACnetConstructedDataPulseConverterAdjustValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPulseConverterAdjustValueParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataPulseConverterAdjustValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataPulseConverterAdjustValue, error) {
		return BACnetConstructedDataPulseConverterAdjustValueParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataPulseConverterAdjustValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPulseConverterAdjustValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPulseConverterAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPulseConverterAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	adjustValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "adjustValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adjustValue' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), adjustValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPulseConverterAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPulseConverterAdjustValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPulseConverterAdjustValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AdjustValue: adjustValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPulseConverterAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPulseConverterAdjustValue")
		}

		// Simple Field (adjustValue)
		if pushErr := writeBuffer.PushContext("adjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for adjustValue")
		}
		_adjustValueErr := writeBuffer.WriteSerializable(ctx, m.GetAdjustValue())
		if popErr := writeBuffer.PopContext("adjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for adjustValue")
		}
		if _adjustValueErr != nil {
			return errors.Wrap(_adjustValueErr, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPulseConverterAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPulseConverterAdjustValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) isBACnetConstructedDataPulseConverterAdjustValue() bool {
	return true
}

func (m *_BACnetConstructedDataPulseConverterAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
