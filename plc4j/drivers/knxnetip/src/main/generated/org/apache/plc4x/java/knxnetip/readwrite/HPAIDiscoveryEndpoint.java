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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class HPAIDiscoveryEndpoint implements Message {

  // Properties.
  protected final HostProtocolCode hostProtocolCode;
  protected final IPAddress ipAddress;
  protected final int ipPort;

  public HPAIDiscoveryEndpoint(HostProtocolCode hostProtocolCode, IPAddress ipAddress, int ipPort) {
    super();
    this.hostProtocolCode = hostProtocolCode;
    this.ipAddress = ipAddress;
    this.ipPort = ipPort;
  }

  public HostProtocolCode getHostProtocolCode() {
    return hostProtocolCode;
  }

  public IPAddress getIpAddress() {
    return ipAddress;
  }

  public int getIpPort() {
    return ipPort;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("HPAIDiscoveryEndpoint");

    // Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short structureLength = (short) (getLengthInBytes());
    writeImplicitField("structureLength", structureLength, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (hostProtocolCode)
    writeSimpleEnumField(
        "hostProtocolCode",
        "HostProtocolCode",
        hostProtocolCode,
        new DataWriterEnumDefault<>(
            HostProtocolCode::getValue,
            HostProtocolCode::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (ipAddress)
    writeSimpleField("ipAddress", ipAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (ipPort)
    writeSimpleField("ipPort", ipPort, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("HPAIDiscoveryEndpoint");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    HPAIDiscoveryEndpoint _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (structureLength)
    lengthInBits += 8;

    // Simple field (hostProtocolCode)
    lengthInBits += 8;

    // Simple field (ipAddress)
    lengthInBits += ipAddress.getLengthInBits();

    // Simple field (ipPort)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static HPAIDiscoveryEndpoint staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static HPAIDiscoveryEndpoint staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("HPAIDiscoveryEndpoint");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short structureLength = readImplicitField("structureLength", readUnsignedShort(readBuffer, 8));

    HostProtocolCode hostProtocolCode =
        readEnumField(
            "hostProtocolCode",
            "HostProtocolCode",
            new DataReaderEnumDefault<>(
                HostProtocolCode::enumForValue, readUnsignedShort(readBuffer, 8)));

    IPAddress ipAddress =
        readSimpleField(
            "ipAddress",
            new DataReaderComplexDefault<>(() -> IPAddress.staticParse(readBuffer), readBuffer));

    int ipPort = readSimpleField("ipPort", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("HPAIDiscoveryEndpoint");
    // Create the instance
    HPAIDiscoveryEndpoint _hPAIDiscoveryEndpoint;
    _hPAIDiscoveryEndpoint = new HPAIDiscoveryEndpoint(hostProtocolCode, ipAddress, ipPort);
    return _hPAIDiscoveryEndpoint;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof HPAIDiscoveryEndpoint)) {
      return false;
    }
    HPAIDiscoveryEndpoint that = (HPAIDiscoveryEndpoint) o;
    return (getHostProtocolCode() == that.getHostProtocolCode())
        && (getIpAddress() == that.getIpAddress())
        && (getIpPort() == that.getIpPort())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHostProtocolCode(), getIpAddress(), getIpPort());
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
