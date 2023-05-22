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

public class BACnetEventParameterFloatingLimit extends BACnetEventParameter implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagUnsignedInteger timeDelay;
  protected final BACnetDeviceObjectPropertyReferenceEnclosed setpointReference;
  protected final BACnetContextTagReal lowDiffLimit;
  protected final BACnetContextTagReal highDiffLimit;
  protected final BACnetContextTagReal deadband;
  protected final BACnetClosingTag closingTag;

  public BACnetEventParameterFloatingLimit(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetContextTagUnsignedInteger timeDelay,
      BACnetDeviceObjectPropertyReferenceEnclosed setpointReference,
      BACnetContextTagReal lowDiffLimit,
      BACnetContextTagReal highDiffLimit,
      BACnetContextTagReal deadband,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.timeDelay = timeDelay;
    this.setpointReference = setpointReference;
    this.lowDiffLimit = lowDiffLimit;
    this.highDiffLimit = highDiffLimit;
    this.deadband = deadband;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagUnsignedInteger getTimeDelay() {
    return timeDelay;
  }

  public BACnetDeviceObjectPropertyReferenceEnclosed getSetpointReference() {
    return setpointReference;
  }

  public BACnetContextTagReal getLowDiffLimit() {
    return lowDiffLimit;
  }

  public BACnetContextTagReal getHighDiffLimit() {
    return highDiffLimit;
  }

  public BACnetContextTagReal getDeadband() {
    return deadband;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetEventParameterFloatingLimit");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (setpointReference)
    writeSimpleField(
        "setpointReference", setpointReference, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (lowDiffLimit)
    writeSimpleField("lowDiffLimit", lowDiffLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (highDiffLimit)
    writeSimpleField("highDiffLimit", highDiffLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (deadband)
    writeSimpleField("deadband", deadband, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterFloatingLimit");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterFloatingLimit _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (setpointReference)
    lengthInBits += setpointReference.getLengthInBits();

    // Simple field (lowDiffLimit)
    lengthInBits += lowDiffLimit.getLengthInBits();

    // Simple field (highDiffLimit)
    lengthInBits += highDiffLimit.getLengthInBits();

    // Simple field (deadband)
    lengthInBits += deadband.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterBuilder staticParseBACnetEventParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterFloatingLimit");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (4)), readBuffer));

    BACnetContextTagUnsignedInteger timeDelay =
        readSimpleField(
            "timeDelay",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetDeviceObjectPropertyReferenceEnclosed setpointReference =
        readSimpleField(
            "setpointReference",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetDeviceObjectPropertyReferenceEnclosed.staticParse(
                        readBuffer, (short) (1)),
                readBuffer));

    BACnetContextTagReal lowDiffLimit =
        readSimpleField(
            "lowDiffLimit",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (2), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagReal highDiffLimit =
        readSimpleField(
            "highDiffLimit",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (3), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagReal deadband =
        readSimpleField(
            "deadband",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (4), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (4)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterFloatingLimit");
    // Create the instance
    return new BACnetEventParameterFloatingLimitBuilderImpl(
        openingTag,
        timeDelay,
        setpointReference,
        lowDiffLimit,
        highDiffLimit,
        deadband,
        closingTag);
  }

  public static class BACnetEventParameterFloatingLimitBuilderImpl
      implements BACnetEventParameter.BACnetEventParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetContextTagUnsignedInteger timeDelay;
    private final BACnetDeviceObjectPropertyReferenceEnclosed setpointReference;
    private final BACnetContextTagReal lowDiffLimit;
    private final BACnetContextTagReal highDiffLimit;
    private final BACnetContextTagReal deadband;
    private final BACnetClosingTag closingTag;

    public BACnetEventParameterFloatingLimitBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetContextTagUnsignedInteger timeDelay,
        BACnetDeviceObjectPropertyReferenceEnclosed setpointReference,
        BACnetContextTagReal lowDiffLimit,
        BACnetContextTagReal highDiffLimit,
        BACnetContextTagReal deadband,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.timeDelay = timeDelay;
      this.setpointReference = setpointReference;
      this.lowDiffLimit = lowDiffLimit;
      this.highDiffLimit = highDiffLimit;
      this.deadband = deadband;
      this.closingTag = closingTag;
    }

    public BACnetEventParameterFloatingLimit build(BACnetTagHeader peekedTagHeader) {
      BACnetEventParameterFloatingLimit bACnetEventParameterFloatingLimit =
          new BACnetEventParameterFloatingLimit(
              peekedTagHeader,
              openingTag,
              timeDelay,
              setpointReference,
              lowDiffLimit,
              highDiffLimit,
              deadband,
              closingTag);
      return bACnetEventParameterFloatingLimit;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterFloatingLimit)) {
      return false;
    }
    BACnetEventParameterFloatingLimit that = (BACnetEventParameterFloatingLimit) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getTimeDelay() == that.getTimeDelay())
        && (getSetpointReference() == that.getSetpointReference())
        && (getLowDiffLimit() == that.getLowDiffLimit())
        && (getHighDiffLimit() == that.getHighDiffLimit())
        && (getDeadband() == that.getDeadband())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getOpeningTag(),
        getTimeDelay(),
        getSetpointReference(),
        getLowDiffLimit(),
        getHighDiffLimit(),
        getDeadband(),
        getClosingTag());
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
