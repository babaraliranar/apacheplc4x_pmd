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

// BACnetConstructedDataNotificationForwarderAll is the corresponding interface of BACnetConstructedDataNotificationForwarderAll
type BACnetConstructedDataNotificationForwarderAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataNotificationForwarderAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNotificationForwarderAll()
	// CreateBuilder creates a BACnetConstructedDataNotificationForwarderAllBuilder
	CreateBACnetConstructedDataNotificationForwarderAllBuilder() BACnetConstructedDataNotificationForwarderAllBuilder
}

// _BACnetConstructedDataNotificationForwarderAll is the data-structure of this message
type _BACnetConstructedDataNotificationForwarderAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataNotificationForwarderAll = (*_BACnetConstructedDataNotificationForwarderAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNotificationForwarderAll)(nil)

// NewBACnetConstructedDataNotificationForwarderAll factory function for _BACnetConstructedDataNotificationForwarderAll
func NewBACnetConstructedDataNotificationForwarderAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNotificationForwarderAll {
	_result := &_BACnetConstructedDataNotificationForwarderAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNotificationForwarderAllBuilder is a builder for BACnetConstructedDataNotificationForwarderAll
type BACnetConstructedDataNotificationForwarderAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataNotificationForwarderAllBuilder
	// Build builds the BACnetConstructedDataNotificationForwarderAll or returns an error if something is wrong
	Build() (BACnetConstructedDataNotificationForwarderAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNotificationForwarderAll
}

// NewBACnetConstructedDataNotificationForwarderAllBuilder() creates a BACnetConstructedDataNotificationForwarderAllBuilder
func NewBACnetConstructedDataNotificationForwarderAllBuilder() BACnetConstructedDataNotificationForwarderAllBuilder {
	return &_BACnetConstructedDataNotificationForwarderAllBuilder{_BACnetConstructedDataNotificationForwarderAll: new(_BACnetConstructedDataNotificationForwarderAll)}
}

type _BACnetConstructedDataNotificationForwarderAllBuilder struct {
	*_BACnetConstructedDataNotificationForwarderAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataNotificationForwarderAllBuilder) = (*_BACnetConstructedDataNotificationForwarderAllBuilder)(nil)

func (m *_BACnetConstructedDataNotificationForwarderAllBuilder) WithMandatoryFields() BACnetConstructedDataNotificationForwarderAllBuilder {
	return m
}

func (m *_BACnetConstructedDataNotificationForwarderAllBuilder) Build() (BACnetConstructedDataNotificationForwarderAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataNotificationForwarderAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataNotificationForwarderAllBuilder) MustBuild() BACnetConstructedDataNotificationForwarderAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataNotificationForwarderAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataNotificationForwarderAllBuilder()
}

// CreateBACnetConstructedDataNotificationForwarderAllBuilder creates a BACnetConstructedDataNotificationForwarderAllBuilder
func (m *_BACnetConstructedDataNotificationForwarderAll) CreateBACnetConstructedDataNotificationForwarderAllBuilder() BACnetConstructedDataNotificationForwarderAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataNotificationForwarderAllBuilder()
	}
	return &_BACnetConstructedDataNotificationForwarderAllBuilder{_BACnetConstructedDataNotificationForwarderAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNotificationForwarderAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_NOTIFICATION_FORWARDER
}

func (m *_BACnetConstructedDataNotificationForwarderAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNotificationForwarderAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNotificationForwarderAll(structType any) BACnetConstructedDataNotificationForwarderAll {
	if casted, ok := structType.(BACnetConstructedDataNotificationForwarderAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNotificationForwarderAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNotificationForwarderAll) GetTypeName() string {
	return "BACnetConstructedDataNotificationForwarderAll"
}

func (m *_BACnetConstructedDataNotificationForwarderAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataNotificationForwarderAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNotificationForwarderAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNotificationForwarderAll BACnetConstructedDataNotificationForwarderAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNotificationForwarderAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNotificationForwarderAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNotificationForwarderAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNotificationForwarderAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNotificationForwarderAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNotificationForwarderAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNotificationForwarderAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNotificationForwarderAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNotificationForwarderAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNotificationForwarderAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNotificationForwarderAll) IsBACnetConstructedDataNotificationForwarderAll() {
}

func (m *_BACnetConstructedDataNotificationForwarderAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNotificationForwarderAll) deepCopy() *_BACnetConstructedDataNotificationForwarderAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNotificationForwarderAllCopy := &_BACnetConstructedDataNotificationForwarderAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNotificationForwarderAllCopy
}

func (m *_BACnetConstructedDataNotificationForwarderAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
