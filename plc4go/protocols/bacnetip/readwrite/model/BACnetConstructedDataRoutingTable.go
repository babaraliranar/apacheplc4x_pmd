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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataRoutingTable is the corresponding interface of BACnetConstructedDataRoutingTable
type BACnetConstructedDataRoutingTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRoutingTable returns RoutingTable (property field)
	GetRoutingTable() []BACnetRouterEntry
}

// BACnetConstructedDataRoutingTableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRoutingTable.
// This is useful for switch cases.
type BACnetConstructedDataRoutingTableExactly interface {
	BACnetConstructedDataRoutingTable
	isBACnetConstructedDataRoutingTable() bool
}

// _BACnetConstructedDataRoutingTable is the data-structure of this message
type _BACnetConstructedDataRoutingTable struct {
	*_BACnetConstructedData
	RoutingTable []BACnetRouterEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRoutingTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRoutingTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ROUTING_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRoutingTable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRoutingTable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRoutingTable) GetRoutingTable() []BACnetRouterEntry {
	return m.RoutingTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRoutingTable factory function for _BACnetConstructedDataRoutingTable
func NewBACnetConstructedDataRoutingTable(routingTable []BACnetRouterEntry, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRoutingTable {
	_result := &_BACnetConstructedDataRoutingTable{
		RoutingTable:           routingTable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRoutingTable(structType any) BACnetConstructedDataRoutingTable {
	if casted, ok := structType.(BACnetConstructedDataRoutingTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRoutingTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRoutingTable) GetTypeName() string {
	return "BACnetConstructedDataRoutingTable"
}

func (m *_BACnetConstructedDataRoutingTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.RoutingTable) > 0 {
		for _, element := range m.RoutingTable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataRoutingTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataRoutingTableParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRoutingTable, error) {
	return BACnetConstructedDataRoutingTableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataRoutingTableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRoutingTable, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRoutingTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRoutingTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (routingTable)
	if pullErr := readBuffer.PullContext("routingTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for routingTable")
	}
	// Terminated array
	var routingTable []BACnetRouterEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetRouterEntryParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'routingTable' field of BACnetConstructedDataRoutingTable")
			}
			routingTable = append(routingTable, _item.(BACnetRouterEntry))
		}
	}
	if closeErr := readBuffer.CloseContext("routingTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for routingTable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRoutingTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRoutingTable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRoutingTable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RoutingTable: routingTable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRoutingTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRoutingTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRoutingTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRoutingTable")
		}

		// Array Field (routingTable)
		if pushErr := writeBuffer.PushContext("routingTable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for routingTable")
		}
		for _curItem, _element := range m.GetRoutingTable() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetRoutingTable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'routingTable' field")
			}
		}
		if popErr := writeBuffer.PopContext("routingTable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for routingTable")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRoutingTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRoutingTable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRoutingTable) isBACnetConstructedDataRoutingTable() bool {
	return true
}

func (m *_BACnetConstructedDataRoutingTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
