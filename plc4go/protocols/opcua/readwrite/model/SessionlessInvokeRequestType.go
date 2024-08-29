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

// SessionlessInvokeRequestType is the corresponding interface of SessionlessInvokeRequestType
type SessionlessInvokeRequestType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetUrisVersion returns UrisVersion (property field)
	GetUrisVersion() uint32
	// GetNoOfNamespaceUris returns NoOfNamespaceUris (property field)
	GetNoOfNamespaceUris() int32
	// GetNamespaceUris returns NamespaceUris (property field)
	GetNamespaceUris() []PascalString
	// GetNoOfServerUris returns NoOfServerUris (property field)
	GetNoOfServerUris() int32
	// GetServerUris returns ServerUris (property field)
	GetServerUris() []PascalString
	// GetNoOfLocaleIds returns NoOfLocaleIds (property field)
	GetNoOfLocaleIds() int32
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetServiceId returns ServiceId (property field)
	GetServiceId() uint32
}

// SessionlessInvokeRequestTypeExactly can be used when we want exactly this type and not a type which fulfills SessionlessInvokeRequestType.
// This is useful for switch cases.
type SessionlessInvokeRequestTypeExactly interface {
	SessionlessInvokeRequestType
	isSessionlessInvokeRequestType() bool
}

// _SessionlessInvokeRequestType is the data-structure of this message
type _SessionlessInvokeRequestType struct {
	*_ExtensionObjectDefinition
	UrisVersion       uint32
	NoOfNamespaceUris int32
	NamespaceUris     []PascalString
	NoOfServerUris    int32
	ServerUris        []PascalString
	NoOfLocaleIds     int32
	LocaleIds         []PascalString
	ServiceId         uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SessionlessInvokeRequestType) GetIdentifier() string {
	return "15903"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SessionlessInvokeRequestType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SessionlessInvokeRequestType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SessionlessInvokeRequestType) GetUrisVersion() uint32 {
	return m.UrisVersion
}

func (m *_SessionlessInvokeRequestType) GetNoOfNamespaceUris() int32 {
	return m.NoOfNamespaceUris
}

func (m *_SessionlessInvokeRequestType) GetNamespaceUris() []PascalString {
	return m.NamespaceUris
}

func (m *_SessionlessInvokeRequestType) GetNoOfServerUris() int32 {
	return m.NoOfServerUris
}

func (m *_SessionlessInvokeRequestType) GetServerUris() []PascalString {
	return m.ServerUris
}

func (m *_SessionlessInvokeRequestType) GetNoOfLocaleIds() int32 {
	return m.NoOfLocaleIds
}

func (m *_SessionlessInvokeRequestType) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_SessionlessInvokeRequestType) GetServiceId() uint32 {
	return m.ServiceId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSessionlessInvokeRequestType factory function for _SessionlessInvokeRequestType
func NewSessionlessInvokeRequestType(urisVersion uint32, noOfNamespaceUris int32, namespaceUris []PascalString, noOfServerUris int32, serverUris []PascalString, noOfLocaleIds int32, localeIds []PascalString, serviceId uint32) *_SessionlessInvokeRequestType {
	_result := &_SessionlessInvokeRequestType{
		UrisVersion:                urisVersion,
		NoOfNamespaceUris:          noOfNamespaceUris,
		NamespaceUris:              namespaceUris,
		NoOfServerUris:             noOfServerUris,
		ServerUris:                 serverUris,
		NoOfLocaleIds:              noOfLocaleIds,
		LocaleIds:                  localeIds,
		ServiceId:                  serviceId,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSessionlessInvokeRequestType(structType any) SessionlessInvokeRequestType {
	if casted, ok := structType.(SessionlessInvokeRequestType); ok {
		return casted
	}
	if casted, ok := structType.(*SessionlessInvokeRequestType); ok {
		return *casted
	}
	return nil
}

func (m *_SessionlessInvokeRequestType) GetTypeName() string {
	return "SessionlessInvokeRequestType"
}

func (m *_SessionlessInvokeRequestType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (urisVersion)
	lengthInBits += 32

	// Simple field (noOfNamespaceUris)
	lengthInBits += 32

	// Array field
	if len(m.NamespaceUris) > 0 {
		for _curItem, element := range m.NamespaceUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NamespaceUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfServerUris)
	lengthInBits += 32

	// Array field
	if len(m.ServerUris) > 0 {
		for _curItem, element := range m.ServerUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (serviceId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SessionlessInvokeRequestType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SessionlessInvokeRequestTypeParse(ctx context.Context, theBytes []byte, identifier string) (SessionlessInvokeRequestType, error) {
	return SessionlessInvokeRequestTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SessionlessInvokeRequestTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SessionlessInvokeRequestType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SessionlessInvokeRequestType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SessionlessInvokeRequestType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (urisVersion)
	_urisVersion, _urisVersionErr := readBuffer.ReadUint32("urisVersion", 32)
	if _urisVersionErr != nil {
		return nil, errors.Wrap(_urisVersionErr, "Error parsing 'urisVersion' field of SessionlessInvokeRequestType")
	}
	urisVersion := _urisVersion

	// Simple Field (noOfNamespaceUris)
	_noOfNamespaceUris, _noOfNamespaceUrisErr := readBuffer.ReadInt32("noOfNamespaceUris", 32)
	if _noOfNamespaceUrisErr != nil {
		return nil, errors.Wrap(_noOfNamespaceUrisErr, "Error parsing 'noOfNamespaceUris' field of SessionlessInvokeRequestType")
	}
	noOfNamespaceUris := _noOfNamespaceUris

	// Array field (namespaceUris)
	if pullErr := readBuffer.PullContext("namespaceUris", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for namespaceUris")
	}
	// Count array
	namespaceUris := make([]PascalString, max(noOfNamespaceUris, 0))
	// This happens when the size is set conditional to 0
	if len(namespaceUris) == 0 {
		namespaceUris = nil
	}
	{
		_numItems := uint16(max(noOfNamespaceUris, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'namespaceUris' field of SessionlessInvokeRequestType")
			}
			namespaceUris[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("namespaceUris", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for namespaceUris")
	}

	// Simple Field (noOfServerUris)
	_noOfServerUris, _noOfServerUrisErr := readBuffer.ReadInt32("noOfServerUris", 32)
	if _noOfServerUrisErr != nil {
		return nil, errors.Wrap(_noOfServerUrisErr, "Error parsing 'noOfServerUris' field of SessionlessInvokeRequestType")
	}
	noOfServerUris := _noOfServerUris

	// Array field (serverUris)
	if pullErr := readBuffer.PullContext("serverUris", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverUris")
	}
	// Count array
	serverUris := make([]PascalString, max(noOfServerUris, 0))
	// This happens when the size is set conditional to 0
	if len(serverUris) == 0 {
		serverUris = nil
	}
	{
		_numItems := uint16(max(noOfServerUris, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'serverUris' field of SessionlessInvokeRequestType")
			}
			serverUris[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("serverUris", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverUris")
	}

	// Simple Field (noOfLocaleIds)
	_noOfLocaleIds, _noOfLocaleIdsErr := readBuffer.ReadInt32("noOfLocaleIds", 32)
	if _noOfLocaleIdsErr != nil {
		return nil, errors.Wrap(_noOfLocaleIdsErr, "Error parsing 'noOfLocaleIds' field of SessionlessInvokeRequestType")
	}
	noOfLocaleIds := _noOfLocaleIds

	// Array field (localeIds)
	if pullErr := readBuffer.PullContext("localeIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localeIds")
	}
	// Count array
	localeIds := make([]PascalString, max(noOfLocaleIds, 0))
	// This happens when the size is set conditional to 0
	if len(localeIds) == 0 {
		localeIds = nil
	}
	{
		_numItems := uint16(max(noOfLocaleIds, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'localeIds' field of SessionlessInvokeRequestType")
			}
			localeIds[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("localeIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localeIds")
	}

	// Simple Field (serviceId)
	_serviceId, _serviceIdErr := readBuffer.ReadUint32("serviceId", 32)
	if _serviceIdErr != nil {
		return nil, errors.Wrap(_serviceIdErr, "Error parsing 'serviceId' field of SessionlessInvokeRequestType")
	}
	serviceId := _serviceId

	if closeErr := readBuffer.CloseContext("SessionlessInvokeRequestType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SessionlessInvokeRequestType")
	}

	// Create a partially initialized instance
	_child := &_SessionlessInvokeRequestType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		UrisVersion:                urisVersion,
		NoOfNamespaceUris:          noOfNamespaceUris,
		NamespaceUris:              namespaceUris,
		NoOfServerUris:             noOfServerUris,
		ServerUris:                 serverUris,
		NoOfLocaleIds:              noOfLocaleIds,
		LocaleIds:                  localeIds,
		ServiceId:                  serviceId,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SessionlessInvokeRequestType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SessionlessInvokeRequestType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SessionlessInvokeRequestType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SessionlessInvokeRequestType")
		}

		// Simple Field (urisVersion)
		urisVersion := uint32(m.GetUrisVersion())
		_urisVersionErr := writeBuffer.WriteUint32("urisVersion", 32, uint32((urisVersion)))
		if _urisVersionErr != nil {
			return errors.Wrap(_urisVersionErr, "Error serializing 'urisVersion' field")
		}

		// Simple Field (noOfNamespaceUris)
		noOfNamespaceUris := int32(m.GetNoOfNamespaceUris())
		_noOfNamespaceUrisErr := writeBuffer.WriteInt32("noOfNamespaceUris", 32, int32((noOfNamespaceUris)))
		if _noOfNamespaceUrisErr != nil {
			return errors.Wrap(_noOfNamespaceUrisErr, "Error serializing 'noOfNamespaceUris' field")
		}

		// Array Field (namespaceUris)
		if pushErr := writeBuffer.PushContext("namespaceUris", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for namespaceUris")
		}
		for _curItem, _element := range m.GetNamespaceUris() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNamespaceUris()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'namespaceUris' field")
			}
		}
		if popErr := writeBuffer.PopContext("namespaceUris", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for namespaceUris")
		}

		// Simple Field (noOfServerUris)
		noOfServerUris := int32(m.GetNoOfServerUris())
		_noOfServerUrisErr := writeBuffer.WriteInt32("noOfServerUris", 32, int32((noOfServerUris)))
		if _noOfServerUrisErr != nil {
			return errors.Wrap(_noOfServerUrisErr, "Error serializing 'noOfServerUris' field")
		}

		// Array Field (serverUris)
		if pushErr := writeBuffer.PushContext("serverUris", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverUris")
		}
		for _curItem, _element := range m.GetServerUris() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServerUris()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'serverUris' field")
			}
		}
		if popErr := writeBuffer.PopContext("serverUris", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverUris")
		}

		// Simple Field (noOfLocaleIds)
		noOfLocaleIds := int32(m.GetNoOfLocaleIds())
		_noOfLocaleIdsErr := writeBuffer.WriteInt32("noOfLocaleIds", 32, int32((noOfLocaleIds)))
		if _noOfLocaleIdsErr != nil {
			return errors.Wrap(_noOfLocaleIdsErr, "Error serializing 'noOfLocaleIds' field")
		}

		// Array Field (localeIds)
		if pushErr := writeBuffer.PushContext("localeIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for localeIds")
		}
		for _curItem, _element := range m.GetLocaleIds() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetLocaleIds()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'localeIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("localeIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for localeIds")
		}

		// Simple Field (serviceId)
		serviceId := uint32(m.GetServiceId())
		_serviceIdErr := writeBuffer.WriteUint32("serviceId", 32, uint32((serviceId)))
		if _serviceIdErr != nil {
			return errors.Wrap(_serviceIdErr, "Error serializing 'serviceId' field")
		}

		if popErr := writeBuffer.PopContext("SessionlessInvokeRequestType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SessionlessInvokeRequestType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SessionlessInvokeRequestType) isSessionlessInvokeRequestType() bool {
	return true
}

func (m *_SessionlessInvokeRequestType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
