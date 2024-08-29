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

public abstract class BACnetUnconfirmedServiceRequest implements Message {

  // Abstract accessors for discriminator values.
  public abstract BACnetUnconfirmedServiceChoice getServiceChoice();

  // Arguments.
  protected final Integer serviceRequestLength;

  public BACnetUnconfirmedServiceRequest(Integer serviceRequestLength) {
    super();
    this.serviceRequestLength = serviceRequestLength;
  }

  protected abstract void serializeBACnetUnconfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetUnconfirmedServiceRequest");

    // Discriminator Field (serviceChoice) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "serviceChoice",
        "BACnetUnconfirmedServiceChoice",
        getServiceChoice(),
        new DataWriterEnumDefault<>(
            BACnetUnconfirmedServiceChoice::getValue,
            BACnetUnconfirmedServiceChoice::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Switch field (Serialize the sub-type)
    serializeBACnetUnconfirmedServiceRequestChild(writeBuffer);

    writeBuffer.popContext("BACnetUnconfirmedServiceRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetUnconfirmedServiceRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (serviceChoice)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static BACnetUnconfirmedServiceRequest staticParse(
      ReadBuffer readBuffer, Integer serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetUnconfirmedServiceRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetUnconfirmedServiceChoice serviceChoice =
        readDiscriminatorEnumField(
            "serviceChoice",
            "BACnetUnconfirmedServiceChoice",
            new DataReaderEnumDefault<>(
                BACnetUnconfirmedServiceChoice::enumForValue, readUnsignedShort(readBuffer, 8)));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetUnconfirmedServiceRequestBuilder builder = null;
    if (EvaluationHelper.equals(serviceChoice, BACnetUnconfirmedServiceChoice.I_AM)) {
      builder =
          BACnetUnconfirmedServiceRequestIAm.staticParseBACnetUnconfirmedServiceRequestBuilder(
              readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetUnconfirmedServiceChoice.I_HAVE)) {
      builder =
          BACnetUnconfirmedServiceRequestIHave.staticParseBACnetUnconfirmedServiceRequestBuilder(
              readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.UNCONFIRMED_COV_NOTIFICATION)) {
      builder =
          BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.UNCONFIRMED_EVENT_NOTIFICATION)) {
      builder =
          BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.UNCONFIRMED_PRIVATE_TRANSFER)) {
      builder =
          BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.UNCONFIRMED_TEXT_MESSAGE)) {
      builder =
          BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.TIME_SYNCHRONIZATION)) {
      builder =
          BACnetUnconfirmedServiceRequestTimeSynchronization
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetUnconfirmedServiceChoice.WHO_HAS)) {
      builder =
          BACnetUnconfirmedServiceRequestWhoHas.staticParseBACnetUnconfirmedServiceRequestBuilder(
              readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetUnconfirmedServiceChoice.WHO_IS)) {
      builder =
          BACnetUnconfirmedServiceRequestWhoIs.staticParseBACnetUnconfirmedServiceRequestBuilder(
              readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.UTC_TIME_SYNCHRONIZATION)) {
      builder =
          BACnetUnconfirmedServiceRequestUTCTimeSynchronization
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetUnconfirmedServiceChoice.WRITE_GROUP)) {
      builder =
          BACnetUnconfirmedServiceRequestWriteGroup
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetUnconfirmedServiceChoice.UNCONFIRMED_COV_NOTIFICATION_MULTIPLE)) {
      builder =
          BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
              .staticParseBACnetUnconfirmedServiceRequestBuilder(readBuffer, serviceRequestLength);
    } else if (true) {
      builder =
          BACnetUnconfirmedServiceRequestUnknown.staticParseBACnetUnconfirmedServiceRequestBuilder(
              readBuffer, serviceRequestLength);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "serviceChoice="
              + serviceChoice
              + "]");
    }

    readBuffer.closeContext("BACnetUnconfirmedServiceRequest");
    // Create the instance
    BACnetUnconfirmedServiceRequest _bACnetUnconfirmedServiceRequest =
        builder.build(serviceRequestLength);

    return _bACnetUnconfirmedServiceRequest;
  }

  public interface BACnetUnconfirmedServiceRequestBuilder {
    BACnetUnconfirmedServiceRequest build(Integer serviceRequestLength);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetUnconfirmedServiceRequest)) {
      return false;
    }
    BACnetUnconfirmedServiceRequest that = (BACnetUnconfirmedServiceRequest) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
