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
package org.apache.plc4x.java.canopen.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class CANOpenFrame implements Message {

  // Constant values.
  public static final Byte ALIGNMENT = 0x00;

  // Properties.
  protected final short nodeId;
  protected final CANOpenService service;
  protected final CANOpenPayload payload;

  public CANOpenFrame(short nodeId, CANOpenService service, CANOpenPayload payload) {
    super();
    this.nodeId = nodeId;
    this.service = service;
    this.payload = payload;
  }

  public short getNodeId() {
    return nodeId;
  }

  public CANOpenService getService() {
    return service;
  }

  public CANOpenPayload getPayload() {
    return payload;
  }

  public byte getAlignment() {
    return ALIGNMENT;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CANOpenFrame");

    // Simple Field (nodeId)
    writeSimpleField(
        "nodeId",
        nodeId,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (service)
    writeSimpleEnumField(
        "service",
        "CANOpenService",
        service,
        new DataWriterEnumDefault<>(
            CANOpenService::getValue, CANOpenService::name, writeUnsignedByte(writeBuffer, 4)),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Const Field (alignment)
    writeConstField(
        "alignment",
        ALIGNMENT,
        writeUnsignedByte(writeBuffer, 4),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (payload)
    writeSimpleField(
        "payload",
        payload,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Padding Field (padding)
    writePaddingField(
        "padding",
        (int) ((8) - ((payload.getLengthInBytes()))),
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("CANOpenFrame");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CANOpenFrame _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nodeId)
    lengthInBits += 8;

    // Simple field (service)
    lengthInBits += 4;

    // Const Field (alignment)
    lengthInBits += 4;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // Padding Field (padding)
    int _timesPadding = (int) ((8) - ((payload.getLengthInBytes())));
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static CANOpenFrame staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CANOpenFrame");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short nodeId =
        readSimpleField(
            "nodeId",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    CANOpenService service =
        readEnumField(
            "service",
            "CANOpenService",
            new DataReaderEnumDefault<>(
                CANOpenService::enumForValue, readUnsignedByte(readBuffer, 4)),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    byte alignment =
        readConstField(
            "alignment",
            readUnsignedByte(readBuffer, 4),
            CANOpenFrame.ALIGNMENT,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    CANOpenPayload payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () -> CANOpenPayload.staticParse(readBuffer, (CANOpenService) (service)),
                readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readPaddingField(
        readUnsignedShort(readBuffer, 8),
        (int) ((8) - ((payload.getLengthInBytes()))),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("CANOpenFrame");
    // Create the instance
    CANOpenFrame _cANOpenFrame;
    _cANOpenFrame = new CANOpenFrame(nodeId, service, payload);
    return _cANOpenFrame;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CANOpenFrame)) {
      return false;
    }
    CANOpenFrame that = (CANOpenFrame) o;
    return (getNodeId() == that.getNodeId())
        && (getService() == that.getService())
        && (getPayload() == that.getPayload())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNodeId(), getService(), getPayload());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
