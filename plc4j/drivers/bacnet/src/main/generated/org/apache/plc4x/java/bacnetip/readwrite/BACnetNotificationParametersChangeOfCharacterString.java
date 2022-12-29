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

public class BACnetNotificationParametersChangeOfCharacterString
    extends BACnetNotificationParameters implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetContextTagCharacterString changedValue;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetContextTagCharacterString alarmValue;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersChangeOfCharacterString(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetContextTagCharacterString changedValue,
      BACnetStatusFlagsTagged statusFlags,
      BACnetContextTagCharacterString alarmValue,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.changedValue = changedValue;
    this.statusFlags = statusFlags;
    this.alarmValue = alarmValue;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetContextTagCharacterString getChangedValue() {
    return changedValue;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetContextTagCharacterString getAlarmValue() {
    return alarmValue;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfCharacterString");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (changedValue)
    writeSimpleField("changedValue", changedValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (alarmValue)
    writeSimpleField("alarmValue", alarmValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfCharacterString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfCharacterString _value = this;

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (changedValue)
    lengthInBits += changedValue.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (alarmValue)
    lengthInBits += alarmValue.getLengthInBits();

    // Simple field (innerClosingTag)
    lengthInBits += innerClosingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfCharacterStringBuilder staticParseBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      Short peekedTagNumber)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfCharacterString");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetContextTagCharacterString changedValue =
        readSimpleField(
            "changedValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagCharacterString alarmValue =
        readSimpleField(
            "alarmValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfCharacterString");
    // Create the instance
    return new BACnetNotificationParametersChangeOfCharacterStringBuilder(
        innerOpeningTag,
        changedValue,
        statusFlags,
        alarmValue,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersChangeOfCharacterStringBuilder
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetContextTagCharacterString changedValue;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetContextTagCharacterString alarmValue;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersChangeOfCharacterStringBuilder(
        BACnetOpeningTag innerOpeningTag,
        BACnetContextTagCharacterString changedValue,
        BACnetStatusFlagsTagged statusFlags,
        BACnetContextTagCharacterString alarmValue,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {

      this.innerOpeningTag = innerOpeningTag;
      this.changedValue = changedValue;
      this.statusFlags = statusFlags;
      this.alarmValue = alarmValue;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersChangeOfCharacterString build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersChangeOfCharacterString
          bACnetNotificationParametersChangeOfCharacterString =
              new BACnetNotificationParametersChangeOfCharacterString(
                  openingTag,
                  peekedTagHeader,
                  closingTag,
                  innerOpeningTag,
                  changedValue,
                  statusFlags,
                  alarmValue,
                  innerClosingTag,
                  tagNumber,
                  objectTypeArgument);
      return bACnetNotificationParametersChangeOfCharacterString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfCharacterString)) {
      return false;
    }
    BACnetNotificationParametersChangeOfCharacterString that =
        (BACnetNotificationParametersChangeOfCharacterString) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getChangedValue() == that.getChangedValue())
        && (getStatusFlags() == that.getStatusFlags())
        && (getAlarmValue() == that.getAlarmValue())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getChangedValue(),
        getStatusFlags(),
        getAlarmValue(),
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
