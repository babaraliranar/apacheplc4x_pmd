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

// NodeAttributes is the corresponding interface of NodeAttributes
type NodeAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSpecifiedAttributes returns SpecifiedAttributes (property field)
	GetSpecifiedAttributes() uint32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetWriteMask returns WriteMask (property field)
	GetWriteMask() uint32
	// GetUserWriteMask returns UserWriteMask (property field)
	GetUserWriteMask() uint32
	// IsNodeAttributes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeAttributes()
	// CreateBuilder creates a NodeAttributesBuilder
	CreateNodeAttributesBuilder() NodeAttributesBuilder
}

// _NodeAttributes is the data-structure of this message
type _NodeAttributes struct {
	ExtensionObjectDefinitionContract
	SpecifiedAttributes uint32
	DisplayName         LocalizedText
	Description         LocalizedText
	WriteMask           uint32
	UserWriteMask       uint32
}

var _ NodeAttributes = (*_NodeAttributes)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NodeAttributes)(nil)

// NewNodeAttributes factory function for _NodeAttributes
func NewNodeAttributes(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32) *_NodeAttributes {
	if displayName == nil {
		panic("displayName of type LocalizedText for NodeAttributes must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for NodeAttributes must not be nil")
	}
	_result := &_NodeAttributes{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SpecifiedAttributes:               specifiedAttributes,
		DisplayName:                       displayName,
		Description:                       description,
		WriteMask:                         writeMask,
		UserWriteMask:                     userWriteMask,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeAttributesBuilder is a builder for NodeAttributes
type NodeAttributesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32) NodeAttributesBuilder
	// WithSpecifiedAttributes adds SpecifiedAttributes (property field)
	WithSpecifiedAttributes(uint32) NodeAttributesBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) NodeAttributesBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) NodeAttributesBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) NodeAttributesBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) NodeAttributesBuilder
	// WithWriteMask adds WriteMask (property field)
	WithWriteMask(uint32) NodeAttributesBuilder
	// WithUserWriteMask adds UserWriteMask (property field)
	WithUserWriteMask(uint32) NodeAttributesBuilder
	// Build builds the NodeAttributes or returns an error if something is wrong
	Build() (NodeAttributes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeAttributes
}

// NewNodeAttributesBuilder() creates a NodeAttributesBuilder
func NewNodeAttributesBuilder() NodeAttributesBuilder {
	return &_NodeAttributesBuilder{_NodeAttributes: new(_NodeAttributes)}
}

type _NodeAttributesBuilder struct {
	*_NodeAttributes

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (NodeAttributesBuilder) = (*_NodeAttributesBuilder)(nil)

func (b *_NodeAttributesBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_NodeAttributesBuilder) WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32) NodeAttributesBuilder {
	return b.WithSpecifiedAttributes(specifiedAttributes).WithDisplayName(displayName).WithDescription(description).WithWriteMask(writeMask).WithUserWriteMask(userWriteMask)
}

func (b *_NodeAttributesBuilder) WithSpecifiedAttributes(specifiedAttributes uint32) NodeAttributesBuilder {
	b.SpecifiedAttributes = specifiedAttributes
	return b
}

func (b *_NodeAttributesBuilder) WithDisplayName(displayName LocalizedText) NodeAttributesBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_NodeAttributesBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) NodeAttributesBuilder {
	builder := builderSupplier(b.DisplayName.CreateLocalizedTextBuilder())
	var err error
	b.DisplayName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_NodeAttributesBuilder) WithDescription(description LocalizedText) NodeAttributesBuilder {
	b.Description = description
	return b
}

func (b *_NodeAttributesBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) NodeAttributesBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_NodeAttributesBuilder) WithWriteMask(writeMask uint32) NodeAttributesBuilder {
	b.WriteMask = writeMask
	return b
}

func (b *_NodeAttributesBuilder) WithUserWriteMask(userWriteMask uint32) NodeAttributesBuilder {
	b.UserWriteMask = userWriteMask
	return b
}

func (b *_NodeAttributesBuilder) Build() (NodeAttributes, error) {
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeAttributes.deepCopy(), nil
}

func (b *_NodeAttributesBuilder) MustBuild() NodeAttributes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NodeAttributesBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_NodeAttributesBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_NodeAttributesBuilder) DeepCopy() any {
	_copy := b.CreateNodeAttributesBuilder().(*_NodeAttributesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeAttributesBuilder creates a NodeAttributesBuilder
func (b *_NodeAttributes) CreateNodeAttributesBuilder() NodeAttributesBuilder {
	if b == nil {
		return NewNodeAttributesBuilder()
	}
	return &_NodeAttributesBuilder{_NodeAttributes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeAttributes) GetExtensionId() int32 {
	return int32(351)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeAttributes) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeAttributes) GetSpecifiedAttributes() uint32 {
	return m.SpecifiedAttributes
}

func (m *_NodeAttributes) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_NodeAttributes) GetDescription() LocalizedText {
	return m.Description
}

func (m *_NodeAttributes) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_NodeAttributes) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeAttributes(structType any) NodeAttributes {
	if casted, ok := structType.(NodeAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*NodeAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_NodeAttributes) GetTypeName() string {
	return "NodeAttributes"
}

func (m *_NodeAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (specifiedAttributes)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (writeMask)
	lengthInBits += 32

	// Simple field (userWriteMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_NodeAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeAttributes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__nodeAttributes NodeAttributes, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedAttributes, err := ReadSimpleField(ctx, "specifiedAttributes", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedAttributes' field"))
	}
	m.SpecifiedAttributes = specifiedAttributes

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	writeMask, err := ReadSimpleField(ctx, "writeMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeMask' field"))
	}
	m.WriteMask = writeMask

	userWriteMask, err := ReadSimpleField(ctx, "userWriteMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userWriteMask' field"))
	}
	m.UserWriteMask = userWriteMask

	if closeErr := readBuffer.CloseContext("NodeAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeAttributes")
	}

	return m, nil
}

func (m *_NodeAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeAttributes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeAttributes")
		}

		if err := WriteSimpleField[uint32](ctx, "specifiedAttributes", m.GetSpecifiedAttributes(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'specifiedAttributes' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleField[uint32](ctx, "writeMask", m.GetWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "userWriteMask", m.GetUserWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'userWriteMask' field")
		}

		if popErr := writeBuffer.PopContext("NodeAttributes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeAttributes")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeAttributes) IsNodeAttributes() {}

func (m *_NodeAttributes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeAttributes) deepCopy() *_NodeAttributes {
	if m == nil {
		return nil
	}
	_NodeAttributesCopy := &_NodeAttributes{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SpecifiedAttributes,
		m.DisplayName.DeepCopy().(LocalizedText),
		m.Description.DeepCopy().(LocalizedText),
		m.WriteMask,
		m.UserWriteMask,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NodeAttributesCopy
}

func (m *_NodeAttributes) String() string {
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
