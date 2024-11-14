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

// S7VarRequestParameterItemAddress is the corresponding interface of S7VarRequestParameterItemAddress
type S7VarRequestParameterItemAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7VarRequestParameterItem
	// GetAddress returns Address (property field)
	GetAddress() S7Address
	// IsS7VarRequestParameterItemAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7VarRequestParameterItemAddress()
	// CreateBuilder creates a S7VarRequestParameterItemAddressBuilder
	CreateS7VarRequestParameterItemAddressBuilder() S7VarRequestParameterItemAddressBuilder
}

// _S7VarRequestParameterItemAddress is the data-structure of this message
type _S7VarRequestParameterItemAddress struct {
	S7VarRequestParameterItemContract
	Address S7Address
}

var _ S7VarRequestParameterItemAddress = (*_S7VarRequestParameterItemAddress)(nil)
var _ S7VarRequestParameterItemRequirements = (*_S7VarRequestParameterItemAddress)(nil)

// NewS7VarRequestParameterItemAddress factory function for _S7VarRequestParameterItemAddress
func NewS7VarRequestParameterItemAddress(address S7Address) *_S7VarRequestParameterItemAddress {
	if address == nil {
		panic("address of type S7Address for S7VarRequestParameterItemAddress must not be nil")
	}
	_result := &_S7VarRequestParameterItemAddress{
		S7VarRequestParameterItemContract: NewS7VarRequestParameterItem(),
		Address:                           address,
	}
	_result.S7VarRequestParameterItemContract.(*_S7VarRequestParameterItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7VarRequestParameterItemAddressBuilder is a builder for S7VarRequestParameterItemAddress
type S7VarRequestParameterItemAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address S7Address) S7VarRequestParameterItemAddressBuilder
	// WithAddress adds Address (property field)
	WithAddress(S7Address) S7VarRequestParameterItemAddressBuilder
	// WithAddressBuilder adds Address (property field) which is build by the builder
	WithAddressBuilder(func(S7AddressBuilder) S7AddressBuilder) S7VarRequestParameterItemAddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7VarRequestParameterItemBuilder
	// Build builds the S7VarRequestParameterItemAddress or returns an error if something is wrong
	Build() (S7VarRequestParameterItemAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7VarRequestParameterItemAddress
}

// NewS7VarRequestParameterItemAddressBuilder() creates a S7VarRequestParameterItemAddressBuilder
func NewS7VarRequestParameterItemAddressBuilder() S7VarRequestParameterItemAddressBuilder {
	return &_S7VarRequestParameterItemAddressBuilder{_S7VarRequestParameterItemAddress: new(_S7VarRequestParameterItemAddress)}
}

type _S7VarRequestParameterItemAddressBuilder struct {
	*_S7VarRequestParameterItemAddress

	parentBuilder *_S7VarRequestParameterItemBuilder

	err *utils.MultiError
}

var _ (S7VarRequestParameterItemAddressBuilder) = (*_S7VarRequestParameterItemAddressBuilder)(nil)

func (b *_S7VarRequestParameterItemAddressBuilder) setParent(contract S7VarRequestParameterItemContract) {
	b.S7VarRequestParameterItemContract = contract
	contract.(*_S7VarRequestParameterItem)._SubType = b._S7VarRequestParameterItemAddress
}

func (b *_S7VarRequestParameterItemAddressBuilder) WithMandatoryFields(address S7Address) S7VarRequestParameterItemAddressBuilder {
	return b.WithAddress(address)
}

func (b *_S7VarRequestParameterItemAddressBuilder) WithAddress(address S7Address) S7VarRequestParameterItemAddressBuilder {
	b.Address = address
	return b
}

func (b *_S7VarRequestParameterItemAddressBuilder) WithAddressBuilder(builderSupplier func(S7AddressBuilder) S7AddressBuilder) S7VarRequestParameterItemAddressBuilder {
	builder := builderSupplier(b.Address.CreateS7AddressBuilder())
	var err error
	b.Address, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "S7AddressBuilder failed"))
	}
	return b
}

func (b *_S7VarRequestParameterItemAddressBuilder) Build() (S7VarRequestParameterItemAddress, error) {
	if b.Address == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'address' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7VarRequestParameterItemAddress.deepCopy(), nil
}

func (b *_S7VarRequestParameterItemAddressBuilder) MustBuild() S7VarRequestParameterItemAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7VarRequestParameterItemAddressBuilder) Done() S7VarRequestParameterItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7VarRequestParameterItemBuilder().(*_S7VarRequestParameterItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7VarRequestParameterItemAddressBuilder) buildForS7VarRequestParameterItem() (S7VarRequestParameterItem, error) {
	return b.Build()
}

func (b *_S7VarRequestParameterItemAddressBuilder) DeepCopy() any {
	_copy := b.CreateS7VarRequestParameterItemAddressBuilder().(*_S7VarRequestParameterItemAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7VarRequestParameterItemAddressBuilder creates a S7VarRequestParameterItemAddressBuilder
func (b *_S7VarRequestParameterItemAddress) CreateS7VarRequestParameterItemAddressBuilder() S7VarRequestParameterItemAddressBuilder {
	if b == nil {
		return NewS7VarRequestParameterItemAddressBuilder()
	}
	return &_S7VarRequestParameterItemAddressBuilder{_S7VarRequestParameterItemAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7VarRequestParameterItemAddress) GetItemType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7VarRequestParameterItemAddress) GetParent() S7VarRequestParameterItemContract {
	return m.S7VarRequestParameterItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7VarRequestParameterItemAddress) GetAddress() S7Address {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7VarRequestParameterItemAddress(structType any) S7VarRequestParameterItemAddress {
	if casted, ok := structType.(S7VarRequestParameterItemAddress); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarRequestParameterItemAddress); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarRequestParameterItemAddress) GetTypeName() string {
	return "S7VarRequestParameterItemAddress"
}

func (m *_S7VarRequestParameterItemAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7VarRequestParameterItemContract.(*_S7VarRequestParameterItem).getLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7VarRequestParameterItemAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7VarRequestParameterItemAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7VarRequestParameterItem) (__s7VarRequestParameterItemAddress S7VarRequestParameterItemAddress, err error) {
	m.S7VarRequestParameterItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItemAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarRequestParameterItemAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemLength, err := ReadImplicitField[uint8](ctx, "itemLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	address, err := ReadSimpleField[S7Address](ctx, "address", ReadComplex[S7Address](S7AddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItemAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarRequestParameterItemAddress")
	}

	return m, nil
}

func (m *_S7VarRequestParameterItemAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7VarRequestParameterItemAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7VarRequestParameterItemAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7VarRequestParameterItemAddress")
		}
		itemLength := uint8(m.GetAddress().GetLengthInBytes(ctx))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		if err := WriteSimpleField[S7Address](ctx, "address", m.GetAddress(), WriteComplex[S7Address](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("S7VarRequestParameterItemAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7VarRequestParameterItemAddress")
		}
		return nil
	}
	return m.S7VarRequestParameterItemContract.(*_S7VarRequestParameterItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7VarRequestParameterItemAddress) IsS7VarRequestParameterItemAddress() {}

func (m *_S7VarRequestParameterItemAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7VarRequestParameterItemAddress) deepCopy() *_S7VarRequestParameterItemAddress {
	if m == nil {
		return nil
	}
	_S7VarRequestParameterItemAddressCopy := &_S7VarRequestParameterItemAddress{
		m.S7VarRequestParameterItemContract.(*_S7VarRequestParameterItem).deepCopy(),
		utils.DeepCopy[S7Address](m.Address),
	}
	_S7VarRequestParameterItemAddressCopy.S7VarRequestParameterItemContract.(*_S7VarRequestParameterItem)._SubType = m
	return _S7VarRequestParameterItemAddressCopy
}

func (m *_S7VarRequestParameterItemAddress) String() string {
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
