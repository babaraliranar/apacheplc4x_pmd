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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
)

	// Code generated by code-generation. DO NOT EDIT.


// AssociatedValueType is the corresponding interface of AssociatedValueType
type AssociatedValueType interface {
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetValueLength returns ValueLength (property field)
	GetValueLength() uint16
	// GetData returns Data (property field)
	GetData() []uint8
}

// AssociatedValueTypeExactly can be used when we want exactly this type and not a type which fulfills AssociatedValueType.
// This is useful for switch cases.
type AssociatedValueTypeExactly interface {
	AssociatedValueType
	isAssociatedValueType() bool
}

// _AssociatedValueType is the data-structure of this message
type _AssociatedValueType struct {
        ReturnCode DataTransportErrorCode
        TransportSize DataTransportSize
        ValueLength uint16
        Data []uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AssociatedValueType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_AssociatedValueType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_AssociatedValueType) GetValueLength() uint16 {
	return m.ValueLength
}

func (m *_AssociatedValueType) GetData() []uint8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAssociatedValueType factory function for _AssociatedValueType
func NewAssociatedValueType( returnCode DataTransportErrorCode , transportSize DataTransportSize , valueLength uint16 , data []uint8 ) *_AssociatedValueType {
return &_AssociatedValueType{ ReturnCode: returnCode , TransportSize: transportSize , ValueLength: valueLength , Data: data }
}

// Deprecated: use the interface for direct cast
func CastAssociatedValueType(structType interface{}) AssociatedValueType {
    if casted, ok := structType.(AssociatedValueType); ok {
		return casted
	}
	if casted, ok := structType.(*AssociatedValueType); ok {
		return *casted
	}
	return nil
}

func (m *_AssociatedValueType) GetTypeName() string {
	return "AssociatedValueType"
}

func (m *_AssociatedValueType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Manual Field (valueLength)
	lengthInBits += uint16(int32(16))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}


func (m *_AssociatedValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AssociatedValueTypeParse(theBytes []byte) (AssociatedValueType, error) {
	return AssociatedValueTypeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AssociatedValueTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedValueType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AssociatedValueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AssociatedValueType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
_returnCode, _returnCodeErr := DataTransportErrorCodeParseWithBuffer(ctx, readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of AssociatedValueType")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
_transportSize, _transportSizeErr := DataTransportSizeParseWithBuffer(ctx, readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field of AssociatedValueType")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Manual Field (valueLength)
	_valueLength, _valueLengthErr := RightShift3(readBuffer)
	if _valueLengthErr != nil {
		return nil, errors.Wrap(_valueLengthErr, "Error parsing 'valueLength' field of AssociatedValueType")
	}
	var valueLength uint16
	if _valueLength != nil {
            valueLength = _valueLength.(uint16)
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Count array
	data := make([]uint8, EventItemLength(readBuffer, valueLength))
	// This happens when the size is set conditional to 0
	if len(data) == 0 {
		data = nil
	}
	{
		_numItems := uint16(EventItemLength(readBuffer, valueLength))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field of AssociatedValueType")
			}
			data[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("AssociatedValueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AssociatedValueType")
	}

	// Create the instance
	return &_AssociatedValueType{
			ReturnCode: returnCode,
			TransportSize: transportSize,
			ValueLength: valueLength,
			Data: data,
		}, nil
}

func (m *_AssociatedValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AssociatedValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("AssociatedValueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AssociatedValueType")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(ctx, m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transportSize")
	}
	_transportSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSize())
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transportSize")
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Manual Field (valueLength)
	_valueLengthErr := LeftShift3(writeBuffer, m.GetValueLength())
	if _valueLengthErr != nil {
		return errors.Wrap(_valueLengthErr, "Error serializing 'valueLength' field")
	}

	// Array Field (data)
	if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for data")
	}
	for _curItem, _element := range m.GetData() {
		_ = _curItem
		_elementErr := writeBuffer.WriteUint8("", 8, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'data' field")
		}
	}
	if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for data")
	}

	if popErr := writeBuffer.PopContext("AssociatedValueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AssociatedValueType")
	}
	return nil
}


func (m *_AssociatedValueType) isAssociatedValueType() bool {
	return true
}

func (m *_AssociatedValueType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



