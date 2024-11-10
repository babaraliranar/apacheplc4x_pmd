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
package org.apache.plc4x.java.opcua.readwrite;

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

public class PublishedEventsDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15584;
  }

  // Properties.
  protected final NodeId eventNotifier;
  protected final List<SimpleAttributeOperand> selectedFields;
  protected final ContentFilter filter;

  public PublishedEventsDataType(
      NodeId eventNotifier, List<SimpleAttributeOperand> selectedFields, ContentFilter filter) {
    super();
    this.eventNotifier = eventNotifier;
    this.selectedFields = selectedFields;
    this.filter = filter;
  }

  public NodeId getEventNotifier() {
    return eventNotifier;
  }

  public List<SimpleAttributeOperand> getSelectedFields() {
    return selectedFields;
  }

  public ContentFilter getFilter() {
    return filter;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PublishedEventsDataType");

    // Simple Field (eventNotifier)
    writeSimpleField("eventNotifier", eventNotifier, writeComplex(writeBuffer));

    // Implicit Field (noOfSelectedFields) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfSelectedFields =
        (int) ((((getSelectedFields()) == (null)) ? -(1) : COUNT(getSelectedFields())));
    writeImplicitField("noOfSelectedFields", noOfSelectedFields, writeSignedInt(writeBuffer, 32));

    // Array Field (selectedFields)
    writeComplexTypeArrayField("selectedFields", selectedFields, writeBuffer);

    // Simple Field (filter)
    writeSimpleField("filter", filter, writeComplex(writeBuffer));

    writeBuffer.popContext("PublishedEventsDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PublishedEventsDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (eventNotifier)
    lengthInBits += eventNotifier.getLengthInBits();

    // Implicit Field (noOfSelectedFields)
    lengthInBits += 32;

    // Array field
    if (selectedFields != null) {
      int i = 0;
      for (SimpleAttributeOperand element : selectedFields) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= selectedFields.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (filter)
    lengthInBits += filter.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("PublishedEventsDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId eventNotifier =
        readSimpleField(
            "eventNotifier", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    int noOfSelectedFields = readImplicitField("noOfSelectedFields", readSignedInt(readBuffer, 32));

    List<SimpleAttributeOperand> selectedFields =
        readCountArrayField(
            "selectedFields",
            readComplex(
                () ->
                    (SimpleAttributeOperand)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (603)),
                readBuffer),
            noOfSelectedFields);

    ContentFilter filter =
        readSimpleField(
            "filter",
            readComplex(
                () ->
                    (ContentFilter) ExtensionObjectDefinition.staticParse(readBuffer, (int) (588)),
                readBuffer));

    readBuffer.closeContext("PublishedEventsDataType");
    // Create the instance
    return new PublishedEventsDataTypeBuilderImpl(eventNotifier, selectedFields, filter);
  }

  public static class PublishedEventsDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId eventNotifier;
    private final List<SimpleAttributeOperand> selectedFields;
    private final ContentFilter filter;

    public PublishedEventsDataTypeBuilderImpl(
        NodeId eventNotifier, List<SimpleAttributeOperand> selectedFields, ContentFilter filter) {
      this.eventNotifier = eventNotifier;
      this.selectedFields = selectedFields;
      this.filter = filter;
    }

    public PublishedEventsDataType build() {
      PublishedEventsDataType publishedEventsDataType =
          new PublishedEventsDataType(eventNotifier, selectedFields, filter);
      return publishedEventsDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PublishedEventsDataType)) {
      return false;
    }
    PublishedEventsDataType that = (PublishedEventsDataType) o;
    return (getEventNotifier() == that.getEventNotifier())
        && (getSelectedFields() == that.getSelectedFields())
        && (getFilter() == that.getFilter())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getEventNotifier(), getSelectedFields(), getFilter());
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
