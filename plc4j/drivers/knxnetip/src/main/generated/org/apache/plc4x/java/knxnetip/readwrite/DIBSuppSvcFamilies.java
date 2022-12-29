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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class DIBSuppSvcFamilies implements Message {

  // Properties.
  protected final short descriptionType;
  protected final List<ServiceId> serviceIds;

  public DIBSuppSvcFamilies(short descriptionType, List<ServiceId> serviceIds) {
    super();
    this.descriptionType = descriptionType;
    this.serviceIds = serviceIds;
  }

  public short getDescriptionType() {
    return descriptionType;
  }

  public List<ServiceId> getServiceIds() {
    return serviceIds;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DIBSuppSvcFamilies");

    // Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short structureLength = (short) (getLengthInBytes());
    writeImplicitField("structureLength", structureLength, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (descriptionType)
    writeSimpleField("descriptionType", descriptionType, writeUnsignedShort(writeBuffer, 8));

    // Array Field (serviceIds)
    writeComplexTypeArrayField("serviceIds", serviceIds, writeBuffer);

    writeBuffer.popContext("DIBSuppSvcFamilies");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DIBSuppSvcFamilies _value = this;

    // Implicit Field (structureLength)
    lengthInBits += 8;

    // Simple field (descriptionType)
    lengthInBits += 8;

    // Array field
    if (serviceIds != null) {
      for (Message element : serviceIds) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static DIBSuppSvcFamilies staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DIBSuppSvcFamilies staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DIBSuppSvcFamilies");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short structureLength = readImplicitField("structureLength", readUnsignedShort(readBuffer, 8));

    short descriptionType = readSimpleField("descriptionType", readUnsignedShort(readBuffer, 8));

    List<ServiceId> serviceIds =
        readLengthArrayField(
            "serviceIds",
            new DataReaderComplexDefault<>(() -> ServiceId.staticParse(readBuffer), readBuffer),
            (structureLength) - (2));

    readBuffer.closeContext("DIBSuppSvcFamilies");
    // Create the instance
    DIBSuppSvcFamilies _dIBSuppSvcFamilies;
    _dIBSuppSvcFamilies = new DIBSuppSvcFamilies(descriptionType, serviceIds);
    return _dIBSuppSvcFamilies;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DIBSuppSvcFamilies)) {
      return false;
    }
    DIBSuppSvcFamilies that = (DIBSuppSvcFamilies) o;
    return (getDescriptionType() == that.getDescriptionType())
        && (getServiceIds() == that.getServiceIds())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getDescriptionType(), getServiceIds());
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
