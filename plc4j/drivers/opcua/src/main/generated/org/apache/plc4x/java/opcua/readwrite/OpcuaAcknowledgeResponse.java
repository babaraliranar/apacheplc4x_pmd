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

public class OpcuaAcknowledgeResponse extends MessagePDU implements Message {

  // Accessors for discriminator values.
  public String getMessageType() {
    return (String) "ACK";
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final String chunk;
  protected final int version;
  protected final int receiveBufferSize;
  protected final int sendBufferSize;
  protected final int maxMessageSize;
  protected final int maxChunkCount;

  public OpcuaAcknowledgeResponse(
      String chunk,
      int version,
      int receiveBufferSize,
      int sendBufferSize,
      int maxMessageSize,
      int maxChunkCount) {
    super();
    this.chunk = chunk;
    this.version = version;
    this.receiveBufferSize = receiveBufferSize;
    this.sendBufferSize = sendBufferSize;
    this.maxMessageSize = maxMessageSize;
    this.maxChunkCount = maxChunkCount;
  }

  public String getChunk() {
    return chunk;
  }

  public int getVersion() {
    return version;
  }

  public int getReceiveBufferSize() {
    return receiveBufferSize;
  }

  public int getSendBufferSize() {
    return sendBufferSize;
  }

  public int getMaxMessageSize() {
    return maxMessageSize;
  }

  public int getMaxChunkCount() {
    return maxChunkCount;
  }

  @Override
  protected void serializeMessagePDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpcuaAcknowledgeResponse");

    // Simple Field (chunk)
    writeSimpleField("chunk", chunk, writeString(writeBuffer, 8));

    // Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int messageSize = (int) (getLengthInBytes());
    writeImplicitField("messageSize", messageSize, writeSignedInt(writeBuffer, 32));

    // Simple Field (version)
    writeSimpleField("version", version, writeSignedInt(writeBuffer, 32));

    // Simple Field (receiveBufferSize)
    writeSimpleField("receiveBufferSize", receiveBufferSize, writeSignedInt(writeBuffer, 32));

    // Simple Field (sendBufferSize)
    writeSimpleField("sendBufferSize", sendBufferSize, writeSignedInt(writeBuffer, 32));

    // Simple Field (maxMessageSize)
    writeSimpleField("maxMessageSize", maxMessageSize, writeSignedInt(writeBuffer, 32));

    // Simple Field (maxChunkCount)
    writeSimpleField("maxChunkCount", maxChunkCount, writeSignedInt(writeBuffer, 32));

    writeBuffer.popContext("OpcuaAcknowledgeResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpcuaAcknowledgeResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (chunk)
    lengthInBits += 8;

    // Implicit Field (messageSize)
    lengthInBits += 32;

    // Simple field (version)
    lengthInBits += 32;

    // Simple field (receiveBufferSize)
    lengthInBits += 32;

    // Simple field (sendBufferSize)
    lengthInBits += 32;

    // Simple field (maxMessageSize)
    lengthInBits += 32;

    // Simple field (maxChunkCount)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static MessagePDUBuilder staticParseMessagePDUBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("OpcuaAcknowledgeResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    String chunk = readSimpleField("chunk", readString(readBuffer, 8));

    int messageSize = readImplicitField("messageSize", readSignedInt(readBuffer, 32));

    int version = readSimpleField("version", readSignedInt(readBuffer, 32));

    int receiveBufferSize = readSimpleField("receiveBufferSize", readSignedInt(readBuffer, 32));

    int sendBufferSize = readSimpleField("sendBufferSize", readSignedInt(readBuffer, 32));

    int maxMessageSize = readSimpleField("maxMessageSize", readSignedInt(readBuffer, 32));

    int maxChunkCount = readSimpleField("maxChunkCount", readSignedInt(readBuffer, 32));

    readBuffer.closeContext("OpcuaAcknowledgeResponse");
    // Create the instance
    return new OpcuaAcknowledgeResponseBuilderImpl(
        chunk, version, receiveBufferSize, sendBufferSize, maxMessageSize, maxChunkCount);
  }

  public static class OpcuaAcknowledgeResponseBuilderImpl implements MessagePDU.MessagePDUBuilder {
    private final String chunk;
    private final int version;
    private final int receiveBufferSize;
    private final int sendBufferSize;
    private final int maxMessageSize;
    private final int maxChunkCount;

    public OpcuaAcknowledgeResponseBuilderImpl(
        String chunk,
        int version,
        int receiveBufferSize,
        int sendBufferSize,
        int maxMessageSize,
        int maxChunkCount) {
      this.chunk = chunk;
      this.version = version;
      this.receiveBufferSize = receiveBufferSize;
      this.sendBufferSize = sendBufferSize;
      this.maxMessageSize = maxMessageSize;
      this.maxChunkCount = maxChunkCount;
    }

    public OpcuaAcknowledgeResponse build() {
      OpcuaAcknowledgeResponse opcuaAcknowledgeResponse =
          new OpcuaAcknowledgeResponse(
              chunk, version, receiveBufferSize, sendBufferSize, maxMessageSize, maxChunkCount);
      return opcuaAcknowledgeResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpcuaAcknowledgeResponse)) {
      return false;
    }
    OpcuaAcknowledgeResponse that = (OpcuaAcknowledgeResponse) o;
    return (getChunk() == that.getChunk())
        && (getVersion() == that.getVersion())
        && (getReceiveBufferSize() == that.getReceiveBufferSize())
        && (getSendBufferSize() == that.getSendBufferSize())
        && (getMaxMessageSize() == that.getMaxMessageSize())
        && (getMaxChunkCount() == that.getMaxChunkCount())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getChunk(),
        getVersion(),
        getReceiveBufferSize(),
        getSendBufferSize(),
        getMaxMessageSize(),
        getMaxChunkCount());
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
