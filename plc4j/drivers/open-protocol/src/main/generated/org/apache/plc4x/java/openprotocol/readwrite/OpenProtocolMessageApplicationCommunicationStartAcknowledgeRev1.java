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
package org.apache.plc4x.java.openprotocol.readwrite;

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

public class OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1
    extends OpenProtocolMessageApplicationCommunicationStartAcknowledge implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 1;
  }

  // Constant values.
  public static final Integer BLOCKIDCELLID = 1;
  public static final Integer BLOCKIDCHANNELID = 2;
  public static final Integer BLOCKIDCONTROLLERNAME = 3;

  // Properties.
  protected final long cellId;
  protected final int channelId;
  protected final String controllerName;

  public OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      long cellId,
      int channelId,
      String controllerName) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.cellId = cellId;
    this.channelId = channelId;
    this.controllerName = controllerName;
  }

  public long getCellId() {
    return cellId;
  }

  public int getChannelId() {
    return channelId;
  }

  public String getControllerName() {
    return controllerName;
  }

  public int getBlockIdCellId() {
    return BLOCKIDCELLID;
  }

  public int getBlockIdChannelId() {
    return BLOCKIDCHANNELID;
  }

  public int getBlockIdControllerName() {
    return BLOCKIDCONTROLLERNAME;
  }

  @Override
  protected void serializeOpenProtocolMessageApplicationCommunicationStartAcknowledgeChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1");

    // Const Field (blockIdCellId)
    writeConstField(
        "blockIdCellId",
        BLOCKIDCELLID,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (cellId)
    writeSimpleField(
        "cellId", cellId, writeUnsignedLong(writeBuffer, 32), WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdChannelId)
    writeConstField(
        "blockIdChannelId",
        BLOCKIDCHANNELID,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (channelId)
    writeSimpleField(
        "channelId",
        channelId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdControllerName)
    writeConstField(
        "blockIdControllerName",
        BLOCKIDCONTROLLERNAME,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (controllerName)
    writeSimpleField(
        "controllerName",
        controllerName,
        writeString(writeBuffer, 200),
        WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (blockIdCellId)
    lengthInBits += 16;

    // Simple field (cellId)
    lengthInBits += 32;

    // Const Field (blockIdChannelId)
    lengthInBits += 16;

    // Simple field (channelId)
    lengthInBits += 16;

    // Const Field (blockIdControllerName)
    lengthInBits += 16;

    // Simple field (controllerName)
    lengthInBits += 200;

    return lengthInBits;
  }

  public static OpenProtocolMessageApplicationCommunicationStartAcknowledgeBuilder
      staticParseOpenProtocolMessageApplicationCommunicationStartAcknowledgeBuilder(
          ReadBuffer readBuffer, Long revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockIdCellId =
        readConstField(
            "blockIdCellId",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1.BLOCKIDCELLID,
            WithOption.WithEncoding("ASCII"));

    long cellId =
        readSimpleField(
            "cellId", readUnsignedLong(readBuffer, 32), WithOption.WithEncoding("ASCII"));

    int blockIdChannelId =
        readConstField(
            "blockIdChannelId",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1.BLOCKIDCHANNELID,
            WithOption.WithEncoding("ASCII"));

    int channelId =
        readSimpleField(
            "channelId", readUnsignedInt(readBuffer, 16), WithOption.WithEncoding("ASCII"));

    int blockIdControllerName =
        readConstField(
            "blockIdControllerName",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1.BLOCKIDCONTROLLERNAME,
            WithOption.WithEncoding("ASCII"));

    String controllerName =
        readSimpleField(
            "controllerName", readString(readBuffer, 200), WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1");
    // Create the instance
    return new OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1BuilderImpl(
        cellId, channelId, controllerName);
  }

  public static class OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1BuilderImpl
      implements OpenProtocolMessageApplicationCommunicationStartAcknowledge
          .OpenProtocolMessageApplicationCommunicationStartAcknowledgeBuilder {
    private final long cellId;
    private final int channelId;
    private final String controllerName;

    public OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1BuilderImpl(
        long cellId, int channelId, String controllerName) {
      this.cellId = cellId;
      this.channelId = channelId;
      this.controllerName = controllerName;
    }

    public OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1
          openProtocolMessageApplicationCommunicationStartAcknowledgeRev1 =
              new OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  cellId,
                  channelId,
                  controllerName);
      return openProtocolMessageApplicationCommunicationStartAcknowledgeRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1)) {
      return false;
    }
    OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1 that =
        (OpenProtocolMessageApplicationCommunicationStartAcknowledgeRev1) o;
    return (getCellId() == that.getCellId())
        && (getChannelId() == that.getChannelId())
        && (getControllerName() == that.getControllerName())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCellId(), getChannelId(), getControllerName());
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
