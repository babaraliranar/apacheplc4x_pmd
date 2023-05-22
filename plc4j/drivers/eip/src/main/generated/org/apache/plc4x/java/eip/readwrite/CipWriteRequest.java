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
package org.apache.plc4x.java.eip.readwrite;

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

public class CipWriteRequest extends CipService implements Message {

  // Accessors for discriminator values.
  public Short getService() {
    return (short) 0x4D;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  public Boolean getConnected() {
    return false;
  }

  // Properties.
  protected final byte[] tag;
  protected final CIPDataTypeCode dataType;
  protected final int elementNb;
  protected final byte[] data;

  public CipWriteRequest(byte[] tag, CIPDataTypeCode dataType, int elementNb, byte[] data) {
    super();
    this.tag = tag;
    this.dataType = dataType;
    this.elementNb = elementNb;
    this.data = data;
  }

  public byte[] getTag() {
    return tag;
  }

  public CIPDataTypeCode getDataType() {
    return dataType;
  }

  public int getElementNb() {
    return elementNb;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CipWriteRequest");

    // Implicit Field (requestPathSize) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short requestPathSize = (short) ((COUNT(getTag())) / (2));
    writeImplicitField("requestPathSize", requestPathSize, writeUnsignedShort(writeBuffer, 8));

    // Array Field (tag)
    writeByteArrayField("tag", tag, writeByteArray(writeBuffer, 8));

    // Simple Field (dataType)
    writeSimpleEnumField(
        "dataType",
        "CIPDataTypeCode",
        dataType,
        new DataWriterEnumDefault<>(
            CIPDataTypeCode::getValue, CIPDataTypeCode::name, writeUnsignedInt(writeBuffer, 16)));

    // Simple Field (elementNb)
    writeSimpleField("elementNb", elementNb, writeUnsignedInt(writeBuffer, 16));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("CipWriteRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipWriteRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (requestPathSize)
    lengthInBits += 8;

    // Array field
    if (tag != null) {
      lengthInBits += 8 * tag.length;
    }

    // Simple field (dataType)
    lengthInBits += 16;

    // Simple field (elementNb)
    lengthInBits += 16;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("CipWriteRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short requestPathSize = readImplicitField("requestPathSize", readUnsignedShort(readBuffer, 8));

    byte[] tag = readBuffer.readByteArray("tag", Math.toIntExact((requestPathSize) * (2)));

    CIPDataTypeCode dataType =
        readEnumField(
            "dataType",
            "CIPDataTypeCode",
            new DataReaderEnumDefault<>(
                CIPDataTypeCode::enumForValue, readUnsignedInt(readBuffer, 16)));

    int elementNb = readSimpleField("elementNb", readUnsignedInt(readBuffer, 16));

    byte[] data =
        readBuffer.readByteArray("data", Math.toIntExact((dataType.getSize()) * (elementNb)));

    readBuffer.closeContext("CipWriteRequest");
    // Create the instance
    return new CipWriteRequestBuilderImpl(tag, dataType, elementNb, data);
  }

  public static class CipWriteRequestBuilderImpl implements CipService.CipServiceBuilder {
    private final byte[] tag;
    private final CIPDataTypeCode dataType;
    private final int elementNb;
    private final byte[] data;

    public CipWriteRequestBuilderImpl(
        byte[] tag, CIPDataTypeCode dataType, int elementNb, byte[] data) {
      this.tag = tag;
      this.dataType = dataType;
      this.elementNb = elementNb;
      this.data = data;
    }

    public CipWriteRequest build() {
      CipWriteRequest cipWriteRequest = new CipWriteRequest(tag, dataType, elementNb, data);
      return cipWriteRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipWriteRequest)) {
      return false;
    }
    CipWriteRequest that = (CipWriteRequest) o;
    return (getTag() == that.getTag())
        && (getDataType() == that.getDataType())
        && (getElementNb() == that.getElementNb())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTag(), getDataType(), getElementNb(), getData());
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
