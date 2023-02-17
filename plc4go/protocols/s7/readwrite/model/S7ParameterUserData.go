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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
)

	// Code generated by code-generation. DO NOT EDIT.


// S7ParameterUserData is the corresponding interface of S7ParameterUserData
type S7ParameterUserData interface {
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7ParameterUserDataItem
}

// S7ParameterUserDataExactly can be used when we want exactly this type and not a type which fulfills S7ParameterUserData.
// This is useful for switch cases.
type S7ParameterUserDataExactly interface {
	S7ParameterUserData
	isS7ParameterUserData() bool
}

// _S7ParameterUserData is the data-structure of this message
type _S7ParameterUserData struct {
	*_S7Parameter
        Items []S7ParameterUserDataItem
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterUserData)  GetParameterType() uint8 {
return 0x00}

func (m *_S7ParameterUserData)  GetMessageType() uint8 {
return 0x07}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterUserData) InitializeParent(parent S7Parameter ) {}

func (m *_S7ParameterUserData)  GetParent() S7Parameter {
	return m._S7Parameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterUserData) GetItems() []S7ParameterUserDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewS7ParameterUserData factory function for _S7ParameterUserData
func NewS7ParameterUserData( items []S7ParameterUserDataItem ) *_S7ParameterUserData {
	_result := &_S7ParameterUserData{
		Items: items,
    	_S7Parameter: NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterUserData(structType interface{}) S7ParameterUserData {
    if casted, ok := structType.(S7ParameterUserData); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterUserData); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterUserData) GetTypeName() string {
	return "S7ParameterUserData"
}

func (m *_S7ParameterUserData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_S7ParameterUserData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7ParameterUserDataParse(theBytes []byte, messageType uint8) (S7ParameterUserData, error) {
	return S7ParameterUserDataParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), messageType)
}

func S7ParameterUserDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterUserData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterUserData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (numItems) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numItems, _numItemsErr := readBuffer.ReadUint8("numItems", 8)
	_ = numItems
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field of S7ParameterUserData")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]S7ParameterUserDataItem, numItems)
	// This happens when the size is set conditional to 0
	if len(items) == 0 {
		items = nil
	}
	{
		_numItems := uint16(numItems)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := S7ParameterUserDataItemParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field of S7ParameterUserData")
			}
			items[_curItem] = _item.(S7ParameterUserDataItem)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7ParameterUserData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserData")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterUserData{
		_S7Parameter: &_S7Parameter{
		},
		Items: items,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterUserData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterUserData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterUserData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterUserData")
		}

	// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numItems := uint8(uint8(len(m.GetItems())))
	_numItemsErr := writeBuffer.WriteUint8("numItems", 8, (numItems))
	if _numItemsErr != nil {
		return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
	}

	// Array Field (items)
	if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for items")
	}
	for _curItem, _element := range m.GetItems() {
		_ = _curItem
		arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'items' field")
		}
	}
	if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for items")
	}

		if popErr := writeBuffer.PopContext("S7ParameterUserData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterUserData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_S7ParameterUserData) isS7ParameterUserData() bool {
	return true
}

func (m *_S7ParameterUserData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



