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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataIPDHCPServer is the data-structure of this message
type BACnetConstructedDataIPDHCPServer struct {
	*BACnetConstructedData
	DhcpServer *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataIPDHCPServer is the corresponding interface of BACnetConstructedDataIPDHCPServer
type IBACnetConstructedDataIPDHCPServer interface {
	IBACnetConstructedData
	// GetDhcpServer returns DhcpServer (property field)
	GetDhcpServer() *BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagOctetString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataIPDHCPServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPDHCPServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPDHCPServer) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPDHCPServer) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPDHCPServer) GetDhcpServer() *BACnetApplicationTagOctetString {
	return m.DhcpServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataIPDHCPServer) GetActualValue() *BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetDhcpServer())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDHCPServer factory function for BACnetConstructedDataIPDHCPServer
func NewBACnetConstructedDataIPDHCPServer(dhcpServer *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataIPDHCPServer {
	_result := &BACnetConstructedDataIPDHCPServer{
		DhcpServer:            dhcpServer,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPDHCPServer(structType interface{}) *BACnetConstructedDataIPDHCPServer {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPServer); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPServer); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPDHCPServer(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPDHCPServer(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPDHCPServer) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPServer"
}

func (m *BACnetConstructedDataIPDHCPServer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPDHCPServer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dhcpServer)
	lengthInBits += m.DhcpServer.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataIPDHCPServer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPDHCPServerParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataIPDHCPServer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dhcpServer)
	if pullErr := readBuffer.PullContext("dhcpServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dhcpServer")
	}
	_dhcpServer, _dhcpServerErr := BACnetApplicationTagParse(readBuffer)
	if _dhcpServerErr != nil {
		return nil, errors.Wrap(_dhcpServerErr, "Error parsing 'dhcpServer' field")
	}
	dhcpServer := CastBACnetApplicationTagOctetString(_dhcpServer)
	if closeErr := readBuffer.CloseContext("dhcpServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dhcpServer")
	}

	// Virtual field
	_actualValue := dhcpServer
	actualValue := CastBACnetApplicationTagOctetString(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPServer")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPDHCPServer{
		DhcpServer:            CastBACnetApplicationTagOctetString(dhcpServer),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPDHCPServer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPServer")
		}

		// Simple Field (dhcpServer)
		if pushErr := writeBuffer.PushContext("dhcpServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dhcpServer")
		}
		_dhcpServerErr := writeBuffer.WriteSerializable(m.DhcpServer)
		if popErr := writeBuffer.PopContext("dhcpServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dhcpServer")
		}
		if _dhcpServerErr != nil {
			return errors.Wrap(_dhcpServerErr, "Error serializing 'dhcpServer' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPServer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPDHCPServer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
