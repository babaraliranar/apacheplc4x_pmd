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

public class BACnetPropertyStatesLiftGroupMode extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetLiftGroupModeTagged liftGroupMode;

  public BACnetPropertyStatesLiftGroupMode(
      BACnetTagHeader peekedTagHeader, BACnetLiftGroupModeTagged liftGroupMode) {
    super(peekedTagHeader);
    this.liftGroupMode = liftGroupMode;
  }

  public BACnetLiftGroupModeTagged getLiftGroupMode() {
    return liftGroupMode;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesLiftGroupMode");

    // Simple Field (liftGroupMode)
    writeSimpleField("liftGroupMode", liftGroupMode, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesLiftGroupMode");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesLiftGroupMode _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (liftGroupMode)
    lengthInBits += liftGroupMode.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBuilder staticParseBACnetPropertyStatesBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesLiftGroupMode");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetLiftGroupModeTagged liftGroupMode =
        readSimpleField(
            "liftGroupMode",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetLiftGroupModeTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesLiftGroupMode");
    // Create the instance
    return new BACnetPropertyStatesLiftGroupModeBuilderImpl(liftGroupMode);
  }

  public static class BACnetPropertyStatesLiftGroupModeBuilderImpl
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetLiftGroupModeTagged liftGroupMode;

    public BACnetPropertyStatesLiftGroupModeBuilderImpl(BACnetLiftGroupModeTagged liftGroupMode) {
      this.liftGroupMode = liftGroupMode;
    }

    public BACnetPropertyStatesLiftGroupMode build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesLiftGroupMode bACnetPropertyStatesLiftGroupMode =
          new BACnetPropertyStatesLiftGroupMode(peekedTagHeader, liftGroupMode);
      return bACnetPropertyStatesLiftGroupMode;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesLiftGroupMode)) {
      return false;
    }
    BACnetPropertyStatesLiftGroupMode that = (BACnetPropertyStatesLiftGroupMode) o;
    return (getLiftGroupMode() == that.getLiftGroupMode()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLiftGroupMode());
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
