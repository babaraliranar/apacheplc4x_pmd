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

public class ExpandedNodeId implements Message {

  // Properties.
  protected final boolean namespaceURISpecified;
  protected final boolean serverIndexSpecified;
  protected final NodeIdTypeDefinition nodeId;
  protected final PascalString namespaceURI;
  protected final Long serverIndex;

  public ExpandedNodeId(
      boolean namespaceURISpecified,
      boolean serverIndexSpecified,
      NodeIdTypeDefinition nodeId,
      PascalString namespaceURI,
      Long serverIndex) {
    super();
    this.namespaceURISpecified = namespaceURISpecified;
    this.serverIndexSpecified = serverIndexSpecified;
    this.nodeId = nodeId;
    this.namespaceURI = namespaceURI;
    this.serverIndex = serverIndex;
  }

  public boolean getNamespaceURISpecified() {
    return namespaceURISpecified;
  }

  public boolean getServerIndexSpecified() {
    return serverIndexSpecified;
  }

  public NodeIdTypeDefinition getNodeId() {
    return nodeId;
  }

  public PascalString getNamespaceURI() {
    return namespaceURI;
  }

  public Long getServerIndex() {
    return serverIndex;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ExpandedNodeId");

    // Simple Field (namespaceURISpecified)
    writeSimpleField("namespaceURISpecified", namespaceURISpecified, writeBoolean(writeBuffer));

    // Simple Field (serverIndexSpecified)
    writeSimpleField("serverIndexSpecified", serverIndexSpecified, writeBoolean(writeBuffer));

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, writeComplex(writeBuffer));

    // Optional Field (namespaceURI) (Can be skipped, if the value is null)
    writeOptionalField("namespaceURI", namespaceURI, writeComplex(writeBuffer));

    // Optional Field (serverIndex) (Can be skipped, if the value is null)
    writeOptionalField("serverIndex", serverIndex, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("ExpandedNodeId");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ExpandedNodeId _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (namespaceURISpecified)
    lengthInBits += 1;

    // Simple field (serverIndexSpecified)
    lengthInBits += 1;

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Optional Field (namespaceURI)
    if (namespaceURI != null) {
      lengthInBits += namespaceURI.getLengthInBits();
    }

    // Optional Field (serverIndex)
    if (serverIndex != null) {
      lengthInBits += 32;
    }

    return lengthInBits;
  }

  public static ExpandedNodeId staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ExpandedNodeId");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean namespaceURISpecified =
        readSimpleField("namespaceURISpecified", readBoolean(readBuffer));

    boolean serverIndexSpecified = readSimpleField("serverIndexSpecified", readBoolean(readBuffer));

    NodeIdTypeDefinition nodeId =
        readSimpleField(
            "nodeId", readComplex(() -> NodeIdTypeDefinition.staticParse(readBuffer), readBuffer));

    PascalString namespaceURI =
        readOptionalField(
            "namespaceURI",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            namespaceURISpecified);

    Long serverIndex =
        readOptionalField("serverIndex", readUnsignedLong(readBuffer, 32), serverIndexSpecified);

    readBuffer.closeContext("ExpandedNodeId");
    // Create the instance
    ExpandedNodeId _expandedNodeId;
    _expandedNodeId =
        new ExpandedNodeId(
            namespaceURISpecified, serverIndexSpecified, nodeId, namespaceURI, serverIndex);
    return _expandedNodeId;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ExpandedNodeId)) {
      return false;
    }
    ExpandedNodeId that = (ExpandedNodeId) o;
    return (getNamespaceURISpecified() == that.getNamespaceURISpecified())
        && (getServerIndexSpecified() == that.getServerIndexSpecified())
        && (getNodeId() == that.getNodeId())
        && (getNamespaceURI() == that.getNamespaceURI())
        && (getServerIndex() == that.getServerIndex())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getNamespaceURISpecified(),
        getServerIndexSpecified(),
        getNodeId(),
        getNamespaceURI(),
        getServerIndex());
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
