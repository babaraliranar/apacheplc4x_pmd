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

// DeleteReferencesResponse is the corresponding interface of DeleteReferencesResponse
type DeleteReferencesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []StatusCode
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsDeleteReferencesResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteReferencesResponse()
	// CreateBuilder creates a DeleteReferencesResponseBuilder
	CreateDeleteReferencesResponseBuilder() DeleteReferencesResponseBuilder
}

// _DeleteReferencesResponse is the data-structure of this message
type _DeleteReferencesResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []StatusCode
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

var _ DeleteReferencesResponse = (*_DeleteReferencesResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteReferencesResponse)(nil)

// NewDeleteReferencesResponse factory function for _DeleteReferencesResponse
func NewDeleteReferencesResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_DeleteReferencesResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for DeleteReferencesResponse must not be nil")
	}
	_result := &_DeleteReferencesResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfResults:                       noOfResults,
		Results:                           results,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteReferencesResponseBuilder is a builder for DeleteReferencesResponse
type DeleteReferencesResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfResults int32, results []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) DeleteReferencesResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ExtensionObjectDefinition) DeleteReferencesResponseBuilder
	// WithNoOfResults adds NoOfResults (property field)
	WithNoOfResults(int32) DeleteReferencesResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...StatusCode) DeleteReferencesResponseBuilder
	// WithNoOfDiagnosticInfos adds NoOfDiagnosticInfos (property field)
	WithNoOfDiagnosticInfos(int32) DeleteReferencesResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) DeleteReferencesResponseBuilder
	// Build builds the DeleteReferencesResponse or returns an error if something is wrong
	Build() (DeleteReferencesResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteReferencesResponse
}

// NewDeleteReferencesResponseBuilder() creates a DeleteReferencesResponseBuilder
func NewDeleteReferencesResponseBuilder() DeleteReferencesResponseBuilder {
	return &_DeleteReferencesResponseBuilder{_DeleteReferencesResponse: new(_DeleteReferencesResponse)}
}

type _DeleteReferencesResponseBuilder struct {
	*_DeleteReferencesResponse

	err *utils.MultiError
}

var _ (DeleteReferencesResponseBuilder) = (*_DeleteReferencesResponseBuilder)(nil)

func (m *_DeleteReferencesResponseBuilder) WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfResults int32, results []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) DeleteReferencesResponseBuilder {
	return m.WithResponseHeader(responseHeader).WithNoOfResults(noOfResults).WithResults(results...).WithNoOfDiagnosticInfos(noOfDiagnosticInfos).WithDiagnosticInfos(diagnosticInfos...)
}

func (m *_DeleteReferencesResponseBuilder) WithResponseHeader(responseHeader ExtensionObjectDefinition) DeleteReferencesResponseBuilder {
	m.ResponseHeader = responseHeader
	return m
}

func (m *_DeleteReferencesResponseBuilder) WithNoOfResults(noOfResults int32) DeleteReferencesResponseBuilder {
	m.NoOfResults = noOfResults
	return m
}

func (m *_DeleteReferencesResponseBuilder) WithResults(results ...StatusCode) DeleteReferencesResponseBuilder {
	m.Results = results
	return m
}

func (m *_DeleteReferencesResponseBuilder) WithNoOfDiagnosticInfos(noOfDiagnosticInfos int32) DeleteReferencesResponseBuilder {
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos
	return m
}

func (m *_DeleteReferencesResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) DeleteReferencesResponseBuilder {
	m.DiagnosticInfos = diagnosticInfos
	return m
}

func (m *_DeleteReferencesResponseBuilder) Build() (DeleteReferencesResponse, error) {
	if m.ResponseHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._DeleteReferencesResponse.deepCopy(), nil
}

func (m *_DeleteReferencesResponseBuilder) MustBuild() DeleteReferencesResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_DeleteReferencesResponseBuilder) DeepCopy() any {
	return m.CreateDeleteReferencesResponseBuilder()
}

// CreateDeleteReferencesResponseBuilder creates a DeleteReferencesResponseBuilder
func (m *_DeleteReferencesResponse) CreateDeleteReferencesResponseBuilder() DeleteReferencesResponseBuilder {
	if m == nil {
		return NewDeleteReferencesResponseBuilder()
	}
	return &_DeleteReferencesResponseBuilder{_DeleteReferencesResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesResponse) GetIdentifier() string {
	return "509"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_DeleteReferencesResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_DeleteReferencesResponse) GetResults() []StatusCode {
	return m.Results
}

func (m *_DeleteReferencesResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_DeleteReferencesResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteReferencesResponse(structType any) DeleteReferencesResponse {
	if casted, ok := structType.(DeleteReferencesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesResponse) GetTypeName() string {
	return "DeleteReferencesResponse"
}

func (m *_DeleteReferencesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteReferencesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteReferencesResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteReferencesResponse DeleteReferencesResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteReferencesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadSimpleField(ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	m.NoOfResults = noOfResults

	results, err := ReadCountArrayField[StatusCode](ctx, "results", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("DeleteReferencesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesResponse")
	}

	return m, nil
}

func (m *_DeleteReferencesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfResults", m.GetNoOfResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesResponse) IsDeleteReferencesResponse() {}

func (m *_DeleteReferencesResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteReferencesResponse) deepCopy() *_DeleteReferencesResponse {
	if m == nil {
		return nil
	}
	_DeleteReferencesResponseCopy := &_DeleteReferencesResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfResults,
		utils.DeepCopySlice[StatusCode, StatusCode](m.Results),
		m.NoOfDiagnosticInfos,
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteReferencesResponseCopy
}

func (m *_DeleteReferencesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
