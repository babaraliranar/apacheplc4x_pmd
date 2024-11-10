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

// TransactionErrorType is the corresponding interface of TransactionErrorType
type TransactionErrorType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetTargetId returns TargetId (property field)
	GetTargetId() NodeId
	// GetError returns Error (property field)
	GetError() StatusCode
	// GetMessage returns Message (property field)
	GetMessage() LocalizedText
	// IsTransactionErrorType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTransactionErrorType()
	// CreateBuilder creates a TransactionErrorTypeBuilder
	CreateTransactionErrorTypeBuilder() TransactionErrorTypeBuilder
}

// _TransactionErrorType is the data-structure of this message
type _TransactionErrorType struct {
	ExtensionObjectDefinitionContract
	TargetId NodeId
	Error    StatusCode
	Message  LocalizedText
}

var _ TransactionErrorType = (*_TransactionErrorType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TransactionErrorType)(nil)

// NewTransactionErrorType factory function for _TransactionErrorType
func NewTransactionErrorType(targetId NodeId, error StatusCode, message LocalizedText) *_TransactionErrorType {
	if targetId == nil {
		panic("targetId of type NodeId for TransactionErrorType must not be nil")
	}
	if error == nil {
		panic("error of type StatusCode for TransactionErrorType must not be nil")
	}
	if message == nil {
		panic("message of type LocalizedText for TransactionErrorType must not be nil")
	}
	_result := &_TransactionErrorType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TargetId:                          targetId,
		Error:                             error,
		Message:                           message,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TransactionErrorTypeBuilder is a builder for TransactionErrorType
type TransactionErrorTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(targetId NodeId, error StatusCode, message LocalizedText) TransactionErrorTypeBuilder
	// WithTargetId adds TargetId (property field)
	WithTargetId(NodeId) TransactionErrorTypeBuilder
	// WithTargetIdBuilder adds TargetId (property field) which is build by the builder
	WithTargetIdBuilder(func(NodeIdBuilder) NodeIdBuilder) TransactionErrorTypeBuilder
	// WithError adds Error (property field)
	WithError(StatusCode) TransactionErrorTypeBuilder
	// WithErrorBuilder adds Error (property field) which is build by the builder
	WithErrorBuilder(func(StatusCodeBuilder) StatusCodeBuilder) TransactionErrorTypeBuilder
	// WithMessage adds Message (property field)
	WithMessage(LocalizedText) TransactionErrorTypeBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) TransactionErrorTypeBuilder
	// Build builds the TransactionErrorType or returns an error if something is wrong
	Build() (TransactionErrorType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TransactionErrorType
}

// NewTransactionErrorTypeBuilder() creates a TransactionErrorTypeBuilder
func NewTransactionErrorTypeBuilder() TransactionErrorTypeBuilder {
	return &_TransactionErrorTypeBuilder{_TransactionErrorType: new(_TransactionErrorType)}
}

type _TransactionErrorTypeBuilder struct {
	*_TransactionErrorType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TransactionErrorTypeBuilder) = (*_TransactionErrorTypeBuilder)(nil)

func (b *_TransactionErrorTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_TransactionErrorTypeBuilder) WithMandatoryFields(targetId NodeId, error StatusCode, message LocalizedText) TransactionErrorTypeBuilder {
	return b.WithTargetId(targetId).WithError(error).WithMessage(message)
}

func (b *_TransactionErrorTypeBuilder) WithTargetId(targetId NodeId) TransactionErrorTypeBuilder {
	b.TargetId = targetId
	return b
}

func (b *_TransactionErrorTypeBuilder) WithTargetIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) TransactionErrorTypeBuilder {
	builder := builderSupplier(b.TargetId.CreateNodeIdBuilder())
	var err error
	b.TargetId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_TransactionErrorTypeBuilder) WithError(error StatusCode) TransactionErrorTypeBuilder {
	b.Error = error
	return b
}

func (b *_TransactionErrorTypeBuilder) WithErrorBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) TransactionErrorTypeBuilder {
	builder := builderSupplier(b.Error.CreateStatusCodeBuilder())
	var err error
	b.Error, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_TransactionErrorTypeBuilder) WithMessage(message LocalizedText) TransactionErrorTypeBuilder {
	b.Message = message
	return b
}

func (b *_TransactionErrorTypeBuilder) WithMessageBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) TransactionErrorTypeBuilder {
	builder := builderSupplier(b.Message.CreateLocalizedTextBuilder())
	var err error
	b.Message, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_TransactionErrorTypeBuilder) Build() (TransactionErrorType, error) {
	if b.TargetId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'targetId' not set"))
	}
	if b.Error == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'error' not set"))
	}
	if b.Message == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TransactionErrorType.deepCopy(), nil
}

func (b *_TransactionErrorTypeBuilder) MustBuild() TransactionErrorType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TransactionErrorTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_TransactionErrorTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TransactionErrorTypeBuilder) DeepCopy() any {
	_copy := b.CreateTransactionErrorTypeBuilder().(*_TransactionErrorTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTransactionErrorTypeBuilder creates a TransactionErrorTypeBuilder
func (b *_TransactionErrorType) CreateTransactionErrorTypeBuilder() TransactionErrorTypeBuilder {
	if b == nil {
		return NewTransactionErrorTypeBuilder()
	}
	return &_TransactionErrorTypeBuilder{_TransactionErrorType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransactionErrorType) GetExtensionId() int32 {
	return int32(32287)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransactionErrorType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransactionErrorType) GetTargetId() NodeId {
	return m.TargetId
}

func (m *_TransactionErrorType) GetError() StatusCode {
	return m.Error
}

func (m *_TransactionErrorType) GetMessage() LocalizedText {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTransactionErrorType(structType any) TransactionErrorType {
	if casted, ok := structType.(TransactionErrorType); ok {
		return casted
	}
	if casted, ok := structType.(*TransactionErrorType); ok {
		return *casted
	}
	return nil
}

func (m *_TransactionErrorType) GetTypeName() string {
	return "TransactionErrorType"
}

func (m *_TransactionErrorType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (targetId)
	lengthInBits += m.TargetId.GetLengthInBits(ctx)

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TransactionErrorType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TransactionErrorType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__transactionErrorType TransactionErrorType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransactionErrorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransactionErrorType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	targetId, err := ReadSimpleField[NodeId](ctx, "targetId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetId' field"))
	}
	m.TargetId = targetId

	error, err := ReadSimpleField[StatusCode](ctx, "error", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	message, err := ReadSimpleField[LocalizedText](ctx, "message", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("TransactionErrorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransactionErrorType")
	}

	return m, nil
}

func (m *_TransactionErrorType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransactionErrorType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransactionErrorType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransactionErrorType")
		}

		if err := WriteSimpleField[NodeId](ctx, "targetId", m.GetTargetId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetId' field")
		}

		if err := WriteSimpleField[StatusCode](ctx, "error", m.GetError(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'error' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "message", m.GetMessage(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("TransactionErrorType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransactionErrorType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransactionErrorType) IsTransactionErrorType() {}

func (m *_TransactionErrorType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TransactionErrorType) deepCopy() *_TransactionErrorType {
	if m == nil {
		return nil
	}
	_TransactionErrorTypeCopy := &_TransactionErrorType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.TargetId.DeepCopy().(NodeId),
		m.Error.DeepCopy().(StatusCode),
		m.Message.DeepCopy().(LocalizedText),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TransactionErrorTypeCopy
}

func (m *_TransactionErrorType) String() string {
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
