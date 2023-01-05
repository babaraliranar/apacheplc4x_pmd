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

public class IODWriteRequestHeader extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.IOD_WRITE_REQUEST_HEADER;
  }

  // Constant values.
  public static final Integer PADFIELD = 0x0000;

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final int sequenceNumber;
  protected final Uuid arUuid;
  protected final long api;
  protected final int slotNumber;
  protected final int subSlotNumber;
  protected final int index;
  protected final long recordDataLength;
  protected final UserData userData;

  public IODWriteRequestHeader(
      short blockVersionHigh,
      short blockVersionLow,
      int sequenceNumber,
      Uuid arUuid,
      long api,
      int slotNumber,
      int subSlotNumber,
      int index,
      long recordDataLength,
      UserData userData) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.sequenceNumber = sequenceNumber;
    this.arUuid = arUuid;
    this.api = api;
    this.slotNumber = slotNumber;
    this.subSlotNumber = subSlotNumber;
    this.index = index;
    this.recordDataLength = recordDataLength;
    this.userData = userData;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public int getSequenceNumber() {
    return sequenceNumber;
  }

  public Uuid getArUuid() {
    return arUuid;
  }

  public long getApi() {
    return api;
  }

  public int getSlotNumber() {
    return slotNumber;
  }

  public int getSubSlotNumber() {
    return subSlotNumber;
  }

  public int getIndex() {
    return index;
  }

  public long getRecordDataLength() {
    return recordDataLength;
  }

  public UserData getUserData() {
    return userData;
  }

  public int getPadField() {
    return PADFIELD;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("IODWriteRequestHeader");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength =
        (int)
            ((((getIndex()) < (0x8000))
                ? ((getLengthInBytes()) - (4)) - (getRecordDataLength())
                : (getLengthInBytes()) - (4)));
    writeImplicitField("blockLength", blockLength, writeUnsignedInt(writeBuffer, 16));

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

    // Simple Field (sequenceNumber)
    writeSimpleField(
        "sequenceNumber",
        sequenceNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (arUuid)
    writeSimpleField(
        "arUuid",
        arUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (api)
    writeSimpleField(
        "api",
        api,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (slotNumber)
    writeSimpleField(
        "slotNumber",
        slotNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (subSlotNumber)
    writeSimpleField(
        "subSlotNumber",
        subSlotNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (padField)
    writeConstField("padField", PADFIELD, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (index)
    writeSimpleField(
        "index",
        index,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (recordDataLength)
    writeSimpleField(
        "recordDataLength",
        recordDataLength,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Padding Field (padding)
    writePaddingField(
        "padding",
        (int)
            ((((index) < (0x8000))
                ? (((((((((64) - (6)) - (2)) - (16)) - (4)) - (2)) - (2)) - (2)) - (2)) - (4)
                : (((((((((64) - (6)) - (2)) - (16)) - (4)) - (2)) - (2)) - (2)) - (2)) - (4))),
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8));

    // Optional Field (userData) (Can be skipped, if the value is null)
    writeOptionalField(
        "userData", userData, new DataWriterComplexDefault<>(writeBuffer), (getIndex()) < (0x8000));

    writeBuffer.popContext("IODWriteRequestHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IODWriteRequestHeader _value = this;

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Simple field (sequenceNumber)
    lengthInBits += 16;

    // Simple field (arUuid)
    lengthInBits += arUuid.getLengthInBits();

    // Simple field (api)
    lengthInBits += 32;

    // Simple field (slotNumber)
    lengthInBits += 16;

    // Simple field (subSlotNumber)
    lengthInBits += 16;

    // Const Field (padField)
    lengthInBits += 16;

    // Simple field (index)
    lengthInBits += 16;

    // Simple field (recordDataLength)
    lengthInBits += 32;

    // Padding Field (padding)
    int _timesPadding =
        (int)
            ((((index) < (0x8000))
                ? (((((((((64) - (6)) - (2)) - (16)) - (4)) - (2)) - (2)) - (2)) - (2)) - (4)
                : (((((((((64) - (6)) - (2)) - (16)) - (4)) - (2)) - (2)) - (2)) - (2)) - (4)));
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    // Optional Field (userData)
    if (userData != null) {
      lengthInBits += userData.getLengthInBits();
    }

    return lengthInBits;
  }

  public static IODWriteRequestHeaderBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("IODWriteRequestHeader");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

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

    int sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Uuid arUuid =
        readSimpleField(
            "arUuid",
            new DataReaderComplexDefault<>(() -> Uuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    long api =
        readSimpleField(
            "api",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int slotNumber =
        readSimpleField(
            "slotNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int subSlotNumber =
        readSimpleField(
            "subSlotNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int padField =
        readConstField(
            "padField",
            readUnsignedInt(readBuffer, 16),
            IODWriteRequestHeader.PADFIELD,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int index =
        readSimpleField(
            "index",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    long recordDataLength =
        readSimpleField(
            "recordDataLength",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readPaddingField(
        readUnsignedShort(readBuffer, 8),
        (int)
            ((((index) < (0x8000))
                ? (((((((((64) - (6)) - (2)) - (16)) - (4)) - (2)) - (2)) - (2)) - (2)) - (4)
                : (((((((((64) - (6)) - (2)) - (16)) - (4)) - (2)) - (2)) - (2)) - (2)) - (4))),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    UserData userData =
        readOptionalField(
            "userData",
            new DataReaderComplexDefault<>(
                () -> UserData.staticParse(readBuffer, (long) (recordDataLength)), readBuffer),
            (index) < (0x8000),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("IODWriteRequestHeader");
    // Create the instance
    return new IODWriteRequestHeaderBuilder(
        blockVersionHigh,
        blockVersionLow,
        sequenceNumber,
        arUuid,
        api,
        slotNumber,
        subSlotNumber,
        index,
        recordDataLength,
        userData);
  }

  public static class IODWriteRequestHeaderBuilder implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final int sequenceNumber;
    private final Uuid arUuid;
    private final long api;
    private final int slotNumber;
    private final int subSlotNumber;
    private final int index;
    private final long recordDataLength;
    private final UserData userData;

    public IODWriteRequestHeaderBuilder(
        short blockVersionHigh,
        short blockVersionLow,
        int sequenceNumber,
        Uuid arUuid,
        long api,
        int slotNumber,
        int subSlotNumber,
        int index,
        long recordDataLength,
        UserData userData) {

      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.sequenceNumber = sequenceNumber;
      this.arUuid = arUuid;
      this.api = api;
      this.slotNumber = slotNumber;
      this.subSlotNumber = subSlotNumber;
      this.index = index;
      this.recordDataLength = recordDataLength;
      this.userData = userData;
    }

    public IODWriteRequestHeader build() {
      IODWriteRequestHeader iODWriteRequestHeader =
          new IODWriteRequestHeader(
              blockVersionHigh,
              blockVersionLow,
              sequenceNumber,
              arUuid,
              api,
              slotNumber,
              subSlotNumber,
              index,
              recordDataLength,
              userData);
      return iODWriteRequestHeader;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IODWriteRequestHeader)) {
      return false;
    }
    IODWriteRequestHeader that = (IODWriteRequestHeader) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getSequenceNumber() == that.getSequenceNumber())
        && (getArUuid() == that.getArUuid())
        && (getApi() == that.getApi())
        && (getSlotNumber() == that.getSlotNumber())
        && (getSubSlotNumber() == that.getSubSlotNumber())
        && (getIndex() == that.getIndex())
        && (getRecordDataLength() == that.getRecordDataLength())
        && (getUserData() == that.getUserData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getSequenceNumber(),
        getArUuid(),
        getApi(),
        getSlotNumber(),
        getSubSlotNumber(),
        getIndex(),
        getRecordDataLength(),
        getUserData());
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
