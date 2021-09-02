/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsReadWriteRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	ReadLength  uint32
	Items       []*AdsMultiRequestItem
	Data        []int8
	Parent      *AdsData
}

// The corresponding interface
type IAdsReadWriteRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadWriteRequest) CommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *AdsReadWriteRequest) Response() bool {
	return false
}

func (m *AdsReadWriteRequest) InitializeParent(parent *AdsData) {
}

func NewAdsReadWriteRequest(indexGroup uint32, indexOffset uint32, readLength uint32, items []*AdsMultiRequestItem, data []int8) *AdsData {
	child := &AdsReadWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  readLength,
		Items:       items,
		Data:        data,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadWriteRequest(structType interface{}) *AdsReadWriteRequest {
	castFunc := func(typ interface{}) *AdsReadWriteRequest {
		if casted, ok := typ.(AdsReadWriteRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadWriteRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadWriteRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadWriteRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadWriteRequest) GetTypeName() string {
	return "AdsReadWriteRequest"
}

func (m *AdsReadWriteRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsReadWriteRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (readLength)
	lengthInBits += 32

	// Implicit Field (writeLength)
	lengthInBits += 32

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsReadWriteRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadWriteRequestParse(readBuffer utils.ReadBuffer) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsReadWriteRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (indexGroup)
	indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field")
	}

	// Simple Field (indexOffset)
	indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field")
	}

	// Simple Field (readLength)
	readLength, _readLengthErr := readBuffer.ReadUint32("readLength", 32)
	if _readLengthErr != nil {
		return nil, errors.Wrap(_readLengthErr, "Error parsing 'readLength' field")
	}

	// Implicit Field (writeLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	writeLength, _writeLengthErr := readBuffer.ReadUint32("writeLength", 32)
	_ = writeLength
	if _writeLengthErr != nil {
		return nil, errors.Wrap(_writeLengthErr, "Error parsing 'writeLength' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*AdsMultiRequestItem, utils.InlineIf(bool(bool(bool(bool(bool((indexGroup) == (61568)))) || bool(bool(bool((indexGroup) == (61569))))) || bool(bool(bool((indexGroup) == (61570))))), func() uint16 { return uint16(indexOffset) }, func() uint16 { return uint16(uint16(0)) }))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(bool(bool(bool(bool(bool((indexGroup) == (61568)))) || bool(bool(bool((indexGroup) == (61569))))) || bool(bool(bool((indexGroup) == (61570))))), func() uint16 { return uint16(indexOffset) }, func() uint16 { return uint16(uint16(0)) })); curItem++ {
		_item, _err := AdsMultiRequestItemParse(readBuffer, indexGroup)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	data := make([]int8, uint16(writeLength)-uint16(uint16(uint16(uint16(len(items)))*uint16(uint16(12)))))
	for curItem := uint16(0); curItem < uint16(uint16(writeLength)-uint16(uint16(uint16(uint16(len(items)))*uint16(uint16(12))))); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsReadWriteRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  readLength,
		Items:       items,
		Data:        data,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadWriteRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadWriteRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.IndexGroup)
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.IndexOffset)
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (readLength)
		readLength := uint32(m.ReadLength)
		_readLengthErr := writeBuffer.WriteUint32("readLength", 32, (readLength))
		if _readLengthErr != nil {
			return errors.Wrap(_readLengthErr, "Error serializing 'readLength' field")
		}

		// Implicit Field (writeLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		writeLength := uint32(uint32(uint32(uint32(uint32(len(m.Items)))*uint32(uint32(utils.InlineIf(bool(bool((m.IndexGroup) == (61570))), func() uint16 { return uint16(uint32(16)) }, func() uint16 { return uint16(uint32(12)) }))))) + uint32(uint32(len(m.Data))))
		_writeLengthErr := writeBuffer.WriteUint32("writeLength", 32, (writeLength))
		if _writeLengthErr != nil {
			return errors.Wrap(_writeLengthErr, "Error serializing 'writeLength' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("AdsReadWriteRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsReadWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
