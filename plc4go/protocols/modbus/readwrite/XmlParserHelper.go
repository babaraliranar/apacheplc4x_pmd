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

package readwrite

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/apache/plc4x/plc4go/protocols/modbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

type ModbusXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m ModbusXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (any, error) {
	switch typeName {
	case "ModbusPDUWriteFileRecordRequestItem":
		return model.ModbusPDUWriteFileRecordRequestItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DataItem":
		dataType, _ := model.ModbusDataTypeByName(parserArguments[0])
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 16)
		if err != nil {
			return nil, err
		}
		numberOfValues := uint16(parsedUint1)
		bigEndian := parserArguments[2] == "true"
		return model.DataItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), dataType, numberOfValues, bigEndian)
	case "ModbusPDUReadFileRecordResponseItem":
		return model.ModbusPDUReadFileRecordResponseItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusDeviceInformationObject":
		return model.ModbusDeviceInformationObjectParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusConstants":
		return model.ModbusConstantsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusPDUWriteFileRecordResponseItem":
		return model.ModbusPDUWriteFileRecordResponseItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusPDU":
		response := parserArguments[0] == "true"
		return model.ModbusPDUParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), response)
	case "ModbusPDUReadFileRecordRequestItem":
		return model.ModbusPDUReadFileRecordRequestItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ModbusADU":
		driverType, _ := model.DriverTypeByName(parserArguments[0])
		response := parserArguments[1] == "true"
		return model.ModbusADUParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), driverType, response)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
