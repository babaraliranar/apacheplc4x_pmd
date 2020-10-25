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
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type ModbusPDUWriteSingleRegisterRequest struct {
    Address uint16
    Value uint16
    ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteSingleRegisterRequest interface {
    IModbusPDU
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m ModbusPDUWriteSingleRegisterRequest) ErrorFlag() bool {
    return false
}

func (m ModbusPDUWriteSingleRegisterRequest) FunctionFlag() uint8 {
    return 0x06
}

func (m ModbusPDUWriteSingleRegisterRequest) Response() bool {
    return false
}

func (m ModbusPDUWriteSingleRegisterRequest) initialize() spi.Message {
    return m
}

func NewModbusPDUWriteSingleRegisterRequest(address uint16, value uint16) ModbusPDUInitializer {
    return &ModbusPDUWriteSingleRegisterRequest{Address: address, Value: value}
}

func CastIModbusPDUWriteSingleRegisterRequest(structType interface{}) IModbusPDUWriteSingleRegisterRequest {
    castFunc := func(typ interface{}) IModbusPDUWriteSingleRegisterRequest {
        if iModbusPDUWriteSingleRegisterRequest, ok := typ.(IModbusPDUWriteSingleRegisterRequest); ok {
            return iModbusPDUWriteSingleRegisterRequest
        }
        return nil
    }
    return castFunc(structType)
}

func CastModbusPDUWriteSingleRegisterRequest(structType interface{}) ModbusPDUWriteSingleRegisterRequest {
    castFunc := func(typ interface{}) ModbusPDUWriteSingleRegisterRequest {
        if sModbusPDUWriteSingleRegisterRequest, ok := typ.(ModbusPDUWriteSingleRegisterRequest); ok {
            return sModbusPDUWriteSingleRegisterRequest
        }
        if sModbusPDUWriteSingleRegisterRequest, ok := typ.(*ModbusPDUWriteSingleRegisterRequest); ok {
            return *sModbusPDUWriteSingleRegisterRequest
        }
        return ModbusPDUWriteSingleRegisterRequest{}
    }
    return castFunc(structType)
}

func (m ModbusPDUWriteSingleRegisterRequest) LengthInBits() uint16 {
    var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

    // Simple field (address)
    lengthInBits += 16

    // Simple field (value)
    lengthInBits += 16

    return lengthInBits
}

func (m ModbusPDUWriteSingleRegisterRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUWriteSingleRegisterRequestParse(io *utils.ReadBuffer) (ModbusPDUInitializer, error) {

    // Simple Field (address)
    address, _addressErr := io.ReadUint16(16)
    if _addressErr != nil {
        return nil, errors.New("Error parsing 'address' field " + _addressErr.Error())
    }

    // Simple Field (value)
    value, _valueErr := io.ReadUint16(16)
    if _valueErr != nil {
        return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
    }

    // Create the instance
    return NewModbusPDUWriteSingleRegisterRequest(address, value), nil
}

func (m ModbusPDUWriteSingleRegisterRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (address)
    address := uint16(m.Address)
    _addressErr := io.WriteUint16(16, (address))
    if _addressErr != nil {
        return errors.New("Error serializing 'address' field " + _addressErr.Error())
    }

    // Simple Field (value)
    value := uint16(m.Value)
    _valueErr := io.WriteUint16(16, (value))
    if _valueErr != nil {
        return errors.New("Error serializing 'value' field " + _valueErr.Error())
    }

        return nil
    }
    return ModbusPDUSerialize(io, m.ModbusPDU, CastIModbusPDU(m), ser)
}

func (m *ModbusPDUWriteSingleRegisterRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
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
            case "address":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Address = data
            case "value":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Value = data
            }
        }
    }
}

func (m ModbusPDUWriteSingleRegisterRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteSingleRegisterRequest"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

