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

// BACnetConstructedDataSubscribedRecipients is the corresponding interface of BACnetConstructedDataSubscribedRecipients
type BACnetConstructedDataSubscribedRecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSubscribedRecipients returns SubscribedRecipients (property field)
	GetSubscribedRecipients() []BACnetEventNotificationSubscription
	// IsBACnetConstructedDataSubscribedRecipients is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSubscribedRecipients()
	// CreateBuilder creates a BACnetConstructedDataSubscribedRecipientsBuilder
	CreateBACnetConstructedDataSubscribedRecipientsBuilder() BACnetConstructedDataSubscribedRecipientsBuilder
}

// _BACnetConstructedDataSubscribedRecipients is the data-structure of this message
type _BACnetConstructedDataSubscribedRecipients struct {
	BACnetConstructedDataContract
	SubscribedRecipients []BACnetEventNotificationSubscription
}

var _ BACnetConstructedDataSubscribedRecipients = (*_BACnetConstructedDataSubscribedRecipients)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSubscribedRecipients)(nil)

// NewBACnetConstructedDataSubscribedRecipients factory function for _BACnetConstructedDataSubscribedRecipients
func NewBACnetConstructedDataSubscribedRecipients(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, subscribedRecipients []BACnetEventNotificationSubscription, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubscribedRecipients {
	_result := &_BACnetConstructedDataSubscribedRecipients{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SubscribedRecipients:          subscribedRecipients,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSubscribedRecipientsBuilder is a builder for BACnetConstructedDataSubscribedRecipients
type BACnetConstructedDataSubscribedRecipientsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscribedRecipients []BACnetEventNotificationSubscription) BACnetConstructedDataSubscribedRecipientsBuilder
	// WithSubscribedRecipients adds SubscribedRecipients (property field)
	WithSubscribedRecipients(...BACnetEventNotificationSubscription) BACnetConstructedDataSubscribedRecipientsBuilder
	// Build builds the BACnetConstructedDataSubscribedRecipients or returns an error if something is wrong
	Build() (BACnetConstructedDataSubscribedRecipients, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSubscribedRecipients
}

// NewBACnetConstructedDataSubscribedRecipientsBuilder() creates a BACnetConstructedDataSubscribedRecipientsBuilder
func NewBACnetConstructedDataSubscribedRecipientsBuilder() BACnetConstructedDataSubscribedRecipientsBuilder {
	return &_BACnetConstructedDataSubscribedRecipientsBuilder{_BACnetConstructedDataSubscribedRecipients: new(_BACnetConstructedDataSubscribedRecipients)}
}

type _BACnetConstructedDataSubscribedRecipientsBuilder struct {
	*_BACnetConstructedDataSubscribedRecipients

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSubscribedRecipientsBuilder) = (*_BACnetConstructedDataSubscribedRecipientsBuilder)(nil)

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) WithMandatoryFields(subscribedRecipients []BACnetEventNotificationSubscription) BACnetConstructedDataSubscribedRecipientsBuilder {
	return b.WithSubscribedRecipients(subscribedRecipients...)
}

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) WithSubscribedRecipients(subscribedRecipients ...BACnetEventNotificationSubscription) BACnetConstructedDataSubscribedRecipientsBuilder {
	b.SubscribedRecipients = subscribedRecipients
	return b
}

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) Build() (BACnetConstructedDataSubscribedRecipients, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSubscribedRecipients.deepCopy(), nil
}

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) MustBuild() BACnetConstructedDataSubscribedRecipients {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSubscribedRecipientsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSubscribedRecipientsBuilder().(*_BACnetConstructedDataSubscribedRecipientsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSubscribedRecipientsBuilder creates a BACnetConstructedDataSubscribedRecipientsBuilder
func (b *_BACnetConstructedDataSubscribedRecipients) CreateBACnetConstructedDataSubscribedRecipientsBuilder() BACnetConstructedDataSubscribedRecipientsBuilder {
	if b == nil {
		return NewBACnetConstructedDataSubscribedRecipientsBuilder()
	}
	return &_BACnetConstructedDataSubscribedRecipientsBuilder{_BACnetConstructedDataSubscribedRecipients: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBSCRIBED_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetSubscribedRecipients() []BACnetEventNotificationSubscription {
	return m.SubscribedRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubscribedRecipients(structType any) BACnetConstructedDataSubscribedRecipients {
	if casted, ok := structType.(BACnetConstructedDataSubscribedRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubscribedRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetTypeName() string {
	return "BACnetConstructedDataSubscribedRecipients"
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.SubscribedRecipients) > 0 {
		for _, element := range m.SubscribedRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSubscribedRecipients) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSubscribedRecipients BACnetConstructedDataSubscribedRecipients, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubscribedRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubscribedRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscribedRecipients, err := ReadTerminatedArrayField[BACnetEventNotificationSubscription](ctx, "subscribedRecipients", ReadComplex[BACnetEventNotificationSubscription](BACnetEventNotificationSubscriptionParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscribedRecipients' field"))
	}
	m.SubscribedRecipients = subscribedRecipients

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubscribedRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubscribedRecipients")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubscribedRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubscribedRecipients")
		}

		if err := WriteComplexTypeArrayField(ctx, "subscribedRecipients", m.GetSubscribedRecipients(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'subscribedRecipients' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubscribedRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubscribedRecipients")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubscribedRecipients) IsBACnetConstructedDataSubscribedRecipients() {}

func (m *_BACnetConstructedDataSubscribedRecipients) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSubscribedRecipients) deepCopy() *_BACnetConstructedDataSubscribedRecipients {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSubscribedRecipientsCopy := &_BACnetConstructedDataSubscribedRecipients{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetEventNotificationSubscription, BACnetEventNotificationSubscription](m.SubscribedRecipients),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSubscribedRecipientsCopy
}

func (m *_BACnetConstructedDataSubscribedRecipients) String() string {
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
