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

// BACnetEventState is an enum
type BACnetEventState uint16

type IBACnetEventState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetEventState_NORMAL                   BACnetEventState = 0
	BACnetEventState_FAULT                    BACnetEventState = 1
	BACnetEventState_OFFNORMAL                BACnetEventState = 2
	BACnetEventState_HIGH_LIMIT               BACnetEventState = 3
	BACnetEventState_LOW_LIMIT                BACnetEventState = 4
	BACnetEventState_LIFE_SAVETY_ALARM        BACnetEventState = 5
	BACnetEventState_VENDOR_PROPRIETARY_VALUE BACnetEventState = 0xFFFF
)

var BACnetEventStateValues []BACnetEventState

func init() {
	_ = errors.New
	BACnetEventStateValues = []BACnetEventState{
		BACnetEventState_NORMAL,
		BACnetEventState_FAULT,
		BACnetEventState_OFFNORMAL,
		BACnetEventState_HIGH_LIMIT,
		BACnetEventState_LOW_LIMIT,
		BACnetEventState_LIFE_SAVETY_ALARM,
		BACnetEventState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEventStateByValue(value uint16) (enum BACnetEventState, ok bool) {
	switch value {
	case 0:
		return BACnetEventState_NORMAL, true
	case 0xFFFF:
		return BACnetEventState_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetEventState_FAULT, true
	case 2:
		return BACnetEventState_OFFNORMAL, true
	case 3:
		return BACnetEventState_HIGH_LIMIT, true
	case 4:
		return BACnetEventState_LOW_LIMIT, true
	case 5:
		return BACnetEventState_LIFE_SAVETY_ALARM, true
	}
	return 0, false
}

func BACnetEventStateByName(value string) (enum BACnetEventState, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetEventState_NORMAL, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetEventState_VENDOR_PROPRIETARY_VALUE, true
	case "FAULT":
		return BACnetEventState_FAULT, true
	case "OFFNORMAL":
		return BACnetEventState_OFFNORMAL, true
	case "HIGH_LIMIT":
		return BACnetEventState_HIGH_LIMIT, true
	case "LOW_LIMIT":
		return BACnetEventState_LOW_LIMIT, true
	case "LIFE_SAVETY_ALARM":
		return BACnetEventState_LIFE_SAVETY_ALARM, true
	}
	return 0, false
}

func BACnetEventStateKnows(value uint16) bool {
	for _, typeValue := range BACnetEventStateValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEventState(structType any) BACnetEventState {
	castFunc := func(typ any) BACnetEventState {
		if sBACnetEventState, ok := typ.(BACnetEventState); ok {
			return sBACnetEventState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEventState) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetEventState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventStateParse(ctx context.Context, theBytes []byte) (BACnetEventState, error) {
	return BACnetEventStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetEventState", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetEventState")
	}
	if enum, ok := BACnetEventStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetEventState")
		return BACnetEventState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetEventState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetEventState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetEventState", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEventState) PLC4XEnumName() string {
	switch e {
	case BACnetEventState_NORMAL:
		return "NORMAL"
	case BACnetEventState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEventState_FAULT:
		return "FAULT"
	case BACnetEventState_OFFNORMAL:
		return "OFFNORMAL"
	case BACnetEventState_HIGH_LIMIT:
		return "HIGH_LIMIT"
	case BACnetEventState_LOW_LIMIT:
		return "LOW_LIMIT"
	case BACnetEventState_LIFE_SAVETY_ALARM:
		return "LIFE_SAVETY_ALARM"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetEventState) String() string {
	return e.PLC4XEnumName()
}
