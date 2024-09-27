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

// BACnetContextTagNull is the corresponding interface of BACnetContextTagNull
type BACnetContextTagNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// IsBACnetContextTagNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagNull()
	// CreateBuilder creates a BACnetContextTagNullBuilder
	CreateBACnetContextTagNullBuilder() BACnetContextTagNullBuilder
}

// _BACnetContextTagNull is the data-structure of this message
type _BACnetContextTagNull struct {
	BACnetContextTagContract
}

var _ BACnetContextTagNull = (*_BACnetContextTagNull)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagNull)(nil)

// NewBACnetContextTagNull factory function for _BACnetContextTagNull
func NewBACnetContextTagNull(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagNull {
	_result := &_BACnetContextTagNull{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagNullBuilder is a builder for BACnetContextTagNull
type BACnetContextTagNullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetContextTagNullBuilder
	// Build builds the BACnetContextTagNull or returns an error if something is wrong
	Build() (BACnetContextTagNull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagNull
}

// NewBACnetContextTagNullBuilder() creates a BACnetContextTagNullBuilder
func NewBACnetContextTagNullBuilder() BACnetContextTagNullBuilder {
	return &_BACnetContextTagNullBuilder{_BACnetContextTagNull: new(_BACnetContextTagNull)}
}

type _BACnetContextTagNullBuilder struct {
	*_BACnetContextTagNull

	err *utils.MultiError
}

var _ (BACnetContextTagNullBuilder) = (*_BACnetContextTagNullBuilder)(nil)

func (m *_BACnetContextTagNullBuilder) WithMandatoryFields() BACnetContextTagNullBuilder {
	return m
}

func (m *_BACnetContextTagNullBuilder) Build() (BACnetContextTagNull, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetContextTagNull.deepCopy(), nil
}

func (m *_BACnetContextTagNullBuilder) MustBuild() BACnetContextTagNull {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetContextTagNullBuilder) DeepCopy() any {
	return m.CreateBACnetContextTagNullBuilder()
}

// CreateBACnetContextTagNullBuilder creates a BACnetContextTagNullBuilder
func (m *_BACnetContextTagNull) CreateBACnetContextTagNullBuilder() BACnetContextTagNullBuilder {
	if m == nil {
		return NewBACnetContextTagNullBuilder()
	}
	return &_BACnetContextTagNullBuilder{_BACnetContextTagNull: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagNull) GetDataType() BACnetDataType {
	return BACnetDataType_NULL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagNull) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagNull(structType any) BACnetContextTagNull {
	if casted, ok := structType.(BACnetContextTagNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagNull) GetTypeName() string {
	return "BACnetContextTagNull"
}

func (m *_BACnetContextTagNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetContextTagNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagNull BACnetContextTagNull, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((header.GetActualLength()) == (0))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "length field should be 0"})
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagNull")
	}

	return m, nil
}

func (m *_BACnetContextTagNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagNull")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagNull")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagNull) IsBACnetContextTagNull() {}

func (m *_BACnetContextTagNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagNull) deepCopy() *_BACnetContextTagNull {
	if m == nil {
		return nil
	}
	_BACnetContextTagNullCopy := &_BACnetContextTagNull{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagNullCopy
}

func (m *_BACnetContextTagNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
