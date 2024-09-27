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

// ApduControlNack is the corresponding interface of ApduControlNack
type ApduControlNack interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduControl
	// IsApduControlNack is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduControlNack()
	// CreateBuilder creates a ApduControlNackBuilder
	CreateApduControlNackBuilder() ApduControlNackBuilder
}

// _ApduControlNack is the data-structure of this message
type _ApduControlNack struct {
	ApduControlContract
}

var _ ApduControlNack = (*_ApduControlNack)(nil)
var _ ApduControlRequirements = (*_ApduControlNack)(nil)

// NewApduControlNack factory function for _ApduControlNack
func NewApduControlNack() *_ApduControlNack {
	_result := &_ApduControlNack{
		ApduControlContract: NewApduControl(),
	}
	_result.ApduControlContract.(*_ApduControl)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduControlNackBuilder is a builder for ApduControlNack
type ApduControlNackBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduControlNackBuilder
	// Build builds the ApduControlNack or returns an error if something is wrong
	Build() (ApduControlNack, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduControlNack
}

// NewApduControlNackBuilder() creates a ApduControlNackBuilder
func NewApduControlNackBuilder() ApduControlNackBuilder {
	return &_ApduControlNackBuilder{_ApduControlNack: new(_ApduControlNack)}
}

type _ApduControlNackBuilder struct {
	*_ApduControlNack

	err *utils.MultiError
}

var _ (ApduControlNackBuilder) = (*_ApduControlNackBuilder)(nil)

func (m *_ApduControlNackBuilder) WithMandatoryFields() ApduControlNackBuilder {
	return m
}

func (m *_ApduControlNackBuilder) Build() (ApduControlNack, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduControlNack.deepCopy(), nil
}

func (m *_ApduControlNackBuilder) MustBuild() ApduControlNack {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduControlNackBuilder) DeepCopy() any {
	return m.CreateApduControlNackBuilder()
}

// CreateApduControlNackBuilder creates a ApduControlNackBuilder
func (m *_ApduControlNack) CreateApduControlNackBuilder() ApduControlNackBuilder {
	if m == nil {
		return NewApduControlNackBuilder()
	}
	return &_ApduControlNackBuilder{_ApduControlNack: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlNack) GetControlType() uint8 {
	return 0x3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlNack) GetParent() ApduControlContract {
	return m.ApduControlContract
}

// Deprecated: use the interface for direct cast
func CastApduControlNack(structType any) ApduControlNack {
	if casted, ok := structType.(ApduControlNack); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlNack); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlNack) GetTypeName() string {
	return "ApduControlNack"
}

func (m *_ApduControlNack) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduControlContract.(*_ApduControl).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduControlNack) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduControlNack) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduControl) (__apduControlNack ApduControlNack, err error) {
	m.ApduControlContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlNack"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlNack")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlNack"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlNack")
	}

	return m, nil
}

func (m *_ApduControlNack) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlNack) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlNack"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlNack")
		}

		if popErr := writeBuffer.PopContext("ApduControlNack"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlNack")
		}
		return nil
	}
	return m.ApduControlContract.(*_ApduControl).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduControlNack) IsApduControlNack() {}

func (m *_ApduControlNack) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduControlNack) deepCopy() *_ApduControlNack {
	if m == nil {
		return nil
	}
	_ApduControlNackCopy := &_ApduControlNack{
		m.ApduControlContract.(*_ApduControl).deepCopy(),
	}
	m.ApduControlContract.(*_ApduControl)._SubType = m
	return _ApduControlNackCopy
}

func (m *_ApduControlNack) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
