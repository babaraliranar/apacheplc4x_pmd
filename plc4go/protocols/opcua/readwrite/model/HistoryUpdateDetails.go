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

// HistoryUpdateDetails is the corresponding interface of HistoryUpdateDetails
type HistoryUpdateDetails interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsHistoryUpdateDetails is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryUpdateDetails()
}

// _HistoryUpdateDetails is the data-structure of this message
type _HistoryUpdateDetails struct {
	ExtensionObjectDefinitionContract
}

var _ HistoryUpdateDetails = (*_HistoryUpdateDetails)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryUpdateDetails)(nil)

// NewHistoryUpdateDetails factory function for _HistoryUpdateDetails
func NewHistoryUpdateDetails() *_HistoryUpdateDetails {
	_result := &_HistoryUpdateDetails{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateDetails) GetIdentifier() string {
	return "679"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateDetails) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastHistoryUpdateDetails(structType any) HistoryUpdateDetails {
	if casted, ok := structType.(HistoryUpdateDetails); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateDetails); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateDetails) GetTypeName() string {
	return "HistoryUpdateDetails"
}

func (m *_HistoryUpdateDetails) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_HistoryUpdateDetails) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryUpdateDetails) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyUpdateDetails HistoryUpdateDetails, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryUpdateDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateDetails")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("HistoryUpdateDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateDetails")
	}

	return m, nil
}

func (m *_HistoryUpdateDetails) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateDetails) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateDetails")
		}

		if popErr := writeBuffer.PopContext("HistoryUpdateDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateDetails")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryUpdateDetails) IsHistoryUpdateDetails() {}

func (m *_HistoryUpdateDetails) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryUpdateDetails) deepCopy() *_HistoryUpdateDetails {
	if m == nil {
		return nil
	}
	_HistoryUpdateDetailsCopy := &_HistoryUpdateDetails{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryUpdateDetailsCopy
}

func (m *_HistoryUpdateDetails) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
