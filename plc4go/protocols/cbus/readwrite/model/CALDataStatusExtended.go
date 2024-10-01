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

// CALDataStatusExtended is the corresponding interface of CALDataStatusExtended
type CALDataStatusExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CALData
	// GetCoding returns Coding (property field)
	GetCoding() StatusCoding
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart (property field)
	GetBlockStart() uint8
	// GetStatusBytes returns StatusBytes (property field)
	GetStatusBytes() []StatusByte
	// GetLevelInformation returns LevelInformation (property field)
	GetLevelInformation() []LevelInformation
	// GetNumberOfStatusBytes returns NumberOfStatusBytes (virtual field)
	GetNumberOfStatusBytes() uint8
	// GetNumberOfLevelInformation returns NumberOfLevelInformation (virtual field)
	GetNumberOfLevelInformation() uint8
	// IsCALDataStatusExtended is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALDataStatusExtended()
	// CreateBuilder creates a CALDataStatusExtendedBuilder
	CreateCALDataStatusExtendedBuilder() CALDataStatusExtendedBuilder
}

// _CALDataStatusExtended is the data-structure of this message
type _CALDataStatusExtended struct {
	CALDataContract
	Coding           StatusCoding
	Application      ApplicationIdContainer
	BlockStart       uint8
	StatusBytes      []StatusByte
	LevelInformation []LevelInformation
}

var _ CALDataStatusExtended = (*_CALDataStatusExtended)(nil)
var _ CALDataRequirements = (*_CALDataStatusExtended)(nil)

// NewCALDataStatusExtended factory function for _CALDataStatusExtended
func NewCALDataStatusExtended(commandTypeContainer CALCommandTypeContainer, additionalData CALData, coding StatusCoding, application ApplicationIdContainer, blockStart uint8, statusBytes []StatusByte, levelInformation []LevelInformation, requestContext RequestContext) *_CALDataStatusExtended {
	_result := &_CALDataStatusExtended{
		CALDataContract:  NewCALData(commandTypeContainer, additionalData, requestContext),
		Coding:           coding,
		Application:      application,
		BlockStart:       blockStart,
		StatusBytes:      statusBytes,
		LevelInformation: levelInformation,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CALDataStatusExtendedBuilder is a builder for CALDataStatusExtended
type CALDataStatusExtendedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(coding StatusCoding, application ApplicationIdContainer, blockStart uint8, statusBytes []StatusByte, levelInformation []LevelInformation) CALDataStatusExtendedBuilder
	// WithCoding adds Coding (property field)
	WithCoding(StatusCoding) CALDataStatusExtendedBuilder
	// WithApplication adds Application (property field)
	WithApplication(ApplicationIdContainer) CALDataStatusExtendedBuilder
	// WithBlockStart adds BlockStart (property field)
	WithBlockStart(uint8) CALDataStatusExtendedBuilder
	// WithStatusBytes adds StatusBytes (property field)
	WithStatusBytes(...StatusByte) CALDataStatusExtendedBuilder
	// WithLevelInformation adds LevelInformation (property field)
	WithLevelInformation(...LevelInformation) CALDataStatusExtendedBuilder
	// Build builds the CALDataStatusExtended or returns an error if something is wrong
	Build() (CALDataStatusExtended, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CALDataStatusExtended
}

// NewCALDataStatusExtendedBuilder() creates a CALDataStatusExtendedBuilder
func NewCALDataStatusExtendedBuilder() CALDataStatusExtendedBuilder {
	return &_CALDataStatusExtendedBuilder{_CALDataStatusExtended: new(_CALDataStatusExtended)}
}

type _CALDataStatusExtendedBuilder struct {
	*_CALDataStatusExtended

	parentBuilder *_CALDataBuilder

	err *utils.MultiError
}

var _ (CALDataStatusExtendedBuilder) = (*_CALDataStatusExtendedBuilder)(nil)

func (b *_CALDataStatusExtendedBuilder) setParent(contract CALDataContract) {
	b.CALDataContract = contract
}

func (b *_CALDataStatusExtendedBuilder) WithMandatoryFields(coding StatusCoding, application ApplicationIdContainer, blockStart uint8, statusBytes []StatusByte, levelInformation []LevelInformation) CALDataStatusExtendedBuilder {
	return b.WithCoding(coding).WithApplication(application).WithBlockStart(blockStart).WithStatusBytes(statusBytes...).WithLevelInformation(levelInformation...)
}

func (b *_CALDataStatusExtendedBuilder) WithCoding(coding StatusCoding) CALDataStatusExtendedBuilder {
	b.Coding = coding
	return b
}

func (b *_CALDataStatusExtendedBuilder) WithApplication(application ApplicationIdContainer) CALDataStatusExtendedBuilder {
	b.Application = application
	return b
}

func (b *_CALDataStatusExtendedBuilder) WithBlockStart(blockStart uint8) CALDataStatusExtendedBuilder {
	b.BlockStart = blockStart
	return b
}

func (b *_CALDataStatusExtendedBuilder) WithStatusBytes(statusBytes ...StatusByte) CALDataStatusExtendedBuilder {
	b.StatusBytes = statusBytes
	return b
}

func (b *_CALDataStatusExtendedBuilder) WithLevelInformation(levelInformation ...LevelInformation) CALDataStatusExtendedBuilder {
	b.LevelInformation = levelInformation
	return b
}

func (b *_CALDataStatusExtendedBuilder) Build() (CALDataStatusExtended, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CALDataStatusExtended.deepCopy(), nil
}

func (b *_CALDataStatusExtendedBuilder) MustBuild() CALDataStatusExtended {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CALDataStatusExtendedBuilder) Done() CALDataBuilder {
	return b.parentBuilder
}

func (b *_CALDataStatusExtendedBuilder) buildForCALData() (CALData, error) {
	return b.Build()
}

func (b *_CALDataStatusExtendedBuilder) DeepCopy() any {
	_copy := b.CreateCALDataStatusExtendedBuilder().(*_CALDataStatusExtendedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCALDataStatusExtendedBuilder creates a CALDataStatusExtendedBuilder
func (b *_CALDataStatusExtended) CreateCALDataStatusExtendedBuilder() CALDataStatusExtendedBuilder {
	if b == nil {
		return NewCALDataStatusExtendedBuilder()
	}
	return &_CALDataStatusExtendedBuilder{_CALDataStatusExtended: b.deepCopy()}
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

func (m *_CALDataStatusExtended) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataStatusExtended) GetCoding() StatusCoding {
	return m.Coding
}

func (m *_CALDataStatusExtended) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CALDataStatusExtended) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *_CALDataStatusExtended) GetStatusBytes() []StatusByte {
	return m.StatusBytes
}

func (m *_CALDataStatusExtended) GetLevelInformation() []LevelInformation {
	return m.LevelInformation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CALDataStatusExtended) GetNumberOfStatusBytes() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(utils.InlineIf((bool(bool((m.GetCoding()) == (StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE))) || bool(bool((m.GetCoding()) == (StatusCoding_BINARY_BY_ELSEWHERE)))), func() any { return uint8((uint8(m.GetCommandTypeContainer().NumBytes()) - uint8(uint8(3)))) }, func() any { return uint8((uint8(0))) }).(uint8))
}

func (m *_CALDataStatusExtended) GetNumberOfLevelInformation() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(utils.InlineIf((bool(bool((m.GetCoding()) == (StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE))) || bool(bool((m.GetCoding()) == (StatusCoding_LEVEL_BY_ELSEWHERE)))), func() any {
		return uint8((uint8((uint8(m.GetCommandTypeContainer().NumBytes()) - uint8(uint8(3)))) / uint8(uint8(2))))
	}, func() any { return uint8((uint8(0))) }).(uint8))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCALDataStatusExtended(structType any) CALDataStatusExtended {
	if casted, ok := structType.(CALDataStatusExtended); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataStatusExtended); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataStatusExtended) GetTypeName() string {
	return "CALDataStatusExtended"
}

func (m *_CALDataStatusExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (coding)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.StatusBytes) > 0 {
		for _curItem, element := range m.StatusBytes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StatusBytes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.LevelInformation) > 0 {
		for _curItem, element := range m.LevelInformation {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LevelInformation), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CALDataStatusExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataStatusExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (__cALDataStatusExtended CALDataStatusExtended, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataStatusExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataStatusExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	coding, err := ReadEnumField[StatusCoding](ctx, "coding", "StatusCoding", ReadEnum(StatusCodingByValue, ReadByte(readBuffer, 8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'coding' field"))
	}
	m.Coding = coding

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}
	m.Application = application

	blockStart, err := ReadSimpleField(ctx, "blockStart", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'blockStart' field"))
	}
	m.BlockStart = blockStart

	numberOfStatusBytes, err := ReadVirtualField[uint8](ctx, "numberOfStatusBytes", (*uint8)(nil), utils.InlineIf((bool(bool((coding) == (StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE))) || bool(bool((coding) == (StatusCoding_BINARY_BY_ELSEWHERE)))), func() any { return uint8((uint8(commandTypeContainer.NumBytes()) - uint8(uint8(3)))) }, func() any { return uint8((uint8(0))) }).(uint8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfStatusBytes' field"))
	}
	_ = numberOfStatusBytes

	numberOfLevelInformation, err := ReadVirtualField[uint8](ctx, "numberOfLevelInformation", (*uint8)(nil), utils.InlineIf((bool(bool((coding) == (StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE))) || bool(bool((coding) == (StatusCoding_LEVEL_BY_ELSEWHERE)))), func() any {
		return uint8((uint8((uint8(commandTypeContainer.NumBytes()) - uint8(uint8(3)))) / uint8(uint8(2))))
	}, func() any { return uint8((uint8(0))) }).(uint8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfLevelInformation' field"))
	}
	_ = numberOfLevelInformation

	statusBytes, err := ReadCountArrayField[StatusByte](ctx, "statusBytes", ReadComplex[StatusByte](StatusByteParseWithBuffer, readBuffer), uint64(numberOfStatusBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusBytes' field"))
	}
	m.StatusBytes = statusBytes

	levelInformation, err := ReadCountArrayField[LevelInformation](ctx, "levelInformation", ReadComplex[LevelInformation](LevelInformationParseWithBuffer, readBuffer), uint64(numberOfLevelInformation))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'levelInformation' field"))
	}
	m.LevelInformation = levelInformation

	if closeErr := readBuffer.CloseContext("CALDataStatusExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataStatusExtended")
	}

	return m, nil
}

func (m *_CALDataStatusExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataStatusExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataStatusExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataStatusExtended")
		}

		if err := WriteSimpleEnumField[StatusCoding](ctx, "coding", "StatusCoding", m.GetCoding(), WriteEnum[StatusCoding, byte](StatusCoding.GetValue, StatusCoding.PLC4XEnumName, WriteByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'coding' field")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", m.GetApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'application' field")
		}

		if err := WriteSimpleField[uint8](ctx, "blockStart", m.GetBlockStart(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'blockStart' field")
		}
		// Virtual field
		numberOfStatusBytes := m.GetNumberOfStatusBytes()
		_ = numberOfStatusBytes
		if _numberOfStatusBytesErr := writeBuffer.WriteVirtual(ctx, "numberOfStatusBytes", m.GetNumberOfStatusBytes()); _numberOfStatusBytesErr != nil {
			return errors.Wrap(_numberOfStatusBytesErr, "Error serializing 'numberOfStatusBytes' field")
		}
		// Virtual field
		numberOfLevelInformation := m.GetNumberOfLevelInformation()
		_ = numberOfLevelInformation
		if _numberOfLevelInformationErr := writeBuffer.WriteVirtual(ctx, "numberOfLevelInformation", m.GetNumberOfLevelInformation()); _numberOfLevelInformationErr != nil {
			return errors.Wrap(_numberOfLevelInformationErr, "Error serializing 'numberOfLevelInformation' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "statusBytes", m.GetStatusBytes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'statusBytes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "levelInformation", m.GetLevelInformation(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'levelInformation' field")
		}

		if popErr := writeBuffer.PopContext("CALDataStatusExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataStatusExtended")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataStatusExtended) IsCALDataStatusExtended() {}

func (m *_CALDataStatusExtended) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CALDataStatusExtended) deepCopy() *_CALDataStatusExtended {
	if m == nil {
		return nil
	}
	_CALDataStatusExtendedCopy := &_CALDataStatusExtended{
		m.CALDataContract.(*_CALData).deepCopy(),
		m.Coding,
		m.Application,
		m.BlockStart,
		utils.DeepCopySlice[StatusByte, StatusByte](m.StatusBytes),
		utils.DeepCopySlice[LevelInformation, LevelInformation](m.LevelInformation),
	}
	m.CALDataContract.(*_CALData)._SubType = m
	return _CALDataStatusExtendedCopy
}

func (m *_CALDataStatusExtended) String() string {
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
