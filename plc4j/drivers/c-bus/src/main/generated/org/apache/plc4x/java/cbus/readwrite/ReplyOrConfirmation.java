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

public abstract class ReplyOrConfirmation implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final byte peekedByte;

  // Arguments.
  protected final CBusOptions cBusOptions;
  protected final RequestContext requestContext;

  public ReplyOrConfirmation(
      byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext) {
    super();
    this.peekedByte = peekedByte;
    this.cBusOptions = cBusOptions;
    this.requestContext = requestContext;
  }

  public byte getPeekedByte() {
    return peekedByte;
  }

  public boolean getIsAlpha() {
    return (boolean) ((((getPeekedByte()) >= (0x67))) && (((getPeekedByte()) <= (0x7A))));
  }

  protected abstract void serializeReplyOrConfirmationChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ReplyOrConfirmation");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isAlpha = getIsAlpha();
    writeBuffer.writeVirtual("isAlpha", isAlpha);

    // Switch field (Serialize the sub-type)
    serializeReplyOrConfirmationChild(writeBuffer);

    writeBuffer.popContext("ReplyOrConfirmation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ReplyOrConfirmation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static ReplyOrConfirmation staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    CBusOptions cBusOptions;
    if (args[0] instanceof CBusOptions) {
      cBusOptions = (CBusOptions) args[0];
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type CBusOptions or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    RequestContext requestContext;
    if (args[1] instanceof RequestContext) {
      requestContext = (RequestContext) args[1];
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type RequestContext or a string which is parseable but was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, cBusOptions, requestContext);
  }

  public static ReplyOrConfirmation staticParse(
      ReadBuffer readBuffer, CBusOptions cBusOptions, RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("ReplyOrConfirmation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte peekedByte = readPeekField("peekedByte", readByte(readBuffer, 8));
    boolean isAlpha =
        readVirtualField(
            "isAlpha", boolean.class, (((peekedByte) >= (0x67))) && (((peekedByte) <= (0x7A))));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    ReplyOrConfirmationBuilder builder = null;
    if (EvaluationHelper.equals(isAlpha, (boolean) false)
        && EvaluationHelper.equals(peekedByte, (byte) 0x21)) {
      builder =
          ServerErrorReply.staticParseReplyOrConfirmationBuilder(
              readBuffer, cBusOptions, requestContext);
    } else if (EvaluationHelper.equals(isAlpha, (boolean) true)) {
      builder =
          ReplyOrConfirmationConfirmation.staticParseReplyOrConfirmationBuilder(
              readBuffer, cBusOptions, requestContext);
    } else if (EvaluationHelper.equals(isAlpha, (boolean) false)) {
      builder =
          ReplyOrConfirmationReply.staticParseReplyOrConfirmationBuilder(
              readBuffer, cBusOptions, requestContext);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "isAlpha="
              + isAlpha
              + " "
              + "peekedByte="
              + peekedByte
              + "]");
    }

    readBuffer.closeContext("ReplyOrConfirmation");
    // Create the instance
    ReplyOrConfirmation _replyOrConfirmation =
        builder.build(peekedByte, cBusOptions, requestContext);
    return _replyOrConfirmation;
  }

  public interface ReplyOrConfirmationBuilder {
    ReplyOrConfirmation build(
        byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ReplyOrConfirmation)) {
      return false;
    }
    ReplyOrConfirmation that = (ReplyOrConfirmation) o;
    return (getPeekedByte() == that.getPeekedByte()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getPeekedByte());
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
