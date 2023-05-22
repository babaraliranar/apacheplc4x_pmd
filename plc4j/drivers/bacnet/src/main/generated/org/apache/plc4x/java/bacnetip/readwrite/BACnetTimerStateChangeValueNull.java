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

public class BACnetTimerStateChangeValueNull extends BACnetTimerStateChangeValue
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagNull nullValue;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetTimerStateChangeValueNull(
      BACnetTagHeader peekedTagHeader,
      BACnetApplicationTagNull nullValue,
      BACnetObjectType objectTypeArgument) {
    super(peekedTagHeader, objectTypeArgument);
    this.nullValue = nullValue;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetApplicationTagNull getNullValue() {
    return nullValue;
  }

  @Override
  protected void serializeBACnetTimerStateChangeValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetTimerStateChangeValueNull");

    // Simple Field (nullValue)
    writeSimpleField("nullValue", nullValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetTimerStateChangeValueNull");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetTimerStateChangeValueNull _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nullValue)
    lengthInBits += nullValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetTimerStateChangeValueBuilder staticParseBACnetTimerStateChangeValueBuilder(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetTimerStateChangeValueNull");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagNull nullValue =
        readSimpleField(
            "nullValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagNull) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetTimerStateChangeValueNull");
    // Create the instance
    return new BACnetTimerStateChangeValueNullBuilderImpl(nullValue, objectTypeArgument);
  }

  public static class BACnetTimerStateChangeValueNullBuilderImpl
      implements BACnetTimerStateChangeValue.BACnetTimerStateChangeValueBuilder {
    private final BACnetApplicationTagNull nullValue;
    private final BACnetObjectType objectTypeArgument;

    public BACnetTimerStateChangeValueNullBuilderImpl(
        BACnetApplicationTagNull nullValue, BACnetObjectType objectTypeArgument) {
      this.nullValue = nullValue;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetTimerStateChangeValueNull build(
        BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument) {
      BACnetTimerStateChangeValueNull bACnetTimerStateChangeValueNull =
          new BACnetTimerStateChangeValueNull(peekedTagHeader, nullValue, objectTypeArgument);
      return bACnetTimerStateChangeValueNull;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTimerStateChangeValueNull)) {
      return false;
    }
    BACnetTimerStateChangeValueNull that = (BACnetTimerStateChangeValueNull) o;
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
