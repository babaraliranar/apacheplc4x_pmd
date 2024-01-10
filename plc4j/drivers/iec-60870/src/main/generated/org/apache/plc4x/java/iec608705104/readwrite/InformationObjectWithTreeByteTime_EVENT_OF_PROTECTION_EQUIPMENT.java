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

public class InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT
    extends InformationObjectWithTreeByteTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.EVENT_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG;
  }

  // Properties.
  protected final TwoOctetBinaryTime cp16Time2a;
  protected final ThreeOctetBinaryTime cp24Time2a;

  public InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT(
      int address, TwoOctetBinaryTime cp16Time2a, ThreeOctetBinaryTime cp24Time2a) {
    super(address);
    this.cp16Time2a = cp16Time2a;
    this.cp24Time2a = cp24Time2a;
  }

  public TwoOctetBinaryTime getCp16Time2a() {
    return cp16Time2a;
  }

  public ThreeOctetBinaryTime getCp24Time2a() {
    return cp24Time2a;
  }

  @Override
  protected void serializeInformationObjectWithTreeByteTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT");

    // Simple Field (cp16Time2a)
    writeSimpleField(
        "cp16Time2a",
        cp16Time2a,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (cp24Time2a)
    writeSimpleField(
        "cp24Time2a",
        cp24Time2a,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (cp16Time2a)
    lengthInBits += cp16Time2a.getLengthInBits();

    // Simple field (cp24Time2a)
    lengthInBits += cp24Time2a.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithTreeByteTimeBuilder
      staticParseInformationObjectWithTreeByteTimeBuilder(
          ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
          throws ParseException {
    readBuffer.pullContext("InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TwoOctetBinaryTime cp16Time2a =
        readSimpleField(
            "cp16Time2a",
            new DataReaderComplexDefault<>(
                () -> TwoOctetBinaryTime.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    ThreeOctetBinaryTime cp24Time2a =
        readSimpleField(
            "cp24Time2a",
            new DataReaderComplexDefault<>(
                () -> ThreeOctetBinaryTime.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT");
    // Create the instance
    return new InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENTBuilderImpl(
        cp16Time2a, cp24Time2a);
  }

  public static class InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENTBuilderImpl
      implements InformationObjectWithTreeByteTime.InformationObjectWithTreeByteTimeBuilder {
    private final TwoOctetBinaryTime cp16Time2a;
    private final ThreeOctetBinaryTime cp24Time2a;

    public InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENTBuilderImpl(
        TwoOctetBinaryTime cp16Time2a, ThreeOctetBinaryTime cp24Time2a) {
      this.cp16Time2a = cp16Time2a;
      this.cp24Time2a = cp24Time2a;
    }

    public InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT build(int address) {
      InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT
          informationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT =
              new InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT(
                  address, cp16Time2a, cp24Time2a);
      return informationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT)) {
      return false;
    }
    InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT that =
        (InformationObjectWithTreeByteTime_EVENT_OF_PROTECTION_EQUIPMENT) o;
    return (getCp16Time2a() == that.getCp16Time2a())
        && (getCp24Time2a() == that.getCp24Time2a())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCp16Time2a(), getCp24Time2a());
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
