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

// MediaTransportControlDataEnumerateCategoriesSelectionTracks is the corresponding interface of MediaTransportControlDataEnumerateCategoriesSelectionTracks
type MediaTransportControlDataEnumerateCategoriesSelectionTracks interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetEnumerateType returns EnumerateType (property field)
	GetEnumerateType() byte
	// GetStart returns Start (property field)
	GetStart() uint8
	// GetIsListCategories returns IsListCategories (virtual field)
	GetIsListCategories() bool
	// GetIsListSelections returns IsListSelections (virtual field)
	GetIsListSelections() bool
	// GetIsListTracks returns IsListTracks (virtual field)
	GetIsListTracks() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
}

// MediaTransportControlDataEnumerateCategoriesSelectionTracksExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataEnumerateCategoriesSelectionTracks.
// This is useful for switch cases.
type MediaTransportControlDataEnumerateCategoriesSelectionTracksExactly interface {
	MediaTransportControlDataEnumerateCategoriesSelectionTracks
	isMediaTransportControlDataEnumerateCategoriesSelectionTracks() bool
}

// _MediaTransportControlDataEnumerateCategoriesSelectionTracks is the data-structure of this message
type _MediaTransportControlDataEnumerateCategoriesSelectionTracks struct {
	*_MediaTransportControlData
	EnumerateType byte
	Start         uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetEnumerateType() byte {
	return m.EnumerateType
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetStart() uint8 {
	return m.Start
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetIsListCategories() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetEnumerateType()) == (0x00)))
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetIsListSelections() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetEnumerateType()) == (0x01)))
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetIsListTracks() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetEnumerateType()) == (0x02)))
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool(!(m.GetIsListCategories())) && bool(!(m.GetIsListSelections()))) && bool(!(m.GetIsListTracks())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataEnumerateCategoriesSelectionTracks factory function for _MediaTransportControlDataEnumerateCategoriesSelectionTracks
func NewMediaTransportControlDataEnumerateCategoriesSelectionTracks(enumerateType byte, start uint8, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataEnumerateCategoriesSelectionTracks {
	_result := &_MediaTransportControlDataEnumerateCategoriesSelectionTracks{
		EnumerateType:              enumerateType,
		Start:                      start,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataEnumerateCategoriesSelectionTracks(structType any) MediaTransportControlDataEnumerateCategoriesSelectionTracks {
	if casted, ok := structType.(MediaTransportControlDataEnumerateCategoriesSelectionTracks); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataEnumerateCategoriesSelectionTracks); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetTypeName() string {
	return "MediaTransportControlDataEnumerateCategoriesSelectionTracks"
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (enumerateType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (start)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataEnumerateCategoriesSelectionTracksParse(theBytes []byte) (MediaTransportControlDataEnumerateCategoriesSelectionTracks, error) {
	return MediaTransportControlDataEnumerateCategoriesSelectionTracksParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataEnumerateCategoriesSelectionTracksParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataEnumerateCategoriesSelectionTracks, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataEnumerateCategoriesSelectionTracks"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataEnumerateCategoriesSelectionTracks")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumerateType)
	_enumerateType, _enumerateTypeErr := readBuffer.ReadByte("enumerateType")
	if _enumerateTypeErr != nil {
		return nil, errors.Wrap(_enumerateTypeErr, "Error parsing 'enumerateType' field of MediaTransportControlDataEnumerateCategoriesSelectionTracks")
	}
	enumerateType := _enumerateType

	// Virtual field
	_isListCategories := bool((enumerateType) == (0x00))
	isListCategories := bool(_isListCategories)
	_ = isListCategories

	// Virtual field
	_isListSelections := bool((enumerateType) == (0x01))
	isListSelections := bool(_isListSelections)
	_ = isListSelections

	// Virtual field
	_isListTracks := bool((enumerateType) == (0x02))
	isListTracks := bool(_isListTracks)
	_ = isListTracks

	// Virtual field
	_isReserved := bool(bool(!(isListCategories)) && bool(!(isListSelections))) && bool(!(isListTracks))
	isReserved := bool(_isReserved)
	_ = isReserved

	// Simple Field (start)
	_start, _startErr := readBuffer.ReadUint8("start", 8)
	if _startErr != nil {
		return nil, errors.Wrap(_startErr, "Error parsing 'start' field of MediaTransportControlDataEnumerateCategoriesSelectionTracks")
	}
	start := _start

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataEnumerateCategoriesSelectionTracks"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataEnumerateCategoriesSelectionTracks")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataEnumerateCategoriesSelectionTracks{
		_MediaTransportControlData: &_MediaTransportControlData{},
		EnumerateType:              enumerateType,
		Start:                      start,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataEnumerateCategoriesSelectionTracks"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataEnumerateCategoriesSelectionTracks")
		}

		// Simple Field (enumerateType)
		enumerateType := byte(m.GetEnumerateType())
		_enumerateTypeErr := writeBuffer.WriteByte("enumerateType", (enumerateType))
		if _enumerateTypeErr != nil {
			return errors.Wrap(_enumerateTypeErr, "Error serializing 'enumerateType' field")
		}
		// Virtual field
		if _isListCategoriesErr := writeBuffer.WriteVirtual(ctx, "isListCategories", m.GetIsListCategories()); _isListCategoriesErr != nil {
			return errors.Wrap(_isListCategoriesErr, "Error serializing 'isListCategories' field")
		}
		// Virtual field
		if _isListSelectionsErr := writeBuffer.WriteVirtual(ctx, "isListSelections", m.GetIsListSelections()); _isListSelectionsErr != nil {
			return errors.Wrap(_isListSelectionsErr, "Error serializing 'isListSelections' field")
		}
		// Virtual field
		if _isListTracksErr := writeBuffer.WriteVirtual(ctx, "isListTracks", m.GetIsListTracks()); _isListTracksErr != nil {
			return errors.Wrap(_isListTracksErr, "Error serializing 'isListTracks' field")
		}
		// Virtual field
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}

		// Simple Field (start)
		start := uint8(m.GetStart())
		_startErr := writeBuffer.WriteUint8("start", 8, (start))
		if _startErr != nil {
			return errors.Wrap(_startErr, "Error serializing 'start' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataEnumerateCategoriesSelectionTracks"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataEnumerateCategoriesSelectionTracks")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) isMediaTransportControlDataEnumerateCategoriesSelectionTracks() bool {
	return true
}

func (m *_MediaTransportControlDataEnumerateCategoriesSelectionTracks) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
