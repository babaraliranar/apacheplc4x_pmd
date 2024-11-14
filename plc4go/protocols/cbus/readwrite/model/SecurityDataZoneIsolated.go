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

// SecurityDataZoneIsolated is the corresponding interface of SecurityDataZoneIsolated
type SecurityDataZoneIsolated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
	// IsSecurityDataZoneIsolated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataZoneIsolated()
	// CreateBuilder creates a SecurityDataZoneIsolatedBuilder
	CreateSecurityDataZoneIsolatedBuilder() SecurityDataZoneIsolatedBuilder
}

// _SecurityDataZoneIsolated is the data-structure of this message
type _SecurityDataZoneIsolated struct {
	SecurityDataContract
	ZoneNumber uint8
}

var _ SecurityDataZoneIsolated = (*_SecurityDataZoneIsolated)(nil)
var _ SecurityDataRequirements = (*_SecurityDataZoneIsolated)(nil)

// NewSecurityDataZoneIsolated factory function for _SecurityDataZoneIsolated
func NewSecurityDataZoneIsolated(commandTypeContainer SecurityCommandTypeContainer, argument byte, zoneNumber uint8) *_SecurityDataZoneIsolated {
	_result := &_SecurityDataZoneIsolated{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneNumber:           zoneNumber,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataZoneIsolatedBuilder is a builder for SecurityDataZoneIsolated
type SecurityDataZoneIsolatedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneNumber uint8) SecurityDataZoneIsolatedBuilder
	// WithZoneNumber adds ZoneNumber (property field)
	WithZoneNumber(uint8) SecurityDataZoneIsolatedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataZoneIsolated or returns an error if something is wrong
	Build() (SecurityDataZoneIsolated, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataZoneIsolated
}

// NewSecurityDataZoneIsolatedBuilder() creates a SecurityDataZoneIsolatedBuilder
func NewSecurityDataZoneIsolatedBuilder() SecurityDataZoneIsolatedBuilder {
	return &_SecurityDataZoneIsolatedBuilder{_SecurityDataZoneIsolated: new(_SecurityDataZoneIsolated)}
}

type _SecurityDataZoneIsolatedBuilder struct {
	*_SecurityDataZoneIsolated

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataZoneIsolatedBuilder) = (*_SecurityDataZoneIsolatedBuilder)(nil)

func (b *_SecurityDataZoneIsolatedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataZoneIsolated
}

func (b *_SecurityDataZoneIsolatedBuilder) WithMandatoryFields(zoneNumber uint8) SecurityDataZoneIsolatedBuilder {
	return b.WithZoneNumber(zoneNumber)
}

func (b *_SecurityDataZoneIsolatedBuilder) WithZoneNumber(zoneNumber uint8) SecurityDataZoneIsolatedBuilder {
	b.ZoneNumber = zoneNumber
	return b
}

func (b *_SecurityDataZoneIsolatedBuilder) Build() (SecurityDataZoneIsolated, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataZoneIsolated.deepCopy(), nil
}

func (b *_SecurityDataZoneIsolatedBuilder) MustBuild() SecurityDataZoneIsolated {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataZoneIsolatedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataZoneIsolatedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataZoneIsolatedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataZoneIsolatedBuilder().(*_SecurityDataZoneIsolatedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataZoneIsolatedBuilder creates a SecurityDataZoneIsolatedBuilder
func (b *_SecurityDataZoneIsolated) CreateSecurityDataZoneIsolatedBuilder() SecurityDataZoneIsolatedBuilder {
	if b == nil {
		return NewSecurityDataZoneIsolatedBuilder()
	}
	return &_SecurityDataZoneIsolatedBuilder{_SecurityDataZoneIsolated: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneIsolated) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneIsolated) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneIsolated(structType any) SecurityDataZoneIsolated {
	if casted, ok := structType.(SecurityDataZoneIsolated); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneIsolated); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneIsolated) GetTypeName() string {
	return "SecurityDataZoneIsolated"
}

func (m *_SecurityDataZoneIsolated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneIsolated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataZoneIsolated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataZoneIsolated SecurityDataZoneIsolated, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneIsolated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneIsolated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneNumber, err := ReadSimpleField(ctx, "zoneNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneNumber' field"))
	}
	m.ZoneNumber = zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneIsolated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneIsolated")
	}

	return m, nil
}

func (m *_SecurityDataZoneIsolated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneIsolated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneIsolated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneIsolated")
		}

		if err := WriteSimpleField[uint8](ctx, "zoneNumber", m.GetZoneNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneIsolated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneIsolated")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneIsolated) IsSecurityDataZoneIsolated() {}

func (m *_SecurityDataZoneIsolated) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataZoneIsolated) deepCopy() *_SecurityDataZoneIsolated {
	if m == nil {
		return nil
	}
	_SecurityDataZoneIsolatedCopy := &_SecurityDataZoneIsolated{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.ZoneNumber,
	}
	_SecurityDataZoneIsolatedCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataZoneIsolatedCopy
}

func (m *_SecurityDataZoneIsolated) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
