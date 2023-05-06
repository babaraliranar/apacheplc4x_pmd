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

// IPAddress is the corresponding interface of IPAddress
type IPAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAddr returns Addr (property field)
	GetAddr() []byte
}

// IPAddressExactly can be used when we want exactly this type and not a type which fulfills IPAddress.
// This is useful for switch cases.
type IPAddressExactly interface {
	IPAddress
	isIPAddress() bool
}

// _IPAddress is the data-structure of this message
type _IPAddress struct {
	Addr []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IPAddress) GetAddr() []byte {
	return m.Addr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIPAddress factory function for _IPAddress
func NewIPAddress(addr []byte) *_IPAddress {
	return &_IPAddress{Addr: addr}
}

// Deprecated: use the interface for direct cast
func CastIPAddress(structType any) IPAddress {
	if casted, ok := structType.(IPAddress); ok {
		return casted
	}
	if casted, ok := structType.(*IPAddress); ok {
		return *casted
	}
	return nil
}

func (m *_IPAddress) GetTypeName() string {
	return "IPAddress"
}

func (m *_IPAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *_IPAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IPAddressParse(theBytes []byte) (IPAddress, error) {
	return IPAddressParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func IPAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IPAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IPAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IPAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (addr)
	numberOfBytesaddr := int(uint16(4))
	addr, _readArrayErr := readBuffer.ReadByteArray("addr", numberOfBytesaddr)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'addr' field of IPAddress")
	}

	if closeErr := readBuffer.CloseContext("IPAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IPAddress")
	}

	// Create the instance
	return &_IPAddress{
		Addr: addr,
	}, nil
}

func (m *_IPAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IPAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("IPAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IPAddress")
	}

	// Array Field (addr)
	// Byte Array field (addr)
	if err := writeBuffer.WriteByteArray("addr", m.GetAddr()); err != nil {
		return errors.Wrap(err, "Error serializing 'addr' field")
	}

	if popErr := writeBuffer.PopContext("IPAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IPAddress")
	}
	return nil
}

func (m *_IPAddress) isIPAddress() bool {
	return true
}

func (m *_IPAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
