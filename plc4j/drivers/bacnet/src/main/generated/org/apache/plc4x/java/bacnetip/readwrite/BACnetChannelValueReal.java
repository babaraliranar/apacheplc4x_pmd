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

public class BACnetChannelValueReal extends BACnetChannelValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagReal realValue;

  public BACnetChannelValueReal(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagReal realValue) {
    super(peekedTagHeader);
    this.realValue = realValue;
  }

  public BACnetApplicationTagReal getRealValue() {
    return realValue;
  }

  @Override
  protected void serializeBACnetChannelValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetChannelValueReal");

    // Simple Field (realValue)
    writeSimpleField("realValue", realValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetChannelValueReal");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetChannelValueReal _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (realValue)
    lengthInBits += realValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetChannelValueBuilder staticParseBACnetChannelValueBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetChannelValueReal");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagReal realValue =
        readSimpleField(
            "realValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagReal) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetChannelValueReal");
    // Create the instance
    return new BACnetChannelValueRealBuilderImpl(realValue);
  }

  public static class BACnetChannelValueRealBuilderImpl
      implements BACnetChannelValue.BACnetChannelValueBuilder {
    private final BACnetApplicationTagReal realValue;

    public BACnetChannelValueRealBuilderImpl(BACnetApplicationTagReal realValue) {
      this.realValue = realValue;
    }

    public BACnetChannelValueReal build(BACnetTagHeader peekedTagHeader) {
      BACnetChannelValueReal bACnetChannelValueReal =
          new BACnetChannelValueReal(peekedTagHeader, realValue);
      return bACnetChannelValueReal;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetChannelValueReal)) {
      return false;
    }
    BACnetChannelValueReal that = (BACnetChannelValueReal) o;
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
