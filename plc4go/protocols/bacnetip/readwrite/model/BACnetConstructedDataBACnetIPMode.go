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

// BACnetConstructedDataBACnetIPMode is the corresponding interface of BACnetConstructedDataBACnetIPMode
type BACnetConstructedDataBACnetIPMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBacnetIpMode returns BacnetIpMode (property field)
	GetBacnetIpMode() BACnetIPModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetIPModeTagged
	// IsBACnetConstructedDataBACnetIPMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPMode()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPModeBuilder
	CreateBACnetConstructedDataBACnetIPModeBuilder() BACnetConstructedDataBACnetIPModeBuilder
}

// _BACnetConstructedDataBACnetIPMode is the data-structure of this message
type _BACnetConstructedDataBACnetIPMode struct {
	BACnetConstructedDataContract
	BacnetIpMode BACnetIPModeTagged
}

var _ BACnetConstructedDataBACnetIPMode = (*_BACnetConstructedDataBACnetIPMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPMode)(nil)

// NewBACnetConstructedDataBACnetIPMode factory function for _BACnetConstructedDataBACnetIPMode
func NewBACnetConstructedDataBACnetIPMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, bacnetIpMode BACnetIPModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPMode {
	if bacnetIpMode == nil {
		panic("bacnetIpMode of type BACnetIPModeTagged for BACnetConstructedDataBACnetIPMode must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BacnetIpMode:                  bacnetIpMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPModeBuilder is a builder for BACnetConstructedDataBACnetIPMode
type BACnetConstructedDataBACnetIPModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bacnetIpMode BACnetIPModeTagged) BACnetConstructedDataBACnetIPModeBuilder
	// WithBacnetIpMode adds BacnetIpMode (property field)
	WithBacnetIpMode(BACnetIPModeTagged) BACnetConstructedDataBACnetIPModeBuilder
	// WithBacnetIpModeBuilder adds BacnetIpMode (property field) which is build by the builder
	WithBacnetIpModeBuilder(func(BACnetIPModeTaggedBuilder) BACnetIPModeTaggedBuilder) BACnetConstructedDataBACnetIPModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBACnetIPMode or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPMode
}

// NewBACnetConstructedDataBACnetIPModeBuilder() creates a BACnetConstructedDataBACnetIPModeBuilder
func NewBACnetConstructedDataBACnetIPModeBuilder() BACnetConstructedDataBACnetIPModeBuilder {
	return &_BACnetConstructedDataBACnetIPModeBuilder{_BACnetConstructedDataBACnetIPMode: new(_BACnetConstructedDataBACnetIPMode)}
}

type _BACnetConstructedDataBACnetIPModeBuilder struct {
	*_BACnetConstructedDataBACnetIPMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPModeBuilder) = (*_BACnetConstructedDataBACnetIPModeBuilder)(nil)

func (b *_BACnetConstructedDataBACnetIPModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBACnetIPMode
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) WithMandatoryFields(bacnetIpMode BACnetIPModeTagged) BACnetConstructedDataBACnetIPModeBuilder {
	return b.WithBacnetIpMode(bacnetIpMode)
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) WithBacnetIpMode(bacnetIpMode BACnetIPModeTagged) BACnetConstructedDataBACnetIPModeBuilder {
	b.BacnetIpMode = bacnetIpMode
	return b
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) WithBacnetIpModeBuilder(builderSupplier func(BACnetIPModeTaggedBuilder) BACnetIPModeTaggedBuilder) BACnetConstructedDataBACnetIPModeBuilder {
	builder := builderSupplier(b.BacnetIpMode.CreateBACnetIPModeTaggedBuilder())
	var err error
	b.BacnetIpMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetIPModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) Build() (BACnetConstructedDataBACnetIPMode, error) {
	if b.BacnetIpMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bacnetIpMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBACnetIPMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) MustBuild() BACnetConstructedDataBACnetIPMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBACnetIPModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBACnetIPModeBuilder().(*_BACnetConstructedDataBACnetIPModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBACnetIPModeBuilder creates a BACnetConstructedDataBACnetIPModeBuilder
func (b *_BACnetConstructedDataBACnetIPMode) CreateBACnetConstructedDataBACnetIPModeBuilder() BACnetConstructedDataBACnetIPModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataBACnetIPModeBuilder()
	}
	return &_BACnetConstructedDataBACnetIPModeBuilder{_BACnetConstructedDataBACnetIPMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetBacnetIpMode() BACnetIPModeTagged {
	return m.BacnetIpMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetActualValue() BACnetIPModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetIPModeTagged(m.GetBacnetIpMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPMode(structType any) BACnetConstructedDataBACnetIPMode {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPMode) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPMode"
}

func (m *_BACnetConstructedDataBACnetIPMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (bacnetIpMode)
	lengthInBits += m.BacnetIpMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPMode BACnetConstructedDataBACnetIPMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetIpMode, err := ReadSimpleField[BACnetIPModeTagged](ctx, "bacnetIpMode", ReadComplex[BACnetIPModeTagged](BACnetIPModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetIpMode' field"))
	}
	m.BacnetIpMode = bacnetIpMode

	actualValue, err := ReadVirtualField[BACnetIPModeTagged](ctx, "actualValue", (*BACnetIPModeTagged)(nil), bacnetIpMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPMode")
		}

		if err := WriteSimpleField[BACnetIPModeTagged](ctx, "bacnetIpMode", m.GetBacnetIpMode(), WriteComplex[BACnetIPModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bacnetIpMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPMode) IsBACnetConstructedDataBACnetIPMode() {}

func (m *_BACnetConstructedDataBACnetIPMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPMode) deepCopy() *_BACnetConstructedDataBACnetIPMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPModeCopy := &_BACnetConstructedDataBACnetIPMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetIPModeTagged](m.BacnetIpMode),
	}
	_BACnetConstructedDataBACnetIPModeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPModeCopy
}

func (m *_BACnetConstructedDataBACnetIPMode) String() string {
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
