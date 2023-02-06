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

public class OpenProtocolMessageSetTimeRev1 extends OpenProtocolMessageSetTime implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 1;
  }

  // Properties.
  protected final String timeToSet;

  public OpenProtocolMessageSetTimeRev1(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      String timeToSet) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.timeToSet = timeToSet;
  }

  public String getTimeToSet() {
    return timeToSet;
  }

  @Override
  protected void serializeOpenProtocolMessageSetTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageSetTimeRev1");

    // Simple Field (timeToSet)
    writeSimpleField(
        "timeToSet", timeToSet, writeString(writeBuffer, 152), WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageSetTimeRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageSetTimeRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (timeToSet)
    lengthInBits += 152;

    return lengthInBits;
  }

  public static OpenProtocolMessageSetTimeBuilder staticParseOpenProtocolMessageSetTimeBuilder(
      ReadBuffer readBuffer, Long revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageSetTimeRev1");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    String timeToSet =
        readSimpleField("timeToSet", readString(readBuffer, 152), WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageSetTimeRev1");
    // Create the instance
    return new OpenProtocolMessageSetTimeRev1BuilderImpl(timeToSet);
  }

  public static class OpenProtocolMessageSetTimeRev1BuilderImpl
      implements OpenProtocolMessageSetTime.OpenProtocolMessageSetTimeBuilder {
    private final String timeToSet;

    public OpenProtocolMessageSetTimeRev1BuilderImpl(String timeToSet) {
      this.timeToSet = timeToSet;
    }

    public OpenProtocolMessageSetTimeRev1 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageSetTimeRev1 openProtocolMessageSetTimeRev1 =
          new OpenProtocolMessageSetTimeRev1(
              midRevision,
              noAckFlag,
              targetStationId,
              targetSpindleId,
              sequenceNumber,
              numberOfMessageParts,
              messagePartNumber,
              timeToSet);
      return openProtocolMessageSetTimeRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageSetTimeRev1)) {
      return false;
    }
    OpenProtocolMessageSetTimeRev1 that = (OpenProtocolMessageSetTimeRev1) o;
    return (getTimeToSet() == that.getTimeToSet()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTimeToSet());
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
