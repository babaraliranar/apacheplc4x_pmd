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

// HVACTemperature is the corresponding interface of HVACTemperature
type HVACTemperature interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTemperatureValue returns TemperatureValue (property field)
	GetTemperatureValue() int16
	// GetTemperatureInCelcius returns TemperatureInCelcius (virtual field)
	GetTemperatureInCelcius() float32
}

// HVACTemperatureExactly can be used when we want exactly this type and not a type which fulfills HVACTemperature.
// This is useful for switch cases.
type HVACTemperatureExactly interface {
	HVACTemperature
	isHVACTemperature() bool
}

// _HVACTemperature is the data-structure of this message
type _HVACTemperature struct {
	TemperatureValue int16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACTemperature) GetTemperatureValue() int16 {
	return m.TemperatureValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACTemperature) GetTemperatureInCelcius() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetTemperatureValue()) / float32(float32(256)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACTemperature factory function for _HVACTemperature
func NewHVACTemperature(temperatureValue int16) *_HVACTemperature {
	return &_HVACTemperature{TemperatureValue: temperatureValue}
}

// Deprecated: use the interface for direct cast
func CastHVACTemperature(structType any) HVACTemperature {
	if casted, ok := structType.(HVACTemperature); ok {
		return casted
	}
	if casted, ok := structType.(*HVACTemperature); ok {
		return *casted
	}
	return nil
}

func (m *_HVACTemperature) GetTypeName() string {
	return "HVACTemperature"
}

func (m *_HVACTemperature) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (temperatureValue)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACTemperature) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACTemperatureParse(ctx context.Context, theBytes []byte) (HVACTemperature, error) {
	return HVACTemperatureParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACTemperatureParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACTemperature, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HVACTemperature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACTemperature")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (temperatureValue)
	_temperatureValue, _temperatureValueErr := readBuffer.ReadInt16("temperatureValue", 16)
	if _temperatureValueErr != nil {
		return nil, errors.Wrap(_temperatureValueErr, "Error parsing 'temperatureValue' field of HVACTemperature")
	}
	temperatureValue := _temperatureValue

	// Virtual field
	_temperatureInCelcius := float32(temperatureValue) / float32(float32(256))
	temperatureInCelcius := float32(_temperatureInCelcius)
	_ = temperatureInCelcius

	if closeErr := readBuffer.CloseContext("HVACTemperature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACTemperature")
	}

	// Create the instance
	return &_HVACTemperature{
		TemperatureValue: temperatureValue,
	}, nil
}

func (m *_HVACTemperature) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACTemperature) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACTemperature"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACTemperature")
	}

	// Simple Field (temperatureValue)
	temperatureValue := int16(m.GetTemperatureValue())
	_temperatureValueErr := writeBuffer.WriteInt16("temperatureValue", 16, int16((temperatureValue)))
	if _temperatureValueErr != nil {
		return errors.Wrap(_temperatureValueErr, "Error serializing 'temperatureValue' field")
	}
	// Virtual field
	temperatureInCelcius := m.GetTemperatureInCelcius()
	_ = temperatureInCelcius
	if _temperatureInCelciusErr := writeBuffer.WriteVirtual(ctx, "temperatureInCelcius", m.GetTemperatureInCelcius()); _temperatureInCelciusErr != nil {
		return errors.Wrap(_temperatureInCelciusErr, "Error serializing 'temperatureInCelcius' field")
	}

	if popErr := writeBuffer.PopContext("HVACTemperature"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACTemperature")
	}
	return nil
}

func (m *_HVACTemperature) isHVACTemperature() bool {
	return true
}

func (m *_HVACTemperature) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
