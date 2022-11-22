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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsDiscoveryBlock is the corresponding interface of AdsDiscoveryBlock
type AdsDiscoveryBlock interface {
	utils.LengthAware
	utils.Serializable
	// GetBlockType returns BlockType (discriminator field)
	GetBlockType() AdsDiscoveryBlockType
}

// AdsDiscoveryBlockExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlock.
// This is useful for switch cases.
type AdsDiscoveryBlockExactly interface {
	AdsDiscoveryBlock
	isAdsDiscoveryBlock() bool
}

// _AdsDiscoveryBlock is the data-structure of this message
type _AdsDiscoveryBlock struct {
	_AdsDiscoveryBlockChildRequirements
}

type _AdsDiscoveryBlockChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetBlockType() AdsDiscoveryBlockType
}

type AdsDiscoveryBlockParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child AdsDiscoveryBlock, serializeChildFunction func() error) error
	GetTypeName() string
}

type AdsDiscoveryBlockChild interface {
	utils.Serializable
	InitializeParent(parent AdsDiscoveryBlock)
	GetParent() *AdsDiscoveryBlock

	GetTypeName() string
	AdsDiscoveryBlock
}

// NewAdsDiscoveryBlock factory function for _AdsDiscoveryBlock
func NewAdsDiscoveryBlock() *_AdsDiscoveryBlock {
	return &_AdsDiscoveryBlock{}
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlock(structType interface{}) AdsDiscoveryBlock {
	if casted, ok := structType.(AdsDiscoveryBlock); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlock); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlock) GetTypeName() string {
	return "AdsDiscoveryBlock"
}

func (m *_AdsDiscoveryBlock) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (blockType)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsDiscoveryBlock) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsDiscoveryBlockParse(theBytes []byte) (AdsDiscoveryBlock, error) {
	return AdsDiscoveryBlockParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockParseWithBuffer(readBuffer utils.ReadBuffer) (AdsDiscoveryBlock, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (blockType) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("blockType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for blockType")
	}
	blockType_temp, _blockTypeErr := AdsDiscoveryBlockTypeParseWithBuffer(readBuffer)
	var blockType AdsDiscoveryBlockType = blockType_temp
	if closeErr := readBuffer.CloseContext("blockType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for blockType")
	}
	if _blockTypeErr != nil {
		return nil, errors.Wrap(_blockTypeErr, "Error parsing 'blockType' field of AdsDiscoveryBlock")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AdsDiscoveryBlockChildSerializeRequirement interface {
		AdsDiscoveryBlock
		InitializeParent(AdsDiscoveryBlock)
		GetParent() AdsDiscoveryBlock
	}
	var _childTemp interface{}
	var _child AdsDiscoveryBlockChildSerializeRequirement
	var typeSwitchError error
	switch {
	case blockType == AdsDiscoveryBlockType_STATUS: // AdsDiscoveryBlockStatus
		_childTemp, typeSwitchError = AdsDiscoveryBlockStatusParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_PASSWORD: // AdsDiscoveryBlockPassword
		_childTemp, typeSwitchError = AdsDiscoveryBlockPasswordParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_VERSION: // AdsDiscoveryBlockVersion
		_childTemp, typeSwitchError = AdsDiscoveryBlockVersionParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_OS_DATA: // AdsDiscoveryBlockOsData
		_childTemp, typeSwitchError = AdsDiscoveryBlockOsDataParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_HOST_NAME: // AdsDiscoveryBlockHostName
		_childTemp, typeSwitchError = AdsDiscoveryBlockHostNameParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_AMS_NET_ID: // AdsDiscoveryBlockAmsNetId
		_childTemp, typeSwitchError = AdsDiscoveryBlockAmsNetIdParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_ROUTE_NAME: // AdsDiscoveryBlockRouteName
		_childTemp, typeSwitchError = AdsDiscoveryBlockRouteNameParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_USER_NAME: // AdsDiscoveryBlockUserName
		_childTemp, typeSwitchError = AdsDiscoveryBlockUserNameParseWithBuffer(readBuffer)
	case blockType == AdsDiscoveryBlockType_FINGERPRINT: // AdsDiscoveryBlockFingerprint
		_childTemp, typeSwitchError = AdsDiscoveryBlockFingerprintParseWithBuffer(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [blockType=%v]", blockType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of AdsDiscoveryBlock")
	}
	_child = _childTemp.(AdsDiscoveryBlockChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlock")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_AdsDiscoveryBlock) SerializeParent(writeBuffer utils.WriteBuffer, child AdsDiscoveryBlock, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsDiscoveryBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlock")
	}

	// Discriminator Field (blockType) (Used as input to a switch field)
	blockType := AdsDiscoveryBlockType(child.GetBlockType())
	if pushErr := writeBuffer.PushContext("blockType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for blockType")
	}
	_blockTypeErr := writeBuffer.WriteSerializable(blockType)
	if popErr := writeBuffer.PopContext("blockType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for blockType")
	}

	if _blockTypeErr != nil {
		return errors.Wrap(_blockTypeErr, "Error serializing 'blockType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsDiscoveryBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlock")
	}
	return nil
}

func (m *_AdsDiscoveryBlock) isAdsDiscoveryBlock() bool {
	return true
}

func (m *_AdsDiscoveryBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
