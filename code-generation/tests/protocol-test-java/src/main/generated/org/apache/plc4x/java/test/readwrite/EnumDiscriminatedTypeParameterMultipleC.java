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
package org.apache.plc4x.java.test.readwrite;

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

public class EnumDiscriminatedTypeParameterMultipleC extends EnumDiscriminatedTypeParameterMultiple
    implements Message {

  // Accessors for discriminator values.
  public EnumType getDiscr1() {
    return EnumType.INT;
  }

  public EnumTypeInt getDiscr2() {
    return EnumTypeInt.INTINT;
  }

  // Properties.
  protected final short simpC;

  public EnumDiscriminatedTypeParameterMultipleC(short simpC) {
    super();
    this.simpC = simpC;
  }

  public short getSimpC() {
    return simpC;
  }

  @Override
  protected void serializeEnumDiscriminatedTypeParameterMultipleChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("EnumDiscriminatedTypeParameterMultipleC");

    // Simple Field (simpC)
    writeSimpleField("simpC", simpC, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("EnumDiscriminatedTypeParameterMultipleC");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EnumDiscriminatedTypeParameterMultipleC _value = this;

    // Simple field (simpC)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static EnumDiscriminatedTypeParameterMultipleBuilder
      staticParseEnumDiscriminatedTypeParameterMultipleBuilder(
          ReadBuffer readBuffer, EnumType discr1, EnumTypeInt discr2) throws ParseException {
    readBuffer.pullContext("EnumDiscriminatedTypeParameterMultipleC");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short simpC = readSimpleField("simpC", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("EnumDiscriminatedTypeParameterMultipleC");
    // Create the instance
    return new EnumDiscriminatedTypeParameterMultipleCBuilderImpl(simpC);
  }

  public static class EnumDiscriminatedTypeParameterMultipleCBuilderImpl
      implements EnumDiscriminatedTypeParameterMultiple
          .EnumDiscriminatedTypeParameterMultipleBuilder {
    private final short simpC;

    public EnumDiscriminatedTypeParameterMultipleCBuilderImpl(short simpC) {

      this.simpC = simpC;
    }

    public EnumDiscriminatedTypeParameterMultipleC build() {
      EnumDiscriminatedTypeParameterMultipleC enumDiscriminatedTypeParameterMultipleC =
          new EnumDiscriminatedTypeParameterMultipleC(simpC);
      return enumDiscriminatedTypeParameterMultipleC;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EnumDiscriminatedTypeParameterMultipleC)) {
      return false;
    }
    EnumDiscriminatedTypeParameterMultipleC that = (EnumDiscriminatedTypeParameterMultipleC) o;
    return (getSimpC() == that.getSimpC()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSimpC());
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
