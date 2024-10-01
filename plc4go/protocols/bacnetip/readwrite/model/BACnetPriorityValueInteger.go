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

// BACnetPriorityValueInteger is the corresponding interface of BACnetPriorityValueInteger
type BACnetPriorityValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// IsBACnetPriorityValueInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueInteger()
	// CreateBuilder creates a BACnetPriorityValueIntegerBuilder
	CreateBACnetPriorityValueIntegerBuilder() BACnetPriorityValueIntegerBuilder
}

// _BACnetPriorityValueInteger is the data-structure of this message
type _BACnetPriorityValueInteger struct {
	BACnetPriorityValueContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetPriorityValueInteger = (*_BACnetPriorityValueInteger)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueInteger)(nil)

// NewBACnetPriorityValueInteger factory function for _BACnetPriorityValueInteger
func NewBACnetPriorityValueInteger(peekedTagHeader BACnetTagHeader, integerValue BACnetApplicationTagSignedInteger, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueInteger {
	if integerValue == nil {
		panic("integerValue of type BACnetApplicationTagSignedInteger for BACnetPriorityValueInteger must not be nil")
	}
	_result := &_BACnetPriorityValueInteger{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		IntegerValue:                integerValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueIntegerBuilder is a builder for BACnetPriorityValueInteger
type BACnetPriorityValueIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetPriorityValueIntegerBuilder
	// WithIntegerValue adds IntegerValue (property field)
	WithIntegerValue(BACnetApplicationTagSignedInteger) BACnetPriorityValueIntegerBuilder
	// WithIntegerValueBuilder adds IntegerValue (property field) which is build by the builder
	WithIntegerValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetPriorityValueIntegerBuilder
	// Build builds the BACnetPriorityValueInteger or returns an error if something is wrong
	Build() (BACnetPriorityValueInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueInteger
}

// NewBACnetPriorityValueIntegerBuilder() creates a BACnetPriorityValueIntegerBuilder
func NewBACnetPriorityValueIntegerBuilder() BACnetPriorityValueIntegerBuilder {
	return &_BACnetPriorityValueIntegerBuilder{_BACnetPriorityValueInteger: new(_BACnetPriorityValueInteger)}
}

type _BACnetPriorityValueIntegerBuilder struct {
	*_BACnetPriorityValueInteger

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueIntegerBuilder) = (*_BACnetPriorityValueIntegerBuilder)(nil)

func (b *_BACnetPriorityValueIntegerBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
}

func (b *_BACnetPriorityValueIntegerBuilder) WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetPriorityValueIntegerBuilder {
	return b.WithIntegerValue(integerValue)
}

func (b *_BACnetPriorityValueIntegerBuilder) WithIntegerValue(integerValue BACnetApplicationTagSignedInteger) BACnetPriorityValueIntegerBuilder {
	b.IntegerValue = integerValue
	return b
}

func (b *_BACnetPriorityValueIntegerBuilder) WithIntegerValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetPriorityValueIntegerBuilder {
	builder := builderSupplier(b.IntegerValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.IntegerValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetPriorityValueIntegerBuilder) Build() (BACnetPriorityValueInteger, error) {
	if b.IntegerValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'integerValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueInteger.deepCopy(), nil
}

func (b *_BACnetPriorityValueIntegerBuilder) MustBuild() BACnetPriorityValueInteger {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPriorityValueIntegerBuilder) Done() BACnetPriorityValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetPriorityValueIntegerBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueIntegerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueIntegerBuilder().(*_BACnetPriorityValueIntegerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueIntegerBuilder creates a BACnetPriorityValueIntegerBuilder
func (b *_BACnetPriorityValueInteger) CreateBACnetPriorityValueIntegerBuilder() BACnetPriorityValueIntegerBuilder {
	if b == nil {
		return NewBACnetPriorityValueIntegerBuilder()
	}
	return &_BACnetPriorityValueIntegerBuilder{_BACnetPriorityValueInteger: b.deepCopy()}
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

func (m *_BACnetPriorityValueInteger) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueInteger(structType any) BACnetPriorityValueInteger {
	if casted, ok := structType.(BACnetPriorityValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueInteger) GetTypeName() string {
	return "BACnetPriorityValueInteger"
}

func (m *_BACnetPriorityValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueInteger BACnetPriorityValueInteger, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueInteger")
	}

	return m, nil
}

func (m *_BACnetPriorityValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueInteger")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueInteger) IsBACnetPriorityValueInteger() {}

func (m *_BACnetPriorityValueInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueInteger) deepCopy() *_BACnetPriorityValueInteger {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueIntegerCopy := &_BACnetPriorityValueInteger{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		m.IntegerValue.DeepCopy().(BACnetApplicationTagSignedInteger),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueIntegerCopy
}

func (m *_BACnetPriorityValueInteger) String() string {
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
