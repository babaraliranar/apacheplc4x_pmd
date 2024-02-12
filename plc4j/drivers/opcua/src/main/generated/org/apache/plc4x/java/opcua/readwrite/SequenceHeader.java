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
package org.apache.plc4x.java.opcua.readwrite;

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

public class SequenceHeader implements Message {

  // Properties.
  protected final int sequenceNumber;
  protected final int requestId;

  public SequenceHeader(int sequenceNumber, int requestId) {
    super();
    this.sequenceNumber = sequenceNumber;
    this.requestId = requestId;
  }

  public int getSequenceNumber() {
    return sequenceNumber;
  }

  public int getRequestId() {
    return requestId;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SequenceHeader");

    // Simple Field (sequenceNumber)
    writeSimpleField("sequenceNumber", sequenceNumber, writeSignedInt(writeBuffer, 32));

    // Simple Field (requestId)
    writeSimpleField("requestId", requestId, writeSignedInt(writeBuffer, 32));

    writeBuffer.popContext("SequenceHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SequenceHeader _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (sequenceNumber)
    lengthInBits += 32;

    // Simple field (requestId)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static SequenceHeader staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static SequenceHeader staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("SequenceHeader");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int sequenceNumber = readSimpleField("sequenceNumber", readSignedInt(readBuffer, 32));

    int requestId = readSimpleField("requestId", readSignedInt(readBuffer, 32));

    readBuffer.closeContext("SequenceHeader");
    // Create the instance
    SequenceHeader _sequenceHeader;
    _sequenceHeader = new SequenceHeader(sequenceNumber, requestId);
    return _sequenceHeader;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SequenceHeader)) {
      return false;
    }
    SequenceHeader that = (SequenceHeader) o;
    return (getSequenceNumber() == that.getSequenceNumber())
        && (getRequestId() == that.getRequestId())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getSequenceNumber(), getRequestId());
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
