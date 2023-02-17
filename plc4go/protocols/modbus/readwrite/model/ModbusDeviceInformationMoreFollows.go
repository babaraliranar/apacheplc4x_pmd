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
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusDeviceInformationMoreFollows is an enum
type ModbusDeviceInformationMoreFollows uint8

type IModbusDeviceInformationMoreFollows interface {
	utils.Serializable
}

const(
	ModbusDeviceInformationMoreFollows_NO_MORE_OBJECTS_AVAILABLE ModbusDeviceInformationMoreFollows = 0x00
	ModbusDeviceInformationMoreFollows_MORE_OBJECTS_AVAILABLE ModbusDeviceInformationMoreFollows = 0xFF
)

var ModbusDeviceInformationMoreFollowsValues []ModbusDeviceInformationMoreFollows

func init() {
	_ = errors.New
	ModbusDeviceInformationMoreFollowsValues = []ModbusDeviceInformationMoreFollows {
		ModbusDeviceInformationMoreFollows_NO_MORE_OBJECTS_AVAILABLE,
		ModbusDeviceInformationMoreFollows_MORE_OBJECTS_AVAILABLE,
	}
}

func ModbusDeviceInformationMoreFollowsByValue(value uint8) (enum ModbusDeviceInformationMoreFollows, ok bool) {
	switch value {
		case 0x00:
			return ModbusDeviceInformationMoreFollows_NO_MORE_OBJECTS_AVAILABLE, true
		case 0xFF:
			return ModbusDeviceInformationMoreFollows_MORE_OBJECTS_AVAILABLE, true
	}
	return 0, false
}

func ModbusDeviceInformationMoreFollowsByName(value string) (enum ModbusDeviceInformationMoreFollows, ok bool) {
	switch value {
	case "NO_MORE_OBJECTS_AVAILABLE":
		return ModbusDeviceInformationMoreFollows_NO_MORE_OBJECTS_AVAILABLE, true
	case "MORE_OBJECTS_AVAILABLE":
		return ModbusDeviceInformationMoreFollows_MORE_OBJECTS_AVAILABLE, true
	}
	return 0, false
}

func ModbusDeviceInformationMoreFollowsKnows(value uint8)  bool {
	for _, typeValue := range ModbusDeviceInformationMoreFollowsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastModbusDeviceInformationMoreFollows(structType interface{}) ModbusDeviceInformationMoreFollows {
	castFunc := func(typ interface{}) ModbusDeviceInformationMoreFollows {
		if sModbusDeviceInformationMoreFollows, ok := typ.(ModbusDeviceInformationMoreFollows); ok {
			return sModbusDeviceInformationMoreFollows
		}
		return 0
	}
	return castFunc(structType)
}

func (m ModbusDeviceInformationMoreFollows) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ModbusDeviceInformationMoreFollows) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusDeviceInformationMoreFollowsParse(ctx context.Context, theBytes []byte) (ModbusDeviceInformationMoreFollows, error) {
	return ModbusDeviceInformationMoreFollowsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusDeviceInformationMoreFollowsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusDeviceInformationMoreFollows, error) {
	val, err := readBuffer.ReadUint8("ModbusDeviceInformationMoreFollows", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ModbusDeviceInformationMoreFollows")
	}
	if enum, ok := ModbusDeviceInformationMoreFollowsByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ModbusDeviceInformationMoreFollows(val), nil
	} else {
		return enum, nil
	}
}

func (e ModbusDeviceInformationMoreFollows) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ModbusDeviceInformationMoreFollows) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ModbusDeviceInformationMoreFollows", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ModbusDeviceInformationMoreFollows) PLC4XEnumName() string {
	switch e {
	case ModbusDeviceInformationMoreFollows_NO_MORE_OBJECTS_AVAILABLE:
		return "NO_MORE_OBJECTS_AVAILABLE"
	case ModbusDeviceInformationMoreFollows_MORE_OBJECTS_AVAILABLE:
		return "MORE_OBJECTS_AVAILABLE"
	}
	return ""
}

func (e ModbusDeviceInformationMoreFollows) String() string {
	return e.PLC4XEnumName()
}

