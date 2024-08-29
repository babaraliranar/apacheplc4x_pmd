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

// OpcuaNodeIdServicesVariableInterface is an enum
type OpcuaNodeIdServicesVariableInterface int32

type IOpcuaNodeIdServicesVariableInterface interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableInterface_InterfaceAdminStatus_EnumValues OpcuaNodeIdServicesVariableInterface = 24236
	OpcuaNodeIdServicesVariableInterface_InterfaceOperStatus_EnumValues  OpcuaNodeIdServicesVariableInterface = 24237
)

var OpcuaNodeIdServicesVariableInterfaceValues []OpcuaNodeIdServicesVariableInterface

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableInterfaceValues = []OpcuaNodeIdServicesVariableInterface{
		OpcuaNodeIdServicesVariableInterface_InterfaceAdminStatus_EnumValues,
		OpcuaNodeIdServicesVariableInterface_InterfaceOperStatus_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableInterfaceByValue(value int32) (enum OpcuaNodeIdServicesVariableInterface, ok bool) {
	switch value {
	case 24236:
		return OpcuaNodeIdServicesVariableInterface_InterfaceAdminStatus_EnumValues, true
	case 24237:
		return OpcuaNodeIdServicesVariableInterface_InterfaceOperStatus_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableInterfaceByName(value string) (enum OpcuaNodeIdServicesVariableInterface, ok bool) {
	switch value {
	case "InterfaceAdminStatus_EnumValues":
		return OpcuaNodeIdServicesVariableInterface_InterfaceAdminStatus_EnumValues, true
	case "InterfaceOperStatus_EnumValues":
		return OpcuaNodeIdServicesVariableInterface_InterfaceOperStatus_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableInterfaceKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableInterfaceValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableInterface(structType any) OpcuaNodeIdServicesVariableInterface {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableInterface {
		if sOpcuaNodeIdServicesVariableInterface, ok := typ.(OpcuaNodeIdServicesVariableInterface); ok {
			return sOpcuaNodeIdServicesVariableInterface
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableInterface) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableInterface) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableInterfaceParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableInterface, error) {
	return OpcuaNodeIdServicesVariableInterfaceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableInterfaceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableInterface, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableInterface", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableInterface")
	}
	if enum, ok := OpcuaNodeIdServicesVariableInterfaceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableInterface")
		return OpcuaNodeIdServicesVariableInterface(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableInterface) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableInterface) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableInterface", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableInterface) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableInterface_InterfaceAdminStatus_EnumValues:
		return "InterfaceAdminStatus_EnumValues"
	case OpcuaNodeIdServicesVariableInterface_InterfaceOperStatus_EnumValues:
		return "InterfaceOperStatus_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableInterface) String() string {
	return e.PLC4XEnumName()
}
