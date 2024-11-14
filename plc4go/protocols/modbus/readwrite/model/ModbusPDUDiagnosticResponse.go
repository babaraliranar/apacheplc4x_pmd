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

// ModbusPDUDiagnosticResponse is the corresponding interface of ModbusPDUDiagnosticResponse
type ModbusPDUDiagnosticResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
	// IsModbusPDUDiagnosticResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUDiagnosticResponse()
	// CreateBuilder creates a ModbusPDUDiagnosticResponseBuilder
	CreateModbusPDUDiagnosticResponseBuilder() ModbusPDUDiagnosticResponseBuilder
}

// _ModbusPDUDiagnosticResponse is the data-structure of this message
type _ModbusPDUDiagnosticResponse struct {
	ModbusPDUContract
	SubFunction uint16
	Data        uint16
}

var _ ModbusPDUDiagnosticResponse = (*_ModbusPDUDiagnosticResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUDiagnosticResponse)(nil)

// NewModbusPDUDiagnosticResponse factory function for _ModbusPDUDiagnosticResponse
func NewModbusPDUDiagnosticResponse(subFunction uint16, data uint16) *_ModbusPDUDiagnosticResponse {
	_result := &_ModbusPDUDiagnosticResponse{
		ModbusPDUContract: NewModbusPDU(),
		SubFunction:       subFunction,
		Data:              data,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUDiagnosticResponseBuilder is a builder for ModbusPDUDiagnosticResponse
type ModbusPDUDiagnosticResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subFunction uint16, data uint16) ModbusPDUDiagnosticResponseBuilder
	// WithSubFunction adds SubFunction (property field)
	WithSubFunction(uint16) ModbusPDUDiagnosticResponseBuilder
	// WithData adds Data (property field)
	WithData(uint16) ModbusPDUDiagnosticResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUDiagnosticResponse or returns an error if something is wrong
	Build() (ModbusPDUDiagnosticResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUDiagnosticResponse
}

// NewModbusPDUDiagnosticResponseBuilder() creates a ModbusPDUDiagnosticResponseBuilder
func NewModbusPDUDiagnosticResponseBuilder() ModbusPDUDiagnosticResponseBuilder {
	return &_ModbusPDUDiagnosticResponseBuilder{_ModbusPDUDiagnosticResponse: new(_ModbusPDUDiagnosticResponse)}
}

type _ModbusPDUDiagnosticResponseBuilder struct {
	*_ModbusPDUDiagnosticResponse

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUDiagnosticResponseBuilder) = (*_ModbusPDUDiagnosticResponseBuilder)(nil)

func (b *_ModbusPDUDiagnosticResponseBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUDiagnosticResponse
}

func (b *_ModbusPDUDiagnosticResponseBuilder) WithMandatoryFields(subFunction uint16, data uint16) ModbusPDUDiagnosticResponseBuilder {
	return b.WithSubFunction(subFunction).WithData(data)
}

func (b *_ModbusPDUDiagnosticResponseBuilder) WithSubFunction(subFunction uint16) ModbusPDUDiagnosticResponseBuilder {
	b.SubFunction = subFunction
	return b
}

func (b *_ModbusPDUDiagnosticResponseBuilder) WithData(data uint16) ModbusPDUDiagnosticResponseBuilder {
	b.Data = data
	return b
}

func (b *_ModbusPDUDiagnosticResponseBuilder) Build() (ModbusPDUDiagnosticResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUDiagnosticResponse.deepCopy(), nil
}

func (b *_ModbusPDUDiagnosticResponseBuilder) MustBuild() ModbusPDUDiagnosticResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUDiagnosticResponseBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUDiagnosticResponseBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUDiagnosticResponseBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUDiagnosticResponseBuilder().(*_ModbusPDUDiagnosticResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUDiagnosticResponseBuilder creates a ModbusPDUDiagnosticResponseBuilder
func (b *_ModbusPDUDiagnosticResponse) CreateModbusPDUDiagnosticResponseBuilder() ModbusPDUDiagnosticResponseBuilder {
	if b == nil {
		return NewModbusPDUDiagnosticResponseBuilder()
	}
	return &_ModbusPDUDiagnosticResponseBuilder{_ModbusPDUDiagnosticResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUDiagnosticResponse) GetFunctionFlag() uint8 {
	return 0x08
}

func (m *_ModbusPDUDiagnosticResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *_ModbusPDUDiagnosticResponse) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUDiagnosticResponse(structType any) ModbusPDUDiagnosticResponse {
	if casted, ok := structType.(ModbusPDUDiagnosticResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUDiagnosticResponse) GetTypeName() string {
	return "ModbusPDUDiagnosticResponse"
}

func (m *_ModbusPDUDiagnosticResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (subFunction)
	lengthInBits += 16

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUDiagnosticResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUDiagnosticResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUDiagnosticResponse ModbusPDUDiagnosticResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUDiagnosticResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subFunction, err := ReadSimpleField(ctx, "subFunction", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subFunction' field"))
	}
	m.SubFunction = subFunction

	data, err := ReadSimpleField(ctx, "data", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUDiagnosticResponse")
	}

	return m, nil
}

func (m *_ModbusPDUDiagnosticResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUDiagnosticResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUDiagnosticResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "subFunction", m.GetSubFunction(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'subFunction' field")
		}

		if err := WriteSimpleField[uint16](ctx, "data", m.GetData(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUDiagnosticResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUDiagnosticResponse) IsModbusPDUDiagnosticResponse() {}

func (m *_ModbusPDUDiagnosticResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUDiagnosticResponse) deepCopy() *_ModbusPDUDiagnosticResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUDiagnosticResponseCopy := &_ModbusPDUDiagnosticResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.SubFunction,
		m.Data,
	}
	_ModbusPDUDiagnosticResponseCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUDiagnosticResponseCopy
}

func (m *_ModbusPDUDiagnosticResponse) String() string {
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
