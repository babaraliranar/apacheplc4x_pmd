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

// BACnetRelationship is an enum
type BACnetRelationship uint16

type IBACnetRelationship interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetRelationship_UNKNOWN                  BACnetRelationship = 0
	BACnetRelationship_DEFAULT                  BACnetRelationship = 1
	BACnetRelationship_CONTAINS                 BACnetRelationship = 2
	BACnetRelationship_CONTAINED_BY             BACnetRelationship = 3
	BACnetRelationship_USES                     BACnetRelationship = 4
	BACnetRelationship_USED_BY                  BACnetRelationship = 5
	BACnetRelationship_COMMANDS                 BACnetRelationship = 6
	BACnetRelationship_COMMANDED_BY             BACnetRelationship = 7
	BACnetRelationship_ADJUSTS                  BACnetRelationship = 8
	BACnetRelationship_ADJUSTED_BY              BACnetRelationship = 9
	BACnetRelationship_INGRESS                  BACnetRelationship = 10
	BACnetRelationship_EGRESS                   BACnetRelationship = 11
	BACnetRelationship_SUPPLIES_AIR             BACnetRelationship = 12
	BACnetRelationship_RECEIVES_AIR             BACnetRelationship = 13
	BACnetRelationship_SUPPLIES_HOT_AIR         BACnetRelationship = 14
	BACnetRelationship_RECEIVES_HOT_AIR         BACnetRelationship = 15
	BACnetRelationship_SUPPLIES_COOL_AIR        BACnetRelationship = 16
	BACnetRelationship_RECEIVES_COOL_AIR        BACnetRelationship = 17
	BACnetRelationship_SUPPLIES_POWER           BACnetRelationship = 18
	BACnetRelationship_RECEIVES_POWER           BACnetRelationship = 19
	BACnetRelationship_SUPPLIES_GAS             BACnetRelationship = 20
	BACnetRelationship_RECEIVES_GAS             BACnetRelationship = 21
	BACnetRelationship_SUPPLIES_WATER           BACnetRelationship = 22
	BACnetRelationship_RECEIVES_WATER           BACnetRelationship = 23
	BACnetRelationship_SUPPLIES_HOT_WATER       BACnetRelationship = 24
	BACnetRelationship_RECEIVES_HOT_WATER       BACnetRelationship = 25
	BACnetRelationship_SUPPLIES_COOL_WATER      BACnetRelationship = 26
	BACnetRelationship_RECEIVES_COOL_WATER      BACnetRelationship = 27
	BACnetRelationship_SUPPLIES_STEAM           BACnetRelationship = 28
	BACnetRelationship_RECEIVES_STEAM           BACnetRelationship = 29
	BACnetRelationship_VENDOR_PROPRIETARY_VALUE BACnetRelationship = 0xFFFF
)

var BACnetRelationshipValues []BACnetRelationship

func init() {
	_ = errors.New
	BACnetRelationshipValues = []BACnetRelationship{
		BACnetRelationship_UNKNOWN,
		BACnetRelationship_DEFAULT,
		BACnetRelationship_CONTAINS,
		BACnetRelationship_CONTAINED_BY,
		BACnetRelationship_USES,
		BACnetRelationship_USED_BY,
		BACnetRelationship_COMMANDS,
		BACnetRelationship_COMMANDED_BY,
		BACnetRelationship_ADJUSTS,
		BACnetRelationship_ADJUSTED_BY,
		BACnetRelationship_INGRESS,
		BACnetRelationship_EGRESS,
		BACnetRelationship_SUPPLIES_AIR,
		BACnetRelationship_RECEIVES_AIR,
		BACnetRelationship_SUPPLIES_HOT_AIR,
		BACnetRelationship_RECEIVES_HOT_AIR,
		BACnetRelationship_SUPPLIES_COOL_AIR,
		BACnetRelationship_RECEIVES_COOL_AIR,
		BACnetRelationship_SUPPLIES_POWER,
		BACnetRelationship_RECEIVES_POWER,
		BACnetRelationship_SUPPLIES_GAS,
		BACnetRelationship_RECEIVES_GAS,
		BACnetRelationship_SUPPLIES_WATER,
		BACnetRelationship_RECEIVES_WATER,
		BACnetRelationship_SUPPLIES_HOT_WATER,
		BACnetRelationship_RECEIVES_HOT_WATER,
		BACnetRelationship_SUPPLIES_COOL_WATER,
		BACnetRelationship_RECEIVES_COOL_WATER,
		BACnetRelationship_SUPPLIES_STEAM,
		BACnetRelationship_RECEIVES_STEAM,
		BACnetRelationship_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetRelationshipByValue(value uint16) (enum BACnetRelationship, ok bool) {
	switch value {
	case 0:
		return BACnetRelationship_UNKNOWN, true
	case 0xFFFF:
		return BACnetRelationship_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetRelationship_DEFAULT, true
	case 10:
		return BACnetRelationship_INGRESS, true
	case 11:
		return BACnetRelationship_EGRESS, true
	case 12:
		return BACnetRelationship_SUPPLIES_AIR, true
	case 13:
		return BACnetRelationship_RECEIVES_AIR, true
	case 14:
		return BACnetRelationship_SUPPLIES_HOT_AIR, true
	case 15:
		return BACnetRelationship_RECEIVES_HOT_AIR, true
	case 16:
		return BACnetRelationship_SUPPLIES_COOL_AIR, true
	case 17:
		return BACnetRelationship_RECEIVES_COOL_AIR, true
	case 18:
		return BACnetRelationship_SUPPLIES_POWER, true
	case 19:
		return BACnetRelationship_RECEIVES_POWER, true
	case 2:
		return BACnetRelationship_CONTAINS, true
	case 20:
		return BACnetRelationship_SUPPLIES_GAS, true
	case 21:
		return BACnetRelationship_RECEIVES_GAS, true
	case 22:
		return BACnetRelationship_SUPPLIES_WATER, true
	case 23:
		return BACnetRelationship_RECEIVES_WATER, true
	case 24:
		return BACnetRelationship_SUPPLIES_HOT_WATER, true
	case 25:
		return BACnetRelationship_RECEIVES_HOT_WATER, true
	case 26:
		return BACnetRelationship_SUPPLIES_COOL_WATER, true
	case 27:
		return BACnetRelationship_RECEIVES_COOL_WATER, true
	case 28:
		return BACnetRelationship_SUPPLIES_STEAM, true
	case 29:
		return BACnetRelationship_RECEIVES_STEAM, true
	case 3:
		return BACnetRelationship_CONTAINED_BY, true
	case 4:
		return BACnetRelationship_USES, true
	case 5:
		return BACnetRelationship_USED_BY, true
	case 6:
		return BACnetRelationship_COMMANDS, true
	case 7:
		return BACnetRelationship_COMMANDED_BY, true
	case 8:
		return BACnetRelationship_ADJUSTS, true
	case 9:
		return BACnetRelationship_ADJUSTED_BY, true
	}
	return 0, false
}

func BACnetRelationshipByName(value string) (enum BACnetRelationship, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetRelationship_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetRelationship_VENDOR_PROPRIETARY_VALUE, true
	case "DEFAULT":
		return BACnetRelationship_DEFAULT, true
	case "INGRESS":
		return BACnetRelationship_INGRESS, true
	case "EGRESS":
		return BACnetRelationship_EGRESS, true
	case "SUPPLIES_AIR":
		return BACnetRelationship_SUPPLIES_AIR, true
	case "RECEIVES_AIR":
		return BACnetRelationship_RECEIVES_AIR, true
	case "SUPPLIES_HOT_AIR":
		return BACnetRelationship_SUPPLIES_HOT_AIR, true
	case "RECEIVES_HOT_AIR":
		return BACnetRelationship_RECEIVES_HOT_AIR, true
	case "SUPPLIES_COOL_AIR":
		return BACnetRelationship_SUPPLIES_COOL_AIR, true
	case "RECEIVES_COOL_AIR":
		return BACnetRelationship_RECEIVES_COOL_AIR, true
	case "SUPPLIES_POWER":
		return BACnetRelationship_SUPPLIES_POWER, true
	case "RECEIVES_POWER":
		return BACnetRelationship_RECEIVES_POWER, true
	case "CONTAINS":
		return BACnetRelationship_CONTAINS, true
	case "SUPPLIES_GAS":
		return BACnetRelationship_SUPPLIES_GAS, true
	case "RECEIVES_GAS":
		return BACnetRelationship_RECEIVES_GAS, true
	case "SUPPLIES_WATER":
		return BACnetRelationship_SUPPLIES_WATER, true
	case "RECEIVES_WATER":
		return BACnetRelationship_RECEIVES_WATER, true
	case "SUPPLIES_HOT_WATER":
		return BACnetRelationship_SUPPLIES_HOT_WATER, true
	case "RECEIVES_HOT_WATER":
		return BACnetRelationship_RECEIVES_HOT_WATER, true
	case "SUPPLIES_COOL_WATER":
		return BACnetRelationship_SUPPLIES_COOL_WATER, true
	case "RECEIVES_COOL_WATER":
		return BACnetRelationship_RECEIVES_COOL_WATER, true
	case "SUPPLIES_STEAM":
		return BACnetRelationship_SUPPLIES_STEAM, true
	case "RECEIVES_STEAM":
		return BACnetRelationship_RECEIVES_STEAM, true
	case "CONTAINED_BY":
		return BACnetRelationship_CONTAINED_BY, true
	case "USES":
		return BACnetRelationship_USES, true
	case "USED_BY":
		return BACnetRelationship_USED_BY, true
	case "COMMANDS":
		return BACnetRelationship_COMMANDS, true
	case "COMMANDED_BY":
		return BACnetRelationship_COMMANDED_BY, true
	case "ADJUSTS":
		return BACnetRelationship_ADJUSTS, true
	case "ADJUSTED_BY":
		return BACnetRelationship_ADJUSTED_BY, true
	}
	return 0, false
}

func BACnetRelationshipKnows(value uint16) bool {
	for _, typeValue := range BACnetRelationshipValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetRelationship(structType any) BACnetRelationship {
	castFunc := func(typ any) BACnetRelationship {
		if sBACnetRelationship, ok := typ.(BACnetRelationship); ok {
			return sBACnetRelationship
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetRelationship) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetRelationship) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRelationshipParse(ctx context.Context, theBytes []byte) (BACnetRelationship, error) {
	return BACnetRelationshipParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRelationshipParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRelationship, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetRelationship", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetRelationship")
	}
	if enum, ok := BACnetRelationshipByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetRelationship")
		return BACnetRelationship(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetRelationship) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetRelationship) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetRelationship", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetRelationship) PLC4XEnumName() string {
	switch e {
	case BACnetRelationship_UNKNOWN:
		return "UNKNOWN"
	case BACnetRelationship_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetRelationship_DEFAULT:
		return "DEFAULT"
	case BACnetRelationship_INGRESS:
		return "INGRESS"
	case BACnetRelationship_EGRESS:
		return "EGRESS"
	case BACnetRelationship_SUPPLIES_AIR:
		return "SUPPLIES_AIR"
	case BACnetRelationship_RECEIVES_AIR:
		return "RECEIVES_AIR"
	case BACnetRelationship_SUPPLIES_HOT_AIR:
		return "SUPPLIES_HOT_AIR"
	case BACnetRelationship_RECEIVES_HOT_AIR:
		return "RECEIVES_HOT_AIR"
	case BACnetRelationship_SUPPLIES_COOL_AIR:
		return "SUPPLIES_COOL_AIR"
	case BACnetRelationship_RECEIVES_COOL_AIR:
		return "RECEIVES_COOL_AIR"
	case BACnetRelationship_SUPPLIES_POWER:
		return "SUPPLIES_POWER"
	case BACnetRelationship_RECEIVES_POWER:
		return "RECEIVES_POWER"
	case BACnetRelationship_CONTAINS:
		return "CONTAINS"
	case BACnetRelationship_SUPPLIES_GAS:
		return "SUPPLIES_GAS"
	case BACnetRelationship_RECEIVES_GAS:
		return "RECEIVES_GAS"
	case BACnetRelationship_SUPPLIES_WATER:
		return "SUPPLIES_WATER"
	case BACnetRelationship_RECEIVES_WATER:
		return "RECEIVES_WATER"
	case BACnetRelationship_SUPPLIES_HOT_WATER:
		return "SUPPLIES_HOT_WATER"
	case BACnetRelationship_RECEIVES_HOT_WATER:
		return "RECEIVES_HOT_WATER"
	case BACnetRelationship_SUPPLIES_COOL_WATER:
		return "SUPPLIES_COOL_WATER"
	case BACnetRelationship_RECEIVES_COOL_WATER:
		return "RECEIVES_COOL_WATER"
	case BACnetRelationship_SUPPLIES_STEAM:
		return "SUPPLIES_STEAM"
	case BACnetRelationship_RECEIVES_STEAM:
		return "RECEIVES_STEAM"
	case BACnetRelationship_CONTAINED_BY:
		return "CONTAINED_BY"
	case BACnetRelationship_USES:
		return "USES"
	case BACnetRelationship_USED_BY:
		return "USED_BY"
	case BACnetRelationship_COMMANDS:
		return "COMMANDS"
	case BACnetRelationship_COMMANDED_BY:
		return "COMMANDED_BY"
	case BACnetRelationship_ADJUSTS:
		return "ADJUSTS"
	case BACnetRelationship_ADJUSTED_BY:
		return "ADJUSTED_BY"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetRelationship) String() string {
	return e.PLC4XEnumName()
}
