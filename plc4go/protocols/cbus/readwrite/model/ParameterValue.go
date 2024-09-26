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

// ParameterValue is the corresponding interface of ParameterValue
type ParameterValue interface {
	ParameterValueContract
	ParameterValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsParameterValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValue()
}

// ParameterValueContract provides a set of functions which can be overwritten by a sub struct
type ParameterValueContract interface {
	// GetNumBytes() returns a parser argument
	GetNumBytes() uint8
	// IsParameterValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValue()
}

// ParameterValueRequirements provides a set of functions which need to be implemented by a sub struct
type ParameterValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() ParameterType
}

// _ParameterValue is the data-structure of this message
type _ParameterValue struct {
	_SubType ParameterValue

	// Arguments.
	NumBytes uint8
}

var _ ParameterValueContract = (*_ParameterValue)(nil)

// NewParameterValue factory function for _ParameterValue
func NewParameterValue(numBytes uint8) *_ParameterValue {
	return &_ParameterValue{NumBytes: numBytes}
}

// Deprecated: use the interface for direct cast
func CastParameterValue(structType any) ParameterValue {
	if casted, ok := structType.(ParameterValue); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValue); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValue) GetTypeName() string {
	return "ParameterValue"
}

func (m *_ParameterValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ParameterValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ParameterValueParse[T ParameterValue](ctx context.Context, theBytes []byte, parameterType ParameterType, numBytes uint8) (T, error) {
	return ParameterValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueParseWithBufferProducer[T ParameterValue](parameterType ParameterType, numBytes uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ParameterValueParseWithBuffer[T](ctx, readBuffer, parameterType, numBytes)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ParameterValueParseWithBuffer[T ParameterValue](ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (T, error) {
	v, err := (&_ParameterValue{NumBytes: numBytes}).parse(ctx, readBuffer, parameterType, numBytes)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ParameterValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (__parameterValue ParameterValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ParameterValue
	switch {
	case parameterType == ParameterType_APPLICATION_ADDRESS_1: // ParameterValueApplicationAddress1
		if _child, err = new(_ParameterValueApplicationAddress1).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueApplicationAddress1 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_APPLICATION_ADDRESS_2: // ParameterValueApplicationAddress2
		if _child, err = new(_ParameterValueApplicationAddress2).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueApplicationAddress2 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_1: // ParameterValueInterfaceOptions1
		if _child, err = new(_ParameterValueInterfaceOptions1).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions1 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_BAUD_RATE_SELECTOR: // ParameterValueBaudRateSelector
		if _child, err = new(_ParameterValueBaudRateSelector).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueBaudRateSelector for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_2: // ParameterValueInterfaceOptions2
		if _child, err = new(_ParameterValueInterfaceOptions2).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions2 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS: // ParameterValueInterfaceOptions1PowerUpSettings
		if _child, err = new(_ParameterValueInterfaceOptions1PowerUpSettings).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions1PowerUpSettings for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_3: // ParameterValueInterfaceOptions3
		if _child, err = new(_ParameterValueInterfaceOptions3).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions3 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_CUSTOM_MANUFACTURER: // ParameterValueCustomManufacturer
		if _child, err = new(_ParameterValueCustomManufacturer).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueCustomManufacturer for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_SERIAL_NUMBER: // ParameterValueSerialNumber
		if _child, err = new(_ParameterValueSerialNumber).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueSerialNumber for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_CUSTOM_TYPE: // ParameterValueCustomTypes
		if _child, err = new(_ParameterValueCustomTypes).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueCustomTypes for type-switch of ParameterValue")
		}
	case 0 == 0: // ParameterValueRaw
		if _child, err = new(_ParameterValueRaw).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueRaw for type-switch of ParameterValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [parameterType=%v]", parameterType)
	}

	if closeErr := readBuffer.CloseContext("ParameterValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValue")
	}

	return _child, nil
}

func (pm *_ParameterValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ParameterValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ParameterValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ParameterValue")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ParameterValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ParameterValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_ParameterValue) GetNumBytes() uint8 {
	return m.NumBytes
}

//
////

func (m *_ParameterValue) IsParameterValue() {}

func (m *_ParameterValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValue) deepCopy() *_ParameterValue {
	if m == nil {
		return nil
	}
	_ParameterValueCopy := &_ParameterValue{
		nil, // will be set by child
		m.NumBytes,
	}
	return _ParameterValueCopy
}
