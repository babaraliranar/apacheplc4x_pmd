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

public class SecurityDataStatusReport2 extends SecurityData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final List<ZoneStatus> zoneStatus;

  public SecurityDataStatusReport2(
      SecurityCommandTypeContainer commandTypeContainer,
      byte argument,
      List<ZoneStatus> zoneStatus) {
    super(commandTypeContainer, argument);
    this.zoneStatus = zoneStatus;
  }

  public List<ZoneStatus> getZoneStatus() {
    return zoneStatus;
  }

  @Override
  protected void serializeSecurityDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SecurityDataStatusReport2");

    // Array Field (zoneStatus)
    writeComplexTypeArrayField("zoneStatus", zoneStatus, writeBuffer);

    writeBuffer.popContext("SecurityDataStatusReport2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SecurityDataStatusReport2 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (zoneStatus != null) {
      int i = 0;
      for (ZoneStatus element : zoneStatus) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= zoneStatus.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static SecurityDataBuilder staticParseSecurityDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("SecurityDataStatusReport2");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    List<ZoneStatus> zoneStatus =
        readCountArrayField(
            "zoneStatus",
            new DataReaderComplexDefault<>(() -> ZoneStatus.staticParse(readBuffer), readBuffer),
            48);

    readBuffer.closeContext("SecurityDataStatusReport2");
    // Create the instance
    return new SecurityDataStatusReport2BuilderImpl(zoneStatus);
  }

  public static class SecurityDataStatusReport2BuilderImpl
      implements SecurityData.SecurityDataBuilder {
    private final List<ZoneStatus> zoneStatus;

    public SecurityDataStatusReport2BuilderImpl(List<ZoneStatus> zoneStatus) {
      this.zoneStatus = zoneStatus;
    }

    public SecurityDataStatusReport2 build(
        SecurityCommandTypeContainer commandTypeContainer, byte argument) {
      SecurityDataStatusReport2 securityDataStatusReport2 =
          new SecurityDataStatusReport2(commandTypeContainer, argument, zoneStatus);
      return securityDataStatusReport2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityDataStatusReport2)) {
      return false;
    }
    SecurityDataStatusReport2 that = (SecurityDataStatusReport2) o;
    return (getZoneStatus() == that.getZoneStatus()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getZoneStatus());
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
