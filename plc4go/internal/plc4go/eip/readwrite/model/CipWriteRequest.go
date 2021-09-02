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
type CipWriteRequest struct {
	RequestPathSize int8
	Tag             []int8
	DataType        CIPDataTypeCode
	ElementNb       uint16
	Data            []int8
	Parent          *CipService
}

// The corresponding interface
type ICipWriteRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CipWriteRequest) Service() uint8 {
	return 0x4D
}

func (m *CipWriteRequest) InitializeParent(parent *CipService) {
}

func NewCipWriteRequest(RequestPathSize int8, tag []int8, dataType CIPDataTypeCode, elementNb uint16, data []int8) *CipService {
	child := &CipWriteRequest{
		RequestPathSize: RequestPathSize,
		Tag:             tag,
		DataType:        dataType,
		ElementNb:       elementNb,
		Data:            data,
		Parent:          NewCipService(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCipWriteRequest(structType interface{}) *CipWriteRequest {
	castFunc := func(typ interface{}) *CipWriteRequest {
		if casted, ok := typ.(CipWriteRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*CipWriteRequest); ok {
			return casted
		}
		if casted, ok := typ.(CipService); ok {
			return CastCipWriteRequest(casted.Child)
		}
		if casted, ok := typ.(*CipService); ok {
			return CastCipWriteRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CipWriteRequest) GetTypeName() string {
	return "CipWriteRequest"
}

func (m *CipWriteRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CipWriteRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (RequestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.Tag) > 0 {
		lengthInBits += 8 * uint16(len(m.Tag))
	}

	// Enum Field (dataType)
	lengthInBits += 16

	// Simple field (elementNb)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *CipWriteRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CipWriteRequestParse(readBuffer utils.ReadBuffer) (*CipService, error) {
	if pullErr := readBuffer.PullContext("CipWriteRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (RequestPathSize)
	RequestPathSize, _RequestPathSizeErr := readBuffer.ReadInt8("RequestPathSize", 8)
	if _RequestPathSizeErr != nil {
		return nil, errors.Wrap(_RequestPathSizeErr, "Error parsing 'RequestPathSize' field")
	}

	// Array field (tag)
	if pullErr := readBuffer.PullContext("tag", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	tag := make([]int8, 0)
	_tagLength := uint16(uint16(RequestPathSize) * uint16(uint16(2)))
	_tagEndPos := readBuffer.GetPos() + uint16(_tagLength)
	for readBuffer.GetPos() < _tagEndPos {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'tag' field")
		}
		tag = append(tag, _item)
	}
	if closeErr := readBuffer.CloseContext("tag", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("dataType"); pullErr != nil {
		return nil, pullErr
	}
	// Enum field (dataType)
	dataType, _dataTypeErr := CIPDataTypeCodeParse(readBuffer)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field")
	}
	if closeErr := readBuffer.CloseContext("dataType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (elementNb)
	elementNb, _elementNbErr := readBuffer.ReadUint16("elementNb", 16)
	if _elementNbErr != nil {
		return nil, errors.Wrap(_elementNbErr, "Error parsing 'elementNb' field")
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]int8, 0)
	_dataLength := uint16(dataType.Size()) * uint16(elementNb)
	_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
	for readBuffer.GetPos() < _dataEndPos {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data = append(data, _item)
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CipWriteRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipWriteRequest{
		RequestPathSize: RequestPathSize,
		Tag:             tag,
		DataType:        dataType,
		ElementNb:       elementNb,
		Data:            data,
		Parent:          &CipService{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *CipWriteRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (RequestPathSize)
		RequestPathSize := int8(m.RequestPathSize)
		_RequestPathSizeErr := writeBuffer.WriteInt8("RequestPathSize", 8, (RequestPathSize))
		if _RequestPathSizeErr != nil {
			return errors.Wrap(_RequestPathSizeErr, "Error serializing 'RequestPathSize' field")
		}

		// Array Field (tag)
		if m.Tag != nil {
			if pushErr := writeBuffer.PushContext("tag", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Tag {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'tag' field")
				}
			}
			if popErr := writeBuffer.PopContext("tag", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if pushErr := writeBuffer.PushContext("dataType"); pushErr != nil {
			return pushErr
		}
		// Enum field (dataType)
		dataType := CastCIPDataTypeCode(m.DataType)
		_dataTypeErr := dataType.Serialize(writeBuffer)
		if _dataTypeErr != nil {
			return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
		}
		if popErr := writeBuffer.PopContext("dataType"); popErr != nil {
			return popErr
		}

		// Simple Field (elementNb)
		elementNb := uint16(m.ElementNb)
		_elementNbErr := writeBuffer.WriteUint16("elementNb", 16, (elementNb))
		if _elementNbErr != nil {
			return errors.Wrap(_elementNbErr, "Error serializing 'elementNb' field")
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

		if popErr := writeBuffer.PopContext("CipWriteRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *CipWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
