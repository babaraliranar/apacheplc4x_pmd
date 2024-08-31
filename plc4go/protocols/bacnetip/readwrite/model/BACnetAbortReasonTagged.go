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
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAbortReasonTagged is the corresponding interface of BACnetAbortReasonTagged
type BACnetAbortReasonTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() BACnetAbortReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetAbortReasonTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetAbortReasonTagged.
// This is useful for switch cases.
type BACnetAbortReasonTaggedExactly interface {
	BACnetAbortReasonTagged
	isBACnetAbortReasonTagged() bool
}

// _BACnetAbortReasonTagged is the data-structure of this message
type _BACnetAbortReasonTagged struct {
	Value            BACnetAbortReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAbortReasonTagged) GetValue() BACnetAbortReason {
	return m.Value
}

func (m *_BACnetAbortReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetAbortReasonTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAbortReasonTagged factory function for _BACnetAbortReasonTagged
func NewBACnetAbortReasonTagged(value BACnetAbortReason, proprietaryValue uint32, actualLength uint32) *_BACnetAbortReasonTagged {
	return &_BACnetAbortReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetAbortReasonTagged(structType any) BACnetAbortReasonTagged {
	if casted, ok := structType.(BACnetAbortReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAbortReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAbortReasonTagged) GetTypeName() string {
	return "BACnetAbortReasonTagged"
}

func (m *_BACnetAbortReasonTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.ActualLength) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.ActualLength) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetAbortReasonTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAbortReasonTaggedParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetAbortReasonTagged, error) {
	return BACnetAbortReasonTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetAbortReasonTaggedParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAbortReasonTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAbortReasonTagged, error) {
		return BACnetAbortReasonTaggedParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetAbortReasonTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetAbortReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAbortReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAbortReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadManualField[BACnetAbortReason](ctx, "value", readBuffer, EnsureType[BACnetAbortReason](ReadEnumGeneric(ctx, readBuffer, actualLength, BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, actualLength, isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetAbortReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAbortReasonTagged")
	}

	// Create the instance
	return &_BACnetAbortReasonTagged{
		ActualLength:     actualLength,
		Value:            value,
		ProprietaryValue: proprietaryValue,
	}, nil
}

func (m *_BACnetAbortReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAbortReasonTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAbortReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAbortReasonTagged")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(ctx, writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAbortReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAbortReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAbortReasonTagged) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetAbortReasonTagged) isBACnetAbortReasonTagged() bool {
	return true
}

func (m *_BACnetAbortReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
