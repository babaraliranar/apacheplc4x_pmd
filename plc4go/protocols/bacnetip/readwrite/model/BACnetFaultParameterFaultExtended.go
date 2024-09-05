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

// BACnetFaultParameterFaultExtended is the corresponding interface of BACnetFaultParameterFaultExtended
type BACnetFaultParameterFaultExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedFaultType returns ExtendedFaultType (property field)
	GetExtendedFaultType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetFaultParameterFaultExtendedParameters
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultExtended is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtended()
}

// _BACnetFaultParameterFaultExtended is the data-structure of this message
type _BACnetFaultParameterFaultExtended struct {
	BACnetFaultParameterContract
	OpeningTag        BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedFaultType BACnetContextTagUnsignedInteger
	Parameters        BACnetFaultParameterFaultExtendedParameters
	ClosingTag        BACnetClosingTag
}

var _ BACnetFaultParameterFaultExtended = (*_BACnetFaultParameterFaultExtended)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultExtended)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtended) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtended) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetFaultParameterFaultExtended) GetExtendedFaultType() BACnetContextTagUnsignedInteger {
	return m.ExtendedFaultType
}

func (m *_BACnetFaultParameterFaultExtended) GetParameters() BACnetFaultParameterFaultExtendedParameters {
	return m.Parameters
}

func (m *_BACnetFaultParameterFaultExtended) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtended factory function for _BACnetFaultParameterFaultExtended
func NewBACnetFaultParameterFaultExtended(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedFaultType BACnetContextTagUnsignedInteger, parameters BACnetFaultParameterFaultExtendedParameters, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtended {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultExtended must not be nil")
	}
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for BACnetFaultParameterFaultExtended must not be nil")
	}
	if extendedFaultType == nil {
		panic("extendedFaultType of type BACnetContextTagUnsignedInteger for BACnetFaultParameterFaultExtended must not be nil")
	}
	if parameters == nil {
		panic("parameters of type BACnetFaultParameterFaultExtendedParameters for BACnetFaultParameterFaultExtended must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultExtended must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtended{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		VendorId:                     vendorId,
		ExtendedFaultType:            extendedFaultType,
		Parameters:                   parameters,
		ClosingTag:                   closingTag,
	}
	_result.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtended(structType any) BACnetFaultParameterFaultExtended {
	if casted, ok := structType.(BACnetFaultParameterFaultExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtended) GetTypeName() string {
	return "BACnetFaultParameterFaultExtended"
}

func (m *_BACnetFaultParameterFaultExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedFaultType)
	lengthInBits += m.ExtendedFaultType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultExtended BACnetFaultParameterFaultExtended, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	extendedFaultType, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedFaultType", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedFaultType' field"))
	}
	m.ExtendedFaultType = extendedFaultType

	parameters, err := ReadSimpleField[BACnetFaultParameterFaultExtendedParameters](ctx, "parameters", ReadComplex[BACnetFaultParameterFaultExtendedParameters](BACnetFaultParameterFaultExtendedParametersParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}
	m.Parameters = parameters

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtended")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtended")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedFaultType", m.GetExtendedFaultType(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedFaultType' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultExtendedParameters](ctx, "parameters", m.GetParameters(), WriteComplex[BACnetFaultParameterFaultExtendedParameters](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameters' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtended")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtended) IsBACnetFaultParameterFaultExtended() {}

func (m *_BACnetFaultParameterFaultExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
