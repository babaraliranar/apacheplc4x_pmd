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

// BACnetEventParameterChangeOfValueCivCriteria is the corresponding interface of BACnetEventParameterChangeOfValueCivCriteria
type BACnetEventParameterChangeOfValueCivCriteria interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetEventParameterChangeOfValueCivCriteriaExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfValueCivCriteria.
// This is useful for switch cases.
type BACnetEventParameterChangeOfValueCivCriteriaExactly interface {
	BACnetEventParameterChangeOfValueCivCriteria
	isBACnetEventParameterChangeOfValueCivCriteria() bool
}

// _BACnetEventParameterChangeOfValueCivCriteria is the data-structure of this message
type _BACnetEventParameterChangeOfValueCivCriteria struct {
	_BACnetEventParameterChangeOfValueCivCriteriaChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

type _BACnetEventParameterChangeOfValueCivCriteriaChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetEventParameterChangeOfValueCivCriteriaParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventParameterChangeOfValueCivCriteria, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetEventParameterChangeOfValueCivCriteriaChild interface {
	utils.Serializable
	InitializeParent(parent BACnetEventParameterChangeOfValueCivCriteria, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetEventParameterChangeOfValueCivCriteria

	GetTypeName() string
	BACnetEventParameterChangeOfValueCivCriteria
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfValueCivCriteria factory function for _BACnetEventParameterChangeOfValueCivCriteria
func NewBACnetEventParameterChangeOfValueCivCriteria(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfValueCivCriteria {
	return &_BACnetEventParameterChangeOfValueCivCriteria{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfValueCivCriteria(structType any) BACnetEventParameterChangeOfValueCivCriteria {
	if casted, ok := structType.(BACnetEventParameterChangeOfValueCivCriteria); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfValueCivCriteria); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetTypeName() string {
	return "BACnetEventParameterChangeOfValueCivCriteria"
}

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfValueCivCriteriaParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfValueCivCriteria, error) {
	return BACnetEventParameterChangeOfValueCivCriteriaParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfValueCivCriteriaParseWithBufferProducer[T BACnetEventParameterChangeOfValueCivCriteria](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetEventParameterChangeOfValueCivCriteriaParseWithBuffer(ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetEventParameterChangeOfValueCivCriteriaParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfValueCivCriteria, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfValueCivCriteria"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfValueCivCriteria")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetEventParameterChangeOfValueCivCriteriaChildSerializeRequirement interface {
		BACnetEventParameterChangeOfValueCivCriteria
		InitializeParent(BACnetEventParameterChangeOfValueCivCriteria, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetEventParameterChangeOfValueCivCriteria
	}
	var _childTemp any
	var _child BACnetEventParameterChangeOfValueCivCriteriaChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetEventParameterChangeOfValueCivCriteriaBitmask
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfValueCivCriteriaBitmaskParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == uint8(1): // BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementParseWithBuffer(ctx, readBuffer, tagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetEventParameterChangeOfValueCivCriteria")
	}
	_child = _childTemp.(BACnetEventParameterChangeOfValueCivCriteriaChildSerializeRequirement)

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfValueCivCriteria"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfValueCivCriteria")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetEventParameterChangeOfValueCivCriteria) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventParameterChangeOfValueCivCriteria, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfValueCivCriteria"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfValueCivCriteria")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfValueCivCriteria"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfValueCivCriteria")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfValueCivCriteria) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfValueCivCriteria) isBACnetEventParameterChangeOfValueCivCriteria() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfValueCivCriteria) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
