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

public class APDUReject extends APDU implements Message {

  // Accessors for discriminator values.
  public ApduType getApduType() {
    return ApduType.REJECT_PDU;
  }

  // Properties.
  protected final short originalInvokeId;
  protected final BACnetRejectReasonTagged rejectReason;

  // Arguments.
  protected final Integer apduLength;
  // Reserved Fields
  private Byte reservedField0;

  public APDUReject(
      short originalInvokeId, BACnetRejectReasonTagged rejectReason, Integer apduLength) {
    super(apduLength);
    this.originalInvokeId = originalInvokeId;
    this.rejectReason = rejectReason;
    this.apduLength = apduLength;
  }

  public short getOriginalInvokeId() {
    return originalInvokeId;
  }

  public BACnetRejectReasonTagged getRejectReason() {
    return rejectReason;
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("APDUReject");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0x00,
        writeUnsignedByte(writeBuffer, 4));

    // Simple Field (originalInvokeId)
    writeSimpleField("originalInvokeId", originalInvokeId, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (rejectReason)
    writeSimpleField("rejectReason", rejectReason, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("APDUReject");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUReject _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 4;

    // Simple field (originalInvokeId)
    lengthInBits += 8;

    // Simple field (rejectReason)
    lengthInBits += rejectReason.getLengthInBits();

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("APDUReject");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 4), (byte) 0x00);

    short originalInvokeId = readSimpleField("originalInvokeId", readUnsignedShort(readBuffer, 8));

    BACnetRejectReasonTagged rejectReason =
        readSimpleField(
            "rejectReason",
            new DataReaderComplexDefault<>(
                () -> BACnetRejectReasonTagged.staticParse(readBuffer, (long) (1L)), readBuffer));

    readBuffer.closeContext("APDUReject");
    // Create the instance
    return new APDURejectBuilderImpl(originalInvokeId, rejectReason, apduLength, reservedField0);
  }

  public static class APDURejectBuilderImpl implements APDU.APDUBuilder {
    private final short originalInvokeId;
    private final BACnetRejectReasonTagged rejectReason;
    private final Integer apduLength;
    private final Byte reservedField0;

    public APDURejectBuilderImpl(
        short originalInvokeId,
        BACnetRejectReasonTagged rejectReason,
        Integer apduLength,
        Byte reservedField0) {
      this.originalInvokeId = originalInvokeId;
      this.rejectReason = rejectReason;
      this.apduLength = apduLength;
      this.reservedField0 = reservedField0;
    }

    public APDUReject build(Integer apduLength) {

      APDUReject aPDUReject = new APDUReject(originalInvokeId, rejectReason, apduLength);
      aPDUReject.reservedField0 = reservedField0;
      return aPDUReject;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUReject)) {
      return false;
    }
    APDUReject that = (APDUReject) o;
    return (getOriginalInvokeId() == that.getOriginalInvokeId())
        && (getRejectReason() == that.getRejectReason())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOriginalInvokeId(), getRejectReason());
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
