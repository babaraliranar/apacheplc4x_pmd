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

public class ServerStatusDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "864";
  }

  // Properties.
  protected final long startTime;
  protected final long currentTime;
  protected final ServerState state;
  protected final ExtensionObjectDefinition buildInfo;
  protected final long secondsTillShutdown;
  protected final LocalizedText shutdownReason;

  public ServerStatusDataType(
      long startTime,
      long currentTime,
      ServerState state,
      ExtensionObjectDefinition buildInfo,
      long secondsTillShutdown,
      LocalizedText shutdownReason) {
    super();
    this.startTime = startTime;
    this.currentTime = currentTime;
    this.state = state;
    this.buildInfo = buildInfo;
    this.secondsTillShutdown = secondsTillShutdown;
    this.shutdownReason = shutdownReason;
  }

  public long getStartTime() {
    return startTime;
  }

  public long getCurrentTime() {
    return currentTime;
  }

  public ServerState getState() {
    return state;
  }

  public ExtensionObjectDefinition getBuildInfo() {
    return buildInfo;
  }

  public long getSecondsTillShutdown() {
    return secondsTillShutdown;
  }

  public LocalizedText getShutdownReason() {
    return shutdownReason;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ServerStatusDataType");

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (currentTime)
    writeSimpleField("currentTime", currentTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (state)
    writeSimpleEnumField(
        "state",
        "ServerState",
        state,
        new DataWriterEnumDefault<>(
            ServerState::getValue, ServerState::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (buildInfo)
    writeSimpleField("buildInfo", buildInfo, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (secondsTillShutdown)
    writeSimpleField(
        "secondsTillShutdown", secondsTillShutdown, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (shutdownReason)
    writeSimpleField("shutdownReason", shutdownReason, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ServerStatusDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ServerStatusDataType _value = this;

    // Simple field (startTime)
    lengthInBits += 64;

    // Simple field (currentTime)
    lengthInBits += 64;

    // Simple field (state)
    lengthInBits += 32;

    // Simple field (buildInfo)
    lengthInBits += buildInfo.getLengthInBits();

    // Simple field (secondsTillShutdown)
    lengthInBits += 32;

    // Simple field (shutdownReason)
    lengthInBits += shutdownReason.getLengthInBits();

    return lengthInBits;
  }

  public static ServerStatusDataTypeBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ServerStatusDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    long startTime = readSimpleField("startTime", readSignedLong(readBuffer, 64));

    long currentTime = readSimpleField("currentTime", readSignedLong(readBuffer, 64));

    ServerState state =
        readEnumField(
            "state",
            "ServerState",
            new DataReaderEnumDefault<>(
                ServerState::enumForValue, readUnsignedLong(readBuffer, 32)));

    ExtensionObjectDefinition buildInfo =
        readSimpleField(
            "buildInfo",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("340")),
                readBuffer));

    long secondsTillShutdown =
        readSimpleField("secondsTillShutdown", readUnsignedLong(readBuffer, 32));

    LocalizedText shutdownReason =
        readSimpleField(
            "shutdownReason",
            new DataReaderComplexDefault<>(
                () -> LocalizedText.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ServerStatusDataType");
    // Create the instance
    return new ServerStatusDataTypeBuilder(
        startTime, currentTime, state, buildInfo, secondsTillShutdown, shutdownReason);
  }

  public static class ServerStatusDataTypeBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long startTime;
    private final long currentTime;
    private final ServerState state;
    private final ExtensionObjectDefinition buildInfo;
    private final long secondsTillShutdown;
    private final LocalizedText shutdownReason;

    public ServerStatusDataTypeBuilder(
        long startTime,
        long currentTime,
        ServerState state,
        ExtensionObjectDefinition buildInfo,
        long secondsTillShutdown,
        LocalizedText shutdownReason) {

      this.startTime = startTime;
      this.currentTime = currentTime;
      this.state = state;
      this.buildInfo = buildInfo;
      this.secondsTillShutdown = secondsTillShutdown;
      this.shutdownReason = shutdownReason;
    }

    public ServerStatusDataType build() {
      ServerStatusDataType serverStatusDataType =
          new ServerStatusDataType(
              startTime, currentTime, state, buildInfo, secondsTillShutdown, shutdownReason);
      return serverStatusDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ServerStatusDataType)) {
      return false;
    }
    ServerStatusDataType that = (ServerStatusDataType) o;
    return (getStartTime() == that.getStartTime())
        && (getCurrentTime() == that.getCurrentTime())
        && (getState() == that.getState())
        && (getBuildInfo() == that.getBuildInfo())
        && (getSecondsTillShutdown() == that.getSecondsTillShutdown())
        && (getShutdownReason() == that.getShutdownReason())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getStartTime(),
        getCurrentTime(),
        getState(),
        getBuildInfo(),
        getSecondsTillShutdown(),
        getShutdownReason());
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
