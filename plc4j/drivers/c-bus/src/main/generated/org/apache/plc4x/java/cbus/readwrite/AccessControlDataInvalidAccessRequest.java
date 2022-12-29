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

public class AccessControlDataInvalidAccessRequest extends AccessControlData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final AccessControlDirection accessControlDirection;
  protected final byte[] data;

  public AccessControlDataInvalidAccessRequest(
      AccessControlCommandTypeContainer commandTypeContainer,
      byte networkId,
      byte accessPointId,
      AccessControlDirection accessControlDirection,
      byte[] data) {
    super(commandTypeContainer, networkId, accessPointId);
    this.accessControlDirection = accessControlDirection;
    this.data = data;
  }

  public AccessControlDirection getAccessControlDirection() {
    return accessControlDirection;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeAccessControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AccessControlDataInvalidAccessRequest");

    // Simple Field (accessControlDirection)
    writeSimpleEnumField(
        "accessControlDirection",
        "AccessControlDirection",
        accessControlDirection,
        new DataWriterEnumDefault<>(
            AccessControlDirection::getValue,
            AccessControlDirection::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("AccessControlDataInvalidAccessRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AccessControlDataInvalidAccessRequest _value = this;

    // Simple field (accessControlDirection)
    lengthInBits += 8;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static AccessControlDataInvalidAccessRequestBuilder staticParseBuilder(
      ReadBuffer readBuffer, AccessControlCommandTypeContainer commandTypeContainer)
      throws ParseException {
    readBuffer.pullContext("AccessControlDataInvalidAccessRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    AccessControlDirection accessControlDirection =
        readEnumField(
            "accessControlDirection",
            "AccessControlDirection",
            new DataReaderEnumDefault<>(
                AccessControlDirection::enumForValue, readUnsignedShort(readBuffer, 8)));

    byte[] data =
        readBuffer.readByteArray(
            "data", Math.toIntExact((commandTypeContainer.getNumBytes()) - (3)));

    readBuffer.closeContext("AccessControlDataInvalidAccessRequest");
    // Create the instance
    return new AccessControlDataInvalidAccessRequestBuilder(accessControlDirection, data);
  }

  public static class AccessControlDataInvalidAccessRequestBuilder
      implements AccessControlData.AccessControlDataBuilder {
    private final AccessControlDirection accessControlDirection;
    private final byte[] data;

    public AccessControlDataInvalidAccessRequestBuilder(
        AccessControlDirection accessControlDirection, byte[] data) {

      this.accessControlDirection = accessControlDirection;
      this.data = data;
    }

    public AccessControlDataInvalidAccessRequest build(
        AccessControlCommandTypeContainer commandTypeContainer,
        byte networkId,
        byte accessPointId) {
      AccessControlDataInvalidAccessRequest accessControlDataInvalidAccessRequest =
          new AccessControlDataInvalidAccessRequest(
              commandTypeContainer, networkId, accessPointId, accessControlDirection, data);
      return accessControlDataInvalidAccessRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AccessControlDataInvalidAccessRequest)) {
      return false;
    }
    AccessControlDataInvalidAccessRequest that = (AccessControlDataInvalidAccessRequest) o;
    return (getAccessControlDirection() == that.getAccessControlDirection())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAccessControlDirection(), getData());
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
