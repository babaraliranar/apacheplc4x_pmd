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

// ConnectedDataItem is the corresponding interface of ConnectedDataItem
type ConnectedDataItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TypeId
	// GetSequenceCount returns SequenceCount (property field)
	GetSequenceCount() uint16
	// GetService returns Service (property field)
	GetService() CipService
}

// ConnectedDataItemExactly can be used when we want exactly this type and not a type which fulfills ConnectedDataItem.
// This is useful for switch cases.
type ConnectedDataItemExactly interface {
	ConnectedDataItem
	isConnectedDataItem() bool
}

// _ConnectedDataItem is the data-structure of this message
type _ConnectedDataItem struct {
	*_TypeId
	SequenceCount uint16
	Service       CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectedDataItem) GetId() uint16 {
	return 0x00B1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectedDataItem) InitializeParent(parent TypeId) {}

func (m *_ConnectedDataItem) GetParent() TypeId {
	return m._TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectedDataItem) GetSequenceCount() uint16 {
	return m.SequenceCount
}

func (m *_ConnectedDataItem) GetService() CipService {
	return m.Service
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectedDataItem factory function for _ConnectedDataItem
func NewConnectedDataItem(sequenceCount uint16, service CipService) *_ConnectedDataItem {
	_result := &_ConnectedDataItem{
		SequenceCount: sequenceCount,
		Service:       service,
		_TypeId:       NewTypeId(),
	}
	_result._TypeId._TypeIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectedDataItem(structType any) ConnectedDataItem {
	if casted, ok := structType.(ConnectedDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectedDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectedDataItem) GetTypeName() string {
	return "ConnectedDataItem"
}

func (m *_ConnectedDataItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (packetSize)
	lengthInBits += 16

	// Simple field (sequenceCount)
	lengthInBits += 16

	// Simple field (service)
	lengthInBits += m.Service.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ConnectedDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectedDataItemParse(ctx context.Context, theBytes []byte) (ConnectedDataItem, error) {
	return ConnectedDataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConnectedDataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectedDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ConnectedDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectedDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (packetSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	packetSize, _packetSizeErr := readBuffer.ReadUint16("packetSize", 16)
	_ = packetSize
	if _packetSizeErr != nil {
		return nil, errors.Wrap(_packetSizeErr, "Error parsing 'packetSize' field of ConnectedDataItem")
	}

	// Simple Field (sequenceCount)
	_sequenceCount, _sequenceCountErr := readBuffer.ReadUint16("sequenceCount", 16)
	if _sequenceCountErr != nil {
		return nil, errors.Wrap(_sequenceCountErr, "Error parsing 'sequenceCount' field of ConnectedDataItem")
	}
	sequenceCount := _sequenceCount

	// Simple Field (service)
	if pullErr := readBuffer.PullContext("service"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for service")
	}
	_service, _serviceErr := CipServiceParseWithBuffer(ctx, readBuffer, bool(bool(true)), uint16(uint16(packetSize)-uint16(uint16(2))))
	if _serviceErr != nil {
		return nil, errors.Wrap(_serviceErr, "Error parsing 'service' field of ConnectedDataItem")
	}
	service := _service.(CipService)
	if closeErr := readBuffer.CloseContext("service"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for service")
	}

	if closeErr := readBuffer.CloseContext("ConnectedDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectedDataItem")
	}

	// Create a partially initialized instance
	_child := &_ConnectedDataItem{
		_TypeId:       &_TypeId{},
		SequenceCount: sequenceCount,
		Service:       service,
	}
	_child._TypeId._TypeIdChildRequirements = _child
	return _child, nil
}

func (m *_ConnectedDataItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectedDataItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectedDataItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectedDataItem")
		}

		// Implicit Field (packetSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		packetSize := uint16(uint16(m.GetService().GetLengthInBytes(ctx)) + uint16(uint16(2)))
		_packetSizeErr := writeBuffer.WriteUint16("packetSize", 16, uint16((packetSize)))
		if _packetSizeErr != nil {
			return errors.Wrap(_packetSizeErr, "Error serializing 'packetSize' field")
		}

		// Simple Field (sequenceCount)
		sequenceCount := uint16(m.GetSequenceCount())
		_sequenceCountErr := writeBuffer.WriteUint16("sequenceCount", 16, uint16((sequenceCount)))
		if _sequenceCountErr != nil {
			return errors.Wrap(_sequenceCountErr, "Error serializing 'sequenceCount' field")
		}

		// Simple Field (service)
		if pushErr := writeBuffer.PushContext("service"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for service")
		}
		_serviceErr := writeBuffer.WriteSerializable(ctx, m.GetService())
		if popErr := writeBuffer.PopContext("service"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for service")
		}
		if _serviceErr != nil {
			return errors.Wrap(_serviceErr, "Error serializing 'service' field")
		}

		if popErr := writeBuffer.PopContext("ConnectedDataItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectedDataItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectedDataItem) isConnectedDataItem() bool {
	return true
}

func (m *_ConnectedDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
