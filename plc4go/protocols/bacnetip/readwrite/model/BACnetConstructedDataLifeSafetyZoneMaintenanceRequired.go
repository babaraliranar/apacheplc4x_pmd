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

// BACnetConstructedDataLifeSafetyZoneMaintenanceRequired is the corresponding interface of BACnetConstructedDataLifeSafetyZoneMaintenanceRequired
type BACnetConstructedDataLifeSafetyZoneMaintenanceRequired interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaintenanceRequired returns MaintenanceRequired (property field)
	GetMaintenanceRequired() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataLifeSafetyZoneMaintenanceRequired is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyZoneMaintenanceRequired()
	// CreateBuilder creates a BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
	CreateBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder() BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
}

// _BACnetConstructedDataLifeSafetyZoneMaintenanceRequired is the data-structure of this message
type _BACnetConstructedDataLifeSafetyZoneMaintenanceRequired struct {
	BACnetConstructedDataContract
	MaintenanceRequired BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataLifeSafetyZoneMaintenanceRequired = (*_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired)(nil)

// NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequired factory function for _BACnetConstructedDataLifeSafetyZoneMaintenanceRequired
func NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequired(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maintenanceRequired BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired {
	if maintenanceRequired == nil {
		panic("maintenanceRequired of type BACnetApplicationTagBoolean for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired must not be nil")
	}
	_result := &_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaintenanceRequired:           maintenanceRequired,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder is a builder for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired
type BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maintenanceRequired BACnetApplicationTagBoolean) BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
	// WithMaintenanceRequired adds MaintenanceRequired (property field)
	WithMaintenanceRequired(BACnetApplicationTagBoolean) BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
	// WithMaintenanceRequiredBuilder adds MaintenanceRequired (property field) which is build by the builder
	WithMaintenanceRequiredBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
	// Build builds the BACnetConstructedDataLifeSafetyZoneMaintenanceRequired or returns an error if something is wrong
	Build() (BACnetConstructedDataLifeSafetyZoneMaintenanceRequired, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLifeSafetyZoneMaintenanceRequired
}

// NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder() creates a BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
func NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder() BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder {
	return &_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder{_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired: new(_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired)}
}

type _BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder struct {
	*_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired

	err *utils.MultiError
}

var _ (BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) = (*_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder)(nil)

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) WithMandatoryFields(maintenanceRequired BACnetApplicationTagBoolean) BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder {
	return m.WithMaintenanceRequired(maintenanceRequired)
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) WithMaintenanceRequired(maintenanceRequired BACnetApplicationTagBoolean) BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder {
	m.MaintenanceRequired = maintenanceRequired
	return m
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) WithMaintenanceRequiredBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder {
	builder := builderSupplier(m.MaintenanceRequired.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.MaintenanceRequired, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) Build() (BACnetConstructedDataLifeSafetyZoneMaintenanceRequired, error) {
	if m.MaintenanceRequired == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'maintenanceRequired' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLifeSafetyZoneMaintenanceRequired.deepCopy(), nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) MustBuild() BACnetConstructedDataLifeSafetyZoneMaintenanceRequired {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder()
}

// CreateBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder creates a BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder
func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) CreateBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder() BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder {
	if m == nil {
		return NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder()
	}
	return &_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredBuilder{_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAINTENANCE_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetMaintenanceRequired() BACnetApplicationTagBoolean {
	return m.MaintenanceRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetMaintenanceRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyZoneMaintenanceRequired(structType any) BACnetConstructedDataLifeSafetyZoneMaintenanceRequired {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneMaintenanceRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneMaintenanceRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maintenanceRequired)
	lengthInBits += m.MaintenanceRequired.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyZoneMaintenanceRequired BACnetConstructedDataLifeSafetyZoneMaintenanceRequired, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maintenanceRequired, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "maintenanceRequired", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maintenanceRequired' field"))
	}
	m.MaintenanceRequired = maintenanceRequired

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), maintenanceRequired)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "maintenanceRequired", m.GetMaintenanceRequired(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maintenanceRequired' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) IsBACnetConstructedDataLifeSafetyZoneMaintenanceRequired() {
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) deepCopy() *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredCopy := &_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MaintenanceRequired.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredCopy
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
