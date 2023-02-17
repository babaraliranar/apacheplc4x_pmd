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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// MeteringDataElectricityConsumption is the corresponding interface of MeteringDataElectricityConsumption
type MeteringDataElectricityConsumption interface {
	utils.LengthAware
	utils.Serializable
	MeteringData
	// GetKWhr returns KWhr (property field)
	GetKWhr() uint32
}

// MeteringDataElectricityConsumptionExactly can be used when we want exactly this type and not a type which fulfills MeteringDataElectricityConsumption.
// This is useful for switch cases.
type MeteringDataElectricityConsumptionExactly interface {
	MeteringDataElectricityConsumption
	isMeteringDataElectricityConsumption() bool
}

// _MeteringDataElectricityConsumption is the data-structure of this message
type _MeteringDataElectricityConsumption struct {
	*_MeteringData
        KWhr uint32
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataElectricityConsumption) InitializeParent(parent MeteringData , commandTypeContainer MeteringCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataElectricityConsumption)  GetParent() MeteringData {
	return m._MeteringData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataElectricityConsumption) GetKWhr() uint32 {
	return m.KWhr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMeteringDataElectricityConsumption factory function for _MeteringDataElectricityConsumption
func NewMeteringDataElectricityConsumption( kWhr uint32 , commandTypeContainer MeteringCommandTypeContainer , argument byte ) *_MeteringDataElectricityConsumption {
	_result := &_MeteringDataElectricityConsumption{
		KWhr: kWhr,
    	_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataElectricityConsumption(structType interface{}) MeteringDataElectricityConsumption {
    if casted, ok := structType.(MeteringDataElectricityConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataElectricityConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataElectricityConsumption) GetTypeName() string {
	return "MeteringDataElectricityConsumption"
}

func (m *_MeteringDataElectricityConsumption) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (kWhr)
	lengthInBits += 32;

	return lengthInBits
}


func (m *_MeteringDataElectricityConsumption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataElectricityConsumptionParse(theBytes []byte) (MeteringDataElectricityConsumption, error) {
	return MeteringDataElectricityConsumptionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataElectricityConsumptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataElectricityConsumption, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataElectricityConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataElectricityConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (kWhr)
_kWhr, _kWhrErr := readBuffer.ReadUint32("kWhr", 32)
	if _kWhrErr != nil {
		return nil, errors.Wrap(_kWhrErr, "Error parsing 'kWhr' field of MeteringDataElectricityConsumption")
	}
	kWhr := _kWhr

	if closeErr := readBuffer.CloseContext("MeteringDataElectricityConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataElectricityConsumption")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataElectricityConsumption{
		_MeteringData: &_MeteringData{
		},
		KWhr: kWhr,
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataElectricityConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataElectricityConsumption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataElectricityConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataElectricityConsumption")
		}

	// Simple Field (kWhr)
	kWhr := uint32(m.GetKWhr())
	_kWhrErr := writeBuffer.WriteUint32("kWhr", 32, (kWhr))
	if _kWhrErr != nil {
		return errors.Wrap(_kWhrErr, "Error serializing 'kWhr' field")
	}

		if popErr := writeBuffer.PopContext("MeteringDataElectricityConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataElectricityConsumption")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_MeteringDataElectricityConsumption) isMeteringDataElectricityConsumption() bool {
	return true
}

func (m *_MeteringDataElectricityConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



