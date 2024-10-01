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

// BACnetConstructedDataActiveCOVMultipleSubscriptions is the corresponding interface of BACnetConstructedDataActiveCOVMultipleSubscriptions
type BACnetConstructedDataActiveCOVMultipleSubscriptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetActiveCOVMultipleSubscriptions returns ActiveCOVMultipleSubscriptions (property field)
	GetActiveCOVMultipleSubscriptions() []BACnetCOVMultipleSubscription
	// IsBACnetConstructedDataActiveCOVMultipleSubscriptions is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataActiveCOVMultipleSubscriptions()
	// CreateBuilder creates a BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder
	CreateBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder() BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder
}

// _BACnetConstructedDataActiveCOVMultipleSubscriptions is the data-structure of this message
type _BACnetConstructedDataActiveCOVMultipleSubscriptions struct {
	BACnetConstructedDataContract
	ActiveCOVMultipleSubscriptions []BACnetCOVMultipleSubscription
}

var _ BACnetConstructedDataActiveCOVMultipleSubscriptions = (*_BACnetConstructedDataActiveCOVMultipleSubscriptions)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataActiveCOVMultipleSubscriptions)(nil)

// NewBACnetConstructedDataActiveCOVMultipleSubscriptions factory function for _BACnetConstructedDataActiveCOVMultipleSubscriptions
func NewBACnetConstructedDataActiveCOVMultipleSubscriptions(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, activeCOVMultipleSubscriptions []BACnetCOVMultipleSubscription, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActiveCOVMultipleSubscriptions {
	_result := &_BACnetConstructedDataActiveCOVMultipleSubscriptions{
		BACnetConstructedDataContract:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ActiveCOVMultipleSubscriptions: activeCOVMultipleSubscriptions,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder is a builder for BACnetConstructedDataActiveCOVMultipleSubscriptions
type BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(activeCOVMultipleSubscriptions []BACnetCOVMultipleSubscription) BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder
	// WithActiveCOVMultipleSubscriptions adds ActiveCOVMultipleSubscriptions (property field)
	WithActiveCOVMultipleSubscriptions(...BACnetCOVMultipleSubscription) BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder
	// Build builds the BACnetConstructedDataActiveCOVMultipleSubscriptions or returns an error if something is wrong
	Build() (BACnetConstructedDataActiveCOVMultipleSubscriptions, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataActiveCOVMultipleSubscriptions
}

// NewBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder() creates a BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder
func NewBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder() BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder {
	return &_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder{_BACnetConstructedDataActiveCOVMultipleSubscriptions: new(_BACnetConstructedDataActiveCOVMultipleSubscriptions)}
}

type _BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder struct {
	*_BACnetConstructedDataActiveCOVMultipleSubscriptions

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) = (*_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder)(nil)

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) WithMandatoryFields(activeCOVMultipleSubscriptions []BACnetCOVMultipleSubscription) BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder {
	return b.WithActiveCOVMultipleSubscriptions(activeCOVMultipleSubscriptions...)
}

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) WithActiveCOVMultipleSubscriptions(activeCOVMultipleSubscriptions ...BACnetCOVMultipleSubscription) BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder {
	b.ActiveCOVMultipleSubscriptions = activeCOVMultipleSubscriptions
	return b
}

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) Build() (BACnetConstructedDataActiveCOVMultipleSubscriptions, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataActiveCOVMultipleSubscriptions.deepCopy(), nil
}

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) MustBuild() BACnetConstructedDataActiveCOVMultipleSubscriptions {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder().(*_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder creates a BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder
func (b *_BACnetConstructedDataActiveCOVMultipleSubscriptions) CreateBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder() BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder {
	if b == nil {
		return NewBACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder()
	}
	return &_BACnetConstructedDataActiveCOVMultipleSubscriptionsBuilder{_BACnetConstructedDataActiveCOVMultipleSubscriptions: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_COV_MULTIPLE_SUBSCRIPTIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetActiveCOVMultipleSubscriptions() []BACnetCOVMultipleSubscription {
	return m.ActiveCOVMultipleSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActiveCOVMultipleSubscriptions(structType any) BACnetConstructedDataActiveCOVMultipleSubscriptions {
	if casted, ok := structType.(BACnetConstructedDataActiveCOVMultipleSubscriptions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveCOVMultipleSubscriptions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetTypeName() string {
	return "BACnetConstructedDataActiveCOVMultipleSubscriptions"
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ActiveCOVMultipleSubscriptions) > 0 {
		for _, element := range m.ActiveCOVMultipleSubscriptions {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataActiveCOVMultipleSubscriptions BACnetConstructedDataActiveCOVMultipleSubscriptions, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveCOVMultipleSubscriptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	activeCOVMultipleSubscriptions, err := ReadTerminatedArrayField[BACnetCOVMultipleSubscription](ctx, "activeCOVMultipleSubscriptions", ReadComplex[BACnetCOVMultipleSubscription](BACnetCOVMultipleSubscriptionParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'activeCOVMultipleSubscriptions' field"))
	}
	m.ActiveCOVMultipleSubscriptions = activeCOVMultipleSubscriptions

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveCOVMultipleSubscriptions")
	}

	return m, nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveCOVMultipleSubscriptions")
		}

		if err := WriteComplexTypeArrayField(ctx, "activeCOVMultipleSubscriptions", m.GetActiveCOVMultipleSubscriptions(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'activeCOVMultipleSubscriptions' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveCOVMultipleSubscriptions")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) IsBACnetConstructedDataActiveCOVMultipleSubscriptions() {
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) deepCopy() *_BACnetConstructedDataActiveCOVMultipleSubscriptions {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataActiveCOVMultipleSubscriptionsCopy := &_BACnetConstructedDataActiveCOVMultipleSubscriptions{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetCOVMultipleSubscription, BACnetCOVMultipleSubscription](m.ActiveCOVMultipleSubscriptions),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataActiveCOVMultipleSubscriptionsCopy
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) String() string {
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
