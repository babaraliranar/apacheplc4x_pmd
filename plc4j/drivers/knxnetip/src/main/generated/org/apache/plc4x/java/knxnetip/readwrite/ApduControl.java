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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public abstract class ApduControl implements Message {

  // Abstract accessors for discriminator values.
  public abstract Byte getControlType();

  public ApduControl() {
    super();
  }

  protected abstract void serializeApduControlChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApduControl");

    // Discriminator Field (controlType) (Used as input to a switch field)
    writeDiscriminatorField("controlType", getControlType(), writeUnsignedByte(writeBuffer, 2));

    // Switch field (Serialize the sub-type)
    serializeApduControlChild(writeBuffer);

    writeBuffer.popContext("ApduControl");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ApduControl _value = this;

    // Discriminator Field (controlType)
    lengthInBits += 2;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static ApduControl staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static ApduControl staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ApduControl");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte controlType = readDiscriminatorField("controlType", readUnsignedByte(readBuffer, 2));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    ApduControlBuilder builder = null;
    if (EvaluationHelper.equals(controlType, (byte) 0x0)) {
      builder = ApduControlConnect.staticParseApduControlBuilder(readBuffer);
    } else if (EvaluationHelper.equals(controlType, (byte) 0x1)) {
      builder = ApduControlDisconnect.staticParseApduControlBuilder(readBuffer);
    } else if (EvaluationHelper.equals(controlType, (byte) 0x2)) {
      builder = ApduControlAck.staticParseApduControlBuilder(readBuffer);
    } else if (EvaluationHelper.equals(controlType, (byte) 0x3)) {
      builder = ApduControlNack.staticParseApduControlBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "controlType="
              + controlType
              + "]");
    }

    readBuffer.closeContext("ApduControl");
    // Create the instance
    ApduControl _apduControl = builder.build();
    return _apduControl;
  }

  public interface ApduControlBuilder {
    ApduControl build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduControl)) {
      return false;
    }
    ApduControl that = (ApduControl) o;
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
