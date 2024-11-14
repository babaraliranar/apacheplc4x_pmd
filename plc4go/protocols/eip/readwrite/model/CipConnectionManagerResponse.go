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

// CipConnectionManagerResponse is the corresponding interface of CipConnectionManagerResponse
type CipConnectionManagerResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetOtConnectionId returns OtConnectionId (property field)
	GetOtConnectionId() uint32
	// GetToConnectionId returns ToConnectionId (property field)
	GetToConnectionId() uint32
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetOtApi returns OtApi (property field)
	GetOtApi() uint32
	// GetToApi returns ToApi (property field)
	GetToApi() uint32
	// IsCipConnectionManagerResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipConnectionManagerResponse()
	// CreateBuilder creates a CipConnectionManagerResponseBuilder
	CreateCipConnectionManagerResponseBuilder() CipConnectionManagerResponseBuilder
}

// _CipConnectionManagerResponse is the data-structure of this message
type _CipConnectionManagerResponse struct {
	CipServiceContract
	OtConnectionId         uint32
	ToConnectionId         uint32
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	OtApi                  uint32
	ToApi                  uint32
	// Reserved Fields
	reservedField0 *uint32
	reservedField1 *uint8
}

var _ CipConnectionManagerResponse = (*_CipConnectionManagerResponse)(nil)
var _ CipServiceRequirements = (*_CipConnectionManagerResponse)(nil)

// NewCipConnectionManagerResponse factory function for _CipConnectionManagerResponse
func NewCipConnectionManagerResponse(otConnectionId uint32, toConnectionId uint32, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, otApi uint32, toApi uint32, serviceLen uint16) *_CipConnectionManagerResponse {
	_result := &_CipConnectionManagerResponse{
		CipServiceContract:     NewCipService(serviceLen),
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		OtApi:                  otApi,
		ToApi:                  toApi,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipConnectionManagerResponseBuilder is a builder for CipConnectionManagerResponse
type CipConnectionManagerResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(otConnectionId uint32, toConnectionId uint32, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, otApi uint32, toApi uint32) CipConnectionManagerResponseBuilder
	// WithOtConnectionId adds OtConnectionId (property field)
	WithOtConnectionId(uint32) CipConnectionManagerResponseBuilder
	// WithToConnectionId adds ToConnectionId (property field)
	WithToConnectionId(uint32) CipConnectionManagerResponseBuilder
	// WithConnectionSerialNumber adds ConnectionSerialNumber (property field)
	WithConnectionSerialNumber(uint16) CipConnectionManagerResponseBuilder
	// WithOriginatorVendorId adds OriginatorVendorId (property field)
	WithOriginatorVendorId(uint16) CipConnectionManagerResponseBuilder
	// WithOriginatorSerialNumber adds OriginatorSerialNumber (property field)
	WithOriginatorSerialNumber(uint32) CipConnectionManagerResponseBuilder
	// WithOtApi adds OtApi (property field)
	WithOtApi(uint32) CipConnectionManagerResponseBuilder
	// WithToApi adds ToApi (property field)
	WithToApi(uint32) CipConnectionManagerResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the CipConnectionManagerResponse or returns an error if something is wrong
	Build() (CipConnectionManagerResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipConnectionManagerResponse
}

// NewCipConnectionManagerResponseBuilder() creates a CipConnectionManagerResponseBuilder
func NewCipConnectionManagerResponseBuilder() CipConnectionManagerResponseBuilder {
	return &_CipConnectionManagerResponseBuilder{_CipConnectionManagerResponse: new(_CipConnectionManagerResponse)}
}

type _CipConnectionManagerResponseBuilder struct {
	*_CipConnectionManagerResponse

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (CipConnectionManagerResponseBuilder) = (*_CipConnectionManagerResponseBuilder)(nil)

func (b *_CipConnectionManagerResponseBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
	contract.(*_CipService)._SubType = b._CipConnectionManagerResponse
}

func (b *_CipConnectionManagerResponseBuilder) WithMandatoryFields(otConnectionId uint32, toConnectionId uint32, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, otApi uint32, toApi uint32) CipConnectionManagerResponseBuilder {
	return b.WithOtConnectionId(otConnectionId).WithToConnectionId(toConnectionId).WithConnectionSerialNumber(connectionSerialNumber).WithOriginatorVendorId(originatorVendorId).WithOriginatorSerialNumber(originatorSerialNumber).WithOtApi(otApi).WithToApi(toApi)
}

func (b *_CipConnectionManagerResponseBuilder) WithOtConnectionId(otConnectionId uint32) CipConnectionManagerResponseBuilder {
	b.OtConnectionId = otConnectionId
	return b
}

func (b *_CipConnectionManagerResponseBuilder) WithToConnectionId(toConnectionId uint32) CipConnectionManagerResponseBuilder {
	b.ToConnectionId = toConnectionId
	return b
}

func (b *_CipConnectionManagerResponseBuilder) WithConnectionSerialNumber(connectionSerialNumber uint16) CipConnectionManagerResponseBuilder {
	b.ConnectionSerialNumber = connectionSerialNumber
	return b
}

func (b *_CipConnectionManagerResponseBuilder) WithOriginatorVendorId(originatorVendorId uint16) CipConnectionManagerResponseBuilder {
	b.OriginatorVendorId = originatorVendorId
	return b
}

func (b *_CipConnectionManagerResponseBuilder) WithOriginatorSerialNumber(originatorSerialNumber uint32) CipConnectionManagerResponseBuilder {
	b.OriginatorSerialNumber = originatorSerialNumber
	return b
}

func (b *_CipConnectionManagerResponseBuilder) WithOtApi(otApi uint32) CipConnectionManagerResponseBuilder {
	b.OtApi = otApi
	return b
}

func (b *_CipConnectionManagerResponseBuilder) WithToApi(toApi uint32) CipConnectionManagerResponseBuilder {
	b.ToApi = toApi
	return b
}

func (b *_CipConnectionManagerResponseBuilder) Build() (CipConnectionManagerResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipConnectionManagerResponse.deepCopy(), nil
}

func (b *_CipConnectionManagerResponseBuilder) MustBuild() CipConnectionManagerResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipConnectionManagerResponseBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_CipConnectionManagerResponseBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_CipConnectionManagerResponseBuilder) DeepCopy() any {
	_copy := b.CreateCipConnectionManagerResponseBuilder().(*_CipConnectionManagerResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipConnectionManagerResponseBuilder creates a CipConnectionManagerResponseBuilder
func (b *_CipConnectionManagerResponse) CreateCipConnectionManagerResponseBuilder() CipConnectionManagerResponseBuilder {
	if b == nil {
		return NewCipConnectionManagerResponseBuilder()
	}
	return &_CipConnectionManagerResponseBuilder{_CipConnectionManagerResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerResponse) GetService() uint8 {
	return 0x5B
}

func (m *_CipConnectionManagerResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectionManagerResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerResponse) GetOtConnectionId() uint32 {
	return m.OtConnectionId
}

func (m *_CipConnectionManagerResponse) GetToConnectionId() uint32 {
	return m.ToConnectionId
}

func (m *_CipConnectionManagerResponse) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerResponse) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerResponse) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerResponse) GetOtApi() uint32 {
	return m.OtApi
}

func (m *_CipConnectionManagerResponse) GetToApi() uint32 {
	return m.ToApi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerResponse(structType any) CipConnectionManagerResponse {
	if casted, ok := structType.(CipConnectionManagerResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerResponse) GetTypeName() string {
	return "CipConnectionManagerResponse"
}

func (m *_CipConnectionManagerResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 24

	// Simple field (otConnectionId)
	lengthInBits += 32

	// Simple field (toConnectionId)
	lengthInBits += 32

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (otApi)
	lengthInBits += 32

	// Simple field (toApi)
	lengthInBits += 32

	// Implicit Field (replySize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipConnectionManagerResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectionManagerResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectionManagerResponse CipConnectionManagerResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedInt(readBuffer, uint8(24)), uint32(0x000000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	otConnectionId, err := ReadSimpleField(ctx, "otConnectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otConnectionId' field"))
	}
	m.OtConnectionId = otConnectionId

	toConnectionId, err := ReadSimpleField(ctx, "toConnectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toConnectionId' field"))
	}
	m.ToConnectionId = toConnectionId

	connectionSerialNumber, err := ReadSimpleField(ctx, "connectionSerialNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSerialNumber' field"))
	}
	m.ConnectionSerialNumber = connectionSerialNumber

	originatorVendorId, err := ReadSimpleField(ctx, "originatorVendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorVendorId' field"))
	}
	m.OriginatorVendorId = originatorVendorId

	originatorSerialNumber, err := ReadSimpleField(ctx, "originatorSerialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorSerialNumber' field"))
	}
	m.OriginatorSerialNumber = originatorSerialNumber

	otApi, err := ReadSimpleField(ctx, "otApi", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otApi' field"))
	}
	m.OtApi = otApi

	toApi, err := ReadSimpleField(ctx, "toApi", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toApi' field"))
	}
	m.ToApi = toApi

	replySize, err := ReadImplicitField[uint8](ctx, "replySize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'replySize' field"))
	}
	_ = replySize

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	if closeErr := readBuffer.CloseContext("CipConnectionManagerResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerResponse")
	}

	return m, nil
}

func (m *_CipConnectionManagerResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerResponse")
		}

		if err := WriteReservedField[uint32](ctx, "reserved", uint32(0x000000), WriteUnsignedInt(writeBuffer, 24)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint32](ctx, "otConnectionId", m.GetOtConnectionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'otConnectionId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "toConnectionId", m.GetToConnectionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'toConnectionId' field")
		}

		if err := WriteSimpleField[uint16](ctx, "connectionSerialNumber", m.GetConnectionSerialNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionSerialNumber' field")
		}

		if err := WriteSimpleField[uint16](ctx, "originatorVendorId", m.GetOriginatorVendorId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorVendorId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "originatorSerialNumber", m.GetOriginatorSerialNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorSerialNumber' field")
		}

		if err := WriteSimpleField[uint32](ctx, "otApi", m.GetOtApi(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'otApi' field")
		}

		if err := WriteSimpleField[uint32](ctx, "toApi", m.GetToApi(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'toApi' field")
		}
		replySize := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(30)))
		if err := WriteImplicitField(ctx, "replySize", replySize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'replySize' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerResponse) IsCipConnectionManagerResponse() {}

func (m *_CipConnectionManagerResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipConnectionManagerResponse) deepCopy() *_CipConnectionManagerResponse {
	if m == nil {
		return nil
	}
	_CipConnectionManagerResponseCopy := &_CipConnectionManagerResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
		m.OtConnectionId,
		m.ToConnectionId,
		m.ConnectionSerialNumber,
		m.OriginatorVendorId,
		m.OriginatorSerialNumber,
		m.OtApi,
		m.ToApi,
		m.reservedField0,
		m.reservedField1,
	}
	_CipConnectionManagerResponseCopy.CipServiceContract.(*_CipService)._SubType = m
	return _CipConnectionManagerResponseCopy
}

func (m *_CipConnectionManagerResponse) String() string {
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
