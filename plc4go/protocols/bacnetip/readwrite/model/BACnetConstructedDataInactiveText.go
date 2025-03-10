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

// BACnetConstructedDataInactiveText is the corresponding interface of BACnetConstructedDataInactiveText
type BACnetConstructedDataInactiveText interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInactiveText returns InactiveText (property field)
	GetInactiveText() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataInactiveText is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataInactiveText()
	// CreateBuilder creates a BACnetConstructedDataInactiveTextBuilder
	CreateBACnetConstructedDataInactiveTextBuilder() BACnetConstructedDataInactiveTextBuilder
}

// _BACnetConstructedDataInactiveText is the data-structure of this message
type _BACnetConstructedDataInactiveText struct {
	BACnetConstructedDataContract
	InactiveText BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataInactiveText = (*_BACnetConstructedDataInactiveText)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataInactiveText)(nil)

// NewBACnetConstructedDataInactiveText factory function for _BACnetConstructedDataInactiveText
func NewBACnetConstructedDataInactiveText(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, inactiveText BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInactiveText {
	if inactiveText == nil {
		panic("inactiveText of type BACnetApplicationTagCharacterString for BACnetConstructedDataInactiveText must not be nil")
	}
	_result := &_BACnetConstructedDataInactiveText{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InactiveText:                  inactiveText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataInactiveTextBuilder is a builder for BACnetConstructedDataInactiveText
type BACnetConstructedDataInactiveTextBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(inactiveText BACnetApplicationTagCharacterString) BACnetConstructedDataInactiveTextBuilder
	// WithInactiveText adds InactiveText (property field)
	WithInactiveText(BACnetApplicationTagCharacterString) BACnetConstructedDataInactiveTextBuilder
	// WithInactiveTextBuilder adds InactiveText (property field) which is build by the builder
	WithInactiveTextBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataInactiveTextBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataInactiveText or returns an error if something is wrong
	Build() (BACnetConstructedDataInactiveText, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataInactiveText
}

// NewBACnetConstructedDataInactiveTextBuilder() creates a BACnetConstructedDataInactiveTextBuilder
func NewBACnetConstructedDataInactiveTextBuilder() BACnetConstructedDataInactiveTextBuilder {
	return &_BACnetConstructedDataInactiveTextBuilder{_BACnetConstructedDataInactiveText: new(_BACnetConstructedDataInactiveText)}
}

type _BACnetConstructedDataInactiveTextBuilder struct {
	*_BACnetConstructedDataInactiveText

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataInactiveTextBuilder) = (*_BACnetConstructedDataInactiveTextBuilder)(nil)

func (b *_BACnetConstructedDataInactiveTextBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataInactiveText
}

func (b *_BACnetConstructedDataInactiveTextBuilder) WithMandatoryFields(inactiveText BACnetApplicationTagCharacterString) BACnetConstructedDataInactiveTextBuilder {
	return b.WithInactiveText(inactiveText)
}

func (b *_BACnetConstructedDataInactiveTextBuilder) WithInactiveText(inactiveText BACnetApplicationTagCharacterString) BACnetConstructedDataInactiveTextBuilder {
	b.InactiveText = inactiveText
	return b
}

func (b *_BACnetConstructedDataInactiveTextBuilder) WithInactiveTextBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataInactiveTextBuilder {
	builder := builderSupplier(b.InactiveText.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.InactiveText, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataInactiveTextBuilder) Build() (BACnetConstructedDataInactiveText, error) {
	if b.InactiveText == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'inactiveText' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataInactiveText.deepCopy(), nil
}

func (b *_BACnetConstructedDataInactiveTextBuilder) MustBuild() BACnetConstructedDataInactiveText {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataInactiveTextBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataInactiveTextBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataInactiveTextBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataInactiveTextBuilder().(*_BACnetConstructedDataInactiveTextBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataInactiveTextBuilder creates a BACnetConstructedDataInactiveTextBuilder
func (b *_BACnetConstructedDataInactiveText) CreateBACnetConstructedDataInactiveTextBuilder() BACnetConstructedDataInactiveTextBuilder {
	if b == nil {
		return NewBACnetConstructedDataInactiveTextBuilder()
	}
	return &_BACnetConstructedDataInactiveTextBuilder{_BACnetConstructedDataInactiveText: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInactiveText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInactiveText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INACTIVE_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInactiveText) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInactiveText) GetInactiveText() BACnetApplicationTagCharacterString {
	return m.InactiveText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInactiveText) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetInactiveText())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInactiveText(structType any) BACnetConstructedDataInactiveText {
	if casted, ok := structType.(BACnetConstructedDataInactiveText); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInactiveText); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInactiveText) GetTypeName() string {
	return "BACnetConstructedDataInactiveText"
}

func (m *_BACnetConstructedDataInactiveText) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (inactiveText)
	lengthInBits += m.InactiveText.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInactiveText) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataInactiveText) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataInactiveText BACnetConstructedDataInactiveText, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInactiveText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInactiveText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	inactiveText, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "inactiveText", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inactiveText' field"))
	}
	m.InactiveText = inactiveText

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), inactiveText)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInactiveText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInactiveText")
	}

	return m, nil
}

func (m *_BACnetConstructedDataInactiveText) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInactiveText) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInactiveText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInactiveText")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "inactiveText", m.GetInactiveText(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'inactiveText' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInactiveText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInactiveText")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInactiveText) IsBACnetConstructedDataInactiveText() {}

func (m *_BACnetConstructedDataInactiveText) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataInactiveText) deepCopy() *_BACnetConstructedDataInactiveText {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataInactiveTextCopy := &_BACnetConstructedDataInactiveText{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.InactiveText),
	}
	_BACnetConstructedDataInactiveTextCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataInactiveTextCopy
}

func (m *_BACnetConstructedDataInactiveText) String() string {
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
