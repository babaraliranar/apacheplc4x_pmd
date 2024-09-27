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

// BACnetConstructedDataArchive is the corresponding interface of BACnetConstructedDataArchive
type BACnetConstructedDataArchive interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetArchive returns Archive (property field)
	GetArchive() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataArchive is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataArchive()
	// CreateBuilder creates a BACnetConstructedDataArchiveBuilder
	CreateBACnetConstructedDataArchiveBuilder() BACnetConstructedDataArchiveBuilder
}

// _BACnetConstructedDataArchive is the data-structure of this message
type _BACnetConstructedDataArchive struct {
	BACnetConstructedDataContract
	Archive BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataArchive = (*_BACnetConstructedDataArchive)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataArchive)(nil)

// NewBACnetConstructedDataArchive factory function for _BACnetConstructedDataArchive
func NewBACnetConstructedDataArchive(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, archive BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataArchive {
	if archive == nil {
		panic("archive of type BACnetApplicationTagBoolean for BACnetConstructedDataArchive must not be nil")
	}
	_result := &_BACnetConstructedDataArchive{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Archive:                       archive,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataArchiveBuilder is a builder for BACnetConstructedDataArchive
type BACnetConstructedDataArchiveBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(archive BACnetApplicationTagBoolean) BACnetConstructedDataArchiveBuilder
	// WithArchive adds Archive (property field)
	WithArchive(BACnetApplicationTagBoolean) BACnetConstructedDataArchiveBuilder
	// WithArchiveBuilder adds Archive (property field) which is build by the builder
	WithArchiveBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataArchiveBuilder
	// Build builds the BACnetConstructedDataArchive or returns an error if something is wrong
	Build() (BACnetConstructedDataArchive, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataArchive
}

// NewBACnetConstructedDataArchiveBuilder() creates a BACnetConstructedDataArchiveBuilder
func NewBACnetConstructedDataArchiveBuilder() BACnetConstructedDataArchiveBuilder {
	return &_BACnetConstructedDataArchiveBuilder{_BACnetConstructedDataArchive: new(_BACnetConstructedDataArchive)}
}

type _BACnetConstructedDataArchiveBuilder struct {
	*_BACnetConstructedDataArchive

	err *utils.MultiError
}

var _ (BACnetConstructedDataArchiveBuilder) = (*_BACnetConstructedDataArchiveBuilder)(nil)

func (m *_BACnetConstructedDataArchiveBuilder) WithMandatoryFields(archive BACnetApplicationTagBoolean) BACnetConstructedDataArchiveBuilder {
	return m.WithArchive(archive)
}

func (m *_BACnetConstructedDataArchiveBuilder) WithArchive(archive BACnetApplicationTagBoolean) BACnetConstructedDataArchiveBuilder {
	m.Archive = archive
	return m
}

func (m *_BACnetConstructedDataArchiveBuilder) WithArchiveBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataArchiveBuilder {
	builder := builderSupplier(m.Archive.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.Archive, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataArchiveBuilder) Build() (BACnetConstructedDataArchive, error) {
	if m.Archive == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'archive' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataArchive.deepCopy(), nil
}

func (m *_BACnetConstructedDataArchiveBuilder) MustBuild() BACnetConstructedDataArchive {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataArchiveBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataArchiveBuilder()
}

// CreateBACnetConstructedDataArchiveBuilder creates a BACnetConstructedDataArchiveBuilder
func (m *_BACnetConstructedDataArchive) CreateBACnetConstructedDataArchiveBuilder() BACnetConstructedDataArchiveBuilder {
	if m == nil {
		return NewBACnetConstructedDataArchiveBuilder()
	}
	return &_BACnetConstructedDataArchiveBuilder{_BACnetConstructedDataArchive: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataArchive) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataArchive) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ARCHIVE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataArchive) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataArchive) GetArchive() BACnetApplicationTagBoolean {
	return m.Archive
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataArchive) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetArchive())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataArchive(structType any) BACnetConstructedDataArchive {
	if casted, ok := structType.(BACnetConstructedDataArchive); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataArchive); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataArchive) GetTypeName() string {
	return "BACnetConstructedDataArchive"
}

func (m *_BACnetConstructedDataArchive) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (archive)
	lengthInBits += m.Archive.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataArchive) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataArchive) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataArchive BACnetConstructedDataArchive, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataArchive"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataArchive")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	archive, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "archive", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'archive' field"))
	}
	m.Archive = archive

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), archive)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataArchive"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataArchive")
	}

	return m, nil
}

func (m *_BACnetConstructedDataArchive) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataArchive) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataArchive"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataArchive")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "archive", m.GetArchive(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'archive' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataArchive"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataArchive")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataArchive) IsBACnetConstructedDataArchive() {}

func (m *_BACnetConstructedDataArchive) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataArchive) deepCopy() *_BACnetConstructedDataArchive {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataArchiveCopy := &_BACnetConstructedDataArchive{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Archive.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataArchiveCopy
}

func (m *_BACnetConstructedDataArchive) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
