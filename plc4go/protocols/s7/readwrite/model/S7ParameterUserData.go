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

// S7ParameterUserData is the corresponding interface of S7ParameterUserData
type S7ParameterUserData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7ParameterUserDataItem
	// IsS7ParameterUserData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterUserData()
	// CreateBuilder creates a S7ParameterUserDataBuilder
	CreateS7ParameterUserDataBuilder() S7ParameterUserDataBuilder
}

// _S7ParameterUserData is the data-structure of this message
type _S7ParameterUserData struct {
	S7ParameterContract
	Items []S7ParameterUserDataItem
}

var _ S7ParameterUserData = (*_S7ParameterUserData)(nil)
var _ S7ParameterRequirements = (*_S7ParameterUserData)(nil)

// NewS7ParameterUserData factory function for _S7ParameterUserData
func NewS7ParameterUserData(items []S7ParameterUserDataItem) *_S7ParameterUserData {
	_result := &_S7ParameterUserData{
		S7ParameterContract: NewS7Parameter(),
		Items:               items,
	}
	_result.S7ParameterContract.(*_S7Parameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterUserDataBuilder is a builder for S7ParameterUserData
type S7ParameterUserDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []S7ParameterUserDataItem) S7ParameterUserDataBuilder
	// WithItems adds Items (property field)
	WithItems(...S7ParameterUserDataItem) S7ParameterUserDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7ParameterBuilder
	// Build builds the S7ParameterUserData or returns an error if something is wrong
	Build() (S7ParameterUserData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterUserData
}

// NewS7ParameterUserDataBuilder() creates a S7ParameterUserDataBuilder
func NewS7ParameterUserDataBuilder() S7ParameterUserDataBuilder {
	return &_S7ParameterUserDataBuilder{_S7ParameterUserData: new(_S7ParameterUserData)}
}

type _S7ParameterUserDataBuilder struct {
	*_S7ParameterUserData

	parentBuilder *_S7ParameterBuilder

	err *utils.MultiError
}

var _ (S7ParameterUserDataBuilder) = (*_S7ParameterUserDataBuilder)(nil)

func (b *_S7ParameterUserDataBuilder) setParent(contract S7ParameterContract) {
	b.S7ParameterContract = contract
	contract.(*_S7Parameter)._SubType = b._S7ParameterUserData
}

func (b *_S7ParameterUserDataBuilder) WithMandatoryFields(items []S7ParameterUserDataItem) S7ParameterUserDataBuilder {
	return b.WithItems(items...)
}

func (b *_S7ParameterUserDataBuilder) WithItems(items ...S7ParameterUserDataItem) S7ParameterUserDataBuilder {
	b.Items = items
	return b
}

func (b *_S7ParameterUserDataBuilder) Build() (S7ParameterUserData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7ParameterUserData.deepCopy(), nil
}

func (b *_S7ParameterUserDataBuilder) MustBuild() S7ParameterUserData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterUserDataBuilder) Done() S7ParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7ParameterBuilder().(*_S7ParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_S7ParameterUserDataBuilder) buildForS7Parameter() (S7Parameter, error) {
	return b.Build()
}

func (b *_S7ParameterUserDataBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterUserDataBuilder().(*_S7ParameterUserDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterUserDataBuilder creates a S7ParameterUserDataBuilder
func (b *_S7ParameterUserData) CreateS7ParameterUserDataBuilder() S7ParameterUserDataBuilder {
	if b == nil {
		return NewS7ParameterUserDataBuilder()
	}
	return &_S7ParameterUserDataBuilder{_S7ParameterUserData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterUserData) GetParameterType() uint8 {
	return 0x00
}

func (m *_S7ParameterUserData) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterUserData) GetParent() S7ParameterContract {
	return m.S7ParameterContract
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

// Deprecated: use the interface for direct cast
func CastS7ParameterUserData(structType any) S7ParameterUserData {
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
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7ParameterUserData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterUserData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterUserData S7ParameterUserData, err error) {
	m.S7ParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterUserData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numItems, err := ReadImplicitField[uint8](ctx, "numItems", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numItems' field"))
	}
	_ = numItems

	items, err := ReadCountArrayField[S7ParameterUserDataItem](ctx, "items", ReadComplex[S7ParameterUserDataItem](S7ParameterUserDataItemParseWithBuffer, readBuffer), uint64(numItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7ParameterUserData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserData")
	}

	return m, nil
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
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterUserData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterUserData")
		}
		numItems := uint8(uint8(len(m.GetItems())))
		if err := WriteImplicitField(ctx, "numItems", numItems, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numItems' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterUserData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterUserData")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterUserData) IsS7ParameterUserData() {}

func (m *_S7ParameterUserData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterUserData) deepCopy() *_S7ParameterUserData {
	if m == nil {
		return nil
	}
	_S7ParameterUserDataCopy := &_S7ParameterUserData{
		m.S7ParameterContract.(*_S7Parameter).deepCopy(),
		utils.DeepCopySlice[S7ParameterUserDataItem, S7ParameterUserDataItem](m.Items),
	}
	_S7ParameterUserDataCopy.S7ParameterContract.(*_S7Parameter)._SubType = m
	return _S7ParameterUserDataCopy
}

func (m *_S7ParameterUserData) String() string {
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
