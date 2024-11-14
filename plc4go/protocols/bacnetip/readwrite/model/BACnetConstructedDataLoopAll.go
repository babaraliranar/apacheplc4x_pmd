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

// BACnetConstructedDataLoopAll is the corresponding interface of BACnetConstructedDataLoopAll
type BACnetConstructedDataLoopAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataLoopAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLoopAll()
	// CreateBuilder creates a BACnetConstructedDataLoopAllBuilder
	CreateBACnetConstructedDataLoopAllBuilder() BACnetConstructedDataLoopAllBuilder
}

// _BACnetConstructedDataLoopAll is the data-structure of this message
type _BACnetConstructedDataLoopAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataLoopAll = (*_BACnetConstructedDataLoopAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLoopAll)(nil)

// NewBACnetConstructedDataLoopAll factory function for _BACnetConstructedDataLoopAll
func NewBACnetConstructedDataLoopAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoopAll {
	_result := &_BACnetConstructedDataLoopAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLoopAllBuilder is a builder for BACnetConstructedDataLoopAll
type BACnetConstructedDataLoopAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataLoopAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLoopAll or returns an error if something is wrong
	Build() (BACnetConstructedDataLoopAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLoopAll
}

// NewBACnetConstructedDataLoopAllBuilder() creates a BACnetConstructedDataLoopAllBuilder
func NewBACnetConstructedDataLoopAllBuilder() BACnetConstructedDataLoopAllBuilder {
	return &_BACnetConstructedDataLoopAllBuilder{_BACnetConstructedDataLoopAll: new(_BACnetConstructedDataLoopAll)}
}

type _BACnetConstructedDataLoopAllBuilder struct {
	*_BACnetConstructedDataLoopAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLoopAllBuilder) = (*_BACnetConstructedDataLoopAllBuilder)(nil)

func (b *_BACnetConstructedDataLoopAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLoopAll
}

func (b *_BACnetConstructedDataLoopAllBuilder) WithMandatoryFields() BACnetConstructedDataLoopAllBuilder {
	return b
}

func (b *_BACnetConstructedDataLoopAllBuilder) Build() (BACnetConstructedDataLoopAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLoopAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataLoopAllBuilder) MustBuild() BACnetConstructedDataLoopAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLoopAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLoopAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLoopAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLoopAllBuilder().(*_BACnetConstructedDataLoopAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLoopAllBuilder creates a BACnetConstructedDataLoopAllBuilder
func (b *_BACnetConstructedDataLoopAll) CreateBACnetConstructedDataLoopAllBuilder() BACnetConstructedDataLoopAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataLoopAllBuilder()
	}
	return &_BACnetConstructedDataLoopAllBuilder{_BACnetConstructedDataLoopAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoopAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LOOP
}

func (m *_BACnetConstructedDataLoopAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoopAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoopAll(structType any) BACnetConstructedDataLoopAll {
	if casted, ok := structType.(BACnetConstructedDataLoopAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoopAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoopAll) GetTypeName() string {
	return "BACnetConstructedDataLoopAll"
}

func (m *_BACnetConstructedDataLoopAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataLoopAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLoopAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLoopAll BACnetConstructedDataLoopAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoopAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoopAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoopAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoopAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLoopAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoopAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoopAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoopAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoopAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoopAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoopAll) IsBACnetConstructedDataLoopAll() {}

func (m *_BACnetConstructedDataLoopAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLoopAll) deepCopy() *_BACnetConstructedDataLoopAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLoopAllCopy := &_BACnetConstructedDataLoopAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataLoopAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLoopAllCopy
}

func (m *_BACnetConstructedDataLoopAll) String() string {
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
