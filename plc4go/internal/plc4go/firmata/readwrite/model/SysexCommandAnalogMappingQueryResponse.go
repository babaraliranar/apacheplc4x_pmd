//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type SysexCommandAnalogMappingQueryResponse struct {
	Pin    uint8
	Parent *SysexCommand
}

// The corresponding interface
type ISysexCommandAnalogMappingQueryResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandAnalogMappingQueryResponse) CommandType() uint8 {
	return 0x69
}

func (m *SysexCommandAnalogMappingQueryResponse) Response() bool {
	return true
}

func (m *SysexCommandAnalogMappingQueryResponse) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandAnalogMappingQueryResponse(pin uint8) *SysexCommand {
	child := &SysexCommandAnalogMappingQueryResponse{
		Pin:    pin,
		Parent: NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandAnalogMappingQueryResponse(structType interface{}) *SysexCommandAnalogMappingQueryResponse {
	castFunc := func(typ interface{}) *SysexCommandAnalogMappingQueryResponse {
		if casted, ok := typ.(SysexCommandAnalogMappingQueryResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandAnalogMappingQueryResponse); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandAnalogMappingQueryResponse(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandAnalogMappingQueryResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandAnalogMappingQueryResponse) GetTypeName() string {
	return "SysexCommandAnalogMappingQueryResponse"
}

func (m *SysexCommandAnalogMappingQueryResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandAnalogMappingQueryResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	return lengthInBits
}

func (m *SysexCommandAnalogMappingQueryResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandAnalogMappingQueryResponseParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingQueryResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pin)
	pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingQueryResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandAnalogMappingQueryResponse{
		Pin:    pin,
		Parent: &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandAnalogMappingQueryResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingQueryResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingQueryResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
