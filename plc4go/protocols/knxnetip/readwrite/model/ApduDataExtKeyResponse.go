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

// ApduDataExtKeyResponse is the corresponding interface of ApduDataExtKeyResponse
type ApduDataExtKeyResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtKeyResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtKeyResponse()
	// CreateBuilder creates a ApduDataExtKeyResponseBuilder
	CreateApduDataExtKeyResponseBuilder() ApduDataExtKeyResponseBuilder
}

// _ApduDataExtKeyResponse is the data-structure of this message
type _ApduDataExtKeyResponse struct {
	ApduDataExtContract
}

var _ ApduDataExtKeyResponse = (*_ApduDataExtKeyResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtKeyResponse)(nil)

// NewApduDataExtKeyResponse factory function for _ApduDataExtKeyResponse
func NewApduDataExtKeyResponse(length uint8) *_ApduDataExtKeyResponse {
	_result := &_ApduDataExtKeyResponse{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtKeyResponseBuilder is a builder for ApduDataExtKeyResponse
type ApduDataExtKeyResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtKeyResponseBuilder
	// Build builds the ApduDataExtKeyResponse or returns an error if something is wrong
	Build() (ApduDataExtKeyResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtKeyResponse
}

// NewApduDataExtKeyResponseBuilder() creates a ApduDataExtKeyResponseBuilder
func NewApduDataExtKeyResponseBuilder() ApduDataExtKeyResponseBuilder {
	return &_ApduDataExtKeyResponseBuilder{_ApduDataExtKeyResponse: new(_ApduDataExtKeyResponse)}
}

type _ApduDataExtKeyResponseBuilder struct {
	*_ApduDataExtKeyResponse

	err *utils.MultiError
}

var _ (ApduDataExtKeyResponseBuilder) = (*_ApduDataExtKeyResponseBuilder)(nil)

func (m *_ApduDataExtKeyResponseBuilder) WithMandatoryFields() ApduDataExtKeyResponseBuilder {
	return m
}

func (m *_ApduDataExtKeyResponseBuilder) Build() (ApduDataExtKeyResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduDataExtKeyResponse.deepCopy(), nil
}

func (m *_ApduDataExtKeyResponseBuilder) MustBuild() ApduDataExtKeyResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduDataExtKeyResponseBuilder) DeepCopy() any {
	return m.CreateApduDataExtKeyResponseBuilder()
}

// CreateApduDataExtKeyResponseBuilder creates a ApduDataExtKeyResponseBuilder
func (m *_ApduDataExtKeyResponse) CreateApduDataExtKeyResponseBuilder() ApduDataExtKeyResponseBuilder {
	if m == nil {
		return NewApduDataExtKeyResponseBuilder()
	}
	return &_ApduDataExtKeyResponseBuilder{_ApduDataExtKeyResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtKeyResponse) GetExtApciType() uint8 {
	return 0x14
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtKeyResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtKeyResponse(structType any) ApduDataExtKeyResponse {
	if casted, ok := structType.(ApduDataExtKeyResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtKeyResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtKeyResponse) GetTypeName() string {
	return "ApduDataExtKeyResponse"
}

func (m *_ApduDataExtKeyResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtKeyResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtKeyResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtKeyResponse ApduDataExtKeyResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtKeyResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtKeyResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtKeyResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtKeyResponse")
	}

	return m, nil
}

func (m *_ApduDataExtKeyResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtKeyResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtKeyResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtKeyResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtKeyResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtKeyResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtKeyResponse) IsApduDataExtKeyResponse() {}

func (m *_ApduDataExtKeyResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtKeyResponse) deepCopy() *_ApduDataExtKeyResponse {
	if m == nil {
		return nil
	}
	_ApduDataExtKeyResponseCopy := &_ApduDataExtKeyResponse{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtKeyResponseCopy
}

func (m *_ApduDataExtKeyResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
