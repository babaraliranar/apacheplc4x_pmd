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

public class ExtensionObject implements Message {

  // Properties.
  protected final ExpandedNodeId typeId;
  protected final ExtensionObjectEncodingMask encodingMask;
  protected final ExtensionObjectDefinition body;

  public ExtensionObject(
      ExpandedNodeId typeId,
      ExtensionObjectEncodingMask encodingMask,
      ExtensionObjectDefinition body) {
    super();
    this.typeId = typeId;
    this.encodingMask = encodingMask;
    this.body = body;
  }

  public ExpandedNodeId getTypeId() {
    return typeId;
  }

  public ExtensionObjectEncodingMask getEncodingMask() {
    return encodingMask;
  }

  public ExtensionObjectDefinition getBody() {
    return body;
  }

  public String getIdentifier() {
    return String.valueOf(getTypeId().getIdentifier());
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ExtensionObject");

    // Simple Field (typeId)
    writeSimpleField("typeId", typeId, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (encodingMask) (Can be skipped, if the value is null)
    writeOptionalField("encodingMask", encodingMask, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    String identifier = getIdentifier();
    writeBuffer.writeVirtual("identifier", identifier);

    // Simple Field (body)
    writeSimpleField("body", body, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ExtensionObject");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ExtensionObject _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (typeId)
    lengthInBits += typeId.getLengthInBits();

    // Optional Field (encodingMask)
    if (encodingMask != null) {
      lengthInBits += encodingMask.getLengthInBits();
    }

    // A virtual field doesn't have any in- or output.

    // Simple field (body)
    lengthInBits += body.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObject staticParse(ReadBuffer readBuffer, Boolean includeEncodingMask)
      throws ParseException {
    readBuffer.pullContext("ExtensionObject");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExpandedNodeId typeId =
        readSimpleField(
            "typeId",
            new DataReaderComplexDefault<>(
                () -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    ExtensionObjectEncodingMask encodingMask =
        readOptionalField(
            "encodingMask",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectEncodingMask.staticParse(readBuffer), readBuffer),
            includeEncodingMask);
    String identifier = readVirtualField("identifier", String.class, typeId.getIdentifier());

    ExtensionObjectDefinition body =
        readSimpleField(
            "body",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) (identifier)),
                readBuffer));

    readBuffer.closeContext("ExtensionObject");
    // Create the instance
    ExtensionObject _extensionObject;
    _extensionObject = new ExtensionObject(typeId, encodingMask, body);
    return _extensionObject;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ExtensionObject)) {
      return false;
    }
    ExtensionObject that = (ExtensionObject) o;
    return (getTypeId() == that.getTypeId())
        && (getEncodingMask() == that.getEncodingMask())
        && (getBody() == that.getBody())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTypeId(), getEncodingMask(), getBody());
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
