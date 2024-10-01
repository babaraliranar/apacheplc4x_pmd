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

// MeteringDataMeasureDrinkingWater is the corresponding interface of MeteringDataMeasureDrinkingWater
type MeteringDataMeasureDrinkingWater interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MeteringData
	// IsMeteringDataMeasureDrinkingWater is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataMeasureDrinkingWater()
	// CreateBuilder creates a MeteringDataMeasureDrinkingWaterBuilder
	CreateMeteringDataMeasureDrinkingWaterBuilder() MeteringDataMeasureDrinkingWaterBuilder
}

// _MeteringDataMeasureDrinkingWater is the data-structure of this message
type _MeteringDataMeasureDrinkingWater struct {
	MeteringDataContract
}

var _ MeteringDataMeasureDrinkingWater = (*_MeteringDataMeasureDrinkingWater)(nil)
var _ MeteringDataRequirements = (*_MeteringDataMeasureDrinkingWater)(nil)

// NewMeteringDataMeasureDrinkingWater factory function for _MeteringDataMeasureDrinkingWater
func NewMeteringDataMeasureDrinkingWater(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureDrinkingWater {
	_result := &_MeteringDataMeasureDrinkingWater{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MeteringDataMeasureDrinkingWaterBuilder is a builder for MeteringDataMeasureDrinkingWater
type MeteringDataMeasureDrinkingWaterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MeteringDataMeasureDrinkingWaterBuilder
	// Build builds the MeteringDataMeasureDrinkingWater or returns an error if something is wrong
	Build() (MeteringDataMeasureDrinkingWater, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MeteringDataMeasureDrinkingWater
}

// NewMeteringDataMeasureDrinkingWaterBuilder() creates a MeteringDataMeasureDrinkingWaterBuilder
func NewMeteringDataMeasureDrinkingWaterBuilder() MeteringDataMeasureDrinkingWaterBuilder {
	return &_MeteringDataMeasureDrinkingWaterBuilder{_MeteringDataMeasureDrinkingWater: new(_MeteringDataMeasureDrinkingWater)}
}

type _MeteringDataMeasureDrinkingWaterBuilder struct {
	*_MeteringDataMeasureDrinkingWater

	parentBuilder *_MeteringDataBuilder

	err *utils.MultiError
}

var _ (MeteringDataMeasureDrinkingWaterBuilder) = (*_MeteringDataMeasureDrinkingWaterBuilder)(nil)

func (b *_MeteringDataMeasureDrinkingWaterBuilder) setParent(contract MeteringDataContract) {
	b.MeteringDataContract = contract
}

func (b *_MeteringDataMeasureDrinkingWaterBuilder) WithMandatoryFields() MeteringDataMeasureDrinkingWaterBuilder {
	return b
}

func (b *_MeteringDataMeasureDrinkingWaterBuilder) Build() (MeteringDataMeasureDrinkingWater, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MeteringDataMeasureDrinkingWater.deepCopy(), nil
}

func (b *_MeteringDataMeasureDrinkingWaterBuilder) MustBuild() MeteringDataMeasureDrinkingWater {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MeteringDataMeasureDrinkingWaterBuilder) Done() MeteringDataBuilder {
	return b.parentBuilder
}

func (b *_MeteringDataMeasureDrinkingWaterBuilder) buildForMeteringData() (MeteringData, error) {
	return b.Build()
}

func (b *_MeteringDataMeasureDrinkingWaterBuilder) DeepCopy() any {
	_copy := b.CreateMeteringDataMeasureDrinkingWaterBuilder().(*_MeteringDataMeasureDrinkingWaterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMeteringDataMeasureDrinkingWaterBuilder creates a MeteringDataMeasureDrinkingWaterBuilder
func (b *_MeteringDataMeasureDrinkingWater) CreateMeteringDataMeasureDrinkingWaterBuilder() MeteringDataMeasureDrinkingWaterBuilder {
	if b == nil {
		return NewMeteringDataMeasureDrinkingWaterBuilder()
	}
	return &_MeteringDataMeasureDrinkingWaterBuilder{_MeteringDataMeasureDrinkingWater: b.deepCopy()}
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

func (m *_MeteringDataMeasureDrinkingWater) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureDrinkingWater(structType any) MeteringDataMeasureDrinkingWater {
	if casted, ok := structType.(MeteringDataMeasureDrinkingWater); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureDrinkingWater); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureDrinkingWater) GetTypeName() string {
	return "MeteringDataMeasureDrinkingWater"
}

func (m *_MeteringDataMeasureDrinkingWater) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureDrinkingWater) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataMeasureDrinkingWater) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataMeasureDrinkingWater MeteringDataMeasureDrinkingWater, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureDrinkingWater"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureDrinkingWater")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureDrinkingWater"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureDrinkingWater")
	}

	return m, nil
}

func (m *_MeteringDataMeasureDrinkingWater) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureDrinkingWater) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureDrinkingWater"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureDrinkingWater")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureDrinkingWater"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureDrinkingWater")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureDrinkingWater) IsMeteringDataMeasureDrinkingWater() {}

func (m *_MeteringDataMeasureDrinkingWater) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataMeasureDrinkingWater) deepCopy() *_MeteringDataMeasureDrinkingWater {
	if m == nil {
		return nil
	}
	_MeteringDataMeasureDrinkingWaterCopy := &_MeteringDataMeasureDrinkingWater{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
	}
	m.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataMeasureDrinkingWaterCopy
}

func (m *_MeteringDataMeasureDrinkingWater) String() string {
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
