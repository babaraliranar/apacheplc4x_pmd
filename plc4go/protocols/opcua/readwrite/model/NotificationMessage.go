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

// NotificationMessage is the corresponding interface of NotificationMessage
type NotificationMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint32
	// GetPublishTime returns PublishTime (property field)
	GetPublishTime() int64
	// GetNoOfNotificationData returns NoOfNotificationData (property field)
	GetNoOfNotificationData() int32
	// GetNotificationData returns NotificationData (property field)
	GetNotificationData() []ExtensionObject
	// IsNotificationMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNotificationMessage()
}

// _NotificationMessage is the data-structure of this message
type _NotificationMessage struct {
	ExtensionObjectDefinitionContract
	SequenceNumber       uint32
	PublishTime          int64
	NoOfNotificationData int32
	NotificationData     []ExtensionObject
}

var _ NotificationMessage = (*_NotificationMessage)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NotificationMessage)(nil)

// NewNotificationMessage factory function for _NotificationMessage
func NewNotificationMessage(sequenceNumber uint32, publishTime int64, noOfNotificationData int32, notificationData []ExtensionObject) *_NotificationMessage {
	_result := &_NotificationMessage{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SequenceNumber:                    sequenceNumber,
		PublishTime:                       publishTime,
		NoOfNotificationData:              noOfNotificationData,
		NotificationData:                  notificationData,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NotificationMessage) GetIdentifier() string {
	return "805"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NotificationMessage) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NotificationMessage) GetSequenceNumber() uint32 {
	return m.SequenceNumber
}

func (m *_NotificationMessage) GetPublishTime() int64 {
	return m.PublishTime
}

func (m *_NotificationMessage) GetNoOfNotificationData() int32 {
	return m.NoOfNotificationData
}

func (m *_NotificationMessage) GetNotificationData() []ExtensionObject {
	return m.NotificationData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNotificationMessage(structType any) NotificationMessage {
	if casted, ok := structType.(NotificationMessage); ok {
		return casted
	}
	if casted, ok := structType.(*NotificationMessage); ok {
		return *casted
	}
	return nil
}

func (m *_NotificationMessage) GetTypeName() string {
	return "NotificationMessage"
}

func (m *_NotificationMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (sequenceNumber)
	lengthInBits += 32

	// Simple field (publishTime)
	lengthInBits += 64

	// Simple field (noOfNotificationData)
	lengthInBits += 32

	// Array field
	if len(m.NotificationData) > 0 {
		for _curItem, element := range m.NotificationData {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NotificationData), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NotificationMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NotificationMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__notificationMessage NotificationMessage, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NotificationMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NotificationMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	publishTime, err := ReadSimpleField(ctx, "publishTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishTime' field"))
	}
	m.PublishTime = publishTime

	noOfNotificationData, err := ReadSimpleField(ctx, "noOfNotificationData", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNotificationData' field"))
	}
	m.NoOfNotificationData = noOfNotificationData

	notificationData, err := ReadCountArrayField[ExtensionObject](ctx, "notificationData", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer), uint64(noOfNotificationData))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationData' field"))
	}
	m.NotificationData = notificationData

	if closeErr := readBuffer.CloseContext("NotificationMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NotificationMessage")
	}

	return m, nil
}

func (m *_NotificationMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NotificationMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotificationMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NotificationMessage")
		}

		if err := WriteSimpleField[uint32](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteSimpleField[int64](ctx, "publishTime", m.GetPublishTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishTime' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNotificationData", m.GetNoOfNotificationData(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNotificationData' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "notificationData", m.GetNotificationData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationData' field")
		}

		if popErr := writeBuffer.PopContext("NotificationMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NotificationMessage")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NotificationMessage) IsNotificationMessage() {}

func (m *_NotificationMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NotificationMessage) deepCopy() *_NotificationMessage {
	if m == nil {
		return nil
	}
	_NotificationMessageCopy := &_NotificationMessage{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SequenceNumber,
		m.PublishTime,
		m.NoOfNotificationData,
		utils.DeepCopySlice[ExtensionObject, ExtensionObject](m.NotificationData),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NotificationMessageCopy
}

func (m *_NotificationMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
