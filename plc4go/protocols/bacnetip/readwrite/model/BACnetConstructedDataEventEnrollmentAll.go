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

// BACnetConstructedDataEventEnrollmentAll is the corresponding interface of BACnetConstructedDataEventEnrollmentAll
type BACnetConstructedDataEventEnrollmentAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataEventEnrollmentAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventEnrollmentAll()
	// CreateBuilder creates a BACnetConstructedDataEventEnrollmentAllBuilder
	CreateBACnetConstructedDataEventEnrollmentAllBuilder() BACnetConstructedDataEventEnrollmentAllBuilder
}

// _BACnetConstructedDataEventEnrollmentAll is the data-structure of this message
type _BACnetConstructedDataEventEnrollmentAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataEventEnrollmentAll = (*_BACnetConstructedDataEventEnrollmentAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventEnrollmentAll)(nil)

// NewBACnetConstructedDataEventEnrollmentAll factory function for _BACnetConstructedDataEventEnrollmentAll
func NewBACnetConstructedDataEventEnrollmentAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventEnrollmentAll {
	_result := &_BACnetConstructedDataEventEnrollmentAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventEnrollmentAllBuilder is a builder for BACnetConstructedDataEventEnrollmentAll
type BACnetConstructedDataEventEnrollmentAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataEventEnrollmentAllBuilder
	// Build builds the BACnetConstructedDataEventEnrollmentAll or returns an error if something is wrong
	Build() (BACnetConstructedDataEventEnrollmentAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventEnrollmentAll
}

// NewBACnetConstructedDataEventEnrollmentAllBuilder() creates a BACnetConstructedDataEventEnrollmentAllBuilder
func NewBACnetConstructedDataEventEnrollmentAllBuilder() BACnetConstructedDataEventEnrollmentAllBuilder {
	return &_BACnetConstructedDataEventEnrollmentAllBuilder{_BACnetConstructedDataEventEnrollmentAll: new(_BACnetConstructedDataEventEnrollmentAll)}
}

type _BACnetConstructedDataEventEnrollmentAllBuilder struct {
	*_BACnetConstructedDataEventEnrollmentAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventEnrollmentAllBuilder) = (*_BACnetConstructedDataEventEnrollmentAllBuilder)(nil)

func (m *_BACnetConstructedDataEventEnrollmentAllBuilder) WithMandatoryFields() BACnetConstructedDataEventEnrollmentAllBuilder {
	return m
}

func (m *_BACnetConstructedDataEventEnrollmentAllBuilder) Build() (BACnetConstructedDataEventEnrollmentAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataEventEnrollmentAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataEventEnrollmentAllBuilder) MustBuild() BACnetConstructedDataEventEnrollmentAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataEventEnrollmentAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataEventEnrollmentAllBuilder()
}

// CreateBACnetConstructedDataEventEnrollmentAllBuilder creates a BACnetConstructedDataEventEnrollmentAllBuilder
func (m *_BACnetConstructedDataEventEnrollmentAll) CreateBACnetConstructedDataEventEnrollmentAllBuilder() BACnetConstructedDataEventEnrollmentAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataEventEnrollmentAllBuilder()
	}
	return &_BACnetConstructedDataEventEnrollmentAllBuilder{_BACnetConstructedDataEventEnrollmentAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventEnrollmentAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_EVENT_ENROLLMENT
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventEnrollmentAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventEnrollmentAll(structType any) BACnetConstructedDataEventEnrollmentAll {
	if casted, ok := structType.(BACnetConstructedDataEventEnrollmentAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventEnrollmentAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetTypeName() string {
	return "BACnetConstructedDataEventEnrollmentAll"
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventEnrollmentAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventEnrollmentAll BACnetConstructedDataEventEnrollmentAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventEnrollmentAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventEnrollmentAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventEnrollmentAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventEnrollmentAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventEnrollmentAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventEnrollmentAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventEnrollmentAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventEnrollmentAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventEnrollmentAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventEnrollmentAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventEnrollmentAll) IsBACnetConstructedDataEventEnrollmentAll() {}

func (m *_BACnetConstructedDataEventEnrollmentAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventEnrollmentAll) deepCopy() *_BACnetConstructedDataEventEnrollmentAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventEnrollmentAllCopy := &_BACnetConstructedDataEventEnrollmentAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventEnrollmentAllCopy
}

func (m *_BACnetConstructedDataEventEnrollmentAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
