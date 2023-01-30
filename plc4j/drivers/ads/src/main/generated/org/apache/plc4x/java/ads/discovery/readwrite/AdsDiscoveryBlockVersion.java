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
package org.apache.plc4x.java.ads.discovery.readwrite;

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

public class AdsDiscoveryBlockVersion extends AdsDiscoveryBlock implements Message {

  // Accessors for discriminator values.
  public AdsDiscoveryBlockType getBlockType() {
    return AdsDiscoveryBlockType.VERSION;
  }

  // Properties.
  protected final byte[] versionData;

  public AdsDiscoveryBlockVersion(byte[] versionData) {
    super();
    this.versionData = versionData;
  }

  public byte[] getVersionData() {
    return versionData;
  }

  @Override
  protected void serializeAdsDiscoveryBlockChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AdsDiscoveryBlockVersion");

    // Implicit Field (versionDataLen) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int versionDataLen = (int) (COUNT(getVersionData()));
    writeImplicitField("versionDataLen", versionDataLen, writeUnsignedInt(writeBuffer, 16));

    // Array Field (versionData)
    writeByteArrayField("versionData", versionData, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("AdsDiscoveryBlockVersion");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsDiscoveryBlockVersion _value = this;

    // Implicit Field (versionDataLen)
    lengthInBits += 16;

    // Array field
    if (versionData != null) {
      lengthInBits += 8 * versionData.length;
    }

    return lengthInBits;
  }

  public static AdsDiscoveryBlockBuilder staticParseAdsDiscoveryBlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsDiscoveryBlockVersion");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int versionDataLen = readImplicitField("versionDataLen", readUnsignedInt(readBuffer, 16));

    byte[] versionData = readBuffer.readByteArray("versionData", Math.toIntExact(versionDataLen));

    readBuffer.closeContext("AdsDiscoveryBlockVersion");
    // Create the instance
    return new AdsDiscoveryBlockVersionBuilderImpl(versionData);
  }

  public static class AdsDiscoveryBlockVersionBuilderImpl
      implements AdsDiscoveryBlock.AdsDiscoveryBlockBuilder {
    private final byte[] versionData;

    public AdsDiscoveryBlockVersionBuilderImpl(byte[] versionData) {
      this.versionData = versionData;
    }

    public AdsDiscoveryBlockVersion build() {
      AdsDiscoveryBlockVersion adsDiscoveryBlockVersion = new AdsDiscoveryBlockVersion(versionData);
      return adsDiscoveryBlockVersion;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsDiscoveryBlockVersion)) {
      return false;
    }
    AdsDiscoveryBlockVersion that = (AdsDiscoveryBlockVersion) o;
    return (getVersionData() == that.getVersionData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getVersionData());
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
