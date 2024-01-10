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

// BrowseDirection is an enum
type BrowseDirection uint32

type IBrowseDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BrowseDirection_browseDirectionForward BrowseDirection = 0
	BrowseDirection_browseDirectionInverse BrowseDirection = 1
	BrowseDirection_browseDirectionBoth    BrowseDirection = 2
	BrowseDirection_browseDirectionInvalid BrowseDirection = 3
)

var BrowseDirectionValues []BrowseDirection

func init() {
	_ = errors.New
	BrowseDirectionValues = []BrowseDirection{
		BrowseDirection_browseDirectionForward,
		BrowseDirection_browseDirectionInverse,
		BrowseDirection_browseDirectionBoth,
		BrowseDirection_browseDirectionInvalid,
	}
}

func BrowseDirectionByValue(value uint32) (enum BrowseDirection, ok bool) {
	switch value {
	case 0:
		return BrowseDirection_browseDirectionForward, true
	case 1:
		return BrowseDirection_browseDirectionInverse, true
	case 2:
		return BrowseDirection_browseDirectionBoth, true
	case 3:
		return BrowseDirection_browseDirectionInvalid, true
	}
	return 0, false
}

func BrowseDirectionByName(value string) (enum BrowseDirection, ok bool) {
	switch value {
	case "browseDirectionForward":
		return BrowseDirection_browseDirectionForward, true
	case "browseDirectionInverse":
		return BrowseDirection_browseDirectionInverse, true
	case "browseDirectionBoth":
		return BrowseDirection_browseDirectionBoth, true
	case "browseDirectionInvalid":
		return BrowseDirection_browseDirectionInvalid, true
	}
	return 0, false
}

func BrowseDirectionKnows(value uint32) bool {
	for _, typeValue := range BrowseDirectionValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBrowseDirection(structType any) BrowseDirection {
	castFunc := func(typ any) BrowseDirection {
		if sBrowseDirection, ok := typ.(BrowseDirection); ok {
			return sBrowseDirection
		}
		return 0
	}
	return castFunc(structType)
}

func (m BrowseDirection) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m BrowseDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowseDirectionParse(ctx context.Context, theBytes []byte) (BrowseDirection, error) {
	return BrowseDirectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BrowseDirectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BrowseDirection, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("BrowseDirection", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BrowseDirection")
	}
	if enum, ok := BrowseDirectionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BrowseDirection")
		return BrowseDirection(val), nil
	} else {
		return enum, nil
	}
}

func (e BrowseDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BrowseDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("BrowseDirection", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BrowseDirection) PLC4XEnumName() string {
	switch e {
	case BrowseDirection_browseDirectionForward:
		return "browseDirectionForward"
	case BrowseDirection_browseDirectionInverse:
		return "browseDirectionInverse"
	case BrowseDirection_browseDirectionBoth:
		return "browseDirectionBoth"
	case BrowseDirection_browseDirectionInvalid:
		return "browseDirectionInvalid"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e BrowseDirection) String() string {
	return e.PLC4XEnumName()
}
