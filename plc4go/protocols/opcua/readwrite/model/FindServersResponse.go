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

// FindServersResponse is the corresponding interface of FindServersResponse
type FindServersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetServers returns Servers (property field)
	GetServers() []ApplicationDescription
	// IsFindServersResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersResponse()
	// CreateBuilder creates a FindServersResponseBuilder
	CreateFindServersResponseBuilder() FindServersResponseBuilder
}

// _FindServersResponse is the data-structure of this message
type _FindServersResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ResponseHeader
	Servers        []ApplicationDescription
}

var _ FindServersResponse = (*_FindServersResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersResponse)(nil)

// NewFindServersResponse factory function for _FindServersResponse
func NewFindServersResponse(responseHeader ResponseHeader, servers []ApplicationDescription) *_FindServersResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for FindServersResponse must not be nil")
	}
	_result := &_FindServersResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		Servers:                           servers,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FindServersResponseBuilder is a builder for FindServersResponse
type FindServersResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, servers []ApplicationDescription) FindServersResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) FindServersResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) FindServersResponseBuilder
	// WithServers adds Servers (property field)
	WithServers(...ApplicationDescription) FindServersResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the FindServersResponse or returns an error if something is wrong
	Build() (FindServersResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FindServersResponse
}

// NewFindServersResponseBuilder() creates a FindServersResponseBuilder
func NewFindServersResponseBuilder() FindServersResponseBuilder {
	return &_FindServersResponseBuilder{_FindServersResponse: new(_FindServersResponse)}
}

type _FindServersResponseBuilder struct {
	*_FindServersResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FindServersResponseBuilder) = (*_FindServersResponseBuilder)(nil)

func (b *_FindServersResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._FindServersResponse
}

func (b *_FindServersResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, servers []ApplicationDescription) FindServersResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithServers(servers...)
}

func (b *_FindServersResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) FindServersResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_FindServersResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) FindServersResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_FindServersResponseBuilder) WithServers(servers ...ApplicationDescription) FindServersResponseBuilder {
	b.Servers = servers
	return b
}

func (b *_FindServersResponseBuilder) Build() (FindServersResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FindServersResponse.deepCopy(), nil
}

func (b *_FindServersResponseBuilder) MustBuild() FindServersResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FindServersResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_FindServersResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FindServersResponseBuilder) DeepCopy() any {
	_copy := b.CreateFindServersResponseBuilder().(*_FindServersResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFindServersResponseBuilder creates a FindServersResponseBuilder
func (b *_FindServersResponse) CreateFindServersResponseBuilder() FindServersResponseBuilder {
	if b == nil {
		return NewFindServersResponseBuilder()
	}
	return &_FindServersResponseBuilder{_FindServersResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersResponse) GetExtensionId() int32 {
	return int32(425)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_FindServersResponse) GetServers() []ApplicationDescription {
	return m.Servers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFindServersResponse(structType any) FindServersResponse {
	if casted, ok := structType.(FindServersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersResponse) GetTypeName() string {
	return "FindServersResponse"
}

func (m *_FindServersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfServers)
	lengthInBits += 32

	// Array field
	if len(m.Servers) > 0 {
		for _curItem, element := range m.Servers {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Servers), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FindServersResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__findServersResponse FindServersResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfServers, err := ReadImplicitField[int32](ctx, "noOfServers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServers' field"))
	}
	_ = noOfServers

	servers, err := ReadCountArrayField[ApplicationDescription](ctx, "servers", ReadComplex[ApplicationDescription](ExtensionObjectDefinitionParseWithBufferProducer[ApplicationDescription]((int32)(int32(310))), readBuffer), uint64(noOfServers))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'servers' field"))
	}
	m.Servers = servers

	if closeErr := readBuffer.CloseContext("FindServersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersResponse")
	}

	return m, nil
}

func (m *_FindServersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}
		noOfServers := int32(utils.InlineIf(bool((m.GetServers()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetServers()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfServers", noOfServers, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServers' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "servers", m.GetServers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'servers' field")
		}

		if popErr := writeBuffer.PopContext("FindServersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersResponse) IsFindServersResponse() {}

func (m *_FindServersResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FindServersResponse) deepCopy() *_FindServersResponse {
	if m == nil {
		return nil
	}
	_FindServersResponseCopy := &_FindServersResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ResponseHeader](m.ResponseHeader),
		utils.DeepCopySlice[ApplicationDescription, ApplicationDescription](m.Servers),
	}
	_FindServersResponseCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FindServersResponseCopy
}

func (m *_FindServersResponse) String() string {
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
