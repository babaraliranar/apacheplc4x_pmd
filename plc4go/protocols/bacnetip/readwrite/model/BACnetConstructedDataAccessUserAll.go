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

// BACnetConstructedDataAccessUserAll is the corresponding interface of BACnetConstructedDataAccessUserAll
type BACnetConstructedDataAccessUserAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAccessUserAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessUserAll()
}

// _BACnetConstructedDataAccessUserAll is the data-structure of this message
type _BACnetConstructedDataAccessUserAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAccessUserAll = (*_BACnetConstructedDataAccessUserAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessUserAll)(nil)

// NewBACnetConstructedDataAccessUserAll factory function for _BACnetConstructedDataAccessUserAll
func NewBACnetConstructedDataAccessUserAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessUserAll {
	_result := &_BACnetConstructedDataAccessUserAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessUserAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_USER
}

func (m *_BACnetConstructedDataAccessUserAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessUserAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessUserAll(structType any) BACnetConstructedDataAccessUserAll {
	if casted, ok := structType.(BACnetConstructedDataAccessUserAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessUserAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessUserAll) GetTypeName() string {
	return "BACnetConstructedDataAccessUserAll"
}

func (m *_BACnetConstructedDataAccessUserAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessUserAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessUserAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessUserAll BACnetConstructedDataAccessUserAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessUserAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessUserAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessUserAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessUserAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessUserAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessUserAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessUserAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessUserAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessUserAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessUserAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessUserAll) IsBACnetConstructedDataAccessUserAll() {}

func (m *_BACnetConstructedDataAccessUserAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessUserAll) deepCopy() *_BACnetConstructedDataAccessUserAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessUserAllCopy := &_BACnetConstructedDataAccessUserAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessUserAllCopy
}

func (m *_BACnetConstructedDataAccessUserAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
