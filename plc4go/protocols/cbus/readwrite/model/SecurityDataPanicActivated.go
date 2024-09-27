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

// SecurityDataPanicActivated is the corresponding interface of SecurityDataPanicActivated
type SecurityDataPanicActivated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataPanicActivated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataPanicActivated()
	// CreateBuilder creates a SecurityDataPanicActivatedBuilder
	CreateSecurityDataPanicActivatedBuilder() SecurityDataPanicActivatedBuilder
}

// _SecurityDataPanicActivated is the data-structure of this message
type _SecurityDataPanicActivated struct {
	SecurityDataContract
}

var _ SecurityDataPanicActivated = (*_SecurityDataPanicActivated)(nil)
var _ SecurityDataRequirements = (*_SecurityDataPanicActivated)(nil)

// NewSecurityDataPanicActivated factory function for _SecurityDataPanicActivated
func NewSecurityDataPanicActivated(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataPanicActivated {
	_result := &_SecurityDataPanicActivated{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataPanicActivatedBuilder is a builder for SecurityDataPanicActivated
type SecurityDataPanicActivatedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataPanicActivatedBuilder
	// Build builds the SecurityDataPanicActivated or returns an error if something is wrong
	Build() (SecurityDataPanicActivated, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataPanicActivated
}

// NewSecurityDataPanicActivatedBuilder() creates a SecurityDataPanicActivatedBuilder
func NewSecurityDataPanicActivatedBuilder() SecurityDataPanicActivatedBuilder {
	return &_SecurityDataPanicActivatedBuilder{_SecurityDataPanicActivated: new(_SecurityDataPanicActivated)}
}

type _SecurityDataPanicActivatedBuilder struct {
	*_SecurityDataPanicActivated

	err *utils.MultiError
}

var _ (SecurityDataPanicActivatedBuilder) = (*_SecurityDataPanicActivatedBuilder)(nil)

func (m *_SecurityDataPanicActivatedBuilder) WithMandatoryFields() SecurityDataPanicActivatedBuilder {
	return m
}

func (m *_SecurityDataPanicActivatedBuilder) Build() (SecurityDataPanicActivated, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataPanicActivated.deepCopy(), nil
}

func (m *_SecurityDataPanicActivatedBuilder) MustBuild() SecurityDataPanicActivated {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataPanicActivatedBuilder) DeepCopy() any {
	return m.CreateSecurityDataPanicActivatedBuilder()
}

// CreateSecurityDataPanicActivatedBuilder creates a SecurityDataPanicActivatedBuilder
func (m *_SecurityDataPanicActivated) CreateSecurityDataPanicActivatedBuilder() SecurityDataPanicActivatedBuilder {
	if m == nil {
		return NewSecurityDataPanicActivatedBuilder()
	}
	return &_SecurityDataPanicActivatedBuilder{_SecurityDataPanicActivated: m.deepCopy()}
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

func (m *_SecurityDataPanicActivated) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataPanicActivated(structType any) SecurityDataPanicActivated {
	if casted, ok := structType.(SecurityDataPanicActivated); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataPanicActivated); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataPanicActivated) GetTypeName() string {
	return "SecurityDataPanicActivated"
}

func (m *_SecurityDataPanicActivated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataPanicActivated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataPanicActivated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataPanicActivated SecurityDataPanicActivated, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataPanicActivated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataPanicActivated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataPanicActivated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataPanicActivated")
	}

	return m, nil
}

func (m *_SecurityDataPanicActivated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataPanicActivated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataPanicActivated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataPanicActivated")
		}

		if popErr := writeBuffer.PopContext("SecurityDataPanicActivated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataPanicActivated")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataPanicActivated) IsSecurityDataPanicActivated() {}

func (m *_SecurityDataPanicActivated) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataPanicActivated) deepCopy() *_SecurityDataPanicActivated {
	if m == nil {
		return nil
	}
	_SecurityDataPanicActivatedCopy := &_SecurityDataPanicActivated{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataPanicActivatedCopy
}

func (m *_SecurityDataPanicActivated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
