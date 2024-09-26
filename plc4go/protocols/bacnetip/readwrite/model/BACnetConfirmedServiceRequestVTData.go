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

// BACnetConfirmedServiceRequestVTData is the corresponding interface of BACnetConfirmedServiceRequestVTData
type BACnetConfirmedServiceRequestVTData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetVtSessionIdentifier returns VtSessionIdentifier (property field)
	GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
	// GetVtNewData returns VtNewData (property field)
	GetVtNewData() BACnetApplicationTagOctetString
	// GetVtDataFlag returns VtDataFlag (property field)
	GetVtDataFlag() BACnetApplicationTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestVTData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestVTData()
}

// _BACnetConfirmedServiceRequestVTData is the data-structure of this message
type _BACnetConfirmedServiceRequestVTData struct {
	BACnetConfirmedServiceRequestContract
	VtSessionIdentifier BACnetApplicationTagUnsignedInteger
	VtNewData           BACnetApplicationTagOctetString
	VtDataFlag          BACnetApplicationTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestVTData = (*_BACnetConfirmedServiceRequestVTData)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestVTData)(nil)

// NewBACnetConfirmedServiceRequestVTData factory function for _BACnetConfirmedServiceRequestVTData
func NewBACnetConfirmedServiceRequestVTData(vtSessionIdentifier BACnetApplicationTagUnsignedInteger, vtNewData BACnetApplicationTagOctetString, vtDataFlag BACnetApplicationTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestVTData {
	if vtSessionIdentifier == nil {
		panic("vtSessionIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetConfirmedServiceRequestVTData must not be nil")
	}
	if vtNewData == nil {
		panic("vtNewData of type BACnetApplicationTagOctetString for BACnetConfirmedServiceRequestVTData must not be nil")
	}
	if vtDataFlag == nil {
		panic("vtDataFlag of type BACnetApplicationTagUnsignedInteger for BACnetConfirmedServiceRequestVTData must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestVTData{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		VtSessionIdentifier:                   vtSessionIdentifier,
		VtNewData:                             vtNewData,
		VtDataFlag:                            vtDataFlag,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTData) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_DATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTData) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTData) GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.VtSessionIdentifier
}

func (m *_BACnetConfirmedServiceRequestVTData) GetVtNewData() BACnetApplicationTagOctetString {
	return m.VtNewData
}

func (m *_BACnetConfirmedServiceRequestVTData) GetVtDataFlag() BACnetApplicationTagUnsignedInteger {
	return m.VtDataFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTData(structType any) BACnetConfirmedServiceRequestVTData {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTData) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTData"
}

func (m *_BACnetConfirmedServiceRequestVTData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (vtSessionIdentifier)
	lengthInBits += m.VtSessionIdentifier.GetLengthInBits(ctx)

	// Simple field (vtNewData)
	lengthInBits += m.VtNewData.GetLengthInBits(ctx)

	// Simple field (vtDataFlag)
	lengthInBits += m.VtDataFlag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestVTData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestVTData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestVTData BACnetConfirmedServiceRequestVTData, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vtSessionIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtSessionIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtSessionIdentifier' field"))
	}
	m.VtSessionIdentifier = vtSessionIdentifier

	vtNewData, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "vtNewData", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtNewData' field"))
	}
	m.VtNewData = vtNewData

	vtDataFlag, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtDataFlag", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtDataFlag' field"))
	}
	m.VtDataFlag = vtDataFlag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTData")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestVTData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTData")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtSessionIdentifier", m.GetVtSessionIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtSessionIdentifier' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "vtNewData", m.GetVtNewData(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtNewData' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtDataFlag", m.GetVtDataFlag(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtDataFlag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTData")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestVTData) IsBACnetConfirmedServiceRequestVTData() {}

func (m *_BACnetConfirmedServiceRequestVTData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestVTData) deepCopy() *_BACnetConfirmedServiceRequestVTData {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestVTDataCopy := &_BACnetConfirmedServiceRequestVTData{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.VtSessionIdentifier.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		m.VtNewData.DeepCopy().(BACnetApplicationTagOctetString),
		m.VtDataFlag.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestVTDataCopy
}

func (m *_BACnetConfirmedServiceRequestVTData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
