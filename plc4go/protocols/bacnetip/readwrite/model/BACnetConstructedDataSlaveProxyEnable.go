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

// BACnetConstructedDataSlaveProxyEnable is the data-structure of this message
type BACnetConstructedDataSlaveProxyEnable struct {
	*BACnetConstructedData
	SlaveProxyEnable *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataSlaveProxyEnable is the corresponding interface of BACnetConstructedDataSlaveProxyEnable
type IBACnetConstructedDataSlaveProxyEnable interface {
	IBACnetConstructedData
	// GetSlaveProxyEnable returns SlaveProxyEnable (property field)
	GetSlaveProxyEnable() *BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataSlaveProxyEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SLAVE_PROXY_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSlaveProxyEnable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSlaveProxyEnable) GetSlaveProxyEnable() *BACnetApplicationTagBoolean {
	return m.SlaveProxyEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataSlaveProxyEnable) GetActualValue() *BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetSlaveProxyEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSlaveProxyEnable factory function for BACnetConstructedDataSlaveProxyEnable
func NewBACnetConstructedDataSlaveProxyEnable(slaveProxyEnable *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataSlaveProxyEnable {
	_result := &BACnetConstructedDataSlaveProxyEnable{
		SlaveProxyEnable:      slaveProxyEnable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSlaveProxyEnable(structType interface{}) *BACnetConstructedDataSlaveProxyEnable {
	if casted, ok := structType.(BACnetConstructedDataSlaveProxyEnable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSlaveProxyEnable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSlaveProxyEnable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSlaveProxyEnable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetTypeName() string {
	return "BACnetConstructedDataSlaveProxyEnable"
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (slaveProxyEnable)
	lengthInBits += m.SlaveProxyEnable.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSlaveProxyEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataSlaveProxyEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSlaveProxyEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSlaveProxyEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (slaveProxyEnable)
	if pullErr := readBuffer.PullContext("slaveProxyEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for slaveProxyEnable")
	}
	_slaveProxyEnable, _slaveProxyEnableErr := BACnetApplicationTagParse(readBuffer)
	if _slaveProxyEnableErr != nil {
		return nil, errors.Wrap(_slaveProxyEnableErr, "Error parsing 'slaveProxyEnable' field")
	}
	slaveProxyEnable := CastBACnetApplicationTagBoolean(_slaveProxyEnable)
	if closeErr := readBuffer.CloseContext("slaveProxyEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for slaveProxyEnable")
	}

	// Virtual field
	_actualValue := slaveProxyEnable
	actualValue := CastBACnetApplicationTagBoolean(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSlaveProxyEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSlaveProxyEnable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSlaveProxyEnable{
		SlaveProxyEnable:      CastBACnetApplicationTagBoolean(slaveProxyEnable),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSlaveProxyEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSlaveProxyEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSlaveProxyEnable")
		}

		// Simple Field (slaveProxyEnable)
		if pushErr := writeBuffer.PushContext("slaveProxyEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for slaveProxyEnable")
		}
		_slaveProxyEnableErr := writeBuffer.WriteSerializable(m.SlaveProxyEnable)
		if popErr := writeBuffer.PopContext("slaveProxyEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for slaveProxyEnable")
		}
		if _slaveProxyEnableErr != nil {
			return errors.Wrap(_slaveProxyEnableErr, "Error serializing 'slaveProxyEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSlaveProxyEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSlaveProxyEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSlaveProxyEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
