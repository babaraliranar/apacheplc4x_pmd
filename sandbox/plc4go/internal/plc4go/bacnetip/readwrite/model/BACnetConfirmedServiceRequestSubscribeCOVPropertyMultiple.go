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
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple struct {
    BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple interface {
    IBACnetConfirmedServiceRequest
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) ServiceChoice() uint8 {
    return 0x1E
}

func (m BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) initialize() spi.Message {
    return m
}

func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple() BACnetConfirmedServiceRequestInitializer {
    return &BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{}
}

func CastIBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(structType interface{}) IBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
    castFunc := func(typ interface{}) IBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
        if iBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, ok := typ.(IBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
            return iBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(structType interface{}) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
    castFunc := func(typ interface{}) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
        if sBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, ok := typ.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
            return sBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
        }
        if sBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, ok := typ.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
            return *sBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
        }
        return BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{}
    }
    return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetConfirmedServiceRequest.LengthInBits()

    return lengthInBits
}

func (m BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleParse(io *utils.ReadBuffer) (BACnetConfirmedServiceRequestInitializer, error) {

    // Create the instance
    return NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(), nil
}

func (m BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetConfirmedServiceRequestSerialize(io, m.BACnetConfirmedServiceRequest, CastIBACnetConfirmedServiceRequest(m), ser)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

