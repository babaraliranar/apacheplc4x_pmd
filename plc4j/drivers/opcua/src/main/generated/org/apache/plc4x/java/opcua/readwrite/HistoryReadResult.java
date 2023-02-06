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

public class HistoryReadResult extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "640";
  }

  // Properties.
  protected final StatusCode statusCode;
  protected final PascalByteString continuationPoint;
  protected final ExtensionObject historyData;

  public HistoryReadResult(
      StatusCode statusCode, PascalByteString continuationPoint, ExtensionObject historyData) {
    super();
    this.statusCode = statusCode;
    this.continuationPoint = continuationPoint;
    this.historyData = historyData;
  }

  public StatusCode getStatusCode() {
    return statusCode;
  }

  public PascalByteString getContinuationPoint() {
    return continuationPoint;
  }

  public ExtensionObject getHistoryData() {
    return historyData;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("HistoryReadResult");

    // Simple Field (statusCode)
    writeSimpleField("statusCode", statusCode, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (continuationPoint)
    writeSimpleField(
        "continuationPoint", continuationPoint, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (historyData)
    writeSimpleField("historyData", historyData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("HistoryReadResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    HistoryReadResult _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (statusCode)
    lengthInBits += statusCode.getLengthInBits();

    // Simple field (continuationPoint)
    lengthInBits += continuationPoint.getLengthInBits();

    // Simple field (historyData)
    lengthInBits += historyData.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("HistoryReadResult");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    StatusCode statusCode =
        readSimpleField(
            "statusCode",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer));

    PascalByteString continuationPoint =
        readSimpleField(
            "continuationPoint",
            new DataReaderComplexDefault<>(
                () -> PascalByteString.staticParse(readBuffer), readBuffer));

    ExtensionObject historyData =
        readSimpleField(
            "historyData",
            new DataReaderComplexDefault<>(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    readBuffer.closeContext("HistoryReadResult");
    // Create the instance
    return new HistoryReadResultBuilderImpl(statusCode, continuationPoint, historyData);
  }

  public static class HistoryReadResultBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final StatusCode statusCode;
    private final PascalByteString continuationPoint;
    private final ExtensionObject historyData;

    public HistoryReadResultBuilderImpl(
        StatusCode statusCode, PascalByteString continuationPoint, ExtensionObject historyData) {
      this.statusCode = statusCode;
      this.continuationPoint = continuationPoint;
      this.historyData = historyData;
    }

    public HistoryReadResult build() {
      HistoryReadResult historyReadResult =
          new HistoryReadResult(statusCode, continuationPoint, historyData);
      return historyReadResult;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof HistoryReadResult)) {
      return false;
    }
    HistoryReadResult that = (HistoryReadResult) o;
    return (getStatusCode() == that.getStatusCode())
        && (getContinuationPoint() == that.getContinuationPoint())
        && (getHistoryData() == that.getHistoryData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getStatusCode(), getContinuationPoint(), getHistoryData());
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
