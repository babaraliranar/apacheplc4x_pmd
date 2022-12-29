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
package org.apache.plc4x.java.cbus.readwrite;

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

public abstract class ErrorReportingSystemCategoryType implements Message {

  // Abstract accessors for discriminator values.
  public abstract ErrorReportingSystemCategoryClass getErrorReportingSystemCategoryClass();

  public ErrorReportingSystemCategoryType() {
    super();
  }

  protected abstract void serializeErrorReportingSystemCategoryTypeChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ErrorReportingSystemCategoryType");

    // Switch field (Serialize the sub-type)
    serializeErrorReportingSystemCategoryTypeChild(writeBuffer);

    writeBuffer.popContext("ErrorReportingSystemCategoryType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ErrorReportingSystemCategoryType _value = this;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static ErrorReportingSystemCategoryType staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    ErrorReportingSystemCategoryClass errorReportingSystemCategoryClass;
    if (args[0] instanceof ErrorReportingSystemCategoryClass) {
      errorReportingSystemCategoryClass = (ErrorReportingSystemCategoryClass) args[0];
    } else if (args[0] instanceof String) {
      errorReportingSystemCategoryClass =
          ErrorReportingSystemCategoryClass.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type ErrorReportingSystemCategoryClass or a string which is"
              + " parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, errorReportingSystemCategoryClass);
  }

  public static ErrorReportingSystemCategoryType staticParse(
      ReadBuffer readBuffer, ErrorReportingSystemCategoryClass errorReportingSystemCategoryClass)
      throws ParseException {
    readBuffer.pullContext("ErrorReportingSystemCategoryType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    ErrorReportingSystemCategoryTypeBuilder builder = null;
    if (EvaluationHelper.equals(
        errorReportingSystemCategoryClass, ErrorReportingSystemCategoryClass.INPUT_UNITS)) {
      builder =
          ErrorReportingSystemCategoryTypeInputUnits.staticParseBuilder(
              readBuffer, errorReportingSystemCategoryClass);
    } else if (EvaluationHelper.equals(
        errorReportingSystemCategoryClass, ErrorReportingSystemCategoryClass.SUPPORT_UNITS)) {
      builder =
          ErrorReportingSystemCategoryTypeSupportUnits.staticParseBuilder(
              readBuffer, errorReportingSystemCategoryClass);
    } else if (EvaluationHelper.equals(
        errorReportingSystemCategoryClass,
        ErrorReportingSystemCategoryClass.BUILDING_MANAGEMENT_SYSTEMS)) {
      builder =
          ErrorReportingSystemCategoryTypeBuildingManagementSystems.staticParseBuilder(
              readBuffer, errorReportingSystemCategoryClass);
    } else if (EvaluationHelper.equals(
        errorReportingSystemCategoryClass, ErrorReportingSystemCategoryClass.OUTPUT_UNITS)) {
      builder =
          ErrorReportingSystemCategoryTypeOutputUnits.staticParseBuilder(
              readBuffer, errorReportingSystemCategoryClass);
    } else if (EvaluationHelper.equals(
        errorReportingSystemCategoryClass, ErrorReportingSystemCategoryClass.CLIMATE_CONTROLLERS)) {
      builder =
          ErrorReportingSystemCategoryTypeClimateControllers.staticParseBuilder(
              readBuffer, errorReportingSystemCategoryClass);
    } else if (true) {
      builder =
          ErrorReportingSystemCategoryTypeReserved.staticParseBuilder(
              readBuffer, errorReportingSystemCategoryClass);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "errorReportingSystemCategoryClass="
              + errorReportingSystemCategoryClass
              + "]");
    }

    readBuffer.closeContext("ErrorReportingSystemCategoryType");
    // Create the instance
    ErrorReportingSystemCategoryType _errorReportingSystemCategoryType = builder.build();
    return _errorReportingSystemCategoryType;
  }

  public static interface ErrorReportingSystemCategoryTypeBuilder {
    ErrorReportingSystemCategoryType build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ErrorReportingSystemCategoryType)) {
      return false;
    }
    ErrorReportingSystemCategoryType that = (ErrorReportingSystemCategoryType) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
