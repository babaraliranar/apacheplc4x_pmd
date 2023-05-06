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

// CustomManufacturer is the corresponding interface of CustomManufacturer
type CustomManufacturer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCustomString returns CustomString (property field)
	GetCustomString() string
}

// CustomManufacturerExactly can be used when we want exactly this type and not a type which fulfills CustomManufacturer.
// This is useful for switch cases.
type CustomManufacturerExactly interface {
	CustomManufacturer
	isCustomManufacturer() bool
}

// _CustomManufacturer is the data-structure of this message
type _CustomManufacturer struct {
	CustomString string

	// Arguments.
	NumBytes uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CustomManufacturer) GetCustomString() string {
	return m.CustomString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCustomManufacturer factory function for _CustomManufacturer
func NewCustomManufacturer(customString string, numBytes uint8) *_CustomManufacturer {
	return &_CustomManufacturer{CustomString: customString, NumBytes: numBytes}
}

// Deprecated: use the interface for direct cast
func CastCustomManufacturer(structType any) CustomManufacturer {
	if casted, ok := structType.(CustomManufacturer); ok {
		return casted
	}
	if casted, ok := structType.(*CustomManufacturer); ok {
		return *casted
	}
	return nil
}

func (m *_CustomManufacturer) GetTypeName() string {
	return "CustomManufacturer"
}

func (m *_CustomManufacturer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (customString)
	lengthInBits += uint16(int32(int32(8)) * int32(m.NumBytes))

	return lengthInBits
}

func (m *_CustomManufacturer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CustomManufacturerParse(theBytes []byte, numBytes uint8) (CustomManufacturer, error) {
	return CustomManufacturerParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), numBytes)
}

func CustomManufacturerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, numBytes uint8) (CustomManufacturer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CustomManufacturer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CustomManufacturer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (customString)
	_customString, _customStringErr := readBuffer.ReadString("customString", uint32((8)*(numBytes)), "UTF-8")
	if _customStringErr != nil {
		return nil, errors.Wrap(_customStringErr, "Error parsing 'customString' field of CustomManufacturer")
	}
	customString := _customString

	if closeErr := readBuffer.CloseContext("CustomManufacturer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CustomManufacturer")
	}

	// Create the instance
	return &_CustomManufacturer{
		NumBytes:     numBytes,
		CustomString: customString,
	}, nil
}

func (m *_CustomManufacturer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CustomManufacturer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CustomManufacturer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CustomManufacturer")
	}

	// Simple Field (customString)
	customString := string(m.GetCustomString())
	_customStringErr := writeBuffer.WriteString("customString", uint32((8)*(m.GetNumBytes())), "UTF-8", (customString))
	if _customStringErr != nil {
		return errors.Wrap(_customStringErr, "Error serializing 'customString' field")
	}

	if popErr := writeBuffer.PopContext("CustomManufacturer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CustomManufacturer")
	}
	return nil
}

////
// Arguments Getter

func (m *_CustomManufacturer) GetNumBytes() uint8 {
	return m.NumBytes
}

//
////

func (m *_CustomManufacturer) isCustomManufacturer() bool {
	return true
}

func (m *_CustomManufacturer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
