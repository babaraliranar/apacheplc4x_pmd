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
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AlarmType is an enum
type AlarmType uint8

type IAlarmType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	AlarmType_SCAN    AlarmType = 0x01
	AlarmType_ALARM_8 AlarmType = 0x02
	AlarmType_ALARM_S AlarmType = 0x04
)

var AlarmTypeValues []AlarmType

func init() {
	_ = errors.New
	AlarmTypeValues = []AlarmType{
		AlarmType_SCAN,
		AlarmType_ALARM_8,
		AlarmType_ALARM_S,
	}
}

func AlarmTypeByValue(value uint8) (enum AlarmType, ok bool) {
	switch value {
	case 0x01:
		return AlarmType_SCAN, true
	case 0x02:
		return AlarmType_ALARM_8, true
	case 0x04:
		return AlarmType_ALARM_S, true
	}
	return 0, false
}

func AlarmTypeByName(value string) (enum AlarmType, ok bool) {
	switch value {
	case "SCAN":
		return AlarmType_SCAN, true
	case "ALARM_8":
		return AlarmType_ALARM_8, true
	case "ALARM_S":
		return AlarmType_ALARM_S, true
	}
	return 0, false
}

func AlarmTypeKnows(value uint8) bool {
	for _, typeValue := range AlarmTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAlarmType(structType any) AlarmType {
	castFunc := func(typ any) AlarmType {
		if sAlarmType, ok := typ.(AlarmType); ok {
			return sAlarmType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AlarmType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m AlarmType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmTypeParse(ctx context.Context, theBytes []byte) (AlarmType, error) {
	return AlarmTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("AlarmType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AlarmType")
	}
	if enum, ok := AlarmTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AlarmType")
		return AlarmType(val), nil
	} else {
		return enum, nil
	}
}

func (e AlarmType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AlarmType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("AlarmType", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AlarmType) PLC4XEnumName() string {
	switch e {
	case AlarmType_SCAN:
		return "SCAN"
	case AlarmType_ALARM_8:
		return "ALARM_8"
	case AlarmType_ALARM_S:
		return "ALARM_S"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e AlarmType) String() string {
	return e.PLC4XEnumName()
}
