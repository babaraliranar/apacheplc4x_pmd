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
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AccessControlCategory is an enum
type AccessControlCategory uint8

type IAccessControlCategory interface {
	utils.Serializable
}

const (
	AccessControlCategory_SYSTEM_ACTIVITY AccessControlCategory = 0x00
	AccessControlCategory_SYSTEM_REQUEST  AccessControlCategory = 0x01
)

var AccessControlCategoryValues []AccessControlCategory

func init() {
	_ = errors.New
	AccessControlCategoryValues = []AccessControlCategory{
		AccessControlCategory_SYSTEM_ACTIVITY,
		AccessControlCategory_SYSTEM_REQUEST,
	}
}

func AccessControlCategoryByValue(value uint8) (enum AccessControlCategory, ok bool) {
	switch value {
	case 0x00:
		return AccessControlCategory_SYSTEM_ACTIVITY, true
	case 0x01:
		return AccessControlCategory_SYSTEM_REQUEST, true
	}
	return 0, false
}

func AccessControlCategoryByName(value string) (enum AccessControlCategory, ok bool) {
	switch value {
	case "SYSTEM_ACTIVITY":
		return AccessControlCategory_SYSTEM_ACTIVITY, true
	case "SYSTEM_REQUEST":
		return AccessControlCategory_SYSTEM_REQUEST, true
	}
	return 0, false
}

func AccessControlCategoryKnows(value uint8) bool {
	for _, typeValue := range AccessControlCategoryValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAccessControlCategory(structType interface{}) AccessControlCategory {
	castFunc := func(typ interface{}) AccessControlCategory {
		if sAccessControlCategory, ok := typ.(AccessControlCategory); ok {
			return sAccessControlCategory
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessControlCategory) GetLengthInBits() uint16 {
	return 4
}

func (m AccessControlCategory) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AccessControlCategoryParse(theBytes []byte) (AccessControlCategory, error) {
	return AccessControlCategoryParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func AccessControlCategoryParseWithBuffer(readBuffer utils.ReadBuffer) (AccessControlCategory, error) {
	val, err := readBuffer.ReadUint8("AccessControlCategory", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessControlCategory")
	}
	if enum, ok := AccessControlCategoryByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AccessControlCategory(val), nil
	} else {
		return enum, nil
	}
}

func (e AccessControlCategory) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AccessControlCategory) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AccessControlCategory", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessControlCategory) PLC4XEnumName() string {
	switch e {
	case AccessControlCategory_SYSTEM_ACTIVITY:
		return "SYSTEM_ACTIVITY"
	case AccessControlCategory_SYSTEM_REQUEST:
		return "SYSTEM_REQUEST"
	}
	return ""
}

func (e AccessControlCategory) String() string {
	return e.PLC4XEnumName()
}
