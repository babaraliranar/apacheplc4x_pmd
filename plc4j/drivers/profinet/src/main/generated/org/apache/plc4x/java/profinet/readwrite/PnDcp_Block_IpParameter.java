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

public class PnDcp_Block_IpParameter extends PnDcp_Block implements Message {

  // Accessors for discriminator values.
  public PnDcp_BlockOptions getOption() {
    return PnDcp_BlockOptions.IP_OPTION;
  }

  public Short getSuboption() {
    return (short) 2;
  }

  // Properties.
  protected final boolean ipConflictDetected;
  protected final boolean setViaDhcp;
  protected final boolean setManually;
  protected final byte[] ipAddress;
  protected final byte[] subnetMask;
  protected final byte[] standardGateway;

  // Reserved Fields
  private Short reservedField0;
  private Short reservedField1;

  public PnDcp_Block_IpParameter(
      boolean ipConflictDetected,
      boolean setViaDhcp,
      boolean setManually,
      byte[] ipAddress,
      byte[] subnetMask,
      byte[] standardGateway) {
    super();
    this.ipConflictDetected = ipConflictDetected;
    this.setViaDhcp = setViaDhcp;
    this.setManually = setManually;
    this.ipAddress = ipAddress;
    this.subnetMask = subnetMask;
    this.standardGateway = standardGateway;
  }

  public boolean getIpConflictDetected() {
    return ipConflictDetected;
  }

  public boolean getSetViaDhcp() {
    return setViaDhcp;
  }

  public boolean getSetManually() {
    return setManually;
  }

  public byte[] getIpAddress() {
    return ipAddress;
  }

  public byte[] getSubnetMask() {
    return subnetMask;
  }

  public byte[] getStandardGateway() {
    return standardGateway;
  }

  @Override
  protected void serializePnDcp_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnDcp_Block_IpParameter");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (ipConflictDetected)
    writeSimpleField(
        "ipConflictDetected",
        ipConflictDetected,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 5),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (setViaDhcp)
    writeSimpleField(
        "setViaDhcp",
        setViaDhcp,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (setManually)
    writeSimpleField(
        "setManually",
        setManually,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (ipAddress)
    writeByteArrayField(
        "ipAddress",
        ipAddress,
        writeByteArray(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (subnetMask)
    writeByteArrayField(
        "subnetMask",
        subnetMask,
        writeByteArray(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (standardGateway)
    writeByteArrayField(
        "standardGateway",
        standardGateway,
        writeByteArray(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnDcp_Block_IpParameter");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnDcp_Block_IpParameter _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (ipConflictDetected)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 5;

    // Simple field (setViaDhcp)
    lengthInBits += 1;

    // Simple field (setManually)
    lengthInBits += 1;

    // Array field
    if (ipAddress != null) {
      lengthInBits += 8 * ipAddress.length;
    }

    // Array field
    if (subnetMask != null) {
      lengthInBits += 8 * subnetMask.length;
    }

    // Array field
    if (standardGateway != null) {
      lengthInBits += 8 * standardGateway.length;
    }

    return lengthInBits;
  }

  public static PnDcp_BlockBuilder staticParsePnDcp_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnDcp_Block_IpParameter");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 8),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean ipConflictDetected =
        readSimpleField(
            "ipConflictDetected",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Short reservedField1 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 5),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean setViaDhcp =
        readSimpleField(
            "setViaDhcp", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean setManually =
        readSimpleField(
            "setManually", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    byte[] ipAddress =
        readBuffer.readByteArray(
            "ipAddress", Math.toIntExact(4), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    byte[] subnetMask =
        readBuffer.readByteArray(
            "subnetMask", Math.toIntExact(4), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    byte[] standardGateway =
        readBuffer.readByteArray(
            "standardGateway", Math.toIntExact(4), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnDcp_Block_IpParameter");
    // Create the instance
    return new PnDcp_Block_IpParameterBuilderImpl(
        ipConflictDetected,
        setViaDhcp,
        setManually,
        ipAddress,
        subnetMask,
        standardGateway,
        reservedField0,
        reservedField1);
  }

  public static class PnDcp_Block_IpParameterBuilderImpl implements PnDcp_Block.PnDcp_BlockBuilder {
    private final boolean ipConflictDetected;
    private final boolean setViaDhcp;
    private final boolean setManually;
    private final byte[] ipAddress;
    private final byte[] subnetMask;
    private final byte[] standardGateway;
    private final Short reservedField0;
    private final Short reservedField1;

    public PnDcp_Block_IpParameterBuilderImpl(
        boolean ipConflictDetected,
        boolean setViaDhcp,
        boolean setManually,
        byte[] ipAddress,
        byte[] subnetMask,
        byte[] standardGateway,
        Short reservedField0,
        Short reservedField1) {
      this.ipConflictDetected = ipConflictDetected;
      this.setViaDhcp = setViaDhcp;
      this.setManually = setManually;
      this.ipAddress = ipAddress;
      this.subnetMask = subnetMask;
      this.standardGateway = standardGateway;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
    }

    public PnDcp_Block_IpParameter build() {
      PnDcp_Block_IpParameter pnDcp_Block_IpParameter =
          new PnDcp_Block_IpParameter(
              ipConflictDetected, setViaDhcp, setManually, ipAddress, subnetMask, standardGateway);
      pnDcp_Block_IpParameter.reservedField0 = reservedField0;
      pnDcp_Block_IpParameter.reservedField1 = reservedField1;
      return pnDcp_Block_IpParameter;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Block_IpParameter)) {
      return false;
    }
    PnDcp_Block_IpParameter that = (PnDcp_Block_IpParameter) o;
    return (getIpConflictDetected() == that.getIpConflictDetected())
        && (getSetViaDhcp() == that.getSetViaDhcp())
        && (getSetManually() == that.getSetManually())
        && (getIpAddress() == that.getIpAddress())
        && (getSubnetMask() == that.getSubnetMask())
        && (getStandardGateway() == that.getStandardGateway())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getIpConflictDetected(),
        getSetViaDhcp(),
        getSetManually(),
        getIpAddress(),
        getSubnetMask(),
        getStandardGateway());
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
