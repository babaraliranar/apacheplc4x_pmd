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

// NetworkProtocolControlInformation is the corresponding interface of NetworkProtocolControlInformation
type NetworkProtocolControlInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStackCounter returns StackCounter (property field)
	GetStackCounter() uint8
	// GetStackDepth returns StackDepth (property field)
	GetStackDepth() uint8
}

// NetworkProtocolControlInformationExactly can be used when we want exactly this type and not a type which fulfills NetworkProtocolControlInformation.
// This is useful for switch cases.
type NetworkProtocolControlInformationExactly interface {
	NetworkProtocolControlInformation
	isNetworkProtocolControlInformation() bool
}

// _NetworkProtocolControlInformation is the data-structure of this message
type _NetworkProtocolControlInformation struct {
	StackCounter uint8
	StackDepth   uint8
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkProtocolControlInformation) GetStackCounter() uint8 {
	return m.StackCounter
}

func (m *_NetworkProtocolControlInformation) GetStackDepth() uint8 {
	return m.StackDepth
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkProtocolControlInformation factory function for _NetworkProtocolControlInformation
func NewNetworkProtocolControlInformation(stackCounter uint8, stackDepth uint8) *_NetworkProtocolControlInformation {
	return &_NetworkProtocolControlInformation{StackCounter: stackCounter, StackDepth: stackDepth}
}

// Deprecated: use the interface for direct cast
func CastNetworkProtocolControlInformation(structType any) NetworkProtocolControlInformation {
	if casted, ok := structType.(NetworkProtocolControlInformation); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkProtocolControlInformation); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkProtocolControlInformation) GetTypeName() string {
	return "NetworkProtocolControlInformation"
}

func (m *_NetworkProtocolControlInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (stackCounter)
	lengthInBits += 3

	// Simple field (stackDepth)
	lengthInBits += 3

	return lengthInBits
}

func (m *_NetworkProtocolControlInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NetworkProtocolControlInformationParse(theBytes []byte) (NetworkProtocolControlInformation, error) {
	return NetworkProtocolControlInformationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func NetworkProtocolControlInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkProtocolControlInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkProtocolControlInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkProtocolControlInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of NetworkProtocolControlInformation")
		}
		if reserved != uint8(0x0) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (stackCounter)
	_stackCounter, _stackCounterErr := readBuffer.ReadUint8("stackCounter", 3)
	if _stackCounterErr != nil {
		return nil, errors.Wrap(_stackCounterErr, "Error parsing 'stackCounter' field of NetworkProtocolControlInformation")
	}
	stackCounter := _stackCounter

	// Simple Field (stackDepth)
	_stackDepth, _stackDepthErr := readBuffer.ReadUint8("stackDepth", 3)
	if _stackDepthErr != nil {
		return nil, errors.Wrap(_stackDepthErr, "Error parsing 'stackDepth' field of NetworkProtocolControlInformation")
	}
	stackDepth := _stackDepth

	if closeErr := readBuffer.CloseContext("NetworkProtocolControlInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkProtocolControlInformation")
	}

	// Create the instance
	return &_NetworkProtocolControlInformation{
		StackCounter:   stackCounter,
		StackDepth:     stackDepth,
		reservedField0: reservedField0,
	}, nil
}

func (m *_NetworkProtocolControlInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkProtocolControlInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NetworkProtocolControlInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkProtocolControlInformation")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x0)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 2, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (stackCounter)
	stackCounter := uint8(m.GetStackCounter())
	_stackCounterErr := writeBuffer.WriteUint8("stackCounter", 3, (stackCounter))
	if _stackCounterErr != nil {
		return errors.Wrap(_stackCounterErr, "Error serializing 'stackCounter' field")
	}

	// Simple Field (stackDepth)
	stackDepth := uint8(m.GetStackDepth())
	_stackDepthErr := writeBuffer.WriteUint8("stackDepth", 3, (stackDepth))
	if _stackDepthErr != nil {
		return errors.Wrap(_stackDepthErr, "Error serializing 'stackDepth' field")
	}

	if popErr := writeBuffer.PopContext("NetworkProtocolControlInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkProtocolControlInformation")
	}
	return nil
}

func (m *_NetworkProtocolControlInformation) isNetworkProtocolControlInformation() bool {
	return true
}

func (m *_NetworkProtocolControlInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
