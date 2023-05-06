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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTrendLogMultipleLogBuffer is the corresponding interface of BACnetConstructedDataTrendLogMultipleLogBuffer
type BACnetConstructedDataTrendLogMultipleLogBuffer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetLogMultipleRecord
}

// BACnetConstructedDataTrendLogMultipleLogBufferExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTrendLogMultipleLogBuffer.
// This is useful for switch cases.
type BACnetConstructedDataTrendLogMultipleLogBufferExactly interface {
	BACnetConstructedDataTrendLogMultipleLogBuffer
	isBACnetConstructedDataTrendLogMultipleLogBuffer() bool
}

// _BACnetConstructedDataTrendLogMultipleLogBuffer is the data-structure of this message
type _BACnetConstructedDataTrendLogMultipleLogBuffer struct {
	*_BACnetConstructedData
	FloorText []BACnetLogMultipleRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG_MULTIPLE
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetFloorText() []BACnetLogMultipleRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrendLogMultipleLogBuffer factory function for _BACnetConstructedDataTrendLogMultipleLogBuffer
func NewBACnetConstructedDataTrendLogMultipleLogBuffer(floorText []BACnetLogMultipleRecord, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogMultipleLogBuffer {
	_result := &_BACnetConstructedDataTrendLogMultipleLogBuffer{
		FloorText:              floorText,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogMultipleLogBuffer(structType any) BACnetConstructedDataTrendLogMultipleLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataTrendLogMultipleLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogMultipleLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataTrendLogMultipleLogBuffer"
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTrendLogMultipleLogBufferParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogMultipleLogBuffer, error) {
	return BACnetConstructedDataTrendLogMultipleLogBufferParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTrendLogMultipleLogBufferParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogMultipleLogBuffer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogMultipleLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (floorText)
	if pullErr := readBuffer.PullContext("floorText", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorText")
	}
	// Terminated array
	var floorText []BACnetLogMultipleRecord
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLogMultipleRecordParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'floorText' field of BACnetConstructedDataTrendLogMultipleLogBuffer")
			}
			floorText = append(floorText, _item.(BACnetLogMultipleRecord))
		}
	}
	if closeErr := readBuffer.CloseContext("floorText", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogMultipleLogBuffer")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTrendLogMultipleLogBuffer{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FloorText: floorText,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogMultipleLogBuffer")
		}

		// Array Field (floorText)
		if pushErr := writeBuffer.PushContext("floorText", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for floorText")
		}
		for _curItem, _element := range m.GetFloorText() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetFloorText()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'floorText' field")
			}
		}
		if popErr := writeBuffer.PopContext("floorText", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for floorText")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogMultipleLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogMultipleLogBuffer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) isBACnetConstructedDataTrendLogMultipleLogBuffer() bool {
	return true
}

func (m *_BACnetConstructedDataTrendLogMultipleLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
