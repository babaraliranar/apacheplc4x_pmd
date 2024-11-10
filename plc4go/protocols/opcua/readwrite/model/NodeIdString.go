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

// NodeIdString is the corresponding interface of NodeIdString
type NodeIdString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() PascalString
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
	// GetNamespace returns Namespace (virtual field)
	GetNamespace() int16
	// IsNodeIdString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdString()
	// CreateBuilder creates a NodeIdStringBuilder
	CreateNodeIdStringBuilder() NodeIdStringBuilder
}

// _NodeIdString is the data-structure of this message
type _NodeIdString struct {
	NodeIdTypeDefinitionContract
	NamespaceIndex uint16
	Id             PascalString
}

var _ NodeIdString = (*_NodeIdString)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdString)(nil)

// NewNodeIdString factory function for _NodeIdString
func NewNodeIdString(namespaceIndex uint16, id PascalString) *_NodeIdString {
	if id == nil {
		panic("id of type PascalString for NodeIdString must not be nil")
	}
	_result := &_NodeIdString{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		NamespaceIndex:               namespaceIndex,
		Id:                           id,
	}
	_result.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeIdStringBuilder is a builder for NodeIdString
type NodeIdStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaceIndex uint16, id PascalString) NodeIdStringBuilder
	// WithNamespaceIndex adds NamespaceIndex (property field)
	WithNamespaceIndex(uint16) NodeIdStringBuilder
	// WithId adds Id (property field)
	WithId(PascalString) NodeIdStringBuilder
	// WithIdBuilder adds Id (property field) which is build by the builder
	WithIdBuilder(func(PascalStringBuilder) PascalStringBuilder) NodeIdStringBuilder
	// Build builds the NodeIdString or returns an error if something is wrong
	Build() (NodeIdString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeIdString
}

// NewNodeIdStringBuilder() creates a NodeIdStringBuilder
func NewNodeIdStringBuilder() NodeIdStringBuilder {
	return &_NodeIdStringBuilder{_NodeIdString: new(_NodeIdString)}
}

type _NodeIdStringBuilder struct {
	*_NodeIdString

	parentBuilder *_NodeIdTypeDefinitionBuilder

	err *utils.MultiError
}

var _ (NodeIdStringBuilder) = (*_NodeIdStringBuilder)(nil)

func (b *_NodeIdStringBuilder) setParent(contract NodeIdTypeDefinitionContract) {
	b.NodeIdTypeDefinitionContract = contract
}

func (b *_NodeIdStringBuilder) WithMandatoryFields(namespaceIndex uint16, id PascalString) NodeIdStringBuilder {
	return b.WithNamespaceIndex(namespaceIndex).WithId(id)
}

func (b *_NodeIdStringBuilder) WithNamespaceIndex(namespaceIndex uint16) NodeIdStringBuilder {
	b.NamespaceIndex = namespaceIndex
	return b
}

func (b *_NodeIdStringBuilder) WithId(id PascalString) NodeIdStringBuilder {
	b.Id = id
	return b
}

func (b *_NodeIdStringBuilder) WithIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) NodeIdStringBuilder {
	builder := builderSupplier(b.Id.CreatePascalStringBuilder())
	var err error
	b.Id, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_NodeIdStringBuilder) Build() (NodeIdString, error) {
	if b.Id == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'id' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeIdString.deepCopy(), nil
}

func (b *_NodeIdStringBuilder) MustBuild() NodeIdString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NodeIdStringBuilder) Done() NodeIdTypeDefinitionBuilder {
	return b.parentBuilder
}

func (b *_NodeIdStringBuilder) buildForNodeIdTypeDefinition() (NodeIdTypeDefinition, error) {
	return b.Build()
}

func (b *_NodeIdStringBuilder) DeepCopy() any {
	_copy := b.CreateNodeIdStringBuilder().(*_NodeIdStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeIdStringBuilder creates a NodeIdStringBuilder
func (b *_NodeIdString) CreateNodeIdStringBuilder() NodeIdStringBuilder {
	if b == nil {
		return NewNodeIdStringBuilder()
	}
	return &_NodeIdStringBuilder{_NodeIdString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdString) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdString) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdString) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdString) GetId() PascalString {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdString) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId().GetStringValue())
}

func (m *_NodeIdString) GetNamespace() int16 {
	ctx := context.Background()
	_ = ctx
	return int16(m.GetNamespaceIndex())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeIdString(structType any) NodeIdString {
	if casted, ok := structType.(NodeIdString); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdString); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdString) GetTypeName() string {
	return "NodeIdString"
}

func (m *_NodeIdString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (id)
	lengthInBits += m.Id.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdString NodeIdString, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	id, err := ReadSimpleField[PascalString](ctx, "id", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id.GetStringValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	namespace, err := ReadVirtualField[int16](ctx, "namespace", (*int16)(nil), namespaceIndex)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespace' field"))
	}
	_ = namespace

	if closeErr := readBuffer.CloseContext("NodeIdString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdString")
	}

	return m, nil
}

func (m *_NodeIdString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdString")
		}

		if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "id", m.GetId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}
		// Virtual field
		namespace := m.GetNamespace()
		_ = namespace
		if _namespaceErr := writeBuffer.WriteVirtual(ctx, "namespace", m.GetNamespace()); _namespaceErr != nil {
			return errors.Wrap(_namespaceErr, "Error serializing 'namespace' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdString")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdString) IsNodeIdString() {}

func (m *_NodeIdString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeIdString) deepCopy() *_NodeIdString {
	if m == nil {
		return nil
	}
	_NodeIdStringCopy := &_NodeIdString{
		m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).deepCopy(),
		m.NamespaceIndex,
		m.Id.DeepCopy().(PascalString),
	}
	m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = m
	return _NodeIdStringCopy
}

func (m *_NodeIdString) String() string {
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
