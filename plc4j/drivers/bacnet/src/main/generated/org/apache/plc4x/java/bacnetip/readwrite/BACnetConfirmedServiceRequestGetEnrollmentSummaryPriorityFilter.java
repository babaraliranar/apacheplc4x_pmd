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

public class BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter implements Message {

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagUnsignedInteger minPriority;
  protected final BACnetContextTagUnsignedInteger maxPriority;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;

  public BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(
      BACnetOpeningTag openingTag,
      BACnetContextTagUnsignedInteger minPriority,
      BACnetContextTagUnsignedInteger maxPriority,
      BACnetClosingTag closingTag,
      Short tagNumber) {
    super();
    this.openingTag = openingTag;
    this.minPriority = minPriority;
    this.maxPriority = maxPriority;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagUnsignedInteger getMinPriority() {
    return minPriority;
  }

  public BACnetContextTagUnsignedInteger getMaxPriority() {
    return maxPriority;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (minPriority)
    writeSimpleField("minPriority", minPriority, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (maxPriority)
    writeSimpleField("maxPriority", maxPriority, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (minPriority)
    lengthInBits += minPriority.getLengthInBits();

    // Simple field (maxPriority)
    lengthInBits += maxPriority.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter staticParse(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    BACnetContextTagUnsignedInteger minPriority =
        readSimpleField(
            "minPriority",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger maxPriority =
        readSimpleField(
            "maxPriority",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter");
    // Create the instance
    BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
        _bACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter;
    _bACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter =
        new BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(
            openingTag, minPriority, maxPriority, closingTag, tagNumber);
    return _bACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter)) {
      return false;
    }
    BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter that =
        (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getMinPriority() == that.getMinPriority())
        && (getMaxPriority() == that.getMaxPriority())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOpeningTag(), getMinPriority(), getMaxPriority(), getClosingTag());
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
