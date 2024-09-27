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

// ReplyOrConfirmationReply is the corresponding interface of ReplyOrConfirmationReply
type ReplyOrConfirmationReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ReplyOrConfirmation
	// GetReply returns Reply (property field)
	GetReply() Reply
	// GetTermination returns Termination (property field)
	GetTermination() ResponseTermination
	// IsReplyOrConfirmationReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReplyOrConfirmationReply()
	// CreateBuilder creates a ReplyOrConfirmationReplyBuilder
	CreateReplyOrConfirmationReplyBuilder() ReplyOrConfirmationReplyBuilder
}

// _ReplyOrConfirmationReply is the data-structure of this message
type _ReplyOrConfirmationReply struct {
	ReplyOrConfirmationContract
	Reply       Reply
	Termination ResponseTermination
}

var _ ReplyOrConfirmationReply = (*_ReplyOrConfirmationReply)(nil)
var _ ReplyOrConfirmationRequirements = (*_ReplyOrConfirmationReply)(nil)

// NewReplyOrConfirmationReply factory function for _ReplyOrConfirmationReply
func NewReplyOrConfirmationReply(peekedByte byte, reply Reply, termination ResponseTermination, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmationReply {
	if reply == nil {
		panic("reply of type Reply for ReplyOrConfirmationReply must not be nil")
	}
	if termination == nil {
		panic("termination of type ResponseTermination for ReplyOrConfirmationReply must not be nil")
	}
	_result := &_ReplyOrConfirmationReply{
		ReplyOrConfirmationContract: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
		Reply:                       reply,
		Termination:                 termination,
	}
	_result.ReplyOrConfirmationContract.(*_ReplyOrConfirmation)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReplyOrConfirmationReplyBuilder is a builder for ReplyOrConfirmationReply
type ReplyOrConfirmationReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reply Reply, termination ResponseTermination) ReplyOrConfirmationReplyBuilder
	// WithReply adds Reply (property field)
	WithReply(Reply) ReplyOrConfirmationReplyBuilder
	// WithTermination adds Termination (property field)
	WithTermination(ResponseTermination) ReplyOrConfirmationReplyBuilder
	// WithTerminationBuilder adds Termination (property field) which is build by the builder
	WithTerminationBuilder(func(ResponseTerminationBuilder) ResponseTerminationBuilder) ReplyOrConfirmationReplyBuilder
	// Build builds the ReplyOrConfirmationReply or returns an error if something is wrong
	Build() (ReplyOrConfirmationReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReplyOrConfirmationReply
}

// NewReplyOrConfirmationReplyBuilder() creates a ReplyOrConfirmationReplyBuilder
func NewReplyOrConfirmationReplyBuilder() ReplyOrConfirmationReplyBuilder {
	return &_ReplyOrConfirmationReplyBuilder{_ReplyOrConfirmationReply: new(_ReplyOrConfirmationReply)}
}

type _ReplyOrConfirmationReplyBuilder struct {
	*_ReplyOrConfirmationReply

	err *utils.MultiError
}

var _ (ReplyOrConfirmationReplyBuilder) = (*_ReplyOrConfirmationReplyBuilder)(nil)

func (m *_ReplyOrConfirmationReplyBuilder) WithMandatoryFields(reply Reply, termination ResponseTermination) ReplyOrConfirmationReplyBuilder {
	return m.WithReply(reply).WithTermination(termination)
}

func (m *_ReplyOrConfirmationReplyBuilder) WithReply(reply Reply) ReplyOrConfirmationReplyBuilder {
	m.Reply = reply
	return m
}

func (m *_ReplyOrConfirmationReplyBuilder) WithTermination(termination ResponseTermination) ReplyOrConfirmationReplyBuilder {
	m.Termination = termination
	return m
}

func (m *_ReplyOrConfirmationReplyBuilder) WithTerminationBuilder(builderSupplier func(ResponseTerminationBuilder) ResponseTerminationBuilder) ReplyOrConfirmationReplyBuilder {
	builder := builderSupplier(m.Termination.CreateResponseTerminationBuilder())
	var err error
	m.Termination, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "ResponseTerminationBuilder failed"))
	}
	return m
}

func (m *_ReplyOrConfirmationReplyBuilder) Build() (ReplyOrConfirmationReply, error) {
	if m.Reply == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'reply' not set"))
	}
	if m.Termination == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'termination' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ReplyOrConfirmationReply.deepCopy(), nil
}

func (m *_ReplyOrConfirmationReplyBuilder) MustBuild() ReplyOrConfirmationReply {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ReplyOrConfirmationReplyBuilder) DeepCopy() any {
	return m.CreateReplyOrConfirmationReplyBuilder()
}

// CreateReplyOrConfirmationReplyBuilder creates a ReplyOrConfirmationReplyBuilder
func (m *_ReplyOrConfirmationReply) CreateReplyOrConfirmationReplyBuilder() ReplyOrConfirmationReplyBuilder {
	if m == nil {
		return NewReplyOrConfirmationReplyBuilder()
	}
	return &_ReplyOrConfirmationReplyBuilder{_ReplyOrConfirmationReply: m.deepCopy()}
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

func (m *_ReplyOrConfirmationReply) GetParent() ReplyOrConfirmationContract {
	return m.ReplyOrConfirmationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmationReply) GetReply() Reply {
	return m.Reply
}

func (m *_ReplyOrConfirmationReply) GetTermination() ResponseTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmationReply(structType any) ReplyOrConfirmationReply {
	if casted, ok := structType.(ReplyOrConfirmationReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmationReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmationReply) GetTypeName() string {
	return "ReplyOrConfirmationReply"
}

func (m *_ReplyOrConfirmationReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).getLengthInBits(ctx))

	// Simple field (reply)
	lengthInBits += m.Reply.GetLengthInBits(ctx)

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReplyOrConfirmationReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReplyOrConfirmationReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ReplyOrConfirmation, cBusOptions CBusOptions, requestContext RequestContext) (__replyOrConfirmationReply ReplyOrConfirmationReply, err error) {
	m.ReplyOrConfirmationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyOrConfirmationReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmationReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reply, err := ReadSimpleField[Reply](ctx, "reply", ReadComplex[Reply](ReplyParseWithBufferProducer[Reply]((CBusOptions)(cBusOptions), (RequestContext)(requestContext)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reply' field"))
	}
	m.Reply = reply

	termination, err := ReadSimpleField[ResponseTermination](ctx, "termination", ReadComplex[ResponseTermination](ResponseTerminationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'termination' field"))
	}
	m.Termination = termination

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmationReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmationReply")
	}

	return m, nil
}

func (m *_ReplyOrConfirmationReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyOrConfirmationReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyOrConfirmationReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmationReply")
		}

		if err := WriteSimpleField[Reply](ctx, "reply", m.GetReply(), WriteComplex[Reply](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reply' field")
		}

		if err := WriteSimpleField[ResponseTermination](ctx, "termination", m.GetTermination(), WriteComplex[ResponseTermination](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'termination' field")
		}

		if popErr := writeBuffer.PopContext("ReplyOrConfirmationReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyOrConfirmationReply")
		}
		return nil
	}
	return m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReplyOrConfirmationReply) IsReplyOrConfirmationReply() {}

func (m *_ReplyOrConfirmationReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReplyOrConfirmationReply) deepCopy() *_ReplyOrConfirmationReply {
	if m == nil {
		return nil
	}
	_ReplyOrConfirmationReplyCopy := &_ReplyOrConfirmationReply{
		m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).deepCopy(),
		m.Reply.DeepCopy().(Reply),
		m.Termination.DeepCopy().(ResponseTermination),
	}
	m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation)._SubType = m
	return _ReplyOrConfirmationReplyCopy
}

func (m *_ReplyOrConfirmationReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
