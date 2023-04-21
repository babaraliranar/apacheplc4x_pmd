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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// ReplyOrConfirmationConfirmation is the corresponding interface of ReplyOrConfirmationConfirmation
type ReplyOrConfirmationConfirmation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ReplyOrConfirmation
	// GetConfirmation returns Confirmation (property field)
	GetConfirmation() Confirmation
	// GetEmbeddedReply returns EmbeddedReply (property field)
	GetEmbeddedReply() ReplyOrConfirmation
}

// ReplyOrConfirmationConfirmationExactly can be used when we want exactly this type and not a type which fulfills ReplyOrConfirmationConfirmation.
// This is useful for switch cases.
type ReplyOrConfirmationConfirmationExactly interface {
	ReplyOrConfirmationConfirmation
	isReplyOrConfirmationConfirmation() bool
}

// _ReplyOrConfirmationConfirmation is the data-structure of this message
type _ReplyOrConfirmationConfirmation struct {
	*_ReplyOrConfirmation
	Confirmation  Confirmation
	EmbeddedReply ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyOrConfirmationConfirmation) InitializeParent(parent ReplyOrConfirmation, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ReplyOrConfirmationConfirmation) GetParent() ReplyOrConfirmation {
	return m._ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmationConfirmation) GetConfirmation() Confirmation {
	return m.Confirmation
}

func (m *_ReplyOrConfirmationConfirmation) GetEmbeddedReply() ReplyOrConfirmation {
	return m.EmbeddedReply
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyOrConfirmationConfirmation factory function for _ReplyOrConfirmationConfirmation
func NewReplyOrConfirmationConfirmation(confirmation Confirmation, embeddedReply ReplyOrConfirmation, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmationConfirmation {
	_result := &_ReplyOrConfirmationConfirmation{
		Confirmation:         confirmation,
		EmbeddedReply:        embeddedReply,
		_ReplyOrConfirmation: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
	}
	_result._ReplyOrConfirmation._ReplyOrConfirmationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmationConfirmation(structType any) ReplyOrConfirmationConfirmation {
	if casted, ok := structType.(ReplyOrConfirmationConfirmation); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmationConfirmation); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmationConfirmation) GetTypeName() string {
	return "ReplyOrConfirmationConfirmation"
}

func (m *_ReplyOrConfirmationConfirmation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (confirmation)
	lengthInBits += m.Confirmation.GetLengthInBits(ctx)

	// Optional Field (embeddedReply)
	if m.EmbeddedReply != nil {
		lengthInBits += m.EmbeddedReply.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_ReplyOrConfirmationConfirmation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReplyOrConfirmationConfirmationParse(theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (ReplyOrConfirmationConfirmation, error) {
	return ReplyOrConfirmationConfirmationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func ReplyOrConfirmationConfirmationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ReplyOrConfirmationConfirmation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyOrConfirmationConfirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmationConfirmation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (confirmation)
	if pullErr := readBuffer.PullContext("confirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for confirmation")
	}
	_confirmation, _confirmationErr := ConfirmationParseWithBuffer(ctx, readBuffer)
	if _confirmationErr != nil {
		return nil, errors.Wrap(_confirmationErr, "Error parsing 'confirmation' field of ReplyOrConfirmationConfirmation")
	}
	confirmation := _confirmation.(Confirmation)
	if closeErr := readBuffer.CloseContext("confirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for confirmation")
	}

	// Optional Field (embeddedReply) (Can be skipped, if a given expression evaluates to false)
	var embeddedReply ReplyOrConfirmation = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("embeddedReply"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for embeddedReply")
		}
		_val, _err := ReplyOrConfirmationParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'embeddedReply' field of ReplyOrConfirmationConfirmation")
		default:
			embeddedReply = _val.(ReplyOrConfirmation)
			if closeErr := readBuffer.CloseContext("embeddedReply"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for embeddedReply")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmationConfirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmationConfirmation")
	}

	// Create a partially initialized instance
	_child := &_ReplyOrConfirmationConfirmation{
		_ReplyOrConfirmation: &_ReplyOrConfirmation{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		Confirmation:  confirmation,
		EmbeddedReply: embeddedReply,
	}
	_child._ReplyOrConfirmation._ReplyOrConfirmationChildRequirements = _child
	return _child, nil
}

func (m *_ReplyOrConfirmationConfirmation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyOrConfirmationConfirmation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyOrConfirmationConfirmation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmationConfirmation")
		}

		// Simple Field (confirmation)
		if pushErr := writeBuffer.PushContext("confirmation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for confirmation")
		}
		_confirmationErr := writeBuffer.WriteSerializable(ctx, m.GetConfirmation())
		if popErr := writeBuffer.PopContext("confirmation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for confirmation")
		}
		if _confirmationErr != nil {
			return errors.Wrap(_confirmationErr, "Error serializing 'confirmation' field")
		}

		// Optional Field (embeddedReply) (Can be skipped, if the value is null)
		var embeddedReply ReplyOrConfirmation = nil
		if m.GetEmbeddedReply() != nil {
			if pushErr := writeBuffer.PushContext("embeddedReply"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for embeddedReply")
			}
			embeddedReply = m.GetEmbeddedReply()
			_embeddedReplyErr := writeBuffer.WriteSerializable(ctx, embeddedReply)
			if popErr := writeBuffer.PopContext("embeddedReply"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for embeddedReply")
			}
			if _embeddedReplyErr != nil {
				return errors.Wrap(_embeddedReplyErr, "Error serializing 'embeddedReply' field")
			}
		}

		if popErr := writeBuffer.PopContext("ReplyOrConfirmationConfirmation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyOrConfirmationConfirmation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReplyOrConfirmationConfirmation) isReplyOrConfirmationConfirmation() bool {
	return true
}

func (m *_ReplyOrConfirmationConfirmation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
