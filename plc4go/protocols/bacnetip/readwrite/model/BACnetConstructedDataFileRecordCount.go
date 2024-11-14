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

// BACnetConstructedDataFileRecordCount is the corresponding interface of BACnetConstructedDataFileRecordCount
type BACnetConstructedDataFileRecordCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRecordCount returns RecordCount (property field)
	GetRecordCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataFileRecordCount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFileRecordCount()
	// CreateBuilder creates a BACnetConstructedDataFileRecordCountBuilder
	CreateBACnetConstructedDataFileRecordCountBuilder() BACnetConstructedDataFileRecordCountBuilder
}

// _BACnetConstructedDataFileRecordCount is the data-structure of this message
type _BACnetConstructedDataFileRecordCount struct {
	BACnetConstructedDataContract
	RecordCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataFileRecordCount = (*_BACnetConstructedDataFileRecordCount)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFileRecordCount)(nil)

// NewBACnetConstructedDataFileRecordCount factory function for _BACnetConstructedDataFileRecordCount
func NewBACnetConstructedDataFileRecordCount(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, recordCount BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFileRecordCount {
	if recordCount == nil {
		panic("recordCount of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataFileRecordCount must not be nil")
	}
	_result := &_BACnetConstructedDataFileRecordCount{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RecordCount:                   recordCount,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFileRecordCountBuilder is a builder for BACnetConstructedDataFileRecordCount
type BACnetConstructedDataFileRecordCountBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recordCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFileRecordCountBuilder
	// WithRecordCount adds RecordCount (property field)
	WithRecordCount(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFileRecordCountBuilder
	// WithRecordCountBuilder adds RecordCount (property field) which is build by the builder
	WithRecordCountBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFileRecordCountBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataFileRecordCount or returns an error if something is wrong
	Build() (BACnetConstructedDataFileRecordCount, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFileRecordCount
}

// NewBACnetConstructedDataFileRecordCountBuilder() creates a BACnetConstructedDataFileRecordCountBuilder
func NewBACnetConstructedDataFileRecordCountBuilder() BACnetConstructedDataFileRecordCountBuilder {
	return &_BACnetConstructedDataFileRecordCountBuilder{_BACnetConstructedDataFileRecordCount: new(_BACnetConstructedDataFileRecordCount)}
}

type _BACnetConstructedDataFileRecordCountBuilder struct {
	*_BACnetConstructedDataFileRecordCount

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFileRecordCountBuilder) = (*_BACnetConstructedDataFileRecordCountBuilder)(nil)

func (b *_BACnetConstructedDataFileRecordCountBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataFileRecordCount
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) WithMandatoryFields(recordCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFileRecordCountBuilder {
	return b.WithRecordCount(recordCount)
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) WithRecordCount(recordCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFileRecordCountBuilder {
	b.RecordCount = recordCount
	return b
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) WithRecordCountBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFileRecordCountBuilder {
	builder := builderSupplier(b.RecordCount.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RecordCount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) Build() (BACnetConstructedDataFileRecordCount, error) {
	if b.RecordCount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'recordCount' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFileRecordCount.deepCopy(), nil
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) MustBuild() BACnetConstructedDataFileRecordCount {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFileRecordCountBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFileRecordCountBuilder().(*_BACnetConstructedDataFileRecordCountBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFileRecordCountBuilder creates a BACnetConstructedDataFileRecordCountBuilder
func (b *_BACnetConstructedDataFileRecordCount) CreateBACnetConstructedDataFileRecordCountBuilder() BACnetConstructedDataFileRecordCountBuilder {
	if b == nil {
		return NewBACnetConstructedDataFileRecordCountBuilder()
	}
	return &_BACnetConstructedDataFileRecordCountBuilder{_BACnetConstructedDataFileRecordCount: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_FILE
}

func (m *_BACnetConstructedDataFileRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFileRecordCount(structType any) BACnetConstructedDataFileRecordCount {
	if casted, ok := structType.(BACnetConstructedDataFileRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFileRecordCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFileRecordCount) GetTypeName() string {
	return "BACnetConstructedDataFileRecordCount"
}

func (m *_BACnetConstructedDataFileRecordCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (recordCount)
	lengthInBits += m.RecordCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFileRecordCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFileRecordCount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFileRecordCount BACnetConstructedDataFileRecordCount, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFileRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFileRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recordCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordCount' field"))
	}
	m.RecordCount = recordCount

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), recordCount)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFileRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFileRecordCount")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFileRecordCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFileRecordCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFileRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFileRecordCount")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordCount", m.GetRecordCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'recordCount' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFileRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFileRecordCount")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFileRecordCount) IsBACnetConstructedDataFileRecordCount() {}

func (m *_BACnetConstructedDataFileRecordCount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFileRecordCount) deepCopy() *_BACnetConstructedDataFileRecordCount {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFileRecordCountCopy := &_BACnetConstructedDataFileRecordCount{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RecordCount),
	}
	_BACnetConstructedDataFileRecordCountCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFileRecordCountCopy
}

func (m *_BACnetConstructedDataFileRecordCount) String() string {
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
