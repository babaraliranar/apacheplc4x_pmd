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

// AnsiExtendedSymbolSegment is the corresponding interface of AnsiExtendedSymbolSegment
type AnsiExtendedSymbolSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DataSegmentType
	// GetSymbol returns Symbol (property field)
	GetSymbol() string
	// GetPad returns Pad (property field)
	GetPad() *uint8
	// IsAnsiExtendedSymbolSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAnsiExtendedSymbolSegment()
	// CreateBuilder creates a AnsiExtendedSymbolSegmentBuilder
	CreateAnsiExtendedSymbolSegmentBuilder() AnsiExtendedSymbolSegmentBuilder
}

// _AnsiExtendedSymbolSegment is the data-structure of this message
type _AnsiExtendedSymbolSegment struct {
	DataSegmentTypeContract
	Symbol string
	Pad    *uint8
}

var _ AnsiExtendedSymbolSegment = (*_AnsiExtendedSymbolSegment)(nil)
var _ DataSegmentTypeRequirements = (*_AnsiExtendedSymbolSegment)(nil)

// NewAnsiExtendedSymbolSegment factory function for _AnsiExtendedSymbolSegment
func NewAnsiExtendedSymbolSegment(symbol string, pad *uint8) *_AnsiExtendedSymbolSegment {
	_result := &_AnsiExtendedSymbolSegment{
		DataSegmentTypeContract: NewDataSegmentType(),
		Symbol:                  symbol,
		Pad:                     pad,
	}
	_result.DataSegmentTypeContract.(*_DataSegmentType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AnsiExtendedSymbolSegmentBuilder is a builder for AnsiExtendedSymbolSegment
type AnsiExtendedSymbolSegmentBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(symbol string) AnsiExtendedSymbolSegmentBuilder
	// WithSymbol adds Symbol (property field)
	WithSymbol(string) AnsiExtendedSymbolSegmentBuilder
	// WithPad adds Pad (property field)
	WithOptionalPad(uint8) AnsiExtendedSymbolSegmentBuilder
	// Build builds the AnsiExtendedSymbolSegment or returns an error if something is wrong
	Build() (AnsiExtendedSymbolSegment, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AnsiExtendedSymbolSegment
}

// NewAnsiExtendedSymbolSegmentBuilder() creates a AnsiExtendedSymbolSegmentBuilder
func NewAnsiExtendedSymbolSegmentBuilder() AnsiExtendedSymbolSegmentBuilder {
	return &_AnsiExtendedSymbolSegmentBuilder{_AnsiExtendedSymbolSegment: new(_AnsiExtendedSymbolSegment)}
}

type _AnsiExtendedSymbolSegmentBuilder struct {
	*_AnsiExtendedSymbolSegment

	parentBuilder *_DataSegmentTypeBuilder

	err *utils.MultiError
}

var _ (AnsiExtendedSymbolSegmentBuilder) = (*_AnsiExtendedSymbolSegmentBuilder)(nil)

func (b *_AnsiExtendedSymbolSegmentBuilder) setParent(contract DataSegmentTypeContract) {
	b.DataSegmentTypeContract = contract
}

func (b *_AnsiExtendedSymbolSegmentBuilder) WithMandatoryFields(symbol string) AnsiExtendedSymbolSegmentBuilder {
	return b.WithSymbol(symbol)
}

func (b *_AnsiExtendedSymbolSegmentBuilder) WithSymbol(symbol string) AnsiExtendedSymbolSegmentBuilder {
	b.Symbol = symbol
	return b
}

func (b *_AnsiExtendedSymbolSegmentBuilder) WithOptionalPad(pad uint8) AnsiExtendedSymbolSegmentBuilder {
	b.Pad = &pad
	return b
}

func (b *_AnsiExtendedSymbolSegmentBuilder) Build() (AnsiExtendedSymbolSegment, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AnsiExtendedSymbolSegment.deepCopy(), nil
}

func (b *_AnsiExtendedSymbolSegmentBuilder) MustBuild() AnsiExtendedSymbolSegment {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AnsiExtendedSymbolSegmentBuilder) Done() DataSegmentTypeBuilder {
	return b.parentBuilder
}

func (b *_AnsiExtendedSymbolSegmentBuilder) buildForDataSegmentType() (DataSegmentType, error) {
	return b.Build()
}

func (b *_AnsiExtendedSymbolSegmentBuilder) DeepCopy() any {
	_copy := b.CreateAnsiExtendedSymbolSegmentBuilder().(*_AnsiExtendedSymbolSegmentBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAnsiExtendedSymbolSegmentBuilder creates a AnsiExtendedSymbolSegmentBuilder
func (b *_AnsiExtendedSymbolSegment) CreateAnsiExtendedSymbolSegmentBuilder() AnsiExtendedSymbolSegmentBuilder {
	if b == nil {
		return NewAnsiExtendedSymbolSegmentBuilder()
	}
	return &_AnsiExtendedSymbolSegmentBuilder{_AnsiExtendedSymbolSegment: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AnsiExtendedSymbolSegment) GetDataSegmentType() uint8 {
	return 0x11
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AnsiExtendedSymbolSegment) GetParent() DataSegmentTypeContract {
	return m.DataSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AnsiExtendedSymbolSegment) GetSymbol() string {
	return m.Symbol
}

func (m *_AnsiExtendedSymbolSegment) GetPad() *uint8 {
	return m.Pad
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAnsiExtendedSymbolSegment(structType any) AnsiExtendedSymbolSegment {
	if casted, ok := structType.(AnsiExtendedSymbolSegment); ok {
		return casted
	}
	if casted, ok := structType.(*AnsiExtendedSymbolSegment); ok {
		return *casted
	}
	return nil
}

func (m *_AnsiExtendedSymbolSegment) GetTypeName() string {
	return "AnsiExtendedSymbolSegment"
}

func (m *_AnsiExtendedSymbolSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DataSegmentTypeContract.(*_DataSegmentType).getLengthInBits(ctx))

	// Implicit Field (dataSize)
	lengthInBits += 8

	// Simple field (symbol)
	lengthInBits += uint16(int32(uint8(len(m.GetSymbol()))) * int32(int32(8)))

	// Optional Field (pad)
	if m.Pad != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_AnsiExtendedSymbolSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AnsiExtendedSymbolSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DataSegmentType) (__ansiExtendedSymbolSegment AnsiExtendedSymbolSegment, err error) {
	m.DataSegmentTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AnsiExtendedSymbolSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AnsiExtendedSymbolSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataSize, err := ReadImplicitField[uint8](ctx, "dataSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSize' field"))
	}
	_ = dataSize

	symbol, err := ReadSimpleField(ctx, "symbol", ReadString(readBuffer, uint32(int32(dataSize)*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'symbol' field"))
	}
	m.Symbol = symbol

	var pad *uint8
	pad, err = ReadOptionalField[uint8](ctx, "pad", ReadUnsignedByte(readBuffer, uint8(8)), bool(((len(symbol))%(2)) != (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pad' field"))
	}
	m.Pad = pad

	if closeErr := readBuffer.CloseContext("AnsiExtendedSymbolSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AnsiExtendedSymbolSegment")
	}

	return m, nil
}

func (m *_AnsiExtendedSymbolSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AnsiExtendedSymbolSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AnsiExtendedSymbolSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AnsiExtendedSymbolSegment")
		}
		dataSize := uint8(uint8(len(m.GetSymbol())))
		if err := WriteImplicitField(ctx, "dataSize", dataSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSize' field")
		}

		if err := WriteSimpleField[string](ctx, "symbol", m.GetSymbol(), WriteString(writeBuffer, int32(int32(uint8(len(m.GetSymbol())))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'symbol' field")
		}

		if err := WriteOptionalField[uint8](ctx, "pad", m.GetPad(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'pad' field")
		}

		if popErr := writeBuffer.PopContext("AnsiExtendedSymbolSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AnsiExtendedSymbolSegment")
		}
		return nil
	}
	return m.DataSegmentTypeContract.(*_DataSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AnsiExtendedSymbolSegment) IsAnsiExtendedSymbolSegment() {}

func (m *_AnsiExtendedSymbolSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AnsiExtendedSymbolSegment) deepCopy() *_AnsiExtendedSymbolSegment {
	if m == nil {
		return nil
	}
	_AnsiExtendedSymbolSegmentCopy := &_AnsiExtendedSymbolSegment{
		m.DataSegmentTypeContract.(*_DataSegmentType).deepCopy(),
		m.Symbol,
		utils.CopyPtr[uint8](m.Pad),
	}
	m.DataSegmentTypeContract.(*_DataSegmentType)._SubType = m
	return _AnsiExtendedSymbolSegmentCopy
}

func (m *_AnsiExtendedSymbolSegment) String() string {
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
