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
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AmsSerialFrame struct {
	MagicCookie        uint16
	TransmitterAddress int8
	ReceiverAddress    int8
	FragmentNumber     int8
	Length             int8
	Userdata           *AmsPacket
	Crc                uint16
}

// The corresponding interface
type IAmsSerialFrame interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

func NewAmsSerialFrame(magicCookie uint16, transmitterAddress int8, receiverAddress int8, fragmentNumber int8, length int8, userdata *AmsPacket, crc uint16) *AmsSerialFrame {
	return &AmsSerialFrame{MagicCookie: magicCookie, TransmitterAddress: transmitterAddress, ReceiverAddress: receiverAddress, FragmentNumber: fragmentNumber, Length: length, Userdata: userdata, Crc: crc}
}

func CastAmsSerialFrame(structType interface{}) *AmsSerialFrame {
	castFunc := func(typ interface{}) *AmsSerialFrame {
		if casted, ok := typ.(AmsSerialFrame); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsSerialFrame); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsSerialFrame) GetTypeName() string {
	return "AmsSerialFrame"
}

func (m *AmsSerialFrame) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (magicCookie)
	lengthInBits += 16

	// Simple field (transmitterAddress)
	lengthInBits += 8

	// Simple field (receiverAddress)
	lengthInBits += 8

	// Simple field (fragmentNumber)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 8

	// Simple field (userdata)
	lengthInBits += m.Userdata.LengthInBits()

	// Simple field (crc)
	lengthInBits += 16

	return lengthInBits
}

func (m *AmsSerialFrame) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AmsSerialFrameParse(io *utils.ReadBuffer) (*AmsSerialFrame, error) {

	// Simple Field (magicCookie)
	magicCookie, _magicCookieErr := io.ReadUint16(16)
	if _magicCookieErr != nil {
		return nil, errors.Wrap(_magicCookieErr, "Error parsing 'magicCookie' field")
	}

	// Simple Field (transmitterAddress)
	transmitterAddress, _transmitterAddressErr := io.ReadInt8(8)
	if _transmitterAddressErr != nil {
		return nil, errors.Wrap(_transmitterAddressErr, "Error parsing 'transmitterAddress' field")
	}

	// Simple Field (receiverAddress)
	receiverAddress, _receiverAddressErr := io.ReadInt8(8)
	if _receiverAddressErr != nil {
		return nil, errors.Wrap(_receiverAddressErr, "Error parsing 'receiverAddress' field")
	}

	// Simple Field (fragmentNumber)
	fragmentNumber, _fragmentNumberErr := io.ReadInt8(8)
	if _fragmentNumberErr != nil {
		return nil, errors.Wrap(_fragmentNumberErr, "Error parsing 'fragmentNumber' field")
	}

	// Simple Field (length)
	length, _lengthErr := io.ReadInt8(8)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (userdata)
	userdata, _userdataErr := AmsPacketParse(io)
	if _userdataErr != nil {
		return nil, errors.Wrap(_userdataErr, "Error parsing 'userdata' field")
	}

	// Simple Field (crc)
	crc, _crcErr := io.ReadUint16(16)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}

	// Create the instance
	return NewAmsSerialFrame(magicCookie, transmitterAddress, receiverAddress, fragmentNumber, length, userdata, crc), nil
}

func (m *AmsSerialFrame) Serialize(io utils.WriteBuffer) error {

	// Simple Field (magicCookie)
	magicCookie := uint16(m.MagicCookie)
	_magicCookieErr := io.WriteUint16(16, (magicCookie))
	if _magicCookieErr != nil {
		return errors.Wrap(_magicCookieErr, "Error serializing 'magicCookie' field")
	}

	// Simple Field (transmitterAddress)
	transmitterAddress := int8(m.TransmitterAddress)
	_transmitterAddressErr := io.WriteInt8(8, (transmitterAddress))
	if _transmitterAddressErr != nil {
		return errors.Wrap(_transmitterAddressErr, "Error serializing 'transmitterAddress' field")
	}

	// Simple Field (receiverAddress)
	receiverAddress := int8(m.ReceiverAddress)
	_receiverAddressErr := io.WriteInt8(8, (receiverAddress))
	if _receiverAddressErr != nil {
		return errors.Wrap(_receiverAddressErr, "Error serializing 'receiverAddress' field")
	}

	// Simple Field (fragmentNumber)
	fragmentNumber := int8(m.FragmentNumber)
	_fragmentNumberErr := io.WriteInt8(8, (fragmentNumber))
	if _fragmentNumberErr != nil {
		return errors.Wrap(_fragmentNumberErr, "Error serializing 'fragmentNumber' field")
	}

	// Simple Field (length)
	length := int8(m.Length)
	_lengthErr := io.WriteInt8(8, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (userdata)
	_userdataErr := m.Userdata.Serialize(io)
	if _userdataErr != nil {
		return errors.Wrap(_userdataErr, "Error serializing 'userdata' field")
	}

	// Simple Field (crc)
	crc := uint16(m.Crc)
	_crcErr := io.WriteUint16(16, (crc))
	if _crcErr != nil {
		return errors.Wrap(_crcErr, "Error serializing 'crc' field")
	}

	return nil
}

func (m *AmsSerialFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "magicCookie":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MagicCookie = data
			case "transmitterAddress":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransmitterAddress = data
			case "receiverAddress":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReceiverAddress = data
			case "fragmentNumber":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.FragmentNumber = data
			case "length":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Length = data
			case "userdata":
				var data AmsPacket
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Userdata = &data
			case "crc":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Crc = data
			}
		}
	}
}

func (m *AmsSerialFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.ads.readwrite.AmsSerialFrame"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MagicCookie, xml.StartElement{Name: xml.Name{Local: "magicCookie"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransmitterAddress, xml.StartElement{Name: xml.Name{Local: "transmitterAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReceiverAddress, xml.StartElement{Name: xml.Name{Local: "receiverAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.FragmentNumber, xml.StartElement{Name: xml.Name{Local: "fragmentNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Length, xml.StartElement{Name: xml.Name{Local: "length"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Userdata, xml.StartElement{Name: xml.Name{Local: "userdata"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Crc, xml.StartElement{Name: xml.Name{Local: "crc"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m AmsSerialFrame) String() string {
	return string(m.Box("AmsSerialFrame", utils.DefaultWidth*2))
}

func (m AmsSerialFrame) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "AmsSerialFrame"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("MagicCookie", m.MagicCookie, width-2))
	boxes = append(boxes, utils.BoxAnything("TransmitterAddress", m.TransmitterAddress, width-2))
	boxes = append(boxes, utils.BoxAnything("ReceiverAddress", m.ReceiverAddress, width-2))
	boxes = append(boxes, utils.BoxAnything("FragmentNumber", m.FragmentNumber, width-2))
	boxes = append(boxes, utils.BoxAnything("Length", m.Length, width-2))
	boxes = append(boxes, utils.BoxAnything("Userdata", m.Userdata, width-2))
	boxes = append(boxes, utils.BoxAnything("Crc", m.Crc, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
