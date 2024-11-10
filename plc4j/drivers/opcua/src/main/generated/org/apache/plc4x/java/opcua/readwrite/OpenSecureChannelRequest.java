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

public class OpenSecureChannelRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 446;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final long clientProtocolVersion;
  protected final SecurityTokenRequestType requestType;
  protected final MessageSecurityMode securityMode;
  protected final PascalByteString clientNonce;
  protected final long requestedLifetime;

  public OpenSecureChannelRequest(
      RequestHeader requestHeader,
      long clientProtocolVersion,
      SecurityTokenRequestType requestType,
      MessageSecurityMode securityMode,
      PascalByteString clientNonce,
      long requestedLifetime) {
    super();
    this.requestHeader = requestHeader;
    this.clientProtocolVersion = clientProtocolVersion;
    this.requestType = requestType;
    this.securityMode = securityMode;
    this.clientNonce = clientNonce;
    this.requestedLifetime = requestedLifetime;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public long getClientProtocolVersion() {
    return clientProtocolVersion;
  }

  public SecurityTokenRequestType getRequestType() {
    return requestType;
  }

  public MessageSecurityMode getSecurityMode() {
    return securityMode;
  }

  public PascalByteString getClientNonce() {
    return clientNonce;
  }

  public long getRequestedLifetime() {
    return requestedLifetime;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenSecureChannelRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Simple Field (clientProtocolVersion)
    writeSimpleField(
        "clientProtocolVersion", clientProtocolVersion, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (requestType)
    writeSimpleEnumField(
        "requestType",
        "SecurityTokenRequestType",
        requestType,
        writeEnum(
            SecurityTokenRequestType::getValue,
            SecurityTokenRequestType::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (securityMode)
    writeSimpleEnumField(
        "securityMode",
        "MessageSecurityMode",
        securityMode,
        writeEnum(
            MessageSecurityMode::getValue,
            MessageSecurityMode::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (clientNonce)
    writeSimpleField("clientNonce", clientNonce, writeComplex(writeBuffer));

    // Simple Field (requestedLifetime)
    writeSimpleField("requestedLifetime", requestedLifetime, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("OpenSecureChannelRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenSecureChannelRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (clientProtocolVersion)
    lengthInBits += 32;

    // Simple field (requestType)
    lengthInBits += 32;

    // Simple field (securityMode)
    lengthInBits += 32;

    // Simple field (clientNonce)
    lengthInBits += clientNonce.getLengthInBits();

    // Simple field (requestedLifetime)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("OpenSecureChannelRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    long clientProtocolVersion =
        readSimpleField("clientProtocolVersion", readUnsignedLong(readBuffer, 32));

    SecurityTokenRequestType requestType =
        readEnumField(
            "requestType",
            "SecurityTokenRequestType",
            readEnum(SecurityTokenRequestType::enumForValue, readUnsignedLong(readBuffer, 32)));

    MessageSecurityMode securityMode =
        readEnumField(
            "securityMode",
            "MessageSecurityMode",
            readEnum(MessageSecurityMode::enumForValue, readUnsignedLong(readBuffer, 32)));

    PascalByteString clientNonce =
        readSimpleField(
            "clientNonce", readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer));

    long requestedLifetime = readSimpleField("requestedLifetime", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("OpenSecureChannelRequest");
    // Create the instance
    return new OpenSecureChannelRequestBuilderImpl(
        requestHeader,
        clientProtocolVersion,
        requestType,
        securityMode,
        clientNonce,
        requestedLifetime);
  }

  public static class OpenSecureChannelRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final long clientProtocolVersion;
    private final SecurityTokenRequestType requestType;
    private final MessageSecurityMode securityMode;
    private final PascalByteString clientNonce;
    private final long requestedLifetime;

    public OpenSecureChannelRequestBuilderImpl(
        RequestHeader requestHeader,
        long clientProtocolVersion,
        SecurityTokenRequestType requestType,
        MessageSecurityMode securityMode,
        PascalByteString clientNonce,
        long requestedLifetime) {
      this.requestHeader = requestHeader;
      this.clientProtocolVersion = clientProtocolVersion;
      this.requestType = requestType;
      this.securityMode = securityMode;
      this.clientNonce = clientNonce;
      this.requestedLifetime = requestedLifetime;
    }

    public OpenSecureChannelRequest build() {
      OpenSecureChannelRequest openSecureChannelRequest =
          new OpenSecureChannelRequest(
              requestHeader,
              clientProtocolVersion,
              requestType,
              securityMode,
              clientNonce,
              requestedLifetime);
      return openSecureChannelRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenSecureChannelRequest)) {
      return false;
    }
    OpenSecureChannelRequest that = (OpenSecureChannelRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getClientProtocolVersion() == that.getClientProtocolVersion())
        && (getRequestType() == that.getRequestType())
        && (getSecurityMode() == that.getSecurityMode())
        && (getClientNonce() == that.getClientNonce())
        && (getRequestedLifetime() == that.getRequestedLifetime())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getClientProtocolVersion(),
        getRequestType(),
        getSecurityMode(),
        getClientNonce(),
        getRequestedLifetime());
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
