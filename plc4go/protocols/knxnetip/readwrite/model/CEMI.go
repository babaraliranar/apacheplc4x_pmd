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

// CEMI is the corresponding interface of CEMI
type CEMI interface {
	CEMIContract
	CEMIRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCEMI is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMI()
	// CreateBuilder creates a CEMIBuilder
	CreateCEMIBuilder() CEMIBuilder
}

// CEMIContract provides a set of functions which can be overwritten by a sub struct
type CEMIContract interface {
	// GetSize() returns a parser argument
	GetSize() uint16
	// IsCEMI is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMI()
	// CreateBuilder creates a CEMIBuilder
	CreateCEMIBuilder() CEMIBuilder
}

// CEMIRequirements provides a set of functions which need to be implemented by a sub struct
type CEMIRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMessageCode returns MessageCode (discriminator field)
	GetMessageCode() uint8
}

// _CEMI is the data-structure of this message
type _CEMI struct {
	_SubType interface {
		CEMIContract
		CEMIRequirements
	}

	// Arguments.
	Size uint16
}

var _ CEMIContract = (*_CEMI)(nil)

// NewCEMI factory function for _CEMI
func NewCEMI(size uint16) *_CEMI {
	return &_CEMI{Size: size}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CEMIBuilder is a builder for CEMI
type CEMIBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() CEMIBuilder
	// AsLBusmonInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLBusmonInd() interface {
		LBusmonIndBuilder
		Done() CEMIBuilder
	}
	// AsLDataReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLDataReq() interface {
		LDataReqBuilder
		Done() CEMIBuilder
	}
	// AsLDataInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLDataInd() interface {
		LDataIndBuilder
		Done() CEMIBuilder
	}
	// AsLDataCon converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLDataCon() interface {
		LDataConBuilder
		Done() CEMIBuilder
	}
	// AsLRawReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLRawReq() interface {
		LRawReqBuilder
		Done() CEMIBuilder
	}
	// AsLRawInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLRawInd() interface {
		LRawIndBuilder
		Done() CEMIBuilder
	}
	// AsLRawCon converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLRawCon() interface {
		LRawConBuilder
		Done() CEMIBuilder
	}
	// AsLPollDataReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLPollDataReq() interface {
		LPollDataReqBuilder
		Done() CEMIBuilder
	}
	// AsLPollDataCon converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsLPollDataCon() interface {
		LPollDataConBuilder
		Done() CEMIBuilder
	}
	// AsTDataConnectedReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsTDataConnectedReq() interface {
		TDataConnectedReqBuilder
		Done() CEMIBuilder
	}
	// AsTDataConnectedInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsTDataConnectedInd() interface {
		TDataConnectedIndBuilder
		Done() CEMIBuilder
	}
	// AsTDataIndividualReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsTDataIndividualReq() interface {
		TDataIndividualReqBuilder
		Done() CEMIBuilder
	}
	// AsTDataIndividualInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsTDataIndividualInd() interface {
		TDataIndividualIndBuilder
		Done() CEMIBuilder
	}
	// AsMPropReadReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMPropReadReq() interface {
		MPropReadReqBuilder
		Done() CEMIBuilder
	}
	// AsMPropReadCon converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMPropReadCon() interface {
		MPropReadConBuilder
		Done() CEMIBuilder
	}
	// AsMPropWriteReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMPropWriteReq() interface {
		MPropWriteReqBuilder
		Done() CEMIBuilder
	}
	// AsMPropWriteCon converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMPropWriteCon() interface {
		MPropWriteConBuilder
		Done() CEMIBuilder
	}
	// AsMPropInfoInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMPropInfoInd() interface {
		MPropInfoIndBuilder
		Done() CEMIBuilder
	}
	// AsMFuncPropCommandReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMFuncPropCommandReq() interface {
		MFuncPropCommandReqBuilder
		Done() CEMIBuilder
	}
	// AsMFuncPropStateReadReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMFuncPropStateReadReq() interface {
		MFuncPropStateReadReqBuilder
		Done() CEMIBuilder
	}
	// AsMFuncPropCon converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMFuncPropCon() interface {
		MFuncPropConBuilder
		Done() CEMIBuilder
	}
	// AsMResetReq converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMResetReq() interface {
		MResetReqBuilder
		Done() CEMIBuilder
	}
	// AsMResetInd converts this build to a subType of CEMI. It is always possible to return to current builder using Done()
	AsMResetInd() interface {
		MResetIndBuilder
		Done() CEMIBuilder
	}
	// Build builds the CEMI or returns an error if something is wrong
	PartialBuild() (CEMIContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() CEMIContract
	// Build builds the CEMI or returns an error if something is wrong
	Build() (CEMI, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CEMI
}

// NewCEMIBuilder() creates a CEMIBuilder
func NewCEMIBuilder() CEMIBuilder {
	return &_CEMIBuilder{_CEMI: new(_CEMI)}
}

type _CEMIChildBuilder interface {
	utils.Copyable
	setParent(CEMIContract)
	buildForCEMI() (CEMI, error)
}

type _CEMIBuilder struct {
	*_CEMI

	childBuilder _CEMIChildBuilder

	err *utils.MultiError
}

var _ (CEMIBuilder) = (*_CEMIBuilder)(nil)

func (b *_CEMIBuilder) WithMandatoryFields() CEMIBuilder {
	return b
}

func (b *_CEMIBuilder) PartialBuild() (CEMIContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CEMI.deepCopy(), nil
}

func (b *_CEMIBuilder) PartialMustBuild() CEMIContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CEMIBuilder) AsLBusmonInd() interface {
	LBusmonIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LBusmonIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLBusmonIndBuilder().(*_LBusmonIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLDataReq() interface {
	LDataReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LDataReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLDataReqBuilder().(*_LDataReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLDataInd() interface {
	LDataIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LDataIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLDataIndBuilder().(*_LDataIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLDataCon() interface {
	LDataConBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LDataConBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLDataConBuilder().(*_LDataConBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLRawReq() interface {
	LRawReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LRawReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLRawReqBuilder().(*_LRawReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLRawInd() interface {
	LRawIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LRawIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLRawIndBuilder().(*_LRawIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLRawCon() interface {
	LRawConBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LRawConBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLRawConBuilder().(*_LRawConBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLPollDataReq() interface {
	LPollDataReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LPollDataReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLPollDataReqBuilder().(*_LPollDataReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsLPollDataCon() interface {
	LPollDataConBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		LPollDataConBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewLPollDataConBuilder().(*_LPollDataConBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsTDataConnectedReq() interface {
	TDataConnectedReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		TDataConnectedReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewTDataConnectedReqBuilder().(*_TDataConnectedReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsTDataConnectedInd() interface {
	TDataConnectedIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		TDataConnectedIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewTDataConnectedIndBuilder().(*_TDataConnectedIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsTDataIndividualReq() interface {
	TDataIndividualReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		TDataIndividualReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewTDataIndividualReqBuilder().(*_TDataIndividualReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsTDataIndividualInd() interface {
	TDataIndividualIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		TDataIndividualIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewTDataIndividualIndBuilder().(*_TDataIndividualIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMPropReadReq() interface {
	MPropReadReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MPropReadReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMPropReadReqBuilder().(*_MPropReadReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMPropReadCon() interface {
	MPropReadConBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MPropReadConBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMPropReadConBuilder().(*_MPropReadConBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMPropWriteReq() interface {
	MPropWriteReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MPropWriteReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMPropWriteReqBuilder().(*_MPropWriteReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMPropWriteCon() interface {
	MPropWriteConBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MPropWriteConBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMPropWriteConBuilder().(*_MPropWriteConBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMPropInfoInd() interface {
	MPropInfoIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MPropInfoIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMPropInfoIndBuilder().(*_MPropInfoIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMFuncPropCommandReq() interface {
	MFuncPropCommandReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MFuncPropCommandReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMFuncPropCommandReqBuilder().(*_MFuncPropCommandReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMFuncPropStateReadReq() interface {
	MFuncPropStateReadReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MFuncPropStateReadReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMFuncPropStateReadReqBuilder().(*_MFuncPropStateReadReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMFuncPropCon() interface {
	MFuncPropConBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MFuncPropConBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMFuncPropConBuilder().(*_MFuncPropConBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMResetReq() interface {
	MResetReqBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MResetReqBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMResetReqBuilder().(*_MResetReqBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) AsMResetInd() interface {
	MResetIndBuilder
	Done() CEMIBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MResetIndBuilder
		Done() CEMIBuilder
	}); ok {
		return cb
	}
	cb := NewMResetIndBuilder().(*_MResetIndBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CEMIBuilder) Build() (CEMI, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCEMI()
}

func (b *_CEMIBuilder) MustBuild() CEMI {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CEMIBuilder) DeepCopy() any {
	_copy := b.CreateCEMIBuilder().(*_CEMIBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_CEMIChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCEMIBuilder creates a CEMIBuilder
func (b *_CEMI) CreateCEMIBuilder() CEMIBuilder {
	if b == nil {
		return NewCEMIBuilder()
	}
	return &_CEMIBuilder{_CEMI: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCEMI(structType any) CEMI {
	if casted, ok := structType.(CEMI); ok {
		return casted
	}
	if casted, ok := structType.(*CEMI); ok {
		return *casted
	}
	return nil
}

func (m *_CEMI) GetTypeName() string {
	return "CEMI"
}

func (m *_CEMI) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CEMI) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_CEMI) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CEMIParse[T CEMI](ctx context.Context, theBytes []byte, size uint16) (T, error) {
	return CEMIParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func CEMIParseWithBufferProducer[T CEMI](size uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CEMIParseWithBuffer[T](ctx, readBuffer, size)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CEMIParseWithBuffer[T CEMI](ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (T, error) {
	v, err := (&_CEMI{Size: size}).parse(ctx, readBuffer, size)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_CEMI) parse(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (__cEMI CEMI, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMI"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMI")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageCode, err := ReadDiscriminatorField[uint8](ctx, "messageCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageCode' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CEMI
	switch {
	case messageCode == 0x2B: // LBusmonInd
		if _child, err = new(_LBusmonInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LBusmonInd for type-switch of CEMI")
		}
	case messageCode == 0x11: // LDataReq
		if _child, err = new(_LDataReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataReq for type-switch of CEMI")
		}
	case messageCode == 0x29: // LDataInd
		if _child, err = new(_LDataInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataInd for type-switch of CEMI")
		}
	case messageCode == 0x2E: // LDataCon
		if _child, err = new(_LDataCon).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataCon for type-switch of CEMI")
		}
	case messageCode == 0x10: // LRawReq
		if _child, err = new(_LRawReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LRawReq for type-switch of CEMI")
		}
	case messageCode == 0x2D: // LRawInd
		if _child, err = new(_LRawInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LRawInd for type-switch of CEMI")
		}
	case messageCode == 0x2F: // LRawCon
		if _child, err = new(_LRawCon).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LRawCon for type-switch of CEMI")
		}
	case messageCode == 0x13: // LPollDataReq
		if _child, err = new(_LPollDataReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LPollDataReq for type-switch of CEMI")
		}
	case messageCode == 0x25: // LPollDataCon
		if _child, err = new(_LPollDataCon).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LPollDataCon for type-switch of CEMI")
		}
	case messageCode == 0x41: // TDataConnectedReq
		if _child, err = new(_TDataConnectedReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataConnectedReq for type-switch of CEMI")
		}
	case messageCode == 0x89: // TDataConnectedInd
		if _child, err = new(_TDataConnectedInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataConnectedInd for type-switch of CEMI")
		}
	case messageCode == 0x4A: // TDataIndividualReq
		if _child, err = new(_TDataIndividualReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataIndividualReq for type-switch of CEMI")
		}
	case messageCode == 0x94: // TDataIndividualInd
		if _child, err = new(_TDataIndividualInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataIndividualInd for type-switch of CEMI")
		}
	case messageCode == 0xFC: // MPropReadReq
		if _child, err = new(_MPropReadReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropReadReq for type-switch of CEMI")
		}
	case messageCode == 0xFB: // MPropReadCon
		if _child, err = new(_MPropReadCon).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropReadCon for type-switch of CEMI")
		}
	case messageCode == 0xF6: // MPropWriteReq
		if _child, err = new(_MPropWriteReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropWriteReq for type-switch of CEMI")
		}
	case messageCode == 0xF5: // MPropWriteCon
		if _child, err = new(_MPropWriteCon).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropWriteCon for type-switch of CEMI")
		}
	case messageCode == 0xF7: // MPropInfoInd
		if _child, err = new(_MPropInfoInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropInfoInd for type-switch of CEMI")
		}
	case messageCode == 0xF8: // MFuncPropCommandReq
		if _child, err = new(_MFuncPropCommandReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MFuncPropCommandReq for type-switch of CEMI")
		}
	case messageCode == 0xF9: // MFuncPropStateReadReq
		if _child, err = new(_MFuncPropStateReadReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MFuncPropStateReadReq for type-switch of CEMI")
		}
	case messageCode == 0xFA: // MFuncPropCon
		if _child, err = new(_MFuncPropCon).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MFuncPropCon for type-switch of CEMI")
		}
	case messageCode == 0xF1: // MResetReq
		if _child, err = new(_MResetReq).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MResetReq for type-switch of CEMI")
		}
	case messageCode == 0xF0: // MResetInd
		if _child, err = new(_MResetInd).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MResetInd for type-switch of CEMI")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageCode=%v]", messageCode)
	}

	if closeErr := readBuffer.CloseContext("CEMI"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMI")
	}

	return _child, nil
}

func (pm *_CEMI) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CEMI, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CEMI"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CEMI")
	}

	if err := WriteDiscriminatorField(ctx, "messageCode", m.GetMessageCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMI"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CEMI")
	}
	return nil
}

////
// Arguments Getter

func (m *_CEMI) GetSize() uint16 {
	return m.Size
}

//
////

func (m *_CEMI) IsCEMI() {}

func (m *_CEMI) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CEMI) deepCopy() *_CEMI {
	if m == nil {
		return nil
	}
	_CEMICopy := &_CEMI{
		nil, // will be set by child
		m.Size,
	}
	return _CEMICopy
}
