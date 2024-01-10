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

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableConversion is an enum
type OpcuaNodeIdServicesVariableConversion int32

type IOpcuaNodeIdServicesVariableConversion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableConversion_ConversionLimitEnum_EnumStrings OpcuaNodeIdServicesVariableConversion = 32437
)

var OpcuaNodeIdServicesVariableConversionValues []OpcuaNodeIdServicesVariableConversion

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableConversionValues = []OpcuaNodeIdServicesVariableConversion{
		OpcuaNodeIdServicesVariableConversion_ConversionLimitEnum_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableConversionByValue(value int32) (enum OpcuaNodeIdServicesVariableConversion, ok bool) {
	switch value {
	case 32437:
		return OpcuaNodeIdServicesVariableConversion_ConversionLimitEnum_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableConversionByName(value string) (enum OpcuaNodeIdServicesVariableConversion, ok bool) {
	switch value {
	case "ConversionLimitEnum_EnumStrings":
		return OpcuaNodeIdServicesVariableConversion_ConversionLimitEnum_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableConversionKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableConversionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableConversion(structType any) OpcuaNodeIdServicesVariableConversion {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableConversion {
		if sOpcuaNodeIdServicesVariableConversion, ok := typ.(OpcuaNodeIdServicesVariableConversion); ok {
			return sOpcuaNodeIdServicesVariableConversion
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableConversion) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableConversion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableConversionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableConversion, error) {
	return OpcuaNodeIdServicesVariableConversionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableConversionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableConversion, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableConversion", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableConversion")
	}
	if enum, ok := OpcuaNodeIdServicesVariableConversionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableConversion")
		return OpcuaNodeIdServicesVariableConversion(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableConversion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableConversion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableConversion", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableConversion) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableConversion_ConversionLimitEnum_EnumStrings:
		return "ConversionLimitEnum_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableConversion) String() string {
	return e.PLC4XEnumName()
}
