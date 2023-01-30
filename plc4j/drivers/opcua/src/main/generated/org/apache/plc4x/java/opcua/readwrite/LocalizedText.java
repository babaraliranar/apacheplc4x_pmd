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
package org.apache.plc4x.java.opcua.readwrite;

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

public class LocalizedText implements Message {

  // Properties.
  protected final boolean textSpecified;
  protected final boolean localeSpecified;
  protected final PascalString locale;
  protected final PascalString text;

  public LocalizedText(
      boolean textSpecified, boolean localeSpecified, PascalString locale, PascalString text) {
    super();
    this.textSpecified = textSpecified;
    this.localeSpecified = localeSpecified;
    this.locale = locale;
    this.text = text;
  }

  public boolean getTextSpecified() {
    return textSpecified;
  }

  public boolean getLocaleSpecified() {
    return localeSpecified;
  }

  public PascalString getLocale() {
    return locale;
  }

  public PascalString getText() {
    return text;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("LocalizedText");

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 6));

    // Simple Field (textSpecified)
    writeSimpleField("textSpecified", textSpecified, writeBoolean(writeBuffer));

    // Simple Field (localeSpecified)
    writeSimpleField("localeSpecified", localeSpecified, writeBoolean(writeBuffer));

    // Optional Field (locale) (Can be skipped, if the value is null)
    writeOptionalField("locale", locale, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (text) (Can be skipped, if the value is null)
    writeOptionalField("text", text, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("LocalizedText");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LocalizedText _value = this;

    // Reserved Field (reserved)
    lengthInBits += 6;

    // Simple field (textSpecified)
    lengthInBits += 1;

    // Simple field (localeSpecified)
    lengthInBits += 1;

    // Optional Field (locale)
    if (locale != null) {
      lengthInBits += locale.getLengthInBits();
    }

    // Optional Field (text)
    if (text != null) {
      lengthInBits += text.getLengthInBits();
    }

    return lengthInBits;
  }

  public static LocalizedText staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static LocalizedText staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("LocalizedText");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 6), (short) 0x00);

    boolean textSpecified = readSimpleField("textSpecified", readBoolean(readBuffer));

    boolean localeSpecified = readSimpleField("localeSpecified", readBoolean(readBuffer));

    PascalString locale =
        readOptionalField(
            "locale",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            localeSpecified);

    PascalString text =
        readOptionalField(
            "text",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            textSpecified);

    readBuffer.closeContext("LocalizedText");
    // Create the instance
    LocalizedText _localizedText;
    _localizedText = new LocalizedText(textSpecified, localeSpecified, locale, text);
    return _localizedText;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LocalizedText)) {
      return false;
    }
    LocalizedText that = (LocalizedText) o;
    return (getTextSpecified() == that.getTextSpecified())
        && (getLocaleSpecified() == that.getLocaleSpecified())
        && (getLocale() == that.getLocale())
        && (getText() == that.getText())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTextSpecified(), getLocaleSpecified(), getLocale(), getText());
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
