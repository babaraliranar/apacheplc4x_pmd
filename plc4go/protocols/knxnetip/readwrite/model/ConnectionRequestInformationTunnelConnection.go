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

// ConnectionRequestInformationTunnelConnection is the corresponding interface of ConnectionRequestInformationTunnelConnection
type ConnectionRequestInformationTunnelConnection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ConnectionRequestInformation
	// GetKnxLayer returns KnxLayer (property field)
	GetKnxLayer() KnxLayer
	// IsConnectionRequestInformationTunnelConnection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionRequestInformationTunnelConnection()
	// CreateBuilder creates a ConnectionRequestInformationTunnelConnectionBuilder
	CreateConnectionRequestInformationTunnelConnectionBuilder() ConnectionRequestInformationTunnelConnectionBuilder
}

// _ConnectionRequestInformationTunnelConnection is the data-structure of this message
type _ConnectionRequestInformationTunnelConnection struct {
	ConnectionRequestInformationContract
	KnxLayer KnxLayer
	// Reserved Fields
	reservedField0 *uint8
}

var _ ConnectionRequestInformationTunnelConnection = (*_ConnectionRequestInformationTunnelConnection)(nil)
var _ ConnectionRequestInformationRequirements = (*_ConnectionRequestInformationTunnelConnection)(nil)

// NewConnectionRequestInformationTunnelConnection factory function for _ConnectionRequestInformationTunnelConnection
func NewConnectionRequestInformationTunnelConnection(knxLayer KnxLayer) *_ConnectionRequestInformationTunnelConnection {
	_result := &_ConnectionRequestInformationTunnelConnection{
		ConnectionRequestInformationContract: NewConnectionRequestInformation(),
		KnxLayer:                             knxLayer,
	}
	_result.ConnectionRequestInformationContract.(*_ConnectionRequestInformation)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionRequestInformationTunnelConnectionBuilder is a builder for ConnectionRequestInformationTunnelConnection
type ConnectionRequestInformationTunnelConnectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(knxLayer KnxLayer) ConnectionRequestInformationTunnelConnectionBuilder
	// WithKnxLayer adds KnxLayer (property field)
	WithKnxLayer(KnxLayer) ConnectionRequestInformationTunnelConnectionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ConnectionRequestInformationBuilder
	// Build builds the ConnectionRequestInformationTunnelConnection or returns an error if something is wrong
	Build() (ConnectionRequestInformationTunnelConnection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionRequestInformationTunnelConnection
}

// NewConnectionRequestInformationTunnelConnectionBuilder() creates a ConnectionRequestInformationTunnelConnectionBuilder
func NewConnectionRequestInformationTunnelConnectionBuilder() ConnectionRequestInformationTunnelConnectionBuilder {
	return &_ConnectionRequestInformationTunnelConnectionBuilder{_ConnectionRequestInformationTunnelConnection: new(_ConnectionRequestInformationTunnelConnection)}
}

type _ConnectionRequestInformationTunnelConnectionBuilder struct {
	*_ConnectionRequestInformationTunnelConnection

	parentBuilder *_ConnectionRequestInformationBuilder

	err *utils.MultiError
}

var _ (ConnectionRequestInformationTunnelConnectionBuilder) = (*_ConnectionRequestInformationTunnelConnectionBuilder)(nil)

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) setParent(contract ConnectionRequestInformationContract) {
	b.ConnectionRequestInformationContract = contract
	contract.(*_ConnectionRequestInformation)._SubType = b._ConnectionRequestInformationTunnelConnection
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) WithMandatoryFields(knxLayer KnxLayer) ConnectionRequestInformationTunnelConnectionBuilder {
	return b.WithKnxLayer(knxLayer)
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) WithKnxLayer(knxLayer KnxLayer) ConnectionRequestInformationTunnelConnectionBuilder {
	b.KnxLayer = knxLayer
	return b
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) Build() (ConnectionRequestInformationTunnelConnection, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectionRequestInformationTunnelConnection.deepCopy(), nil
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) MustBuild() ConnectionRequestInformationTunnelConnection {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) Done() ConnectionRequestInformationBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewConnectionRequestInformationBuilder().(*_ConnectionRequestInformationBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) buildForConnectionRequestInformation() (ConnectionRequestInformation, error) {
	return b.Build()
}

func (b *_ConnectionRequestInformationTunnelConnectionBuilder) DeepCopy() any {
	_copy := b.CreateConnectionRequestInformationTunnelConnectionBuilder().(*_ConnectionRequestInformationTunnelConnectionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectionRequestInformationTunnelConnectionBuilder creates a ConnectionRequestInformationTunnelConnectionBuilder
func (b *_ConnectionRequestInformationTunnelConnection) CreateConnectionRequestInformationTunnelConnectionBuilder() ConnectionRequestInformationTunnelConnectionBuilder {
	if b == nil {
		return NewConnectionRequestInformationTunnelConnectionBuilder()
	}
	return &_ConnectionRequestInformationTunnelConnectionBuilder{_ConnectionRequestInformationTunnelConnection: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetConnectionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetParent() ConnectionRequestInformationContract {
	return m.ConnectionRequestInformationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetKnxLayer() KnxLayer {
	return m.KnxLayer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformationTunnelConnection(structType any) ConnectionRequestInformationTunnelConnection {
	if casted, ok := structType.(ConnectionRequestInformationTunnelConnection); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformationTunnelConnection); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformationTunnelConnection) GetTypeName() string {
	return "ConnectionRequestInformationTunnelConnection"
}

func (m *_ConnectionRequestInformationTunnelConnection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).getLengthInBits(ctx))

	// Simple field (knxLayer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionRequestInformationTunnelConnection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionRequestInformationTunnelConnection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ConnectionRequestInformation) (__connectionRequestInformationTunnelConnection ConnectionRequestInformationTunnelConnection, err error) {
	m.ConnectionRequestInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationTunnelConnection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformationTunnelConnection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	knxLayer, err := ReadEnumField[KnxLayer](ctx, "knxLayer", "KnxLayer", ReadEnum(KnxLayerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'knxLayer' field"))
	}
	m.KnxLayer = knxLayer

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationTunnelConnection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformationTunnelConnection")
	}

	return m, nil
}

func (m *_ConnectionRequestInformationTunnelConnection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequestInformationTunnelConnection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationTunnelConnection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformationTunnelConnection")
		}

		if err := WriteSimpleEnumField[KnxLayer](ctx, "knxLayer", "KnxLayer", m.GetKnxLayer(), WriteEnum[KnxLayer, uint8](KnxLayer.GetValue, KnxLayer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'knxLayer' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationTunnelConnection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequestInformationTunnelConnection")
		}
		return nil
	}
	return m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionRequestInformationTunnelConnection) IsConnectionRequestInformationTunnelConnection() {
}

func (m *_ConnectionRequestInformationTunnelConnection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionRequestInformationTunnelConnection) deepCopy() *_ConnectionRequestInformationTunnelConnection {
	if m == nil {
		return nil
	}
	_ConnectionRequestInformationTunnelConnectionCopy := &_ConnectionRequestInformationTunnelConnection{
		m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).deepCopy(),
		m.KnxLayer,
		m.reservedField0,
	}
	_ConnectionRequestInformationTunnelConnectionCopy.ConnectionRequestInformationContract.(*_ConnectionRequestInformation)._SubType = m
	return _ConnectionRequestInformationTunnelConnectionCopy
}

func (m *_ConnectionRequestInformationTunnelConnection) String() string {
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
