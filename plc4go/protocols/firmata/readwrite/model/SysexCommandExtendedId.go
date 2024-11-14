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

// SysexCommandExtendedId is the corresponding interface of SysexCommandExtendedId
type SysexCommandExtendedId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// GetId returns Id (property field)
	GetId() []int8
	// IsSysexCommandExtendedId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandExtendedId()
	// CreateBuilder creates a SysexCommandExtendedIdBuilder
	CreateSysexCommandExtendedIdBuilder() SysexCommandExtendedIdBuilder
}

// _SysexCommandExtendedId is the data-structure of this message
type _SysexCommandExtendedId struct {
	SysexCommandContract
	Id []int8
}

var _ SysexCommandExtendedId = (*_SysexCommandExtendedId)(nil)
var _ SysexCommandRequirements = (*_SysexCommandExtendedId)(nil)

// NewSysexCommandExtendedId factory function for _SysexCommandExtendedId
func NewSysexCommandExtendedId(id []int8) *_SysexCommandExtendedId {
	_result := &_SysexCommandExtendedId{
		SysexCommandContract: NewSysexCommand(),
		Id:                   id,
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandExtendedIdBuilder is a builder for SysexCommandExtendedId
type SysexCommandExtendedIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(id []int8) SysexCommandExtendedIdBuilder
	// WithId adds Id (property field)
	WithId(...int8) SysexCommandExtendedIdBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SysexCommandBuilder
	// Build builds the SysexCommandExtendedId or returns an error if something is wrong
	Build() (SysexCommandExtendedId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandExtendedId
}

// NewSysexCommandExtendedIdBuilder() creates a SysexCommandExtendedIdBuilder
func NewSysexCommandExtendedIdBuilder() SysexCommandExtendedIdBuilder {
	return &_SysexCommandExtendedIdBuilder{_SysexCommandExtendedId: new(_SysexCommandExtendedId)}
}

type _SysexCommandExtendedIdBuilder struct {
	*_SysexCommandExtendedId

	parentBuilder *_SysexCommandBuilder

	err *utils.MultiError
}

var _ (SysexCommandExtendedIdBuilder) = (*_SysexCommandExtendedIdBuilder)(nil)

func (b *_SysexCommandExtendedIdBuilder) setParent(contract SysexCommandContract) {
	b.SysexCommandContract = contract
	contract.(*_SysexCommand)._SubType = b._SysexCommandExtendedId
}

func (b *_SysexCommandExtendedIdBuilder) WithMandatoryFields(id []int8) SysexCommandExtendedIdBuilder {
	return b.WithId(id...)
}

func (b *_SysexCommandExtendedIdBuilder) WithId(id ...int8) SysexCommandExtendedIdBuilder {
	b.Id = id
	return b
}

func (b *_SysexCommandExtendedIdBuilder) Build() (SysexCommandExtendedId, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommandExtendedId.deepCopy(), nil
}

func (b *_SysexCommandExtendedIdBuilder) MustBuild() SysexCommandExtendedId {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SysexCommandExtendedIdBuilder) Done() SysexCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSysexCommandBuilder().(*_SysexCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_SysexCommandExtendedIdBuilder) buildForSysexCommand() (SysexCommand, error) {
	return b.Build()
}

func (b *_SysexCommandExtendedIdBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandExtendedIdBuilder().(*_SysexCommandExtendedIdBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandExtendedIdBuilder creates a SysexCommandExtendedIdBuilder
func (b *_SysexCommandExtendedId) CreateSysexCommandExtendedIdBuilder() SysexCommandExtendedIdBuilder {
	if b == nil {
		return NewSysexCommandExtendedIdBuilder()
	}
	return &_SysexCommandExtendedIdBuilder{_SysexCommandExtendedId: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandExtendedId) GetCommandType() uint8 {
	return 0x00
}

func (m *_SysexCommandExtendedId) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandExtendedId) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandExtendedId) GetId() []int8 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSysexCommandExtendedId(structType any) SysexCommandExtendedId {
	if casted, ok := structType.(SysexCommandExtendedId); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandExtendedId); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandExtendedId) GetTypeName() string {
	return "SysexCommandExtendedId"
}

func (m *_SysexCommandExtendedId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	// Array field
	if len(m.Id) > 0 {
		lengthInBits += 8 * uint16(len(m.Id))
	}

	return lengthInBits
}

func (m *_SysexCommandExtendedId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandExtendedId) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandExtendedId SysexCommandExtendedId, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandExtendedId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandExtendedId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	id, err := ReadCountArrayField[int8](ctx, "id", ReadSignedByte(readBuffer, uint8(8)), uint64(int32(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	if closeErr := readBuffer.CloseContext("SysexCommandExtendedId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandExtendedId")
	}

	return m, nil
}

func (m *_SysexCommandExtendedId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandExtendedId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandExtendedId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandExtendedId")
		}

		if err := WriteSimpleTypeArrayField(ctx, "id", m.GetId(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandExtendedId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandExtendedId")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandExtendedId) IsSysexCommandExtendedId() {}

func (m *_SysexCommandExtendedId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandExtendedId) deepCopy() *_SysexCommandExtendedId {
	if m == nil {
		return nil
	}
	_SysexCommandExtendedIdCopy := &_SysexCommandExtendedId{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
		utils.DeepCopySlice[int8, int8](m.Id),
	}
	_SysexCommandExtendedIdCopy.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandExtendedIdCopy
}

func (m *_SysexCommandExtendedId) String() string {
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
