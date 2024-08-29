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

public abstract class CBusPointToPointToMultiPointCommand implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BridgeAddress bridgeAddress;
  protected final NetworkRoute networkRoute;
  protected final byte peekedApplication;

  // Arguments.
  protected final CBusOptions cBusOptions;

  public CBusPointToPointToMultiPointCommand(
      BridgeAddress bridgeAddress,
      NetworkRoute networkRoute,
      byte peekedApplication,
      CBusOptions cBusOptions) {
    super();
    this.bridgeAddress = bridgeAddress;
    this.networkRoute = networkRoute;
    this.peekedApplication = peekedApplication;
    this.cBusOptions = cBusOptions;
  }

  public BridgeAddress getBridgeAddress() {
    return bridgeAddress;
  }

  public NetworkRoute getNetworkRoute() {
    return networkRoute;
  }

  public byte getPeekedApplication() {
    return peekedApplication;
  }

  protected abstract void serializeCBusPointToPointToMultiPointCommandChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CBusPointToPointToMultiPointCommand");

    // Simple Field (bridgeAddress)
    writeSimpleField("bridgeAddress", bridgeAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (networkRoute)
    writeSimpleField("networkRoute", networkRoute, new DataWriterComplexDefault<>(writeBuffer));

    // Switch field (Serialize the sub-type)
    serializeCBusPointToPointToMultiPointCommandChild(writeBuffer);

    writeBuffer.popContext("CBusPointToPointToMultiPointCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CBusPointToPointToMultiPointCommand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (bridgeAddress)
    lengthInBits += bridgeAddress.getLengthInBits();

    // Simple field (networkRoute)
    lengthInBits += networkRoute.getLengthInBits();

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static CBusPointToPointToMultiPointCommand staticParse(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("CBusPointToPointToMultiPointCommand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BridgeAddress bridgeAddress =
        readSimpleField(
            "bridgeAddress",
            new DataReaderComplexDefault<>(
                () -> BridgeAddress.staticParse(readBuffer), readBuffer));

    NetworkRoute networkRoute =
        readSimpleField(
            "networkRoute",
            new DataReaderComplexDefault<>(() -> NetworkRoute.staticParse(readBuffer), readBuffer));

    byte peekedApplication = readPeekField("peekedApplication", readByte(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    CBusPointToPointToMultiPointCommandBuilder builder = null;
    if (EvaluationHelper.equals(peekedApplication, (byte) 0xFF)) {
      builder =
          CBusPointToPointToMultiPointCommandStatus
              .staticParseCBusPointToPointToMultiPointCommandBuilder(readBuffer, cBusOptions);
    } else if (true) {
      builder =
          CBusPointToPointToMultiPointCommandNormal
              .staticParseCBusPointToPointToMultiPointCommandBuilder(readBuffer, cBusOptions);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedApplication="
              + peekedApplication
              + "]");
    }

    readBuffer.closeContext("CBusPointToPointToMultiPointCommand");
    // Create the instance
    CBusPointToPointToMultiPointCommand _cBusPointToPointToMultiPointCommand =
        builder.build(bridgeAddress, networkRoute, peekedApplication, cBusOptions);
    return _cBusPointToPointToMultiPointCommand;
  }

  public interface CBusPointToPointToMultiPointCommandBuilder {
    CBusPointToPointToMultiPointCommand build(
        BridgeAddress bridgeAddress,
        NetworkRoute networkRoute,
        byte peekedApplication,
        CBusOptions cBusOptions);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusPointToPointToMultiPointCommand)) {
      return false;
    }
    CBusPointToPointToMultiPointCommand that = (CBusPointToPointToMultiPointCommand) o;
    return (getBridgeAddress() == that.getBridgeAddress())
        && (getNetworkRoute() == that.getNetworkRoute())
        && (getPeekedApplication() == that.getPeekedApplication())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getBridgeAddress(), getNetworkRoute(), getPeekedApplication());
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
