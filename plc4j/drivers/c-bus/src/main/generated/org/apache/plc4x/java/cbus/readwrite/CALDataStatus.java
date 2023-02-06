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

public class CALDataStatus extends CALData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final ApplicationIdContainer application;
  protected final short blockStart;
  protected final List<StatusByte> statusBytes;

  // Arguments.
  protected final RequestContext requestContext;

  public CALDataStatus(
      CALCommandTypeContainer commandTypeContainer,
      CALData additionalData,
      ApplicationIdContainer application,
      short blockStart,
      List<StatusByte> statusBytes,
      RequestContext requestContext) {
    super(commandTypeContainer, additionalData, requestContext);
    this.application = application;
    this.blockStart = blockStart;
    this.statusBytes = statusBytes;
    this.requestContext = requestContext;
  }

  public ApplicationIdContainer getApplication() {
    return application;
  }

  public short getBlockStart() {
    return blockStart;
  }

  public List<StatusByte> getStatusBytes() {
    return statusBytes;
  }

  @Override
  protected void serializeCALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CALDataStatus");

    // Simple Field (application)
    writeSimpleEnumField(
        "application",
        "ApplicationIdContainer",
        application,
        new DataWriterEnumDefault<>(
            ApplicationIdContainer::getValue,
            ApplicationIdContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (blockStart)
    writeSimpleField("blockStart", blockStart, writeUnsignedShort(writeBuffer, 8));

    // Array Field (statusBytes)
    writeComplexTypeArrayField("statusBytes", statusBytes, writeBuffer);

    writeBuffer.popContext("CALDataStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CALDataStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (application)
    lengthInBits += 8;

    // Simple field (blockStart)
    lengthInBits += 8;

    // Array field
    if (statusBytes != null) {
      int i = 0;
      for (StatusByte element : statusBytes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= statusBytes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static CALDataBuilder staticParseCALDataBuilder(
      ReadBuffer readBuffer,
      CALCommandTypeContainer commandTypeContainer,
      RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("CALDataStatus");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ApplicationIdContainer application =
        readEnumField(
            "application",
            "ApplicationIdContainer",
            new DataReaderEnumDefault<>(
                ApplicationIdContainer::enumForValue, readUnsignedShort(readBuffer, 8)));

    short blockStart = readSimpleField("blockStart", readUnsignedShort(readBuffer, 8));

    List<StatusByte> statusBytes =
        readCountArrayField(
            "statusBytes",
            new DataReaderComplexDefault<>(() -> StatusByte.staticParse(readBuffer), readBuffer),
            (commandTypeContainer.getNumBytes()) - (2));

    readBuffer.closeContext("CALDataStatus");
    // Create the instance
    return new CALDataStatusBuilderImpl(application, blockStart, statusBytes, requestContext);
  }

  public static class CALDataStatusBuilderImpl implements CALData.CALDataBuilder {
    private final ApplicationIdContainer application;
    private final short blockStart;
    private final List<StatusByte> statusBytes;
    private final RequestContext requestContext;

    public CALDataStatusBuilderImpl(
        ApplicationIdContainer application,
        short blockStart,
        List<StatusByte> statusBytes,
        RequestContext requestContext) {
      this.application = application;
      this.blockStart = blockStart;
      this.statusBytes = statusBytes;
      this.requestContext = requestContext;
    }

    public CALDataStatus build(
        CALCommandTypeContainer commandTypeContainer,
        CALData additionalData,
        RequestContext requestContext) {
      CALDataStatus cALDataStatus =
          new CALDataStatus(
              commandTypeContainer,
              additionalData,
              application,
              blockStart,
              statusBytes,
              requestContext);
      return cALDataStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CALDataStatus)) {
      return false;
    }
    CALDataStatus that = (CALDataStatus) o;
    return (getApplication() == that.getApplication())
        && (getBlockStart() == that.getBlockStart())
        && (getStatusBytes() == that.getStatusBytes())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getApplication(), getBlockStart(), getStatusBytes());
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
