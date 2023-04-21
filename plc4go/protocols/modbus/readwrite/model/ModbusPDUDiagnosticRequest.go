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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUDiagnosticRequest is the corresponding interface of ModbusPDUDiagnosticRequest
type ModbusPDUDiagnosticRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
}

// ModbusPDUDiagnosticRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUDiagnosticRequest.
// This is useful for switch cases.
type ModbusPDUDiagnosticRequestExactly interface {
	ModbusPDUDiagnosticRequest
	isModbusPDUDiagnosticRequest() bool
}

// _ModbusPDUDiagnosticRequest is the data-structure of this message
type _ModbusPDUDiagnosticRequest struct {
	*_ModbusPDU
	SubFunction uint16
	Data        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUDiagnosticRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUDiagnosticRequest) GetFunctionFlag() uint8 {
	return 0x08
}

func (m *_ModbusPDUDiagnosticRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUDiagnosticRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUDiagnosticRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUDiagnosticRequest) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *_ModbusPDUDiagnosticRequest) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUDiagnosticRequest factory function for _ModbusPDUDiagnosticRequest
func NewModbusPDUDiagnosticRequest(subFunction uint16, data uint16) *_ModbusPDUDiagnosticRequest {
	_result := &_ModbusPDUDiagnosticRequest{
		SubFunction: subFunction,
		Data:        data,
		_ModbusPDU:  NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUDiagnosticRequest(structType any) ModbusPDUDiagnosticRequest {
	if casted, ok := structType.(ModbusPDUDiagnosticRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUDiagnosticRequest) GetTypeName() string {
	return "ModbusPDUDiagnosticRequest"
}

func (m *_ModbusPDUDiagnosticRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subFunction)
	lengthInBits += 16

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUDiagnosticRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUDiagnosticRequestParse(theBytes []byte, response bool) (ModbusPDUDiagnosticRequest, error) {
	return ModbusPDUDiagnosticRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUDiagnosticRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUDiagnosticRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUDiagnosticRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subFunction)
	_subFunction, _subFunctionErr := readBuffer.ReadUint16("subFunction", 16)
	if _subFunctionErr != nil {
		return nil, errors.Wrap(_subFunctionErr, "Error parsing 'subFunction' field of ModbusPDUDiagnosticRequest")
	}
	subFunction := _subFunction

	// Simple Field (data)
	_data, _dataErr := readBuffer.ReadUint16("data", 16)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field of ModbusPDUDiagnosticRequest")
	}
	data := _data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUDiagnosticRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUDiagnosticRequest{
		_ModbusPDU:  &_ModbusPDU{},
		SubFunction: subFunction,
		Data:        data,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUDiagnosticRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUDiagnosticRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUDiagnosticRequest")
		}

		// Simple Field (subFunction)
		subFunction := uint16(m.GetSubFunction())
		_subFunctionErr := writeBuffer.WriteUint16("subFunction", 16, (subFunction))
		if _subFunctionErr != nil {
			return errors.Wrap(_subFunctionErr, "Error serializing 'subFunction' field")
		}

		// Simple Field (data)
		data := uint16(m.GetData())
		_dataErr := writeBuffer.WriteUint16("data", 16, (data))
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUDiagnosticRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUDiagnosticRequest) isModbusPDUDiagnosticRequest() bool {
	return true
}

func (m *_ModbusPDUDiagnosticRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
