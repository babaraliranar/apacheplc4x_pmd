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

// BACnetConstructedDataUpdateKeySetTimeout is the corresponding interface of BACnetConstructedDataUpdateKeySetTimeout
type BACnetConstructedDataUpdateKeySetTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUpdateKeySetTimeout returns UpdateKeySetTimeout (property field)
	GetUpdateKeySetTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataUpdateKeySetTimeout is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUpdateKeySetTimeout()
}

// _BACnetConstructedDataUpdateKeySetTimeout is the data-structure of this message
type _BACnetConstructedDataUpdateKeySetTimeout struct {
	BACnetConstructedDataContract
	UpdateKeySetTimeout BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataUpdateKeySetTimeout = (*_BACnetConstructedDataUpdateKeySetTimeout)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUpdateKeySetTimeout)(nil)

// NewBACnetConstructedDataUpdateKeySetTimeout factory function for _BACnetConstructedDataUpdateKeySetTimeout
func NewBACnetConstructedDataUpdateKeySetTimeout(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, updateKeySetTimeout BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUpdateKeySetTimeout {
	if updateKeySetTimeout == nil {
		panic("updateKeySetTimeout of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataUpdateKeySetTimeout must not be nil")
	}
	_result := &_BACnetConstructedDataUpdateKeySetTimeout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		UpdateKeySetTimeout:           updateKeySetTimeout,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_KEY_SET_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetUpdateKeySetTimeout() BACnetApplicationTagUnsignedInteger {
	return m.UpdateKeySetTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetUpdateKeySetTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUpdateKeySetTimeout(structType any) BACnetConstructedDataUpdateKeySetTimeout {
	if casted, ok := structType.(BACnetConstructedDataUpdateKeySetTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUpdateKeySetTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetTypeName() string {
	return "BACnetConstructedDataUpdateKeySetTimeout"
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (updateKeySetTimeout)
	lengthInBits += m.UpdateKeySetTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUpdateKeySetTimeout BACnetConstructedDataUpdateKeySetTimeout, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUpdateKeySetTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUpdateKeySetTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	updateKeySetTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "updateKeySetTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateKeySetTimeout' field"))
	}
	m.UpdateKeySetTimeout = updateKeySetTimeout

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), updateKeySetTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUpdateKeySetTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUpdateKeySetTimeout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUpdateKeySetTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUpdateKeySetTimeout")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "updateKeySetTimeout", m.GetUpdateKeySetTimeout(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'updateKeySetTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUpdateKeySetTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUpdateKeySetTimeout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) IsBACnetConstructedDataUpdateKeySetTimeout() {}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) deepCopy() *_BACnetConstructedDataUpdateKeySetTimeout {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUpdateKeySetTimeoutCopy := &_BACnetConstructedDataUpdateKeySetTimeout{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.UpdateKeySetTimeout.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUpdateKeySetTimeoutCopy
}

func (m *_BACnetConstructedDataUpdateKeySetTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
