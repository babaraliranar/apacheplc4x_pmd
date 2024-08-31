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

// BACnetUnconfirmedServiceRequestWhoHasObject is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHasObject
type BACnetUnconfirmedServiceRequestWhoHasObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetUnconfirmedServiceRequestWhoHasObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWhoHasObject.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWhoHasObjectExactly interface {
	BACnetUnconfirmedServiceRequestWhoHasObject
	isBACnetUnconfirmedServiceRequestWhoHasObject() bool
}

// _BACnetUnconfirmedServiceRequestWhoHasObject is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHasObject struct {
	_BACnetUnconfirmedServiceRequestWhoHasObjectChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetUnconfirmedServiceRequestWhoHasObjectChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetUnconfirmedServiceRequestWhoHasObjectParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetUnconfirmedServiceRequestWhoHasObject, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetUnconfirmedServiceRequestWhoHasObjectChild interface {
	utils.Serializable
	InitializeParent(parent BACnetUnconfirmedServiceRequestWhoHasObject, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetUnconfirmedServiceRequestWhoHasObject

	GetTypeName() string
	BACnetUnconfirmedServiceRequestWhoHasObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHasObject factory function for _BACnetUnconfirmedServiceRequestWhoHasObject
func NewBACnetUnconfirmedServiceRequestWhoHasObject(peekedTagHeader BACnetTagHeader) *_BACnetUnconfirmedServiceRequestWhoHasObject {
	return &_BACnetUnconfirmedServiceRequestWhoHasObject{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHasObject(structType any) BACnetUnconfirmedServiceRequestWhoHasObject {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHasObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHasObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHasObject"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestWhoHasObjectParse(ctx context.Context, theBytes []byte) (BACnetUnconfirmedServiceRequestWhoHasObject, error) {
	return BACnetUnconfirmedServiceRequestWhoHasObjectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetUnconfirmedServiceRequestWhoHasObjectParseWithBufferProducer[T BACnetUnconfirmedServiceRequestWhoHasObject]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetUnconfirmedServiceRequestWhoHasObjectParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetUnconfirmedServiceRequestWhoHasObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoHasObject, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHasObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHasObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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
	type BACnetUnconfirmedServiceRequestWhoHasObjectChildSerializeRequirement interface {
		BACnetUnconfirmedServiceRequestWhoHasObject
		InitializeParent(BACnetUnconfirmedServiceRequestWhoHasObject, BACnetTagHeader)
		GetParent() BACnetUnconfirmedServiceRequestWhoHasObject
	}
	var _childTemp any
	var _child BACnetUnconfirmedServiceRequestWhoHasObjectChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(2): // BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(3): // BACnetUnconfirmedServiceRequestWhoHasObjectName
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestWhoHasObjectNameParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetUnconfirmedServiceRequestWhoHasObject")
	}
	_child = _childTemp.(BACnetUnconfirmedServiceRequestWhoHasObjectChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHasObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHasObject")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetUnconfirmedServiceRequestWhoHasObject) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetUnconfirmedServiceRequestWhoHasObject, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHasObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHasObject")
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

	if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHasObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHasObject")
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) isBACnetUnconfirmedServiceRequestWhoHasObject() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
