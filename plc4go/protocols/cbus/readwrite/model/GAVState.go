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

// GAVState is an enum
type GAVState uint8

type IGAVState interface {
	utils.Serializable
}

const (
	GAVState_DOES_NOT_EXIST GAVState = 0
	GAVState_ON             GAVState = 1
	GAVState_OFF            GAVState = 2
	GAVState_ERROR          GAVState = 3
)

var GAVStateValues []GAVState

func init() {
	_ = errors.New
	GAVStateValues = []GAVState{
		GAVState_DOES_NOT_EXIST,
		GAVState_ON,
		GAVState_OFF,
		GAVState_ERROR,
	}
}

func GAVStateByValue(value uint8) (enum GAVState, ok bool) {
	switch value {
	case 0:
		return GAVState_DOES_NOT_EXIST, true
	case 1:
		return GAVState_ON, true
	case 2:
		return GAVState_OFF, true
	case 3:
		return GAVState_ERROR, true
	}
	return 0, false
}

func GAVStateByName(value string) (enum GAVState, ok bool) {
	switch value {
	case "DOES_NOT_EXIST":
		return GAVState_DOES_NOT_EXIST, true
	case "ON":
		return GAVState_ON, true
	case "OFF":
		return GAVState_OFF, true
	case "ERROR":
		return GAVState_ERROR, true
	}
	return 0, false
}

func GAVStateKnows(value uint8) bool {
	for _, typeValue := range GAVStateValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastGAVState(structType interface{}) GAVState {
	castFunc := func(typ interface{}) GAVState {
		if sGAVState, ok := typ.(GAVState); ok {
			return sGAVState
		}
		return 0
	}
	return castFunc(structType)
}

func (m GAVState) GetLengthInBits() uint16 {
	return 2
}

func (m GAVState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func GAVStateParse(theBytes []byte) (GAVState, error) {
	return GAVStateParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func GAVStateParseWithBuffer(readBuffer utils.ReadBuffer) (GAVState, error) {
	val, err := readBuffer.ReadUint8("GAVState", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading GAVState")
	}
	if enum, ok := GAVStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return GAVState(val), nil
	} else {
		return enum, nil
	}
}

func (e GAVState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e GAVState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("GAVState", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e GAVState) PLC4XEnumName() string {
	switch e {
	case GAVState_DOES_NOT_EXIST:
		return "DOES_NOT_EXIST"
	case GAVState_ON:
		return "ON"
	case GAVState_OFF:
		return "OFF"
	case GAVState_ERROR:
		return "ERROR"
	}
	return ""
}

func (e GAVState) String() string {
	return e.PLC4XEnumName()
}
