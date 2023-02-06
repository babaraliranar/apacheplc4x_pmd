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
package org.apache.plc4x.java.simulated.readwrite;

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

public class Dummy implements Message {

  // Properties.
  protected final int dummy;

  public Dummy(int dummy) {
    super();
    this.dummy = dummy;
  }

  public int getDummy() {
    return dummy;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("Dummy");

    // Simple Field (dummy)
    writeSimpleField(
        "dummy",
        dummy,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("Dummy");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Dummy _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dummy)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static Dummy staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static Dummy staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Dummy");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int dummy =
        readSimpleField(
            "dummy",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("Dummy");
    // Create the instance
    Dummy _dummy;
    _dummy = new Dummy(dummy);
    return _dummy;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Dummy)) {
      return false;
    }
    Dummy that = (Dummy) o;
    return (getDummy() == that.getDummy()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getDummy());
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
