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

// ReaderGroupTransportDataType is the corresponding interface of ReaderGroupTransportDataType
type ReaderGroupTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// ReaderGroupTransportDataTypeExactly can be used when we want exactly this type and not a type which fulfills ReaderGroupTransportDataType.
// This is useful for switch cases.
type ReaderGroupTransportDataTypeExactly interface {
	ReaderGroupTransportDataType
	isReaderGroupTransportDataType() bool
}

// _ReaderGroupTransportDataType is the data-structure of this message
type _ReaderGroupTransportDataType struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReaderGroupTransportDataType) GetIdentifier() string {
	return "15623"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReaderGroupTransportDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ReaderGroupTransportDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewReaderGroupTransportDataType factory function for _ReaderGroupTransportDataType
func NewReaderGroupTransportDataType() *_ReaderGroupTransportDataType {
	_result := &_ReaderGroupTransportDataType{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReaderGroupTransportDataType(structType any) ReaderGroupTransportDataType {
	if casted, ok := structType.(ReaderGroupTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReaderGroupTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReaderGroupTransportDataType) GetTypeName() string {
	return "ReaderGroupTransportDataType"
}

func (m *_ReaderGroupTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ReaderGroupTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReaderGroupTransportDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ReaderGroupTransportDataType, error) {
	return ReaderGroupTransportDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ReaderGroupTransportDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ReaderGroupTransportDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReaderGroupTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReaderGroupTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ReaderGroupTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReaderGroupTransportDataType")
	}

	// Create a partially initialized instance
	_child := &_ReaderGroupTransportDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ReaderGroupTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReaderGroupTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReaderGroupTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReaderGroupTransportDataType")
		}

		if popErr := writeBuffer.PopContext("ReaderGroupTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReaderGroupTransportDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReaderGroupTransportDataType) isReaderGroupTransportDataType() bool {
	return true
}

func (m *_ReaderGroupTransportDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
