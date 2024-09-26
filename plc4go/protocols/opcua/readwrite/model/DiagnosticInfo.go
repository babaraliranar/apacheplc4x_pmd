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

// DiagnosticInfo is the corresponding interface of DiagnosticInfo
type DiagnosticInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetInnerDiagnosticInfoSpecified returns InnerDiagnosticInfoSpecified (property field)
	GetInnerDiagnosticInfoSpecified() bool
	// GetInnerStatusCodeSpecified returns InnerStatusCodeSpecified (property field)
	GetInnerStatusCodeSpecified() bool
	// GetAdditionalInfoSpecified returns AdditionalInfoSpecified (property field)
	GetAdditionalInfoSpecified() bool
	// GetLocaleSpecified returns LocaleSpecified (property field)
	GetLocaleSpecified() bool
	// GetLocalizedTextSpecified returns LocalizedTextSpecified (property field)
	GetLocalizedTextSpecified() bool
	// GetNamespaceURISpecified returns NamespaceURISpecified (property field)
	GetNamespaceURISpecified() bool
	// GetSymbolicIdSpecified returns SymbolicIdSpecified (property field)
	GetSymbolicIdSpecified() bool
	// GetSymbolicId returns SymbolicId (property field)
	GetSymbolicId() *int32
	// GetNamespaceURI returns NamespaceURI (property field)
	GetNamespaceURI() *int32
	// GetLocale returns Locale (property field)
	GetLocale() *int32
	// GetLocalizedText returns LocalizedText (property field)
	GetLocalizedText() *int32
	// GetAdditionalInfo returns AdditionalInfo (property field)
	GetAdditionalInfo() PascalString
	// GetInnerStatusCode returns InnerStatusCode (property field)
	GetInnerStatusCode() StatusCode
	// GetInnerDiagnosticInfo returns InnerDiagnosticInfo (property field)
	GetInnerDiagnosticInfo() DiagnosticInfo
	// IsDiagnosticInfo is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDiagnosticInfo()
}

// _DiagnosticInfo is the data-structure of this message
type _DiagnosticInfo struct {
	InnerDiagnosticInfoSpecified bool
	InnerStatusCodeSpecified     bool
	AdditionalInfoSpecified      bool
	LocaleSpecified              bool
	LocalizedTextSpecified       bool
	NamespaceURISpecified        bool
	SymbolicIdSpecified          bool
	SymbolicId                   *int32
	NamespaceURI                 *int32
	Locale                       *int32
	LocalizedText                *int32
	AdditionalInfo               PascalString
	InnerStatusCode              StatusCode
	InnerDiagnosticInfo          DiagnosticInfo
	// Reserved Fields
	reservedField0 *bool
}

var _ DiagnosticInfo = (*_DiagnosticInfo)(nil)

// NewDiagnosticInfo factory function for _DiagnosticInfo
func NewDiagnosticInfo(innerDiagnosticInfoSpecified bool, innerStatusCodeSpecified bool, additionalInfoSpecified bool, localeSpecified bool, localizedTextSpecified bool, namespaceURISpecified bool, symbolicIdSpecified bool, symbolicId *int32, namespaceURI *int32, locale *int32, localizedText *int32, additionalInfo PascalString, innerStatusCode StatusCode, innerDiagnosticInfo DiagnosticInfo) *_DiagnosticInfo {
	return &_DiagnosticInfo{InnerDiagnosticInfoSpecified: innerDiagnosticInfoSpecified, InnerStatusCodeSpecified: innerStatusCodeSpecified, AdditionalInfoSpecified: additionalInfoSpecified, LocaleSpecified: localeSpecified, LocalizedTextSpecified: localizedTextSpecified, NamespaceURISpecified: namespaceURISpecified, SymbolicIdSpecified: symbolicIdSpecified, SymbolicId: symbolicId, NamespaceURI: namespaceURI, Locale: locale, LocalizedText: localizedText, AdditionalInfo: additionalInfo, InnerStatusCode: innerStatusCode, InnerDiagnosticInfo: innerDiagnosticInfo}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DiagnosticInfo) GetInnerDiagnosticInfoSpecified() bool {
	return m.InnerDiagnosticInfoSpecified
}

func (m *_DiagnosticInfo) GetInnerStatusCodeSpecified() bool {
	return m.InnerStatusCodeSpecified
}

func (m *_DiagnosticInfo) GetAdditionalInfoSpecified() bool {
	return m.AdditionalInfoSpecified
}

func (m *_DiagnosticInfo) GetLocaleSpecified() bool {
	return m.LocaleSpecified
}

func (m *_DiagnosticInfo) GetLocalizedTextSpecified() bool {
	return m.LocalizedTextSpecified
}

func (m *_DiagnosticInfo) GetNamespaceURISpecified() bool {
	return m.NamespaceURISpecified
}

func (m *_DiagnosticInfo) GetSymbolicIdSpecified() bool {
	return m.SymbolicIdSpecified
}

func (m *_DiagnosticInfo) GetSymbolicId() *int32 {
	return m.SymbolicId
}

func (m *_DiagnosticInfo) GetNamespaceURI() *int32 {
	return m.NamespaceURI
}

func (m *_DiagnosticInfo) GetLocale() *int32 {
	return m.Locale
}

func (m *_DiagnosticInfo) GetLocalizedText() *int32 {
	return m.LocalizedText
}

func (m *_DiagnosticInfo) GetAdditionalInfo() PascalString {
	return m.AdditionalInfo
}

func (m *_DiagnosticInfo) GetInnerStatusCode() StatusCode {
	return m.InnerStatusCode
}

func (m *_DiagnosticInfo) GetInnerDiagnosticInfo() DiagnosticInfo {
	return m.InnerDiagnosticInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDiagnosticInfo(structType any) DiagnosticInfo {
	if casted, ok := structType.(DiagnosticInfo); ok {
		return casted
	}
	if casted, ok := structType.(*DiagnosticInfo); ok {
		return *casted
	}
	return nil
}

func (m *_DiagnosticInfo) GetTypeName() string {
	return "DiagnosticInfo"
}

func (m *_DiagnosticInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (innerDiagnosticInfoSpecified)
	lengthInBits += 1

	// Simple field (innerStatusCodeSpecified)
	lengthInBits += 1

	// Simple field (additionalInfoSpecified)
	lengthInBits += 1

	// Simple field (localeSpecified)
	lengthInBits += 1

	// Simple field (localizedTextSpecified)
	lengthInBits += 1

	// Simple field (namespaceURISpecified)
	lengthInBits += 1

	// Simple field (symbolicIdSpecified)
	lengthInBits += 1

	// Optional Field (symbolicId)
	if m.SymbolicId != nil {
		lengthInBits += 32
	}

	// Optional Field (namespaceURI)
	if m.NamespaceURI != nil {
		lengthInBits += 32
	}

	// Optional Field (locale)
	if m.Locale != nil {
		lengthInBits += 32
	}

	// Optional Field (localizedText)
	if m.LocalizedText != nil {
		lengthInBits += 32
	}

	// Optional Field (additionalInfo)
	if m.AdditionalInfo != nil {
		lengthInBits += m.AdditionalInfo.GetLengthInBits(ctx)
	}

	// Optional Field (innerStatusCode)
	if m.InnerStatusCode != nil {
		lengthInBits += m.InnerStatusCode.GetLengthInBits(ctx)
	}

	// Optional Field (innerDiagnosticInfo)
	if m.InnerDiagnosticInfo != nil {
		lengthInBits += m.InnerDiagnosticInfo.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_DiagnosticInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DiagnosticInfoParse(ctx context.Context, theBytes []byte) (DiagnosticInfo, error) {
	return DiagnosticInfoParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DiagnosticInfoParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DiagnosticInfo, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DiagnosticInfo, error) {
		return DiagnosticInfoParseWithBuffer(ctx, readBuffer)
	}
}

func DiagnosticInfoParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DiagnosticInfo, error) {
	v, err := (&_DiagnosticInfo{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DiagnosticInfo) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__diagnosticInfo DiagnosticInfo, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DiagnosticInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DiagnosticInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	innerDiagnosticInfoSpecified, err := ReadSimpleField(ctx, "innerDiagnosticInfoSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerDiagnosticInfoSpecified' field"))
	}
	m.InnerDiagnosticInfoSpecified = innerDiagnosticInfoSpecified

	innerStatusCodeSpecified, err := ReadSimpleField(ctx, "innerStatusCodeSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerStatusCodeSpecified' field"))
	}
	m.InnerStatusCodeSpecified = innerStatusCodeSpecified

	additionalInfoSpecified, err := ReadSimpleField(ctx, "additionalInfoSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInfoSpecified' field"))
	}
	m.AdditionalInfoSpecified = additionalInfoSpecified

	localeSpecified, err := ReadSimpleField(ctx, "localeSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeSpecified' field"))
	}
	m.LocaleSpecified = localeSpecified

	localizedTextSpecified, err := ReadSimpleField(ctx, "localizedTextSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localizedTextSpecified' field"))
	}
	m.LocalizedTextSpecified = localizedTextSpecified

	namespaceURISpecified, err := ReadSimpleField(ctx, "namespaceURISpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceURISpecified' field"))
	}
	m.NamespaceURISpecified = namespaceURISpecified

	symbolicIdSpecified, err := ReadSimpleField(ctx, "symbolicIdSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'symbolicIdSpecified' field"))
	}
	m.SymbolicIdSpecified = symbolicIdSpecified

	var symbolicId *int32
	symbolicId, err = ReadOptionalField[int32](ctx, "symbolicId", ReadSignedInt(readBuffer, uint8(32)), symbolicIdSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'symbolicId' field"))
	}
	m.SymbolicId = symbolicId

	var namespaceURI *int32
	namespaceURI, err = ReadOptionalField[int32](ctx, "namespaceURI", ReadSignedInt(readBuffer, uint8(32)), namespaceURISpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceURI' field"))
	}
	m.NamespaceURI = namespaceURI

	var locale *int32
	locale, err = ReadOptionalField[int32](ctx, "locale", ReadSignedInt(readBuffer, uint8(32)), localeSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'locale' field"))
	}
	m.Locale = locale

	var localizedText *int32
	localizedText, err = ReadOptionalField[int32](ctx, "localizedText", ReadSignedInt(readBuffer, uint8(32)), localizedTextSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localizedText' field"))
	}
	m.LocalizedText = localizedText

	var additionalInfo PascalString
	_additionalInfo, err := ReadOptionalField[PascalString](ctx, "additionalInfo", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), additionalInfoSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInfo' field"))
	}
	if _additionalInfo != nil {
		additionalInfo = *_additionalInfo
		m.AdditionalInfo = additionalInfo
	}

	var innerStatusCode StatusCode
	_innerStatusCode, err := ReadOptionalField[StatusCode](ctx, "innerStatusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), innerStatusCodeSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerStatusCode' field"))
	}
	if _innerStatusCode != nil {
		innerStatusCode = *_innerStatusCode
		m.InnerStatusCode = innerStatusCode
	}

	var innerDiagnosticInfo DiagnosticInfo
	_innerDiagnosticInfo, err := ReadOptionalField[DiagnosticInfo](ctx, "innerDiagnosticInfo", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), innerDiagnosticInfoSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerDiagnosticInfo' field"))
	}
	if _innerDiagnosticInfo != nil {
		innerDiagnosticInfo = *_innerDiagnosticInfo
		m.InnerDiagnosticInfo = innerDiagnosticInfo
	}

	if closeErr := readBuffer.CloseContext("DiagnosticInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DiagnosticInfo")
	}

	return m, nil
}

func (m *_DiagnosticInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DiagnosticInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DiagnosticInfo"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DiagnosticInfo")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "innerDiagnosticInfoSpecified", m.GetInnerDiagnosticInfoSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'innerDiagnosticInfoSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "innerStatusCodeSpecified", m.GetInnerStatusCodeSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'innerStatusCodeSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "additionalInfoSpecified", m.GetAdditionalInfoSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'additionalInfoSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "localeSpecified", m.GetLocaleSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localeSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "localizedTextSpecified", m.GetLocalizedTextSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localizedTextSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "namespaceURISpecified", m.GetNamespaceURISpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'namespaceURISpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "symbolicIdSpecified", m.GetSymbolicIdSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'symbolicIdSpecified' field")
	}

	if err := WriteOptionalField[int32](ctx, "symbolicId", m.GetSymbolicId(), WriteSignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'symbolicId' field")
	}

	if err := WriteOptionalField[int32](ctx, "namespaceURI", m.GetNamespaceURI(), WriteSignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'namespaceURI' field")
	}

	if err := WriteOptionalField[int32](ctx, "locale", m.GetLocale(), WriteSignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'locale' field")
	}

	if err := WriteOptionalField[int32](ctx, "localizedText", m.GetLocalizedText(), WriteSignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'localizedText' field")
	}

	if err := WriteOptionalField[PascalString](ctx, "additionalInfo", GetRef(m.GetAdditionalInfo()), WriteComplex[PascalString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'additionalInfo' field")
	}

	if err := WriteOptionalField[StatusCode](ctx, "innerStatusCode", GetRef(m.GetInnerStatusCode()), WriteComplex[StatusCode](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'innerStatusCode' field")
	}

	if err := WriteOptionalField[DiagnosticInfo](ctx, "innerDiagnosticInfo", GetRef(m.GetInnerDiagnosticInfo()), WriteComplex[DiagnosticInfo](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'innerDiagnosticInfo' field")
	}

	if popErr := writeBuffer.PopContext("DiagnosticInfo"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DiagnosticInfo")
	}
	return nil
}

func (m *_DiagnosticInfo) IsDiagnosticInfo() {}

func (m *_DiagnosticInfo) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DiagnosticInfo) deepCopy() *_DiagnosticInfo {
	if m == nil {
		return nil
	}
	_DiagnosticInfoCopy := &_DiagnosticInfo{
		m.InnerDiagnosticInfoSpecified,
		m.InnerStatusCodeSpecified,
		m.AdditionalInfoSpecified,
		m.LocaleSpecified,
		m.LocalizedTextSpecified,
		m.NamespaceURISpecified,
		m.SymbolicIdSpecified,
		utils.CopyPtr[int32](m.SymbolicId),
		utils.CopyPtr[int32](m.NamespaceURI),
		utils.CopyPtr[int32](m.Locale),
		utils.CopyPtr[int32](m.LocalizedText),
		m.AdditionalInfo.DeepCopy().(PascalString),
		m.InnerStatusCode.DeepCopy().(StatusCode),
		m.InnerDiagnosticInfo.DeepCopy().(DiagnosticInfo),
		m.reservedField0,
	}
	return _DiagnosticInfoCopy
}

func (m *_DiagnosticInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
