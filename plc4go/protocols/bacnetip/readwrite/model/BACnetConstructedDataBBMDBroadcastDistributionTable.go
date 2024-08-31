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

// BACnetConstructedDataBBMDBroadcastDistributionTable is the corresponding interface of BACnetConstructedDataBBMDBroadcastDistributionTable
type BACnetConstructedDataBBMDBroadcastDistributionTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBbmdBroadcastDistributionTable returns BbmdBroadcastDistributionTable (property field)
	GetBbmdBroadcastDistributionTable() []BACnetBDTEntry
}

// BACnetConstructedDataBBMDBroadcastDistributionTableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBBMDBroadcastDistributionTable.
// This is useful for switch cases.
type BACnetConstructedDataBBMDBroadcastDistributionTableExactly interface {
	BACnetConstructedDataBBMDBroadcastDistributionTable
	isBACnetConstructedDataBBMDBroadcastDistributionTable() bool
}

// _BACnetConstructedDataBBMDBroadcastDistributionTable is the data-structure of this message
type _BACnetConstructedDataBBMDBroadcastDistributionTable struct {
	*_BACnetConstructedData
	BbmdBroadcastDistributionTable []BACnetBDTEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_BROADCAST_DISTRIBUTION_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetBbmdBroadcastDistributionTable() []BACnetBDTEntry {
	return m.BbmdBroadcastDistributionTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBBMDBroadcastDistributionTable factory function for _BACnetConstructedDataBBMDBroadcastDistributionTable
func NewBACnetConstructedDataBBMDBroadcastDistributionTable(bbmdBroadcastDistributionTable []BACnetBDTEntry, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBBMDBroadcastDistributionTable {
	_result := &_BACnetConstructedDataBBMDBroadcastDistributionTable{
		BbmdBroadcastDistributionTable: bbmdBroadcastDistributionTable,
		_BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBBMDBroadcastDistributionTable(structType any) BACnetConstructedDataBBMDBroadcastDistributionTable {
	if casted, ok := structType.(BACnetConstructedDataBBMDBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDBroadcastDistributionTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetTypeName() string {
	return "BACnetConstructedDataBBMDBroadcastDistributionTable"
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.BbmdBroadcastDistributionTable) > 0 {
		for _, element := range m.BbmdBroadcastDistributionTable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBBMDBroadcastDistributionTableParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDBroadcastDistributionTable, error) {
	return BACnetConstructedDataBBMDBroadcastDistributionTableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBBMDBroadcastDistributionTableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDBroadcastDistributionTable, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDBroadcastDistributionTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bbmdBroadcastDistributionTable, err := ReadTerminatedArrayField[BACnetBDTEntry](ctx, "bbmdBroadcastDistributionTable", ReadComplex[BACnetBDTEntry](BACnetBDTEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bbmdBroadcastDistributionTable' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDBroadcastDistributionTable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBBMDBroadcastDistributionTable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BbmdBroadcastDistributionTable: bbmdBroadcastDistributionTable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDBroadcastDistributionTable")
		}

		// Array Field (bbmdBroadcastDistributionTable)
		if pushErr := writeBuffer.PushContext("bbmdBroadcastDistributionTable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bbmdBroadcastDistributionTable")
		}
		for _curItem, _element := range m.GetBbmdBroadcastDistributionTable() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetBbmdBroadcastDistributionTable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'bbmdBroadcastDistributionTable' field")
			}
		}
		if popErr := writeBuffer.PopContext("bbmdBroadcastDistributionTable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bbmdBroadcastDistributionTable")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDBroadcastDistributionTable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) isBACnetConstructedDataBBMDBroadcastDistributionTable() bool {
	return true
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
