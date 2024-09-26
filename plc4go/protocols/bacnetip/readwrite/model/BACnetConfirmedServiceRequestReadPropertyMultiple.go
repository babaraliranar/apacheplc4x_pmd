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

// BACnetConfirmedServiceRequestReadPropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestReadPropertyMultiple
type BACnetConfirmedServiceRequestReadPropertyMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []BACnetReadAccessSpecification
	// IsBACnetConfirmedServiceRequestReadPropertyMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadPropertyMultiple()
}

// _BACnetConfirmedServiceRequestReadPropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestReadPropertyMultiple struct {
	BACnetConfirmedServiceRequestContract
	Data []BACnetReadAccessSpecification

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestReadPropertyMultiple = (*_BACnetConfirmedServiceRequestReadPropertyMultiple)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestReadPropertyMultiple)(nil)

// NewBACnetConfirmedServiceRequestReadPropertyMultiple factory function for _BACnetConfirmedServiceRequestReadPropertyMultiple
func NewBACnetConfirmedServiceRequestReadPropertyMultiple(data []BACnetReadAccessSpecification, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReadPropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestReadPropertyMultiple{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		Data:                                  data,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetData() []BACnetReadAccessSpecification {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadPropertyMultiple(structType any) BACnetConfirmedServiceRequestReadPropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadPropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestReadPropertyMultiple BACnetConfirmedServiceRequestReadPropertyMultiple, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := ReadLengthArrayField[BACnetReadAccessSpecification](ctx, "data", ReadComplex[BACnetReadAccessSpecification](BACnetReadAccessSpecificationParseWithBuffer, readBuffer), int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadPropertyMultiple")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadPropertyMultiple")
		}

		if err := WriteComplexTypeArrayField(ctx, "data", m.GetData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadPropertyMultiple")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) IsBACnetConfirmedServiceRequestReadPropertyMultiple() {
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) deepCopy() *_BACnetConfirmedServiceRequestReadPropertyMultiple {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReadPropertyMultipleCopy := &_BACnetConfirmedServiceRequestReadPropertyMultiple{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopySlice[BACnetReadAccessSpecification, BACnetReadAccessSpecification](m.Data),
		m.ServiceRequestPayloadLength,
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestReadPropertyMultipleCopy
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
