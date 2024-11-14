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

// SALDataTemperatureBroadcast is the corresponding interface of SALDataTemperatureBroadcast
type SALDataTemperatureBroadcast interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetTemperatureBroadcastData returns TemperatureBroadcastData (property field)
	GetTemperatureBroadcastData() TemperatureBroadcastData
	// IsSALDataTemperatureBroadcast is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataTemperatureBroadcast()
	// CreateBuilder creates a SALDataTemperatureBroadcastBuilder
	CreateSALDataTemperatureBroadcastBuilder() SALDataTemperatureBroadcastBuilder
}

// _SALDataTemperatureBroadcast is the data-structure of this message
type _SALDataTemperatureBroadcast struct {
	SALDataContract
	TemperatureBroadcastData TemperatureBroadcastData
}

var _ SALDataTemperatureBroadcast = (*_SALDataTemperatureBroadcast)(nil)
var _ SALDataRequirements = (*_SALDataTemperatureBroadcast)(nil)

// NewSALDataTemperatureBroadcast factory function for _SALDataTemperatureBroadcast
func NewSALDataTemperatureBroadcast(salData SALData, temperatureBroadcastData TemperatureBroadcastData) *_SALDataTemperatureBroadcast {
	if temperatureBroadcastData == nil {
		panic("temperatureBroadcastData of type TemperatureBroadcastData for SALDataTemperatureBroadcast must not be nil")
	}
	_result := &_SALDataTemperatureBroadcast{
		SALDataContract:          NewSALData(salData),
		TemperatureBroadcastData: temperatureBroadcastData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataTemperatureBroadcastBuilder is a builder for SALDataTemperatureBroadcast
type SALDataTemperatureBroadcastBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(temperatureBroadcastData TemperatureBroadcastData) SALDataTemperatureBroadcastBuilder
	// WithTemperatureBroadcastData adds TemperatureBroadcastData (property field)
	WithTemperatureBroadcastData(TemperatureBroadcastData) SALDataTemperatureBroadcastBuilder
	// WithTemperatureBroadcastDataBuilder adds TemperatureBroadcastData (property field) which is build by the builder
	WithTemperatureBroadcastDataBuilder(func(TemperatureBroadcastDataBuilder) TemperatureBroadcastDataBuilder) SALDataTemperatureBroadcastBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SALDataBuilder
	// Build builds the SALDataTemperatureBroadcast or returns an error if something is wrong
	Build() (SALDataTemperatureBroadcast, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataTemperatureBroadcast
}

// NewSALDataTemperatureBroadcastBuilder() creates a SALDataTemperatureBroadcastBuilder
func NewSALDataTemperatureBroadcastBuilder() SALDataTemperatureBroadcastBuilder {
	return &_SALDataTemperatureBroadcastBuilder{_SALDataTemperatureBroadcast: new(_SALDataTemperatureBroadcast)}
}

type _SALDataTemperatureBroadcastBuilder struct {
	*_SALDataTemperatureBroadcast

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataTemperatureBroadcastBuilder) = (*_SALDataTemperatureBroadcastBuilder)(nil)

func (b *_SALDataTemperatureBroadcastBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
	contract.(*_SALData)._SubType = b._SALDataTemperatureBroadcast
}

func (b *_SALDataTemperatureBroadcastBuilder) WithMandatoryFields(temperatureBroadcastData TemperatureBroadcastData) SALDataTemperatureBroadcastBuilder {
	return b.WithTemperatureBroadcastData(temperatureBroadcastData)
}

func (b *_SALDataTemperatureBroadcastBuilder) WithTemperatureBroadcastData(temperatureBroadcastData TemperatureBroadcastData) SALDataTemperatureBroadcastBuilder {
	b.TemperatureBroadcastData = temperatureBroadcastData
	return b
}

func (b *_SALDataTemperatureBroadcastBuilder) WithTemperatureBroadcastDataBuilder(builderSupplier func(TemperatureBroadcastDataBuilder) TemperatureBroadcastDataBuilder) SALDataTemperatureBroadcastBuilder {
	builder := builderSupplier(b.TemperatureBroadcastData.CreateTemperatureBroadcastDataBuilder())
	var err error
	b.TemperatureBroadcastData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "TemperatureBroadcastDataBuilder failed"))
	}
	return b
}

func (b *_SALDataTemperatureBroadcastBuilder) Build() (SALDataTemperatureBroadcast, error) {
	if b.TemperatureBroadcastData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'temperatureBroadcastData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataTemperatureBroadcast.deepCopy(), nil
}

func (b *_SALDataTemperatureBroadcastBuilder) MustBuild() SALDataTemperatureBroadcast {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SALDataTemperatureBroadcastBuilder) Done() SALDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSALDataBuilder().(*_SALDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SALDataTemperatureBroadcastBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataTemperatureBroadcastBuilder) DeepCopy() any {
	_copy := b.CreateSALDataTemperatureBroadcastBuilder().(*_SALDataTemperatureBroadcastBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataTemperatureBroadcastBuilder creates a SALDataTemperatureBroadcastBuilder
func (b *_SALDataTemperatureBroadcast) CreateSALDataTemperatureBroadcastBuilder() SALDataTemperatureBroadcastBuilder {
	if b == nil {
		return NewSALDataTemperatureBroadcastBuilder()
	}
	return &_SALDataTemperatureBroadcastBuilder{_SALDataTemperatureBroadcast: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTemperatureBroadcast) GetApplicationId() ApplicationId {
	return ApplicationId_TEMPERATURE_BROADCAST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTemperatureBroadcast) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTemperatureBroadcast) GetTemperatureBroadcastData() TemperatureBroadcastData {
	return m.TemperatureBroadcastData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataTemperatureBroadcast(structType any) SALDataTemperatureBroadcast {
	if casted, ok := structType.(SALDataTemperatureBroadcast); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTemperatureBroadcast); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTemperatureBroadcast) GetTypeName() string {
	return "SALDataTemperatureBroadcast"
}

func (m *_SALDataTemperatureBroadcast) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (temperatureBroadcastData)
	lengthInBits += m.TemperatureBroadcastData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTemperatureBroadcast) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataTemperatureBroadcast) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataTemperatureBroadcast SALDataTemperatureBroadcast, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTemperatureBroadcast"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTemperatureBroadcast")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	temperatureBroadcastData, err := ReadSimpleField[TemperatureBroadcastData](ctx, "temperatureBroadcastData", ReadComplex[TemperatureBroadcastData](TemperatureBroadcastDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'temperatureBroadcastData' field"))
	}
	m.TemperatureBroadcastData = temperatureBroadcastData

	if closeErr := readBuffer.CloseContext("SALDataTemperatureBroadcast"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTemperatureBroadcast")
	}

	return m, nil
}

func (m *_SALDataTemperatureBroadcast) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTemperatureBroadcast) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTemperatureBroadcast"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTemperatureBroadcast")
		}

		if err := WriteSimpleField[TemperatureBroadcastData](ctx, "temperatureBroadcastData", m.GetTemperatureBroadcastData(), WriteComplex[TemperatureBroadcastData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'temperatureBroadcastData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTemperatureBroadcast"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTemperatureBroadcast")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTemperatureBroadcast) IsSALDataTemperatureBroadcast() {}

func (m *_SALDataTemperatureBroadcast) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataTemperatureBroadcast) deepCopy() *_SALDataTemperatureBroadcast {
	if m == nil {
		return nil
	}
	_SALDataTemperatureBroadcastCopy := &_SALDataTemperatureBroadcast{
		m.SALDataContract.(*_SALData).deepCopy(),
		utils.DeepCopy[TemperatureBroadcastData](m.TemperatureBroadcastData),
	}
	_SALDataTemperatureBroadcastCopy.SALDataContract.(*_SALData)._SubType = m
	return _SALDataTemperatureBroadcastCopy
}

func (m *_SALDataTemperatureBroadcast) String() string {
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
