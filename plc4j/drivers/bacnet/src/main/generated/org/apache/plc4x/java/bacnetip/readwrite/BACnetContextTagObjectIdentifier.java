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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetContextTagObjectIdentifier extends BACnetContextTag implements Message {

  // Accessors for discriminator values.
  public BACnetDataType getDataType() {
    return BACnetDataType.BACNET_OBJECT_IDENTIFIER;
  }

  // Properties.
  protected final BACnetTagPayloadObjectIdentifier payload;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetContextTagObjectIdentifier(
      BACnetTagHeader header, BACnetTagPayloadObjectIdentifier payload, Short tagNumberArgument) {
    super(header, tagNumberArgument);
    this.payload = payload;
    this.tagNumberArgument = tagNumberArgument;
  }

  public BACnetTagPayloadObjectIdentifier getPayload() {
    return payload;
  }

  public BACnetObjectType getObjectType() {
    return (BACnetObjectType) (getPayload().getObjectType());
  }

  public long getInstanceNumber() {
    return (long) (getPayload().getInstanceNumber());
  }

  @Override
  protected void serializeBACnetContextTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetContextTagObjectIdentifier");

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetObjectType objectType = getObjectType();
    writeBuffer.writeVirtual("objectType", objectType);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    long instanceNumber = getInstanceNumber();
    writeBuffer.writeVirtual("instanceNumber", instanceNumber);

    writeBuffer.popContext("BACnetContextTagObjectIdentifier");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetContextTagObjectIdentifier _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetContextTagBuilder staticParseBACnetContextTagBuilder(
      ReadBuffer readBuffer, Short tagNumberArgument, BACnetDataType dataType)
      throws ParseException {
    readBuffer.pullContext("BACnetContextTagObjectIdentifier");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagPayloadObjectIdentifier payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () -> BACnetTagPayloadObjectIdentifier.staticParse(readBuffer), readBuffer));
    BACnetObjectType objectType =
        readVirtualField("objectType", BACnetObjectType.class, payload.getObjectType());
    long instanceNumber =
        readVirtualField("instanceNumber", long.class, payload.getInstanceNumber());

    readBuffer.closeContext("BACnetContextTagObjectIdentifier");
    // Create the instance
    return new BACnetContextTagObjectIdentifierBuilderImpl(payload, tagNumberArgument);
  }

  public static class BACnetContextTagObjectIdentifierBuilderImpl
      implements BACnetContextTag.BACnetContextTagBuilder {
    private final BACnetTagPayloadObjectIdentifier payload;
    private final Short tagNumberArgument;

    public BACnetContextTagObjectIdentifierBuilderImpl(
        BACnetTagPayloadObjectIdentifier payload, Short tagNumberArgument) {
      this.payload = payload;
      this.tagNumberArgument = tagNumberArgument;
    }

    public BACnetContextTagObjectIdentifier build(BACnetTagHeader header, Short tagNumberArgument) {
      BACnetContextTagObjectIdentifier bACnetContextTagObjectIdentifier =
          new BACnetContextTagObjectIdentifier(header, payload, tagNumberArgument);
      return bACnetContextTagObjectIdentifier;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetContextTagObjectIdentifier)) {
      return false;
    }
    BACnetContextTagObjectIdentifier that = (BACnetContextTagObjectIdentifier) o;
    return (getPayload() == that.getPayload()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPayload());
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
