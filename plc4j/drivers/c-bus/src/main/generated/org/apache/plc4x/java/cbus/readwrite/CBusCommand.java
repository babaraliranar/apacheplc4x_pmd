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

public abstract class CBusCommand implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final CBusHeader header;

  // Arguments.
  protected final CBusOptions cBusOptions;

  public CBusCommand(CBusHeader header, CBusOptions cBusOptions) {
    super();
    this.header = header;
    this.cBusOptions = cBusOptions;
  }

  public CBusHeader getHeader() {
    return header;
  }

  public boolean getIsDeviceManagement() {
    return (boolean) (getHeader().getDp());
  }

  public DestinationAddressType getDestinationAddressType() {
    return (DestinationAddressType) (getHeader().getDestinationAddressType());
  }

  protected abstract void serializeCBusCommandChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CBusCommand");

    // Simple Field (header)
    writeSimpleField("header", header, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isDeviceManagement = getIsDeviceManagement();
    writeBuffer.writeVirtual("isDeviceManagement", isDeviceManagement);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    DestinationAddressType destinationAddressType = getDestinationAddressType();
    writeBuffer.writeVirtual("destinationAddressType", destinationAddressType);

    // Switch field (Serialize the sub-type)
    serializeCBusCommandChild(writeBuffer);

    writeBuffer.popContext("CBusCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CBusCommand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static CBusCommand staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    CBusOptions cBusOptions;
    if (args[0] instanceof CBusOptions) {
      cBusOptions = (CBusOptions) args[0];
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type CBusOptions or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, cBusOptions);
  }

  public static CBusCommand staticParse(ReadBuffer readBuffer, CBusOptions cBusOptions)
      throws ParseException {
    readBuffer.pullContext("CBusCommand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    CBusHeader header =
        readSimpleField(
            "header",
            new DataReaderComplexDefault<>(() -> CBusHeader.staticParse(readBuffer), readBuffer));
    boolean isDeviceManagement =
        readVirtualField("isDeviceManagement", boolean.class, header.getDp());
    DestinationAddressType destinationAddressType =
        readVirtualField(
            "destinationAddressType",
            DestinationAddressType.class,
            header.getDestinationAddressType());

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    CBusCommandBuilder builder = null;
    if (true && EvaluationHelper.equals(isDeviceManagement, (boolean) true)) {
      builder = CBusCommandDeviceManagement.staticParseCBusCommandBuilder(readBuffer, cBusOptions);
    } else if (EvaluationHelper.equals(
        destinationAddressType, DestinationAddressType.PointToPointToMultiPoint)) {
      builder =
          CBusCommandPointToPointToMultiPoint.staticParseCBusCommandBuilder(
              readBuffer, cBusOptions);
    } else if (EvaluationHelper.equals(
        destinationAddressType, DestinationAddressType.PointToMultiPoint)) {
      builder = CBusCommandPointToMultiPoint.staticParseCBusCommandBuilder(readBuffer, cBusOptions);
    } else if (EvaluationHelper.equals(
        destinationAddressType, DestinationAddressType.PointToPoint)) {
      builder = CBusCommandPointToPoint.staticParseCBusCommandBuilder(readBuffer, cBusOptions);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "destinationAddressType="
              + destinationAddressType
              + " "
              + "isDeviceManagement="
              + isDeviceManagement
              + "]");
    }

    readBuffer.closeContext("CBusCommand");
    // Create the instance
    CBusCommand _cBusCommand = builder.build(header, cBusOptions);
    return _cBusCommand;
  }

  public interface CBusCommandBuilder {
    CBusCommand build(CBusHeader header, CBusOptions cBusOptions);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusCommand)) {
      return false;
    }
    CBusCommand that = (CBusCommand) o;
    return (getHeader() == that.getHeader()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader());
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
