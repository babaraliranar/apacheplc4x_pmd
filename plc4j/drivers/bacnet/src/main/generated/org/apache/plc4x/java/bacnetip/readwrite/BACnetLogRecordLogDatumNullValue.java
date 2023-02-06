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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetLogRecordLogDatumNullValue extends BACnetLogRecordLogDatum implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagNull nullValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetLogRecordLogDatumNullValue(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetContextTagNull nullValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.nullValue = nullValue;
    this.tagNumber = tagNumber;
  }

  public BACnetContextTagNull getNullValue() {
    return nullValue;
  }

  @Override
  protected void serializeBACnetLogRecordLogDatumChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetLogRecordLogDatumNullValue");

    // Simple Field (nullValue)
    writeSimpleField("nullValue", nullValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogRecordLogDatumNullValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogRecordLogDatumNullValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nullValue)
    lengthInBits += nullValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogRecordLogDatumBuilder staticParseBACnetLogRecordLogDatumBuilder(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetLogRecordLogDatumNullValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagNull nullValue =
        readSimpleField(
            "nullValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagNull)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (7), (BACnetDataType) (BACnetDataType.NULL)),
                readBuffer));

    readBuffer.closeContext("BACnetLogRecordLogDatumNullValue");
    // Create the instance
    return new BACnetLogRecordLogDatumNullValueBuilderImpl(nullValue, tagNumber);
  }

  public static class BACnetLogRecordLogDatumNullValueBuilderImpl
      implements BACnetLogRecordLogDatum.BACnetLogRecordLogDatumBuilder {
    private final BACnetContextTagNull nullValue;
    private final Short tagNumber;

    public BACnetLogRecordLogDatumNullValueBuilderImpl(
        BACnetContextTagNull nullValue, Short tagNumber) {
      this.nullValue = nullValue;
      this.tagNumber = tagNumber;
    }

    public BACnetLogRecordLogDatumNullValue build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetLogRecordLogDatumNullValue bACnetLogRecordLogDatumNullValue =
          new BACnetLogRecordLogDatumNullValue(
              openingTag, peekedTagHeader, closingTag, nullValue, tagNumber);
      return bACnetLogRecordLogDatumNullValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogRecordLogDatumNullValue)) {
      return false;
    }
    BACnetLogRecordLogDatumNullValue that = (BACnetLogRecordLogDatumNullValue) o;
    return (getNullValue() == that.getNullValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNullValue());
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
