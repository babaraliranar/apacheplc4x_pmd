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

public class FirmataCommandSetPinMode extends FirmataCommand implements Message {

  // Accessors for discriminator values.
  public Byte getCommandCode() {
    return (byte) 0x4;
  }

  // Properties.
  protected final short pin;
  protected final PinMode mode;

  public FirmataCommandSetPinMode(short pin, PinMode mode) {
    super();
    this.pin = pin;
    this.mode = mode;
  }

  public short getPin() {
    return pin;
  }

  public PinMode getMode() {
    return mode;
  }

  @Override
  protected void serializeFirmataCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("FirmataCommandSetPinMode");

    // Simple Field (pin)
    writeSimpleField("pin", pin, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (mode)
    writeSimpleEnumField(
        "mode",
        "PinMode",
        mode,
        new DataWriterEnumDefault<>(
            PinMode::getValue, PinMode::name, writeUnsignedShort(writeBuffer, 8)));

    writeBuffer.popContext("FirmataCommandSetPinMode");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FirmataCommandSetPinMode _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (pin)
    lengthInBits += 8;

    // Simple field (mode)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static FirmataCommandBuilder staticParseFirmataCommandBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("FirmataCommandSetPinMode");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short pin = readSimpleField("pin", readUnsignedShort(readBuffer, 8));

    PinMode mode =
        readEnumField(
            "mode",
            "PinMode",
            new DataReaderEnumDefault<>(PinMode::enumForValue, readUnsignedShort(readBuffer, 8)));

    readBuffer.closeContext("FirmataCommandSetPinMode");
    // Create the instance
    return new FirmataCommandSetPinModeBuilderImpl(pin, mode);
  }

  public static class FirmataCommandSetPinModeBuilderImpl
      implements FirmataCommand.FirmataCommandBuilder {
    private final short pin;
    private final PinMode mode;

    public FirmataCommandSetPinModeBuilderImpl(short pin, PinMode mode) {
      this.pin = pin;
      this.mode = mode;
    }

    public FirmataCommandSetPinMode build() {
      FirmataCommandSetPinMode firmataCommandSetPinMode = new FirmataCommandSetPinMode(pin, mode);
      return firmataCommandSetPinMode;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FirmataCommandSetPinMode)) {
      return false;
    }
    FirmataCommandSetPinMode that = (FirmataCommandSetPinMode) o;
    return (getPin() == that.getPin())
        && (getMode() == that.getMode())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPin(), getMode());
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
