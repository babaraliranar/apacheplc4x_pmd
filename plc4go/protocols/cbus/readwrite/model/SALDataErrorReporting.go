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

// SALDataErrorReporting is the corresponding interface of SALDataErrorReporting
type SALDataErrorReporting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetErrorReportingData returns ErrorReportingData (property field)
	GetErrorReportingData() ErrorReportingData
	// IsSALDataErrorReporting is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataErrorReporting()
	// CreateBuilder creates a SALDataErrorReportingBuilder
	CreateSALDataErrorReportingBuilder() SALDataErrorReportingBuilder
}

// _SALDataErrorReporting is the data-structure of this message
type _SALDataErrorReporting struct {
	SALDataContract
	ErrorReportingData ErrorReportingData
}

var _ SALDataErrorReporting = (*_SALDataErrorReporting)(nil)
var _ SALDataRequirements = (*_SALDataErrorReporting)(nil)

// NewSALDataErrorReporting factory function for _SALDataErrorReporting
func NewSALDataErrorReporting(salData SALData, errorReportingData ErrorReportingData) *_SALDataErrorReporting {
	if errorReportingData == nil {
		panic("errorReportingData of type ErrorReportingData for SALDataErrorReporting must not be nil")
	}
	_result := &_SALDataErrorReporting{
		SALDataContract:    NewSALData(salData),
		ErrorReportingData: errorReportingData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataErrorReportingBuilder is a builder for SALDataErrorReporting
type SALDataErrorReportingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorReportingData ErrorReportingData) SALDataErrorReportingBuilder
	// WithErrorReportingData adds ErrorReportingData (property field)
	WithErrorReportingData(ErrorReportingData) SALDataErrorReportingBuilder
	// Build builds the SALDataErrorReporting or returns an error if something is wrong
	Build() (SALDataErrorReporting, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataErrorReporting
}

// NewSALDataErrorReportingBuilder() creates a SALDataErrorReportingBuilder
func NewSALDataErrorReportingBuilder() SALDataErrorReportingBuilder {
	return &_SALDataErrorReportingBuilder{_SALDataErrorReporting: new(_SALDataErrorReporting)}
}

type _SALDataErrorReportingBuilder struct {
	*_SALDataErrorReporting

	err *utils.MultiError
}

var _ (SALDataErrorReportingBuilder) = (*_SALDataErrorReportingBuilder)(nil)

func (m *_SALDataErrorReportingBuilder) WithMandatoryFields(errorReportingData ErrorReportingData) SALDataErrorReportingBuilder {
	return m.WithErrorReportingData(errorReportingData)
}

func (m *_SALDataErrorReportingBuilder) WithErrorReportingData(errorReportingData ErrorReportingData) SALDataErrorReportingBuilder {
	m.ErrorReportingData = errorReportingData
	return m
}

func (m *_SALDataErrorReportingBuilder) Build() (SALDataErrorReporting, error) {
	if m.ErrorReportingData == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'errorReportingData' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SALDataErrorReporting.deepCopy(), nil
}

func (m *_SALDataErrorReportingBuilder) MustBuild() SALDataErrorReporting {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SALDataErrorReportingBuilder) DeepCopy() any {
	return m.CreateSALDataErrorReportingBuilder()
}

// CreateSALDataErrorReportingBuilder creates a SALDataErrorReportingBuilder
func (m *_SALDataErrorReporting) CreateSALDataErrorReportingBuilder() SALDataErrorReportingBuilder {
	if m == nil {
		return NewSALDataErrorReportingBuilder()
	}
	return &_SALDataErrorReportingBuilder{_SALDataErrorReporting: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataErrorReporting) GetApplicationId() ApplicationId {
	return ApplicationId_ERROR_REPORTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataErrorReporting) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataErrorReporting) GetErrorReportingData() ErrorReportingData {
	return m.ErrorReportingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataErrorReporting(structType any) SALDataErrorReporting {
	if casted, ok := structType.(SALDataErrorReporting); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataErrorReporting); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataErrorReporting) GetTypeName() string {
	return "SALDataErrorReporting"
}

func (m *_SALDataErrorReporting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (errorReportingData)
	lengthInBits += m.ErrorReportingData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataErrorReporting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataErrorReporting) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataErrorReporting SALDataErrorReporting, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataErrorReporting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataErrorReporting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorReportingData, err := ReadSimpleField[ErrorReportingData](ctx, "errorReportingData", ReadComplex[ErrorReportingData](ErrorReportingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorReportingData' field"))
	}
	m.ErrorReportingData = errorReportingData

	if closeErr := readBuffer.CloseContext("SALDataErrorReporting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataErrorReporting")
	}

	return m, nil
}

func (m *_SALDataErrorReporting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataErrorReporting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataErrorReporting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataErrorReporting")
		}

		if err := WriteSimpleField[ErrorReportingData](ctx, "errorReportingData", m.GetErrorReportingData(), WriteComplex[ErrorReportingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorReportingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataErrorReporting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataErrorReporting")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataErrorReporting) IsSALDataErrorReporting() {}

func (m *_SALDataErrorReporting) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataErrorReporting) deepCopy() *_SALDataErrorReporting {
	if m == nil {
		return nil
	}
	_SALDataErrorReportingCopy := &_SALDataErrorReporting{
		m.SALDataContract.(*_SALData).deepCopy(),
		m.ErrorReportingData.DeepCopy().(ErrorReportingData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataErrorReportingCopy
}

func (m *_SALDataErrorReporting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
