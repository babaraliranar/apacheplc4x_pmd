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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// NetworkAddressDataType is the corresponding interface of NetworkAddressDataType
type NetworkAddressDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNetworkInterface returns NetworkInterface (property field)
	GetNetworkInterface() PascalString
	// IsNetworkAddressDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNetworkAddressDataType()
}

// _NetworkAddressDataType is the data-structure of this message
type _NetworkAddressDataType struct {
	ExtensionObjectDefinitionContract
	NetworkInterface PascalString
}

var _ NetworkAddressDataType = (*_NetworkAddressDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NetworkAddressDataType)(nil)

// NewNetworkAddressDataType factory function for _NetworkAddressDataType
func NewNetworkAddressDataType(networkInterface PascalString) *_NetworkAddressDataType {
	if networkInterface == nil {
		panic("networkInterface of type PascalString for NetworkAddressDataType must not be nil")
	}
	_result := &_NetworkAddressDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NetworkInterface:                  networkInterface,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NetworkAddressDataType) GetIdentifier() string {
	return "15504"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NetworkAddressDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkAddressDataType) GetNetworkInterface() PascalString {
	return m.NetworkInterface
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNetworkAddressDataType(structType any) NetworkAddressDataType {
	if casted, ok := structType.(NetworkAddressDataType); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkAddressDataType); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkAddressDataType) GetTypeName() string {
	return "NetworkAddressDataType"
}

func (m *_NetworkAddressDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (networkInterface)
	lengthInBits += m.NetworkInterface.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_NetworkAddressDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NetworkAddressDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__networkAddressDataType NetworkAddressDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkAddressDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkAddressDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkInterface, err := ReadSimpleField[PascalString](ctx, "networkInterface", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkInterface' field"))
	}
	m.NetworkInterface = networkInterface

	if closeErr := readBuffer.CloseContext("NetworkAddressDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkAddressDataType")
	}

	return m, nil
}

func (m *_NetworkAddressDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkAddressDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NetworkAddressDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NetworkAddressDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "networkInterface", m.GetNetworkInterface(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkInterface' field")
		}

		if popErr := writeBuffer.PopContext("NetworkAddressDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NetworkAddressDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NetworkAddressDataType) IsNetworkAddressDataType() {}

func (m *_NetworkAddressDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NetworkAddressDataType) deepCopy() *_NetworkAddressDataType {
	if m == nil {
		return nil
	}
	_NetworkAddressDataTypeCopy := &_NetworkAddressDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NetworkInterface.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NetworkAddressDataTypeCopy
}

func (m *_NetworkAddressDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
