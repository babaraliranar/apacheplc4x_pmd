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

public class BACnetAccessRuleLocationSpecifierTagged implements Message {

  // Properties.
  protected final BACnetTagHeader header;
  protected final BACnetAccessRuleLocationSpecifier value;

  // Arguments.
  protected final Short tagNumber;
  protected final TagClass tagClass;

  public BACnetAccessRuleLocationSpecifierTagged(
      BACnetTagHeader header,
      BACnetAccessRuleLocationSpecifier value,
      Short tagNumber,
      TagClass tagClass) {
    super();
    this.header = header;
    this.value = value;
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public BACnetAccessRuleLocationSpecifier getValue() {
    return value;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetAccessRuleLocationSpecifierTagged");

    // Simple Field (header)
    writeSimpleField("header", header, new DataWriterComplexDefault<>(writeBuffer));

    // Manual Field (value)
    writeManualField(
        "value",
        () ->
            org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.writeEnumGeneric(
                writeBuffer, value),
        writeBuffer);

    writeBuffer.popContext("BACnetAccessRuleLocationSpecifierTagged");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetAccessRuleLocationSpecifierTagged _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // Manual Field (value)
    lengthInBits += (header.getActualLength()) * (8);

    return lengthInBits;
  }

  public static BACnetAccessRuleLocationSpecifierTagged staticParse(
      ReadBuffer readBuffer, Short tagNumber, TagClass tagClass) throws ParseException {
    readBuffer.pullContext("BACnetAccessRuleLocationSpecifierTagged");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getTagClass()) == (tagClass))) {
      throw new ParseValidationException("tag class doesn't match");
    }
    // Validation
    if (!((((header.getTagClass()) == (TagClass.APPLICATION_TAGS)))
        || (((header.getActualTagNumber()) == (tagNumber))))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }

    BACnetAccessRuleLocationSpecifier value =
        readManualField(
            "value",
            readBuffer,
            () ->
                (BACnetAccessRuleLocationSpecifier)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .readEnumGenericFailing(
                            readBuffer,
                            header.getActualLength(),
                            BACnetAccessRuleLocationSpecifier.SPECIFIED)));

    readBuffer.closeContext("BACnetAccessRuleLocationSpecifierTagged");
    // Create the instance
    BACnetAccessRuleLocationSpecifierTagged _bACnetAccessRuleLocationSpecifierTagged;
    _bACnetAccessRuleLocationSpecifierTagged =
        new BACnetAccessRuleLocationSpecifierTagged(header, value, tagNumber, tagClass);
    return _bACnetAccessRuleLocationSpecifierTagged;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetAccessRuleLocationSpecifierTagged)) {
      return false;
    }
    BACnetAccessRuleLocationSpecifierTagged that = (BACnetAccessRuleLocationSpecifierTagged) o;
    return (getHeader() == that.getHeader()) && (getValue() == that.getValue()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader(), getValue());
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
