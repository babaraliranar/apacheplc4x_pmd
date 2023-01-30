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

public class BACnetConstructedDataTrendLogLogDeviceObjectProperty extends BACnetConstructedData
    implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return BACnetObjectType.TREND_LOG;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.LOG_DEVICE_OBJECT_PROPERTY;
  }

  // Properties.
  protected final BACnetDeviceObjectPropertyReference logDeviceObjectProperty;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataTrendLogLogDeviceObjectProperty(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetDeviceObjectPropertyReference logDeviceObjectProperty,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.logDeviceObjectProperty = logDeviceObjectProperty;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetDeviceObjectPropertyReference getLogDeviceObjectProperty() {
    return logDeviceObjectProperty;
  }

  public BACnetDeviceObjectPropertyReference getActualValue() {
    return (BACnetDeviceObjectPropertyReference) (getLogDeviceObjectProperty());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty");

    // Simple Field (logDeviceObjectProperty)
    writeSimpleField(
        "logDeviceObjectProperty",
        logDeviceObjectProperty,
        new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetDeviceObjectPropertyReference actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataTrendLogLogDeviceObjectProperty _value = this;

    // Simple field (logDeviceObjectProperty)
    lengthInBits += logDeviceObjectProperty.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetDeviceObjectPropertyReference logDeviceObjectProperty =
        readSimpleField(
            "logDeviceObjectProperty",
            new DataReaderComplexDefault<>(
                () -> BACnetDeviceObjectPropertyReference.staticParse(readBuffer), readBuffer));
    BACnetDeviceObjectPropertyReference actualValue =
        readVirtualField(
            "actualValue", BACnetDeviceObjectPropertyReference.class, logDeviceObjectProperty);

    readBuffer.closeContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty");
    // Create the instance
    return new BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilderImpl(
        logDeviceObjectProperty, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetDeviceObjectPropertyReference logDeviceObjectProperty;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilderImpl(
        BACnetDeviceObjectPropertyReference logDeviceObjectProperty,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.logDeviceObjectProperty = logDeviceObjectProperty;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataTrendLogLogDeviceObjectProperty build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataTrendLogLogDeviceObjectProperty
          bACnetConstructedDataTrendLogLogDeviceObjectProperty =
              new BACnetConstructedDataTrendLogLogDeviceObjectProperty(
                  openingTag,
                  peekedTagHeader,
                  closingTag,
                  logDeviceObjectProperty,
                  tagNumber,
                  arrayIndexArgument);
      return bACnetConstructedDataTrendLogLogDeviceObjectProperty;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataTrendLogLogDeviceObjectProperty)) {
      return false;
    }
    BACnetConstructedDataTrendLogLogDeviceObjectProperty that =
        (BACnetConstructedDataTrendLogLogDeviceObjectProperty) o;
    return (getLogDeviceObjectProperty() == that.getLogDeviceObjectProperty())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLogDeviceObjectProperty());
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
