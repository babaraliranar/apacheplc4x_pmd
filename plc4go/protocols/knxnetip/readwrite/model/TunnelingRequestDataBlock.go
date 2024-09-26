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

// TunnelingRequestDataBlock is the corresponding interface of TunnelingRequestDataBlock
type TunnelingRequestDataBlock interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
	// IsTunnelingRequestDataBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTunnelingRequestDataBlock()
}

// _TunnelingRequestDataBlock is the data-structure of this message
type _TunnelingRequestDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ TunnelingRequestDataBlock = (*_TunnelingRequestDataBlock)(nil)

// NewTunnelingRequestDataBlock factory function for _TunnelingRequestDataBlock
func NewTunnelingRequestDataBlock(communicationChannelId uint8, sequenceCounter uint8) *_TunnelingRequestDataBlock {
	return &_TunnelingRequestDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingRequestDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_TunnelingRequestDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTunnelingRequestDataBlock(structType any) TunnelingRequestDataBlock {
	if casted, ok := structType.(TunnelingRequestDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingRequestDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingRequestDataBlock) GetTypeName() string {
	return "TunnelingRequestDataBlock"
}

func (m *_TunnelingRequestDataBlock) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TunnelingRequestDataBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TunnelingRequestDataBlockParse(ctx context.Context, theBytes []byte) (TunnelingRequestDataBlock, error) {
	return TunnelingRequestDataBlockParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TunnelingRequestDataBlockParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingRequestDataBlock, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingRequestDataBlock, error) {
		return TunnelingRequestDataBlockParseWithBuffer(ctx, readBuffer)
	}
}

func TunnelingRequestDataBlockParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingRequestDataBlock, error) {
	v, err := (&_TunnelingRequestDataBlock{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_TunnelingRequestDataBlock) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__tunnelingRequestDataBlock TunnelingRequestDataBlock, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingRequestDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingRequestDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	sequenceCounter, err := ReadSimpleField(ctx, "sequenceCounter", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceCounter' field"))
	}
	m.SequenceCounter = sequenceCounter

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("TunnelingRequestDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingRequestDataBlock")
	}

	return m, nil
}

func (m *_TunnelingRequestDataBlock) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TunnelingRequestDataBlock) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TunnelingRequestDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TunnelingRequestDataBlock")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sequenceCounter", m.GetSequenceCounter(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sequenceCounter' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if popErr := writeBuffer.PopContext("TunnelingRequestDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TunnelingRequestDataBlock")
	}
	return nil
}

func (m *_TunnelingRequestDataBlock) IsTunnelingRequestDataBlock() {}

func (m *_TunnelingRequestDataBlock) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TunnelingRequestDataBlock) deepCopy() *_TunnelingRequestDataBlock {
	if m == nil {
		return nil
	}
	_TunnelingRequestDataBlockCopy := &_TunnelingRequestDataBlock{
		m.CommunicationChannelId,
		m.SequenceCounter,
		m.reservedField0,
	}
	return _TunnelingRequestDataBlockCopy
}

func (m *_TunnelingRequestDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
