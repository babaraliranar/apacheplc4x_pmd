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

// RepublishResponse is the corresponding interface of RepublishResponse
type RepublishResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNotificationMessage returns NotificationMessage (property field)
	GetNotificationMessage() ExtensionObjectDefinition
	// IsRepublishResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRepublishResponse()
}

// _RepublishResponse is the data-structure of this message
type _RepublishResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ExtensionObjectDefinition
	NotificationMessage ExtensionObjectDefinition
}

var _ RepublishResponse = (*_RepublishResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RepublishResponse)(nil)

// NewRepublishResponse factory function for _RepublishResponse
func NewRepublishResponse(responseHeader ExtensionObjectDefinition, notificationMessage ExtensionObjectDefinition) *_RepublishResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for RepublishResponse must not be nil")
	}
	if notificationMessage == nil {
		panic("notificationMessage of type ExtensionObjectDefinition for RepublishResponse must not be nil")
	}
	_result := &_RepublishResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NotificationMessage:               notificationMessage,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RepublishResponse) GetIdentifier() string {
	return "835"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RepublishResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RepublishResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_RepublishResponse) GetNotificationMessage() ExtensionObjectDefinition {
	return m.NotificationMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRepublishResponse(structType any) RepublishResponse {
	if casted, ok := structType.(RepublishResponse); ok {
		return casted
	}
	if casted, ok := structType.(*RepublishResponse); ok {
		return *casted
	}
	return nil
}

func (m *_RepublishResponse) GetTypeName() string {
	return "RepublishResponse"
}

func (m *_RepublishResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (notificationMessage)
	lengthInBits += m.NotificationMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RepublishResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RepublishResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__republishResponse RepublishResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RepublishResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RepublishResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	notificationMessage, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "notificationMessage", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("805")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationMessage' field"))
	}
	m.NotificationMessage = notificationMessage

	if closeErr := readBuffer.CloseContext("RepublishResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RepublishResponse")
	}

	return m, nil
}

func (m *_RepublishResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RepublishResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RepublishResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RepublishResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "notificationMessage", m.GetNotificationMessage(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationMessage' field")
		}

		if popErr := writeBuffer.PopContext("RepublishResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RepublishResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RepublishResponse) IsRepublishResponse() {}

func (m *_RepublishResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RepublishResponse) deepCopy() *_RepublishResponse {
	if m == nil {
		return nil
	}
	_RepublishResponseCopy := &_RepublishResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NotificationMessage.DeepCopy().(ExtensionObjectDefinition),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RepublishResponseCopy
}

func (m *_RepublishResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
