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

public abstract class ErrorReportingData implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final ErrorReportingCommandTypeContainer commandTypeContainer;

  public ErrorReportingData(ErrorReportingCommandTypeContainer commandTypeContainer) {
    super();
    this.commandTypeContainer = commandTypeContainer;
  }

  public ErrorReportingCommandTypeContainer getCommandTypeContainer() {
    return commandTypeContainer;
  }

  public ErrorReportingCommandType getCommandType() {
    return (ErrorReportingCommandType) (getCommandTypeContainer().getCommandType());
  }

  protected abstract void serializeErrorReportingDataChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ErrorReportingData");

    // Simple Field (commandTypeContainer)
    writeSimpleEnumField(
        "commandTypeContainer",
        "ErrorReportingCommandTypeContainer",
        commandTypeContainer,
        new DataWriterEnumDefault<>(
            ErrorReportingCommandTypeContainer::getValue,
            ErrorReportingCommandTypeContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    ErrorReportingCommandType commandType = getCommandType();
    writeBuffer.writeVirtual("commandType", commandType);

    // Switch field (Serialize the sub-type)
    serializeErrorReportingDataChild(writeBuffer);

    writeBuffer.popContext("ErrorReportingData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ErrorReportingData _value = this;

    // Simple field (commandTypeContainer)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static ErrorReportingData staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static ErrorReportingData staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ErrorReportingData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    // Validation
    if (!(org.apache.plc4x.java.cbus.readwrite.utils.StaticHelper
        .knowsErrorReportingCommandTypeContainer(readBuffer))) {
      throw new ParseAssertException("no command type could be found");
    }

    ErrorReportingCommandTypeContainer commandTypeContainer =
        readEnumField(
            "commandTypeContainer",
            "ErrorReportingCommandTypeContainer",
            new DataReaderEnumDefault<>(
                ErrorReportingCommandTypeContainer::enumForValue,
                readUnsignedShort(readBuffer, 8)));
    ErrorReportingCommandType commandType =
        readVirtualField(
            "commandType", ErrorReportingCommandType.class, commandTypeContainer.getCommandType());

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    ErrorReportingDataBuilder builder = null;
    if (true) {
      builder = ErrorReportingDataGeneric.staticParseErrorReportingDataBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandType="
              + commandType
              + "]");
    }

    readBuffer.closeContext("ErrorReportingData");
    // Create the instance
    ErrorReportingData _errorReportingData = builder.build(commandTypeContainer);
    return _errorReportingData;
  }

  public interface ErrorReportingDataBuilder {
    ErrorReportingData build(ErrorReportingCommandTypeContainer commandTypeContainer);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ErrorReportingData)) {
      return false;
    }
    ErrorReportingData that = (ErrorReportingData) o;
    return (getCommandTypeContainer() == that.getCommandTypeContainer()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCommandTypeContainer());
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
