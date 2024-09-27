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

// BACnetConstructedDataLastPriority is the corresponding interface of BACnetConstructedDataLastPriority
type BACnetConstructedDataLastPriority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastPriority returns LastPriority (property field)
	GetLastPriority() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataLastPriority is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastPriority()
	// CreateBuilder creates a BACnetConstructedDataLastPriorityBuilder
	CreateBACnetConstructedDataLastPriorityBuilder() BACnetConstructedDataLastPriorityBuilder
}

// _BACnetConstructedDataLastPriority is the data-structure of this message
type _BACnetConstructedDataLastPriority struct {
	BACnetConstructedDataContract
	LastPriority BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataLastPriority = (*_BACnetConstructedDataLastPriority)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastPriority)(nil)

// NewBACnetConstructedDataLastPriority factory function for _BACnetConstructedDataLastPriority
func NewBACnetConstructedDataLastPriority(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastPriority BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastPriority {
	if lastPriority == nil {
		panic("lastPriority of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataLastPriority must not be nil")
	}
	_result := &_BACnetConstructedDataLastPriority{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastPriority:                  lastPriority,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastPriorityBuilder is a builder for BACnetConstructedDataLastPriority
type BACnetConstructedDataLastPriorityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastPriority BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLastPriorityBuilder
	// WithLastPriority adds LastPriority (property field)
	WithLastPriority(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLastPriorityBuilder
	// WithLastPriorityBuilder adds LastPriority (property field) which is build by the builder
	WithLastPriorityBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLastPriorityBuilder
	// Build builds the BACnetConstructedDataLastPriority or returns an error if something is wrong
	Build() (BACnetConstructedDataLastPriority, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastPriority
}

// NewBACnetConstructedDataLastPriorityBuilder() creates a BACnetConstructedDataLastPriorityBuilder
func NewBACnetConstructedDataLastPriorityBuilder() BACnetConstructedDataLastPriorityBuilder {
	return &_BACnetConstructedDataLastPriorityBuilder{_BACnetConstructedDataLastPriority: new(_BACnetConstructedDataLastPriority)}
}

type _BACnetConstructedDataLastPriorityBuilder struct {
	*_BACnetConstructedDataLastPriority

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastPriorityBuilder) = (*_BACnetConstructedDataLastPriorityBuilder)(nil)

func (m *_BACnetConstructedDataLastPriorityBuilder) WithMandatoryFields(lastPriority BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLastPriorityBuilder {
	return m.WithLastPriority(lastPriority)
}

func (m *_BACnetConstructedDataLastPriorityBuilder) WithLastPriority(lastPriority BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLastPriorityBuilder {
	m.LastPriority = lastPriority
	return m
}

func (m *_BACnetConstructedDataLastPriorityBuilder) WithLastPriorityBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLastPriorityBuilder {
	builder := builderSupplier(m.LastPriority.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.LastPriority, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLastPriorityBuilder) Build() (BACnetConstructedDataLastPriority, error) {
	if m.LastPriority == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lastPriority' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLastPriority.deepCopy(), nil
}

func (m *_BACnetConstructedDataLastPriorityBuilder) MustBuild() BACnetConstructedDataLastPriority {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLastPriorityBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLastPriorityBuilder()
}

// CreateBACnetConstructedDataLastPriorityBuilder creates a BACnetConstructedDataLastPriorityBuilder
func (m *_BACnetConstructedDataLastPriority) CreateBACnetConstructedDataLastPriorityBuilder() BACnetConstructedDataLastPriorityBuilder {
	if m == nil {
		return NewBACnetConstructedDataLastPriorityBuilder()
	}
	return &_BACnetConstructedDataLastPriorityBuilder{_BACnetConstructedDataLastPriority: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastPriority) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastPriority) GetLastPriority() BACnetApplicationTagUnsignedInteger {
	return m.LastPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastPriority) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLastPriority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastPriority(structType any) BACnetConstructedDataLastPriority {
	if casted, ok := structType.(BACnetConstructedDataLastPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastPriority) GetTypeName() string {
	return "BACnetConstructedDataLastPriority"
}

func (m *_BACnetConstructedDataLastPriority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastPriority)
	lengthInBits += m.LastPriority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastPriority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastPriority) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastPriority BACnetConstructedDataLastPriority, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastPriority, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lastPriority", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastPriority' field"))
	}
	m.LastPriority = lastPriority

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), lastPriority)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastPriority")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastPriority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastPriority")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lastPriority", m.GetLastPriority(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastPriority' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastPriority")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastPriority) IsBACnetConstructedDataLastPriority() {}

func (m *_BACnetConstructedDataLastPriority) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastPriority) deepCopy() *_BACnetConstructedDataLastPriority {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastPriorityCopy := &_BACnetConstructedDataLastPriority{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LastPriority.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastPriorityCopy
}

func (m *_BACnetConstructedDataLastPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
