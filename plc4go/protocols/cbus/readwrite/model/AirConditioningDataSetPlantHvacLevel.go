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
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AirConditioningDataSetPlantHvacLevel is the corresponding interface of AirConditioningDataSetPlantHvacLevel
type AirConditioningDataSetPlantHvacLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
	// GetHvacType returns HvacType (property field)
	GetHvacType() HVACType
	// GetLevel returns Level (property field)
	GetLevel() HVACTemperature
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
}

// AirConditioningDataSetPlantHvacLevelExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetPlantHvacLevel.
// This is useful for switch cases.
type AirConditioningDataSetPlantHvacLevelExactly interface {
	AirConditioningDataSetPlantHvacLevel
	isAirConditioningDataSetPlantHvacLevel() bool
}

// _AirConditioningDataSetPlantHvacLevel is the data-structure of this message
type _AirConditioningDataSetPlantHvacLevel struct {
	*_AirConditioningData
	ZoneGroup        byte
	ZoneList         HVACZoneList
	HvacModeAndFlags HVACModeAndFlags
	HvacType         HVACType
	Level            HVACTemperature
	RawLevel         HVACRawLevels
	AuxLevel         HVACAuxiliaryLevel
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetPlantHvacLevel) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetPlantHvacLevel) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetHvacType() HVACType {
	return m.HvacType
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetLevel() HVACTemperature {
	return m.Level
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetPlantHvacLevel factory function for _AirConditioningDataSetPlantHvacLevel
func NewAirConditioningDataSetPlantHvacLevel(zoneGroup byte, zoneList HVACZoneList, hvacModeAndFlags HVACModeAndFlags, hvacType HVACType, level HVACTemperature, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetPlantHvacLevel {
	_result := &_AirConditioningDataSetPlantHvacLevel{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		HvacModeAndFlags:     hvacModeAndFlags,
		HvacType:             hvacType,
		Level:                level,
		RawLevel:             rawLevel,
		AuxLevel:             auxLevel,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetPlantHvacLevel(structType any) AirConditioningDataSetPlantHvacLevel {
	if casted, ok := structType.(AirConditioningDataSetPlantHvacLevel); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetPlantHvacLevel); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetTypeName() string {
	return "AirConditioningDataSetPlantHvacLevel"
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	// Simple field (hvacType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataSetPlantHvacLevelParse(ctx context.Context, theBytes []byte) (AirConditioningDataSetPlantHvacLevel, error) {
	return AirConditioningDataSetPlantHvacLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataSetPlantHvacLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataSetPlantHvacLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AirConditioningDataSetPlantHvacLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetPlantHvacLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataSetPlantHvacLevel")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParseWithBuffer(ctx, readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataSetPlantHvacLevel")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (hvacModeAndFlags)
	if pullErr := readBuffer.PullContext("hvacModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacModeAndFlags")
	}
	_hvacModeAndFlags, _hvacModeAndFlagsErr := HVACModeAndFlagsParseWithBuffer(ctx, readBuffer)
	if _hvacModeAndFlagsErr != nil {
		return nil, errors.Wrap(_hvacModeAndFlagsErr, "Error parsing 'hvacModeAndFlags' field of AirConditioningDataSetPlantHvacLevel")
	}
	hvacModeAndFlags := _hvacModeAndFlags.(HVACModeAndFlags)
	if closeErr := readBuffer.CloseContext("hvacModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacModeAndFlags")
	}

	// Simple Field (hvacType)
	if pullErr := readBuffer.PullContext("hvacType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacType")
	}
	_hvacType, _hvacTypeErr := HVACTypeParseWithBuffer(ctx, readBuffer)
	if _hvacTypeErr != nil {
		return nil, errors.Wrap(_hvacTypeErr, "Error parsing 'hvacType' field of AirConditioningDataSetPlantHvacLevel")
	}
	hvacType := _hvacType
	if closeErr := readBuffer.CloseContext("hvacType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacType")
	}

	// Optional Field (level) (Can be skipped, if a given expression evaluates to false)
	var level HVACTemperature = nil
	if hvacModeAndFlags.GetIsLevelTemperature() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("level"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for level")
		}
		_val, _err := HVACTemperatureParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'level' field of AirConditioningDataSetPlantHvacLevel")
		default:
			level = _val.(HVACTemperature)
			if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for level")
			}
		}
	}

	// Optional Field (rawLevel) (Can be skipped, if a given expression evaluates to false)
	var rawLevel HVACRawLevels = nil
	if hvacModeAndFlags.GetIsLevelRaw() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("rawLevel"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for rawLevel")
		}
		_val, _err := HVACRawLevelsParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawLevel' field of AirConditioningDataSetPlantHvacLevel")
		default:
			rawLevel = _val.(HVACRawLevels)
			if closeErr := readBuffer.CloseContext("rawLevel"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for rawLevel")
			}
		}
	}

	// Optional Field (auxLevel) (Can be skipped, if a given expression evaluates to false)
	var auxLevel HVACAuxiliaryLevel = nil
	if hvacModeAndFlags.GetIsAuxLevelUsed() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("auxLevel"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for auxLevel")
		}
		_val, _err := HVACAuxiliaryLevelParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'auxLevel' field of AirConditioningDataSetPlantHvacLevel")
		default:
			auxLevel = _val.(HVACAuxiliaryLevel)
			if closeErr := readBuffer.CloseContext("auxLevel"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for auxLevel")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetPlantHvacLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetPlantHvacLevel")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetPlantHvacLevel{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		HvacModeAndFlags:     hvacModeAndFlags,
		HvacType:             hvacType,
		Level:                level,
		RawLevel:             rawLevel,
		AuxLevel:             auxLevel,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetPlantHvacLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetPlantHvacLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetPlantHvacLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetPlantHvacLevel")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(ctx, m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (hvacModeAndFlags)
		if pushErr := writeBuffer.PushContext("hvacModeAndFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hvacModeAndFlags")
		}
		_hvacModeAndFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetHvacModeAndFlags())
		if popErr := writeBuffer.PopContext("hvacModeAndFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hvacModeAndFlags")
		}
		if _hvacModeAndFlagsErr != nil {
			return errors.Wrap(_hvacModeAndFlagsErr, "Error serializing 'hvacModeAndFlags' field")
		}

		// Simple Field (hvacType)
		if pushErr := writeBuffer.PushContext("hvacType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hvacType")
		}
		_hvacTypeErr := writeBuffer.WriteSerializable(ctx, m.GetHvacType())
		if popErr := writeBuffer.PopContext("hvacType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hvacType")
		}
		if _hvacTypeErr != nil {
			return errors.Wrap(_hvacTypeErr, "Error serializing 'hvacType' field")
		}

		// Optional Field (level) (Can be skipped, if the value is null)
		var level HVACTemperature = nil
		if m.GetLevel() != nil {
			if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for level")
			}
			level = m.GetLevel()
			_levelErr := writeBuffer.WriteSerializable(ctx, level)
			if popErr := writeBuffer.PopContext("level"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for level")
			}
			if _levelErr != nil {
				return errors.Wrap(_levelErr, "Error serializing 'level' field")
			}
		}

		// Optional Field (rawLevel) (Can be skipped, if the value is null)
		var rawLevel HVACRawLevels = nil
		if m.GetRawLevel() != nil {
			if pushErr := writeBuffer.PushContext("rawLevel"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for rawLevel")
			}
			rawLevel = m.GetRawLevel()
			_rawLevelErr := writeBuffer.WriteSerializable(ctx, rawLevel)
			if popErr := writeBuffer.PopContext("rawLevel"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for rawLevel")
			}
			if _rawLevelErr != nil {
				return errors.Wrap(_rawLevelErr, "Error serializing 'rawLevel' field")
			}
		}

		// Optional Field (auxLevel) (Can be skipped, if the value is null)
		var auxLevel HVACAuxiliaryLevel = nil
		if m.GetAuxLevel() != nil {
			if pushErr := writeBuffer.PushContext("auxLevel"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for auxLevel")
			}
			auxLevel = m.GetAuxLevel()
			_auxLevelErr := writeBuffer.WriteSerializable(ctx, auxLevel)
			if popErr := writeBuffer.PopContext("auxLevel"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for auxLevel")
			}
			if _auxLevelErr != nil {
				return errors.Wrap(_auxLevelErr, "Error serializing 'auxLevel' field")
			}
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetPlantHvacLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetPlantHvacLevel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetPlantHvacLevel) isAirConditioningDataSetPlantHvacLevel() bool {
	return true
}

func (m *_AirConditioningDataSetPlantHvacLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
