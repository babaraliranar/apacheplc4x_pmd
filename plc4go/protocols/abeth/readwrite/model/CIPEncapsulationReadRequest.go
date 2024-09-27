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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CIPEncapsulationReadRequest is the corresponding interface of CIPEncapsulationReadRequest
type CIPEncapsulationReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CIPEncapsulationPacket
	// GetRequest returns Request (property field)
	GetRequest() DF1RequestMessage
	// IsCIPEncapsulationReadRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCIPEncapsulationReadRequest()
	// CreateBuilder creates a CIPEncapsulationReadRequestBuilder
	CreateCIPEncapsulationReadRequestBuilder() CIPEncapsulationReadRequestBuilder
}

// _CIPEncapsulationReadRequest is the data-structure of this message
type _CIPEncapsulationReadRequest struct {
	CIPEncapsulationPacketContract
	Request DF1RequestMessage
}

var _ CIPEncapsulationReadRequest = (*_CIPEncapsulationReadRequest)(nil)
var _ CIPEncapsulationPacketRequirements = (*_CIPEncapsulationReadRequest)(nil)

// NewCIPEncapsulationReadRequest factory function for _CIPEncapsulationReadRequest
func NewCIPEncapsulationReadRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32, request DF1RequestMessage) *_CIPEncapsulationReadRequest {
	if request == nil {
		panic("request of type DF1RequestMessage for CIPEncapsulationReadRequest must not be nil")
	}
	_result := &_CIPEncapsulationReadRequest{
		CIPEncapsulationPacketContract: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
		Request:                        request,
	}
	_result.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CIPEncapsulationReadRequestBuilder is a builder for CIPEncapsulationReadRequest
type CIPEncapsulationReadRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(request DF1RequestMessage) CIPEncapsulationReadRequestBuilder
	// WithRequest adds Request (property field)
	WithRequest(DF1RequestMessage) CIPEncapsulationReadRequestBuilder
	// Build builds the CIPEncapsulationReadRequest or returns an error if something is wrong
	Build() (CIPEncapsulationReadRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CIPEncapsulationReadRequest
}

// NewCIPEncapsulationReadRequestBuilder() creates a CIPEncapsulationReadRequestBuilder
func NewCIPEncapsulationReadRequestBuilder() CIPEncapsulationReadRequestBuilder {
	return &_CIPEncapsulationReadRequestBuilder{_CIPEncapsulationReadRequest: new(_CIPEncapsulationReadRequest)}
}

type _CIPEncapsulationReadRequestBuilder struct {
	*_CIPEncapsulationReadRequest

	err *utils.MultiError
}

var _ (CIPEncapsulationReadRequestBuilder) = (*_CIPEncapsulationReadRequestBuilder)(nil)

func (m *_CIPEncapsulationReadRequestBuilder) WithMandatoryFields(request DF1RequestMessage) CIPEncapsulationReadRequestBuilder {
	return m.WithRequest(request)
}

func (m *_CIPEncapsulationReadRequestBuilder) WithRequest(request DF1RequestMessage) CIPEncapsulationReadRequestBuilder {
	m.Request = request
	return m
}

func (m *_CIPEncapsulationReadRequestBuilder) Build() (CIPEncapsulationReadRequest, error) {
	if m.Request == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'request' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._CIPEncapsulationReadRequest.deepCopy(), nil
}

func (m *_CIPEncapsulationReadRequestBuilder) MustBuild() CIPEncapsulationReadRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_CIPEncapsulationReadRequestBuilder) DeepCopy() any {
	return m.CreateCIPEncapsulationReadRequestBuilder()
}

// CreateCIPEncapsulationReadRequestBuilder creates a CIPEncapsulationReadRequestBuilder
func (m *_CIPEncapsulationReadRequest) CreateCIPEncapsulationReadRequestBuilder() CIPEncapsulationReadRequestBuilder {
	if m == nil {
		return NewCIPEncapsulationReadRequestBuilder()
	}
	return &_CIPEncapsulationReadRequestBuilder{_CIPEncapsulationReadRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationReadRequest) GetCommandType() uint16 {
	return 0x0107
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationReadRequest) GetParent() CIPEncapsulationPacketContract {
	return m.CIPEncapsulationPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationReadRequest) GetRequest() DF1RequestMessage {
	return m.Request
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationReadRequest(structType any) CIPEncapsulationReadRequest {
	if casted, ok := structType.(CIPEncapsulationReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationReadRequest) GetTypeName() string {
	return "CIPEncapsulationReadRequest"
}

func (m *_CIPEncapsulationReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).getLengthInBits(ctx))

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CIPEncapsulationReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CIPEncapsulationReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CIPEncapsulationPacket) (__cIPEncapsulationReadRequest CIPEncapsulationReadRequest, err error) {
	m.CIPEncapsulationPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	request, err := ReadSimpleField[DF1RequestMessage](ctx, "request", ReadComplex[DF1RequestMessage](DF1RequestMessageParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'request' field"))
	}
	m.Request = request

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationReadRequest")
	}

	return m, nil
}

func (m *_CIPEncapsulationReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPEncapsulationReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationReadRequest")
		}

		if err := WriteSimpleField[DF1RequestMessage](ctx, "request", m.GetRequest(), WriteComplex[DF1RequestMessage](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'request' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationReadRequest")
		}
		return nil
	}
	return m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CIPEncapsulationReadRequest) IsCIPEncapsulationReadRequest() {}

func (m *_CIPEncapsulationReadRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CIPEncapsulationReadRequest) deepCopy() *_CIPEncapsulationReadRequest {
	if m == nil {
		return nil
	}
	_CIPEncapsulationReadRequestCopy := &_CIPEncapsulationReadRequest{
		m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).deepCopy(),
		m.Request.DeepCopy().(DF1RequestMessage),
	}
	m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket)._SubType = m
	return _CIPEncapsulationReadRequestCopy
}

func (m *_CIPEncapsulationReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
