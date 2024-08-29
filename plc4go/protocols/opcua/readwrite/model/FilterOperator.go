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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FilterOperator is an enum
type FilterOperator uint32

type IFilterOperator interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	FilterOperator_filterOperatorEquals             FilterOperator = 0
	FilterOperator_filterOperatorIsNull             FilterOperator = 1
	FilterOperator_filterOperatorGreaterThan        FilterOperator = 2
	FilterOperator_filterOperatorLessThan           FilterOperator = 3
	FilterOperator_filterOperatorGreaterThanOrEqual FilterOperator = 4
	FilterOperator_filterOperatorLessThanOrEqual    FilterOperator = 5
	FilterOperator_filterOperatorLike               FilterOperator = 6
	FilterOperator_filterOperatorNot                FilterOperator = 7
	FilterOperator_filterOperatorBetween            FilterOperator = 8
	FilterOperator_filterOperatorInList             FilterOperator = 9
	FilterOperator_filterOperatorAnd                FilterOperator = 10
	FilterOperator_filterOperatorOr                 FilterOperator = 11
	FilterOperator_filterOperatorCast               FilterOperator = 12
	FilterOperator_filterOperatorInView             FilterOperator = 13
	FilterOperator_filterOperatorOfType             FilterOperator = 14
	FilterOperator_filterOperatorRelatedTo          FilterOperator = 15
	FilterOperator_filterOperatorBitwiseAnd         FilterOperator = 16
	FilterOperator_filterOperatorBitwiseOr          FilterOperator = 17
)

var FilterOperatorValues []FilterOperator

func init() {
	_ = errors.New
	FilterOperatorValues = []FilterOperator{
		FilterOperator_filterOperatorEquals,
		FilterOperator_filterOperatorIsNull,
		FilterOperator_filterOperatorGreaterThan,
		FilterOperator_filterOperatorLessThan,
		FilterOperator_filterOperatorGreaterThanOrEqual,
		FilterOperator_filterOperatorLessThanOrEqual,
		FilterOperator_filterOperatorLike,
		FilterOperator_filterOperatorNot,
		FilterOperator_filterOperatorBetween,
		FilterOperator_filterOperatorInList,
		FilterOperator_filterOperatorAnd,
		FilterOperator_filterOperatorOr,
		FilterOperator_filterOperatorCast,
		FilterOperator_filterOperatorInView,
		FilterOperator_filterOperatorOfType,
		FilterOperator_filterOperatorRelatedTo,
		FilterOperator_filterOperatorBitwiseAnd,
		FilterOperator_filterOperatorBitwiseOr,
	}
}

func FilterOperatorByValue(value uint32) (enum FilterOperator, ok bool) {
	switch value {
	case 0:
		return FilterOperator_filterOperatorEquals, true
	case 1:
		return FilterOperator_filterOperatorIsNull, true
	case 10:
		return FilterOperator_filterOperatorAnd, true
	case 11:
		return FilterOperator_filterOperatorOr, true
	case 12:
		return FilterOperator_filterOperatorCast, true
	case 13:
		return FilterOperator_filterOperatorInView, true
	case 14:
		return FilterOperator_filterOperatorOfType, true
	case 15:
		return FilterOperator_filterOperatorRelatedTo, true
	case 16:
		return FilterOperator_filterOperatorBitwiseAnd, true
	case 17:
		return FilterOperator_filterOperatorBitwiseOr, true
	case 2:
		return FilterOperator_filterOperatorGreaterThan, true
	case 3:
		return FilterOperator_filterOperatorLessThan, true
	case 4:
		return FilterOperator_filterOperatorGreaterThanOrEqual, true
	case 5:
		return FilterOperator_filterOperatorLessThanOrEqual, true
	case 6:
		return FilterOperator_filterOperatorLike, true
	case 7:
		return FilterOperator_filterOperatorNot, true
	case 8:
		return FilterOperator_filterOperatorBetween, true
	case 9:
		return FilterOperator_filterOperatorInList, true
	}
	return 0, false
}

func FilterOperatorByName(value string) (enum FilterOperator, ok bool) {
	switch value {
	case "filterOperatorEquals":
		return FilterOperator_filterOperatorEquals, true
	case "filterOperatorIsNull":
		return FilterOperator_filterOperatorIsNull, true
	case "filterOperatorAnd":
		return FilterOperator_filterOperatorAnd, true
	case "filterOperatorOr":
		return FilterOperator_filterOperatorOr, true
	case "filterOperatorCast":
		return FilterOperator_filterOperatorCast, true
	case "filterOperatorInView":
		return FilterOperator_filterOperatorInView, true
	case "filterOperatorOfType":
		return FilterOperator_filterOperatorOfType, true
	case "filterOperatorRelatedTo":
		return FilterOperator_filterOperatorRelatedTo, true
	case "filterOperatorBitwiseAnd":
		return FilterOperator_filterOperatorBitwiseAnd, true
	case "filterOperatorBitwiseOr":
		return FilterOperator_filterOperatorBitwiseOr, true
	case "filterOperatorGreaterThan":
		return FilterOperator_filterOperatorGreaterThan, true
	case "filterOperatorLessThan":
		return FilterOperator_filterOperatorLessThan, true
	case "filterOperatorGreaterThanOrEqual":
		return FilterOperator_filterOperatorGreaterThanOrEqual, true
	case "filterOperatorLessThanOrEqual":
		return FilterOperator_filterOperatorLessThanOrEqual, true
	case "filterOperatorLike":
		return FilterOperator_filterOperatorLike, true
	case "filterOperatorNot":
		return FilterOperator_filterOperatorNot, true
	case "filterOperatorBetween":
		return FilterOperator_filterOperatorBetween, true
	case "filterOperatorInList":
		return FilterOperator_filterOperatorInList, true
	}
	return 0, false
}

func FilterOperatorKnows(value uint32) bool {
	for _, typeValue := range FilterOperatorValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastFilterOperator(structType any) FilterOperator {
	castFunc := func(typ any) FilterOperator {
		if sFilterOperator, ok := typ.(FilterOperator); ok {
			return sFilterOperator
		}
		return 0
	}
	return castFunc(structType)
}

func (m FilterOperator) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m FilterOperator) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FilterOperatorParse(ctx context.Context, theBytes []byte) (FilterOperator, error) {
	return FilterOperatorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func FilterOperatorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (FilterOperator, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("FilterOperator", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading FilterOperator")
	}
	if enum, ok := FilterOperatorByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for FilterOperator")
		return FilterOperator(val), nil
	} else {
		return enum, nil
	}
}

func (e FilterOperator) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e FilterOperator) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("FilterOperator", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e FilterOperator) PLC4XEnumName() string {
	switch e {
	case FilterOperator_filterOperatorEquals:
		return "filterOperatorEquals"
	case FilterOperator_filterOperatorIsNull:
		return "filterOperatorIsNull"
	case FilterOperator_filterOperatorAnd:
		return "filterOperatorAnd"
	case FilterOperator_filterOperatorOr:
		return "filterOperatorOr"
	case FilterOperator_filterOperatorCast:
		return "filterOperatorCast"
	case FilterOperator_filterOperatorInView:
		return "filterOperatorInView"
	case FilterOperator_filterOperatorOfType:
		return "filterOperatorOfType"
	case FilterOperator_filterOperatorRelatedTo:
		return "filterOperatorRelatedTo"
	case FilterOperator_filterOperatorBitwiseAnd:
		return "filterOperatorBitwiseAnd"
	case FilterOperator_filterOperatorBitwiseOr:
		return "filterOperatorBitwiseOr"
	case FilterOperator_filterOperatorGreaterThan:
		return "filterOperatorGreaterThan"
	case FilterOperator_filterOperatorLessThan:
		return "filterOperatorLessThan"
	case FilterOperator_filterOperatorGreaterThanOrEqual:
		return "filterOperatorGreaterThanOrEqual"
	case FilterOperator_filterOperatorLessThanOrEqual:
		return "filterOperatorLessThanOrEqual"
	case FilterOperator_filterOperatorLike:
		return "filterOperatorLike"
	case FilterOperator_filterOperatorNot:
		return "filterOperatorNot"
	case FilterOperator_filterOperatorBetween:
		return "filterOperatorBetween"
	case FilterOperator_filterOperatorInList:
		return "filterOperatorInList"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e FilterOperator) String() string {
	return e.PLC4XEnumName()
}
