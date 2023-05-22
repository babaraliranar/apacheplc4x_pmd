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

public class PcDcp_GetSet_Pdu extends PnDcp_Pdu implements Message {

  // Accessors for discriminator values.

  // Constant values.
  public static final Short SERVICEID = 0x04;

  // Properties.
  protected final boolean notSupported;
  protected final boolean response;
  protected final long xid;
  protected final List<PnDcp_Block> blocks;

  // Reserved Fields
  private Short reservedField0;
  private Byte reservedField1;
  private Integer reservedField2;

  public PcDcp_GetSet_Pdu(
      int frameIdValue,
      boolean notSupported,
      boolean response,
      long xid,
      List<PnDcp_Block> blocks) {
    super(frameIdValue);
    this.notSupported = notSupported;
    this.response = response;
    this.xid = xid;
    this.blocks = blocks;
  }

  public boolean getNotSupported() {
    return notSupported;
  }

  public boolean getResponse() {
    return response;
  }

  public long getXid() {
    return xid;
  }

  public List<PnDcp_Block> getBlocks() {
    return blocks;
  }

  public short getServiceId() {
    return SERVICEID;
  }

  @Override
  protected void serializePnDcp_PduChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PcDcp_GetSet_Pdu");

    // Const Field (serviceId)
    writeConstField(
        "serviceId",
        SERVICEID,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 5),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (notSupported)
    writeSimpleField(
        "notSupported",
        notSupported,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (byte) 0x00,
        writeUnsignedByte(writeBuffer, 1),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (response)
    writeSimpleField(
        "response",
        response,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (xid)
    writeSimpleField(
        "xid",
        xid,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField2 != null ? reservedField2 : (int) 0x0000,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (dcpDataLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int dcpDataLength = (int) ((getLengthInBytes()) - (12));
    writeImplicitField(
        "dcpDataLength",
        dcpDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (blocks)
    writeComplexTypeArrayField(
        "blocks", blocks, writeBuffer, WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PcDcp_GetSet_Pdu");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PcDcp_GetSet_Pdu _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (serviceId)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 5;

    // Simple field (notSupported)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (response)
    lengthInBits += 1;

    // Simple field (xid)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Implicit Field (dcpDataLength)
    lengthInBits += 16;

    // Array field
    if (blocks != null) {
      for (Message element : blocks) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static PnDcp_PduBuilder staticParsePnDcp_PduBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PcDcp_GetSet_Pdu");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short serviceId =
        readConstField(
            "serviceId",
            readUnsignedShort(readBuffer, 8),
            PcDcp_GetSet_Pdu.SERVICEID,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Short reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 5),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean notSupported =
        readSimpleField(
            "notSupported",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Byte reservedField1 =
        readReservedField(
            "reserved",
            readUnsignedByte(readBuffer, 1),
            (byte) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean response =
        readSimpleField(
            "response", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    long xid =
        readSimpleField(
            "xid",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Integer reservedField2 =
        readReservedField(
            "reserved",
            readUnsignedInt(readBuffer, 16),
            (int) 0x0000,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int dcpDataLength =
        readImplicitField(
            "dcpDataLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<PnDcp_Block> blocks =
        readLengthArrayField(
            "blocks",
            new DataReaderComplexDefault<>(() -> PnDcp_Block.staticParse(readBuffer), readBuffer),
            dcpDataLength,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PcDcp_GetSet_Pdu");
    // Create the instance
    return new PcDcp_GetSet_PduBuilderImpl(
        notSupported, response, xid, blocks, reservedField0, reservedField1, reservedField2);
  }

  public static class PcDcp_GetSet_PduBuilderImpl implements PnDcp_Pdu.PnDcp_PduBuilder {
    private final boolean notSupported;
    private final boolean response;
    private final long xid;
    private final List<PnDcp_Block> blocks;
    private final Short reservedField0;
    private final Byte reservedField1;
    private final Integer reservedField2;

    public PcDcp_GetSet_PduBuilderImpl(
        boolean notSupported,
        boolean response,
        long xid,
        List<PnDcp_Block> blocks,
        Short reservedField0,
        Byte reservedField1,
        Integer reservedField2) {
      this.notSupported = notSupported;
      this.response = response;
      this.xid = xid;
      this.blocks = blocks;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
      this.reservedField2 = reservedField2;
    }

    public PcDcp_GetSet_Pdu build(int frameIdValue) {
      PcDcp_GetSet_Pdu pcDcp_GetSet_Pdu =
          new PcDcp_GetSet_Pdu(frameIdValue, notSupported, response, xid, blocks);
      pcDcp_GetSet_Pdu.reservedField0 = reservedField0;
      pcDcp_GetSet_Pdu.reservedField1 = reservedField1;
      pcDcp_GetSet_Pdu.reservedField2 = reservedField2;
      return pcDcp_GetSet_Pdu;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PcDcp_GetSet_Pdu)) {
      return false;
    }
    PcDcp_GetSet_Pdu that = (PcDcp_GetSet_Pdu) o;
    return (getNotSupported() == that.getNotSupported())
        && (getResponse() == that.getResponse())
        && (getXid() == that.getXid())
        && (getBlocks() == that.getBlocks())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNotSupported(), getResponse(), getXid(), getBlocks());
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
