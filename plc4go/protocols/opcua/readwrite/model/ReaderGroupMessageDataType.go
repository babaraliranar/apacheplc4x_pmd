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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ReaderGroupMessageDataType is the corresponding interface of ReaderGroupMessageDataType
type ReaderGroupMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsReaderGroupMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReaderGroupMessageDataType()
	// CreateBuilder creates a ReaderGroupMessageDataTypeBuilder
	CreateReaderGroupMessageDataTypeBuilder() ReaderGroupMessageDataTypeBuilder
}

// _ReaderGroupMessageDataType is the data-structure of this message
type _ReaderGroupMessageDataType struct {
	ExtensionObjectDefinitionContract
}

var _ ReaderGroupMessageDataType = (*_ReaderGroupMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReaderGroupMessageDataType)(nil)

// NewReaderGroupMessageDataType factory function for _ReaderGroupMessageDataType
func NewReaderGroupMessageDataType() *_ReaderGroupMessageDataType {
	_result := &_ReaderGroupMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReaderGroupMessageDataTypeBuilder is a builder for ReaderGroupMessageDataType
type ReaderGroupMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ReaderGroupMessageDataTypeBuilder
	// Build builds the ReaderGroupMessageDataType or returns an error if something is wrong
	Build() (ReaderGroupMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReaderGroupMessageDataType
}

// NewReaderGroupMessageDataTypeBuilder() creates a ReaderGroupMessageDataTypeBuilder
func NewReaderGroupMessageDataTypeBuilder() ReaderGroupMessageDataTypeBuilder {
	return &_ReaderGroupMessageDataTypeBuilder{_ReaderGroupMessageDataType: new(_ReaderGroupMessageDataType)}
}

type _ReaderGroupMessageDataTypeBuilder struct {
	*_ReaderGroupMessageDataType

	err *utils.MultiError
}

var _ (ReaderGroupMessageDataTypeBuilder) = (*_ReaderGroupMessageDataTypeBuilder)(nil)

func (m *_ReaderGroupMessageDataTypeBuilder) WithMandatoryFields() ReaderGroupMessageDataTypeBuilder {
	return m
}

func (m *_ReaderGroupMessageDataTypeBuilder) Build() (ReaderGroupMessageDataType, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ReaderGroupMessageDataType.deepCopy(), nil
}

func (m *_ReaderGroupMessageDataTypeBuilder) MustBuild() ReaderGroupMessageDataType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ReaderGroupMessageDataTypeBuilder) DeepCopy() any {
	return m.CreateReaderGroupMessageDataTypeBuilder()
}

// CreateReaderGroupMessageDataTypeBuilder creates a ReaderGroupMessageDataTypeBuilder
func (m *_ReaderGroupMessageDataType) CreateReaderGroupMessageDataTypeBuilder() ReaderGroupMessageDataTypeBuilder {
	if m == nil {
		return NewReaderGroupMessageDataTypeBuilder()
	}
	return &_ReaderGroupMessageDataTypeBuilder{_ReaderGroupMessageDataType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReaderGroupMessageDataType) GetIdentifier() string {
	return "15624"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReaderGroupMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastReaderGroupMessageDataType(structType any) ReaderGroupMessageDataType {
	if casted, ok := structType.(ReaderGroupMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReaderGroupMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReaderGroupMessageDataType) GetTypeName() string {
	return "ReaderGroupMessageDataType"
}

func (m *_ReaderGroupMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ReaderGroupMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReaderGroupMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__readerGroupMessageDataType ReaderGroupMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReaderGroupMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReaderGroupMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ReaderGroupMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReaderGroupMessageDataType")
	}

	return m, nil
}

func (m *_ReaderGroupMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReaderGroupMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReaderGroupMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReaderGroupMessageDataType")
		}

		if popErr := writeBuffer.PopContext("ReaderGroupMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReaderGroupMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReaderGroupMessageDataType) IsReaderGroupMessageDataType() {}

func (m *_ReaderGroupMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReaderGroupMessageDataType) deepCopy() *_ReaderGroupMessageDataType {
	if m == nil {
		return nil
	}
	_ReaderGroupMessageDataTypeCopy := &_ReaderGroupMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReaderGroupMessageDataTypeCopy
}

func (m *_ReaderGroupMessageDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
