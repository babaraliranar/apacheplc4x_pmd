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

// CIPEncapsulationReadResponse is the corresponding interface of CIPEncapsulationReadResponse
type CIPEncapsulationReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CIPEncapsulationPacket
	// GetResponse returns Response (property field)
	GetResponse() DF1ResponseMessage
	// IsCIPEncapsulationReadResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCIPEncapsulationReadResponse()
	// CreateBuilder creates a CIPEncapsulationReadResponseBuilder
	CreateCIPEncapsulationReadResponseBuilder() CIPEncapsulationReadResponseBuilder
}

// _CIPEncapsulationReadResponse is the data-structure of this message
type _CIPEncapsulationReadResponse struct {
	CIPEncapsulationPacketContract
	Response DF1ResponseMessage

	// Arguments.
	PacketLen uint16
}

var _ CIPEncapsulationReadResponse = (*_CIPEncapsulationReadResponse)(nil)
var _ CIPEncapsulationPacketRequirements = (*_CIPEncapsulationReadResponse)(nil)

// NewCIPEncapsulationReadResponse factory function for _CIPEncapsulationReadResponse
func NewCIPEncapsulationReadResponse(sessionHandle uint32, status uint32, senderContext []uint8, options uint32, response DF1ResponseMessage, packetLen uint16) *_CIPEncapsulationReadResponse {
	if response == nil {
		panic("response of type DF1ResponseMessage for CIPEncapsulationReadResponse must not be nil")
	}
	_result := &_CIPEncapsulationReadResponse{
		CIPEncapsulationPacketContract: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
		Response:                       response,
	}
	_result.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CIPEncapsulationReadResponseBuilder is a builder for CIPEncapsulationReadResponse
type CIPEncapsulationReadResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(response DF1ResponseMessage) CIPEncapsulationReadResponseBuilder
	// WithResponse adds Response (property field)
	WithResponse(DF1ResponseMessage) CIPEncapsulationReadResponseBuilder
	// WithResponseBuilder adds Response (property field) which is build by the builder
	WithResponseBuilder(func(DF1ResponseMessageBuilder) DF1ResponseMessageBuilder) CIPEncapsulationReadResponseBuilder
	// WithArgPacketLen sets a parser argument
	WithArgPacketLen(uint16) CIPEncapsulationReadResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CIPEncapsulationPacketBuilder
	// Build builds the CIPEncapsulationReadResponse or returns an error if something is wrong
	Build() (CIPEncapsulationReadResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CIPEncapsulationReadResponse
}

// NewCIPEncapsulationReadResponseBuilder() creates a CIPEncapsulationReadResponseBuilder
func NewCIPEncapsulationReadResponseBuilder() CIPEncapsulationReadResponseBuilder {
	return &_CIPEncapsulationReadResponseBuilder{_CIPEncapsulationReadResponse: new(_CIPEncapsulationReadResponse)}
}

type _CIPEncapsulationReadResponseBuilder struct {
	*_CIPEncapsulationReadResponse

	parentBuilder *_CIPEncapsulationPacketBuilder

	err *utils.MultiError
}

var _ (CIPEncapsulationReadResponseBuilder) = (*_CIPEncapsulationReadResponseBuilder)(nil)

func (b *_CIPEncapsulationReadResponseBuilder) setParent(contract CIPEncapsulationPacketContract) {
	b.CIPEncapsulationPacketContract = contract
	contract.(*_CIPEncapsulationPacket)._SubType = b._CIPEncapsulationReadResponse
}

func (b *_CIPEncapsulationReadResponseBuilder) WithMandatoryFields(response DF1ResponseMessage) CIPEncapsulationReadResponseBuilder {
	return b.WithResponse(response)
}

func (b *_CIPEncapsulationReadResponseBuilder) WithResponse(response DF1ResponseMessage) CIPEncapsulationReadResponseBuilder {
	b.Response = response
	return b
}

func (b *_CIPEncapsulationReadResponseBuilder) WithResponseBuilder(builderSupplier func(DF1ResponseMessageBuilder) DF1ResponseMessageBuilder) CIPEncapsulationReadResponseBuilder {
	builder := builderSupplier(b.Response.CreateDF1ResponseMessageBuilder())
	var err error
	b.Response, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DF1ResponseMessageBuilder failed"))
	}
	return b
}

func (b *_CIPEncapsulationReadResponseBuilder) WithArgPacketLen(packetLen uint16) CIPEncapsulationReadResponseBuilder {
	b.PacketLen = packetLen
	return b
}

func (b *_CIPEncapsulationReadResponseBuilder) Build() (CIPEncapsulationReadResponse, error) {
	if b.Response == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'response' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CIPEncapsulationReadResponse.deepCopy(), nil
}

func (b *_CIPEncapsulationReadResponseBuilder) MustBuild() CIPEncapsulationReadResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CIPEncapsulationReadResponseBuilder) Done() CIPEncapsulationPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCIPEncapsulationPacketBuilder().(*_CIPEncapsulationPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_CIPEncapsulationReadResponseBuilder) buildForCIPEncapsulationPacket() (CIPEncapsulationPacket, error) {
	return b.Build()
}

func (b *_CIPEncapsulationReadResponseBuilder) DeepCopy() any {
	_copy := b.CreateCIPEncapsulationReadResponseBuilder().(*_CIPEncapsulationReadResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCIPEncapsulationReadResponseBuilder creates a CIPEncapsulationReadResponseBuilder
func (b *_CIPEncapsulationReadResponse) CreateCIPEncapsulationReadResponseBuilder() CIPEncapsulationReadResponseBuilder {
	if b == nil {
		return NewCIPEncapsulationReadResponseBuilder()
	}
	return &_CIPEncapsulationReadResponseBuilder{_CIPEncapsulationReadResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationReadResponse) GetCommandType() uint16 {
	return 0x0207
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationReadResponse) GetParent() CIPEncapsulationPacketContract {
	return m.CIPEncapsulationPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationReadResponse) GetResponse() DF1ResponseMessage {
	return m.Response
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationReadResponse(structType any) CIPEncapsulationReadResponse {
	if casted, ok := structType.(CIPEncapsulationReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationReadResponse) GetTypeName() string {
	return "CIPEncapsulationReadResponse"
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).getLengthInBits(ctx))

	// Simple field (response)
	lengthInBits += m.Response.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CIPEncapsulationReadResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CIPEncapsulationPacket, packetLen uint16) (__cIPEncapsulationReadResponse CIPEncapsulationReadResponse, err error) {
	m.CIPEncapsulationPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	response, err := ReadSimpleField[DF1ResponseMessage](ctx, "response", ReadComplex[DF1ResponseMessage](DF1ResponseMessageParseWithBufferProducer[DF1ResponseMessage]((uint16)(packetLen)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'response' field"))
	}
	m.Response = response

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationReadResponse")
	}

	return m, nil
}

func (m *_CIPEncapsulationReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPEncapsulationReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationReadResponse")
		}

		if err := WriteSimpleField[DF1ResponseMessage](ctx, "response", m.GetResponse(), WriteComplex[DF1ResponseMessage](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'response' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationReadResponse")
		}
		return nil
	}
	return m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_CIPEncapsulationReadResponse) GetPacketLen() uint16 {
	return m.PacketLen
}

//
////

func (m *_CIPEncapsulationReadResponse) IsCIPEncapsulationReadResponse() {}

func (m *_CIPEncapsulationReadResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CIPEncapsulationReadResponse) deepCopy() *_CIPEncapsulationReadResponse {
	if m == nil {
		return nil
	}
	_CIPEncapsulationReadResponseCopy := &_CIPEncapsulationReadResponse{
		m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).deepCopy(),
		utils.DeepCopy[DF1ResponseMessage](m.Response),
		m.PacketLen,
	}
	_CIPEncapsulationReadResponseCopy.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket)._SubType = m
	return _CIPEncapsulationReadResponseCopy
}

func (m *_CIPEncapsulationReadResponse) String() string {
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
