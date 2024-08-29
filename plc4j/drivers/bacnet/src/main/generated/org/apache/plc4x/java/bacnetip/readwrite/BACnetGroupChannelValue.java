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

public class BACnetGroupChannelValue implements Message {

  // Properties.
  protected final BACnetContextTagUnsignedInteger channel;
  protected final BACnetContextTagUnsignedInteger overridingPriority;
  protected final BACnetChannelValue value;

  public BACnetGroupChannelValue(
      BACnetContextTagUnsignedInteger channel,
      BACnetContextTagUnsignedInteger overridingPriority,
      BACnetChannelValue value) {
    super();
    this.channel = channel;
    this.overridingPriority = overridingPriority;
    this.value = value;
  }

  public BACnetContextTagUnsignedInteger getChannel() {
    return channel;
  }

  public BACnetContextTagUnsignedInteger getOverridingPriority() {
    return overridingPriority;
  }

  public BACnetChannelValue getValue() {
    return value;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetGroupChannelValue");

    // Simple Field (channel)
    writeSimpleField("channel", channel, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (overridingPriority) (Can be skipped, if the value is null)
    writeOptionalField(
        "overridingPriority", overridingPriority, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (value)
    writeSimpleField("value", value, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetGroupChannelValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetGroupChannelValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (channel)
    lengthInBits += channel.getLengthInBits();

    // Optional Field (overridingPriority)
    if (overridingPriority != null) {
      lengthInBits += overridingPriority.getLengthInBits();
    }

    // Simple field (value)
    lengthInBits += value.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetGroupChannelValue staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetGroupChannelValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger channel =
        readSimpleField(
            "channel",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger overridingPriority =
        readOptionalField(
            "overridingPriority",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetChannelValue value =
        readSimpleField(
            "value",
            new DataReaderComplexDefault<>(
                () -> BACnetChannelValue.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("BACnetGroupChannelValue");
    // Create the instance
    BACnetGroupChannelValue _bACnetGroupChannelValue;
    _bACnetGroupChannelValue = new BACnetGroupChannelValue(channel, overridingPriority, value);
    return _bACnetGroupChannelValue;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetGroupChannelValue)) {
      return false;
    }
    BACnetGroupChannelValue that = (BACnetGroupChannelValue) o;
    return (getChannel() == that.getChannel())
        && (getOverridingPriority() == that.getOverridingPriority())
        && (getValue() == that.getValue())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getChannel(), getOverridingPriority(), getValue());
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
