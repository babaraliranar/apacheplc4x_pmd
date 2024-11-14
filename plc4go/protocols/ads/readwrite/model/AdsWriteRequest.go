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

// AdsWriteRequest is the corresponding interface of AdsWriteRequest
type AdsWriteRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetData returns Data (property field)
	GetData() []byte
	// IsAdsWriteRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsWriteRequest()
	// CreateBuilder creates a AdsWriteRequestBuilder
	CreateAdsWriteRequestBuilder() AdsWriteRequestBuilder
}

// _AdsWriteRequest is the data-structure of this message
type _AdsWriteRequest struct {
	AmsPacketContract
	IndexGroup  uint32
	IndexOffset uint32
	Data        []byte
}

var _ AdsWriteRequest = (*_AdsWriteRequest)(nil)
var _ AmsPacketRequirements = (*_AdsWriteRequest)(nil)

// NewAdsWriteRequest factory function for _AdsWriteRequest
func NewAdsWriteRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, indexGroup uint32, indexOffset uint32, data []byte) *_AdsWriteRequest {
	_result := &_AdsWriteRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		IndexGroup:        indexGroup,
		IndexOffset:       indexOffset,
		Data:              data,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsWriteRequestBuilder is a builder for AdsWriteRequest
type AdsWriteRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(indexGroup uint32, indexOffset uint32, data []byte) AdsWriteRequestBuilder
	// WithIndexGroup adds IndexGroup (property field)
	WithIndexGroup(uint32) AdsWriteRequestBuilder
	// WithIndexOffset adds IndexOffset (property field)
	WithIndexOffset(uint32) AdsWriteRequestBuilder
	// WithData adds Data (property field)
	WithData(...byte) AdsWriteRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AmsPacketBuilder
	// Build builds the AdsWriteRequest or returns an error if something is wrong
	Build() (AdsWriteRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsWriteRequest
}

// NewAdsWriteRequestBuilder() creates a AdsWriteRequestBuilder
func NewAdsWriteRequestBuilder() AdsWriteRequestBuilder {
	return &_AdsWriteRequestBuilder{_AdsWriteRequest: new(_AdsWriteRequest)}
}

type _AdsWriteRequestBuilder struct {
	*_AdsWriteRequest

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsWriteRequestBuilder) = (*_AdsWriteRequestBuilder)(nil)

func (b *_AdsWriteRequestBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
	contract.(*_AmsPacket)._SubType = b._AdsWriteRequest
}

func (b *_AdsWriteRequestBuilder) WithMandatoryFields(indexGroup uint32, indexOffset uint32, data []byte) AdsWriteRequestBuilder {
	return b.WithIndexGroup(indexGroup).WithIndexOffset(indexOffset).WithData(data...)
}

func (b *_AdsWriteRequestBuilder) WithIndexGroup(indexGroup uint32) AdsWriteRequestBuilder {
	b.IndexGroup = indexGroup
	return b
}

func (b *_AdsWriteRequestBuilder) WithIndexOffset(indexOffset uint32) AdsWriteRequestBuilder {
	b.IndexOffset = indexOffset
	return b
}

func (b *_AdsWriteRequestBuilder) WithData(data ...byte) AdsWriteRequestBuilder {
	b.Data = data
	return b
}

func (b *_AdsWriteRequestBuilder) Build() (AdsWriteRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsWriteRequest.deepCopy(), nil
}

func (b *_AdsWriteRequestBuilder) MustBuild() AdsWriteRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsWriteRequestBuilder) Done() AmsPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAmsPacketBuilder().(*_AmsPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsWriteRequestBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsWriteRequestBuilder) DeepCopy() any {
	_copy := b.CreateAdsWriteRequestBuilder().(*_AdsWriteRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsWriteRequestBuilder creates a AdsWriteRequestBuilder
func (b *_AdsWriteRequest) CreateAdsWriteRequestBuilder() AdsWriteRequestBuilder {
	if b == nil {
		return NewAdsWriteRequestBuilder()
	}
	return &_AdsWriteRequestBuilder{_AdsWriteRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsWriteRequest) GetCommandId() CommandId {
	return CommandId_ADS_WRITE
}

func (m *_AdsWriteRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsWriteRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsWriteRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsWriteRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsWriteRequest(structType any) AdsWriteRequest {
	if casted, ok := structType.(AdsWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsWriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsWriteRequest) GetTypeName() string {
	return "AdsWriteRequest"
}

func (m *_AdsWriteRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsWriteRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsWriteRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsWriteRequest AdsWriteRequest, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsWriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsWriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	indexGroup, err := ReadSimpleField(ctx, "indexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexGroup' field"))
	}
	m.IndexGroup = indexGroup

	indexOffset, err := ReadSimpleField(ctx, "indexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexOffset' field"))
	}
	m.IndexOffset = indexOffset

	length, err := ReadImplicitField[uint32](ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	data, err := readBuffer.ReadByteArray("data", int(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AdsWriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsWriteRequest")
	}

	return m, nil
}

func (m *_AdsWriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsWriteRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsWriteRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "indexGroup", m.GetIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "indexOffset", m.GetIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexOffset' field")
		}
		length := uint32(uint32(len(m.GetData())))
		if err := WriteImplicitField(ctx, "length", length, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsWriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsWriteRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsWriteRequest) IsAdsWriteRequest() {}

func (m *_AdsWriteRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsWriteRequest) deepCopy() *_AdsWriteRequest {
	if m == nil {
		return nil
	}
	_AdsWriteRequestCopy := &_AdsWriteRequest{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
		m.IndexGroup,
		m.IndexOffset,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	_AdsWriteRequestCopy.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsWriteRequestCopy
}

func (m *_AdsWriteRequest) String() string {
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
