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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetEventTimestampsEnclosed is the corresponding interface of BACnetEventTimestampsEnclosed
type BACnetEventTimestampsEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetEventTimestamps returns EventTimestamps (property field)
	GetEventTimestamps() BACnetEventTimestamps
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventTimestampsEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetEventTimestampsEnclosed.
// This is useful for switch cases.
type BACnetEventTimestampsEnclosedExactly interface {
	BACnetEventTimestampsEnclosed
	isBACnetEventTimestampsEnclosed() bool
}

// _BACnetEventTimestampsEnclosed is the data-structure of this message
type _BACnetEventTimestampsEnclosed struct {
        OpeningTag BACnetOpeningTag
        EventTimestamps BACnetEventTimestamps
        ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventTimestampsEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventTimestampsEnclosed) GetEventTimestamps() BACnetEventTimestamps {
	return m.EventTimestamps
}

func (m *_BACnetEventTimestampsEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventTimestampsEnclosed factory function for _BACnetEventTimestampsEnclosed
func NewBACnetEventTimestampsEnclosed( openingTag BACnetOpeningTag , eventTimestamps BACnetEventTimestamps , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetEventTimestampsEnclosed {
return &_BACnetEventTimestampsEnclosed{ OpeningTag: openingTag , EventTimestamps: eventTimestamps , ClosingTag: closingTag , TagNumber: tagNumber }
}

// Deprecated: use the interface for direct cast
func CastBACnetEventTimestampsEnclosed(structType interface{}) BACnetEventTimestampsEnclosed {
    if casted, ok := structType.(BACnetEventTimestampsEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventTimestampsEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventTimestampsEnclosed) GetTypeName() string {
	return "BACnetEventTimestampsEnclosed"
}

func (m *_BACnetEventTimestampsEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (eventTimestamps)
	lengthInBits += m.EventTimestamps.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetEventTimestampsEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventTimestampsEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetEventTimestampsEnclosed, error) {
	return BACnetEventTimestampsEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventTimestampsEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventTimestampsEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTimestampsEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTimestampsEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer , uint8( tagNumber ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventTimestampsEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (eventTimestamps)
	if pullErr := readBuffer.PullContext("eventTimestamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventTimestamps")
	}
_eventTimestamps, _eventTimestampsErr := BACnetEventTimestampsParseWithBuffer(ctx, readBuffer)
	if _eventTimestampsErr != nil {
		return nil, errors.Wrap(_eventTimestampsErr, "Error parsing 'eventTimestamps' field of BACnetEventTimestampsEnclosed")
	}
	eventTimestamps := _eventTimestamps.(BACnetEventTimestamps)
	if closeErr := readBuffer.CloseContext("eventTimestamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventTimestamps")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer , uint8( tagNumber ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventTimestampsEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventTimestampsEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTimestampsEnclosed")
	}

	// Create the instance
	return &_BACnetEventTimestampsEnclosed{
            TagNumber: tagNumber,
			OpeningTag: openingTag,
			EventTimestamps: eventTimestamps,
			ClosingTag: closingTag,
		}, nil
}

func (m *_BACnetEventTimestampsEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventTimestampsEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetEventTimestampsEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTimestampsEnclosed")
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

	// Simple Field (eventTimestamps)
	if pushErr := writeBuffer.PushContext("eventTimestamps"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventTimestamps")
	}
	_eventTimestampsErr := writeBuffer.WriteSerializable(ctx, m.GetEventTimestamps())
	if popErr := writeBuffer.PopContext("eventTimestamps"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventTimestamps")
	}
	if _eventTimestampsErr != nil {
		return errors.Wrap(_eventTimestampsErr, "Error serializing 'eventTimestamps' field")
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

	if popErr := writeBuffer.PopContext("BACnetEventTimestampsEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTimestampsEnclosed")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetEventTimestampsEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}
//
////

func (m *_BACnetEventTimestampsEnclosed) isBACnetEventTimestampsEnclosed() bool {
	return true
}

func (m *_BACnetEventTimestampsEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



