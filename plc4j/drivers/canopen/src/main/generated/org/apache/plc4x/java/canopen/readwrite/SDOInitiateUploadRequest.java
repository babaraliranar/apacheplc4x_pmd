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
package org.apache.plc4x.java.canopen.readwrite;

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

public class SDOInitiateUploadRequest extends SDORequest implements Message {

  // Accessors for discriminator values.
  public SDORequestCommand getCommand() {
    return SDORequestCommand.INITIATE_UPLOAD;
  }

  // Properties.
  protected final IndexAddress address;

  public SDOInitiateUploadRequest(IndexAddress address) {
    super();
    this.address = address;
  }

  public IndexAddress getAddress() {
    return address;
  }

  @Override
  protected void serializeSDORequestChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SDOInitiateUploadRequest");

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 5));

    // Simple Field (address)
    writeSimpleField("address", address, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (int) 0x00, writeSignedInt(writeBuffer, 32));

    writeBuffer.popContext("SDOInitiateUploadRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SDOInitiateUploadRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 5;

    // Simple field (address)
    lengthInBits += address.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static SDORequestBuilder staticParseSDORequestBuilder(
      ReadBuffer readBuffer, SDORequestCommand command) throws ParseException {
    readBuffer.pullContext("SDOInitiateUploadRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 5), (short) 0x00);

    IndexAddress address =
        readSimpleField(
            "address",
            new DataReaderComplexDefault<>(() -> IndexAddress.staticParse(readBuffer), readBuffer));

    Integer reservedField1 =
        readReservedField("reserved", readSignedInt(readBuffer, 32), (int) 0x00);

    readBuffer.closeContext("SDOInitiateUploadRequest");
    // Create the instance
    return new SDOInitiateUploadRequestBuilderImpl(address);
  }

  public static class SDOInitiateUploadRequestBuilderImpl implements SDORequest.SDORequestBuilder {
    private final IndexAddress address;

    public SDOInitiateUploadRequestBuilderImpl(IndexAddress address) {
      this.address = address;
    }

    public SDOInitiateUploadRequest build() {
      SDOInitiateUploadRequest sDOInitiateUploadRequest = new SDOInitiateUploadRequest(address);
      return sDOInitiateUploadRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SDOInitiateUploadRequest)) {
      return false;
    }
    SDOInitiateUploadRequest that = (SDOInitiateUploadRequest) o;
    return (getAddress() == that.getAddress()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAddress());
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
