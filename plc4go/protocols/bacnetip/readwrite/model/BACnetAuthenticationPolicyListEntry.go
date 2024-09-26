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

// BACnetAuthenticationPolicyListEntry is the corresponding interface of BACnetAuthenticationPolicyListEntry
type BACnetAuthenticationPolicyListEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetCredentialDataInput returns CredentialDataInput (property field)
	GetCredentialDataInput() BACnetDeviceObjectReferenceEnclosed
	// GetIndex returns Index (property field)
	GetIndex() BACnetContextTagUnsignedInteger
	// IsBACnetAuthenticationPolicyListEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAuthenticationPolicyListEntry()
}

// _BACnetAuthenticationPolicyListEntry is the data-structure of this message
type _BACnetAuthenticationPolicyListEntry struct {
	CredentialDataInput BACnetDeviceObjectReferenceEnclosed
	Index               BACnetContextTagUnsignedInteger
}

var _ BACnetAuthenticationPolicyListEntry = (*_BACnetAuthenticationPolicyListEntry)(nil)

// NewBACnetAuthenticationPolicyListEntry factory function for _BACnetAuthenticationPolicyListEntry
func NewBACnetAuthenticationPolicyListEntry(credentialDataInput BACnetDeviceObjectReferenceEnclosed, index BACnetContextTagUnsignedInteger) *_BACnetAuthenticationPolicyListEntry {
	if credentialDataInput == nil {
		panic("credentialDataInput of type BACnetDeviceObjectReferenceEnclosed for BACnetAuthenticationPolicyListEntry must not be nil")
	}
	if index == nil {
		panic("index of type BACnetContextTagUnsignedInteger for BACnetAuthenticationPolicyListEntry must not be nil")
	}
	return &_BACnetAuthenticationPolicyListEntry{CredentialDataInput: credentialDataInput, Index: index}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationPolicyListEntry) GetCredentialDataInput() BACnetDeviceObjectReferenceEnclosed {
	return m.CredentialDataInput
}

func (m *_BACnetAuthenticationPolicyListEntry) GetIndex() BACnetContextTagUnsignedInteger {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationPolicyListEntry(structType any) BACnetAuthenticationPolicyListEntry {
	if casted, ok := structType.(BACnetAuthenticationPolicyListEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicyListEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationPolicyListEntry) GetTypeName() string {
	return "BACnetAuthenticationPolicyListEntry"
}

func (m *_BACnetAuthenticationPolicyListEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (credentialDataInput)
	lengthInBits += m.CredentialDataInput.GetLengthInBits(ctx)

	// Simple field (index)
	lengthInBits += m.Index.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationPolicyListEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationPolicyListEntryParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationPolicyListEntry, error) {
	return BACnetAuthenticationPolicyListEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationPolicyListEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyListEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyListEntry, error) {
		return BACnetAuthenticationPolicyListEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAuthenticationPolicyListEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyListEntry, error) {
	v, err := (&_BACnetAuthenticationPolicyListEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAuthenticationPolicyListEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAuthenticationPolicyListEntry BACnetAuthenticationPolicyListEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicyListEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicyListEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	credentialDataInput, err := ReadSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "credentialDataInput", ReadComplex[BACnetDeviceObjectReferenceEnclosed](BACnetDeviceObjectReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'credentialDataInput' field"))
	}
	m.CredentialDataInput = credentialDataInput

	index, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "index", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}
	m.Index = index

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicyListEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicyListEntry")
	}

	return m, nil
}

func (m *_BACnetAuthenticationPolicyListEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationPolicyListEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicyListEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicyListEntry")
	}

	if err := WriteSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "credentialDataInput", m.GetCredentialDataInput(), WriteComplex[BACnetDeviceObjectReferenceEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'credentialDataInput' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "index", m.GetIndex(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'index' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicyListEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationPolicyListEntry")
	}
	return nil
}

func (m *_BACnetAuthenticationPolicyListEntry) IsBACnetAuthenticationPolicyListEntry() {}

func (m *_BACnetAuthenticationPolicyListEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAuthenticationPolicyListEntry) deepCopy() *_BACnetAuthenticationPolicyListEntry {
	if m == nil {
		return nil
	}
	_BACnetAuthenticationPolicyListEntryCopy := &_BACnetAuthenticationPolicyListEntry{
		m.CredentialDataInput.DeepCopy().(BACnetDeviceObjectReferenceEnclosed),
		m.Index.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	return _BACnetAuthenticationPolicyListEntryCopy
}

func (m *_BACnetAuthenticationPolicyListEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
