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

public class BACnetUnconfirmedServiceRequestUTCTimeSynchronization
    extends BACnetUnconfirmedServiceRequest implements Message {

  // Accessors for discriminator values.
  public BACnetUnconfirmedServiceChoice getServiceChoice() {
    return BACnetUnconfirmedServiceChoice.UTC_TIME_SYNCHRONIZATION;
  }

  // Properties.
  protected final BACnetApplicationTagDate synchronizedDate;
  protected final BACnetApplicationTagTime synchronizedTime;

  // Arguments.
  protected final Integer serviceRequestLength;

  public BACnetUnconfirmedServiceRequestUTCTimeSynchronization(
      BACnetApplicationTagDate synchronizedDate,
      BACnetApplicationTagTime synchronizedTime,
      Integer serviceRequestLength) {
    super(serviceRequestLength);
    this.synchronizedDate = synchronizedDate;
    this.synchronizedTime = synchronizedTime;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetApplicationTagDate getSynchronizedDate() {
    return synchronizedDate;
  }

  public BACnetApplicationTagTime getSynchronizedTime() {
    return synchronizedTime;
  }

  @Override
  protected void serializeBACnetUnconfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization");

    // Simple Field (synchronizedDate)
    writeSimpleField(
        "synchronizedDate", synchronizedDate, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (synchronizedTime)
    writeSimpleField(
        "synchronizedTime", synchronizedTime, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetUnconfirmedServiceRequestUTCTimeSynchronization _value = this;

    // Simple field (synchronizedDate)
    lengthInBits += synchronizedDate.getLengthInBits();

    // Simple field (synchronizedTime)
    lengthInBits += synchronizedTime.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetUnconfirmedServiceRequestBuilder
      staticParseBACnetUnconfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Integer serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagDate synchronizedDate =
        readSimpleField(
            "synchronizedDate",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagDate) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagTime synchronizedTime =
        readSimpleField(
            "synchronizedTime",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagTime) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization");
    // Create the instance
    return new BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilderImpl(
        synchronizedDate, synchronizedTime, serviceRequestLength);
  }

  public static class BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilderImpl
      implements BACnetUnconfirmedServiceRequest.BACnetUnconfirmedServiceRequestBuilder {
    private final BACnetApplicationTagDate synchronizedDate;
    private final BACnetApplicationTagTime synchronizedTime;
    private final Integer serviceRequestLength;

    public BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilderImpl(
        BACnetApplicationTagDate synchronizedDate,
        BACnetApplicationTagTime synchronizedTime,
        Integer serviceRequestLength) {
      this.synchronizedDate = synchronizedDate;
      this.synchronizedTime = synchronizedTime;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetUnconfirmedServiceRequestUTCTimeSynchronization build(
        Integer serviceRequestLength) {
      BACnetUnconfirmedServiceRequestUTCTimeSynchronization
          bACnetUnconfirmedServiceRequestUTCTimeSynchronization =
              new BACnetUnconfirmedServiceRequestUTCTimeSynchronization(
                  synchronizedDate, synchronizedTime, serviceRequestLength);
      return bACnetUnconfirmedServiceRequestUTCTimeSynchronization;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetUnconfirmedServiceRequestUTCTimeSynchronization)) {
      return false;
    }
    BACnetUnconfirmedServiceRequestUTCTimeSynchronization that =
        (BACnetUnconfirmedServiceRequestUTCTimeSynchronization) o;
    return (getSynchronizedDate() == that.getSynchronizedDate())
        && (getSynchronizedTime() == that.getSynchronizedTime())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSynchronizedDate(), getSynchronizedTime());
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
