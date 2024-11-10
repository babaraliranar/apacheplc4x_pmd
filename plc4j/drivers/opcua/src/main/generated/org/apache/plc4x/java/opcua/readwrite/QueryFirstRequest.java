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

public class QueryFirstRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 615;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final ViewDescription view;
  protected final List<NodeTypeDescription> nodeTypes;
  protected final ContentFilter filter;
  protected final long maxDataSetsToReturn;
  protected final long maxReferencesToReturn;

  public QueryFirstRequest(
      RequestHeader requestHeader,
      ViewDescription view,
      List<NodeTypeDescription> nodeTypes,
      ContentFilter filter,
      long maxDataSetsToReturn,
      long maxReferencesToReturn) {
    super();
    this.requestHeader = requestHeader;
    this.view = view;
    this.nodeTypes = nodeTypes;
    this.filter = filter;
    this.maxDataSetsToReturn = maxDataSetsToReturn;
    this.maxReferencesToReturn = maxReferencesToReturn;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public ViewDescription getView() {
    return view;
  }

  public List<NodeTypeDescription> getNodeTypes() {
    return nodeTypes;
  }

  public ContentFilter getFilter() {
    return filter;
  }

  public long getMaxDataSetsToReturn() {
    return maxDataSetsToReturn;
  }

  public long getMaxReferencesToReturn() {
    return maxReferencesToReturn;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("QueryFirstRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Simple Field (view)
    writeSimpleField("view", view, writeComplex(writeBuffer));

    // Implicit Field (noOfNodeTypes) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfNodeTypes = (int) ((((getNodeTypes()) == (null)) ? -(1) : COUNT(getNodeTypes())));
    writeImplicitField("noOfNodeTypes", noOfNodeTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (nodeTypes)
    writeComplexTypeArrayField("nodeTypes", nodeTypes, writeBuffer);

    // Simple Field (filter)
    writeSimpleField("filter", filter, writeComplex(writeBuffer));

    // Simple Field (maxDataSetsToReturn)
    writeSimpleField(
        "maxDataSetsToReturn", maxDataSetsToReturn, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxReferencesToReturn)
    writeSimpleField(
        "maxReferencesToReturn", maxReferencesToReturn, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("QueryFirstRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    QueryFirstRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (view)
    lengthInBits += view.getLengthInBits();

    // Implicit Field (noOfNodeTypes)
    lengthInBits += 32;

    // Array field
    if (nodeTypes != null) {
      int i = 0;
      for (NodeTypeDescription element : nodeTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= nodeTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (filter)
    lengthInBits += filter.getLengthInBits();

    // Simple field (maxDataSetsToReturn)
    lengthInBits += 32;

    // Simple field (maxReferencesToReturn)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("QueryFirstRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    ViewDescription view =
        readSimpleField(
            "view",
            readComplex(
                () ->
                    (ViewDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (513)),
                readBuffer));

    int noOfNodeTypes = readImplicitField("noOfNodeTypes", readSignedInt(readBuffer, 32));

    List<NodeTypeDescription> nodeTypes =
        readCountArrayField(
            "nodeTypes",
            readComplex(
                () ->
                    (NodeTypeDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (575)),
                readBuffer),
            noOfNodeTypes);

    ContentFilter filter =
        readSimpleField(
            "filter",
            readComplex(
                () ->
                    (ContentFilter) ExtensionObjectDefinition.staticParse(readBuffer, (int) (588)),
                readBuffer));

    long maxDataSetsToReturn =
        readSimpleField("maxDataSetsToReturn", readUnsignedLong(readBuffer, 32));

    long maxReferencesToReturn =
        readSimpleField("maxReferencesToReturn", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("QueryFirstRequest");
    // Create the instance
    return new QueryFirstRequestBuilderImpl(
        requestHeader, view, nodeTypes, filter, maxDataSetsToReturn, maxReferencesToReturn);
  }

  public static class QueryFirstRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final ViewDescription view;
    private final List<NodeTypeDescription> nodeTypes;
    private final ContentFilter filter;
    private final long maxDataSetsToReturn;
    private final long maxReferencesToReturn;

    public QueryFirstRequestBuilderImpl(
        RequestHeader requestHeader,
        ViewDescription view,
        List<NodeTypeDescription> nodeTypes,
        ContentFilter filter,
        long maxDataSetsToReturn,
        long maxReferencesToReturn) {
      this.requestHeader = requestHeader;
      this.view = view;
      this.nodeTypes = nodeTypes;
      this.filter = filter;
      this.maxDataSetsToReturn = maxDataSetsToReturn;
      this.maxReferencesToReturn = maxReferencesToReturn;
    }

    public QueryFirstRequest build() {
      QueryFirstRequest queryFirstRequest =
          new QueryFirstRequest(
              requestHeader, view, nodeTypes, filter, maxDataSetsToReturn, maxReferencesToReturn);
      return queryFirstRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QueryFirstRequest)) {
      return false;
    }
    QueryFirstRequest that = (QueryFirstRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getView() == that.getView())
        && (getNodeTypes() == that.getNodeTypes())
        && (getFilter() == that.getFilter())
        && (getMaxDataSetsToReturn() == that.getMaxDataSetsToReturn())
        && (getMaxReferencesToReturn() == that.getMaxReferencesToReturn())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getView(),
        getNodeTypes(),
        getFilter(),
        getMaxDataSetsToReturn(),
        getMaxReferencesToReturn());
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
