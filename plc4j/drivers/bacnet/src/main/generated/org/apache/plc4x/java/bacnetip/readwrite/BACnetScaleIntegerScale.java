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

public class BACnetScaleIntegerScale extends BACnetScale implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagSignedInteger integerScale;

  public BACnetScaleIntegerScale(
      BACnetTagHeader peekedTagHeader, BACnetContextTagSignedInteger integerScale) {
    super(peekedTagHeader);
    this.integerScale = integerScale;
  }

  public BACnetContextTagSignedInteger getIntegerScale() {
    return integerScale;
  }

  @Override
  protected void serializeBACnetScaleChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetScaleIntegerScale");

    // Simple Field (integerScale)
    writeSimpleField("integerScale", integerScale, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetScaleIntegerScale");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetScaleIntegerScale _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (integerScale)
    lengthInBits += integerScale.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetScaleBuilder staticParseBACnetScaleBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetScaleIntegerScale");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagSignedInteger integerScale =
        readSimpleField(
            "integerScale",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagSignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.SIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetScaleIntegerScale");
    // Create the instance
    return new BACnetScaleIntegerScaleBuilderImpl(integerScale);
  }

  public static class BACnetScaleIntegerScaleBuilderImpl implements BACnetScale.BACnetScaleBuilder {
    private final BACnetContextTagSignedInteger integerScale;

    public BACnetScaleIntegerScaleBuilderImpl(BACnetContextTagSignedInteger integerScale) {
      this.integerScale = integerScale;
    }

    public BACnetScaleIntegerScale build(BACnetTagHeader peekedTagHeader) {
      BACnetScaleIntegerScale bACnetScaleIntegerScale =
          new BACnetScaleIntegerScale(peekedTagHeader, integerScale);
      return bACnetScaleIntegerScale;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetScaleIntegerScale)) {
      return false;
    }
    BACnetScaleIntegerScale that = (BACnetScaleIntegerScale) o;
    return (getIntegerScale() == that.getIntegerScale()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIntegerScale());
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
