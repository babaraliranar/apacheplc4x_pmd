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

// ServerDiagnosticsSummaryDataType is the corresponding interface of ServerDiagnosticsSummaryDataType
type ServerDiagnosticsSummaryDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetServerViewCount returns ServerViewCount (property field)
	GetServerViewCount() uint32
	// GetCurrentSessionCount returns CurrentSessionCount (property field)
	GetCurrentSessionCount() uint32
	// GetCumulatedSessionCount returns CumulatedSessionCount (property field)
	GetCumulatedSessionCount() uint32
	// GetSecurityRejectedSessionCount returns SecurityRejectedSessionCount (property field)
	GetSecurityRejectedSessionCount() uint32
	// GetRejectedSessionCount returns RejectedSessionCount (property field)
	GetRejectedSessionCount() uint32
	// GetSessionTimeoutCount returns SessionTimeoutCount (property field)
	GetSessionTimeoutCount() uint32
	// GetSessionAbortCount returns SessionAbortCount (property field)
	GetSessionAbortCount() uint32
	// GetCurrentSubscriptionCount returns CurrentSubscriptionCount (property field)
	GetCurrentSubscriptionCount() uint32
	// GetCumulatedSubscriptionCount returns CumulatedSubscriptionCount (property field)
	GetCumulatedSubscriptionCount() uint32
	// GetPublishingIntervalCount returns PublishingIntervalCount (property field)
	GetPublishingIntervalCount() uint32
	// GetSecurityRejectedRequestsCount returns SecurityRejectedRequestsCount (property field)
	GetSecurityRejectedRequestsCount() uint32
	// GetRejectedRequestsCount returns RejectedRequestsCount (property field)
	GetRejectedRequestsCount() uint32
	// IsServerDiagnosticsSummaryDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServerDiagnosticsSummaryDataType()
	// CreateBuilder creates a ServerDiagnosticsSummaryDataTypeBuilder
	CreateServerDiagnosticsSummaryDataTypeBuilder() ServerDiagnosticsSummaryDataTypeBuilder
}

// _ServerDiagnosticsSummaryDataType is the data-structure of this message
type _ServerDiagnosticsSummaryDataType struct {
	ExtensionObjectDefinitionContract
	ServerViewCount               uint32
	CurrentSessionCount           uint32
	CumulatedSessionCount         uint32
	SecurityRejectedSessionCount  uint32
	RejectedSessionCount          uint32
	SessionTimeoutCount           uint32
	SessionAbortCount             uint32
	CurrentSubscriptionCount      uint32
	CumulatedSubscriptionCount    uint32
	PublishingIntervalCount       uint32
	SecurityRejectedRequestsCount uint32
	RejectedRequestsCount         uint32
}

var _ ServerDiagnosticsSummaryDataType = (*_ServerDiagnosticsSummaryDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServerDiagnosticsSummaryDataType)(nil)

// NewServerDiagnosticsSummaryDataType factory function for _ServerDiagnosticsSummaryDataType
func NewServerDiagnosticsSummaryDataType(serverViewCount uint32, currentSessionCount uint32, cumulatedSessionCount uint32, securityRejectedSessionCount uint32, rejectedSessionCount uint32, sessionTimeoutCount uint32, sessionAbortCount uint32, currentSubscriptionCount uint32, cumulatedSubscriptionCount uint32, publishingIntervalCount uint32, securityRejectedRequestsCount uint32, rejectedRequestsCount uint32) *_ServerDiagnosticsSummaryDataType {
	_result := &_ServerDiagnosticsSummaryDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ServerViewCount:                   serverViewCount,
		CurrentSessionCount:               currentSessionCount,
		CumulatedSessionCount:             cumulatedSessionCount,
		SecurityRejectedSessionCount:      securityRejectedSessionCount,
		RejectedSessionCount:              rejectedSessionCount,
		SessionTimeoutCount:               sessionTimeoutCount,
		SessionAbortCount:                 sessionAbortCount,
		CurrentSubscriptionCount:          currentSubscriptionCount,
		CumulatedSubscriptionCount:        cumulatedSubscriptionCount,
		PublishingIntervalCount:           publishingIntervalCount,
		SecurityRejectedRequestsCount:     securityRejectedRequestsCount,
		RejectedRequestsCount:             rejectedRequestsCount,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ServerDiagnosticsSummaryDataTypeBuilder is a builder for ServerDiagnosticsSummaryDataType
type ServerDiagnosticsSummaryDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(serverViewCount uint32, currentSessionCount uint32, cumulatedSessionCount uint32, securityRejectedSessionCount uint32, rejectedSessionCount uint32, sessionTimeoutCount uint32, sessionAbortCount uint32, currentSubscriptionCount uint32, cumulatedSubscriptionCount uint32, publishingIntervalCount uint32, securityRejectedRequestsCount uint32, rejectedRequestsCount uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithServerViewCount adds ServerViewCount (property field)
	WithServerViewCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithCurrentSessionCount adds CurrentSessionCount (property field)
	WithCurrentSessionCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithCumulatedSessionCount adds CumulatedSessionCount (property field)
	WithCumulatedSessionCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithSecurityRejectedSessionCount adds SecurityRejectedSessionCount (property field)
	WithSecurityRejectedSessionCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithRejectedSessionCount adds RejectedSessionCount (property field)
	WithRejectedSessionCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithSessionTimeoutCount adds SessionTimeoutCount (property field)
	WithSessionTimeoutCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithSessionAbortCount adds SessionAbortCount (property field)
	WithSessionAbortCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithCurrentSubscriptionCount adds CurrentSubscriptionCount (property field)
	WithCurrentSubscriptionCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithCumulatedSubscriptionCount adds CumulatedSubscriptionCount (property field)
	WithCumulatedSubscriptionCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithPublishingIntervalCount adds PublishingIntervalCount (property field)
	WithPublishingIntervalCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithSecurityRejectedRequestsCount adds SecurityRejectedRequestsCount (property field)
	WithSecurityRejectedRequestsCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// WithRejectedRequestsCount adds RejectedRequestsCount (property field)
	WithRejectedRequestsCount(uint32) ServerDiagnosticsSummaryDataTypeBuilder
	// Build builds the ServerDiagnosticsSummaryDataType or returns an error if something is wrong
	Build() (ServerDiagnosticsSummaryDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ServerDiagnosticsSummaryDataType
}

// NewServerDiagnosticsSummaryDataTypeBuilder() creates a ServerDiagnosticsSummaryDataTypeBuilder
func NewServerDiagnosticsSummaryDataTypeBuilder() ServerDiagnosticsSummaryDataTypeBuilder {
	return &_ServerDiagnosticsSummaryDataTypeBuilder{_ServerDiagnosticsSummaryDataType: new(_ServerDiagnosticsSummaryDataType)}
}

type _ServerDiagnosticsSummaryDataTypeBuilder struct {
	*_ServerDiagnosticsSummaryDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ServerDiagnosticsSummaryDataTypeBuilder) = (*_ServerDiagnosticsSummaryDataTypeBuilder)(nil)

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithMandatoryFields(serverViewCount uint32, currentSessionCount uint32, cumulatedSessionCount uint32, securityRejectedSessionCount uint32, rejectedSessionCount uint32, sessionTimeoutCount uint32, sessionAbortCount uint32, currentSubscriptionCount uint32, cumulatedSubscriptionCount uint32, publishingIntervalCount uint32, securityRejectedRequestsCount uint32, rejectedRequestsCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	return b.WithServerViewCount(serverViewCount).WithCurrentSessionCount(currentSessionCount).WithCumulatedSessionCount(cumulatedSessionCount).WithSecurityRejectedSessionCount(securityRejectedSessionCount).WithRejectedSessionCount(rejectedSessionCount).WithSessionTimeoutCount(sessionTimeoutCount).WithSessionAbortCount(sessionAbortCount).WithCurrentSubscriptionCount(currentSubscriptionCount).WithCumulatedSubscriptionCount(cumulatedSubscriptionCount).WithPublishingIntervalCount(publishingIntervalCount).WithSecurityRejectedRequestsCount(securityRejectedRequestsCount).WithRejectedRequestsCount(rejectedRequestsCount)
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithServerViewCount(serverViewCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.ServerViewCount = serverViewCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithCurrentSessionCount(currentSessionCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.CurrentSessionCount = currentSessionCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithCumulatedSessionCount(cumulatedSessionCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.CumulatedSessionCount = cumulatedSessionCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithSecurityRejectedSessionCount(securityRejectedSessionCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.SecurityRejectedSessionCount = securityRejectedSessionCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithRejectedSessionCount(rejectedSessionCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.RejectedSessionCount = rejectedSessionCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithSessionTimeoutCount(sessionTimeoutCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.SessionTimeoutCount = sessionTimeoutCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithSessionAbortCount(sessionAbortCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.SessionAbortCount = sessionAbortCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithCurrentSubscriptionCount(currentSubscriptionCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.CurrentSubscriptionCount = currentSubscriptionCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithCumulatedSubscriptionCount(cumulatedSubscriptionCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.CumulatedSubscriptionCount = cumulatedSubscriptionCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithPublishingIntervalCount(publishingIntervalCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.PublishingIntervalCount = publishingIntervalCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithSecurityRejectedRequestsCount(securityRejectedRequestsCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.SecurityRejectedRequestsCount = securityRejectedRequestsCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) WithRejectedRequestsCount(rejectedRequestsCount uint32) ServerDiagnosticsSummaryDataTypeBuilder {
	b.RejectedRequestsCount = rejectedRequestsCount
	return b
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) Build() (ServerDiagnosticsSummaryDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ServerDiagnosticsSummaryDataType.deepCopy(), nil
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) MustBuild() ServerDiagnosticsSummaryDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ServerDiagnosticsSummaryDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ServerDiagnosticsSummaryDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateServerDiagnosticsSummaryDataTypeBuilder().(*_ServerDiagnosticsSummaryDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateServerDiagnosticsSummaryDataTypeBuilder creates a ServerDiagnosticsSummaryDataTypeBuilder
func (b *_ServerDiagnosticsSummaryDataType) CreateServerDiagnosticsSummaryDataTypeBuilder() ServerDiagnosticsSummaryDataTypeBuilder {
	if b == nil {
		return NewServerDiagnosticsSummaryDataTypeBuilder()
	}
	return &_ServerDiagnosticsSummaryDataTypeBuilder{_ServerDiagnosticsSummaryDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServerDiagnosticsSummaryDataType) GetIdentifier() string {
	return "861"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerDiagnosticsSummaryDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServerDiagnosticsSummaryDataType) GetServerViewCount() uint32 {
	return m.ServerViewCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCurrentSessionCount() uint32 {
	return m.CurrentSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCumulatedSessionCount() uint32 {
	return m.CumulatedSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSecurityRejectedSessionCount() uint32 {
	return m.SecurityRejectedSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetRejectedSessionCount() uint32 {
	return m.RejectedSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSessionTimeoutCount() uint32 {
	return m.SessionTimeoutCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSessionAbortCount() uint32 {
	return m.SessionAbortCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCurrentSubscriptionCount() uint32 {
	return m.CurrentSubscriptionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCumulatedSubscriptionCount() uint32 {
	return m.CumulatedSubscriptionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetPublishingIntervalCount() uint32 {
	return m.PublishingIntervalCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSecurityRejectedRequestsCount() uint32 {
	return m.SecurityRejectedRequestsCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetRejectedRequestsCount() uint32 {
	return m.RejectedRequestsCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastServerDiagnosticsSummaryDataType(structType any) ServerDiagnosticsSummaryDataType {
	if casted, ok := structType.(ServerDiagnosticsSummaryDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServerDiagnosticsSummaryDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServerDiagnosticsSummaryDataType) GetTypeName() string {
	return "ServerDiagnosticsSummaryDataType"
}

func (m *_ServerDiagnosticsSummaryDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (serverViewCount)
	lengthInBits += 32

	// Simple field (currentSessionCount)
	lengthInBits += 32

	// Simple field (cumulatedSessionCount)
	lengthInBits += 32

	// Simple field (securityRejectedSessionCount)
	lengthInBits += 32

	// Simple field (rejectedSessionCount)
	lengthInBits += 32

	// Simple field (sessionTimeoutCount)
	lengthInBits += 32

	// Simple field (sessionAbortCount)
	lengthInBits += 32

	// Simple field (currentSubscriptionCount)
	lengthInBits += 32

	// Simple field (cumulatedSubscriptionCount)
	lengthInBits += 32

	// Simple field (publishingIntervalCount)
	lengthInBits += 32

	// Simple field (securityRejectedRequestsCount)
	lengthInBits += 32

	// Simple field (rejectedRequestsCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ServerDiagnosticsSummaryDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServerDiagnosticsSummaryDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__serverDiagnosticsSummaryDataType ServerDiagnosticsSummaryDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServerDiagnosticsSummaryDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerDiagnosticsSummaryDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serverViewCount, err := ReadSimpleField(ctx, "serverViewCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverViewCount' field"))
	}
	m.ServerViewCount = serverViewCount

	currentSessionCount, err := ReadSimpleField(ctx, "currentSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentSessionCount' field"))
	}
	m.CurrentSessionCount = currentSessionCount

	cumulatedSessionCount, err := ReadSimpleField(ctx, "cumulatedSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cumulatedSessionCount' field"))
	}
	m.CumulatedSessionCount = cumulatedSessionCount

	securityRejectedSessionCount, err := ReadSimpleField(ctx, "securityRejectedSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityRejectedSessionCount' field"))
	}
	m.SecurityRejectedSessionCount = securityRejectedSessionCount

	rejectedSessionCount, err := ReadSimpleField(ctx, "rejectedSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectedSessionCount' field"))
	}
	m.RejectedSessionCount = rejectedSessionCount

	sessionTimeoutCount, err := ReadSimpleField(ctx, "sessionTimeoutCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionTimeoutCount' field"))
	}
	m.SessionTimeoutCount = sessionTimeoutCount

	sessionAbortCount, err := ReadSimpleField(ctx, "sessionAbortCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionAbortCount' field"))
	}
	m.SessionAbortCount = sessionAbortCount

	currentSubscriptionCount, err := ReadSimpleField(ctx, "currentSubscriptionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentSubscriptionCount' field"))
	}
	m.CurrentSubscriptionCount = currentSubscriptionCount

	cumulatedSubscriptionCount, err := ReadSimpleField(ctx, "cumulatedSubscriptionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cumulatedSubscriptionCount' field"))
	}
	m.CumulatedSubscriptionCount = cumulatedSubscriptionCount

	publishingIntervalCount, err := ReadSimpleField(ctx, "publishingIntervalCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingIntervalCount' field"))
	}
	m.PublishingIntervalCount = publishingIntervalCount

	securityRejectedRequestsCount, err := ReadSimpleField(ctx, "securityRejectedRequestsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityRejectedRequestsCount' field"))
	}
	m.SecurityRejectedRequestsCount = securityRejectedRequestsCount

	rejectedRequestsCount, err := ReadSimpleField(ctx, "rejectedRequestsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectedRequestsCount' field"))
	}
	m.RejectedRequestsCount = rejectedRequestsCount

	if closeErr := readBuffer.CloseContext("ServerDiagnosticsSummaryDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerDiagnosticsSummaryDataType")
	}

	return m, nil
}

func (m *_ServerDiagnosticsSummaryDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerDiagnosticsSummaryDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerDiagnosticsSummaryDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerDiagnosticsSummaryDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "serverViewCount", m.GetServerViewCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverViewCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "currentSessionCount", m.GetCurrentSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cumulatedSessionCount", m.GetCumulatedSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cumulatedSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "securityRejectedSessionCount", m.GetSecurityRejectedSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityRejectedSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "rejectedSessionCount", m.GetRejectedSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectedSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "sessionTimeoutCount", m.GetSessionTimeoutCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionTimeoutCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "sessionAbortCount", m.GetSessionAbortCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionAbortCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "currentSubscriptionCount", m.GetCurrentSubscriptionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentSubscriptionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cumulatedSubscriptionCount", m.GetCumulatedSubscriptionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cumulatedSubscriptionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "publishingIntervalCount", m.GetPublishingIntervalCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingIntervalCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "securityRejectedRequestsCount", m.GetSecurityRejectedRequestsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityRejectedRequestsCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "rejectedRequestsCount", m.GetRejectedRequestsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectedRequestsCount' field")
		}

		if popErr := writeBuffer.PopContext("ServerDiagnosticsSummaryDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerDiagnosticsSummaryDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerDiagnosticsSummaryDataType) IsServerDiagnosticsSummaryDataType() {}

func (m *_ServerDiagnosticsSummaryDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ServerDiagnosticsSummaryDataType) deepCopy() *_ServerDiagnosticsSummaryDataType {
	if m == nil {
		return nil
	}
	_ServerDiagnosticsSummaryDataTypeCopy := &_ServerDiagnosticsSummaryDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ServerViewCount,
		m.CurrentSessionCount,
		m.CumulatedSessionCount,
		m.SecurityRejectedSessionCount,
		m.RejectedSessionCount,
		m.SessionTimeoutCount,
		m.SessionAbortCount,
		m.CurrentSubscriptionCount,
		m.CumulatedSubscriptionCount,
		m.PublishingIntervalCount,
		m.SecurityRejectedRequestsCount,
		m.RejectedRequestsCount,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ServerDiagnosticsSummaryDataTypeCopy
}

func (m *_ServerDiagnosticsSummaryDataType) String() string {
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
