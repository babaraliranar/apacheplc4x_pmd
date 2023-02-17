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


// BACnetKeyIdentifier is the corresponding interface of BACnetKeyIdentifier
type BACnetKeyIdentifier interface {
	utils.LengthAware
	utils.Serializable
	// GetAlgorithm returns Algorithm (property field)
	GetAlgorithm() BACnetContextTagUnsignedInteger
	// GetKeyId returns KeyId (property field)
	GetKeyId() BACnetContextTagUnsignedInteger
}

// BACnetKeyIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetKeyIdentifier.
// This is useful for switch cases.
type BACnetKeyIdentifierExactly interface {
	BACnetKeyIdentifier
	isBACnetKeyIdentifier() bool
}

// _BACnetKeyIdentifier is the data-structure of this message
type _BACnetKeyIdentifier struct {
        Algorithm BACnetContextTagUnsignedInteger
        KeyId BACnetContextTagUnsignedInteger
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetKeyIdentifier) GetAlgorithm() BACnetContextTagUnsignedInteger {
	return m.Algorithm
}

func (m *_BACnetKeyIdentifier) GetKeyId() BACnetContextTagUnsignedInteger {
	return m.KeyId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetKeyIdentifier factory function for _BACnetKeyIdentifier
func NewBACnetKeyIdentifier( algorithm BACnetContextTagUnsignedInteger , keyId BACnetContextTagUnsignedInteger ) *_BACnetKeyIdentifier {
return &_BACnetKeyIdentifier{ Algorithm: algorithm , KeyId: keyId }
}

// Deprecated: use the interface for direct cast
func CastBACnetKeyIdentifier(structType interface{}) BACnetKeyIdentifier {
    if casted, ok := structType.(BACnetKeyIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetKeyIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetKeyIdentifier) GetTypeName() string {
	return "BACnetKeyIdentifier"
}

func (m *_BACnetKeyIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (algorithm)
	lengthInBits += m.Algorithm.GetLengthInBits(ctx)

	// Simple field (keyId)
	lengthInBits += m.KeyId.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetKeyIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetKeyIdentifierParse(theBytes []byte) (BACnetKeyIdentifier, error) {
	return BACnetKeyIdentifierParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetKeyIdentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetKeyIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetKeyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetKeyIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (algorithm)
	if pullErr := readBuffer.PullContext("algorithm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for algorithm")
	}
_algorithm, _algorithmErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _algorithmErr != nil {
		return nil, errors.Wrap(_algorithmErr, "Error parsing 'algorithm' field of BACnetKeyIdentifier")
	}
	algorithm := _algorithm.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("algorithm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for algorithm")
	}

	// Simple Field (keyId)
	if pullErr := readBuffer.PullContext("keyId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for keyId")
	}
_keyId, _keyIdErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(1) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _keyIdErr != nil {
		return nil, errors.Wrap(_keyIdErr, "Error parsing 'keyId' field of BACnetKeyIdentifier")
	}
	keyId := _keyId.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("keyId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for keyId")
	}

	if closeErr := readBuffer.CloseContext("BACnetKeyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetKeyIdentifier")
	}

	// Create the instance
	return &_BACnetKeyIdentifier{
			Algorithm: algorithm,
			KeyId: keyId,
		}, nil
}

func (m *_BACnetKeyIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetKeyIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetKeyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetKeyIdentifier")
	}

	// Simple Field (algorithm)
	if pushErr := writeBuffer.PushContext("algorithm"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for algorithm")
	}
	_algorithmErr := writeBuffer.WriteSerializable(ctx, m.GetAlgorithm())
	if popErr := writeBuffer.PopContext("algorithm"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for algorithm")
	}
	if _algorithmErr != nil {
		return errors.Wrap(_algorithmErr, "Error serializing 'algorithm' field")
	}

	// Simple Field (keyId)
	if pushErr := writeBuffer.PushContext("keyId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for keyId")
	}
	_keyIdErr := writeBuffer.WriteSerializable(ctx, m.GetKeyId())
	if popErr := writeBuffer.PopContext("keyId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for keyId")
	}
	if _keyIdErr != nil {
		return errors.Wrap(_keyIdErr, "Error serializing 'keyId' field")
	}

	if popErr := writeBuffer.PopContext("BACnetKeyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetKeyIdentifier")
	}
	return nil
}


func (m *_BACnetKeyIdentifier) isBACnetKeyIdentifier() bool {
	return true
}

func (m *_BACnetKeyIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



