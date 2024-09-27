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

// COTPPacketDisconnectResponse is the corresponding interface of COTPPacketDisconnectResponse
type COTPPacketDisconnectResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// IsCOTPPacketDisconnectResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPPacketDisconnectResponse()
	// CreateBuilder creates a COTPPacketDisconnectResponseBuilder
	CreateCOTPPacketDisconnectResponseBuilder() COTPPacketDisconnectResponseBuilder
}

// _COTPPacketDisconnectResponse is the data-structure of this message
type _COTPPacketDisconnectResponse struct {
	COTPPacketContract
	DestinationReference uint16
	SourceReference      uint16
}

var _ COTPPacketDisconnectResponse = (*_COTPPacketDisconnectResponse)(nil)
var _ COTPPacketRequirements = (*_COTPPacketDisconnectResponse)(nil)

// NewCOTPPacketDisconnectResponse factory function for _COTPPacketDisconnectResponse
func NewCOTPPacketDisconnectResponse(parameters []COTPParameter, payload S7Message, destinationReference uint16, sourceReference uint16, cotpLen uint16) *_COTPPacketDisconnectResponse {
	_result := &_COTPPacketDisconnectResponse{
		COTPPacketContract:   NewCOTPPacket(parameters, payload, cotpLen),
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
	}
	_result.COTPPacketContract.(*_COTPPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPPacketDisconnectResponseBuilder is a builder for COTPPacketDisconnectResponse
type COTPPacketDisconnectResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationReference uint16, sourceReference uint16) COTPPacketDisconnectResponseBuilder
	// WithDestinationReference adds DestinationReference (property field)
	WithDestinationReference(uint16) COTPPacketDisconnectResponseBuilder
	// WithSourceReference adds SourceReference (property field)
	WithSourceReference(uint16) COTPPacketDisconnectResponseBuilder
	// Build builds the COTPPacketDisconnectResponse or returns an error if something is wrong
	Build() (COTPPacketDisconnectResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPPacketDisconnectResponse
}

// NewCOTPPacketDisconnectResponseBuilder() creates a COTPPacketDisconnectResponseBuilder
func NewCOTPPacketDisconnectResponseBuilder() COTPPacketDisconnectResponseBuilder {
	return &_COTPPacketDisconnectResponseBuilder{_COTPPacketDisconnectResponse: new(_COTPPacketDisconnectResponse)}
}

type _COTPPacketDisconnectResponseBuilder struct {
	*_COTPPacketDisconnectResponse

	err *utils.MultiError
}

var _ (COTPPacketDisconnectResponseBuilder) = (*_COTPPacketDisconnectResponseBuilder)(nil)

func (m *_COTPPacketDisconnectResponseBuilder) WithMandatoryFields(destinationReference uint16, sourceReference uint16) COTPPacketDisconnectResponseBuilder {
	return m.WithDestinationReference(destinationReference).WithSourceReference(sourceReference)
}

func (m *_COTPPacketDisconnectResponseBuilder) WithDestinationReference(destinationReference uint16) COTPPacketDisconnectResponseBuilder {
	m.DestinationReference = destinationReference
	return m
}

func (m *_COTPPacketDisconnectResponseBuilder) WithSourceReference(sourceReference uint16) COTPPacketDisconnectResponseBuilder {
	m.SourceReference = sourceReference
	return m
}

func (m *_COTPPacketDisconnectResponseBuilder) Build() (COTPPacketDisconnectResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._COTPPacketDisconnectResponse.deepCopy(), nil
}

func (m *_COTPPacketDisconnectResponseBuilder) MustBuild() COTPPacketDisconnectResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_COTPPacketDisconnectResponseBuilder) DeepCopy() any {
	return m.CreateCOTPPacketDisconnectResponseBuilder()
}

// CreateCOTPPacketDisconnectResponseBuilder creates a COTPPacketDisconnectResponseBuilder
func (m *_COTPPacketDisconnectResponse) CreateCOTPPacketDisconnectResponseBuilder() COTPPacketDisconnectResponseBuilder {
	if m == nil {
		return NewCOTPPacketDisconnectResponseBuilder()
	}
	return &_COTPPacketDisconnectResponseBuilder{_COTPPacketDisconnectResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketDisconnectResponse) GetTpduCode() uint8 {
	return 0xC0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketDisconnectResponse) GetParent() COTPPacketContract {
	return m.COTPPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketDisconnectResponse) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketDisconnectResponse) GetSourceReference() uint16 {
	return m.SourceReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPPacketDisconnectResponse(structType any) COTPPacketDisconnectResponse {
	if casted, ok := structType.(COTPPacketDisconnectResponse); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketDisconnectResponse); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketDisconnectResponse) GetTypeName() string {
	return "COTPPacketDisconnectResponse"
}

func (m *_COTPPacketDisconnectResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPPacketContract.(*_COTPPacket).getLengthInBits(ctx))

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	return lengthInBits
}

func (m *_COTPPacketDisconnectResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPPacketDisconnectResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPPacket, cotpLen uint16) (__cOTPPacketDisconnectResponse COTPPacketDisconnectResponse, err error) {
	m.COTPPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketDisconnectResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketDisconnectResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationReference, err := ReadSimpleField(ctx, "destinationReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationReference' field"))
	}
	m.DestinationReference = destinationReference

	sourceReference, err := ReadSimpleField(ctx, "sourceReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceReference' field"))
	}
	m.SourceReference = sourceReference

	if closeErr := readBuffer.CloseContext("COTPPacketDisconnectResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketDisconnectResponse")
	}

	return m, nil
}

func (m *_COTPPacketDisconnectResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketDisconnectResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketDisconnectResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketDisconnectResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationReference", m.GetDestinationReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationReference' field")
		}

		if err := WriteSimpleField[uint16](ctx, "sourceReference", m.GetSourceReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceReference' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketDisconnectResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketDisconnectResponse")
		}
		return nil
	}
	return m.COTPPacketContract.(*_COTPPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketDisconnectResponse) IsCOTPPacketDisconnectResponse() {}

func (m *_COTPPacketDisconnectResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPPacketDisconnectResponse) deepCopy() *_COTPPacketDisconnectResponse {
	if m == nil {
		return nil
	}
	_COTPPacketDisconnectResponseCopy := &_COTPPacketDisconnectResponse{
		m.COTPPacketContract.(*_COTPPacket).deepCopy(),
		m.DestinationReference,
		m.SourceReference,
	}
	m.COTPPacketContract.(*_COTPPacket)._SubType = m
	return _COTPPacketDisconnectResponseCopy
}

func (m *_COTPPacketDisconnectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
