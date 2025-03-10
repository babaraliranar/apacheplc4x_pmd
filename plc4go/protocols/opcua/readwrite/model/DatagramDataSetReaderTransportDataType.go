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

// DatagramDataSetReaderTransportDataType is the corresponding interface of DatagramDataSetReaderTransportDataType
type DatagramDataSetReaderTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetAddress returns Address (property field)
	GetAddress() ExtensionObject
	// GetQosCategory returns QosCategory (property field)
	GetQosCategory() PascalString
	// GetDatagramQos returns DatagramQos (property field)
	GetDatagramQos() []ExtensionObject
	// GetTopic returns Topic (property field)
	GetTopic() PascalString
	// IsDatagramDataSetReaderTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDatagramDataSetReaderTransportDataType()
	// CreateBuilder creates a DatagramDataSetReaderTransportDataTypeBuilder
	CreateDatagramDataSetReaderTransportDataTypeBuilder() DatagramDataSetReaderTransportDataTypeBuilder
}

// _DatagramDataSetReaderTransportDataType is the data-structure of this message
type _DatagramDataSetReaderTransportDataType struct {
	ExtensionObjectDefinitionContract
	Address     ExtensionObject
	QosCategory PascalString
	DatagramQos []ExtensionObject
	Topic       PascalString
}

var _ DatagramDataSetReaderTransportDataType = (*_DatagramDataSetReaderTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DatagramDataSetReaderTransportDataType)(nil)

// NewDatagramDataSetReaderTransportDataType factory function for _DatagramDataSetReaderTransportDataType
func NewDatagramDataSetReaderTransportDataType(address ExtensionObject, qosCategory PascalString, datagramQos []ExtensionObject, topic PascalString) *_DatagramDataSetReaderTransportDataType {
	if address == nil {
		panic("address of type ExtensionObject for DatagramDataSetReaderTransportDataType must not be nil")
	}
	if qosCategory == nil {
		panic("qosCategory of type PascalString for DatagramDataSetReaderTransportDataType must not be nil")
	}
	if topic == nil {
		panic("topic of type PascalString for DatagramDataSetReaderTransportDataType must not be nil")
	}
	_result := &_DatagramDataSetReaderTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Address:                           address,
		QosCategory:                       qosCategory,
		DatagramQos:                       datagramQos,
		Topic:                             topic,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DatagramDataSetReaderTransportDataTypeBuilder is a builder for DatagramDataSetReaderTransportDataType
type DatagramDataSetReaderTransportDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address ExtensionObject, qosCategory PascalString, datagramQos []ExtensionObject, topic PascalString) DatagramDataSetReaderTransportDataTypeBuilder
	// WithAddress adds Address (property field)
	WithAddress(ExtensionObject) DatagramDataSetReaderTransportDataTypeBuilder
	// WithAddressBuilder adds Address (property field) which is build by the builder
	WithAddressBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) DatagramDataSetReaderTransportDataTypeBuilder
	// WithQosCategory adds QosCategory (property field)
	WithQosCategory(PascalString) DatagramDataSetReaderTransportDataTypeBuilder
	// WithQosCategoryBuilder adds QosCategory (property field) which is build by the builder
	WithQosCategoryBuilder(func(PascalStringBuilder) PascalStringBuilder) DatagramDataSetReaderTransportDataTypeBuilder
	// WithDatagramQos adds DatagramQos (property field)
	WithDatagramQos(...ExtensionObject) DatagramDataSetReaderTransportDataTypeBuilder
	// WithTopic adds Topic (property field)
	WithTopic(PascalString) DatagramDataSetReaderTransportDataTypeBuilder
	// WithTopicBuilder adds Topic (property field) which is build by the builder
	WithTopicBuilder(func(PascalStringBuilder) PascalStringBuilder) DatagramDataSetReaderTransportDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DatagramDataSetReaderTransportDataType or returns an error if something is wrong
	Build() (DatagramDataSetReaderTransportDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DatagramDataSetReaderTransportDataType
}

// NewDatagramDataSetReaderTransportDataTypeBuilder() creates a DatagramDataSetReaderTransportDataTypeBuilder
func NewDatagramDataSetReaderTransportDataTypeBuilder() DatagramDataSetReaderTransportDataTypeBuilder {
	return &_DatagramDataSetReaderTransportDataTypeBuilder{_DatagramDataSetReaderTransportDataType: new(_DatagramDataSetReaderTransportDataType)}
}

type _DatagramDataSetReaderTransportDataTypeBuilder struct {
	*_DatagramDataSetReaderTransportDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DatagramDataSetReaderTransportDataTypeBuilder) = (*_DatagramDataSetReaderTransportDataTypeBuilder)(nil)

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._DatagramDataSetReaderTransportDataType
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithMandatoryFields(address ExtensionObject, qosCategory PascalString, datagramQos []ExtensionObject, topic PascalString) DatagramDataSetReaderTransportDataTypeBuilder {
	return b.WithAddress(address).WithQosCategory(qosCategory).WithDatagramQos(datagramQos...).WithTopic(topic)
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithAddress(address ExtensionObject) DatagramDataSetReaderTransportDataTypeBuilder {
	b.Address = address
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithAddressBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) DatagramDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.Address.CreateExtensionObjectBuilder())
	var err error
	b.Address, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithQosCategory(qosCategory PascalString) DatagramDataSetReaderTransportDataTypeBuilder {
	b.QosCategory = qosCategory
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithQosCategoryBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) DatagramDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.QosCategory.CreatePascalStringBuilder())
	var err error
	b.QosCategory, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithDatagramQos(datagramQos ...ExtensionObject) DatagramDataSetReaderTransportDataTypeBuilder {
	b.DatagramQos = datagramQos
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithTopic(topic PascalString) DatagramDataSetReaderTransportDataTypeBuilder {
	b.Topic = topic
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) WithTopicBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) DatagramDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.Topic.CreatePascalStringBuilder())
	var err error
	b.Topic, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) Build() (DatagramDataSetReaderTransportDataType, error) {
	if b.Address == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'address' not set"))
	}
	if b.QosCategory == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'qosCategory' not set"))
	}
	if b.Topic == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'topic' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DatagramDataSetReaderTransportDataType.deepCopy(), nil
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) MustBuild() DatagramDataSetReaderTransportDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DatagramDataSetReaderTransportDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateDatagramDataSetReaderTransportDataTypeBuilder().(*_DatagramDataSetReaderTransportDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDatagramDataSetReaderTransportDataTypeBuilder creates a DatagramDataSetReaderTransportDataTypeBuilder
func (b *_DatagramDataSetReaderTransportDataType) CreateDatagramDataSetReaderTransportDataTypeBuilder() DatagramDataSetReaderTransportDataTypeBuilder {
	if b == nil {
		return NewDatagramDataSetReaderTransportDataTypeBuilder()
	}
	return &_DatagramDataSetReaderTransportDataTypeBuilder{_DatagramDataSetReaderTransportDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DatagramDataSetReaderTransportDataType) GetExtensionId() int32 {
	return int32(23616)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DatagramDataSetReaderTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DatagramDataSetReaderTransportDataType) GetAddress() ExtensionObject {
	return m.Address
}

func (m *_DatagramDataSetReaderTransportDataType) GetQosCategory() PascalString {
	return m.QosCategory
}

func (m *_DatagramDataSetReaderTransportDataType) GetDatagramQos() []ExtensionObject {
	return m.DatagramQos
}

func (m *_DatagramDataSetReaderTransportDataType) GetTopic() PascalString {
	return m.Topic
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDatagramDataSetReaderTransportDataType(structType any) DatagramDataSetReaderTransportDataType {
	if casted, ok := structType.(DatagramDataSetReaderTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DatagramDataSetReaderTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DatagramDataSetReaderTransportDataType) GetTypeName() string {
	return "DatagramDataSetReaderTransportDataType"
}

func (m *_DatagramDataSetReaderTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	// Simple field (qosCategory)
	lengthInBits += m.QosCategory.GetLengthInBits(ctx)

	// Implicit Field (noOfDatagramQos)
	lengthInBits += 32

	// Array field
	if len(m.DatagramQos) > 0 {
		for _curItem, element := range m.DatagramQos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DatagramQos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (topic)
	lengthInBits += m.Topic.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DatagramDataSetReaderTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DatagramDataSetReaderTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__datagramDataSetReaderTransportDataType DatagramDataSetReaderTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DatagramDataSetReaderTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DatagramDataSetReaderTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField[ExtensionObject](ctx, "address", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	qosCategory, err := ReadSimpleField[PascalString](ctx, "qosCategory", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'qosCategory' field"))
	}
	m.QosCategory = qosCategory

	noOfDatagramQos, err := ReadImplicitField[int32](ctx, "noOfDatagramQos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDatagramQos' field"))
	}
	_ = noOfDatagramQos

	datagramQos, err := ReadCountArrayField[ExtensionObject](ctx, "datagramQos", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer), uint64(noOfDatagramQos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'datagramQos' field"))
	}
	m.DatagramQos = datagramQos

	topic, err := ReadSimpleField[PascalString](ctx, "topic", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'topic' field"))
	}
	m.Topic = topic

	if closeErr := readBuffer.CloseContext("DatagramDataSetReaderTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DatagramDataSetReaderTransportDataType")
	}

	return m, nil
}

func (m *_DatagramDataSetReaderTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DatagramDataSetReaderTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DatagramDataSetReaderTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DatagramDataSetReaderTransportDataType")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "address", m.GetAddress(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "qosCategory", m.GetQosCategory(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'qosCategory' field")
		}
		noOfDatagramQos := int32(utils.InlineIf(bool((m.GetDatagramQos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDatagramQos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDatagramQos", noOfDatagramQos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDatagramQos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "datagramQos", m.GetDatagramQos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'datagramQos' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "topic", m.GetTopic(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'topic' field")
		}

		if popErr := writeBuffer.PopContext("DatagramDataSetReaderTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DatagramDataSetReaderTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DatagramDataSetReaderTransportDataType) IsDatagramDataSetReaderTransportDataType() {}

func (m *_DatagramDataSetReaderTransportDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DatagramDataSetReaderTransportDataType) deepCopy() *_DatagramDataSetReaderTransportDataType {
	if m == nil {
		return nil
	}
	_DatagramDataSetReaderTransportDataTypeCopy := &_DatagramDataSetReaderTransportDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ExtensionObject](m.Address),
		utils.DeepCopy[PascalString](m.QosCategory),
		utils.DeepCopySlice[ExtensionObject, ExtensionObject](m.DatagramQos),
		utils.DeepCopy[PascalString](m.Topic),
	}
	_DatagramDataSetReaderTransportDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DatagramDataSetReaderTransportDataTypeCopy
}

func (m *_DatagramDataSetReaderTransportDataType) String() string {
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
