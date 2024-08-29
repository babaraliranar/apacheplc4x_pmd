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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MonitoringParameters is the corresponding interface of MonitoringParameters
type MonitoringParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetClientHandle returns ClientHandle (property field)
	GetClientHandle() uint32
	// GetSamplingInterval returns SamplingInterval (property field)
	GetSamplingInterval() float64
	// GetFilter returns Filter (property field)
	GetFilter() ExtensionObject
	// GetQueueSize returns QueueSize (property field)
	GetQueueSize() uint32
	// GetDiscardOldest returns DiscardOldest (property field)
	GetDiscardOldest() bool
}

// MonitoringParametersExactly can be used when we want exactly this type and not a type which fulfills MonitoringParameters.
// This is useful for switch cases.
type MonitoringParametersExactly interface {
	MonitoringParameters
	isMonitoringParameters() bool
}

// _MonitoringParameters is the data-structure of this message
type _MonitoringParameters struct {
	*_ExtensionObjectDefinition
	ClientHandle     uint32
	SamplingInterval float64
	Filter           ExtensionObject
	QueueSize        uint32
	DiscardOldest    bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoringParameters) GetIdentifier() string {
	return "742"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoringParameters) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoringParameters) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoringParameters) GetClientHandle() uint32 {
	return m.ClientHandle
}

func (m *_MonitoringParameters) GetSamplingInterval() float64 {
	return m.SamplingInterval
}

func (m *_MonitoringParameters) GetFilter() ExtensionObject {
	return m.Filter
}

func (m *_MonitoringParameters) GetQueueSize() uint32 {
	return m.QueueSize
}

func (m *_MonitoringParameters) GetDiscardOldest() bool {
	return m.DiscardOldest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoringParameters factory function for _MonitoringParameters
func NewMonitoringParameters(clientHandle uint32, samplingInterval float64, filter ExtensionObject, queueSize uint32, discardOldest bool) *_MonitoringParameters {
	_result := &_MonitoringParameters{
		ClientHandle:               clientHandle,
		SamplingInterval:           samplingInterval,
		Filter:                     filter,
		QueueSize:                  queueSize,
		DiscardOldest:              discardOldest,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoringParameters(structType any) MonitoringParameters {
	if casted, ok := structType.(MonitoringParameters); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoringParameters); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoringParameters) GetTypeName() string {
	return "MonitoringParameters"
}

func (m *_MonitoringParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (clientHandle)
	lengthInBits += 32

	// Simple field (samplingInterval)
	lengthInBits += 64

	// Simple field (filter)
	lengthInBits += m.Filter.GetLengthInBits(ctx)

	// Simple field (queueSize)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (discardOldest)
	lengthInBits += 1

	return lengthInBits
}

func (m *_MonitoringParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoringParametersParse(ctx context.Context, theBytes []byte, identifier string) (MonitoringParameters, error) {
	return MonitoringParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoringParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoringParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoringParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoringParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (clientHandle)
	_clientHandle, _clientHandleErr := readBuffer.ReadUint32("clientHandle", 32)
	if _clientHandleErr != nil {
		return nil, errors.Wrap(_clientHandleErr, "Error parsing 'clientHandle' field of MonitoringParameters")
	}
	clientHandle := _clientHandle

	// Simple Field (samplingInterval)
	_samplingInterval, _samplingIntervalErr := readBuffer.ReadFloat64("samplingInterval", 64)
	if _samplingIntervalErr != nil {
		return nil, errors.Wrap(_samplingIntervalErr, "Error parsing 'samplingInterval' field of MonitoringParameters")
	}
	samplingInterval := _samplingInterval

	// Simple Field (filter)
	if pullErr := readBuffer.PullContext("filter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for filter")
	}
	_filter, _filterErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _filterErr != nil {
		return nil, errors.Wrap(_filterErr, "Error parsing 'filter' field of MonitoringParameters")
	}
	filter := _filter.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("filter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for filter")
	}

	// Simple Field (queueSize)
	_queueSize, _queueSizeErr := readBuffer.ReadUint32("queueSize", 32)
	if _queueSizeErr != nil {
		return nil, errors.Wrap(_queueSizeErr, "Error parsing 'queueSize' field of MonitoringParameters")
	}
	queueSize := _queueSize

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of MonitoringParameters")
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

	// Simple Field (discardOldest)
	_discardOldest, _discardOldestErr := readBuffer.ReadBit("discardOldest")
	if _discardOldestErr != nil {
		return nil, errors.Wrap(_discardOldestErr, "Error parsing 'discardOldest' field of MonitoringParameters")
	}
	discardOldest := _discardOldest

	if closeErr := readBuffer.CloseContext("MonitoringParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoringParameters")
	}

	// Create a partially initialized instance
	_child := &_MonitoringParameters{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ClientHandle:               clientHandle,
		SamplingInterval:           samplingInterval,
		Filter:                     filter,
		QueueSize:                  queueSize,
		DiscardOldest:              discardOldest,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoringParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoringParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoringParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoringParameters")
		}

		// Simple Field (clientHandle)
		clientHandle := uint32(m.GetClientHandle())
		_clientHandleErr := writeBuffer.WriteUint32("clientHandle", 32, uint32((clientHandle)))
		if _clientHandleErr != nil {
			return errors.Wrap(_clientHandleErr, "Error serializing 'clientHandle' field")
		}

		// Simple Field (samplingInterval)
		samplingInterval := float64(m.GetSamplingInterval())
		_samplingIntervalErr := writeBuffer.WriteFloat64("samplingInterval", 64, (samplingInterval))
		if _samplingIntervalErr != nil {
			return errors.Wrap(_samplingIntervalErr, "Error serializing 'samplingInterval' field")
		}

		// Simple Field (filter)
		if pushErr := writeBuffer.PushContext("filter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filter")
		}
		_filterErr := writeBuffer.WriteSerializable(ctx, m.GetFilter())
		if popErr := writeBuffer.PopContext("filter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filter")
		}
		if _filterErr != nil {
			return errors.Wrap(_filterErr, "Error serializing 'filter' field")
		}

		// Simple Field (queueSize)
		queueSize := uint32(m.GetQueueSize())
		_queueSizeErr := writeBuffer.WriteUint32("queueSize", 32, uint32((queueSize)))
		if _queueSizeErr != nil {
			return errors.Wrap(_queueSizeErr, "Error serializing 'queueSize' field")
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
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (discardOldest)
		discardOldest := bool(m.GetDiscardOldest())
		_discardOldestErr := writeBuffer.WriteBit("discardOldest", (discardOldest))
		if _discardOldestErr != nil {
			return errors.Wrap(_discardOldestErr, "Error serializing 'discardOldest' field")
		}

		if popErr := writeBuffer.PopContext("MonitoringParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoringParameters")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoringParameters) isMonitoringParameters() bool {
	return true
}

func (m *_MonitoringParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
