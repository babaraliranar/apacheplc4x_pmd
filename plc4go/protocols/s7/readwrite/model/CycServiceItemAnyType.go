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

// CycServiceItemAnyType is the corresponding interface of CycServiceItemAnyType
type CycServiceItemAnyType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CycServiceItemType
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() TransportSize
	// GetLength returns Length (property field)
	GetLength() uint16
	// GetDbNumber returns DbNumber (property field)
	GetDbNumber() uint16
	// GetMemoryArea returns MemoryArea (property field)
	GetMemoryArea() MemoryArea
	// GetAddress returns Address (property field)
	GetAddress() uint32
}

// CycServiceItemAnyTypeExactly can be used when we want exactly this type and not a type which fulfills CycServiceItemAnyType.
// This is useful for switch cases.
type CycServiceItemAnyTypeExactly interface {
	CycServiceItemAnyType
	isCycServiceItemAnyType() bool
}

// _CycServiceItemAnyType is the data-structure of this message
type _CycServiceItemAnyType struct {
	*_CycServiceItemType
	TransportSize TransportSize
	Length        uint16
	DbNumber      uint16
	MemoryArea    MemoryArea
	Address       uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CycServiceItemAnyType) InitializeParent(parent CycServiceItemType, byteLength uint8, syntaxId uint8) {
	m.ByteLength = byteLength
	m.SyntaxId = syntaxId
}

func (m *_CycServiceItemAnyType) GetParent() CycServiceItemType {
	return m._CycServiceItemType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CycServiceItemAnyType) GetTransportSize() TransportSize {
	return m.TransportSize
}

func (m *_CycServiceItemAnyType) GetLength() uint16 {
	return m.Length
}

func (m *_CycServiceItemAnyType) GetDbNumber() uint16 {
	return m.DbNumber
}

func (m *_CycServiceItemAnyType) GetMemoryArea() MemoryArea {
	return m.MemoryArea
}

func (m *_CycServiceItemAnyType) GetAddress() uint32 {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCycServiceItemAnyType factory function for _CycServiceItemAnyType
func NewCycServiceItemAnyType(transportSize TransportSize, length uint16, dbNumber uint16, memoryArea MemoryArea, address uint32, byteLength uint8, syntaxId uint8) *_CycServiceItemAnyType {
	_result := &_CycServiceItemAnyType{
		TransportSize:       transportSize,
		Length:              length,
		DbNumber:            dbNumber,
		MemoryArea:          memoryArea,
		Address:             address,
		_CycServiceItemType: NewCycServiceItemType(byteLength, syntaxId),
	}
	_result._CycServiceItemType._CycServiceItemTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCycServiceItemAnyType(structType any) CycServiceItemAnyType {
	if casted, ok := structType.(CycServiceItemAnyType); ok {
		return casted
	}
	if casted, ok := structType.(*CycServiceItemAnyType); ok {
		return *casted
	}
	return nil
}

func (m *_CycServiceItemAnyType) GetTypeName() string {
	return "CycServiceItemAnyType"
}

func (m *_CycServiceItemAnyType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Enum Field (transportSize)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 16

	// Simple field (dbNumber)
	lengthInBits += 16

	// Simple field (memoryArea)
	lengthInBits += 8

	// Simple field (address)
	lengthInBits += 24

	return lengthInBits
}

func (m *_CycServiceItemAnyType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CycServiceItemAnyTypeParse(ctx context.Context, theBytes []byte) (CycServiceItemAnyType, error) {
	return CycServiceItemAnyTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CycServiceItemAnyTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CycServiceItemAnyType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CycServiceItemAnyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CycServiceItemAnyType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	// Enum field (transportSize)
	transportSizeCode, _transportSizeCodeErr := readBuffer.ReadUint8("TransportSize", 8)
	if _transportSizeCodeErr != nil {
		return nil, errors.Wrap(_transportSizeCodeErr, "Error serializing 'transportSize' field")
	}
	transportSize, _transportSizeErr := TransportSizeFirstEnumForFieldCode(transportSizeCode)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadUint16("length", 16)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of CycServiceItemAnyType")
	}
	length := _length

	// Simple Field (dbNumber)
	_dbNumber, _dbNumberErr := readBuffer.ReadUint16("dbNumber", 16)
	if _dbNumberErr != nil {
		return nil, errors.Wrap(_dbNumberErr, "Error parsing 'dbNumber' field of CycServiceItemAnyType")
	}
	dbNumber := _dbNumber

	// Simple Field (memoryArea)
	if pullErr := readBuffer.PullContext("memoryArea"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for memoryArea")
	}
	_memoryArea, _memoryAreaErr := MemoryAreaParseWithBuffer(ctx, readBuffer)
	if _memoryAreaErr != nil {
		return nil, errors.Wrap(_memoryAreaErr, "Error parsing 'memoryArea' field of CycServiceItemAnyType")
	}
	memoryArea := _memoryArea
	if closeErr := readBuffer.CloseContext("memoryArea"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for memoryArea")
	}

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint32("address", 24)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of CycServiceItemAnyType")
	}
	address := _address

	if closeErr := readBuffer.CloseContext("CycServiceItemAnyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CycServiceItemAnyType")
	}

	// Create a partially initialized instance
	_child := &_CycServiceItemAnyType{
		_CycServiceItemType: &_CycServiceItemType{},
		TransportSize:       transportSize,
		Length:              length,
		DbNumber:            dbNumber,
		MemoryArea:          memoryArea,
		Address:             address,
	}
	_child._CycServiceItemType._CycServiceItemTypeChildRequirements = _child
	return _child, nil
}

func (m *_CycServiceItemAnyType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CycServiceItemAnyType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CycServiceItemAnyType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CycServiceItemAnyType")
		}

		if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportSize")
		}
		// Enum field (transportSize)
		_transportSizeErr := writeBuffer.WriteUint8("TransportSize", 8, uint8(m.TransportSize.Code()), utils.WithAdditionalStringRepresentation(m.GetTransportSize().PLC4XEnumName()))
		if _transportSizeErr != nil {
			return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
		}
		if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportSize")
		}

		// Simple Field (length)
		length := uint16(m.GetLength())
		_lengthErr := writeBuffer.WriteUint16("length", 16, uint16((length)))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (dbNumber)
		dbNumber := uint16(m.GetDbNumber())
		_dbNumberErr := writeBuffer.WriteUint16("dbNumber", 16, uint16((dbNumber)))
		if _dbNumberErr != nil {
			return errors.Wrap(_dbNumberErr, "Error serializing 'dbNumber' field")
		}

		// Simple Field (memoryArea)
		if pushErr := writeBuffer.PushContext("memoryArea"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for memoryArea")
		}
		_memoryAreaErr := writeBuffer.WriteSerializable(ctx, m.GetMemoryArea())
		if popErr := writeBuffer.PopContext("memoryArea"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for memoryArea")
		}
		if _memoryAreaErr != nil {
			return errors.Wrap(_memoryAreaErr, "Error serializing 'memoryArea' field")
		}

		// Simple Field (address)
		address := uint32(m.GetAddress())
		_addressErr := writeBuffer.WriteUint32("address", 24, uint32((address)))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("CycServiceItemAnyType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CycServiceItemAnyType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CycServiceItemAnyType) isCycServiceItemAnyType() bool {
	return true
}

func (m *_CycServiceItemAnyType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
