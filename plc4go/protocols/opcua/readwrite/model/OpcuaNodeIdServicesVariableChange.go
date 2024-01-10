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

// OpcuaNodeIdServicesVariableChange is an enum
type OpcuaNodeIdServicesVariableChange int32

type IOpcuaNodeIdServicesVariableChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableChange_ChangePasswordMethodType_InputArguments OpcuaNodeIdServicesVariableChange = 24289
)

var OpcuaNodeIdServicesVariableChangeValues []OpcuaNodeIdServicesVariableChange

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableChangeValues = []OpcuaNodeIdServicesVariableChange{
		OpcuaNodeIdServicesVariableChange_ChangePasswordMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableChangeByValue(value int32) (enum OpcuaNodeIdServicesVariableChange, ok bool) {
	switch value {
	case 24289:
		return OpcuaNodeIdServicesVariableChange_ChangePasswordMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableChangeByName(value string) (enum OpcuaNodeIdServicesVariableChange, ok bool) {
	switch value {
	case "ChangePasswordMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableChange_ChangePasswordMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableChangeKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableChangeValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableChange(structType any) OpcuaNodeIdServicesVariableChange {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableChange {
		if sOpcuaNodeIdServicesVariableChange, ok := typ.(OpcuaNodeIdServicesVariableChange); ok {
			return sOpcuaNodeIdServicesVariableChange
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableChange) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableChangeParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableChange, error) {
	return OpcuaNodeIdServicesVariableChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableChange, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableChange", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableChange")
	}
	if enum, ok := OpcuaNodeIdServicesVariableChangeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableChange")
		return OpcuaNodeIdServicesVariableChange(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableChange", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableChange) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableChange_ChangePasswordMethodType_InputArguments:
		return "ChangePasswordMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableChange) String() string {
	return e.PLC4XEnumName()
}
