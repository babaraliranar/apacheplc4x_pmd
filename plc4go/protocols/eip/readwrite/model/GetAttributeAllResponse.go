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

// GetAttributeAllResponse is the corresponding interface of GetAttributeAllResponse
type GetAttributeAllResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// GetAttributes returns Attributes (property field)
	GetAttributes() CIPAttributes
	// IsGetAttributeAllResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetAttributeAllResponse()
	// CreateBuilder creates a GetAttributeAllResponseBuilder
	CreateGetAttributeAllResponseBuilder() GetAttributeAllResponseBuilder
}

// _GetAttributeAllResponse is the data-structure of this message
type _GetAttributeAllResponse struct {
	CipServiceContract
	Status     uint8
	ExtStatus  uint8
	Attributes CIPAttributes
	// Reserved Fields
	reservedField0 *uint8
}

var _ GetAttributeAllResponse = (*_GetAttributeAllResponse)(nil)
var _ CipServiceRequirements = (*_GetAttributeAllResponse)(nil)

// NewGetAttributeAllResponse factory function for _GetAttributeAllResponse
func NewGetAttributeAllResponse(status uint8, extStatus uint8, attributes CIPAttributes, serviceLen uint16) *_GetAttributeAllResponse {
	_result := &_GetAttributeAllResponse{
		CipServiceContract: NewCipService(serviceLen),
		Status:             status,
		ExtStatus:          extStatus,
		Attributes:         attributes,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GetAttributeAllResponseBuilder is a builder for GetAttributeAllResponse
type GetAttributeAllResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(status uint8, extStatus uint8) GetAttributeAllResponseBuilder
	// WithStatus adds Status (property field)
	WithStatus(uint8) GetAttributeAllResponseBuilder
	// WithExtStatus adds ExtStatus (property field)
	WithExtStatus(uint8) GetAttributeAllResponseBuilder
	// WithAttributes adds Attributes (property field)
	WithOptionalAttributes(CIPAttributes) GetAttributeAllResponseBuilder
	// WithOptionalAttributesBuilder adds Attributes (property field) which is build by the builder
	WithOptionalAttributesBuilder(func(CIPAttributesBuilder) CIPAttributesBuilder) GetAttributeAllResponseBuilder
	// Build builds the GetAttributeAllResponse or returns an error if something is wrong
	Build() (GetAttributeAllResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GetAttributeAllResponse
}

// NewGetAttributeAllResponseBuilder() creates a GetAttributeAllResponseBuilder
func NewGetAttributeAllResponseBuilder() GetAttributeAllResponseBuilder {
	return &_GetAttributeAllResponseBuilder{_GetAttributeAllResponse: new(_GetAttributeAllResponse)}
}

type _GetAttributeAllResponseBuilder struct {
	*_GetAttributeAllResponse

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (GetAttributeAllResponseBuilder) = (*_GetAttributeAllResponseBuilder)(nil)

func (b *_GetAttributeAllResponseBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
}

func (b *_GetAttributeAllResponseBuilder) WithMandatoryFields(status uint8, extStatus uint8) GetAttributeAllResponseBuilder {
	return b.WithStatus(status).WithExtStatus(extStatus)
}

func (b *_GetAttributeAllResponseBuilder) WithStatus(status uint8) GetAttributeAllResponseBuilder {
	b.Status = status
	return b
}

func (b *_GetAttributeAllResponseBuilder) WithExtStatus(extStatus uint8) GetAttributeAllResponseBuilder {
	b.ExtStatus = extStatus
	return b
}

func (b *_GetAttributeAllResponseBuilder) WithOptionalAttributes(attributes CIPAttributes) GetAttributeAllResponseBuilder {
	b.Attributes = attributes
	return b
}

func (b *_GetAttributeAllResponseBuilder) WithOptionalAttributesBuilder(builderSupplier func(CIPAttributesBuilder) CIPAttributesBuilder) GetAttributeAllResponseBuilder {
	builder := builderSupplier(b.Attributes.CreateCIPAttributesBuilder())
	var err error
	b.Attributes, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CIPAttributesBuilder failed"))
	}
	return b
}

func (b *_GetAttributeAllResponseBuilder) Build() (GetAttributeAllResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._GetAttributeAllResponse.deepCopy(), nil
}

func (b *_GetAttributeAllResponseBuilder) MustBuild() GetAttributeAllResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_GetAttributeAllResponseBuilder) Done() CipServiceBuilder {
	return b.parentBuilder
}

func (b *_GetAttributeAllResponseBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_GetAttributeAllResponseBuilder) DeepCopy() any {
	_copy := b.CreateGetAttributeAllResponseBuilder().(*_GetAttributeAllResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGetAttributeAllResponseBuilder creates a GetAttributeAllResponseBuilder
func (b *_GetAttributeAllResponse) CreateGetAttributeAllResponseBuilder() GetAttributeAllResponseBuilder {
	if b == nil {
		return NewGetAttributeAllResponseBuilder()
	}
	return &_GetAttributeAllResponseBuilder{_GetAttributeAllResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeAllResponse) GetService() uint8 {
	return 0x01
}

func (m *_GetAttributeAllResponse) GetResponse() bool {
	return bool(true)
}

func (m *_GetAttributeAllResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeAllResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetAttributeAllResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_GetAttributeAllResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *_GetAttributeAllResponse) GetAttributes() CIPAttributes {
	return m.Attributes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGetAttributeAllResponse(structType any) GetAttributeAllResponse {
	if casted, ok := structType.(GetAttributeAllResponse); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeAllResponse); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeAllResponse) GetTypeName() string {
	return "GetAttributeAllResponse"
}

func (m *_GetAttributeAllResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Optional Field (attributes)
	if m.Attributes != nil {
		lengthInBits += m.Attributes.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_GetAttributeAllResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeAllResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__getAttributeAllResponse GetAttributeAllResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeAllResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeAllResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	extStatus, err := ReadSimpleField(ctx, "extStatus", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extStatus' field"))
	}
	m.ExtStatus = extStatus

	var attributes CIPAttributes
	_attributes, err := ReadOptionalField[CIPAttributes](ctx, "attributes", ReadComplex[CIPAttributes](CIPAttributesParseWithBufferProducer((uint16)(uint16(serviceLen)-uint16(uint16(4)))), readBuffer), bool(((serviceLen)-(4)) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributes' field"))
	}
	if _attributes != nil {
		attributes = *_attributes
		m.Attributes = attributes
	}

	if closeErr := readBuffer.CloseContext("GetAttributeAllResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeAllResponse")
	}

	return m, nil
}

func (m *_GetAttributeAllResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeAllResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeAllResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeAllResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "extStatus", m.GetExtStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'extStatus' field")
		}

		if err := WriteOptionalField[CIPAttributes](ctx, "attributes", GetRef(m.GetAttributes()), WriteComplex[CIPAttributes](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'attributes' field")
		}

		if popErr := writeBuffer.PopContext("GetAttributeAllResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeAllResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeAllResponse) IsGetAttributeAllResponse() {}

func (m *_GetAttributeAllResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GetAttributeAllResponse) deepCopy() *_GetAttributeAllResponse {
	if m == nil {
		return nil
	}
	_GetAttributeAllResponseCopy := &_GetAttributeAllResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
		m.Status,
		m.ExtStatus,
		m.Attributes.DeepCopy().(CIPAttributes),
		m.reservedField0,
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _GetAttributeAllResponseCopy
}

func (m *_GetAttributeAllResponse) String() string {
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
