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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DescriptionResponse is the corresponding interface of DescriptionResponse
type DescriptionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetDibDeviceInfo returns DibDeviceInfo (property field)
	GetDibDeviceInfo() DIBDeviceInfo
	// GetDibSuppSvcFamilies returns DibSuppSvcFamilies (property field)
	GetDibSuppSvcFamilies() DIBSuppSvcFamilies
	// IsDescriptionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDescriptionResponse()
	// CreateBuilder creates a DescriptionResponseBuilder
	CreateDescriptionResponseBuilder() DescriptionResponseBuilder
}

// _DescriptionResponse is the data-structure of this message
type _DescriptionResponse struct {
	KnxNetIpMessageContract
	DibDeviceInfo      DIBDeviceInfo
	DibSuppSvcFamilies DIBSuppSvcFamilies
}

var _ DescriptionResponse = (*_DescriptionResponse)(nil)
var _ KnxNetIpMessageRequirements = (*_DescriptionResponse)(nil)

// NewDescriptionResponse factory function for _DescriptionResponse
func NewDescriptionResponse(dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) *_DescriptionResponse {
	if dibDeviceInfo == nil {
		panic("dibDeviceInfo of type DIBDeviceInfo for DescriptionResponse must not be nil")
	}
	if dibSuppSvcFamilies == nil {
		panic("dibSuppSvcFamilies of type DIBSuppSvcFamilies for DescriptionResponse must not be nil")
	}
	_result := &_DescriptionResponse{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		DibDeviceInfo:           dibDeviceInfo,
		DibSuppSvcFamilies:      dibSuppSvcFamilies,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DescriptionResponseBuilder is a builder for DescriptionResponse
type DescriptionResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) DescriptionResponseBuilder
	// WithDibDeviceInfo adds DibDeviceInfo (property field)
	WithDibDeviceInfo(DIBDeviceInfo) DescriptionResponseBuilder
	// WithDibDeviceInfoBuilder adds DibDeviceInfo (property field) which is build by the builder
	WithDibDeviceInfoBuilder(func(DIBDeviceInfoBuilder) DIBDeviceInfoBuilder) DescriptionResponseBuilder
	// WithDibSuppSvcFamilies adds DibSuppSvcFamilies (property field)
	WithDibSuppSvcFamilies(DIBSuppSvcFamilies) DescriptionResponseBuilder
	// WithDibSuppSvcFamiliesBuilder adds DibSuppSvcFamilies (property field) which is build by the builder
	WithDibSuppSvcFamiliesBuilder(func(DIBSuppSvcFamiliesBuilder) DIBSuppSvcFamiliesBuilder) DescriptionResponseBuilder
	// Build builds the DescriptionResponse or returns an error if something is wrong
	Build() (DescriptionResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DescriptionResponse
}

// NewDescriptionResponseBuilder() creates a DescriptionResponseBuilder
func NewDescriptionResponseBuilder() DescriptionResponseBuilder {
	return &_DescriptionResponseBuilder{_DescriptionResponse: new(_DescriptionResponse)}
}

type _DescriptionResponseBuilder struct {
	*_DescriptionResponse

	err *utils.MultiError
}

var _ (DescriptionResponseBuilder) = (*_DescriptionResponseBuilder)(nil)

func (m *_DescriptionResponseBuilder) WithMandatoryFields(dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) DescriptionResponseBuilder {
	return m.WithDibDeviceInfo(dibDeviceInfo).WithDibSuppSvcFamilies(dibSuppSvcFamilies)
}

func (m *_DescriptionResponseBuilder) WithDibDeviceInfo(dibDeviceInfo DIBDeviceInfo) DescriptionResponseBuilder {
	m.DibDeviceInfo = dibDeviceInfo
	return m
}

func (m *_DescriptionResponseBuilder) WithDibDeviceInfoBuilder(builderSupplier func(DIBDeviceInfoBuilder) DIBDeviceInfoBuilder) DescriptionResponseBuilder {
	builder := builderSupplier(m.DibDeviceInfo.CreateDIBDeviceInfoBuilder())
	var err error
	m.DibDeviceInfo, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "DIBDeviceInfoBuilder failed"))
	}
	return m
}

func (m *_DescriptionResponseBuilder) WithDibSuppSvcFamilies(dibSuppSvcFamilies DIBSuppSvcFamilies) DescriptionResponseBuilder {
	m.DibSuppSvcFamilies = dibSuppSvcFamilies
	return m
}

func (m *_DescriptionResponseBuilder) WithDibSuppSvcFamiliesBuilder(builderSupplier func(DIBSuppSvcFamiliesBuilder) DIBSuppSvcFamiliesBuilder) DescriptionResponseBuilder {
	builder := builderSupplier(m.DibSuppSvcFamilies.CreateDIBSuppSvcFamiliesBuilder())
	var err error
	m.DibSuppSvcFamilies, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "DIBSuppSvcFamiliesBuilder failed"))
	}
	return m
}

func (m *_DescriptionResponseBuilder) Build() (DescriptionResponse, error) {
	if m.DibDeviceInfo == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'dibDeviceInfo' not set"))
	}
	if m.DibSuppSvcFamilies == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'dibSuppSvcFamilies' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._DescriptionResponse.deepCopy(), nil
}

func (m *_DescriptionResponseBuilder) MustBuild() DescriptionResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_DescriptionResponseBuilder) DeepCopy() any {
	return m.CreateDescriptionResponseBuilder()
}

// CreateDescriptionResponseBuilder creates a DescriptionResponseBuilder
func (m *_DescriptionResponse) CreateDescriptionResponseBuilder() DescriptionResponseBuilder {
	if m == nil {
		return NewDescriptionResponseBuilder()
	}
	return &_DescriptionResponseBuilder{_DescriptionResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DescriptionResponse) GetMsgType() uint16 {
	return 0x0204
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DescriptionResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DescriptionResponse) GetDibDeviceInfo() DIBDeviceInfo {
	return m.DibDeviceInfo
}

func (m *_DescriptionResponse) GetDibSuppSvcFamilies() DIBSuppSvcFamilies {
	return m.DibSuppSvcFamilies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDescriptionResponse(structType any) DescriptionResponse {
	if casted, ok := structType.(DescriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DescriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DescriptionResponse) GetTypeName() string {
	return "DescriptionResponse"
}

func (m *_DescriptionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.GetLengthInBits(ctx)

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DescriptionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DescriptionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__descriptionResponse DescriptionResponse, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DescriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DescriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dibDeviceInfo, err := ReadSimpleField[DIBDeviceInfo](ctx, "dibDeviceInfo", ReadComplex[DIBDeviceInfo](DIBDeviceInfoParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dibDeviceInfo' field"))
	}
	m.DibDeviceInfo = dibDeviceInfo

	dibSuppSvcFamilies, err := ReadSimpleField[DIBSuppSvcFamilies](ctx, "dibSuppSvcFamilies", ReadComplex[DIBSuppSvcFamilies](DIBSuppSvcFamiliesParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dibSuppSvcFamilies' field"))
	}
	m.DibSuppSvcFamilies = dibSuppSvcFamilies

	if closeErr := readBuffer.CloseContext("DescriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DescriptionResponse")
	}

	return m, nil
}

func (m *_DescriptionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DescriptionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DescriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DescriptionResponse")
		}

		if err := WriteSimpleField[DIBDeviceInfo](ctx, "dibDeviceInfo", m.GetDibDeviceInfo(), WriteComplex[DIBDeviceInfo](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'dibDeviceInfo' field")
		}

		if err := WriteSimpleField[DIBSuppSvcFamilies](ctx, "dibSuppSvcFamilies", m.GetDibSuppSvcFamilies(), WriteComplex[DIBSuppSvcFamilies](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := writeBuffer.PopContext("DescriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DescriptionResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DescriptionResponse) IsDescriptionResponse() {}

func (m *_DescriptionResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DescriptionResponse) deepCopy() *_DescriptionResponse {
	if m == nil {
		return nil
	}
	_DescriptionResponseCopy := &_DescriptionResponse{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		m.DibDeviceInfo.DeepCopy().(DIBDeviceInfo),
		m.DibSuppSvcFamilies.DeepCopy().(DIBSuppSvcFamilies),
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _DescriptionResponseCopy
}

func (m *_DescriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
