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

// MonitoringParameters is the corresponding interface of MonitoringParameters
type MonitoringParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetClientHandle returns ClientHandle (property field)
	GetClientHandle() uint32
	// GetSamplingInterval returns SamplingInterval (property field)
	GetSamplingInterval() float64
	// GetFilter returns Filter (property field)
	GetFilter() ExtensionObject
	// GetQueueSize returns QueueSize (property field)
	GetQueueSize() uint32
	// GetDiscardOldest returns DiscardOldest (property field)
	GetDiscardOldest() bool
	// IsMonitoringParameters is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMonitoringParameters()
	// CreateBuilder creates a MonitoringParametersBuilder
	CreateMonitoringParametersBuilder() MonitoringParametersBuilder
}

// _MonitoringParameters is the data-structure of this message
type _MonitoringParameters struct {
	ExtensionObjectDefinitionContract
	ClientHandle     uint32
	SamplingInterval float64
	Filter           ExtensionObject
	QueueSize        uint32
	DiscardOldest    bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ MonitoringParameters = (*_MonitoringParameters)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_MonitoringParameters)(nil)

// NewMonitoringParameters factory function for _MonitoringParameters
func NewMonitoringParameters(clientHandle uint32, samplingInterval float64, filter ExtensionObject, queueSize uint32, discardOldest bool) *_MonitoringParameters {
	if filter == nil {
		panic("filter of type ExtensionObject for MonitoringParameters must not be nil")
	}
	_result := &_MonitoringParameters{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ClientHandle:                      clientHandle,
		SamplingInterval:                  samplingInterval,
		Filter:                            filter,
		QueueSize:                         queueSize,
		DiscardOldest:                     discardOldest,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MonitoringParametersBuilder is a builder for MonitoringParameters
type MonitoringParametersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(clientHandle uint32, samplingInterval float64, filter ExtensionObject, queueSize uint32, discardOldest bool) MonitoringParametersBuilder
	// WithClientHandle adds ClientHandle (property field)
	WithClientHandle(uint32) MonitoringParametersBuilder
	// WithSamplingInterval adds SamplingInterval (property field)
	WithSamplingInterval(float64) MonitoringParametersBuilder
	// WithFilter adds Filter (property field)
	WithFilter(ExtensionObject) MonitoringParametersBuilder
	// WithFilterBuilder adds Filter (property field) which is build by the builder
	WithFilterBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) MonitoringParametersBuilder
	// WithQueueSize adds QueueSize (property field)
	WithQueueSize(uint32) MonitoringParametersBuilder
	// WithDiscardOldest adds DiscardOldest (property field)
	WithDiscardOldest(bool) MonitoringParametersBuilder
	// Build builds the MonitoringParameters or returns an error if something is wrong
	Build() (MonitoringParameters, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MonitoringParameters
}

// NewMonitoringParametersBuilder() creates a MonitoringParametersBuilder
func NewMonitoringParametersBuilder() MonitoringParametersBuilder {
	return &_MonitoringParametersBuilder{_MonitoringParameters: new(_MonitoringParameters)}
}

type _MonitoringParametersBuilder struct {
	*_MonitoringParameters

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (MonitoringParametersBuilder) = (*_MonitoringParametersBuilder)(nil)

func (b *_MonitoringParametersBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_MonitoringParametersBuilder) WithMandatoryFields(clientHandle uint32, samplingInterval float64, filter ExtensionObject, queueSize uint32, discardOldest bool) MonitoringParametersBuilder {
	return b.WithClientHandle(clientHandle).WithSamplingInterval(samplingInterval).WithFilter(filter).WithQueueSize(queueSize).WithDiscardOldest(discardOldest)
}

func (b *_MonitoringParametersBuilder) WithClientHandle(clientHandle uint32) MonitoringParametersBuilder {
	b.ClientHandle = clientHandle
	return b
}

func (b *_MonitoringParametersBuilder) WithSamplingInterval(samplingInterval float64) MonitoringParametersBuilder {
	b.SamplingInterval = samplingInterval
	return b
}

func (b *_MonitoringParametersBuilder) WithFilter(filter ExtensionObject) MonitoringParametersBuilder {
	b.Filter = filter
	return b
}

func (b *_MonitoringParametersBuilder) WithFilterBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) MonitoringParametersBuilder {
	builder := builderSupplier(b.Filter.CreateExtensionObjectBuilder())
	var err error
	b.Filter, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_MonitoringParametersBuilder) WithQueueSize(queueSize uint32) MonitoringParametersBuilder {
	b.QueueSize = queueSize
	return b
}

func (b *_MonitoringParametersBuilder) WithDiscardOldest(discardOldest bool) MonitoringParametersBuilder {
	b.DiscardOldest = discardOldest
	return b
}

func (b *_MonitoringParametersBuilder) Build() (MonitoringParameters, error) {
	if b.Filter == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'filter' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MonitoringParameters.deepCopy(), nil
}

func (b *_MonitoringParametersBuilder) MustBuild() MonitoringParameters {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MonitoringParametersBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_MonitoringParametersBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_MonitoringParametersBuilder) DeepCopy() any {
	_copy := b.CreateMonitoringParametersBuilder().(*_MonitoringParametersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMonitoringParametersBuilder creates a MonitoringParametersBuilder
func (b *_MonitoringParameters) CreateMonitoringParametersBuilder() MonitoringParametersBuilder {
	if b == nil {
		return NewMonitoringParametersBuilder()
	}
	return &_MonitoringParametersBuilder{_MonitoringParameters: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoringParameters) GetExtensionId() int32 {
	return int32(742)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoringParameters) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoringParameters) GetClientHandle() uint32 {
	return m.ClientHandle
}

func (m *_MonitoringParameters) GetSamplingInterval() float64 {
	return m.SamplingInterval
}

func (m *_MonitoringParameters) GetFilter() ExtensionObject {
	return m.Filter
}

func (m *_MonitoringParameters) GetQueueSize() uint32 {
	return m.QueueSize
}

func (m *_MonitoringParameters) GetDiscardOldest() bool {
	return m.DiscardOldest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMonitoringParameters(structType any) MonitoringParameters {
	if casted, ok := structType.(MonitoringParameters); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoringParameters); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoringParameters) GetTypeName() string {
	return "MonitoringParameters"
}

func (m *_MonitoringParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (clientHandle)
	lengthInBits += 32

	// Simple field (samplingInterval)
	lengthInBits += 64

	// Simple field (filter)
	lengthInBits += m.Filter.GetLengthInBits(ctx)

	// Simple field (queueSize)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (discardOldest)
	lengthInBits += 1

	return lengthInBits
}

func (m *_MonitoringParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoringParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__monitoringParameters MonitoringParameters, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoringParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoringParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	clientHandle, err := ReadSimpleField(ctx, "clientHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientHandle' field"))
	}
	m.ClientHandle = clientHandle

	samplingInterval, err := ReadSimpleField(ctx, "samplingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'samplingInterval' field"))
	}
	m.SamplingInterval = samplingInterval

	filter, err := ReadSimpleField[ExtensionObject](ctx, "filter", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'filter' field"))
	}
	m.Filter = filter

	queueSize, err := ReadSimpleField(ctx, "queueSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queueSize' field"))
	}
	m.QueueSize = queueSize

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	discardOldest, err := ReadSimpleField(ctx, "discardOldest", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discardOldest' field"))
	}
	m.DiscardOldest = discardOldest

	if closeErr := readBuffer.CloseContext("MonitoringParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoringParameters")
	}

	return m, nil
}

func (m *_MonitoringParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoringParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoringParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoringParameters")
		}

		if err := WriteSimpleField[uint32](ctx, "clientHandle", m.GetClientHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientHandle' field")
		}

		if err := WriteSimpleField[float64](ctx, "samplingInterval", m.GetSamplingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'samplingInterval' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "filter", m.GetFilter(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'filter' field")
		}

		if err := WriteSimpleField[uint32](ctx, "queueSize", m.GetQueueSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'queueSize' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "discardOldest", m.GetDiscardOldest(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'discardOldest' field")
		}

		if popErr := writeBuffer.PopContext("MonitoringParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoringParameters")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoringParameters) IsMonitoringParameters() {}

func (m *_MonitoringParameters) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MonitoringParameters) deepCopy() *_MonitoringParameters {
	if m == nil {
		return nil
	}
	_MonitoringParametersCopy := &_MonitoringParameters{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ClientHandle,
		m.SamplingInterval,
		m.Filter.DeepCopy().(ExtensionObject),
		m.QueueSize,
		m.DiscardOldest,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _MonitoringParametersCopy
}

func (m *_MonitoringParameters) String() string {
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
