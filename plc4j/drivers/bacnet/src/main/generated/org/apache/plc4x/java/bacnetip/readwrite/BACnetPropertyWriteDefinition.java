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

public class BACnetPropertyWriteDefinition implements Message {

  // Properties.
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger arrayIndex;
  protected final BACnetConstructedData propertyValue;
  protected final BACnetContextTagUnsignedInteger priority;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetPropertyWriteDefinition(
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger arrayIndex,
      BACnetConstructedData propertyValue,
      BACnetContextTagUnsignedInteger priority,
      BACnetObjectType objectTypeArgument) {
    super();
    this.propertyIdentifier = propertyIdentifier;
    this.arrayIndex = arrayIndex;
    this.propertyValue = propertyValue;
    this.priority = priority;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getArrayIndex() {
    return arrayIndex;
  }

  public BACnetConstructedData getPropertyValue() {
    return propertyValue;
  }

  public BACnetContextTagUnsignedInteger getPriority() {
    return priority;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPropertyWriteDefinition");

    // Simple Field (propertyIdentifier)
    writeSimpleField(
        "propertyIdentifier", propertyIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (arrayIndex) (Can be skipped, if the value is null)
    writeOptionalField("arrayIndex", arrayIndex, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (propertyValue) (Can be skipped, if the value is null)
    writeOptionalField("propertyValue", propertyValue, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (priority) (Can be skipped, if the value is null)
    writeOptionalField("priority", priority, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyWriteDefinition");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetPropertyWriteDefinition _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (arrayIndex)
    if (arrayIndex != null) {
      lengthInBits += arrayIndex.getLengthInBits();
    }

    // Optional Field (propertyValue)
    if (propertyValue != null) {
      lengthInBits += propertyValue.getLengthInBits();
    }

    // Optional Field (priority)
    if (priority != null) {
      lengthInBits += priority.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetPropertyWriteDefinition staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    BACnetObjectType objectTypeArgument;
    if (args[0] instanceof BACnetObjectType) {
      objectTypeArgument = (BACnetObjectType) args[0];
    } else if (args[0] instanceof String) {
      objectTypeArgument = BACnetObjectType.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type BACnetObjectType or a string which is parseable but"
              + " was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, objectTypeArgument);
  }

  public static BACnetPropertyWriteDefinition staticParse(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetPropertyWriteDefinition");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetPropertyIdentifierTagged propertyIdentifier =
        readSimpleField(
            "propertyIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPropertyIdentifierTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger arrayIndex =
        readOptionalField(
            "arrayIndex",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData propertyValue =
        readOptionalField(
            "propertyValue",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (2),
                        (BACnetObjectType) (objectTypeArgument),
                        (BACnetPropertyIdentifier) (propertyIdentifier.getValue()),
                        (BACnetTagPayloadUnsignedInteger)
                            (((((arrayIndex) != (null)) ? arrayIndex.getPayload() : null)))),
                readBuffer));

    BACnetContextTagUnsignedInteger priority =
        readOptionalField(
            "priority",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyWriteDefinition");
    // Create the instance
    BACnetPropertyWriteDefinition _bACnetPropertyWriteDefinition;
    _bACnetPropertyWriteDefinition =
        new BACnetPropertyWriteDefinition(
            propertyIdentifier, arrayIndex, propertyValue, priority, objectTypeArgument);
    return _bACnetPropertyWriteDefinition;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyWriteDefinition)) {
      return false;
    }
    BACnetPropertyWriteDefinition that = (BACnetPropertyWriteDefinition) o;
    return (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getArrayIndex() == that.getArrayIndex())
        && (getPropertyValue() == that.getPropertyValue())
        && (getPriority() == that.getPriority())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getPropertyIdentifier(), getArrayIndex(), getPropertyValue(), getPriority());
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
