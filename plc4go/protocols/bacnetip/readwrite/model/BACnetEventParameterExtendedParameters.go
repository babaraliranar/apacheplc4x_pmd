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

// BACnetEventParameterExtendedParameters is the corresponding interface of BACnetEventParameterExtendedParameters
type BACnetEventParameterExtendedParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetReference returns Reference (property field)
	GetReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetIsOpeningTag returns IsOpeningTag (virtual field)
	GetIsOpeningTag() bool
	// GetIsClosingTag returns IsClosingTag (virtual field)
	GetIsClosingTag() bool
	// IsBACnetEventParameterExtendedParameters is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterExtendedParameters()
}

// _BACnetEventParameterExtendedParameters is the data-structure of this message
type _BACnetEventParameterExtendedParameters struct {
	OpeningTag           BACnetOpeningTag
	PeekedTagHeader      BACnetTagHeader
	NullValue            BACnetApplicationTagNull
	RealValue            BACnetApplicationTagReal
	UnsignedValue        BACnetApplicationTagUnsignedInteger
	BooleanValue         BACnetApplicationTagBoolean
	IntegerValue         BACnetApplicationTagSignedInteger
	DoubleValue          BACnetApplicationTagDouble
	OctetStringValue     BACnetApplicationTagOctetString
	CharacterStringValue BACnetApplicationTagCharacterString
	BitStringValue       BACnetApplicationTagBitString
	EnumeratedValue      BACnetApplicationTagEnumerated
	DateValue            BACnetApplicationTagDate
	TimeValue            BACnetApplicationTagTime
	ObjectIdentifier     BACnetApplicationTagObjectIdentifier
	Reference            BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterExtendedParameters = (*_BACnetEventParameterExtendedParameters)(nil)

// NewBACnetEventParameterExtendedParameters factory function for _BACnetEventParameterExtendedParameters
func NewBACnetEventParameterExtendedParameters(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull, realValue BACnetApplicationTagReal, unsignedValue BACnetApplicationTagUnsignedInteger, booleanValue BACnetApplicationTagBoolean, integerValue BACnetApplicationTagSignedInteger, doubleValue BACnetApplicationTagDouble, octetStringValue BACnetApplicationTagOctetString, characterStringValue BACnetApplicationTagCharacterString, bitStringValue BACnetApplicationTagBitString, enumeratedValue BACnetApplicationTagEnumerated, dateValue BACnetApplicationTagDate, timeValue BACnetApplicationTagTime, objectIdentifier BACnetApplicationTagObjectIdentifier, reference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterExtendedParameters {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterExtendedParameters must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetEventParameterExtendedParameters must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterExtendedParameters must not be nil")
	}
	return &_BACnetEventParameterExtendedParameters{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, NullValue: nullValue, RealValue: realValue, UnsignedValue: unsignedValue, BooleanValue: booleanValue, IntegerValue: integerValue, DoubleValue: doubleValue, OctetStringValue: octetStringValue, CharacterStringValue: characterStringValue, BitStringValue: bitStringValue, EnumeratedValue: enumeratedValue, DateValue: dateValue, TimeValue: timeValue, ObjectIdentifier: objectIdentifier, Reference: reference, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterExtendedParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterExtendedParameters) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetEventParameterExtendedParameters) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

func (m *_BACnetEventParameterExtendedParameters) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

func (m *_BACnetEventParameterExtendedParameters) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

func (m *_BACnetEventParameterExtendedParameters) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

func (m *_BACnetEventParameterExtendedParameters) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

func (m *_BACnetEventParameterExtendedParameters) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

func (m *_BACnetEventParameterExtendedParameters) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

func (m *_BACnetEventParameterExtendedParameters) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

func (m *_BACnetEventParameterExtendedParameters) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

func (m *_BACnetEventParameterExtendedParameters) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

func (m *_BACnetEventParameterExtendedParameters) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

func (m *_BACnetEventParameterExtendedParameters) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *_BACnetEventParameterExtendedParameters) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetEventParameterExtendedParameters) GetReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.Reference
}

func (m *_BACnetEventParameterExtendedParameters) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetEventParameterExtendedParameters) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	nullValue := m.GetNullValue()
	_ = nullValue
	realValue := m.GetRealValue()
	_ = realValue
	unsignedValue := m.GetUnsignedValue()
	_ = unsignedValue
	booleanValue := m.GetBooleanValue()
	_ = booleanValue
	integerValue := m.GetIntegerValue()
	_ = integerValue
	doubleValue := m.GetDoubleValue()
	_ = doubleValue
	octetStringValue := m.GetOctetStringValue()
	_ = octetStringValue
	characterStringValue := m.GetCharacterStringValue()
	_ = characterStringValue
	bitStringValue := m.GetBitStringValue()
	_ = bitStringValue
	enumeratedValue := m.GetEnumeratedValue()
	_ = enumeratedValue
	dateValue := m.GetDateValue()
	_ = dateValue
	timeValue := m.GetTimeValue()
	_ = timeValue
	objectIdentifier := m.GetObjectIdentifier()
	_ = objectIdentifier
	reference := m.GetReference()
	_ = reference
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetEventParameterExtendedParameters) GetIsOpeningTag() bool {
	ctx := context.Background()
	_ = ctx
	nullValue := m.GetNullValue()
	_ = nullValue
	realValue := m.GetRealValue()
	_ = realValue
	unsignedValue := m.GetUnsignedValue()
	_ = unsignedValue
	booleanValue := m.GetBooleanValue()
	_ = booleanValue
	integerValue := m.GetIntegerValue()
	_ = integerValue
	doubleValue := m.GetDoubleValue()
	_ = doubleValue
	octetStringValue := m.GetOctetStringValue()
	_ = octetStringValue
	characterStringValue := m.GetCharacterStringValue()
	_ = characterStringValue
	bitStringValue := m.GetBitStringValue()
	_ = bitStringValue
	enumeratedValue := m.GetEnumeratedValue()
	_ = enumeratedValue
	dateValue := m.GetDateValue()
	_ = dateValue
	timeValue := m.GetTimeValue()
	_ = timeValue
	objectIdentifier := m.GetObjectIdentifier()
	_ = objectIdentifier
	reference := m.GetReference()
	_ = reference
	return bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x6)))
}

func (m *_BACnetEventParameterExtendedParameters) GetIsClosingTag() bool {
	ctx := context.Background()
	_ = ctx
	nullValue := m.GetNullValue()
	_ = nullValue
	realValue := m.GetRealValue()
	_ = realValue
	unsignedValue := m.GetUnsignedValue()
	_ = unsignedValue
	booleanValue := m.GetBooleanValue()
	_ = booleanValue
	integerValue := m.GetIntegerValue()
	_ = integerValue
	doubleValue := m.GetDoubleValue()
	_ = doubleValue
	octetStringValue := m.GetOctetStringValue()
	_ = octetStringValue
	characterStringValue := m.GetCharacterStringValue()
	_ = characterStringValue
	bitStringValue := m.GetBitStringValue()
	_ = bitStringValue
	enumeratedValue := m.GetEnumeratedValue()
	_ = enumeratedValue
	dateValue := m.GetDateValue()
	_ = dateValue
	timeValue := m.GetTimeValue()
	_ = timeValue
	objectIdentifier := m.GetObjectIdentifier()
	_ = objectIdentifier
	reference := m.GetReference()
	_ = reference
	return bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x7)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterExtendedParameters(structType any) BACnetEventParameterExtendedParameters {
	if casted, ok := structType.(BACnetEventParameterExtendedParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterExtendedParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterExtendedParameters) GetTypeName() string {
	return "BACnetEventParameterExtendedParameters"
}

func (m *_BACnetEventParameterExtendedParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (nullValue)
	if m.NullValue != nil {
		lengthInBits += m.NullValue.GetLengthInBits(ctx)
	}

	// Optional Field (realValue)
	if m.RealValue != nil {
		lengthInBits += m.RealValue.GetLengthInBits(ctx)
	}

	// Optional Field (unsignedValue)
	if m.UnsignedValue != nil {
		lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)
	}

	// Optional Field (booleanValue)
	if m.BooleanValue != nil {
		lengthInBits += m.BooleanValue.GetLengthInBits(ctx)
	}

	// Optional Field (integerValue)
	if m.IntegerValue != nil {
		lengthInBits += m.IntegerValue.GetLengthInBits(ctx)
	}

	// Optional Field (doubleValue)
	if m.DoubleValue != nil {
		lengthInBits += m.DoubleValue.GetLengthInBits(ctx)
	}

	// Optional Field (octetStringValue)
	if m.OctetStringValue != nil {
		lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)
	}

	// Optional Field (characterStringValue)
	if m.CharacterStringValue != nil {
		lengthInBits += m.CharacterStringValue.GetLengthInBits(ctx)
	}

	// Optional Field (bitStringValue)
	if m.BitStringValue != nil {
		lengthInBits += m.BitStringValue.GetLengthInBits(ctx)
	}

	// Optional Field (enumeratedValue)
	if m.EnumeratedValue != nil {
		lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)
	}

	// Optional Field (dateValue)
	if m.DateValue != nil {
		lengthInBits += m.DateValue.GetLengthInBits(ctx)
	}

	// Optional Field (timeValue)
	if m.TimeValue != nil {
		lengthInBits += m.TimeValue.GetLengthInBits(ctx)
	}

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)
	}

	// Optional Field (reference)
	if m.Reference != nil {
		lengthInBits += m.Reference.GetLengthInBits(ctx)
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterExtendedParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterExtendedParametersParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterExtendedParameters, error) {
	return BACnetEventParameterExtendedParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterExtendedParametersParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterExtendedParameters, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterExtendedParameters, error) {
		return BACnetEventParameterExtendedParametersParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterExtendedParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterExtendedParameters, error) {
	v, err := (&_BACnetEventParameterExtendedParameters{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventParameterExtendedParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterExtendedParameters BACnetEventParameterExtendedParameters, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterExtendedParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterExtendedParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	isOpeningTag, err := ReadVirtualField[bool](ctx, "isOpeningTag", (*bool)(nil), bool((peekedTagHeader.GetLengthValueType()) == (0x6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isOpeningTag' field"))
	}
	_ = isOpeningTag

	isClosingTag, err := ReadVirtualField[bool](ctx, "isClosingTag", (*bool)(nil), bool((peekedTagHeader.GetLengthValueType()) == (0x7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isClosingTag' field"))
	}
	_ = isClosingTag

	var nullValue BACnetApplicationTagNull
	_nullValue, err := ReadOptionalField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x0))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	if _nullValue != nil {
		nullValue = *_nullValue
		m.NullValue = nullValue
	}

	var realValue BACnetApplicationTagReal
	_realValue, err := ReadOptionalField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x4))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	if _realValue != nil {
		realValue = *_realValue
		m.RealValue = realValue
	}

	var unsignedValue BACnetApplicationTagUnsignedInteger
	_unsignedValue, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x2))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	if _unsignedValue != nil {
		unsignedValue = *_unsignedValue
		m.UnsignedValue = unsignedValue
	}

	var booleanValue BACnetApplicationTagBoolean
	_booleanValue, err := ReadOptionalField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x1))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	if _booleanValue != nil {
		booleanValue = *_booleanValue
		m.BooleanValue = booleanValue
	}

	var integerValue BACnetApplicationTagSignedInteger
	_integerValue, err := ReadOptionalField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x3))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	if _integerValue != nil {
		integerValue = *_integerValue
		m.IntegerValue = integerValue
	}

	var doubleValue BACnetApplicationTagDouble
	_doubleValue, err := ReadOptionalField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x5))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	if _doubleValue != nil {
		doubleValue = *_doubleValue
		m.DoubleValue = doubleValue
	}

	var octetStringValue BACnetApplicationTagOctetString
	_octetStringValue, err := ReadOptionalField[BACnetApplicationTagOctetString](ctx, "octetStringValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x6))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octetStringValue' field"))
	}
	if _octetStringValue != nil {
		octetStringValue = *_octetStringValue
		m.OctetStringValue = octetStringValue
	}

	var characterStringValue BACnetApplicationTagCharacterString
	_characterStringValue, err := ReadOptionalField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x7))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterStringValue' field"))
	}
	if _characterStringValue != nil {
		characterStringValue = *_characterStringValue
		m.CharacterStringValue = characterStringValue
	}

	var bitStringValue BACnetApplicationTagBitString
	_bitStringValue, err := ReadOptionalField[BACnetApplicationTagBitString](ctx, "bitStringValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x8))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	if _bitStringValue != nil {
		bitStringValue = *_bitStringValue
		m.BitStringValue = bitStringValue
	}

	var enumeratedValue BACnetApplicationTagEnumerated
	_enumeratedValue, err := ReadOptionalField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", ReadComplex[BACnetApplicationTagEnumerated](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagEnumerated](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x9))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumeratedValue' field"))
	}
	if _enumeratedValue != nil {
		enumeratedValue = *_enumeratedValue
		m.EnumeratedValue = enumeratedValue
	}

	var dateValue BACnetApplicationTagDate
	_dateValue, err := ReadOptionalField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer), bool(bool(bool((peekedTagNumber) == (0xA))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	if _dateValue != nil {
		dateValue = *_dateValue
		m.DateValue = dateValue
	}

	var timeValue BACnetApplicationTagTime
	_timeValue, err := ReadOptionalField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer), bool(bool(bool((peekedTagNumber) == (0xB))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	if _timeValue != nil {
		timeValue = *_timeValue
		m.TimeValue = timeValue
	}

	var objectIdentifier BACnetApplicationTagObjectIdentifier
	_objectIdentifier, err := ReadOptionalField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer), bool(bool((peekedTagNumber) == (0xC))) && bool(!(isOpeningTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	if _objectIdentifier != nil {
		objectIdentifier = *_objectIdentifier
		m.ObjectIdentifier = objectIdentifier
	}

	var reference BACnetDeviceObjectPropertyReferenceEnclosed
	_reference, err := ReadOptionalField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "reference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer), bool(isOpeningTag) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reference' field"))
	}
	if _reference != nil {
		reference = *_reference
		m.Reference = reference
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterExtendedParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterExtendedParameters")
	}

	return m, nil
}

func (m *_BACnetEventParameterExtendedParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterExtendedParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterExtendedParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterExtendedParameters")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	isOpeningTag := m.GetIsOpeningTag()
	_ = isOpeningTag
	if _isOpeningTagErr := writeBuffer.WriteVirtual(ctx, "isOpeningTag", m.GetIsOpeningTag()); _isOpeningTagErr != nil {
		return errors.Wrap(_isOpeningTagErr, "Error serializing 'isOpeningTag' field")
	}
	// Virtual field
	isClosingTag := m.GetIsClosingTag()
	_ = isClosingTag
	if _isClosingTagErr := writeBuffer.WriteVirtual(ctx, "isClosingTag", m.GetIsClosingTag()); _isClosingTagErr != nil {
		return errors.Wrap(_isClosingTagErr, "Error serializing 'isClosingTag' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagNull](ctx, "nullValue", GetRef(m.GetNullValue()), WriteComplex[BACnetApplicationTagNull](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'nullValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagReal](ctx, "realValue", GetRef(m.GetRealValue()), WriteComplex[BACnetApplicationTagReal](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'realValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", GetRef(m.GetUnsignedValue()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'unsignedValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagBoolean](ctx, "booleanValue", GetRef(m.GetBooleanValue()), WriteComplex[BACnetApplicationTagBoolean](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'booleanValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagSignedInteger](ctx, "integerValue", GetRef(m.GetIntegerValue()), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'integerValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagDouble](ctx, "doubleValue", GetRef(m.GetDoubleValue()), WriteComplex[BACnetApplicationTagDouble](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'doubleValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagOctetString](ctx, "octetStringValue", GetRef(m.GetOctetStringValue()), WriteComplex[BACnetApplicationTagOctetString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'octetStringValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", GetRef(m.GetCharacterStringValue()), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'characterStringValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagBitString](ctx, "bitStringValue", GetRef(m.GetBitStringValue()), WriteComplex[BACnetApplicationTagBitString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'bitStringValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", GetRef(m.GetEnumeratedValue()), WriteComplex[BACnetApplicationTagEnumerated](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'enumeratedValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagDate](ctx, "dateValue", GetRef(m.GetDateValue()), WriteComplex[BACnetApplicationTagDate](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'dateValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagTime](ctx, "timeValue", GetRef(m.GetTimeValue()), WriteComplex[BACnetApplicationTagTime](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'timeValue' field")
	}

	if err := WriteOptionalField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", GetRef(m.GetObjectIdentifier()), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteOptionalField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "reference", GetRef(m.GetReference()), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'reference' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterExtendedParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterExtendedParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterExtendedParameters) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterExtendedParameters) IsBACnetEventParameterExtendedParameters() {}

func (m *_BACnetEventParameterExtendedParameters) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterExtendedParameters) deepCopy() *_BACnetEventParameterExtendedParameters {
	if m == nil {
		return nil
	}
	_BACnetEventParameterExtendedParametersCopy := &_BACnetEventParameterExtendedParameters{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.NullValue.DeepCopy().(BACnetApplicationTagNull),
		m.RealValue.DeepCopy().(BACnetApplicationTagReal),
		m.UnsignedValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		m.BooleanValue.DeepCopy().(BACnetApplicationTagBoolean),
		m.IntegerValue.DeepCopy().(BACnetApplicationTagSignedInteger),
		m.DoubleValue.DeepCopy().(BACnetApplicationTagDouble),
		m.OctetStringValue.DeepCopy().(BACnetApplicationTagOctetString),
		m.CharacterStringValue.DeepCopy().(BACnetApplicationTagCharacterString),
		m.BitStringValue.DeepCopy().(BACnetApplicationTagBitString),
		m.EnumeratedValue.DeepCopy().(BACnetApplicationTagEnumerated),
		m.DateValue.DeepCopy().(BACnetApplicationTagDate),
		m.TimeValue.DeepCopy().(BACnetApplicationTagTime),
		m.ObjectIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
		m.Reference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetEventParameterExtendedParametersCopy
}

func (m *_BACnetEventParameterExtendedParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
