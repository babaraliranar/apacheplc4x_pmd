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

public class CallMethodResult extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "709";
  }

  // Properties.
  protected final StatusCode statusCode;
  protected final int noOfInputArgumentResults;
  protected final List<StatusCode> inputArgumentResults;
  protected final int noOfInputArgumentDiagnosticInfos;
  protected final List<DiagnosticInfo> inputArgumentDiagnosticInfos;
  protected final int noOfOutputArguments;
  protected final List<Variant> outputArguments;

  public CallMethodResult(
      StatusCode statusCode,
      int noOfInputArgumentResults,
      List<StatusCode> inputArgumentResults,
      int noOfInputArgumentDiagnosticInfos,
      List<DiagnosticInfo> inputArgumentDiagnosticInfos,
      int noOfOutputArguments,
      List<Variant> outputArguments) {
    super();
    this.statusCode = statusCode;
    this.noOfInputArgumentResults = noOfInputArgumentResults;
    this.inputArgumentResults = inputArgumentResults;
    this.noOfInputArgumentDiagnosticInfos = noOfInputArgumentDiagnosticInfos;
    this.inputArgumentDiagnosticInfos = inputArgumentDiagnosticInfos;
    this.noOfOutputArguments = noOfOutputArguments;
    this.outputArguments = outputArguments;
  }

  public StatusCode getStatusCode() {
    return statusCode;
  }

  public int getNoOfInputArgumentResults() {
    return noOfInputArgumentResults;
  }

  public List<StatusCode> getInputArgumentResults() {
    return inputArgumentResults;
  }

  public int getNoOfInputArgumentDiagnosticInfos() {
    return noOfInputArgumentDiagnosticInfos;
  }

  public List<DiagnosticInfo> getInputArgumentDiagnosticInfos() {
    return inputArgumentDiagnosticInfos;
  }

  public int getNoOfOutputArguments() {
    return noOfOutputArguments;
  }

  public List<Variant> getOutputArguments() {
    return outputArguments;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CallMethodResult");

    // Simple Field (statusCode)
    writeSimpleField("statusCode", statusCode, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfInputArgumentResults)
    writeSimpleField(
        "noOfInputArgumentResults", noOfInputArgumentResults, writeSignedInt(writeBuffer, 32));

    // Array Field (inputArgumentResults)
    writeComplexTypeArrayField("inputArgumentResults", inputArgumentResults, writeBuffer);

    // Simple Field (noOfInputArgumentDiagnosticInfos)
    writeSimpleField(
        "noOfInputArgumentDiagnosticInfos",
        noOfInputArgumentDiagnosticInfos,
        writeSignedInt(writeBuffer, 32));

    // Array Field (inputArgumentDiagnosticInfos)
    writeComplexTypeArrayField(
        "inputArgumentDiagnosticInfos", inputArgumentDiagnosticInfos, writeBuffer);

    // Simple Field (noOfOutputArguments)
    writeSimpleField("noOfOutputArguments", noOfOutputArguments, writeSignedInt(writeBuffer, 32));

    // Array Field (outputArguments)
    writeComplexTypeArrayField("outputArguments", outputArguments, writeBuffer);

    writeBuffer.popContext("CallMethodResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CallMethodResult _value = this;

    // Simple field (statusCode)
    lengthInBits += statusCode.getLengthInBits();

    // Simple field (noOfInputArgumentResults)
    lengthInBits += 32;

    // Array field
    if (inputArgumentResults != null) {
      int i = 0;
      for (StatusCode element : inputArgumentResults) {
        boolean last = ++i >= inputArgumentResults.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfInputArgumentDiagnosticInfos)
    lengthInBits += 32;

    // Array field
    if (inputArgumentDiagnosticInfos != null) {
      int i = 0;
      for (DiagnosticInfo element : inputArgumentDiagnosticInfos) {
        boolean last = ++i >= inputArgumentDiagnosticInfos.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfOutputArguments)
    lengthInBits += 32;

    // Array field
    if (outputArguments != null) {
      int i = 0;
      for (Variant element : outputArguments) {
        boolean last = ++i >= outputArguments.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static CallMethodResultBuilder staticParseBuilder(ReadBuffer readBuffer, String identifier)
      throws ParseException {
    readBuffer.pullContext("CallMethodResult");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    StatusCode statusCode =
        readSimpleField(
            "statusCode",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer));

    int noOfInputArgumentResults =
        readSimpleField("noOfInputArgumentResults", readSignedInt(readBuffer, 32));

    List<StatusCode> inputArgumentResults =
        readCountArrayField(
            "inputArgumentResults",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer),
            noOfInputArgumentResults);

    int noOfInputArgumentDiagnosticInfos =
        readSimpleField("noOfInputArgumentDiagnosticInfos", readSignedInt(readBuffer, 32));

    List<DiagnosticInfo> inputArgumentDiagnosticInfos =
        readCountArrayField(
            "inputArgumentDiagnosticInfos",
            new DataReaderComplexDefault<>(
                () -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            noOfInputArgumentDiagnosticInfos);

    int noOfOutputArguments = readSimpleField("noOfOutputArguments", readSignedInt(readBuffer, 32));

    List<Variant> outputArguments =
        readCountArrayField(
            "outputArguments",
            new DataReaderComplexDefault<>(() -> Variant.staticParse(readBuffer), readBuffer),
            noOfOutputArguments);

    readBuffer.closeContext("CallMethodResult");
    // Create the instance
    return new CallMethodResultBuilder(
        statusCode,
        noOfInputArgumentResults,
        inputArgumentResults,
        noOfInputArgumentDiagnosticInfos,
        inputArgumentDiagnosticInfos,
        noOfOutputArguments,
        outputArguments);
  }

  public static class CallMethodResultBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final StatusCode statusCode;
    private final int noOfInputArgumentResults;
    private final List<StatusCode> inputArgumentResults;
    private final int noOfInputArgumentDiagnosticInfos;
    private final List<DiagnosticInfo> inputArgumentDiagnosticInfos;
    private final int noOfOutputArguments;
    private final List<Variant> outputArguments;

    public CallMethodResultBuilder(
        StatusCode statusCode,
        int noOfInputArgumentResults,
        List<StatusCode> inputArgumentResults,
        int noOfInputArgumentDiagnosticInfos,
        List<DiagnosticInfo> inputArgumentDiagnosticInfos,
        int noOfOutputArguments,
        List<Variant> outputArguments) {

      this.statusCode = statusCode;
      this.noOfInputArgumentResults = noOfInputArgumentResults;
      this.inputArgumentResults = inputArgumentResults;
      this.noOfInputArgumentDiagnosticInfos = noOfInputArgumentDiagnosticInfos;
      this.inputArgumentDiagnosticInfos = inputArgumentDiagnosticInfos;
      this.noOfOutputArguments = noOfOutputArguments;
      this.outputArguments = outputArguments;
    }

    public CallMethodResult build() {
      CallMethodResult callMethodResult =
          new CallMethodResult(
              statusCode,
              noOfInputArgumentResults,
              inputArgumentResults,
              noOfInputArgumentDiagnosticInfos,
              inputArgumentDiagnosticInfos,
              noOfOutputArguments,
              outputArguments);
      return callMethodResult;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CallMethodResult)) {
      return false;
    }
    CallMethodResult that = (CallMethodResult) o;
    return (getStatusCode() == that.getStatusCode())
        && (getNoOfInputArgumentResults() == that.getNoOfInputArgumentResults())
        && (getInputArgumentResults() == that.getInputArgumentResults())
        && (getNoOfInputArgumentDiagnosticInfos() == that.getNoOfInputArgumentDiagnosticInfos())
        && (getInputArgumentDiagnosticInfos() == that.getInputArgumentDiagnosticInfos())
        && (getNoOfOutputArguments() == that.getNoOfOutputArguments())
        && (getOutputArguments() == that.getOutputArguments())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getStatusCode(),
        getNoOfInputArgumentResults(),
        getInputArgumentResults(),
        getNoOfInputArgumentDiagnosticInfos(),
        getInputArgumentDiagnosticInfos(),
        getNoOfOutputArguments(),
        getOutputArguments());
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
