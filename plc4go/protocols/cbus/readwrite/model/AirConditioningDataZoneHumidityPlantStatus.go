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

// AirConditioningDataZoneHumidityPlantStatus is the corresponding interface of AirConditioningDataZoneHumidityPlantStatus
type AirConditioningDataZoneHumidityPlantStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidityType returns HumidityType (property field)
	GetHumidityType() HVACHumidityType
	// GetHumidityStatus returns HumidityStatus (property field)
	GetHumidityStatus() HVACHumidityStatusFlags
	// GetHumidityErrorCode returns HumidityErrorCode (property field)
	GetHumidityErrorCode() HVACHumidityError
	// IsAirConditioningDataZoneHumidityPlantStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataZoneHumidityPlantStatus()
	// CreateBuilder creates a AirConditioningDataZoneHumidityPlantStatusBuilder
	CreateAirConditioningDataZoneHumidityPlantStatusBuilder() AirConditioningDataZoneHumidityPlantStatusBuilder
}

// _AirConditioningDataZoneHumidityPlantStatus is the data-structure of this message
type _AirConditioningDataZoneHumidityPlantStatus struct {
	AirConditioningDataContract
	ZoneGroup         byte
	ZoneList          HVACZoneList
	HumidityType      HVACHumidityType
	HumidityStatus    HVACHumidityStatusFlags
	HumidityErrorCode HVACHumidityError
}

var _ AirConditioningDataZoneHumidityPlantStatus = (*_AirConditioningDataZoneHumidityPlantStatus)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataZoneHumidityPlantStatus)(nil)

// NewAirConditioningDataZoneHumidityPlantStatus factory function for _AirConditioningDataZoneHumidityPlantStatus
func NewAirConditioningDataZoneHumidityPlantStatus(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, humidityType HVACHumidityType, humidityStatus HVACHumidityStatusFlags, humidityErrorCode HVACHumidityError) *_AirConditioningDataZoneHumidityPlantStatus {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataZoneHumidityPlantStatus must not be nil")
	}
	if humidityStatus == nil {
		panic("humidityStatus of type HVACHumidityStatusFlags for AirConditioningDataZoneHumidityPlantStatus must not be nil")
	}
	_result := &_AirConditioningDataZoneHumidityPlantStatus{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HumidityType:                humidityType,
		HumidityStatus:              humidityStatus,
		HumidityErrorCode:           humidityErrorCode,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataZoneHumidityPlantStatusBuilder is a builder for AirConditioningDataZoneHumidityPlantStatus
type AirConditioningDataZoneHumidityPlantStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidityType HVACHumidityType, humidityStatus HVACHumidityStatusFlags, humidityErrorCode HVACHumidityError) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithHumidityType adds HumidityType (property field)
	WithHumidityType(HVACHumidityType) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithHumidityStatus adds HumidityStatus (property field)
	WithHumidityStatus(HVACHumidityStatusFlags) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithHumidityStatusBuilder adds HumidityStatus (property field) which is build by the builder
	WithHumidityStatusBuilder(func(HVACHumidityStatusFlagsBuilder) HVACHumidityStatusFlagsBuilder) AirConditioningDataZoneHumidityPlantStatusBuilder
	// WithHumidityErrorCode adds HumidityErrorCode (property field)
	WithHumidityErrorCode(HVACHumidityError) AirConditioningDataZoneHumidityPlantStatusBuilder
	// Build builds the AirConditioningDataZoneHumidityPlantStatus or returns an error if something is wrong
	Build() (AirConditioningDataZoneHumidityPlantStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataZoneHumidityPlantStatus
}

// NewAirConditioningDataZoneHumidityPlantStatusBuilder() creates a AirConditioningDataZoneHumidityPlantStatusBuilder
func NewAirConditioningDataZoneHumidityPlantStatusBuilder() AirConditioningDataZoneHumidityPlantStatusBuilder {
	return &_AirConditioningDataZoneHumidityPlantStatusBuilder{_AirConditioningDataZoneHumidityPlantStatus: new(_AirConditioningDataZoneHumidityPlantStatus)}
}

type _AirConditioningDataZoneHumidityPlantStatusBuilder struct {
	*_AirConditioningDataZoneHumidityPlantStatus

	err *utils.MultiError
}

var _ (AirConditioningDataZoneHumidityPlantStatusBuilder) = (*_AirConditioningDataZoneHumidityPlantStatusBuilder)(nil)

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidityType HVACHumidityType, humidityStatus HVACHumidityStatusFlags, humidityErrorCode HVACHumidityError) AirConditioningDataZoneHumidityPlantStatusBuilder {
	return m.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithHumidityType(humidityType).WithHumidityStatus(humidityStatus).WithHumidityErrorCode(humidityErrorCode)
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataZoneHumidityPlantStatusBuilder {
	m.ZoneGroup = zoneGroup
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataZoneHumidityPlantStatusBuilder {
	m.ZoneList = zoneList
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataZoneHumidityPlantStatusBuilder {
	builder := builderSupplier(m.ZoneList.CreateHVACZoneListBuilder())
	var err error
	m.ZoneList, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "HVACZoneListBuilder failed"))
	}
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithHumidityType(humidityType HVACHumidityType) AirConditioningDataZoneHumidityPlantStatusBuilder {
	m.HumidityType = humidityType
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithHumidityStatus(humidityStatus HVACHumidityStatusFlags) AirConditioningDataZoneHumidityPlantStatusBuilder {
	m.HumidityStatus = humidityStatus
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithHumidityStatusBuilder(builderSupplier func(HVACHumidityStatusFlagsBuilder) HVACHumidityStatusFlagsBuilder) AirConditioningDataZoneHumidityPlantStatusBuilder {
	builder := builderSupplier(m.HumidityStatus.CreateHVACHumidityStatusFlagsBuilder())
	var err error
	m.HumidityStatus, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "HVACHumidityStatusFlagsBuilder failed"))
	}
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) WithHumidityErrorCode(humidityErrorCode HVACHumidityError) AirConditioningDataZoneHumidityPlantStatusBuilder {
	m.HumidityErrorCode = humidityErrorCode
	return m
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) Build() (AirConditioningDataZoneHumidityPlantStatus, error) {
	if m.ZoneList == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if m.HumidityStatus == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'humidityStatus' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AirConditioningDataZoneHumidityPlantStatus.deepCopy(), nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) MustBuild() AirConditioningDataZoneHumidityPlantStatus {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AirConditioningDataZoneHumidityPlantStatusBuilder) DeepCopy() any {
	return m.CreateAirConditioningDataZoneHumidityPlantStatusBuilder()
}

// CreateAirConditioningDataZoneHumidityPlantStatusBuilder creates a AirConditioningDataZoneHumidityPlantStatusBuilder
func (m *_AirConditioningDataZoneHumidityPlantStatus) CreateAirConditioningDataZoneHumidityPlantStatusBuilder() AirConditioningDataZoneHumidityPlantStatusBuilder {
	if m == nil {
		return NewAirConditioningDataZoneHumidityPlantStatusBuilder()
	}
	return &_AirConditioningDataZoneHumidityPlantStatusBuilder{_AirConditioningDataZoneHumidityPlantStatus: m.deepCopy()}
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

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetHumidityType() HVACHumidityType {
	return m.HumidityType
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetHumidityStatus() HVACHumidityStatusFlags {
	return m.HumidityStatus
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetHumidityErrorCode() HVACHumidityError {
	return m.HumidityErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHumidityPlantStatus(structType any) AirConditioningDataZoneHumidityPlantStatus {
	if casted, ok := structType.(AirConditioningDataZoneHumidityPlantStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHumidityPlantStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetTypeName() string {
	return "AirConditioningDataZoneHumidityPlantStatus"
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidityType)
	lengthInBits += 8

	// Simple field (humidityStatus)
	lengthInBits += m.HumidityStatus.GetLengthInBits(ctx)

	// Simple field (humidityErrorCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataZoneHumidityPlantStatus AirConditioningDataZoneHumidityPlantStatus, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHumidityPlantStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHumidityPlantStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	humidityType, err := ReadEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", ReadEnum(HVACHumidityTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityType' field"))
	}
	m.HumidityType = humidityType

	humidityStatus, err := ReadSimpleField[HVACHumidityStatusFlags](ctx, "humidityStatus", ReadComplex[HVACHumidityStatusFlags](HVACHumidityStatusFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityStatus' field"))
	}
	m.HumidityStatus = humidityStatus

	humidityErrorCode, err := ReadEnumField[HVACHumidityError](ctx, "humidityErrorCode", "HVACHumidityError", ReadEnum(HVACHumidityErrorByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityErrorCode' field"))
	}
	m.HumidityErrorCode = humidityErrorCode

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHumidityPlantStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHumidityPlantStatus")
	}

	return m, nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHumidityPlantStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHumidityPlantStatus")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", m.GetHumidityType(), WriteEnum[HVACHumidityType, uint8](HVACHumidityType.GetValue, HVACHumidityType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityType' field")
		}

		if err := WriteSimpleField[HVACHumidityStatusFlags](ctx, "humidityStatus", m.GetHumidityStatus(), WriteComplex[HVACHumidityStatusFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityStatus' field")
		}

		if err := WriteSimpleEnumField[HVACHumidityError](ctx, "humidityErrorCode", "HVACHumidityError", m.GetHumidityErrorCode(), WriteEnum[HVACHumidityError, uint8](HVACHumidityError.GetValue, HVACHumidityError.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityErrorCode' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHumidityPlantStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHumidityPlantStatus")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) IsAirConditioningDataZoneHumidityPlantStatus() {
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) deepCopy() *_AirConditioningDataZoneHumidityPlantStatus {
	if m == nil {
		return nil
	}
	_AirConditioningDataZoneHumidityPlantStatusCopy := &_AirConditioningDataZoneHumidityPlantStatus{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		m.ZoneList.DeepCopy().(HVACZoneList),
		m.HumidityType,
		m.HumidityStatus.DeepCopy().(HVACHumidityStatusFlags),
		m.HumidityErrorCode,
	}
	m.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataZoneHumidityPlantStatusCopy
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
