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
)

	// Code generated by code-generation. DO NOT EDIT.


// LDataExtended is the corresponding interface of LDataExtended
type LDataExtended interface {
	utils.LengthAware
	utils.Serializable
	LDataFrame
	// GetGroupAddress returns GroupAddress (property field)
	GetGroupAddress() bool
	// GetHopCount returns HopCount (property field)
	GetHopCount() uint8
	// GetExtendedFrameFormat returns ExtendedFrameFormat (property field)
	GetExtendedFrameFormat() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() KnxAddress
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() []byte
	// GetApdu returns Apdu (property field)
	GetApdu() Apdu
}

// LDataExtendedExactly can be used when we want exactly this type and not a type which fulfills LDataExtended.
// This is useful for switch cases.
type LDataExtendedExactly interface {
	LDataExtended
	isLDataExtended() bool
}

// _LDataExtended is the data-structure of this message
type _LDataExtended struct {
	*_LDataFrame
        GroupAddress bool
        HopCount uint8
        ExtendedFrameFormat uint8
        SourceAddress KnxAddress
        DestinationAddress []byte
        Apdu Apdu
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataExtended)  GetNotAckFrame() bool {
return bool(true)}

func (m *_LDataExtended)  GetPolling() bool {
return bool(false)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataExtended) InitializeParent(parent LDataFrame , frameType bool , notRepeated bool , priority CEMIPriority , acknowledgeRequested bool , errorFlag bool ) {	m.FrameType = frameType
	m.NotRepeated = notRepeated
	m.Priority = priority
	m.AcknowledgeRequested = acknowledgeRequested
	m.ErrorFlag = errorFlag
}

func (m *_LDataExtended)  GetParent() LDataFrame {
	return m._LDataFrame
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataExtended) GetGroupAddress() bool {
	return m.GroupAddress
}

func (m *_LDataExtended) GetHopCount() uint8 {
	return m.HopCount
}

func (m *_LDataExtended) GetExtendedFrameFormat() uint8 {
	return m.ExtendedFrameFormat
}

func (m *_LDataExtended) GetSourceAddress() KnxAddress {
	return m.SourceAddress
}

func (m *_LDataExtended) GetDestinationAddress() []byte {
	return m.DestinationAddress
}

func (m *_LDataExtended) GetApdu() Apdu {
	return m.Apdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewLDataExtended factory function for _LDataExtended
func NewLDataExtended( groupAddress bool , hopCount uint8 , extendedFrameFormat uint8 , sourceAddress KnxAddress , destinationAddress []byte , apdu Apdu , frameType bool , notRepeated bool , priority CEMIPriority , acknowledgeRequested bool , errorFlag bool ) *_LDataExtended {
	_result := &_LDataExtended{
		GroupAddress: groupAddress,
		HopCount: hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress: sourceAddress,
		DestinationAddress: destinationAddress,
		Apdu: apdu,
    	_LDataFrame: NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	_result._LDataFrame._LDataFrameChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLDataExtended(structType interface{}) LDataExtended {
    if casted, ok := structType.(LDataExtended); ok {
		return casted
	}
	if casted, ok := structType.(*LDataExtended); ok {
		return *casted
	}
	return nil
}

func (m *_LDataExtended) GetTypeName() string {
	return "LDataExtended"
}

func (m *_LDataExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (groupAddress)
	lengthInBits += 1;

	// Simple field (hopCount)
	lengthInBits += 3;

	// Simple field (extendedFrameFormat)
	lengthInBits += 4;

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.GetLengthInBits(ctx)

	// Array field
	if len(m.DestinationAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.DestinationAddress))
	}

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (apdu)
	lengthInBits += m.Apdu.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_LDataExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LDataExtendedParse(theBytes []byte) (LDataExtended, error) {
	return LDataExtendedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func LDataExtendedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LDataExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (groupAddress)
_groupAddress, _groupAddressErr := readBuffer.ReadBit("groupAddress")
	if _groupAddressErr != nil {
		return nil, errors.Wrap(_groupAddressErr, "Error parsing 'groupAddress' field of LDataExtended")
	}
	groupAddress := _groupAddress

	// Simple Field (hopCount)
_hopCount, _hopCountErr := readBuffer.ReadUint8("hopCount", 3)
	if _hopCountErr != nil {
		return nil, errors.Wrap(_hopCountErr, "Error parsing 'hopCount' field of LDataExtended")
	}
	hopCount := _hopCount

	// Simple Field (extendedFrameFormat)
_extendedFrameFormat, _extendedFrameFormatErr := readBuffer.ReadUint8("extendedFrameFormat", 4)
	if _extendedFrameFormatErr != nil {
		return nil, errors.Wrap(_extendedFrameFormatErr, "Error parsing 'extendedFrameFormat' field of LDataExtended")
	}
	extendedFrameFormat := _extendedFrameFormat

	// Simple Field (sourceAddress)
	if pullErr := readBuffer.PullContext("sourceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceAddress")
	}
_sourceAddress, _sourceAddressErr := KnxAddressParseWithBuffer(ctx, readBuffer)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field of LDataExtended")
	}
	sourceAddress := _sourceAddress.(KnxAddress)
	if closeErr := readBuffer.CloseContext("sourceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceAddress")
	}
	// Byte Array field (destinationAddress)
	numberOfBytesdestinationAddress := int(uint16(2))
	destinationAddress, _readArrayErr := readBuffer.ReadByteArray("destinationAddress", numberOfBytesdestinationAddress)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'destinationAddress' field of LDataExtended")
	}

	// Implicit Field (dataLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint8("dataLength", 8)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field of LDataExtended")
	}

	// Simple Field (apdu)
	if pullErr := readBuffer.PullContext("apdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for apdu")
	}
_apdu, _apduErr := ApduParseWithBuffer(ctx, readBuffer , uint8( dataLength ) )
	if _apduErr != nil {
		return nil, errors.Wrap(_apduErr, "Error parsing 'apdu' field of LDataExtended")
	}
	apdu := _apdu.(Apdu)
	if closeErr := readBuffer.CloseContext("apdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for apdu")
	}

	if closeErr := readBuffer.CloseContext("LDataExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataExtended")
	}

	// Create a partially initialized instance
	_child := &_LDataExtended{
		_LDataFrame: &_LDataFrame{
		},
		GroupAddress: groupAddress,
		HopCount: hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress: sourceAddress,
		DestinationAddress: destinationAddress,
		Apdu: apdu,
	}
	_child._LDataFrame._LDataFrameChildRequirements = _child
	return _child, nil
}

func (m *_LDataExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataExtended")
		}

	// Simple Field (groupAddress)
	groupAddress := bool(m.GetGroupAddress())
	_groupAddressErr := writeBuffer.WriteBit("groupAddress", (groupAddress))
	if _groupAddressErr != nil {
		return errors.Wrap(_groupAddressErr, "Error serializing 'groupAddress' field")
	}

	// Simple Field (hopCount)
	hopCount := uint8(m.GetHopCount())
	_hopCountErr := writeBuffer.WriteUint8("hopCount", 3, (hopCount))
	if _hopCountErr != nil {
		return errors.Wrap(_hopCountErr, "Error serializing 'hopCount' field")
	}

	// Simple Field (extendedFrameFormat)
	extendedFrameFormat := uint8(m.GetExtendedFrameFormat())
	_extendedFrameFormatErr := writeBuffer.WriteUint8("extendedFrameFormat", 4, (extendedFrameFormat))
	if _extendedFrameFormatErr != nil {
		return errors.Wrap(_extendedFrameFormatErr, "Error serializing 'extendedFrameFormat' field")
	}

	// Simple Field (sourceAddress)
	if pushErr := writeBuffer.PushContext("sourceAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for sourceAddress")
	}
	_sourceAddressErr := writeBuffer.WriteSerializable(ctx, m.GetSourceAddress())
	if popErr := writeBuffer.PopContext("sourceAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for sourceAddress")
	}
	if _sourceAddressErr != nil {
		return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
	}

	// Array Field (destinationAddress)
	// Byte Array field (destinationAddress)
	if err := writeBuffer.WriteByteArray("destinationAddress", m.GetDestinationAddress()); err != nil {
		return errors.Wrap(err, "Error serializing 'destinationAddress' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint8(uint8(m.GetApdu().GetLengthInBytes(ctx)) - uint8(uint8(1)))
	_dataLengthErr := writeBuffer.WriteUint8("dataLength", 8, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Simple Field (apdu)
	if pushErr := writeBuffer.PushContext("apdu"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for apdu")
	}
	_apduErr := writeBuffer.WriteSerializable(ctx, m.GetApdu())
	if popErr := writeBuffer.PopContext("apdu"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for apdu")
	}
	if _apduErr != nil {
		return errors.Wrap(_apduErr, "Error serializing 'apdu' field")
	}

		if popErr := writeBuffer.PopContext("LDataExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataExtended")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_LDataExtended) isLDataExtended() bool {
	return true
}

func (m *_LDataExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



