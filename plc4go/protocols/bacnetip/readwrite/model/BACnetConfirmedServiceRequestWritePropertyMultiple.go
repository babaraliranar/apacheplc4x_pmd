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

// BACnetConfirmedServiceRequestWritePropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestWritePropertyMultiple
type BACnetConfirmedServiceRequestWritePropertyMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []BACnetWriteAccessSpecification
	// IsBACnetConfirmedServiceRequestWritePropertyMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestWritePropertyMultiple()
	// CreateBuilder creates a BACnetConfirmedServiceRequestWritePropertyMultipleBuilder
	CreateBACnetConfirmedServiceRequestWritePropertyMultipleBuilder() BACnetConfirmedServiceRequestWritePropertyMultipleBuilder
}

// _BACnetConfirmedServiceRequestWritePropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestWritePropertyMultiple struct {
	BACnetConfirmedServiceRequestContract
	Data []BACnetWriteAccessSpecification

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestWritePropertyMultiple = (*_BACnetConfirmedServiceRequestWritePropertyMultiple)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestWritePropertyMultiple)(nil)

// NewBACnetConfirmedServiceRequestWritePropertyMultiple factory function for _BACnetConfirmedServiceRequestWritePropertyMultiple
func NewBACnetConfirmedServiceRequestWritePropertyMultiple(data []BACnetWriteAccessSpecification, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestWritePropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestWritePropertyMultiple{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		Data:                                  data,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestWritePropertyMultipleBuilder is a builder for BACnetConfirmedServiceRequestWritePropertyMultiple
type BACnetConfirmedServiceRequestWritePropertyMultipleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(data []BACnetWriteAccessSpecification) BACnetConfirmedServiceRequestWritePropertyMultipleBuilder
	// WithData adds Data (property field)
	WithData(...BACnetWriteAccessSpecification) BACnetConfirmedServiceRequestWritePropertyMultipleBuilder
	// Build builds the BACnetConfirmedServiceRequestWritePropertyMultiple or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestWritePropertyMultiple, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestWritePropertyMultiple
}

// NewBACnetConfirmedServiceRequestWritePropertyMultipleBuilder() creates a BACnetConfirmedServiceRequestWritePropertyMultipleBuilder
func NewBACnetConfirmedServiceRequestWritePropertyMultipleBuilder() BACnetConfirmedServiceRequestWritePropertyMultipleBuilder {
	return &_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder{_BACnetConfirmedServiceRequestWritePropertyMultiple: new(_BACnetConfirmedServiceRequestWritePropertyMultiple)}
}

type _BACnetConfirmedServiceRequestWritePropertyMultipleBuilder struct {
	*_BACnetConfirmedServiceRequestWritePropertyMultiple

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) = (*_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) WithMandatoryFields(data []BACnetWriteAccessSpecification) BACnetConfirmedServiceRequestWritePropertyMultipleBuilder {
	return b.WithData(data...)
}

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) WithData(data ...BACnetWriteAccessSpecification) BACnetConfirmedServiceRequestWritePropertyMultipleBuilder {
	b.Data = data
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) Build() (BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestWritePropertyMultiple.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) MustBuild() BACnetConfirmedServiceRequestWritePropertyMultiple {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestWritePropertyMultipleBuilder().(*_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestWritePropertyMultipleBuilder creates a BACnetConfirmedServiceRequestWritePropertyMultipleBuilder
func (b *_BACnetConfirmedServiceRequestWritePropertyMultiple) CreateBACnetConfirmedServiceRequestWritePropertyMultipleBuilder() BACnetConfirmedServiceRequestWritePropertyMultipleBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestWritePropertyMultipleBuilder()
	}
	return &_BACnetConfirmedServiceRequestWritePropertyMultipleBuilder{_BACnetConfirmedServiceRequestWritePropertyMultiple: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetData() []BACnetWriteAccessSpecification {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestWritePropertyMultiple(structType any) BACnetConfirmedServiceRequestWritePropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWritePropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestWritePropertyMultiple BACnetConfirmedServiceRequestWritePropertyMultiple, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestWritePropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := ReadLengthArrayField[BACnetWriteAccessSpecification](ctx, "data", ReadComplex[BACnetWriteAccessSpecification](BACnetWriteAccessSpecificationParseWithBuffer, readBuffer), int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestWritePropertyMultiple")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestWritePropertyMultiple")
		}

		if err := WriteComplexTypeArrayField(ctx, "data", m.GetData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestWritePropertyMultiple")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) IsBACnetConfirmedServiceRequestWritePropertyMultiple() {
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) deepCopy() *_BACnetConfirmedServiceRequestWritePropertyMultiple {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestWritePropertyMultipleCopy := &_BACnetConfirmedServiceRequestWritePropertyMultiple{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopySlice[BACnetWriteAccessSpecification, BACnetWriteAccessSpecification](m.Data),
		m.ServiceRequestPayloadLength,
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestWritePropertyMultipleCopy
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
