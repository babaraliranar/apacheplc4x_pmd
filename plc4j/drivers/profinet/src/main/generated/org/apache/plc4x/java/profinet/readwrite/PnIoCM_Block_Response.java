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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCM_Block_Response extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.IOX_BLOCK_RES;
  }

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final Uuid arUuid;
  protected final int sessionKey;
  protected final int controlCommand;
  protected final int controlBlockProperties;

  // Reserved Fields
  private Integer reservedField0;
  private Integer reservedField1;

  public PnIoCM_Block_Response(
      short blockVersionHigh,
      short blockVersionLow,
      Uuid arUuid,
      int sessionKey,
      int controlCommand,
      int controlBlockProperties) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.arUuid = arUuid;
    this.sessionKey = sessionKey;
    this.controlCommand = controlCommand;
    this.controlBlockProperties = controlBlockProperties;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public Uuid getArUuid() {
    return arUuid;
  }

  public int getSessionKey() {
    return sessionKey;
  }

  public int getControlCommand() {
    return controlCommand;
  }

  public int getControlBlockProperties() {
    return controlBlockProperties;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnIoCM_Block_Response");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (int) 0x0000,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (arUuid)
    writeSimpleField(
        "arUuid",
        arUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (sessionKey)
    writeSimpleField(
        "sessionKey",
        sessionKey,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (int) 0x0000,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (controlCommand)
    writeSimpleField(
        "controlCommand",
        controlCommand,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (controlBlockProperties)
    writeSimpleField(
        "controlBlockProperties",
        controlBlockProperties,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCM_Block_Response");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCM_Block_Response _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Simple field (arUuid)
    lengthInBits += arUuid.getLengthInBits();

    // Simple field (sessionKey)
    lengthInBits += 16;

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Simple field (controlCommand)
    lengthInBits += 16;

    // Simple field (controlBlockProperties)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCM_Block_Response");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Integer reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedInt(readBuffer, 16),
            (int) 0x0000,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Uuid arUuid =
        readSimpleField(
            "arUuid",
            new DataReaderComplexDefault<>(() -> Uuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int sessionKey =
        readSimpleField(
            "sessionKey",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Integer reservedField1 =
        readReservedField(
            "reserved",
            readUnsignedInt(readBuffer, 16),
            (int) 0x0000,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int controlCommand =
        readSimpleField(
            "controlCommand",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int controlBlockProperties =
        readSimpleField(
            "controlBlockProperties",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCM_Block_Response");
    // Create the instance
    return new PnIoCM_Block_ResponseBuilderImpl(
        blockVersionHigh,
        blockVersionLow,
        arUuid,
        sessionKey,
        controlCommand,
        controlBlockProperties,
        reservedField0,
        reservedField1);
  }

  public static class PnIoCM_Block_ResponseBuilderImpl implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final Uuid arUuid;
    private final int sessionKey;
    private final int controlCommand;
    private final int controlBlockProperties;
    private final Integer reservedField0;
    private final Integer reservedField1;

    public PnIoCM_Block_ResponseBuilderImpl(
        short blockVersionHigh,
        short blockVersionLow,
        Uuid arUuid,
        int sessionKey,
        int controlCommand,
        int controlBlockProperties,
        Integer reservedField0,
        Integer reservedField1) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.arUuid = arUuid;
      this.sessionKey = sessionKey;
      this.controlCommand = controlCommand;
      this.controlBlockProperties = controlBlockProperties;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
    }

    public PnIoCM_Block_Response build() {
      PnIoCM_Block_Response pnIoCM_Block_Response =
          new PnIoCM_Block_Response(
              blockVersionHigh,
              blockVersionLow,
              arUuid,
              sessionKey,
              controlCommand,
              controlBlockProperties);
      pnIoCM_Block_Response.reservedField0 = reservedField0;
      pnIoCM_Block_Response.reservedField1 = reservedField1;
      return pnIoCM_Block_Response;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCM_Block_Response)) {
      return false;
    }
    PnIoCM_Block_Response that = (PnIoCM_Block_Response) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getArUuid() == that.getArUuid())
        && (getSessionKey() == that.getSessionKey())
        && (getControlCommand() == that.getControlCommand())
        && (getControlBlockProperties() == that.getControlBlockProperties())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getArUuid(),
        getSessionKey(),
        getControlCommand(),
        getControlBlockProperties());
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
