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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesReasonForHalt is the corresponding interface of BACnetPropertyStatesReasonForHalt
type BACnetPropertyStatesReasonForHalt interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetReasonForHalt returns ReasonForHalt (property field)
	GetReasonForHalt() BACnetProgramErrorTagged
}

// BACnetPropertyStatesReasonForHaltExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesReasonForHalt.
// This is useful for switch cases.
type BACnetPropertyStatesReasonForHaltExactly interface {
	BACnetPropertyStatesReasonForHalt
	isBACnetPropertyStatesReasonForHalt() bool
}

// _BACnetPropertyStatesReasonForHalt is the data-structure of this message
type _BACnetPropertyStatesReasonForHalt struct {
	*_BACnetPropertyStates
	ReasonForHalt BACnetProgramErrorTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesReasonForHalt) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesReasonForHalt) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesReasonForHalt) GetReasonForHalt() BACnetProgramErrorTagged {
	return m.ReasonForHalt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesReasonForHalt factory function for _BACnetPropertyStatesReasonForHalt
func NewBACnetPropertyStatesReasonForHalt(reasonForHalt BACnetProgramErrorTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesReasonForHalt {
	_result := &_BACnetPropertyStatesReasonForHalt{
		ReasonForHalt:         reasonForHalt,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesReasonForHalt(structType interface{}) BACnetPropertyStatesReasonForHalt {
	if casted, ok := structType.(BACnetPropertyStatesReasonForHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesReasonForHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesReasonForHalt) GetTypeName() string {
	return "BACnetPropertyStatesReasonForHalt"
}

func (m *_BACnetPropertyStatesReasonForHalt) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesReasonForHalt) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reasonForHalt)
	lengthInBits += m.ReasonForHalt.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesReasonForHalt) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesReasonForHaltParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesReasonForHalt, error) {
	return BACnetPropertyStatesReasonForHaltParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), peekedTagNumber) // TODO: get endianness from mspec
}

func BACnetPropertyStatesReasonForHaltParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesReasonForHalt, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesReasonForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesReasonForHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reasonForHalt)
	if pullErr := readBuffer.PullContext("reasonForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reasonForHalt")
	}
	_reasonForHalt, _reasonForHaltErr := BACnetProgramErrorTaggedParseWithBuffer(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _reasonForHaltErr != nil {
		return nil, errors.Wrap(_reasonForHaltErr, "Error parsing 'reasonForHalt' field of BACnetPropertyStatesReasonForHalt")
	}
	reasonForHalt := _reasonForHalt.(BACnetProgramErrorTagged)
	if closeErr := readBuffer.CloseContext("reasonForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reasonForHalt")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesReasonForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesReasonForHalt")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesReasonForHalt{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		ReasonForHalt:         reasonForHalt,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesReasonForHalt) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesReasonForHalt) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesReasonForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesReasonForHalt")
		}

		// Simple Field (reasonForHalt)
		if pushErr := writeBuffer.PushContext("reasonForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reasonForHalt")
		}
		_reasonForHaltErr := writeBuffer.WriteSerializable(m.GetReasonForHalt())
		if popErr := writeBuffer.PopContext("reasonForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reasonForHalt")
		}
		if _reasonForHaltErr != nil {
			return errors.Wrap(_reasonForHaltErr, "Error serializing 'reasonForHalt' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesReasonForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesReasonForHalt")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesReasonForHalt) isBACnetPropertyStatesReasonForHalt() bool {
	return true
}

func (m *_BACnetPropertyStatesReasonForHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
