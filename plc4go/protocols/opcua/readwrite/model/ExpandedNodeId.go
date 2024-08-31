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

// ExpandedNodeId is the corresponding interface of ExpandedNodeId
type ExpandedNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceURISpecified returns NamespaceURISpecified (property field)
	GetNamespaceURISpecified() bool
	// GetServerIndexSpecified returns ServerIndexSpecified (property field)
	GetServerIndexSpecified() bool
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeIdTypeDefinition
	// GetNamespaceURI returns NamespaceURI (property field)
	GetNamespaceURI() PascalString
	// GetServerIndex returns ServerIndex (property field)
	GetServerIndex() *uint32
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// ExpandedNodeIdExactly can be used when we want exactly this type and not a type which fulfills ExpandedNodeId.
// This is useful for switch cases.
type ExpandedNodeIdExactly interface {
	ExpandedNodeId
	isExpandedNodeId() bool
}

// _ExpandedNodeId is the data-structure of this message
type _ExpandedNodeId struct {
	NamespaceURISpecified bool
	ServerIndexSpecified  bool
	NodeId                NodeIdTypeDefinition
	NamespaceURI          PascalString
	ServerIndex           *uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExpandedNodeId) GetNamespaceURISpecified() bool {
	return m.NamespaceURISpecified
}

func (m *_ExpandedNodeId) GetServerIndexSpecified() bool {
	return m.ServerIndexSpecified
}

func (m *_ExpandedNodeId) GetNodeId() NodeIdTypeDefinition {
	return m.NodeId
}

func (m *_ExpandedNodeId) GetNamespaceURI() PascalString {
	return m.NamespaceURI
}

func (m *_ExpandedNodeId) GetServerIndex() *uint32 {
	return m.ServerIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ExpandedNodeId) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	namespaceURI := m.NamespaceURI
	_ = namespaceURI
	serverIndex := m.ServerIndex
	_ = serverIndex
	return fmt.Sprintf("%v", m.GetNodeId().GetIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExpandedNodeId factory function for _ExpandedNodeId
func NewExpandedNodeId(namespaceURISpecified bool, serverIndexSpecified bool, nodeId NodeIdTypeDefinition, namespaceURI PascalString, serverIndex *uint32) *_ExpandedNodeId {
	return &_ExpandedNodeId{NamespaceURISpecified: namespaceURISpecified, ServerIndexSpecified: serverIndexSpecified, NodeId: nodeId, NamespaceURI: namespaceURI, ServerIndex: serverIndex}
}

// Deprecated: use the interface for direct cast
func CastExpandedNodeId(structType any) ExpandedNodeId {
	if casted, ok := structType.(ExpandedNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*ExpandedNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_ExpandedNodeId) GetTypeName() string {
	return "ExpandedNodeId"
}

func (m *_ExpandedNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceURISpecified)
	lengthInBits += 1

	// Simple field (serverIndexSpecified)
	lengthInBits += 1

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Optional Field (namespaceURI)
	if m.NamespaceURI != nil {
		lengthInBits += m.NamespaceURI.GetLengthInBits(ctx)
	}

	// Optional Field (serverIndex)
	if m.ServerIndex != nil {
		lengthInBits += 32
	}

	return lengthInBits
}

func (m *_ExpandedNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ExpandedNodeIdParse(ctx context.Context, theBytes []byte) (ExpandedNodeId, error) {
	return ExpandedNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ExpandedNodeIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ExpandedNodeId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ExpandedNodeId, error) {
		return ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	}
}

func ExpandedNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ExpandedNodeId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ExpandedNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExpandedNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceURISpecified, err := ReadSimpleField(ctx, "namespaceURISpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceURISpecified' field"))
	}

	serverIndexSpecified, err := ReadSimpleField(ctx, "serverIndexSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverIndexSpecified' field"))
	}

	nodeId, err := ReadSimpleField[NodeIdTypeDefinition](ctx, "nodeId", ReadComplex[NodeIdTypeDefinition](NodeIdTypeDefinitionParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), nodeId.GetIdentifier())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	_namespaceURI, err := ReadOptionalField[PascalString](ctx, "namespaceURI", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), namespaceURISpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceURI' field"))
	}
	var namespaceURI PascalString
	if _namespaceURI != nil {
		namespaceURI = *_namespaceURI
	}

	serverIndex, err := ReadOptionalField[uint32](ctx, "serverIndex", ReadUnsignedInt(readBuffer, uint8(32)), serverIndexSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverIndex' field"))
	}

	if closeErr := readBuffer.CloseContext("ExpandedNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExpandedNodeId")
	}

	// Create the instance
	return &_ExpandedNodeId{
		NamespaceURISpecified: namespaceURISpecified,
		ServerIndexSpecified:  serverIndexSpecified,
		NodeId:                nodeId,
		NamespaceURI:          namespaceURI,
		ServerIndex:           serverIndex,
	}, nil
}

func (m *_ExpandedNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExpandedNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ExpandedNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExpandedNodeId")
	}

	// Simple Field (namespaceURISpecified)
	namespaceURISpecified := bool(m.GetNamespaceURISpecified())
	_namespaceURISpecifiedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("namespaceURISpecified", (namespaceURISpecified))
	if _namespaceURISpecifiedErr != nil {
		return errors.Wrap(_namespaceURISpecifiedErr, "Error serializing 'namespaceURISpecified' field")
	}

	// Simple Field (serverIndexSpecified)
	serverIndexSpecified := bool(m.GetServerIndexSpecified())
	_serverIndexSpecifiedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("serverIndexSpecified", (serverIndexSpecified))
	if _serverIndexSpecifiedErr != nil {
		return errors.Wrap(_serverIndexSpecifiedErr, "Error serializing 'serverIndexSpecified' field")
	}

	// Simple Field (nodeId)
	if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for nodeId")
	}
	_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
	if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for nodeId")
	}
	if _nodeIdErr != nil {
		return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
	}
	// Virtual field
	identifier := m.GetIdentifier()
	_ = identifier
	if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
		return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
	}

	// Optional Field (namespaceURI) (Can be skipped, if the value is null)
	var namespaceURI PascalString = nil
	if m.GetNamespaceURI() != nil {
		if pushErr := writeBuffer.PushContext("namespaceURI"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for namespaceURI")
		}
		namespaceURI = m.GetNamespaceURI()
		_namespaceURIErr := writeBuffer.WriteSerializable(ctx, namespaceURI)
		if popErr := writeBuffer.PopContext("namespaceURI"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for namespaceURI")
		}
		if _namespaceURIErr != nil {
			return errors.Wrap(_namespaceURIErr, "Error serializing 'namespaceURI' field")
		}
	}

	// Optional Field (serverIndex) (Can be skipped, if the value is null)
	var serverIndex *uint32 = nil
	if m.GetServerIndex() != nil {
		serverIndex = m.GetServerIndex()
		_serverIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("serverIndex", 32, uint32(*(serverIndex)))
		if _serverIndexErr != nil {
			return errors.Wrap(_serverIndexErr, "Error serializing 'serverIndex' field")
		}
	}

	if popErr := writeBuffer.PopContext("ExpandedNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExpandedNodeId")
	}
	return nil
}

func (m *_ExpandedNodeId) isExpandedNodeId() bool {
	return true
}

func (m *_ExpandedNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
