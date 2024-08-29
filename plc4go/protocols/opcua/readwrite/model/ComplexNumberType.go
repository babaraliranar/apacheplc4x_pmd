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

// ComplexNumberType is the corresponding interface of ComplexNumberType
type ComplexNumberType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetReal returns Real (property field)
	GetReal() float32
	// GetImaginary returns Imaginary (property field)
	GetImaginary() float32
}

// ComplexNumberTypeExactly can be used when we want exactly this type and not a type which fulfills ComplexNumberType.
// This is useful for switch cases.
type ComplexNumberTypeExactly interface {
	ComplexNumberType
	isComplexNumberType() bool
}

// _ComplexNumberType is the data-structure of this message
type _ComplexNumberType struct {
	*_ExtensionObjectDefinition
	Real      float32
	Imaginary float32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComplexNumberType) GetIdentifier() string {
	return "12173"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComplexNumberType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ComplexNumberType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComplexNumberType) GetReal() float32 {
	return m.Real
}

func (m *_ComplexNumberType) GetImaginary() float32 {
	return m.Imaginary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewComplexNumberType factory function for _ComplexNumberType
func NewComplexNumberType(real float32, imaginary float32) *_ComplexNumberType {
	_result := &_ComplexNumberType{
		Real:                       real,
		Imaginary:                  imaginary,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastComplexNumberType(structType any) ComplexNumberType {
	if casted, ok := structType.(ComplexNumberType); ok {
		return casted
	}
	if casted, ok := structType.(*ComplexNumberType); ok {
		return *casted
	}
	return nil
}

func (m *_ComplexNumberType) GetTypeName() string {
	return "ComplexNumberType"
}

func (m *_ComplexNumberType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (real)
	lengthInBits += 32

	// Simple field (imaginary)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ComplexNumberType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ComplexNumberTypeParse(ctx context.Context, theBytes []byte, identifier string) (ComplexNumberType, error) {
	return ComplexNumberTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ComplexNumberTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ComplexNumberType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ComplexNumberType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComplexNumberType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (real)
	_real, _realErr := readBuffer.ReadFloat32("real", 32)
	if _realErr != nil {
		return nil, errors.Wrap(_realErr, "Error parsing 'real' field of ComplexNumberType")
	}
	real := _real

	// Simple Field (imaginary)
	_imaginary, _imaginaryErr := readBuffer.ReadFloat32("imaginary", 32)
	if _imaginaryErr != nil {
		return nil, errors.Wrap(_imaginaryErr, "Error parsing 'imaginary' field of ComplexNumberType")
	}
	imaginary := _imaginary

	if closeErr := readBuffer.CloseContext("ComplexNumberType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComplexNumberType")
	}

	// Create a partially initialized instance
	_child := &_ComplexNumberType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Real:                       real,
		Imaginary:                  imaginary,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ComplexNumberType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComplexNumberType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComplexNumberType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComplexNumberType")
		}

		// Simple Field (real)
		real := float32(m.GetReal())
		_realErr := writeBuffer.WriteFloat32("real", 32, (real))
		if _realErr != nil {
			return errors.Wrap(_realErr, "Error serializing 'real' field")
		}

		// Simple Field (imaginary)
		imaginary := float32(m.GetImaginary())
		_imaginaryErr := writeBuffer.WriteFloat32("imaginary", 32, (imaginary))
		if _imaginaryErr != nil {
			return errors.Wrap(_imaginaryErr, "Error serializing 'imaginary' field")
		}

		if popErr := writeBuffer.PopContext("ComplexNumberType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComplexNumberType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComplexNumberType) isComplexNumberType() bool {
	return true
}

func (m *_ComplexNumberType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
