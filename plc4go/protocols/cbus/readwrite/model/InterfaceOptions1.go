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

// InterfaceOptions1 is the corresponding interface of InterfaceOptions1
type InterfaceOptions1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetIdmon returns Idmon (property field)
	GetIdmon() bool
	// GetMonitor returns Monitor (property field)
	GetMonitor() bool
	// GetSmart returns Smart (property field)
	GetSmart() bool
	// GetSrchk returns Srchk (property field)
	GetSrchk() bool
	// GetXonXoff returns XonXoff (property field)
	GetXonXoff() bool
	// GetConnect returns Connect (property field)
	GetConnect() bool
}

// InterfaceOptions1Exactly can be used when we want exactly this type and not a type which fulfills InterfaceOptions1.
// This is useful for switch cases.
type InterfaceOptions1Exactly interface {
	InterfaceOptions1
	isInterfaceOptions1() bool
}

// _InterfaceOptions1 is the data-structure of this message
type _InterfaceOptions1 struct {
	Idmon   bool
	Monitor bool
	Smart   bool
	Srchk   bool
	XonXoff bool
	Connect bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions1) GetIdmon() bool {
	return m.Idmon
}

func (m *_InterfaceOptions1) GetMonitor() bool {
	return m.Monitor
}

func (m *_InterfaceOptions1) GetSmart() bool {
	return m.Smart
}

func (m *_InterfaceOptions1) GetSrchk() bool {
	return m.Srchk
}

func (m *_InterfaceOptions1) GetXonXoff() bool {
	return m.XonXoff
}

func (m *_InterfaceOptions1) GetConnect() bool {
	return m.Connect
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewInterfaceOptions1 factory function for _InterfaceOptions1
func NewInterfaceOptions1(idmon bool, monitor bool, smart bool, srchk bool, xonXoff bool, connect bool) *_InterfaceOptions1 {
	return &_InterfaceOptions1{Idmon: idmon, Monitor: monitor, Smart: smart, Srchk: srchk, XonXoff: xonXoff, Connect: connect}
}

// Deprecated: use the interface for direct cast
func CastInterfaceOptions1(structType any) InterfaceOptions1 {
	if casted, ok := structType.(InterfaceOptions1); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions1); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions1) GetTypeName() string {
	return "InterfaceOptions1"
}

func (m *_InterfaceOptions1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (idmon)
	lengthInBits += 1

	// Simple field (monitor)
	lengthInBits += 1

	// Simple field (smart)
	lengthInBits += 1

	// Simple field (srchk)
	lengthInBits += 1

	// Simple field (xonXoff)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (connect)
	lengthInBits += 1

	return lengthInBits
}

func (m *_InterfaceOptions1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InterfaceOptions1Parse(theBytes []byte) (InterfaceOptions1, error) {
	return InterfaceOptions1ParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions1ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InterfaceOptions1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions1")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (idmon)
	_idmon, _idmonErr := readBuffer.ReadBit("idmon")
	if _idmonErr != nil {
		return nil, errors.Wrap(_idmonErr, "Error parsing 'idmon' field of InterfaceOptions1")
	}
	idmon := _idmon

	// Simple Field (monitor)
	_monitor, _monitorErr := readBuffer.ReadBit("monitor")
	if _monitorErr != nil {
		return nil, errors.Wrap(_monitorErr, "Error parsing 'monitor' field of InterfaceOptions1")
	}
	monitor := _monitor

	// Simple Field (smart)
	_smart, _smartErr := readBuffer.ReadBit("smart")
	if _smartErr != nil {
		return nil, errors.Wrap(_smartErr, "Error parsing 'smart' field of InterfaceOptions1")
	}
	smart := _smart

	// Simple Field (srchk)
	_srchk, _srchkErr := readBuffer.ReadBit("srchk")
	if _srchkErr != nil {
		return nil, errors.Wrap(_srchkErr, "Error parsing 'srchk' field of InterfaceOptions1")
	}
	srchk := _srchk

	// Simple Field (xonXoff)
	_xonXoff, _xonXoffErr := readBuffer.ReadBit("xonXoff")
	if _xonXoffErr != nil {
		return nil, errors.Wrap(_xonXoffErr, "Error parsing 'xonXoff' field of InterfaceOptions1")
	}
	xonXoff := _xonXoff

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions1")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (connect)
	_connect, _connectErr := readBuffer.ReadBit("connect")
	if _connectErr != nil {
		return nil, errors.Wrap(_connectErr, "Error parsing 'connect' field of InterfaceOptions1")
	}
	connect := _connect

	if closeErr := readBuffer.CloseContext("InterfaceOptions1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions1")
	}

	// Create the instance
	return &_InterfaceOptions1{
		Idmon:          idmon,
		Monitor:        monitor,
		Smart:          smart,
		Srchk:          srchk,
		XonXoff:        xonXoff,
		Connect:        connect,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
	}, nil
}

func (m *_InterfaceOptions1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("InterfaceOptions1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions1")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (idmon)
	idmon := bool(m.GetIdmon())
	_idmonErr := writeBuffer.WriteBit("idmon", (idmon))
	if _idmonErr != nil {
		return errors.Wrap(_idmonErr, "Error serializing 'idmon' field")
	}

	// Simple Field (monitor)
	monitor := bool(m.GetMonitor())
	_monitorErr := writeBuffer.WriteBit("monitor", (monitor))
	if _monitorErr != nil {
		return errors.Wrap(_monitorErr, "Error serializing 'monitor' field")
	}

	// Simple Field (smart)
	smart := bool(m.GetSmart())
	_smartErr := writeBuffer.WriteBit("smart", (smart))
	if _smartErr != nil {
		return errors.Wrap(_smartErr, "Error serializing 'smart' field")
	}

	// Simple Field (srchk)
	srchk := bool(m.GetSrchk())
	_srchkErr := writeBuffer.WriteBit("srchk", (srchk))
	if _srchkErr != nil {
		return errors.Wrap(_srchkErr, "Error serializing 'srchk' field")
	}

	// Simple Field (xonXoff)
	xonXoff := bool(m.GetXonXoff())
	_xonXoffErr := writeBuffer.WriteBit("xonXoff", (xonXoff))
	if _xonXoffErr != nil {
		return errors.Wrap(_xonXoffErr, "Error serializing 'xonXoff' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField1 != nil {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (connect)
	connect := bool(m.GetConnect())
	_connectErr := writeBuffer.WriteBit("connect", (connect))
	if _connectErr != nil {
		return errors.Wrap(_connectErr, "Error serializing 'connect' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions1")
	}
	return nil
}

func (m *_InterfaceOptions1) isInterfaceOptions1() bool {
	return true
}

func (m *_InterfaceOptions1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
