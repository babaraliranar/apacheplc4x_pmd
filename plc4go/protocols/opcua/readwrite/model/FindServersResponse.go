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
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfServers returns NoOfServers (property field)
	GetNoOfServers() int32
	// GetServers returns Servers (property field)
	GetServers() []ExtensionObjectDefinition
	// IsFindServersResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersResponse()
	// CreateBuilder creates a FindServersResponseBuilder
	CreateFindServersResponseBuilder() FindServersResponseBuilder
}

// _FindServersResponse is the data-structure of this message
type _FindServersResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
	NoOfServers    int32
	Servers        []ExtensionObjectDefinition
}

var _ FindServersResponse = (*_FindServersResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersResponse)(nil)

// NewFindServersResponse factory function for _FindServersResponse
func NewFindServersResponse(responseHeader ExtensionObjectDefinition, noOfServers int32, servers []ExtensionObjectDefinition) *_FindServersResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for FindServersResponse must not be nil")
	}
	_result := &_FindServersResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfServers:                       noOfServers,
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
	WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfServers int32, servers []ExtensionObjectDefinition) FindServersResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ExtensionObjectDefinition) FindServersResponseBuilder
	// WithNoOfServers adds NoOfServers (property field)
	WithNoOfServers(int32) FindServersResponseBuilder
	// WithServers adds Servers (property field)
	WithServers(...ExtensionObjectDefinition) FindServersResponseBuilder
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

	err *utils.MultiError
}

var _ (FindServersResponseBuilder) = (*_FindServersResponseBuilder)(nil)

func (m *_FindServersResponseBuilder) WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfServers int32, servers []ExtensionObjectDefinition) FindServersResponseBuilder {
	return m.WithResponseHeader(responseHeader).WithNoOfServers(noOfServers).WithServers(servers...)
}

func (m *_FindServersResponseBuilder) WithResponseHeader(responseHeader ExtensionObjectDefinition) FindServersResponseBuilder {
	m.ResponseHeader = responseHeader
	return m
}

func (m *_FindServersResponseBuilder) WithNoOfServers(noOfServers int32) FindServersResponseBuilder {
	m.NoOfServers = noOfServers
	return m
}

func (m *_FindServersResponseBuilder) WithServers(servers ...ExtensionObjectDefinition) FindServersResponseBuilder {
	m.Servers = servers
	return m
}

func (m *_FindServersResponseBuilder) Build() (FindServersResponse, error) {
	if m.ResponseHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._FindServersResponse.deepCopy(), nil
}

func (m *_FindServersResponseBuilder) MustBuild() FindServersResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_FindServersResponseBuilder) DeepCopy() any {
	return m.CreateFindServersResponseBuilder()
}

// CreateFindServersResponseBuilder creates a FindServersResponseBuilder
func (m *_FindServersResponse) CreateFindServersResponseBuilder() FindServersResponseBuilder {
	if m == nil {
		return NewFindServersResponseBuilder()
	}
	return &_FindServersResponseBuilder{_FindServersResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersResponse) GetIdentifier() string {
	return "425"
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

func (m *_FindServersResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_FindServersResponse) GetNoOfServers() int32 {
	return m.NoOfServers
}

func (m *_FindServersResponse) GetServers() []ExtensionObjectDefinition {
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

	// Simple field (noOfServers)
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

func (m *_FindServersResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__findServersResponse FindServersResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfServers, err := ReadSimpleField(ctx, "noOfServers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServers' field"))
	}
	m.NoOfServers = noOfServers

	servers, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "servers", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("310")), readBuffer), uint64(noOfServers))
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

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServers", m.GetNoOfServers(), WriteSignedInt(writeBuffer, 32)); err != nil {
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
		m.ResponseHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfServers,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.Servers),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FindServersResponseCopy
}

func (m *_FindServersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
