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

// EndpointConfiguration is the corresponding interface of EndpointConfiguration
type EndpointConfiguration interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetOperationTimeout returns OperationTimeout (property field)
	GetOperationTimeout() int32
	// GetUseBinaryEncoding returns UseBinaryEncoding (property field)
	GetUseBinaryEncoding() bool
	// GetMaxStringLength returns MaxStringLength (property field)
	GetMaxStringLength() int32
	// GetMaxByteStringLength returns MaxByteStringLength (property field)
	GetMaxByteStringLength() int32
	// GetMaxArrayLength returns MaxArrayLength (property field)
	GetMaxArrayLength() int32
	// GetMaxMessageSize returns MaxMessageSize (property field)
	GetMaxMessageSize() int32
	// GetMaxBufferSize returns MaxBufferSize (property field)
	GetMaxBufferSize() int32
	// GetChannelLifetime returns ChannelLifetime (property field)
	GetChannelLifetime() int32
	// GetSecurityTokenLifetime returns SecurityTokenLifetime (property field)
	GetSecurityTokenLifetime() int32
}

// EndpointConfigurationExactly can be used when we want exactly this type and not a type which fulfills EndpointConfiguration.
// This is useful for switch cases.
type EndpointConfigurationExactly interface {
	EndpointConfiguration
	isEndpointConfiguration() bool
}

// _EndpointConfiguration is the data-structure of this message
type _EndpointConfiguration struct {
	*_ExtensionObjectDefinition
	OperationTimeout      int32
	UseBinaryEncoding     bool
	MaxStringLength       int32
	MaxByteStringLength   int32
	MaxArrayLength        int32
	MaxMessageSize        int32
	MaxBufferSize         int32
	ChannelLifetime       int32
	SecurityTokenLifetime int32
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EndpointConfiguration) GetIdentifier() string {
	return "333"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EndpointConfiguration) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_EndpointConfiguration) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EndpointConfiguration) GetOperationTimeout() int32 {
	return m.OperationTimeout
}

func (m *_EndpointConfiguration) GetUseBinaryEncoding() bool {
	return m.UseBinaryEncoding
}

func (m *_EndpointConfiguration) GetMaxStringLength() int32 {
	return m.MaxStringLength
}

func (m *_EndpointConfiguration) GetMaxByteStringLength() int32 {
	return m.MaxByteStringLength
}

func (m *_EndpointConfiguration) GetMaxArrayLength() int32 {
	return m.MaxArrayLength
}

func (m *_EndpointConfiguration) GetMaxMessageSize() int32 {
	return m.MaxMessageSize
}

func (m *_EndpointConfiguration) GetMaxBufferSize() int32 {
	return m.MaxBufferSize
}

func (m *_EndpointConfiguration) GetChannelLifetime() int32 {
	return m.ChannelLifetime
}

func (m *_EndpointConfiguration) GetSecurityTokenLifetime() int32 {
	return m.SecurityTokenLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEndpointConfiguration factory function for _EndpointConfiguration
func NewEndpointConfiguration(operationTimeout int32, useBinaryEncoding bool, maxStringLength int32, maxByteStringLength int32, maxArrayLength int32, maxMessageSize int32, maxBufferSize int32, channelLifetime int32, securityTokenLifetime int32) *_EndpointConfiguration {
	_result := &_EndpointConfiguration{
		OperationTimeout:           operationTimeout,
		UseBinaryEncoding:          useBinaryEncoding,
		MaxStringLength:            maxStringLength,
		MaxByteStringLength:        maxByteStringLength,
		MaxArrayLength:             maxArrayLength,
		MaxMessageSize:             maxMessageSize,
		MaxBufferSize:              maxBufferSize,
		ChannelLifetime:            channelLifetime,
		SecurityTokenLifetime:      securityTokenLifetime,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEndpointConfiguration(structType any) EndpointConfiguration {
	if casted, ok := structType.(EndpointConfiguration); ok {
		return casted
	}
	if casted, ok := structType.(*EndpointConfiguration); ok {
		return *casted
	}
	return nil
}

func (m *_EndpointConfiguration) GetTypeName() string {
	return "EndpointConfiguration"
}

func (m *_EndpointConfiguration) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (operationTimeout)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (useBinaryEncoding)
	lengthInBits += 1

	// Simple field (maxStringLength)
	lengthInBits += 32

	// Simple field (maxByteStringLength)
	lengthInBits += 32

	// Simple field (maxArrayLength)
	lengthInBits += 32

	// Simple field (maxMessageSize)
	lengthInBits += 32

	// Simple field (maxBufferSize)
	lengthInBits += 32

	// Simple field (channelLifetime)
	lengthInBits += 32

	// Simple field (securityTokenLifetime)
	lengthInBits += 32

	return lengthInBits
}

func (m *_EndpointConfiguration) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EndpointConfigurationParse(ctx context.Context, theBytes []byte, identifier string) (EndpointConfiguration, error) {
	return EndpointConfigurationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func EndpointConfigurationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (EndpointConfiguration, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("EndpointConfiguration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EndpointConfiguration")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (operationTimeout)
	_operationTimeout, _operationTimeoutErr := readBuffer.ReadInt32("operationTimeout", 32)
	if _operationTimeoutErr != nil {
		return nil, errors.Wrap(_operationTimeoutErr, "Error parsing 'operationTimeout' field of EndpointConfiguration")
	}
	operationTimeout := _operationTimeout

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of EndpointConfiguration")
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

	// Simple Field (useBinaryEncoding)
	_useBinaryEncoding, _useBinaryEncodingErr := readBuffer.ReadBit("useBinaryEncoding")
	if _useBinaryEncodingErr != nil {
		return nil, errors.Wrap(_useBinaryEncodingErr, "Error parsing 'useBinaryEncoding' field of EndpointConfiguration")
	}
	useBinaryEncoding := _useBinaryEncoding

	// Simple Field (maxStringLength)
	_maxStringLength, _maxStringLengthErr := readBuffer.ReadInt32("maxStringLength", 32)
	if _maxStringLengthErr != nil {
		return nil, errors.Wrap(_maxStringLengthErr, "Error parsing 'maxStringLength' field of EndpointConfiguration")
	}
	maxStringLength := _maxStringLength

	// Simple Field (maxByteStringLength)
	_maxByteStringLength, _maxByteStringLengthErr := readBuffer.ReadInt32("maxByteStringLength", 32)
	if _maxByteStringLengthErr != nil {
		return nil, errors.Wrap(_maxByteStringLengthErr, "Error parsing 'maxByteStringLength' field of EndpointConfiguration")
	}
	maxByteStringLength := _maxByteStringLength

	// Simple Field (maxArrayLength)
	_maxArrayLength, _maxArrayLengthErr := readBuffer.ReadInt32("maxArrayLength", 32)
	if _maxArrayLengthErr != nil {
		return nil, errors.Wrap(_maxArrayLengthErr, "Error parsing 'maxArrayLength' field of EndpointConfiguration")
	}
	maxArrayLength := _maxArrayLength

	// Simple Field (maxMessageSize)
	_maxMessageSize, _maxMessageSizeErr := readBuffer.ReadInt32("maxMessageSize", 32)
	if _maxMessageSizeErr != nil {
		return nil, errors.Wrap(_maxMessageSizeErr, "Error parsing 'maxMessageSize' field of EndpointConfiguration")
	}
	maxMessageSize := _maxMessageSize

	// Simple Field (maxBufferSize)
	_maxBufferSize, _maxBufferSizeErr := readBuffer.ReadInt32("maxBufferSize", 32)
	if _maxBufferSizeErr != nil {
		return nil, errors.Wrap(_maxBufferSizeErr, "Error parsing 'maxBufferSize' field of EndpointConfiguration")
	}
	maxBufferSize := _maxBufferSize

	// Simple Field (channelLifetime)
	_channelLifetime, _channelLifetimeErr := readBuffer.ReadInt32("channelLifetime", 32)
	if _channelLifetimeErr != nil {
		return nil, errors.Wrap(_channelLifetimeErr, "Error parsing 'channelLifetime' field of EndpointConfiguration")
	}
	channelLifetime := _channelLifetime

	// Simple Field (securityTokenLifetime)
	_securityTokenLifetime, _securityTokenLifetimeErr := readBuffer.ReadInt32("securityTokenLifetime", 32)
	if _securityTokenLifetimeErr != nil {
		return nil, errors.Wrap(_securityTokenLifetimeErr, "Error parsing 'securityTokenLifetime' field of EndpointConfiguration")
	}
	securityTokenLifetime := _securityTokenLifetime

	if closeErr := readBuffer.CloseContext("EndpointConfiguration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EndpointConfiguration")
	}

	// Create a partially initialized instance
	_child := &_EndpointConfiguration{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		OperationTimeout:           operationTimeout,
		UseBinaryEncoding:          useBinaryEncoding,
		MaxStringLength:            maxStringLength,
		MaxByteStringLength:        maxByteStringLength,
		MaxArrayLength:             maxArrayLength,
		MaxMessageSize:             maxMessageSize,
		MaxBufferSize:              maxBufferSize,
		ChannelLifetime:            channelLifetime,
		SecurityTokenLifetime:      securityTokenLifetime,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_EndpointConfiguration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EndpointConfiguration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EndpointConfiguration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EndpointConfiguration")
		}

		// Simple Field (operationTimeout)
		operationTimeout := int32(m.GetOperationTimeout())
		_operationTimeoutErr := writeBuffer.WriteInt32("operationTimeout", 32, int32((operationTimeout)))
		if _operationTimeoutErr != nil {
			return errors.Wrap(_operationTimeoutErr, "Error serializing 'operationTimeout' field")
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

		// Simple Field (useBinaryEncoding)
		useBinaryEncoding := bool(m.GetUseBinaryEncoding())
		_useBinaryEncodingErr := writeBuffer.WriteBit("useBinaryEncoding", (useBinaryEncoding))
		if _useBinaryEncodingErr != nil {
			return errors.Wrap(_useBinaryEncodingErr, "Error serializing 'useBinaryEncoding' field")
		}

		// Simple Field (maxStringLength)
		maxStringLength := int32(m.GetMaxStringLength())
		_maxStringLengthErr := writeBuffer.WriteInt32("maxStringLength", 32, int32((maxStringLength)))
		if _maxStringLengthErr != nil {
			return errors.Wrap(_maxStringLengthErr, "Error serializing 'maxStringLength' field")
		}

		// Simple Field (maxByteStringLength)
		maxByteStringLength := int32(m.GetMaxByteStringLength())
		_maxByteStringLengthErr := writeBuffer.WriteInt32("maxByteStringLength", 32, int32((maxByteStringLength)))
		if _maxByteStringLengthErr != nil {
			return errors.Wrap(_maxByteStringLengthErr, "Error serializing 'maxByteStringLength' field")
		}

		// Simple Field (maxArrayLength)
		maxArrayLength := int32(m.GetMaxArrayLength())
		_maxArrayLengthErr := writeBuffer.WriteInt32("maxArrayLength", 32, int32((maxArrayLength)))
		if _maxArrayLengthErr != nil {
			return errors.Wrap(_maxArrayLengthErr, "Error serializing 'maxArrayLength' field")
		}

		// Simple Field (maxMessageSize)
		maxMessageSize := int32(m.GetMaxMessageSize())
		_maxMessageSizeErr := writeBuffer.WriteInt32("maxMessageSize", 32, int32((maxMessageSize)))
		if _maxMessageSizeErr != nil {
			return errors.Wrap(_maxMessageSizeErr, "Error serializing 'maxMessageSize' field")
		}

		// Simple Field (maxBufferSize)
		maxBufferSize := int32(m.GetMaxBufferSize())
		_maxBufferSizeErr := writeBuffer.WriteInt32("maxBufferSize", 32, int32((maxBufferSize)))
		if _maxBufferSizeErr != nil {
			return errors.Wrap(_maxBufferSizeErr, "Error serializing 'maxBufferSize' field")
		}

		// Simple Field (channelLifetime)
		channelLifetime := int32(m.GetChannelLifetime())
		_channelLifetimeErr := writeBuffer.WriteInt32("channelLifetime", 32, int32((channelLifetime)))
		if _channelLifetimeErr != nil {
			return errors.Wrap(_channelLifetimeErr, "Error serializing 'channelLifetime' field")
		}

		// Simple Field (securityTokenLifetime)
		securityTokenLifetime := int32(m.GetSecurityTokenLifetime())
		_securityTokenLifetimeErr := writeBuffer.WriteInt32("securityTokenLifetime", 32, int32((securityTokenLifetime)))
		if _securityTokenLifetimeErr != nil {
			return errors.Wrap(_securityTokenLifetimeErr, "Error serializing 'securityTokenLifetime' field")
		}

		if popErr := writeBuffer.PopContext("EndpointConfiguration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EndpointConfiguration")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EndpointConfiguration) isEndpointConfiguration() bool {
	return true
}

func (m *_EndpointConfiguration) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
