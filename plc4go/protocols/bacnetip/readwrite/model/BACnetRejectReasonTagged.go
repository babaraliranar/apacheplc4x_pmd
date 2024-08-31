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

// BACnetRejectReasonTagged is the corresponding interface of BACnetRejectReasonTagged
type BACnetRejectReasonTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() BACnetRejectReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetRejectReasonTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetRejectReasonTagged.
// This is useful for switch cases.
type BACnetRejectReasonTaggedExactly interface {
	BACnetRejectReasonTagged
	isBACnetRejectReasonTagged() bool
}

// _BACnetRejectReasonTagged is the data-structure of this message
type _BACnetRejectReasonTagged struct {
	Value            BACnetRejectReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRejectReasonTagged) GetValue() BACnetRejectReason {
	return m.Value
}

func (m *_BACnetRejectReasonTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetRejectReasonTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRejectReasonTagged factory function for _BACnetRejectReasonTagged
func NewBACnetRejectReasonTagged(value BACnetRejectReason, proprietaryValue uint32, actualLength uint32) *_BACnetRejectReasonTagged {
	return &_BACnetRejectReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetRejectReasonTagged(structType any) BACnetRejectReasonTagged {
	if casted, ok := structType.(BACnetRejectReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRejectReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRejectReasonTagged) GetTypeName() string {
	return "BACnetRejectReasonTagged"
}

func (m *_BACnetRejectReasonTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.ActualLength) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.ActualLength) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetRejectReasonTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRejectReasonTaggedParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetRejectReasonTagged, error) {
	return BACnetRejectReasonTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetRejectReasonTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetRejectReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetRejectReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRejectReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadManualField[BACnetRejectReason](ctx, "value", readBuffer, EnsureType[BACnetRejectReason](ReadEnumGeneric(ctx, readBuffer, actualLength, BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetRejectReason_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, actualLength, isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetRejectReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRejectReasonTagged")
	}

	// Create the instance
	return &_BACnetRejectReasonTagged{
		ActualLength:     actualLength,
		Value:            value,
		ProprietaryValue: proprietaryValue,
	}, nil
}

func (m *_BACnetRejectReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRejectReasonTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRejectReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRejectReasonTagged")
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

	if popErr := writeBuffer.PopContext("BACnetRejectReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRejectReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRejectReasonTagged) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetRejectReasonTagged) isBACnetRejectReasonTagged() bool {
	return true
}

func (m *_BACnetRejectReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
