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
package org.apache.plc4x.java.cbus.readwrite;

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

public class IdentifyReplyCommandDSIStatus extends IdentifyReplyCommand implements Message {

  // Accessors for discriminator values.
  public Attribute getAttribute() {
    return Attribute.DSIStatus;
  }

  // Properties.
  protected final ChannelStatus channelStatus1;
  protected final ChannelStatus channelStatus2;
  protected final ChannelStatus channelStatus3;
  protected final ChannelStatus channelStatus4;
  protected final ChannelStatus channelStatus5;
  protected final ChannelStatus channelStatus6;
  protected final ChannelStatus channelStatus7;
  protected final ChannelStatus channelStatus8;
  protected final UnitStatus unitStatus;
  protected final byte dimmingUCRevisionNumber;

  // Arguments.
  protected final Short numBytes;

  public IdentifyReplyCommandDSIStatus(
      ChannelStatus channelStatus1,
      ChannelStatus channelStatus2,
      ChannelStatus channelStatus3,
      ChannelStatus channelStatus4,
      ChannelStatus channelStatus5,
      ChannelStatus channelStatus6,
      ChannelStatus channelStatus7,
      ChannelStatus channelStatus8,
      UnitStatus unitStatus,
      byte dimmingUCRevisionNumber,
      Short numBytes) {
    super(numBytes);
    this.channelStatus1 = channelStatus1;
    this.channelStatus2 = channelStatus2;
    this.channelStatus3 = channelStatus3;
    this.channelStatus4 = channelStatus4;
    this.channelStatus5 = channelStatus5;
    this.channelStatus6 = channelStatus6;
    this.channelStatus7 = channelStatus7;
    this.channelStatus8 = channelStatus8;
    this.unitStatus = unitStatus;
    this.dimmingUCRevisionNumber = dimmingUCRevisionNumber;
    this.numBytes = numBytes;
  }

  public ChannelStatus getChannelStatus1() {
    return channelStatus1;
  }

  public ChannelStatus getChannelStatus2() {
    return channelStatus2;
  }

  public ChannelStatus getChannelStatus3() {
    return channelStatus3;
  }

  public ChannelStatus getChannelStatus4() {
    return channelStatus4;
  }

  public ChannelStatus getChannelStatus5() {
    return channelStatus5;
  }

  public ChannelStatus getChannelStatus6() {
    return channelStatus6;
  }

  public ChannelStatus getChannelStatus7() {
    return channelStatus7;
  }

  public ChannelStatus getChannelStatus8() {
    return channelStatus8;
  }

  public UnitStatus getUnitStatus() {
    return unitStatus;
  }

  public byte getDimmingUCRevisionNumber() {
    return dimmingUCRevisionNumber;
  }

  @Override
  protected void serializeIdentifyReplyCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("IdentifyReplyCommandDSIStatus");

    // Simple Field (channelStatus1)
    writeSimpleEnumField(
        "channelStatus1",
        "ChannelStatus",
        channelStatus1,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus2)
    writeSimpleEnumField(
        "channelStatus2",
        "ChannelStatus",
        channelStatus2,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus3)
    writeSimpleEnumField(
        "channelStatus3",
        "ChannelStatus",
        channelStatus3,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus4)
    writeSimpleEnumField(
        "channelStatus4",
        "ChannelStatus",
        channelStatus4,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus5)
    writeSimpleEnumField(
        "channelStatus5",
        "ChannelStatus",
        channelStatus5,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus6)
    writeSimpleEnumField(
        "channelStatus6",
        "ChannelStatus",
        channelStatus6,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus7)
    writeSimpleEnumField(
        "channelStatus7",
        "ChannelStatus",
        channelStatus7,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (channelStatus8)
    writeSimpleEnumField(
        "channelStatus8",
        "ChannelStatus",
        channelStatus8,
        new DataWriterEnumDefault<>(
            ChannelStatus::getValue, ChannelStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (unitStatus)
    writeSimpleEnumField(
        "unitStatus",
        "UnitStatus",
        unitStatus,
        new DataWriterEnumDefault<>(
            UnitStatus::getValue, UnitStatus::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (dimmingUCRevisionNumber)
    writeSimpleField("dimmingUCRevisionNumber", dimmingUCRevisionNumber, writeByte(writeBuffer, 8));

    writeBuffer.popContext("IdentifyReplyCommandDSIStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IdentifyReplyCommandDSIStatus _value = this;

    // Simple field (channelStatus1)
    lengthInBits += 8;

    // Simple field (channelStatus2)
    lengthInBits += 8;

    // Simple field (channelStatus3)
    lengthInBits += 8;

    // Simple field (channelStatus4)
    lengthInBits += 8;

    // Simple field (channelStatus5)
    lengthInBits += 8;

    // Simple field (channelStatus6)
    lengthInBits += 8;

    // Simple field (channelStatus7)
    lengthInBits += 8;

    // Simple field (channelStatus8)
    lengthInBits += 8;

    // Simple field (unitStatus)
    lengthInBits += 8;

    // Simple field (dimmingUCRevisionNumber)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static IdentifyReplyCommandBuilder staticParseIdentifyReplyCommandBuilder(
      ReadBuffer readBuffer, Attribute attribute, Short numBytes) throws ParseException {
    readBuffer.pullContext("IdentifyReplyCommandDSIStatus");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ChannelStatus channelStatus1 =
        readEnumField(
            "channelStatus1",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus2 =
        readEnumField(
            "channelStatus2",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus3 =
        readEnumField(
            "channelStatus3",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus4 =
        readEnumField(
            "channelStatus4",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus5 =
        readEnumField(
            "channelStatus5",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus6 =
        readEnumField(
            "channelStatus6",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus7 =
        readEnumField(
            "channelStatus7",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    ChannelStatus channelStatus8 =
        readEnumField(
            "channelStatus8",
            "ChannelStatus",
            new DataReaderEnumDefault<>(
                ChannelStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    UnitStatus unitStatus =
        readEnumField(
            "unitStatus",
            "UnitStatus",
            new DataReaderEnumDefault<>(
                UnitStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    byte dimmingUCRevisionNumber =
        readSimpleField("dimmingUCRevisionNumber", readByte(readBuffer, 8));

    readBuffer.closeContext("IdentifyReplyCommandDSIStatus");
    // Create the instance
    return new IdentifyReplyCommandDSIStatusBuilderImpl(
        channelStatus1,
        channelStatus2,
        channelStatus3,
        channelStatus4,
        channelStatus5,
        channelStatus6,
        channelStatus7,
        channelStatus8,
        unitStatus,
        dimmingUCRevisionNumber,
        numBytes);
  }

  public static class IdentifyReplyCommandDSIStatusBuilderImpl
      implements IdentifyReplyCommand.IdentifyReplyCommandBuilder {
    private final ChannelStatus channelStatus1;
    private final ChannelStatus channelStatus2;
    private final ChannelStatus channelStatus3;
    private final ChannelStatus channelStatus4;
    private final ChannelStatus channelStatus5;
    private final ChannelStatus channelStatus6;
    private final ChannelStatus channelStatus7;
    private final ChannelStatus channelStatus8;
    private final UnitStatus unitStatus;
    private final byte dimmingUCRevisionNumber;
    private final Short numBytes;

    public IdentifyReplyCommandDSIStatusBuilderImpl(
        ChannelStatus channelStatus1,
        ChannelStatus channelStatus2,
        ChannelStatus channelStatus3,
        ChannelStatus channelStatus4,
        ChannelStatus channelStatus5,
        ChannelStatus channelStatus6,
        ChannelStatus channelStatus7,
        ChannelStatus channelStatus8,
        UnitStatus unitStatus,
        byte dimmingUCRevisionNumber,
        Short numBytes) {
      this.channelStatus1 = channelStatus1;
      this.channelStatus2 = channelStatus2;
      this.channelStatus3 = channelStatus3;
      this.channelStatus4 = channelStatus4;
      this.channelStatus5 = channelStatus5;
      this.channelStatus6 = channelStatus6;
      this.channelStatus7 = channelStatus7;
      this.channelStatus8 = channelStatus8;
      this.unitStatus = unitStatus;
      this.dimmingUCRevisionNumber = dimmingUCRevisionNumber;
      this.numBytes = numBytes;
    }

    public IdentifyReplyCommandDSIStatus build(Short numBytes) {

      IdentifyReplyCommandDSIStatus identifyReplyCommandDSIStatus =
          new IdentifyReplyCommandDSIStatus(
              channelStatus1,
              channelStatus2,
              channelStatus3,
              channelStatus4,
              channelStatus5,
              channelStatus6,
              channelStatus7,
              channelStatus8,
              unitStatus,
              dimmingUCRevisionNumber,
              numBytes);
      return identifyReplyCommandDSIStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IdentifyReplyCommandDSIStatus)) {
      return false;
    }
    IdentifyReplyCommandDSIStatus that = (IdentifyReplyCommandDSIStatus) o;
    return (getChannelStatus1() == that.getChannelStatus1())
        && (getChannelStatus2() == that.getChannelStatus2())
        && (getChannelStatus3() == that.getChannelStatus3())
        && (getChannelStatus4() == that.getChannelStatus4())
        && (getChannelStatus5() == that.getChannelStatus5())
        && (getChannelStatus6() == that.getChannelStatus6())
        && (getChannelStatus7() == that.getChannelStatus7())
        && (getChannelStatus8() == that.getChannelStatus8())
        && (getUnitStatus() == that.getUnitStatus())
        && (getDimmingUCRevisionNumber() == that.getDimmingUCRevisionNumber())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getChannelStatus1(),
        getChannelStatus2(),
        getChannelStatus3(),
        getChannelStatus4(),
        getChannelStatus5(),
        getChannelStatus6(),
        getChannelStatus7(),
        getChannelStatus8(),
        getUnitStatus(),
        getDimmingUCRevisionNumber());
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
