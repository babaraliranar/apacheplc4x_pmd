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

public class GuidNodeId implements Message {

  // Properties.
  protected final int namespaceIndex;
  protected final GuidValue identifier;

  public GuidNodeId(int namespaceIndex, GuidValue identifier) {
    super();
    this.namespaceIndex = namespaceIndex;
    this.identifier = identifier;
  }

  public int getNamespaceIndex() {
    return namespaceIndex;
  }

  public GuidValue getIdentifier() {
    return identifier;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("GuidNodeId");

    // Simple Field (namespaceIndex)
    writeSimpleField("namespaceIndex", namespaceIndex, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (identifier)
    writeSimpleField("identifier", identifier, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("GuidNodeId");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    GuidNodeId _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (namespaceIndex)
    lengthInBits += 16;

    // Simple field (identifier)
    lengthInBits += identifier.getLengthInBits();

    return lengthInBits;
  }

  public static GuidNodeId staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("GuidNodeId");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int namespaceIndex = readSimpleField("namespaceIndex", readUnsignedInt(readBuffer, 16));

    GuidValue identifier =
        readSimpleField(
            "identifier",
            new DataReaderComplexDefault<>(() -> GuidValue.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("GuidNodeId");
    // Create the instance
    GuidNodeId _guidNodeId;
    _guidNodeId = new GuidNodeId(namespaceIndex, identifier);
    return _guidNodeId;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof GuidNodeId)) {
      return false;
    }
    GuidNodeId that = (GuidNodeId) o;
    return (getNamespaceIndex() == that.getNamespaceIndex())
        && (getIdentifier() == that.getIdentifier())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNamespaceIndex(), getIdentifier());
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
