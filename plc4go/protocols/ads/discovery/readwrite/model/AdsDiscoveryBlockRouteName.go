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

// AdsDiscoveryBlockRouteName is the corresponding interface of AdsDiscoveryBlockRouteName
type AdsDiscoveryBlockRouteName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetRouteName returns RouteName (property field)
	GetRouteName() AmsString
	// IsAdsDiscoveryBlockRouteName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockRouteName()
}

// _AdsDiscoveryBlockRouteName is the data-structure of this message
type _AdsDiscoveryBlockRouteName struct {
	AdsDiscoveryBlockContract
	RouteName AmsString
}

var _ AdsDiscoveryBlockRouteName = (*_AdsDiscoveryBlockRouteName)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockRouteName)(nil)

// NewAdsDiscoveryBlockRouteName factory function for _AdsDiscoveryBlockRouteName
func NewAdsDiscoveryBlockRouteName(routeName AmsString) *_AdsDiscoveryBlockRouteName {
	if routeName == nil {
		panic("routeName of type AmsString for AdsDiscoveryBlockRouteName must not be nil")
	}
	_result := &_AdsDiscoveryBlockRouteName{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		RouteName:                 routeName,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockRouteName) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_ROUTE_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockRouteName) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockRouteName) GetRouteName() AmsString {
	return m.RouteName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockRouteName(structType any) AdsDiscoveryBlockRouteName {
	if casted, ok := structType.(AdsDiscoveryBlockRouteName); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockRouteName); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockRouteName) GetTypeName() string {
	return "AdsDiscoveryBlockRouteName"
}

func (m *_AdsDiscoveryBlockRouteName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Simple field (routeName)
	lengthInBits += m.RouteName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockRouteName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockRouteName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockRouteName AdsDiscoveryBlockRouteName, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockRouteName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockRouteName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	routeName, err := ReadSimpleField[AmsString](ctx, "routeName", ReadComplex[AmsString](AmsStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'routeName' field"))
	}
	m.RouteName = routeName

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockRouteName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockRouteName")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockRouteName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockRouteName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockRouteName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockRouteName")
		}

		if err := WriteSimpleField[AmsString](ctx, "routeName", m.GetRouteName(), WriteComplex[AmsString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'routeName' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockRouteName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockRouteName")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockRouteName) IsAdsDiscoveryBlockRouteName() {}

func (m *_AdsDiscoveryBlockRouteName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockRouteName) deepCopy() *_AdsDiscoveryBlockRouteName {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockRouteNameCopy := &_AdsDiscoveryBlockRouteName{
		m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).deepCopy(),
		m.RouteName.DeepCopy().(AmsString),
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockRouteNameCopy
}

func (m *_AdsDiscoveryBlockRouteName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
