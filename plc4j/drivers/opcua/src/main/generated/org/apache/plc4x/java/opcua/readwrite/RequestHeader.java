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

public class RequestHeader extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "391";
  }

  // Properties.
  protected final NodeId authenticationToken;
  protected final long timestamp;
  protected final long requestHandle;
  protected final long returnDiagnostics;
  protected final PascalString auditEntryId;
  protected final long timeoutHint;
  protected final ExtensionObject additionalHeader;

  public RequestHeader(
      NodeId authenticationToken,
      long timestamp,
      long requestHandle,
      long returnDiagnostics,
      PascalString auditEntryId,
      long timeoutHint,
      ExtensionObject additionalHeader) {
    super();
    this.authenticationToken = authenticationToken;
    this.timestamp = timestamp;
    this.requestHandle = requestHandle;
    this.returnDiagnostics = returnDiagnostics;
    this.auditEntryId = auditEntryId;
    this.timeoutHint = timeoutHint;
    this.additionalHeader = additionalHeader;
  }

  public NodeId getAuthenticationToken() {
    return authenticationToken;
  }

  public long getTimestamp() {
    return timestamp;
  }

  public long getRequestHandle() {
    return requestHandle;
  }

  public long getReturnDiagnostics() {
    return returnDiagnostics;
  }

  public PascalString getAuditEntryId() {
    return auditEntryId;
  }

  public long getTimeoutHint() {
    return timeoutHint;
  }

  public ExtensionObject getAdditionalHeader() {
    return additionalHeader;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RequestHeader");

    // Simple Field (authenticationToken)
    writeSimpleField(
        "authenticationToken", authenticationToken, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timestamp)
    writeSimpleField("timestamp", timestamp, writeSignedLong(writeBuffer, 64));

    // Simple Field (requestHandle)
    writeSimpleField("requestHandle", requestHandle, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (returnDiagnostics)
    writeSimpleField("returnDiagnostics", returnDiagnostics, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (auditEntryId)
    writeSimpleField("auditEntryId", auditEntryId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeoutHint)
    writeSimpleField("timeoutHint", timeoutHint, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (additionalHeader)
    writeSimpleField(
        "additionalHeader", additionalHeader, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("RequestHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RequestHeader _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (authenticationToken)
    lengthInBits += authenticationToken.getLengthInBits();

    // Simple field (timestamp)
    lengthInBits += 64;

    // Simple field (requestHandle)
    lengthInBits += 32;

    // Simple field (returnDiagnostics)
    lengthInBits += 32;

    // Simple field (auditEntryId)
    lengthInBits += auditEntryId.getLengthInBits();

    // Simple field (timeoutHint)
    lengthInBits += 32;

    // Simple field (additionalHeader)
    lengthInBits += additionalHeader.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RequestHeader");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId authenticationToken =
        readSimpleField(
            "authenticationToken",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    long timestamp = readSimpleField("timestamp", readSignedLong(readBuffer, 64));

    long requestHandle = readSimpleField("requestHandle", readUnsignedLong(readBuffer, 32));

    long returnDiagnostics = readSimpleField("returnDiagnostics", readUnsignedLong(readBuffer, 32));

    PascalString auditEntryId =
        readSimpleField(
            "auditEntryId",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    long timeoutHint = readSimpleField("timeoutHint", readUnsignedLong(readBuffer, 32));

    ExtensionObject additionalHeader =
        readSimpleField(
            "additionalHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    readBuffer.closeContext("RequestHeader");
    // Create the instance
    return new RequestHeaderBuilderImpl(
        authenticationToken,
        timestamp,
        requestHandle,
        returnDiagnostics,
        auditEntryId,
        timeoutHint,
        additionalHeader);
  }

  public static class RequestHeaderBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId authenticationToken;
    private final long timestamp;
    private final long requestHandle;
    private final long returnDiagnostics;
    private final PascalString auditEntryId;
    private final long timeoutHint;
    private final ExtensionObject additionalHeader;

    public RequestHeaderBuilderImpl(
        NodeId authenticationToken,
        long timestamp,
        long requestHandle,
        long returnDiagnostics,
        PascalString auditEntryId,
        long timeoutHint,
        ExtensionObject additionalHeader) {
      this.authenticationToken = authenticationToken;
      this.timestamp = timestamp;
      this.requestHandle = requestHandle;
      this.returnDiagnostics = returnDiagnostics;
      this.auditEntryId = auditEntryId;
      this.timeoutHint = timeoutHint;
      this.additionalHeader = additionalHeader;
    }

    public RequestHeader build() {
      RequestHeader requestHeader =
          new RequestHeader(
              authenticationToken,
              timestamp,
              requestHandle,
              returnDiagnostics,
              auditEntryId,
              timeoutHint,
              additionalHeader);
      return requestHeader;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RequestHeader)) {
      return false;
    }
    RequestHeader that = (RequestHeader) o;
    return (getAuthenticationToken() == that.getAuthenticationToken())
        && (getTimestamp() == that.getTimestamp())
        && (getRequestHandle() == that.getRequestHandle())
        && (getReturnDiagnostics() == that.getReturnDiagnostics())
        && (getAuditEntryId() == that.getAuditEntryId())
        && (getTimeoutHint() == that.getTimeoutHint())
        && (getAdditionalHeader() == that.getAdditionalHeader())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getAuthenticationToken(),
        getTimestamp(),
        getRequestHandle(),
        getReturnDiagnostics(),
        getAuditEntryId(),
        getTimeoutHint(),
        getAdditionalHeader());
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
