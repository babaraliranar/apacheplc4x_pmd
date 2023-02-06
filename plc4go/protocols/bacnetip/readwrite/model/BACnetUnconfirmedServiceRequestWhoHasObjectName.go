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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetUnconfirmedServiceRequestWhoHasObjectName is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHasObjectName
type BACnetUnconfirmedServiceRequestWhoHasObjectName interface {
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequestWhoHasObject
	// GetObjectName returns ObjectName (property field)
	GetObjectName() BACnetContextTagCharacterString
}

// BACnetUnconfirmedServiceRequestWhoHasObjectNameExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWhoHasObjectName.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWhoHasObjectNameExactly interface {
	BACnetUnconfirmedServiceRequestWhoHasObjectName
	isBACnetUnconfirmedServiceRequestWhoHasObjectName() bool
}

// _BACnetUnconfirmedServiceRequestWhoHasObjectName is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHasObjectName struct {
	*_BACnetUnconfirmedServiceRequestWhoHasObject
	ObjectName BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) InitializeParent(parent BACnetUnconfirmedServiceRequestWhoHasObject, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetParent() BACnetUnconfirmedServiceRequestWhoHasObject {
	return m._BACnetUnconfirmedServiceRequestWhoHasObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetObjectName() BACnetContextTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHasObjectName factory function for _BACnetUnconfirmedServiceRequestWhoHasObjectName
func NewBACnetUnconfirmedServiceRequestWhoHasObjectName(objectName BACnetContextTagCharacterString, peekedTagHeader BACnetTagHeader) *_BACnetUnconfirmedServiceRequestWhoHasObjectName {
	_result := &_BACnetUnconfirmedServiceRequestWhoHasObjectName{
		ObjectName: objectName,
		_BACnetUnconfirmedServiceRequestWhoHasObject: NewBACnetUnconfirmedServiceRequestWhoHasObject(peekedTagHeader),
	}
	_result._BACnetUnconfirmedServiceRequestWhoHasObject._BACnetUnconfirmedServiceRequestWhoHasObjectChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHasObjectName(structType interface{}) BACnetUnconfirmedServiceRequestWhoHasObjectName {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHasObjectName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHasObjectName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHasObjectName"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestWhoHasObjectNameParse(theBytes []byte) (BACnetUnconfirmedServiceRequestWhoHasObjectName, error) {
	return BACnetUnconfirmedServiceRequestWhoHasObjectNameParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetUnconfirmedServiceRequestWhoHasObjectNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoHasObjectName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHasObjectName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectName)
	if pullErr := readBuffer.PullContext("objectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectName")
	}
	_objectName, _objectNameErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _objectNameErr != nil {
		return nil, errors.Wrap(_objectNameErr, "Error parsing 'objectName' field of BACnetUnconfirmedServiceRequestWhoHasObjectName")
	}
	objectName := _objectName.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("objectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectName")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHasObjectName")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestWhoHasObjectName{
		_BACnetUnconfirmedServiceRequestWhoHasObject: &_BACnetUnconfirmedServiceRequestWhoHasObject{},
		ObjectName: objectName,
	}
	_child._BACnetUnconfirmedServiceRequestWhoHasObject._BACnetUnconfirmedServiceRequestWhoHasObjectChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHasObjectName")
		}

		// Simple Field (objectName)
		if pushErr := writeBuffer.PushContext("objectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectName")
		}
		_objectNameErr := writeBuffer.WriteSerializable(ctx, m.GetObjectName())
		if popErr := writeBuffer.PopContext("objectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectName")
		}
		if _objectNameErr != nil {
			return errors.Wrap(_objectNameErr, "Error serializing 'objectName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHasObjectName")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) isBACnetUnconfirmedServiceRequestWhoHasObjectName() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
