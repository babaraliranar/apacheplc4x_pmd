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

// BACnetPriorityValueReal is the corresponding interface of BACnetPriorityValueReal
type BACnetPriorityValueReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// IsBACnetPriorityValueReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueReal()
	// CreateBuilder creates a BACnetPriorityValueRealBuilder
	CreateBACnetPriorityValueRealBuilder() BACnetPriorityValueRealBuilder
}

// _BACnetPriorityValueReal is the data-structure of this message
type _BACnetPriorityValueReal struct {
	BACnetPriorityValueContract
	RealValue BACnetApplicationTagReal
}

var _ BACnetPriorityValueReal = (*_BACnetPriorityValueReal)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueReal)(nil)

// NewBACnetPriorityValueReal factory function for _BACnetPriorityValueReal
func NewBACnetPriorityValueReal(peekedTagHeader BACnetTagHeader, realValue BACnetApplicationTagReal, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueReal {
	if realValue == nil {
		panic("realValue of type BACnetApplicationTagReal for BACnetPriorityValueReal must not be nil")
	}
	_result := &_BACnetPriorityValueReal{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		RealValue:                   realValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueRealBuilder is a builder for BACnetPriorityValueReal
type BACnetPriorityValueRealBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetPriorityValueRealBuilder
	// WithRealValue adds RealValue (property field)
	WithRealValue(BACnetApplicationTagReal) BACnetPriorityValueRealBuilder
	// WithRealValueBuilder adds RealValue (property field) which is build by the builder
	WithRealValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetPriorityValueRealBuilder
	// Build builds the BACnetPriorityValueReal or returns an error if something is wrong
	Build() (BACnetPriorityValueReal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueReal
}

// NewBACnetPriorityValueRealBuilder() creates a BACnetPriorityValueRealBuilder
func NewBACnetPriorityValueRealBuilder() BACnetPriorityValueRealBuilder {
	return &_BACnetPriorityValueRealBuilder{_BACnetPriorityValueReal: new(_BACnetPriorityValueReal)}
}

type _BACnetPriorityValueRealBuilder struct {
	*_BACnetPriorityValueReal

	err *utils.MultiError
}

var _ (BACnetPriorityValueRealBuilder) = (*_BACnetPriorityValueRealBuilder)(nil)

func (m *_BACnetPriorityValueRealBuilder) WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetPriorityValueRealBuilder {
	return m.WithRealValue(realValue)
}

func (m *_BACnetPriorityValueRealBuilder) WithRealValue(realValue BACnetApplicationTagReal) BACnetPriorityValueRealBuilder {
	m.RealValue = realValue
	return m
}

func (m *_BACnetPriorityValueRealBuilder) WithRealValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetPriorityValueRealBuilder {
	builder := builderSupplier(m.RealValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.RealValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetPriorityValueRealBuilder) Build() (BACnetPriorityValueReal, error) {
	if m.RealValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'realValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPriorityValueReal.deepCopy(), nil
}

func (m *_BACnetPriorityValueRealBuilder) MustBuild() BACnetPriorityValueReal {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPriorityValueRealBuilder) DeepCopy() any {
	return m.CreateBACnetPriorityValueRealBuilder()
}

// CreateBACnetPriorityValueRealBuilder creates a BACnetPriorityValueRealBuilder
func (m *_BACnetPriorityValueReal) CreateBACnetPriorityValueRealBuilder() BACnetPriorityValueRealBuilder {
	if m == nil {
		return NewBACnetPriorityValueRealBuilder()
	}
	return &_BACnetPriorityValueRealBuilder{_BACnetPriorityValueReal: m.deepCopy()}
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

func (m *_BACnetPriorityValueReal) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueReal(structType any) BACnetPriorityValueReal {
	if casted, ok := structType.(BACnetPriorityValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueReal) GetTypeName() string {
	return "BACnetPriorityValueReal"
}

func (m *_BACnetPriorityValueReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueReal BACnetPriorityValueReal, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	m.RealValue = realValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueReal")
	}

	return m, nil
}

func (m *_BACnetPriorityValueReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueReal")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueReal")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueReal) IsBACnetPriorityValueReal() {}

func (m *_BACnetPriorityValueReal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueReal) deepCopy() *_BACnetPriorityValueReal {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueRealCopy := &_BACnetPriorityValueReal{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		m.RealValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueRealCopy
}

func (m *_BACnetPriorityValueReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
