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

public class CBusPointToMultiPointCommandStatus extends CBusPointToMultiPointCommand
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final StatusRequest statusRequest;

  // Arguments.
  protected final CBusOptions cBusOptions;
  // Reserved Fields
  private Byte reservedField0;
  private Byte reservedField1;

  public CBusPointToMultiPointCommandStatus(
      byte peekedApplication, StatusRequest statusRequest, CBusOptions cBusOptions) {
    super(peekedApplication, cBusOptions);
    this.statusRequest = statusRequest;
    this.cBusOptions = cBusOptions;
  }

  public StatusRequest getStatusRequest() {
    return statusRequest;
  }

  @Override
  protected void serializeCBusPointToMultiPointCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CBusPointToMultiPointCommandStatus");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0xFF,
        writeByte(writeBuffer, 8));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (byte) 0x00,
        writeByte(writeBuffer, 8));

    // Simple Field (statusRequest)
    writeSimpleField("statusRequest", statusRequest, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CBusPointToMultiPointCommandStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CBusPointToMultiPointCommandStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (statusRequest)
    lengthInBits += statusRequest.getLengthInBits();

    return lengthInBits;
  }

  public static CBusPointToMultiPointCommandBuilder staticParseCBusPointToMultiPointCommandBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("CBusPointToMultiPointCommandStatus");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0xFF);

    Byte reservedField1 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0x00);

    StatusRequest statusRequest =
        readSimpleField(
            "statusRequest",
            new DataReaderComplexDefault<>(
                () -> StatusRequest.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("CBusPointToMultiPointCommandStatus");
    // Create the instance
    return new CBusPointToMultiPointCommandStatusBuilderImpl(
        statusRequest, cBusOptions, reservedField0, reservedField1);
  }

  public static class CBusPointToMultiPointCommandStatusBuilderImpl
      implements CBusPointToMultiPointCommand.CBusPointToMultiPointCommandBuilder {
    private final StatusRequest statusRequest;
    private final CBusOptions cBusOptions;
    private final Byte reservedField0;
    private final Byte reservedField1;

    public CBusPointToMultiPointCommandStatusBuilderImpl(
        StatusRequest statusRequest,
        CBusOptions cBusOptions,
        Byte reservedField0,
        Byte reservedField1) {
      this.statusRequest = statusRequest;
      this.cBusOptions = cBusOptions;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
    }

    public CBusPointToMultiPointCommandStatus build(
        byte peekedApplication, CBusOptions cBusOptions) {
      CBusPointToMultiPointCommandStatus cBusPointToMultiPointCommandStatus =
          new CBusPointToMultiPointCommandStatus(peekedApplication, statusRequest, cBusOptions);
      cBusPointToMultiPointCommandStatus.reservedField0 = reservedField0;
      cBusPointToMultiPointCommandStatus.reservedField1 = reservedField1;
      return cBusPointToMultiPointCommandStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusPointToMultiPointCommandStatus)) {
      return false;
    }
    CBusPointToMultiPointCommandStatus that = (CBusPointToMultiPointCommandStatus) o;
    return (getStatusRequest() == that.getStatusRequest()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatusRequest());
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
