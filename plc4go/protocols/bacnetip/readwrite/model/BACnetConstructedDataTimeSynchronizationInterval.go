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

// BACnetConstructedDataTimeSynchronizationInterval is the corresponding interface of BACnetConstructedDataTimeSynchronizationInterval
type BACnetConstructedDataTimeSynchronizationInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTimeSynchronization returns TimeSynchronization (property field)
	GetTimeSynchronization() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTimeSynchronizationInterval is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeSynchronizationInterval()
}

// _BACnetConstructedDataTimeSynchronizationInterval is the data-structure of this message
type _BACnetConstructedDataTimeSynchronizationInterval struct {
	BACnetConstructedDataContract
	TimeSynchronization BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTimeSynchronizationInterval = (*_BACnetConstructedDataTimeSynchronizationInterval)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeSynchronizationInterval)(nil)

// NewBACnetConstructedDataTimeSynchronizationInterval factory function for _BACnetConstructedDataTimeSynchronizationInterval
func NewBACnetConstructedDataTimeSynchronizationInterval(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeSynchronization BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeSynchronizationInterval {
	if timeSynchronization == nil {
		panic("timeSynchronization of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTimeSynchronizationInterval must not be nil")
	}
	_result := &_BACnetConstructedDataTimeSynchronizationInterval{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeSynchronization:           timeSynchronization,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_SYNCHRONIZATION_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetTimeSynchronization() BACnetApplicationTagUnsignedInteger {
	return m.TimeSynchronization
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTimeSynchronization())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeSynchronizationInterval(structType any) BACnetConstructedDataTimeSynchronizationInterval {
	if casted, ok := structType.(BACnetConstructedDataTimeSynchronizationInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeSynchronizationInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetTypeName() string {
	return "BACnetConstructedDataTimeSynchronizationInterval"
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeSynchronization)
	lengthInBits += m.TimeSynchronization.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeSynchronizationInterval BACnetConstructedDataTimeSynchronizationInterval, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeSynchronizationInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeSynchronizationInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeSynchronization, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeSynchronization", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeSynchronization' field"))
	}
	m.TimeSynchronization = timeSynchronization

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), timeSynchronization)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeSynchronizationInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeSynchronizationInterval")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeSynchronizationInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeSynchronizationInterval")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeSynchronization", m.GetTimeSynchronization(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeSynchronization' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeSynchronizationInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeSynchronizationInterval")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) IsBACnetConstructedDataTimeSynchronizationInterval() {
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) deepCopy() *_BACnetConstructedDataTimeSynchronizationInterval {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeSynchronizationIntervalCopy := &_BACnetConstructedDataTimeSynchronizationInterval{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.TimeSynchronization.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeSynchronizationIntervalCopy
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
