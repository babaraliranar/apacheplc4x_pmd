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

// BACnetChannelValueUnsigned is the corresponding interface of BACnetChannelValueUnsigned
type BACnetChannelValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetChannelValueUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueUnsigned()
}

// _BACnetChannelValueUnsigned is the data-structure of this message
type _BACnetChannelValueUnsigned struct {
	BACnetChannelValueContract
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetChannelValueUnsigned = (*_BACnetChannelValueUnsigned)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueUnsigned)(nil)

// NewBACnetChannelValueUnsigned factory function for _BACnetChannelValueUnsigned
func NewBACnetChannelValueUnsigned(peekedTagHeader BACnetTagHeader, unsignedValue BACnetApplicationTagUnsignedInteger) *_BACnetChannelValueUnsigned {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetApplicationTagUnsignedInteger for BACnetChannelValueUnsigned must not be nil")
	}
	_result := &_BACnetChannelValueUnsigned{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		UnsignedValue:              unsignedValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueUnsigned) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueUnsigned(structType any) BACnetChannelValueUnsigned {
	if casted, ok := structType.(BACnetChannelValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueUnsigned) GetTypeName() string {
	return "BACnetChannelValueUnsigned"
}

func (m *_BACnetChannelValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueUnsigned BACnetChannelValueUnsigned, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueUnsigned")
	}

	return m, nil
}

func (m *_BACnetChannelValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueUnsigned")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueUnsigned")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueUnsigned) IsBACnetChannelValueUnsigned() {}

func (m *_BACnetChannelValueUnsigned) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueUnsigned) deepCopy() *_BACnetChannelValueUnsigned {
	if m == nil {
		return nil
	}
	_BACnetChannelValueUnsignedCopy := &_BACnetChannelValueUnsigned{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		m.UnsignedValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueUnsignedCopy
}

func (m *_BACnetChannelValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
