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

// NodeIdTypeDefinition is the corresponding interface of NodeIdTypeDefinition
type NodeIdTypeDefinition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNodeType returns NodeType (discriminator field)
	GetNodeType() NodeIdType
	// GetIdentifier returns Identifier (abstract field)
	GetIdentifier() string
}

// NodeIdTypeDefinitionExactly can be used when we want exactly this type and not a type which fulfills NodeIdTypeDefinition.
// This is useful for switch cases.
type NodeIdTypeDefinitionExactly interface {
	NodeIdTypeDefinition
	isNodeIdTypeDefinition() bool
}

// _NodeIdTypeDefinition is the data-structure of this message
type _NodeIdTypeDefinition struct {
	_NodeIdTypeDefinitionChildRequirements
}

type _NodeIdTypeDefinitionChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetNodeType() NodeIdType
	// GetIdentifier returns Identifier (abstract field)
	GetIdentifier() string
}

type NodeIdTypeDefinitionParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NodeIdTypeDefinition, serializeChildFunction func() error) error
	GetTypeName() string
}

type NodeIdTypeDefinitionChild interface {
	utils.Serializable
	InitializeParent(parent NodeIdTypeDefinition)
	GetParent() *NodeIdTypeDefinition

	GetTypeName() string
	NodeIdTypeDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for abstract fields.
///////////////////////

func (m *_NodeIdTypeDefinition) GetIdentifier() string {
	return m._NodeIdTypeDefinitionChildRequirements.GetIdentifier()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdTypeDefinition factory function for _NodeIdTypeDefinition
func NewNodeIdTypeDefinition() *_NodeIdTypeDefinition {
	return &_NodeIdTypeDefinition{}
}

// Deprecated: use the interface for direct cast
func CastNodeIdTypeDefinition(structType any) NodeIdTypeDefinition {
	if casted, ok := structType.(NodeIdTypeDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdTypeDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdTypeDefinition) GetTypeName() string {
	return "NodeIdTypeDefinition"
}

func (m *_NodeIdTypeDefinition) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (nodeType)
	lengthInBits += 6

	return lengthInBits
}

func (m *_NodeIdTypeDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdTypeDefinitionParse(ctx context.Context, theBytes []byte) (NodeIdTypeDefinition, error) {
	return NodeIdTypeDefinitionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdTypeDefinitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdTypeDefinition, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeIdTypeDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdTypeDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (nodeType) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("nodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeType")
	}
	nodeType_temp, _nodeTypeErr := NodeIdTypeParseWithBuffer(ctx, readBuffer)
	var nodeType NodeIdType = nodeType_temp
	if closeErr := readBuffer.CloseContext("nodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeType")
	}
	if _nodeTypeErr != nil {
		return nil, errors.Wrap(_nodeTypeErr, "Error parsing 'nodeType' field of NodeIdTypeDefinition")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type NodeIdTypeDefinitionChildSerializeRequirement interface {
		NodeIdTypeDefinition
		InitializeParent(NodeIdTypeDefinition)
		GetParent() NodeIdTypeDefinition
	}
	var _childTemp any
	var _child NodeIdTypeDefinitionChildSerializeRequirement
	var typeSwitchError error
	switch {
	case nodeType == NodeIdType_nodeIdTypeTwoByte: // NodeIdTwoByte
		_childTemp, typeSwitchError = NodeIdTwoByteParseWithBuffer(ctx, readBuffer)
	case nodeType == NodeIdType_nodeIdTypeFourByte: // NodeIdFourByte
		_childTemp, typeSwitchError = NodeIdFourByteParseWithBuffer(ctx, readBuffer)
	case nodeType == NodeIdType_nodeIdTypeNumeric: // NodeIdNumeric
		_childTemp, typeSwitchError = NodeIdNumericParseWithBuffer(ctx, readBuffer)
	case nodeType == NodeIdType_nodeIdTypeString: // NodeIdString
		_childTemp, typeSwitchError = NodeIdStringParseWithBuffer(ctx, readBuffer)
	case nodeType == NodeIdType_nodeIdTypeGuid: // NodeIdGuid
		_childTemp, typeSwitchError = NodeIdGuidParseWithBuffer(ctx, readBuffer)
	case nodeType == NodeIdType_nodeIdTypeByteString: // NodeIdByteString
		_childTemp, typeSwitchError = NodeIdByteStringParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [nodeType=%v]", nodeType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of NodeIdTypeDefinition")
	}
	_child = _childTemp.(NodeIdTypeDefinitionChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("NodeIdTypeDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdTypeDefinition")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_NodeIdTypeDefinition) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NodeIdTypeDefinition, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NodeIdTypeDefinition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NodeIdTypeDefinition")
	}

	// Discriminator Field (nodeType) (Used as input to a switch field)
	nodeType := NodeIdType(child.GetNodeType())
	if pushErr := writeBuffer.PushContext("nodeType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for nodeType")
	}
	_nodeTypeErr := writeBuffer.WriteSerializable(ctx, nodeType)
	if popErr := writeBuffer.PopContext("nodeType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for nodeType")
	}

	if _nodeTypeErr != nil {
		return errors.Wrap(_nodeTypeErr, "Error serializing 'nodeType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NodeIdTypeDefinition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NodeIdTypeDefinition")
	}
	return nil
}

func (m *_NodeIdTypeDefinition) isNodeIdTypeDefinition() bool {
	return true
}

func (m *_NodeIdTypeDefinition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
