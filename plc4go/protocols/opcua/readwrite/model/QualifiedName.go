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

// QualifiedName is the corresponding interface of QualifiedName
type QualifiedName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetName returns Name (property field)
	GetName() PascalString
}

// QualifiedNameExactly can be used when we want exactly this type and not a type which fulfills QualifiedName.
// This is useful for switch cases.
type QualifiedNameExactly interface {
	QualifiedName
	isQualifiedName() bool
}

// _QualifiedName is the data-structure of this message
type _QualifiedName struct {
	NamespaceIndex uint16
	Name           PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QualifiedName) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_QualifiedName) GetName() PascalString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQualifiedName factory function for _QualifiedName
func NewQualifiedName(namespaceIndex uint16, name PascalString) *_QualifiedName {
	return &_QualifiedName{NamespaceIndex: namespaceIndex, Name: name}
}

// Deprecated: use the interface for direct cast
func CastQualifiedName(structType any) QualifiedName {
	if casted, ok := structType.(QualifiedName); ok {
		return casted
	}
	if casted, ok := structType.(*QualifiedName); ok {
		return *casted
	}
	return nil
}

func (m *_QualifiedName) GetTypeName() string {
	return "QualifiedName"
}

func (m *_QualifiedName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QualifiedName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QualifiedNameParse(ctx context.Context, theBytes []byte) (QualifiedName, error) {
	return QualifiedNameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func QualifiedNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (QualifiedName, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QualifiedName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QualifiedName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint16("namespaceIndex", 16)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of QualifiedName")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of QualifiedName")
	}
	name := _name.(PascalString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	if closeErr := readBuffer.CloseContext("QualifiedName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QualifiedName")
	}

	// Create the instance
	return &_QualifiedName{
		NamespaceIndex: namespaceIndex,
		Name:           name,
	}, nil
}

func (m *_QualifiedName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QualifiedName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("QualifiedName"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for QualifiedName")
	}

	// Simple Field (namespaceIndex)
	namespaceIndex := uint16(m.GetNamespaceIndex())
	_namespaceIndexErr := writeBuffer.WriteUint16("namespaceIndex", 16, uint16((namespaceIndex)))
	if _namespaceIndexErr != nil {
		return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
	}

	// Simple Field (name)
	if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for name")
	}
	_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
	if popErr := writeBuffer.PopContext("name"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for name")
	}
	if _nameErr != nil {
		return errors.Wrap(_nameErr, "Error serializing 'name' field")
	}

	if popErr := writeBuffer.PopContext("QualifiedName"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for QualifiedName")
	}
	return nil
}

func (m *_QualifiedName) isQualifiedName() bool {
	return true
}

func (m *_QualifiedName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
