//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type APDUAbort struct {
	Server           bool
	OriginalInvokeId uint8
	AbortReason      uint8
	Parent           *APDU
}

// The corresponding interface
type IAPDUAbort interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUAbort) ApduType() uint8 {
	return 0x7
}

func (m *APDUAbort) InitializeParent(parent *APDU) {
}

func NewAPDUAbort(server bool, originalInvokeId uint8, abortReason uint8) *APDU {
	child := &APDUAbort{
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      abortReason,
		Parent:           NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUAbort(structType interface{}) *APDUAbort {
	castFunc := func(typ interface{}) *APDUAbort {
		if casted, ok := typ.(APDUAbort); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUAbort); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUAbort(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUAbort(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUAbort) GetTypeName() string {
	return "APDUAbort"
}

func (m *APDUAbort) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUAbort) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (abortReason)
	lengthInBits += 8

	return lengthInBits
}

func (m *APDUAbort) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUAbortParse(readBuffer utils.ReadBuffer) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUAbort"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 3)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (server)
	server, _serverErr := readBuffer.ReadBit("server")
	if _serverErr != nil {
		return nil, errors.Wrap(_serverErr, "Error parsing 'server' field")
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}

	// Simple Field (abortReason)
	abortReason, _abortReasonErr := readBuffer.ReadUint8("abortReason", 8)
	if _abortReasonErr != nil {
		return nil, errors.Wrap(_abortReasonErr, "Error parsing 'abortReason' field")
	}

	if closeErr := readBuffer.CloseContext("APDUAbort"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUAbort{
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      abortReason,
		Parent:           &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUAbort) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUAbort"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 3, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (server)
		server := bool(m.Server)
		_serverErr := writeBuffer.WriteBit("server", (server))
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (abortReason)
		abortReason := uint8(m.AbortReason)
		_abortReasonErr := writeBuffer.WriteUint8("abortReason", 8, (abortReason))
		if _abortReasonErr != nil {
			return errors.Wrap(_abortReasonErr, "Error serializing 'abortReason' field")
		}

		if popErr := writeBuffer.PopContext("APDUAbort"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
