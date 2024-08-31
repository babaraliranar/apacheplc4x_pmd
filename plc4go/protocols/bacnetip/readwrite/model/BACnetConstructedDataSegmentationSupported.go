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

// BACnetConstructedDataSegmentationSupported is the corresponding interface of BACnetConstructedDataSegmentationSupported
type BACnetConstructedDataSegmentationSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSegmentationSupported returns SegmentationSupported (property field)
	GetSegmentationSupported() BACnetSegmentationTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSegmentationTagged
}

// BACnetConstructedDataSegmentationSupportedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSegmentationSupported.
// This is useful for switch cases.
type BACnetConstructedDataSegmentationSupportedExactly interface {
	BACnetConstructedDataSegmentationSupported
	isBACnetConstructedDataSegmentationSupported() bool
}

// _BACnetConstructedDataSegmentationSupported is the data-structure of this message
type _BACnetConstructedDataSegmentationSupported struct {
	*_BACnetConstructedData
	SegmentationSupported BACnetSegmentationTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSegmentationSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SEGMENTATION_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSegmentationSupported) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetSegmentationSupported() BACnetSegmentationTagged {
	return m.SegmentationSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetActualValue() BACnetSegmentationTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetSegmentationTagged(m.GetSegmentationSupported())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSegmentationSupported factory function for _BACnetConstructedDataSegmentationSupported
func NewBACnetConstructedDataSegmentationSupported(segmentationSupported BACnetSegmentationTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSegmentationSupported {
	_result := &_BACnetConstructedDataSegmentationSupported{
		SegmentationSupported:  segmentationSupported,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSegmentationSupported(structType any) BACnetConstructedDataSegmentationSupported {
	if casted, ok := structType.(BACnetConstructedDataSegmentationSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSegmentationSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSegmentationSupported) GetTypeName() string {
	return "BACnetConstructedDataSegmentationSupported"
}

func (m *_BACnetConstructedDataSegmentationSupported) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (segmentationSupported)
	lengthInBits += m.SegmentationSupported.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSegmentationSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSegmentationSupportedParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSegmentationSupported, error) {
	return BACnetConstructedDataSegmentationSupportedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSegmentationSupportedParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataSegmentationSupported, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataSegmentationSupported, error) {
		return BACnetConstructedDataSegmentationSupportedParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataSegmentationSupportedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSegmentationSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSegmentationSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSegmentationSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentationSupported, err := ReadSimpleField[BACnetSegmentationTagged](ctx, "segmentationSupported", ReadComplex[BACnetSegmentationTagged](BACnetSegmentationTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentationSupported' field"))
	}

	actualValue, err := ReadVirtualField[BACnetSegmentationTagged](ctx, "actualValue", (*BACnetSegmentationTagged)(nil), segmentationSupported)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSegmentationSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSegmentationSupported")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSegmentationSupported{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SegmentationSupported: segmentationSupported,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSegmentationSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSegmentationSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSegmentationSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSegmentationSupported")
		}

		// Simple Field (segmentationSupported)
		if pushErr := writeBuffer.PushContext("segmentationSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for segmentationSupported")
		}
		_segmentationSupportedErr := writeBuffer.WriteSerializable(ctx, m.GetSegmentationSupported())
		if popErr := writeBuffer.PopContext("segmentationSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for segmentationSupported")
		}
		if _segmentationSupportedErr != nil {
			return errors.Wrap(_segmentationSupportedErr, "Error serializing 'segmentationSupported' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSegmentationSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSegmentationSupported")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSegmentationSupported) isBACnetConstructedDataSegmentationSupported() bool {
	return true
}

func (m *_BACnetConstructedDataSegmentationSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
