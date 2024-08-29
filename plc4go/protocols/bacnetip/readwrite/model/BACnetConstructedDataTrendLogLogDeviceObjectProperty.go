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

// BACnetConstructedDataTrendLogLogDeviceObjectProperty is the corresponding interface of BACnetConstructedDataTrendLogLogDeviceObjectProperty
type BACnetConstructedDataTrendLogLogDeviceObjectProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLogDeviceObjectProperty returns LogDeviceObjectProperty (property field)
	GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectPropertyReference
}

// BACnetConstructedDataTrendLogLogDeviceObjectPropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTrendLogLogDeviceObjectProperty.
// This is useful for switch cases.
type BACnetConstructedDataTrendLogLogDeviceObjectPropertyExactly interface {
	BACnetConstructedDataTrendLogLogDeviceObjectProperty
	isBACnetConstructedDataTrendLogLogDeviceObjectProperty() bool
}

// _BACnetConstructedDataTrendLogLogDeviceObjectProperty is the data-structure of this message
type _BACnetConstructedDataTrendLogLogDeviceObjectProperty struct {
	*_BACnetConstructedData
	LogDeviceObjectProperty BACnetDeviceObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_DEVICE_OBJECT_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference {
	return m.LogDeviceObjectProperty
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetActualValue() BACnetDeviceObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectPropertyReference(m.GetLogDeviceObjectProperty())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrendLogLogDeviceObjectProperty factory function for _BACnetConstructedDataTrendLogLogDeviceObjectProperty
func NewBACnetConstructedDataTrendLogLogDeviceObjectProperty(logDeviceObjectProperty BACnetDeviceObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	_result := &_BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		LogDeviceObjectProperty: logDeviceObjectProperty,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogLogDeviceObjectProperty(structType any) BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	if casted, ok := structType.(BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetTypeName() string {
	return "BACnetConstructedDataTrendLogLogDeviceObjectProperty"
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (logDeviceObjectProperty)
	lengthInBits += m.LogDeviceObjectProperty.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTrendLogLogDeviceObjectPropertyParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogLogDeviceObjectProperty, error) {
	return BACnetConstructedDataTrendLogLogDeviceObjectPropertyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTrendLogLogDeviceObjectPropertyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogLogDeviceObjectProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logDeviceObjectProperty)
	if pullErr := readBuffer.PullContext("logDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logDeviceObjectProperty")
	}
	_logDeviceObjectProperty, _logDeviceObjectPropertyErr := BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx, readBuffer)
	if _logDeviceObjectPropertyErr != nil {
		return nil, errors.Wrap(_logDeviceObjectPropertyErr, "Error parsing 'logDeviceObjectProperty' field of BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}
	logDeviceObjectProperty := _logDeviceObjectProperty.(BACnetDeviceObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("logDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logDeviceObjectProperty")
	}

	// Virtual field
	_actualValue := logDeviceObjectProperty
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LogDeviceObjectProperty: logDeviceObjectProperty,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}

		// Simple Field (logDeviceObjectProperty)
		if pushErr := writeBuffer.PushContext("logDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logDeviceObjectProperty")
		}
		_logDeviceObjectPropertyErr := writeBuffer.WriteSerializable(ctx, m.GetLogDeviceObjectProperty())
		if popErr := writeBuffer.PopContext("logDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logDeviceObjectProperty")
		}
		if _logDeviceObjectPropertyErr != nil {
			return errors.Wrap(_logDeviceObjectPropertyErr, "Error serializing 'logDeviceObjectProperty' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) isBACnetConstructedDataTrendLogLogDeviceObjectProperty() bool {
	return true
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
