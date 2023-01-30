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

public class ApduDataContainer extends Apdu implements Message {

  // Accessors for discriminator values.
  public Byte getControl() {
    return (byte) 0;
  }

  // Properties.
  protected final ApduData dataApdu;

  public ApduDataContainer(boolean numbered, byte counter, ApduData dataApdu) {
    super(numbered, counter);
    this.dataApdu = dataApdu;
  }

  public ApduData getDataApdu() {
    return dataApdu;
  }

  @Override
  protected void serializeApduChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApduDataContainer");

    // Simple Field (dataApdu)
    writeSimpleField("dataApdu", dataApdu, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ApduDataContainer");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduDataContainer _value = this;

    // Simple field (dataApdu)
    lengthInBits += dataApdu.getLengthInBits();

    return lengthInBits;
  }

  public static ApduBuilder staticParseApduBuilder(ReadBuffer readBuffer, Short dataLength)
      throws ParseException {
    readBuffer.pullContext("ApduDataContainer");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ApduData dataApdu =
        readSimpleField(
            "dataApdu",
            new DataReaderComplexDefault<>(
                () -> ApduData.staticParse(readBuffer, (short) (dataLength)), readBuffer));

    readBuffer.closeContext("ApduDataContainer");
    // Create the instance
    return new ApduDataContainerBuilderImpl(dataApdu);
  }

  public static class ApduDataContainerBuilderImpl implements Apdu.ApduBuilder {
    private final ApduData dataApdu;

    public ApduDataContainerBuilderImpl(ApduData dataApdu) {
      this.dataApdu = dataApdu;
    }

    public ApduDataContainer build(boolean numbered, byte counter) {
      ApduDataContainer apduDataContainer = new ApduDataContainer(numbered, counter, dataApdu);
      return apduDataContainer;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduDataContainer)) {
      return false;
    }
    ApduDataContainer that = (ApduDataContainer) o;
    return (getDataApdu() == that.getDataApdu()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDataApdu());
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
