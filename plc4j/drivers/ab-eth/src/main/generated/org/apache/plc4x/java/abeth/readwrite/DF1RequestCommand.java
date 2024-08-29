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
package org.apache.plc4x.java.abeth.readwrite;

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

public abstract class DF1RequestCommand implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getFunctionCode();

  public DF1RequestCommand() {
    super();
  }

  protected abstract void serializeDF1RequestCommandChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DF1RequestCommand");

    // Discriminator Field (functionCode) (Used as input to a switch field)
    writeDiscriminatorField("functionCode", getFunctionCode(), writeUnsignedShort(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeDF1RequestCommandChild(writeBuffer);

    writeBuffer.popContext("DF1RequestCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DF1RequestCommand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (functionCode)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static DF1RequestCommand staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DF1RequestCommand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short functionCode = readDiscriminatorField("functionCode", readUnsignedShort(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    DF1RequestCommandBuilder builder = null;
    if (EvaluationHelper.equals(functionCode, (short) 0xA2)) {
      builder = DF1RequestProtectedTypedLogicalRead.staticParseDF1RequestCommandBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "functionCode="
              + functionCode
              + "]");
    }

    readBuffer.closeContext("DF1RequestCommand");
    // Create the instance
    DF1RequestCommand _dF1RequestCommand = builder.build();
    return _dF1RequestCommand;
  }

  public interface DF1RequestCommandBuilder {
    DF1RequestCommand build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DF1RequestCommand)) {
      return false;
    }
    DF1RequestCommand that = (DF1RequestCommand) o;
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
