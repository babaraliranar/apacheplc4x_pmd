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

// BACnetConstructedDataProtocolVersion is the corresponding interface of BACnetConstructedDataProtocolVersion
type BACnetConstructedDataProtocolVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetProtocolVersion returns ProtocolVersion (property field)
	GetProtocolVersion() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataProtocolVersion is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProtocolVersion()
	// CreateBuilder creates a BACnetConstructedDataProtocolVersionBuilder
	CreateBACnetConstructedDataProtocolVersionBuilder() BACnetConstructedDataProtocolVersionBuilder
}

// _BACnetConstructedDataProtocolVersion is the data-structure of this message
type _BACnetConstructedDataProtocolVersion struct {
	BACnetConstructedDataContract
	ProtocolVersion BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataProtocolVersion = (*_BACnetConstructedDataProtocolVersion)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProtocolVersion)(nil)

// NewBACnetConstructedDataProtocolVersion factory function for _BACnetConstructedDataProtocolVersion
func NewBACnetConstructedDataProtocolVersion(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, protocolVersion BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolVersion {
	if protocolVersion == nil {
		panic("protocolVersion of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataProtocolVersion must not be nil")
	}
	_result := &_BACnetConstructedDataProtocolVersion{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProtocolVersion:               protocolVersion,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProtocolVersionBuilder is a builder for BACnetConstructedDataProtocolVersion
type BACnetConstructedDataProtocolVersionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(protocolVersion BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolVersionBuilder
	// WithProtocolVersion adds ProtocolVersion (property field)
	WithProtocolVersion(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolVersionBuilder
	// WithProtocolVersionBuilder adds ProtocolVersion (property field) which is build by the builder
	WithProtocolVersionBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataProtocolVersionBuilder
	// Build builds the BACnetConstructedDataProtocolVersion or returns an error if something is wrong
	Build() (BACnetConstructedDataProtocolVersion, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProtocolVersion
}

// NewBACnetConstructedDataProtocolVersionBuilder() creates a BACnetConstructedDataProtocolVersionBuilder
func NewBACnetConstructedDataProtocolVersionBuilder() BACnetConstructedDataProtocolVersionBuilder {
	return &_BACnetConstructedDataProtocolVersionBuilder{_BACnetConstructedDataProtocolVersion: new(_BACnetConstructedDataProtocolVersion)}
}

type _BACnetConstructedDataProtocolVersionBuilder struct {
	*_BACnetConstructedDataProtocolVersion

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataProtocolVersionBuilder) = (*_BACnetConstructedDataProtocolVersionBuilder)(nil)

func (b *_BACnetConstructedDataProtocolVersionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) WithMandatoryFields(protocolVersion BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolVersionBuilder {
	return b.WithProtocolVersion(protocolVersion)
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) WithProtocolVersion(protocolVersion BACnetApplicationTagUnsignedInteger) BACnetConstructedDataProtocolVersionBuilder {
	b.ProtocolVersion = protocolVersion
	return b
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) WithProtocolVersionBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataProtocolVersionBuilder {
	builder := builderSupplier(b.ProtocolVersion.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ProtocolVersion, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) Build() (BACnetConstructedDataProtocolVersion, error) {
	if b.ProtocolVersion == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'protocolVersion' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataProtocolVersion.deepCopy(), nil
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) MustBuild() BACnetConstructedDataProtocolVersion {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataProtocolVersionBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataProtocolVersionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataProtocolVersionBuilder().(*_BACnetConstructedDataProtocolVersionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataProtocolVersionBuilder creates a BACnetConstructedDataProtocolVersionBuilder
func (b *_BACnetConstructedDataProtocolVersion) CreateBACnetConstructedDataProtocolVersionBuilder() BACnetConstructedDataProtocolVersionBuilder {
	if b == nil {
		return NewBACnetConstructedDataProtocolVersionBuilder()
	}
	return &_BACnetConstructedDataProtocolVersionBuilder{_BACnetConstructedDataProtocolVersion: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolVersion) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetProtocolVersion() BACnetApplicationTagUnsignedInteger {
	return m.ProtocolVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetProtocolVersion())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolVersion(structType any) BACnetConstructedDataProtocolVersion {
	if casted, ok := structType.(BACnetConstructedDataProtocolVersion); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolVersion); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolVersion) GetTypeName() string {
	return "BACnetConstructedDataProtocolVersion"
}

func (m *_BACnetConstructedDataProtocolVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (protocolVersion)
	lengthInBits += m.ProtocolVersion.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProtocolVersion) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProtocolVersion BACnetConstructedDataProtocolVersion, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolVersion, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "protocolVersion", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersion' field"))
	}
	m.ProtocolVersion = protocolVersion

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), protocolVersion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolVersion")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProtocolVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolVersion")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "protocolVersion", m.GetProtocolVersion(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolVersion' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolVersion")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolVersion) IsBACnetConstructedDataProtocolVersion() {}

func (m *_BACnetConstructedDataProtocolVersion) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProtocolVersion) deepCopy() *_BACnetConstructedDataProtocolVersion {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProtocolVersionCopy := &_BACnetConstructedDataProtocolVersion{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ProtocolVersion.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProtocolVersionCopy
}

func (m *_BACnetConstructedDataProtocolVersion) String() string {
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
