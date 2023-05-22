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

public abstract class BACnetEventParameter implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BACnetTagHeader peekedTagHeader;

  public BACnetEventParameter(BACnetTagHeader peekedTagHeader) {
    super();
    this.peekedTagHeader = peekedTagHeader;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  protected abstract void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetEventParameter");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Switch field (Serialize the sub-type)
    serializeBACnetEventParameterChild(writeBuffer);

    writeBuffer.popContext("BACnetEventParameter");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetEventParameter _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static BACnetEventParameter staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetEventParameter staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameter");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader peekedTagHeader =
        readPeekField(
            "peekedTagHeader",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    short peekedTagNumber =
        readVirtualField("peekedTagNumber", short.class, peekedTagHeader.getActualTagNumber());

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetEventParameterBuilder builder = null;
    if (EvaluationHelper.equals(peekedTagNumber, (short) 0)) {
      builder =
          BACnetEventParameterChangeOfBitstring.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 1)) {
      builder =
          BACnetEventParameterChangeOfState.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 2)) {
      builder =
          BACnetEventParameterChangeOfValue.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 3)) {
      builder =
          BACnetEventParameterCommandFailure.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 4)) {
      builder =
          BACnetEventParameterFloatingLimit.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 5)) {
      builder = BACnetEventParameterOutOfRange.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 8)) {
      builder =
          BACnetEventParameterChangeOfLifeSavety.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 9)) {
      builder = BACnetEventParameterExtended.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 10)) {
      builder = BACnetEventParameterBufferReady.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 11)) {
      builder =
          BACnetEventParameterUnsignedRange.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 13)) {
      builder = BACnetEventParameterAccessEvent.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 14)) {
      builder =
          BACnetEventParameterDoubleOutOfRange.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 15)) {
      builder =
          BACnetEventParameterSignedOutOfRange.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 16)) {
      builder =
          BACnetEventParameterUnsignedOutOfRange.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 17)) {
      builder =
          BACnetEventParameterChangeOfCharacterString.staticParseBACnetEventParameterBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 18)) {
      builder =
          BACnetEventParameterChangeOfStatusFlags.staticParseBACnetEventParameterBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 20)) {
      builder = BACnetEventParameterNone.staticParseBACnetEventParameterBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 21)) {
      builder =
          BACnetEventParameterChangeOfDiscreteValue.staticParseBACnetEventParameterBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 22)) {
      builder =
          BACnetEventParameterChangeOfTimer.staticParseBACnetEventParameterBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedTagNumber="
              + peekedTagNumber
              + "]");
    }

    readBuffer.closeContext("BACnetEventParameter");
    // Create the instance
    BACnetEventParameter _bACnetEventParameter = builder.build(peekedTagHeader);
    return _bACnetEventParameter;
  }

  public interface BACnetEventParameterBuilder {
    BACnetEventParameter build(BACnetTagHeader peekedTagHeader);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameter)) {
      return false;
    }
    BACnetEventParameter that = (BACnetEventParameter) o;
    return (getPeekedTagHeader() == that.getPeekedTagHeader()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getPeekedTagHeader());
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
