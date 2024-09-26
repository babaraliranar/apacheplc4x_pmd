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

// GroupObjectDescriptorRealisationType7 is the corresponding interface of GroupObjectDescriptorRealisationType7
type GroupObjectDescriptorRealisationType7 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetDataAddress returns DataAddress (property field)
	GetDataAddress() uint16
	// GetUpdateEnable returns UpdateEnable (property field)
	GetUpdateEnable() bool
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
	// IsGroupObjectDescriptorRealisationType7 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGroupObjectDescriptorRealisationType7()
}

// _GroupObjectDescriptorRealisationType7 is the data-structure of this message
type _GroupObjectDescriptorRealisationType7 struct {
	DataAddress           uint16
	UpdateEnable          bool
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

var _ GroupObjectDescriptorRealisationType7 = (*_GroupObjectDescriptorRealisationType7)(nil)

// NewGroupObjectDescriptorRealisationType7 factory function for _GroupObjectDescriptorRealisationType7
func NewGroupObjectDescriptorRealisationType7(dataAddress uint16, updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *_GroupObjectDescriptorRealisationType7 {
	return &_GroupObjectDescriptorRealisationType7{DataAddress: dataAddress, UpdateEnable: updateEnable, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GroupObjectDescriptorRealisationType7) GetDataAddress() uint16 {
	return m.DataAddress
}

func (m *_GroupObjectDescriptorRealisationType7) GetUpdateEnable() bool {
	return m.UpdateEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_GroupObjectDescriptorRealisationType7) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationType7(structType any) GroupObjectDescriptorRealisationType7 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType7); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType7); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType7) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType7"
}

func (m *_GroupObjectDescriptorRealisationType7) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataAddress)
	lengthInBits += 16

	// Simple field (updateEnable)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationType7) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GroupObjectDescriptorRealisationType7Parse(ctx context.Context, theBytes []byte) (GroupObjectDescriptorRealisationType7, error) {
	return GroupObjectDescriptorRealisationType7ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GroupObjectDescriptorRealisationType7ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType7, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType7, error) {
		return GroupObjectDescriptorRealisationType7ParseWithBuffer(ctx, readBuffer)
	}
}

func GroupObjectDescriptorRealisationType7ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType7, error) {
	v, err := (&_GroupObjectDescriptorRealisationType7{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_GroupObjectDescriptorRealisationType7) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__groupObjectDescriptorRealisationType7 GroupObjectDescriptorRealisationType7, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType7"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType7")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataAddress, err := ReadSimpleField(ctx, "dataAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataAddress' field"))
	}
	m.DataAddress = dataAddress

	updateEnable, err := ReadSimpleField(ctx, "updateEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateEnable' field"))
	}
	m.UpdateEnable = updateEnable

	transmitEnable, err := ReadSimpleField(ctx, "transmitEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transmitEnable' field"))
	}
	m.TransmitEnable = transmitEnable

	segmentSelectorEnable, err := ReadSimpleField(ctx, "segmentSelectorEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentSelectorEnable' field"))
	}
	m.SegmentSelectorEnable = segmentSelectorEnable

	writeEnable, err := ReadSimpleField(ctx, "writeEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeEnable' field"))
	}
	m.WriteEnable = writeEnable

	readEnable, err := ReadSimpleField(ctx, "readEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readEnable' field"))
	}
	m.ReadEnable = readEnable

	communicationEnable, err := ReadSimpleField(ctx, "communicationEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationEnable' field"))
	}
	m.CommunicationEnable = communicationEnable

	priority, err := ReadEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", ReadEnum(CEMIPriorityByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	valueType, err := ReadEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", ReadEnum(ComObjectValueTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueType' field"))
	}
	m.ValueType = valueType

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType7"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType7")
	}

	return m, nil
}

func (m *_GroupObjectDescriptorRealisationType7) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GroupObjectDescriptorRealisationType7) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType7"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType7")
	}

	if err := WriteSimpleField[uint16](ctx, "dataAddress", m.GetDataAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataAddress' field")
	}

	if err := WriteSimpleField[bool](ctx, "updateEnable", m.GetUpdateEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'updateEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "transmitEnable", m.GetTransmitEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'transmitEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "segmentSelectorEnable", m.GetSegmentSelectorEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'segmentSelectorEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "writeEnable", m.GetWriteEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'writeEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "readEnable", m.GetReadEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'readEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "communicationEnable", m.GetCommunicationEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationEnable' field")
	}

	if err := WriteSimpleEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", m.GetPriority(), WriteEnum[CEMIPriority, uint8](CEMIPriority.GetValue, CEMIPriority.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if err := WriteSimpleEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", m.GetValueType(), WriteEnum[ComObjectValueType, uint8](ComObjectValueType.GetValue, ComObjectValueType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType7"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType7")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType7) IsGroupObjectDescriptorRealisationType7() {}

func (m *_GroupObjectDescriptorRealisationType7) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GroupObjectDescriptorRealisationType7) deepCopy() *_GroupObjectDescriptorRealisationType7 {
	if m == nil {
		return nil
	}
	_GroupObjectDescriptorRealisationType7Copy := &_GroupObjectDescriptorRealisationType7{
		m.DataAddress,
		m.UpdateEnable,
		m.TransmitEnable,
		m.SegmentSelectorEnable,
		m.WriteEnable,
		m.ReadEnable,
		m.CommunicationEnable,
		m.Priority,
		m.ValueType,
	}
	return _GroupObjectDescriptorRealisationType7Copy
}

func (m *_GroupObjectDescriptorRealisationType7) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
