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

public abstract class MonitoredSAL implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final byte salType;

  // Arguments.
  protected final CBusOptions cBusOptions;

  public MonitoredSAL(byte salType, CBusOptions cBusOptions) {
    super();
    this.salType = salType;
    this.cBusOptions = cBusOptions;
  }

  public byte getSalType() {
    return salType;
  }

  protected abstract void serializeMonitoredSALChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MonitoredSAL");

    // Switch field (Serialize the sub-type)
    serializeMonitoredSALChild(writeBuffer);

    writeBuffer.popContext("MonitoredSAL");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    MonitoredSAL _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static MonitoredSAL staticParse(ReadBuffer readBuffer, CBusOptions cBusOptions)
      throws ParseException {
    readBuffer.pullContext("MonitoredSAL");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte salType = readPeekField("salType", readByte(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    MonitoredSALBuilder builder = null;
    if (EvaluationHelper.equals(salType, (byte) 0x05)) {
      builder =
          MonitoredSALLongFormSmartMode.staticParseMonitoredSALBuilder(readBuffer, cBusOptions);
    } else if (true) {
      builder =
          MonitoredSALShortFormBasicMode.staticParseMonitoredSALBuilder(readBuffer, cBusOptions);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "salType=" + salType + "]");
    }

    readBuffer.closeContext("MonitoredSAL");
    // Create the instance
    MonitoredSAL _monitoredSAL = builder.build(salType, cBusOptions);
    return _monitoredSAL;
  }

  public interface MonitoredSALBuilder {
    MonitoredSAL build(byte salType, CBusOptions cBusOptions);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoredSAL)) {
      return false;
    }
    MonitoredSAL that = (MonitoredSAL) o;
    return (getSalType() == that.getSalType()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getSalType());
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
