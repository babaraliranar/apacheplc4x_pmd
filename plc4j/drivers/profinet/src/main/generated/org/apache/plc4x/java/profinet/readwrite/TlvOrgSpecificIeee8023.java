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

public class TlvOrgSpecificIeee8023 extends TlvOrganizationSpecificUnit implements Message {

  // Accessors for discriminator values.
  public Long getUniqueCode() {
    return (long) 0x00120F;
  }

  // Properties.
  protected final short subType;
  protected final short negotiationSupport;
  protected final int negotiationCapability;
  protected final int operationalMauType;

  public TlvOrgSpecificIeee8023(
      short subType, short negotiationSupport, int negotiationCapability, int operationalMauType) {
    super();
    this.subType = subType;
    this.negotiationSupport = negotiationSupport;
    this.negotiationCapability = negotiationCapability;
    this.operationalMauType = operationalMauType;
  }

  public short getSubType() {
    return subType;
  }

  public short getNegotiationSupport() {
    return negotiationSupport;
  }

  public int getNegotiationCapability() {
    return negotiationCapability;
  }

  public int getOperationalMauType() {
    return operationalMauType;
  }

  @Override
  protected void serializeTlvOrganizationSpecificUnitChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TlvOrgSpecificIeee8023");

    // Simple Field (subType)
    writeSimpleField("subType", subType, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (negotiationSupport)
    writeSimpleField("negotiationSupport", negotiationSupport, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (negotiationCapability)
    writeSimpleField(
        "negotiationCapability", negotiationCapability, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (operationalMauType)
    writeSimpleField("operationalMauType", operationalMauType, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("TlvOrgSpecificIeee8023");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvOrgSpecificIeee8023 _value = this;

    // Simple field (subType)
    lengthInBits += 8;

    // Simple field (negotiationSupport)
    lengthInBits += 8;

    // Simple field (negotiationCapability)
    lengthInBits += 16;

    // Simple field (operationalMauType)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static TlvOrgSpecificIeee8023Builder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("TlvOrgSpecificIeee8023");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short subType = readSimpleField("subType", readUnsignedShort(readBuffer, 8));

    short negotiationSupport =
        readSimpleField("negotiationSupport", readUnsignedShort(readBuffer, 8));

    int negotiationCapability =
        readSimpleField("negotiationCapability", readUnsignedInt(readBuffer, 16));

    int operationalMauType = readSimpleField("operationalMauType", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("TlvOrgSpecificIeee8023");
    // Create the instance
    return new TlvOrgSpecificIeee8023Builder(
        subType, negotiationSupport, negotiationCapability, operationalMauType);
  }

  public static class TlvOrgSpecificIeee8023Builder
      implements TlvOrganizationSpecificUnit.TlvOrganizationSpecificUnitBuilder {
    private final short subType;
    private final short negotiationSupport;
    private final int negotiationCapability;
    private final int operationalMauType;

    public TlvOrgSpecificIeee8023Builder(
        short subType,
        short negotiationSupport,
        int negotiationCapability,
        int operationalMauType) {

      this.subType = subType;
      this.negotiationSupport = negotiationSupport;
      this.negotiationCapability = negotiationCapability;
      this.operationalMauType = operationalMauType;
    }

    public TlvOrgSpecificIeee8023 build() {
      TlvOrgSpecificIeee8023 tlvOrgSpecificIeee8023 =
          new TlvOrgSpecificIeee8023(
              subType, negotiationSupport, negotiationCapability, operationalMauType);
      return tlvOrgSpecificIeee8023;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvOrgSpecificIeee8023)) {
      return false;
    }
    TlvOrgSpecificIeee8023 that = (TlvOrgSpecificIeee8023) o;
    return (getSubType() == that.getSubType())
        && (getNegotiationSupport() == that.getNegotiationSupport())
        && (getNegotiationCapability() == that.getNegotiationCapability())
        && (getOperationalMauType() == that.getOperationalMauType())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSubType(),
        getNegotiationSupport(),
        getNegotiationCapability(),
        getOperationalMauType());
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
