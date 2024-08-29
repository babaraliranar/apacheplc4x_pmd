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

// OpcuaNodeIdServicesVariableNetwork is an enum
type OpcuaNodeIdServicesVariableNetwork int32

type IOpcuaNodeIdServicesVariableNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_Selections            OpcuaNodeIdServicesVariableNetwork = 17582
	OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_SelectionDescriptions OpcuaNodeIdServicesVariableNetwork = 17583
	OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_RestrictToList        OpcuaNodeIdServicesVariableNetwork = 17584
	OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface                       OpcuaNodeIdServicesVariableNetwork = 21146
	OpcuaNodeIdServicesVariableNetwork_NetworkAddressUrlType_Url                                 OpcuaNodeIdServicesVariableNetwork = 21149
)

var OpcuaNodeIdServicesVariableNetworkValues []OpcuaNodeIdServicesVariableNetwork

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableNetworkValues = []OpcuaNodeIdServicesVariableNetwork{
		OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_Selections,
		OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_SelectionDescriptions,
		OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_RestrictToList,
		OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface,
		OpcuaNodeIdServicesVariableNetwork_NetworkAddressUrlType_Url,
	}
}

func OpcuaNodeIdServicesVariableNetworkByValue(value int32) (enum OpcuaNodeIdServicesVariableNetwork, ok bool) {
	switch value {
	case 17582:
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_Selections, true
	case 17583:
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_SelectionDescriptions, true
	case 17584:
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_RestrictToList, true
	case 21146:
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface, true
	case 21149:
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressUrlType_Url, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNetworkByName(value string) (enum OpcuaNodeIdServicesVariableNetwork, ok bool) {
	switch value {
	case "NetworkAddressType_NetworkInterface_Selections":
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_Selections, true
	case "NetworkAddressType_NetworkInterface_SelectionDescriptions":
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_SelectionDescriptions, true
	case "NetworkAddressType_NetworkInterface_RestrictToList":
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_RestrictToList, true
	case "NetworkAddressType_NetworkInterface":
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface, true
	case "NetworkAddressUrlType_Url":
		return OpcuaNodeIdServicesVariableNetwork_NetworkAddressUrlType_Url, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNetworkKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableNetworkValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableNetwork(structType any) OpcuaNodeIdServicesVariableNetwork {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableNetwork {
		if sOpcuaNodeIdServicesVariableNetwork, ok := typ.(OpcuaNodeIdServicesVariableNetwork); ok {
			return sOpcuaNodeIdServicesVariableNetwork
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableNetwork) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableNetworkParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableNetwork, error) {
	return OpcuaNodeIdServicesVariableNetworkParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableNetwork, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableNetwork", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableNetwork")
	}
	if enum, ok := OpcuaNodeIdServicesVariableNetworkByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableNetwork")
		return OpcuaNodeIdServicesVariableNetwork(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableNetwork", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableNetwork) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_Selections:
		return "NetworkAddressType_NetworkInterface_Selections"
	case OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_SelectionDescriptions:
		return "NetworkAddressType_NetworkInterface_SelectionDescriptions"
	case OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface_RestrictToList:
		return "NetworkAddressType_NetworkInterface_RestrictToList"
	case OpcuaNodeIdServicesVariableNetwork_NetworkAddressType_NetworkInterface:
		return "NetworkAddressType_NetworkInterface"
	case OpcuaNodeIdServicesVariableNetwork_NetworkAddressUrlType_Url:
		return "NetworkAddressUrlType_Url"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableNetwork) String() string {
	return e.PLC4XEnumName()
}
