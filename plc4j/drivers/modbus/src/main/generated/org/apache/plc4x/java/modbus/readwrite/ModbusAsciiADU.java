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
package org.apache.plc4x.java.modbus.readwrite;

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

public class ModbusAsciiADU extends ModbusADU implements Message {

  // Accessors for discriminator values.
  public DriverType getDriverType() {
    return DriverType.MODBUS_ASCII;
  }

  // Properties.
  protected final short address;
  protected final ModbusPDU pdu;

  // Arguments.
  protected final Boolean response;

  public ModbusAsciiADU(short address, ModbusPDU pdu, Boolean response) {
    super(response);
    this.address = address;
    this.pdu = pdu;
    this.response = response;
  }

  public short getAddress() {
    return address;
  }

  public ModbusPDU getPdu() {
    return pdu;
  }

  @Override
  protected void serializeModbusADUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModbusAsciiADU");

    // Simple Field (address)
    writeSimpleField(
        "address",
        address,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (pdu)
    writeSimpleField(
        "pdu",
        pdu,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Checksum Field (checksum) (Calculated)
    writeChecksumField(
        "crc",
        (short)
            (org.apache.plc4x.java.modbus.readwrite.utils.StaticHelper.asciiLrcCheck(address, pdu)),
        writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("ModbusAsciiADU");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusAsciiADU _value = this;

    // Simple field (address)
    lengthInBits += 8;

    // Simple field (pdu)
    lengthInBits += pdu.getLengthInBits();

    // Checksum Field (checksum)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ModbusAsciiADUBuilder staticParseBuilder(
      ReadBuffer readBuffer, DriverType driverType, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusAsciiADU");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short address =
        readSimpleField(
            "address",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    ModbusPDU pdu =
        readSimpleField(
            "pdu",
            new DataReaderComplexDefault<>(
                () -> ModbusPDU.staticParse(readBuffer, (boolean) (response)), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short crc =
        readChecksumField(
            "crc",
            readUnsignedShort(readBuffer, 8),
            (short)
                (org.apache.plc4x.java.modbus.readwrite.utils.StaticHelper.asciiLrcCheck(
                    address, pdu)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("ModbusAsciiADU");
    // Create the instance
    return new ModbusAsciiADUBuilder(address, pdu, response);
  }

  public static class ModbusAsciiADUBuilder implements ModbusADU.ModbusADUBuilder {
    private final short address;
    private final ModbusPDU pdu;
    private final Boolean response;

    public ModbusAsciiADUBuilder(short address, ModbusPDU pdu, Boolean response) {

      this.address = address;
      this.pdu = pdu;
      this.response = response;
    }

    public ModbusAsciiADU build(Boolean response) {

      ModbusAsciiADU modbusAsciiADU = new ModbusAsciiADU(address, pdu, response);
      return modbusAsciiADU;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusAsciiADU)) {
      return false;
    }
    ModbusAsciiADU that = (ModbusAsciiADU) o;
    return (getAddress() == that.getAddress())
        && (getPdu() == that.getPdu())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAddress(), getPdu());
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
