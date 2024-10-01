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

// MeteringDataDrinkingWaterConsumption is the corresponding interface of MeteringDataDrinkingWaterConsumption
type MeteringDataDrinkingWaterConsumption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MeteringData
	// GetKL returns KL (property field)
	GetKL() uint32
	// IsMeteringDataDrinkingWaterConsumption is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataDrinkingWaterConsumption()
	// CreateBuilder creates a MeteringDataDrinkingWaterConsumptionBuilder
	CreateMeteringDataDrinkingWaterConsumptionBuilder() MeteringDataDrinkingWaterConsumptionBuilder
}

// _MeteringDataDrinkingWaterConsumption is the data-structure of this message
type _MeteringDataDrinkingWaterConsumption struct {
	MeteringDataContract
	KL uint32
}

var _ MeteringDataDrinkingWaterConsumption = (*_MeteringDataDrinkingWaterConsumption)(nil)
var _ MeteringDataRequirements = (*_MeteringDataDrinkingWaterConsumption)(nil)

// NewMeteringDataDrinkingWaterConsumption factory function for _MeteringDataDrinkingWaterConsumption
func NewMeteringDataDrinkingWaterConsumption(commandTypeContainer MeteringCommandTypeContainer, argument byte, kL uint32) *_MeteringDataDrinkingWaterConsumption {
	_result := &_MeteringDataDrinkingWaterConsumption{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
		KL:                   kL,
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MeteringDataDrinkingWaterConsumptionBuilder is a builder for MeteringDataDrinkingWaterConsumption
type MeteringDataDrinkingWaterConsumptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(kL uint32) MeteringDataDrinkingWaterConsumptionBuilder
	// WithKL adds KL (property field)
	WithKL(uint32) MeteringDataDrinkingWaterConsumptionBuilder
	// Build builds the MeteringDataDrinkingWaterConsumption or returns an error if something is wrong
	Build() (MeteringDataDrinkingWaterConsumption, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MeteringDataDrinkingWaterConsumption
}

// NewMeteringDataDrinkingWaterConsumptionBuilder() creates a MeteringDataDrinkingWaterConsumptionBuilder
func NewMeteringDataDrinkingWaterConsumptionBuilder() MeteringDataDrinkingWaterConsumptionBuilder {
	return &_MeteringDataDrinkingWaterConsumptionBuilder{_MeteringDataDrinkingWaterConsumption: new(_MeteringDataDrinkingWaterConsumption)}
}

type _MeteringDataDrinkingWaterConsumptionBuilder struct {
	*_MeteringDataDrinkingWaterConsumption

	parentBuilder *_MeteringDataBuilder

	err *utils.MultiError
}

var _ (MeteringDataDrinkingWaterConsumptionBuilder) = (*_MeteringDataDrinkingWaterConsumptionBuilder)(nil)

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) setParent(contract MeteringDataContract) {
	b.MeteringDataContract = contract
}

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) WithMandatoryFields(kL uint32) MeteringDataDrinkingWaterConsumptionBuilder {
	return b.WithKL(kL)
}

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) WithKL(kL uint32) MeteringDataDrinkingWaterConsumptionBuilder {
	b.KL = kL
	return b
}

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) Build() (MeteringDataDrinkingWaterConsumption, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MeteringDataDrinkingWaterConsumption.deepCopy(), nil
}

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) MustBuild() MeteringDataDrinkingWaterConsumption {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MeteringDataDrinkingWaterConsumptionBuilder) Done() MeteringDataBuilder {
	return b.parentBuilder
}

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) buildForMeteringData() (MeteringData, error) {
	return b.Build()
}

func (b *_MeteringDataDrinkingWaterConsumptionBuilder) DeepCopy() any {
	_copy := b.CreateMeteringDataDrinkingWaterConsumptionBuilder().(*_MeteringDataDrinkingWaterConsumptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMeteringDataDrinkingWaterConsumptionBuilder creates a MeteringDataDrinkingWaterConsumptionBuilder
func (b *_MeteringDataDrinkingWaterConsumption) CreateMeteringDataDrinkingWaterConsumptionBuilder() MeteringDataDrinkingWaterConsumptionBuilder {
	if b == nil {
		return NewMeteringDataDrinkingWaterConsumptionBuilder()
	}
	return &_MeteringDataDrinkingWaterConsumptionBuilder{_MeteringDataDrinkingWaterConsumption: b.deepCopy()}
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

func (m *_MeteringDataDrinkingWaterConsumption) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataDrinkingWaterConsumption) GetKL() uint32 {
	return m.KL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMeteringDataDrinkingWaterConsumption(structType any) MeteringDataDrinkingWaterConsumption {
	if casted, ok := structType.(MeteringDataDrinkingWaterConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataDrinkingWaterConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataDrinkingWaterConsumption) GetTypeName() string {
	return "MeteringDataDrinkingWaterConsumption"
}

func (m *_MeteringDataDrinkingWaterConsumption) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	// Simple field (kL)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MeteringDataDrinkingWaterConsumption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataDrinkingWaterConsumption) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataDrinkingWaterConsumption MeteringDataDrinkingWaterConsumption, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataDrinkingWaterConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataDrinkingWaterConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	kL, err := ReadSimpleField(ctx, "kL", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'kL' field"))
	}
	m.KL = kL

	if closeErr := readBuffer.CloseContext("MeteringDataDrinkingWaterConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataDrinkingWaterConsumption")
	}

	return m, nil
}

func (m *_MeteringDataDrinkingWaterConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataDrinkingWaterConsumption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataDrinkingWaterConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataDrinkingWaterConsumption")
		}

		if err := WriteSimpleField[uint32](ctx, "kL", m.GetKL(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'kL' field")
		}

		if popErr := writeBuffer.PopContext("MeteringDataDrinkingWaterConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataDrinkingWaterConsumption")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataDrinkingWaterConsumption) IsMeteringDataDrinkingWaterConsumption() {}

func (m *_MeteringDataDrinkingWaterConsumption) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataDrinkingWaterConsumption) deepCopy() *_MeteringDataDrinkingWaterConsumption {
	if m == nil {
		return nil
	}
	_MeteringDataDrinkingWaterConsumptionCopy := &_MeteringDataDrinkingWaterConsumption{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
		m.KL,
	}
	m.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataDrinkingWaterConsumptionCopy
}

func (m *_MeteringDataDrinkingWaterConsumption) String() string {
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
