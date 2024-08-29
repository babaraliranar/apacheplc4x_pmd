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

// DeleteNodesItem is the corresponding interface of DeleteNodesItem
type DeleteNodesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetDeleteTargetReferences returns DeleteTargetReferences (property field)
	GetDeleteTargetReferences() bool
}

// DeleteNodesItemExactly can be used when we want exactly this type and not a type which fulfills DeleteNodesItem.
// This is useful for switch cases.
type DeleteNodesItemExactly interface {
	DeleteNodesItem
	isDeleteNodesItem() bool
}

// _DeleteNodesItem is the data-structure of this message
type _DeleteNodesItem struct {
	*_ExtensionObjectDefinition
	NodeId                 NodeId
	DeleteTargetReferences bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteNodesItem) GetIdentifier() string {
	return "384"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteNodesItem) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteNodesItem) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteNodesItem) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_DeleteNodesItem) GetDeleteTargetReferences() bool {
	return m.DeleteTargetReferences
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteNodesItem factory function for _DeleteNodesItem
func NewDeleteNodesItem(nodeId NodeId, deleteTargetReferences bool) *_DeleteNodesItem {
	_result := &_DeleteNodesItem{
		NodeId:                     nodeId,
		DeleteTargetReferences:     deleteTargetReferences,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteNodesItem(structType any) DeleteNodesItem {
	if casted, ok := structType.(DeleteNodesItem); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteNodesItem); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteNodesItem) GetTypeName() string {
	return "DeleteNodesItem"
}

func (m *_DeleteNodesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteTargetReferences)
	lengthInBits += 1

	return lengthInBits
}

func (m *_DeleteNodesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteNodesItemParse(ctx context.Context, theBytes []byte, identifier string) (DeleteNodesItem, error) {
	return DeleteNodesItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteNodesItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteNodesItem, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeleteNodesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteNodesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of DeleteNodesItem")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DeleteNodesItem")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (deleteTargetReferences)
	_deleteTargetReferences, _deleteTargetReferencesErr := readBuffer.ReadBit("deleteTargetReferences")
	if _deleteTargetReferencesErr != nil {
		return nil, errors.Wrap(_deleteTargetReferencesErr, "Error parsing 'deleteTargetReferences' field of DeleteNodesItem")
	}
	deleteTargetReferences := _deleteTargetReferences

	if closeErr := readBuffer.CloseContext("DeleteNodesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteNodesItem")
	}

	// Create a partially initialized instance
	_child := &_DeleteNodesItem{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NodeId:                     nodeId,
		DeleteTargetReferences:     deleteTargetReferences,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteNodesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteNodesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteNodesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteNodesItem")
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

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (deleteTargetReferences)
		deleteTargetReferences := bool(m.GetDeleteTargetReferences())
		_deleteTargetReferencesErr := writeBuffer.WriteBit("deleteTargetReferences", (deleteTargetReferences))
		if _deleteTargetReferencesErr != nil {
			return errors.Wrap(_deleteTargetReferencesErr, "Error serializing 'deleteTargetReferences' field")
		}

		if popErr := writeBuffer.PopContext("DeleteNodesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteNodesItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteNodesItem) isDeleteNodesItem() bool {
	return true
}

func (m *_DeleteNodesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
