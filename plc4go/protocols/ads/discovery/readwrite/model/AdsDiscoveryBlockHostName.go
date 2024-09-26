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

// AdsDiscoveryBlockHostName is the corresponding interface of AdsDiscoveryBlockHostName
type AdsDiscoveryBlockHostName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetHostName returns HostName (property field)
	GetHostName() AmsString
	// IsAdsDiscoveryBlockHostName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockHostName()
}

// _AdsDiscoveryBlockHostName is the data-structure of this message
type _AdsDiscoveryBlockHostName struct {
	AdsDiscoveryBlockContract
	HostName AmsString
}

var _ AdsDiscoveryBlockHostName = (*_AdsDiscoveryBlockHostName)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockHostName)(nil)

// NewAdsDiscoveryBlockHostName factory function for _AdsDiscoveryBlockHostName
func NewAdsDiscoveryBlockHostName(hostName AmsString) *_AdsDiscoveryBlockHostName {
	if hostName == nil {
		panic("hostName of type AmsString for AdsDiscoveryBlockHostName must not be nil")
	}
	_result := &_AdsDiscoveryBlockHostName{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		HostName:                  hostName,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockHostName) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_HOST_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockHostName) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockHostName) GetHostName() AmsString {
	return m.HostName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockHostName(structType any) AdsDiscoveryBlockHostName {
	if casted, ok := structType.(AdsDiscoveryBlockHostName); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockHostName); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockHostName) GetTypeName() string {
	return "AdsDiscoveryBlockHostName"
}

func (m *_AdsDiscoveryBlockHostName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Simple field (hostName)
	lengthInBits += m.HostName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockHostName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockHostName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockHostName AdsDiscoveryBlockHostName, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockHostName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockHostName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hostName, err := ReadSimpleField[AmsString](ctx, "hostName", ReadComplex[AmsString](AmsStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hostName' field"))
	}
	m.HostName = hostName

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockHostName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockHostName")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockHostName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockHostName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockHostName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockHostName")
		}

		if err := WriteSimpleField[AmsString](ctx, "hostName", m.GetHostName(), WriteComplex[AmsString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hostName' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockHostName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockHostName")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockHostName) IsAdsDiscoveryBlockHostName() {}

func (m *_AdsDiscoveryBlockHostName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockHostName) deepCopy() *_AdsDiscoveryBlockHostName {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockHostNameCopy := &_AdsDiscoveryBlockHostName{
		m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).deepCopy(),
		m.HostName.DeepCopy().(AmsString),
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockHostNameCopy
}

func (m *_AdsDiscoveryBlockHostName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
