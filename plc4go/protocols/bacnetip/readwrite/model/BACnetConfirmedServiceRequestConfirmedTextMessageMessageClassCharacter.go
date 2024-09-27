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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetCharacterValue returns CharacterValue (property field)
	GetCharacterValue() BACnetContextTagCharacterString
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter struct {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	CharacterValue BACnetContextTagCharacterString
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter)(nil)
var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter)(nil)

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, characterValue BACnetContextTagCharacterString, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if characterValue == nil {
		panic("characterValue of type BACnetContextTagCharacterString for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
		CharacterValue: characterValue,
	}
	_result.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder is a builder for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(characterValue BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// WithCharacterValue adds CharacterValue (property field)
	WithCharacterValue(BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// WithCharacterValueBuilder adds CharacterValue (property field) which is build by the builder
	WithCharacterValueBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
}

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter: new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter)}
}

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) WithMandatoryFields(characterValue BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	return m.WithCharacterValue(characterValue)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) WithCharacterValue(characterValue BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	m.CharacterValue = characterValue
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) WithCharacterValueBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	builder := builderSupplier(m.CharacterValue.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	m.CharacterValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error) {
	if m.CharacterValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'characterValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder()
}

// CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetParent() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract {
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetCharacterValue() BACnetContextTagCharacterString {
	return m.CharacterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).getLengthInBits(ctx))

	// Simple field (characterValue)
	lengthInBits += m.CharacterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, tagNumber uint8) (__bACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, err error) {
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterValue, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "characterValue", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterValue' field"))
	}
	m.CharacterValue = characterValue

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "characterValue", m.GetCharacterValue(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) deepCopy() *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterCopy := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).deepCopy(),
		m.CharacterValue.DeepCopy().(BACnetContextTagCharacterString),
	}
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = m
	return _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
