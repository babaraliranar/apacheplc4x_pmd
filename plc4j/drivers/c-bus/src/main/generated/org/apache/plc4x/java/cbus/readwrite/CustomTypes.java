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

public class CustomTypes implements Message {

  // Properties.
  protected final String customString;

  // Arguments.
  protected final Short numBytes;

  public CustomTypes(String customString, Short numBytes) {
    super();
    this.customString = customString;
    this.numBytes = numBytes;
  }

  public String getCustomString() {
    return customString;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CustomTypes");

    // Simple Field (customString)
    writeSimpleField("customString", customString, writeString(writeBuffer, (8) * (numBytes)));

    writeBuffer.popContext("CustomTypes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CustomTypes _value = this;

    // Simple field (customString)
    lengthInBits += (8) * (numBytes);

    return lengthInBits;
  }

  public static CustomTypes staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Short numBytes;
    if (args[0] instanceof Short) {
      numBytes = (Short) args[0];
    } else if (args[0] instanceof String) {
      numBytes = Short.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Short or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, numBytes);
  }

  public static CustomTypes staticParse(ReadBuffer readBuffer, Short numBytes)
      throws ParseException {
    readBuffer.pullContext("CustomTypes");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    String customString = readSimpleField("customString", readString(readBuffer, (8) * (numBytes)));

    readBuffer.closeContext("CustomTypes");
    // Create the instance
    CustomTypes _customTypes;
    _customTypes = new CustomTypes(customString, numBytes);
    return _customTypes;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CustomTypes)) {
      return false;
    }
    CustomTypes that = (CustomTypes) o;
    return (getCustomString() == that.getCustomString()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCustomString());
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
