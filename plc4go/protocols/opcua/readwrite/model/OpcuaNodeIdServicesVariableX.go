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

// OpcuaNodeIdServicesVariableX is an enum
type OpcuaNodeIdServicesVariableX int32

type IOpcuaNodeIdServicesVariableX interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableX_XYArrayItemType_XAxisDefinition OpcuaNodeIdServicesVariableX = 12046
)

var OpcuaNodeIdServicesVariableXValues []OpcuaNodeIdServicesVariableX

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableXValues = []OpcuaNodeIdServicesVariableX{
		OpcuaNodeIdServicesVariableX_XYArrayItemType_XAxisDefinition,
	}
}

func OpcuaNodeIdServicesVariableXByValue(value int32) (enum OpcuaNodeIdServicesVariableX, ok bool) {
	switch value {
	case 12046:
		return OpcuaNodeIdServicesVariableX_XYArrayItemType_XAxisDefinition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableXByName(value string) (enum OpcuaNodeIdServicesVariableX, ok bool) {
	switch value {
	case "XYArrayItemType_XAxisDefinition":
		return OpcuaNodeIdServicesVariableX_XYArrayItemType_XAxisDefinition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableXKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableXValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableX(structType any) OpcuaNodeIdServicesVariableX {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableX {
		if sOpcuaNodeIdServicesVariableX, ok := typ.(OpcuaNodeIdServicesVariableX); ok {
			return sOpcuaNodeIdServicesVariableX
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableX) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableX) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableXParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableX, error) {
	return OpcuaNodeIdServicesVariableXParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableXParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableX, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableX", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableX")
	}
	if enum, ok := OpcuaNodeIdServicesVariableXByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableX")
		return OpcuaNodeIdServicesVariableX(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableX) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableX) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableX", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableX) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableX_XYArrayItemType_XAxisDefinition:
		return "XYArrayItemType_XAxisDefinition"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableX) String() string {
	return e.PLC4XEnumName()
}
