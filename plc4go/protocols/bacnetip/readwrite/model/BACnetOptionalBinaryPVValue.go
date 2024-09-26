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

// BACnetOptionalBinaryPVValue is the corresponding interface of BACnetOptionalBinaryPVValue
type BACnetOptionalBinaryPVValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetOptionalBinaryPV
	// GetBinaryPv returns BinaryPv (property field)
	GetBinaryPv() BACnetBinaryPVTagged
	// IsBACnetOptionalBinaryPVValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalBinaryPVValue()
}

// _BACnetOptionalBinaryPVValue is the data-structure of this message
type _BACnetOptionalBinaryPVValue struct {
	BACnetOptionalBinaryPVContract
	BinaryPv BACnetBinaryPVTagged
}

var _ BACnetOptionalBinaryPVValue = (*_BACnetOptionalBinaryPVValue)(nil)
var _ BACnetOptionalBinaryPVRequirements = (*_BACnetOptionalBinaryPVValue)(nil)

// NewBACnetOptionalBinaryPVValue factory function for _BACnetOptionalBinaryPVValue
func NewBACnetOptionalBinaryPVValue(peekedTagHeader BACnetTagHeader, binaryPv BACnetBinaryPVTagged) *_BACnetOptionalBinaryPVValue {
	if binaryPv == nil {
		panic("binaryPv of type BACnetBinaryPVTagged for BACnetOptionalBinaryPVValue must not be nil")
	}
	_result := &_BACnetOptionalBinaryPVValue{
		BACnetOptionalBinaryPVContract: NewBACnetOptionalBinaryPV(peekedTagHeader),
		BinaryPv:                       binaryPv,
	}
	_result.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV)._SubType = _result
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

func (m *_BACnetOptionalBinaryPVValue) GetParent() BACnetOptionalBinaryPVContract {
	return m.BACnetOptionalBinaryPVContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalBinaryPVValue) GetBinaryPv() BACnetBinaryPVTagged {
	return m.BinaryPv
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOptionalBinaryPVValue(structType any) BACnetOptionalBinaryPVValue {
	if casted, ok := structType.(BACnetOptionalBinaryPVValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPVValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalBinaryPVValue) GetTypeName() string {
	return "BACnetOptionalBinaryPVValue"
}

func (m *_BACnetOptionalBinaryPVValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV).getLengthInBits(ctx))

	// Simple field (binaryPv)
	lengthInBits += m.BinaryPv.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalBinaryPVValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalBinaryPVValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalBinaryPV) (__bACnetOptionalBinaryPVValue BACnetOptionalBinaryPVValue, err error) {
	m.BACnetOptionalBinaryPVContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalBinaryPVValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalBinaryPVValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	binaryPv, err := ReadSimpleField[BACnetBinaryPVTagged](ctx, "binaryPv", ReadComplex[BACnetBinaryPVTagged](BACnetBinaryPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'binaryPv' field"))
	}
	m.BinaryPv = binaryPv

	if closeErr := readBuffer.CloseContext("BACnetOptionalBinaryPVValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalBinaryPVValue")
	}

	return m, nil
}

func (m *_BACnetOptionalBinaryPVValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalBinaryPVValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalBinaryPVValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalBinaryPVValue")
		}

		if err := WriteSimpleField[BACnetBinaryPVTagged](ctx, "binaryPv", m.GetBinaryPv(), WriteComplex[BACnetBinaryPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'binaryPv' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalBinaryPVValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalBinaryPVValue")
		}
		return nil
	}
	return m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalBinaryPVValue) IsBACnetOptionalBinaryPVValue() {}

func (m *_BACnetOptionalBinaryPVValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalBinaryPVValue) deepCopy() *_BACnetOptionalBinaryPVValue {
	if m == nil {
		return nil
	}
	_BACnetOptionalBinaryPVValueCopy := &_BACnetOptionalBinaryPVValue{
		m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV).deepCopy(),
		m.BinaryPv.DeepCopy().(BACnetBinaryPVTagged),
	}
	m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV)._SubType = m
	return _BACnetOptionalBinaryPVValueCopy
}

func (m *_BACnetOptionalBinaryPVValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
