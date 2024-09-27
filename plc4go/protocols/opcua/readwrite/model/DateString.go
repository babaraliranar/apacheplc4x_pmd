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

// DateString is the corresponding interface of DateString
type DateString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsDateString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDateString()
	// CreateBuilder creates a DateStringBuilder
	CreateDateStringBuilder() DateStringBuilder
}

// _DateString is the data-structure of this message
type _DateString struct {
}

var _ DateString = (*_DateString)(nil)

// NewDateString factory function for _DateString
func NewDateString() *_DateString {
	return &_DateString{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DateStringBuilder is a builder for DateString
type DateStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DateStringBuilder
	// Build builds the DateString or returns an error if something is wrong
	Build() (DateString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DateString
}

// NewDateStringBuilder() creates a DateStringBuilder
func NewDateStringBuilder() DateStringBuilder {
	return &_DateStringBuilder{_DateString: new(_DateString)}
}

type _DateStringBuilder struct {
	*_DateString

	err *utils.MultiError
}

var _ (DateStringBuilder) = (*_DateStringBuilder)(nil)

func (m *_DateStringBuilder) WithMandatoryFields() DateStringBuilder {
	return m
}

func (m *_DateStringBuilder) Build() (DateString, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._DateString.deepCopy(), nil
}

func (m *_DateStringBuilder) MustBuild() DateString {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_DateStringBuilder) DeepCopy() any {
	return m.CreateDateStringBuilder()
}

// CreateDateStringBuilder creates a DateStringBuilder
func (m *_DateString) CreateDateStringBuilder() DateStringBuilder {
	if m == nil {
		return NewDateStringBuilder()
	}
	return &_DateStringBuilder{_DateString: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDateString(structType any) DateString {
	if casted, ok := structType.(DateString); ok {
		return casted
	}
	if casted, ok := structType.(*DateString); ok {
		return *casted
	}
	return nil
}

func (m *_DateString) GetTypeName() string {
	return "DateString"
}

func (m *_DateString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_DateString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DateStringParse(ctx context.Context, theBytes []byte) (DateString, error) {
	return DateStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DateStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DateString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DateString, error) {
		return DateStringParseWithBuffer(ctx, readBuffer)
	}
}

func DateStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DateString, error) {
	v, err := (&_DateString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DateString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dateString DateString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DateString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DateString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DateString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DateString")
	}

	return m, nil
}

func (m *_DateString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DateString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DateString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DateString")
	}

	if popErr := writeBuffer.PopContext("DateString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DateString")
	}
	return nil
}

func (m *_DateString) IsDateString() {}

func (m *_DateString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DateString) deepCopy() *_DateString {
	if m == nil {
		return nil
	}
	_DateStringCopy := &_DateString{}
	return _DateStringCopy
}

func (m *_DateString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
