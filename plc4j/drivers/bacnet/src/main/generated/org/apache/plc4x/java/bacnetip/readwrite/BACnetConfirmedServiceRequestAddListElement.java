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

public class BACnetConfirmedServiceRequestAddListElement extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.ADD_LIST_ELEMENT;
  }

  // Properties.
  protected final BACnetContextTagObjectIdentifier objectIdentifier;
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger arrayIndex;
  protected final BACnetConstructedData listOfElements;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestAddListElement(
      BACnetContextTagObjectIdentifier objectIdentifier,
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger arrayIndex,
      BACnetConstructedData listOfElements,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.objectIdentifier = objectIdentifier;
    this.propertyIdentifier = propertyIdentifier;
    this.arrayIndex = arrayIndex;
    this.listOfElements = listOfElements;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getArrayIndex() {
    return arrayIndex;
  }

  public BACnetConstructedData getListOfElements() {
    return listOfElements;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestAddListElement");

    // Simple Field (objectIdentifier)
    writeSimpleField("objectIdentifier", objectIdentifier, writeComplex(writeBuffer));

    // Simple Field (propertyIdentifier)
    writeSimpleField("propertyIdentifier", propertyIdentifier, writeComplex(writeBuffer));

    // Optional Field (arrayIndex) (Can be skipped, if the value is null)
    writeOptionalField("arrayIndex", arrayIndex, writeComplex(writeBuffer));

    // Optional Field (listOfElements) (Can be skipped, if the value is null)
    writeOptionalField("listOfElements", listOfElements, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestAddListElement");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestAddListElement _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (arrayIndex)
    if (arrayIndex != null) {
      lengthInBits += arrayIndex.getLengthInBits();
    }

    // Optional Field (listOfElements)
    if (listOfElements != null) {
      lengthInBits += listOfElements.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestAddListElement");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetPropertyIdentifierTagged propertyIdentifier =
        readSimpleField(
            "propertyIdentifier",
            readComplex(
                () ->
                    BACnetPropertyIdentifierTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger arrayIndex =
        readOptionalField(
            "arrayIndex",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData listOfElements =
        readOptionalField(
            "listOfElements",
            readComplex(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (3),
                        (BACnetObjectType) (objectIdentifier.getObjectType()),
                        (BACnetPropertyIdentifier) (propertyIdentifier.getValue()),
                        (BACnetTagPayloadUnsignedInteger)
                            (((((arrayIndex) != (null)) ? arrayIndex.getPayload() : null)))),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestAddListElement");
    // Create the instance
    return new BACnetConfirmedServiceRequestAddListElementBuilderImpl(
        objectIdentifier, propertyIdentifier, arrayIndex, listOfElements, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestAddListElementBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagObjectIdentifier objectIdentifier;
    private final BACnetPropertyIdentifierTagged propertyIdentifier;
    private final BACnetContextTagUnsignedInteger arrayIndex;
    private final BACnetConstructedData listOfElements;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestAddListElementBuilderImpl(
        BACnetContextTagObjectIdentifier objectIdentifier,
        BACnetPropertyIdentifierTagged propertyIdentifier,
        BACnetContextTagUnsignedInteger arrayIndex,
        BACnetConstructedData listOfElements,
        Long serviceRequestLength) {
      this.objectIdentifier = objectIdentifier;
      this.propertyIdentifier = propertyIdentifier;
      this.arrayIndex = arrayIndex;
      this.listOfElements = listOfElements;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestAddListElement build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestAddListElement bACnetConfirmedServiceRequestAddListElement =
          new BACnetConfirmedServiceRequestAddListElement(
              objectIdentifier,
              propertyIdentifier,
              arrayIndex,
              listOfElements,
              serviceRequestLength);
      return bACnetConfirmedServiceRequestAddListElement;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestAddListElement)) {
      return false;
    }
    BACnetConfirmedServiceRequestAddListElement that =
        (BACnetConfirmedServiceRequestAddListElement) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getArrayIndex() == that.getArrayIndex())
        && (getListOfElements() == that.getListOfElements())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getObjectIdentifier(),
        getPropertyIdentifier(),
        getArrayIndex(),
        getListOfElements());
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
