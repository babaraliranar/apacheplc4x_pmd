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

// BACnetOptionalBinaryPV is the corresponding interface of BACnetOptionalBinaryPV
type BACnetOptionalBinaryPV interface {
	BACnetOptionalBinaryPVContract
	BACnetOptionalBinaryPVRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetOptionalBinaryPV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalBinaryPV()
}

// BACnetOptionalBinaryPVContract provides a set of functions which can be overwritten by a sub struct
type BACnetOptionalBinaryPVContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetOptionalBinaryPV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalBinaryPV()
}

// BACnetOptionalBinaryPVRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetOptionalBinaryPVRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetOptionalBinaryPV is the data-structure of this message
type _BACnetOptionalBinaryPV struct {
	_SubType        BACnetOptionalBinaryPV
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetOptionalBinaryPVContract = (*_BACnetOptionalBinaryPV)(nil)

// NewBACnetOptionalBinaryPV factory function for _BACnetOptionalBinaryPV
func NewBACnetOptionalBinaryPV(peekedTagHeader BACnetTagHeader) *_BACnetOptionalBinaryPV {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetOptionalBinaryPV must not be nil")
	}
	return &_BACnetOptionalBinaryPV{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalBinaryPV) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetOptionalBinaryPV) GetPeekedTagNumber() uint8 {
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
func CastBACnetOptionalBinaryPV(structType any) BACnetOptionalBinaryPV {
	if casted, ok := structType.(BACnetOptionalBinaryPV); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPV); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalBinaryPV) GetTypeName() string {
	return "BACnetOptionalBinaryPV"
}

func (m *_BACnetOptionalBinaryPV) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalBinaryPV) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetOptionalBinaryPVParse[T BACnetOptionalBinaryPV](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetOptionalBinaryPVParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalBinaryPVParseWithBufferProducer[T BACnetOptionalBinaryPV]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetOptionalBinaryPVParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetOptionalBinaryPVParseWithBuffer[T BACnetOptionalBinaryPV](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetOptionalBinaryPV{}).parse(ctx, readBuffer)
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

func (m *_BACnetOptionalBinaryPV) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetOptionalBinaryPV BACnetOptionalBinaryPV, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalBinaryPV"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalBinaryPV")
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
	var _child BACnetOptionalBinaryPV
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalBinaryPVNull
		if _child, err = new(_BACnetOptionalBinaryPVNull).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalBinaryPVNull for type-switch of BACnetOptionalBinaryPV")
		}
	case 0 == 0: // BACnetOptionalBinaryPVValue
		if _child, err = new(_BACnetOptionalBinaryPVValue).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalBinaryPVValue for type-switch of BACnetOptionalBinaryPV")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalBinaryPV"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalBinaryPV")
	}

	return _child, nil
}

func (pm *_BACnetOptionalBinaryPV) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetOptionalBinaryPV, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetOptionalBinaryPV"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalBinaryPV")
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

	if popErr := writeBuffer.PopContext("BACnetOptionalBinaryPV"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalBinaryPV")
	}
	return nil
}

func (m *_BACnetOptionalBinaryPV) IsBACnetOptionalBinaryPV() {}

func (m *_BACnetOptionalBinaryPV) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalBinaryPV) deepCopy() *_BACnetOptionalBinaryPV {
	if m == nil {
		return nil
	}
	_BACnetOptionalBinaryPVCopy := &_BACnetOptionalBinaryPV{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetOptionalBinaryPVCopy
}
