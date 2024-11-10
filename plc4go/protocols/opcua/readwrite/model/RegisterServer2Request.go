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

// RegisterServer2Request is the corresponding interface of RegisterServer2Request
type RegisterServer2Request interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetServer returns Server (property field)
	GetServer() RegisteredServer
	// GetDiscoveryConfiguration returns DiscoveryConfiguration (property field)
	GetDiscoveryConfiguration() []ExtensionObject
	// IsRegisterServer2Request is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRegisterServer2Request()
	// CreateBuilder creates a RegisterServer2RequestBuilder
	CreateRegisterServer2RequestBuilder() RegisterServer2RequestBuilder
}

// _RegisterServer2Request is the data-structure of this message
type _RegisterServer2Request struct {
	ExtensionObjectDefinitionContract
	RequestHeader          RequestHeader
	Server                 RegisteredServer
	DiscoveryConfiguration []ExtensionObject
}

var _ RegisterServer2Request = (*_RegisterServer2Request)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterServer2Request)(nil)

// NewRegisterServer2Request factory function for _RegisterServer2Request
func NewRegisterServer2Request(requestHeader RequestHeader, server RegisteredServer, discoveryConfiguration []ExtensionObject) *_RegisterServer2Request {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for RegisterServer2Request must not be nil")
	}
	if server == nil {
		panic("server of type RegisteredServer for RegisterServer2Request must not be nil")
	}
	_result := &_RegisterServer2Request{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		Server:                            server,
		DiscoveryConfiguration:            discoveryConfiguration,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RegisterServer2RequestBuilder is a builder for RegisterServer2Request
type RegisterServer2RequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, server RegisteredServer, discoveryConfiguration []ExtensionObject) RegisterServer2RequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) RegisterServer2RequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) RegisterServer2RequestBuilder
	// WithServer adds Server (property field)
	WithServer(RegisteredServer) RegisterServer2RequestBuilder
	// WithServerBuilder adds Server (property field) which is build by the builder
	WithServerBuilder(func(RegisteredServerBuilder) RegisteredServerBuilder) RegisterServer2RequestBuilder
	// WithDiscoveryConfiguration adds DiscoveryConfiguration (property field)
	WithDiscoveryConfiguration(...ExtensionObject) RegisterServer2RequestBuilder
	// Build builds the RegisterServer2Request or returns an error if something is wrong
	Build() (RegisterServer2Request, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RegisterServer2Request
}

// NewRegisterServer2RequestBuilder() creates a RegisterServer2RequestBuilder
func NewRegisterServer2RequestBuilder() RegisterServer2RequestBuilder {
	return &_RegisterServer2RequestBuilder{_RegisterServer2Request: new(_RegisterServer2Request)}
}

type _RegisterServer2RequestBuilder struct {
	*_RegisterServer2Request

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RegisterServer2RequestBuilder) = (*_RegisterServer2RequestBuilder)(nil)

func (b *_RegisterServer2RequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RegisterServer2RequestBuilder) WithMandatoryFields(requestHeader RequestHeader, server RegisteredServer, discoveryConfiguration []ExtensionObject) RegisterServer2RequestBuilder {
	return b.WithRequestHeader(requestHeader).WithServer(server).WithDiscoveryConfiguration(discoveryConfiguration...)
}

func (b *_RegisterServer2RequestBuilder) WithRequestHeader(requestHeader RequestHeader) RegisterServer2RequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_RegisterServer2RequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) RegisterServer2RequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_RegisterServer2RequestBuilder) WithServer(server RegisteredServer) RegisterServer2RequestBuilder {
	b.Server = server
	return b
}

func (b *_RegisterServer2RequestBuilder) WithServerBuilder(builderSupplier func(RegisteredServerBuilder) RegisteredServerBuilder) RegisterServer2RequestBuilder {
	builder := builderSupplier(b.Server.CreateRegisteredServerBuilder())
	var err error
	b.Server, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RegisteredServerBuilder failed"))
	}
	return b
}

func (b *_RegisterServer2RequestBuilder) WithDiscoveryConfiguration(discoveryConfiguration ...ExtensionObject) RegisterServer2RequestBuilder {
	b.DiscoveryConfiguration = discoveryConfiguration
	return b
}

func (b *_RegisterServer2RequestBuilder) Build() (RegisterServer2Request, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.Server == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'server' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RegisterServer2Request.deepCopy(), nil
}

func (b *_RegisterServer2RequestBuilder) MustBuild() RegisterServer2Request {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RegisterServer2RequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RegisterServer2RequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RegisterServer2RequestBuilder) DeepCopy() any {
	_copy := b.CreateRegisterServer2RequestBuilder().(*_RegisterServer2RequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRegisterServer2RequestBuilder creates a RegisterServer2RequestBuilder
func (b *_RegisterServer2Request) CreateRegisterServer2RequestBuilder() RegisterServer2RequestBuilder {
	if b == nil {
		return NewRegisterServer2RequestBuilder()
	}
	return &_RegisterServer2RequestBuilder{_RegisterServer2Request: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterServer2Request) GetExtensionId() int32 {
	return int32(12195)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterServer2Request) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterServer2Request) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_RegisterServer2Request) GetServer() RegisteredServer {
	return m.Server
}

func (m *_RegisterServer2Request) GetDiscoveryConfiguration() []ExtensionObject {
	return m.DiscoveryConfiguration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRegisterServer2Request(structType any) RegisterServer2Request {
	if casted, ok := structType.(RegisterServer2Request); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterServer2Request); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterServer2Request) GetTypeName() string {
	return "RegisterServer2Request"
}

func (m *_RegisterServer2Request) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (server)
	lengthInBits += m.Server.GetLengthInBits(ctx)

	// Implicit Field (noOfDiscoveryConfiguration)
	lengthInBits += 32

	// Array field
	if len(m.DiscoveryConfiguration) > 0 {
		for _curItem, element := range m.DiscoveryConfiguration {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiscoveryConfiguration), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RegisterServer2Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RegisterServer2Request) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__registerServer2Request RegisterServer2Request, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterServer2Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterServer2Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	server, err := ReadSimpleField[RegisteredServer](ctx, "server", ReadComplex[RegisteredServer](ExtensionObjectDefinitionParseWithBufferProducer[RegisteredServer]((int32)(int32(434))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'server' field"))
	}
	m.Server = server

	noOfDiscoveryConfiguration, err := ReadImplicitField[int32](ctx, "noOfDiscoveryConfiguration", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiscoveryConfiguration' field"))
	}
	_ = noOfDiscoveryConfiguration

	discoveryConfiguration, err := ReadCountArrayField[ExtensionObject](ctx, "discoveryConfiguration", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer), uint64(noOfDiscoveryConfiguration))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discoveryConfiguration' field"))
	}
	m.DiscoveryConfiguration = discoveryConfiguration

	if closeErr := readBuffer.CloseContext("RegisterServer2Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterServer2Request")
	}

	return m, nil
}

func (m *_RegisterServer2Request) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterServer2Request) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterServer2Request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterServer2Request")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[RegisteredServer](ctx, "server", m.GetServer(), WriteComplex[RegisteredServer](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'server' field")
		}
		noOfDiscoveryConfiguration := int32(utils.InlineIf(bool((m.GetDiscoveryConfiguration()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiscoveryConfiguration()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiscoveryConfiguration", noOfDiscoveryConfiguration, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiscoveryConfiguration' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "discoveryConfiguration", m.GetDiscoveryConfiguration(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'discoveryConfiguration' field")
		}

		if popErr := writeBuffer.PopContext("RegisterServer2Request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterServer2Request")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterServer2Request) IsRegisterServer2Request() {}

func (m *_RegisterServer2Request) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RegisterServer2Request) deepCopy() *_RegisterServer2Request {
	if m == nil {
		return nil
	}
	_RegisterServer2RequestCopy := &_RegisterServer2Request{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(RequestHeader),
		m.Server.DeepCopy().(RegisteredServer),
		utils.DeepCopySlice[ExtensionObject, ExtensionObject](m.DiscoveryConfiguration),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RegisterServer2RequestCopy
}

func (m *_RegisterServer2Request) String() string {
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
