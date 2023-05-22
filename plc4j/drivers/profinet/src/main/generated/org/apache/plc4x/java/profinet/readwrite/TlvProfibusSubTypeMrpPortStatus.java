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
package org.apache.plc4x.java.profinet.readwrite;

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

public class TlvProfibusSubTypeMrpPortStatus extends TlvOrgSpecificProfibusUnit implements Message {

  // Accessors for discriminator values.
  public TlvProfibusSubType getSubType() {
    return TlvProfibusSubType.MRP_PORT_STATUS;
  }

  // Properties.
  protected final Uuid macAddress;
  protected final int Status;

  public TlvProfibusSubTypeMrpPortStatus(Uuid macAddress, int Status) {
    super();
    this.macAddress = macAddress;
    this.Status = Status;
  }

  public Uuid getMacAddress() {
    return macAddress;
  }

  public int getStatus() {
    return Status;
  }

  @Override
  protected void serializeTlvOrgSpecificProfibusUnitChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvProfibusSubTypeMrpPortStatus");

    // Simple Field (macAddress)
    writeSimpleField("macAddress", macAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (Status)
    writeSimpleField("Status", Status, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("TlvProfibusSubTypeMrpPortStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvProfibusSubTypeMrpPortStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (macAddress)
    lengthInBits += macAddress.getLengthInBits();

    // Simple field (Status)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static TlvOrgSpecificProfibusUnitBuilder staticParseTlvOrgSpecificProfibusUnitBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("TlvProfibusSubTypeMrpPortStatus");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Uuid macAddress =
        readSimpleField(
            "macAddress",
            new DataReaderComplexDefault<>(() -> Uuid.staticParse(readBuffer), readBuffer));

    int Status = readSimpleField("Status", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("TlvProfibusSubTypeMrpPortStatus");
    // Create the instance
    return new TlvProfibusSubTypeMrpPortStatusBuilderImpl(macAddress, Status);
  }

  public static class TlvProfibusSubTypeMrpPortStatusBuilderImpl
      implements TlvOrgSpecificProfibusUnit.TlvOrgSpecificProfibusUnitBuilder {
    private final Uuid macAddress;
    private final int Status;

    public TlvProfibusSubTypeMrpPortStatusBuilderImpl(Uuid macAddress, int Status) {
      this.macAddress = macAddress;
      this.Status = Status;
    }

    public TlvProfibusSubTypeMrpPortStatus build() {
      TlvProfibusSubTypeMrpPortStatus tlvProfibusSubTypeMrpPortStatus =
          new TlvProfibusSubTypeMrpPortStatus(macAddress, Status);
      return tlvProfibusSubTypeMrpPortStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvProfibusSubTypeMrpPortStatus)) {
      return false;
    }
    TlvProfibusSubTypeMrpPortStatus that = (TlvProfibusSubTypeMrpPortStatus) o;
    return (getMacAddress() == that.getMacAddress())
        && (getStatus() == that.getStatus())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMacAddress(), getStatus());
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
