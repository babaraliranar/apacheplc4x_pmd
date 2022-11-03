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

// AirConditioningCommandType is an enum
type AirConditioningCommandType uint8

type IAirConditioningCommandType interface {
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	AirConditioningCommandType_SET_ZONE_GROUP_OFF             AirConditioningCommandType = 0x00
	AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS         AirConditioningCommandType = 0x01
	AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS     AirConditioningCommandType = 0x02
	AirConditioningCommandType_ZONE_TEMPERATURE               AirConditioningCommandType = 0x03
	AirConditioningCommandType_ZONE_HUMIDITY                  AirConditioningCommandType = 0x04
	AirConditioningCommandType_REFRESH                        AirConditioningCommandType = 0x05
	AirConditioningCommandType_SET_ZONE_HVAC_MODE             AirConditioningCommandType = 0x06
	AirConditioningCommandType_SET_PLANT_HVAC_LEVEL           AirConditioningCommandType = 0x07
	AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE         AirConditioningCommandType = 0x08
	AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL       AirConditioningCommandType = 0x09
	AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT     AirConditioningCommandType = 0x0A
	AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT     AirConditioningCommandType = 0x0B
	AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT         AirConditioningCommandType = 0x0C
	AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT AirConditioningCommandType = 0x0D
	AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT AirConditioningCommandType = 0x0E
	AirConditioningCommandType_SET_ZONE_GROUP_ON              AirConditioningCommandType = 0x0F
	AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT     AirConditioningCommandType = 0x10
	AirConditioningCommandType_HVAC_SCHEDULE_ENTRY            AirConditioningCommandType = 0x11
	AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY        AirConditioningCommandType = 0x12
)

var AirConditioningCommandTypeValues []AirConditioningCommandType

func init() {
	_ = errors.New
	AirConditioningCommandTypeValues = []AirConditioningCommandType{
		AirConditioningCommandType_SET_ZONE_GROUP_OFF,
		AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS,
		AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS,
		AirConditioningCommandType_ZONE_TEMPERATURE,
		AirConditioningCommandType_ZONE_HUMIDITY,
		AirConditioningCommandType_REFRESH,
		AirConditioningCommandType_SET_ZONE_HVAC_MODE,
		AirConditioningCommandType_SET_PLANT_HVAC_LEVEL,
		AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE,
		AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL,
		AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT,
		AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT,
		AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT,
		AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT,
		AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT,
		AirConditioningCommandType_SET_ZONE_GROUP_ON,
		AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT,
		AirConditioningCommandType_HVAC_SCHEDULE_ENTRY,
		AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY,
	}
}

func (e AirConditioningCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0
		}
	case 0x01:
		{ /* '0x01' */
			return 5
		}
	case 0x02:
		{ /* '0x02' */
			return 5
		}
	case 0x03:
		{ /* '0x03' */
			return 4
		}
	case 0x04:
		{ /* '0x04' */
			return 4
		}
	case 0x05:
		{ /* '0x05' */
			return 1
		}
	case 0x06:
		{ /* '0x06' */
			return 5
		}
	case 0x07:
		{ /* '0x07' */
			return 5
		}
	case 0x08:
		{ /* '0x08' */
			return 5
		}
	case 0x09:
		{ /* '0x09' */
			return 5
		}
	case 0x0A:
		{ /* '0x0A' */
			return 4
		}
	case 0x0B:
		{ /* '0x0B' */
			return 4
		}
	case 0x0C:
		{ /* '0x0C' */
			return 4
		}
	case 0x0D:
		{ /* '0x0D' */
			return 4
		}
	case 0x0E:
		{ /* '0x0E' */
			return 4
		}
	case 0x0F:
		{ /* '0x0F' */
			return 1
		}
	case 0x10:
		{ /* '0x10' */
			return 1
		}
	case 0x11:
		{ /* '0x11' */
			return 7
		}
	case 0x12:
		{ /* '0x12' */
			return 7
		}
	default:
		{
			return 0
		}
	}
}

func AirConditioningCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (AirConditioningCommandType, error) {
	for _, sizeValue := range AirConditioningCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func AirConditioningCommandTypeByValue(value uint8) (enum AirConditioningCommandType, ok bool) {
	switch value {
	case 0x00:
		return AirConditioningCommandType_SET_ZONE_GROUP_OFF, true
	case 0x01:
		return AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS, true
	case 0x02:
		return AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS, true
	case 0x03:
		return AirConditioningCommandType_ZONE_TEMPERATURE, true
	case 0x04:
		return AirConditioningCommandType_ZONE_HUMIDITY, true
	case 0x05:
		return AirConditioningCommandType_REFRESH, true
	case 0x06:
		return AirConditioningCommandType_SET_ZONE_HVAC_MODE, true
	case 0x07:
		return AirConditioningCommandType_SET_PLANT_HVAC_LEVEL, true
	case 0x08:
		return AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE, true
	case 0x09:
		return AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL, true
	case 0x0A:
		return AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT, true
	case 0x0B:
		return AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT, true
	case 0x0C:
		return AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT, true
	case 0x0D:
		return AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT, true
	case 0x0E:
		return AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT, true
	case 0x0F:
		return AirConditioningCommandType_SET_ZONE_GROUP_ON, true
	case 0x10:
		return AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT, true
	case 0x11:
		return AirConditioningCommandType_HVAC_SCHEDULE_ENTRY, true
	case 0x12:
		return AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY, true
	}
	return 0, false
}

func AirConditioningCommandTypeByName(value string) (enum AirConditioningCommandType, ok bool) {
	switch value {
	case "SET_ZONE_GROUP_OFF":
		return AirConditioningCommandType_SET_ZONE_GROUP_OFF, true
	case "ZONE_HVAC_PLANT_STATUS":
		return AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS, true
	case "ZONE_HUMIDITY_PLANT_STATUS":
		return AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS, true
	case "ZONE_TEMPERATURE":
		return AirConditioningCommandType_ZONE_TEMPERATURE, true
	case "ZONE_HUMIDITY":
		return AirConditioningCommandType_ZONE_HUMIDITY, true
	case "REFRESH":
		return AirConditioningCommandType_REFRESH, true
	case "SET_ZONE_HVAC_MODE":
		return AirConditioningCommandType_SET_ZONE_HVAC_MODE, true
	case "SET_PLANT_HVAC_LEVEL":
		return AirConditioningCommandType_SET_PLANT_HVAC_LEVEL, true
	case "SET_ZONE_HUMIDITY_MODE":
		return AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE, true
	case "SET_PLANT_HUMIDITY_LEVEL":
		return AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL, true
	case "SET_HVAC_UPPER_GUARD_LIMIT":
		return AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT, true
	case "SET_HVAC_LOWER_GUARD_LIMIT":
		return AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT, true
	case "SET_HVAC_SETBACK_LIMIT":
		return AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT, true
	case "SET_HUMIDITY_UPPER_GUARD_LIMIT":
		return AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT, true
	case "SET_HUMIDITY_LOWER_GUARD_LIMIT":
		return AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT, true
	case "SET_ZONE_GROUP_ON":
		return AirConditioningCommandType_SET_ZONE_GROUP_ON, true
	case "SET_HUMIDITY_SETBACK_LIMIT":
		return AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT, true
	case "HVAC_SCHEDULE_ENTRY":
		return AirConditioningCommandType_HVAC_SCHEDULE_ENTRY, true
	case "HUMIDITY_SCHEDULE_ENTRY":
		return AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY, true
	}
	return 0, false
}

func AirConditioningCommandTypeKnows(value uint8) bool {
	for _, typeValue := range AirConditioningCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAirConditioningCommandType(structType interface{}) AirConditioningCommandType {
	castFunc := func(typ interface{}) AirConditioningCommandType {
		if sAirConditioningCommandType, ok := typ.(AirConditioningCommandType); ok {
			return sAirConditioningCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AirConditioningCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m AirConditioningCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningCommandTypeParse(theBytes []byte) (AirConditioningCommandType, error) {
	return AirConditioningCommandTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func AirConditioningCommandTypeParseWithBuffer(readBuffer utils.ReadBuffer) (AirConditioningCommandType, error) {
	val, err := readBuffer.ReadUint8("AirConditioningCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AirConditioningCommandType")
	}
	if enum, ok := AirConditioningCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AirConditioningCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e AirConditioningCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AirConditioningCommandType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AirConditioningCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AirConditioningCommandType) PLC4XEnumName() string {
	switch e {
	case AirConditioningCommandType_SET_ZONE_GROUP_OFF:
		return "SET_ZONE_GROUP_OFF"
	case AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS:
		return "ZONE_HVAC_PLANT_STATUS"
	case AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS:
		return "ZONE_HUMIDITY_PLANT_STATUS"
	case AirConditioningCommandType_ZONE_TEMPERATURE:
		return "ZONE_TEMPERATURE"
	case AirConditioningCommandType_ZONE_HUMIDITY:
		return "ZONE_HUMIDITY"
	case AirConditioningCommandType_REFRESH:
		return "REFRESH"
	case AirConditioningCommandType_SET_ZONE_HVAC_MODE:
		return "SET_ZONE_HVAC_MODE"
	case AirConditioningCommandType_SET_PLANT_HVAC_LEVEL:
		return "SET_PLANT_HVAC_LEVEL"
	case AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE:
		return "SET_ZONE_HUMIDITY_MODE"
	case AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL:
		return "SET_PLANT_HUMIDITY_LEVEL"
	case AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT:
		return "SET_HVAC_UPPER_GUARD_LIMIT"
	case AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT:
		return "SET_HVAC_LOWER_GUARD_LIMIT"
	case AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT:
		return "SET_HVAC_SETBACK_LIMIT"
	case AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT:
		return "SET_HUMIDITY_UPPER_GUARD_LIMIT"
	case AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT:
		return "SET_HUMIDITY_LOWER_GUARD_LIMIT"
	case AirConditioningCommandType_SET_ZONE_GROUP_ON:
		return "SET_ZONE_GROUP_ON"
	case AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT:
		return "SET_HUMIDITY_SETBACK_LIMIT"
	case AirConditioningCommandType_HVAC_SCHEDULE_ENTRY:
		return "HVAC_SCHEDULE_ENTRY"
	case AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY:
		return "HUMIDITY_SCHEDULE_ENTRY"
	}
	return ""
}

func (e AirConditioningCommandType) String() string {
	return e.PLC4XEnumName()
}
