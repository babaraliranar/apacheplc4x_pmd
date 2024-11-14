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

// ReferenceDescription is the corresponding interface of ReferenceDescription
type ReferenceDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetNodeId returns NodeId (property field)
	GetNodeId() ExpandedNodeId
	// GetBrowseName returns BrowseName (property field)
	GetBrowseName() QualifiedName
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetNodeClass returns NodeClass (property field)
	GetNodeClass() NodeClass
	// GetTypeDefinition returns TypeDefinition (property field)
	GetTypeDefinition() ExpandedNodeId
	// IsReferenceDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReferenceDescription()
	// CreateBuilder creates a ReferenceDescriptionBuilder
	CreateReferenceDescriptionBuilder() ReferenceDescriptionBuilder
}

// _ReferenceDescription is the data-structure of this message
type _ReferenceDescription struct {
	ExtensionObjectDefinitionContract
	ReferenceTypeId NodeId
	IsForward       bool
	NodeId          ExpandedNodeId
	BrowseName      QualifiedName
	DisplayName     LocalizedText
	NodeClass       NodeClass
	TypeDefinition  ExpandedNodeId
	// Reserved Fields
	reservedField0 *uint8
}

var _ ReferenceDescription = (*_ReferenceDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReferenceDescription)(nil)

// NewReferenceDescription factory function for _ReferenceDescription
func NewReferenceDescription(referenceTypeId NodeId, isForward bool, nodeId ExpandedNodeId, browseName QualifiedName, displayName LocalizedText, nodeClass NodeClass, typeDefinition ExpandedNodeId) *_ReferenceDescription {
	if referenceTypeId == nil {
		panic("referenceTypeId of type NodeId for ReferenceDescription must not be nil")
	}
	if nodeId == nil {
		panic("nodeId of type ExpandedNodeId for ReferenceDescription must not be nil")
	}
	if browseName == nil {
		panic("browseName of type QualifiedName for ReferenceDescription must not be nil")
	}
	if displayName == nil {
		panic("displayName of type LocalizedText for ReferenceDescription must not be nil")
	}
	if typeDefinition == nil {
		panic("typeDefinition of type ExpandedNodeId for ReferenceDescription must not be nil")
	}
	_result := &_ReferenceDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ReferenceTypeId:                   referenceTypeId,
		IsForward:                         isForward,
		NodeId:                            nodeId,
		BrowseName:                        browseName,
		DisplayName:                       displayName,
		NodeClass:                         nodeClass,
		TypeDefinition:                    typeDefinition,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReferenceDescriptionBuilder is a builder for ReferenceDescription
type ReferenceDescriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referenceTypeId NodeId, isForward bool, nodeId ExpandedNodeId, browseName QualifiedName, displayName LocalizedText, nodeClass NodeClass, typeDefinition ExpandedNodeId) ReferenceDescriptionBuilder
	// WithReferenceTypeId adds ReferenceTypeId (property field)
	WithReferenceTypeId(NodeId) ReferenceDescriptionBuilder
	// WithReferenceTypeIdBuilder adds ReferenceTypeId (property field) which is build by the builder
	WithReferenceTypeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) ReferenceDescriptionBuilder
	// WithIsForward adds IsForward (property field)
	WithIsForward(bool) ReferenceDescriptionBuilder
	// WithNodeId adds NodeId (property field)
	WithNodeId(ExpandedNodeId) ReferenceDescriptionBuilder
	// WithNodeIdBuilder adds NodeId (property field) which is build by the builder
	WithNodeIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ReferenceDescriptionBuilder
	// WithBrowseName adds BrowseName (property field)
	WithBrowseName(QualifiedName) ReferenceDescriptionBuilder
	// WithBrowseNameBuilder adds BrowseName (property field) which is build by the builder
	WithBrowseNameBuilder(func(QualifiedNameBuilder) QualifiedNameBuilder) ReferenceDescriptionBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) ReferenceDescriptionBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) ReferenceDescriptionBuilder
	// WithNodeClass adds NodeClass (property field)
	WithNodeClass(NodeClass) ReferenceDescriptionBuilder
	// WithTypeDefinition adds TypeDefinition (property field)
	WithTypeDefinition(ExpandedNodeId) ReferenceDescriptionBuilder
	// WithTypeDefinitionBuilder adds TypeDefinition (property field) which is build by the builder
	WithTypeDefinitionBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ReferenceDescriptionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ReferenceDescription or returns an error if something is wrong
	Build() (ReferenceDescription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReferenceDescription
}

// NewReferenceDescriptionBuilder() creates a ReferenceDescriptionBuilder
func NewReferenceDescriptionBuilder() ReferenceDescriptionBuilder {
	return &_ReferenceDescriptionBuilder{_ReferenceDescription: new(_ReferenceDescription)}
}

type _ReferenceDescriptionBuilder struct {
	*_ReferenceDescription

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReferenceDescriptionBuilder) = (*_ReferenceDescriptionBuilder)(nil)

func (b *_ReferenceDescriptionBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ReferenceDescription
}

func (b *_ReferenceDescriptionBuilder) WithMandatoryFields(referenceTypeId NodeId, isForward bool, nodeId ExpandedNodeId, browseName QualifiedName, displayName LocalizedText, nodeClass NodeClass, typeDefinition ExpandedNodeId) ReferenceDescriptionBuilder {
	return b.WithReferenceTypeId(referenceTypeId).WithIsForward(isForward).WithNodeId(nodeId).WithBrowseName(browseName).WithDisplayName(displayName).WithNodeClass(nodeClass).WithTypeDefinition(typeDefinition)
}

func (b *_ReferenceDescriptionBuilder) WithReferenceTypeId(referenceTypeId NodeId) ReferenceDescriptionBuilder {
	b.ReferenceTypeId = referenceTypeId
	return b
}

func (b *_ReferenceDescriptionBuilder) WithReferenceTypeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ReferenceDescriptionBuilder {
	builder := builderSupplier(b.ReferenceTypeId.CreateNodeIdBuilder())
	var err error
	b.ReferenceTypeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionBuilder) WithIsForward(isForward bool) ReferenceDescriptionBuilder {
	b.IsForward = isForward
	return b
}

func (b *_ReferenceDescriptionBuilder) WithNodeId(nodeId ExpandedNodeId) ReferenceDescriptionBuilder {
	b.NodeId = nodeId
	return b
}

func (b *_ReferenceDescriptionBuilder) WithNodeIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ReferenceDescriptionBuilder {
	builder := builderSupplier(b.NodeId.CreateExpandedNodeIdBuilder())
	var err error
	b.NodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionBuilder) WithBrowseName(browseName QualifiedName) ReferenceDescriptionBuilder {
	b.BrowseName = browseName
	return b
}

func (b *_ReferenceDescriptionBuilder) WithBrowseNameBuilder(builderSupplier func(QualifiedNameBuilder) QualifiedNameBuilder) ReferenceDescriptionBuilder {
	builder := builderSupplier(b.BrowseName.CreateQualifiedNameBuilder())
	var err error
	b.BrowseName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "QualifiedNameBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionBuilder) WithDisplayName(displayName LocalizedText) ReferenceDescriptionBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_ReferenceDescriptionBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) ReferenceDescriptionBuilder {
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

func (b *_ReferenceDescriptionBuilder) WithNodeClass(nodeClass NodeClass) ReferenceDescriptionBuilder {
	b.NodeClass = nodeClass
	return b
}

func (b *_ReferenceDescriptionBuilder) WithTypeDefinition(typeDefinition ExpandedNodeId) ReferenceDescriptionBuilder {
	b.TypeDefinition = typeDefinition
	return b
}

func (b *_ReferenceDescriptionBuilder) WithTypeDefinitionBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ReferenceDescriptionBuilder {
	builder := builderSupplier(b.TypeDefinition.CreateExpandedNodeIdBuilder())
	var err error
	b.TypeDefinition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionBuilder) Build() (ReferenceDescription, error) {
	if b.ReferenceTypeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceTypeId' not set"))
	}
	if b.NodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeId' not set"))
	}
	if b.BrowseName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'browseName' not set"))
	}
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.TypeDefinition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'typeDefinition' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReferenceDescription.deepCopy(), nil
}

func (b *_ReferenceDescriptionBuilder) MustBuild() ReferenceDescription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReferenceDescriptionBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ReferenceDescriptionBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReferenceDescriptionBuilder) DeepCopy() any {
	_copy := b.CreateReferenceDescriptionBuilder().(*_ReferenceDescriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReferenceDescriptionBuilder creates a ReferenceDescriptionBuilder
func (b *_ReferenceDescription) CreateReferenceDescriptionBuilder() ReferenceDescriptionBuilder {
	if b == nil {
		return NewReferenceDescriptionBuilder()
	}
	return &_ReferenceDescriptionBuilder{_ReferenceDescription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReferenceDescription) GetExtensionId() int32 {
	return int32(520)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReferenceDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReferenceDescription) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_ReferenceDescription) GetIsForward() bool {
	return m.IsForward
}

func (m *_ReferenceDescription) GetNodeId() ExpandedNodeId {
	return m.NodeId
}

func (m *_ReferenceDescription) GetBrowseName() QualifiedName {
	return m.BrowseName
}

func (m *_ReferenceDescription) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_ReferenceDescription) GetNodeClass() NodeClass {
	return m.NodeClass
}

func (m *_ReferenceDescription) GetTypeDefinition() ExpandedNodeId {
	return m.TypeDefinition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReferenceDescription(structType any) ReferenceDescription {
	if casted, ok := structType.(ReferenceDescription); ok {
		return casted
	}
	if casted, ok := structType.(*ReferenceDescription); ok {
		return *casted
	}
	return nil
}

func (m *_ReferenceDescription) GetTypeName() string {
	return "ReferenceDescription"
}

func (m *_ReferenceDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (browseName)
	lengthInBits += m.BrowseName.GetLengthInBits(ctx)

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (nodeClass)
	lengthInBits += 32

	// Simple field (typeDefinition)
	lengthInBits += m.TypeDefinition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReferenceDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReferenceDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__referenceDescription ReferenceDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReferenceDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReferenceDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}
	m.ReferenceTypeId = referenceTypeId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	isForward, err := ReadSimpleField(ctx, "isForward", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isForward' field"))
	}
	m.IsForward = isForward

	nodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "nodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	browseName, err := ReadSimpleField[QualifiedName](ctx, "browseName", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'browseName' field"))
	}
	m.BrowseName = browseName

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	nodeClass, err := ReadEnumField[NodeClass](ctx, "nodeClass", "NodeClass", ReadEnum(NodeClassByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeClass' field"))
	}
	m.NodeClass = nodeClass

	typeDefinition, err := ReadSimpleField[ExpandedNodeId](ctx, "typeDefinition", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinition' field"))
	}
	m.TypeDefinition = typeDefinition

	if closeErr := readBuffer.CloseContext("ReferenceDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReferenceDescription")
	}

	return m, nil
}

func (m *_ReferenceDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReferenceDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReferenceDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReferenceDescription")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isForward", m.GetIsForward(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isForward' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "browseName", m.GetBrowseName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'browseName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleEnumField[NodeClass](ctx, "nodeClass", "NodeClass", m.GetNodeClass(), WriteEnum[NodeClass, uint32](NodeClass.GetValue, NodeClass.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeClass' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "typeDefinition", m.GetTypeDefinition(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinition' field")
		}

		if popErr := writeBuffer.PopContext("ReferenceDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReferenceDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReferenceDescription) IsReferenceDescription() {}

func (m *_ReferenceDescription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReferenceDescription) deepCopy() *_ReferenceDescription {
	if m == nil {
		return nil
	}
	_ReferenceDescriptionCopy := &_ReferenceDescription{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.ReferenceTypeId),
		m.IsForward,
		utils.DeepCopy[ExpandedNodeId](m.NodeId),
		utils.DeepCopy[QualifiedName](m.BrowseName),
		utils.DeepCopy[LocalizedText](m.DisplayName),
		m.NodeClass,
		utils.DeepCopy[ExpandedNodeId](m.TypeDefinition),
		m.reservedField0,
	}
	_ReferenceDescriptionCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReferenceDescriptionCopy
}

func (m *_ReferenceDescription) String() string {
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
