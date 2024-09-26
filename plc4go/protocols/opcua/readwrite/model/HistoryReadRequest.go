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

// HistoryReadRequest is the corresponding interface of HistoryReadRequest
type HistoryReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetHistoryReadDetails returns HistoryReadDetails (property field)
	GetHistoryReadDetails() ExtensionObject
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetReleaseContinuationPoints returns ReleaseContinuationPoints (property field)
	GetReleaseContinuationPoints() bool
	// GetNoOfNodesToRead returns NoOfNodesToRead (property field)
	GetNoOfNodesToRead() int32
	// GetNodesToRead returns NodesToRead (property field)
	GetNodesToRead() []ExtensionObjectDefinition
	// IsHistoryReadRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryReadRequest()
}

// _HistoryReadRequest is the data-structure of this message
type _HistoryReadRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader             ExtensionObjectDefinition
	HistoryReadDetails        ExtensionObject
	TimestampsToReturn        TimestampsToReturn
	ReleaseContinuationPoints bool
	NoOfNodesToRead           int32
	NodesToRead               []ExtensionObjectDefinition
	// Reserved Fields
	reservedField0 *uint8
}

var _ HistoryReadRequest = (*_HistoryReadRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryReadRequest)(nil)

// NewHistoryReadRequest factory function for _HistoryReadRequest
func NewHistoryReadRequest(requestHeader ExtensionObjectDefinition, historyReadDetails ExtensionObject, timestampsToReturn TimestampsToReturn, releaseContinuationPoints bool, noOfNodesToRead int32, nodesToRead []ExtensionObjectDefinition) *_HistoryReadRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for HistoryReadRequest must not be nil")
	}
	if historyReadDetails == nil {
		panic("historyReadDetails of type ExtensionObject for HistoryReadRequest must not be nil")
	}
	_result := &_HistoryReadRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		HistoryReadDetails:                historyReadDetails,
		TimestampsToReturn:                timestampsToReturn,
		ReleaseContinuationPoints:         releaseContinuationPoints,
		NoOfNodesToRead:                   noOfNodesToRead,
		NodesToRead:                       nodesToRead,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadRequest) GetIdentifier() string {
	return "664"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_HistoryReadRequest) GetHistoryReadDetails() ExtensionObject {
	return m.HistoryReadDetails
}

func (m *_HistoryReadRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_HistoryReadRequest) GetReleaseContinuationPoints() bool {
	return m.ReleaseContinuationPoints
}

func (m *_HistoryReadRequest) GetNoOfNodesToRead() int32 {
	return m.NoOfNodesToRead
}

func (m *_HistoryReadRequest) GetNodesToRead() []ExtensionObjectDefinition {
	return m.NodesToRead
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryReadRequest(structType any) HistoryReadRequest {
	if casted, ok := structType.(HistoryReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadRequest) GetTypeName() string {
	return "HistoryReadRequest"
}

func (m *_HistoryReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (historyReadDetails)
	lengthInBits += m.HistoryReadDetails.GetLengthInBits(ctx)

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (releaseContinuationPoints)
	lengthInBits += 1

	// Simple field (noOfNodesToRead)
	lengthInBits += 32

	// Array field
	if len(m.NodesToRead) > 0 {
		for _curItem, element := range m.NodesToRead {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToRead), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyReadRequest HistoryReadRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	historyReadDetails, err := ReadSimpleField[ExtensionObject](ctx, "historyReadDetails", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'historyReadDetails' field"))
	}
	m.HistoryReadDetails = historyReadDetails

	timestampsToReturn, err := ReadEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", ReadEnum(TimestampsToReturnByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampsToReturn' field"))
	}
	m.TimestampsToReturn = timestampsToReturn

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	releaseContinuationPoints, err := ReadSimpleField(ctx, "releaseContinuationPoints", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'releaseContinuationPoints' field"))
	}
	m.ReleaseContinuationPoints = releaseContinuationPoints

	noOfNodesToRead, err := ReadSimpleField(ctx, "noOfNodesToRead", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToRead' field"))
	}
	m.NoOfNodesToRead = noOfNodesToRead

	nodesToRead, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToRead", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("637")), readBuffer), uint64(noOfNodesToRead))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToRead' field"))
	}
	m.NodesToRead = nodesToRead

	if closeErr := readBuffer.CloseContext("HistoryReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadRequest")
	}

	return m, nil
}

func (m *_HistoryReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "historyReadDetails", m.GetHistoryReadDetails(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'historyReadDetails' field")
		}

		if err := WriteSimpleEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", m.GetTimestampsToReturn(), WriteEnum[TimestampsToReturn, uint32](TimestampsToReturn.GetValue, TimestampsToReturn.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'timestampsToReturn' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "releaseContinuationPoints", m.GetReleaseContinuationPoints(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'releaseContinuationPoints' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToRead", m.GetNoOfNodesToRead(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToRead' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToRead", m.GetNodesToRead(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToRead' field")
		}

		if popErr := writeBuffer.PopContext("HistoryReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadRequest) IsHistoryReadRequest() {}

func (m *_HistoryReadRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryReadRequest) deepCopy() *_HistoryReadRequest {
	if m == nil {
		return nil
	}
	_HistoryReadRequestCopy := &_HistoryReadRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.HistoryReadDetails.DeepCopy().(ExtensionObject),
		m.TimestampsToReturn,
		m.ReleaseContinuationPoints,
		m.NoOfNodesToRead,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.NodesToRead),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryReadRequestCopy
}

func (m *_HistoryReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
