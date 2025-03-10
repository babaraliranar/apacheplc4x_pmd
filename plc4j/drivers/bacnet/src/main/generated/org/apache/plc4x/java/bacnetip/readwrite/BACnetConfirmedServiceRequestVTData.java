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

public class BACnetConfirmedServiceRequestVTData extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.VT_DATA;
  }

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger vtSessionIdentifier;
  protected final BACnetApplicationTagOctetString vtNewData;
  protected final BACnetApplicationTagUnsignedInteger vtDataFlag;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestVTData(
      BACnetApplicationTagUnsignedInteger vtSessionIdentifier,
      BACnetApplicationTagOctetString vtNewData,
      BACnetApplicationTagUnsignedInteger vtDataFlag,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.vtSessionIdentifier = vtSessionIdentifier;
    this.vtNewData = vtNewData;
    this.vtDataFlag = vtDataFlag;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetApplicationTagUnsignedInteger getVtSessionIdentifier() {
    return vtSessionIdentifier;
  }

  public BACnetApplicationTagOctetString getVtNewData() {
    return vtNewData;
  }

  public BACnetApplicationTagUnsignedInteger getVtDataFlag() {
    return vtDataFlag;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestVTData");

    // Simple Field (vtSessionIdentifier)
    writeSimpleField("vtSessionIdentifier", vtSessionIdentifier, writeComplex(writeBuffer));

    // Simple Field (vtNewData)
    writeSimpleField("vtNewData", vtNewData, writeComplex(writeBuffer));

    // Simple Field (vtDataFlag)
    writeSimpleField("vtDataFlag", vtDataFlag, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestVTData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestVTData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (vtSessionIdentifier)
    lengthInBits += vtSessionIdentifier.getLengthInBits();

    // Simple field (vtNewData)
    lengthInBits += vtNewData.getLengthInBits();

    // Simple field (vtDataFlag)
    lengthInBits += vtDataFlag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestVTData");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagUnsignedInteger vtSessionIdentifier =
        readSimpleField(
            "vtSessionIdentifier",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagOctetString vtNewData =
        readSimpleField(
            "vtNewData",
            readComplex(
                () ->
                    (BACnetApplicationTagOctetString) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagUnsignedInteger vtDataFlag =
        readSimpleField(
            "vtDataFlag",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestVTData");
    // Create the instance
    return new BACnetConfirmedServiceRequestVTDataBuilderImpl(
        vtSessionIdentifier, vtNewData, vtDataFlag, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestVTDataBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetApplicationTagUnsignedInteger vtSessionIdentifier;
    private final BACnetApplicationTagOctetString vtNewData;
    private final BACnetApplicationTagUnsignedInteger vtDataFlag;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestVTDataBuilderImpl(
        BACnetApplicationTagUnsignedInteger vtSessionIdentifier,
        BACnetApplicationTagOctetString vtNewData,
        BACnetApplicationTagUnsignedInteger vtDataFlag,
        Long serviceRequestLength) {
      this.vtSessionIdentifier = vtSessionIdentifier;
      this.vtNewData = vtNewData;
      this.vtDataFlag = vtDataFlag;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestVTData build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestVTData bACnetConfirmedServiceRequestVTData =
          new BACnetConfirmedServiceRequestVTData(
              vtSessionIdentifier, vtNewData, vtDataFlag, serviceRequestLength);
      return bACnetConfirmedServiceRequestVTData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestVTData)) {
      return false;
    }
    BACnetConfirmedServiceRequestVTData that = (BACnetConfirmedServiceRequestVTData) o;
    return (getVtSessionIdentifier() == that.getVtSessionIdentifier())
        && (getVtNewData() == that.getVtNewData())
        && (getVtDataFlag() == that.getVtDataFlag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getVtSessionIdentifier(), getVtNewData(), getVtDataFlag());
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
