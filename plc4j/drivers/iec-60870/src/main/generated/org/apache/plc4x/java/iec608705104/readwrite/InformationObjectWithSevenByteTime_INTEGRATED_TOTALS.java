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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObjectWithSevenByteTime_INTEGRATED_TOTALS
    extends InformationObjectWithSevenByteTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.INTEGRATED_TOTALS_WITH_TIME_TAG_CP56TIME2A;
  }

  // Properties.
  protected final BinaryCounterReading bcr;
  protected final SevenOctetBinaryTime cp56Time2a;

  public InformationObjectWithSevenByteTime_INTEGRATED_TOTALS(
      int address, BinaryCounterReading bcr, SevenOctetBinaryTime cp56Time2a) {
    super(address);
    this.bcr = bcr;
    this.cp56Time2a = cp56Time2a;
  }

  public BinaryCounterReading getBcr() {
    return bcr;
  }

  public SevenOctetBinaryTime getCp56Time2a() {
    return cp56Time2a;
  }

  @Override
  protected void serializeInformationObjectWithSevenByteTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithSevenByteTime_INTEGRATED_TOTALS");

    // Simple Field (bcr)
    writeSimpleField(
        "bcr",
        bcr,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (cp56Time2a)
    writeSimpleField(
        "cp56Time2a",
        cp56Time2a,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithSevenByteTime_INTEGRATED_TOTALS");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithSevenByteTime_INTEGRATED_TOTALS _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (bcr)
    lengthInBits += bcr.getLengthInBits();

    // Simple field (cp56Time2a)
    lengthInBits += cp56Time2a.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithSevenByteTimeBuilder
      staticParseInformationObjectWithSevenByteTimeBuilder(
          ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
          throws ParseException {
    readBuffer.pullContext("InformationObjectWithSevenByteTime_INTEGRATED_TOTALS");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BinaryCounterReading bcr =
        readSimpleField(
            "bcr",
            new DataReaderComplexDefault<>(
                () -> BinaryCounterReading.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    SevenOctetBinaryTime cp56Time2a =
        readSimpleField(
            "cp56Time2a",
            new DataReaderComplexDefault<>(
                () -> SevenOctetBinaryTime.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithSevenByteTime_INTEGRATED_TOTALS");
    // Create the instance
    return new InformationObjectWithSevenByteTime_INTEGRATED_TOTALSBuilderImpl(bcr, cp56Time2a);
  }

  public static class InformationObjectWithSevenByteTime_INTEGRATED_TOTALSBuilderImpl
      implements InformationObjectWithSevenByteTime.InformationObjectWithSevenByteTimeBuilder {
    private final BinaryCounterReading bcr;
    private final SevenOctetBinaryTime cp56Time2a;

    public InformationObjectWithSevenByteTime_INTEGRATED_TOTALSBuilderImpl(
        BinaryCounterReading bcr, SevenOctetBinaryTime cp56Time2a) {
      this.bcr = bcr;
      this.cp56Time2a = cp56Time2a;
    }

    public InformationObjectWithSevenByteTime_INTEGRATED_TOTALS build(int address) {
      InformationObjectWithSevenByteTime_INTEGRATED_TOTALS
          informationObjectWithSevenByteTime_INTEGRATED_TOTALS =
              new InformationObjectWithSevenByteTime_INTEGRATED_TOTALS(address, bcr, cp56Time2a);
      return informationObjectWithSevenByteTime_INTEGRATED_TOTALS;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithSevenByteTime_INTEGRATED_TOTALS)) {
      return false;
    }
    InformationObjectWithSevenByteTime_INTEGRATED_TOTALS that =
        (InformationObjectWithSevenByteTime_INTEGRATED_TOTALS) o;
    return (getBcr() == that.getBcr())
        && (getCp56Time2a() == that.getCp56Time2a())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBcr(), getCp56Time2a());
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
