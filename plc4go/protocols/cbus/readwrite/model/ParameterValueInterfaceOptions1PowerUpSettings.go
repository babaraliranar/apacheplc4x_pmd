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


// ParameterValueInterfaceOptions1PowerUpSettings is the corresponding interface of ParameterValueInterfaceOptions1PowerUpSettings
type ParameterValueInterfaceOptions1PowerUpSettings interface {
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() InterfaceOptions1PowerUpSettings
}

// ParameterValueInterfaceOptions1PowerUpSettingsExactly can be used when we want exactly this type and not a type which fulfills ParameterValueInterfaceOptions1PowerUpSettings.
// This is useful for switch cases.
type ParameterValueInterfaceOptions1PowerUpSettingsExactly interface {
	ParameterValueInterfaceOptions1PowerUpSettings
	isParameterValueInterfaceOptions1PowerUpSettings() bool
}

// _ParameterValueInterfaceOptions1PowerUpSettings is the data-structure of this message
type _ParameterValueInterfaceOptions1PowerUpSettings struct {
	*_ParameterValue
        Value InterfaceOptions1PowerUpSettings
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueInterfaceOptions1PowerUpSettings)  GetParameterType() ParameterType {
return ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) InitializeParent(parent ParameterValue ) {}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings)  GetParent() ParameterValue {
	return m._ParameterValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetValue() InterfaceOptions1PowerUpSettings {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewParameterValueInterfaceOptions1PowerUpSettings factory function for _ParameterValueInterfaceOptions1PowerUpSettings
func NewParameterValueInterfaceOptions1PowerUpSettings( value InterfaceOptions1PowerUpSettings , numBytes uint8 ) *_ParameterValueInterfaceOptions1PowerUpSettings {
	_result := &_ParameterValueInterfaceOptions1PowerUpSettings{
		Value: value,
    	_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueInterfaceOptions1PowerUpSettings(structType interface{}) ParameterValueInterfaceOptions1PowerUpSettings {
    if casted, ok := structType.(ParameterValueInterfaceOptions1PowerUpSettings); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueInterfaceOptions1PowerUpSettings); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetTypeName() string {
	return "ParameterValueInterfaceOptions1PowerUpSettings"
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueInterfaceOptions1PowerUpSettingsParse(theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValueInterfaceOptions1PowerUpSettings, error) {
	return ParameterValueInterfaceOptions1PowerUpSettingsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueInterfaceOptions1PowerUpSettingsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueInterfaceOptions1PowerUpSettings, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueInterfaceOptions1PowerUpSettings"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueInterfaceOptions1PowerUpSettings")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if (!(bool((numBytes) >= ((1))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"InterfaceOptions1PowerUpSettings has exactly one byte"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
_value, _valueErr := InterfaceOptions1PowerUpSettingsParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueInterfaceOptions1PowerUpSettings")
	}
	value := _value.(InterfaceOptions1PowerUpSettings)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueInterfaceOptions1PowerUpSettings"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueInterfaceOptions1PowerUpSettings")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueInterfaceOptions1PowerUpSettings{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueInterfaceOptions1PowerUpSettings"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueInterfaceOptions1PowerUpSettings")
		}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("ParameterValueInterfaceOptions1PowerUpSettings"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueInterfaceOptions1PowerUpSettings")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ParameterValueInterfaceOptions1PowerUpSettings) isParameterValueInterfaceOptions1PowerUpSettings() bool {
	return true
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



