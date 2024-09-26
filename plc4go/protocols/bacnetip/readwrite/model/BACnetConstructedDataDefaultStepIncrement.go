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

// BACnetConstructedDataDefaultStepIncrement is the corresponding interface of BACnetConstructedDataDefaultStepIncrement
type BACnetConstructedDataDefaultStepIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDefaultStepIncrement returns DefaultStepIncrement (property field)
	GetDefaultStepIncrement() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataDefaultStepIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDefaultStepIncrement()
}

// _BACnetConstructedDataDefaultStepIncrement is the data-structure of this message
type _BACnetConstructedDataDefaultStepIncrement struct {
	BACnetConstructedDataContract
	DefaultStepIncrement BACnetApplicationTagReal
}

var _ BACnetConstructedDataDefaultStepIncrement = (*_BACnetConstructedDataDefaultStepIncrement)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDefaultStepIncrement)(nil)

// NewBACnetConstructedDataDefaultStepIncrement factory function for _BACnetConstructedDataDefaultStepIncrement
func NewBACnetConstructedDataDefaultStepIncrement(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, defaultStepIncrement BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDefaultStepIncrement {
	if defaultStepIncrement == nil {
		panic("defaultStepIncrement of type BACnetApplicationTagReal for BACnetConstructedDataDefaultStepIncrement must not be nil")
	}
	_result := &_BACnetConstructedDataDefaultStepIncrement{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DefaultStepIncrement:          defaultStepIncrement,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_STEP_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetDefaultStepIncrement() BACnetApplicationTagReal {
	return m.DefaultStepIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetDefaultStepIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDefaultStepIncrement(structType any) BACnetConstructedDataDefaultStepIncrement {
	if casted, ok := structType.(BACnetConstructedDataDefaultStepIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultStepIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetTypeName() string {
	return "BACnetConstructedDataDefaultStepIncrement"
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (defaultStepIncrement)
	lengthInBits += m.DefaultStepIncrement.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDefaultStepIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDefaultStepIncrement BACnetConstructedDataDefaultStepIncrement, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultStepIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDefaultStepIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	defaultStepIncrement, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "defaultStepIncrement", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'defaultStepIncrement' field"))
	}
	m.DefaultStepIncrement = defaultStepIncrement

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), defaultStepIncrement)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultStepIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDefaultStepIncrement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDefaultStepIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDefaultStepIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultStepIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDefaultStepIncrement")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "defaultStepIncrement", m.GetDefaultStepIncrement(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'defaultStepIncrement' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultStepIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDefaultStepIncrement")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDefaultStepIncrement) IsBACnetConstructedDataDefaultStepIncrement() {}

func (m *_BACnetConstructedDataDefaultStepIncrement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDefaultStepIncrement) deepCopy() *_BACnetConstructedDataDefaultStepIncrement {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDefaultStepIncrementCopy := &_BACnetConstructedDataDefaultStepIncrement{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DefaultStepIncrement.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDefaultStepIncrementCopy
}

func (m *_BACnetConstructedDataDefaultStepIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
