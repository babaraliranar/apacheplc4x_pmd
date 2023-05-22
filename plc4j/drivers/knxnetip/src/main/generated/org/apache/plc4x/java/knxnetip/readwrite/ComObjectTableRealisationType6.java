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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ComObjectTableRealisationType6 extends ComObjectTable implements Message {

  // Accessors for discriminator values.
  public FirmwareType getFirmwareType() {
    return FirmwareType.SYSTEM_300;
  }

  // Properties.
  protected final GroupObjectDescriptorRealisationType6 comObjectDescriptors;

  public ComObjectTableRealisationType6(
      GroupObjectDescriptorRealisationType6 comObjectDescriptors) {
    super();
    this.comObjectDescriptors = comObjectDescriptors;
  }

  public GroupObjectDescriptorRealisationType6 getComObjectDescriptors() {
    return comObjectDescriptors;
  }

  @Override
  protected void serializeComObjectTableChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ComObjectTableRealisationType6");

    // Simple Field (comObjectDescriptors)
    writeSimpleField(
        "comObjectDescriptors", comObjectDescriptors, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ComObjectTableRealisationType6");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ComObjectTableRealisationType6 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (comObjectDescriptors)
    lengthInBits += comObjectDescriptors.getLengthInBits();

    return lengthInBits;
  }

  public static ComObjectTableBuilder staticParseComObjectTableBuilder(
      ReadBuffer readBuffer, FirmwareType firmwareType) throws ParseException {
    readBuffer.pullContext("ComObjectTableRealisationType6");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    GroupObjectDescriptorRealisationType6 comObjectDescriptors =
        readSimpleField(
            "comObjectDescriptors",
            new DataReaderComplexDefault<>(
                () -> GroupObjectDescriptorRealisationType6.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ComObjectTableRealisationType6");
    // Create the instance
    return new ComObjectTableRealisationType6BuilderImpl(comObjectDescriptors);
  }

  public static class ComObjectTableRealisationType6BuilderImpl
      implements ComObjectTable.ComObjectTableBuilder {
    private final GroupObjectDescriptorRealisationType6 comObjectDescriptors;

    public ComObjectTableRealisationType6BuilderImpl(
        GroupObjectDescriptorRealisationType6 comObjectDescriptors) {
      this.comObjectDescriptors = comObjectDescriptors;
    }

    public ComObjectTableRealisationType6 build() {
      ComObjectTableRealisationType6 comObjectTableRealisationType6 =
          new ComObjectTableRealisationType6(comObjectDescriptors);
      return comObjectTableRealisationType6;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ComObjectTableRealisationType6)) {
      return false;
    }
    ComObjectTableRealisationType6 that = (ComObjectTableRealisationType6) o;
    return (getComObjectDescriptors() == that.getComObjectDescriptors())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getComObjectDescriptors());
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
