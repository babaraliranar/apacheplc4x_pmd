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

// OpcuaOpenRequest is the corresponding interface of OpcuaOpenRequest
type OpcuaOpenRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MessagePDU
	// GetOpenRequest returns OpenRequest (property field)
	GetOpenRequest() OpenChannelMessage
	// GetMessage returns Message (property field)
	GetMessage() Payload
	// IsOpcuaOpenRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaOpenRequest()
	// CreateBuilder creates a OpcuaOpenRequestBuilder
	CreateOpcuaOpenRequestBuilder() OpcuaOpenRequestBuilder
}

// _OpcuaOpenRequest is the data-structure of this message
type _OpcuaOpenRequest struct {
	MessagePDUContract
	OpenRequest OpenChannelMessage
	Message     Payload

	// Arguments.
	TotalLength uint32
}

var _ OpcuaOpenRequest = (*_OpcuaOpenRequest)(nil)
var _ MessagePDURequirements = (*_OpcuaOpenRequest)(nil)

// NewOpcuaOpenRequest factory function for _OpcuaOpenRequest
func NewOpcuaOpenRequest(chunk ChunkType, openRequest OpenChannelMessage, message Payload, totalLength uint32, binary bool) *_OpcuaOpenRequest {
	if openRequest == nil {
		panic("openRequest of type OpenChannelMessage for OpcuaOpenRequest must not be nil")
	}
	if message == nil {
		panic("message of type Payload for OpcuaOpenRequest must not be nil")
	}
	_result := &_OpcuaOpenRequest{
		MessagePDUContract: NewMessagePDU(chunk, binary),
		OpenRequest:        openRequest,
		Message:            message,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaOpenRequestBuilder is a builder for OpcuaOpenRequest
type OpcuaOpenRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openRequest OpenChannelMessage, message Payload) OpcuaOpenRequestBuilder
	// WithOpenRequest adds OpenRequest (property field)
	WithOpenRequest(OpenChannelMessage) OpcuaOpenRequestBuilder
	// WithOpenRequestBuilder adds OpenRequest (property field) which is build by the builder
	WithOpenRequestBuilder(func(OpenChannelMessageBuilder) OpenChannelMessageBuilder) OpcuaOpenRequestBuilder
	// WithMessage adds Message (property field)
	WithMessage(Payload) OpcuaOpenRequestBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(PayloadBuilder) PayloadBuilder) OpcuaOpenRequestBuilder
	// WithArgTotalLength sets a parser argument
	WithArgTotalLength(uint32) OpcuaOpenRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MessagePDUBuilder
	// Build builds the OpcuaOpenRequest or returns an error if something is wrong
	Build() (OpcuaOpenRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaOpenRequest
}

// NewOpcuaOpenRequestBuilder() creates a OpcuaOpenRequestBuilder
func NewOpcuaOpenRequestBuilder() OpcuaOpenRequestBuilder {
	return &_OpcuaOpenRequestBuilder{_OpcuaOpenRequest: new(_OpcuaOpenRequest)}
}

type _OpcuaOpenRequestBuilder struct {
	*_OpcuaOpenRequest

	parentBuilder *_MessagePDUBuilder

	err *utils.MultiError
}

var _ (OpcuaOpenRequestBuilder) = (*_OpcuaOpenRequestBuilder)(nil)

func (b *_OpcuaOpenRequestBuilder) setParent(contract MessagePDUContract) {
	b.MessagePDUContract = contract
	contract.(*_MessagePDU)._SubType = b._OpcuaOpenRequest
}

func (b *_OpcuaOpenRequestBuilder) WithMandatoryFields(openRequest OpenChannelMessage, message Payload) OpcuaOpenRequestBuilder {
	return b.WithOpenRequest(openRequest).WithMessage(message)
}

func (b *_OpcuaOpenRequestBuilder) WithOpenRequest(openRequest OpenChannelMessage) OpcuaOpenRequestBuilder {
	b.OpenRequest = openRequest
	return b
}

func (b *_OpcuaOpenRequestBuilder) WithOpenRequestBuilder(builderSupplier func(OpenChannelMessageBuilder) OpenChannelMessageBuilder) OpcuaOpenRequestBuilder {
	builder := builderSupplier(b.OpenRequest.CreateOpenChannelMessageBuilder())
	var err error
	b.OpenRequest, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OpenChannelMessageBuilder failed"))
	}
	return b
}

func (b *_OpcuaOpenRequestBuilder) WithMessage(message Payload) OpcuaOpenRequestBuilder {
	b.Message = message
	return b
}

func (b *_OpcuaOpenRequestBuilder) WithMessageBuilder(builderSupplier func(PayloadBuilder) PayloadBuilder) OpcuaOpenRequestBuilder {
	builder := builderSupplier(b.Message.CreatePayloadBuilder())
	var err error
	b.Message, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PayloadBuilder failed"))
	}
	return b
}

func (b *_OpcuaOpenRequestBuilder) WithArgTotalLength(totalLength uint32) OpcuaOpenRequestBuilder {
	b.TotalLength = totalLength
	return b
}

func (b *_OpcuaOpenRequestBuilder) Build() (OpcuaOpenRequest, error) {
	if b.OpenRequest == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openRequest' not set"))
	}
	if b.Message == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpcuaOpenRequest.deepCopy(), nil
}

func (b *_OpcuaOpenRequestBuilder) MustBuild() OpcuaOpenRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_OpcuaOpenRequestBuilder) Done() MessagePDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMessagePDUBuilder().(*_MessagePDUBuilder)
	}
	return b.parentBuilder
}

func (b *_OpcuaOpenRequestBuilder) buildForMessagePDU() (MessagePDU, error) {
	return b.Build()
}

func (b *_OpcuaOpenRequestBuilder) DeepCopy() any {
	_copy := b.CreateOpcuaOpenRequestBuilder().(*_OpcuaOpenRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpcuaOpenRequestBuilder creates a OpcuaOpenRequestBuilder
func (b *_OpcuaOpenRequest) CreateOpcuaOpenRequestBuilder() OpcuaOpenRequestBuilder {
	if b == nil {
		return NewOpcuaOpenRequestBuilder()
	}
	return &_OpcuaOpenRequestBuilder{_OpcuaOpenRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaOpenRequest) GetMessageType() string {
	return "OPN"
}

func (m *_OpcuaOpenRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaOpenRequest) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaOpenRequest) GetOpenRequest() OpenChannelMessage {
	return m.OpenRequest
}

func (m *_OpcuaOpenRequest) GetMessage() Payload {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaOpenRequest(structType any) OpcuaOpenRequest {
	if casted, ok := structType.(OpcuaOpenRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaOpenRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaOpenRequest) GetTypeName() string {
	return "OpcuaOpenRequest"
}

func (m *_OpcuaOpenRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (openRequest)
	lengthInBits += m.OpenRequest.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaOpenRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaOpenRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, totalLength uint32, response bool, binary bool) (__opcuaOpenRequest OpcuaOpenRequest, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaOpenRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaOpenRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openRequest, err := ReadSimpleField[OpenChannelMessage](ctx, "openRequest", ReadComplex[OpenChannelMessage](OpenChannelMessageParseWithBufferProducer[OpenChannelMessage]((bool)(response)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openRequest' field"))
	}
	m.OpenRequest = openRequest

	message, err := ReadSimpleField[Payload](ctx, "message", ReadComplex[Payload](PayloadParseWithBufferProducer[Payload]((bool)(binary), (uint32)(uint32(uint32(totalLength)-uint32(openRequest.GetLengthInBytes(ctx)))-uint32(uint32(16)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaOpenRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaOpenRequest")
	}

	return m, nil
}

func (m *_OpcuaOpenRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaOpenRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaOpenRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaOpenRequest")
		}

		if err := WriteSimpleField[OpenChannelMessage](ctx, "openRequest", m.GetOpenRequest(), WriteComplex[OpenChannelMessage](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openRequest' field")
		}

		if err := WriteSimpleField[Payload](ctx, "message", m.GetMessage(), WriteComplex[Payload](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaOpenRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaOpenRequest")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_OpcuaOpenRequest) GetTotalLength() uint32 {
	return m.TotalLength
}

//
////

func (m *_OpcuaOpenRequest) IsOpcuaOpenRequest() {}

func (m *_OpcuaOpenRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaOpenRequest) deepCopy() *_OpcuaOpenRequest {
	if m == nil {
		return nil
	}
	_OpcuaOpenRequestCopy := &_OpcuaOpenRequest{
		m.MessagePDUContract.(*_MessagePDU).deepCopy(),
		utils.DeepCopy[OpenChannelMessage](m.OpenRequest),
		utils.DeepCopy[Payload](m.Message),
		m.TotalLength,
	}
	_OpcuaOpenRequestCopy.MessagePDUContract.(*_MessagePDU)._SubType = m
	return _OpcuaOpenRequestCopy
}

func (m *_OpcuaOpenRequest) String() string {
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
