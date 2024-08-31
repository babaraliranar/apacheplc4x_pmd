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

// BACnetConstructedDataObjectIdentifier is the corresponding interface of BACnetConstructedDataObjectIdentifier
type BACnetConstructedDataObjectIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
}

// BACnetConstructedDataObjectIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataObjectIdentifier.
// This is useful for switch cases.
type BACnetConstructedDataObjectIdentifierExactly interface {
	BACnetConstructedDataObjectIdentifier
	isBACnetConstructedDataObjectIdentifier() bool
}

// _BACnetConstructedDataObjectIdentifier is the data-structure of this message
type _BACnetConstructedDataObjectIdentifier struct {
	*_BACnetConstructedData
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataObjectIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataObjectIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataObjectIdentifier) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataObjectIdentifier) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataObjectIdentifier) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataObjectIdentifier) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetObjectIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataObjectIdentifier factory function for _BACnetConstructedDataObjectIdentifier
func NewBACnetConstructedDataObjectIdentifier(objectIdentifier BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataObjectIdentifier {
	_result := &_BACnetConstructedDataObjectIdentifier{
		ObjectIdentifier:       objectIdentifier,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataObjectIdentifier(structType any) BACnetConstructedDataObjectIdentifier {
	if casted, ok := structType.(BACnetConstructedDataObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataObjectIdentifier) GetTypeName() string {
	return "BACnetConstructedDataObjectIdentifier"
}

func (m *_BACnetConstructedDataObjectIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataObjectIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataObjectIdentifierParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectIdentifier, error) {
	return BACnetConstructedDataObjectIdentifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataObjectIdentifierParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataObjectIdentifier, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataObjectIdentifier, error) {
		return BACnetConstructedDataObjectIdentifierParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataObjectIdentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagObjectIdentifier](ctx, "actualValue", (*BACnetApplicationTagObjectIdentifier)(nil), objectIdentifier)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataObjectIdentifier{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ObjectIdentifier: objectIdentifier,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataObjectIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataObjectIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectIdentifier")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectIdentifier")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataObjectIdentifier) isBACnetConstructedDataObjectIdentifier() bool {
	return true
}

func (m *_BACnetConstructedDataObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
