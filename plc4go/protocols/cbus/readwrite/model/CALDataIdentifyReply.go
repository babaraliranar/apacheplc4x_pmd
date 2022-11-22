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

// CALDataIdentifyReply is the corresponding interface of CALDataIdentifyReply
type CALDataIdentifyReply interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetAttribute returns Attribute (property field)
	GetAttribute() Attribute
	// GetIdentifyReplyCommand returns IdentifyReplyCommand (property field)
	GetIdentifyReplyCommand() IdentifyReplyCommand
}

// CALDataIdentifyReplyExactly can be used when we want exactly this type and not a type which fulfills CALDataIdentifyReply.
// This is useful for switch cases.
type CALDataIdentifyReplyExactly interface {
	CALDataIdentifyReply
	isCALDataIdentifyReply() bool
}

// _CALDataIdentifyReply is the data-structure of this message
type _CALDataIdentifyReply struct {
	*_CALData
	Attribute            Attribute
	IdentifyReplyCommand IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataIdentifyReply) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataIdentifyReply) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataIdentifyReply) GetAttribute() Attribute {
	return m.Attribute
}

func (m *_CALDataIdentifyReply) GetIdentifyReplyCommand() IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataIdentifyReply factory function for _CALDataIdentifyReply
func NewCALDataIdentifyReply(attribute Attribute, identifyReplyCommand IdentifyReplyCommand, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataIdentifyReply {
	_result := &_CALDataIdentifyReply{
		Attribute:            attribute,
		IdentifyReplyCommand: identifyReplyCommand,
		_CALData:             NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataIdentifyReply(structType interface{}) CALDataIdentifyReply {
	if casted, ok := structType.(CALDataIdentifyReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataIdentifyReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataIdentifyReply) GetTypeName() string {
	return "CALDataIdentifyReply"
}

func (m *_CALDataIdentifyReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataIdentifyReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (attribute)
	lengthInBits += 8

	// Simple field (identifyReplyCommand)
	lengthInBits += m.IdentifyReplyCommand.GetLengthInBits()

	return lengthInBits
}

func (m *_CALDataIdentifyReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataIdentifyReplyParse(theBytes []byte, requestContext RequestContext, commandTypeContainer CALCommandTypeContainer) (CALDataIdentifyReply, error) {
	return CALDataIdentifyReplyParseWithBuffer(utils.NewReadBufferByteBased(theBytes), requestContext, commandTypeContainer)
}

func CALDataIdentifyReplyParseWithBuffer(readBuffer utils.ReadBuffer, requestContext RequestContext, commandTypeContainer CALCommandTypeContainer) (CALDataIdentifyReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataIdentifyReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataIdentifyReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (attribute)
	if pullErr := readBuffer.PullContext("attribute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for attribute")
	}
	_attribute, _attributeErr := AttributeParseWithBuffer(readBuffer)
	if _attributeErr != nil {
		return nil, errors.Wrap(_attributeErr, "Error parsing 'attribute' field of CALDataIdentifyReply")
	}
	attribute := _attribute
	if closeErr := readBuffer.CloseContext("attribute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for attribute")
	}

	// Simple Field (identifyReplyCommand)
	if pullErr := readBuffer.PullContext("identifyReplyCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for identifyReplyCommand")
	}
	_identifyReplyCommand, _identifyReplyCommandErr := IdentifyReplyCommandParseWithBuffer(readBuffer, Attribute(attribute), uint8(uint8(commandTypeContainer.NumBytes())-uint8(uint8(1))))
	if _identifyReplyCommandErr != nil {
		return nil, errors.Wrap(_identifyReplyCommandErr, "Error parsing 'identifyReplyCommand' field of CALDataIdentifyReply")
	}
	identifyReplyCommand := _identifyReplyCommand.(IdentifyReplyCommand)
	if closeErr := readBuffer.CloseContext("identifyReplyCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for identifyReplyCommand")
	}

	if closeErr := readBuffer.CloseContext("CALDataIdentifyReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataIdentifyReply")
	}

	// Create a partially initialized instance
	_child := &_CALDataIdentifyReply{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		Attribute:            attribute,
		IdentifyReplyCommand: identifyReplyCommand,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataIdentifyReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataIdentifyReply) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataIdentifyReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataIdentifyReply")
		}

		// Simple Field (attribute)
		if pushErr := writeBuffer.PushContext("attribute"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for attribute")
		}
		_attributeErr := writeBuffer.WriteSerializable(m.GetAttribute())
		if popErr := writeBuffer.PopContext("attribute"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for attribute")
		}
		if _attributeErr != nil {
			return errors.Wrap(_attributeErr, "Error serializing 'attribute' field")
		}

		// Simple Field (identifyReplyCommand)
		if pushErr := writeBuffer.PushContext("identifyReplyCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for identifyReplyCommand")
		}
		_identifyReplyCommandErr := writeBuffer.WriteSerializable(m.GetIdentifyReplyCommand())
		if popErr := writeBuffer.PopContext("identifyReplyCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for identifyReplyCommand")
		}
		if _identifyReplyCommandErr != nil {
			return errors.Wrap(_identifyReplyCommandErr, "Error serializing 'identifyReplyCommand' field")
		}

		if popErr := writeBuffer.PopContext("CALDataIdentifyReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataIdentifyReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataIdentifyReply) isCALDataIdentifyReply() bool {
	return true
}

func (m *_CALDataIdentifyReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
