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

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class AmsPacket implements Message {

  // Abstract accessors for discriminator values.
  public abstract CommandId getCommandId();

  public abstract Boolean getResponse();

  // Constant values.
  public static final Boolean INITCOMMAND = false;
  public static final Boolean UPDCOMMAND = false;
  public static final Boolean TIMESTAMPADDED = false;
  public static final Boolean HIGHPRIORITYCOMMAND = false;
  public static final Boolean SYSTEMCOMMAND = false;
  public static final Boolean ADSCOMMAND = true;
  public static final Boolean NORETURN = false;
  public static final Boolean BROADCAST = false;

  // Properties.
  protected final AmsNetId targetAmsNetId;
  protected final int targetAmsPort;
  protected final AmsNetId sourceAmsNetId;
  protected final int sourceAmsPort;
  protected final long errorCode;
  protected final long invokeId;

  public AmsPacket(
      AmsNetId targetAmsNetId,
      int targetAmsPort,
      AmsNetId sourceAmsNetId,
      int sourceAmsPort,
      long errorCode,
      long invokeId) {
    super();
    this.targetAmsNetId = targetAmsNetId;
    this.targetAmsPort = targetAmsPort;
    this.sourceAmsNetId = sourceAmsNetId;
    this.sourceAmsPort = sourceAmsPort;
    this.errorCode = errorCode;
    this.invokeId = invokeId;
  }

  public AmsNetId getTargetAmsNetId() {
    return targetAmsNetId;
  }

  public int getTargetAmsPort() {
    return targetAmsPort;
  }

  public AmsNetId getSourceAmsNetId() {
    return sourceAmsNetId;
  }

  public int getSourceAmsPort() {
    return sourceAmsPort;
  }

  public long getErrorCode() {
    return errorCode;
  }

  public long getInvokeId() {
    return invokeId;
  }

  public boolean getInitCommand() {
    return INITCOMMAND;
  }

  public boolean getUpdCommand() {
    return UPDCOMMAND;
  }

  public boolean getTimestampAdded() {
    return TIMESTAMPADDED;
  }

  public boolean getHighPriorityCommand() {
    return HIGHPRIORITYCOMMAND;
  }

  public boolean getSystemCommand() {
    return SYSTEMCOMMAND;
  }

  public boolean getAdsCommand() {
    return ADSCOMMAND;
  }

  public boolean getNoReturn() {
    return NORETURN;
  }

  public boolean getBroadcast() {
    return BROADCAST;
  }

  protected abstract void serializeAmsPacketChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AmsPacket");

    // Simple Field (targetAmsNetId)
    writeSimpleField("targetAmsNetId", targetAmsNetId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (targetAmsPort)
    writeSimpleField("targetAmsPort", targetAmsPort, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (sourceAmsNetId)
    writeSimpleField("sourceAmsNetId", sourceAmsNetId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (sourceAmsPort)
    writeSimpleField("sourceAmsPort", sourceAmsPort, writeUnsignedInt(writeBuffer, 16));

    // Discriminator Field (commandId) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "commandId",
        "CommandId",
        getCommandId(),
        new DataWriterEnumDefault<>(
            CommandId::getValue, CommandId::name, writeUnsignedInt(writeBuffer, 16)));

    // Const Field (initCommand)
    writeConstField("initCommand", INITCOMMAND, writeBoolean(writeBuffer));

    // Const Field (updCommand)
    writeConstField("updCommand", UPDCOMMAND, writeBoolean(writeBuffer));

    // Const Field (timestampAdded)
    writeConstField("timestampAdded", TIMESTAMPADDED, writeBoolean(writeBuffer));

    // Const Field (highPriorityCommand)
    writeConstField("highPriorityCommand", HIGHPRIORITYCOMMAND, writeBoolean(writeBuffer));

    // Const Field (systemCommand)
    writeConstField("systemCommand", SYSTEMCOMMAND, writeBoolean(writeBuffer));

    // Const Field (adsCommand)
    writeConstField("adsCommand", ADSCOMMAND, writeBoolean(writeBuffer));

    // Const Field (noReturn)
    writeConstField("noReturn", NORETURN, writeBoolean(writeBuffer));

    // Discriminator Field (response) (Used as input to a switch field)
    writeDiscriminatorField("response", getResponse(), writeBoolean(writeBuffer));

    // Const Field (broadcast)
    writeConstField("broadcast", BROADCAST, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x0, writeSignedByte(writeBuffer, 7));

    // Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    long length = (long) ((getLengthInBytes()) - (32L));
    writeImplicitField("length", length, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (errorCode)
    writeSimpleField("errorCode", errorCode, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (invokeId)
    writeSimpleField("invokeId", invokeId, writeUnsignedLong(writeBuffer, 32));

    // Switch field (Serialize the sub-type)
    serializeAmsPacketChild(writeBuffer);

    writeBuffer.popContext("AmsPacket");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AmsPacket _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (targetAmsNetId)
    lengthInBits += targetAmsNetId.getLengthInBits();

    // Simple field (targetAmsPort)
    lengthInBits += 16;

    // Simple field (sourceAmsNetId)
    lengthInBits += sourceAmsNetId.getLengthInBits();

    // Simple field (sourceAmsPort)
    lengthInBits += 16;

    // Discriminator Field (commandId)
    lengthInBits += 16;

    // Const Field (initCommand)
    lengthInBits += 1;

    // Const Field (updCommand)
    lengthInBits += 1;

    // Const Field (timestampAdded)
    lengthInBits += 1;

    // Const Field (highPriorityCommand)
    lengthInBits += 1;

    // Const Field (systemCommand)
    lengthInBits += 1;

    // Const Field (adsCommand)
    lengthInBits += 1;

    // Const Field (noReturn)
    lengthInBits += 1;

    // Discriminator Field (response)
    lengthInBits += 1;

    // Const Field (broadcast)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Implicit Field (length)
    lengthInBits += 32;

    // Simple field (errorCode)
    lengthInBits += 32;

    // Simple field (invokeId)
    lengthInBits += 32;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static AmsPacket staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AmsPacket staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AmsPacket");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    AmsNetId targetAmsNetId =
        readSimpleField(
            "targetAmsNetId",
            new DataReaderComplexDefault<>(() -> AmsNetId.staticParse(readBuffer), readBuffer));

    int targetAmsPort = readSimpleField("targetAmsPort", readUnsignedInt(readBuffer, 16));

    AmsNetId sourceAmsNetId =
        readSimpleField(
            "sourceAmsNetId",
            new DataReaderComplexDefault<>(() -> AmsNetId.staticParse(readBuffer), readBuffer));

    int sourceAmsPort = readSimpleField("sourceAmsPort", readUnsignedInt(readBuffer, 16));

    CommandId commandId =
        readDiscriminatorField(
            "commandId",
            new DataReaderEnumDefault<>(CommandId::enumForValue, readUnsignedInt(readBuffer, 16)));

    boolean initCommand =
        readConstField("initCommand", readBoolean(readBuffer), AmsPacket.INITCOMMAND);

    boolean updCommand =
        readConstField("updCommand", readBoolean(readBuffer), AmsPacket.UPDCOMMAND);

    boolean timestampAdded =
        readConstField("timestampAdded", readBoolean(readBuffer), AmsPacket.TIMESTAMPADDED);

    boolean highPriorityCommand =
        readConstField(
            "highPriorityCommand", readBoolean(readBuffer), AmsPacket.HIGHPRIORITYCOMMAND);

    boolean systemCommand =
        readConstField("systemCommand", readBoolean(readBuffer), AmsPacket.SYSTEMCOMMAND);

    boolean adsCommand =
        readConstField("adsCommand", readBoolean(readBuffer), AmsPacket.ADSCOMMAND);

    boolean noReturn = readConstField("noReturn", readBoolean(readBuffer), AmsPacket.NORETURN);

    boolean response = readDiscriminatorField("response", readBoolean(readBuffer));

    boolean broadcast = readConstField("broadcast", readBoolean(readBuffer), AmsPacket.BROADCAST);

    Byte reservedField0 = readReservedField("reserved", readSignedByte(readBuffer, 7), (byte) 0x0);

    long length = readImplicitField("length", readUnsignedLong(readBuffer, 32));

    long errorCode = readSimpleField("errorCode", readUnsignedLong(readBuffer, 32));

    long invokeId = readSimpleField("invokeId", readUnsignedLong(readBuffer, 32));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    AmsPacketBuilder builder = null;
    if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.INVALID)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsInvalidRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.INVALID)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsInvalidResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ_DEVICE_INFO)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsReadDeviceInfoRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ_DEVICE_INFO)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsReadDeviceInfoResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsReadRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsReadResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_WRITE)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsWriteRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_WRITE)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsWriteResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ_STATE)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsReadStateRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ_STATE)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsReadStateResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_WRITE_CONTROL)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsWriteControlRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_WRITE_CONTROL)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsWriteControlResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_ADD_DEVICE_NOTIFICATION)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsAddDeviceNotificationRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_ADD_DEVICE_NOTIFICATION)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsAddDeviceNotificationResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_DELETE_DEVICE_NOTIFICATION)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsDeleteDeviceNotificationRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_DELETE_DEVICE_NOTIFICATION)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsDeleteDeviceNotificationResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_DEVICE_NOTIFICATION)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsDeviceNotificationRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_DEVICE_NOTIFICATION)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsDeviceNotificationResponse.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ_WRITE)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = AdsReadWriteRequest.staticParseAmsPacketBuilder(readBuffer);
    } else if (EvaluationHelper.equals(errorCode, (long) 0x00000000L)
        && EvaluationHelper.equals(commandId, CommandId.ADS_READ_WRITE)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = AdsReadWriteResponse.staticParseAmsPacketBuilder(readBuffer);
    } else {
      builder = ErrorResponse.staticParseAmsPacketBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "errorCode="
              + errorCode
              + " "
              + "commandId="
              + commandId
              + " "
              + "response="
              + response
              + "]");
    }

    readBuffer.closeContext("AmsPacket");
    // Create the instance
    AmsPacket _amsPacket =
        builder.build(
            targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId);
    return _amsPacket;
  }

  public interface AmsPacketBuilder {
    AmsPacket build(
        AmsNetId targetAmsNetId,
        int targetAmsPort,
        AmsNetId sourceAmsNetId,
        int sourceAmsPort,
        long errorCode,
        long invokeId);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AmsPacket)) {
      return false;
    }
    AmsPacket that = (AmsPacket) o;
    return (getTargetAmsNetId() == that.getTargetAmsNetId())
        && (getTargetAmsPort() == that.getTargetAmsPort())
        && (getSourceAmsNetId() == that.getSourceAmsNetId())
        && (getSourceAmsPort() == that.getSourceAmsPort())
        && (getErrorCode() == that.getErrorCode())
        && (getInvokeId() == that.getInvokeId())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getTargetAmsNetId(),
        getTargetAmsPort(),
        getSourceAmsNetId(),
        getSourceAmsPort(),
        getErrorCode(),
        getInvokeId());
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
