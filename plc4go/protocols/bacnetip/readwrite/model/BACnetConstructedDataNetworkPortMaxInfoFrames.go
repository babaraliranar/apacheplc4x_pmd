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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataNetworkPortMaxInfoFrames is the corresponding interface of BACnetConstructedDataNetworkPortMaxInfoFrames
type BACnetConstructedDataNetworkPortMaxInfoFrames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataNetworkPortMaxInfoFramesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNetworkPortMaxInfoFrames.
// This is useful for switch cases.
type BACnetConstructedDataNetworkPortMaxInfoFramesExactly interface {
	BACnetConstructedDataNetworkPortMaxInfoFrames
	isBACnetConstructedDataNetworkPortMaxInfoFrames() bool
}

// _BACnetConstructedDataNetworkPortMaxInfoFrames is the data-structure of this message
type _BACnetConstructedDataNetworkPortMaxInfoFrames struct {
	*_BACnetConstructedData
	MaxInfoFrames BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_NETWORK_PORT
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_INFO_FRAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxInfoFrames())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkPortMaxInfoFrames factory function for _BACnetConstructedDataNetworkPortMaxInfoFrames
func NewBACnetConstructedDataNetworkPortMaxInfoFrames(maxInfoFrames BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkPortMaxInfoFrames {
	_result := &_BACnetConstructedDataNetworkPortMaxInfoFrames{
		MaxInfoFrames:          maxInfoFrames,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkPortMaxInfoFrames(structType any) BACnetConstructedDataNetworkPortMaxInfoFrames {
	if casted, ok := structType.(BACnetConstructedDataNetworkPortMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkPortMaxInfoFrames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataNetworkPortMaxInfoFrames"
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataNetworkPortMaxInfoFramesParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkPortMaxInfoFrames, error) {
	return BACnetConstructedDataNetworkPortMaxInfoFramesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataNetworkPortMaxInfoFramesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkPortMaxInfoFrames, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkPortMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxInfoFrames)
	if pullErr := readBuffer.PullContext("maxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxInfoFrames")
	}
	_maxInfoFrames, _maxInfoFramesErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _maxInfoFramesErr != nil {
		return nil, errors.Wrap(_maxInfoFramesErr, "Error parsing 'maxInfoFrames' field of BACnetConstructedDataNetworkPortMaxInfoFrames")
	}
	maxInfoFrames := _maxInfoFrames.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxInfoFrames")
	}

	// Virtual field
	_actualValue := maxInfoFrames
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkPortMaxInfoFrames")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNetworkPortMaxInfoFrames{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxInfoFrames: maxInfoFrames,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkPortMaxInfoFrames")
		}

		// Simple Field (maxInfoFrames)
		if pushErr := writeBuffer.PushContext("maxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxInfoFrames")
		}
		_maxInfoFramesErr := writeBuffer.WriteSerializable(ctx, m.GetMaxInfoFrames())
		if popErr := writeBuffer.PopContext("maxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxInfoFrames")
		}
		if _maxInfoFramesErr != nil {
			return errors.Wrap(_maxInfoFramesErr, "Error serializing 'maxInfoFrames' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkPortMaxInfoFrames")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) isBACnetConstructedDataNetworkPortMaxInfoFrames() bool {
	return true
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
