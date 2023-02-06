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

public class BACnetFaultParameterFaultExtendedParametersEntryEnumerated
    extends BACnetFaultParameterFaultExtendedParametersEntry implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagEnumerated enumeratedValue;

  public BACnetFaultParameterFaultExtendedParametersEntryEnumerated(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagEnumerated enumeratedValue) {
    super(peekedTagHeader);
    this.enumeratedValue = enumeratedValue;
  }

  public BACnetApplicationTagEnumerated getEnumeratedValue() {
    return enumeratedValue;
  }

  @Override
  protected void serializeBACnetFaultParameterFaultExtendedParametersEntryChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated");

    // Simple Field (enumeratedValue)
    writeSimpleField(
        "enumeratedValue", enumeratedValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetFaultParameterFaultExtendedParametersEntryEnumerated _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (enumeratedValue)
    lengthInBits += enumeratedValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetFaultParameterFaultExtendedParametersEntryBuilder
      staticParseBACnetFaultParameterFaultExtendedParametersEntryBuilder(ReadBuffer readBuffer)
          throws ParseException {
    readBuffer.pullContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagEnumerated enumeratedValue =
        readSimpleField(
            "enumeratedValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagEnumerated) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated");
    // Create the instance
    return new BACnetFaultParameterFaultExtendedParametersEntryEnumeratedBuilderImpl(
        enumeratedValue);
  }

  public static class BACnetFaultParameterFaultExtendedParametersEntryEnumeratedBuilderImpl
      implements BACnetFaultParameterFaultExtendedParametersEntry
          .BACnetFaultParameterFaultExtendedParametersEntryBuilder {
    private final BACnetApplicationTagEnumerated enumeratedValue;

    public BACnetFaultParameterFaultExtendedParametersEntryEnumeratedBuilderImpl(
        BACnetApplicationTagEnumerated enumeratedValue) {
      this.enumeratedValue = enumeratedValue;
    }

    public BACnetFaultParameterFaultExtendedParametersEntryEnumerated build(
        BACnetTagHeader peekedTagHeader) {
      BACnetFaultParameterFaultExtendedParametersEntryEnumerated
          bACnetFaultParameterFaultExtendedParametersEntryEnumerated =
              new BACnetFaultParameterFaultExtendedParametersEntryEnumerated(
                  peekedTagHeader, enumeratedValue);
      return bACnetFaultParameterFaultExtendedParametersEntryEnumerated;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetFaultParameterFaultExtendedParametersEntryEnumerated)) {
      return false;
    }
    BACnetFaultParameterFaultExtendedParametersEntryEnumerated that =
        (BACnetFaultParameterFaultExtendedParametersEntryEnumerated) o;
    return (getEnumeratedValue() == that.getEnumeratedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getEnumeratedValue());
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
