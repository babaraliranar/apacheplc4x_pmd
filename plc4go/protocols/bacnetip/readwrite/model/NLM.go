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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// NLM is the corresponding interface of NLM
type NLM interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetIsVendorProprietaryMessage returns IsVendorProprietaryMessage (virtual field)
	GetIsVendorProprietaryMessage() bool
}

// NLMExactly can be used when we want exactly this type and not a type which fulfills NLM.
// This is useful for switch cases.
type NLMExactly interface {
	NLM
	isNLM() bool
}

// _NLM is the data-structure of this message
type _NLM struct {
	_NLMChildRequirements

	// Arguments.
	ApduLength uint16
}

type _NLMChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() uint8
	GetIsVendorProprietaryMessage() bool
}

type NLMParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NLM, serializeChildFunction func() error) error
	GetTypeName() string
}

type NLMChild interface {
	utils.Serializable
	InitializeParent(parent NLM)
	GetParent() *NLM

	GetTypeName() string
	NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NLM) GetIsVendorProprietaryMessage() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMessageType()) >= (128)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLM factory function for _NLM
func NewNLM(apduLength uint16) *_NLM {
	return &_NLM{ApduLength: apduLength}
}

// Deprecated: use the interface for direct cast
func CastNLM(structType any) NLM {
	if casted, ok := structType.(NLM); ok {
		return casted
	}
	if casted, ok := structType.(*NLM); ok {
		return *casted
	}
	return nil
}

func (m *_NLM) GetTypeName() string {
	return "NLM"
}

func (m *_NLM) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NLM) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLM, error) {
	return NLMParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLM, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLM"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLM")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field of NLM")
	}

	// Virtual field
	_isVendorProprietaryMessage := bool((messageType) >= (128))
	isVendorProprietaryMessage := bool(_isVendorProprietaryMessage)
	_ = isVendorProprietaryMessage

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type NLMChildSerializeRequirement interface {
		NLM
		InitializeParent(NLM)
		GetParent() NLM
	}
	var _childTemp any
	var _child NLMChildSerializeRequirement
	var typeSwitchError error
	switch {
	case messageType == 0x00: // NLMWhoIsRouterToNetwork
		_childTemp, typeSwitchError = NLMWhoIsRouterToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x01: // NLMIAmRouterToNetwork
		_childTemp, typeSwitchError = NLMIAmRouterToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x02: // NLMICouldBeRouterToNetwork
		_childTemp, typeSwitchError = NLMICouldBeRouterToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x03: // NLMRejectMessageToNetwork
		_childTemp, typeSwitchError = NLMRejectMessageToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x04: // NLMRouterBusyToNetwork
		_childTemp, typeSwitchError = NLMRouterBusyToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x05: // NLMRouterAvailableToNetwork
		_childTemp, typeSwitchError = NLMRouterAvailableToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x06: // NLMInitializeRoutingTable
		_childTemp, typeSwitchError = NLMInitializeRoutingTableParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x07: // NLMInitializeRoutingTableAck
		_childTemp, typeSwitchError = NLMInitializeRoutingTableAckParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x08: // NLMEstablishConnectionToNetwork
		_childTemp, typeSwitchError = NLMEstablishConnectionToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x09: // NLMDisconnectConnectionToNetwork
		_childTemp, typeSwitchError = NLMDisconnectConnectionToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x0A: // NLMChallengeRequest
		_childTemp, typeSwitchError = NLMChallengeRequestParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x0B: // NLMSecurityPayload
		_childTemp, typeSwitchError = NLMSecurityPayloadParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x0C: // NLMSecurityResponse
		_childTemp, typeSwitchError = NLMSecurityResponseParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x0D: // NLMRequestKeyUpdate
		_childTemp, typeSwitchError = NLMRequestKeyUpdateParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x0E: // NLMUpdateKeyUpdate
		_childTemp, typeSwitchError = NLMUpdateKeyUpdateParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x0F: // NLMUpdateKeyDistributionKey
		_childTemp, typeSwitchError = NLMUpdateKeyDistributionKeyParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x10: // NLMRequestMasterKey
		_childTemp, typeSwitchError = NLMRequestMasterKeyParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x11: // NLMSetMasterKey
		_childTemp, typeSwitchError = NLMSetMasterKeyParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x12: // NLMWhatIsNetworkNumber
		_childTemp, typeSwitchError = NLMWhatIsNetworkNumberParseWithBuffer(ctx, readBuffer, apduLength)
	case messageType == 0x13: // NLMNetworkNumberIs
		_childTemp, typeSwitchError = NLMNetworkNumberIsParseWithBuffer(ctx, readBuffer, apduLength)
	case 0 == 0 && isVendorProprietaryMessage == bool(false): // NLMReserved
		_childTemp, typeSwitchError = NLMReservedParseWithBuffer(ctx, readBuffer, apduLength)
	case 0 == 0: // NLMVendorProprietaryMessage
		_childTemp, typeSwitchError = NLMVendorProprietaryMessageParseWithBuffer(ctx, readBuffer, apduLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageType=%v, isVendorProprietaryMessage=%v]", messageType, isVendorProprietaryMessage)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of NLM")
	}
	_child = _childTemp.(NLMChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("NLM"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLM")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_NLM) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NLM, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLM"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLM")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 8, uint8((messageType)))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}
	// Virtual field
	isVendorProprietaryMessage := m.GetIsVendorProprietaryMessage()
	_ = isVendorProprietaryMessage
	if _isVendorProprietaryMessageErr := writeBuffer.WriteVirtual(ctx, "isVendorProprietaryMessage", m.GetIsVendorProprietaryMessage()); _isVendorProprietaryMessageErr != nil {
		return errors.Wrap(_isVendorProprietaryMessageErr, "Error serializing 'isVendorProprietaryMessage' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NLM"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLM")
	}
	return nil
}

////
// Arguments Getter

func (m *_NLM) GetApduLength() uint16 {
	return m.ApduLength
}

//
////

func (m *_NLM) isNLM() bool {
	return true
}

func (m *_NLM) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
