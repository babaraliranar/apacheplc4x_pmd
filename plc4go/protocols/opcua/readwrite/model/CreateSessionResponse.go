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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// CreateSessionResponse is the corresponding interface of CreateSessionResponse
type CreateSessionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetSessionId returns SessionId (property field)
	GetSessionId() NodeId
	// GetAuthenticationToken returns AuthenticationToken (property field)
	GetAuthenticationToken() NodeId
	// GetRevisedSessionTimeout returns RevisedSessionTimeout (property field)
	GetRevisedSessionTimeout() float64
	// GetServerNonce returns ServerNonce (property field)
	GetServerNonce() PascalByteString
	// GetServerCertificate returns ServerCertificate (property field)
	GetServerCertificate() PascalByteString
	// GetNoOfServerEndpoints returns NoOfServerEndpoints (property field)
	GetNoOfServerEndpoints() int32
	// GetServerEndpoints returns ServerEndpoints (property field)
	GetServerEndpoints() []ExtensionObjectDefinition
	// GetNoOfServerSoftwareCertificates returns NoOfServerSoftwareCertificates (property field)
	GetNoOfServerSoftwareCertificates() int32
	// GetServerSoftwareCertificates returns ServerSoftwareCertificates (property field)
	GetServerSoftwareCertificates() []ExtensionObjectDefinition
	// GetServerSignature returns ServerSignature (property field)
	GetServerSignature() ExtensionObjectDefinition
	// GetMaxRequestMessageSize returns MaxRequestMessageSize (property field)
	GetMaxRequestMessageSize() uint32
}

// CreateSessionResponseExactly can be used when we want exactly this type and not a type which fulfills CreateSessionResponse.
// This is useful for switch cases.
type CreateSessionResponseExactly interface {
	CreateSessionResponse
	isCreateSessionResponse() bool
}

// _CreateSessionResponse is the data-structure of this message
type _CreateSessionResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader                 ExtensionObjectDefinition
	SessionId                      NodeId
	AuthenticationToken            NodeId
	RevisedSessionTimeout          float64
	ServerNonce                    PascalByteString
	ServerCertificate              PascalByteString
	NoOfServerEndpoints            int32
	ServerEndpoints                []ExtensionObjectDefinition
	NoOfServerSoftwareCertificates int32
	ServerSoftwareCertificates     []ExtensionObjectDefinition
	ServerSignature                ExtensionObjectDefinition
	MaxRequestMessageSize          uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSessionResponse) GetIdentifier() string {
	return "464"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSessionResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CreateSessionResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSessionResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_CreateSessionResponse) GetSessionId() NodeId {
	return m.SessionId
}

func (m *_CreateSessionResponse) GetAuthenticationToken() NodeId {
	return m.AuthenticationToken
}

func (m *_CreateSessionResponse) GetRevisedSessionTimeout() float64 {
	return m.RevisedSessionTimeout
}

func (m *_CreateSessionResponse) GetServerNonce() PascalByteString {
	return m.ServerNonce
}

func (m *_CreateSessionResponse) GetServerCertificate() PascalByteString {
	return m.ServerCertificate
}

func (m *_CreateSessionResponse) GetNoOfServerEndpoints() int32 {
	return m.NoOfServerEndpoints
}

func (m *_CreateSessionResponse) GetServerEndpoints() []ExtensionObjectDefinition {
	return m.ServerEndpoints
}

func (m *_CreateSessionResponse) GetNoOfServerSoftwareCertificates() int32 {
	return m.NoOfServerSoftwareCertificates
}

func (m *_CreateSessionResponse) GetServerSoftwareCertificates() []ExtensionObjectDefinition {
	return m.ServerSoftwareCertificates
}

func (m *_CreateSessionResponse) GetServerSignature() ExtensionObjectDefinition {
	return m.ServerSignature
}

func (m *_CreateSessionResponse) GetMaxRequestMessageSize() uint32 {
	return m.MaxRequestMessageSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateSessionResponse factory function for _CreateSessionResponse
func NewCreateSessionResponse(responseHeader ExtensionObjectDefinition, sessionId NodeId, authenticationToken NodeId, revisedSessionTimeout float64, serverNonce PascalByteString, serverCertificate PascalByteString, noOfServerEndpoints int32, serverEndpoints []ExtensionObjectDefinition, noOfServerSoftwareCertificates int32, serverSoftwareCertificates []ExtensionObjectDefinition, serverSignature ExtensionObjectDefinition, maxRequestMessageSize uint32) *_CreateSessionResponse {
	_result := &_CreateSessionResponse{
		ResponseHeader:                 responseHeader,
		SessionId:                      sessionId,
		AuthenticationToken:            authenticationToken,
		RevisedSessionTimeout:          revisedSessionTimeout,
		ServerNonce:                    serverNonce,
		ServerCertificate:              serverCertificate,
		NoOfServerEndpoints:            noOfServerEndpoints,
		ServerEndpoints:                serverEndpoints,
		NoOfServerSoftwareCertificates: noOfServerSoftwareCertificates,
		ServerSoftwareCertificates:     serverSoftwareCertificates,
		ServerSignature:                serverSignature,
		MaxRequestMessageSize:          maxRequestMessageSize,
		_ExtensionObjectDefinition:     NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateSessionResponse(structType any) CreateSessionResponse {
	if casted, ok := structType.(CreateSessionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSessionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSessionResponse) GetTypeName() string {
	return "CreateSessionResponse"
}

func (m *_CreateSessionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (sessionId)
	lengthInBits += m.SessionId.GetLengthInBits(ctx)

	// Simple field (authenticationToken)
	lengthInBits += m.AuthenticationToken.GetLengthInBits(ctx)

	// Simple field (revisedSessionTimeout)
	lengthInBits += 64

	// Simple field (serverNonce)
	lengthInBits += m.ServerNonce.GetLengthInBits(ctx)

	// Simple field (serverCertificate)
	lengthInBits += m.ServerCertificate.GetLengthInBits(ctx)

	// Simple field (noOfServerEndpoints)
	lengthInBits += 32

	// Array field
	if len(m.ServerEndpoints) > 0 {
		for _curItem, element := range m.ServerEndpoints {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerEndpoints), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfServerSoftwareCertificates)
	lengthInBits += 32

	// Array field
	if len(m.ServerSoftwareCertificates) > 0 {
		for _curItem, element := range m.ServerSoftwareCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerSoftwareCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (serverSignature)
	lengthInBits += m.ServerSignature.GetLengthInBits(ctx)

	// Simple field (maxRequestMessageSize)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CreateSessionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CreateSessionResponseParse(ctx context.Context, theBytes []byte, identifier string) (CreateSessionResponse, error) {
	return CreateSessionResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CreateSessionResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CreateSessionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CreateSessionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSessionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of CreateSessionResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (sessionId)
	if pullErr := readBuffer.PullContext("sessionId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sessionId")
	}
	_sessionId, _sessionIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _sessionIdErr != nil {
		return nil, errors.Wrap(_sessionIdErr, "Error parsing 'sessionId' field of CreateSessionResponse")
	}
	sessionId := _sessionId.(NodeId)
	if closeErr := readBuffer.CloseContext("sessionId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sessionId")
	}

	// Simple Field (authenticationToken)
	if pullErr := readBuffer.PullContext("authenticationToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationToken")
	}
	_authenticationToken, _authenticationTokenErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _authenticationTokenErr != nil {
		return nil, errors.Wrap(_authenticationTokenErr, "Error parsing 'authenticationToken' field of CreateSessionResponse")
	}
	authenticationToken := _authenticationToken.(NodeId)
	if closeErr := readBuffer.CloseContext("authenticationToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationToken")
	}

	// Simple Field (revisedSessionTimeout)
	_revisedSessionTimeout, _revisedSessionTimeoutErr := readBuffer.ReadFloat64("revisedSessionTimeout", 64)
	if _revisedSessionTimeoutErr != nil {
		return nil, errors.Wrap(_revisedSessionTimeoutErr, "Error parsing 'revisedSessionTimeout' field of CreateSessionResponse")
	}
	revisedSessionTimeout := _revisedSessionTimeout

	// Simple Field (serverNonce)
	if pullErr := readBuffer.PullContext("serverNonce"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverNonce")
	}
	_serverNonce, _serverNonceErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _serverNonceErr != nil {
		return nil, errors.Wrap(_serverNonceErr, "Error parsing 'serverNonce' field of CreateSessionResponse")
	}
	serverNonce := _serverNonce.(PascalByteString)
	if closeErr := readBuffer.CloseContext("serverNonce"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverNonce")
	}

	// Simple Field (serverCertificate)
	if pullErr := readBuffer.PullContext("serverCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverCertificate")
	}
	_serverCertificate, _serverCertificateErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _serverCertificateErr != nil {
		return nil, errors.Wrap(_serverCertificateErr, "Error parsing 'serverCertificate' field of CreateSessionResponse")
	}
	serverCertificate := _serverCertificate.(PascalByteString)
	if closeErr := readBuffer.CloseContext("serverCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverCertificate")
	}

	// Simple Field (noOfServerEndpoints)
	_noOfServerEndpoints, _noOfServerEndpointsErr := readBuffer.ReadInt32("noOfServerEndpoints", 32)
	if _noOfServerEndpointsErr != nil {
		return nil, errors.Wrap(_noOfServerEndpointsErr, "Error parsing 'noOfServerEndpoints' field of CreateSessionResponse")
	}
	noOfServerEndpoints := _noOfServerEndpoints

	// Array field (serverEndpoints)
	if pullErr := readBuffer.PullContext("serverEndpoints", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverEndpoints")
	}
	// Count array
	serverEndpoints := make([]ExtensionObjectDefinition, utils.Max(noOfServerEndpoints, 0))
	// This happens when the size is set conditional to 0
	if len(serverEndpoints) == 0 {
		serverEndpoints = nil
	}
	{
		_numItems := uint16(utils.Max(noOfServerEndpoints, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "314")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'serverEndpoints' field of CreateSessionResponse")
			}
			serverEndpoints[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("serverEndpoints", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverEndpoints")
	}

	// Simple Field (noOfServerSoftwareCertificates)
	_noOfServerSoftwareCertificates, _noOfServerSoftwareCertificatesErr := readBuffer.ReadInt32("noOfServerSoftwareCertificates", 32)
	if _noOfServerSoftwareCertificatesErr != nil {
		return nil, errors.Wrap(_noOfServerSoftwareCertificatesErr, "Error parsing 'noOfServerSoftwareCertificates' field of CreateSessionResponse")
	}
	noOfServerSoftwareCertificates := _noOfServerSoftwareCertificates

	// Array field (serverSoftwareCertificates)
	if pullErr := readBuffer.PullContext("serverSoftwareCertificates", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverSoftwareCertificates")
	}
	// Count array
	serverSoftwareCertificates := make([]ExtensionObjectDefinition, utils.Max(noOfServerSoftwareCertificates, 0))
	// This happens when the size is set conditional to 0
	if len(serverSoftwareCertificates) == 0 {
		serverSoftwareCertificates = nil
	}
	{
		_numItems := uint16(utils.Max(noOfServerSoftwareCertificates, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "346")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'serverSoftwareCertificates' field of CreateSessionResponse")
			}
			serverSoftwareCertificates[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("serverSoftwareCertificates", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverSoftwareCertificates")
	}

	// Simple Field (serverSignature)
	if pullErr := readBuffer.PullContext("serverSignature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverSignature")
	}
	_serverSignature, _serverSignatureErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("458"))
	if _serverSignatureErr != nil {
		return nil, errors.Wrap(_serverSignatureErr, "Error parsing 'serverSignature' field of CreateSessionResponse")
	}
	serverSignature := _serverSignature.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("serverSignature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverSignature")
	}

	// Simple Field (maxRequestMessageSize)
	_maxRequestMessageSize, _maxRequestMessageSizeErr := readBuffer.ReadUint32("maxRequestMessageSize", 32)
	if _maxRequestMessageSizeErr != nil {
		return nil, errors.Wrap(_maxRequestMessageSizeErr, "Error parsing 'maxRequestMessageSize' field of CreateSessionResponse")
	}
	maxRequestMessageSize := _maxRequestMessageSize

	if closeErr := readBuffer.CloseContext("CreateSessionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSessionResponse")
	}

	// Create a partially initialized instance
	_child := &_CreateSessionResponse{
		_ExtensionObjectDefinition:     &_ExtensionObjectDefinition{},
		ResponseHeader:                 responseHeader,
		SessionId:                      sessionId,
		AuthenticationToken:            authenticationToken,
		RevisedSessionTimeout:          revisedSessionTimeout,
		ServerNonce:                    serverNonce,
		ServerCertificate:              serverCertificate,
		NoOfServerEndpoints:            noOfServerEndpoints,
		ServerEndpoints:                serverEndpoints,
		NoOfServerSoftwareCertificates: noOfServerSoftwareCertificates,
		ServerSoftwareCertificates:     serverSoftwareCertificates,
		ServerSignature:                serverSignature,
		MaxRequestMessageSize:          maxRequestMessageSize,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CreateSessionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSessionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSessionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSessionResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (sessionId)
		if pushErr := writeBuffer.PushContext("sessionId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sessionId")
		}
		_sessionIdErr := writeBuffer.WriteSerializable(ctx, m.GetSessionId())
		if popErr := writeBuffer.PopContext("sessionId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sessionId")
		}
		if _sessionIdErr != nil {
			return errors.Wrap(_sessionIdErr, "Error serializing 'sessionId' field")
		}

		// Simple Field (authenticationToken)
		if pushErr := writeBuffer.PushContext("authenticationToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authenticationToken")
		}
		_authenticationTokenErr := writeBuffer.WriteSerializable(ctx, m.GetAuthenticationToken())
		if popErr := writeBuffer.PopContext("authenticationToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authenticationToken")
		}
		if _authenticationTokenErr != nil {
			return errors.Wrap(_authenticationTokenErr, "Error serializing 'authenticationToken' field")
		}

		// Simple Field (revisedSessionTimeout)
		revisedSessionTimeout := float64(m.GetRevisedSessionTimeout())
		_revisedSessionTimeoutErr := writeBuffer.WriteFloat64("revisedSessionTimeout", 64, (revisedSessionTimeout))
		if _revisedSessionTimeoutErr != nil {
			return errors.Wrap(_revisedSessionTimeoutErr, "Error serializing 'revisedSessionTimeout' field")
		}

		// Simple Field (serverNonce)
		if pushErr := writeBuffer.PushContext("serverNonce"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverNonce")
		}
		_serverNonceErr := writeBuffer.WriteSerializable(ctx, m.GetServerNonce())
		if popErr := writeBuffer.PopContext("serverNonce"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverNonce")
		}
		if _serverNonceErr != nil {
			return errors.Wrap(_serverNonceErr, "Error serializing 'serverNonce' field")
		}

		// Simple Field (serverCertificate)
		if pushErr := writeBuffer.PushContext("serverCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverCertificate")
		}
		_serverCertificateErr := writeBuffer.WriteSerializable(ctx, m.GetServerCertificate())
		if popErr := writeBuffer.PopContext("serverCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverCertificate")
		}
		if _serverCertificateErr != nil {
			return errors.Wrap(_serverCertificateErr, "Error serializing 'serverCertificate' field")
		}

		// Simple Field (noOfServerEndpoints)
		noOfServerEndpoints := int32(m.GetNoOfServerEndpoints())
		_noOfServerEndpointsErr := writeBuffer.WriteInt32("noOfServerEndpoints", 32, int32((noOfServerEndpoints)))
		if _noOfServerEndpointsErr != nil {
			return errors.Wrap(_noOfServerEndpointsErr, "Error serializing 'noOfServerEndpoints' field")
		}

		// Array Field (serverEndpoints)
		if pushErr := writeBuffer.PushContext("serverEndpoints", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverEndpoints")
		}
		for _curItem, _element := range m.GetServerEndpoints() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServerEndpoints()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'serverEndpoints' field")
			}
		}
		if popErr := writeBuffer.PopContext("serverEndpoints", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverEndpoints")
		}

		// Simple Field (noOfServerSoftwareCertificates)
		noOfServerSoftwareCertificates := int32(m.GetNoOfServerSoftwareCertificates())
		_noOfServerSoftwareCertificatesErr := writeBuffer.WriteInt32("noOfServerSoftwareCertificates", 32, int32((noOfServerSoftwareCertificates)))
		if _noOfServerSoftwareCertificatesErr != nil {
			return errors.Wrap(_noOfServerSoftwareCertificatesErr, "Error serializing 'noOfServerSoftwareCertificates' field")
		}

		// Array Field (serverSoftwareCertificates)
		if pushErr := writeBuffer.PushContext("serverSoftwareCertificates", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverSoftwareCertificates")
		}
		for _curItem, _element := range m.GetServerSoftwareCertificates() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServerSoftwareCertificates()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'serverSoftwareCertificates' field")
			}
		}
		if popErr := writeBuffer.PopContext("serverSoftwareCertificates", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverSoftwareCertificates")
		}

		// Simple Field (serverSignature)
		if pushErr := writeBuffer.PushContext("serverSignature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverSignature")
		}
		_serverSignatureErr := writeBuffer.WriteSerializable(ctx, m.GetServerSignature())
		if popErr := writeBuffer.PopContext("serverSignature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverSignature")
		}
		if _serverSignatureErr != nil {
			return errors.Wrap(_serverSignatureErr, "Error serializing 'serverSignature' field")
		}

		// Simple Field (maxRequestMessageSize)
		maxRequestMessageSize := uint32(m.GetMaxRequestMessageSize())
		_maxRequestMessageSizeErr := writeBuffer.WriteUint32("maxRequestMessageSize", 32, uint32((maxRequestMessageSize)))
		if _maxRequestMessageSizeErr != nil {
			return errors.Wrap(_maxRequestMessageSizeErr, "Error serializing 'maxRequestMessageSize' field")
		}

		if popErr := writeBuffer.PopContext("CreateSessionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSessionResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSessionResponse) isCreateSessionResponse() bool {
	return true
}

func (m *_CreateSessionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
