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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestAtomicWriteFile is the corresponding interface of BACnetConfirmedServiceRequestAtomicWriteFile
type BACnetConfirmedServiceRequestAtomicWriteFile interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetApplicationTagSignedInteger
	// GetFileData returns FileData (property field)
	GetFileData() BACnetApplicationTagOctetString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetConfirmedServiceRequestAtomicWriteFileExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestAtomicWriteFile.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestAtomicWriteFileExactly interface {
	BACnetConfirmedServiceRequestAtomicWriteFile
	isBACnetConfirmedServiceRequestAtomicWriteFile() bool
}

// _BACnetConfirmedServiceRequestAtomicWriteFile is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicWriteFile struct {
	*_BACnetConfirmedServiceRequest
	DeviceIdentifier  BACnetApplicationTagObjectIdentifier
	OpeningTag        BACnetOpeningTag
	FileStartPosition BACnetApplicationTagSignedInteger
	FileData          BACnetApplicationTagOctetString
	ClosingTag        BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetFileStartPosition() BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetFileData() BACnetApplicationTagOctetString {
	return m.FileData
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicWriteFile factory function for _BACnetConfirmedServiceRequestAtomicWriteFile
func NewBACnetConfirmedServiceRequestAtomicWriteFile(deviceIdentifier BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, fileStartPosition BACnetApplicationTagSignedInteger, fileData BACnetApplicationTagOctetString, closingTag BACnetClosingTag, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAtomicWriteFile {
	_result := &_BACnetConfirmedServiceRequestAtomicWriteFile{
		DeviceIdentifier:               deviceIdentifier,
		OpeningTag:                     openingTag,
		FileStartPosition:              fileStartPosition,
		FileData:                       fileData,
		ClosingTag:                     closingTag,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicWriteFile(structType interface{}) BACnetConfirmedServiceRequestAtomicWriteFile {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicWriteFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicWriteFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicWriteFile"
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits()

	// Optional Field (openingTag)
	if m.OpeningTag != nil {
		lengthInBits += m.OpeningTag.GetLengthInBits()
	}

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits()

	// Simple field (fileData)
	lengthInBits += m.FileData.GetLengthInBits()

	// Optional Field (closingTag)
	if m.ClosingTag != nil {
		lengthInBits += m.ClosingTag.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicWriteFileParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestAtomicWriteFile, error) {
	return BACnetConfirmedServiceRequestAtomicWriteFileParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), serviceRequestLength) // TODO: get endianness from mspec
}

func BACnetConfirmedServiceRequestAtomicWriteFileParseWithBuffer(readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestAtomicWriteFile, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicWriteFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
	}
	_deviceIdentifier, _deviceIdentifierErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field of BACnetConfirmedServiceRequestAtomicWriteFile")
	}
	deviceIdentifier := _deviceIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
	}

	// Optional Field (openingTag) (Can be skipped, if a given expression evaluates to false)
	var openingTag BACnetOpeningTag = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
		}
		_val, _err := BACnetOpeningTagParseWithBuffer(readBuffer, uint8(0))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'openingTag' field of BACnetConfirmedServiceRequestAtomicWriteFile")
		default:
			openingTag = _val.(BACnetOpeningTag)
			if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for openingTag")
			}
		}
	}

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileStartPosition")
	}
	_fileStartPosition, _fileStartPositionErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field of BACnetConfirmedServiceRequestAtomicWriteFile")
	}
	fileStartPosition := _fileStartPosition.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileStartPosition")
	}

	// Simple Field (fileData)
	if pullErr := readBuffer.PullContext("fileData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileData")
	}
	_fileData, _fileDataErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _fileDataErr != nil {
		return nil, errors.Wrap(_fileDataErr, "Error parsing 'fileData' field of BACnetConfirmedServiceRequestAtomicWriteFile")
	}
	fileData := _fileData.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("fileData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileData")
	}

	// Optional Field (closingTag) (Can be skipped, if a given expression evaluates to false)
	var closingTag BACnetClosingTag = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
		}
		_val, _err := BACnetClosingTagParseWithBuffer(readBuffer, uint8(0))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'closingTag' field of BACnetConfirmedServiceRequestAtomicWriteFile")
		default:
			closingTag = _val.(BACnetClosingTag)
			if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for closingTag")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicWriteFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicWriteFile")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestAtomicWriteFile{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		DeviceIdentifier:  deviceIdentifier,
		OpeningTag:        openingTag,
		FileStartPosition: fileStartPosition,
		FileData:          fileData,
		ClosingTag:        closingTag,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicWriteFile")
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		_deviceIdentifierErr := writeBuffer.WriteSerializable(m.GetDeviceIdentifier())
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}

		// Optional Field (openingTag) (Can be skipped, if the value is null)
		var openingTag BACnetOpeningTag = nil
		if m.GetOpeningTag() != nil {
			if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for openingTag")
			}
			openingTag = m.GetOpeningTag()
			_openingTagErr := writeBuffer.WriteSerializable(openingTag)
			if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for openingTag")
			}
			if _openingTagErr != nil {
				return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
			}
		}

		// Simple Field (fileStartPosition)
		if pushErr := writeBuffer.PushContext("fileStartPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileStartPosition")
		}
		_fileStartPositionErr := writeBuffer.WriteSerializable(m.GetFileStartPosition())
		if popErr := writeBuffer.PopContext("fileStartPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileStartPosition")
		}
		if _fileStartPositionErr != nil {
			return errors.Wrap(_fileStartPositionErr, "Error serializing 'fileStartPosition' field")
		}

		// Simple Field (fileData)
		if pushErr := writeBuffer.PushContext("fileData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileData")
		}
		_fileDataErr := writeBuffer.WriteSerializable(m.GetFileData())
		if popErr := writeBuffer.PopContext("fileData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileData")
		}
		if _fileDataErr != nil {
			return errors.Wrap(_fileDataErr, "Error serializing 'fileData' field")
		}

		// Optional Field (closingTag) (Can be skipped, if the value is null)
		var closingTag BACnetClosingTag = nil
		if m.GetClosingTag() != nil {
			if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for closingTag")
			}
			closingTag = m.GetClosingTag()
			_closingTagErr := writeBuffer.WriteSerializable(closingTag)
			if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for closingTag")
			}
			if _closingTagErr != nil {
				return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicWriteFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicWriteFile")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) isBACnetConfirmedServiceRequestAtomicWriteFile() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
