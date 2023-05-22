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

public class BACnetLogDataLogDataEntryRealValue extends BACnetLogDataLogDataEntry
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagReal realValue;

  public BACnetLogDataLogDataEntryRealValue(
      BACnetTagHeader peekedTagHeader, BACnetContextTagReal realValue) {
    super(peekedTagHeader);
    this.realValue = realValue;
  }

  public BACnetContextTagReal getRealValue() {
    return realValue;
  }

  @Override
  protected void serializeBACnetLogDataLogDataEntryChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLogDataLogDataEntryRealValue");

    // Simple Field (realValue)
    writeSimpleField("realValue", realValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogDataLogDataEntryRealValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogDataLogDataEntryRealValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (realValue)
    lengthInBits += realValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogDataLogDataEntryBuilder staticParseBACnetLogDataLogDataEntryBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLogDataLogDataEntryRealValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagReal realValue =
        readSimpleField(
            "realValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (1), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    readBuffer.closeContext("BACnetLogDataLogDataEntryRealValue");
    // Create the instance
    return new BACnetLogDataLogDataEntryRealValueBuilderImpl(realValue);
  }

  public static class BACnetLogDataLogDataEntryRealValueBuilderImpl
      implements BACnetLogDataLogDataEntry.BACnetLogDataLogDataEntryBuilder {
    private final BACnetContextTagReal realValue;

    public BACnetLogDataLogDataEntryRealValueBuilderImpl(BACnetContextTagReal realValue) {
      this.realValue = realValue;
    }

    public BACnetLogDataLogDataEntryRealValue build(BACnetTagHeader peekedTagHeader) {
      BACnetLogDataLogDataEntryRealValue bACnetLogDataLogDataEntryRealValue =
          new BACnetLogDataLogDataEntryRealValue(peekedTagHeader, realValue);
      return bACnetLogDataLogDataEntryRealValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogDataLogDataEntryRealValue)) {
      return false;
    }
    BACnetLogDataLogDataEntryRealValue that = (BACnetLogDataLogDataEntryRealValue) o;
    return (getRealValue() == that.getRealValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRealValue());
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
