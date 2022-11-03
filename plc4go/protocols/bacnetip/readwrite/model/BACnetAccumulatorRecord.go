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

// BACnetAccumulatorRecord is the corresponding interface of BACnetAccumulatorRecord
type BACnetAccumulatorRecord interface {
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetContextTagSignedInteger
	// GetAccumulatedValue returns AccumulatedValue (property field)
	GetAccumulatedValue() BACnetContextTagSignedInteger
	// GetAccumulatorStatus returns AccumulatorStatus (property field)
	GetAccumulatorStatus() BACnetAccumulatorRecordAccumulatorStatusTagged
}

// BACnetAccumulatorRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetAccumulatorRecord.
// This is useful for switch cases.
type BACnetAccumulatorRecordExactly interface {
	BACnetAccumulatorRecord
	isBACnetAccumulatorRecord() bool
}

// _BACnetAccumulatorRecord is the data-structure of this message
type _BACnetAccumulatorRecord struct {
	Timestamp         BACnetDateTimeEnclosed
	PresentValue      BACnetContextTagSignedInteger
	AccumulatedValue  BACnetContextTagSignedInteger
	AccumulatorStatus BACnetAccumulatorRecordAccumulatorStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccumulatorRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetAccumulatorRecord) GetPresentValue() BACnetContextTagSignedInteger {
	return m.PresentValue
}

func (m *_BACnetAccumulatorRecord) GetAccumulatedValue() BACnetContextTagSignedInteger {
	return m.AccumulatedValue
}

func (m *_BACnetAccumulatorRecord) GetAccumulatorStatus() BACnetAccumulatorRecordAccumulatorStatusTagged {
	return m.AccumulatorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccumulatorRecord factory function for _BACnetAccumulatorRecord
func NewBACnetAccumulatorRecord(timestamp BACnetDateTimeEnclosed, presentValue BACnetContextTagSignedInteger, accumulatedValue BACnetContextTagSignedInteger, accumulatorStatus BACnetAccumulatorRecordAccumulatorStatusTagged) *_BACnetAccumulatorRecord {
	return &_BACnetAccumulatorRecord{Timestamp: timestamp, PresentValue: presentValue, AccumulatedValue: accumulatedValue, AccumulatorStatus: accumulatorStatus}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccumulatorRecord(structType interface{}) BACnetAccumulatorRecord {
	if casted, ok := structType.(BACnetAccumulatorRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccumulatorRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccumulatorRecord) GetTypeName() string {
	return "BACnetAccumulatorRecord"
}

func (m *_BACnetAccumulatorRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAccumulatorRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits()

	// Simple field (accumulatedValue)
	lengthInBits += m.AccumulatedValue.GetLengthInBits()

	// Simple field (accumulatorStatus)
	lengthInBits += m.AccumulatorStatus.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetAccumulatorRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccumulatorRecordParse(theBytes []byte) (BACnetAccumulatorRecord, error) {
	return BACnetAccumulatorRecordParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetAccumulatorRecordParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccumulatorRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccumulatorRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
	}
	_timestamp, _timestampErr := BACnetDateTimeEnclosedParseWithBuffer(readBuffer, uint8(uint8(0)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of BACnetAccumulatorRecord")
	}
	timestamp := _timestamp.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamp")
	}

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for presentValue")
	}
	_presentValue, _presentValueErr := BACnetContextTagParseWithBuffer(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field of BACnetAccumulatorRecord")
	}
	presentValue := _presentValue.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for presentValue")
	}

	// Simple Field (accumulatedValue)
	if pullErr := readBuffer.PullContext("accumulatedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accumulatedValue")
	}
	_accumulatedValue, _accumulatedValueErr := BACnetContextTagParseWithBuffer(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _accumulatedValueErr != nil {
		return nil, errors.Wrap(_accumulatedValueErr, "Error parsing 'accumulatedValue' field of BACnetAccumulatorRecord")
	}
	accumulatedValue := _accumulatedValue.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("accumulatedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accumulatedValue")
	}

	// Simple Field (accumulatorStatus)
	if pullErr := readBuffer.PullContext("accumulatorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accumulatorStatus")
	}
	_accumulatorStatus, _accumulatorStatusErr := BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBuffer(readBuffer, uint8(uint8(3)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _accumulatorStatusErr != nil {
		return nil, errors.Wrap(_accumulatorStatusErr, "Error parsing 'accumulatorStatus' field of BACnetAccumulatorRecord")
	}
	accumulatorStatus := _accumulatorStatus.(BACnetAccumulatorRecordAccumulatorStatusTagged)
	if closeErr := readBuffer.CloseContext("accumulatorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accumulatorStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetAccumulatorRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccumulatorRecord")
	}

	// Create the instance
	return &_BACnetAccumulatorRecord{
		Timestamp:         timestamp,
		PresentValue:      presentValue,
		AccumulatedValue:  accumulatedValue,
		AccumulatorStatus: accumulatorStatus,
	}, nil
}

func (m *_BACnetAccumulatorRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccumulatorRecord) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAccumulatorRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccumulatorRecord")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamp")
	}
	_timestampErr := writeBuffer.WriteSerializable(m.GetTimestamp())
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamp")
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (presentValue)
	if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for presentValue")
	}
	_presentValueErr := writeBuffer.WriteSerializable(m.GetPresentValue())
	if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for presentValue")
	}
	if _presentValueErr != nil {
		return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
	}

	// Simple Field (accumulatedValue)
	if pushErr := writeBuffer.PushContext("accumulatedValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for accumulatedValue")
	}
	_accumulatedValueErr := writeBuffer.WriteSerializable(m.GetAccumulatedValue())
	if popErr := writeBuffer.PopContext("accumulatedValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for accumulatedValue")
	}
	if _accumulatedValueErr != nil {
		return errors.Wrap(_accumulatedValueErr, "Error serializing 'accumulatedValue' field")
	}

	// Simple Field (accumulatorStatus)
	if pushErr := writeBuffer.PushContext("accumulatorStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for accumulatorStatus")
	}
	_accumulatorStatusErr := writeBuffer.WriteSerializable(m.GetAccumulatorStatus())
	if popErr := writeBuffer.PopContext("accumulatorStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for accumulatorStatus")
	}
	if _accumulatorStatusErr != nil {
		return errors.Wrap(_accumulatorStatusErr, "Error serializing 'accumulatorStatus' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccumulatorRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccumulatorRecord")
	}
	return nil
}

func (m *_BACnetAccumulatorRecord) isBACnetAccumulatorRecord() bool {
	return true
}

func (m *_BACnetAccumulatorRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
