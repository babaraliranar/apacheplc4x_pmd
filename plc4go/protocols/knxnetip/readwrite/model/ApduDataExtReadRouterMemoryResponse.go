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

// ApduDataExtReadRouterMemoryResponse is the corresponding interface of ApduDataExtReadRouterMemoryResponse
type ApduDataExtReadRouterMemoryResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtReadRouterMemoryResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtReadRouterMemoryResponse()
	// CreateBuilder creates a ApduDataExtReadRouterMemoryResponseBuilder
	CreateApduDataExtReadRouterMemoryResponseBuilder() ApduDataExtReadRouterMemoryResponseBuilder
}

// _ApduDataExtReadRouterMemoryResponse is the data-structure of this message
type _ApduDataExtReadRouterMemoryResponse struct {
	ApduDataExtContract
}

var _ ApduDataExtReadRouterMemoryResponse = (*_ApduDataExtReadRouterMemoryResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtReadRouterMemoryResponse)(nil)

// NewApduDataExtReadRouterMemoryResponse factory function for _ApduDataExtReadRouterMemoryResponse
func NewApduDataExtReadRouterMemoryResponse(length uint8) *_ApduDataExtReadRouterMemoryResponse {
	_result := &_ApduDataExtReadRouterMemoryResponse{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtReadRouterMemoryResponseBuilder is a builder for ApduDataExtReadRouterMemoryResponse
type ApduDataExtReadRouterMemoryResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtReadRouterMemoryResponseBuilder
	// Build builds the ApduDataExtReadRouterMemoryResponse or returns an error if something is wrong
	Build() (ApduDataExtReadRouterMemoryResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtReadRouterMemoryResponse
}

// NewApduDataExtReadRouterMemoryResponseBuilder() creates a ApduDataExtReadRouterMemoryResponseBuilder
func NewApduDataExtReadRouterMemoryResponseBuilder() ApduDataExtReadRouterMemoryResponseBuilder {
	return &_ApduDataExtReadRouterMemoryResponseBuilder{_ApduDataExtReadRouterMemoryResponse: new(_ApduDataExtReadRouterMemoryResponse)}
}

type _ApduDataExtReadRouterMemoryResponseBuilder struct {
	*_ApduDataExtReadRouterMemoryResponse

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtReadRouterMemoryResponseBuilder) = (*_ApduDataExtReadRouterMemoryResponseBuilder)(nil)

func (b *_ApduDataExtReadRouterMemoryResponseBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtReadRouterMemoryResponseBuilder) WithMandatoryFields() ApduDataExtReadRouterMemoryResponseBuilder {
	return b
}

func (b *_ApduDataExtReadRouterMemoryResponseBuilder) Build() (ApduDataExtReadRouterMemoryResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtReadRouterMemoryResponse.deepCopy(), nil
}

func (b *_ApduDataExtReadRouterMemoryResponseBuilder) MustBuild() ApduDataExtReadRouterMemoryResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtReadRouterMemoryResponseBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtReadRouterMemoryResponseBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtReadRouterMemoryResponseBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtReadRouterMemoryResponseBuilder().(*_ApduDataExtReadRouterMemoryResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtReadRouterMemoryResponseBuilder creates a ApduDataExtReadRouterMemoryResponseBuilder
func (b *_ApduDataExtReadRouterMemoryResponse) CreateApduDataExtReadRouterMemoryResponseBuilder() ApduDataExtReadRouterMemoryResponseBuilder {
	if b == nil {
		return NewApduDataExtReadRouterMemoryResponseBuilder()
	}
	return &_ApduDataExtReadRouterMemoryResponseBuilder{_ApduDataExtReadRouterMemoryResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterMemoryResponse) GetExtApciType() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterMemoryResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterMemoryResponse(structType any) ApduDataExtReadRouterMemoryResponse {
	if casted, ok := structType.(ApduDataExtReadRouterMemoryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterMemoryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterMemoryResponse) GetTypeName() string {
	return "ApduDataExtReadRouterMemoryResponse"
}

func (m *_ApduDataExtReadRouterMemoryResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRouterMemoryResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtReadRouterMemoryResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtReadRouterMemoryResponse ApduDataExtReadRouterMemoryResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterMemoryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterMemoryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterMemoryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterMemoryResponse")
	}

	return m, nil
}

func (m *_ApduDataExtReadRouterMemoryResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterMemoryResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterMemoryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterMemoryResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterMemoryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterMemoryResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterMemoryResponse) IsApduDataExtReadRouterMemoryResponse() {}

func (m *_ApduDataExtReadRouterMemoryResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtReadRouterMemoryResponse) deepCopy() *_ApduDataExtReadRouterMemoryResponse {
	if m == nil {
		return nil
	}
	_ApduDataExtReadRouterMemoryResponseCopy := &_ApduDataExtReadRouterMemoryResponse{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtReadRouterMemoryResponseCopy
}

func (m *_ApduDataExtReadRouterMemoryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
