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

// IdentifyReplyCommandDSIStatus is the corresponding interface of IdentifyReplyCommandDSIStatus
type IdentifyReplyCommandDSIStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetChannelStatus1 returns ChannelStatus1 (property field)
	GetChannelStatus1() ChannelStatus
	// GetChannelStatus2 returns ChannelStatus2 (property field)
	GetChannelStatus2() ChannelStatus
	// GetChannelStatus3 returns ChannelStatus3 (property field)
	GetChannelStatus3() ChannelStatus
	// GetChannelStatus4 returns ChannelStatus4 (property field)
	GetChannelStatus4() ChannelStatus
	// GetChannelStatus5 returns ChannelStatus5 (property field)
	GetChannelStatus5() ChannelStatus
	// GetChannelStatus6 returns ChannelStatus6 (property field)
	GetChannelStatus6() ChannelStatus
	// GetChannelStatus7 returns ChannelStatus7 (property field)
	GetChannelStatus7() ChannelStatus
	// GetChannelStatus8 returns ChannelStatus8 (property field)
	GetChannelStatus8() ChannelStatus
	// GetUnitStatus returns UnitStatus (property field)
	GetUnitStatus() UnitStatus
	// GetDimmingUCRevisionNumber returns DimmingUCRevisionNumber (property field)
	GetDimmingUCRevisionNumber() byte
	// IsIdentifyReplyCommandDSIStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandDSIStatus()
	// CreateBuilder creates a IdentifyReplyCommandDSIStatusBuilder
	CreateIdentifyReplyCommandDSIStatusBuilder() IdentifyReplyCommandDSIStatusBuilder
}

// _IdentifyReplyCommandDSIStatus is the data-structure of this message
type _IdentifyReplyCommandDSIStatus struct {
	IdentifyReplyCommandContract
	ChannelStatus1          ChannelStatus
	ChannelStatus2          ChannelStatus
	ChannelStatus3          ChannelStatus
	ChannelStatus4          ChannelStatus
	ChannelStatus5          ChannelStatus
	ChannelStatus6          ChannelStatus
	ChannelStatus7          ChannelStatus
	ChannelStatus8          ChannelStatus
	UnitStatus              UnitStatus
	DimmingUCRevisionNumber byte
}

var _ IdentifyReplyCommandDSIStatus = (*_IdentifyReplyCommandDSIStatus)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandDSIStatus)(nil)

// NewIdentifyReplyCommandDSIStatus factory function for _IdentifyReplyCommandDSIStatus
func NewIdentifyReplyCommandDSIStatus(channelStatus1 ChannelStatus, channelStatus2 ChannelStatus, channelStatus3 ChannelStatus, channelStatus4 ChannelStatus, channelStatus5 ChannelStatus, channelStatus6 ChannelStatus, channelStatus7 ChannelStatus, channelStatus8 ChannelStatus, unitStatus UnitStatus, dimmingUCRevisionNumber byte, numBytes uint8) *_IdentifyReplyCommandDSIStatus {
	_result := &_IdentifyReplyCommandDSIStatus{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		ChannelStatus1:               channelStatus1,
		ChannelStatus2:               channelStatus2,
		ChannelStatus3:               channelStatus3,
		ChannelStatus4:               channelStatus4,
		ChannelStatus5:               channelStatus5,
		ChannelStatus6:               channelStatus6,
		ChannelStatus7:               channelStatus7,
		ChannelStatus8:               channelStatus8,
		UnitStatus:                   unitStatus,
		DimmingUCRevisionNumber:      dimmingUCRevisionNumber,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandDSIStatusBuilder is a builder for IdentifyReplyCommandDSIStatus
type IdentifyReplyCommandDSIStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(channelStatus1 ChannelStatus, channelStatus2 ChannelStatus, channelStatus3 ChannelStatus, channelStatus4 ChannelStatus, channelStatus5 ChannelStatus, channelStatus6 ChannelStatus, channelStatus7 ChannelStatus, channelStatus8 ChannelStatus, unitStatus UnitStatus, dimmingUCRevisionNumber byte) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus1 adds ChannelStatus1 (property field)
	WithChannelStatus1(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus2 adds ChannelStatus2 (property field)
	WithChannelStatus2(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus3 adds ChannelStatus3 (property field)
	WithChannelStatus3(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus4 adds ChannelStatus4 (property field)
	WithChannelStatus4(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus5 adds ChannelStatus5 (property field)
	WithChannelStatus5(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus6 adds ChannelStatus6 (property field)
	WithChannelStatus6(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus7 adds ChannelStatus7 (property field)
	WithChannelStatus7(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithChannelStatus8 adds ChannelStatus8 (property field)
	WithChannelStatus8(ChannelStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithUnitStatus adds UnitStatus (property field)
	WithUnitStatus(UnitStatus) IdentifyReplyCommandDSIStatusBuilder
	// WithDimmingUCRevisionNumber adds DimmingUCRevisionNumber (property field)
	WithDimmingUCRevisionNumber(byte) IdentifyReplyCommandDSIStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() IdentifyReplyCommandBuilder
	// Build builds the IdentifyReplyCommandDSIStatus or returns an error if something is wrong
	Build() (IdentifyReplyCommandDSIStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandDSIStatus
}

// NewIdentifyReplyCommandDSIStatusBuilder() creates a IdentifyReplyCommandDSIStatusBuilder
func NewIdentifyReplyCommandDSIStatusBuilder() IdentifyReplyCommandDSIStatusBuilder {
	return &_IdentifyReplyCommandDSIStatusBuilder{_IdentifyReplyCommandDSIStatus: new(_IdentifyReplyCommandDSIStatus)}
}

type _IdentifyReplyCommandDSIStatusBuilder struct {
	*_IdentifyReplyCommandDSIStatus

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandDSIStatusBuilder) = (*_IdentifyReplyCommandDSIStatusBuilder)(nil)

func (b *_IdentifyReplyCommandDSIStatusBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
	contract.(*_IdentifyReplyCommand)._SubType = b._IdentifyReplyCommandDSIStatus
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithMandatoryFields(channelStatus1 ChannelStatus, channelStatus2 ChannelStatus, channelStatus3 ChannelStatus, channelStatus4 ChannelStatus, channelStatus5 ChannelStatus, channelStatus6 ChannelStatus, channelStatus7 ChannelStatus, channelStatus8 ChannelStatus, unitStatus UnitStatus, dimmingUCRevisionNumber byte) IdentifyReplyCommandDSIStatusBuilder {
	return b.WithChannelStatus1(channelStatus1).WithChannelStatus2(channelStatus2).WithChannelStatus3(channelStatus3).WithChannelStatus4(channelStatus4).WithChannelStatus5(channelStatus5).WithChannelStatus6(channelStatus6).WithChannelStatus7(channelStatus7).WithChannelStatus8(channelStatus8).WithUnitStatus(unitStatus).WithDimmingUCRevisionNumber(dimmingUCRevisionNumber)
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus1(channelStatus1 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus1 = channelStatus1
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus2(channelStatus2 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus2 = channelStatus2
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus3(channelStatus3 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus3 = channelStatus3
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus4(channelStatus4 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus4 = channelStatus4
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus5(channelStatus5 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus5 = channelStatus5
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus6(channelStatus6 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus6 = channelStatus6
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus7(channelStatus7 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus7 = channelStatus7
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithChannelStatus8(channelStatus8 ChannelStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.ChannelStatus8 = channelStatus8
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithUnitStatus(unitStatus UnitStatus) IdentifyReplyCommandDSIStatusBuilder {
	b.UnitStatus = unitStatus
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) WithDimmingUCRevisionNumber(dimmingUCRevisionNumber byte) IdentifyReplyCommandDSIStatusBuilder {
	b.DimmingUCRevisionNumber = dimmingUCRevisionNumber
	return b
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) Build() (IdentifyReplyCommandDSIStatus, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandDSIStatus.deepCopy(), nil
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) MustBuild() IdentifyReplyCommandDSIStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) Done() IdentifyReplyCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewIdentifyReplyCommandBuilder().(*_IdentifyReplyCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandDSIStatusBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandDSIStatusBuilder().(*_IdentifyReplyCommandDSIStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandDSIStatusBuilder creates a IdentifyReplyCommandDSIStatusBuilder
func (b *_IdentifyReplyCommandDSIStatus) CreateIdentifyReplyCommandDSIStatusBuilder() IdentifyReplyCommandDSIStatusBuilder {
	if b == nil {
		return NewIdentifyReplyCommandDSIStatusBuilder()
	}
	return &_IdentifyReplyCommandDSIStatusBuilder{_IdentifyReplyCommandDSIStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandDSIStatus) GetAttribute() Attribute {
	return Attribute_DSIStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandDSIStatus) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus1() ChannelStatus {
	return m.ChannelStatus1
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus2() ChannelStatus {
	return m.ChannelStatus2
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus3() ChannelStatus {
	return m.ChannelStatus3
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus4() ChannelStatus {
	return m.ChannelStatus4
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus5() ChannelStatus {
	return m.ChannelStatus5
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus6() ChannelStatus {
	return m.ChannelStatus6
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus7() ChannelStatus {
	return m.ChannelStatus7
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus8() ChannelStatus {
	return m.ChannelStatus8
}

func (m *_IdentifyReplyCommandDSIStatus) GetUnitStatus() UnitStatus {
	return m.UnitStatus
}

func (m *_IdentifyReplyCommandDSIStatus) GetDimmingUCRevisionNumber() byte {
	return m.DimmingUCRevisionNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandDSIStatus(structType any) IdentifyReplyCommandDSIStatus {
	if casted, ok := structType.(IdentifyReplyCommandDSIStatus); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandDSIStatus); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandDSIStatus) GetTypeName() string {
	return "IdentifyReplyCommandDSIStatus"
}

func (m *_IdentifyReplyCommandDSIStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Simple field (channelStatus1)
	lengthInBits += 8

	// Simple field (channelStatus2)
	lengthInBits += 8

	// Simple field (channelStatus3)
	lengthInBits += 8

	// Simple field (channelStatus4)
	lengthInBits += 8

	// Simple field (channelStatus5)
	lengthInBits += 8

	// Simple field (channelStatus6)
	lengthInBits += 8

	// Simple field (channelStatus7)
	lengthInBits += 8

	// Simple field (channelStatus8)
	lengthInBits += 8

	// Simple field (unitStatus)
	lengthInBits += 8

	// Simple field (dimmingUCRevisionNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandDSIStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandDSIStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandDSIStatus IdentifyReplyCommandDSIStatus, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandDSIStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandDSIStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	channelStatus1, err := ReadEnumField[ChannelStatus](ctx, "channelStatus1", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus1' field"))
	}
	m.ChannelStatus1 = channelStatus1

	channelStatus2, err := ReadEnumField[ChannelStatus](ctx, "channelStatus2", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus2' field"))
	}
	m.ChannelStatus2 = channelStatus2

	channelStatus3, err := ReadEnumField[ChannelStatus](ctx, "channelStatus3", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus3' field"))
	}
	m.ChannelStatus3 = channelStatus3

	channelStatus4, err := ReadEnumField[ChannelStatus](ctx, "channelStatus4", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus4' field"))
	}
	m.ChannelStatus4 = channelStatus4

	channelStatus5, err := ReadEnumField[ChannelStatus](ctx, "channelStatus5", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus5' field"))
	}
	m.ChannelStatus5 = channelStatus5

	channelStatus6, err := ReadEnumField[ChannelStatus](ctx, "channelStatus6", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus6' field"))
	}
	m.ChannelStatus6 = channelStatus6

	channelStatus7, err := ReadEnumField[ChannelStatus](ctx, "channelStatus7", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus7' field"))
	}
	m.ChannelStatus7 = channelStatus7

	channelStatus8, err := ReadEnumField[ChannelStatus](ctx, "channelStatus8", "ChannelStatus", ReadEnum(ChannelStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelStatus8' field"))
	}
	m.ChannelStatus8 = channelStatus8

	unitStatus, err := ReadEnumField[UnitStatus](ctx, "unitStatus", "UnitStatus", ReadEnum(UnitStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitStatus' field"))
	}
	m.UnitStatus = unitStatus

	dimmingUCRevisionNumber, err := ReadSimpleField(ctx, "dimmingUCRevisionNumber", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dimmingUCRevisionNumber' field"))
	}
	m.DimmingUCRevisionNumber = dimmingUCRevisionNumber

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandDSIStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandDSIStatus")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandDSIStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandDSIStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandDSIStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandDSIStatus")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus1", "ChannelStatus", m.GetChannelStatus1(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus1' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus2", "ChannelStatus", m.GetChannelStatus2(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus2' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus3", "ChannelStatus", m.GetChannelStatus3(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus3' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus4", "ChannelStatus", m.GetChannelStatus4(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus4' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus5", "ChannelStatus", m.GetChannelStatus5(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus5' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus6", "ChannelStatus", m.GetChannelStatus6(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus6' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus7", "ChannelStatus", m.GetChannelStatus7(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus7' field")
		}

		if err := WriteSimpleEnumField[ChannelStatus](ctx, "channelStatus8", "ChannelStatus", m.GetChannelStatus8(), WriteEnum[ChannelStatus, uint8](ChannelStatus.GetValue, ChannelStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'channelStatus8' field")
		}

		if err := WriteSimpleEnumField[UnitStatus](ctx, "unitStatus", "UnitStatus", m.GetUnitStatus(), WriteEnum[UnitStatus, uint8](UnitStatus.GetValue, UnitStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'unitStatus' field")
		}

		if err := WriteSimpleField[byte](ctx, "dimmingUCRevisionNumber", m.GetDimmingUCRevisionNumber(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'dimmingUCRevisionNumber' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandDSIStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandDSIStatus")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandDSIStatus) IsIdentifyReplyCommandDSIStatus() {}

func (m *_IdentifyReplyCommandDSIStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandDSIStatus) deepCopy() *_IdentifyReplyCommandDSIStatus {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandDSIStatusCopy := &_IdentifyReplyCommandDSIStatus{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		m.ChannelStatus1,
		m.ChannelStatus2,
		m.ChannelStatus3,
		m.ChannelStatus4,
		m.ChannelStatus5,
		m.ChannelStatus6,
		m.ChannelStatus7,
		m.ChannelStatus8,
		m.UnitStatus,
		m.DimmingUCRevisionNumber,
	}
	_IdentifyReplyCommandDSIStatusCopy.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandDSIStatusCopy
}

func (m *_IdentifyReplyCommandDSIStatus) String() string {
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
