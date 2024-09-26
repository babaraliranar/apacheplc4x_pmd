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

// DataSetWriterDataType is the corresponding interface of DataSetWriterDataType
type DataSetWriterDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetDataSetWriterId returns DataSetWriterId (property field)
	GetDataSetWriterId() uint16
	// GetDataSetFieldContentMask returns DataSetFieldContentMask (property field)
	GetDataSetFieldContentMask() DataSetFieldContentMask
	// GetKeyFrameCount returns KeyFrameCount (property field)
	GetKeyFrameCount() uint32
	// GetDataSetName returns DataSetName (property field)
	GetDataSetName() PascalString
	// GetNoOfDataSetWriterProperties returns NoOfDataSetWriterProperties (property field)
	GetNoOfDataSetWriterProperties() int32
	// GetDataSetWriterProperties returns DataSetWriterProperties (property field)
	GetDataSetWriterProperties() []ExtensionObjectDefinition
	// GetTransportSettings returns TransportSettings (property field)
	GetTransportSettings() ExtensionObject
	// GetMessageSettings returns MessageSettings (property field)
	GetMessageSettings() ExtensionObject
	// IsDataSetWriterDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSetWriterDataType()
}

// _DataSetWriterDataType is the data-structure of this message
type _DataSetWriterDataType struct {
	ExtensionObjectDefinitionContract
	Name                        PascalString
	Enabled                     bool
	DataSetWriterId             uint16
	DataSetFieldContentMask     DataSetFieldContentMask
	KeyFrameCount               uint32
	DataSetName                 PascalString
	NoOfDataSetWriterProperties int32
	DataSetWriterProperties     []ExtensionObjectDefinition
	TransportSettings           ExtensionObject
	MessageSettings             ExtensionObject
	// Reserved Fields
	reservedField0 *uint8
}

var _ DataSetWriterDataType = (*_DataSetWriterDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetWriterDataType)(nil)

// NewDataSetWriterDataType factory function for _DataSetWriterDataType
func NewDataSetWriterDataType(name PascalString, enabled bool, dataSetWriterId uint16, dataSetFieldContentMask DataSetFieldContentMask, keyFrameCount uint32, dataSetName PascalString, noOfDataSetWriterProperties int32, dataSetWriterProperties []ExtensionObjectDefinition, transportSettings ExtensionObject, messageSettings ExtensionObject) *_DataSetWriterDataType {
	if name == nil {
		panic("name of type PascalString for DataSetWriterDataType must not be nil")
	}
	if dataSetName == nil {
		panic("dataSetName of type PascalString for DataSetWriterDataType must not be nil")
	}
	if transportSettings == nil {
		panic("transportSettings of type ExtensionObject for DataSetWriterDataType must not be nil")
	}
	if messageSettings == nil {
		panic("messageSettings of type ExtensionObject for DataSetWriterDataType must not be nil")
	}
	_result := &_DataSetWriterDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Enabled:                           enabled,
		DataSetWriterId:                   dataSetWriterId,
		DataSetFieldContentMask:           dataSetFieldContentMask,
		KeyFrameCount:                     keyFrameCount,
		DataSetName:                       dataSetName,
		NoOfDataSetWriterProperties:       noOfDataSetWriterProperties,
		DataSetWriterProperties:           dataSetWriterProperties,
		TransportSettings:                 transportSettings,
		MessageSettings:                   messageSettings,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetWriterDataType) GetIdentifier() string {
	return "15599"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetWriterDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataSetWriterDataType) GetName() PascalString {
	return m.Name
}

func (m *_DataSetWriterDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_DataSetWriterDataType) GetDataSetWriterId() uint16 {
	return m.DataSetWriterId
}

func (m *_DataSetWriterDataType) GetDataSetFieldContentMask() DataSetFieldContentMask {
	return m.DataSetFieldContentMask
}

func (m *_DataSetWriterDataType) GetKeyFrameCount() uint32 {
	return m.KeyFrameCount
}

func (m *_DataSetWriterDataType) GetDataSetName() PascalString {
	return m.DataSetName
}

func (m *_DataSetWriterDataType) GetNoOfDataSetWriterProperties() int32 {
	return m.NoOfDataSetWriterProperties
}

func (m *_DataSetWriterDataType) GetDataSetWriterProperties() []ExtensionObjectDefinition {
	return m.DataSetWriterProperties
}

func (m *_DataSetWriterDataType) GetTransportSettings() ExtensionObject {
	return m.TransportSettings
}

func (m *_DataSetWriterDataType) GetMessageSettings() ExtensionObject {
	return m.MessageSettings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataSetWriterDataType(structType any) DataSetWriterDataType {
	if casted, ok := structType.(DataSetWriterDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetWriterDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetWriterDataType) GetTypeName() string {
	return "DataSetWriterDataType"
}

func (m *_DataSetWriterDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (dataSetWriterId)
	lengthInBits += 16

	// Simple field (dataSetFieldContentMask)
	lengthInBits += 32

	// Simple field (keyFrameCount)
	lengthInBits += 32

	// Simple field (dataSetName)
	lengthInBits += m.DataSetName.GetLengthInBits(ctx)

	// Simple field (noOfDataSetWriterProperties)
	lengthInBits += 32

	// Array field
	if len(m.DataSetWriterProperties) > 0 {
		for _curItem, element := range m.DataSetWriterProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataSetWriterProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportSettings)
	lengthInBits += m.TransportSettings.GetLengthInBits(ctx)

	// Simple field (messageSettings)
	lengthInBits += m.MessageSettings.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataSetWriterDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetWriterDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataSetWriterDataType DataSetWriterDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetWriterDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetWriterDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	dataSetWriterId, err := ReadSimpleField(ctx, "dataSetWriterId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetWriterId' field"))
	}
	m.DataSetWriterId = dataSetWriterId

	dataSetFieldContentMask, err := ReadEnumField[DataSetFieldContentMask](ctx, "dataSetFieldContentMask", "DataSetFieldContentMask", ReadEnum(DataSetFieldContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetFieldContentMask' field"))
	}
	m.DataSetFieldContentMask = dataSetFieldContentMask

	keyFrameCount, err := ReadSimpleField(ctx, "keyFrameCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyFrameCount' field"))
	}
	m.KeyFrameCount = keyFrameCount

	dataSetName, err := ReadSimpleField[PascalString](ctx, "dataSetName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetName' field"))
	}
	m.DataSetName = dataSetName

	noOfDataSetWriterProperties, err := ReadSimpleField(ctx, "noOfDataSetWriterProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataSetWriterProperties' field"))
	}
	m.NoOfDataSetWriterProperties = noOfDataSetWriterProperties

	dataSetWriterProperties, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "dataSetWriterProperties", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("14535")), readBuffer), uint64(noOfDataSetWriterProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetWriterProperties' field"))
	}
	m.DataSetWriterProperties = dataSetWriterProperties

	transportSettings, err := ReadSimpleField[ExtensionObject](ctx, "transportSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSettings' field"))
	}
	m.TransportSettings = transportSettings

	messageSettings, err := ReadSimpleField[ExtensionObject](ctx, "messageSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageSettings' field"))
	}
	m.MessageSettings = messageSettings

	if closeErr := readBuffer.CloseContext("DataSetWriterDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetWriterDataType")
	}

	return m, nil
}

func (m *_DataSetWriterDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetWriterDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetWriterDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetWriterDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}

		if err := WriteSimpleField[uint16](ctx, "dataSetWriterId", m.GetDataSetWriterId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetWriterId' field")
		}

		if err := WriteSimpleEnumField[DataSetFieldContentMask](ctx, "dataSetFieldContentMask", "DataSetFieldContentMask", m.GetDataSetFieldContentMask(), WriteEnum[DataSetFieldContentMask, uint32](DataSetFieldContentMask.GetValue, DataSetFieldContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetFieldContentMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "keyFrameCount", m.GetKeyFrameCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'keyFrameCount' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "dataSetName", m.GetDataSetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetName' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDataSetWriterProperties", m.GetNoOfDataSetWriterProperties(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataSetWriterProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataSetWriterProperties", m.GetDataSetWriterProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetWriterProperties' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "transportSettings", m.GetTransportSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSettings' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "messageSettings", m.GetMessageSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'messageSettings' field")
		}

		if popErr := writeBuffer.PopContext("DataSetWriterDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetWriterDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetWriterDataType) IsDataSetWriterDataType() {}

func (m *_DataSetWriterDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSetWriterDataType) deepCopy() *_DataSetWriterDataType {
	if m == nil {
		return nil
	}
	_DataSetWriterDataTypeCopy := &_DataSetWriterDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Name.DeepCopy().(PascalString),
		m.Enabled,
		m.DataSetWriterId,
		m.DataSetFieldContentMask,
		m.KeyFrameCount,
		m.DataSetName.DeepCopy().(PascalString),
		m.NoOfDataSetWriterProperties,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.DataSetWriterProperties),
		m.TransportSettings.DeepCopy().(ExtensionObject),
		m.MessageSettings.DeepCopy().(ExtensionObject),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataSetWriterDataTypeCopy
}

func (m *_DataSetWriterDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
