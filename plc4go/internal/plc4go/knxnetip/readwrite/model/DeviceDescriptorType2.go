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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type DeviceDescriptorType2 struct {
	ManufacturerId uint16
	DeviceType     uint16
	Version        uint8
	ReadSupported  bool
	WriteSupported bool
	LogicalTagBase uint8
	ChannelInfo1   *ChannelInformation
	ChannelInfo2   *ChannelInformation
	ChannelInfo3   *ChannelInformation
	ChannelInfo4   *ChannelInformation
}

// The corresponding interface
type IDeviceDescriptorType2 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewDeviceDescriptorType2(manufacturerId uint16, deviceType uint16, version uint8, readSupported bool, writeSupported bool, logicalTagBase uint8, channelInfo1 *ChannelInformation, channelInfo2 *ChannelInformation, channelInfo3 *ChannelInformation, channelInfo4 *ChannelInformation) *DeviceDescriptorType2 {
	return &DeviceDescriptorType2{ManufacturerId: manufacturerId, DeviceType: deviceType, Version: version, ReadSupported: readSupported, WriteSupported: writeSupported, LogicalTagBase: logicalTagBase, ChannelInfo1: channelInfo1, ChannelInfo2: channelInfo2, ChannelInfo3: channelInfo3, ChannelInfo4: channelInfo4}
}

func CastDeviceDescriptorType2(structType interface{}) *DeviceDescriptorType2 {
	castFunc := func(typ interface{}) *DeviceDescriptorType2 {
		if casted, ok := typ.(DeviceDescriptorType2); ok {
			return &casted
		}
		if casted, ok := typ.(*DeviceDescriptorType2); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DeviceDescriptorType2) GetTypeName() string {
	return "DeviceDescriptorType2"
}

func (m *DeviceDescriptorType2) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DeviceDescriptorType2) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (manufacturerId)
	lengthInBits += 16

	// Simple field (deviceType)
	lengthInBits += 16

	// Simple field (version)
	lengthInBits += 8

	// Simple field (readSupported)
	lengthInBits += 1

	// Simple field (writeSupported)
	lengthInBits += 1

	// Simple field (logicalTagBase)
	lengthInBits += 6

	// Simple field (channelInfo1)
	lengthInBits += m.ChannelInfo1.LengthInBits()

	// Simple field (channelInfo2)
	lengthInBits += m.ChannelInfo2.LengthInBits()

	// Simple field (channelInfo3)
	lengthInBits += m.ChannelInfo3.LengthInBits()

	// Simple field (channelInfo4)
	lengthInBits += m.ChannelInfo4.LengthInBits()

	return lengthInBits
}

func (m *DeviceDescriptorType2) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceDescriptorType2Parse(readBuffer utils.ReadBuffer) (*DeviceDescriptorType2, error) {
	if pullErr := readBuffer.PullContext("DeviceDescriptorType2"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (manufacturerId)
	manufacturerId, _manufacturerIdErr := readBuffer.ReadUint16("manufacturerId", 16)
	if _manufacturerIdErr != nil {
		return nil, errors.Wrap(_manufacturerIdErr, "Error parsing 'manufacturerId' field")
	}

	// Simple Field (deviceType)
	deviceType, _deviceTypeErr := readBuffer.ReadUint16("deviceType", 16)
	if _deviceTypeErr != nil {
		return nil, errors.Wrap(_deviceTypeErr, "Error parsing 'deviceType' field")
	}

	// Simple Field (version)
	version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}

	// Simple Field (readSupported)
	readSupported, _readSupportedErr := readBuffer.ReadBit("readSupported")
	if _readSupportedErr != nil {
		return nil, errors.Wrap(_readSupportedErr, "Error parsing 'readSupported' field")
	}

	// Simple Field (writeSupported)
	writeSupported, _writeSupportedErr := readBuffer.ReadBit("writeSupported")
	if _writeSupportedErr != nil {
		return nil, errors.Wrap(_writeSupportedErr, "Error parsing 'writeSupported' field")
	}

	// Simple Field (logicalTagBase)
	logicalTagBase, _logicalTagBaseErr := readBuffer.ReadUint8("logicalTagBase", 6)
	if _logicalTagBaseErr != nil {
		return nil, errors.Wrap(_logicalTagBaseErr, "Error parsing 'logicalTagBase' field")
	}

	if pullErr := readBuffer.PullContext("channelInfo1"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo1)
	channelInfo1, _channelInfo1Err := ChannelInformationParse(readBuffer)
	if _channelInfo1Err != nil {
		return nil, errors.Wrap(_channelInfo1Err, "Error parsing 'channelInfo1' field")
	}
	if closeErr := readBuffer.CloseContext("channelInfo1"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("channelInfo2"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo2)
	channelInfo2, _channelInfo2Err := ChannelInformationParse(readBuffer)
	if _channelInfo2Err != nil {
		return nil, errors.Wrap(_channelInfo2Err, "Error parsing 'channelInfo2' field")
	}
	if closeErr := readBuffer.CloseContext("channelInfo2"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("channelInfo3"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo3)
	channelInfo3, _channelInfo3Err := ChannelInformationParse(readBuffer)
	if _channelInfo3Err != nil {
		return nil, errors.Wrap(_channelInfo3Err, "Error parsing 'channelInfo3' field")
	}
	if closeErr := readBuffer.CloseContext("channelInfo3"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("channelInfo4"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (channelInfo4)
	channelInfo4, _channelInfo4Err := ChannelInformationParse(readBuffer)
	if _channelInfo4Err != nil {
		return nil, errors.Wrap(_channelInfo4Err, "Error parsing 'channelInfo4' field")
	}
	if closeErr := readBuffer.CloseContext("channelInfo4"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DeviceDescriptorType2"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDeviceDescriptorType2(manufacturerId, deviceType, version, readSupported, writeSupported, logicalTagBase, channelInfo1, channelInfo2, channelInfo3, channelInfo4), nil
}

func (m *DeviceDescriptorType2) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("DeviceDescriptorType2"); pushErr != nil {
		return pushErr
	}

	// Simple Field (manufacturerId)
	manufacturerId := uint16(m.ManufacturerId)
	_manufacturerIdErr := writeBuffer.WriteUint16("manufacturerId", 16, (manufacturerId))
	if _manufacturerIdErr != nil {
		return errors.Wrap(_manufacturerIdErr, "Error serializing 'manufacturerId' field")
	}

	// Simple Field (deviceType)
	deviceType := uint16(m.DeviceType)
	_deviceTypeErr := writeBuffer.WriteUint16("deviceType", 16, (deviceType))
	if _deviceTypeErr != nil {
		return errors.Wrap(_deviceTypeErr, "Error serializing 'deviceType' field")
	}

	// Simple Field (version)
	version := uint8(m.Version)
	_versionErr := writeBuffer.WriteUint8("version", 8, (version))
	if _versionErr != nil {
		return errors.Wrap(_versionErr, "Error serializing 'version' field")
	}

	// Simple Field (readSupported)
	readSupported := bool(m.ReadSupported)
	_readSupportedErr := writeBuffer.WriteBit("readSupported", (readSupported))
	if _readSupportedErr != nil {
		return errors.Wrap(_readSupportedErr, "Error serializing 'readSupported' field")
	}

	// Simple Field (writeSupported)
	writeSupported := bool(m.WriteSupported)
	_writeSupportedErr := writeBuffer.WriteBit("writeSupported", (writeSupported))
	if _writeSupportedErr != nil {
		return errors.Wrap(_writeSupportedErr, "Error serializing 'writeSupported' field")
	}

	// Simple Field (logicalTagBase)
	logicalTagBase := uint8(m.LogicalTagBase)
	_logicalTagBaseErr := writeBuffer.WriteUint8("logicalTagBase", 6, (logicalTagBase))
	if _logicalTagBaseErr != nil {
		return errors.Wrap(_logicalTagBaseErr, "Error serializing 'logicalTagBase' field")
	}

	// Simple Field (channelInfo1)
	if pushErr := writeBuffer.PushContext("channelInfo1"); pushErr != nil {
		return pushErr
	}
	_channelInfo1Err := m.ChannelInfo1.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("channelInfo1"); popErr != nil {
		return popErr
	}
	if _channelInfo1Err != nil {
		return errors.Wrap(_channelInfo1Err, "Error serializing 'channelInfo1' field")
	}

	// Simple Field (channelInfo2)
	if pushErr := writeBuffer.PushContext("channelInfo2"); pushErr != nil {
		return pushErr
	}
	_channelInfo2Err := m.ChannelInfo2.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("channelInfo2"); popErr != nil {
		return popErr
	}
	if _channelInfo2Err != nil {
		return errors.Wrap(_channelInfo2Err, "Error serializing 'channelInfo2' field")
	}

	// Simple Field (channelInfo3)
	if pushErr := writeBuffer.PushContext("channelInfo3"); pushErr != nil {
		return pushErr
	}
	_channelInfo3Err := m.ChannelInfo3.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("channelInfo3"); popErr != nil {
		return popErr
	}
	if _channelInfo3Err != nil {
		return errors.Wrap(_channelInfo3Err, "Error serializing 'channelInfo3' field")
	}

	// Simple Field (channelInfo4)
	if pushErr := writeBuffer.PushContext("channelInfo4"); pushErr != nil {
		return pushErr
	}
	_channelInfo4Err := m.ChannelInfo4.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("channelInfo4"); popErr != nil {
		return popErr
	}
	if _channelInfo4Err != nil {
		return errors.Wrap(_channelInfo4Err, "Error serializing 'channelInfo4' field")
	}

	if popErr := writeBuffer.PopContext("DeviceDescriptorType2"); popErr != nil {
		return popErr
	}
	return nil
}
