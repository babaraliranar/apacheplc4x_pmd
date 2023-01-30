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

public class BACnetConfirmedServiceRequestReadPropertyConditional
    extends BACnetConfirmedServiceRequest implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.READ_PROPERTY_CONDITIONAL;
  }

  // Properties.
  protected final byte[] bytesOfRemovedService;

  // Arguments.
  protected final Long serviceRequestPayloadLength;
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestReadPropertyConditional(
      byte[] bytesOfRemovedService, Long serviceRequestPayloadLength, Long serviceRequestLength) {
    super(serviceRequestLength);
    this.bytesOfRemovedService = bytesOfRemovedService;
    this.serviceRequestPayloadLength = serviceRequestPayloadLength;
    this.serviceRequestLength = serviceRequestLength;
  }

  public byte[] getBytesOfRemovedService() {
    return bytesOfRemovedService;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestReadPropertyConditional");

    // Array Field (bytesOfRemovedService)
    writeByteArrayField(
        "bytesOfRemovedService", bytesOfRemovedService, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("BACnetConfirmedServiceRequestReadPropertyConditional");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestReadPropertyConditional _value = this;

    // Array field
    if (bytesOfRemovedService != null) {
      lengthInBits += 8 * bytesOfRemovedService.length;
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestPayloadLength, Long serviceRequestLength)
          throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestReadPropertyConditional");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte[] bytesOfRemovedService =
        readBuffer.readByteArray(
            "bytesOfRemovedService", Math.toIntExact(serviceRequestPayloadLength));

    readBuffer.closeContext("BACnetConfirmedServiceRequestReadPropertyConditional");
    // Create the instance
    return new BACnetConfirmedServiceRequestReadPropertyConditionalBuilderImpl(
        bytesOfRemovedService, serviceRequestPayloadLength, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestReadPropertyConditionalBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final byte[] bytesOfRemovedService;
    private final Long serviceRequestPayloadLength;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestReadPropertyConditionalBuilderImpl(
        byte[] bytesOfRemovedService, Long serviceRequestPayloadLength, Long serviceRequestLength) {
      this.bytesOfRemovedService = bytesOfRemovedService;
      this.serviceRequestPayloadLength = serviceRequestPayloadLength;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestReadPropertyConditional build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestReadPropertyConditional
          bACnetConfirmedServiceRequestReadPropertyConditional =
              new BACnetConfirmedServiceRequestReadPropertyConditional(
                  bytesOfRemovedService, serviceRequestPayloadLength, serviceRequestLength);
      return bACnetConfirmedServiceRequestReadPropertyConditional;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestReadPropertyConditional)) {
      return false;
    }
    BACnetConfirmedServiceRequestReadPropertyConditional that =
        (BACnetConfirmedServiceRequestReadPropertyConditional) o;
    return (getBytesOfRemovedService() == that.getBytesOfRemovedService())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBytesOfRemovedService());
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
