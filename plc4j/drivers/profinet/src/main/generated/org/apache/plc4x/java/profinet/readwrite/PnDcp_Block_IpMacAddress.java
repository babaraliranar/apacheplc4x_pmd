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

public class PnDcp_Block_IpMacAddress extends PnDcp_Block implements Message {

  // Accessors for discriminator values.
  public PnDcp_BlockOptions getOption() {
    return PnDcp_BlockOptions.IP_OPTION;
  }

  public Short getSuboption() {
    return (short) 1;
  }

  // Properties.
  protected final MacAddress macAddress;

  public PnDcp_Block_IpMacAddress(MacAddress macAddress) {
    super();
    this.macAddress = macAddress;
  }

  public MacAddress getMacAddress() {
    return macAddress;
  }

  @Override
  protected void serializePnDcp_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnDcp_Block_IpMacAddress");

    // Reserved Field (reserved)
    writeReservedField("reserved", (int) 0x0000, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (macAddress)
    writeSimpleField("macAddress", macAddress, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("PnDcp_Block_IpMacAddress");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnDcp_Block_IpMacAddress _value = this;

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Simple field (macAddress)
    lengthInBits += macAddress.getLengthInBits();

    return lengthInBits;
  }

  public static PnDcp_BlockBuilder staticParsePnDcp_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnDcp_Block_IpMacAddress");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Integer reservedField0 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 16), (int) 0x0000);

    MacAddress macAddress =
        readSimpleField(
            "macAddress",
            new DataReaderComplexDefault<>(() -> MacAddress.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("PnDcp_Block_IpMacAddress");
    // Create the instance
    return new PnDcp_Block_IpMacAddressBuilderImpl(macAddress);
  }

  public static class PnDcp_Block_IpMacAddressBuilderImpl
      implements PnDcp_Block.PnDcp_BlockBuilder {
    private final MacAddress macAddress;

    public PnDcp_Block_IpMacAddressBuilderImpl(MacAddress macAddress) {
      this.macAddress = macAddress;
    }

    public PnDcp_Block_IpMacAddress build() {
      PnDcp_Block_IpMacAddress pnDcp_Block_IpMacAddress = new PnDcp_Block_IpMacAddress(macAddress);
      return pnDcp_Block_IpMacAddress;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Block_IpMacAddress)) {
      return false;
    }
    PnDcp_Block_IpMacAddress that = (PnDcp_Block_IpMacAddress) o;
    return (getMacAddress() == that.getMacAddress()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMacAddress());
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
