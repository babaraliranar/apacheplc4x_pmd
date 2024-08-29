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

public abstract class BACnetValueSource implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BACnetTagHeader peekedTagHeader;

  public BACnetValueSource(BACnetTagHeader peekedTagHeader) {
    super();
    this.peekedTagHeader = peekedTagHeader;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  protected abstract void serializeBACnetValueSourceChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetValueSource");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Switch field (Serialize the sub-type)
    serializeBACnetValueSourceChild(writeBuffer);

    writeBuffer.popContext("BACnetValueSource");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetValueSource _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static BACnetValueSource staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetValueSource");
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
    BACnetValueSourceBuilder builder = null;
    if (EvaluationHelper.equals(peekedTagNumber, (short) 0)) {
      builder = BACnetValueSourceNone.staticParseBACnetValueSourceBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 1)) {
      builder = BACnetValueSourceObject.staticParseBACnetValueSourceBuilder(readBuffer);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 2)) {
      builder = BACnetValueSourceAddress.staticParseBACnetValueSourceBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedTagNumber="
              + peekedTagNumber
              + "]");
    }

    readBuffer.closeContext("BACnetValueSource");
    // Create the instance
    BACnetValueSource _bACnetValueSource = builder.build(peekedTagHeader);
    return _bACnetValueSource;
  }

  public interface BACnetValueSourceBuilder {
    BACnetValueSource build(BACnetTagHeader peekedTagHeader);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetValueSource)) {
      return false;
    }
    BACnetValueSource that = (BACnetValueSource) o;
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
