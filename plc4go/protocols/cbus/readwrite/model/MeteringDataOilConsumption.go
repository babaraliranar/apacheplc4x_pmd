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

// MeteringDataOilConsumption is the corresponding interface of MeteringDataOilConsumption
type MeteringDataOilConsumption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MeteringData
	// GetL returns L (property field)
	GetL() uint32
	// IsMeteringDataOilConsumption is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataOilConsumption()
}

// _MeteringDataOilConsumption is the data-structure of this message
type _MeteringDataOilConsumption struct {
	MeteringDataContract
	L uint32
}

var _ MeteringDataOilConsumption = (*_MeteringDataOilConsumption)(nil)
var _ MeteringDataRequirements = (*_MeteringDataOilConsumption)(nil)

// NewMeteringDataOilConsumption factory function for _MeteringDataOilConsumption
func NewMeteringDataOilConsumption(commandTypeContainer MeteringCommandTypeContainer, argument byte, L uint32) *_MeteringDataOilConsumption {
	_result := &_MeteringDataOilConsumption{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
		L:                    L,
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataOilConsumption) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataOilConsumption) GetL() uint32 {
	return m.L
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMeteringDataOilConsumption(structType any) MeteringDataOilConsumption {
	if casted, ok := structType.(MeteringDataOilConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataOilConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataOilConsumption) GetTypeName() string {
	return "MeteringDataOilConsumption"
}

func (m *_MeteringDataOilConsumption) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	// Simple field (L)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MeteringDataOilConsumption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataOilConsumption) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataOilConsumption MeteringDataOilConsumption, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataOilConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataOilConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	L, err := ReadSimpleField(ctx, "L", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'L' field"))
	}
	m.L = L

	if closeErr := readBuffer.CloseContext("MeteringDataOilConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataOilConsumption")
	}

	return m, nil
}

func (m *_MeteringDataOilConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataOilConsumption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataOilConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataOilConsumption")
		}

		if err := WriteSimpleField[uint32](ctx, "L", m.GetL(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'L' field")
		}

		if popErr := writeBuffer.PopContext("MeteringDataOilConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataOilConsumption")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataOilConsumption) IsMeteringDataOilConsumption() {}

func (m *_MeteringDataOilConsumption) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataOilConsumption) deepCopy() *_MeteringDataOilConsumption {
	if m == nil {
		return nil
	}
	_MeteringDataOilConsumptionCopy := &_MeteringDataOilConsumption{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
		m.L,
	}
	m.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataOilConsumptionCopy
}

func (m *_MeteringDataOilConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
