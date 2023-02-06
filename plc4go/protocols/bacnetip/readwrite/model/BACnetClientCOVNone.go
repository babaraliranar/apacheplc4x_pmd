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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetClientCOVNone is the corresponding interface of BACnetClientCOVNone
type BACnetClientCOVNone interface {
	utils.LengthAware
	utils.Serializable
	BACnetClientCOV
	// GetDefaultIncrement returns DefaultIncrement (property field)
	GetDefaultIncrement() BACnetApplicationTagNull
}

// BACnetClientCOVNoneExactly can be used when we want exactly this type and not a type which fulfills BACnetClientCOVNone.
// This is useful for switch cases.
type BACnetClientCOVNoneExactly interface {
	BACnetClientCOVNone
	isBACnetClientCOVNone() bool
}

// _BACnetClientCOVNone is the data-structure of this message
type _BACnetClientCOVNone struct {
	*_BACnetClientCOV
	DefaultIncrement BACnetApplicationTagNull
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetClientCOVNone) InitializeParent(parent BACnetClientCOV, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetClientCOVNone) GetParent() BACnetClientCOV {
	return m._BACnetClientCOV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetClientCOVNone) GetDefaultIncrement() BACnetApplicationTagNull {
	return m.DefaultIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetClientCOVNone factory function for _BACnetClientCOVNone
func NewBACnetClientCOVNone(defaultIncrement BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader) *_BACnetClientCOVNone {
	_result := &_BACnetClientCOVNone{
		DefaultIncrement: defaultIncrement,
		_BACnetClientCOV: NewBACnetClientCOV(peekedTagHeader),
	}
	_result._BACnetClientCOV._BACnetClientCOVChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetClientCOVNone(structType interface{}) BACnetClientCOVNone {
	if casted, ok := structType.(BACnetClientCOVNone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetClientCOVNone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetClientCOVNone) GetTypeName() string {
	return "BACnetClientCOVNone"
}

func (m *_BACnetClientCOVNone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (defaultIncrement)
	lengthInBits += m.DefaultIncrement.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetClientCOVNone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetClientCOVNoneParse(theBytes []byte) (BACnetClientCOVNone, error) {
	return BACnetClientCOVNoneParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetClientCOVNoneParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetClientCOVNone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClientCOVNone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetClientCOVNone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (defaultIncrement)
	if pullErr := readBuffer.PullContext("defaultIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for defaultIncrement")
	}
	_defaultIncrement, _defaultIncrementErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _defaultIncrementErr != nil {
		return nil, errors.Wrap(_defaultIncrementErr, "Error parsing 'defaultIncrement' field of BACnetClientCOVNone")
	}
	defaultIncrement := _defaultIncrement.(BACnetApplicationTagNull)
	if closeErr := readBuffer.CloseContext("defaultIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for defaultIncrement")
	}

	if closeErr := readBuffer.CloseContext("BACnetClientCOVNone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetClientCOVNone")
	}

	// Create a partially initialized instance
	_child := &_BACnetClientCOVNone{
		_BACnetClientCOV: &_BACnetClientCOV{},
		DefaultIncrement: defaultIncrement,
	}
	_child._BACnetClientCOV._BACnetClientCOVChildRequirements = _child
	return _child, nil
}

func (m *_BACnetClientCOVNone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetClientCOVNone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetClientCOVNone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetClientCOVNone")
		}

		// Simple Field (defaultIncrement)
		if pushErr := writeBuffer.PushContext("defaultIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for defaultIncrement")
		}
		_defaultIncrementErr := writeBuffer.WriteSerializable(ctx, m.GetDefaultIncrement())
		if popErr := writeBuffer.PopContext("defaultIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for defaultIncrement")
		}
		if _defaultIncrementErr != nil {
			return errors.Wrap(_defaultIncrementErr, "Error serializing 'defaultIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetClientCOVNone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetClientCOVNone")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetClientCOVNone) isBACnetClientCOVNone() bool {
	return true
}

func (m *_BACnetClientCOVNone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
