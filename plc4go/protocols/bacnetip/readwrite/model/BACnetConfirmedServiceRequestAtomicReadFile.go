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

// BACnetConfirmedServiceRequestAtomicReadFile is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFile
type BACnetConfirmedServiceRequestAtomicReadFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetFileIdentifier returns FileIdentifier (property field)
	GetFileIdentifier() BACnetApplicationTagObjectIdentifier
	// GetAccessMethod returns AccessMethod (property field)
	GetAccessMethod() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// IsBACnetConfirmedServiceRequestAtomicReadFile is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicReadFile()
}

// _BACnetConfirmedServiceRequestAtomicReadFile is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFile struct {
	BACnetConfirmedServiceRequestContract
	FileIdentifier BACnetApplicationTagObjectIdentifier
	AccessMethod   BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
}

var _ BACnetConfirmedServiceRequestAtomicReadFile = (*_BACnetConfirmedServiceRequestAtomicReadFile)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestAtomicReadFile)(nil)

// NewBACnetConfirmedServiceRequestAtomicReadFile factory function for _BACnetConfirmedServiceRequestAtomicReadFile
func NewBACnetConfirmedServiceRequestAtomicReadFile(fileIdentifier BACnetApplicationTagObjectIdentifier, accessMethod BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAtomicReadFile {
	if fileIdentifier == nil {
		panic("fileIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetConfirmedServiceRequestAtomicReadFile must not be nil")
	}
	if accessMethod == nil {
		panic("accessMethod of type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord for BACnetConfirmedServiceRequestAtomicReadFile must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAtomicReadFile{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		FileIdentifier:                        fileIdentifier,
		AccessMethod:                          accessMethod,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetFileIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.FileIdentifier
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetAccessMethod() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFile(structType any) BACnetConfirmedServiceRequestAtomicReadFile {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFile"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (fileIdentifier)
	lengthInBits += m.FileIdentifier.GetLengthInBits(ctx)

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAtomicReadFile BACnetConfirmedServiceRequestAtomicReadFile, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "fileIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileIdentifier' field"))
	}
	m.FileIdentifier = fileIdentifier

	accessMethod, err := ReadSimpleField[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](ctx, "accessMethod", ReadComplex[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessMethod' field"))
	}
	m.AccessMethod = accessMethod

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFile")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFile")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "fileIdentifier", m.GetFileIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileIdentifier' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](ctx, "accessMethod", m.GetAccessMethod(), WriteComplex[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFile")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) IsBACnetConfirmedServiceRequestAtomicReadFile() {
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) deepCopy() *_BACnetConfirmedServiceRequestAtomicReadFile {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAtomicReadFileCopy := &_BACnetConfirmedServiceRequestAtomicReadFile{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.FileIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
		m.AccessMethod.DeepCopy().(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestAtomicReadFileCopy
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
