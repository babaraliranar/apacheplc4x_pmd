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

// BACnetApplicationTagDouble is the corresponding interface of BACnetApplicationTagDouble
type BACnetApplicationTagDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float64
	// IsBACnetApplicationTagDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagDouble()
	// CreateBuilder creates a BACnetApplicationTagDoubleBuilder
	CreateBACnetApplicationTagDoubleBuilder() BACnetApplicationTagDoubleBuilder
}

// _BACnetApplicationTagDouble is the data-structure of this message
type _BACnetApplicationTagDouble struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadDouble
}

var _ BACnetApplicationTagDouble = (*_BACnetApplicationTagDouble)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagDouble)(nil)

// NewBACnetApplicationTagDouble factory function for _BACnetApplicationTagDouble
func NewBACnetApplicationTagDouble(header BACnetTagHeader, payload BACnetTagPayloadDouble) *_BACnetApplicationTagDouble {
	if payload == nil {
		panic("payload of type BACnetTagPayloadDouble for BACnetApplicationTagDouble must not be nil")
	}
	_result := &_BACnetApplicationTagDouble{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagDoubleBuilder is a builder for BACnetApplicationTagDouble
type BACnetApplicationTagDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadDoubleBuilder) BACnetTagPayloadDoubleBuilder) BACnetApplicationTagDoubleBuilder
	// Build builds the BACnetApplicationTagDouble or returns an error if something is wrong
	Build() (BACnetApplicationTagDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagDouble
}

// NewBACnetApplicationTagDoubleBuilder() creates a BACnetApplicationTagDoubleBuilder
func NewBACnetApplicationTagDoubleBuilder() BACnetApplicationTagDoubleBuilder {
	return &_BACnetApplicationTagDoubleBuilder{_BACnetApplicationTagDouble: new(_BACnetApplicationTagDouble)}
}

type _BACnetApplicationTagDoubleBuilder struct {
	*_BACnetApplicationTagDouble

	err *utils.MultiError
}

var _ (BACnetApplicationTagDoubleBuilder) = (*_BACnetApplicationTagDoubleBuilder)(nil)

func (m *_BACnetApplicationTagDoubleBuilder) WithMandatoryFields(payload BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder {
	return m.WithPayload(payload)
}

func (m *_BACnetApplicationTagDoubleBuilder) WithPayload(payload BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder {
	m.Payload = payload
	return m
}

func (m *_BACnetApplicationTagDoubleBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadDoubleBuilder) BACnetTagPayloadDoubleBuilder) BACnetApplicationTagDoubleBuilder {
	builder := builderSupplier(m.Payload.CreateBACnetTagPayloadDoubleBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagPayloadDoubleBuilder failed"))
	}
	return m
}

func (m *_BACnetApplicationTagDoubleBuilder) Build() (BACnetApplicationTagDouble, error) {
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetApplicationTagDouble.deepCopy(), nil
}

func (m *_BACnetApplicationTagDoubleBuilder) MustBuild() BACnetApplicationTagDouble {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetApplicationTagDoubleBuilder) DeepCopy() any {
	return m.CreateBACnetApplicationTagDoubleBuilder()
}

// CreateBACnetApplicationTagDoubleBuilder creates a BACnetApplicationTagDoubleBuilder
func (m *_BACnetApplicationTagDouble) CreateBACnetApplicationTagDoubleBuilder() BACnetApplicationTagDoubleBuilder {
	if m == nil {
		return NewBACnetApplicationTagDoubleBuilder()
	}
	return &_BACnetApplicationTagDoubleBuilder{_BACnetApplicationTagDouble: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagDouble) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagDouble) GetPayload() BACnetTagPayloadDouble {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagDouble) GetActualValue() float64 {
	ctx := context.Background()
	_ = ctx
	return float64(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagDouble(structType any) BACnetApplicationTagDouble {
	if casted, ok := structType.(BACnetApplicationTagDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagDouble) GetTypeName() string {
	return "BACnetApplicationTagDouble"
}

func (m *_BACnetApplicationTagDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag) (__bACnetApplicationTagDouble BACnetApplicationTagDouble, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadDouble](ctx, "payload", ReadComplex[BACnetTagPayloadDouble](BACnetTagPayloadDoubleParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[float64](ctx, "actualValue", (*float64)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagDouble")
	}

	return m, nil
}

func (m *_BACnetApplicationTagDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagDouble")
		}

		if err := WriteSimpleField[BACnetTagPayloadDouble](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagDouble")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagDouble) IsBACnetApplicationTagDouble() {}

func (m *_BACnetApplicationTagDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagDouble) deepCopy() *_BACnetApplicationTagDouble {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagDoubleCopy := &_BACnetApplicationTagDouble{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadDouble),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagDoubleCopy
}

func (m *_BACnetApplicationTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
