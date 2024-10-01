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

// Code generated by "plc4xGenerator -type=MessageCodec"; DO NOT EDIT.

package cbus

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *MessageCodec) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *MessageCodec) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("MessageCodec"); err != nil {
		return err
	}
	if err := d.DefaultCodec.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}

	if d.requestContext != nil {
		if serializableField, ok := any(d.requestContext).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("requestContext"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("requestContext"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.requestContext)
			if err := writeBuffer.WriteString("requestContext", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if d.cbusOptions != nil {
		if serializableField, ok := any(d.cbusOptions).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("cbusOptions"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("cbusOptions"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.cbusOptions)
			if err := writeBuffer.WriteString("cbusOptions", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	_monitoredMMIs_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.monitoredMMIs))
	if err := writeBuffer.WriteString("monitoredMMIs", uint32(len(_monitoredMMIs_plx4gen_description)*8), _monitoredMMIs_plx4gen_description); err != nil {
		return err
	}

	_monitoredSALs_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.monitoredSALs))
	if err := writeBuffer.WriteString("monitoredSALs", uint32(len(_monitoredSALs_plx4gen_description)*8), _monitoredSALs_plx4gen_description); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint32("lastPackageHash", 32, d.lastPackageHash.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint64("hashEncountered", 64, d.hashEncountered.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint64("currentlyReportedServerErrors", 64, d.currentlyReportedServerErrors.Load()); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("MessageCodec"); err != nil {
		return err
	}
	return nil
}

func (d *MessageCodec) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	wb := utils.NewWriteBufferBoxBased(utils.WithWriteBufferBoxBasedMergeSingleBoxes(), utils.WithWriteBufferBoxBasedOmitEmptyBoxes())
	if err := wb.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
