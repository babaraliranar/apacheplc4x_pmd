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
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetSegmentation is an enum
type BACnetSegmentation uint8

type IBACnetSegmentation interface {
	utils.Serializable
}

const (
	BACnetSegmentation_SEGMENTED_BOTH     BACnetSegmentation = 0
	BACnetSegmentation_SEGMENTED_TRANSMIT BACnetSegmentation = 1
	BACnetSegmentation_SEGMENTED_RECEIVE  BACnetSegmentation = 2
	BACnetSegmentation_NO_SEGMENTATION    BACnetSegmentation = 3
)

var BACnetSegmentationValues []BACnetSegmentation

func init() {
	_ = errors.New
	BACnetSegmentationValues = []BACnetSegmentation{
		BACnetSegmentation_SEGMENTED_BOTH,
		BACnetSegmentation_SEGMENTED_TRANSMIT,
		BACnetSegmentation_SEGMENTED_RECEIVE,
		BACnetSegmentation_NO_SEGMENTATION,
	}
}

func BACnetSegmentationByValue(value uint8) (enum BACnetSegmentation, ok bool) {
	switch value {
	case 0:
		return BACnetSegmentation_SEGMENTED_BOTH, true
	case 1:
		return BACnetSegmentation_SEGMENTED_TRANSMIT, true
	case 2:
		return BACnetSegmentation_SEGMENTED_RECEIVE, true
	case 3:
		return BACnetSegmentation_NO_SEGMENTATION, true
	}
	return 0, false
}

func BACnetSegmentationByName(value string) (enum BACnetSegmentation, ok bool) {
	switch value {
	case "SEGMENTED_BOTH":
		return BACnetSegmentation_SEGMENTED_BOTH, true
	case "SEGMENTED_TRANSMIT":
		return BACnetSegmentation_SEGMENTED_TRANSMIT, true
	case "SEGMENTED_RECEIVE":
		return BACnetSegmentation_SEGMENTED_RECEIVE, true
	case "NO_SEGMENTATION":
		return BACnetSegmentation_NO_SEGMENTATION, true
	}
	return 0, false
}

func BACnetSegmentationKnows(value uint8) bool {
	for _, typeValue := range BACnetSegmentationValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetSegmentation(structType interface{}) BACnetSegmentation {
	castFunc := func(typ interface{}) BACnetSegmentation {
		if sBACnetSegmentation, ok := typ.(BACnetSegmentation); ok {
			return sBACnetSegmentation
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetSegmentation) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetSegmentation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSegmentationParse(theBytes []byte) (BACnetSegmentation, error) {
	return BACnetSegmentationParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetSegmentationParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetSegmentation, error) {
	val, err := readBuffer.ReadUint8("BACnetSegmentation", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetSegmentation")
	}
	if enum, ok := BACnetSegmentationByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetSegmentation(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetSegmentation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetSegmentation) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetSegmentation", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetSegmentation) PLC4XEnumName() string {
	switch e {
	case BACnetSegmentation_SEGMENTED_BOTH:
		return "SEGMENTED_BOTH"
	case BACnetSegmentation_SEGMENTED_TRANSMIT:
		return "SEGMENTED_TRANSMIT"
	case BACnetSegmentation_SEGMENTED_RECEIVE:
		return "SEGMENTED_RECEIVE"
	case BACnetSegmentation_NO_SEGMENTATION:
		return "NO_SEGMENTATION"
	}
	return ""
}

func (e BACnetSegmentation) String() string {
	return e.PLC4XEnumName()
}
