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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SysexCommandPinStateResponse is the corresponding interface of SysexCommandPinStateResponse
type SysexCommandPinStateResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetPinMode returns PinMode (property field)
	GetPinMode() uint8
	// GetPinState returns PinState (property field)
	GetPinState() uint8
	// IsSysexCommandPinStateResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandPinStateResponse()
	// CreateBuilder creates a SysexCommandPinStateResponseBuilder
	CreateSysexCommandPinStateResponseBuilder() SysexCommandPinStateResponseBuilder
}

// _SysexCommandPinStateResponse is the data-structure of this message
type _SysexCommandPinStateResponse struct {
	SysexCommandContract
	Pin      uint8
	PinMode  uint8
	PinState uint8
}

var _ SysexCommandPinStateResponse = (*_SysexCommandPinStateResponse)(nil)
var _ SysexCommandRequirements = (*_SysexCommandPinStateResponse)(nil)

// NewSysexCommandPinStateResponse factory function for _SysexCommandPinStateResponse
func NewSysexCommandPinStateResponse(pin uint8, pinMode uint8, pinState uint8) *_SysexCommandPinStateResponse {
	_result := &_SysexCommandPinStateResponse{
		SysexCommandContract: NewSysexCommand(),
		Pin:                  pin,
		PinMode:              pinMode,
		PinState:             pinState,
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandPinStateResponseBuilder is a builder for SysexCommandPinStateResponse
type SysexCommandPinStateResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pin uint8, pinMode uint8, pinState uint8) SysexCommandPinStateResponseBuilder
	// WithPin adds Pin (property field)
	WithPin(uint8) SysexCommandPinStateResponseBuilder
	// WithPinMode adds PinMode (property field)
	WithPinMode(uint8) SysexCommandPinStateResponseBuilder
	// WithPinState adds PinState (property field)
	WithPinState(uint8) SysexCommandPinStateResponseBuilder
	// Build builds the SysexCommandPinStateResponse or returns an error if something is wrong
	Build() (SysexCommandPinStateResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandPinStateResponse
}

// NewSysexCommandPinStateResponseBuilder() creates a SysexCommandPinStateResponseBuilder
func NewSysexCommandPinStateResponseBuilder() SysexCommandPinStateResponseBuilder {
	return &_SysexCommandPinStateResponseBuilder{_SysexCommandPinStateResponse: new(_SysexCommandPinStateResponse)}
}

type _SysexCommandPinStateResponseBuilder struct {
	*_SysexCommandPinStateResponse

	err *utils.MultiError
}

var _ (SysexCommandPinStateResponseBuilder) = (*_SysexCommandPinStateResponseBuilder)(nil)

func (m *_SysexCommandPinStateResponseBuilder) WithMandatoryFields(pin uint8, pinMode uint8, pinState uint8) SysexCommandPinStateResponseBuilder {
	return m.WithPin(pin).WithPinMode(pinMode).WithPinState(pinState)
}

func (m *_SysexCommandPinStateResponseBuilder) WithPin(pin uint8) SysexCommandPinStateResponseBuilder {
	m.Pin = pin
	return m
}

func (m *_SysexCommandPinStateResponseBuilder) WithPinMode(pinMode uint8) SysexCommandPinStateResponseBuilder {
	m.PinMode = pinMode
	return m
}

func (m *_SysexCommandPinStateResponseBuilder) WithPinState(pinState uint8) SysexCommandPinStateResponseBuilder {
	m.PinState = pinState
	return m
}

func (m *_SysexCommandPinStateResponseBuilder) Build() (SysexCommandPinStateResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SysexCommandPinStateResponse.deepCopy(), nil
}

func (m *_SysexCommandPinStateResponseBuilder) MustBuild() SysexCommandPinStateResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SysexCommandPinStateResponseBuilder) DeepCopy() any {
	return m.CreateSysexCommandPinStateResponseBuilder()
}

// CreateSysexCommandPinStateResponseBuilder creates a SysexCommandPinStateResponseBuilder
func (m *_SysexCommandPinStateResponse) CreateSysexCommandPinStateResponseBuilder() SysexCommandPinStateResponseBuilder {
	if m == nil {
		return NewSysexCommandPinStateResponseBuilder()
	}
	return &_SysexCommandPinStateResponseBuilder{_SysexCommandPinStateResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandPinStateResponse) GetCommandType() uint8 {
	return 0x6E
}

func (m *_SysexCommandPinStateResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandPinStateResponse) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandPinStateResponse) GetPin() uint8 {
	return m.Pin
}

func (m *_SysexCommandPinStateResponse) GetPinMode() uint8 {
	return m.PinMode
}

func (m *_SysexCommandPinStateResponse) GetPinState() uint8 {
	return m.PinState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSysexCommandPinStateResponse(structType any) SysexCommandPinStateResponse {
	if casted, ok := structType.(SysexCommandPinStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandPinStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandPinStateResponse) GetTypeName() string {
	return "SysexCommandPinStateResponse"
}

func (m *_SysexCommandPinStateResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 8

	// Simple field (pinMode)
	lengthInBits += 8

	// Simple field (pinState)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommandPinStateResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandPinStateResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandPinStateResponse SysexCommandPinStateResponse, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandPinStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandPinStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}
	m.Pin = pin

	pinMode, err := ReadSimpleField(ctx, "pinMode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pinMode' field"))
	}
	m.PinMode = pinMode

	pinState, err := ReadSimpleField(ctx, "pinState", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pinState' field"))
	}
	m.PinState = pinState

	if closeErr := readBuffer.CloseContext("SysexCommandPinStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandPinStateResponse")
	}

	return m, nil
}

func (m *_SysexCommandPinStateResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandPinStateResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandPinStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandPinStateResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if err := WriteSimpleField[uint8](ctx, "pinMode", m.GetPinMode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pinMode' field")
		}

		if err := WriteSimpleField[uint8](ctx, "pinState", m.GetPinState(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pinState' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandPinStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandPinStateResponse")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandPinStateResponse) IsSysexCommandPinStateResponse() {}

func (m *_SysexCommandPinStateResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandPinStateResponse) deepCopy() *_SysexCommandPinStateResponse {
	if m == nil {
		return nil
	}
	_SysexCommandPinStateResponseCopy := &_SysexCommandPinStateResponse{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
		m.Pin,
		m.PinMode,
		m.PinState,
	}
	m.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandPinStateResponseCopy
}

func (m *_SysexCommandPinStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
