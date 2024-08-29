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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// UnregisterNodesResponse is the corresponding interface of UnregisterNodesResponse
type UnregisterNodesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
}

// UnregisterNodesResponseExactly can be used when we want exactly this type and not a type which fulfills UnregisterNodesResponse.
// This is useful for switch cases.
type UnregisterNodesResponseExactly interface {
	UnregisterNodesResponse
	isUnregisterNodesResponse() bool
}

// _UnregisterNodesResponse is the data-structure of this message
type _UnregisterNodesResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnregisterNodesResponse) GetIdentifier() string {
	return "569"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnregisterNodesResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_UnregisterNodesResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnregisterNodesResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUnregisterNodesResponse factory function for _UnregisterNodesResponse
func NewUnregisterNodesResponse(responseHeader ExtensionObjectDefinition) *_UnregisterNodesResponse {
	_result := &_UnregisterNodesResponse{
		ResponseHeader:             responseHeader,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUnregisterNodesResponse(structType any) UnregisterNodesResponse {
	if casted, ok := structType.(UnregisterNodesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*UnregisterNodesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_UnregisterNodesResponse) GetTypeName() string {
	return "UnregisterNodesResponse"
}

func (m *_UnregisterNodesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UnregisterNodesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UnregisterNodesResponseParse(ctx context.Context, theBytes []byte, identifier string) (UnregisterNodesResponse, error) {
	return UnregisterNodesResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func UnregisterNodesResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (UnregisterNodesResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("UnregisterNodesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnregisterNodesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of UnregisterNodesResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	if closeErr := readBuffer.CloseContext("UnregisterNodesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnregisterNodesResponse")
	}

	// Create a partially initialized instance
	_child := &_UnregisterNodesResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_UnregisterNodesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnregisterNodesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnregisterNodesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnregisterNodesResponse")
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

		if popErr := writeBuffer.PopContext("UnregisterNodesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnregisterNodesResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UnregisterNodesResponse) isUnregisterNodesResponse() bool {
	return true
}

func (m *_UnregisterNodesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
