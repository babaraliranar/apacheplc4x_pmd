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

// OpcuaNodeIdServicesVariableDeadband is an enum
type OpcuaNodeIdServicesVariableDeadband int32

type IOpcuaNodeIdServicesVariableDeadband interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableDeadband_DeadbandType_EnumStrings OpcuaNodeIdServicesVariableDeadband = 7610
)

var OpcuaNodeIdServicesVariableDeadbandValues []OpcuaNodeIdServicesVariableDeadband

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableDeadbandValues = []OpcuaNodeIdServicesVariableDeadband{
		OpcuaNodeIdServicesVariableDeadband_DeadbandType_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableDeadbandByValue(value int32) (enum OpcuaNodeIdServicesVariableDeadband, ok bool) {
	switch value {
	case 7610:
		return OpcuaNodeIdServicesVariableDeadband_DeadbandType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDeadbandByName(value string) (enum OpcuaNodeIdServicesVariableDeadband, ok bool) {
	switch value {
	case "DeadbandType_EnumStrings":
		return OpcuaNodeIdServicesVariableDeadband_DeadbandType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDeadbandKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableDeadbandValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableDeadband(structType any) OpcuaNodeIdServicesVariableDeadband {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableDeadband {
		if sOpcuaNodeIdServicesVariableDeadband, ok := typ.(OpcuaNodeIdServicesVariableDeadband); ok {
			return sOpcuaNodeIdServicesVariableDeadband
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableDeadband) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableDeadband) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableDeadbandParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableDeadband, error) {
	return OpcuaNodeIdServicesVariableDeadbandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableDeadbandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableDeadband, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableDeadband", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableDeadband")
	}
	if enum, ok := OpcuaNodeIdServicesVariableDeadbandByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableDeadband")
		return OpcuaNodeIdServicesVariableDeadband(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableDeadband) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableDeadband) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableDeadband", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableDeadband) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableDeadband_DeadbandType_EnumStrings:
		return "DeadbandType_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableDeadband) String() string {
	return e.PLC4XEnumName()
}
