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

public class APDUUnknown extends APDU implements Message {

  // Accessors for discriminator values.
  public ApduType getApduType() {
    return null;
  }

  // Properties.
  protected final byte unknownTypeRest;
  protected final byte[] unknownBytes;

  // Arguments.
  protected final Integer apduLength;

  public APDUUnknown(byte unknownTypeRest, byte[] unknownBytes, Integer apduLength) {
    super(apduLength);
    this.unknownTypeRest = unknownTypeRest;
    this.unknownBytes = unknownBytes;
    this.apduLength = apduLength;
  }

  public byte getUnknownTypeRest() {
    return unknownTypeRest;
  }

  public byte[] getUnknownBytes() {
    return unknownBytes;
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("APDUUnknown");

    // Simple Field (unknownTypeRest)
    writeSimpleField("unknownTypeRest", unknownTypeRest, writeUnsignedByte(writeBuffer, 4));

    // Array Field (unknownBytes)
    writeByteArrayField("unknownBytes", unknownBytes, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("APDUUnknown");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUUnknown _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (unknownTypeRest)
    lengthInBits += 4;

    // Array field
    if (unknownBytes != null) {
      lengthInBits += 8 * unknownBytes.length;
    }

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("APDUUnknown");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte unknownTypeRest = readSimpleField("unknownTypeRest", readUnsignedByte(readBuffer, 4));

    byte[] unknownBytes =
        readBuffer.readByteArray(
            "unknownBytes", Math.toIntExact(((((apduLength) > (0))) ? apduLength : 0)));

    readBuffer.closeContext("APDUUnknown");
    // Create the instance
    return new APDUUnknownBuilderImpl(unknownTypeRest, unknownBytes, apduLength);
  }

  public static class APDUUnknownBuilderImpl implements APDU.APDUBuilder {
    private final byte unknownTypeRest;
    private final byte[] unknownBytes;
    private final Integer apduLength;

    public APDUUnknownBuilderImpl(byte unknownTypeRest, byte[] unknownBytes, Integer apduLength) {
      this.unknownTypeRest = unknownTypeRest;
      this.unknownBytes = unknownBytes;
      this.apduLength = apduLength;
    }

    public APDUUnknown build(Integer apduLength) {

      APDUUnknown aPDUUnknown = new APDUUnknown(unknownTypeRest, unknownBytes, apduLength);
      return aPDUUnknown;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUUnknown)) {
      return false;
    }
    APDUUnknown that = (APDUUnknown) o;
    return (getUnknownTypeRest() == that.getUnknownTypeRest())
        && (getUnknownBytes() == that.getUnknownBytes())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnknownTypeRest(), getUnknownBytes());
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
