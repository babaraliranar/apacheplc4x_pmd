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


// AmsNetId is the corresponding interface of AmsNetId
type AmsNetId interface {
	utils.LengthAware
	utils.Serializable
	// GetOctet1 returns Octet1 (property field)
	GetOctet1() uint8
	// GetOctet2 returns Octet2 (property field)
	GetOctet2() uint8
	// GetOctet3 returns Octet3 (property field)
	GetOctet3() uint8
	// GetOctet4 returns Octet4 (property field)
	GetOctet4() uint8
	// GetOctet5 returns Octet5 (property field)
	GetOctet5() uint8
	// GetOctet6 returns Octet6 (property field)
	GetOctet6() uint8
}

// AmsNetIdExactly can be used when we want exactly this type and not a type which fulfills AmsNetId.
// This is useful for switch cases.
type AmsNetIdExactly interface {
	AmsNetId
	isAmsNetId() bool
}

// _AmsNetId is the data-structure of this message
type _AmsNetId struct {
        Octet1 uint8
        Octet2 uint8
        Octet3 uint8
        Octet4 uint8
        Octet5 uint8
        Octet6 uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsNetId) GetOctet1() uint8 {
	return m.Octet1
}

func (m *_AmsNetId) GetOctet2() uint8 {
	return m.Octet2
}

func (m *_AmsNetId) GetOctet3() uint8 {
	return m.Octet3
}

func (m *_AmsNetId) GetOctet4() uint8 {
	return m.Octet4
}

func (m *_AmsNetId) GetOctet5() uint8 {
	return m.Octet5
}

func (m *_AmsNetId) GetOctet6() uint8 {
	return m.Octet6
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAmsNetId factory function for _AmsNetId
func NewAmsNetId( octet1 uint8 , octet2 uint8 , octet3 uint8 , octet4 uint8 , octet5 uint8 , octet6 uint8 ) *_AmsNetId {
return &_AmsNetId{ Octet1: octet1 , Octet2: octet2 , Octet3: octet3 , Octet4: octet4 , Octet5: octet5 , Octet6: octet6 }
}

// Deprecated: use the interface for direct cast
func CastAmsNetId(structType interface{}) AmsNetId {
    if casted, ok := structType.(AmsNetId); ok {
		return casted
	}
	if casted, ok := structType.(*AmsNetId); ok {
		return *casted
	}
	return nil
}

func (m *_AmsNetId) GetTypeName() string {
	return "AmsNetId"
}

func (m *_AmsNetId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (octet1)
	lengthInBits += 8;

	// Simple field (octet2)
	lengthInBits += 8;

	// Simple field (octet3)
	lengthInBits += 8;

	// Simple field (octet4)
	lengthInBits += 8;

	// Simple field (octet5)
	lengthInBits += 8;

	// Simple field (octet6)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_AmsNetId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsNetIdParse(theBytes []byte) (AmsNetId, error) {
	return AmsNetIdParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AmsNetIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsNetId, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsNetId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (octet1)
_octet1, _octet1Err := readBuffer.ReadUint8("octet1", 8)
	if _octet1Err != nil {
		return nil, errors.Wrap(_octet1Err, "Error parsing 'octet1' field of AmsNetId")
	}
	octet1 := _octet1

	// Simple Field (octet2)
_octet2, _octet2Err := readBuffer.ReadUint8("octet2", 8)
	if _octet2Err != nil {
		return nil, errors.Wrap(_octet2Err, "Error parsing 'octet2' field of AmsNetId")
	}
	octet2 := _octet2

	// Simple Field (octet3)
_octet3, _octet3Err := readBuffer.ReadUint8("octet3", 8)
	if _octet3Err != nil {
		return nil, errors.Wrap(_octet3Err, "Error parsing 'octet3' field of AmsNetId")
	}
	octet3 := _octet3

	// Simple Field (octet4)
_octet4, _octet4Err := readBuffer.ReadUint8("octet4", 8)
	if _octet4Err != nil {
		return nil, errors.Wrap(_octet4Err, "Error parsing 'octet4' field of AmsNetId")
	}
	octet4 := _octet4

	// Simple Field (octet5)
_octet5, _octet5Err := readBuffer.ReadUint8("octet5", 8)
	if _octet5Err != nil {
		return nil, errors.Wrap(_octet5Err, "Error parsing 'octet5' field of AmsNetId")
	}
	octet5 := _octet5

	// Simple Field (octet6)
_octet6, _octet6Err := readBuffer.ReadUint8("octet6", 8)
	if _octet6Err != nil {
		return nil, errors.Wrap(_octet6Err, "Error parsing 'octet6' field of AmsNetId")
	}
	octet6 := _octet6

	if closeErr := readBuffer.CloseContext("AmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsNetId")
	}

	// Create the instance
	return &_AmsNetId{
			Octet1: octet1,
			Octet2: octet2,
			Octet3: octet3,
			Octet4: octet4,
			Octet5: octet5,
			Octet6: octet6,
		}, nil
}

func (m *_AmsNetId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsNetId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("AmsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsNetId")
	}

	// Simple Field (octet1)
	octet1 := uint8(m.GetOctet1())
	_octet1Err := writeBuffer.WriteUint8("octet1", 8, (octet1))
	if _octet1Err != nil {
		return errors.Wrap(_octet1Err, "Error serializing 'octet1' field")
	}

	// Simple Field (octet2)
	octet2 := uint8(m.GetOctet2())
	_octet2Err := writeBuffer.WriteUint8("octet2", 8, (octet2))
	if _octet2Err != nil {
		return errors.Wrap(_octet2Err, "Error serializing 'octet2' field")
	}

	// Simple Field (octet3)
	octet3 := uint8(m.GetOctet3())
	_octet3Err := writeBuffer.WriteUint8("octet3", 8, (octet3))
	if _octet3Err != nil {
		return errors.Wrap(_octet3Err, "Error serializing 'octet3' field")
	}

	// Simple Field (octet4)
	octet4 := uint8(m.GetOctet4())
	_octet4Err := writeBuffer.WriteUint8("octet4", 8, (octet4))
	if _octet4Err != nil {
		return errors.Wrap(_octet4Err, "Error serializing 'octet4' field")
	}

	// Simple Field (octet5)
	octet5 := uint8(m.GetOctet5())
	_octet5Err := writeBuffer.WriteUint8("octet5", 8, (octet5))
	if _octet5Err != nil {
		return errors.Wrap(_octet5Err, "Error serializing 'octet5' field")
	}

	// Simple Field (octet6)
	octet6 := uint8(m.GetOctet6())
	_octet6Err := writeBuffer.WriteUint8("octet6", 8, (octet6))
	if _octet6Err != nil {
		return errors.Wrap(_octet6Err, "Error serializing 'octet6' field")
	}

	if popErr := writeBuffer.PopContext("AmsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsNetId")
	}
	return nil
}


func (m *_AmsNetId) isAmsNetId() bool {
	return true
}

func (m *_AmsNetId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



