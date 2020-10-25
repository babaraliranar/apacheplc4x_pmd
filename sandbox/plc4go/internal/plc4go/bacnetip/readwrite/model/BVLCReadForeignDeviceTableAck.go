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
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BVLCReadForeignDeviceTableAck struct {
    BVLC
}

// The corresponding interface
type IBVLCReadForeignDeviceTableAck interface {
    IBVLC
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BVLCReadForeignDeviceTableAck) BvlcFunction() uint8 {
    return 0x07
}

func (m BVLCReadForeignDeviceTableAck) initialize() spi.Message {
    return m
}

func NewBVLCReadForeignDeviceTableAck() BVLCInitializer {
    return &BVLCReadForeignDeviceTableAck{}
}

func CastIBVLCReadForeignDeviceTableAck(structType interface{}) IBVLCReadForeignDeviceTableAck {
    castFunc := func(typ interface{}) IBVLCReadForeignDeviceTableAck {
        if iBVLCReadForeignDeviceTableAck, ok := typ.(IBVLCReadForeignDeviceTableAck); ok {
            return iBVLCReadForeignDeviceTableAck
        }
        return nil
    }
    return castFunc(structType)
}

func CastBVLCReadForeignDeviceTableAck(structType interface{}) BVLCReadForeignDeviceTableAck {
    castFunc := func(typ interface{}) BVLCReadForeignDeviceTableAck {
        if sBVLCReadForeignDeviceTableAck, ok := typ.(BVLCReadForeignDeviceTableAck); ok {
            return sBVLCReadForeignDeviceTableAck
        }
        if sBVLCReadForeignDeviceTableAck, ok := typ.(*BVLCReadForeignDeviceTableAck); ok {
            return *sBVLCReadForeignDeviceTableAck
        }
        return BVLCReadForeignDeviceTableAck{}
    }
    return castFunc(structType)
}

func (m BVLCReadForeignDeviceTableAck) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BVLC.LengthInBits()

    return lengthInBits
}

func (m BVLCReadForeignDeviceTableAck) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BVLCReadForeignDeviceTableAckParse(io *utils.ReadBuffer) (BVLCInitializer, error) {

    // Create the instance
    return NewBVLCReadForeignDeviceTableAck(), nil
}

func (m BVLCReadForeignDeviceTableAck) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BVLCSerialize(io, m.BVLC, CastIBVLC(m), ser)
}

func (m *BVLCReadForeignDeviceTableAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            }
        }
    }
}

func (m BVLCReadForeignDeviceTableAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTableAck"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

