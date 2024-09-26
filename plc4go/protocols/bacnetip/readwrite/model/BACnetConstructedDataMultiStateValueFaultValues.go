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

// BACnetConstructedDataMultiStateValueFaultValues is the corresponding interface of BACnetConstructedDataMultiStateValueFaultValues
type BACnetConstructedDataMultiStateValueFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultValues returns FaultValues (property field)
	GetFaultValues() []BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMultiStateValueFaultValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMultiStateValueFaultValues()
}

// _BACnetConstructedDataMultiStateValueFaultValues is the data-structure of this message
type _BACnetConstructedDataMultiStateValueFaultValues struct {
	BACnetConstructedDataContract
	FaultValues []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMultiStateValueFaultValues = (*_BACnetConstructedDataMultiStateValueFaultValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMultiStateValueFaultValues)(nil)

// NewBACnetConstructedDataMultiStateValueFaultValues factory function for _BACnetConstructedDataMultiStateValueFaultValues
func NewBACnetConstructedDataMultiStateValueFaultValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultValues []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateValueFaultValues {
	_result := &_BACnetConstructedDataMultiStateValueFaultValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultValues:                   faultValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_VALUE
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetFaultValues() []BACnetApplicationTagUnsignedInteger {
	return m.FaultValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateValueFaultValues(structType any) BACnetConstructedDataMultiStateValueFaultValues {
	if casted, ok := structType.(BACnetConstructedDataMultiStateValueFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateValueFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetTypeName() string {
	return "BACnetConstructedDataMultiStateValueFaultValues"
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FaultValues) > 0 {
		for _, element := range m.FaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateValueFaultValues BACnetConstructedDataMultiStateValueFaultValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateValueFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateValueFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultValues, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "faultValues", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultValues' field"))
	}
	m.FaultValues = faultValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateValueFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateValueFaultValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateValueFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateValueFaultValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultValues", m.GetFaultValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateValueFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateValueFaultValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) IsBACnetConstructedDataMultiStateValueFaultValues() {
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) deepCopy() *_BACnetConstructedDataMultiStateValueFaultValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMultiStateValueFaultValuesCopy := &_BACnetConstructedDataMultiStateValueFaultValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.FaultValues),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMultiStateValueFaultValuesCopy
}

func (m *_BACnetConstructedDataMultiStateValueFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
