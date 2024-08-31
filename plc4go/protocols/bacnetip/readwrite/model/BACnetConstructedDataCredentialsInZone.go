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

// BACnetConstructedDataCredentialsInZone is the corresponding interface of BACnetConstructedDataCredentialsInZone
type BACnetConstructedDataCredentialsInZone interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCredentialsInZone returns CredentialsInZone (property field)
	GetCredentialsInZone() []BACnetDeviceObjectReference
}

// BACnetConstructedDataCredentialsInZoneExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCredentialsInZone.
// This is useful for switch cases.
type BACnetConstructedDataCredentialsInZoneExactly interface {
	BACnetConstructedDataCredentialsInZone
	isBACnetConstructedDataCredentialsInZone() bool
}

// _BACnetConstructedDataCredentialsInZone is the data-structure of this message
type _BACnetConstructedDataCredentialsInZone struct {
	*_BACnetConstructedData
	CredentialsInZone []BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialsInZone) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIALS_IN_ZONE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCredentialsInZone) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetCredentialsInZone() []BACnetDeviceObjectReference {
	return m.CredentialsInZone
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialsInZone factory function for _BACnetConstructedDataCredentialsInZone
func NewBACnetConstructedDataCredentialsInZone(credentialsInZone []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialsInZone {
	_result := &_BACnetConstructedDataCredentialsInZone{
		CredentialsInZone:      credentialsInZone,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialsInZone(structType any) BACnetConstructedDataCredentialsInZone {
	if casted, ok := structType.(BACnetConstructedDataCredentialsInZone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialsInZone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialsInZone) GetTypeName() string {
	return "BACnetConstructedDataCredentialsInZone"
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.CredentialsInZone) > 0 {
		for _, element := range m.CredentialsInZone {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCredentialsInZoneParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialsInZone, error) {
	return BACnetConstructedDataCredentialsInZoneParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCredentialsInZoneParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialsInZone, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialsInZone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialsInZone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	credentialsInZone, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "credentialsInZone", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'credentialsInZone' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialsInZone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialsInZone")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCredentialsInZone{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CredentialsInZone: credentialsInZone,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCredentialsInZone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCredentialsInZone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialsInZone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialsInZone")
		}

		// Array Field (credentialsInZone)
		if pushErr := writeBuffer.PushContext("credentialsInZone", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for credentialsInZone")
		}
		for _curItem, _element := range m.GetCredentialsInZone() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetCredentialsInZone()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'credentialsInZone' field")
			}
		}
		if popErr := writeBuffer.PopContext("credentialsInZone", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for credentialsInZone")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialsInZone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialsInZone")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialsInZone) isBACnetConstructedDataCredentialsInZone() bool {
	return true
}

func (m *_BACnetConstructedDataCredentialsInZone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
