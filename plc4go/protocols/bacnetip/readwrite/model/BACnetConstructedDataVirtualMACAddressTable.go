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

// BACnetConstructedDataVirtualMACAddressTable is the corresponding interface of BACnetConstructedDataVirtualMACAddressTable
type BACnetConstructedDataVirtualMACAddressTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetVirtualMacAddressTable returns VirtualMacAddressTable (property field)
	GetVirtualMacAddressTable() []BACnetVMACEntry
	// IsBACnetConstructedDataVirtualMACAddressTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataVirtualMACAddressTable()
	// CreateBuilder creates a BACnetConstructedDataVirtualMACAddressTableBuilder
	CreateBACnetConstructedDataVirtualMACAddressTableBuilder() BACnetConstructedDataVirtualMACAddressTableBuilder
}

// _BACnetConstructedDataVirtualMACAddressTable is the data-structure of this message
type _BACnetConstructedDataVirtualMACAddressTable struct {
	BACnetConstructedDataContract
	VirtualMacAddressTable []BACnetVMACEntry
}

var _ BACnetConstructedDataVirtualMACAddressTable = (*_BACnetConstructedDataVirtualMACAddressTable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataVirtualMACAddressTable)(nil)

// NewBACnetConstructedDataVirtualMACAddressTable factory function for _BACnetConstructedDataVirtualMACAddressTable
func NewBACnetConstructedDataVirtualMACAddressTable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, virtualMacAddressTable []BACnetVMACEntry, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVirtualMACAddressTable {
	_result := &_BACnetConstructedDataVirtualMACAddressTable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		VirtualMacAddressTable:        virtualMacAddressTable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataVirtualMACAddressTableBuilder is a builder for BACnetConstructedDataVirtualMACAddressTable
type BACnetConstructedDataVirtualMACAddressTableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(virtualMacAddressTable []BACnetVMACEntry) BACnetConstructedDataVirtualMACAddressTableBuilder
	// WithVirtualMacAddressTable adds VirtualMacAddressTable (property field)
	WithVirtualMacAddressTable(...BACnetVMACEntry) BACnetConstructedDataVirtualMACAddressTableBuilder
	// Build builds the BACnetConstructedDataVirtualMACAddressTable or returns an error if something is wrong
	Build() (BACnetConstructedDataVirtualMACAddressTable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataVirtualMACAddressTable
}

// NewBACnetConstructedDataVirtualMACAddressTableBuilder() creates a BACnetConstructedDataVirtualMACAddressTableBuilder
func NewBACnetConstructedDataVirtualMACAddressTableBuilder() BACnetConstructedDataVirtualMACAddressTableBuilder {
	return &_BACnetConstructedDataVirtualMACAddressTableBuilder{_BACnetConstructedDataVirtualMACAddressTable: new(_BACnetConstructedDataVirtualMACAddressTable)}
}

type _BACnetConstructedDataVirtualMACAddressTableBuilder struct {
	*_BACnetConstructedDataVirtualMACAddressTable

	err *utils.MultiError
}

var _ (BACnetConstructedDataVirtualMACAddressTableBuilder) = (*_BACnetConstructedDataVirtualMACAddressTableBuilder)(nil)

func (m *_BACnetConstructedDataVirtualMACAddressTableBuilder) WithMandatoryFields(virtualMacAddressTable []BACnetVMACEntry) BACnetConstructedDataVirtualMACAddressTableBuilder {
	return m.WithVirtualMacAddressTable(virtualMacAddressTable...)
}

func (m *_BACnetConstructedDataVirtualMACAddressTableBuilder) WithVirtualMacAddressTable(virtualMacAddressTable ...BACnetVMACEntry) BACnetConstructedDataVirtualMACAddressTableBuilder {
	m.VirtualMacAddressTable = virtualMacAddressTable
	return m
}

func (m *_BACnetConstructedDataVirtualMACAddressTableBuilder) Build() (BACnetConstructedDataVirtualMACAddressTable, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataVirtualMACAddressTable.deepCopy(), nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTableBuilder) MustBuild() BACnetConstructedDataVirtualMACAddressTable {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataVirtualMACAddressTableBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataVirtualMACAddressTableBuilder()
}

// CreateBACnetConstructedDataVirtualMACAddressTableBuilder creates a BACnetConstructedDataVirtualMACAddressTableBuilder
func (m *_BACnetConstructedDataVirtualMACAddressTable) CreateBACnetConstructedDataVirtualMACAddressTableBuilder() BACnetConstructedDataVirtualMACAddressTableBuilder {
	if m == nil {
		return NewBACnetConstructedDataVirtualMACAddressTableBuilder()
	}
	return &_BACnetConstructedDataVirtualMACAddressTableBuilder{_BACnetConstructedDataVirtualMACAddressTable: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VIRTUAL_MAC_ADDRESS_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetVirtualMacAddressTable() []BACnetVMACEntry {
	return m.VirtualMacAddressTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVirtualMACAddressTable(structType any) BACnetConstructedDataVirtualMACAddressTable {
	if casted, ok := structType.(BACnetConstructedDataVirtualMACAddressTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVirtualMACAddressTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetTypeName() string {
	return "BACnetConstructedDataVirtualMACAddressTable"
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.VirtualMacAddressTable) > 0 {
		for _, element := range m.VirtualMacAddressTable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataVirtualMACAddressTable BACnetConstructedDataVirtualMACAddressTable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVirtualMACAddressTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVirtualMACAddressTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	virtualMacAddressTable, err := ReadTerminatedArrayField[BACnetVMACEntry](ctx, "virtualMacAddressTable", ReadComplex[BACnetVMACEntry](BACnetVMACEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'virtualMacAddressTable' field"))
	}
	m.VirtualMacAddressTable = virtualMacAddressTable

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVirtualMACAddressTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVirtualMACAddressTable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVirtualMACAddressTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVirtualMACAddressTable")
		}

		if err := WriteComplexTypeArrayField(ctx, "virtualMacAddressTable", m.GetVirtualMacAddressTable(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'virtualMacAddressTable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVirtualMACAddressTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVirtualMACAddressTable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) IsBACnetConstructedDataVirtualMACAddressTable() {
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) deepCopy() *_BACnetConstructedDataVirtualMACAddressTable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataVirtualMACAddressTableCopy := &_BACnetConstructedDataVirtualMACAddressTable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetVMACEntry, BACnetVMACEntry](m.VirtualMacAddressTable),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataVirtualMACAddressTableCopy
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
