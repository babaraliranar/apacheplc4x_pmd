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

// BACnetConstructedDataProtocolRevision is the corresponding interface of BACnetConstructedDataProtocolRevision
type BACnetConstructedDataProtocolRevision interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetProtocolRevision returns ProtocolRevision (property field)
	GetProtocolRevision() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataProtocolRevision is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProtocolRevision()
	// CreateBuilder creates a BACnetConstructedDataProtocolRevisionBuilder
	CreateBACnetConstructedDataProtocolRevisionBuilder() BACnetConstructedDataProtocolRevisionBuilder
}

// _BACnetConstructedDataProtocolRevision is the data-structure of this message
type _BACnetConstructedDataProtocolRevision struct {
	BACnetConstructedDataContract
	ProtocolRevision BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataProtocolRevision = (*_BACnetConstructedDataProtocolRevision)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProtocolRevision)(nil)

// NewBACnetConstructedDataProtocolRevision factory function for _BACnetConstructedDataProtocolRevision
func NewBACnetConstructedDataProtocolRevision(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, protocolRevision BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolRevision {
	if protocolRevision == nil {
		panic("protocolRevision of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataProtocolRevision must not be nil")
	}
	_result := &_BACnetConstructedDataProtocolRevision{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProtocolRevision:              protocolRevision,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProtocolRevisionBuilder is a builder for BACnetConstructedDataProtocolRevision
type BACnetConstructedDataProtocolRevisionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(protocolRevision BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolRevisionBuilder
	// WithProtocolRevision adds ProtocolRevision (property field)
	WithProtocolRevision(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolRevisionBuilder
	// WithProtocolRevisionBuilder adds ProtocolRevision (property field) which is build by the builder
	WithProtocolRevisionBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataProtocolRevisionBuilder
	// Build builds the BACnetConstructedDataProtocolRevision or returns an error if something is wrong
	Build() (BACnetConstructedDataProtocolRevision, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProtocolRevision
}

// NewBACnetConstructedDataProtocolRevisionBuilder() creates a BACnetConstructedDataProtocolRevisionBuilder
func NewBACnetConstructedDataProtocolRevisionBuilder() BACnetConstructedDataProtocolRevisionBuilder {
	return &_BACnetConstructedDataProtocolRevisionBuilder{_BACnetConstructedDataProtocolRevision: new(_BACnetConstructedDataProtocolRevision)}
}

type _BACnetConstructedDataProtocolRevisionBuilder struct {
	*_BACnetConstructedDataProtocolRevision

	err *utils.MultiError
}

var _ (BACnetConstructedDataProtocolRevisionBuilder) = (*_BACnetConstructedDataProtocolRevisionBuilder)(nil)

func (m *_BACnetConstructedDataProtocolRevisionBuilder) WithMandatoryFields(protocolRevision BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolRevisionBuilder {
	return m.WithProtocolRevision(protocolRevision)
}

func (m *_BACnetConstructedDataProtocolRevisionBuilder) WithProtocolRevision(protocolRevision BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolRevisionBuilder {
	m.ProtocolRevision = protocolRevision
	return m
}

func (m *_BACnetConstructedDataProtocolRevisionBuilder) WithProtocolRevisionBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataProtocolRevisionBuilder {
	builder := builderSupplier(m.ProtocolRevision.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.ProtocolRevision, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataProtocolRevisionBuilder) Build() (BACnetConstructedDataProtocolRevision, error) {
	if m.ProtocolRevision == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'protocolRevision' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataProtocolRevision.deepCopy(), nil
}

func (m *_BACnetConstructedDataProtocolRevisionBuilder) MustBuild() BACnetConstructedDataProtocolRevision {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataProtocolRevisionBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataProtocolRevisionBuilder()
}

// CreateBACnetConstructedDataProtocolRevisionBuilder creates a BACnetConstructedDataProtocolRevisionBuilder
func (m *_BACnetConstructedDataProtocolRevision) CreateBACnetConstructedDataProtocolRevisionBuilder() BACnetConstructedDataProtocolRevisionBuilder {
	if m == nil {
		return NewBACnetConstructedDataProtocolRevisionBuilder()
	}
	return &_BACnetConstructedDataProtocolRevisionBuilder{_BACnetConstructedDataProtocolRevision: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolRevision) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolRevision) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_REVISION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolRevision) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolRevision) GetProtocolRevision() BACnetApplicationTagUnsignedInteger {
	return m.ProtocolRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolRevision) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetProtocolRevision())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolRevision(structType any) BACnetConstructedDataProtocolRevision {
	if casted, ok := structType.(BACnetConstructedDataProtocolRevision); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolRevision); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolRevision) GetTypeName() string {
	return "BACnetConstructedDataProtocolRevision"
}

func (m *_BACnetConstructedDataProtocolRevision) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (protocolRevision)
	lengthInBits += m.ProtocolRevision.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolRevision) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProtocolRevision) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProtocolRevision BACnetConstructedDataProtocolRevision, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolRevision")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolRevision, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "protocolRevision", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolRevision' field"))
	}
	m.ProtocolRevision = protocolRevision

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), protocolRevision)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolRevision")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProtocolRevision) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolRevision) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolRevision")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "protocolRevision", m.GetProtocolRevision(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolRevision' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolRevision")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolRevision) IsBACnetConstructedDataProtocolRevision() {}

func (m *_BACnetConstructedDataProtocolRevision) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProtocolRevision) deepCopy() *_BACnetConstructedDataProtocolRevision {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProtocolRevisionCopy := &_BACnetConstructedDataProtocolRevision{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ProtocolRevision.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProtocolRevisionCopy
}

func (m *_BACnetConstructedDataProtocolRevision) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
