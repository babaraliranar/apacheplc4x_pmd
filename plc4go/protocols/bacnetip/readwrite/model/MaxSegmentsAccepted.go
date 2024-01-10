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

// MaxSegmentsAccepted is an enum
type MaxSegmentsAccepted uint8

type IMaxSegmentsAccepted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MaxSegments() uint8
}

const (
	MaxSegmentsAccepted_UNSPECIFIED           MaxSegmentsAccepted = 0x0
	MaxSegmentsAccepted_NUM_SEGMENTS_02       MaxSegmentsAccepted = 0x1
	MaxSegmentsAccepted_NUM_SEGMENTS_04       MaxSegmentsAccepted = 0x2
	MaxSegmentsAccepted_NUM_SEGMENTS_08       MaxSegmentsAccepted = 0x3
	MaxSegmentsAccepted_NUM_SEGMENTS_16       MaxSegmentsAccepted = 0x4
	MaxSegmentsAccepted_NUM_SEGMENTS_32       MaxSegmentsAccepted = 0x5
	MaxSegmentsAccepted_NUM_SEGMENTS_64       MaxSegmentsAccepted = 0x6
	MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS MaxSegmentsAccepted = 0x7
)

var MaxSegmentsAcceptedValues []MaxSegmentsAccepted

func init() {
	_ = errors.New
	MaxSegmentsAcceptedValues = []MaxSegmentsAccepted{
		MaxSegmentsAccepted_UNSPECIFIED,
		MaxSegmentsAccepted_NUM_SEGMENTS_02,
		MaxSegmentsAccepted_NUM_SEGMENTS_04,
		MaxSegmentsAccepted_NUM_SEGMENTS_08,
		MaxSegmentsAccepted_NUM_SEGMENTS_16,
		MaxSegmentsAccepted_NUM_SEGMENTS_32,
		MaxSegmentsAccepted_NUM_SEGMENTS_64,
		MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS,
	}
}

func (e MaxSegmentsAccepted) MaxSegments() uint8 {
	switch e {
	case 0x0:
		{ /* '0x0' */
			return 255
		}
	case 0x1:
		{ /* '0x1' */
			return 2
		}
	case 0x2:
		{ /* '0x2' */
			return 4
		}
	case 0x3:
		{ /* '0x3' */
			return 8
		}
	case 0x4:
		{ /* '0x4' */
			return 16
		}
	case 0x5:
		{ /* '0x5' */
			return 32
		}
	case 0x6:
		{ /* '0x6' */
			return 64
		}
	case 0x7:
		{ /* '0x7' */
			return 255
		}
	default:
		{
			return 0
		}
	}
}

func MaxSegmentsAcceptedFirstEnumForFieldMaxSegments(value uint8) (MaxSegmentsAccepted, error) {
	for _, sizeValue := range MaxSegmentsAcceptedValues {
		if sizeValue.MaxSegments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing MaxSegments not found", value)
}
func MaxSegmentsAcceptedByValue(value uint8) (enum MaxSegmentsAccepted, ok bool) {
	switch value {
	case 0x0:
		return MaxSegmentsAccepted_UNSPECIFIED, true
	case 0x1:
		return MaxSegmentsAccepted_NUM_SEGMENTS_02, true
	case 0x2:
		return MaxSegmentsAccepted_NUM_SEGMENTS_04, true
	case 0x3:
		return MaxSegmentsAccepted_NUM_SEGMENTS_08, true
	case 0x4:
		return MaxSegmentsAccepted_NUM_SEGMENTS_16, true
	case 0x5:
		return MaxSegmentsAccepted_NUM_SEGMENTS_32, true
	case 0x6:
		return MaxSegmentsAccepted_NUM_SEGMENTS_64, true
	case 0x7:
		return MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS, true
	}
	return 0, false
}

func MaxSegmentsAcceptedByName(value string) (enum MaxSegmentsAccepted, ok bool) {
	switch value {
	case "UNSPECIFIED":
		return MaxSegmentsAccepted_UNSPECIFIED, true
	case "NUM_SEGMENTS_02":
		return MaxSegmentsAccepted_NUM_SEGMENTS_02, true
	case "NUM_SEGMENTS_04":
		return MaxSegmentsAccepted_NUM_SEGMENTS_04, true
	case "NUM_SEGMENTS_08":
		return MaxSegmentsAccepted_NUM_SEGMENTS_08, true
	case "NUM_SEGMENTS_16":
		return MaxSegmentsAccepted_NUM_SEGMENTS_16, true
	case "NUM_SEGMENTS_32":
		return MaxSegmentsAccepted_NUM_SEGMENTS_32, true
	case "NUM_SEGMENTS_64":
		return MaxSegmentsAccepted_NUM_SEGMENTS_64, true
	case "MORE_THAN_64_SEGMENTS":
		return MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS, true
	}
	return 0, false
}

func MaxSegmentsAcceptedKnows(value uint8) bool {
	for _, typeValue := range MaxSegmentsAcceptedValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMaxSegmentsAccepted(structType any) MaxSegmentsAccepted {
	castFunc := func(typ any) MaxSegmentsAccepted {
		if sMaxSegmentsAccepted, ok := typ.(MaxSegmentsAccepted); ok {
			return sMaxSegmentsAccepted
		}
		return 0
	}
	return castFunc(structType)
}

func (m MaxSegmentsAccepted) GetLengthInBits(ctx context.Context) uint16 {
	return 3
}

func (m MaxSegmentsAccepted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MaxSegmentsAcceptedParse(ctx context.Context, theBytes []byte) (MaxSegmentsAccepted, error) {
	return MaxSegmentsAcceptedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MaxSegmentsAcceptedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MaxSegmentsAccepted, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("MaxSegmentsAccepted", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MaxSegmentsAccepted")
	}
	if enum, ok := MaxSegmentsAcceptedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for MaxSegmentsAccepted")
		return MaxSegmentsAccepted(val), nil
	} else {
		return enum, nil
	}
}

func (e MaxSegmentsAccepted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MaxSegmentsAccepted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("MaxSegmentsAccepted", 3, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MaxSegmentsAccepted) PLC4XEnumName() string {
	switch e {
	case MaxSegmentsAccepted_UNSPECIFIED:
		return "UNSPECIFIED"
	case MaxSegmentsAccepted_NUM_SEGMENTS_02:
		return "NUM_SEGMENTS_02"
	case MaxSegmentsAccepted_NUM_SEGMENTS_04:
		return "NUM_SEGMENTS_04"
	case MaxSegmentsAccepted_NUM_SEGMENTS_08:
		return "NUM_SEGMENTS_08"
	case MaxSegmentsAccepted_NUM_SEGMENTS_16:
		return "NUM_SEGMENTS_16"
	case MaxSegmentsAccepted_NUM_SEGMENTS_32:
		return "NUM_SEGMENTS_32"
	case MaxSegmentsAccepted_NUM_SEGMENTS_64:
		return "NUM_SEGMENTS_64"
	case MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS:
		return "MORE_THAN_64_SEGMENTS"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e MaxSegmentsAccepted) String() string {
	return e.PLC4XEnumName()
}
