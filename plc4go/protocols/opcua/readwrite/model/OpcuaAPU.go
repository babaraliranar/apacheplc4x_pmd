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

// OpcuaAPU is the corresponding interface of OpcuaAPU
type OpcuaAPU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetMessage returns Message (property field)
	GetMessage() MessagePDU
	// IsOpcuaAPU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaAPU()
	// CreateBuilder creates a OpcuaAPUBuilder
	CreateOpcuaAPUBuilder() OpcuaAPUBuilder
}

// _OpcuaAPU is the data-structure of this message
type _OpcuaAPU struct {
	Message MessagePDU

	// Arguments.
	Response       bool
	BinaryEncoding bool
}

var _ OpcuaAPU = (*_OpcuaAPU)(nil)

// NewOpcuaAPU factory function for _OpcuaAPU
func NewOpcuaAPU(message MessagePDU, response bool, binaryEncoding bool) *_OpcuaAPU {
	if message == nil {
		panic("message of type MessagePDU for OpcuaAPU must not be nil")
	}
	return &_OpcuaAPU{Message: message, Response: response, BinaryEncoding: binaryEncoding}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaAPUBuilder is a builder for OpcuaAPU
type OpcuaAPUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(message MessagePDU) OpcuaAPUBuilder
	// WithMessage adds Message (property field)
	WithMessage(MessagePDU) OpcuaAPUBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(MessagePDUBuilder) MessagePDUBuilder) OpcuaAPUBuilder
	// WithArgResponse sets a parser argument
	WithArgResponse(bool) OpcuaAPUBuilder
	// WithArgBinaryEncoding sets a parser argument
	WithArgBinaryEncoding(bool) OpcuaAPUBuilder
	// Build builds the OpcuaAPU or returns an error if something is wrong
	Build() (OpcuaAPU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaAPU
}

// NewOpcuaAPUBuilder() creates a OpcuaAPUBuilder
func NewOpcuaAPUBuilder() OpcuaAPUBuilder {
	return &_OpcuaAPUBuilder{_OpcuaAPU: new(_OpcuaAPU)}
}

type _OpcuaAPUBuilder struct {
	*_OpcuaAPU

	err *utils.MultiError
}

var _ (OpcuaAPUBuilder) = (*_OpcuaAPUBuilder)(nil)

func (b *_OpcuaAPUBuilder) WithMandatoryFields(message MessagePDU) OpcuaAPUBuilder {
	return b.WithMessage(message)
}

func (b *_OpcuaAPUBuilder) WithMessage(message MessagePDU) OpcuaAPUBuilder {
	b.Message = message
	return b
}

func (b *_OpcuaAPUBuilder) WithMessageBuilder(builderSupplier func(MessagePDUBuilder) MessagePDUBuilder) OpcuaAPUBuilder {
	builder := builderSupplier(b.Message.CreateMessagePDUBuilder())
	var err error
	b.Message, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "MessagePDUBuilder failed"))
	}
	return b
}

func (b *_OpcuaAPUBuilder) WithArgResponse(response bool) OpcuaAPUBuilder {
	b.Response = response
	return b
}
func (b *_OpcuaAPUBuilder) WithArgBinaryEncoding(binaryEncoding bool) OpcuaAPUBuilder {
	b.BinaryEncoding = binaryEncoding
	return b
}

func (b *_OpcuaAPUBuilder) Build() (OpcuaAPU, error) {
	if b.Message == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpcuaAPU.deepCopy(), nil
}

func (b *_OpcuaAPUBuilder) MustBuild() OpcuaAPU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_OpcuaAPUBuilder) DeepCopy() any {
	_copy := b.CreateOpcuaAPUBuilder().(*_OpcuaAPUBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpcuaAPUBuilder creates a OpcuaAPUBuilder
func (b *_OpcuaAPU) CreateOpcuaAPUBuilder() OpcuaAPUBuilder {
	if b == nil {
		return NewOpcuaAPUBuilder()
	}
	return &_OpcuaAPUBuilder{_OpcuaAPU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaAPU) GetMessage() MessagePDU {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaAPU(structType any) OpcuaAPU {
	if casted, ok := structType.(OpcuaAPU); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaAPU); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaAPU) GetTypeName() string {
	return "OpcuaAPU"
}

func (m *_OpcuaAPU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaAPU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaAPUParse(ctx context.Context, theBytes []byte, response bool, binaryEncoding bool) (OpcuaAPU, error) {
	return OpcuaAPUParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)), response, binaryEncoding)
}

func OpcuaAPUParseWithBufferProducer(response bool, binaryEncoding bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaAPU, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaAPU, error) {
		return OpcuaAPUParseWithBuffer(ctx, readBuffer, response, binaryEncoding)
	}
}

func OpcuaAPUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool, binaryEncoding bool) (OpcuaAPU, error) {
	v, err := (&_OpcuaAPU{Response: response, BinaryEncoding: binaryEncoding}).parse(ctx, readBuffer, response, binaryEncoding)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_OpcuaAPU) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool, binaryEncoding bool) (__opcuaAPU OpcuaAPU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaAPU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaAPU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	message, err := ReadSimpleField[MessagePDU](ctx, "message", ReadComplex[MessagePDU](MessagePDUParseWithBufferProducer[MessagePDU]((bool)(response), (bool)(binaryEncoding)), readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaAPU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaAPU")
	}

	return m, nil
}

func (m *_OpcuaAPU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaAPU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("OpcuaAPU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for OpcuaAPU")
	}

	if err := WriteSimpleField[MessagePDU](ctx, "message", m.GetMessage(), WriteComplex[MessagePDU](writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'message' field")
	}

	if popErr := writeBuffer.PopContext("OpcuaAPU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for OpcuaAPU")
	}
	return nil
}

////
// Arguments Getter

func (m *_OpcuaAPU) GetResponse() bool {
	return m.Response
}
func (m *_OpcuaAPU) GetBinaryEncoding() bool {
	return m.BinaryEncoding
}

//
////

func (m *_OpcuaAPU) IsOpcuaAPU() {}

func (m *_OpcuaAPU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaAPU) deepCopy() *_OpcuaAPU {
	if m == nil {
		return nil
	}
	_OpcuaAPUCopy := &_OpcuaAPU{
		utils.DeepCopy[MessagePDU](m.Message),
		m.Response,
		m.BinaryEncoding,
	}
	return _OpcuaAPUCopy
}

func (m *_OpcuaAPU) String() string {
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
