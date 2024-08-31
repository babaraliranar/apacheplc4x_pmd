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
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ReplyEncodedReply is the corresponding interface of ReplyEncodedReply
type ReplyEncodedReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Reply
	// GetEncodedReply returns EncodedReply (property field)
	GetEncodedReply() EncodedReply
	// GetChksum returns Chksum (property field)
	GetChksum() Checksum
	// GetEncodedReplyDecoded returns EncodedReplyDecoded (virtual field)
	GetEncodedReplyDecoded() EncodedReply
	// GetChksumDecoded returns ChksumDecoded (virtual field)
	GetChksumDecoded() Checksum
}

// ReplyEncodedReplyExactly can be used when we want exactly this type and not a type which fulfills ReplyEncodedReply.
// This is useful for switch cases.
type ReplyEncodedReplyExactly interface {
	ReplyEncodedReply
	isReplyEncodedReply() bool
}

// _ReplyEncodedReply is the data-structure of this message
type _ReplyEncodedReply struct {
	*_Reply
	EncodedReply EncodedReply
	Chksum       Checksum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyEncodedReply) InitializeParent(parent Reply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ReplyEncodedReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyEncodedReply) GetEncodedReply() EncodedReply {
	return m.EncodedReply
}

func (m *_ReplyEncodedReply) GetChksum() Checksum {
	return m.Chksum
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ReplyEncodedReply) GetEncodedReplyDecoded() EncodedReply {
	ctx := context.Background()
	_ = ctx
	return CastEncodedReply(m.GetEncodedReply())
}

func (m *_ReplyEncodedReply) GetChksumDecoded() Checksum {
	ctx := context.Background()
	_ = ctx
	return CastChecksum(m.GetChksum())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyEncodedReply factory function for _ReplyEncodedReply
func NewReplyEncodedReply(encodedReply EncodedReply, chksum Checksum, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyEncodedReply {
	_result := &_ReplyEncodedReply{
		EncodedReply: encodedReply,
		Chksum:       chksum,
		_Reply:       NewReply(peekedByte, cBusOptions, requestContext),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyEncodedReply(structType any) ReplyEncodedReply {
	if casted, ok := structType.(ReplyEncodedReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyEncodedReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyEncodedReply) GetTypeName() string {
	return "ReplyEncodedReply"
}

func (m *_ReplyEncodedReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Manual Field (encodedReply)
	lengthInBits += uint16(int32((int32(m.GetEncodedReply().GetLengthInBytes(ctx)) * int32(int32(2)))) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Manual Field (chksum)
	lengthInBits += uint16(utils.InlineIf((m.CBusOptions.GetSrchk()), func() any { return int32((int32(16))) }, func() any { return int32((int32(0))) }).(int32))

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ReplyEncodedReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReplyEncodedReplyParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (ReplyEncodedReply, error) {
	return ReplyEncodedReplyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func ReplyEncodedReplyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ReplyEncodedReply, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReplyEncodedReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyEncodedReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	encodedReply, err := ReadManualField[EncodedReply](ctx, "encodedReply", readBuffer, EnsureType[EncodedReply](ReadEncodedReply(ctx, readBuffer, cBusOptions, requestContext, cBusOptions.GetSrchk())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encodedReply' field"))
	}

	// Virtual field
	_encodedReplyDecoded := encodedReply
	encodedReplyDecoded := _encodedReplyDecoded
	_ = encodedReplyDecoded

	chksum, err := ReadManualField[Checksum](ctx, "chksum", readBuffer, EnsureType[Checksum](ReadAndValidateChecksum(ctx, readBuffer, encodedReply, cBusOptions.GetSrchk())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chksum' field"))
	}

	// Virtual field
	_chksumDecoded := chksum
	chksumDecoded := _chksumDecoded
	_ = chksumDecoded

	if closeErr := readBuffer.CloseContext("ReplyEncodedReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyEncodedReply")
	}

	// Create a partially initialized instance
	_child := &_ReplyEncodedReply{
		_Reply: &_Reply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		EncodedReply: encodedReply,
		Chksum:       chksum,
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_ReplyEncodedReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyEncodedReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyEncodedReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyEncodedReply")
		}

		// Manual Field (encodedReply)
		_encodedReplyErr := WriteEncodedReply(ctx, writeBuffer, m.GetEncodedReply())
		if _encodedReplyErr != nil {
			return errors.Wrap(_encodedReplyErr, "Error serializing 'encodedReply' field")
		}
		// Virtual field
		encodedReplyDecoded := m.GetEncodedReplyDecoded()
		_ = encodedReplyDecoded
		if _encodedReplyDecodedErr := writeBuffer.WriteVirtual(ctx, "encodedReplyDecoded", m.GetEncodedReplyDecoded()); _encodedReplyDecodedErr != nil {
			return errors.Wrap(_encodedReplyDecodedErr, "Error serializing 'encodedReplyDecoded' field")
		}

		// Manual Field (chksum)
		_chksumErr := CalculateChecksum(ctx, writeBuffer, m.GetEncodedReply(), m.CBusOptions.GetSrchk())
		if _chksumErr != nil {
			return errors.Wrap(_chksumErr, "Error serializing 'chksum' field")
		}
		// Virtual field
		chksumDecoded := m.GetChksumDecoded()
		_ = chksumDecoded
		if _chksumDecodedErr := writeBuffer.WriteVirtual(ctx, "chksumDecoded", m.GetChksumDecoded()); _chksumDecodedErr != nil {
			return errors.Wrap(_chksumDecodedErr, "Error serializing 'chksumDecoded' field")
		}

		if popErr := writeBuffer.PopContext("ReplyEncodedReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyEncodedReply")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReplyEncodedReply) isReplyEncodedReply() bool {
	return true
}

func (m *_ReplyEncodedReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
