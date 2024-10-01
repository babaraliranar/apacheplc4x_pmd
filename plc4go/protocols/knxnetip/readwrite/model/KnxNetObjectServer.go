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

// KnxNetObjectServer is the corresponding interface of KnxNetObjectServer
type KnxNetObjectServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// IsKnxNetObjectServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetObjectServer()
	// CreateBuilder creates a KnxNetObjectServerBuilder
	CreateKnxNetObjectServerBuilder() KnxNetObjectServerBuilder
}

// _KnxNetObjectServer is the data-structure of this message
type _KnxNetObjectServer struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetObjectServer = (*_KnxNetObjectServer)(nil)
var _ ServiceIdRequirements = (*_KnxNetObjectServer)(nil)

// NewKnxNetObjectServer factory function for _KnxNetObjectServer
func NewKnxNetObjectServer(version uint8) *_KnxNetObjectServer {
	_result := &_KnxNetObjectServer{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	_result.ServiceIdContract.(*_ServiceId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// KnxNetObjectServerBuilder is a builder for KnxNetObjectServer
type KnxNetObjectServerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(version uint8) KnxNetObjectServerBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint8) KnxNetObjectServerBuilder
	// Build builds the KnxNetObjectServer or returns an error if something is wrong
	Build() (KnxNetObjectServer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() KnxNetObjectServer
}

// NewKnxNetObjectServerBuilder() creates a KnxNetObjectServerBuilder
func NewKnxNetObjectServerBuilder() KnxNetObjectServerBuilder {
	return &_KnxNetObjectServerBuilder{_KnxNetObjectServer: new(_KnxNetObjectServer)}
}

type _KnxNetObjectServerBuilder struct {
	*_KnxNetObjectServer

	parentBuilder *_ServiceIdBuilder

	err *utils.MultiError
}

var _ (KnxNetObjectServerBuilder) = (*_KnxNetObjectServerBuilder)(nil)

func (b *_KnxNetObjectServerBuilder) setParent(contract ServiceIdContract) {
	b.ServiceIdContract = contract
}

func (b *_KnxNetObjectServerBuilder) WithMandatoryFields(version uint8) KnxNetObjectServerBuilder {
	return b.WithVersion(version)
}

func (b *_KnxNetObjectServerBuilder) WithVersion(version uint8) KnxNetObjectServerBuilder {
	b.Version = version
	return b
}

func (b *_KnxNetObjectServerBuilder) Build() (KnxNetObjectServer, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._KnxNetObjectServer.deepCopy(), nil
}

func (b *_KnxNetObjectServerBuilder) MustBuild() KnxNetObjectServer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_KnxNetObjectServerBuilder) Done() ServiceIdBuilder {
	return b.parentBuilder
}

func (b *_KnxNetObjectServerBuilder) buildForServiceId() (ServiceId, error) {
	return b.Build()
}

func (b *_KnxNetObjectServerBuilder) DeepCopy() any {
	_copy := b.CreateKnxNetObjectServerBuilder().(*_KnxNetObjectServerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateKnxNetObjectServerBuilder creates a KnxNetObjectServerBuilder
func (b *_KnxNetObjectServer) CreateKnxNetObjectServerBuilder() KnxNetObjectServerBuilder {
	if b == nil {
		return NewKnxNetObjectServerBuilder()
	}
	return &_KnxNetObjectServerBuilder{_KnxNetObjectServer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetObjectServer) GetServiceType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetObjectServer) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetObjectServer) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetObjectServer(structType any) KnxNetObjectServer {
	if casted, ok := structType.(KnxNetObjectServer); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetObjectServer); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetObjectServer) GetTypeName() string {
	return "KnxNetObjectServer"
}

func (m *_KnxNetObjectServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetObjectServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetObjectServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetObjectServer KnxNetObjectServer, err error) {
	m.ServiceIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetObjectServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetObjectServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetObjectServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetObjectServer")
	}

	return m, nil
}

func (m *_KnxNetObjectServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetObjectServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetObjectServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetObjectServer")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetObjectServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetObjectServer")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetObjectServer) IsKnxNetObjectServer() {}

func (m *_KnxNetObjectServer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_KnxNetObjectServer) deepCopy() *_KnxNetObjectServer {
	if m == nil {
		return nil
	}
	_KnxNetObjectServerCopy := &_KnxNetObjectServer{
		m.ServiceIdContract.(*_ServiceId).deepCopy(),
		m.Version,
	}
	m.ServiceIdContract.(*_ServiceId)._SubType = m
	return _KnxNetObjectServerCopy
}

func (m *_KnxNetObjectServer) String() string {
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
