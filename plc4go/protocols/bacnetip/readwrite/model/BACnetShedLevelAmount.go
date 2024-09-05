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

// BACnetShedLevelAmount is the corresponding interface of BACnetShedLevelAmount
type BACnetShedLevelAmount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetShedLevel
	// GetAmount returns Amount (property field)
	GetAmount() BACnetContextTagReal
	// IsBACnetShedLevelAmount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevelAmount()
}

// _BACnetShedLevelAmount is the data-structure of this message
type _BACnetShedLevelAmount struct {
	BACnetShedLevelContract
	Amount BACnetContextTagReal
}

var _ BACnetShedLevelAmount = (*_BACnetShedLevelAmount)(nil)
var _ BACnetShedLevelRequirements = (*_BACnetShedLevelAmount)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetShedLevelAmount) GetParent() BACnetShedLevelContract {
	return m.BACnetShedLevelContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelAmount) GetAmount() BACnetContextTagReal {
	return m.Amount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetShedLevelAmount factory function for _BACnetShedLevelAmount
func NewBACnetShedLevelAmount(amount BACnetContextTagReal, peekedTagHeader BACnetTagHeader) *_BACnetShedLevelAmount {
	if amount == nil {
		panic("amount of type BACnetContextTagReal for BACnetShedLevelAmount must not be nil")
	}
	_result := &_BACnetShedLevelAmount{
		BACnetShedLevelContract: NewBACnetShedLevel(peekedTagHeader),
		Amount:                  amount,
	}
	_result.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelAmount(structType any) BACnetShedLevelAmount {
	if casted, ok := structType.(BACnetShedLevelAmount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelAmount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelAmount) GetTypeName() string {
	return "BACnetShedLevelAmount"
}

func (m *_BACnetShedLevelAmount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetShedLevelContract.(*_BACnetShedLevel).getLengthInBits(ctx))

	// Simple field (amount)
	lengthInBits += m.Amount.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetShedLevelAmount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetShedLevelAmount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetShedLevel) (__bACnetShedLevelAmount BACnetShedLevelAmount, err error) {
	m.BACnetShedLevelContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelAmount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelAmount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	amount, err := ReadSimpleField[BACnetContextTagReal](ctx, "amount", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'amount' field"))
	}
	m.Amount = amount

	if closeErr := readBuffer.CloseContext("BACnetShedLevelAmount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelAmount")
	}

	return m, nil
}

func (m *_BACnetShedLevelAmount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelAmount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelAmount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelAmount")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "amount", m.GetAmount(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'amount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelAmount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelAmount")
		}
		return nil
	}
	return m.BACnetShedLevelContract.(*_BACnetShedLevel).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetShedLevelAmount) IsBACnetShedLevelAmount() {}

func (m *_BACnetShedLevelAmount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
