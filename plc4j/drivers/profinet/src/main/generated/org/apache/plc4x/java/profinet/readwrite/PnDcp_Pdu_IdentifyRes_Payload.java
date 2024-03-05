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
package org.apache.plc4x.java.profinet.readwrite;

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

public abstract class PnDcp_Pdu_IdentifyRes_Payload implements Message {

  // Abstract accessors for discriminator values.
  public abstract PnDcp_ServiceId getServiceId();

  public abstract Boolean getServiceTypeResponse();

  // Properties.
  protected final PnDcp_ServiceType serviceType;
  protected final long xid;
  protected final int responseDelayFactorOrPadding;

  public PnDcp_Pdu_IdentifyRes_Payload(
      PnDcp_ServiceType serviceType, long xid, int responseDelayFactorOrPadding) {
    super();
    this.serviceType = serviceType;
    this.xid = xid;
    this.responseDelayFactorOrPadding = responseDelayFactorOrPadding;
  }

  public PnDcp_ServiceType getServiceType() {
    return serviceType;
  }

  public long getXid() {
    return xid;
  }

  public int getResponseDelayFactorOrPadding() {
    return responseDelayFactorOrPadding;
  }

  protected abstract void serializePnDcp_Pdu_IdentifyRes_PayloadChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnDcp_Pdu_IdentifyRes_Payload");

    // Discriminator Field (serviceId) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "serviceId",
        "PnDcp_ServiceId",
        getServiceId(),
        new DataWriterEnumDefault<>(
            PnDcp_ServiceId::getValue, PnDcp_ServiceId::name, writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (serviceType)
    writeSimpleField(
        "serviceType",
        serviceType,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (xid)
    writeSimpleField(
        "xid",
        xid,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (responseDelayFactorOrPadding)
    writeSimpleField(
        "responseDelayFactorOrPadding",
        responseDelayFactorOrPadding,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (dcpDataLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int dcpDataLength = (int) ((getLengthInBytes()) - (12));
    writeImplicitField(
        "dcpDataLength",
        dcpDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch field (Serialize the sub-type)
    serializePnDcp_Pdu_IdentifyRes_PayloadChild(writeBuffer);

    writeBuffer.popContext("PnDcp_Pdu_IdentifyRes_Payload");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnDcp_Pdu_IdentifyRes_Payload _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (serviceId)
    lengthInBits += 8;

    // Simple field (serviceType)
    lengthInBits += serviceType.getLengthInBits();

    // Simple field (xid)
    lengthInBits += 32;

    // Simple field (responseDelayFactorOrPadding)
    lengthInBits += 16;

    // Implicit Field (dcpDataLength)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static PnDcp_Pdu_IdentifyRes_Payload staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnDcp_Pdu_IdentifyRes_Payload staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnDcp_Pdu_IdentifyRes_Payload");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PnDcp_ServiceId serviceId =
        readDiscriminatorEnumField(
            "serviceId",
            "PnDcp_ServiceId",
            new DataReaderEnumDefault<>(
                PnDcp_ServiceId::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    PnDcp_ServiceType serviceType =
        readSimpleField(
            "serviceType",
            new DataReaderComplexDefault<>(
                () -> PnDcp_ServiceType.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    long xid =
        readSimpleField(
            "xid",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int responseDelayFactorOrPadding =
        readSimpleField(
            "responseDelayFactorOrPadding",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int dcpDataLength =
        readImplicitField(
            "dcpDataLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    PnDcp_Pdu_IdentifyRes_PayloadBuilder builder = null;
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "serviceId="
              + serviceId
              + " "
              + "serviceTyperesponse="
              + serviceType.getResponse()
              + "]");
    }

    readBuffer.closeContext("PnDcp_Pdu_IdentifyRes_Payload");
    // Create the instance
    PnDcp_Pdu_IdentifyRes_Payload _pnDcp_Pdu_IdentifyRes_Payload =
        builder.build(serviceType, xid, responseDelayFactorOrPadding);
    return _pnDcp_Pdu_IdentifyRes_Payload;
  }

  public interface PnDcp_Pdu_IdentifyRes_PayloadBuilder {
    PnDcp_Pdu_IdentifyRes_Payload build(
        PnDcp_ServiceType serviceType, long xid, int responseDelayFactorOrPadding);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Pdu_IdentifyRes_Payload)) {
      return false;
    }
    PnDcp_Pdu_IdentifyRes_Payload that = (PnDcp_Pdu_IdentifyRes_Payload) o;
    return (getServiceType() == that.getServiceType())
        && (getXid() == that.getXid())
        && (getResponseDelayFactorOrPadding() == that.getResponseDelayFactorOrPadding())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getServiceType(), getXid(), getResponseDelayFactorOrPadding());
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
