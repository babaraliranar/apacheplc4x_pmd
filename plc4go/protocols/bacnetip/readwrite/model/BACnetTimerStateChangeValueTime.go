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

// BACnetTimerStateChangeValueTime is the corresponding interface of BACnetTimerStateChangeValueTime
type BACnetTimerStateChangeValueTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// IsBACnetTimerStateChangeValueTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueTime()
	// CreateBuilder creates a BACnetTimerStateChangeValueTimeBuilder
	CreateBACnetTimerStateChangeValueTimeBuilder() BACnetTimerStateChangeValueTimeBuilder
}

// _BACnetTimerStateChangeValueTime is the data-structure of this message
type _BACnetTimerStateChangeValueTime struct {
	BACnetTimerStateChangeValueContract
	TimeValue BACnetApplicationTagTime
}

var _ BACnetTimerStateChangeValueTime = (*_BACnetTimerStateChangeValueTime)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueTime)(nil)

// NewBACnetTimerStateChangeValueTime factory function for _BACnetTimerStateChangeValueTime
func NewBACnetTimerStateChangeValueTime(peekedTagHeader BACnetTagHeader, timeValue BACnetApplicationTagTime, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueTime {
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetTimerStateChangeValueTime must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueTime{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		TimeValue:                           timeValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueTimeBuilder is a builder for BACnetTimerStateChangeValueTime
type BACnetTimerStateChangeValueTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetTimerStateChangeValueTimeBuilder
	// WithTimeValue adds TimeValue (property field)
	WithTimeValue(BACnetApplicationTagTime) BACnetTimerStateChangeValueTimeBuilder
	// WithTimeValueBuilder adds TimeValue (property field) which is build by the builder
	WithTimeValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetTimerStateChangeValueTimeBuilder
	// Build builds the BACnetTimerStateChangeValueTime or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueTime
}

// NewBACnetTimerStateChangeValueTimeBuilder() creates a BACnetTimerStateChangeValueTimeBuilder
func NewBACnetTimerStateChangeValueTimeBuilder() BACnetTimerStateChangeValueTimeBuilder {
	return &_BACnetTimerStateChangeValueTimeBuilder{_BACnetTimerStateChangeValueTime: new(_BACnetTimerStateChangeValueTime)}
}

type _BACnetTimerStateChangeValueTimeBuilder struct {
	*_BACnetTimerStateChangeValueTime

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueTimeBuilder) = (*_BACnetTimerStateChangeValueTimeBuilder)(nil)

func (b *_BACnetTimerStateChangeValueTimeBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetTimerStateChangeValueTimeBuilder {
	return b.WithTimeValue(timeValue)
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) WithTimeValue(timeValue BACnetApplicationTagTime) BACnetTimerStateChangeValueTimeBuilder {
	b.TimeValue = timeValue
	return b
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) WithTimeValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetTimerStateChangeValueTimeBuilder {
	builder := builderSupplier(b.TimeValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.TimeValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) Build() (BACnetTimerStateChangeValueTime, error) {
	if b.TimeValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueTime.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) MustBuild() BACnetTimerStateChangeValueTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetTimerStateChangeValueTimeBuilder) Done() BACnetTimerStateChangeValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueTimeBuilder().(*_BACnetTimerStateChangeValueTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueTimeBuilder creates a BACnetTimerStateChangeValueTimeBuilder
func (b *_BACnetTimerStateChangeValueTime) CreateBACnetTimerStateChangeValueTimeBuilder() BACnetTimerStateChangeValueTimeBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueTimeBuilder()
	}
	return &_BACnetTimerStateChangeValueTimeBuilder{_BACnetTimerStateChangeValueTime: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueTime) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueTime(structType any) BACnetTimerStateChangeValueTime {
	if casted, ok := structType.(BACnetTimerStateChangeValueTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueTime) GetTypeName() string {
	return "BACnetTimerStateChangeValueTime"
}

func (m *_BACnetTimerStateChangeValueTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueTime BACnetTimerStateChangeValueTime, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueTime")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueTime")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueTime) IsBACnetTimerStateChangeValueTime() {}

func (m *_BACnetTimerStateChangeValueTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueTime) deepCopy() *_BACnetTimerStateChangeValueTime {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueTimeCopy := &_BACnetTimerStateChangeValueTime{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		m.TimeValue.DeepCopy().(BACnetApplicationTagTime),
	}
	m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueTimeCopy
}

func (m *_BACnetTimerStateChangeValueTime) String() string {
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
