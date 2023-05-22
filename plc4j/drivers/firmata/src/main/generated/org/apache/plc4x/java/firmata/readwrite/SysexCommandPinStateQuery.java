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
package org.apache.plc4x.java.firmata.readwrite;

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

public class SysexCommandPinStateQuery extends SysexCommand implements Message {

  // Accessors for discriminator values.
  public Short getCommandType() {
    return (short) 0x6D;
  }

  public Boolean getResponse() {
    return false;
  }

  // Properties.
  protected final short pin;

  public SysexCommandPinStateQuery(short pin) {
    super();
    this.pin = pin;
  }

  public short getPin() {
    return pin;
  }

  @Override
  protected void serializeSysexCommandChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SysexCommandPinStateQuery");

    // Simple Field (pin)
    writeSimpleField("pin", pin, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("SysexCommandPinStateQuery");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SysexCommandPinStateQuery _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (pin)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static SysexCommandBuilder staticParseSysexCommandBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("SysexCommandPinStateQuery");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short pin = readSimpleField("pin", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("SysexCommandPinStateQuery");
    // Create the instance
    return new SysexCommandPinStateQueryBuilderImpl(pin);
  }

  public static class SysexCommandPinStateQueryBuilderImpl
      implements SysexCommand.SysexCommandBuilder {
    private final short pin;

    public SysexCommandPinStateQueryBuilderImpl(short pin) {
      this.pin = pin;
    }

    public SysexCommandPinStateQuery build() {
      SysexCommandPinStateQuery sysexCommandPinStateQuery = new SysexCommandPinStateQuery(pin);
      return sysexCommandPinStateQuery;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SysexCommandPinStateQuery)) {
      return false;
    }
    SysexCommandPinStateQuery that = (SysexCommandPinStateQuery) o;
    return (getPin() == that.getPin()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPin());
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
