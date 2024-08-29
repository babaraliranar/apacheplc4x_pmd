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
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetSubscription returns Subscription (property field)
	GetSubscription() uint8
	// GetMagicKey returns MagicKey (property field)
	GetMagicKey() string
	// GetAlarmtype returns Alarmtype (property field)
	GetAlarmtype() *AlarmStateType
	// GetReserve returns Reserve (property field)
	GetReserve() *uint8
}

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestExactly interface {
	S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
	isS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest() bool
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest struct {
	*_S7PayloadUserDataItem
	Subscription uint8
	MagicKey     string
	Alarmtype    *AlarmStateType
	Reserve      *uint8
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetCpuSubfunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetSubscription() uint8 {
	return m.Subscription
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetMagicKey() string {
	return m.MagicKey
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetAlarmtype() *AlarmStateType {
	return m.Alarmtype
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetReserve() *uint8 {
	return m.Reserve
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest(subscription uint8, magicKey string, alarmtype *AlarmStateType, reserve *uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest{
		Subscription:           subscription,
		MagicKey:               magicKey,
		Alarmtype:              alarmtype,
		Reserve:                reserve,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest(structType any) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subscription)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (magicKey)
	lengthInBits += 64

	// Optional Field (alarmtype)
	if m.Alarmtype != nil {
		lengthInBits += 8
	}

	// Optional Field (reserve)
	if m.Reserve != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest, error) {
	return S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscription)
	_subscription, _subscriptionErr := readBuffer.ReadUint8("subscription", 8)
	if _subscriptionErr != nil {
		return nil, errors.Wrap(_subscriptionErr, "Error parsing 'subscription' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
	}
	subscription := _subscription

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (magicKey)
	_magicKey, _magicKeyErr := readBuffer.ReadString("magicKey", uint32(64), "UTF-8")
	if _magicKeyErr != nil {
		return nil, errors.Wrap(_magicKeyErr, "Error parsing 'magicKey' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
	}
	magicKey := _magicKey

	// Optional Field (alarmtype) (Can be skipped, if a given expression evaluates to false)
	var alarmtype *AlarmStateType = nil
	if bool((subscription) >= (128)) {
		if pullErr := readBuffer.PullContext("alarmtype"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for alarmtype")
		}
		currentPos = positionAware.GetPos()
		_val, _err := AlarmStateTypeParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alarmtype' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		default:
			alarmtype = &_val
		}
		if closeErr := readBuffer.CloseContext("alarmtype"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for alarmtype")
		}
	}

	// Optional Field (reserve) (Can be skipped, if a given expression evaluates to false)
	var reserve *uint8 = nil
	if bool((subscription) >= (128)) {
		currentPos = positionAware.GetPos()
		_val, _err := readBuffer.ReadUint8("reserve", 8)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'reserve' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		default:
			reserve = &_val
		}
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Subscription:           subscription,
		MagicKey:               magicKey,
		Alarmtype:              alarmtype,
		Reserve:                reserve,
		reservedField0:         reservedField0,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		}

		// Simple Field (subscription)
		subscription := uint8(m.GetSubscription())
		_subscriptionErr := writeBuffer.WriteUint8("subscription", 8, uint8((subscription)))
		if _subscriptionErr != nil {
			return errors.Wrap(_subscriptionErr, "Error serializing 'subscription' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (magicKey)
		magicKey := string(m.GetMagicKey())
		_magicKeyErr := writeBuffer.WriteString("magicKey", uint32(64), "UTF-8", (magicKey))
		if _magicKeyErr != nil {
			return errors.Wrap(_magicKeyErr, "Error serializing 'magicKey' field")
		}

		// Optional Field (alarmtype) (Can be skipped, if the value is null)
		var alarmtype *AlarmStateType = nil
		if m.GetAlarmtype() != nil {
			if pushErr := writeBuffer.PushContext("alarmtype"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for alarmtype")
			}
			alarmtype = m.GetAlarmtype()
			_alarmtypeErr := writeBuffer.WriteSerializable(ctx, alarmtype)
			if popErr := writeBuffer.PopContext("alarmtype"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for alarmtype")
			}
			if _alarmtypeErr != nil {
				return errors.Wrap(_alarmtypeErr, "Error serializing 'alarmtype' field")
			}
		}

		// Optional Field (reserve) (Can be skipped, if the value is null)
		var reserve *uint8 = nil
		if m.GetReserve() != nil {
			reserve = m.GetReserve()
			_reserveErr := writeBuffer.WriteUint8("reserve", 8, uint8(*(reserve)))
			if _reserveErr != nil {
				return errors.Wrap(_reserveErr, "Error serializing 'reserve' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) isS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
