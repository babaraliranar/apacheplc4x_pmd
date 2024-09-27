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

// BACnetConstructedDataBACnetIPUDPPort is the corresponding interface of BACnetConstructedDataBACnetIPUDPPort
type BACnetConstructedDataBACnetIPUDPPort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpUdpPort returns IpUdpPort (property field)
	GetIpUdpPort() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataBACnetIPUDPPort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPUDPPort()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPUDPPortBuilder
	CreateBACnetConstructedDataBACnetIPUDPPortBuilder() BACnetConstructedDataBACnetIPUDPPortBuilder
}

// _BACnetConstructedDataBACnetIPUDPPort is the data-structure of this message
type _BACnetConstructedDataBACnetIPUDPPort struct {
	BACnetConstructedDataContract
	IpUdpPort BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataBACnetIPUDPPort = (*_BACnetConstructedDataBACnetIPUDPPort)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPUDPPort)(nil)

// NewBACnetConstructedDataBACnetIPUDPPort factory function for _BACnetConstructedDataBACnetIPUDPPort
func NewBACnetConstructedDataBACnetIPUDPPort(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipUdpPort BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPUDPPort {
	if ipUdpPort == nil {
		panic("ipUdpPort of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataBACnetIPUDPPort must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPUDPPort{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IpUdpPort:                     ipUdpPort,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPUDPPortBuilder is a builder for BACnetConstructedDataBACnetIPUDPPort
type BACnetConstructedDataBACnetIPUDPPortBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipUdpPort BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPUDPPortBuilder
	// WithIpUdpPort adds IpUdpPort (property field)
	WithIpUdpPort(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPUDPPortBuilder
	// WithIpUdpPortBuilder adds IpUdpPort (property field) which is build by the builder
	WithIpUdpPortBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBACnetIPUDPPortBuilder
	// Build builds the BACnetConstructedDataBACnetIPUDPPort or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPUDPPort, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPUDPPort
}

// NewBACnetConstructedDataBACnetIPUDPPortBuilder() creates a BACnetConstructedDataBACnetIPUDPPortBuilder
func NewBACnetConstructedDataBACnetIPUDPPortBuilder() BACnetConstructedDataBACnetIPUDPPortBuilder {
	return &_BACnetConstructedDataBACnetIPUDPPortBuilder{_BACnetConstructedDataBACnetIPUDPPort: new(_BACnetConstructedDataBACnetIPUDPPort)}
}

type _BACnetConstructedDataBACnetIPUDPPortBuilder struct {
	*_BACnetConstructedDataBACnetIPUDPPort

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPUDPPortBuilder) = (*_BACnetConstructedDataBACnetIPUDPPortBuilder)(nil)

func (m *_BACnetConstructedDataBACnetIPUDPPortBuilder) WithMandatoryFields(ipUdpPort BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPUDPPortBuilder {
	return m.WithIpUdpPort(ipUdpPort)
}

func (m *_BACnetConstructedDataBACnetIPUDPPortBuilder) WithIpUdpPort(ipUdpPort BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPUDPPortBuilder {
	m.IpUdpPort = ipUdpPort
	return m
}

func (m *_BACnetConstructedDataBACnetIPUDPPortBuilder) WithIpUdpPortBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBACnetIPUDPPortBuilder {
	builder := builderSupplier(m.IpUdpPort.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.IpUdpPort, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataBACnetIPUDPPortBuilder) Build() (BACnetConstructedDataBACnetIPUDPPort, error) {
	if m.IpUdpPort == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'ipUdpPort' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataBACnetIPUDPPort.deepCopy(), nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPortBuilder) MustBuild() BACnetConstructedDataBACnetIPUDPPort {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataBACnetIPUDPPortBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataBACnetIPUDPPortBuilder()
}

// CreateBACnetConstructedDataBACnetIPUDPPortBuilder creates a BACnetConstructedDataBACnetIPUDPPortBuilder
func (m *_BACnetConstructedDataBACnetIPUDPPort) CreateBACnetConstructedDataBACnetIPUDPPortBuilder() BACnetConstructedDataBACnetIPUDPPortBuilder {
	if m == nil {
		return NewBACnetConstructedDataBACnetIPUDPPortBuilder()
	}
	return &_BACnetConstructedDataBACnetIPUDPPortBuilder{_BACnetConstructedDataBACnetIPUDPPort: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_UDP_PORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetIpUdpPort() BACnetApplicationTagUnsignedInteger {
	return m.IpUdpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpUdpPort())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPUDPPort(structType any) BACnetConstructedDataBACnetIPUDPPort {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPUDPPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPUDPPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPUDPPort"
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipUdpPort)
	lengthInBits += m.IpUdpPort.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPUDPPort BACnetConstructedDataBACnetIPUDPPort, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPUDPPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPUDPPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipUdpPort, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipUdpPort", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipUdpPort' field"))
	}
	m.IpUdpPort = ipUdpPort

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), ipUdpPort)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPUDPPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPUDPPort")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPUDPPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPUDPPort")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipUdpPort", m.GetIpUdpPort(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipUdpPort' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPUDPPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPUDPPort")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) IsBACnetConstructedDataBACnetIPUDPPort() {}

func (m *_BACnetConstructedDataBACnetIPUDPPort) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) deepCopy() *_BACnetConstructedDataBACnetIPUDPPort {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPUDPPortCopy := &_BACnetConstructedDataBACnetIPUDPPort{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.IpUdpPort.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPUDPPortCopy
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
