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
package org.apache.plc4x.java.canopen.readwrite;

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

public class SDOInitiateSegmentedReservedResponse extends SDOInitiateUploadResponsePayload
    implements Message {

  // Accessors for discriminator values.
  public Boolean getExpedited() {
    return (boolean) false;
  }

  public Boolean getIndicated() {
    return (boolean) false;
  }

  public SDOInitiateSegmentedReservedResponse() {
    super();
  }

  @Override
  protected void serializeSDOInitiateUploadResponsePayloadChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SDOInitiateSegmentedReservedResponse");

    // Reserved Field (reserved)
    writeReservedField("reserved", (int) 0x00, writeSignedInt(writeBuffer, 32));

    writeBuffer.popContext("SDOInitiateSegmentedReservedResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SDOInitiateSegmentedReservedResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static SDOInitiateUploadResponsePayloadBuilder
      staticParseSDOInitiateUploadResponsePayloadBuilder(
          ReadBuffer readBuffer, Boolean expedited, Boolean indicated, Byte size)
          throws ParseException {
    readBuffer.pullContext("SDOInitiateSegmentedReservedResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Integer reservedField0 =
        readReservedField("reserved", readSignedInt(readBuffer, 32), (int) 0x00);

    readBuffer.closeContext("SDOInitiateSegmentedReservedResponse");
    // Create the instance
    return new SDOInitiateSegmentedReservedResponseBuilderImpl();
  }

  public static class SDOInitiateSegmentedReservedResponseBuilderImpl
      implements SDOInitiateUploadResponsePayload.SDOInitiateUploadResponsePayloadBuilder {

    public SDOInitiateSegmentedReservedResponseBuilderImpl() {}

    public SDOInitiateSegmentedReservedResponse build() {
      SDOInitiateSegmentedReservedResponse sDOInitiateSegmentedReservedResponse =
          new SDOInitiateSegmentedReservedResponse();
      return sDOInitiateSegmentedReservedResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SDOInitiateSegmentedReservedResponse)) {
      return false;
    }
    SDOInitiateSegmentedReservedResponse that = (SDOInitiateSegmentedReservedResponse) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
