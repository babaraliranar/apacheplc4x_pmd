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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// MediaTransportControlDataCategoryName is the corresponding interface of MediaTransportControlDataCategoryName
type MediaTransportControlDataCategoryName interface {
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetCategoryName returns CategoryName (property field)
	GetCategoryName() string
}

// MediaTransportControlDataCategoryNameExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataCategoryName.
// This is useful for switch cases.
type MediaTransportControlDataCategoryNameExactly interface {
	MediaTransportControlDataCategoryName
	isMediaTransportControlDataCategoryName() bool
}

// _MediaTransportControlDataCategoryName is the data-structure of this message
type _MediaTransportControlDataCategoryName struct {
	*_MediaTransportControlData
	CategoryName string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataCategoryName) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataCategoryName) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataCategoryName) GetCategoryName() string {
	return m.CategoryName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataCategoryName factory function for _MediaTransportControlDataCategoryName
func NewMediaTransportControlDataCategoryName(categoryName string, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataCategoryName {
	_result := &_MediaTransportControlDataCategoryName{
		CategoryName:               categoryName,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataCategoryName(structType interface{}) MediaTransportControlDataCategoryName {
	if casted, ok := structType.(MediaTransportControlDataCategoryName); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataCategoryName); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataCategoryName) GetTypeName() string {
	return "MediaTransportControlDataCategoryName"
}

func (m *_MediaTransportControlDataCategoryName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MediaTransportControlDataCategoryName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (categoryName)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_MediaTransportControlDataCategoryName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MediaTransportControlDataCategoryNameParse(theBytes []byte, commandTypeContainer MediaTransportControlCommandTypeContainer) (MediaTransportControlDataCategoryName, error) {
	return MediaTransportControlDataCategoryNameParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), commandTypeContainer) // TODO: get endianness from mspec
}

func MediaTransportControlDataCategoryNameParseWithBuffer(readBuffer utils.ReadBuffer, commandTypeContainer MediaTransportControlCommandTypeContainer) (MediaTransportControlDataCategoryName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataCategoryName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataCategoryName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (categoryName)
	_categoryName, _categoryNameErr := readBuffer.ReadString("categoryName", uint32(((commandTypeContainer.NumBytes())-(1))*(8)), "UTF-8")
	if _categoryNameErr != nil {
		return nil, errors.Wrap(_categoryNameErr, "Error parsing 'categoryName' field of MediaTransportControlDataCategoryName")
	}
	categoryName := _categoryName

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataCategoryName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataCategoryName")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataCategoryName{
		_MediaTransportControlData: &_MediaTransportControlData{},
		CategoryName:               categoryName,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataCategoryName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataCategoryName) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataCategoryName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataCategoryName")
		}

		// Simple Field (categoryName)
		categoryName := string(m.GetCategoryName())
		_categoryNameErr := writeBuffer.WriteString("categoryName", uint32(((m.GetCommandTypeContainer().NumBytes())-(1))*(8)), "UTF-8", (categoryName))
		if _categoryNameErr != nil {
			return errors.Wrap(_categoryNameErr, "Error serializing 'categoryName' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataCategoryName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataCategoryName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataCategoryName) isMediaTransportControlDataCategoryName() bool {
	return true
}

func (m *_MediaTransportControlDataCategoryName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
