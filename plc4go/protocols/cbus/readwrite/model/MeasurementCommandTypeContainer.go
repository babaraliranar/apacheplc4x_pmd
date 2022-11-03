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

// MeasurementCommandTypeContainer is an enum
type MeasurementCommandTypeContainer uint8

type IMeasurementCommandTypeContainer interface {
	utils.Serializable
	NumBytes() uint8
	CommandType() MeasurementCommandType
}

const (
	MeasurementCommandTypeContainer_MeasurementCommandChannelMeasurementData MeasurementCommandTypeContainer = 0x0E
)

var MeasurementCommandTypeContainerValues []MeasurementCommandTypeContainer

func init() {
	_ = errors.New
	MeasurementCommandTypeContainerValues = []MeasurementCommandTypeContainer{
		MeasurementCommandTypeContainer_MeasurementCommandChannelMeasurementData,
	}
}

func (e MeasurementCommandTypeContainer) NumBytes() uint8 {
	switch e {
	case 0x0E:
		{ /* '0x0E' */
			return 6
		}
	default:
		{
			return 0
		}
	}
}

func MeasurementCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (MeasurementCommandTypeContainer, error) {
	for _, sizeValue := range MeasurementCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e MeasurementCommandTypeContainer) CommandType() MeasurementCommandType {
	switch e {
	case 0x0E:
		{ /* '0x0E' */
			return MeasurementCommandType_MEASUREMENT_EVENT
		}
	default:
		{
			return 0
		}
	}
}

func MeasurementCommandTypeContainerFirstEnumForFieldCommandType(value MeasurementCommandType) (MeasurementCommandTypeContainer, error) {
	for _, sizeValue := range MeasurementCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func MeasurementCommandTypeContainerByValue(value uint8) (enum MeasurementCommandTypeContainer, ok bool) {
	switch value {
	case 0x0E:
		return MeasurementCommandTypeContainer_MeasurementCommandChannelMeasurementData, true
	}
	return 0, false
}

func MeasurementCommandTypeContainerByName(value string) (enum MeasurementCommandTypeContainer, ok bool) {
	switch value {
	case "MeasurementCommandChannelMeasurementData":
		return MeasurementCommandTypeContainer_MeasurementCommandChannelMeasurementData, true
	}
	return 0, false
}

func MeasurementCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range MeasurementCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMeasurementCommandTypeContainer(structType interface{}) MeasurementCommandTypeContainer {
	castFunc := func(typ interface{}) MeasurementCommandTypeContainer {
		if sMeasurementCommandTypeContainer, ok := typ.(MeasurementCommandTypeContainer); ok {
			return sMeasurementCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m MeasurementCommandTypeContainer) GetLengthInBits() uint16 {
	return 8
}

func (m MeasurementCommandTypeContainer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeasurementCommandTypeContainerParse(theBytes []byte) (MeasurementCommandTypeContainer, error) {
	return MeasurementCommandTypeContainerParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func MeasurementCommandTypeContainerParseWithBuffer(readBuffer utils.ReadBuffer) (MeasurementCommandTypeContainer, error) {
	val, err := readBuffer.ReadUint8("MeasurementCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MeasurementCommandTypeContainer")
	}
	if enum, ok := MeasurementCommandTypeContainerByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return MeasurementCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e MeasurementCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MeasurementCommandTypeContainer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MeasurementCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MeasurementCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case MeasurementCommandTypeContainer_MeasurementCommandChannelMeasurementData:
		return "MeasurementCommandChannelMeasurementData"
	}
	return ""
}

func (e MeasurementCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
