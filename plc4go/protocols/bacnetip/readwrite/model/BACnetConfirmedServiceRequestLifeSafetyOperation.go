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

// BACnetConfirmedServiceRequestLifeSafetyOperation is the corresponding interface of BACnetConfirmedServiceRequestLifeSafetyOperation
type BACnetConfirmedServiceRequestLifeSafetyOperation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetRequestingProcessIdentifier returns RequestingProcessIdentifier (property field)
	GetRequestingProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetRequestingSource returns RequestingSource (property field)
	GetRequestingSource() BACnetContextTagCharacterString
	// GetRequest returns Request (property field)
	GetRequest() BACnetLifeSafetyOperationTagged
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// IsBACnetConfirmedServiceRequestLifeSafetyOperation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestLifeSafetyOperation()
}

// _BACnetConfirmedServiceRequestLifeSafetyOperation is the data-structure of this message
type _BACnetConfirmedServiceRequestLifeSafetyOperation struct {
	BACnetConfirmedServiceRequestContract
	RequestingProcessIdentifier BACnetContextTagUnsignedInteger
	RequestingSource            BACnetContextTagCharacterString
	Request                     BACnetLifeSafetyOperationTagged
	ObjectIdentifier            BACnetContextTagObjectIdentifier
}

var _ BACnetConfirmedServiceRequestLifeSafetyOperation = (*_BACnetConfirmedServiceRequestLifeSafetyOperation)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestLifeSafetyOperation)(nil)

// NewBACnetConfirmedServiceRequestLifeSafetyOperation factory function for _BACnetConfirmedServiceRequestLifeSafetyOperation
func NewBACnetConfirmedServiceRequestLifeSafetyOperation(requestingProcessIdentifier BACnetContextTagUnsignedInteger, requestingSource BACnetContextTagCharacterString, request BACnetLifeSafetyOperationTagged, objectIdentifier BACnetContextTagObjectIdentifier, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestLifeSafetyOperation {
	if requestingProcessIdentifier == nil {
		panic("requestingProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestLifeSafetyOperation must not be nil")
	}
	if requestingSource == nil {
		panic("requestingSource of type BACnetContextTagCharacterString for BACnetConfirmedServiceRequestLifeSafetyOperation must not be nil")
	}
	if request == nil {
		panic("request of type BACnetLifeSafetyOperationTagged for BACnetConfirmedServiceRequestLifeSafetyOperation must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestLifeSafetyOperation{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		RequestingProcessIdentifier:           requestingProcessIdentifier,
		RequestingSource:                      requestingSource,
		Request:                               request,
		ObjectIdentifier:                      objectIdentifier,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetRequestingProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.RequestingProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetRequestingSource() BACnetContextTagCharacterString {
	return m.RequestingSource
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetRequest() BACnetLifeSafetyOperationTagged {
	return m.Request
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestLifeSafetyOperation(structType any) BACnetConfirmedServiceRequestLifeSafetyOperation {
	if casted, ok := structType.(BACnetConfirmedServiceRequestLifeSafetyOperation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestLifeSafetyOperation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetTypeName() string {
	return "BACnetConfirmedServiceRequestLifeSafetyOperation"
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (requestingProcessIdentifier)
	lengthInBits += m.RequestingProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (requestingSource)
	lengthInBits += m.RequestingSource.GetLengthInBits(ctx)

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits(ctx)

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestLifeSafetyOperation BACnetConfirmedServiceRequestLifeSafetyOperation, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestLifeSafetyOperation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestingProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "requestingProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestingProcessIdentifier' field"))
	}
	m.RequestingProcessIdentifier = requestingProcessIdentifier

	requestingSource, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "requestingSource", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestingSource' field"))
	}
	m.RequestingSource = requestingSource

	request, err := ReadSimpleField[BACnetLifeSafetyOperationTagged](ctx, "request", ReadComplex[BACnetLifeSafetyOperationTagged](BACnetLifeSafetyOperationTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'request' field"))
	}
	m.Request = request

	var objectIdentifier BACnetContextTagObjectIdentifier
	_objectIdentifier, err := ReadOptionalField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	if _objectIdentifier != nil {
		objectIdentifier = *_objectIdentifier
		m.ObjectIdentifier = objectIdentifier
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestLifeSafetyOperation")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestLifeSafetyOperation")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "requestingProcessIdentifier", m.GetRequestingProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestingProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "requestingSource", m.GetRequestingSource(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestingSource' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyOperationTagged](ctx, "request", m.GetRequest(), WriteComplex[BACnetLifeSafetyOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'request' field")
		}

		if err := WriteOptionalField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", GetRef(m.GetObjectIdentifier()), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestLifeSafetyOperation")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) IsBACnetConfirmedServiceRequestLifeSafetyOperation() {
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) deepCopy() *_BACnetConfirmedServiceRequestLifeSafetyOperation {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestLifeSafetyOperationCopy := &_BACnetConfirmedServiceRequestLifeSafetyOperation{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.RequestingProcessIdentifier.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.RequestingSource.DeepCopy().(BACnetContextTagCharacterString),
		m.Request.DeepCopy().(BACnetLifeSafetyOperationTagged),
		m.ObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestLifeSafetyOperationCopy
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
