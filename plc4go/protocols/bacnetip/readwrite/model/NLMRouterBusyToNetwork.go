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

// NLMRouterBusyToNetwork is the corresponding interface of NLMRouterBusyToNetwork
type NLMRouterBusyToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetDestinationNetworkAddresses returns DestinationNetworkAddresses (property field)
	GetDestinationNetworkAddresses() []uint16
	// IsNLMRouterBusyToNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMRouterBusyToNetwork()
	// CreateBuilder creates a NLMRouterBusyToNetworkBuilder
	CreateNLMRouterBusyToNetworkBuilder() NLMRouterBusyToNetworkBuilder
}

// _NLMRouterBusyToNetwork is the data-structure of this message
type _NLMRouterBusyToNetwork struct {
	NLMContract
	DestinationNetworkAddresses []uint16
}

var _ NLMRouterBusyToNetwork = (*_NLMRouterBusyToNetwork)(nil)
var _ NLMRequirements = (*_NLMRouterBusyToNetwork)(nil)

// NewNLMRouterBusyToNetwork factory function for _NLMRouterBusyToNetwork
func NewNLMRouterBusyToNetwork(destinationNetworkAddresses []uint16, apduLength uint16) *_NLMRouterBusyToNetwork {
	_result := &_NLMRouterBusyToNetwork{
		NLMContract:                 NewNLM(apduLength),
		DestinationNetworkAddresses: destinationNetworkAddresses,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMRouterBusyToNetworkBuilder is a builder for NLMRouterBusyToNetwork
type NLMRouterBusyToNetworkBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationNetworkAddresses []uint16) NLMRouterBusyToNetworkBuilder
	// WithDestinationNetworkAddresses adds DestinationNetworkAddresses (property field)
	WithDestinationNetworkAddresses(...uint16) NLMRouterBusyToNetworkBuilder
	// Build builds the NLMRouterBusyToNetwork or returns an error if something is wrong
	Build() (NLMRouterBusyToNetwork, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMRouterBusyToNetwork
}

// NewNLMRouterBusyToNetworkBuilder() creates a NLMRouterBusyToNetworkBuilder
func NewNLMRouterBusyToNetworkBuilder() NLMRouterBusyToNetworkBuilder {
	return &_NLMRouterBusyToNetworkBuilder{_NLMRouterBusyToNetwork: new(_NLMRouterBusyToNetwork)}
}

type _NLMRouterBusyToNetworkBuilder struct {
	*_NLMRouterBusyToNetwork

	err *utils.MultiError
}

var _ (NLMRouterBusyToNetworkBuilder) = (*_NLMRouterBusyToNetworkBuilder)(nil)

func (m *_NLMRouterBusyToNetworkBuilder) WithMandatoryFields(destinationNetworkAddresses []uint16) NLMRouterBusyToNetworkBuilder {
	return m.WithDestinationNetworkAddresses(destinationNetworkAddresses...)
}

func (m *_NLMRouterBusyToNetworkBuilder) WithDestinationNetworkAddresses(destinationNetworkAddresses ...uint16) NLMRouterBusyToNetworkBuilder {
	m.DestinationNetworkAddresses = destinationNetworkAddresses
	return m
}

func (m *_NLMRouterBusyToNetworkBuilder) Build() (NLMRouterBusyToNetwork, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._NLMRouterBusyToNetwork.deepCopy(), nil
}

func (m *_NLMRouterBusyToNetworkBuilder) MustBuild() NLMRouterBusyToNetwork {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_NLMRouterBusyToNetworkBuilder) DeepCopy() any {
	return m.CreateNLMRouterBusyToNetworkBuilder()
}

// CreateNLMRouterBusyToNetworkBuilder creates a NLMRouterBusyToNetworkBuilder
func (m *_NLMRouterBusyToNetwork) CreateNLMRouterBusyToNetworkBuilder() NLMRouterBusyToNetworkBuilder {
	if m == nil {
		return NewNLMRouterBusyToNetworkBuilder()
	}
	return &_NLMRouterBusyToNetworkBuilder{_NLMRouterBusyToNetwork: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRouterBusyToNetwork) GetMessageType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRouterBusyToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRouterBusyToNetwork) GetDestinationNetworkAddresses() []uint16 {
	return m.DestinationNetworkAddresses
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMRouterBusyToNetwork(structType any) NLMRouterBusyToNetwork {
	if casted, ok := structType.(NLMRouterBusyToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRouterBusyToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRouterBusyToNetwork) GetTypeName() string {
	return "NLMRouterBusyToNetwork"
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Array field
	if len(m.DestinationNetworkAddresses) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddresses))
	}

	return lengthInBits
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMRouterBusyToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMRouterBusyToNetwork NLMRouterBusyToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRouterBusyToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRouterBusyToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationNetworkAddresses, err := ReadLengthArrayField[uint16](ctx, "destinationNetworkAddresses", ReadUnsignedShort(readBuffer, uint8(16)), int(int32(apduLength)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddresses' field"))
	}
	m.DestinationNetworkAddresses = destinationNetworkAddresses

	if closeErr := readBuffer.CloseContext("NLMRouterBusyToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRouterBusyToNetwork")
	}

	return m, nil
}

func (m *_NLMRouterBusyToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRouterBusyToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRouterBusyToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRouterBusyToNetwork")
		}

		if err := WriteSimpleTypeArrayField(ctx, "destinationNetworkAddresses", m.GetDestinationNetworkAddresses(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddresses' field")
		}

		if popErr := writeBuffer.PopContext("NLMRouterBusyToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRouterBusyToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMRouterBusyToNetwork) IsNLMRouterBusyToNetwork() {}

func (m *_NLMRouterBusyToNetwork) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMRouterBusyToNetwork) deepCopy() *_NLMRouterBusyToNetwork {
	if m == nil {
		return nil
	}
	_NLMRouterBusyToNetworkCopy := &_NLMRouterBusyToNetwork{
		m.NLMContract.(*_NLM).deepCopy(),
		utils.DeepCopySlice[uint16, uint16](m.DestinationNetworkAddresses),
	}
	m.NLMContract.(*_NLM)._SubType = m
	return _NLMRouterBusyToNetworkCopy
}

func (m *_NLMRouterBusyToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
