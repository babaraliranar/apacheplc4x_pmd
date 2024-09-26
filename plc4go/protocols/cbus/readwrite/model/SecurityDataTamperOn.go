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

// SecurityDataTamperOn is the corresponding interface of SecurityDataTamperOn
type SecurityDataTamperOn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataTamperOn is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataTamperOn()
}

// _SecurityDataTamperOn is the data-structure of this message
type _SecurityDataTamperOn struct {
	SecurityDataContract
}

var _ SecurityDataTamperOn = (*_SecurityDataTamperOn)(nil)
var _ SecurityDataRequirements = (*_SecurityDataTamperOn)(nil)

// NewSecurityDataTamperOn factory function for _SecurityDataTamperOn
func NewSecurityDataTamperOn(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataTamperOn {
	_result := &_SecurityDataTamperOn{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataTamperOn) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataTamperOn(structType any) SecurityDataTamperOn {
	if casted, ok := structType.(SecurityDataTamperOn); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataTamperOn); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataTamperOn) GetTypeName() string {
	return "SecurityDataTamperOn"
}

func (m *_SecurityDataTamperOn) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataTamperOn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataTamperOn) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataTamperOn SecurityDataTamperOn, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataTamperOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataTamperOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataTamperOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataTamperOn")
	}

	return m, nil
}

func (m *_SecurityDataTamperOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataTamperOn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataTamperOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataTamperOn")
		}

		if popErr := writeBuffer.PopContext("SecurityDataTamperOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataTamperOn")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataTamperOn) IsSecurityDataTamperOn() {}

func (m *_SecurityDataTamperOn) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataTamperOn) deepCopy() *_SecurityDataTamperOn {
	if m == nil {
		return nil
	}
	_SecurityDataTamperOnCopy := &_SecurityDataTamperOn{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataTamperOnCopy
}

func (m *_SecurityDataTamperOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
