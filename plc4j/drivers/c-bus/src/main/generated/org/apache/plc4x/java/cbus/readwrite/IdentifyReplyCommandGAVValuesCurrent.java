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
package org.apache.plc4x.java.cbus.readwrite;

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

public class IdentifyReplyCommandGAVValuesCurrent extends IdentifyReplyCommand implements Message {

  // Accessors for discriminator values.
  public Attribute getAttribute() {
    return Attribute.GAVValuesCurrent;
  }

  // Properties.
  protected final byte[] values;

  // Arguments.
  protected final Short numBytes;

  public IdentifyReplyCommandGAVValuesCurrent(byte[] values, Short numBytes) {
    super(numBytes);
    this.values = values;
    this.numBytes = numBytes;
  }

  public byte[] getValues() {
    return values;
  }

  @Override
  protected void serializeIdentifyReplyCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("IdentifyReplyCommandGAVValuesCurrent");

    // Array Field (values)
    writeByteArrayField("values", values, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("IdentifyReplyCommandGAVValuesCurrent");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IdentifyReplyCommandGAVValuesCurrent _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (values != null) {
      lengthInBits += 8 * values.length;
    }

    return lengthInBits;
  }

  public static IdentifyReplyCommandBuilder staticParseIdentifyReplyCommandBuilder(
      ReadBuffer readBuffer, Attribute attribute, Short numBytes) throws ParseException {
    readBuffer.pullContext("IdentifyReplyCommandGAVValuesCurrent");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte[] values = readBuffer.readByteArray("values", Math.toIntExact(numBytes));

    readBuffer.closeContext("IdentifyReplyCommandGAVValuesCurrent");
    // Create the instance
    return new IdentifyReplyCommandGAVValuesCurrentBuilderImpl(values, numBytes);
  }

  public static class IdentifyReplyCommandGAVValuesCurrentBuilderImpl
      implements IdentifyReplyCommand.IdentifyReplyCommandBuilder {
    private final byte[] values;
    private final Short numBytes;

    public IdentifyReplyCommandGAVValuesCurrentBuilderImpl(byte[] values, Short numBytes) {
      this.values = values;
      this.numBytes = numBytes;
    }

    public IdentifyReplyCommandGAVValuesCurrent build(Short numBytes) {

      IdentifyReplyCommandGAVValuesCurrent identifyReplyCommandGAVValuesCurrent =
          new IdentifyReplyCommandGAVValuesCurrent(values, numBytes);
      return identifyReplyCommandGAVValuesCurrent;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IdentifyReplyCommandGAVValuesCurrent)) {
      return false;
    }
    IdentifyReplyCommandGAVValuesCurrent that = (IdentifyReplyCommandGAVValuesCurrent) o;
    return (getValues() == that.getValues()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getValues());
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
