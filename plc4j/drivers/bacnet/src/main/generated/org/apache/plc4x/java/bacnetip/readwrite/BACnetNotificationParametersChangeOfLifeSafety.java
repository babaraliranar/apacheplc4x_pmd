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

public class BACnetNotificationParametersChangeOfLifeSafety extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetLifeSafetyStateTagged newState;
  protected final BACnetLifeSafetyModeTagged newMode;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetLifeSafetyOperationTagged operationExpected;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersChangeOfLifeSafety(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetLifeSafetyStateTagged newState,
      BACnetLifeSafetyModeTagged newMode,
      BACnetStatusFlagsTagged statusFlags,
      BACnetLifeSafetyOperationTagged operationExpected,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.newState = newState;
    this.newMode = newMode;
    this.statusFlags = statusFlags;
    this.operationExpected = operationExpected;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetLifeSafetyStateTagged getNewState() {
    return newState;
  }

  public BACnetLifeSafetyModeTagged getNewMode() {
    return newMode;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetLifeSafetyOperationTagged getOperationExpected() {
    return operationExpected;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfLifeSafety");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (newState)
    writeSimpleField("newState", newState, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (newMode)
    writeSimpleField("newMode", newMode, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (operationExpected)
    writeSimpleField(
        "operationExpected", operationExpected, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfLifeSafety");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfLifeSafety _value = this;

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (newState)
    lengthInBits += newState.getLengthInBits();

    // Simple field (newMode)
    lengthInBits += newMode.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (operationExpected)
    lengthInBits += operationExpected.getLengthInBits();

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
    readBuffer.pullContext("BACnetNotificationParametersChangeOfLifeSafety");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetLifeSafetyStateTagged newState =
        readSimpleField(
            "newState",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetLifeSafetyStateTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetLifeSafetyModeTagged newMode =
        readSimpleField(
            "newMode",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetLifeSafetyModeTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (2), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetLifeSafetyOperationTagged operationExpected =
        readSimpleField(
            "operationExpected",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetLifeSafetyOperationTagged.staticParse(
                        readBuffer, (short) (3), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfLifeSafety");
    // Create the instance
    return new BACnetNotificationParametersChangeOfLifeSafetyBuilderImpl(
        innerOpeningTag,
        newState,
        newMode,
        statusFlags,
        operationExpected,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersChangeOfLifeSafetyBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetLifeSafetyStateTagged newState;
    private final BACnetLifeSafetyModeTagged newMode;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetLifeSafetyOperationTagged operationExpected;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersChangeOfLifeSafetyBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetLifeSafetyStateTagged newState,
        BACnetLifeSafetyModeTagged newMode,
        BACnetStatusFlagsTagged statusFlags,
        BACnetLifeSafetyOperationTagged operationExpected,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.newState = newState;
      this.newMode = newMode;
      this.statusFlags = statusFlags;
      this.operationExpected = operationExpected;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersChangeOfLifeSafety build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersChangeOfLifeSafety
          bACnetNotificationParametersChangeOfLifeSafety =
              new BACnetNotificationParametersChangeOfLifeSafety(
                  openingTag,
                  peekedTagHeader,
                  closingTag,
                  innerOpeningTag,
                  newState,
                  newMode,
                  statusFlags,
                  operationExpected,
                  innerClosingTag,
                  tagNumber,
                  objectTypeArgument);
      return bACnetNotificationParametersChangeOfLifeSafety;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfLifeSafety)) {
      return false;
    }
    BACnetNotificationParametersChangeOfLifeSafety that =
        (BACnetNotificationParametersChangeOfLifeSafety) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getNewState() == that.getNewState())
        && (getNewMode() == that.getNewMode())
        && (getStatusFlags() == that.getStatusFlags())
        && (getOperationExpected() == that.getOperationExpected())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getNewState(),
        getNewMode(),
        getStatusFlags(),
        getOperationExpected(),
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
