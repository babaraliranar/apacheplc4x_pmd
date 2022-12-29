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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ModifySubscriptionResponse extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "796";
  }

  // Properties.
  protected final ExtensionObjectDefinition responseHeader;
  protected final double revisedPublishingInterval;
  protected final long revisedLifetimeCount;
  protected final long revisedMaxKeepAliveCount;

  public ModifySubscriptionResponse(
      ExtensionObjectDefinition responseHeader,
      double revisedPublishingInterval,
      long revisedLifetimeCount,
      long revisedMaxKeepAliveCount) {
    super();
    this.responseHeader = responseHeader;
    this.revisedPublishingInterval = revisedPublishingInterval;
    this.revisedLifetimeCount = revisedLifetimeCount;
    this.revisedMaxKeepAliveCount = revisedMaxKeepAliveCount;
  }

  public ExtensionObjectDefinition getResponseHeader() {
    return responseHeader;
  }

  public double getRevisedPublishingInterval() {
    return revisedPublishingInterval;
  }

  public long getRevisedLifetimeCount() {
    return revisedLifetimeCount;
  }

  public long getRevisedMaxKeepAliveCount() {
    return revisedMaxKeepAliveCount;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModifySubscriptionResponse");

    // Simple Field (responseHeader)
    writeSimpleField("responseHeader", responseHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (revisedPublishingInterval)
    writeSimpleField(
        "revisedPublishingInterval", revisedPublishingInterval, writeDouble(writeBuffer, 64));

    // Simple Field (revisedLifetimeCount)
    writeSimpleField(
        "revisedLifetimeCount", revisedLifetimeCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (revisedMaxKeepAliveCount)
    writeSimpleField(
        "revisedMaxKeepAliveCount", revisedMaxKeepAliveCount, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("ModifySubscriptionResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModifySubscriptionResponse _value = this;

    // Simple field (responseHeader)
    lengthInBits += responseHeader.getLengthInBits();

    // Simple field (revisedPublishingInterval)
    lengthInBits += 64;

    // Simple field (revisedLifetimeCount)
    lengthInBits += 32;

    // Simple field (revisedMaxKeepAliveCount)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ModifySubscriptionResponseBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ModifySubscriptionResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ExtensionObjectDefinition responseHeader =
        readSimpleField(
            "responseHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("394")),
                readBuffer));

    double revisedPublishingInterval =
        readSimpleField("revisedPublishingInterval", readDouble(readBuffer, 64));

    long revisedLifetimeCount =
        readSimpleField("revisedLifetimeCount", readUnsignedLong(readBuffer, 32));

    long revisedMaxKeepAliveCount =
        readSimpleField("revisedMaxKeepAliveCount", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("ModifySubscriptionResponse");
    // Create the instance
    return new ModifySubscriptionResponseBuilder(
        responseHeader, revisedPublishingInterval, revisedLifetimeCount, revisedMaxKeepAliveCount);
  }

  public static class ModifySubscriptionResponseBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition responseHeader;
    private final double revisedPublishingInterval;
    private final long revisedLifetimeCount;
    private final long revisedMaxKeepAliveCount;

    public ModifySubscriptionResponseBuilder(
        ExtensionObjectDefinition responseHeader,
        double revisedPublishingInterval,
        long revisedLifetimeCount,
        long revisedMaxKeepAliveCount) {

      this.responseHeader = responseHeader;
      this.revisedPublishingInterval = revisedPublishingInterval;
      this.revisedLifetimeCount = revisedLifetimeCount;
      this.revisedMaxKeepAliveCount = revisedMaxKeepAliveCount;
    }

    public ModifySubscriptionResponse build() {
      ModifySubscriptionResponse modifySubscriptionResponse =
          new ModifySubscriptionResponse(
              responseHeader,
              revisedPublishingInterval,
              revisedLifetimeCount,
              revisedMaxKeepAliveCount);
      return modifySubscriptionResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModifySubscriptionResponse)) {
      return false;
    }
    ModifySubscriptionResponse that = (ModifySubscriptionResponse) o;
    return (getResponseHeader() == that.getResponseHeader())
        && (getRevisedPublishingInterval() == that.getRevisedPublishingInterval())
        && (getRevisedLifetimeCount() == that.getRevisedLifetimeCount())
        && (getRevisedMaxKeepAliveCount() == that.getRevisedMaxKeepAliveCount())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getResponseHeader(),
        getRevisedPublishingInterval(),
        getRevisedLifetimeCount(),
        getRevisedMaxKeepAliveCount());
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
