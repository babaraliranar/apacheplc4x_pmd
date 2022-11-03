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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtPropertyDescriptionResponse is the corresponding interface of ApduDataExtPropertyDescriptionResponse
type ApduDataExtPropertyDescriptionResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint8
	// GetWriteEnabled returns WriteEnabled (property field)
	GetWriteEnabled() bool
	// GetPropertyDataType returns PropertyDataType (property field)
	GetPropertyDataType() KnxPropertyDataType
	// GetMaxNrOfElements returns MaxNrOfElements (property field)
	GetMaxNrOfElements() uint16
	// GetReadLevel returns ReadLevel (property field)
	GetReadLevel() AccessLevel
	// GetWriteLevel returns WriteLevel (property field)
	GetWriteLevel() AccessLevel
}

// ApduDataExtPropertyDescriptionResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyDescriptionResponse.
// This is useful for switch cases.
type ApduDataExtPropertyDescriptionResponseExactly interface {
	ApduDataExtPropertyDescriptionResponse
	isApduDataExtPropertyDescriptionResponse() bool
}

// _ApduDataExtPropertyDescriptionResponse is the data-structure of this message
type _ApduDataExtPropertyDescriptionResponse struct {
	*_ApduDataExt
	ObjectIndex      uint8
	PropertyId       uint8
	Index            uint8
	WriteEnabled     bool
	PropertyDataType KnxPropertyDataType
	MaxNrOfElements  uint16
	ReadLevel        AccessLevel
	WriteLevel       AccessLevel
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionResponse) GetExtApciType() uint8 {
	return 0x19
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyDescriptionResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtPropertyDescriptionResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionResponse) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetIndex() uint8 {
	return m.Index
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetWriteEnabled() bool {
	return m.WriteEnabled
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetPropertyDataType() KnxPropertyDataType {
	return m.PropertyDataType
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetMaxNrOfElements() uint16 {
	return m.MaxNrOfElements
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetReadLevel() AccessLevel {
	return m.ReadLevel
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetWriteLevel() AccessLevel {
	return m.WriteLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyDescriptionResponse factory function for _ApduDataExtPropertyDescriptionResponse
func NewApduDataExtPropertyDescriptionResponse(objectIndex uint8, propertyId uint8, index uint8, writeEnabled bool, propertyDataType KnxPropertyDataType, maxNrOfElements uint16, readLevel AccessLevel, writeLevel AccessLevel, length uint8) *_ApduDataExtPropertyDescriptionResponse {
	_result := &_ApduDataExtPropertyDescriptionResponse{
		ObjectIndex:      objectIndex,
		PropertyId:       propertyId,
		Index:            index,
		WriteEnabled:     writeEnabled,
		PropertyDataType: propertyDataType,
		MaxNrOfElements:  maxNrOfElements,
		ReadLevel:        readLevel,
		WriteLevel:       writeLevel,
		_ApduDataExt:     NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyDescriptionResponse(structType interface{}) ApduDataExtPropertyDescriptionResponse {
	if casted, ok := structType.(ApduDataExtPropertyDescriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyDescriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionResponse"
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	// Simple field (writeEnabled)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (propertyDataType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (maxNrOfElements)
	lengthInBits += 12

	// Simple field (readLevel)
	lengthInBits += 4

	// Simple field (writeLevel)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtPropertyDescriptionResponseParse(theBytes []byte, length uint8) (ApduDataExtPropertyDescriptionResponse, error) {
	return ApduDataExtPropertyDescriptionResponseParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), length) // TODO: get endianness from mspec
}

func ApduDataExtPropertyDescriptionResponseParseWithBuffer(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtPropertyDescriptionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyDescriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyDescriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIndex)
	_objectIndex, _objectIndexErr := readBuffer.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field of ApduDataExtPropertyDescriptionResponse")
	}
	objectIndex := _objectIndex

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field of ApduDataExtPropertyDescriptionResponse")
	}
	propertyId := _propertyId

	// Simple Field (index)
	_index, _indexErr := readBuffer.ReadUint8("index", 8)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field of ApduDataExtPropertyDescriptionResponse")
	}
	index := _index

	// Simple Field (writeEnabled)
	_writeEnabled, _writeEnabledErr := readBuffer.ReadBit("writeEnabled")
	if _writeEnabledErr != nil {
		return nil, errors.Wrap(_writeEnabledErr, "Error parsing 'writeEnabled' field of ApduDataExtPropertyDescriptionResponse")
	}
	writeEnabled := _writeEnabled

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ApduDataExtPropertyDescriptionResponse")
		}
		if reserved != uint8(0x0) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (propertyDataType)
	if pullErr := readBuffer.PullContext("propertyDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyDataType")
	}
	_propertyDataType, _propertyDataTypeErr := KnxPropertyDataTypeParseWithBuffer(readBuffer)
	if _propertyDataTypeErr != nil {
		return nil, errors.Wrap(_propertyDataTypeErr, "Error parsing 'propertyDataType' field of ApduDataExtPropertyDescriptionResponse")
	}
	propertyDataType := _propertyDataType
	if closeErr := readBuffer.CloseContext("propertyDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyDataType")
	}

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ApduDataExtPropertyDescriptionResponse")
		}
		if reserved != uint8(0x0) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (maxNrOfElements)
	_maxNrOfElements, _maxNrOfElementsErr := readBuffer.ReadUint16("maxNrOfElements", 12)
	if _maxNrOfElementsErr != nil {
		return nil, errors.Wrap(_maxNrOfElementsErr, "Error parsing 'maxNrOfElements' field of ApduDataExtPropertyDescriptionResponse")
	}
	maxNrOfElements := _maxNrOfElements

	// Simple Field (readLevel)
	if pullErr := readBuffer.PullContext("readLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for readLevel")
	}
	_readLevel, _readLevelErr := AccessLevelParseWithBuffer(readBuffer)
	if _readLevelErr != nil {
		return nil, errors.Wrap(_readLevelErr, "Error parsing 'readLevel' field of ApduDataExtPropertyDescriptionResponse")
	}
	readLevel := _readLevel
	if closeErr := readBuffer.CloseContext("readLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for readLevel")
	}

	// Simple Field (writeLevel)
	if pullErr := readBuffer.PullContext("writeLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writeLevel")
	}
	_writeLevel, _writeLevelErr := AccessLevelParseWithBuffer(readBuffer)
	if _writeLevelErr != nil {
		return nil, errors.Wrap(_writeLevelErr, "Error parsing 'writeLevel' field of ApduDataExtPropertyDescriptionResponse")
	}
	writeLevel := _writeLevel
	if closeErr := readBuffer.CloseContext("writeLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writeLevel")
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyDescriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyDescriptionResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtPropertyDescriptionResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
		ObjectIndex:      objectIndex,
		PropertyId:       propertyId,
		Index:            index,
		WriteEnabled:     writeEnabled,
		PropertyDataType: propertyDataType,
		MaxNrOfElements:  maxNrOfElements,
		ReadLevel:        readLevel,
		WriteLevel:       writeLevel,
		reservedField0:   reservedField0,
		reservedField1:   reservedField1,
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtPropertyDescriptionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyDescriptionResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyDescriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyDescriptionResponse")
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.GetObjectIndex())
		_objectIndexErr := writeBuffer.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.GetPropertyId())
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (index)
		index := uint8(m.GetIndex())
		_indexErr := writeBuffer.WriteUint8("index", 8, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		// Simple Field (writeEnabled)
		writeEnabled := bool(m.GetWriteEnabled())
		_writeEnabledErr := writeBuffer.WriteBit("writeEnabled", (writeEnabled))
		if _writeEnabledErr != nil {
			return errors.Wrap(_writeEnabledErr, "Error serializing 'writeEnabled' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x0)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x0),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 1, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (propertyDataType)
		if pushErr := writeBuffer.PushContext("propertyDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyDataType")
		}
		_propertyDataTypeErr := writeBuffer.WriteSerializable(m.GetPropertyDataType())
		if popErr := writeBuffer.PopContext("propertyDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyDataType")
		}
		if _propertyDataTypeErr != nil {
			return errors.Wrap(_propertyDataTypeErr, "Error serializing 'propertyDataType' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x0)
			if m.reservedField1 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x0),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint8("reserved", 4, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (maxNrOfElements)
		maxNrOfElements := uint16(m.GetMaxNrOfElements())
		_maxNrOfElementsErr := writeBuffer.WriteUint16("maxNrOfElements", 12, (maxNrOfElements))
		if _maxNrOfElementsErr != nil {
			return errors.Wrap(_maxNrOfElementsErr, "Error serializing 'maxNrOfElements' field")
		}

		// Simple Field (readLevel)
		if pushErr := writeBuffer.PushContext("readLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for readLevel")
		}
		_readLevelErr := writeBuffer.WriteSerializable(m.GetReadLevel())
		if popErr := writeBuffer.PopContext("readLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for readLevel")
		}
		if _readLevelErr != nil {
			return errors.Wrap(_readLevelErr, "Error serializing 'readLevel' field")
		}

		// Simple Field (writeLevel)
		if pushErr := writeBuffer.PushContext("writeLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for writeLevel")
		}
		_writeLevelErr := writeBuffer.WriteSerializable(m.GetWriteLevel())
		if popErr := writeBuffer.PopContext("writeLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for writeLevel")
		}
		if _writeLevelErr != nil {
			return errors.Wrap(_writeLevelErr, "Error serializing 'writeLevel' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyDescriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyDescriptionResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyDescriptionResponse) isApduDataExtPropertyDescriptionResponse() bool {
	return true
}

func (m *_ApduDataExtPropertyDescriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
