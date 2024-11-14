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

// BACnetChannelValueBoolean is the corresponding interface of BACnetChannelValueBoolean
type BACnetChannelValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// IsBACnetChannelValueBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueBoolean()
	// CreateBuilder creates a BACnetChannelValueBooleanBuilder
	CreateBACnetChannelValueBooleanBuilder() BACnetChannelValueBooleanBuilder
}

// _BACnetChannelValueBoolean is the data-structure of this message
type _BACnetChannelValueBoolean struct {
	BACnetChannelValueContract
	BooleanValue BACnetApplicationTagBoolean
}

var _ BACnetChannelValueBoolean = (*_BACnetChannelValueBoolean)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueBoolean)(nil)

// NewBACnetChannelValueBoolean factory function for _BACnetChannelValueBoolean
func NewBACnetChannelValueBoolean(peekedTagHeader BACnetTagHeader, booleanValue BACnetApplicationTagBoolean) *_BACnetChannelValueBoolean {
	if booleanValue == nil {
		panic("booleanValue of type BACnetApplicationTagBoolean for BACnetChannelValueBoolean must not be nil")
	}
	_result := &_BACnetChannelValueBoolean{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		BooleanValue:               booleanValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueBooleanBuilder is a builder for BACnetChannelValueBoolean
type BACnetChannelValueBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetChannelValueBooleanBuilder
	// WithBooleanValue adds BooleanValue (property field)
	WithBooleanValue(BACnetApplicationTagBoolean) BACnetChannelValueBooleanBuilder
	// WithBooleanValueBuilder adds BooleanValue (property field) which is build by the builder
	WithBooleanValueBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetChannelValueBooleanBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetChannelValueBuilder
	// Build builds the BACnetChannelValueBoolean or returns an error if something is wrong
	Build() (BACnetChannelValueBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueBoolean
}

// NewBACnetChannelValueBooleanBuilder() creates a BACnetChannelValueBooleanBuilder
func NewBACnetChannelValueBooleanBuilder() BACnetChannelValueBooleanBuilder {
	return &_BACnetChannelValueBooleanBuilder{_BACnetChannelValueBoolean: new(_BACnetChannelValueBoolean)}
}

type _BACnetChannelValueBooleanBuilder struct {
	*_BACnetChannelValueBoolean

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueBooleanBuilder) = (*_BACnetChannelValueBooleanBuilder)(nil)

func (b *_BACnetChannelValueBooleanBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
	contract.(*_BACnetChannelValue)._SubType = b._BACnetChannelValueBoolean
}

func (b *_BACnetChannelValueBooleanBuilder) WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetChannelValueBooleanBuilder {
	return b.WithBooleanValue(booleanValue)
}

func (b *_BACnetChannelValueBooleanBuilder) WithBooleanValue(booleanValue BACnetApplicationTagBoolean) BACnetChannelValueBooleanBuilder {
	b.BooleanValue = booleanValue
	return b
}

func (b *_BACnetChannelValueBooleanBuilder) WithBooleanValueBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetChannelValueBooleanBuilder {
	builder := builderSupplier(b.BooleanValue.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.BooleanValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetChannelValueBooleanBuilder) Build() (BACnetChannelValueBoolean, error) {
	if b.BooleanValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'booleanValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueBoolean.deepCopy(), nil
}

func (b *_BACnetChannelValueBooleanBuilder) MustBuild() BACnetChannelValueBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetChannelValueBooleanBuilder) Done() BACnetChannelValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetChannelValueBuilder().(*_BACnetChannelValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetChannelValueBooleanBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueBooleanBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueBooleanBuilder().(*_BACnetChannelValueBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueBooleanBuilder creates a BACnetChannelValueBooleanBuilder
func (b *_BACnetChannelValueBoolean) CreateBACnetChannelValueBooleanBuilder() BACnetChannelValueBooleanBuilder {
	if b == nil {
		return NewBACnetChannelValueBooleanBuilder()
	}
	return &_BACnetChannelValueBooleanBuilder{_BACnetChannelValueBoolean: b.deepCopy()}
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

func (m *_BACnetChannelValueBoolean) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueBoolean(structType any) BACnetChannelValueBoolean {
	if casted, ok := structType.(BACnetChannelValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueBoolean) GetTypeName() string {
	return "BACnetChannelValueBoolean"
}

func (m *_BACnetChannelValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueBoolean BACnetChannelValueBoolean, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueBoolean")
	}

	return m, nil
}

func (m *_BACnetChannelValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueBoolean")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueBoolean")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueBoolean) IsBACnetChannelValueBoolean() {}

func (m *_BACnetChannelValueBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueBoolean) deepCopy() *_BACnetChannelValueBoolean {
	if m == nil {
		return nil
	}
	_BACnetChannelValueBooleanCopy := &_BACnetChannelValueBoolean{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.BooleanValue),
	}
	_BACnetChannelValueBooleanCopy.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueBooleanCopy
}

func (m *_BACnetChannelValueBoolean) String() string {
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
