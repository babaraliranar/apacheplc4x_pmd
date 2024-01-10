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

// BACnetDoorSecuredStatus is an enum
type BACnetDoorSecuredStatus uint8

type IBACnetDoorSecuredStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetDoorSecuredStatus_SECURED   BACnetDoorSecuredStatus = 0
	BACnetDoorSecuredStatus_UNSECURED BACnetDoorSecuredStatus = 1
	BACnetDoorSecuredStatus_UNKNOWN   BACnetDoorSecuredStatus = 2
)

var BACnetDoorSecuredStatusValues []BACnetDoorSecuredStatus

func init() {
	_ = errors.New
	BACnetDoorSecuredStatusValues = []BACnetDoorSecuredStatus{
		BACnetDoorSecuredStatus_SECURED,
		BACnetDoorSecuredStatus_UNSECURED,
		BACnetDoorSecuredStatus_UNKNOWN,
	}
}

func BACnetDoorSecuredStatusByValue(value uint8) (enum BACnetDoorSecuredStatus, ok bool) {
	switch value {
	case 0:
		return BACnetDoorSecuredStatus_SECURED, true
	case 1:
		return BACnetDoorSecuredStatus_UNSECURED, true
	case 2:
		return BACnetDoorSecuredStatus_UNKNOWN, true
	}
	return 0, false
}

func BACnetDoorSecuredStatusByName(value string) (enum BACnetDoorSecuredStatus, ok bool) {
	switch value {
	case "SECURED":
		return BACnetDoorSecuredStatus_SECURED, true
	case "UNSECURED":
		return BACnetDoorSecuredStatus_UNSECURED, true
	case "UNKNOWN":
		return BACnetDoorSecuredStatus_UNKNOWN, true
	}
	return 0, false
}

func BACnetDoorSecuredStatusKnows(value uint8) bool {
	for _, typeValue := range BACnetDoorSecuredStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetDoorSecuredStatus(structType any) BACnetDoorSecuredStatus {
	castFunc := func(typ any) BACnetDoorSecuredStatus {
		if sBACnetDoorSecuredStatus, ok := typ.(BACnetDoorSecuredStatus); ok {
			return sBACnetDoorSecuredStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDoorSecuredStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetDoorSecuredStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDoorSecuredStatusParse(ctx context.Context, theBytes []byte) (BACnetDoorSecuredStatus, error) {
	return BACnetDoorSecuredStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDoorSecuredStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDoorSecuredStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetDoorSecuredStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDoorSecuredStatus")
	}
	if enum, ok := BACnetDoorSecuredStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetDoorSecuredStatus")
		return BACnetDoorSecuredStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDoorSecuredStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetDoorSecuredStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetDoorSecuredStatus", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDoorSecuredStatus) PLC4XEnumName() string {
	switch e {
	case BACnetDoorSecuredStatus_SECURED:
		return "SECURED"
	case BACnetDoorSecuredStatus_UNSECURED:
		return "UNSECURED"
	case BACnetDoorSecuredStatus_UNKNOWN:
		return "UNKNOWN"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetDoorSecuredStatus) String() string {
	return e.PLC4XEnumName()
}
