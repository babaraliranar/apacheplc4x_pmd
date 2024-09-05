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

// Code generated by "plc4xGenerator -type=defaultConnection"; DO NOT EDIT.

package _default

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *defaultConnection) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *defaultConnection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("defaultConnection"); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("defaultTtl", uint32(len(fmt.Sprintf("%s", d.defaultTtl))*8), fmt.Sprintf("%s", d.defaultTtl)); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("connected", d.connected.Load()); err != nil {
		return err
	}

	if d.tagHandler != nil {
		if serializableField, ok := any(d.tagHandler).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("tagHandler"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("tagHandler"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.tagHandler)
			if err := writeBuffer.WriteString("tagHandler", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if d.valueHandler != nil {
		if serializableField, ok := any(d.valueHandler).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("valueHandler"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("valueHandler"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.valueHandler)
			if err := writeBuffer.WriteString("valueHandler", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("defaultConnection"); err != nil {
		return err
	}
	return nil
}

func (d *defaultConnection) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
