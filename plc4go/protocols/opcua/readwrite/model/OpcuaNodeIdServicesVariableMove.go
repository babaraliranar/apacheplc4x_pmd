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

// OpcuaNodeIdServicesVariableMove is an enum
type OpcuaNodeIdServicesVariableMove int32

type IOpcuaNodeIdServicesVariableMove interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_InputArguments  OpcuaNodeIdServicesVariableMove = 13351
	OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_OutputArguments OpcuaNodeIdServicesVariableMove = 13352
)

var OpcuaNodeIdServicesVariableMoveValues []OpcuaNodeIdServicesVariableMove

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableMoveValues = []OpcuaNodeIdServicesVariableMove{
		OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_InputArguments,
		OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableMoveByValue(value int32) (enum OpcuaNodeIdServicesVariableMove, ok bool) {
	switch value {
	case 13351:
		return OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_InputArguments, true
	case 13352:
		return OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableMoveByName(value string) (enum OpcuaNodeIdServicesVariableMove, ok bool) {
	switch value {
	case "MoveOrCopyMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_InputArguments, true
	case "MoveOrCopyMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableMoveKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableMoveValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableMove(structType any) OpcuaNodeIdServicesVariableMove {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableMove {
		if sOpcuaNodeIdServicesVariableMove, ok := typ.(OpcuaNodeIdServicesVariableMove); ok {
			return sOpcuaNodeIdServicesVariableMove
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableMove) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableMove) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableMoveParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableMove, error) {
	return OpcuaNodeIdServicesVariableMoveParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableMoveParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableMove, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableMove", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableMove")
	}
	if enum, ok := OpcuaNodeIdServicesVariableMoveByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableMove")
		return OpcuaNodeIdServicesVariableMove(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableMove) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableMove) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableMove", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableMove) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_InputArguments:
		return "MoveOrCopyMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableMove_MoveOrCopyMethodType_OutputArguments:
		return "MoveOrCopyMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableMove) String() string {
	return e.PLC4XEnumName()
}
