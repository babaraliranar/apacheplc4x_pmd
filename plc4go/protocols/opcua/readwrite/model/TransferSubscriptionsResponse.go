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

// TransferSubscriptionsResponse is the corresponding interface of TransferSubscriptionsResponse
type TransferSubscriptionsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetResults returns Results (property field)
	GetResults() []TransferResult
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsTransferSubscriptionsResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTransferSubscriptionsResponse()
	// CreateBuilder creates a TransferSubscriptionsResponseBuilder
	CreateTransferSubscriptionsResponseBuilder() TransferSubscriptionsResponseBuilder
}

// _TransferSubscriptionsResponse is the data-structure of this message
type _TransferSubscriptionsResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader  ResponseHeader
	Results         []TransferResult
	DiagnosticInfos []DiagnosticInfo
}

var _ TransferSubscriptionsResponse = (*_TransferSubscriptionsResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TransferSubscriptionsResponse)(nil)

// NewTransferSubscriptionsResponse factory function for _TransferSubscriptionsResponse
func NewTransferSubscriptionsResponse(responseHeader ResponseHeader, results []TransferResult, diagnosticInfos []DiagnosticInfo) *_TransferSubscriptionsResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for TransferSubscriptionsResponse must not be nil")
	}
	_result := &_TransferSubscriptionsResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		Results:                           results,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TransferSubscriptionsResponseBuilder is a builder for TransferSubscriptionsResponse
type TransferSubscriptionsResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, results []TransferResult, diagnosticInfos []DiagnosticInfo) TransferSubscriptionsResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) TransferSubscriptionsResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) TransferSubscriptionsResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...TransferResult) TransferSubscriptionsResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) TransferSubscriptionsResponseBuilder
	// Build builds the TransferSubscriptionsResponse or returns an error if something is wrong
	Build() (TransferSubscriptionsResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TransferSubscriptionsResponse
}

// NewTransferSubscriptionsResponseBuilder() creates a TransferSubscriptionsResponseBuilder
func NewTransferSubscriptionsResponseBuilder() TransferSubscriptionsResponseBuilder {
	return &_TransferSubscriptionsResponseBuilder{_TransferSubscriptionsResponse: new(_TransferSubscriptionsResponse)}
}

type _TransferSubscriptionsResponseBuilder struct {
	*_TransferSubscriptionsResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TransferSubscriptionsResponseBuilder) = (*_TransferSubscriptionsResponseBuilder)(nil)

func (b *_TransferSubscriptionsResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_TransferSubscriptionsResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, results []TransferResult, diagnosticInfos []DiagnosticInfo) TransferSubscriptionsResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithResults(results...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_TransferSubscriptionsResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) TransferSubscriptionsResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_TransferSubscriptionsResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) TransferSubscriptionsResponseBuilder {
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

func (b *_TransferSubscriptionsResponseBuilder) WithResults(results ...TransferResult) TransferSubscriptionsResponseBuilder {
	b.Results = results
	return b
}

func (b *_TransferSubscriptionsResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) TransferSubscriptionsResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_TransferSubscriptionsResponseBuilder) Build() (TransferSubscriptionsResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TransferSubscriptionsResponse.deepCopy(), nil
}

func (b *_TransferSubscriptionsResponseBuilder) MustBuild() TransferSubscriptionsResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TransferSubscriptionsResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_TransferSubscriptionsResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TransferSubscriptionsResponseBuilder) DeepCopy() any {
	_copy := b.CreateTransferSubscriptionsResponseBuilder().(*_TransferSubscriptionsResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTransferSubscriptionsResponseBuilder creates a TransferSubscriptionsResponseBuilder
func (b *_TransferSubscriptionsResponse) CreateTransferSubscriptionsResponseBuilder() TransferSubscriptionsResponseBuilder {
	if b == nil {
		return NewTransferSubscriptionsResponseBuilder()
	}
	return &_TransferSubscriptionsResponseBuilder{_TransferSubscriptionsResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransferSubscriptionsResponse) GetExtensionId() int32 {
	return int32(844)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransferSubscriptionsResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransferSubscriptionsResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_TransferSubscriptionsResponse) GetResults() []TransferResult {
	return m.Results
}

func (m *_TransferSubscriptionsResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTransferSubscriptionsResponse(structType any) TransferSubscriptionsResponse {
	if casted, ok := structType.(TransferSubscriptionsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*TransferSubscriptionsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_TransferSubscriptionsResponse) GetTypeName() string {
	return "TransferSubscriptionsResponse"
}

func (m *_TransferSubscriptionsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfResults)
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

	// Implicit Field (noOfDiagnosticInfos)
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

func (m *_TransferSubscriptionsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TransferSubscriptionsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__transferSubscriptionsResponse TransferSubscriptionsResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransferSubscriptionsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransferSubscriptionsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadImplicitField[int32](ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	_ = noOfResults

	results, err := ReadCountArrayField[TransferResult](ctx, "results", ReadComplex[TransferResult](ExtensionObjectDefinitionParseWithBufferProducer[TransferResult]((int32)(int32(838))), readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadImplicitField[int32](ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	_ = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("TransferSubscriptionsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransferSubscriptionsResponse")
	}

	return m, nil
}

func (m *_TransferSubscriptionsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransferSubscriptionsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransferSubscriptionsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransferSubscriptionsResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}
		noOfResults := int32(utils.InlineIf(bool((m.GetResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfResults", noOfResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("TransferSubscriptionsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransferSubscriptionsResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransferSubscriptionsResponse) IsTransferSubscriptionsResponse() {}

func (m *_TransferSubscriptionsResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TransferSubscriptionsResponse) deepCopy() *_TransferSubscriptionsResponse {
	if m == nil {
		return nil
	}
	_TransferSubscriptionsResponseCopy := &_TransferSubscriptionsResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ResponseHeader),
		utils.DeepCopySlice[TransferResult, TransferResult](m.Results),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TransferSubscriptionsResponseCopy
}

func (m *_TransferSubscriptionsResponse) String() string {
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
