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
package org.apache.plc4x.java.ads.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class AdsAddDeviceNotificationRequest extends AmsPacket implements Message {

  // Accessors for discriminator values.
  public CommandId getCommandId() {
    return CommandId.ADS_ADD_DEVICE_NOTIFICATION;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Properties.
  protected final long indexGroup;
  protected final long indexOffset;
  protected final long length;
  protected final AdsTransMode transmissionMode;
  protected final long maxDelayInMs;
  protected final long cycleTimeInMs;

  public AdsAddDeviceNotificationRequest(
      AmsNetId targetAmsNetId,
      int targetAmsPort,
      AmsNetId sourceAmsNetId,
      int sourceAmsPort,
      long errorCode,
      long invokeId,
      long indexGroup,
      long indexOffset,
      long length,
      AdsTransMode transmissionMode,
      long maxDelayInMs,
      long cycleTimeInMs) {
    super(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId);
    this.indexGroup = indexGroup;
    this.indexOffset = indexOffset;
    this.length = length;
    this.transmissionMode = transmissionMode;
    this.maxDelayInMs = maxDelayInMs;
    this.cycleTimeInMs = cycleTimeInMs;
  }

  public long getIndexGroup() {
    return indexGroup;
  }

  public long getIndexOffset() {
    return indexOffset;
  }

  public long getLength() {
    return length;
  }

  public AdsTransMode getTransmissionMode() {
    return transmissionMode;
  }

  public long getMaxDelayInMs() {
    return maxDelayInMs;
  }

  public long getCycleTimeInMs() {
    return cycleTimeInMs;
  }

  @Override
  protected void serializeAmsPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdsAddDeviceNotificationRequest");

    // Simple Field (indexGroup)
    writeSimpleField("indexGroup", indexGroup, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (indexOffset)
    writeSimpleField("indexOffset", indexOffset, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (length)
    writeSimpleField("length", length, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (transmissionMode)
    writeSimpleEnumField(
        "transmissionMode",
        "AdsTransMode",
        transmissionMode,
        new DataWriterEnumDefault<>(
            AdsTransMode::getValue, AdsTransMode::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (maxDelayInMs)
    writeSimpleField("maxDelayInMs", maxDelayInMs, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (cycleTimeInMs)
    writeSimpleField("cycleTimeInMs", cycleTimeInMs, writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved", BigInteger.valueOf(0x0000), writeUnsignedBigInteger(writeBuffer, 64));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved", BigInteger.valueOf(0x0000), writeUnsignedBigInteger(writeBuffer, 64));

    writeBuffer.popContext("AdsAddDeviceNotificationRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsAddDeviceNotificationRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (indexGroup)
    lengthInBits += 32;

    // Simple field (indexOffset)
    lengthInBits += 32;

    // Simple field (length)
    lengthInBits += 32;

    // Simple field (transmissionMode)
    lengthInBits += 32;

    // Simple field (maxDelayInMs)
    lengthInBits += 32;

    // Simple field (cycleTimeInMs)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 64;

    // Reserved Field (reserved)
    lengthInBits += 64;

    return lengthInBits;
  }

  public static AmsPacketBuilder staticParseAmsPacketBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsAddDeviceNotificationRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long indexGroup = readSimpleField("indexGroup", readUnsignedLong(readBuffer, 32));

    long indexOffset = readSimpleField("indexOffset", readUnsignedLong(readBuffer, 32));

    long length = readSimpleField("length", readUnsignedLong(readBuffer, 32));

    AdsTransMode transmissionMode =
        readEnumField(
            "transmissionMode",
            "AdsTransMode",
            new DataReaderEnumDefault<>(
                AdsTransMode::enumForValue, readUnsignedLong(readBuffer, 32)));

    long maxDelayInMs = readSimpleField("maxDelayInMs", readUnsignedLong(readBuffer, 32));

    long cycleTimeInMs = readSimpleField("cycleTimeInMs", readUnsignedLong(readBuffer, 32));

    BigInteger reservedField0 =
        readReservedField(
            "reserved", readUnsignedBigInteger(readBuffer, 64), BigInteger.valueOf(0x0000));

    BigInteger reservedField1 =
        readReservedField(
            "reserved", readUnsignedBigInteger(readBuffer, 64), BigInteger.valueOf(0x0000));

    readBuffer.closeContext("AdsAddDeviceNotificationRequest");
    // Create the instance
    return new AdsAddDeviceNotificationRequestBuilderImpl(
        indexGroup, indexOffset, length, transmissionMode, maxDelayInMs, cycleTimeInMs);
  }

  public static class AdsAddDeviceNotificationRequestBuilderImpl
      implements AmsPacket.AmsPacketBuilder {
    private final long indexGroup;
    private final long indexOffset;
    private final long length;
    private final AdsTransMode transmissionMode;
    private final long maxDelayInMs;
    private final long cycleTimeInMs;

    public AdsAddDeviceNotificationRequestBuilderImpl(
        long indexGroup,
        long indexOffset,
        long length,
        AdsTransMode transmissionMode,
        long maxDelayInMs,
        long cycleTimeInMs) {
      this.indexGroup = indexGroup;
      this.indexOffset = indexOffset;
      this.length = length;
      this.transmissionMode = transmissionMode;
      this.maxDelayInMs = maxDelayInMs;
      this.cycleTimeInMs = cycleTimeInMs;
    }

    public AdsAddDeviceNotificationRequest build(
        AmsNetId targetAmsNetId,
        int targetAmsPort,
        AmsNetId sourceAmsNetId,
        int sourceAmsPort,
        long errorCode,
        long invokeId) {
      AdsAddDeviceNotificationRequest adsAddDeviceNotificationRequest =
          new AdsAddDeviceNotificationRequest(
              targetAmsNetId,
              targetAmsPort,
              sourceAmsNetId,
              sourceAmsPort,
              errorCode,
              invokeId,
              indexGroup,
              indexOffset,
              length,
              transmissionMode,
              maxDelayInMs,
              cycleTimeInMs);
      return adsAddDeviceNotificationRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsAddDeviceNotificationRequest)) {
      return false;
    }
    AdsAddDeviceNotificationRequest that = (AdsAddDeviceNotificationRequest) o;
    return (getIndexGroup() == that.getIndexGroup())
        && (getIndexOffset() == that.getIndexOffset())
        && (getLength() == that.getLength())
        && (getTransmissionMode() == that.getTransmissionMode())
        && (getMaxDelayInMs() == that.getMaxDelayInMs())
        && (getCycleTimeInMs() == that.getCycleTimeInMs())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getIndexGroup(),
        getIndexOffset(),
        getLength(),
        getTransmissionMode(),
        getMaxDelayInMs(),
        getCycleTimeInMs());
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
