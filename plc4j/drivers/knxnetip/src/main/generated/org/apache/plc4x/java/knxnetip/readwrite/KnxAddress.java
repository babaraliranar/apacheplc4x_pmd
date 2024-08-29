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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class KnxAddress implements Message {

  // Properties.
  protected final byte mainGroup;
  protected final byte middleGroup;
  protected final short subGroup;

  public KnxAddress(byte mainGroup, byte middleGroup, short subGroup) {
    super();
    this.mainGroup = mainGroup;
    this.middleGroup = middleGroup;
    this.subGroup = subGroup;
  }

  public byte getMainGroup() {
    return mainGroup;
  }

  public byte getMiddleGroup() {
    return middleGroup;
  }

  public short getSubGroup() {
    return subGroup;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("KnxAddress");

    // Simple Field (mainGroup)
    writeSimpleField("mainGroup", mainGroup, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (middleGroup)
    writeSimpleField("middleGroup", middleGroup, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (subGroup)
    writeSimpleField("subGroup", subGroup, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("KnxAddress");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    KnxAddress _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (mainGroup)
    lengthInBits += 4;

    // Simple field (middleGroup)
    lengthInBits += 4;

    // Simple field (subGroup)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static KnxAddress staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("KnxAddress");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte mainGroup = readSimpleField("mainGroup", readUnsignedByte(readBuffer, 4));

    byte middleGroup = readSimpleField("middleGroup", readUnsignedByte(readBuffer, 4));

    short subGroup = readSimpleField("subGroup", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("KnxAddress");
    // Create the instance
    KnxAddress _knxAddress;
    _knxAddress = new KnxAddress(mainGroup, middleGroup, subGroup);
    return _knxAddress;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof KnxAddress)) {
      return false;
    }
    KnxAddress that = (KnxAddress) o;
    return (getMainGroup() == that.getMainGroup())
        && (getMiddleGroup() == that.getMiddleGroup())
        && (getSubGroup() == that.getSubGroup())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getMainGroup(), getMiddleGroup(), getSubGroup());
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
