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

// MediaTransportControlDataRewind is the corresponding interface of MediaTransportControlDataRewind
type MediaTransportControlDataRewind interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsCeaseRewind returns IsCeaseRewind (virtual field)
	GetIsCeaseRewind() bool
	// GetIs2x returns Is2x (virtual field)
	GetIs2x() bool
	// GetIs4x returns Is4x (virtual field)
	GetIs4x() bool
	// GetIs8x returns Is8x (virtual field)
	GetIs8x() bool
	// GetIs16x returns Is16x (virtual field)
	GetIs16x() bool
	// GetIs32x returns Is32x (virtual field)
	GetIs32x() bool
	// GetIs64x returns Is64x (virtual field)
	GetIs64x() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
}

// MediaTransportControlDataRewindExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataRewind.
// This is useful for switch cases.
type MediaTransportControlDataRewindExactly interface {
	MediaTransportControlDataRewind
	isMediaTransportControlDataRewind() bool
}

// _MediaTransportControlDataRewind is the data-structure of this message
type _MediaTransportControlDataRewind struct {
	*_MediaTransportControlData
	Operation byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataRewind) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataRewind) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataRewind) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataRewind) GetIsCeaseRewind() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataRewind) GetIs2x() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x02)))
}

func (m *_MediaTransportControlDataRewind) GetIs4x() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x04)))
}

func (m *_MediaTransportControlDataRewind) GetIs8x() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x06)))
}

func (m *_MediaTransportControlDataRewind) GetIs16x() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x08)))
}

func (m *_MediaTransportControlDataRewind) GetIs32x() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x0A)))
}

func (m *_MediaTransportControlDataRewind) GetIs64x() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x0C)))
}

func (m *_MediaTransportControlDataRewind) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool(bool(bool(bool(bool(!(m.GetIsCeaseRewind())) && bool(!(m.GetIs2x()))) && bool(!(m.GetIs4x()))) && bool(!(m.GetIs8x()))) && bool(!(m.GetIs16x()))) && bool(!(m.GetIs32x()))) && bool(!(m.GetIs64x())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataRewind factory function for _MediaTransportControlDataRewind
func NewMediaTransportControlDataRewind(operation byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataRewind {
	_result := &_MediaTransportControlDataRewind{
		Operation:                  operation,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataRewind(structType any) MediaTransportControlDataRewind {
	if casted, ok := structType.(MediaTransportControlDataRewind); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataRewind); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataRewind) GetTypeName() string {
	return "MediaTransportControlDataRewind"
}

func (m *_MediaTransportControlDataRewind) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataRewind) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataRewindParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataRewind, error) {
	return MediaTransportControlDataRewindParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataRewindParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataRewind, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataRewind, error) {
		return MediaTransportControlDataRewindParseWithBuffer(ctx, readBuffer)
	}
}

func MediaTransportControlDataRewindParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataRewind, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MediaTransportControlDataRewind"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataRewind")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operation, err := ReadSimpleField(ctx, "operation", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operation' field"))
	}

	isCeaseRewind, err := ReadVirtualField[bool](ctx, "isCeaseRewind", (*bool)(nil), bool((operation) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isCeaseRewind' field"))
	}
	_ = isCeaseRewind

	is2x, err := ReadVirtualField[bool](ctx, "is2x", (*bool)(nil), bool((operation) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'is2x' field"))
	}
	_ = is2x

	is4x, err := ReadVirtualField[bool](ctx, "is4x", (*bool)(nil), bool((operation) == (0x04)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'is4x' field"))
	}
	_ = is4x

	is8x, err := ReadVirtualField[bool](ctx, "is8x", (*bool)(nil), bool((operation) == (0x06)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'is8x' field"))
	}
	_ = is8x

	is16x, err := ReadVirtualField[bool](ctx, "is16x", (*bool)(nil), bool((operation) == (0x08)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'is16x' field"))
	}
	_ = is16x

	is32x, err := ReadVirtualField[bool](ctx, "is32x", (*bool)(nil), bool((operation) == (0x0A)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'is32x' field"))
	}
	_ = is32x

	is64x, err := ReadVirtualField[bool](ctx, "is64x", (*bool)(nil), bool((operation) == (0x0C)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'is64x' field"))
	}
	_ = is64x

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool(bool(bool(bool(bool(!(isCeaseRewind)) && bool(!(is2x))) && bool(!(is4x))) && bool(!(is8x))) && bool(!(is16x))) && bool(!(is32x))) && bool(!(is64x)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataRewind"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataRewind")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataRewind{
		_MediaTransportControlData: &_MediaTransportControlData{},
		Operation:                  operation,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataRewind) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataRewind) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataRewind"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataRewind")
		}

		// Simple Field (operation)
		operation := byte(m.GetOperation())
		_operationErr := /*TODO: migrate me*/ writeBuffer.WriteByte("operation", (operation))
		if _operationErr != nil {
			return errors.Wrap(_operationErr, "Error serializing 'operation' field")
		}
		// Virtual field
		isCeaseRewind := m.GetIsCeaseRewind()
		_ = isCeaseRewind
		if _isCeaseRewindErr := writeBuffer.WriteVirtual(ctx, "isCeaseRewind", m.GetIsCeaseRewind()); _isCeaseRewindErr != nil {
			return errors.Wrap(_isCeaseRewindErr, "Error serializing 'isCeaseRewind' field")
		}
		// Virtual field
		is2x := m.GetIs2x()
		_ = is2x
		if _is2xErr := writeBuffer.WriteVirtual(ctx, "is2x", m.GetIs2x()); _is2xErr != nil {
			return errors.Wrap(_is2xErr, "Error serializing 'is2x' field")
		}
		// Virtual field
		is4x := m.GetIs4x()
		_ = is4x
		if _is4xErr := writeBuffer.WriteVirtual(ctx, "is4x", m.GetIs4x()); _is4xErr != nil {
			return errors.Wrap(_is4xErr, "Error serializing 'is4x' field")
		}
		// Virtual field
		is8x := m.GetIs8x()
		_ = is8x
		if _is8xErr := writeBuffer.WriteVirtual(ctx, "is8x", m.GetIs8x()); _is8xErr != nil {
			return errors.Wrap(_is8xErr, "Error serializing 'is8x' field")
		}
		// Virtual field
		is16x := m.GetIs16x()
		_ = is16x
		if _is16xErr := writeBuffer.WriteVirtual(ctx, "is16x", m.GetIs16x()); _is16xErr != nil {
			return errors.Wrap(_is16xErr, "Error serializing 'is16x' field")
		}
		// Virtual field
		is32x := m.GetIs32x()
		_ = is32x
		if _is32xErr := writeBuffer.WriteVirtual(ctx, "is32x", m.GetIs32x()); _is32xErr != nil {
			return errors.Wrap(_is32xErr, "Error serializing 'is32x' field")
		}
		// Virtual field
		is64x := m.GetIs64x()
		_ = is64x
		if _is64xErr := writeBuffer.WriteVirtual(ctx, "is64x", m.GetIs64x()); _is64xErr != nil {
			return errors.Wrap(_is64xErr, "Error serializing 'is64x' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataRewind"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataRewind")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataRewind) isMediaTransportControlDataRewind() bool {
	return true
}

func (m *_MediaTransportControlDataRewind) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
