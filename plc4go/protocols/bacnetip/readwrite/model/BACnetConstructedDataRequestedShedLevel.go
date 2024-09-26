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

// BACnetConstructedDataRequestedShedLevel is the corresponding interface of BACnetConstructedDataRequestedShedLevel
type BACnetConstructedDataRequestedShedLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRequestedShedLevel returns RequestedShedLevel (property field)
	GetRequestedShedLevel() BACnetShedLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetShedLevel
	// IsBACnetConstructedDataRequestedShedLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRequestedShedLevel()
}

// _BACnetConstructedDataRequestedShedLevel is the data-structure of this message
type _BACnetConstructedDataRequestedShedLevel struct {
	BACnetConstructedDataContract
	RequestedShedLevel BACnetShedLevel
}

var _ BACnetConstructedDataRequestedShedLevel = (*_BACnetConstructedDataRequestedShedLevel)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRequestedShedLevel)(nil)

// NewBACnetConstructedDataRequestedShedLevel factory function for _BACnetConstructedDataRequestedShedLevel
func NewBACnetConstructedDataRequestedShedLevel(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, requestedShedLevel BACnetShedLevel, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRequestedShedLevel {
	if requestedShedLevel == nil {
		panic("requestedShedLevel of type BACnetShedLevel for BACnetConstructedDataRequestedShedLevel must not be nil")
	}
	_result := &_BACnetConstructedDataRequestedShedLevel{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RequestedShedLevel:            requestedShedLevel,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REQUESTED_SHED_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetRequestedShedLevel() BACnetShedLevel {
	return m.RequestedShedLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedShedLevel) GetActualValue() BACnetShedLevel {
	ctx := context.Background()
	_ = ctx
	return CastBACnetShedLevel(m.GetRequestedShedLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRequestedShedLevel(structType any) BACnetConstructedDataRequestedShedLevel {
	if casted, ok := structType.(BACnetConstructedDataRequestedShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRequestedShedLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetTypeName() string {
	return "BACnetConstructedDataRequestedShedLevel"
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (requestedShedLevel)
	lengthInBits += m.RequestedShedLevel.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRequestedShedLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRequestedShedLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRequestedShedLevel BACnetConstructedDataRequestedShedLevel, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRequestedShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRequestedShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestedShedLevel, err := ReadSimpleField[BACnetShedLevel](ctx, "requestedShedLevel", ReadComplex[BACnetShedLevel](BACnetShedLevelParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedShedLevel' field"))
	}
	m.RequestedShedLevel = requestedShedLevel

	actualValue, err := ReadVirtualField[BACnetShedLevel](ctx, "actualValue", (*BACnetShedLevel)(nil), requestedShedLevel)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRequestedShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRequestedShedLevel")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRequestedShedLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRequestedShedLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRequestedShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRequestedShedLevel")
		}

		if err := WriteSimpleField[BACnetShedLevel](ctx, "requestedShedLevel", m.GetRequestedShedLevel(), WriteComplex[BACnetShedLevel](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedShedLevel' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRequestedShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRequestedShedLevel")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRequestedShedLevel) IsBACnetConstructedDataRequestedShedLevel() {}

func (m *_BACnetConstructedDataRequestedShedLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRequestedShedLevel) deepCopy() *_BACnetConstructedDataRequestedShedLevel {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRequestedShedLevelCopy := &_BACnetConstructedDataRequestedShedLevel{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.RequestedShedLevel.DeepCopy().(BACnetShedLevel),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRequestedShedLevelCopy
}

func (m *_BACnetConstructedDataRequestedShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
