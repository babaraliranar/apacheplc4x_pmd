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
package org.apache.plc4x.java.cbus.readwrite;

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

public class InterfaceOptions2 implements Message {

  // Properties.
  protected final boolean burden;
  protected final boolean clockGen;

  // Reserved Fields
  private Boolean reservedField0;
  private Boolean reservedField1;
  private Boolean reservedField2;
  private Boolean reservedField3;
  private Boolean reservedField4;
  private Boolean reservedField5;

  public InterfaceOptions2(boolean burden, boolean clockGen) {
    super();
    this.burden = burden;
    this.clockGen = clockGen;
  }

  public boolean getBurden() {
    return burden;
  }

  public boolean getClockGen() {
    return clockGen;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("InterfaceOptions2");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (boolean) false,
        writeBoolean(writeBuffer));

    // Simple Field (burden)
    writeSimpleField("burden", burden, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (boolean) false,
        writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField2 != null ? reservedField2 : (boolean) false,
        writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField3 != null ? reservedField3 : (boolean) false,
        writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField4 != null ? reservedField4 : (boolean) false,
        writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField5 != null ? reservedField5 : (boolean) false,
        writeBoolean(writeBuffer));

    // Simple Field (clockGen)
    writeSimpleField("clockGen", clockGen, writeBoolean(writeBuffer));

    writeBuffer.popContext("InterfaceOptions2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    InterfaceOptions2 _value = this;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (burden)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (clockGen)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static InterfaceOptions2 staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static InterfaceOptions2 staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("InterfaceOptions2");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Boolean reservedField0 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    boolean burden = readSimpleField("burden", readBoolean(readBuffer));

    Boolean reservedField1 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    Boolean reservedField2 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    Boolean reservedField3 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    Boolean reservedField4 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    Boolean reservedField5 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    boolean clockGen = readSimpleField("clockGen", readBoolean(readBuffer));

    readBuffer.closeContext("InterfaceOptions2");
    // Create the instance
    InterfaceOptions2 _interfaceOptions2;
    _interfaceOptions2 = new InterfaceOptions2(burden, clockGen);
    _interfaceOptions2.reservedField0 = reservedField0;
    _interfaceOptions2.reservedField1 = reservedField1;
    _interfaceOptions2.reservedField2 = reservedField2;
    _interfaceOptions2.reservedField3 = reservedField3;
    _interfaceOptions2.reservedField4 = reservedField4;
    _interfaceOptions2.reservedField5 = reservedField5;
    return _interfaceOptions2;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InterfaceOptions2)) {
      return false;
    }
    InterfaceOptions2 that = (InterfaceOptions2) o;
    return (getBurden() == that.getBurden()) && (getClockGen() == that.getClockGen()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getBurden(), getClockGen());
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
