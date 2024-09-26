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

// BACnetOptionalUnsigned is the corresponding interface of BACnetOptionalUnsigned
type BACnetOptionalUnsigned interface {
	BACnetOptionalUnsignedContract
	BACnetOptionalUnsignedRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetOptionalUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalUnsigned()
}

// BACnetOptionalUnsignedContract provides a set of functions which can be overwritten by a sub struct
type BACnetOptionalUnsignedContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetOptionalUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalUnsigned()
}

// BACnetOptionalUnsignedRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetOptionalUnsignedRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetOptionalUnsigned is the data-structure of this message
type _BACnetOptionalUnsigned struct {
	_SubType        BACnetOptionalUnsigned
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetOptionalUnsignedContract = (*_BACnetOptionalUnsigned)(nil)

// NewBACnetOptionalUnsigned factory function for _BACnetOptionalUnsigned
func NewBACnetOptionalUnsigned(peekedTagHeader BACnetTagHeader) *_BACnetOptionalUnsigned {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetOptionalUnsigned must not be nil")
	}
	return &_BACnetOptionalUnsigned{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalUnsigned) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetOptionalUnsigned) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOptionalUnsigned(structType any) BACnetOptionalUnsigned {
	if casted, ok := structType.(BACnetOptionalUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalUnsigned) GetTypeName() string {
	return "BACnetOptionalUnsigned"
}

func (m *_BACnetOptionalUnsigned) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetOptionalUnsignedParse[T BACnetOptionalUnsigned](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetOptionalUnsignedParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalUnsignedParseWithBufferProducer[T BACnetOptionalUnsigned]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetOptionalUnsignedParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetOptionalUnsignedParseWithBuffer[T BACnetOptionalUnsigned](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetOptionalUnsigned{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetOptionalUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetOptionalUnsigned BACnetOptionalUnsigned, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetOptionalUnsigned
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalUnsignedNull
		if _child, err = new(_BACnetOptionalUnsignedNull).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalUnsignedNull for type-switch of BACnetOptionalUnsigned")
		}
	case 0 == 0: // BACnetOptionalUnsignedValue
		if _child, err = new(_BACnetOptionalUnsignedValue).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalUnsignedValue for type-switch of BACnetOptionalUnsigned")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalUnsigned")
	}

	return _child, nil
}

func (pm *_BACnetOptionalUnsigned) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetOptionalUnsigned, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetOptionalUnsigned"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalUnsigned")
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

	if popErr := writeBuffer.PopContext("BACnetOptionalUnsigned"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalUnsigned")
	}
	return nil
}

func (m *_BACnetOptionalUnsigned) IsBACnetOptionalUnsigned() {}

func (m *_BACnetOptionalUnsigned) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalUnsigned) deepCopy() *_BACnetOptionalUnsigned {
	if m == nil {
		return nil
	}
	_BACnetOptionalUnsignedCopy := &_BACnetOptionalUnsigned{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetOptionalUnsignedCopy
}
