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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// COTPTpduSize is an enum
type COTPTpduSize uint8

type ICOTPTpduSize interface {
	utils.Serializable
	SizeInBytes() uint16
}

const(
	COTPTpduSize_SIZE_128 COTPTpduSize = 0x07
	COTPTpduSize_SIZE_256 COTPTpduSize = 0x08
	COTPTpduSize_SIZE_512 COTPTpduSize = 0x09
	COTPTpduSize_SIZE_1024 COTPTpduSize = 0x0a
	COTPTpduSize_SIZE_2048 COTPTpduSize = 0x0b
	COTPTpduSize_SIZE_4096 COTPTpduSize = 0x0c
	COTPTpduSize_SIZE_8192 COTPTpduSize = 0x0d
)

var COTPTpduSizeValues []COTPTpduSize

func init() {
	_ = errors.New
	COTPTpduSizeValues = []COTPTpduSize {
		COTPTpduSize_SIZE_128,
		COTPTpduSize_SIZE_256,
		COTPTpduSize_SIZE_512,
		COTPTpduSize_SIZE_1024,
		COTPTpduSize_SIZE_2048,
		COTPTpduSize_SIZE_4096,
		COTPTpduSize_SIZE_8192,
	}
}


func (e COTPTpduSize) SizeInBytes() uint16 {
	switch e  {
		case 0x07: { /* '0x07' */
            return 128
		}
		case 0x08: { /* '0x08' */
            return 256
		}
		case 0x09: { /* '0x09' */
            return 512
		}
		case 0x0a: { /* '0x0a' */
            return 1024
		}
		case 0x0b: { /* '0x0b' */
            return 2048
		}
		case 0x0c: { /* '0x0c' */
            return 4096
		}
		case 0x0d: { /* '0x0d' */
            return 8192
		}
		default: {
			return 0
		}
	}
}

func COTPTpduSizeFirstEnumForFieldSizeInBytes(value uint16) (COTPTpduSize, error) {
	for _, sizeValue := range COTPTpduSizeValues {
		if sizeValue.SizeInBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing SizeInBytes not found", value)
}
func COTPTpduSizeByValue(value uint8) (enum COTPTpduSize, ok bool) {
	switch value {
		case 0x07:
			return COTPTpduSize_SIZE_128, true
		case 0x08:
			return COTPTpduSize_SIZE_256, true
		case 0x09:
			return COTPTpduSize_SIZE_512, true
		case 0x0a:
			return COTPTpduSize_SIZE_1024, true
		case 0x0b:
			return COTPTpduSize_SIZE_2048, true
		case 0x0c:
			return COTPTpduSize_SIZE_4096, true
		case 0x0d:
			return COTPTpduSize_SIZE_8192, true
	}
	return 0, false
}

func COTPTpduSizeByName(value string) (enum COTPTpduSize, ok bool) {
	switch value {
	case "SIZE_128":
		return COTPTpduSize_SIZE_128, true
	case "SIZE_256":
		return COTPTpduSize_SIZE_256, true
	case "SIZE_512":
		return COTPTpduSize_SIZE_512, true
	case "SIZE_1024":
		return COTPTpduSize_SIZE_1024, true
	case "SIZE_2048":
		return COTPTpduSize_SIZE_2048, true
	case "SIZE_4096":
		return COTPTpduSize_SIZE_4096, true
	case "SIZE_8192":
		return COTPTpduSize_SIZE_8192, true
	}
	return 0, false
}

func COTPTpduSizeKnows(value uint8)  bool {
	for _, typeValue := range COTPTpduSizeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastCOTPTpduSize(structType interface{}) COTPTpduSize {
	castFunc := func(typ interface{}) COTPTpduSize {
		if sCOTPTpduSize, ok := typ.(COTPTpduSize); ok {
			return sCOTPTpduSize
		}
		return 0
	}
	return castFunc(structType)
}

func (m COTPTpduSize) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m COTPTpduSize) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPTpduSizeParse(ctx context.Context, theBytes []byte) (COTPTpduSize, error) {
	return COTPTpduSizeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func COTPTpduSizeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (COTPTpduSize, error) {
	val, err := readBuffer.ReadUint8("COTPTpduSize", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading COTPTpduSize")
	}
	if enum, ok := COTPTpduSizeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return COTPTpduSize(val), nil
	} else {
		return enum, nil
	}
}

func (e COTPTpduSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e COTPTpduSize) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("COTPTpduSize", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e COTPTpduSize) PLC4XEnumName() string {
	switch e {
	case COTPTpduSize_SIZE_128:
		return "SIZE_128"
	case COTPTpduSize_SIZE_256:
		return "SIZE_256"
	case COTPTpduSize_SIZE_512:
		return "SIZE_512"
	case COTPTpduSize_SIZE_1024:
		return "SIZE_1024"
	case COTPTpduSize_SIZE_2048:
		return "SIZE_2048"
	case COTPTpduSize_SIZE_4096:
		return "SIZE_4096"
	case COTPTpduSize_SIZE_8192:
		return "SIZE_8192"
	}
	return ""
}

func (e COTPTpduSize) String() string {
	return e.PLC4XEnumName()
}

