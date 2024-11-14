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

// FieldMetaData is the corresponding interface of FieldMetaData
type FieldMetaData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetFieldFlags returns FieldFlags (property field)
	GetFieldFlags() DataSetFieldFlags
	// GetBuiltInType returns BuiltInType (property field)
	GetBuiltInType() uint8
	// GetDataType returns DataType (property field)
	GetDataType() NodeId
	// GetValueRank returns ValueRank (property field)
	GetValueRank() int32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() []uint32
	// GetMaxStringLength returns MaxStringLength (property field)
	GetMaxStringLength() uint32
	// GetDataSetFieldId returns DataSetFieldId (property field)
	GetDataSetFieldId() GuidValue
	// GetProperties returns Properties (property field)
	GetProperties() []KeyValuePair
	// IsFieldMetaData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFieldMetaData()
	// CreateBuilder creates a FieldMetaDataBuilder
	CreateFieldMetaDataBuilder() FieldMetaDataBuilder
}

// _FieldMetaData is the data-structure of this message
type _FieldMetaData struct {
	ExtensionObjectDefinitionContract
	Name            PascalString
	Description     LocalizedText
	FieldFlags      DataSetFieldFlags
	BuiltInType     uint8
	DataType        NodeId
	ValueRank       int32
	ArrayDimensions []uint32
	MaxStringLength uint32
	DataSetFieldId  GuidValue
	Properties      []KeyValuePair
}

var _ FieldMetaData = (*_FieldMetaData)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FieldMetaData)(nil)

// NewFieldMetaData factory function for _FieldMetaData
func NewFieldMetaData(name PascalString, description LocalizedText, fieldFlags DataSetFieldFlags, builtInType uint8, dataType NodeId, valueRank int32, arrayDimensions []uint32, maxStringLength uint32, dataSetFieldId GuidValue, properties []KeyValuePair) *_FieldMetaData {
	if name == nil {
		panic("name of type PascalString for FieldMetaData must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for FieldMetaData must not be nil")
	}
	if dataType == nil {
		panic("dataType of type NodeId for FieldMetaData must not be nil")
	}
	if dataSetFieldId == nil {
		panic("dataSetFieldId of type GuidValue for FieldMetaData must not be nil")
	}
	_result := &_FieldMetaData{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Description:                       description,
		FieldFlags:                        fieldFlags,
		BuiltInType:                       builtInType,
		DataType:                          dataType,
		ValueRank:                         valueRank,
		ArrayDimensions:                   arrayDimensions,
		MaxStringLength:                   maxStringLength,
		DataSetFieldId:                    dataSetFieldId,
		Properties:                        properties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FieldMetaDataBuilder is a builder for FieldMetaData
type FieldMetaDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(name PascalString, description LocalizedText, fieldFlags DataSetFieldFlags, builtInType uint8, dataType NodeId, valueRank int32, arrayDimensions []uint32, maxStringLength uint32, dataSetFieldId GuidValue, properties []KeyValuePair) FieldMetaDataBuilder
	// WithName adds Name (property field)
	WithName(PascalString) FieldMetaDataBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) FieldMetaDataBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) FieldMetaDataBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) FieldMetaDataBuilder
	// WithFieldFlags adds FieldFlags (property field)
	WithFieldFlags(DataSetFieldFlags) FieldMetaDataBuilder
	// WithBuiltInType adds BuiltInType (property field)
	WithBuiltInType(uint8) FieldMetaDataBuilder
	// WithDataType adds DataType (property field)
	WithDataType(NodeId) FieldMetaDataBuilder
	// WithDataTypeBuilder adds DataType (property field) which is build by the builder
	WithDataTypeBuilder(func(NodeIdBuilder) NodeIdBuilder) FieldMetaDataBuilder
	// WithValueRank adds ValueRank (property field)
	WithValueRank(int32) FieldMetaDataBuilder
	// WithArrayDimensions adds ArrayDimensions (property field)
	WithArrayDimensions(...uint32) FieldMetaDataBuilder
	// WithMaxStringLength adds MaxStringLength (property field)
	WithMaxStringLength(uint32) FieldMetaDataBuilder
	// WithDataSetFieldId adds DataSetFieldId (property field)
	WithDataSetFieldId(GuidValue) FieldMetaDataBuilder
	// WithDataSetFieldIdBuilder adds DataSetFieldId (property field) which is build by the builder
	WithDataSetFieldIdBuilder(func(GuidValueBuilder) GuidValueBuilder) FieldMetaDataBuilder
	// WithProperties adds Properties (property field)
	WithProperties(...KeyValuePair) FieldMetaDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the FieldMetaData or returns an error if something is wrong
	Build() (FieldMetaData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FieldMetaData
}

// NewFieldMetaDataBuilder() creates a FieldMetaDataBuilder
func NewFieldMetaDataBuilder() FieldMetaDataBuilder {
	return &_FieldMetaDataBuilder{_FieldMetaData: new(_FieldMetaData)}
}

type _FieldMetaDataBuilder struct {
	*_FieldMetaData

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FieldMetaDataBuilder) = (*_FieldMetaDataBuilder)(nil)

func (b *_FieldMetaDataBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._FieldMetaData
}

func (b *_FieldMetaDataBuilder) WithMandatoryFields(name PascalString, description LocalizedText, fieldFlags DataSetFieldFlags, builtInType uint8, dataType NodeId, valueRank int32, arrayDimensions []uint32, maxStringLength uint32, dataSetFieldId GuidValue, properties []KeyValuePair) FieldMetaDataBuilder {
	return b.WithName(name).WithDescription(description).WithFieldFlags(fieldFlags).WithBuiltInType(builtInType).WithDataType(dataType).WithValueRank(valueRank).WithArrayDimensions(arrayDimensions...).WithMaxStringLength(maxStringLength).WithDataSetFieldId(dataSetFieldId).WithProperties(properties...)
}

func (b *_FieldMetaDataBuilder) WithName(name PascalString) FieldMetaDataBuilder {
	b.Name = name
	return b
}

func (b *_FieldMetaDataBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) FieldMetaDataBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_FieldMetaDataBuilder) WithDescription(description LocalizedText) FieldMetaDataBuilder {
	b.Description = description
	return b
}

func (b *_FieldMetaDataBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) FieldMetaDataBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_FieldMetaDataBuilder) WithFieldFlags(fieldFlags DataSetFieldFlags) FieldMetaDataBuilder {
	b.FieldFlags = fieldFlags
	return b
}

func (b *_FieldMetaDataBuilder) WithBuiltInType(builtInType uint8) FieldMetaDataBuilder {
	b.BuiltInType = builtInType
	return b
}

func (b *_FieldMetaDataBuilder) WithDataType(dataType NodeId) FieldMetaDataBuilder {
	b.DataType = dataType
	return b
}

func (b *_FieldMetaDataBuilder) WithDataTypeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) FieldMetaDataBuilder {
	builder := builderSupplier(b.DataType.CreateNodeIdBuilder())
	var err error
	b.DataType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_FieldMetaDataBuilder) WithValueRank(valueRank int32) FieldMetaDataBuilder {
	b.ValueRank = valueRank
	return b
}

func (b *_FieldMetaDataBuilder) WithArrayDimensions(arrayDimensions ...uint32) FieldMetaDataBuilder {
	b.ArrayDimensions = arrayDimensions
	return b
}

func (b *_FieldMetaDataBuilder) WithMaxStringLength(maxStringLength uint32) FieldMetaDataBuilder {
	b.MaxStringLength = maxStringLength
	return b
}

func (b *_FieldMetaDataBuilder) WithDataSetFieldId(dataSetFieldId GuidValue) FieldMetaDataBuilder {
	b.DataSetFieldId = dataSetFieldId
	return b
}

func (b *_FieldMetaDataBuilder) WithDataSetFieldIdBuilder(builderSupplier func(GuidValueBuilder) GuidValueBuilder) FieldMetaDataBuilder {
	builder := builderSupplier(b.DataSetFieldId.CreateGuidValueBuilder())
	var err error
	b.DataSetFieldId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "GuidValueBuilder failed"))
	}
	return b
}

func (b *_FieldMetaDataBuilder) WithProperties(properties ...KeyValuePair) FieldMetaDataBuilder {
	b.Properties = properties
	return b
}

func (b *_FieldMetaDataBuilder) Build() (FieldMetaData, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.DataType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataType' not set"))
	}
	if b.DataSetFieldId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataSetFieldId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FieldMetaData.deepCopy(), nil
}

func (b *_FieldMetaDataBuilder) MustBuild() FieldMetaData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FieldMetaDataBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_FieldMetaDataBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FieldMetaDataBuilder) DeepCopy() any {
	_copy := b.CreateFieldMetaDataBuilder().(*_FieldMetaDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFieldMetaDataBuilder creates a FieldMetaDataBuilder
func (b *_FieldMetaData) CreateFieldMetaDataBuilder() FieldMetaDataBuilder {
	if b == nil {
		return NewFieldMetaDataBuilder()
	}
	return &_FieldMetaDataBuilder{_FieldMetaData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FieldMetaData) GetExtensionId() int32 {
	return int32(14526)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FieldMetaData) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FieldMetaData) GetName() PascalString {
	return m.Name
}

func (m *_FieldMetaData) GetDescription() LocalizedText {
	return m.Description
}

func (m *_FieldMetaData) GetFieldFlags() DataSetFieldFlags {
	return m.FieldFlags
}

func (m *_FieldMetaData) GetBuiltInType() uint8 {
	return m.BuiltInType
}

func (m *_FieldMetaData) GetDataType() NodeId {
	return m.DataType
}

func (m *_FieldMetaData) GetValueRank() int32 {
	return m.ValueRank
}

func (m *_FieldMetaData) GetArrayDimensions() []uint32 {
	return m.ArrayDimensions
}

func (m *_FieldMetaData) GetMaxStringLength() uint32 {
	return m.MaxStringLength
}

func (m *_FieldMetaData) GetDataSetFieldId() GuidValue {
	return m.DataSetFieldId
}

func (m *_FieldMetaData) GetProperties() []KeyValuePair {
	return m.Properties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFieldMetaData(structType any) FieldMetaData {
	if casted, ok := structType.(FieldMetaData); ok {
		return casted
	}
	if casted, ok := structType.(*FieldMetaData); ok {
		return *casted
	}
	return nil
}

func (m *_FieldMetaData) GetTypeName() string {
	return "FieldMetaData"
}

func (m *_FieldMetaData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (fieldFlags)
	lengthInBits += 16

	// Simple field (builtInType)
	lengthInBits += 8

	// Simple field (dataType)
	lengthInBits += m.DataType.GetLengthInBits(ctx)

	// Simple field (valueRank)
	lengthInBits += 32

	// Implicit Field (noOfArrayDimensions)
	lengthInBits += 32

	// Array field
	if len(m.ArrayDimensions) > 0 {
		lengthInBits += 32 * uint16(len(m.ArrayDimensions))
	}

	// Simple field (maxStringLength)
	lengthInBits += 32

	// Simple field (dataSetFieldId)
	lengthInBits += m.DataSetFieldId.GetLengthInBits(ctx)

	// Implicit Field (noOfProperties)
	lengthInBits += 32

	// Array field
	if len(m.Properties) > 0 {
		for _curItem, element := range m.Properties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Properties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FieldMetaData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FieldMetaData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__fieldMetaData FieldMetaData, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FieldMetaData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FieldMetaData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	fieldFlags, err := ReadEnumField[DataSetFieldFlags](ctx, "fieldFlags", "DataSetFieldFlags", ReadEnum(DataSetFieldFlagsByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fieldFlags' field"))
	}
	m.FieldFlags = fieldFlags

	builtInType, err := ReadSimpleField(ctx, "builtInType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'builtInType' field"))
	}
	m.BuiltInType = builtInType

	dataType, err := ReadSimpleField[NodeId](ctx, "dataType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	valueRank, err := ReadSimpleField(ctx, "valueRank", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueRank' field"))
	}
	m.ValueRank = valueRank

	noOfArrayDimensions, err := ReadImplicitField[int32](ctx, "noOfArrayDimensions", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfArrayDimensions' field"))
	}
	_ = noOfArrayDimensions

	arrayDimensions, err := ReadCountArrayField[uint32](ctx, "arrayDimensions", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfArrayDimensions))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensions' field"))
	}
	m.ArrayDimensions = arrayDimensions

	maxStringLength, err := ReadSimpleField(ctx, "maxStringLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxStringLength' field"))
	}
	m.MaxStringLength = maxStringLength

	dataSetFieldId, err := ReadSimpleField[GuidValue](ctx, "dataSetFieldId", ReadComplex[GuidValue](GuidValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetFieldId' field"))
	}
	m.DataSetFieldId = dataSetFieldId

	noOfProperties, err := ReadImplicitField[int32](ctx, "noOfProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfProperties' field"))
	}
	_ = noOfProperties

	properties, err := ReadCountArrayField[KeyValuePair](ctx, "properties", ReadComplex[KeyValuePair](ExtensionObjectDefinitionParseWithBufferProducer[KeyValuePair]((int32)(int32(14535))), readBuffer), uint64(noOfProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'properties' field"))
	}
	m.Properties = properties

	if closeErr := readBuffer.CloseContext("FieldMetaData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FieldMetaData")
	}

	return m, nil
}

func (m *_FieldMetaData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FieldMetaData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FieldMetaData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FieldMetaData")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleEnumField[DataSetFieldFlags](ctx, "fieldFlags", "DataSetFieldFlags", m.GetFieldFlags(), WriteEnum[DataSetFieldFlags, uint16](DataSetFieldFlags.GetValue, DataSetFieldFlags.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
			return errors.Wrap(err, "Error serializing 'fieldFlags' field")
		}

		if err := WriteSimpleField[uint8](ctx, "builtInType", m.GetBuiltInType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'builtInType' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "dataType", m.GetDataType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataType' field")
		}

		if err := WriteSimpleField[int32](ctx, "valueRank", m.GetValueRank(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'valueRank' field")
		}
		noOfArrayDimensions := int32(utils.InlineIf(bool((m.GetArrayDimensions()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetArrayDimensions()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfArrayDimensions", noOfArrayDimensions, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfArrayDimensions' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "arrayDimensions", m.GetArrayDimensions(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayDimensions' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxStringLength", m.GetMaxStringLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxStringLength' field")
		}

		if err := WriteSimpleField[GuidValue](ctx, "dataSetFieldId", m.GetDataSetFieldId(), WriteComplex[GuidValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetFieldId' field")
		}
		noOfProperties := int32(utils.InlineIf(bool((m.GetProperties()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetProperties()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfProperties", noOfProperties, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "properties", m.GetProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'properties' field")
		}

		if popErr := writeBuffer.PopContext("FieldMetaData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FieldMetaData")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FieldMetaData) IsFieldMetaData() {}

func (m *_FieldMetaData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FieldMetaData) deepCopy() *_FieldMetaData {
	if m == nil {
		return nil
	}
	_FieldMetaDataCopy := &_FieldMetaData{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.Name),
		utils.DeepCopy[LocalizedText](m.Description),
		m.FieldFlags,
		m.BuiltInType,
		utils.DeepCopy[NodeId](m.DataType),
		m.ValueRank,
		utils.DeepCopySlice[uint32, uint32](m.ArrayDimensions),
		m.MaxStringLength,
		utils.DeepCopy[GuidValue](m.DataSetFieldId),
		utils.DeepCopySlice[KeyValuePair, KeyValuePair](m.Properties),
	}
	_FieldMetaDataCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FieldMetaDataCopy
}

func (m *_FieldMetaData) String() string {
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
