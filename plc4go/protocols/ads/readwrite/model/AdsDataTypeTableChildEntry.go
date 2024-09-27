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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AdsDataTypeTableChildEntry_PROPERTYNAMETERMINATOR uint8 = 0x00
const AdsDataTypeTableChildEntry_DATATYPENAMETERMINATOR uint8 = 0x00
const AdsDataTypeTableChildEntry_COMMENTTERMINATOR uint8 = 0x00

// AdsDataTypeTableChildEntry is the corresponding interface of AdsDataTypeTableChildEntry
type AdsDataTypeTableChildEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetEntryLength returns EntryLength (property field)
	GetEntryLength() uint32
	// GetVersion returns Version (property field)
	GetVersion() uint32
	// GetHashValue returns HashValue (property field)
	GetHashValue() uint32
	// GetTypeHashValue returns TypeHashValue (property field)
	GetTypeHashValue() uint32
	// GetSize returns Size (property field)
	GetSize() uint32
	// GetOffset returns Offset (property field)
	GetOffset() uint32
	// GetDataType returns DataType (property field)
	GetDataType() uint32
	// GetFlags returns Flags (property field)
	GetFlags() uint32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() uint16
	// GetNumChildren returns NumChildren (property field)
	GetNumChildren() uint16
	// GetPropertyName returns PropertyName (property field)
	GetPropertyName() string
	// GetDataTypeName returns DataTypeName (property field)
	GetDataTypeName() string
	// GetComment returns Comment (property field)
	GetComment() string
	// GetArrayInfo returns ArrayInfo (property field)
	GetArrayInfo() []AdsDataTypeArrayInfo
	// GetChildren returns Children (property field)
	GetChildren() []AdsDataTypeTableEntry
	// GetRest returns Rest (property field)
	GetRest() []byte
	// IsAdsDataTypeTableChildEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDataTypeTableChildEntry()
	// CreateBuilder creates a AdsDataTypeTableChildEntryBuilder
	CreateAdsDataTypeTableChildEntryBuilder() AdsDataTypeTableChildEntryBuilder
}

// _AdsDataTypeTableChildEntry is the data-structure of this message
type _AdsDataTypeTableChildEntry struct {
	EntryLength     uint32
	Version         uint32
	HashValue       uint32
	TypeHashValue   uint32
	Size            uint32
	Offset          uint32
	DataType        uint32
	Flags           uint32
	ArrayDimensions uint16
	NumChildren     uint16
	PropertyName    string
	DataTypeName    string
	Comment         string
	ArrayInfo       []AdsDataTypeArrayInfo
	Children        []AdsDataTypeTableEntry
	Rest            []byte
}

var _ AdsDataTypeTableChildEntry = (*_AdsDataTypeTableChildEntry)(nil)

// NewAdsDataTypeTableChildEntry factory function for _AdsDataTypeTableChildEntry
func NewAdsDataTypeTableChildEntry(entryLength uint32, version uint32, hashValue uint32, typeHashValue uint32, size uint32, offset uint32, dataType uint32, flags uint32, arrayDimensions uint16, numChildren uint16, propertyName string, dataTypeName string, comment string, arrayInfo []AdsDataTypeArrayInfo, children []AdsDataTypeTableEntry, rest []byte) *_AdsDataTypeTableChildEntry {
	return &_AdsDataTypeTableChildEntry{EntryLength: entryLength, Version: version, HashValue: hashValue, TypeHashValue: typeHashValue, Size: size, Offset: offset, DataType: dataType, Flags: flags, ArrayDimensions: arrayDimensions, NumChildren: numChildren, PropertyName: propertyName, DataTypeName: dataTypeName, Comment: comment, ArrayInfo: arrayInfo, Children: children, Rest: rest}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDataTypeTableChildEntryBuilder is a builder for AdsDataTypeTableChildEntry
type AdsDataTypeTableChildEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(entryLength uint32, version uint32, hashValue uint32, typeHashValue uint32, size uint32, offset uint32, dataType uint32, flags uint32, arrayDimensions uint16, numChildren uint16, propertyName string, dataTypeName string, comment string, arrayInfo []AdsDataTypeArrayInfo, children []AdsDataTypeTableEntry, rest []byte) AdsDataTypeTableChildEntryBuilder
	// WithEntryLength adds EntryLength (property field)
	WithEntryLength(uint32) AdsDataTypeTableChildEntryBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint32) AdsDataTypeTableChildEntryBuilder
	// WithHashValue adds HashValue (property field)
	WithHashValue(uint32) AdsDataTypeTableChildEntryBuilder
	// WithTypeHashValue adds TypeHashValue (property field)
	WithTypeHashValue(uint32) AdsDataTypeTableChildEntryBuilder
	// WithSize adds Size (property field)
	WithSize(uint32) AdsDataTypeTableChildEntryBuilder
	// WithOffset adds Offset (property field)
	WithOffset(uint32) AdsDataTypeTableChildEntryBuilder
	// WithDataType adds DataType (property field)
	WithDataType(uint32) AdsDataTypeTableChildEntryBuilder
	// WithFlags adds Flags (property field)
	WithFlags(uint32) AdsDataTypeTableChildEntryBuilder
	// WithArrayDimensions adds ArrayDimensions (property field)
	WithArrayDimensions(uint16) AdsDataTypeTableChildEntryBuilder
	// WithNumChildren adds NumChildren (property field)
	WithNumChildren(uint16) AdsDataTypeTableChildEntryBuilder
	// WithPropertyName adds PropertyName (property field)
	WithPropertyName(string) AdsDataTypeTableChildEntryBuilder
	// WithDataTypeName adds DataTypeName (property field)
	WithDataTypeName(string) AdsDataTypeTableChildEntryBuilder
	// WithComment adds Comment (property field)
	WithComment(string) AdsDataTypeTableChildEntryBuilder
	// WithArrayInfo adds ArrayInfo (property field)
	WithArrayInfo(...AdsDataTypeArrayInfo) AdsDataTypeTableChildEntryBuilder
	// WithChildren adds Children (property field)
	WithChildren(...AdsDataTypeTableEntry) AdsDataTypeTableChildEntryBuilder
	// WithRest adds Rest (property field)
	WithRest(...byte) AdsDataTypeTableChildEntryBuilder
	// Build builds the AdsDataTypeTableChildEntry or returns an error if something is wrong
	Build() (AdsDataTypeTableChildEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDataTypeTableChildEntry
}

// NewAdsDataTypeTableChildEntryBuilder() creates a AdsDataTypeTableChildEntryBuilder
func NewAdsDataTypeTableChildEntryBuilder() AdsDataTypeTableChildEntryBuilder {
	return &_AdsDataTypeTableChildEntryBuilder{_AdsDataTypeTableChildEntry: new(_AdsDataTypeTableChildEntry)}
}

type _AdsDataTypeTableChildEntryBuilder struct {
	*_AdsDataTypeTableChildEntry

	err *utils.MultiError
}

var _ (AdsDataTypeTableChildEntryBuilder) = (*_AdsDataTypeTableChildEntryBuilder)(nil)

func (m *_AdsDataTypeTableChildEntryBuilder) WithMandatoryFields(entryLength uint32, version uint32, hashValue uint32, typeHashValue uint32, size uint32, offset uint32, dataType uint32, flags uint32, arrayDimensions uint16, numChildren uint16, propertyName string, dataTypeName string, comment string, arrayInfo []AdsDataTypeArrayInfo, children []AdsDataTypeTableEntry, rest []byte) AdsDataTypeTableChildEntryBuilder {
	return m.WithEntryLength(entryLength).WithVersion(version).WithHashValue(hashValue).WithTypeHashValue(typeHashValue).WithSize(size).WithOffset(offset).WithDataType(dataType).WithFlags(flags).WithArrayDimensions(arrayDimensions).WithNumChildren(numChildren).WithPropertyName(propertyName).WithDataTypeName(dataTypeName).WithComment(comment).WithArrayInfo(arrayInfo...).WithChildren(children...).WithRest(rest...)
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithEntryLength(entryLength uint32) AdsDataTypeTableChildEntryBuilder {
	m.EntryLength = entryLength
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithVersion(version uint32) AdsDataTypeTableChildEntryBuilder {
	m.Version = version
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithHashValue(hashValue uint32) AdsDataTypeTableChildEntryBuilder {
	m.HashValue = hashValue
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithTypeHashValue(typeHashValue uint32) AdsDataTypeTableChildEntryBuilder {
	m.TypeHashValue = typeHashValue
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithSize(size uint32) AdsDataTypeTableChildEntryBuilder {
	m.Size = size
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithOffset(offset uint32) AdsDataTypeTableChildEntryBuilder {
	m.Offset = offset
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithDataType(dataType uint32) AdsDataTypeTableChildEntryBuilder {
	m.DataType = dataType
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithFlags(flags uint32) AdsDataTypeTableChildEntryBuilder {
	m.Flags = flags
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithArrayDimensions(arrayDimensions uint16) AdsDataTypeTableChildEntryBuilder {
	m.ArrayDimensions = arrayDimensions
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithNumChildren(numChildren uint16) AdsDataTypeTableChildEntryBuilder {
	m.NumChildren = numChildren
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithPropertyName(propertyName string) AdsDataTypeTableChildEntryBuilder {
	m.PropertyName = propertyName
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithDataTypeName(dataTypeName string) AdsDataTypeTableChildEntryBuilder {
	m.DataTypeName = dataTypeName
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithComment(comment string) AdsDataTypeTableChildEntryBuilder {
	m.Comment = comment
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithArrayInfo(arrayInfo ...AdsDataTypeArrayInfo) AdsDataTypeTableChildEntryBuilder {
	m.ArrayInfo = arrayInfo
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithChildren(children ...AdsDataTypeTableEntry) AdsDataTypeTableChildEntryBuilder {
	m.Children = children
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) WithRest(rest ...byte) AdsDataTypeTableChildEntryBuilder {
	m.Rest = rest
	return m
}

func (m *_AdsDataTypeTableChildEntryBuilder) Build() (AdsDataTypeTableChildEntry, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AdsDataTypeTableChildEntry.deepCopy(), nil
}

func (m *_AdsDataTypeTableChildEntryBuilder) MustBuild() AdsDataTypeTableChildEntry {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AdsDataTypeTableChildEntryBuilder) DeepCopy() any {
	return m.CreateAdsDataTypeTableChildEntryBuilder()
}

// CreateAdsDataTypeTableChildEntryBuilder creates a AdsDataTypeTableChildEntryBuilder
func (m *_AdsDataTypeTableChildEntry) CreateAdsDataTypeTableChildEntryBuilder() AdsDataTypeTableChildEntryBuilder {
	if m == nil {
		return NewAdsDataTypeTableChildEntryBuilder()
	}
	return &_AdsDataTypeTableChildEntryBuilder{_AdsDataTypeTableChildEntry: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDataTypeTableChildEntry) GetEntryLength() uint32 {
	return m.EntryLength
}

func (m *_AdsDataTypeTableChildEntry) GetVersion() uint32 {
	return m.Version
}

func (m *_AdsDataTypeTableChildEntry) GetHashValue() uint32 {
	return m.HashValue
}

func (m *_AdsDataTypeTableChildEntry) GetTypeHashValue() uint32 {
	return m.TypeHashValue
}

func (m *_AdsDataTypeTableChildEntry) GetSize() uint32 {
	return m.Size
}

func (m *_AdsDataTypeTableChildEntry) GetOffset() uint32 {
	return m.Offset
}

func (m *_AdsDataTypeTableChildEntry) GetDataType() uint32 {
	return m.DataType
}

func (m *_AdsDataTypeTableChildEntry) GetFlags() uint32 {
	return m.Flags
}

func (m *_AdsDataTypeTableChildEntry) GetArrayDimensions() uint16 {
	return m.ArrayDimensions
}

func (m *_AdsDataTypeTableChildEntry) GetNumChildren() uint16 {
	return m.NumChildren
}

func (m *_AdsDataTypeTableChildEntry) GetPropertyName() string {
	return m.PropertyName
}

func (m *_AdsDataTypeTableChildEntry) GetDataTypeName() string {
	return m.DataTypeName
}

func (m *_AdsDataTypeTableChildEntry) GetComment() string {
	return m.Comment
}

func (m *_AdsDataTypeTableChildEntry) GetArrayInfo() []AdsDataTypeArrayInfo {
	return m.ArrayInfo
}

func (m *_AdsDataTypeTableChildEntry) GetChildren() []AdsDataTypeTableEntry {
	return m.Children
}

func (m *_AdsDataTypeTableChildEntry) GetRest() []byte {
	return m.Rest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDataTypeTableChildEntry) GetPropertyNameTerminator() uint8 {
	return AdsDataTypeTableChildEntry_PROPERTYNAMETERMINATOR
}

func (m *_AdsDataTypeTableChildEntry) GetDataTypeNameTerminator() uint8 {
	return AdsDataTypeTableChildEntry_DATATYPENAMETERMINATOR
}

func (m *_AdsDataTypeTableChildEntry) GetCommentTerminator() uint8 {
	return AdsDataTypeTableChildEntry_COMMENTTERMINATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDataTypeTableChildEntry(structType any) AdsDataTypeTableChildEntry {
	if casted, ok := structType.(AdsDataTypeTableChildEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDataTypeTableChildEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDataTypeTableChildEntry) GetTypeName() string {
	return "AdsDataTypeTableChildEntry"
}

func (m *_AdsDataTypeTableChildEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (entryLength)
	lengthInBits += 32

	// Simple field (version)
	lengthInBits += 32

	// Simple field (hashValue)
	lengthInBits += 32

	// Simple field (typeHashValue)
	lengthInBits += 32

	// Simple field (size)
	lengthInBits += 32

	// Simple field (offset)
	lengthInBits += 32

	// Simple field (dataType)
	lengthInBits += 32

	// Simple field (flags)
	lengthInBits += 32

	// Implicit Field (propertyNameLength)
	lengthInBits += 16

	// Implicit Field (dataTypeNameLength)
	lengthInBits += 16

	// Implicit Field (commentLength)
	lengthInBits += 16

	// Simple field (arrayDimensions)
	lengthInBits += 16

	// Simple field (numChildren)
	lengthInBits += 16

	// Simple field (propertyName)
	lengthInBits += uint16(int32(uint16(len(m.GetPropertyName()))) * int32(int32(8)))

	// Const Field (propertyNameTerminator)
	lengthInBits += 8

	// Simple field (dataTypeName)
	lengthInBits += uint16(int32(uint16(len(m.GetDataTypeName()))) * int32(int32(8)))

	// Const Field (dataTypeNameTerminator)
	lengthInBits += 8

	// Simple field (comment)
	lengthInBits += uint16(int32(uint16(len(m.GetComment()))) * int32(int32(8)))

	// Const Field (commentTerminator)
	lengthInBits += 8

	// Array field
	if len(m.ArrayInfo) > 0 {
		for _curItem, element := range m.ArrayInfo {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ArrayInfo), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.Children) > 0 {
		for _curItem, element := range m.Children {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Children), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.Rest) > 0 {
		lengthInBits += 8 * uint16(len(m.Rest))
	}

	return lengthInBits
}

func (m *_AdsDataTypeTableChildEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDataTypeTableChildEntryParse(ctx context.Context, theBytes []byte) (AdsDataTypeTableChildEntry, error) {
	return AdsDataTypeTableChildEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsDataTypeTableChildEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeTableChildEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeTableChildEntry, error) {
		return AdsDataTypeTableChildEntryParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDataTypeTableChildEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeTableChildEntry, error) {
	v, err := (&_AdsDataTypeTableChildEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AdsDataTypeTableChildEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__adsDataTypeTableChildEntry AdsDataTypeTableChildEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDataTypeTableChildEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDataTypeTableChildEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos

	entryLength, err := ReadSimpleField(ctx, "entryLength", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entryLength' field"))
	}
	m.EntryLength = entryLength

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	hashValue, err := ReadSimpleField(ctx, "hashValue", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hashValue' field"))
	}
	m.HashValue = hashValue

	typeHashValue, err := ReadSimpleField(ctx, "typeHashValue", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeHashValue' field"))
	}
	m.TypeHashValue = typeHashValue

	size, err := ReadSimpleField(ctx, "size", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'size' field"))
	}
	m.Size = size

	offset, err := ReadSimpleField(ctx, "offset", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offset' field"))
	}
	m.Offset = offset

	dataType, err := ReadSimpleField(ctx, "dataType", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	flags, err := ReadSimpleField(ctx, "flags", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flags' field"))
	}
	m.Flags = flags

	propertyNameLength, err := ReadImplicitField[uint16](ctx, "propertyNameLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyNameLength' field"))
	}
	_ = propertyNameLength

	dataTypeNameLength, err := ReadImplicitField[uint16](ctx, "dataTypeNameLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeNameLength' field"))
	}
	_ = dataTypeNameLength

	commentLength, err := ReadImplicitField[uint16](ctx, "commentLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commentLength' field"))
	}
	_ = commentLength

	arrayDimensions, err := ReadSimpleField(ctx, "arrayDimensions", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensions' field"))
	}
	m.ArrayDimensions = arrayDimensions

	numChildren, err := ReadSimpleField(ctx, "numChildren", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numChildren' field"))
	}
	m.NumChildren = numChildren

	propertyName, err := ReadSimpleField(ctx, "propertyName", ReadString(readBuffer, uint32(int32(propertyNameLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyName' field"))
	}
	m.PropertyName = propertyName

	propertyNameTerminator, err := ReadConstField[uint8](ctx, "propertyNameTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsDataTypeTableChildEntry_PROPERTYNAMETERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyNameTerminator' field"))
	}
	_ = propertyNameTerminator

	dataTypeName, err := ReadSimpleField(ctx, "dataTypeName", ReadString(readBuffer, uint32(int32(dataTypeNameLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeName' field"))
	}
	m.DataTypeName = dataTypeName

	dataTypeNameTerminator, err := ReadConstField[uint8](ctx, "dataTypeNameTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsDataTypeTableChildEntry_DATATYPENAMETERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeNameTerminator' field"))
	}
	_ = dataTypeNameTerminator

	comment, err := ReadSimpleField(ctx, "comment", ReadString(readBuffer, uint32(int32(commentLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comment' field"))
	}
	m.Comment = comment

	commentTerminator, err := ReadConstField[uint8](ctx, "commentTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsDataTypeTableChildEntry_COMMENTTERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commentTerminator' field"))
	}
	_ = commentTerminator

	arrayInfo, err := ReadCountArrayField[AdsDataTypeArrayInfo](ctx, "arrayInfo", ReadComplex[AdsDataTypeArrayInfo](AdsDataTypeArrayInfoParseWithBuffer, readBuffer), uint64(arrayDimensions), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayInfo' field"))
	}
	m.ArrayInfo = arrayInfo

	children, err := ReadCountArrayField[AdsDataTypeTableEntry](ctx, "children", ReadComplex[AdsDataTypeTableEntry](AdsDataTypeTableEntryParseWithBuffer, readBuffer), uint64(numChildren), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'children' field"))
	}
	m.Children = children

	rest, err := readBuffer.ReadByteArray("rest", int(int32(entryLength)-int32((positionAware.GetPos()-startPos))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rest' field"))
	}
	m.Rest = rest

	if closeErr := readBuffer.CloseContext("AdsDataTypeTableChildEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDataTypeTableChildEntry")
	}

	return m, nil
}

func (m *_AdsDataTypeTableChildEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDataTypeTableChildEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDataTypeTableChildEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDataTypeTableChildEntry")
	}

	if err := WriteSimpleField[uint32](ctx, "entryLength", m.GetEntryLength(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'entryLength' field")
	}

	if err := WriteSimpleField[uint32](ctx, "version", m.GetVersion(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'version' field")
	}

	if err := WriteSimpleField[uint32](ctx, "hashValue", m.GetHashValue(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'hashValue' field")
	}

	if err := WriteSimpleField[uint32](ctx, "typeHashValue", m.GetTypeHashValue(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'typeHashValue' field")
	}

	if err := WriteSimpleField[uint32](ctx, "size", m.GetSize(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'size' field")
	}

	if err := WriteSimpleField[uint32](ctx, "offset", m.GetOffset(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'offset' field")
	}

	if err := WriteSimpleField[uint32](ctx, "dataType", m.GetDataType(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataType' field")
	}

	if err := WriteSimpleField[uint32](ctx, "flags", m.GetFlags(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flags' field")
	}
	propertyNameLength := uint16(uint16(len(m.GetPropertyName())))
	if err := WriteImplicitField(ctx, "propertyNameLength", propertyNameLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyNameLength' field")
	}
	dataTypeNameLength := uint16(uint16(len(m.GetDataTypeName())))
	if err := WriteImplicitField(ctx, "dataTypeNameLength", dataTypeNameLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeNameLength' field")
	}
	commentLength := uint16(uint16(len(m.GetComment())))
	if err := WriteImplicitField(ctx, "commentLength", commentLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commentLength' field")
	}

	if err := WriteSimpleField[uint16](ctx, "arrayDimensions", m.GetArrayDimensions(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayDimensions' field")
	}

	if err := WriteSimpleField[uint16](ctx, "numChildren", m.GetNumChildren(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'numChildren' field")
	}

	if err := WriteSimpleField[string](ctx, "propertyName", m.GetPropertyName(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetPropertyName())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyName' field")
	}

	if err := WriteConstField(ctx, "propertyNameTerminator", AdsDataTypeTableChildEntry_PROPERTYNAMETERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyNameTerminator' field")
	}

	if err := WriteSimpleField[string](ctx, "dataTypeName", m.GetDataTypeName(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetDataTypeName())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeName' field")
	}

	if err := WriteConstField(ctx, "dataTypeNameTerminator", AdsDataTypeTableChildEntry_DATATYPENAMETERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeNameTerminator' field")
	}

	if err := WriteSimpleField[string](ctx, "comment", m.GetComment(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetComment())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'comment' field")
	}

	if err := WriteConstField(ctx, "commentTerminator", AdsDataTypeTableChildEntry_COMMENTTERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commentTerminator' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "arrayInfo", m.GetArrayInfo(), writeBuffer, codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayInfo' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "children", m.GetChildren(), writeBuffer, codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'children' field")
	}

	if err := WriteByteArrayField(ctx, "rest", m.GetRest(), WriteByteArray(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'rest' field")
	}

	if popErr := writeBuffer.PopContext("AdsDataTypeTableChildEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDataTypeTableChildEntry")
	}
	return nil
}

func (m *_AdsDataTypeTableChildEntry) IsAdsDataTypeTableChildEntry() {}

func (m *_AdsDataTypeTableChildEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDataTypeTableChildEntry) deepCopy() *_AdsDataTypeTableChildEntry {
	if m == nil {
		return nil
	}
	_AdsDataTypeTableChildEntryCopy := &_AdsDataTypeTableChildEntry{
		m.EntryLength,
		m.Version,
		m.HashValue,
		m.TypeHashValue,
		m.Size,
		m.Offset,
		m.DataType,
		m.Flags,
		m.ArrayDimensions,
		m.NumChildren,
		m.PropertyName,
		m.DataTypeName,
		m.Comment,
		utils.DeepCopySlice[AdsDataTypeArrayInfo, AdsDataTypeArrayInfo](m.ArrayInfo),
		utils.DeepCopySlice[AdsDataTypeTableEntry, AdsDataTypeTableEntry](m.Children),
		utils.DeepCopySlice[byte, byte](m.Rest),
	}
	return _AdsDataTypeTableChildEntryCopy
}

func (m *_AdsDataTypeTableChildEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
