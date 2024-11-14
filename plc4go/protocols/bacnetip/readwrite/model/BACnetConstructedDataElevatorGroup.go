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

// BACnetConstructedDataElevatorGroup is the corresponding interface of BACnetConstructedDataElevatorGroup
type BACnetConstructedDataElevatorGroup interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetElevatorGroup returns ElevatorGroup (property field)
	GetElevatorGroup() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetConstructedDataElevatorGroup is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataElevatorGroup()
	// CreateBuilder creates a BACnetConstructedDataElevatorGroupBuilder
	CreateBACnetConstructedDataElevatorGroupBuilder() BACnetConstructedDataElevatorGroupBuilder
}

// _BACnetConstructedDataElevatorGroup is the data-structure of this message
type _BACnetConstructedDataElevatorGroup struct {
	BACnetConstructedDataContract
	ElevatorGroup BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataElevatorGroup = (*_BACnetConstructedDataElevatorGroup)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataElevatorGroup)(nil)

// NewBACnetConstructedDataElevatorGroup factory function for _BACnetConstructedDataElevatorGroup
func NewBACnetConstructedDataElevatorGroup(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, elevatorGroup BACnetApplicationTagObjectIdentifier, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataElevatorGroup {
	if elevatorGroup == nil {
		panic("elevatorGroup of type BACnetApplicationTagObjectIdentifier for BACnetConstructedDataElevatorGroup must not be nil")
	}
	_result := &_BACnetConstructedDataElevatorGroup{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ElevatorGroup:                 elevatorGroup,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataElevatorGroupBuilder is a builder for BACnetConstructedDataElevatorGroup
type BACnetConstructedDataElevatorGroupBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(elevatorGroup BACnetApplicationTagObjectIdentifier) BACnetConstructedDataElevatorGroupBuilder
	// WithElevatorGroup adds ElevatorGroup (property field)
	WithElevatorGroup(BACnetApplicationTagObjectIdentifier) BACnetConstructedDataElevatorGroupBuilder
	// WithElevatorGroupBuilder adds ElevatorGroup (property field) which is build by the builder
	WithElevatorGroupBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConstructedDataElevatorGroupBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataElevatorGroup or returns an error if something is wrong
	Build() (BACnetConstructedDataElevatorGroup, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataElevatorGroup
}

// NewBACnetConstructedDataElevatorGroupBuilder() creates a BACnetConstructedDataElevatorGroupBuilder
func NewBACnetConstructedDataElevatorGroupBuilder() BACnetConstructedDataElevatorGroupBuilder {
	return &_BACnetConstructedDataElevatorGroupBuilder{_BACnetConstructedDataElevatorGroup: new(_BACnetConstructedDataElevatorGroup)}
}

type _BACnetConstructedDataElevatorGroupBuilder struct {
	*_BACnetConstructedDataElevatorGroup

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataElevatorGroupBuilder) = (*_BACnetConstructedDataElevatorGroupBuilder)(nil)

func (b *_BACnetConstructedDataElevatorGroupBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataElevatorGroup
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) WithMandatoryFields(elevatorGroup BACnetApplicationTagObjectIdentifier) BACnetConstructedDataElevatorGroupBuilder {
	return b.WithElevatorGroup(elevatorGroup)
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) WithElevatorGroup(elevatorGroup BACnetApplicationTagObjectIdentifier) BACnetConstructedDataElevatorGroupBuilder {
	b.ElevatorGroup = elevatorGroup
	return b
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) WithElevatorGroupBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConstructedDataElevatorGroupBuilder {
	builder := builderSupplier(b.ElevatorGroup.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.ElevatorGroup, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) Build() (BACnetConstructedDataElevatorGroup, error) {
	if b.ElevatorGroup == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'elevatorGroup' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataElevatorGroup.deepCopy(), nil
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) MustBuild() BACnetConstructedDataElevatorGroup {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataElevatorGroupBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataElevatorGroupBuilder().(*_BACnetConstructedDataElevatorGroupBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataElevatorGroupBuilder creates a BACnetConstructedDataElevatorGroupBuilder
func (b *_BACnetConstructedDataElevatorGroup) CreateBACnetConstructedDataElevatorGroupBuilder() BACnetConstructedDataElevatorGroupBuilder {
	if b == nil {
		return NewBACnetConstructedDataElevatorGroupBuilder()
	}
	return &_BACnetConstructedDataElevatorGroupBuilder{_BACnetConstructedDataElevatorGroup: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataElevatorGroup) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataElevatorGroup) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ELEVATOR_GROUP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataElevatorGroup) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataElevatorGroup) GetElevatorGroup() BACnetApplicationTagObjectIdentifier {
	return m.ElevatorGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataElevatorGroup) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetElevatorGroup())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataElevatorGroup(structType any) BACnetConstructedDataElevatorGroup {
	if casted, ok := structType.(BACnetConstructedDataElevatorGroup); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataElevatorGroup); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataElevatorGroup) GetTypeName() string {
	return "BACnetConstructedDataElevatorGroup"
}

func (m *_BACnetConstructedDataElevatorGroup) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (elevatorGroup)
	lengthInBits += m.ElevatorGroup.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataElevatorGroup) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataElevatorGroup) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataElevatorGroup BACnetConstructedDataElevatorGroup, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElevatorGroup"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataElevatorGroup")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	elevatorGroup, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "elevatorGroup", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elevatorGroup' field"))
	}
	m.ElevatorGroup = elevatorGroup

	actualValue, err := ReadVirtualField[BACnetApplicationTagObjectIdentifier](ctx, "actualValue", (*BACnetApplicationTagObjectIdentifier)(nil), elevatorGroup)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElevatorGroup"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataElevatorGroup")
	}

	return m, nil
}

func (m *_BACnetConstructedDataElevatorGroup) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataElevatorGroup) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataElevatorGroup"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataElevatorGroup")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "elevatorGroup", m.GetElevatorGroup(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'elevatorGroup' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataElevatorGroup"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataElevatorGroup")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataElevatorGroup) IsBACnetConstructedDataElevatorGroup() {}

func (m *_BACnetConstructedDataElevatorGroup) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataElevatorGroup) deepCopy() *_BACnetConstructedDataElevatorGroup {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataElevatorGroupCopy := &_BACnetConstructedDataElevatorGroup{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagObjectIdentifier](m.ElevatorGroup),
	}
	_BACnetConstructedDataElevatorGroupCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataElevatorGroupCopy
}

func (m *_BACnetConstructedDataElevatorGroup) String() string {
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
