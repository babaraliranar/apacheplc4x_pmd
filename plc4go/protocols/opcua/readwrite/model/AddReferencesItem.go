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

// AddReferencesItem is the corresponding interface of AddReferencesItem
type AddReferencesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSourceNodeId returns SourceNodeId (property field)
	GetSourceNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetServerUri returns TargetServerUri (property field)
	GetTargetServerUri() PascalString
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() ExpandedNodeId
	// GetTargetNodeClass returns TargetNodeClass (property field)
	GetTargetNodeClass() NodeClass
}

// AddReferencesItemExactly can be used when we want exactly this type and not a type which fulfills AddReferencesItem.
// This is useful for switch cases.
type AddReferencesItemExactly interface {
	AddReferencesItem
	isAddReferencesItem() bool
}

// _AddReferencesItem is the data-structure of this message
type _AddReferencesItem struct {
	*_ExtensionObjectDefinition
	SourceNodeId    NodeId
	ReferenceTypeId NodeId
	IsForward       bool
	TargetServerUri PascalString
	TargetNodeId    ExpandedNodeId
	TargetNodeClass NodeClass
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddReferencesItem) GetIdentifier() string {
	return "381"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddReferencesItem) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_AddReferencesItem) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddReferencesItem) GetSourceNodeId() NodeId {
	return m.SourceNodeId
}

func (m *_AddReferencesItem) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_AddReferencesItem) GetIsForward() bool {
	return m.IsForward
}

func (m *_AddReferencesItem) GetTargetServerUri() PascalString {
	return m.TargetServerUri
}

func (m *_AddReferencesItem) GetTargetNodeId() ExpandedNodeId {
	return m.TargetNodeId
}

func (m *_AddReferencesItem) GetTargetNodeClass() NodeClass {
	return m.TargetNodeClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAddReferencesItem factory function for _AddReferencesItem
func NewAddReferencesItem(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetServerUri PascalString, targetNodeId ExpandedNodeId, targetNodeClass NodeClass) *_AddReferencesItem {
	_result := &_AddReferencesItem{
		SourceNodeId:               sourceNodeId,
		ReferenceTypeId:            referenceTypeId,
		IsForward:                  isForward,
		TargetServerUri:            targetServerUri,
		TargetNodeId:               targetNodeId,
		TargetNodeClass:            targetNodeClass,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAddReferencesItem(structType any) AddReferencesItem {
	if casted, ok := structType.(AddReferencesItem); ok {
		return casted
	}
	if casted, ok := structType.(*AddReferencesItem); ok {
		return *casted
	}
	return nil
}

func (m *_AddReferencesItem) GetTypeName() string {
	return "AddReferencesItem"
}

func (m *_AddReferencesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sourceNodeId)
	lengthInBits += m.SourceNodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetServerUri)
	lengthInBits += m.TargetServerUri.GetLengthInBits(ctx)

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Simple field (targetNodeClass)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AddReferencesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AddReferencesItemParse(ctx context.Context, theBytes []byte, identifier string) (AddReferencesItem, error) {
	return AddReferencesItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AddReferencesItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (AddReferencesItem, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AddReferencesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddReferencesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sourceNodeId)
	if pullErr := readBuffer.PullContext("sourceNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceNodeId")
	}
	_sourceNodeId, _sourceNodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _sourceNodeIdErr != nil {
		return nil, errors.Wrap(_sourceNodeIdErr, "Error parsing 'sourceNodeId' field of AddReferencesItem")
	}
	sourceNodeId := _sourceNodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("sourceNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceNodeId")
	}

	// Simple Field (referenceTypeId)
	if pullErr := readBuffer.PullContext("referenceTypeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTypeId")
	}
	_referenceTypeId, _referenceTypeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _referenceTypeIdErr != nil {
		return nil, errors.Wrap(_referenceTypeIdErr, "Error parsing 'referenceTypeId' field of AddReferencesItem")
	}
	referenceTypeId := _referenceTypeId.(NodeId)
	if closeErr := readBuffer.CloseContext("referenceTypeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTypeId")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AddReferencesItem")
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

	// Simple Field (isForward)
	_isForward, _isForwardErr := readBuffer.ReadBit("isForward")
	if _isForwardErr != nil {
		return nil, errors.Wrap(_isForwardErr, "Error parsing 'isForward' field of AddReferencesItem")
	}
	isForward := _isForward

	// Simple Field (targetServerUri)
	if pullErr := readBuffer.PullContext("targetServerUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetServerUri")
	}
	_targetServerUri, _targetServerUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _targetServerUriErr != nil {
		return nil, errors.Wrap(_targetServerUriErr, "Error parsing 'targetServerUri' field of AddReferencesItem")
	}
	targetServerUri := _targetServerUri.(PascalString)
	if closeErr := readBuffer.CloseContext("targetServerUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetServerUri")
	}

	// Simple Field (targetNodeId)
	if pullErr := readBuffer.PullContext("targetNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetNodeId")
	}
	_targetNodeId, _targetNodeIdErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _targetNodeIdErr != nil {
		return nil, errors.Wrap(_targetNodeIdErr, "Error parsing 'targetNodeId' field of AddReferencesItem")
	}
	targetNodeId := _targetNodeId.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("targetNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetNodeId")
	}

	// Simple Field (targetNodeClass)
	if pullErr := readBuffer.PullContext("targetNodeClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetNodeClass")
	}
	_targetNodeClass, _targetNodeClassErr := NodeClassParseWithBuffer(ctx, readBuffer)
	if _targetNodeClassErr != nil {
		return nil, errors.Wrap(_targetNodeClassErr, "Error parsing 'targetNodeClass' field of AddReferencesItem")
	}
	targetNodeClass := _targetNodeClass
	if closeErr := readBuffer.CloseContext("targetNodeClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetNodeClass")
	}

	if closeErr := readBuffer.CloseContext("AddReferencesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddReferencesItem")
	}

	// Create a partially initialized instance
	_child := &_AddReferencesItem{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SourceNodeId:               sourceNodeId,
		ReferenceTypeId:            referenceTypeId,
		IsForward:                  isForward,
		TargetServerUri:            targetServerUri,
		TargetNodeId:               targetNodeId,
		TargetNodeClass:            targetNodeClass,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_AddReferencesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddReferencesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddReferencesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddReferencesItem")
		}

		// Simple Field (sourceNodeId)
		if pushErr := writeBuffer.PushContext("sourceNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sourceNodeId")
		}
		_sourceNodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetSourceNodeId())
		if popErr := writeBuffer.PopContext("sourceNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sourceNodeId")
		}
		if _sourceNodeIdErr != nil {
			return errors.Wrap(_sourceNodeIdErr, "Error serializing 'sourceNodeId' field")
		}

		// Simple Field (referenceTypeId)
		if pushErr := writeBuffer.PushContext("referenceTypeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceTypeId")
		}
		_referenceTypeIdErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceTypeId())
		if popErr := writeBuffer.PopContext("referenceTypeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceTypeId")
		}
		if _referenceTypeIdErr != nil {
			return errors.Wrap(_referenceTypeIdErr, "Error serializing 'referenceTypeId' field")
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

		// Simple Field (isForward)
		isForward := bool(m.GetIsForward())
		_isForwardErr := writeBuffer.WriteBit("isForward", (isForward))
		if _isForwardErr != nil {
			return errors.Wrap(_isForwardErr, "Error serializing 'isForward' field")
		}

		// Simple Field (targetServerUri)
		if pushErr := writeBuffer.PushContext("targetServerUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetServerUri")
		}
		_targetServerUriErr := writeBuffer.WriteSerializable(ctx, m.GetTargetServerUri())
		if popErr := writeBuffer.PopContext("targetServerUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetServerUri")
		}
		if _targetServerUriErr != nil {
			return errors.Wrap(_targetServerUriErr, "Error serializing 'targetServerUri' field")
		}

		// Simple Field (targetNodeId)
		if pushErr := writeBuffer.PushContext("targetNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetNodeId")
		}
		_targetNodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetTargetNodeId())
		if popErr := writeBuffer.PopContext("targetNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetNodeId")
		}
		if _targetNodeIdErr != nil {
			return errors.Wrap(_targetNodeIdErr, "Error serializing 'targetNodeId' field")
		}

		// Simple Field (targetNodeClass)
		if pushErr := writeBuffer.PushContext("targetNodeClass"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetNodeClass")
		}
		_targetNodeClassErr := writeBuffer.WriteSerializable(ctx, m.GetTargetNodeClass())
		if popErr := writeBuffer.PopContext("targetNodeClass"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetNodeClass")
		}
		if _targetNodeClassErr != nil {
			return errors.Wrap(_targetNodeClassErr, "Error serializing 'targetNodeClass' field")
		}

		if popErr := writeBuffer.PopContext("AddReferencesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddReferencesItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddReferencesItem) isAddReferencesItem() bool {
	return true
}

func (m *_AddReferencesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
