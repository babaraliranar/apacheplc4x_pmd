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

// PortSegment is the corresponding interface of PortSegment
type PortSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	PathSegment
	// GetSegmentType returns SegmentType (property field)
	GetSegmentType() PortSegmentType
	// IsPortSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortSegment()
	// CreateBuilder creates a PortSegmentBuilder
	CreatePortSegmentBuilder() PortSegmentBuilder
}

// _PortSegment is the data-structure of this message
type _PortSegment struct {
	PathSegmentContract
	SegmentType PortSegmentType
}

var _ PortSegment = (*_PortSegment)(nil)
var _ PathSegmentRequirements = (*_PortSegment)(nil)

// NewPortSegment factory function for _PortSegment
func NewPortSegment(segmentType PortSegmentType) *_PortSegment {
	if segmentType == nil {
		panic("segmentType of type PortSegmentType for PortSegment must not be nil")
	}
	_result := &_PortSegment{
		PathSegmentContract: NewPathSegment(),
		SegmentType:         segmentType,
	}
	_result.PathSegmentContract.(*_PathSegment)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PortSegmentBuilder is a builder for PortSegment
type PortSegmentBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(segmentType PortSegmentType) PortSegmentBuilder
	// WithSegmentType adds SegmentType (property field)
	WithSegmentType(PortSegmentType) PortSegmentBuilder
	// WithSegmentTypeBuilder adds SegmentType (property field) which is build by the builder
	WithSegmentTypeBuilder(func(PortSegmentTypeBuilder) PortSegmentTypeBuilder) PortSegmentBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() PathSegmentBuilder
	// Build builds the PortSegment or returns an error if something is wrong
	Build() (PortSegment, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PortSegment
}

// NewPortSegmentBuilder() creates a PortSegmentBuilder
func NewPortSegmentBuilder() PortSegmentBuilder {
	return &_PortSegmentBuilder{_PortSegment: new(_PortSegment)}
}

type _PortSegmentBuilder struct {
	*_PortSegment

	parentBuilder *_PathSegmentBuilder

	err *utils.MultiError
}

var _ (PortSegmentBuilder) = (*_PortSegmentBuilder)(nil)

func (b *_PortSegmentBuilder) setParent(contract PathSegmentContract) {
	b.PathSegmentContract = contract
	contract.(*_PathSegment)._SubType = b._PortSegment
}

func (b *_PortSegmentBuilder) WithMandatoryFields(segmentType PortSegmentType) PortSegmentBuilder {
	return b.WithSegmentType(segmentType)
}

func (b *_PortSegmentBuilder) WithSegmentType(segmentType PortSegmentType) PortSegmentBuilder {
	b.SegmentType = segmentType
	return b
}

func (b *_PortSegmentBuilder) WithSegmentTypeBuilder(builderSupplier func(PortSegmentTypeBuilder) PortSegmentTypeBuilder) PortSegmentBuilder {
	builder := builderSupplier(b.SegmentType.CreatePortSegmentTypeBuilder())
	var err error
	b.SegmentType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PortSegmentTypeBuilder failed"))
	}
	return b
}

func (b *_PortSegmentBuilder) Build() (PortSegment, error) {
	if b.SegmentType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'segmentType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PortSegment.deepCopy(), nil
}

func (b *_PortSegmentBuilder) MustBuild() PortSegment {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PortSegmentBuilder) Done() PathSegmentBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewPathSegmentBuilder().(*_PathSegmentBuilder)
	}
	return b.parentBuilder
}

func (b *_PortSegmentBuilder) buildForPathSegment() (PathSegment, error) {
	return b.Build()
}

func (b *_PortSegmentBuilder) DeepCopy() any {
	_copy := b.CreatePortSegmentBuilder().(*_PortSegmentBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePortSegmentBuilder creates a PortSegmentBuilder
func (b *_PortSegment) CreatePortSegmentBuilder() PortSegmentBuilder {
	if b == nil {
		return NewPortSegmentBuilder()
	}
	return &_PortSegmentBuilder{_PortSegment: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegment) GetPathSegment() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegment) GetParent() PathSegmentContract {
	return m.PathSegmentContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegment) GetSegmentType() PortSegmentType {
	return m.SegmentType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPortSegment(structType any) PortSegment {
	if casted, ok := structType.(PortSegment); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegment); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegment) GetTypeName() string {
	return "PortSegment"
}

func (m *_PortSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PathSegmentContract.(*_PathSegment).getLengthInBits(ctx))

	// Simple field (segmentType)
	lengthInBits += m.SegmentType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PortSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PathSegment) (__portSegment PortSegment, err error) {
	m.PathSegmentContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentType, err := ReadSimpleField[PortSegmentType](ctx, "segmentType", ReadComplex[PortSegmentType](PortSegmentTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentType' field"))
	}
	m.SegmentType = segmentType

	if closeErr := readBuffer.CloseContext("PortSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegment")
	}

	return m, nil
}

func (m *_PortSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegment")
		}

		if err := WriteSimpleField[PortSegmentType](ctx, "segmentType", m.GetSegmentType(), WriteComplex[PortSegmentType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentType' field")
		}

		if popErr := writeBuffer.PopContext("PortSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegment")
		}
		return nil
	}
	return m.PathSegmentContract.(*_PathSegment).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegment) IsPortSegment() {}

func (m *_PortSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PortSegment) deepCopy() *_PortSegment {
	if m == nil {
		return nil
	}
	_PortSegmentCopy := &_PortSegment{
		m.PathSegmentContract.(*_PathSegment).deepCopy(),
		utils.DeepCopy[PortSegmentType](m.SegmentType),
	}
	_PortSegmentCopy.PathSegmentContract.(*_PathSegment)._SubType = m
	return _PortSegmentCopy
}

func (m *_PortSegment) String() string {
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
