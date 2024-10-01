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

// BrowseResult is the corresponding interface of BrowseResult
type BrowseResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetContinuationPoint returns ContinuationPoint (property field)
	GetContinuationPoint() PascalByteString
	// GetNoOfReferences returns NoOfReferences (property field)
	GetNoOfReferences() int32
	// GetReferences returns References (property field)
	GetReferences() []ExtensionObjectDefinition
	// IsBrowseResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrowseResult()
	// CreateBuilder creates a BrowseResultBuilder
	CreateBrowseResultBuilder() BrowseResultBuilder
}

// _BrowseResult is the data-structure of this message
type _BrowseResult struct {
	ExtensionObjectDefinitionContract
	StatusCode        StatusCode
	ContinuationPoint PascalByteString
	NoOfReferences    int32
	References        []ExtensionObjectDefinition
}

var _ BrowseResult = (*_BrowseResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowseResult)(nil)

// NewBrowseResult factory function for _BrowseResult
func NewBrowseResult(statusCode StatusCode, continuationPoint PascalByteString, noOfReferences int32, references []ExtensionObjectDefinition) *_BrowseResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for BrowseResult must not be nil")
	}
	if continuationPoint == nil {
		panic("continuationPoint of type PascalByteString for BrowseResult must not be nil")
	}
	_result := &_BrowseResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		ContinuationPoint:                 continuationPoint,
		NoOfReferences:                    noOfReferences,
		References:                        references,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrowseResultBuilder is a builder for BrowseResult
type BrowseResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, continuationPoint PascalByteString, noOfReferences int32, references []ExtensionObjectDefinition) BrowseResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) BrowseResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) BrowseResultBuilder
	// WithContinuationPoint adds ContinuationPoint (property field)
	WithContinuationPoint(PascalByteString) BrowseResultBuilder
	// WithContinuationPointBuilder adds ContinuationPoint (property field) which is build by the builder
	WithContinuationPointBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) BrowseResultBuilder
	// WithNoOfReferences adds NoOfReferences (property field)
	WithNoOfReferences(int32) BrowseResultBuilder
	// WithReferences adds References (property field)
	WithReferences(...ExtensionObjectDefinition) BrowseResultBuilder
	// Build builds the BrowseResult or returns an error if something is wrong
	Build() (BrowseResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrowseResult
}

// NewBrowseResultBuilder() creates a BrowseResultBuilder
func NewBrowseResultBuilder() BrowseResultBuilder {
	return &_BrowseResultBuilder{_BrowseResult: new(_BrowseResult)}
}

type _BrowseResultBuilder struct {
	*_BrowseResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrowseResultBuilder) = (*_BrowseResultBuilder)(nil)

func (b *_BrowseResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_BrowseResultBuilder) WithMandatoryFields(statusCode StatusCode, continuationPoint PascalByteString, noOfReferences int32, references []ExtensionObjectDefinition) BrowseResultBuilder {
	return b.WithStatusCode(statusCode).WithContinuationPoint(continuationPoint).WithNoOfReferences(noOfReferences).WithReferences(references...)
}

func (b *_BrowseResultBuilder) WithStatusCode(statusCode StatusCode) BrowseResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_BrowseResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) BrowseResultBuilder {
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

func (b *_BrowseResultBuilder) WithContinuationPoint(continuationPoint PascalByteString) BrowseResultBuilder {
	b.ContinuationPoint = continuationPoint
	return b
}

func (b *_BrowseResultBuilder) WithContinuationPointBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) BrowseResultBuilder {
	builder := builderSupplier(b.ContinuationPoint.CreatePascalByteStringBuilder())
	var err error
	b.ContinuationPoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_BrowseResultBuilder) WithNoOfReferences(noOfReferences int32) BrowseResultBuilder {
	b.NoOfReferences = noOfReferences
	return b
}

func (b *_BrowseResultBuilder) WithReferences(references ...ExtensionObjectDefinition) BrowseResultBuilder {
	b.References = references
	return b
}

func (b *_BrowseResultBuilder) Build() (BrowseResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.ContinuationPoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'continuationPoint' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrowseResult.deepCopy(), nil
}

func (b *_BrowseResultBuilder) MustBuild() BrowseResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BrowseResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_BrowseResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrowseResultBuilder) DeepCopy() any {
	_copy := b.CreateBrowseResultBuilder().(*_BrowseResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrowseResultBuilder creates a BrowseResultBuilder
func (b *_BrowseResult) CreateBrowseResultBuilder() BrowseResultBuilder {
	if b == nil {
		return NewBrowseResultBuilder()
	}
	return &_BrowseResultBuilder{_BrowseResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseResult) GetIdentifier() string {
	return "524"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_BrowseResult) GetContinuationPoint() PascalByteString {
	return m.ContinuationPoint
}

func (m *_BrowseResult) GetNoOfReferences() int32 {
	return m.NoOfReferences
}

func (m *_BrowseResult) GetReferences() []ExtensionObjectDefinition {
	return m.References
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrowseResult(structType any) BrowseResult {
	if casted, ok := structType.(BrowseResult); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseResult); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseResult) GetTypeName() string {
	return "BrowseResult"
}

func (m *_BrowseResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (continuationPoint)
	lengthInBits += m.ContinuationPoint.GetLengthInBits(ctx)

	// Simple field (noOfReferences)
	lengthInBits += 32

	// Array field
	if len(m.References) > 0 {
		for _curItem, element := range m.References {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.References), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowseResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowseResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__browseResult BrowseResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowseResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	continuationPoint, err := ReadSimpleField[PascalByteString](ctx, "continuationPoint", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'continuationPoint' field"))
	}
	m.ContinuationPoint = continuationPoint

	noOfReferences, err := ReadSimpleField(ctx, "noOfReferences", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferences' field"))
	}
	m.NoOfReferences = noOfReferences

	references, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "references", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("520")), readBuffer), uint64(noOfReferences))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'references' field"))
	}
	m.References = references

	if closeErr := readBuffer.CloseContext("BrowseResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseResult")
	}

	return m, nil
}

func (m *_BrowseResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "continuationPoint", m.GetContinuationPoint(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'continuationPoint' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfReferences", m.GetNoOfReferences(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferences' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "references", m.GetReferences(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'references' field")
		}

		if popErr := writeBuffer.PopContext("BrowseResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseResult) IsBrowseResult() {}

func (m *_BrowseResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrowseResult) deepCopy() *_BrowseResult {
	if m == nil {
		return nil
	}
	_BrowseResultCopy := &_BrowseResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.ContinuationPoint.DeepCopy().(PascalByteString),
		m.NoOfReferences,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.References),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrowseResultCopy
}

func (m *_BrowseResult) String() string {
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
