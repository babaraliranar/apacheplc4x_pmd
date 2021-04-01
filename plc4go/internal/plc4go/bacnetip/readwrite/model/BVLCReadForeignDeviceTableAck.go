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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BVLCReadForeignDeviceTableAck struct {
	Parent *BVLC
}

// The corresponding interface
type IBVLCReadForeignDeviceTableAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCReadForeignDeviceTableAck) BvlcFunction() uint8 {
	return 0x07
}

func (m *BVLCReadForeignDeviceTableAck) InitializeParent(parent *BVLC) {
}

func NewBVLCReadForeignDeviceTableAck() *BVLC {
	child := &BVLCReadForeignDeviceTableAck{
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCReadForeignDeviceTableAck(structType interface{}) *BVLCReadForeignDeviceTableAck {
	castFunc := func(typ interface{}) *BVLCReadForeignDeviceTableAck {
		if casted, ok := typ.(BVLCReadForeignDeviceTableAck); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCReadForeignDeviceTableAck); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCReadForeignDeviceTableAck(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCReadForeignDeviceTableAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCReadForeignDeviceTableAck) GetTypeName() string {
	return "BVLCReadForeignDeviceTableAck"
}

func (m *BVLCReadForeignDeviceTableAck) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *BVLCReadForeignDeviceTableAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCReadForeignDeviceTableAckParse(io *utils.ReadBuffer) (*BVLC, error) {

	// Create a partially initialized instance
	_child := &BVLCReadForeignDeviceTableAck{
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCReadForeignDeviceTableAck) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BVLCReadForeignDeviceTableAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (m *BVLCReadForeignDeviceTableAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BVLCReadForeignDeviceTableAck) String() string {
	return string(m.Box("BVLCReadForeignDeviceTableAck", utils.DefaultWidth*2))
}

func (m BVLCReadForeignDeviceTableAck) Box(name string, width int) utils.AsciiBox {
	boxes := make([]utils.AsciiBox, 0)
	return utils.BoxString(name, string(utils.AlignBoxes(boxes, width-2)), width)
}
