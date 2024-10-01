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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCWriteBroadcastDistributionTable is the corresponding interface of BVLCWriteBroadcastDistributionTable
type BVLCWriteBroadcastDistributionTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// GetTable returns Table (property field)
	GetTable() []BVLCBroadcastDistributionTableEntry
	// IsBVLCWriteBroadcastDistributionTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCWriteBroadcastDistributionTable()
	// CreateBuilder creates a BVLCWriteBroadcastDistributionTableBuilder
	CreateBVLCWriteBroadcastDistributionTableBuilder() BVLCWriteBroadcastDistributionTableBuilder
}

// _BVLCWriteBroadcastDistributionTable is the data-structure of this message
type _BVLCWriteBroadcastDistributionTable struct {
	BVLCContract
	Table []BVLCBroadcastDistributionTableEntry

	// Arguments.
	BvlcPayloadLength uint16
}

var _ BVLCWriteBroadcastDistributionTable = (*_BVLCWriteBroadcastDistributionTable)(nil)
var _ BVLCRequirements = (*_BVLCWriteBroadcastDistributionTable)(nil)

// NewBVLCWriteBroadcastDistributionTable factory function for _BVLCWriteBroadcastDistributionTable
func NewBVLCWriteBroadcastDistributionTable(table []BVLCBroadcastDistributionTableEntry, bvlcPayloadLength uint16) *_BVLCWriteBroadcastDistributionTable {
	_result := &_BVLCWriteBroadcastDistributionTable{
		BVLCContract: NewBVLC(),
		Table:        table,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCWriteBroadcastDistributionTableBuilder is a builder for BVLCWriteBroadcastDistributionTable
type BVLCWriteBroadcastDistributionTableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(table []BVLCBroadcastDistributionTableEntry) BVLCWriteBroadcastDistributionTableBuilder
	// WithTable adds Table (property field)
	WithTable(...BVLCBroadcastDistributionTableEntry) BVLCWriteBroadcastDistributionTableBuilder
	// Build builds the BVLCWriteBroadcastDistributionTable or returns an error if something is wrong
	Build() (BVLCWriteBroadcastDistributionTable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCWriteBroadcastDistributionTable
}

// NewBVLCWriteBroadcastDistributionTableBuilder() creates a BVLCWriteBroadcastDistributionTableBuilder
func NewBVLCWriteBroadcastDistributionTableBuilder() BVLCWriteBroadcastDistributionTableBuilder {
	return &_BVLCWriteBroadcastDistributionTableBuilder{_BVLCWriteBroadcastDistributionTable: new(_BVLCWriteBroadcastDistributionTable)}
}

type _BVLCWriteBroadcastDistributionTableBuilder struct {
	*_BVLCWriteBroadcastDistributionTable

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCWriteBroadcastDistributionTableBuilder) = (*_BVLCWriteBroadcastDistributionTableBuilder)(nil)

func (b *_BVLCWriteBroadcastDistributionTableBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
}

func (b *_BVLCWriteBroadcastDistributionTableBuilder) WithMandatoryFields(table []BVLCBroadcastDistributionTableEntry) BVLCWriteBroadcastDistributionTableBuilder {
	return b.WithTable(table...)
}

func (b *_BVLCWriteBroadcastDistributionTableBuilder) WithTable(table ...BVLCBroadcastDistributionTableEntry) BVLCWriteBroadcastDistributionTableBuilder {
	b.Table = table
	return b
}

func (b *_BVLCWriteBroadcastDistributionTableBuilder) Build() (BVLCWriteBroadcastDistributionTable, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCWriteBroadcastDistributionTable.deepCopy(), nil
}

func (b *_BVLCWriteBroadcastDistributionTableBuilder) MustBuild() BVLCWriteBroadcastDistributionTable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BVLCWriteBroadcastDistributionTableBuilder) Done() BVLCBuilder {
	return b.parentBuilder
}

func (b *_BVLCWriteBroadcastDistributionTableBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCWriteBroadcastDistributionTableBuilder) DeepCopy() any {
	_copy := b.CreateBVLCWriteBroadcastDistributionTableBuilder().(*_BVLCWriteBroadcastDistributionTableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCWriteBroadcastDistributionTableBuilder creates a BVLCWriteBroadcastDistributionTableBuilder
func (b *_BVLCWriteBroadcastDistributionTable) CreateBVLCWriteBroadcastDistributionTableBuilder() BVLCWriteBroadcastDistributionTableBuilder {
	if b == nil {
		return NewBVLCWriteBroadcastDistributionTableBuilder()
	}
	return &_BVLCWriteBroadcastDistributionTableBuilder{_BVLCWriteBroadcastDistributionTable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCWriteBroadcastDistributionTable) GetBvlcFunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCWriteBroadcastDistributionTable) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCWriteBroadcastDistributionTable) GetTable() []BVLCBroadcastDistributionTableEntry {
	return m.Table
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCWriteBroadcastDistributionTable(structType any) BVLCWriteBroadcastDistributionTable {
	if casted, ok := structType.(BVLCWriteBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCWriteBroadcastDistributionTable); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCWriteBroadcastDistributionTable) GetTypeName() string {
	return "BVLCWriteBroadcastDistributionTable"
}

func (m *_BVLCWriteBroadcastDistributionTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BVLCWriteBroadcastDistributionTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCWriteBroadcastDistributionTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC, bvlcPayloadLength uint16) (__bVLCWriteBroadcastDistributionTable BVLCWriteBroadcastDistributionTable, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCWriteBroadcastDistributionTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCWriteBroadcastDistributionTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	table, err := ReadLengthArrayField[BVLCBroadcastDistributionTableEntry](ctx, "table", ReadComplex[BVLCBroadcastDistributionTableEntry](BVLCBroadcastDistributionTableEntryParseWithBuffer, readBuffer), int(bvlcPayloadLength), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'table' field"))
	}
	m.Table = table

	if closeErr := readBuffer.CloseContext("BVLCWriteBroadcastDistributionTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCWriteBroadcastDistributionTable")
	}

	return m, nil
}

func (m *_BVLCWriteBroadcastDistributionTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCWriteBroadcastDistributionTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCWriteBroadcastDistributionTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCWriteBroadcastDistributionTable")
		}

		if err := WriteComplexTypeArrayField(ctx, "table", m.GetTable(), writeBuffer, codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'table' field")
		}

		if popErr := writeBuffer.PopContext("BVLCWriteBroadcastDistributionTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCWriteBroadcastDistributionTable")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCWriteBroadcastDistributionTable) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCWriteBroadcastDistributionTable) IsBVLCWriteBroadcastDistributionTable() {}

func (m *_BVLCWriteBroadcastDistributionTable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCWriteBroadcastDistributionTable) deepCopy() *_BVLCWriteBroadcastDistributionTable {
	if m == nil {
		return nil
	}
	_BVLCWriteBroadcastDistributionTableCopy := &_BVLCWriteBroadcastDistributionTable{
		m.BVLCContract.(*_BVLC).deepCopy(),
		utils.DeepCopySlice[BVLCBroadcastDistributionTableEntry, BVLCBroadcastDistributionTableEntry](m.Table),
		m.BvlcPayloadLength,
	}
	m.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCWriteBroadcastDistributionTableCopy
}

func (m *_BVLCWriteBroadcastDistributionTable) String() string {
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
