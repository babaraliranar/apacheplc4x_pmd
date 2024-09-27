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

// RegisterNodesRequest is the corresponding interface of RegisterNodesRequest
type RegisterNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToRegister returns NoOfNodesToRegister (property field)
	GetNoOfNodesToRegister() int32
	// GetNodesToRegister returns NodesToRegister (property field)
	GetNodesToRegister() []NodeId
	// IsRegisterNodesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRegisterNodesRequest()
	// CreateBuilder creates a RegisterNodesRequestBuilder
	CreateRegisterNodesRequestBuilder() RegisterNodesRequestBuilder
}

// _RegisterNodesRequest is the data-structure of this message
type _RegisterNodesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader       ExtensionObjectDefinition
	NoOfNodesToRegister int32
	NodesToRegister     []NodeId
}

var _ RegisterNodesRequest = (*_RegisterNodesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterNodesRequest)(nil)

// NewRegisterNodesRequest factory function for _RegisterNodesRequest
func NewRegisterNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToRegister int32, nodesToRegister []NodeId) *_RegisterNodesRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for RegisterNodesRequest must not be nil")
	}
	_result := &_RegisterNodesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfNodesToRegister:               noOfNodesToRegister,
		NodesToRegister:                   nodesToRegister,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RegisterNodesRequestBuilder is a builder for RegisterNodesRequest
type RegisterNodesRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfNodesToRegister int32, nodesToRegister []NodeId) RegisterNodesRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) RegisterNodesRequestBuilder
	// WithNoOfNodesToRegister adds NoOfNodesToRegister (property field)
	WithNoOfNodesToRegister(int32) RegisterNodesRequestBuilder
	// WithNodesToRegister adds NodesToRegister (property field)
	WithNodesToRegister(...NodeId) RegisterNodesRequestBuilder
	// Build builds the RegisterNodesRequest or returns an error if something is wrong
	Build() (RegisterNodesRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RegisterNodesRequest
}

// NewRegisterNodesRequestBuilder() creates a RegisterNodesRequestBuilder
func NewRegisterNodesRequestBuilder() RegisterNodesRequestBuilder {
	return &_RegisterNodesRequestBuilder{_RegisterNodesRequest: new(_RegisterNodesRequest)}
}

type _RegisterNodesRequestBuilder struct {
	*_RegisterNodesRequest

	err *utils.MultiError
}

var _ (RegisterNodesRequestBuilder) = (*_RegisterNodesRequestBuilder)(nil)

func (m *_RegisterNodesRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfNodesToRegister int32, nodesToRegister []NodeId) RegisterNodesRequestBuilder {
	return m.WithRequestHeader(requestHeader).WithNoOfNodesToRegister(noOfNodesToRegister).WithNodesToRegister(nodesToRegister...)
}

func (m *_RegisterNodesRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) RegisterNodesRequestBuilder {
	m.RequestHeader = requestHeader
	return m
}

func (m *_RegisterNodesRequestBuilder) WithNoOfNodesToRegister(noOfNodesToRegister int32) RegisterNodesRequestBuilder {
	m.NoOfNodesToRegister = noOfNodesToRegister
	return m
}

func (m *_RegisterNodesRequestBuilder) WithNodesToRegister(nodesToRegister ...NodeId) RegisterNodesRequestBuilder {
	m.NodesToRegister = nodesToRegister
	return m
}

func (m *_RegisterNodesRequestBuilder) Build() (RegisterNodesRequest, error) {
	if m.RequestHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._RegisterNodesRequest.deepCopy(), nil
}

func (m *_RegisterNodesRequestBuilder) MustBuild() RegisterNodesRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_RegisterNodesRequestBuilder) DeepCopy() any {
	return m.CreateRegisterNodesRequestBuilder()
}

// CreateRegisterNodesRequestBuilder creates a RegisterNodesRequestBuilder
func (m *_RegisterNodesRequest) CreateRegisterNodesRequestBuilder() RegisterNodesRequestBuilder {
	if m == nil {
		return NewRegisterNodesRequestBuilder()
	}
	return &_RegisterNodesRequestBuilder{_RegisterNodesRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterNodesRequest) GetIdentifier() string {
	return "560"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterNodesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_RegisterNodesRequest) GetNoOfNodesToRegister() int32 {
	return m.NoOfNodesToRegister
}

func (m *_RegisterNodesRequest) GetNodesToRegister() []NodeId {
	return m.NodesToRegister
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRegisterNodesRequest(structType any) RegisterNodesRequest {
	if casted, ok := structType.(RegisterNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterNodesRequest) GetTypeName() string {
	return "RegisterNodesRequest"
}

func (m *_RegisterNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToRegister)
	lengthInBits += 32

	// Array field
	if len(m.NodesToRegister) > 0 {
		for _curItem, element := range m.NodesToRegister {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToRegister), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RegisterNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RegisterNodesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__registerNodesRequest RegisterNodesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToRegister, err := ReadSimpleField(ctx, "noOfNodesToRegister", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToRegister' field"))
	}
	m.NoOfNodesToRegister = noOfNodesToRegister

	nodesToRegister, err := ReadCountArrayField[NodeId](ctx, "nodesToRegister", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer), uint64(noOfNodesToRegister))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToRegister' field"))
	}
	m.NodesToRegister = nodesToRegister

	if closeErr := readBuffer.CloseContext("RegisterNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterNodesRequest")
	}

	return m, nil
}

func (m *_RegisterNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterNodesRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToRegister", m.GetNoOfNodesToRegister(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToRegister' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToRegister", m.GetNodesToRegister(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToRegister' field")
		}

		if popErr := writeBuffer.PopContext("RegisterNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterNodesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterNodesRequest) IsRegisterNodesRequest() {}

func (m *_RegisterNodesRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RegisterNodesRequest) deepCopy() *_RegisterNodesRequest {
	if m == nil {
		return nil
	}
	_RegisterNodesRequestCopy := &_RegisterNodesRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfNodesToRegister,
		utils.DeepCopySlice[NodeId, NodeId](m.NodesToRegister),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RegisterNodesRequestCopy
}

func (m *_RegisterNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
