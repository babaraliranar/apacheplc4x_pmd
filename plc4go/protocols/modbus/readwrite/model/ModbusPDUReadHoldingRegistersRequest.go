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

// ModbusPDUReadHoldingRegistersRequest is the corresponding interface of ModbusPDUReadHoldingRegistersRequest
type ModbusPDUReadHoldingRegistersRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// IsModbusPDUReadHoldingRegistersRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadHoldingRegistersRequest()
	// CreateBuilder creates a ModbusPDUReadHoldingRegistersRequestBuilder
	CreateModbusPDUReadHoldingRegistersRequestBuilder() ModbusPDUReadHoldingRegistersRequestBuilder
}

// _ModbusPDUReadHoldingRegistersRequest is the data-structure of this message
type _ModbusPDUReadHoldingRegistersRequest struct {
	ModbusPDUContract
	StartingAddress uint16
	Quantity        uint16
}

var _ ModbusPDUReadHoldingRegistersRequest = (*_ModbusPDUReadHoldingRegistersRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadHoldingRegistersRequest)(nil)

// NewModbusPDUReadHoldingRegistersRequest factory function for _ModbusPDUReadHoldingRegistersRequest
func NewModbusPDUReadHoldingRegistersRequest(startingAddress uint16, quantity uint16) *_ModbusPDUReadHoldingRegistersRequest {
	_result := &_ModbusPDUReadHoldingRegistersRequest{
		ModbusPDUContract: NewModbusPDU(),
		StartingAddress:   startingAddress,
		Quantity:          quantity,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadHoldingRegistersRequestBuilder is a builder for ModbusPDUReadHoldingRegistersRequest
type ModbusPDUReadHoldingRegistersRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUReadHoldingRegistersRequestBuilder
	// WithStartingAddress adds StartingAddress (property field)
	WithStartingAddress(uint16) ModbusPDUReadHoldingRegistersRequestBuilder
	// WithQuantity adds Quantity (property field)
	WithQuantity(uint16) ModbusPDUReadHoldingRegistersRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUReadHoldingRegistersRequest or returns an error if something is wrong
	Build() (ModbusPDUReadHoldingRegistersRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadHoldingRegistersRequest
}

// NewModbusPDUReadHoldingRegistersRequestBuilder() creates a ModbusPDUReadHoldingRegistersRequestBuilder
func NewModbusPDUReadHoldingRegistersRequestBuilder() ModbusPDUReadHoldingRegistersRequestBuilder {
	return &_ModbusPDUReadHoldingRegistersRequestBuilder{_ModbusPDUReadHoldingRegistersRequest: new(_ModbusPDUReadHoldingRegistersRequest)}
}

type _ModbusPDUReadHoldingRegistersRequestBuilder struct {
	*_ModbusPDUReadHoldingRegistersRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadHoldingRegistersRequestBuilder) = (*_ModbusPDUReadHoldingRegistersRequestBuilder)(nil)

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUReadHoldingRegistersRequest
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUReadHoldingRegistersRequestBuilder {
	return b.WithStartingAddress(startingAddress).WithQuantity(quantity)
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) WithStartingAddress(startingAddress uint16) ModbusPDUReadHoldingRegistersRequestBuilder {
	b.StartingAddress = startingAddress
	return b
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) WithQuantity(quantity uint16) ModbusPDUReadHoldingRegistersRequestBuilder {
	b.Quantity = quantity
	return b
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) Build() (ModbusPDUReadHoldingRegistersRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadHoldingRegistersRequest.deepCopy(), nil
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) MustBuild() ModbusPDUReadHoldingRegistersRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadHoldingRegistersRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadHoldingRegistersRequestBuilder().(*_ModbusPDUReadHoldingRegistersRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadHoldingRegistersRequestBuilder creates a ModbusPDUReadHoldingRegistersRequestBuilder
func (b *_ModbusPDUReadHoldingRegistersRequest) CreateModbusPDUReadHoldingRegistersRequestBuilder() ModbusPDUReadHoldingRegistersRequestBuilder {
	if b == nil {
		return NewModbusPDUReadHoldingRegistersRequestBuilder()
	}
	return &_ModbusPDUReadHoldingRegistersRequestBuilder{_ModbusPDUReadHoldingRegistersRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x03
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadHoldingRegistersRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadHoldingRegistersRequest(structType any) ModbusPDUReadHoldingRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadHoldingRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadHoldingRegistersRequest"
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadHoldingRegistersRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadHoldingRegistersRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadHoldingRegistersRequest ModbusPDUReadHoldingRegistersRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadHoldingRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadHoldingRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingAddress, err := ReadSimpleField(ctx, "startingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingAddress' field"))
	}
	m.StartingAddress = startingAddress

	quantity, err := ReadSimpleField(ctx, "quantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'quantity' field"))
	}
	m.Quantity = quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadHoldingRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadHoldingRegistersRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadHoldingRegistersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadHoldingRegistersRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadHoldingRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadHoldingRegistersRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "startingAddress", m.GetStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "quantity", m.GetQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadHoldingRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadHoldingRegistersRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadHoldingRegistersRequest) IsModbusPDUReadHoldingRegistersRequest() {}

func (m *_ModbusPDUReadHoldingRegistersRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadHoldingRegistersRequest) deepCopy() *_ModbusPDUReadHoldingRegistersRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReadHoldingRegistersRequestCopy := &_ModbusPDUReadHoldingRegistersRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.StartingAddress,
		m.Quantity,
	}
	_ModbusPDUReadHoldingRegistersRequestCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadHoldingRegistersRequestCopy
}

func (m *_ModbusPDUReadHoldingRegistersRequest) String() string {
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
