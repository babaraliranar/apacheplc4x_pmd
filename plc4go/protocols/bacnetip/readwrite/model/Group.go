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

// Group is the corresponding interface of Group
type Group interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() ReadableProperty
	// GetObjectName returns ObjectName (property field)
	GetObjectName() ReadableProperty
	// GetObjectType returns ObjectType (property field)
	GetObjectType() ReadableProperty
	// GetDescription returns Description (property field)
	GetDescription() OptionalProperty
	// GetListOfGroupMembers returns ListOfGroupMembers (property field)
	GetListOfGroupMembers() ReadableProperty
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() ReadableProperty
	// GetPropertyList returns PropertyList (property field)
	GetPropertyList() ReadableProperty
	// GetTags returns Tags (property field)
	GetTags() OptionalProperty
	// GetProfileLocation returns ProfileLocation (property field)
	GetProfileLocation() OptionalProperty
	// GetProfileName returns ProfileName (property field)
	GetProfileName() OptionalProperty
	// IsGroup is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGroup()
	// CreateBuilder creates a GroupBuilder
	CreateGroupBuilder() GroupBuilder
}

// _Group is the data-structure of this message
type _Group struct {
	ObjectIdentifier   ReadableProperty
	ObjectName         ReadableProperty
	ObjectType         ReadableProperty
	Description        OptionalProperty
	ListOfGroupMembers ReadableProperty
	PresentValue       ReadableProperty
	PropertyList       ReadableProperty
	Tags               OptionalProperty
	ProfileLocation    OptionalProperty
	ProfileName        OptionalProperty
}

var _ Group = (*_Group)(nil)

// NewGroup factory function for _Group
func NewGroup(objectIdentifier ReadableProperty, objectName ReadableProperty, objectType ReadableProperty, description OptionalProperty, listOfGroupMembers ReadableProperty, presentValue ReadableProperty, propertyList ReadableProperty, tags OptionalProperty, profileLocation OptionalProperty, profileName OptionalProperty) *_Group {
	if objectIdentifier == nil {
		panic("objectIdentifier of type ReadableProperty for Group must not be nil")
	}
	if objectName == nil {
		panic("objectName of type ReadableProperty for Group must not be nil")
	}
	if objectType == nil {
		panic("objectType of type ReadableProperty for Group must not be nil")
	}
	if description == nil {
		panic("description of type OptionalProperty for Group must not be nil")
	}
	if listOfGroupMembers == nil {
		panic("listOfGroupMembers of type ReadableProperty for Group must not be nil")
	}
	if presentValue == nil {
		panic("presentValue of type ReadableProperty for Group must not be nil")
	}
	if propertyList == nil {
		panic("propertyList of type ReadableProperty for Group must not be nil")
	}
	if tags == nil {
		panic("tags of type OptionalProperty for Group must not be nil")
	}
	if profileLocation == nil {
		panic("profileLocation of type OptionalProperty for Group must not be nil")
	}
	if profileName == nil {
		panic("profileName of type OptionalProperty for Group must not be nil")
	}
	return &_Group{ObjectIdentifier: objectIdentifier, ObjectName: objectName, ObjectType: objectType, Description: description, ListOfGroupMembers: listOfGroupMembers, PresentValue: presentValue, PropertyList: propertyList, Tags: tags, ProfileLocation: profileLocation, ProfileName: profileName}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GroupBuilder is a builder for Group
type GroupBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier ReadableProperty, objectName ReadableProperty, objectType ReadableProperty, description OptionalProperty, listOfGroupMembers ReadableProperty, presentValue ReadableProperty, propertyList ReadableProperty, tags OptionalProperty, profileLocation OptionalProperty, profileName OptionalProperty) GroupBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(ReadableProperty) GroupBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder
	// WithObjectName adds ObjectName (property field)
	WithObjectName(ReadableProperty) GroupBuilder
	// WithObjectNameBuilder adds ObjectName (property field) which is build by the builder
	WithObjectNameBuilder(func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder
	// WithObjectType adds ObjectType (property field)
	WithObjectType(ReadableProperty) GroupBuilder
	// WithObjectTypeBuilder adds ObjectType (property field) which is build by the builder
	WithObjectTypeBuilder(func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder
	// WithDescription adds Description (property field)
	WithDescription(OptionalProperty) GroupBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder
	// WithListOfGroupMembers adds ListOfGroupMembers (property field)
	WithListOfGroupMembers(ReadableProperty) GroupBuilder
	// WithListOfGroupMembersBuilder adds ListOfGroupMembers (property field) which is build by the builder
	WithListOfGroupMembersBuilder(func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(ReadableProperty) GroupBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder
	// WithPropertyList adds PropertyList (property field)
	WithPropertyList(ReadableProperty) GroupBuilder
	// WithPropertyListBuilder adds PropertyList (property field) which is build by the builder
	WithPropertyListBuilder(func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder
	// WithTags adds Tags (property field)
	WithTags(OptionalProperty) GroupBuilder
	// WithTagsBuilder adds Tags (property field) which is build by the builder
	WithTagsBuilder(func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder
	// WithProfileLocation adds ProfileLocation (property field)
	WithProfileLocation(OptionalProperty) GroupBuilder
	// WithProfileLocationBuilder adds ProfileLocation (property field) which is build by the builder
	WithProfileLocationBuilder(func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder
	// WithProfileName adds ProfileName (property field)
	WithProfileName(OptionalProperty) GroupBuilder
	// WithProfileNameBuilder adds ProfileName (property field) which is build by the builder
	WithProfileNameBuilder(func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder
	// Build builds the Group or returns an error if something is wrong
	Build() (Group, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Group
}

// NewGroupBuilder() creates a GroupBuilder
func NewGroupBuilder() GroupBuilder {
	return &_GroupBuilder{_Group: new(_Group)}
}

type _GroupBuilder struct {
	*_Group

	err *utils.MultiError
}

var _ (GroupBuilder) = (*_GroupBuilder)(nil)

func (b *_GroupBuilder) WithMandatoryFields(objectIdentifier ReadableProperty, objectName ReadableProperty, objectType ReadableProperty, description OptionalProperty, listOfGroupMembers ReadableProperty, presentValue ReadableProperty, propertyList ReadableProperty, tags OptionalProperty, profileLocation OptionalProperty, profileName OptionalProperty) GroupBuilder {
	return b.WithObjectIdentifier(objectIdentifier).WithObjectName(objectName).WithObjectType(objectType).WithDescription(description).WithListOfGroupMembers(listOfGroupMembers).WithPresentValue(presentValue).WithPropertyList(propertyList).WithTags(tags).WithProfileLocation(profileLocation).WithProfileName(profileName)
}

func (b *_GroupBuilder) WithObjectIdentifier(objectIdentifier ReadableProperty) GroupBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_GroupBuilder) WithObjectIdentifierBuilder(builderSupplier func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateReadablePropertyBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReadablePropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithObjectName(objectName ReadableProperty) GroupBuilder {
	b.ObjectName = objectName
	return b
}

func (b *_GroupBuilder) WithObjectNameBuilder(builderSupplier func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.ObjectName.CreateReadablePropertyBuilder())
	var err error
	b.ObjectName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReadablePropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithObjectType(objectType ReadableProperty) GroupBuilder {
	b.ObjectType = objectType
	return b
}

func (b *_GroupBuilder) WithObjectTypeBuilder(builderSupplier func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.ObjectType.CreateReadablePropertyBuilder())
	var err error
	b.ObjectType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReadablePropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithDescription(description OptionalProperty) GroupBuilder {
	b.Description = description
	return b
}

func (b *_GroupBuilder) WithDescriptionBuilder(builderSupplier func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.Description.CreateOptionalPropertyBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OptionalPropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithListOfGroupMembers(listOfGroupMembers ReadableProperty) GroupBuilder {
	b.ListOfGroupMembers = listOfGroupMembers
	return b
}

func (b *_GroupBuilder) WithListOfGroupMembersBuilder(builderSupplier func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.ListOfGroupMembers.CreateReadablePropertyBuilder())
	var err error
	b.ListOfGroupMembers, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReadablePropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithPresentValue(presentValue ReadableProperty) GroupBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_GroupBuilder) WithPresentValueBuilder(builderSupplier func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.PresentValue.CreateReadablePropertyBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReadablePropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithPropertyList(propertyList ReadableProperty) GroupBuilder {
	b.PropertyList = propertyList
	return b
}

func (b *_GroupBuilder) WithPropertyListBuilder(builderSupplier func(ReadablePropertyBuilder) ReadablePropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.PropertyList.CreateReadablePropertyBuilder())
	var err error
	b.PropertyList, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReadablePropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithTags(tags OptionalProperty) GroupBuilder {
	b.Tags = tags
	return b
}

func (b *_GroupBuilder) WithTagsBuilder(builderSupplier func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.Tags.CreateOptionalPropertyBuilder())
	var err error
	b.Tags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OptionalPropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithProfileLocation(profileLocation OptionalProperty) GroupBuilder {
	b.ProfileLocation = profileLocation
	return b
}

func (b *_GroupBuilder) WithProfileLocationBuilder(builderSupplier func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.ProfileLocation.CreateOptionalPropertyBuilder())
	var err error
	b.ProfileLocation, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OptionalPropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) WithProfileName(profileName OptionalProperty) GroupBuilder {
	b.ProfileName = profileName
	return b
}

func (b *_GroupBuilder) WithProfileNameBuilder(builderSupplier func(OptionalPropertyBuilder) OptionalPropertyBuilder) GroupBuilder {
	builder := builderSupplier(b.ProfileName.CreateOptionalPropertyBuilder())
	var err error
	b.ProfileName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OptionalPropertyBuilder failed"))
	}
	return b
}

func (b *_GroupBuilder) Build() (Group, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.ObjectName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectName' not set"))
	}
	if b.ObjectType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectType' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.ListOfGroupMembers == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfGroupMembers' not set"))
	}
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.PropertyList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyList' not set"))
	}
	if b.Tags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'tags' not set"))
	}
	if b.ProfileLocation == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'profileLocation' not set"))
	}
	if b.ProfileName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'profileName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Group.deepCopy(), nil
}

func (b *_GroupBuilder) MustBuild() Group {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_GroupBuilder) DeepCopy() any {
	_copy := b.CreateGroupBuilder().(*_GroupBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGroupBuilder creates a GroupBuilder
func (b *_Group) CreateGroupBuilder() GroupBuilder {
	if b == nil {
		return NewGroupBuilder()
	}
	return &_GroupBuilder{_Group: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Group) GetObjectIdentifier() ReadableProperty {
	return m.ObjectIdentifier
}

func (m *_Group) GetObjectName() ReadableProperty {
	return m.ObjectName
}

func (m *_Group) GetObjectType() ReadableProperty {
	return m.ObjectType
}

func (m *_Group) GetDescription() OptionalProperty {
	return m.Description
}

func (m *_Group) GetListOfGroupMembers() ReadableProperty {
	return m.ListOfGroupMembers
}

func (m *_Group) GetPresentValue() ReadableProperty {
	return m.PresentValue
}

func (m *_Group) GetPropertyList() ReadableProperty {
	return m.PropertyList
}

func (m *_Group) GetTags() OptionalProperty {
	return m.Tags
}

func (m *_Group) GetProfileLocation() OptionalProperty {
	return m.ProfileLocation
}

func (m *_Group) GetProfileName() OptionalProperty {
	return m.ProfileName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGroup(structType any) Group {
	if casted, ok := structType.(Group); ok {
		return casted
	}
	if casted, ok := structType.(*Group); ok {
		return *casted
	}
	return nil
}

func (m *_Group) GetTypeName() string {
	return "Group"
}

func (m *_Group) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits(ctx)

	// Simple field (objectType)
	lengthInBits += m.ObjectType.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (listOfGroupMembers)
	lengthInBits += m.ListOfGroupMembers.GetLengthInBits(ctx)

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// Simple field (propertyList)
	lengthInBits += m.PropertyList.GetLengthInBits(ctx)

	// Simple field (tags)
	lengthInBits += m.Tags.GetLengthInBits(ctx)

	// Simple field (profileLocation)
	lengthInBits += m.ProfileLocation.GetLengthInBits(ctx)

	// Simple field (profileName)
	lengthInBits += m.ProfileName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Group) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GroupParse(ctx context.Context, theBytes []byte) (Group, error) {
	return GroupParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GroupParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Group, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Group, error) {
		return GroupParseWithBuffer(ctx, readBuffer)
	}
}

func GroupParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Group, error) {
	v, err := (&_Group{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_Group) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__group Group, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Group"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Group")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[ReadableProperty](ctx, "objectIdentifier", ReadComplex[ReadableProperty](ReadablePropertyParseWithBufferProducer((string)("BACnetObjectIdentifier")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	objectName, err := ReadSimpleField[ReadableProperty](ctx, "objectName", ReadComplex[ReadableProperty](ReadablePropertyParseWithBufferProducer((string)("CharacterString")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectName' field"))
	}
	m.ObjectName = objectName

	objectType, err := ReadSimpleField[ReadableProperty](ctx, "objectType", ReadComplex[ReadableProperty](ReadablePropertyParseWithBufferProducer((string)("BACnetObjectType")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectType' field"))
	}
	m.ObjectType = objectType

	description, err := ReadSimpleField[OptionalProperty](ctx, "description", ReadComplex[OptionalProperty](OptionalPropertyParseWithBufferProducer((string)("CharacterString")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	listOfGroupMembers, err := ReadSimpleField[ReadableProperty](ctx, "listOfGroupMembers", ReadComplex[ReadableProperty](ReadablePropertyParseWithBufferProducer((string)("BACnetLIST of ReadAccessSpecification")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfGroupMembers' field"))
	}
	m.ListOfGroupMembers = listOfGroupMembers

	presentValue, err := ReadSimpleField[ReadableProperty](ctx, "presentValue", ReadComplex[ReadableProperty](ReadablePropertyParseWithBufferProducer((string)("BACnetLIST of ReadAccessResult")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	propertyList, err := ReadSimpleField[ReadableProperty](ctx, "propertyList", ReadComplex[ReadableProperty](ReadablePropertyParseWithBufferProducer((string)("BACnetARRAY[N] of BACnetPropertyIdentifier")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyList' field"))
	}
	m.PropertyList = propertyList

	tags, err := ReadSimpleField[OptionalProperty](ctx, "tags", ReadComplex[OptionalProperty](OptionalPropertyParseWithBufferProducer((string)("BACnetARRAY[N] of BACnetNameValue")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tags' field"))
	}
	m.Tags = tags

	profileLocation, err := ReadSimpleField[OptionalProperty](ctx, "profileLocation", ReadComplex[OptionalProperty](OptionalPropertyParseWithBufferProducer((string)("CharacterString")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'profileLocation' field"))
	}
	m.ProfileLocation = profileLocation

	profileName, err := ReadSimpleField[OptionalProperty](ctx, "profileName", ReadComplex[OptionalProperty](OptionalPropertyParseWithBufferProducer((string)("CharacterString")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'profileName' field"))
	}
	m.ProfileName = profileName

	if closeErr := readBuffer.CloseContext("Group"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Group")
	}

	return m, nil
}

func (m *_Group) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Group) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Group"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Group")
	}

	if err := WriteSimpleField[ReadableProperty](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[ReadableProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteSimpleField[ReadableProperty](ctx, "objectName", m.GetObjectName(), WriteComplex[ReadableProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectName' field")
	}

	if err := WriteSimpleField[ReadableProperty](ctx, "objectType", m.GetObjectType(), WriteComplex[ReadableProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectType' field")
	}

	if err := WriteSimpleField[OptionalProperty](ctx, "description", m.GetDescription(), WriteComplex[OptionalProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'description' field")
	}

	if err := WriteSimpleField[ReadableProperty](ctx, "listOfGroupMembers", m.GetListOfGroupMembers(), WriteComplex[ReadableProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfGroupMembers' field")
	}

	if err := WriteSimpleField[ReadableProperty](ctx, "presentValue", m.GetPresentValue(), WriteComplex[ReadableProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'presentValue' field")
	}

	if err := WriteSimpleField[ReadableProperty](ctx, "propertyList", m.GetPropertyList(), WriteComplex[ReadableProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyList' field")
	}

	if err := WriteSimpleField[OptionalProperty](ctx, "tags", m.GetTags(), WriteComplex[OptionalProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'tags' field")
	}

	if err := WriteSimpleField[OptionalProperty](ctx, "profileLocation", m.GetProfileLocation(), WriteComplex[OptionalProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'profileLocation' field")
	}

	if err := WriteSimpleField[OptionalProperty](ctx, "profileName", m.GetProfileName(), WriteComplex[OptionalProperty](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'profileName' field")
	}

	if popErr := writeBuffer.PopContext("Group"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Group")
	}
	return nil
}

func (m *_Group) IsGroup() {}

func (m *_Group) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Group) deepCopy() *_Group {
	if m == nil {
		return nil
	}
	_GroupCopy := &_Group{
		m.ObjectIdentifier.DeepCopy().(ReadableProperty),
		m.ObjectName.DeepCopy().(ReadableProperty),
		m.ObjectType.DeepCopy().(ReadableProperty),
		m.Description.DeepCopy().(OptionalProperty),
		m.ListOfGroupMembers.DeepCopy().(ReadableProperty),
		m.PresentValue.DeepCopy().(ReadableProperty),
		m.PropertyList.DeepCopy().(ReadableProperty),
		m.Tags.DeepCopy().(OptionalProperty),
		m.ProfileLocation.DeepCopy().(OptionalProperty),
		m.ProfileName.DeepCopy().(OptionalProperty),
	}
	return _GroupCopy
}

func (m *_Group) String() string {
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
