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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MResetReq is the corresponding interface of MResetReq
type MResetReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMResetReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMResetReq()
	// CreateBuilder creates a MResetReqBuilder
	CreateMResetReqBuilder() MResetReqBuilder
}

// _MResetReq is the data-structure of this message
type _MResetReq struct {
	CEMIContract
}

var _ MResetReq = (*_MResetReq)(nil)
var _ CEMIRequirements = (*_MResetReq)(nil)

// NewMResetReq factory function for _MResetReq
func NewMResetReq(size uint16) *_MResetReq {
	_result := &_MResetReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MResetReqBuilder is a builder for MResetReq
type MResetReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MResetReqBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the MResetReq or returns an error if something is wrong
	Build() (MResetReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MResetReq
}

// NewMResetReqBuilder() creates a MResetReqBuilder
func NewMResetReqBuilder() MResetReqBuilder {
	return &_MResetReqBuilder{_MResetReq: new(_MResetReq)}
}

type _MResetReqBuilder struct {
	*_MResetReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MResetReqBuilder) = (*_MResetReqBuilder)(nil)

func (b *_MResetReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._MResetReq
}

func (b *_MResetReqBuilder) WithMandatoryFields() MResetReqBuilder {
	return b
}

func (b *_MResetReqBuilder) Build() (MResetReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MResetReq.deepCopy(), nil
}

func (b *_MResetReqBuilder) MustBuild() MResetReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MResetReqBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_MResetReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MResetReqBuilder) DeepCopy() any {
	_copy := b.CreateMResetReqBuilder().(*_MResetReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMResetReqBuilder creates a MResetReqBuilder
func (b *_MResetReq) CreateMResetReqBuilder() MResetReqBuilder {
	if b == nil {
		return NewMResetReqBuilder()
	}
	return &_MResetReqBuilder{_MResetReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MResetReq) GetMessageCode() uint8 {
	return 0xF1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MResetReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMResetReq(structType any) MResetReq {
	if casted, ok := structType.(MResetReq); ok {
		return casted
	}
	if casted, ok := structType.(*MResetReq); ok {
		return *casted
	}
	return nil
}

func (m *_MResetReq) GetTypeName() string {
	return "MResetReq"
}

func (m *_MResetReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MResetReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MResetReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mResetReq MResetReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MResetReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MResetReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MResetReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MResetReq")
	}

	return m, nil
}

func (m *_MResetReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MResetReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MResetReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MResetReq")
		}

		if popErr := writeBuffer.PopContext("MResetReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MResetReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MResetReq) IsMResetReq() {}

func (m *_MResetReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MResetReq) deepCopy() *_MResetReq {
	if m == nil {
		return nil
	}
	_MResetReqCopy := &_MResetReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	_MResetReqCopy.CEMIContract.(*_CEMI)._SubType = m
	return _MResetReqCopy
}

func (m *_MResetReq) String() string {
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
