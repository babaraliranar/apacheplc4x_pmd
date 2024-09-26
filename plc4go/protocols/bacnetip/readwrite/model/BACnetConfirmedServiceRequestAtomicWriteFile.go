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

// BACnetConfirmedServiceRequestAtomicWriteFile is the corresponding interface of BACnetConfirmedServiceRequestAtomicWriteFile
type BACnetConfirmedServiceRequestAtomicWriteFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
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
	// IsBACnetConfirmedServiceRequestAtomicWriteFile is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicWriteFile()
}

// _BACnetConfirmedServiceRequestAtomicWriteFile is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicWriteFile struct {
	BACnetConfirmedServiceRequestContract
	DeviceIdentifier  BACnetApplicationTagObjectIdentifier
	OpeningTag        BACnetOpeningTag
	FileStartPosition BACnetApplicationTagSignedInteger
	FileData          BACnetApplicationTagOctetString
	ClosingTag        BACnetClosingTag
}

var _ BACnetConfirmedServiceRequestAtomicWriteFile = (*_BACnetConfirmedServiceRequestAtomicWriteFile)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestAtomicWriteFile)(nil)

// NewBACnetConfirmedServiceRequestAtomicWriteFile factory function for _BACnetConfirmedServiceRequestAtomicWriteFile
func NewBACnetConfirmedServiceRequestAtomicWriteFile(deviceIdentifier BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, fileStartPosition BACnetApplicationTagSignedInteger, fileData BACnetApplicationTagOctetString, closingTag BACnetClosingTag, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAtomicWriteFile {
	if deviceIdentifier == nil {
		panic("deviceIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetConfirmedServiceRequestAtomicWriteFile must not be nil")
	}
	if fileStartPosition == nil {
		panic("fileStartPosition of type BACnetApplicationTagSignedInteger for BACnetConfirmedServiceRequestAtomicWriteFile must not be nil")
	}
	if fileData == nil {
		panic("fileData of type BACnetApplicationTagOctetString for BACnetConfirmedServiceRequestAtomicWriteFile must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAtomicWriteFile{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		DeviceIdentifier:                      deviceIdentifier,
		OpeningTag:                            openingTag,
		FileStartPosition:                     fileStartPosition,
		FileData:                              fileData,
		ClosingTag:                            closingTag,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
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

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
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

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicWriteFile(structType any) BACnetConfirmedServiceRequestAtomicWriteFile {
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

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)

	// Optional Field (openingTag)
	if m.OpeningTag != nil {
		lengthInBits += m.OpeningTag.GetLengthInBits(ctx)
	}

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	// Simple field (fileData)
	lengthInBits += m.FileData.GetLengthInBits(ctx)

	// Optional Field (closingTag)
	if m.ClosingTag != nil {
		lengthInBits += m.ClosingTag.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAtomicWriteFile BACnetConfirmedServiceRequestAtomicWriteFile, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicWriteFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}
	m.DeviceIdentifier = deviceIdentifier

	var openingTag BACnetOpeningTag
	_openingTag, err := ReadOptionalField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	if _openingTag != nil {
		openingTag = *_openingTag
		m.OpeningTag = openingTag
	}

	fileStartPosition, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	fileData, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "fileData", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileData' field"))
	}
	m.FileData = fileData

	var closingTag BACnetClosingTag
	_closingTag, err := ReadOptionalField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	if _closingTag != nil {
		closingTag = *_closingTag
		m.ClosingTag = closingTag
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicWriteFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicWriteFile")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicWriteFile")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", m.GetDeviceIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceIdentifier' field")
		}

		if err := WriteOptionalField[BACnetOpeningTag](ctx, "openingTag", GetRef(m.GetOpeningTag()), WriteComplex[BACnetOpeningTag](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "fileData", m.GetFileData(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileData' field")
		}

		if err := WriteOptionalField[BACnetClosingTag](ctx, "closingTag", GetRef(m.GetClosingTag()), WriteComplex[BACnetClosingTag](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicWriteFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicWriteFile")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) IsBACnetConfirmedServiceRequestAtomicWriteFile() {
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) deepCopy() *_BACnetConfirmedServiceRequestAtomicWriteFile {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAtomicWriteFileCopy := &_BACnetConfirmedServiceRequestAtomicWriteFile{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.DeviceIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.FileStartPosition.DeepCopy().(BACnetApplicationTagSignedInteger),
		m.FileData.DeepCopy().(BACnetApplicationTagOctetString),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestAtomicWriteFileCopy
}

func (m *_BACnetConfirmedServiceRequestAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
