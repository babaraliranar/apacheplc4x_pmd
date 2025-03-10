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

public class BinaryCounterReading implements Message {

  // Properties.
  protected final long counterValue;
  protected final boolean counterValid;
  protected final boolean counterAdjusted;
  protected final boolean carry;
  protected final byte sequenceNumber;

  public BinaryCounterReading(
      long counterValue,
      boolean counterValid,
      boolean counterAdjusted,
      boolean carry,
      byte sequenceNumber) {
    super();
    this.counterValue = counterValue;
    this.counterValid = counterValid;
    this.counterAdjusted = counterAdjusted;
    this.carry = carry;
    this.sequenceNumber = sequenceNumber;
  }

  public long getCounterValue() {
    return counterValue;
  }

  public boolean getCounterValid() {
    return counterValid;
  }

  public boolean getCounterAdjusted() {
    return counterAdjusted;
  }

  public boolean getCarry() {
    return carry;
  }

  public byte getSequenceNumber() {
    return sequenceNumber;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BinaryCounterReading");

    // Simple Field (counterValue)
    writeSimpleField(
        "counterValue",
        counterValue,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (counterValid)
    writeSimpleField(
        "counterValid",
        counterValid,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (counterAdjusted)
    writeSimpleField(
        "counterAdjusted",
        counterAdjusted,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (carry)
    writeSimpleField(
        "carry",
        carry,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (sequenceNumber)
    writeSimpleField(
        "sequenceNumber",
        sequenceNumber,
        writeUnsignedByte(writeBuffer, 5),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("BinaryCounterReading");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BinaryCounterReading _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (counterValue)
    lengthInBits += 32;

    // Simple field (counterValid)
    lengthInBits += 1;

    // Simple field (counterAdjusted)
    lengthInBits += 1;

    // Simple field (carry)
    lengthInBits += 1;

    // Simple field (sequenceNumber)
    lengthInBits += 5;

    return lengthInBits;
  }

  public static BinaryCounterReading staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BinaryCounterReading");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long counterValue =
        readSimpleField(
            "counterValue",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean counterValid =
        readSimpleField(
            "counterValid",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean counterAdjusted =
        readSimpleField(
            "counterAdjusted",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean carry =
        readSimpleField(
            "carry", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    byte sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            readUnsignedByte(readBuffer, 5),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("BinaryCounterReading");
    // Create the instance
    BinaryCounterReading _binaryCounterReading;
    _binaryCounterReading =
        new BinaryCounterReading(
            counterValue, counterValid, counterAdjusted, carry, sequenceNumber);
    return _binaryCounterReading;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BinaryCounterReading)) {
      return false;
    }
    BinaryCounterReading that = (BinaryCounterReading) o;
    return (getCounterValue() == that.getCounterValue())
        && (getCounterValid() == that.getCounterValid())
        && (getCounterAdjusted() == that.getCounterAdjusted())
        && (getCarry() == that.getCarry())
        && (getSequenceNumber() == that.getSequenceNumber())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getCounterValue(),
        getCounterValid(),
        getCounterAdjusted(),
        getCarry(),
        getSequenceNumber());
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
