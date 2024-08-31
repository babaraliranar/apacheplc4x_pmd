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

// BACnetLiftCarCallListFloorList is the corresponding interface of BACnetLiftCarCallListFloorList
type BACnetLiftCarCallListFloorList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetFloorNumbers returns FloorNumbers (property field)
	GetFloorNumbers() []BACnetApplicationTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetLiftCarCallListFloorListExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftCarCallListFloorList.
// This is useful for switch cases.
type BACnetLiftCarCallListFloorListExactly interface {
	BACnetLiftCarCallListFloorList
	isBACnetLiftCarCallListFloorList() bool
}

// _BACnetLiftCarCallListFloorList is the data-structure of this message
type _BACnetLiftCarCallListFloorList struct {
	OpeningTag   BACnetOpeningTag
	FloorNumbers []BACnetApplicationTagUnsignedInteger
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarCallListFloorList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLiftCarCallListFloorList) GetFloorNumbers() []BACnetApplicationTagUnsignedInteger {
	return m.FloorNumbers
}

func (m *_BACnetLiftCarCallListFloorList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLiftCarCallListFloorList factory function for _BACnetLiftCarCallListFloorList
func NewBACnetLiftCarCallListFloorList(openingTag BACnetOpeningTag, floorNumbers []BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLiftCarCallListFloorList {
	return &_BACnetLiftCarCallListFloorList{OpeningTag: openingTag, FloorNumbers: floorNumbers, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarCallListFloorList(structType any) BACnetLiftCarCallListFloorList {
	if casted, ok := structType.(BACnetLiftCarCallListFloorList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarCallListFloorList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarCallListFloorList) GetTypeName() string {
	return "BACnetLiftCarCallListFloorList"
}

func (m *_BACnetLiftCarCallListFloorList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.FloorNumbers) > 0 {
		for _, element := range m.FloorNumbers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLiftCarCallListFloorList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarCallListFloorListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetLiftCarCallListFloorList, error) {
	return BACnetLiftCarCallListFloorListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLiftCarCallListFloorListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLiftCarCallListFloorList, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLiftCarCallListFloorList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarCallListFloorList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetLiftCarCallListFloorList")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	floorNumbers, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "floorNumbers", ReadComplex[BACnetApplicationTagUnsignedInteger](func(ctx context.Context, buffer utils.ReadBuffer) (BACnetApplicationTagUnsignedInteger, error) {
		v, err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		if err != nil {
			return nil, err
		}
		return v.(BACnetApplicationTagUnsignedInteger), nil
	}, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorNumbers' field"))
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetLiftCarCallListFloorList")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftCarCallListFloorList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarCallListFloorList")
	}

	// Create the instance
	return &_BACnetLiftCarCallListFloorList{
		TagNumber:    tagNumber,
		OpeningTag:   openingTag,
		FloorNumbers: floorNumbers,
		ClosingTag:   closingTag,
	}, nil
}

func (m *_BACnetLiftCarCallListFloorList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarCallListFloorList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLiftCarCallListFloorList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarCallListFloorList")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (floorNumbers)
	if pushErr := writeBuffer.PushContext("floorNumbers", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for floorNumbers")
	}
	for _curItem, _element := range m.GetFloorNumbers() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetFloorNumbers()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'floorNumbers' field")
		}
	}
	if popErr := writeBuffer.PopContext("floorNumbers", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for floorNumbers")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLiftCarCallListFloorList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarCallListFloorList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLiftCarCallListFloorList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLiftCarCallListFloorList) isBACnetLiftCarCallListFloorList() bool {
	return true
}

func (m *_BACnetLiftCarCallListFloorList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
