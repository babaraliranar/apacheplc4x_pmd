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

// BACnetSpecialEventPeriodCalendarEntry is the corresponding interface of BACnetSpecialEventPeriodCalendarEntry
type BACnetSpecialEventPeriodCalendarEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetSpecialEventPeriod
	// GetCalendarEntry returns CalendarEntry (property field)
	GetCalendarEntry() BACnetCalendarEntryEnclosed
}

// BACnetSpecialEventPeriodCalendarEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetSpecialEventPeriodCalendarEntry.
// This is useful for switch cases.
type BACnetSpecialEventPeriodCalendarEntryExactly interface {
	BACnetSpecialEventPeriodCalendarEntry
	isBACnetSpecialEventPeriodCalendarEntry() bool
}

// _BACnetSpecialEventPeriodCalendarEntry is the data-structure of this message
type _BACnetSpecialEventPeriodCalendarEntry struct {
	*_BACnetSpecialEventPeriod
	CalendarEntry BACnetCalendarEntryEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetSpecialEventPeriodCalendarEntry) InitializeParent(parent BACnetSpecialEventPeriod, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetParent() BACnetSpecialEventPeriod {
	return m._BACnetSpecialEventPeriod
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetCalendarEntry() BACnetCalendarEntryEnclosed {
	return m.CalendarEntry
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventPeriodCalendarEntry factory function for _BACnetSpecialEventPeriodCalendarEntry
func NewBACnetSpecialEventPeriodCalendarEntry(calendarEntry BACnetCalendarEntryEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetSpecialEventPeriodCalendarEntry {
	_result := &_BACnetSpecialEventPeriodCalendarEntry{
		CalendarEntry:             calendarEntry,
		_BACnetSpecialEventPeriod: NewBACnetSpecialEventPeriod(peekedTagHeader),
	}
	_result._BACnetSpecialEventPeriod._BACnetSpecialEventPeriodChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventPeriodCalendarEntry(structType any) BACnetSpecialEventPeriodCalendarEntry {
	if casted, ok := structType.(BACnetSpecialEventPeriodCalendarEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriodCalendarEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetTypeName() string {
	return "BACnetSpecialEventPeriodCalendarEntry"
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (calendarEntry)
	lengthInBits += m.CalendarEntry.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventPeriodCalendarEntryParse(ctx context.Context, theBytes []byte) (BACnetSpecialEventPeriodCalendarEntry, error) {
	return BACnetSpecialEventPeriodCalendarEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetSpecialEventPeriodCalendarEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEventPeriodCalendarEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetSpecialEventPeriodCalendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventPeriodCalendarEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (calendarEntry)
	if pullErr := readBuffer.PullContext("calendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calendarEntry")
	}
	_calendarEntry, _calendarEntryErr := BACnetCalendarEntryEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _calendarEntryErr != nil {
		return nil, errors.Wrap(_calendarEntryErr, "Error parsing 'calendarEntry' field of BACnetSpecialEventPeriodCalendarEntry")
	}
	calendarEntry := _calendarEntry.(BACnetCalendarEntryEnclosed)
	if closeErr := readBuffer.CloseContext("calendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calendarEntry")
	}

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventPeriodCalendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventPeriodCalendarEntry")
	}

	// Create a partially initialized instance
	_child := &_BACnetSpecialEventPeriodCalendarEntry{
		_BACnetSpecialEventPeriod: &_BACnetSpecialEventPeriod{},
		CalendarEntry:             calendarEntry,
	}
	_child._BACnetSpecialEventPeriod._BACnetSpecialEventPeriodChildRequirements = _child
	return _child, nil
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetSpecialEventPeriodCalendarEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventPeriodCalendarEntry")
		}

		// Simple Field (calendarEntry)
		if pushErr := writeBuffer.PushContext("calendarEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for calendarEntry")
		}
		_calendarEntryErr := writeBuffer.WriteSerializable(ctx, m.GetCalendarEntry())
		if popErr := writeBuffer.PopContext("calendarEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for calendarEntry")
		}
		if _calendarEntryErr != nil {
			return errors.Wrap(_calendarEntryErr, "Error serializing 'calendarEntry' field")
		}

		if popErr := writeBuffer.PopContext("BACnetSpecialEventPeriodCalendarEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetSpecialEventPeriodCalendarEntry")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) isBACnetSpecialEventPeriodCalendarEntry() bool {
	return true
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
