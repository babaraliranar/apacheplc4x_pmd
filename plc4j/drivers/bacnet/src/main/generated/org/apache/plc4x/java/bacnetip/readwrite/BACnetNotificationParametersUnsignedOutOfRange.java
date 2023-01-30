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

public class BACnetNotificationParametersUnsignedOutOfRange extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetContextTagUnsignedInteger exceedingValue;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetContextTagUnsignedInteger deadband;
  protected final BACnetContextTagUnsignedInteger exceededLimit;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersUnsignedOutOfRange(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetContextTagUnsignedInteger exceedingValue,
      BACnetStatusFlagsTagged statusFlags,
      BACnetContextTagUnsignedInteger deadband,
      BACnetContextTagUnsignedInteger exceededLimit,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.exceedingValue = exceedingValue;
    this.statusFlags = statusFlags;
    this.deadband = deadband;
    this.exceededLimit = exceededLimit;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetContextTagUnsignedInteger getExceedingValue() {
    return exceedingValue;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetContextTagUnsignedInteger getDeadband() {
    return deadband;
  }

  public BACnetContextTagUnsignedInteger getExceededLimit() {
    return exceededLimit;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersUnsignedOutOfRange");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (exceedingValue)
    writeSimpleField("exceedingValue", exceedingValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (deadband)
    writeSimpleField("deadband", deadband, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (exceededLimit)
    writeSimpleField("exceededLimit", exceededLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersUnsignedOutOfRange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersUnsignedOutOfRange _value = this;

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (exceedingValue)
    lengthInBits += exceedingValue.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (deadband)
    lengthInBits += deadband.getLengthInBits();

    // Simple field (exceededLimit)
    lengthInBits += exceededLimit.getLengthInBits();

    // Simple field (innerClosingTag)
    lengthInBits += innerClosingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersBuilder staticParseBACnetNotificationParametersBuilder(
      ReadBuffer readBuffer,
      Short peekedTagNumber,
      Short tagNumber,
      BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersUnsignedOutOfRange");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetContextTagUnsignedInteger exceedingValue =
        readSimpleField(
            "exceedingValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger deadband =
        readSimpleField(
            "deadband",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger exceededLimit =
        readSimpleField(
            "exceededLimit",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersUnsignedOutOfRange");
    // Create the instance
    return new BACnetNotificationParametersUnsignedOutOfRangeBuilderImpl(
        innerOpeningTag,
        exceedingValue,
        statusFlags,
        deadband,
        exceededLimit,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersUnsignedOutOfRangeBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetContextTagUnsignedInteger exceedingValue;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetContextTagUnsignedInteger deadband;
    private final BACnetContextTagUnsignedInteger exceededLimit;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersUnsignedOutOfRangeBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetContextTagUnsignedInteger exceedingValue,
        BACnetStatusFlagsTagged statusFlags,
        BACnetContextTagUnsignedInteger deadband,
        BACnetContextTagUnsignedInteger exceededLimit,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.exceedingValue = exceedingValue;
      this.statusFlags = statusFlags;
      this.deadband = deadband;
      this.exceededLimit = exceededLimit;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersUnsignedOutOfRange build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersUnsignedOutOfRange
          bACnetNotificationParametersUnsignedOutOfRange =
              new BACnetNotificationParametersUnsignedOutOfRange(
                  openingTag,
                  peekedTagHeader,
                  closingTag,
                  innerOpeningTag,
                  exceedingValue,
                  statusFlags,
                  deadband,
                  exceededLimit,
                  innerClosingTag,
                  tagNumber,
                  objectTypeArgument);
      return bACnetNotificationParametersUnsignedOutOfRange;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersUnsignedOutOfRange)) {
      return false;
    }
    BACnetNotificationParametersUnsignedOutOfRange that =
        (BACnetNotificationParametersUnsignedOutOfRange) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getExceedingValue() == that.getExceedingValue())
        && (getStatusFlags() == that.getStatusFlags())
        && (getDeadband() == that.getDeadband())
        && (getExceededLimit() == that.getExceededLimit())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getExceedingValue(),
        getStatusFlags(),
        getDeadband(),
        getExceededLimit(),
        getInnerClosingTag());
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
