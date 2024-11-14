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

// BACnetPropertyStatesSystemStatus is the corresponding interface of BACnetPropertyStatesSystemStatus
type BACnetPropertyStatesSystemStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetSystemStatus returns SystemStatus (property field)
	GetSystemStatus() BACnetDeviceStatusTagged
	// IsBACnetPropertyStatesSystemStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesSystemStatus()
	// CreateBuilder creates a BACnetPropertyStatesSystemStatusBuilder
	CreateBACnetPropertyStatesSystemStatusBuilder() BACnetPropertyStatesSystemStatusBuilder
}

// _BACnetPropertyStatesSystemStatus is the data-structure of this message
type _BACnetPropertyStatesSystemStatus struct {
	BACnetPropertyStatesContract
	SystemStatus BACnetDeviceStatusTagged
}

var _ BACnetPropertyStatesSystemStatus = (*_BACnetPropertyStatesSystemStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesSystemStatus)(nil)

// NewBACnetPropertyStatesSystemStatus factory function for _BACnetPropertyStatesSystemStatus
func NewBACnetPropertyStatesSystemStatus(peekedTagHeader BACnetTagHeader, systemStatus BACnetDeviceStatusTagged) *_BACnetPropertyStatesSystemStatus {
	if systemStatus == nil {
		panic("systemStatus of type BACnetDeviceStatusTagged for BACnetPropertyStatesSystemStatus must not be nil")
	}
	_result := &_BACnetPropertyStatesSystemStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		SystemStatus:                 systemStatus,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesSystemStatusBuilder is a builder for BACnetPropertyStatesSystemStatus
type BACnetPropertyStatesSystemStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(systemStatus BACnetDeviceStatusTagged) BACnetPropertyStatesSystemStatusBuilder
	// WithSystemStatus adds SystemStatus (property field)
	WithSystemStatus(BACnetDeviceStatusTagged) BACnetPropertyStatesSystemStatusBuilder
	// WithSystemStatusBuilder adds SystemStatus (property field) which is build by the builder
	WithSystemStatusBuilder(func(BACnetDeviceStatusTaggedBuilder) BACnetDeviceStatusTaggedBuilder) BACnetPropertyStatesSystemStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesSystemStatus or returns an error if something is wrong
	Build() (BACnetPropertyStatesSystemStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesSystemStatus
}

// NewBACnetPropertyStatesSystemStatusBuilder() creates a BACnetPropertyStatesSystemStatusBuilder
func NewBACnetPropertyStatesSystemStatusBuilder() BACnetPropertyStatesSystemStatusBuilder {
	return &_BACnetPropertyStatesSystemStatusBuilder{_BACnetPropertyStatesSystemStatus: new(_BACnetPropertyStatesSystemStatus)}
}

type _BACnetPropertyStatesSystemStatusBuilder struct {
	*_BACnetPropertyStatesSystemStatus

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesSystemStatusBuilder) = (*_BACnetPropertyStatesSystemStatusBuilder)(nil)

func (b *_BACnetPropertyStatesSystemStatusBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesSystemStatus
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) WithMandatoryFields(systemStatus BACnetDeviceStatusTagged) BACnetPropertyStatesSystemStatusBuilder {
	return b.WithSystemStatus(systemStatus)
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) WithSystemStatus(systemStatus BACnetDeviceStatusTagged) BACnetPropertyStatesSystemStatusBuilder {
	b.SystemStatus = systemStatus
	return b
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) WithSystemStatusBuilder(builderSupplier func(BACnetDeviceStatusTaggedBuilder) BACnetDeviceStatusTaggedBuilder) BACnetPropertyStatesSystemStatusBuilder {
	builder := builderSupplier(b.SystemStatus.CreateBACnetDeviceStatusTaggedBuilder())
	var err error
	b.SystemStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceStatusTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) Build() (BACnetPropertyStatesSystemStatus, error) {
	if b.SystemStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'systemStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesSystemStatus.deepCopy(), nil
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) MustBuild() BACnetPropertyStatesSystemStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesSystemStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesSystemStatusBuilder().(*_BACnetPropertyStatesSystemStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesSystemStatusBuilder creates a BACnetPropertyStatesSystemStatusBuilder
func (b *_BACnetPropertyStatesSystemStatus) CreateBACnetPropertyStatesSystemStatusBuilder() BACnetPropertyStatesSystemStatusBuilder {
	if b == nil {
		return NewBACnetPropertyStatesSystemStatusBuilder()
	}
	return &_BACnetPropertyStatesSystemStatusBuilder{_BACnetPropertyStatesSystemStatus: b.deepCopy()}
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

func (m *_BACnetPropertyStatesSystemStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesSystemStatus) GetSystemStatus() BACnetDeviceStatusTagged {
	return m.SystemStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesSystemStatus(structType any) BACnetPropertyStatesSystemStatus {
	if casted, ok := structType.(BACnetPropertyStatesSystemStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSystemStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesSystemStatus) GetTypeName() string {
	return "BACnetPropertyStatesSystemStatus"
}

func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (systemStatus)
	lengthInBits += m.SystemStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesSystemStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesSystemStatus BACnetPropertyStatesSystemStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSystemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSystemStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	systemStatus, err := ReadSimpleField[BACnetDeviceStatusTagged](ctx, "systemStatus", ReadComplex[BACnetDeviceStatusTagged](BACnetDeviceStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemStatus' field"))
	}
	m.SystemStatus = systemStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSystemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSystemStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesSystemStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesSystemStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSystemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSystemStatus")
		}

		if err := WriteSimpleField[BACnetDeviceStatusTagged](ctx, "systemStatus", m.GetSystemStatus(), WriteComplex[BACnetDeviceStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'systemStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSystemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSystemStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesSystemStatus) IsBACnetPropertyStatesSystemStatus() {}

func (m *_BACnetPropertyStatesSystemStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesSystemStatus) deepCopy() *_BACnetPropertyStatesSystemStatus {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesSystemStatusCopy := &_BACnetPropertyStatesSystemStatus{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetDeviceStatusTagged](m.SystemStatus),
	}
	_BACnetPropertyStatesSystemStatusCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesSystemStatusCopy
}

func (m *_BACnetPropertyStatesSystemStatus) String() string {
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
