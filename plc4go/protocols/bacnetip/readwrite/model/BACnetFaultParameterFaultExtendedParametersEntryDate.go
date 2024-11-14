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

// BACnetFaultParameterFaultExtendedParametersEntryDate is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryDate
type BACnetFaultParameterFaultExtendedParametersEntryDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// IsBACnetFaultParameterFaultExtendedParametersEntryDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryDate()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryDateBuilder() BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryDate is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryDate struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	DateValue BACnetApplicationTagDate
}

var _ BACnetFaultParameterFaultExtendedParametersEntryDate = (*_BACnetFaultParameterFaultExtendedParametersEntryDate)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryDate)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryDate factory function for _BACnetFaultParameterFaultExtendedParametersEntryDate
func NewBACnetFaultParameterFaultExtendedParametersEntryDate(peekedTagHeader BACnetTagHeader, dateValue BACnetApplicationTagDate) *_BACnetFaultParameterFaultExtendedParametersEntryDate {
	if dateValue == nil {
		panic("dateValue of type BACnetApplicationTagDate for BACnetFaultParameterFaultExtendedParametersEntryDate must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryDate{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		DateValue: dateValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryDateBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryDate
type BACnetFaultParameterFaultExtendedParametersEntryDateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateValue BACnetApplicationTagDate) BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
	// WithDateValue adds DateValue (property field)
	WithDateValue(BACnetApplicationTagDate) BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
	// WithDateValueBuilder adds DateValue (property field) which is build by the builder
	WithDateValueBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryDate or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryDate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryDate
}

// NewBACnetFaultParameterFaultExtendedParametersEntryDateBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryDateBuilder() BACnetFaultParameterFaultExtendedParametersEntryDateBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder{_BACnetFaultParameterFaultExtendedParametersEntryDate: new(_BACnetFaultParameterFaultExtendedParametersEntryDate)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryDateBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryDate

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryDate
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) WithMandatoryFields(dateValue BACnetApplicationTagDate) BACnetFaultParameterFaultExtendedParametersEntryDateBuilder {
	return b.WithDateValue(dateValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) WithDateValue(dateValue BACnetApplicationTagDate) BACnetFaultParameterFaultExtendedParametersEntryDateBuilder {
	b.DateValue = dateValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) WithDateValueBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetFaultParameterFaultExtendedParametersEntryDateBuilder {
	builder := builderSupplier(b.DateValue.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.DateValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryDate, error) {
	if b.DateValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dateValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryDate.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryDate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryDateBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryDateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryDateBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryDate) CreateBACnetFaultParameterFaultExtendedParametersEntryDateBuilder() BACnetFaultParameterFaultExtendedParametersEntryDateBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryDateBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryDateBuilder{_BACnetFaultParameterFaultExtendedParametersEntryDate: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryDate(structType any) BACnetFaultParameterFaultExtendedParametersEntryDate {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryDate"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryDate BACnetFaultParameterFaultExtendedParametersEntryDate, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	m.DateValue = dateValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryDate")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryDate")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "dateValue", m.GetDateValue(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryDate")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) IsBACnetFaultParameterFaultExtendedParametersEntryDate() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryDate {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryDateCopy := &_BACnetFaultParameterFaultExtendedParametersEntryDate{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDate](m.DateValue),
	}
	_BACnetFaultParameterFaultExtendedParametersEntryDateCopy.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryDateCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) String() string {
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
