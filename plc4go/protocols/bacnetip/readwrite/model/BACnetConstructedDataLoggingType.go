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

// BACnetConstructedDataLoggingType is the corresponding interface of BACnetConstructedDataLoggingType
type BACnetConstructedDataLoggingType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLoggingType returns LoggingType (property field)
	GetLoggingType() BACnetLoggingTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLoggingTypeTagged
	// IsBACnetConstructedDataLoggingType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLoggingType()
}

// _BACnetConstructedDataLoggingType is the data-structure of this message
type _BACnetConstructedDataLoggingType struct {
	BACnetConstructedDataContract
	LoggingType BACnetLoggingTypeTagged
}

var _ BACnetConstructedDataLoggingType = (*_BACnetConstructedDataLoggingType)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLoggingType)(nil)

// NewBACnetConstructedDataLoggingType factory function for _BACnetConstructedDataLoggingType
func NewBACnetConstructedDataLoggingType(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, loggingType BACnetLoggingTypeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoggingType {
	if loggingType == nil {
		panic("loggingType of type BACnetLoggingTypeTagged for BACnetConstructedDataLoggingType must not be nil")
	}
	_result := &_BACnetConstructedDataLoggingType{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LoggingType:                   loggingType,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoggingType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLoggingType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOGGING_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoggingType) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingType) GetLoggingType() BACnetLoggingTypeTagged {
	return m.LoggingType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingType) GetActualValue() BACnetLoggingTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLoggingTypeTagged(m.GetLoggingType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoggingType(structType any) BACnetConstructedDataLoggingType {
	if casted, ok := structType.(BACnetConstructedDataLoggingType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoggingType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoggingType) GetTypeName() string {
	return "BACnetConstructedDataLoggingType"
}

func (m *_BACnetConstructedDataLoggingType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (loggingType)
	lengthInBits += m.LoggingType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLoggingType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLoggingType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLoggingType BACnetConstructedDataLoggingType, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoggingType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoggingType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	loggingType, err := ReadSimpleField[BACnetLoggingTypeTagged](ctx, "loggingType", ReadComplex[BACnetLoggingTypeTagged](BACnetLoggingTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'loggingType' field"))
	}
	m.LoggingType = loggingType

	actualValue, err := ReadVirtualField[BACnetLoggingTypeTagged](ctx, "actualValue", (*BACnetLoggingTypeTagged)(nil), loggingType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoggingType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoggingType")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLoggingType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoggingType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoggingType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoggingType")
		}

		if err := WriteSimpleField[BACnetLoggingTypeTagged](ctx, "loggingType", m.GetLoggingType(), WriteComplex[BACnetLoggingTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'loggingType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoggingType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoggingType")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoggingType) IsBACnetConstructedDataLoggingType() {}

func (m *_BACnetConstructedDataLoggingType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLoggingType) deepCopy() *_BACnetConstructedDataLoggingType {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLoggingTypeCopy := &_BACnetConstructedDataLoggingType{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LoggingType.DeepCopy().(BACnetLoggingTypeTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLoggingTypeCopy
}

func (m *_BACnetConstructedDataLoggingType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
