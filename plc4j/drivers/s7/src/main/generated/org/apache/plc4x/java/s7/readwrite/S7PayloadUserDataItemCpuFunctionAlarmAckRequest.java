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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7PayloadUserDataItemCpuFunctionAlarmAckRequest extends S7PayloadUserDataItem
    implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x04;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x04;
  }

  public Short getCpuSubfunction() {
    return (short) 0x0b;
  }

  // Constant values.
  public static final Short FUNCTIONID = 0x09;

  // Properties.
  protected final List<AlarmMessageObjectAckType> messageObjects;

  public S7PayloadUserDataItemCpuFunctionAlarmAckRequest(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      List<AlarmMessageObjectAckType> messageObjects) {
    super(returnCode, transportSize, dataLength);
    this.messageObjects = messageObjects;
  }

  public List<AlarmMessageObjectAckType> getMessageObjects() {
    return messageObjects;
  }

  public short getFunctionId() {
    return FUNCTIONID;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest");

    // Const Field (functionId)
    writeConstField("functionId", FUNCTIONID, writeUnsignedShort(writeBuffer, 8));

    // Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short numberOfObjects = (short) (COUNT(getMessageObjects()));
    writeImplicitField("numberOfObjects", numberOfObjects, writeUnsignedShort(writeBuffer, 8));

    // Array Field (messageObjects)
    writeComplexTypeArrayField("messageObjects", messageObjects, writeBuffer);

    writeBuffer.popContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemCpuFunctionAlarmAckRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (functionId)
    lengthInBits += 8;

    // Implicit Field (numberOfObjects)
    lengthInBits += 8;

    // Array field
    if (messageObjects != null) {
      int i = 0;
      for (AlarmMessageObjectAckType element : messageObjects) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= messageObjects.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionGroup, Byte cpuFunctionType, Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short functionId =
        readConstField(
            "functionId",
            readUnsignedShort(readBuffer, 8),
            S7PayloadUserDataItemCpuFunctionAlarmAckRequest.FUNCTIONID);

    short numberOfObjects = readImplicitField("numberOfObjects", readUnsignedShort(readBuffer, 8));

    List<AlarmMessageObjectAckType> messageObjects =
        readCountArrayField(
            "messageObjects",
            new DataReaderComplexDefault<>(
                () -> AlarmMessageObjectAckType.staticParse(readBuffer), readBuffer),
            numberOfObjects);

    readBuffer.closeContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest");
    // Create the instance
    return new S7PayloadUserDataItemCpuFunctionAlarmAckRequestBuilderImpl(messageObjects);
  }

  public static class S7PayloadUserDataItemCpuFunctionAlarmAckRequestBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final List<AlarmMessageObjectAckType> messageObjects;

    public S7PayloadUserDataItemCpuFunctionAlarmAckRequestBuilderImpl(
        List<AlarmMessageObjectAckType> messageObjects) {
      this.messageObjects = messageObjects;
    }

    public S7PayloadUserDataItemCpuFunctionAlarmAckRequest build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadUserDataItemCpuFunctionAlarmAckRequest
          s7PayloadUserDataItemCpuFunctionAlarmAckRequest =
              new S7PayloadUserDataItemCpuFunctionAlarmAckRequest(
                  returnCode, transportSize, dataLength, messageObjects);
      return s7PayloadUserDataItemCpuFunctionAlarmAckRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemCpuFunctionAlarmAckRequest)) {
      return false;
    }
    S7PayloadUserDataItemCpuFunctionAlarmAckRequest that =
        (S7PayloadUserDataItemCpuFunctionAlarmAckRequest) o;
    return (getMessageObjects() == that.getMessageObjects()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMessageObjects());
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
