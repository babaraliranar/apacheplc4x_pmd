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

// StatusResult is the corresponding interface of StatusResult
type StatusResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetDiagnosticInfo returns DiagnosticInfo (property field)
	GetDiagnosticInfo() DiagnosticInfo
	// IsStatusResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusResult()
	// CreateBuilder creates a StatusResultBuilder
	CreateStatusResultBuilder() StatusResultBuilder
}

// _StatusResult is the data-structure of this message
type _StatusResult struct {
	ExtensionObjectDefinitionContract
	StatusCode     StatusCode
	DiagnosticInfo DiagnosticInfo
}

var _ StatusResult = (*_StatusResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_StatusResult)(nil)

// NewStatusResult factory function for _StatusResult
func NewStatusResult(statusCode StatusCode, diagnosticInfo DiagnosticInfo) *_StatusResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for StatusResult must not be nil")
	}
	if diagnosticInfo == nil {
		panic("diagnosticInfo of type DiagnosticInfo for StatusResult must not be nil")
	}
	_result := &_StatusResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		DiagnosticInfo:                    diagnosticInfo,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StatusResultBuilder is a builder for StatusResult
type StatusResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, diagnosticInfo DiagnosticInfo) StatusResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) StatusResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) StatusResultBuilder
	// WithDiagnosticInfo adds DiagnosticInfo (property field)
	WithDiagnosticInfo(DiagnosticInfo) StatusResultBuilder
	// WithDiagnosticInfoBuilder adds DiagnosticInfo (property field) which is build by the builder
	WithDiagnosticInfoBuilder(func(DiagnosticInfoBuilder) DiagnosticInfoBuilder) StatusResultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the StatusResult or returns an error if something is wrong
	Build() (StatusResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StatusResult
}

// NewStatusResultBuilder() creates a StatusResultBuilder
func NewStatusResultBuilder() StatusResultBuilder {
	return &_StatusResultBuilder{_StatusResult: new(_StatusResult)}
}

type _StatusResultBuilder struct {
	*_StatusResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (StatusResultBuilder) = (*_StatusResultBuilder)(nil)

func (b *_StatusResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._StatusResult
}

func (b *_StatusResultBuilder) WithMandatoryFields(statusCode StatusCode, diagnosticInfo DiagnosticInfo) StatusResultBuilder {
	return b.WithStatusCode(statusCode).WithDiagnosticInfo(diagnosticInfo)
}

func (b *_StatusResultBuilder) WithStatusCode(statusCode StatusCode) StatusResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_StatusResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) StatusResultBuilder {
	builder := builderSupplier(b.StatusCode.CreateStatusCodeBuilder())
	var err error
	b.StatusCode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_StatusResultBuilder) WithDiagnosticInfo(diagnosticInfo DiagnosticInfo) StatusResultBuilder {
	b.DiagnosticInfo = diagnosticInfo
	return b
}

func (b *_StatusResultBuilder) WithDiagnosticInfoBuilder(builderSupplier func(DiagnosticInfoBuilder) DiagnosticInfoBuilder) StatusResultBuilder {
	builder := builderSupplier(b.DiagnosticInfo.CreateDiagnosticInfoBuilder())
	var err error
	b.DiagnosticInfo, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DiagnosticInfoBuilder failed"))
	}
	return b
}

func (b *_StatusResultBuilder) Build() (StatusResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.DiagnosticInfo == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'diagnosticInfo' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StatusResult.deepCopy(), nil
}

func (b *_StatusResultBuilder) MustBuild() StatusResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StatusResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_StatusResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_StatusResultBuilder) DeepCopy() any {
	_copy := b.CreateStatusResultBuilder().(*_StatusResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStatusResultBuilder creates a StatusResultBuilder
func (b *_StatusResult) CreateStatusResultBuilder() StatusResultBuilder {
	if b == nil {
		return NewStatusResultBuilder()
	}
	return &_StatusResultBuilder{_StatusResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_StatusResult) GetExtensionId() int32 {
	return int32(301)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_StatusResult) GetDiagnosticInfo() DiagnosticInfo {
	return m.DiagnosticInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStatusResult(structType any) StatusResult {
	if casted, ok := structType.(StatusResult); ok {
		return casted
	}
	if casted, ok := structType.(*StatusResult); ok {
		return *casted
	}
	return nil
}

func (m *_StatusResult) GetTypeName() string {
	return "StatusResult"
}

func (m *_StatusResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (diagnosticInfo)
	lengthInBits += m.DiagnosticInfo.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_StatusResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_StatusResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__statusResult StatusResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	diagnosticInfo, err := ReadSimpleField[DiagnosticInfo](ctx, "diagnosticInfo", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfo' field"))
	}
	m.DiagnosticInfo = diagnosticInfo

	if closeErr := readBuffer.CloseContext("StatusResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusResult")
	}

	return m, nil
}

func (m *_StatusResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[DiagnosticInfo](ctx, "diagnosticInfo", m.GetDiagnosticInfo(), WriteComplex[DiagnosticInfo](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfo' field")
		}

		if popErr := writeBuffer.PopContext("StatusResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusResult) IsStatusResult() {}

func (m *_StatusResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StatusResult) deepCopy() *_StatusResult {
	if m == nil {
		return nil
	}
	_StatusResultCopy := &_StatusResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[StatusCode](m.StatusCode),
		utils.DeepCopy[DiagnosticInfo](m.DiagnosticInfo),
	}
	_StatusResultCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _StatusResultCopy
}

func (m *_StatusResult) String() string {
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
