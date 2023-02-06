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

public class BACnetLogRecord implements Message {

  // Properties.
  protected final BACnetDateTimeEnclosed timestamp;
  protected final BACnetLogRecordLogDatum logDatum;
  protected final BACnetStatusFlagsTagged statusFlags;

  public BACnetLogRecord(
      BACnetDateTimeEnclosed timestamp,
      BACnetLogRecordLogDatum logDatum,
      BACnetStatusFlagsTagged statusFlags) {
    super();
    this.timestamp = timestamp;
    this.logDatum = logDatum;
    this.statusFlags = statusFlags;
  }

  public BACnetDateTimeEnclosed getTimestamp() {
    return timestamp;
  }

  public BACnetLogRecordLogDatum getLogDatum() {
    return logDatum;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetLogRecord");

    // Simple Field (timestamp)
    writeSimpleField("timestamp", timestamp, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (logDatum)
    writeSimpleField("logDatum", logDatum, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (statusFlags) (Can be skipped, if the value is null)
    writeOptionalField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogRecord");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetLogRecord _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (timestamp)
    lengthInBits += timestamp.getLengthInBits();

    // Simple field (logDatum)
    lengthInBits += logDatum.getLengthInBits();

    // Optional Field (statusFlags)
    if (statusFlags != null) {
      lengthInBits += statusFlags.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetLogRecord staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetLogRecord staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLogRecord");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetDateTimeEnclosed timestamp =
        readSimpleField(
            "timestamp",
            new DataReaderComplexDefault<>(
                () -> BACnetDateTimeEnclosed.staticParse(readBuffer, (short) (0)), readBuffer));

    BACnetLogRecordLogDatum logDatum =
        readSimpleField(
            "logDatum",
            new DataReaderComplexDefault<>(
                () -> BACnetLogRecordLogDatum.staticParse(readBuffer, (short) (1)), readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readOptionalField(
            "statusFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (2), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetLogRecord");
    // Create the instance
    BACnetLogRecord _bACnetLogRecord;
    _bACnetLogRecord = new BACnetLogRecord(timestamp, logDatum, statusFlags);
    return _bACnetLogRecord;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogRecord)) {
      return false;
    }
    BACnetLogRecord that = (BACnetLogRecord) o;
    return (getTimestamp() == that.getTimestamp())
        && (getLogDatum() == that.getLogDatum())
        && (getStatusFlags() == that.getStatusFlags())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTimestamp(), getLogDatum(), getStatusFlags());
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
