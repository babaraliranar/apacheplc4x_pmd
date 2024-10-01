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

// CreateSessionRequest is the corresponding interface of CreateSessionRequest
type CreateSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetClientDescription returns ClientDescription (property field)
	GetClientDescription() ExtensionObjectDefinition
	// GetServerUri returns ServerUri (property field)
	GetServerUri() PascalString
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetSessionName returns SessionName (property field)
	GetSessionName() PascalString
	// GetClientNonce returns ClientNonce (property field)
	GetClientNonce() PascalByteString
	// GetClientCertificate returns ClientCertificate (property field)
	GetClientCertificate() PascalByteString
	// GetRequestedSessionTimeout returns RequestedSessionTimeout (property field)
	GetRequestedSessionTimeout() float64
	// GetMaxResponseMessageSize returns MaxResponseMessageSize (property field)
	GetMaxResponseMessageSize() uint32
	// IsCreateSessionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateSessionRequest()
	// CreateBuilder creates a CreateSessionRequestBuilder
	CreateCreateSessionRequestBuilder() CreateSessionRequestBuilder
}

// _CreateSessionRequest is the data-structure of this message
type _CreateSessionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader           ExtensionObjectDefinition
	ClientDescription       ExtensionObjectDefinition
	ServerUri               PascalString
	EndpointUrl             PascalString
	SessionName             PascalString
	ClientNonce             PascalByteString
	ClientCertificate       PascalByteString
	RequestedSessionTimeout float64
	MaxResponseMessageSize  uint32
}

var _ CreateSessionRequest = (*_CreateSessionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateSessionRequest)(nil)

// NewCreateSessionRequest factory function for _CreateSessionRequest
func NewCreateSessionRequest(requestHeader ExtensionObjectDefinition, clientDescription ExtensionObjectDefinition, serverUri PascalString, endpointUrl PascalString, sessionName PascalString, clientNonce PascalByteString, clientCertificate PascalByteString, requestedSessionTimeout float64, maxResponseMessageSize uint32) *_CreateSessionRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for CreateSessionRequest must not be nil")
	}
	if clientDescription == nil {
		panic("clientDescription of type ExtensionObjectDefinition for CreateSessionRequest must not be nil")
	}
	if serverUri == nil {
		panic("serverUri of type PascalString for CreateSessionRequest must not be nil")
	}
	if endpointUrl == nil {
		panic("endpointUrl of type PascalString for CreateSessionRequest must not be nil")
	}
	if sessionName == nil {
		panic("sessionName of type PascalString for CreateSessionRequest must not be nil")
	}
	if clientNonce == nil {
		panic("clientNonce of type PascalByteString for CreateSessionRequest must not be nil")
	}
	if clientCertificate == nil {
		panic("clientCertificate of type PascalByteString for CreateSessionRequest must not be nil")
	}
	_result := &_CreateSessionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ClientDescription:                 clientDescription,
		ServerUri:                         serverUri,
		EndpointUrl:                       endpointUrl,
		SessionName:                       sessionName,
		ClientNonce:                       clientNonce,
		ClientCertificate:                 clientCertificate,
		RequestedSessionTimeout:           requestedSessionTimeout,
		MaxResponseMessageSize:            maxResponseMessageSize,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CreateSessionRequestBuilder is a builder for CreateSessionRequest
type CreateSessionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, clientDescription ExtensionObjectDefinition, serverUri PascalString, endpointUrl PascalString, sessionName PascalString, clientNonce PascalByteString, clientCertificate PascalByteString, requestedSessionTimeout float64, maxResponseMessageSize uint32) CreateSessionRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) CreateSessionRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CreateSessionRequestBuilder
	// WithClientDescription adds ClientDescription (property field)
	WithClientDescription(ExtensionObjectDefinition) CreateSessionRequestBuilder
	// WithClientDescriptionBuilder adds ClientDescription (property field) which is build by the builder
	WithClientDescriptionBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CreateSessionRequestBuilder
	// WithServerUri adds ServerUri (property field)
	WithServerUri(PascalString) CreateSessionRequestBuilder
	// WithServerUriBuilder adds ServerUri (property field) which is build by the builder
	WithServerUriBuilder(func(PascalStringBuilder) PascalStringBuilder) CreateSessionRequestBuilder
	// WithEndpointUrl adds EndpointUrl (property field)
	WithEndpointUrl(PascalString) CreateSessionRequestBuilder
	// WithEndpointUrlBuilder adds EndpointUrl (property field) which is build by the builder
	WithEndpointUrlBuilder(func(PascalStringBuilder) PascalStringBuilder) CreateSessionRequestBuilder
	// WithSessionName adds SessionName (property field)
	WithSessionName(PascalString) CreateSessionRequestBuilder
	// WithSessionNameBuilder adds SessionName (property field) which is build by the builder
	WithSessionNameBuilder(func(PascalStringBuilder) PascalStringBuilder) CreateSessionRequestBuilder
	// WithClientNonce adds ClientNonce (property field)
	WithClientNonce(PascalByteString) CreateSessionRequestBuilder
	// WithClientNonceBuilder adds ClientNonce (property field) which is build by the builder
	WithClientNonceBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) CreateSessionRequestBuilder
	// WithClientCertificate adds ClientCertificate (property field)
	WithClientCertificate(PascalByteString) CreateSessionRequestBuilder
	// WithClientCertificateBuilder adds ClientCertificate (property field) which is build by the builder
	WithClientCertificateBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) CreateSessionRequestBuilder
	// WithRequestedSessionTimeout adds RequestedSessionTimeout (property field)
	WithRequestedSessionTimeout(float64) CreateSessionRequestBuilder
	// WithMaxResponseMessageSize adds MaxResponseMessageSize (property field)
	WithMaxResponseMessageSize(uint32) CreateSessionRequestBuilder
	// Build builds the CreateSessionRequest or returns an error if something is wrong
	Build() (CreateSessionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CreateSessionRequest
}

// NewCreateSessionRequestBuilder() creates a CreateSessionRequestBuilder
func NewCreateSessionRequestBuilder() CreateSessionRequestBuilder {
	return &_CreateSessionRequestBuilder{_CreateSessionRequest: new(_CreateSessionRequest)}
}

type _CreateSessionRequestBuilder struct {
	*_CreateSessionRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CreateSessionRequestBuilder) = (*_CreateSessionRequestBuilder)(nil)

func (b *_CreateSessionRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CreateSessionRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, clientDescription ExtensionObjectDefinition, serverUri PascalString, endpointUrl PascalString, sessionName PascalString, clientNonce PascalByteString, clientCertificate PascalByteString, requestedSessionTimeout float64, maxResponseMessageSize uint32) CreateSessionRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithClientDescription(clientDescription).WithServerUri(serverUri).WithEndpointUrl(endpointUrl).WithSessionName(sessionName).WithClientNonce(clientNonce).WithClientCertificate(clientCertificate).WithRequestedSessionTimeout(requestedSessionTimeout).WithMaxResponseMessageSize(maxResponseMessageSize)
}

func (b *_CreateSessionRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) CreateSessionRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CreateSessionRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithClientDescription(clientDescription ExtensionObjectDefinition) CreateSessionRequestBuilder {
	b.ClientDescription = clientDescription
	return b
}

func (b *_CreateSessionRequestBuilder) WithClientDescriptionBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.ClientDescription.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.ClientDescription, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithServerUri(serverUri PascalString) CreateSessionRequestBuilder {
	b.ServerUri = serverUri
	return b
}

func (b *_CreateSessionRequestBuilder) WithServerUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.ServerUri.CreatePascalStringBuilder())
	var err error
	b.ServerUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithEndpointUrl(endpointUrl PascalString) CreateSessionRequestBuilder {
	b.EndpointUrl = endpointUrl
	return b
}

func (b *_CreateSessionRequestBuilder) WithEndpointUrlBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.EndpointUrl.CreatePascalStringBuilder())
	var err error
	b.EndpointUrl, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithSessionName(sessionName PascalString) CreateSessionRequestBuilder {
	b.SessionName = sessionName
	return b
}

func (b *_CreateSessionRequestBuilder) WithSessionNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.SessionName.CreatePascalStringBuilder())
	var err error
	b.SessionName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithClientNonce(clientNonce PascalByteString) CreateSessionRequestBuilder {
	b.ClientNonce = clientNonce
	return b
}

func (b *_CreateSessionRequestBuilder) WithClientNonceBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.ClientNonce.CreatePascalByteStringBuilder())
	var err error
	b.ClientNonce, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithClientCertificate(clientCertificate PascalByteString) CreateSessionRequestBuilder {
	b.ClientCertificate = clientCertificate
	return b
}

func (b *_CreateSessionRequestBuilder) WithClientCertificateBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) CreateSessionRequestBuilder {
	builder := builderSupplier(b.ClientCertificate.CreatePascalByteStringBuilder())
	var err error
	b.ClientCertificate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_CreateSessionRequestBuilder) WithRequestedSessionTimeout(requestedSessionTimeout float64) CreateSessionRequestBuilder {
	b.RequestedSessionTimeout = requestedSessionTimeout
	return b
}

func (b *_CreateSessionRequestBuilder) WithMaxResponseMessageSize(maxResponseMessageSize uint32) CreateSessionRequestBuilder {
	b.MaxResponseMessageSize = maxResponseMessageSize
	return b
}

func (b *_CreateSessionRequestBuilder) Build() (CreateSessionRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.ClientDescription == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'clientDescription' not set"))
	}
	if b.ServerUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'serverUri' not set"))
	}
	if b.EndpointUrl == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'endpointUrl' not set"))
	}
	if b.SessionName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'sessionName' not set"))
	}
	if b.ClientNonce == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'clientNonce' not set"))
	}
	if b.ClientCertificate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'clientCertificate' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CreateSessionRequest.deepCopy(), nil
}

func (b *_CreateSessionRequestBuilder) MustBuild() CreateSessionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CreateSessionRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CreateSessionRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CreateSessionRequestBuilder) DeepCopy() any {
	_copy := b.CreateCreateSessionRequestBuilder().(*_CreateSessionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCreateSessionRequestBuilder creates a CreateSessionRequestBuilder
func (b *_CreateSessionRequest) CreateCreateSessionRequestBuilder() CreateSessionRequestBuilder {
	if b == nil {
		return NewCreateSessionRequestBuilder()
	}
	return &_CreateSessionRequestBuilder{_CreateSessionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSessionRequest) GetIdentifier() string {
	return "461"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSessionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSessionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateSessionRequest) GetClientDescription() ExtensionObjectDefinition {
	return m.ClientDescription
}

func (m *_CreateSessionRequest) GetServerUri() PascalString {
	return m.ServerUri
}

func (m *_CreateSessionRequest) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_CreateSessionRequest) GetSessionName() PascalString {
	return m.SessionName
}

func (m *_CreateSessionRequest) GetClientNonce() PascalByteString {
	return m.ClientNonce
}

func (m *_CreateSessionRequest) GetClientCertificate() PascalByteString {
	return m.ClientCertificate
}

func (m *_CreateSessionRequest) GetRequestedSessionTimeout() float64 {
	return m.RequestedSessionTimeout
}

func (m *_CreateSessionRequest) GetMaxResponseMessageSize() uint32 {
	return m.MaxResponseMessageSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCreateSessionRequest(structType any) CreateSessionRequest {
	if casted, ok := structType.(CreateSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSessionRequest) GetTypeName() string {
	return "CreateSessionRequest"
}

func (m *_CreateSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientDescription)
	lengthInBits += m.ClientDescription.GetLengthInBits(ctx)

	// Simple field (serverUri)
	lengthInBits += m.ServerUri.GetLengthInBits(ctx)

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (sessionName)
	lengthInBits += m.SessionName.GetLengthInBits(ctx)

	// Simple field (clientNonce)
	lengthInBits += m.ClientNonce.GetLengthInBits(ctx)

	// Simple field (clientCertificate)
	lengthInBits += m.ClientCertificate.GetLengthInBits(ctx)

	// Simple field (requestedSessionTimeout)
	lengthInBits += 64

	// Simple field (maxResponseMessageSize)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CreateSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateSessionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__createSessionRequest CreateSessionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	clientDescription, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "clientDescription", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("310")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientDescription' field"))
	}
	m.ClientDescription = clientDescription

	serverUri, err := ReadSimpleField[PascalString](ctx, "serverUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverUri' field"))
	}
	m.ServerUri = serverUri

	endpointUrl, err := ReadSimpleField[PascalString](ctx, "endpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpointUrl' field"))
	}
	m.EndpointUrl = endpointUrl

	sessionName, err := ReadSimpleField[PascalString](ctx, "sessionName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionName' field"))
	}
	m.SessionName = sessionName

	clientNonce, err := ReadSimpleField[PascalByteString](ctx, "clientNonce", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientNonce' field"))
	}
	m.ClientNonce = clientNonce

	clientCertificate, err := ReadSimpleField[PascalByteString](ctx, "clientCertificate", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientCertificate' field"))
	}
	m.ClientCertificate = clientCertificate

	requestedSessionTimeout, err := ReadSimpleField(ctx, "requestedSessionTimeout", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedSessionTimeout' field"))
	}
	m.RequestedSessionTimeout = requestedSessionTimeout

	maxResponseMessageSize, err := ReadSimpleField(ctx, "maxResponseMessageSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxResponseMessageSize' field"))
	}
	m.MaxResponseMessageSize = maxResponseMessageSize

	if closeErr := readBuffer.CloseContext("CreateSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSessionRequest")
	}

	return m, nil
}

func (m *_CreateSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSessionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "clientDescription", m.GetClientDescription(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientDescription' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "serverUri", m.GetServerUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpointUrl", m.GetEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpointUrl' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "sessionName", m.GetSessionName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionName' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "clientNonce", m.GetClientNonce(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientNonce' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "clientCertificate", m.GetClientCertificate(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientCertificate' field")
		}

		if err := WriteSimpleField[float64](ctx, "requestedSessionTimeout", m.GetRequestedSessionTimeout(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedSessionTimeout' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxResponseMessageSize", m.GetMaxResponseMessageSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxResponseMessageSize' field")
		}

		if popErr := writeBuffer.PopContext("CreateSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSessionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSessionRequest) IsCreateSessionRequest() {}

func (m *_CreateSessionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CreateSessionRequest) deepCopy() *_CreateSessionRequest {
	if m == nil {
		return nil
	}
	_CreateSessionRequestCopy := &_CreateSessionRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.ClientDescription.DeepCopy().(ExtensionObjectDefinition),
		m.ServerUri.DeepCopy().(PascalString),
		m.EndpointUrl.DeepCopy().(PascalString),
		m.SessionName.DeepCopy().(PascalString),
		m.ClientNonce.DeepCopy().(PascalByteString),
		m.ClientCertificate.DeepCopy().(PascalByteString),
		m.RequestedSessionTimeout,
		m.MaxResponseMessageSize,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CreateSessionRequestCopy
}

func (m *_CreateSessionRequest) String() string {
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
