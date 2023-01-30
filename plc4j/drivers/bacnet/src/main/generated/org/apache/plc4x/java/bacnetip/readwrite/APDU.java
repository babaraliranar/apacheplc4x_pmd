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

public abstract class APDU implements Message {

  // Abstract accessors for discriminator values.
  public abstract ApduType getApduType();

  // Arguments.
  protected final Integer apduLength;

  public APDU(Integer apduLength) {
    super();
    this.apduLength = apduLength;
  }

  protected abstract void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("APDU");

    // Discriminator Field (apduType) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "apduType",
        "ApduType",
        getApduType(),
        new DataWriterEnumDefault<>(
            ApduType::getValue, ApduType::name, writeUnsignedByte(writeBuffer, 4)));

    // Switch field (Serialize the sub-type)
    serializeAPDUChild(writeBuffer);

    writeBuffer.popContext("APDU");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    APDU _value = this;

    // Discriminator Field (apduType)
    lengthInBits += 4;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static APDU staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Integer apduLength;
    if (args[0] instanceof Integer) {
      apduLength = (Integer) args[0];
    } else if (args[0] instanceof String) {
      apduLength = Integer.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Integer or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, apduLength);
  }

  public static APDU staticParse(ReadBuffer readBuffer, Integer apduLength) throws ParseException {
    readBuffer.pullContext("APDU");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ApduType apduType =
        readDiscriminatorField(
            "apduType",
            new DataReaderEnumDefault<>(ApduType::enumForValue, readUnsignedByte(readBuffer, 4)));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    APDUBuilder builder = null;
    if (EvaluationHelper.equals(apduType, ApduType.CONFIRMED_REQUEST_PDU)) {
      builder = APDUConfirmedRequest.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.UNCONFIRMED_REQUEST_PDU)) {
      builder = APDUUnconfirmedRequest.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.SIMPLE_ACK_PDU)) {
      builder = APDUSimpleAck.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.COMPLEX_ACK_PDU)) {
      builder = APDUComplexAck.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.SEGMENT_ACK_PDU)) {
      builder = APDUSegmentAck.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.ERROR_PDU)) {
      builder = APDUError.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.REJECT_PDU)) {
      builder = APDUReject.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(apduType, ApduType.ABORT_PDU)) {
      builder = APDUAbort.staticParseAPDUBuilder(readBuffer, apduLength);
    } else if (true) {
      builder = APDUUnknown.staticParseAPDUBuilder(readBuffer, apduLength);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "apduType="
              + apduType
              + "]");
    }

    readBuffer.closeContext("APDU");
    // Create the instance
    APDU _aPDU = builder.build(apduLength);

    return _aPDU;
  }

  public interface APDUBuilder {
    APDU build(Integer apduLength);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDU)) {
      return false;
    }
    APDU that = (APDU) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
