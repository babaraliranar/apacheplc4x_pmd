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

// BinaryPayload is the corresponding interface of BinaryPayload
type BinaryPayload interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Payload
	// GetPayload returns Payload (property field)
	GetPayload() []byte
	// IsBinaryPayload is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBinaryPayload()
	// CreateBuilder creates a BinaryPayloadBuilder
	CreateBinaryPayloadBuilder() BinaryPayloadBuilder
}

// _BinaryPayload is the data-structure of this message
type _BinaryPayload struct {
	PayloadContract
	Payload []byte
}

var _ BinaryPayload = (*_BinaryPayload)(nil)
var _ PayloadRequirements = (*_BinaryPayload)(nil)

// NewBinaryPayload factory function for _BinaryPayload
func NewBinaryPayload(sequenceHeader SequenceHeader, payload []byte, byteCount uint32) *_BinaryPayload {
	_result := &_BinaryPayload{
		PayloadContract: NewPayload(sequenceHeader, byteCount),
		Payload:         payload,
	}
	_result.PayloadContract.(*_Payload)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BinaryPayloadBuilder is a builder for BinaryPayload
type BinaryPayloadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload []byte) BinaryPayloadBuilder
	// WithPayload adds Payload (property field)
	WithPayload(...byte) BinaryPayloadBuilder
	// Build builds the BinaryPayload or returns an error if something is wrong
	Build() (BinaryPayload, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BinaryPayload
}

// NewBinaryPayloadBuilder() creates a BinaryPayloadBuilder
func NewBinaryPayloadBuilder() BinaryPayloadBuilder {
	return &_BinaryPayloadBuilder{_BinaryPayload: new(_BinaryPayload)}
}

type _BinaryPayloadBuilder struct {
	*_BinaryPayload

	err *utils.MultiError
}

var _ (BinaryPayloadBuilder) = (*_BinaryPayloadBuilder)(nil)

func (m *_BinaryPayloadBuilder) WithMandatoryFields(payload []byte) BinaryPayloadBuilder {
	return m.WithPayload(payload...)
}

func (m *_BinaryPayloadBuilder) WithPayload(payload ...byte) BinaryPayloadBuilder {
	m.Payload = payload
	return m
}

func (m *_BinaryPayloadBuilder) Build() (BinaryPayload, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BinaryPayload.deepCopy(), nil
}

func (m *_BinaryPayloadBuilder) MustBuild() BinaryPayload {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BinaryPayloadBuilder) DeepCopy() any {
	return m.CreateBinaryPayloadBuilder()
}

// CreateBinaryPayloadBuilder creates a BinaryPayloadBuilder
func (m *_BinaryPayload) CreateBinaryPayloadBuilder() BinaryPayloadBuilder {
	if m == nil {
		return NewBinaryPayloadBuilder()
	}
	return &_BinaryPayloadBuilder{_BinaryPayload: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BinaryPayload) GetExtensible() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BinaryPayload) GetParent() PayloadContract {
	return m.PayloadContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BinaryPayload) GetPayload() []byte {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBinaryPayload(structType any) BinaryPayload {
	if casted, ok := structType.(BinaryPayload); ok {
		return casted
	}
	if casted, ok := structType.(*BinaryPayload); ok {
		return *casted
	}
	return nil
}

func (m *_BinaryPayload) GetTypeName() string {
	return "BinaryPayload"
}

func (m *_BinaryPayload) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PayloadContract.(*_Payload).getLengthInBits(ctx))

	// Array field
	if len(m.Payload) > 0 {
		lengthInBits += 8 * uint16(len(m.Payload))
	}

	return lengthInBits
}

func (m *_BinaryPayload) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BinaryPayload) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Payload, extensible bool, byteCount uint32) (__binaryPayload BinaryPayload, err error) {
	m.PayloadContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BinaryPayload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BinaryPayload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := readBuffer.ReadByteArray("payload", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BinaryPayload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BinaryPayload")
	}

	return m, nil
}

func (m *_BinaryPayload) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BinaryPayload) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BinaryPayload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BinaryPayload")
		}

		if err := WriteByteArrayField(ctx, "payload", m.GetPayload(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BinaryPayload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BinaryPayload")
		}
		return nil
	}
	return m.PayloadContract.(*_Payload).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BinaryPayload) IsBinaryPayload() {}

func (m *_BinaryPayload) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BinaryPayload) deepCopy() *_BinaryPayload {
	if m == nil {
		return nil
	}
	_BinaryPayloadCopy := &_BinaryPayload{
		m.PayloadContract.(*_Payload).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Payload),
	}
	m.PayloadContract.(*_Payload)._SubType = m
	return _BinaryPayloadCopy
}

func (m *_BinaryPayload) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
