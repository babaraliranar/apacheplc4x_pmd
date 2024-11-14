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

// UadpDataSetReaderMessageDataType is the corresponding interface of UadpDataSetReaderMessageDataType
type UadpDataSetReaderMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetGroupVersion returns GroupVersion (property field)
	GetGroupVersion() uint32
	// GetNetworkMessageNumber returns NetworkMessageNumber (property field)
	GetNetworkMessageNumber() uint16
	// GetDataSetOffset returns DataSetOffset (property field)
	GetDataSetOffset() uint16
	// GetDataSetClassId returns DataSetClassId (property field)
	GetDataSetClassId() GuidValue
	// GetNetworkMessageContentMask returns NetworkMessageContentMask (property field)
	GetNetworkMessageContentMask() UadpNetworkMessageContentMask
	// GetDataSetMessageContentMask returns DataSetMessageContentMask (property field)
	GetDataSetMessageContentMask() UadpDataSetMessageContentMask
	// GetPublishingInterval returns PublishingInterval (property field)
	GetPublishingInterval() float64
	// GetReceiveOffset returns ReceiveOffset (property field)
	GetReceiveOffset() float64
	// GetProcessingOffset returns ProcessingOffset (property field)
	GetProcessingOffset() float64
	// IsUadpDataSetReaderMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUadpDataSetReaderMessageDataType()
	// CreateBuilder creates a UadpDataSetReaderMessageDataTypeBuilder
	CreateUadpDataSetReaderMessageDataTypeBuilder() UadpDataSetReaderMessageDataTypeBuilder
}

// _UadpDataSetReaderMessageDataType is the data-structure of this message
type _UadpDataSetReaderMessageDataType struct {
	ExtensionObjectDefinitionContract
	GroupVersion              uint32
	NetworkMessageNumber      uint16
	DataSetOffset             uint16
	DataSetClassId            GuidValue
	NetworkMessageContentMask UadpNetworkMessageContentMask
	DataSetMessageContentMask UadpDataSetMessageContentMask
	PublishingInterval        float64
	ReceiveOffset             float64
	ProcessingOffset          float64
}

var _ UadpDataSetReaderMessageDataType = (*_UadpDataSetReaderMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UadpDataSetReaderMessageDataType)(nil)

// NewUadpDataSetReaderMessageDataType factory function for _UadpDataSetReaderMessageDataType
func NewUadpDataSetReaderMessageDataType(groupVersion uint32, networkMessageNumber uint16, dataSetOffset uint16, dataSetClassId GuidValue, networkMessageContentMask UadpNetworkMessageContentMask, dataSetMessageContentMask UadpDataSetMessageContentMask, publishingInterval float64, receiveOffset float64, processingOffset float64) *_UadpDataSetReaderMessageDataType {
	if dataSetClassId == nil {
		panic("dataSetClassId of type GuidValue for UadpDataSetReaderMessageDataType must not be nil")
	}
	_result := &_UadpDataSetReaderMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		GroupVersion:                      groupVersion,
		NetworkMessageNumber:              networkMessageNumber,
		DataSetOffset:                     dataSetOffset,
		DataSetClassId:                    dataSetClassId,
		NetworkMessageContentMask:         networkMessageContentMask,
		DataSetMessageContentMask:         dataSetMessageContentMask,
		PublishingInterval:                publishingInterval,
		ReceiveOffset:                     receiveOffset,
		ProcessingOffset:                  processingOffset,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UadpDataSetReaderMessageDataTypeBuilder is a builder for UadpDataSetReaderMessageDataType
type UadpDataSetReaderMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(groupVersion uint32, networkMessageNumber uint16, dataSetOffset uint16, dataSetClassId GuidValue, networkMessageContentMask UadpNetworkMessageContentMask, dataSetMessageContentMask UadpDataSetMessageContentMask, publishingInterval float64, receiveOffset float64, processingOffset float64) UadpDataSetReaderMessageDataTypeBuilder
	// WithGroupVersion adds GroupVersion (property field)
	WithGroupVersion(uint32) UadpDataSetReaderMessageDataTypeBuilder
	// WithNetworkMessageNumber adds NetworkMessageNumber (property field)
	WithNetworkMessageNumber(uint16) UadpDataSetReaderMessageDataTypeBuilder
	// WithDataSetOffset adds DataSetOffset (property field)
	WithDataSetOffset(uint16) UadpDataSetReaderMessageDataTypeBuilder
	// WithDataSetClassId adds DataSetClassId (property field)
	WithDataSetClassId(GuidValue) UadpDataSetReaderMessageDataTypeBuilder
	// WithDataSetClassIdBuilder adds DataSetClassId (property field) which is build by the builder
	WithDataSetClassIdBuilder(func(GuidValueBuilder) GuidValueBuilder) UadpDataSetReaderMessageDataTypeBuilder
	// WithNetworkMessageContentMask adds NetworkMessageContentMask (property field)
	WithNetworkMessageContentMask(UadpNetworkMessageContentMask) UadpDataSetReaderMessageDataTypeBuilder
	// WithDataSetMessageContentMask adds DataSetMessageContentMask (property field)
	WithDataSetMessageContentMask(UadpDataSetMessageContentMask) UadpDataSetReaderMessageDataTypeBuilder
	// WithPublishingInterval adds PublishingInterval (property field)
	WithPublishingInterval(float64) UadpDataSetReaderMessageDataTypeBuilder
	// WithReceiveOffset adds ReceiveOffset (property field)
	WithReceiveOffset(float64) UadpDataSetReaderMessageDataTypeBuilder
	// WithProcessingOffset adds ProcessingOffset (property field)
	WithProcessingOffset(float64) UadpDataSetReaderMessageDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the UadpDataSetReaderMessageDataType or returns an error if something is wrong
	Build() (UadpDataSetReaderMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UadpDataSetReaderMessageDataType
}

// NewUadpDataSetReaderMessageDataTypeBuilder() creates a UadpDataSetReaderMessageDataTypeBuilder
func NewUadpDataSetReaderMessageDataTypeBuilder() UadpDataSetReaderMessageDataTypeBuilder {
	return &_UadpDataSetReaderMessageDataTypeBuilder{_UadpDataSetReaderMessageDataType: new(_UadpDataSetReaderMessageDataType)}
}

type _UadpDataSetReaderMessageDataTypeBuilder struct {
	*_UadpDataSetReaderMessageDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UadpDataSetReaderMessageDataTypeBuilder) = (*_UadpDataSetReaderMessageDataTypeBuilder)(nil)

func (b *_UadpDataSetReaderMessageDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._UadpDataSetReaderMessageDataType
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithMandatoryFields(groupVersion uint32, networkMessageNumber uint16, dataSetOffset uint16, dataSetClassId GuidValue, networkMessageContentMask UadpNetworkMessageContentMask, dataSetMessageContentMask UadpDataSetMessageContentMask, publishingInterval float64, receiveOffset float64, processingOffset float64) UadpDataSetReaderMessageDataTypeBuilder {
	return b.WithGroupVersion(groupVersion).WithNetworkMessageNumber(networkMessageNumber).WithDataSetOffset(dataSetOffset).WithDataSetClassId(dataSetClassId).WithNetworkMessageContentMask(networkMessageContentMask).WithDataSetMessageContentMask(dataSetMessageContentMask).WithPublishingInterval(publishingInterval).WithReceiveOffset(receiveOffset).WithProcessingOffset(processingOffset)
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithGroupVersion(groupVersion uint32) UadpDataSetReaderMessageDataTypeBuilder {
	b.GroupVersion = groupVersion
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithNetworkMessageNumber(networkMessageNumber uint16) UadpDataSetReaderMessageDataTypeBuilder {
	b.NetworkMessageNumber = networkMessageNumber
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithDataSetOffset(dataSetOffset uint16) UadpDataSetReaderMessageDataTypeBuilder {
	b.DataSetOffset = dataSetOffset
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithDataSetClassId(dataSetClassId GuidValue) UadpDataSetReaderMessageDataTypeBuilder {
	b.DataSetClassId = dataSetClassId
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithDataSetClassIdBuilder(builderSupplier func(GuidValueBuilder) GuidValueBuilder) UadpDataSetReaderMessageDataTypeBuilder {
	builder := builderSupplier(b.DataSetClassId.CreateGuidValueBuilder())
	var err error
	b.DataSetClassId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "GuidValueBuilder failed"))
	}
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithNetworkMessageContentMask(networkMessageContentMask UadpNetworkMessageContentMask) UadpDataSetReaderMessageDataTypeBuilder {
	b.NetworkMessageContentMask = networkMessageContentMask
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithDataSetMessageContentMask(dataSetMessageContentMask UadpDataSetMessageContentMask) UadpDataSetReaderMessageDataTypeBuilder {
	b.DataSetMessageContentMask = dataSetMessageContentMask
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithPublishingInterval(publishingInterval float64) UadpDataSetReaderMessageDataTypeBuilder {
	b.PublishingInterval = publishingInterval
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithReceiveOffset(receiveOffset float64) UadpDataSetReaderMessageDataTypeBuilder {
	b.ReceiveOffset = receiveOffset
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) WithProcessingOffset(processingOffset float64) UadpDataSetReaderMessageDataTypeBuilder {
	b.ProcessingOffset = processingOffset
	return b
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) Build() (UadpDataSetReaderMessageDataType, error) {
	if b.DataSetClassId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataSetClassId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UadpDataSetReaderMessageDataType.deepCopy(), nil
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) MustBuild() UadpDataSetReaderMessageDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UadpDataSetReaderMessageDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateUadpDataSetReaderMessageDataTypeBuilder().(*_UadpDataSetReaderMessageDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUadpDataSetReaderMessageDataTypeBuilder creates a UadpDataSetReaderMessageDataTypeBuilder
func (b *_UadpDataSetReaderMessageDataType) CreateUadpDataSetReaderMessageDataTypeBuilder() UadpDataSetReaderMessageDataTypeBuilder {
	if b == nil {
		return NewUadpDataSetReaderMessageDataTypeBuilder()
	}
	return &_UadpDataSetReaderMessageDataTypeBuilder{_UadpDataSetReaderMessageDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UadpDataSetReaderMessageDataType) GetExtensionId() int32 {
	return int32(15655)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UadpDataSetReaderMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UadpDataSetReaderMessageDataType) GetGroupVersion() uint32 {
	return m.GroupVersion
}

func (m *_UadpDataSetReaderMessageDataType) GetNetworkMessageNumber() uint16 {
	return m.NetworkMessageNumber
}

func (m *_UadpDataSetReaderMessageDataType) GetDataSetOffset() uint16 {
	return m.DataSetOffset
}

func (m *_UadpDataSetReaderMessageDataType) GetDataSetClassId() GuidValue {
	return m.DataSetClassId
}

func (m *_UadpDataSetReaderMessageDataType) GetNetworkMessageContentMask() UadpNetworkMessageContentMask {
	return m.NetworkMessageContentMask
}

func (m *_UadpDataSetReaderMessageDataType) GetDataSetMessageContentMask() UadpDataSetMessageContentMask {
	return m.DataSetMessageContentMask
}

func (m *_UadpDataSetReaderMessageDataType) GetPublishingInterval() float64 {
	return m.PublishingInterval
}

func (m *_UadpDataSetReaderMessageDataType) GetReceiveOffset() float64 {
	return m.ReceiveOffset
}

func (m *_UadpDataSetReaderMessageDataType) GetProcessingOffset() float64 {
	return m.ProcessingOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUadpDataSetReaderMessageDataType(structType any) UadpDataSetReaderMessageDataType {
	if casted, ok := structType.(UadpDataSetReaderMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*UadpDataSetReaderMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_UadpDataSetReaderMessageDataType) GetTypeName() string {
	return "UadpDataSetReaderMessageDataType"
}

func (m *_UadpDataSetReaderMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (groupVersion)
	lengthInBits += 32

	// Simple field (networkMessageNumber)
	lengthInBits += 16

	// Simple field (dataSetOffset)
	lengthInBits += 16

	// Simple field (dataSetClassId)
	lengthInBits += m.DataSetClassId.GetLengthInBits(ctx)

	// Simple field (networkMessageContentMask)
	lengthInBits += 32

	// Simple field (dataSetMessageContentMask)
	lengthInBits += 32

	// Simple field (publishingInterval)
	lengthInBits += 64

	// Simple field (receiveOffset)
	lengthInBits += 64

	// Simple field (processingOffset)
	lengthInBits += 64

	return lengthInBits
}

func (m *_UadpDataSetReaderMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UadpDataSetReaderMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__uadpDataSetReaderMessageDataType UadpDataSetReaderMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UadpDataSetReaderMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UadpDataSetReaderMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	groupVersion, err := ReadSimpleField(ctx, "groupVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupVersion' field"))
	}
	m.GroupVersion = groupVersion

	networkMessageNumber, err := ReadSimpleField(ctx, "networkMessageNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkMessageNumber' field"))
	}
	m.NetworkMessageNumber = networkMessageNumber

	dataSetOffset, err := ReadSimpleField(ctx, "dataSetOffset", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetOffset' field"))
	}
	m.DataSetOffset = dataSetOffset

	dataSetClassId, err := ReadSimpleField[GuidValue](ctx, "dataSetClassId", ReadComplex[GuidValue](GuidValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetClassId' field"))
	}
	m.DataSetClassId = dataSetClassId

	networkMessageContentMask, err := ReadEnumField[UadpNetworkMessageContentMask](ctx, "networkMessageContentMask", "UadpNetworkMessageContentMask", ReadEnum(UadpNetworkMessageContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkMessageContentMask' field"))
	}
	m.NetworkMessageContentMask = networkMessageContentMask

	dataSetMessageContentMask, err := ReadEnumField[UadpDataSetMessageContentMask](ctx, "dataSetMessageContentMask", "UadpDataSetMessageContentMask", ReadEnum(UadpDataSetMessageContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetMessageContentMask' field"))
	}
	m.DataSetMessageContentMask = dataSetMessageContentMask

	publishingInterval, err := ReadSimpleField(ctx, "publishingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingInterval' field"))
	}
	m.PublishingInterval = publishingInterval

	receiveOffset, err := ReadSimpleField(ctx, "receiveOffset", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiveOffset' field"))
	}
	m.ReceiveOffset = receiveOffset

	processingOffset, err := ReadSimpleField(ctx, "processingOffset", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processingOffset' field"))
	}
	m.ProcessingOffset = processingOffset

	if closeErr := readBuffer.CloseContext("UadpDataSetReaderMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UadpDataSetReaderMessageDataType")
	}

	return m, nil
}

func (m *_UadpDataSetReaderMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UadpDataSetReaderMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UadpDataSetReaderMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UadpDataSetReaderMessageDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "groupVersion", m.GetGroupVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'groupVersion' field")
		}

		if err := WriteSimpleField[uint16](ctx, "networkMessageNumber", m.GetNetworkMessageNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkMessageNumber' field")
		}

		if err := WriteSimpleField[uint16](ctx, "dataSetOffset", m.GetDataSetOffset(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetOffset' field")
		}

		if err := WriteSimpleField[GuidValue](ctx, "dataSetClassId", m.GetDataSetClassId(), WriteComplex[GuidValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetClassId' field")
		}

		if err := WriteSimpleEnumField[UadpNetworkMessageContentMask](ctx, "networkMessageContentMask", "UadpNetworkMessageContentMask", m.GetNetworkMessageContentMask(), WriteEnum[UadpNetworkMessageContentMask, uint32](UadpNetworkMessageContentMask.GetValue, UadpNetworkMessageContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'networkMessageContentMask' field")
		}

		if err := WriteSimpleEnumField[UadpDataSetMessageContentMask](ctx, "dataSetMessageContentMask", "UadpDataSetMessageContentMask", m.GetDataSetMessageContentMask(), WriteEnum[UadpDataSetMessageContentMask, uint32](UadpDataSetMessageContentMask.GetValue, UadpDataSetMessageContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetMessageContentMask' field")
		}

		if err := WriteSimpleField[float64](ctx, "publishingInterval", m.GetPublishingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingInterval' field")
		}

		if err := WriteSimpleField[float64](ctx, "receiveOffset", m.GetReceiveOffset(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'receiveOffset' field")
		}

		if err := WriteSimpleField[float64](ctx, "processingOffset", m.GetProcessingOffset(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'processingOffset' field")
		}

		if popErr := writeBuffer.PopContext("UadpDataSetReaderMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UadpDataSetReaderMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UadpDataSetReaderMessageDataType) IsUadpDataSetReaderMessageDataType() {}

func (m *_UadpDataSetReaderMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UadpDataSetReaderMessageDataType) deepCopy() *_UadpDataSetReaderMessageDataType {
	if m == nil {
		return nil
	}
	_UadpDataSetReaderMessageDataTypeCopy := &_UadpDataSetReaderMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.GroupVersion,
		m.NetworkMessageNumber,
		m.DataSetOffset,
		utils.DeepCopy[GuidValue](m.DataSetClassId),
		m.NetworkMessageContentMask,
		m.DataSetMessageContentMask,
		m.PublishingInterval,
		m.ReceiveOffset,
		m.ProcessingOffset,
	}
	_UadpDataSetReaderMessageDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UadpDataSetReaderMessageDataTypeCopy
}

func (m *_UadpDataSetReaderMessageDataType) String() string {
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
