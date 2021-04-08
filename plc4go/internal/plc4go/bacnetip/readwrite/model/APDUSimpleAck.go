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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type APDUSimpleAck struct {
	OriginalInvokeId uint8
	ServiceChoice    uint8
	Parent           *APDU
}

// The corresponding interface
type IAPDUSimpleAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUSimpleAck) ApduType() uint8 {
	return 0x2
}

func (m *APDUSimpleAck) InitializeParent(parent *APDU) {
}

func NewAPDUSimpleAck(originalInvokeId uint8, serviceChoice uint8) *APDU {
	child := &APDUSimpleAck{
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
		Parent:           NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUSimpleAck(structType interface{}) *APDUSimpleAck {
	castFunc := func(typ interface{}) *APDUSimpleAck {
		if casted, ok := typ.(APDUSimpleAck); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUSimpleAck); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUSimpleAck(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUSimpleAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUSimpleAck) GetTypeName() string {
	return "APDUSimpleAck"
}

func (m *APDUSimpleAck) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *APDUSimpleAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUSimpleAckParse(io *utils.ReadBuffer) (*APDU, error) {

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8(4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := io.ReadUint8(8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}

	// Simple Field (serviceChoice)
	serviceChoice, _serviceChoiceErr := io.ReadUint8(8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Create a partially initialized instance
	_child := &APDUSimpleAck{
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
		Parent:           &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUSimpleAck) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8(4, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := io.WriteUint8(8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (serviceChoice)
		serviceChoice := uint8(m.ServiceChoice)
		_serviceChoiceErr := io.WriteUint8(8, (serviceChoice))
		if _serviceChoiceErr != nil {
			return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *APDUSimpleAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "originalInvokeId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.OriginalInvokeId = data
			case "serviceChoice":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ServiceChoice = data
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

func (m *APDUSimpleAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.OriginalInvokeId, xml.StartElement{Name: xml.Name{Local: "originalInvokeId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ServiceChoice, xml.StartElement{Name: xml.Name{Local: "serviceChoice"}}); err != nil {
		return err
	}
	return nil
}

func (m APDUSimpleAck) String() string {
	return string(m.Box("APDUSimpleAck", utils.DefaultWidth*2))
}

func (m APDUSimpleAck) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "APDUSimpleAck"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("OriginalInvokeId", m.OriginalInvokeId, width-2))
	boxes = append(boxes, utils.BoxAnything("ServiceChoice", m.ServiceChoice, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
