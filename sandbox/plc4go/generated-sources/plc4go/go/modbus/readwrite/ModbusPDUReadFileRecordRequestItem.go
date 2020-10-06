//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package readwrite

import "plc4x.apache.org/plc4go-modbus-driver/0.8.0/src/plc4go/spi"

type ModbusPDUReadFileRecordRequestItem struct {
	referenceType uint8
	fileNumber    uint16
	recordNumber  uint16
	recordLength  uint16
}

func (m ModbusPDUReadFileRecordRequestItem) lengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Simple field (recordLength)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUReadFileRecordRequestItem) lengthInBytes() uint16 {
	return m.lengthInBits() / 8
}

func (m ModbusPDUReadFileRecordRequestItem) parse(io spi.ReadBuffer) {
	// TODO: Implement ...
}

func (m ModbusPDUReadFileRecordRequestItem) serialize(io spi.WriteBuffer) {
	// TODO: Implement ...
}
