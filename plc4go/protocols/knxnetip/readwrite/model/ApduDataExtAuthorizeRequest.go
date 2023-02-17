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


// ApduDataExtAuthorizeRequest is the corresponding interface of ApduDataExtAuthorizeRequest
type ApduDataExtAuthorizeRequest interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetLevel returns Level (property field)
	GetLevel() uint8
	// GetData returns Data (property field)
	GetData() []byte
}

// ApduDataExtAuthorizeRequestExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtAuthorizeRequest.
// This is useful for switch cases.
type ApduDataExtAuthorizeRequestExactly interface {
	ApduDataExtAuthorizeRequest
	isApduDataExtAuthorizeRequest() bool
}

// _ApduDataExtAuthorizeRequest is the data-structure of this message
type _ApduDataExtAuthorizeRequest struct {
	*_ApduDataExt
        Level uint8
        Data []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtAuthorizeRequest)  GetExtApciType() uint8 {
return 0x11}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtAuthorizeRequest) InitializeParent(parent ApduDataExt ) {}

func (m *_ApduDataExtAuthorizeRequest)  GetParent() ApduDataExt {
	return m._ApduDataExt
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtAuthorizeRequest) GetLevel() uint8 {
	return m.Level
}

func (m *_ApduDataExtAuthorizeRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewApduDataExtAuthorizeRequest factory function for _ApduDataExtAuthorizeRequest
func NewApduDataExtAuthorizeRequest( level uint8 , data []byte , length uint8 ) *_ApduDataExtAuthorizeRequest {
	_result := &_ApduDataExtAuthorizeRequest{
		Level: level,
		Data: data,
    	_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtAuthorizeRequest(structType interface{}) ApduDataExtAuthorizeRequest {
    if casted, ok := structType.(ApduDataExtAuthorizeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtAuthorizeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtAuthorizeRequest) GetTypeName() string {
	return "ApduDataExtAuthorizeRequest"
}

func (m *_ApduDataExtAuthorizeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (level)
	lengthInBits += 8;

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}


func (m *_ApduDataExtAuthorizeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtAuthorizeRequestParse(theBytes []byte, length uint8) (ApduDataExtAuthorizeRequest, error) {
	return ApduDataExtAuthorizeRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtAuthorizeRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtAuthorizeRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtAuthorizeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtAuthorizeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (level)
_level, _levelErr := readBuffer.ReadUint8("level", 8)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of ApduDataExtAuthorizeRequest")
	}
	level := _level
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(4))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ApduDataExtAuthorizeRequest")
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtAuthorizeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtAuthorizeRequest")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtAuthorizeRequest{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
		Level: level,
		Data: data,
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtAuthorizeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtAuthorizeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtAuthorizeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtAuthorizeRequest")
		}

	// Simple Field (level)
	level := uint8(m.GetLevel())
	_levelErr := writeBuffer.WriteUint8("level", 8, (level))
	if _levelErr != nil {
		return errors.Wrap(_levelErr, "Error serializing 'level' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

		if popErr := writeBuffer.PopContext("ApduDataExtAuthorizeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtAuthorizeRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ApduDataExtAuthorizeRequest) isApduDataExtAuthorizeRequest() bool {
	return true
}

func (m *_ApduDataExtAuthorizeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



