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

public class OpenProtocolMessageResultTracesCurveRev1 extends OpenProtocolMessageResultTracesCurve
    implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 1;
  }

  public OpenProtocolMessageResultTracesCurveRev1(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
  }

  @Override
  protected void serializeOpenProtocolMessageResultTracesCurveChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageResultTracesCurveRev1");

    writeBuffer.popContext("OpenProtocolMessageResultTracesCurveRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageResultTracesCurveRev1 _value = this;

    return lengthInBits;
  }

  public static OpenProtocolMessageResultTracesCurveBuilder
      staticParseOpenProtocolMessageResultTracesCurveBuilder(ReadBuffer readBuffer, Long revision)
          throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageResultTracesCurveRev1");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    readBuffer.closeContext("OpenProtocolMessageResultTracesCurveRev1");
    // Create the instance
    return new OpenProtocolMessageResultTracesCurveRev1BuilderImpl();
  }

  public static class OpenProtocolMessageResultTracesCurveRev1BuilderImpl
      implements OpenProtocolMessageResultTracesCurve.OpenProtocolMessageResultTracesCurveBuilder {

    public OpenProtocolMessageResultTracesCurveRev1BuilderImpl() {}

    public OpenProtocolMessageResultTracesCurveRev1 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageResultTracesCurveRev1 openProtocolMessageResultTracesCurveRev1 =
          new OpenProtocolMessageResultTracesCurveRev1(
              midRevision,
              noAckFlag,
              targetStationId,
              targetSpindleId,
              sequenceNumber,
              numberOfMessageParts,
              messagePartNumber);
      return openProtocolMessageResultTracesCurveRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageResultTracesCurveRev1)) {
      return false;
    }
    OpenProtocolMessageResultTracesCurveRev1 that = (OpenProtocolMessageResultTracesCurveRev1) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
