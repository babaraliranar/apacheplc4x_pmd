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

// BACnetLogDataLogStatus is the corresponding interface of BACnetLogDataLogStatus
type BACnetLogDataLogStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogData
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
	// IsBACnetLogDataLogStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogStatus()
	// CreateBuilder creates a BACnetLogDataLogStatusBuilder
	CreateBACnetLogDataLogStatusBuilder() BACnetLogDataLogStatusBuilder
}

// _BACnetLogDataLogStatus is the data-structure of this message
type _BACnetLogDataLogStatus struct {
	BACnetLogDataContract
	LogStatus BACnetLogStatusTagged
}

var _ BACnetLogDataLogStatus = (*_BACnetLogDataLogStatus)(nil)
var _ BACnetLogDataRequirements = (*_BACnetLogDataLogStatus)(nil)

// NewBACnetLogDataLogStatus factory function for _BACnetLogDataLogStatus
func NewBACnetLogDataLogStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, logStatus BACnetLogStatusTagged, tagNumber uint8) *_BACnetLogDataLogStatus {
	if logStatus == nil {
		panic("logStatus of type BACnetLogStatusTagged for BACnetLogDataLogStatus must not be nil")
	}
	_result := &_BACnetLogDataLogStatus{
		BACnetLogDataContract: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
		LogStatus:             logStatus,
	}
	_result.BACnetLogDataContract.(*_BACnetLogData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogDataLogStatusBuilder is a builder for BACnetLogDataLogStatus
type BACnetLogDataLogStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(logStatus BACnetLogStatusTagged) BACnetLogDataLogStatusBuilder
	// WithLogStatus adds LogStatus (property field)
	WithLogStatus(BACnetLogStatusTagged) BACnetLogDataLogStatusBuilder
	// WithLogStatusBuilder adds LogStatus (property field) which is build by the builder
	WithLogStatusBuilder(func(BACnetLogStatusTaggedBuilder) BACnetLogStatusTaggedBuilder) BACnetLogDataLogStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogDataBuilder
	// Build builds the BACnetLogDataLogStatus or returns an error if something is wrong
	Build() (BACnetLogDataLogStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogDataLogStatus
}

// NewBACnetLogDataLogStatusBuilder() creates a BACnetLogDataLogStatusBuilder
func NewBACnetLogDataLogStatusBuilder() BACnetLogDataLogStatusBuilder {
	return &_BACnetLogDataLogStatusBuilder{_BACnetLogDataLogStatus: new(_BACnetLogDataLogStatus)}
}

type _BACnetLogDataLogStatusBuilder struct {
	*_BACnetLogDataLogStatus

	parentBuilder *_BACnetLogDataBuilder

	err *utils.MultiError
}

var _ (BACnetLogDataLogStatusBuilder) = (*_BACnetLogDataLogStatusBuilder)(nil)

func (b *_BACnetLogDataLogStatusBuilder) setParent(contract BACnetLogDataContract) {
	b.BACnetLogDataContract = contract
	contract.(*_BACnetLogData)._SubType = b._BACnetLogDataLogStatus
}

func (b *_BACnetLogDataLogStatusBuilder) WithMandatoryFields(logStatus BACnetLogStatusTagged) BACnetLogDataLogStatusBuilder {
	return b.WithLogStatus(logStatus)
}

func (b *_BACnetLogDataLogStatusBuilder) WithLogStatus(logStatus BACnetLogStatusTagged) BACnetLogDataLogStatusBuilder {
	b.LogStatus = logStatus
	return b
}

func (b *_BACnetLogDataLogStatusBuilder) WithLogStatusBuilder(builderSupplier func(BACnetLogStatusTaggedBuilder) BACnetLogStatusTaggedBuilder) BACnetLogDataLogStatusBuilder {
	builder := builderSupplier(b.LogStatus.CreateBACnetLogStatusTaggedBuilder())
	var err error
	b.LogStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLogStatusTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogStatusBuilder) Build() (BACnetLogDataLogStatus, error) {
	if b.LogStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'logStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogDataLogStatus.deepCopy(), nil
}

func (b *_BACnetLogDataLogStatusBuilder) MustBuild() BACnetLogDataLogStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogDataLogStatusBuilder) Done() BACnetLogDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogDataBuilder().(*_BACnetLogDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogDataLogStatusBuilder) buildForBACnetLogData() (BACnetLogData, error) {
	return b.Build()
}

func (b *_BACnetLogDataLogStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogDataLogStatusBuilder().(*_BACnetLogDataLogStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogDataLogStatusBuilder creates a BACnetLogDataLogStatusBuilder
func (b *_BACnetLogDataLogStatus) CreateBACnetLogDataLogStatusBuilder() BACnetLogDataLogStatusBuilder {
	if b == nil {
		return NewBACnetLogDataLogStatusBuilder()
	}
	return &_BACnetLogDataLogStatusBuilder{_BACnetLogDataLogStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogStatus) GetParent() BACnetLogDataContract {
	return m.BACnetLogDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogStatus(structType any) BACnetLogDataLogStatus {
	if casted, ok := structType.(BACnetLogDataLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogStatus) GetTypeName() string {
	return "BACnetLogDataLogStatus"
}

func (m *_BACnetLogDataLogStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataContract.(*_BACnetLogData).getLengthInBits(ctx))

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogData, tagNumber uint8) (__bACnetLogDataLogStatus BACnetLogDataLogStatus, err error) {
	m.BACnetLogDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logStatus, err := ReadSimpleField[BACnetLogStatusTagged](ctx, "logStatus", ReadComplex[BACnetLogStatusTagged](BACnetLogStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logStatus' field"))
	}
	m.LogStatus = logStatus

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogStatus")
	}

	return m, nil
}

func (m *_BACnetLogDataLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogStatus")
		}

		if err := WriteSimpleField[BACnetLogStatusTagged](ctx, "logStatus", m.GetLogStatus(), WriteComplex[BACnetLogStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogStatus")
		}
		return nil
	}
	return m.BACnetLogDataContract.(*_BACnetLogData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogStatus) IsBACnetLogDataLogStatus() {}

func (m *_BACnetLogDataLogStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogStatus) deepCopy() *_BACnetLogDataLogStatus {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogStatusCopy := &_BACnetLogDataLogStatus{
		m.BACnetLogDataContract.(*_BACnetLogData).deepCopy(),
		utils.DeepCopy[BACnetLogStatusTagged](m.LogStatus),
	}
	_BACnetLogDataLogStatusCopy.BACnetLogDataContract.(*_BACnetLogData)._SubType = m
	return _BACnetLogDataLogStatusCopy
}

func (m *_BACnetLogDataLogStatus) String() string {
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
