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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7PayloadReadVarResponse extends S7Payload implements Message {

  // Accessors for discriminator values.
  public Short getParameterParameterType() {
    return (short) 0x04;
  }

  public Short getMessageType() {
    return (short) 0x03;
  }

  // Properties.
  protected final List<S7VarPayloadDataItem> items;

  // Arguments.
  protected final S7Parameter parameter;

  public S7PayloadReadVarResponse(List<S7VarPayloadDataItem> items, S7Parameter parameter) {
    super(parameter);
    this.items = items;
    this.parameter = parameter;
  }

  public List<S7VarPayloadDataItem> getItems() {
    return items;
  }

  @Override
  protected void serializeS7PayloadChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadReadVarResponse");

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("S7PayloadReadVarResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadReadVarResponse _value = this;

    // Array field
    if (items != null) {
      int i = 0;
      for (S7VarPayloadDataItem element : items) {
        boolean last = ++i >= items.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7PayloadReadVarResponseBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short messageType, S7Parameter parameter) throws ParseException {
    readBuffer.pullContext("S7PayloadReadVarResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    List<S7VarPayloadDataItem> items =
        readCountArrayField(
            "items",
            new DataReaderComplexDefault<>(
                () -> S7VarPayloadDataItem.staticParse(readBuffer), readBuffer),
            CAST(parameter, S7ParameterReadVarResponse.class).getNumItems());

    readBuffer.closeContext("S7PayloadReadVarResponse");
    // Create the instance
    return new S7PayloadReadVarResponseBuilder(items, parameter);
  }

  public static class S7PayloadReadVarResponseBuilder implements S7Payload.S7PayloadBuilder {
    private final List<S7VarPayloadDataItem> items;
    private final S7Parameter parameter;

    public S7PayloadReadVarResponseBuilder(
        List<S7VarPayloadDataItem> items, S7Parameter parameter) {

      this.items = items;
      this.parameter = parameter;
    }

    public S7PayloadReadVarResponse build(S7Parameter parameter) {

      S7PayloadReadVarResponse s7PayloadReadVarResponse =
          new S7PayloadReadVarResponse(items, parameter);
      return s7PayloadReadVarResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadReadVarResponse)) {
      return false;
    }
    S7PayloadReadVarResponse that = (S7PayloadReadVarResponse) o;
    return (getItems() == that.getItems()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getItems());
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
