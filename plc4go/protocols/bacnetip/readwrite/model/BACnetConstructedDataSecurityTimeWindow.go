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

// BACnetConstructedDataSecurityTimeWindow is the corresponding interface of BACnetConstructedDataSecurityTimeWindow
type BACnetConstructedDataSecurityTimeWindow interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSecurityTimeWindow returns SecurityTimeWindow (property field)
	GetSecurityTimeWindow() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataSecurityTimeWindow is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSecurityTimeWindow()
	// CreateBuilder creates a BACnetConstructedDataSecurityTimeWindowBuilder
	CreateBACnetConstructedDataSecurityTimeWindowBuilder() BACnetConstructedDataSecurityTimeWindowBuilder
}

// _BACnetConstructedDataSecurityTimeWindow is the data-structure of this message
type _BACnetConstructedDataSecurityTimeWindow struct {
	BACnetConstructedDataContract
	SecurityTimeWindow BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataSecurityTimeWindow = (*_BACnetConstructedDataSecurityTimeWindow)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSecurityTimeWindow)(nil)

// NewBACnetConstructedDataSecurityTimeWindow factory function for _BACnetConstructedDataSecurityTimeWindow
func NewBACnetConstructedDataSecurityTimeWindow(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, securityTimeWindow BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSecurityTimeWindow {
	if securityTimeWindow == nil {
		panic("securityTimeWindow of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataSecurityTimeWindow must not be nil")
	}
	_result := &_BACnetConstructedDataSecurityTimeWindow{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SecurityTimeWindow:            securityTimeWindow,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSecurityTimeWindowBuilder is a builder for BACnetConstructedDataSecurityTimeWindow
type BACnetConstructedDataSecurityTimeWindowBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(securityTimeWindow BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSecurityTimeWindowBuilder
	// WithSecurityTimeWindow adds SecurityTimeWindow (property field)
	WithSecurityTimeWindow(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSecurityTimeWindowBuilder
	// WithSecurityTimeWindowBuilder adds SecurityTimeWindow (property field) which is build by the builder
	WithSecurityTimeWindowBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSecurityTimeWindowBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataSecurityTimeWindow or returns an error if something is wrong
	Build() (BACnetConstructedDataSecurityTimeWindow, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSecurityTimeWindow
}

// NewBACnetConstructedDataSecurityTimeWindowBuilder() creates a BACnetConstructedDataSecurityTimeWindowBuilder
func NewBACnetConstructedDataSecurityTimeWindowBuilder() BACnetConstructedDataSecurityTimeWindowBuilder {
	return &_BACnetConstructedDataSecurityTimeWindowBuilder{_BACnetConstructedDataSecurityTimeWindow: new(_BACnetConstructedDataSecurityTimeWindow)}
}

type _BACnetConstructedDataSecurityTimeWindowBuilder struct {
	*_BACnetConstructedDataSecurityTimeWindow

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSecurityTimeWindowBuilder) = (*_BACnetConstructedDataSecurityTimeWindowBuilder)(nil)

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataSecurityTimeWindow
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) WithMandatoryFields(securityTimeWindow BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSecurityTimeWindowBuilder {
	return b.WithSecurityTimeWindow(securityTimeWindow)
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) WithSecurityTimeWindow(securityTimeWindow BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSecurityTimeWindowBuilder {
	b.SecurityTimeWindow = securityTimeWindow
	return b
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) WithSecurityTimeWindowBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSecurityTimeWindowBuilder {
	builder := builderSupplier(b.SecurityTimeWindow.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.SecurityTimeWindow, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) Build() (BACnetConstructedDataSecurityTimeWindow, error) {
	if b.SecurityTimeWindow == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityTimeWindow' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSecurityTimeWindow.deepCopy(), nil
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) MustBuild() BACnetConstructedDataSecurityTimeWindow {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSecurityTimeWindowBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSecurityTimeWindowBuilder().(*_BACnetConstructedDataSecurityTimeWindowBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSecurityTimeWindowBuilder creates a BACnetConstructedDataSecurityTimeWindowBuilder
func (b *_BACnetConstructedDataSecurityTimeWindow) CreateBACnetConstructedDataSecurityTimeWindowBuilder() BACnetConstructedDataSecurityTimeWindowBuilder {
	if b == nil {
		return NewBACnetConstructedDataSecurityTimeWindowBuilder()
	}
	return &_BACnetConstructedDataSecurityTimeWindowBuilder{_BACnetConstructedDataSecurityTimeWindow: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURITY_TIME_WINDOW
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetSecurityTimeWindow() BACnetApplicationTagUnsignedInteger {
	return m.SecurityTimeWindow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSecurityTimeWindow) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetSecurityTimeWindow())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSecurityTimeWindow(structType any) BACnetConstructedDataSecurityTimeWindow {
	if casted, ok := structType.(BACnetConstructedDataSecurityTimeWindow); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecurityTimeWindow); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetTypeName() string {
	return "BACnetConstructedDataSecurityTimeWindow"
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (securityTimeWindow)
	lengthInBits += m.SecurityTimeWindow.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSecurityTimeWindow) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSecurityTimeWindow) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSecurityTimeWindow BACnetConstructedDataSecurityTimeWindow, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecurityTimeWindow"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecurityTimeWindow")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securityTimeWindow, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "securityTimeWindow", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityTimeWindow' field"))
	}
	m.SecurityTimeWindow = securityTimeWindow

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), securityTimeWindow)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecurityTimeWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecurityTimeWindow")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSecurityTimeWindow) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSecurityTimeWindow) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecurityTimeWindow"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecurityTimeWindow")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "securityTimeWindow", m.GetSecurityTimeWindow(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityTimeWindow' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecurityTimeWindow"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecurityTimeWindow")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSecurityTimeWindow) IsBACnetConstructedDataSecurityTimeWindow() {}

func (m *_BACnetConstructedDataSecurityTimeWindow) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSecurityTimeWindow) deepCopy() *_BACnetConstructedDataSecurityTimeWindow {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSecurityTimeWindowCopy := &_BACnetConstructedDataSecurityTimeWindow{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.SecurityTimeWindow),
	}
	_BACnetConstructedDataSecurityTimeWindowCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSecurityTimeWindowCopy
}

func (m *_BACnetConstructedDataSecurityTimeWindow) String() string {
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
