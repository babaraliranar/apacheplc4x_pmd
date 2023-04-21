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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable is an enum
type BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable uint8

type IBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE             BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable = 0
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE            BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable = 1
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable = 2
)

var BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableValues []BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable

func init() {
	_ = errors.New
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableValues = []BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable{
		BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE,
		BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE,
		BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION,
	}
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableByValue(value uint8) (enum BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, ok bool) {
	switch value {
	case 0:
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE, true
	case 1:
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE, true
	case 2:
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableByName(value string) (enum BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, ok bool) {
	switch value {
	case "ENABLE":
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE, true
	case "DISABLE":
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE, true
	case "DISABLE_INITIATION":
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable(structType any) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
	castFunc := func(typ any) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
		if sBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, ok := typ.(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable); ok {
			return sBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableParse(ctx context.Context, theBytes []byte) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, error) {
	return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable")
	}
	if enum, ok := BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) PLC4XEnumName() string {
	switch e {
	case BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE:
		return "ENABLE"
	case BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE:
		return "DISABLE"
	case BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION:
		return "DISABLE_INITIATION"
	}
	return ""
}

func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) String() string {
	return e.PLC4XEnumName()
}
