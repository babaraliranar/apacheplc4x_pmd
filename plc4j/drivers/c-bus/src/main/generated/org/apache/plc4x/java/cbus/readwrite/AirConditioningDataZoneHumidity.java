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

public class AirConditioningDataZoneHumidity extends AirConditioningData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte zoneGroup;
  protected final HVACZoneList zoneList;
  protected final HVACHumidity humidity;
  protected final HVACSensorStatus sensorStatus;

  public AirConditioningDataZoneHumidity(
      AirConditioningCommandTypeContainer commandTypeContainer,
      byte zoneGroup,
      HVACZoneList zoneList,
      HVACHumidity humidity,
      HVACSensorStatus sensorStatus) {
    super(commandTypeContainer);
    this.zoneGroup = zoneGroup;
    this.zoneList = zoneList;
    this.humidity = humidity;
    this.sensorStatus = sensorStatus;
  }

  public byte getZoneGroup() {
    return zoneGroup;
  }

  public HVACZoneList getZoneList() {
    return zoneList;
  }

  public HVACHumidity getHumidity() {
    return humidity;
  }

  public HVACSensorStatus getSensorStatus() {
    return sensorStatus;
  }

  @Override
  protected void serializeAirConditioningDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AirConditioningDataZoneHumidity");

    // Simple Field (zoneGroup)
    writeSimpleField("zoneGroup", zoneGroup, writeByte(writeBuffer, 8));

    // Simple Field (zoneList)
    writeSimpleField("zoneList", zoneList, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (humidity)
    writeSimpleField("humidity", humidity, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (sensorStatus)
    writeSimpleEnumField(
        "sensorStatus",
        "HVACSensorStatus",
        sensorStatus,
        new DataWriterEnumDefault<>(
            HVACSensorStatus::getValue,
            HVACSensorStatus::name,
            writeUnsignedShort(writeBuffer, 8)));

    writeBuffer.popContext("AirConditioningDataZoneHumidity");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AirConditioningDataZoneHumidity _value = this;

    // Simple field (zoneGroup)
    lengthInBits += 8;

    // Simple field (zoneList)
    lengthInBits += zoneList.getLengthInBits();

    // Simple field (humidity)
    lengthInBits += humidity.getLengthInBits();

    // Simple field (sensorStatus)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static AirConditioningDataZoneHumidityBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AirConditioningDataZoneHumidity");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte zoneGroup = readSimpleField("zoneGroup", readByte(readBuffer, 8));

    HVACZoneList zoneList =
        readSimpleField(
            "zoneList",
            new DataReaderComplexDefault<>(() -> HVACZoneList.staticParse(readBuffer), readBuffer));

    HVACHumidity humidity =
        readSimpleField(
            "humidity",
            new DataReaderComplexDefault<>(() -> HVACHumidity.staticParse(readBuffer), readBuffer));

    HVACSensorStatus sensorStatus =
        readEnumField(
            "sensorStatus",
            "HVACSensorStatus",
            new DataReaderEnumDefault<>(
                HVACSensorStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    readBuffer.closeContext("AirConditioningDataZoneHumidity");
    // Create the instance
    return new AirConditioningDataZoneHumidityBuilder(zoneGroup, zoneList, humidity, sensorStatus);
  }

  public static class AirConditioningDataZoneHumidityBuilder
      implements AirConditioningData.AirConditioningDataBuilder {
    private final byte zoneGroup;
    private final HVACZoneList zoneList;
    private final HVACHumidity humidity;
    private final HVACSensorStatus sensorStatus;

    public AirConditioningDataZoneHumidityBuilder(
        byte zoneGroup,
        HVACZoneList zoneList,
        HVACHumidity humidity,
        HVACSensorStatus sensorStatus) {

      this.zoneGroup = zoneGroup;
      this.zoneList = zoneList;
      this.humidity = humidity;
      this.sensorStatus = sensorStatus;
    }

    public AirConditioningDataZoneHumidity build(
        AirConditioningCommandTypeContainer commandTypeContainer) {
      AirConditioningDataZoneHumidity airConditioningDataZoneHumidity =
          new AirConditioningDataZoneHumidity(
              commandTypeContainer, zoneGroup, zoneList, humidity, sensorStatus);
      return airConditioningDataZoneHumidity;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AirConditioningDataZoneHumidity)) {
      return false;
    }
    AirConditioningDataZoneHumidity that = (AirConditioningDataZoneHumidity) o;
    return (getZoneGroup() == that.getZoneGroup())
        && (getZoneList() == that.getZoneList())
        && (getHumidity() == that.getHumidity())
        && (getSensorStatus() == that.getSensorStatus())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getZoneGroup(), getZoneList(), getHumidity(), getSensorStatus());
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
