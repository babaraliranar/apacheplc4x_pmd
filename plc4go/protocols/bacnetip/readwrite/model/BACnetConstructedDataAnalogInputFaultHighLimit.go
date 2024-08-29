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

// BACnetConstructedDataAnalogInputFaultHighLimit is the corresponding interface of BACnetConstructedDataAnalogInputFaultHighLimit
type BACnetConstructedDataAnalogInputFaultHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataAnalogInputFaultHighLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAnalogInputFaultHighLimit.
// This is useful for switch cases.
type BACnetConstructedDataAnalogInputFaultHighLimitExactly interface {
	BACnetConstructedDataAnalogInputFaultHighLimit
	isBACnetConstructedDataAnalogInputFaultHighLimit() bool
}

// _BACnetConstructedDataAnalogInputFaultHighLimit is the data-structure of this message
type _BACnetConstructedDataAnalogInputFaultHighLimit struct {
	*_BACnetConstructedData
	FaultHighLimit BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_INPUT
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetFaultHighLimit() BACnetApplicationTagReal {
	return m.FaultHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFaultHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAnalogInputFaultHighLimit factory function for _BACnetConstructedDataAnalogInputFaultHighLimit
func NewBACnetConstructedDataAnalogInputFaultHighLimit(faultHighLimit BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogInputFaultHighLimit {
	_result := &_BACnetConstructedDataAnalogInputFaultHighLimit{
		FaultHighLimit:         faultHighLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogInputFaultHighLimit(structType any) BACnetConstructedDataAnalogInputFaultHighLimit {
	if casted, ok := structType.(BACnetConstructedDataAnalogInputFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogInputFaultHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataAnalogInputFaultHighLimit"
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAnalogInputFaultHighLimitParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogInputFaultHighLimit, error) {
	return BACnetConstructedDataAnalogInputFaultHighLimitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAnalogInputFaultHighLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogInputFaultHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogInputFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogInputFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultHighLimit)
	if pullErr := readBuffer.PullContext("faultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultHighLimit")
	}
	_faultHighLimit, _faultHighLimitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _faultHighLimitErr != nil {
		return nil, errors.Wrap(_faultHighLimitErr, "Error parsing 'faultHighLimit' field of BACnetConstructedDataAnalogInputFaultHighLimit")
	}
	faultHighLimit := _faultHighLimit.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("faultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultHighLimit")
	}

	// Virtual field
	_actualValue := faultHighLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogInputFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogInputFaultHighLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAnalogInputFaultHighLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultHighLimit: faultHighLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogInputFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogInputFaultHighLimit")
		}

		// Simple Field (faultHighLimit)
		if pushErr := writeBuffer.PushContext("faultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultHighLimit")
		}
		_faultHighLimitErr := writeBuffer.WriteSerializable(ctx, m.GetFaultHighLimit())
		if popErr := writeBuffer.PopContext("faultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultHighLimit")
		}
		if _faultHighLimitErr != nil {
			return errors.Wrap(_faultHighLimitErr, "Error serializing 'faultHighLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogInputFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogInputFaultHighLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) isBACnetConstructedDataAnalogInputFaultHighLimit() bool {
	return true
}

func (m *_BACnetConstructedDataAnalogInputFaultHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
