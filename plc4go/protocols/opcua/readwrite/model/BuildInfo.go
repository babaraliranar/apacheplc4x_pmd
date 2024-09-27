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

// BuildInfo is the corresponding interface of BuildInfo
type BuildInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetProductUri returns ProductUri (property field)
	GetProductUri() PascalString
	// GetManufacturerName returns ManufacturerName (property field)
	GetManufacturerName() PascalString
	// GetProductName returns ProductName (property field)
	GetProductName() PascalString
	// GetSoftwareVersion returns SoftwareVersion (property field)
	GetSoftwareVersion() PascalString
	// GetBuildNumber returns BuildNumber (property field)
	GetBuildNumber() PascalString
	// GetBuildDate returns BuildDate (property field)
	GetBuildDate() int64
	// IsBuildInfo is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBuildInfo()
	// CreateBuilder creates a BuildInfoBuilder
	CreateBuildInfoBuilder() BuildInfoBuilder
}

// _BuildInfo is the data-structure of this message
type _BuildInfo struct {
	ExtensionObjectDefinitionContract
	ProductUri       PascalString
	ManufacturerName PascalString
	ProductName      PascalString
	SoftwareVersion  PascalString
	BuildNumber      PascalString
	BuildDate        int64
}

var _ BuildInfo = (*_BuildInfo)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BuildInfo)(nil)

// NewBuildInfo factory function for _BuildInfo
func NewBuildInfo(productUri PascalString, manufacturerName PascalString, productName PascalString, softwareVersion PascalString, buildNumber PascalString, buildDate int64) *_BuildInfo {
	if productUri == nil {
		panic("productUri of type PascalString for BuildInfo must not be nil")
	}
	if manufacturerName == nil {
		panic("manufacturerName of type PascalString for BuildInfo must not be nil")
	}
	if productName == nil {
		panic("productName of type PascalString for BuildInfo must not be nil")
	}
	if softwareVersion == nil {
		panic("softwareVersion of type PascalString for BuildInfo must not be nil")
	}
	if buildNumber == nil {
		panic("buildNumber of type PascalString for BuildInfo must not be nil")
	}
	_result := &_BuildInfo{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ProductUri:                        productUri,
		ManufacturerName:                  manufacturerName,
		ProductName:                       productName,
		SoftwareVersion:                   softwareVersion,
		BuildNumber:                       buildNumber,
		BuildDate:                         buildDate,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BuildInfoBuilder is a builder for BuildInfo
type BuildInfoBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(productUri PascalString, manufacturerName PascalString, productName PascalString, softwareVersion PascalString, buildNumber PascalString, buildDate int64) BuildInfoBuilder
	// WithProductUri adds ProductUri (property field)
	WithProductUri(PascalString) BuildInfoBuilder
	// WithProductUriBuilder adds ProductUri (property field) which is build by the builder
	WithProductUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder
	// WithManufacturerName adds ManufacturerName (property field)
	WithManufacturerName(PascalString) BuildInfoBuilder
	// WithManufacturerNameBuilder adds ManufacturerName (property field) which is build by the builder
	WithManufacturerNameBuilder(func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder
	// WithProductName adds ProductName (property field)
	WithProductName(PascalString) BuildInfoBuilder
	// WithProductNameBuilder adds ProductName (property field) which is build by the builder
	WithProductNameBuilder(func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder
	// WithSoftwareVersion adds SoftwareVersion (property field)
	WithSoftwareVersion(PascalString) BuildInfoBuilder
	// WithSoftwareVersionBuilder adds SoftwareVersion (property field) which is build by the builder
	WithSoftwareVersionBuilder(func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder
	// WithBuildNumber adds BuildNumber (property field)
	WithBuildNumber(PascalString) BuildInfoBuilder
	// WithBuildNumberBuilder adds BuildNumber (property field) which is build by the builder
	WithBuildNumberBuilder(func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder
	// WithBuildDate adds BuildDate (property field)
	WithBuildDate(int64) BuildInfoBuilder
	// Build builds the BuildInfo or returns an error if something is wrong
	Build() (BuildInfo, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BuildInfo
}

// NewBuildInfoBuilder() creates a BuildInfoBuilder
func NewBuildInfoBuilder() BuildInfoBuilder {
	return &_BuildInfoBuilder{_BuildInfo: new(_BuildInfo)}
}

type _BuildInfoBuilder struct {
	*_BuildInfo

	err *utils.MultiError
}

var _ (BuildInfoBuilder) = (*_BuildInfoBuilder)(nil)

func (m *_BuildInfoBuilder) WithMandatoryFields(productUri PascalString, manufacturerName PascalString, productName PascalString, softwareVersion PascalString, buildNumber PascalString, buildDate int64) BuildInfoBuilder {
	return m.WithProductUri(productUri).WithManufacturerName(manufacturerName).WithProductName(productName).WithSoftwareVersion(softwareVersion).WithBuildNumber(buildNumber).WithBuildDate(buildDate)
}

func (m *_BuildInfoBuilder) WithProductUri(productUri PascalString) BuildInfoBuilder {
	m.ProductUri = productUri
	return m
}

func (m *_BuildInfoBuilder) WithProductUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder {
	builder := builderSupplier(m.ProductUri.CreatePascalStringBuilder())
	var err error
	m.ProductUri, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_BuildInfoBuilder) WithManufacturerName(manufacturerName PascalString) BuildInfoBuilder {
	m.ManufacturerName = manufacturerName
	return m
}

func (m *_BuildInfoBuilder) WithManufacturerNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder {
	builder := builderSupplier(m.ManufacturerName.CreatePascalStringBuilder())
	var err error
	m.ManufacturerName, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_BuildInfoBuilder) WithProductName(productName PascalString) BuildInfoBuilder {
	m.ProductName = productName
	return m
}

func (m *_BuildInfoBuilder) WithProductNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder {
	builder := builderSupplier(m.ProductName.CreatePascalStringBuilder())
	var err error
	m.ProductName, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_BuildInfoBuilder) WithSoftwareVersion(softwareVersion PascalString) BuildInfoBuilder {
	m.SoftwareVersion = softwareVersion
	return m
}

func (m *_BuildInfoBuilder) WithSoftwareVersionBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder {
	builder := builderSupplier(m.SoftwareVersion.CreatePascalStringBuilder())
	var err error
	m.SoftwareVersion, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_BuildInfoBuilder) WithBuildNumber(buildNumber PascalString) BuildInfoBuilder {
	m.BuildNumber = buildNumber
	return m
}

func (m *_BuildInfoBuilder) WithBuildNumberBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BuildInfoBuilder {
	builder := builderSupplier(m.BuildNumber.CreatePascalStringBuilder())
	var err error
	m.BuildNumber, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_BuildInfoBuilder) WithBuildDate(buildDate int64) BuildInfoBuilder {
	m.BuildDate = buildDate
	return m
}

func (m *_BuildInfoBuilder) Build() (BuildInfo, error) {
	if m.ProductUri == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'productUri' not set"))
	}
	if m.ManufacturerName == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'manufacturerName' not set"))
	}
	if m.ProductName == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'productName' not set"))
	}
	if m.SoftwareVersion == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'softwareVersion' not set"))
	}
	if m.BuildNumber == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'buildNumber' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BuildInfo.deepCopy(), nil
}

func (m *_BuildInfoBuilder) MustBuild() BuildInfo {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BuildInfoBuilder) DeepCopy() any {
	return m.CreateBuildInfoBuilder()
}

// CreateBuildInfoBuilder creates a BuildInfoBuilder
func (m *_BuildInfo) CreateBuildInfoBuilder() BuildInfoBuilder {
	if m == nil {
		return NewBuildInfoBuilder()
	}
	return &_BuildInfoBuilder{_BuildInfo: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BuildInfo) GetIdentifier() string {
	return "340"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BuildInfo) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BuildInfo) GetProductUri() PascalString {
	return m.ProductUri
}

func (m *_BuildInfo) GetManufacturerName() PascalString {
	return m.ManufacturerName
}

func (m *_BuildInfo) GetProductName() PascalString {
	return m.ProductName
}

func (m *_BuildInfo) GetSoftwareVersion() PascalString {
	return m.SoftwareVersion
}

func (m *_BuildInfo) GetBuildNumber() PascalString {
	return m.BuildNumber
}

func (m *_BuildInfo) GetBuildDate() int64 {
	return m.BuildDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBuildInfo(structType any) BuildInfo {
	if casted, ok := structType.(BuildInfo); ok {
		return casted
	}
	if casted, ok := structType.(*BuildInfo); ok {
		return *casted
	}
	return nil
}

func (m *_BuildInfo) GetTypeName() string {
	return "BuildInfo"
}

func (m *_BuildInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (productUri)
	lengthInBits += m.ProductUri.GetLengthInBits(ctx)

	// Simple field (manufacturerName)
	lengthInBits += m.ManufacturerName.GetLengthInBits(ctx)

	// Simple field (productName)
	lengthInBits += m.ProductName.GetLengthInBits(ctx)

	// Simple field (softwareVersion)
	lengthInBits += m.SoftwareVersion.GetLengthInBits(ctx)

	// Simple field (buildNumber)
	lengthInBits += m.BuildNumber.GetLengthInBits(ctx)

	// Simple field (buildDate)
	lengthInBits += 64

	return lengthInBits
}

func (m *_BuildInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BuildInfo) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__buildInfo BuildInfo, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BuildInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BuildInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	productUri, err := ReadSimpleField[PascalString](ctx, "productUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productUri' field"))
	}
	m.ProductUri = productUri

	manufacturerName, err := ReadSimpleField[PascalString](ctx, "manufacturerName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'manufacturerName' field"))
	}
	m.ManufacturerName = manufacturerName

	productName, err := ReadSimpleField[PascalString](ctx, "productName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productName' field"))
	}
	m.ProductName = productName

	softwareVersion, err := ReadSimpleField[PascalString](ctx, "softwareVersion", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'softwareVersion' field"))
	}
	m.SoftwareVersion = softwareVersion

	buildNumber, err := ReadSimpleField[PascalString](ctx, "buildNumber", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'buildNumber' field"))
	}
	m.BuildNumber = buildNumber

	buildDate, err := ReadSimpleField(ctx, "buildDate", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'buildDate' field"))
	}
	m.BuildDate = buildDate

	if closeErr := readBuffer.CloseContext("BuildInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BuildInfo")
	}

	return m, nil
}

func (m *_BuildInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BuildInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BuildInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BuildInfo")
		}

		if err := WriteSimpleField[PascalString](ctx, "productUri", m.GetProductUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'productUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "manufacturerName", m.GetManufacturerName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'manufacturerName' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "productName", m.GetProductName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'productName' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "softwareVersion", m.GetSoftwareVersion(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'softwareVersion' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "buildNumber", m.GetBuildNumber(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'buildNumber' field")
		}

		if err := WriteSimpleField[int64](ctx, "buildDate", m.GetBuildDate(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'buildDate' field")
		}

		if popErr := writeBuffer.PopContext("BuildInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BuildInfo")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BuildInfo) IsBuildInfo() {}

func (m *_BuildInfo) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BuildInfo) deepCopy() *_BuildInfo {
	if m == nil {
		return nil
	}
	_BuildInfoCopy := &_BuildInfo{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ProductUri.DeepCopy().(PascalString),
		m.ManufacturerName.DeepCopy().(PascalString),
		m.ProductName.DeepCopy().(PascalString),
		m.SoftwareVersion.DeepCopy().(PascalString),
		m.BuildNumber.DeepCopy().(PascalString),
		m.BuildDate,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BuildInfoCopy
}

func (m *_BuildInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
