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

public class NodeIdNumeric extends NodeIdTypeDefinition implements Message {

  // Accessors for discriminator values.
  public NodeIdType getNodeType() {
    return NodeIdType.nodeIdTypeNumeric;
  }

  // Properties.
  protected final int namespaceIndex;
  protected final long id;

  public NodeIdNumeric(int namespaceIndex, long id) {
    super();
    this.namespaceIndex = namespaceIndex;
    this.id = id;
  }

  public int getNamespaceIndex() {
    return namespaceIndex;
  }

  public long getId() {
    return id;
  }

  public String getIdentifier() {
    return String.valueOf(getId());
  }

  @Override
  protected void serializeNodeIdTypeDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("NodeIdNumeric");

    // Simple Field (namespaceIndex)
    writeSimpleField("namespaceIndex", namespaceIndex, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (id)
    writeSimpleField("id", id, writeUnsignedLong(writeBuffer, 32));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    String identifier = getIdentifier();
    writeBuffer.writeVirtual("identifier", identifier);

    writeBuffer.popContext("NodeIdNumeric");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NodeIdNumeric _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (namespaceIndex)
    lengthInBits += 16;

    // Simple field (id)
    lengthInBits += 32;

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static NodeIdTypeDefinitionBuilder staticParseNodeIdTypeDefinitionBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("NodeIdNumeric");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int namespaceIndex = readSimpleField("namespaceIndex", readUnsignedInt(readBuffer, 16));

    long id = readSimpleField("id", readUnsignedLong(readBuffer, 32));
    String identifier = readVirtualField("identifier", String.class, id);

    readBuffer.closeContext("NodeIdNumeric");
    // Create the instance
    return new NodeIdNumericBuilderImpl(namespaceIndex, id);
  }

  public static class NodeIdNumericBuilderImpl
      implements NodeIdTypeDefinition.NodeIdTypeDefinitionBuilder {
    private final int namespaceIndex;
    private final long id;

    public NodeIdNumericBuilderImpl(int namespaceIndex, long id) {
      this.namespaceIndex = namespaceIndex;
      this.id = id;
    }

    public NodeIdNumeric build() {
      NodeIdNumeric nodeIdNumeric = new NodeIdNumeric(namespaceIndex, id);
      return nodeIdNumeric;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NodeIdNumeric)) {
      return false;
    }
    NodeIdNumeric that = (NodeIdNumeric) o;
    return (getNamespaceIndex() == that.getNamespaceIndex())
        && (getId() == that.getId())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNamespaceIndex(), getId());
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
