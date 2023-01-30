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

public class BACnetNotificationParametersChangeOfValue extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetNotificationParametersChangeOfValueNewValue newValue;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersChangeOfValue(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetNotificationParametersChangeOfValueNewValue newValue,
      BACnetStatusFlagsTagged statusFlags,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.newValue = newValue;
    this.statusFlags = statusFlags;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetNotificationParametersChangeOfValueNewValue getNewValue() {
    return newValue;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfValue");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (newValue)
    writeSimpleField("newValue", newValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfValue _value = this;

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (newValue)
    lengthInBits += newValue.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

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
    readBuffer.pullContext("BACnetNotificationParametersChangeOfValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetNotificationParametersChangeOfValueNewValue newValue =
        readSimpleField(
            "newValue",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetNotificationParametersChangeOfValueNewValue.staticParse(
                        readBuffer, (short) (0)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfValue");
    // Create the instance
    return new BACnetNotificationParametersChangeOfValueBuilderImpl(
        innerOpeningTag, newValue, statusFlags, innerClosingTag, tagNumber, objectTypeArgument);
  }

  public static class BACnetNotificationParametersChangeOfValueBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetNotificationParametersChangeOfValueNewValue newValue;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersChangeOfValueBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetNotificationParametersChangeOfValueNewValue newValue,
        BACnetStatusFlagsTagged statusFlags,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.newValue = newValue;
      this.statusFlags = statusFlags;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersChangeOfValue build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersChangeOfValue bACnetNotificationParametersChangeOfValue =
          new BACnetNotificationParametersChangeOfValue(
              openingTag,
              peekedTagHeader,
              closingTag,
              innerOpeningTag,
              newValue,
              statusFlags,
              innerClosingTag,
              tagNumber,
              objectTypeArgument);
      return bACnetNotificationParametersChangeOfValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfValue)) {
      return false;
    }
    BACnetNotificationParametersChangeOfValue that = (BACnetNotificationParametersChangeOfValue) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getNewValue() == that.getNewValue())
        && (getStatusFlags() == that.getStatusFlags())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getNewValue(),
        getStatusFlags(),
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
