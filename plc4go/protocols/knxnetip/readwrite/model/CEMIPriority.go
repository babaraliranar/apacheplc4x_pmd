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

// CEMIPriority is an enum
type CEMIPriority uint8

type ICEMIPriority interface {
	utils.Serializable
}

const (
	CEMIPriority_SYSTEM CEMIPriority = 0x0
	CEMIPriority_NORMAL CEMIPriority = 0x1
	CEMIPriority_URGENT CEMIPriority = 0x2
	CEMIPriority_LOW    CEMIPriority = 0x3
)

var CEMIPriorityValues []CEMIPriority

func init() {
	_ = errors.New
	CEMIPriorityValues = []CEMIPriority{
		CEMIPriority_SYSTEM,
		CEMIPriority_NORMAL,
		CEMIPriority_URGENT,
		CEMIPriority_LOW,
	}
}

func CEMIPriorityByValue(value uint8) (enum CEMIPriority, ok bool) {
	switch value {
	case 0x0:
		return CEMIPriority_SYSTEM, true
	case 0x1:
		return CEMIPriority_NORMAL, true
	case 0x2:
		return CEMIPriority_URGENT, true
	case 0x3:
		return CEMIPriority_LOW, true
	}
	return 0, false
}

func CEMIPriorityByName(value string) (enum CEMIPriority, ok bool) {
	switch value {
	case "SYSTEM":
		return CEMIPriority_SYSTEM, true
	case "NORMAL":
		return CEMIPriority_NORMAL, true
	case "URGENT":
		return CEMIPriority_URGENT, true
	case "LOW":
		return CEMIPriority_LOW, true
	}
	return 0, false
}

func CEMIPriorityKnows(value uint8) bool {
	for _, typeValue := range CEMIPriorityValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCEMIPriority(structType interface{}) CEMIPriority {
	castFunc := func(typ interface{}) CEMIPriority {
		if sCEMIPriority, ok := typ.(CEMIPriority); ok {
			return sCEMIPriority
		}
		return 0
	}
	return castFunc(structType)
}

func (m CEMIPriority) GetLengthInBits() uint16 {
	return 2
}

func (m CEMIPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIPriorityParse(theBytes []byte) (CEMIPriority, error) {
	return CEMIPriorityParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func CEMIPriorityParseWithBuffer(readBuffer utils.ReadBuffer) (CEMIPriority, error) {
	val, err := readBuffer.ReadUint8("CEMIPriority", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CEMIPriority")
	}
	if enum, ok := CEMIPriorityByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return CEMIPriority(val), nil
	} else {
		return enum, nil
	}
}

func (e CEMIPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CEMIPriority) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("CEMIPriority", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CEMIPriority) PLC4XEnumName() string {
	switch e {
	case CEMIPriority_SYSTEM:
		return "SYSTEM"
	case CEMIPriority_NORMAL:
		return "NORMAL"
	case CEMIPriority_URGENT:
		return "URGENT"
	case CEMIPriority_LOW:
		return "LOW"
	}
	return ""
}

func (e CEMIPriority) String() string {
	return e.PLC4XEnumName()
}
