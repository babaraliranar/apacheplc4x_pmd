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

// ApduDataExtNetworkParameterResponse is the corresponding interface of ApduDataExtNetworkParameterResponse
type ApduDataExtNetworkParameterResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtNetworkParameterResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtNetworkParameterResponse.
// This is useful for switch cases.
type ApduDataExtNetworkParameterResponseExactly interface {
	ApduDataExtNetworkParameterResponse
	isApduDataExtNetworkParameterResponse() bool
}

// _ApduDataExtNetworkParameterResponse is the data-structure of this message
type _ApduDataExtNetworkParameterResponse struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtNetworkParameterResponse) GetExtApciType() uint8 {
	return 0x1B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtNetworkParameterResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtNetworkParameterResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtNetworkParameterResponse factory function for _ApduDataExtNetworkParameterResponse
func NewApduDataExtNetworkParameterResponse(length uint8) *_ApduDataExtNetworkParameterResponse {
	_result := &_ApduDataExtNetworkParameterResponse{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtNetworkParameterResponse(structType interface{}) ApduDataExtNetworkParameterResponse {
	if casted, ok := structType.(ApduDataExtNetworkParameterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtNetworkParameterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtNetworkParameterResponse) GetTypeName() string {
	return "ApduDataExtNetworkParameterResponse"
}

func (m *_ApduDataExtNetworkParameterResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtNetworkParameterResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtNetworkParameterResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtNetworkParameterResponseParse(theBytes []byte, length uint8) (ApduDataExtNetworkParameterResponse, error) {
	return ApduDataExtNetworkParameterResponseParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), length) // TODO: get endianness from mspec
}

func ApduDataExtNetworkParameterResponseParseWithBuffer(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtNetworkParameterResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtNetworkParameterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtNetworkParameterResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtNetworkParameterResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtNetworkParameterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtNetworkParameterResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtNetworkParameterResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtNetworkParameterResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtNetworkParameterResponse) isApduDataExtNetworkParameterResponse() bool {
	return true
}

func (m *_ApduDataExtNetworkParameterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
