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

public class BACnetPropertyAccessResultAccessResultPropertyAccessError
    extends BACnetPropertyAccessResultAccessResult implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final ErrorEnclosed propertyAccessError;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;
  protected final BACnetPropertyIdentifier propertyIdentifierArgument;
  protected final BACnetTagPayloadUnsignedInteger propertyArrayIndexArgument;

  public BACnetPropertyAccessResultAccessResultPropertyAccessError(
      BACnetTagHeader peekedTagHeader,
      ErrorEnclosed propertyAccessError,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger propertyArrayIndexArgument) {
    super(
        peekedTagHeader,
        objectTypeArgument,
        propertyIdentifierArgument,
        propertyArrayIndexArgument);
    this.propertyAccessError = propertyAccessError;
    this.objectTypeArgument = objectTypeArgument;
    this.propertyIdentifierArgument = propertyIdentifierArgument;
    this.propertyArrayIndexArgument = propertyArrayIndexArgument;
  }

  public ErrorEnclosed getPropertyAccessError() {
    return propertyAccessError;
  }

  @Override
  protected void serializeBACnetPropertyAccessResultAccessResultChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPropertyAccessResultAccessResultPropertyAccessError");

    // Simple Field (propertyAccessError)
    writeSimpleField(
        "propertyAccessError", propertyAccessError, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyAccessResultAccessResultPropertyAccessError");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyAccessResultAccessResultPropertyAccessError _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (propertyAccessError)
    lengthInBits += propertyAccessError.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyAccessResultAccessResultBuilder
      staticParseBACnetPropertyAccessResultAccessResultBuilder(
          ReadBuffer readBuffer,
          BACnetObjectType objectTypeArgument,
          BACnetPropertyIdentifier propertyIdentifierArgument,
          BACnetTagPayloadUnsignedInteger propertyArrayIndexArgument)
          throws ParseException {
    readBuffer.pullContext("BACnetPropertyAccessResultAccessResultPropertyAccessError");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ErrorEnclosed propertyAccessError =
        readSimpleField(
            "propertyAccessError",
            new DataReaderComplexDefault<>(
                () -> ErrorEnclosed.staticParse(readBuffer, (short) (5)), readBuffer));

    readBuffer.closeContext("BACnetPropertyAccessResultAccessResultPropertyAccessError");
    // Create the instance
    return new BACnetPropertyAccessResultAccessResultPropertyAccessErrorBuilderImpl(
        propertyAccessError,
        objectTypeArgument,
        propertyIdentifierArgument,
        propertyArrayIndexArgument);
  }

  public static class BACnetPropertyAccessResultAccessResultPropertyAccessErrorBuilderImpl
      implements BACnetPropertyAccessResultAccessResult
          .BACnetPropertyAccessResultAccessResultBuilder {
    private final ErrorEnclosed propertyAccessError;
    private final BACnetObjectType objectTypeArgument;
    private final BACnetPropertyIdentifier propertyIdentifierArgument;
    private final BACnetTagPayloadUnsignedInteger propertyArrayIndexArgument;

    public BACnetPropertyAccessResultAccessResultPropertyAccessErrorBuilderImpl(
        ErrorEnclosed propertyAccessError,
        BACnetObjectType objectTypeArgument,
        BACnetPropertyIdentifier propertyIdentifierArgument,
        BACnetTagPayloadUnsignedInteger propertyArrayIndexArgument) {
      this.propertyAccessError = propertyAccessError;
      this.objectTypeArgument = objectTypeArgument;
      this.propertyIdentifierArgument = propertyIdentifierArgument;
      this.propertyArrayIndexArgument = propertyArrayIndexArgument;
    }

    public BACnetPropertyAccessResultAccessResultPropertyAccessError build(
        BACnetTagHeader peekedTagHeader,
        BACnetObjectType objectTypeArgument,
        BACnetPropertyIdentifier propertyIdentifierArgument,
        BACnetTagPayloadUnsignedInteger propertyArrayIndexArgument) {
      BACnetPropertyAccessResultAccessResultPropertyAccessError
          bACnetPropertyAccessResultAccessResultPropertyAccessError =
              new BACnetPropertyAccessResultAccessResultPropertyAccessError(
                  peekedTagHeader,
                  propertyAccessError,
                  objectTypeArgument,
                  propertyIdentifierArgument,
                  propertyArrayIndexArgument);
      return bACnetPropertyAccessResultAccessResultPropertyAccessError;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyAccessResultAccessResultPropertyAccessError)) {
      return false;
    }
    BACnetPropertyAccessResultAccessResultPropertyAccessError that =
        (BACnetPropertyAccessResultAccessResultPropertyAccessError) o;
    return (getPropertyAccessError() == that.getPropertyAccessError())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPropertyAccessError());
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
