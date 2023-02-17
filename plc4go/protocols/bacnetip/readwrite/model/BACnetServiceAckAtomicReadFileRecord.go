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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetServiceAckAtomicReadFileRecord is the corresponding interface of BACnetServiceAckAtomicReadFileRecord
type BACnetServiceAckAtomicReadFileRecord interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAckAtomicReadFileStreamOrRecord
	// GetFileStartRecord returns FileStartRecord (property field)
	GetFileStartRecord() BACnetApplicationTagSignedInteger
	// GetReturnedRecordCount returns ReturnedRecordCount (property field)
	GetReturnedRecordCount() BACnetApplicationTagUnsignedInteger
	// GetFileRecordData returns FileRecordData (property field)
	GetFileRecordData() []BACnetApplicationTagOctetString
}

// BACnetServiceAckAtomicReadFileRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAtomicReadFileRecord.
// This is useful for switch cases.
type BACnetServiceAckAtomicReadFileRecordExactly interface {
	BACnetServiceAckAtomicReadFileRecord
	isBACnetServiceAckAtomicReadFileRecord() bool
}

// _BACnetServiceAckAtomicReadFileRecord is the data-structure of this message
type _BACnetServiceAckAtomicReadFileRecord struct {
	*_BACnetServiceAckAtomicReadFileStreamOrRecord
        FileStartRecord BACnetApplicationTagSignedInteger
        ReturnedRecordCount BACnetApplicationTagUnsignedInteger
        FileRecordData []BACnetApplicationTagOctetString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicReadFileRecord) InitializeParent(parent BACnetServiceAckAtomicReadFileStreamOrRecord , peekedTagHeader BACnetTagHeader , openingTag BACnetOpeningTag , closingTag BACnetClosingTag ) {	m.PeekedTagHeader = peekedTagHeader
	m.OpeningTag = openingTag
	m.ClosingTag = closingTag
}

func (m *_BACnetServiceAckAtomicReadFileRecord)  GetParent() BACnetServiceAckAtomicReadFileStreamOrRecord {
	return m._BACnetServiceAckAtomicReadFileStreamOrRecord
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileRecord) GetFileStartRecord() BACnetApplicationTagSignedInteger {
	return m.FileStartRecord
}

func (m *_BACnetServiceAckAtomicReadFileRecord) GetReturnedRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.ReturnedRecordCount
}

func (m *_BACnetServiceAckAtomicReadFileRecord) GetFileRecordData() []BACnetApplicationTagOctetString {
	return m.FileRecordData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetServiceAckAtomicReadFileRecord factory function for _BACnetServiceAckAtomicReadFileRecord
func NewBACnetServiceAckAtomicReadFileRecord( fileStartRecord BACnetApplicationTagSignedInteger , returnedRecordCount BACnetApplicationTagUnsignedInteger , fileRecordData []BACnetApplicationTagOctetString , peekedTagHeader BACnetTagHeader , openingTag BACnetOpeningTag , closingTag BACnetClosingTag ) *_BACnetServiceAckAtomicReadFileRecord {
	_result := &_BACnetServiceAckAtomicReadFileRecord{
		FileStartRecord: fileStartRecord,
		ReturnedRecordCount: returnedRecordCount,
		FileRecordData: fileRecordData,
    	_BACnetServiceAckAtomicReadFileStreamOrRecord: NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
	}
	_result._BACnetServiceAckAtomicReadFileStreamOrRecord._BACnetServiceAckAtomicReadFileStreamOrRecordChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFileRecord(structType interface{}) BACnetServiceAckAtomicReadFileRecord {
    if casted, ok := structType.(BACnetServiceAckAtomicReadFileRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileRecord) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileRecord"
}

func (m *_BACnetServiceAckAtomicReadFileRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (fileStartRecord)
	lengthInBits += m.FileStartRecord.GetLengthInBits(ctx)

	// Simple field (returnedRecordCount)
	lengthInBits += m.ReturnedRecordCount.GetLengthInBits(ctx)

	// Array field
	if len(m.FileRecordData) > 0 {
		for _curItem, element := range m.FileRecordData {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.FileRecordData), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_BACnetServiceAckAtomicReadFileRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckAtomicReadFileRecordParse(theBytes []byte) (BACnetServiceAckAtomicReadFileRecord, error) {
	return BACnetServiceAckAtomicReadFileRecordParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetServiceAckAtomicReadFileRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckAtomicReadFileRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFileRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileStartRecord)
	if pullErr := readBuffer.PullContext("fileStartRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileStartRecord")
	}
_fileStartRecord, _fileStartRecordErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _fileStartRecordErr != nil {
		return nil, errors.Wrap(_fileStartRecordErr, "Error parsing 'fileStartRecord' field of BACnetServiceAckAtomicReadFileRecord")
	}
	fileStartRecord := _fileStartRecord.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("fileStartRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileStartRecord")
	}

	// Simple Field (returnedRecordCount)
	if pullErr := readBuffer.PullContext("returnedRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnedRecordCount")
	}
_returnedRecordCount, _returnedRecordCountErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _returnedRecordCountErr != nil {
		return nil, errors.Wrap(_returnedRecordCountErr, "Error parsing 'returnedRecordCount' field of BACnetServiceAckAtomicReadFileRecord")
	}
	returnedRecordCount := _returnedRecordCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("returnedRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnedRecordCount")
	}

	// Array field (fileRecordData)
	if pullErr := readBuffer.PullContext("fileRecordData", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileRecordData")
	}
	// Count array
	fileRecordData := make([]BACnetApplicationTagOctetString, returnedRecordCount.GetPayload().GetActualValue())
	// This happens when the size is set conditional to 0
	if len(fileRecordData) == 0 {
		fileRecordData = nil
	}
	{
		_numItems := uint16(returnedRecordCount.GetPayload().GetActualValue())
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := BACnetApplicationTagParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'fileRecordData' field of BACnetServiceAckAtomicReadFileRecord")
			}
			fileRecordData[_curItem] = _item.(BACnetApplicationTagOctetString)
		}
	}
	if closeErr := readBuffer.CloseContext("fileRecordData", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileRecordData")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFileRecord")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckAtomicReadFileRecord{
		_BACnetServiceAckAtomicReadFileStreamOrRecord: &_BACnetServiceAckAtomicReadFileStreamOrRecord{
		},
		FileStartRecord: fileStartRecord,
		ReturnedRecordCount: returnedRecordCount,
		FileRecordData: fileRecordData,
	}
	_child._BACnetServiceAckAtomicReadFileStreamOrRecord._BACnetServiceAckAtomicReadFileStreamOrRecordChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckAtomicReadFileRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicReadFileRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFileRecord")
		}

	// Simple Field (fileStartRecord)
	if pushErr := writeBuffer.PushContext("fileStartRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for fileStartRecord")
	}
	_fileStartRecordErr := writeBuffer.WriteSerializable(ctx, m.GetFileStartRecord())
	if popErr := writeBuffer.PopContext("fileStartRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for fileStartRecord")
	}
	if _fileStartRecordErr != nil {
		return errors.Wrap(_fileStartRecordErr, "Error serializing 'fileStartRecord' field")
	}

	// Simple Field (returnedRecordCount)
	if pushErr := writeBuffer.PushContext("returnedRecordCount"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnedRecordCount")
	}
	_returnedRecordCountErr := writeBuffer.WriteSerializable(ctx, m.GetReturnedRecordCount())
	if popErr := writeBuffer.PopContext("returnedRecordCount"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnedRecordCount")
	}
	if _returnedRecordCountErr != nil {
		return errors.Wrap(_returnedRecordCountErr, "Error serializing 'returnedRecordCount' field")
	}

	// Array Field (fileRecordData)
	if pushErr := writeBuffer.PushContext("fileRecordData", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for fileRecordData")
	}
	for _curItem, _element := range m.GetFileRecordData() {
		_ = _curItem
		arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetFileRecordData()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'fileRecordData' field")
		}
	}
	if popErr := writeBuffer.PopContext("fileRecordData", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for fileRecordData")
	}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFileRecord")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetServiceAckAtomicReadFileRecord) isBACnetServiceAckAtomicReadFileRecord() bool {
	return true
}

func (m *_BACnetServiceAckAtomicReadFileRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



