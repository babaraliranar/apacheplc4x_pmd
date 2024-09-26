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

// BACnetConstructedDataIPv6ZoneIndex is the corresponding interface of BACnetConstructedDataIPv6ZoneIndex
type BACnetConstructedDataIPv6ZoneIndex interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6ZoneIndex returns Ipv6ZoneIndex (property field)
	GetIpv6ZoneIndex() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataIPv6ZoneIndex is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6ZoneIndex()
}

// _BACnetConstructedDataIPv6ZoneIndex is the data-structure of this message
type _BACnetConstructedDataIPv6ZoneIndex struct {
	BACnetConstructedDataContract
	Ipv6ZoneIndex BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataIPv6ZoneIndex = (*_BACnetConstructedDataIPv6ZoneIndex)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6ZoneIndex)(nil)

// NewBACnetConstructedDataIPv6ZoneIndex factory function for _BACnetConstructedDataIPv6ZoneIndex
func NewBACnetConstructedDataIPv6ZoneIndex(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6ZoneIndex BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6ZoneIndex {
	if ipv6ZoneIndex == nil {
		panic("ipv6ZoneIndex of type BACnetApplicationTagCharacterString for BACnetConstructedDataIPv6ZoneIndex must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6ZoneIndex{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6ZoneIndex:                 ipv6ZoneIndex,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_ZONE_INDEX
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetIpv6ZoneIndex() BACnetApplicationTagCharacterString {
	return m.Ipv6ZoneIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetIpv6ZoneIndex())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6ZoneIndex(structType any) BACnetConstructedDataIPv6ZoneIndex {
	if casted, ok := structType.(BACnetConstructedDataIPv6ZoneIndex); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6ZoneIndex); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetTypeName() string {
	return "BACnetConstructedDataIPv6ZoneIndex"
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6ZoneIndex)
	lengthInBits += m.Ipv6ZoneIndex.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6ZoneIndex BACnetConstructedDataIPv6ZoneIndex, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6ZoneIndex"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6ZoneIndex")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6ZoneIndex, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "ipv6ZoneIndex", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6ZoneIndex' field"))
	}
	m.Ipv6ZoneIndex = ipv6ZoneIndex

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), ipv6ZoneIndex)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6ZoneIndex"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6ZoneIndex")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6ZoneIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6ZoneIndex")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "ipv6ZoneIndex", m.GetIpv6ZoneIndex(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6ZoneIndex' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6ZoneIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6ZoneIndex")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) IsBACnetConstructedDataIPv6ZoneIndex() {}

func (m *_BACnetConstructedDataIPv6ZoneIndex) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) deepCopy() *_BACnetConstructedDataIPv6ZoneIndex {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6ZoneIndexCopy := &_BACnetConstructedDataIPv6ZoneIndex{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Ipv6ZoneIndex.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6ZoneIndexCopy
}

func (m *_BACnetConstructedDataIPv6ZoneIndex) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
