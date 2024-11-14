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

// ApduDataDeviceDescriptorRead is the corresponding interface of ApduDataDeviceDescriptorRead
type ApduDataDeviceDescriptorRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// GetDescriptorType returns DescriptorType (property field)
	GetDescriptorType() uint8
	// IsApduDataDeviceDescriptorRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataDeviceDescriptorRead()
	// CreateBuilder creates a ApduDataDeviceDescriptorReadBuilder
	CreateApduDataDeviceDescriptorReadBuilder() ApduDataDeviceDescriptorReadBuilder
}

// _ApduDataDeviceDescriptorRead is the data-structure of this message
type _ApduDataDeviceDescriptorRead struct {
	ApduDataContract
	DescriptorType uint8
}

var _ ApduDataDeviceDescriptorRead = (*_ApduDataDeviceDescriptorRead)(nil)
var _ ApduDataRequirements = (*_ApduDataDeviceDescriptorRead)(nil)

// NewApduDataDeviceDescriptorRead factory function for _ApduDataDeviceDescriptorRead
func NewApduDataDeviceDescriptorRead(descriptorType uint8, dataLength uint8) *_ApduDataDeviceDescriptorRead {
	_result := &_ApduDataDeviceDescriptorRead{
		ApduDataContract: NewApduData(dataLength),
		DescriptorType:   descriptorType,
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataDeviceDescriptorReadBuilder is a builder for ApduDataDeviceDescriptorRead
type ApduDataDeviceDescriptorReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(descriptorType uint8) ApduDataDeviceDescriptorReadBuilder
	// WithDescriptorType adds DescriptorType (property field)
	WithDescriptorType(uint8) ApduDataDeviceDescriptorReadBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataBuilder
	// Build builds the ApduDataDeviceDescriptorRead or returns an error if something is wrong
	Build() (ApduDataDeviceDescriptorRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataDeviceDescriptorRead
}

// NewApduDataDeviceDescriptorReadBuilder() creates a ApduDataDeviceDescriptorReadBuilder
func NewApduDataDeviceDescriptorReadBuilder() ApduDataDeviceDescriptorReadBuilder {
	return &_ApduDataDeviceDescriptorReadBuilder{_ApduDataDeviceDescriptorRead: new(_ApduDataDeviceDescriptorRead)}
}

type _ApduDataDeviceDescriptorReadBuilder struct {
	*_ApduDataDeviceDescriptorRead

	parentBuilder *_ApduDataBuilder

	err *utils.MultiError
}

var _ (ApduDataDeviceDescriptorReadBuilder) = (*_ApduDataDeviceDescriptorReadBuilder)(nil)

func (b *_ApduDataDeviceDescriptorReadBuilder) setParent(contract ApduDataContract) {
	b.ApduDataContract = contract
	contract.(*_ApduData)._SubType = b._ApduDataDeviceDescriptorRead
}

func (b *_ApduDataDeviceDescriptorReadBuilder) WithMandatoryFields(descriptorType uint8) ApduDataDeviceDescriptorReadBuilder {
	return b.WithDescriptorType(descriptorType)
}

func (b *_ApduDataDeviceDescriptorReadBuilder) WithDescriptorType(descriptorType uint8) ApduDataDeviceDescriptorReadBuilder {
	b.DescriptorType = descriptorType
	return b
}

func (b *_ApduDataDeviceDescriptorReadBuilder) Build() (ApduDataDeviceDescriptorRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataDeviceDescriptorRead.deepCopy(), nil
}

func (b *_ApduDataDeviceDescriptorReadBuilder) MustBuild() ApduDataDeviceDescriptorRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataDeviceDescriptorReadBuilder) Done() ApduDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataBuilder().(*_ApduDataBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataDeviceDescriptorReadBuilder) buildForApduData() (ApduData, error) {
	return b.Build()
}

func (b *_ApduDataDeviceDescriptorReadBuilder) DeepCopy() any {
	_copy := b.CreateApduDataDeviceDescriptorReadBuilder().(*_ApduDataDeviceDescriptorReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataDeviceDescriptorReadBuilder creates a ApduDataDeviceDescriptorReadBuilder
func (b *_ApduDataDeviceDescriptorRead) CreateApduDataDeviceDescriptorReadBuilder() ApduDataDeviceDescriptorReadBuilder {
	if b == nil {
		return NewApduDataDeviceDescriptorReadBuilder()
	}
	return &_ApduDataDeviceDescriptorReadBuilder{_ApduDataDeviceDescriptorRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataDeviceDescriptorRead) GetApciType() uint8 {
	return 0xC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataDeviceDescriptorRead) GetParent() ApduDataContract {
	return m.ApduDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataDeviceDescriptorRead) GetDescriptorType() uint8 {
	return m.DescriptorType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataDeviceDescriptorRead(structType any) ApduDataDeviceDescriptorRead {
	if casted, ok := structType.(ApduDataDeviceDescriptorRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataDeviceDescriptorRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataDeviceDescriptorRead) GetTypeName() string {
	return "ApduDataDeviceDescriptorRead"
}

func (m *_ApduDataDeviceDescriptorRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	// Simple field (descriptorType)
	lengthInBits += 6

	return lengthInBits
}

func (m *_ApduDataDeviceDescriptorRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataDeviceDescriptorRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataDeviceDescriptorRead ApduDataDeviceDescriptorRead, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataDeviceDescriptorRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	descriptorType, err := ReadSimpleField(ctx, "descriptorType", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'descriptorType' field"))
	}
	m.DescriptorType = descriptorType

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataDeviceDescriptorRead")
	}

	return m, nil
}

func (m *_ApduDataDeviceDescriptorRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataDeviceDescriptorRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataDeviceDescriptorRead")
		}

		if err := WriteSimpleField[uint8](ctx, "descriptorType", m.GetDescriptorType(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'descriptorType' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataDeviceDescriptorRead")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataDeviceDescriptorRead) IsApduDataDeviceDescriptorRead() {}

func (m *_ApduDataDeviceDescriptorRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataDeviceDescriptorRead) deepCopy() *_ApduDataDeviceDescriptorRead {
	if m == nil {
		return nil
	}
	_ApduDataDeviceDescriptorReadCopy := &_ApduDataDeviceDescriptorRead{
		m.ApduDataContract.(*_ApduData).deepCopy(),
		m.DescriptorType,
	}
	_ApduDataDeviceDescriptorReadCopy.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataDeviceDescriptorReadCopy
}

func (m *_ApduDataDeviceDescriptorRead) String() string {
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
