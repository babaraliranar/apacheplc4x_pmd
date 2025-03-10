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

// OpcuaNodeIdServicesVariableView is an enum
type OpcuaNodeIdServicesVariableView int32

type IOpcuaNodeIdServicesVariableView interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableView_ViewVersion OpcuaNodeIdServicesVariableView = 12170
)

var OpcuaNodeIdServicesVariableViewValues []OpcuaNodeIdServicesVariableView

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableViewValues = []OpcuaNodeIdServicesVariableView{
		OpcuaNodeIdServicesVariableView_ViewVersion,
	}
}

func OpcuaNodeIdServicesVariableViewByValue(value int32) (enum OpcuaNodeIdServicesVariableView, ok bool) {
	switch value {
	case 12170:
		return OpcuaNodeIdServicesVariableView_ViewVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableViewByName(value string) (enum OpcuaNodeIdServicesVariableView, ok bool) {
	switch value {
	case "ViewVersion":
		return OpcuaNodeIdServicesVariableView_ViewVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableViewKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableViewValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableView(structType any) OpcuaNodeIdServicesVariableView {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableView {
		if sOpcuaNodeIdServicesVariableView, ok := typ.(OpcuaNodeIdServicesVariableView); ok {
			return sOpcuaNodeIdServicesVariableView
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableView) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableView) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableViewParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableView, error) {
	return OpcuaNodeIdServicesVariableViewParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableViewParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableView, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableView", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableView")
	}
	if enum, ok := OpcuaNodeIdServicesVariableViewByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableView")
		return OpcuaNodeIdServicesVariableView(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableView) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableView) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableView", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableView) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableView) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableView_ViewVersion:
		return "ViewVersion"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableView) String() string {
	return e.PLC4XEnumName()
}
