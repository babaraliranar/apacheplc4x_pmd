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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
    "reflect"
    "strings"
)

// The data-structure of this message
type AdsMultiRequestItem struct {
    Child IAdsMultiRequestItemChild
    IAdsMultiRequestItem
    IAdsMultiRequestItemParent
}

// The corresponding interface
type IAdsMultiRequestItem interface {
    IndexGroup() uint32
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IAdsMultiRequestItemParent interface {
    SerializeParent(io utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error
    GetTypeName() string
}

type IAdsMultiRequestItemChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *AdsMultiRequestItem)
    GetTypeName() string
    IAdsMultiRequestItem
}

func NewAdsMultiRequestItem() *AdsMultiRequestItem {
    return &AdsMultiRequestItem{}
}

func CastAdsMultiRequestItem(structType interface{}) *AdsMultiRequestItem {
    castFunc := func(typ interface{}) *AdsMultiRequestItem {
        if casted, ok := typ.(AdsMultiRequestItem); ok {
            return &casted
        }
        if casted, ok := typ.(*AdsMultiRequestItem); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AdsMultiRequestItem) GetTypeName() string {
    return "AdsMultiRequestItem"
}

func (m *AdsMultiRequestItem) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *AdsMultiRequestItem) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AdsMultiRequestItemParse(io *utils.ReadBuffer, indexGroup uint32) (*AdsMultiRequestItem, error) {

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *AdsMultiRequestItem
    var typeSwitchError error
    switch {
    case indexGroup == 61568L:
        _parent, typeSwitchError = AdsMultiRequestItemReadParse(io)
    case indexGroup == 61569L:
        _parent, typeSwitchError = AdsMultiRequestItemWriteParse(io)
    case indexGroup == 61570L:
        _parent, typeSwitchError = AdsMultiRequestItemReadWriteParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *AdsMultiRequestItem) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *AdsMultiRequestItem) SerializeParent(io utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error {

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *AdsMultiRequestItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.ads.readwrite.AdsMultiRequestItemRead":
                        var dt *AdsMultiRequestItemRead
                        if m.Child != nil {
                            dt = m.Child.(*AdsMultiRequestItemRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.ads.readwrite.AdsMultiRequestItemWrite":
                        var dt *AdsMultiRequestItemWrite
                        if m.Child != nil {
                            dt = m.Child.(*AdsMultiRequestItemWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.ads.readwrite.AdsMultiRequestItemReadWrite":
                        var dt *AdsMultiRequestItemReadWrite
                        if m.Child != nil {
                            dt = m.Child.(*AdsMultiRequestItemReadWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                }
            }
        }
    }
}

func (m *AdsMultiRequestItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.ads.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    marshaller, ok := m.Child.(xml.Marshaler)
    if !ok {
        return errors.New("child is not castable to Marshaler")
    }
    if err := marshaller.MarshalXML(e, start); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

