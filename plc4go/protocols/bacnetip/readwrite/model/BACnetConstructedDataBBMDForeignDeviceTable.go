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

// BACnetConstructedDataBBMDForeignDeviceTable is the corresponding interface of BACnetConstructedDataBBMDForeignDeviceTable
type BACnetConstructedDataBBMDForeignDeviceTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBbmdForeignDeviceTable returns BbmdForeignDeviceTable (property field)
	GetBbmdForeignDeviceTable() []BACnetBDTEntry
	// IsBACnetConstructedDataBBMDForeignDeviceTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBBMDForeignDeviceTable()
}

// _BACnetConstructedDataBBMDForeignDeviceTable is the data-structure of this message
type _BACnetConstructedDataBBMDForeignDeviceTable struct {
	BACnetConstructedDataContract
	BbmdForeignDeviceTable []BACnetBDTEntry
}

var _ BACnetConstructedDataBBMDForeignDeviceTable = (*_BACnetConstructedDataBBMDForeignDeviceTable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBBMDForeignDeviceTable)(nil)

// NewBACnetConstructedDataBBMDForeignDeviceTable factory function for _BACnetConstructedDataBBMDForeignDeviceTable
func NewBACnetConstructedDataBBMDForeignDeviceTable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, bbmdForeignDeviceTable []BACnetBDTEntry, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBBMDForeignDeviceTable {
	_result := &_BACnetConstructedDataBBMDForeignDeviceTable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BbmdForeignDeviceTable:        bbmdForeignDeviceTable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_FOREIGN_DEVICE_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetBbmdForeignDeviceTable() []BACnetBDTEntry {
	return m.BbmdForeignDeviceTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBBMDForeignDeviceTable(structType any) BACnetConstructedDataBBMDForeignDeviceTable {
	if casted, ok := structType.(BACnetConstructedDataBBMDForeignDeviceTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDForeignDeviceTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetTypeName() string {
	return "BACnetConstructedDataBBMDForeignDeviceTable"
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.BbmdForeignDeviceTable) > 0 {
		for _, element := range m.BbmdForeignDeviceTable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBBMDForeignDeviceTable BACnetConstructedDataBBMDForeignDeviceTable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDForeignDeviceTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDForeignDeviceTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bbmdForeignDeviceTable, err := ReadTerminatedArrayField[BACnetBDTEntry](ctx, "bbmdForeignDeviceTable", ReadComplex[BACnetBDTEntry](BACnetBDTEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bbmdForeignDeviceTable' field"))
	}
	m.BbmdForeignDeviceTable = bbmdForeignDeviceTable

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDForeignDeviceTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDForeignDeviceTable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDForeignDeviceTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDForeignDeviceTable")
		}

		if err := WriteComplexTypeArrayField(ctx, "bbmdForeignDeviceTable", m.GetBbmdForeignDeviceTable(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'bbmdForeignDeviceTable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDForeignDeviceTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDForeignDeviceTable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) IsBACnetConstructedDataBBMDForeignDeviceTable() {
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) deepCopy() *_BACnetConstructedDataBBMDForeignDeviceTable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBBMDForeignDeviceTableCopy := &_BACnetConstructedDataBBMDForeignDeviceTable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetBDTEntry, BACnetBDTEntry](m.BbmdForeignDeviceTable),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBBMDForeignDeviceTableCopy
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
