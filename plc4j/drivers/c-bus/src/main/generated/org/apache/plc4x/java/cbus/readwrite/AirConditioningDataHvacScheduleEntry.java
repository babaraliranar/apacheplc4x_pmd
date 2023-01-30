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

public class AirConditioningDataHvacScheduleEntry extends AirConditioningData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte zoneGroup;
  protected final HVACZoneList zoneList;
  protected final short entry;
  protected final byte format;
  protected final HVACModeAndFlags hvacModeAndFlags;
  protected final HVACStartTime startTime;
  protected final HVACTemperature level;
  protected final HVACRawLevels rawLevel;

  public AirConditioningDataHvacScheduleEntry(
      AirConditioningCommandTypeContainer commandTypeContainer,
      byte zoneGroup,
      HVACZoneList zoneList,
      short entry,
      byte format,
      HVACModeAndFlags hvacModeAndFlags,
      HVACStartTime startTime,
      HVACTemperature level,
      HVACRawLevels rawLevel) {
    super(commandTypeContainer);
    this.zoneGroup = zoneGroup;
    this.zoneList = zoneList;
    this.entry = entry;
    this.format = format;
    this.hvacModeAndFlags = hvacModeAndFlags;
    this.startTime = startTime;
    this.level = level;
    this.rawLevel = rawLevel;
  }

  public byte getZoneGroup() {
    return zoneGroup;
  }

  public HVACZoneList getZoneList() {
    return zoneList;
  }

  public short getEntry() {
    return entry;
  }

  public byte getFormat() {
    return format;
  }

  public HVACModeAndFlags getHvacModeAndFlags() {
    return hvacModeAndFlags;
  }

  public HVACStartTime getStartTime() {
    return startTime;
  }

  public HVACTemperature getLevel() {
    return level;
  }

  public HVACRawLevels getRawLevel() {
    return rawLevel;
  }

  @Override
  protected void serializeAirConditioningDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AirConditioningDataHvacScheduleEntry");

    // Simple Field (zoneGroup)
    writeSimpleField("zoneGroup", zoneGroup, writeByte(writeBuffer, 8));

    // Simple Field (zoneList)
    writeSimpleField("zoneList", zoneList, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (entry)
    writeSimpleField("entry", entry, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (format)
    writeSimpleField("format", format, writeByte(writeBuffer, 8));

    // Simple Field (hvacModeAndFlags)
    writeSimpleField(
        "hvacModeAndFlags", hvacModeAndFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (level) (Can be skipped, if the value is null)
    writeOptionalField(
        "level",
        level,
        new DataWriterComplexDefault<>(writeBuffer),
        getHvacModeAndFlags().getIsLevelTemperature());

    // Optional Field (rawLevel) (Can be skipped, if the value is null)
    writeOptionalField(
        "rawLevel",
        rawLevel,
        new DataWriterComplexDefault<>(writeBuffer),
        getHvacModeAndFlags().getIsLevelRaw());

    writeBuffer.popContext("AirConditioningDataHvacScheduleEntry");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AirConditioningDataHvacScheduleEntry _value = this;

    // Simple field (zoneGroup)
    lengthInBits += 8;

    // Simple field (zoneList)
    lengthInBits += zoneList.getLengthInBits();

    // Simple field (entry)
    lengthInBits += 8;

    // Simple field (format)
    lengthInBits += 8;

    // Simple field (hvacModeAndFlags)
    lengthInBits += hvacModeAndFlags.getLengthInBits();

    // Simple field (startTime)
    lengthInBits += startTime.getLengthInBits();

    // Optional Field (level)
    if (level != null) {
      lengthInBits += level.getLengthInBits();
    }

    // Optional Field (rawLevel)
    if (rawLevel != null) {
      lengthInBits += rawLevel.getLengthInBits();
    }

    return lengthInBits;
  }

  public static AirConditioningDataBuilder staticParseAirConditioningDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AirConditioningDataHvacScheduleEntry");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte zoneGroup = readSimpleField("zoneGroup", readByte(readBuffer, 8));

    HVACZoneList zoneList =
        readSimpleField(
            "zoneList",
            new DataReaderComplexDefault<>(() -> HVACZoneList.staticParse(readBuffer), readBuffer));

    short entry = readSimpleField("entry", readUnsignedShort(readBuffer, 8));

    byte format = readSimpleField("format", readByte(readBuffer, 8));

    HVACModeAndFlags hvacModeAndFlags =
        readSimpleField(
            "hvacModeAndFlags",
            new DataReaderComplexDefault<>(
                () -> HVACModeAndFlags.staticParse(readBuffer), readBuffer));

    HVACStartTime startTime =
        readSimpleField(
            "startTime",
            new DataReaderComplexDefault<>(
                () -> HVACStartTime.staticParse(readBuffer), readBuffer));

    HVACTemperature level =
        readOptionalField(
            "level",
            new DataReaderComplexDefault<>(
                () -> HVACTemperature.staticParse(readBuffer), readBuffer),
            hvacModeAndFlags.getIsLevelTemperature());

    HVACRawLevels rawLevel =
        readOptionalField(
            "rawLevel",
            new DataReaderComplexDefault<>(() -> HVACRawLevels.staticParse(readBuffer), readBuffer),
            hvacModeAndFlags.getIsLevelRaw());

    readBuffer.closeContext("AirConditioningDataHvacScheduleEntry");
    // Create the instance
    return new AirConditioningDataHvacScheduleEntryBuilderImpl(
        zoneGroup, zoneList, entry, format, hvacModeAndFlags, startTime, level, rawLevel);
  }

  public static class AirConditioningDataHvacScheduleEntryBuilderImpl
      implements AirConditioningData.AirConditioningDataBuilder {
    private final byte zoneGroup;
    private final HVACZoneList zoneList;
    private final short entry;
    private final byte format;
    private final HVACModeAndFlags hvacModeAndFlags;
    private final HVACStartTime startTime;
    private final HVACTemperature level;
    private final HVACRawLevels rawLevel;

    public AirConditioningDataHvacScheduleEntryBuilderImpl(
        byte zoneGroup,
        HVACZoneList zoneList,
        short entry,
        byte format,
        HVACModeAndFlags hvacModeAndFlags,
        HVACStartTime startTime,
        HVACTemperature level,
        HVACRawLevels rawLevel) {
      this.zoneGroup = zoneGroup;
      this.zoneList = zoneList;
      this.entry = entry;
      this.format = format;
      this.hvacModeAndFlags = hvacModeAndFlags;
      this.startTime = startTime;
      this.level = level;
      this.rawLevel = rawLevel;
    }

    public AirConditioningDataHvacScheduleEntry build(
        AirConditioningCommandTypeContainer commandTypeContainer) {
      AirConditioningDataHvacScheduleEntry airConditioningDataHvacScheduleEntry =
          new AirConditioningDataHvacScheduleEntry(
              commandTypeContainer,
              zoneGroup,
              zoneList,
              entry,
              format,
              hvacModeAndFlags,
              startTime,
              level,
              rawLevel);
      return airConditioningDataHvacScheduleEntry;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AirConditioningDataHvacScheduleEntry)) {
      return false;
    }
    AirConditioningDataHvacScheduleEntry that = (AirConditioningDataHvacScheduleEntry) o;
    return (getZoneGroup() == that.getZoneGroup())
        && (getZoneList() == that.getZoneList())
        && (getEntry() == that.getEntry())
        && (getFormat() == that.getFormat())
        && (getHvacModeAndFlags() == that.getHvacModeAndFlags())
        && (getStartTime() == that.getStartTime())
        && (getLevel() == that.getLevel())
        && (getRawLevel() == that.getRawLevel())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getZoneGroup(),
        getZoneList(),
        getEntry(),
        getFormat(),
        getHvacModeAndFlags(),
        getStartTime(),
        getLevel(),
        getRawLevel());
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
