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

// BACnetChannelValueBitString is the corresponding interface of BACnetChannelValueBitString
type BACnetChannelValueBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
	// IsBACnetChannelValueBitString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueBitString()
	// CreateBuilder creates a BACnetChannelValueBitStringBuilder
	CreateBACnetChannelValueBitStringBuilder() BACnetChannelValueBitStringBuilder
}

// _BACnetChannelValueBitString is the data-structure of this message
type _BACnetChannelValueBitString struct {
	BACnetChannelValueContract
	BitStringValue BACnetApplicationTagBitString
}

var _ BACnetChannelValueBitString = (*_BACnetChannelValueBitString)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueBitString)(nil)

// NewBACnetChannelValueBitString factory function for _BACnetChannelValueBitString
func NewBACnetChannelValueBitString(peekedTagHeader BACnetTagHeader, bitStringValue BACnetApplicationTagBitString) *_BACnetChannelValueBitString {
	if bitStringValue == nil {
		panic("bitStringValue of type BACnetApplicationTagBitString for BACnetChannelValueBitString must not be nil")
	}
	_result := &_BACnetChannelValueBitString{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		BitStringValue:             bitStringValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueBitStringBuilder is a builder for BACnetChannelValueBitString
type BACnetChannelValueBitStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bitStringValue BACnetApplicationTagBitString) BACnetChannelValueBitStringBuilder
	// WithBitStringValue adds BitStringValue (property field)
	WithBitStringValue(BACnetApplicationTagBitString) BACnetChannelValueBitStringBuilder
	// WithBitStringValueBuilder adds BitStringValue (property field) which is build by the builder
	WithBitStringValueBuilder(func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetChannelValueBitStringBuilder
	// Build builds the BACnetChannelValueBitString or returns an error if something is wrong
	Build() (BACnetChannelValueBitString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueBitString
}

// NewBACnetChannelValueBitStringBuilder() creates a BACnetChannelValueBitStringBuilder
func NewBACnetChannelValueBitStringBuilder() BACnetChannelValueBitStringBuilder {
	return &_BACnetChannelValueBitStringBuilder{_BACnetChannelValueBitString: new(_BACnetChannelValueBitString)}
}

type _BACnetChannelValueBitStringBuilder struct {
	*_BACnetChannelValueBitString

	err *utils.MultiError
}

var _ (BACnetChannelValueBitStringBuilder) = (*_BACnetChannelValueBitStringBuilder)(nil)

func (m *_BACnetChannelValueBitStringBuilder) WithMandatoryFields(bitStringValue BACnetApplicationTagBitString) BACnetChannelValueBitStringBuilder {
	return m.WithBitStringValue(bitStringValue)
}

func (m *_BACnetChannelValueBitStringBuilder) WithBitStringValue(bitStringValue BACnetApplicationTagBitString) BACnetChannelValueBitStringBuilder {
	m.BitStringValue = bitStringValue
	return m
}

func (m *_BACnetChannelValueBitStringBuilder) WithBitStringValueBuilder(builderSupplier func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetChannelValueBitStringBuilder {
	builder := builderSupplier(m.BitStringValue.CreateBACnetApplicationTagBitStringBuilder())
	var err error
	m.BitStringValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBitStringBuilder failed"))
	}
	return m
}

func (m *_BACnetChannelValueBitStringBuilder) Build() (BACnetChannelValueBitString, error) {
	if m.BitStringValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'bitStringValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetChannelValueBitString.deepCopy(), nil
}

func (m *_BACnetChannelValueBitStringBuilder) MustBuild() BACnetChannelValueBitString {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetChannelValueBitStringBuilder) DeepCopy() any {
	return m.CreateBACnetChannelValueBitStringBuilder()
}

// CreateBACnetChannelValueBitStringBuilder creates a BACnetChannelValueBitStringBuilder
func (m *_BACnetChannelValueBitString) CreateBACnetChannelValueBitStringBuilder() BACnetChannelValueBitStringBuilder {
	if m == nil {
		return NewBACnetChannelValueBitStringBuilder()
	}
	return &_BACnetChannelValueBitStringBuilder{_BACnetChannelValueBitString: m.deepCopy()}
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

func (m *_BACnetChannelValueBitString) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueBitString) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueBitString(structType any) BACnetChannelValueBitString {
	if casted, ok := structType.(BACnetChannelValueBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueBitString) GetTypeName() string {
	return "BACnetChannelValueBitString"
}

func (m *_BACnetChannelValueBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueBitString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueBitString BACnetChannelValueBitString, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bitStringValue, err := ReadSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	m.BitStringValue = bitStringValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueBitString")
	}

	return m, nil
}

func (m *_BACnetChannelValueBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueBitString")
		}

		if err := WriteSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", m.GetBitStringValue(), WriteComplex[BACnetApplicationTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueBitString")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueBitString) IsBACnetChannelValueBitString() {}

func (m *_BACnetChannelValueBitString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueBitString) deepCopy() *_BACnetChannelValueBitString {
	if m == nil {
		return nil
	}
	_BACnetChannelValueBitStringCopy := &_BACnetChannelValueBitString{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		m.BitStringValue.DeepCopy().(BACnetApplicationTagBitString),
	}
	m.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueBitStringCopy
}

func (m *_BACnetChannelValueBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
