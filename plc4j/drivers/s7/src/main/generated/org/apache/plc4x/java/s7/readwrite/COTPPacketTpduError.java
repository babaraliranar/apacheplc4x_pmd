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

public class COTPPacketTpduError extends COTPPacket implements Message {

  // Accessors for discriminator values.
  public Short getTpduCode() {
    return (short) 0x70;
  }

  // Properties.
  protected final int destinationReference;
  protected final short rejectCause;

  public COTPPacketTpduError(
      List<COTPParameter> parameters,
      S7Message payload,
      int destinationReference,
      short rejectCause) {
    super(parameters, payload);
    this.destinationReference = destinationReference;
    this.rejectCause = rejectCause;
  }

  public int getDestinationReference() {
    return destinationReference;
  }

  public short getRejectCause() {
    return rejectCause;
  }

  @Override
  protected void serializeCOTPPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("COTPPacketTpduError");

    // Simple Field (destinationReference)
    writeSimpleField(
        "destinationReference", destinationReference, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (rejectCause)
    writeSimpleField("rejectCause", rejectCause, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("COTPPacketTpduError");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    COTPPacketTpduError _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (destinationReference)
    lengthInBits += 16;

    // Simple field (rejectCause)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static COTPPacketBuilder staticParseCOTPPacketBuilder(
      ReadBuffer readBuffer, Integer cotpLen) throws ParseException {
    readBuffer.pullContext("COTPPacketTpduError");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int destinationReference =
        readSimpleField("destinationReference", readUnsignedInt(readBuffer, 16));

    short rejectCause = readSimpleField("rejectCause", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("COTPPacketTpduError");
    // Create the instance
    return new COTPPacketTpduErrorBuilderImpl(destinationReference, rejectCause);
  }

  public static class COTPPacketTpduErrorBuilderImpl implements COTPPacket.COTPPacketBuilder {
    private final int destinationReference;
    private final short rejectCause;

    public COTPPacketTpduErrorBuilderImpl(int destinationReference, short rejectCause) {
      this.destinationReference = destinationReference;
      this.rejectCause = rejectCause;
    }

    public COTPPacketTpduError build(List<COTPParameter> parameters, S7Message payload) {
      COTPPacketTpduError cOTPPacketTpduError =
          new COTPPacketTpduError(parameters, payload, destinationReference, rejectCause);
      return cOTPPacketTpduError;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof COTPPacketTpduError)) {
      return false;
    }
    COTPPacketTpduError that = (COTPPacketTpduError) o;
    return (getDestinationReference() == that.getDestinationReference())
        && (getRejectCause() == that.getRejectCause())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDestinationReference(), getRejectCause());
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
