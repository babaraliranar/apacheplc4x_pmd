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

public class HVACHumidityModeAndFlags implements Message {

  // Properties.
  protected final boolean auxiliaryLevel;
  protected final boolean guard;
  protected final boolean setback;
  protected final boolean level;
  protected final HVACHumidityModeAndFlagsMode mode;

  // Reserved Fields
  private Boolean reservedField0;

  public HVACHumidityModeAndFlags(
      boolean auxiliaryLevel,
      boolean guard,
      boolean setback,
      boolean level,
      HVACHumidityModeAndFlagsMode mode) {
    super();
    this.auxiliaryLevel = auxiliaryLevel;
    this.guard = guard;
    this.setback = setback;
    this.level = level;
    this.mode = mode;
  }

  public boolean getAuxiliaryLevel() {
    return auxiliaryLevel;
  }

  public boolean getGuard() {
    return guard;
  }

  public boolean getSetback() {
    return setback;
  }

  public boolean getLevel() {
    return level;
  }

  public HVACHumidityModeAndFlagsMode getMode() {
    return mode;
  }

  public boolean getIsAuxLevelUnused() {
    return (boolean) (!(getAuxiliaryLevel()));
  }

  public boolean getIsAuxLevelUsed() {
    return (boolean) (getAuxiliaryLevel());
  }

  public boolean getIsGuardDisabled() {
    return (boolean) (!(getGuard()));
  }

  public boolean getIsGuardEnabled() {
    return (boolean) (getGuard());
  }

  public boolean getIsSetbackDisabled() {
    return (boolean) (!(getSetback()));
  }

  public boolean getIsSetbackEnabled() {
    return (boolean) (getSetback());
  }

  public boolean getIsLevelHumidity() {
    return (boolean) (!(getLevel()));
  }

  public boolean getIsLevelRaw() {
    return (boolean) (getLevel());
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("HVACHumidityModeAndFlags");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (boolean) false,
        writeBoolean(writeBuffer));

    // Simple Field (auxiliaryLevel)
    writeSimpleField("auxiliaryLevel", auxiliaryLevel, writeBoolean(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isAuxLevelUnused = getIsAuxLevelUnused();
    writeBuffer.writeVirtual("isAuxLevelUnused", isAuxLevelUnused);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isAuxLevelUsed = getIsAuxLevelUsed();
    writeBuffer.writeVirtual("isAuxLevelUsed", isAuxLevelUsed);

    // Simple Field (guard)
    writeSimpleField("guard", guard, writeBoolean(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isGuardDisabled = getIsGuardDisabled();
    writeBuffer.writeVirtual("isGuardDisabled", isGuardDisabled);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isGuardEnabled = getIsGuardEnabled();
    writeBuffer.writeVirtual("isGuardEnabled", isGuardEnabled);

    // Simple Field (setback)
    writeSimpleField("setback", setback, writeBoolean(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isSetbackDisabled = getIsSetbackDisabled();
    writeBuffer.writeVirtual("isSetbackDisabled", isSetbackDisabled);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isSetbackEnabled = getIsSetbackEnabled();
    writeBuffer.writeVirtual("isSetbackEnabled", isSetbackEnabled);

    // Simple Field (level)
    writeSimpleField("level", level, writeBoolean(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isLevelHumidity = getIsLevelHumidity();
    writeBuffer.writeVirtual("isLevelHumidity", isLevelHumidity);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isLevelRaw = getIsLevelRaw();
    writeBuffer.writeVirtual("isLevelRaw", isLevelRaw);

    // Simple Field (mode)
    writeSimpleEnumField(
        "mode",
        "HVACHumidityModeAndFlagsMode",
        mode,
        new DataWriterEnumDefault<>(
            HVACHumidityModeAndFlagsMode::getValue,
            HVACHumidityModeAndFlagsMode::name,
            writeUnsignedByte(writeBuffer, 3)));

    writeBuffer.popContext("HVACHumidityModeAndFlags");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    HVACHumidityModeAndFlags _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (auxiliaryLevel)
    lengthInBits += 1;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (guard)
    lengthInBits += 1;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (setback)
    lengthInBits += 1;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (level)
    lengthInBits += 1;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (mode)
    lengthInBits += 3;

    return lengthInBits;
  }

  public static HVACHumidityModeAndFlags staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static HVACHumidityModeAndFlags staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("HVACHumidityModeAndFlags");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Boolean reservedField0 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    boolean auxiliaryLevel = readSimpleField("auxiliaryLevel", readBoolean(readBuffer));
    boolean isAuxLevelUnused =
        readVirtualField("isAuxLevelUnused", boolean.class, !(auxiliaryLevel));
    boolean isAuxLevelUsed = readVirtualField("isAuxLevelUsed", boolean.class, auxiliaryLevel);

    boolean guard = readSimpleField("guard", readBoolean(readBuffer));
    boolean isGuardDisabled = readVirtualField("isGuardDisabled", boolean.class, !(guard));
    boolean isGuardEnabled = readVirtualField("isGuardEnabled", boolean.class, guard);

    boolean setback = readSimpleField("setback", readBoolean(readBuffer));
    boolean isSetbackDisabled = readVirtualField("isSetbackDisabled", boolean.class, !(setback));
    boolean isSetbackEnabled = readVirtualField("isSetbackEnabled", boolean.class, setback);

    boolean level = readSimpleField("level", readBoolean(readBuffer));
    boolean isLevelHumidity = readVirtualField("isLevelHumidity", boolean.class, !(level));
    boolean isLevelRaw = readVirtualField("isLevelRaw", boolean.class, level);

    HVACHumidityModeAndFlagsMode mode =
        readEnumField(
            "mode",
            "HVACHumidityModeAndFlagsMode",
            new DataReaderEnumDefault<>(
                HVACHumidityModeAndFlagsMode::enumForValue, readUnsignedByte(readBuffer, 3)));

    readBuffer.closeContext("HVACHumidityModeAndFlags");
    // Create the instance
    HVACHumidityModeAndFlags _hVACHumidityModeAndFlags;
    _hVACHumidityModeAndFlags =
        new HVACHumidityModeAndFlags(auxiliaryLevel, guard, setback, level, mode);
    _hVACHumidityModeAndFlags.reservedField0 = reservedField0;
    return _hVACHumidityModeAndFlags;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof HVACHumidityModeAndFlags)) {
      return false;
    }
    HVACHumidityModeAndFlags that = (HVACHumidityModeAndFlags) o;
    return (getAuxiliaryLevel() == that.getAuxiliaryLevel())
        && (getGuard() == that.getGuard())
        && (getSetback() == that.getSetback())
        && (getLevel() == that.getLevel())
        && (getMode() == that.getMode())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getAuxiliaryLevel(), getGuard(), getSetback(), getLevel(), getMode());
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
