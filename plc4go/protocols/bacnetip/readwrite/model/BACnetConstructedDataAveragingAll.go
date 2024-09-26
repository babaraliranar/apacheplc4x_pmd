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

// BACnetConstructedDataAveragingAll is the corresponding interface of BACnetConstructedDataAveragingAll
type BACnetConstructedDataAveragingAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAveragingAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAveragingAll()
}

// _BACnetConstructedDataAveragingAll is the data-structure of this message
type _BACnetConstructedDataAveragingAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAveragingAll = (*_BACnetConstructedDataAveragingAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAveragingAll)(nil)

// NewBACnetConstructedDataAveragingAll factory function for _BACnetConstructedDataAveragingAll
func NewBACnetConstructedDataAveragingAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAveragingAll {
	_result := &_BACnetConstructedDataAveragingAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAveragingAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_AVERAGING
}

func (m *_BACnetConstructedDataAveragingAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAveragingAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAveragingAll(structType any) BACnetConstructedDataAveragingAll {
	if casted, ok := structType.(BACnetConstructedDataAveragingAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAveragingAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAveragingAll) GetTypeName() string {
	return "BACnetConstructedDataAveragingAll"
}

func (m *_BACnetConstructedDataAveragingAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAveragingAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAveragingAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAveragingAll BACnetConstructedDataAveragingAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAveragingAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAveragingAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAveragingAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAveragingAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAveragingAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAveragingAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAveragingAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAveragingAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAveragingAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAveragingAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAveragingAll) IsBACnetConstructedDataAveragingAll() {}

func (m *_BACnetConstructedDataAveragingAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAveragingAll) deepCopy() *_BACnetConstructedDataAveragingAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAveragingAllCopy := &_BACnetConstructedDataAveragingAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAveragingAllCopy
}

func (m *_BACnetConstructedDataAveragingAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
