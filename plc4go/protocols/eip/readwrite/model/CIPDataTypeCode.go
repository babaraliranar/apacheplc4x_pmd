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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CIPDataTypeCode is an enum
type CIPDataTypeCode uint16

type ICIPDataTypeCode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Size() uint8
}

const (
	CIPDataTypeCode_BOOL          CIPDataTypeCode = 0x00C1
	CIPDataTypeCode_SINT          CIPDataTypeCode = 0x00C2
	CIPDataTypeCode_INT           CIPDataTypeCode = 0x00C3
	CIPDataTypeCode_DINT          CIPDataTypeCode = 0x00C4
	CIPDataTypeCode_LINT          CIPDataTypeCode = 0x00C5
	CIPDataTypeCode_USINT         CIPDataTypeCode = 0x00C6
	CIPDataTypeCode_UINT          CIPDataTypeCode = 0x00C7
	CIPDataTypeCode_UDINT         CIPDataTypeCode = 0x00C8
	CIPDataTypeCode_ULINT         CIPDataTypeCode = 0x00C9
	CIPDataTypeCode_REAL          CIPDataTypeCode = 0x00CA
	CIPDataTypeCode_LREAL         CIPDataTypeCode = 0x00CB
	CIPDataTypeCode_STIME         CIPDataTypeCode = 0x00CC
	CIPDataTypeCode_DATE          CIPDataTypeCode = 0x00CD
	CIPDataTypeCode_TIME_OF_DAY   CIPDataTypeCode = 0x00CE
	CIPDataTypeCode_DATE_AND_TIME CIPDataTypeCode = 0x00CF
	CIPDataTypeCode_STRING        CIPDataTypeCode = 0x00D0
	CIPDataTypeCode_BYTE          CIPDataTypeCode = 0x00D1
	CIPDataTypeCode_WORD          CIPDataTypeCode = 0x00D2
	CIPDataTypeCode_DWORD         CIPDataTypeCode = 0x00D3
	CIPDataTypeCode_LWORD         CIPDataTypeCode = 0x00D3
	CIPDataTypeCode_STRING2       CIPDataTypeCode = 0x00D5
	CIPDataTypeCode_FTIME         CIPDataTypeCode = 0x00D6
	CIPDataTypeCode_LTIME         CIPDataTypeCode = 0x00D7
	CIPDataTypeCode_ITIME         CIPDataTypeCode = 0x00D8
	CIPDataTypeCode_STRINGN       CIPDataTypeCode = 0x00D9
	CIPDataTypeCode_SHORT_STRING  CIPDataTypeCode = 0x00DA
	CIPDataTypeCode_TIME          CIPDataTypeCode = 0x00DB
	CIPDataTypeCode_EPATH         CIPDataTypeCode = 0x00DC
	CIPDataTypeCode_ENGUNIT       CIPDataTypeCode = 0x00DD
	CIPDataTypeCode_STRINGI       CIPDataTypeCode = 0x00DD
	CIPDataTypeCode_STRUCTURED    CIPDataTypeCode = 0x02A0
)

var CIPDataTypeCodeValues []CIPDataTypeCode

func init() {
	_ = errors.New
	CIPDataTypeCodeValues = []CIPDataTypeCode{
		CIPDataTypeCode_BOOL,
		CIPDataTypeCode_SINT,
		CIPDataTypeCode_INT,
		CIPDataTypeCode_DINT,
		CIPDataTypeCode_LINT,
		CIPDataTypeCode_USINT,
		CIPDataTypeCode_UINT,
		CIPDataTypeCode_UDINT,
		CIPDataTypeCode_ULINT,
		CIPDataTypeCode_REAL,
		CIPDataTypeCode_LREAL,
		CIPDataTypeCode_STIME,
		CIPDataTypeCode_DATE,
		CIPDataTypeCode_TIME_OF_DAY,
		CIPDataTypeCode_DATE_AND_TIME,
		CIPDataTypeCode_STRING,
		CIPDataTypeCode_BYTE,
		CIPDataTypeCode_WORD,
		CIPDataTypeCode_DWORD,
		CIPDataTypeCode_LWORD,
		CIPDataTypeCode_STRING2,
		CIPDataTypeCode_FTIME,
		CIPDataTypeCode_LTIME,
		CIPDataTypeCode_ITIME,
		CIPDataTypeCode_STRINGN,
		CIPDataTypeCode_SHORT_STRING,
		CIPDataTypeCode_TIME,
		CIPDataTypeCode_EPATH,
		CIPDataTypeCode_ENGUNIT,
		CIPDataTypeCode_STRINGI,
		CIPDataTypeCode_STRUCTURED,
	}
}

func (e CIPDataTypeCode) Size() uint8 {
	switch e {
	case 0x00C1:
		{ /* '0X00C1' */
			return 1
		}
	case 0x00C2:
		{ /* '0X00C2' */
			return 1
		}
	case 0x00C3:
		{ /* '0X00C3' */
			return 2
		}
	case 0x00C4:
		{ /* '0X00C4' */
			return 4
		}
	case 0x00C5:
		{ /* '0X00C5' */
			return 8
		}
	case 0x00C6:
		{ /* '0X00C6' */
			return 1
		}
	case 0x00C7:
		{ /* '0X00C7' */
			return 2
		}
	case 0x00C8:
		{ /* '0X00C8' */
			return 4
		}
	case 0x00C9:
		{ /* '0X00C9' */
			return 8
		}
	case 0x00CA:
		{ /* '0X00CA' */
			return 4
		}
	case 0x00CB:
		{ /* '0X00CB' */
			return 8
		}
	case 0x00CC:
		{ /* '0X00CC' */
			return 4
		}
	case 0x00CD:
		{ /* '0X00CD' */
			return 2
		}
	case 0x00CE:
		{ /* '0X00CE' */
			return 4
		}
	case 0x00CF:
		{ /* '0X00CF' */
			return 6
		}
	case 0x00D0:
		{ /* '0X00D0' */
			return 0
		}
	case 0x00D1:
		{ /* '0X00D1' */
			return 1
		}
	case 0x00D2:
		{ /* '0X00D2' */
			return 2
		}
	case 0x00D3:
		{ /* '0X00D3' */
			return 4
		}
	case 0x00D5:
		{ /* '0X00D5' */
			return 0
		}
	case 0x00D6:
		{ /* '0X00D6' */
			return 4
		}
	case 0x00D7:
		{ /* '0X00D7' */
			return 8
		}
	case 0x00D8:
		{ /* '0X00D8' */
			return 2
		}
	case 0x00D9:
		{ /* '0X00D9' */
			return 0
		}
	case 0x00DA:
		{ /* '0X00DA' */
			return 0
		}
	case 0x00DB:
		{ /* '0X00DB' */
			return 4
		}
	case 0x00DC:
		{ /* '0X00DC' */
			return 0
		}
	case 0x00DD:
		{ /* '0X00DD' */
			return 0
		}
	case 0x02A0:
		{ /* '0X02A0' */
			return 88
		}
	default:
		{
			return 0
		}
	}
}

func CIPDataTypeCodeFirstEnumForFieldSize(value uint8) (CIPDataTypeCode, error) {
	for _, sizeValue := range CIPDataTypeCodeValues {
		if sizeValue.Size() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing Size not found", value)
}
func CIPDataTypeCodeByValue(value uint16) (enum CIPDataTypeCode, ok bool) {
	switch value {
	case 0x00C1:
		return CIPDataTypeCode_BOOL, true
	case 0x00C2:
		return CIPDataTypeCode_SINT, true
	case 0x00C3:
		return CIPDataTypeCode_INT, true
	case 0x00C4:
		return CIPDataTypeCode_DINT, true
	case 0x00C5:
		return CIPDataTypeCode_LINT, true
	case 0x00C6:
		return CIPDataTypeCode_USINT, true
	case 0x00C7:
		return CIPDataTypeCode_UINT, true
	case 0x00C8:
		return CIPDataTypeCode_UDINT, true
	case 0x00C9:
		return CIPDataTypeCode_ULINT, true
	case 0x00CA:
		return CIPDataTypeCode_REAL, true
	case 0x00CB:
		return CIPDataTypeCode_LREAL, true
	case 0x00CC:
		return CIPDataTypeCode_STIME, true
	case 0x00CD:
		return CIPDataTypeCode_DATE, true
	case 0x00CE:
		return CIPDataTypeCode_TIME_OF_DAY, true
	case 0x00CF:
		return CIPDataTypeCode_DATE_AND_TIME, true
	case 0x00D0:
		return CIPDataTypeCode_STRING, true
	case 0x00D1:
		return CIPDataTypeCode_BYTE, true
	case 0x00D2:
		return CIPDataTypeCode_WORD, true
	case 0x00D3:
		return CIPDataTypeCode_DWORD, true
	case 0x00D5:
		return CIPDataTypeCode_STRING2, true
	case 0x00D6:
		return CIPDataTypeCode_FTIME, true
	case 0x00D7:
		return CIPDataTypeCode_LTIME, true
	case 0x00D8:
		return CIPDataTypeCode_ITIME, true
	case 0x00D9:
		return CIPDataTypeCode_STRINGN, true
	case 0x00DA:
		return CIPDataTypeCode_SHORT_STRING, true
	case 0x00DB:
		return CIPDataTypeCode_TIME, true
	case 0x00DC:
		return CIPDataTypeCode_EPATH, true
	case 0x00DD:
		return CIPDataTypeCode_ENGUNIT, true
	case 0x02A0:
		return CIPDataTypeCode_STRUCTURED, true
	}
	return 0, false
}

func CIPDataTypeCodeByName(value string) (enum CIPDataTypeCode, ok bool) {
	switch value {
	case "BOOL":
		return CIPDataTypeCode_BOOL, true
	case "SINT":
		return CIPDataTypeCode_SINT, true
	case "INT":
		return CIPDataTypeCode_INT, true
	case "DINT":
		return CIPDataTypeCode_DINT, true
	case "LINT":
		return CIPDataTypeCode_LINT, true
	case "USINT":
		return CIPDataTypeCode_USINT, true
	case "UINT":
		return CIPDataTypeCode_UINT, true
	case "UDINT":
		return CIPDataTypeCode_UDINT, true
	case "ULINT":
		return CIPDataTypeCode_ULINT, true
	case "REAL":
		return CIPDataTypeCode_REAL, true
	case "LREAL":
		return CIPDataTypeCode_LREAL, true
	case "STIME":
		return CIPDataTypeCode_STIME, true
	case "DATE":
		return CIPDataTypeCode_DATE, true
	case "TIME_OF_DAY":
		return CIPDataTypeCode_TIME_OF_DAY, true
	case "DATE_AND_TIME":
		return CIPDataTypeCode_DATE_AND_TIME, true
	case "STRING":
		return CIPDataTypeCode_STRING, true
	case "BYTE":
		return CIPDataTypeCode_BYTE, true
	case "WORD":
		return CIPDataTypeCode_WORD, true
	case "DWORD":
		return CIPDataTypeCode_DWORD, true
	case "STRING2":
		return CIPDataTypeCode_STRING2, true
	case "FTIME":
		return CIPDataTypeCode_FTIME, true
	case "LTIME":
		return CIPDataTypeCode_LTIME, true
	case "ITIME":
		return CIPDataTypeCode_ITIME, true
	case "STRINGN":
		return CIPDataTypeCode_STRINGN, true
	case "SHORT_STRING":
		return CIPDataTypeCode_SHORT_STRING, true
	case "TIME":
		return CIPDataTypeCode_TIME, true
	case "EPATH":
		return CIPDataTypeCode_EPATH, true
	case "ENGUNIT":
		return CIPDataTypeCode_ENGUNIT, true
	case "STRUCTURED":
		return CIPDataTypeCode_STRUCTURED, true
	}
	return 0, false
}

func CIPDataTypeCodeKnows(value uint16) bool {
	for _, typeValue := range CIPDataTypeCodeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCIPDataTypeCode(structType any) CIPDataTypeCode {
	castFunc := func(typ any) CIPDataTypeCode {
		if sCIPDataTypeCode, ok := typ.(CIPDataTypeCode); ok {
			return sCIPDataTypeCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m CIPDataTypeCode) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m CIPDataTypeCode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPDataTypeCodeParse(ctx context.Context, theBytes []byte) (CIPDataTypeCode, error) {
	return CIPDataTypeCodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CIPDataTypeCodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CIPDataTypeCode, error) {
	val, err := readBuffer.ReadUint16("CIPDataTypeCode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CIPDataTypeCode")
	}
	if enum, ok := CIPDataTypeCodeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return CIPDataTypeCode(val), nil
	} else {
		return enum, nil
	}
}

func (e CIPDataTypeCode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CIPDataTypeCode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("CIPDataTypeCode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CIPDataTypeCode) PLC4XEnumName() string {
	switch e {
	case CIPDataTypeCode_BOOL:
		return "BOOL"
	case CIPDataTypeCode_SINT:
		return "SINT"
	case CIPDataTypeCode_INT:
		return "INT"
	case CIPDataTypeCode_DINT:
		return "DINT"
	case CIPDataTypeCode_LINT:
		return "LINT"
	case CIPDataTypeCode_USINT:
		return "USINT"
	case CIPDataTypeCode_UINT:
		return "UINT"
	case CIPDataTypeCode_UDINT:
		return "UDINT"
	case CIPDataTypeCode_ULINT:
		return "ULINT"
	case CIPDataTypeCode_REAL:
		return "REAL"
	case CIPDataTypeCode_LREAL:
		return "LREAL"
	case CIPDataTypeCode_STIME:
		return "STIME"
	case CIPDataTypeCode_DATE:
		return "DATE"
	case CIPDataTypeCode_TIME_OF_DAY:
		return "TIME_OF_DAY"
	case CIPDataTypeCode_DATE_AND_TIME:
		return "DATE_AND_TIME"
	case CIPDataTypeCode_STRING:
		return "STRING"
	case CIPDataTypeCode_BYTE:
		return "BYTE"
	case CIPDataTypeCode_WORD:
		return "WORD"
	case CIPDataTypeCode_DWORD:
		return "DWORD"
	case CIPDataTypeCode_STRING2:
		return "STRING2"
	case CIPDataTypeCode_FTIME:
		return "FTIME"
	case CIPDataTypeCode_LTIME:
		return "LTIME"
	case CIPDataTypeCode_ITIME:
		return "ITIME"
	case CIPDataTypeCode_STRINGN:
		return "STRINGN"
	case CIPDataTypeCode_SHORT_STRING:
		return "SHORT_STRING"
	case CIPDataTypeCode_TIME:
		return "TIME"
	case CIPDataTypeCode_EPATH:
		return "EPATH"
	case CIPDataTypeCode_ENGUNIT:
		return "ENGUNIT"
	case CIPDataTypeCode_STRUCTURED:
		return "STRUCTURED"
	}
	return ""
}

func (e CIPDataTypeCode) String() string {
	return e.PLC4XEnumName()
}
