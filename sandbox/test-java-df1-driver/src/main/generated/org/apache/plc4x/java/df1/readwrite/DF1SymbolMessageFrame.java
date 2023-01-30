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
package org.apache.plc4x.java.df1.readwrite;

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

public class DF1SymbolMessageFrame extends DF1Symbol implements Message {

  // Accessors for discriminator values.
  public Short getSymbolType() {
    return (short) 0x02;
  }

  // Constant values.
  public static final Short MESSAGEEND = 0x10;
  public static final Short ENDTRANSACTION = 0x03;

  // Properties.
  protected final short destinationAddress;
  protected final short sourceAddress;
  protected final DF1Command command;

  public DF1SymbolMessageFrame(short destinationAddress, short sourceAddress, DF1Command command) {
    super();
    this.destinationAddress = destinationAddress;
    this.sourceAddress = sourceAddress;
    this.command = command;
  }

  public short getDestinationAddress() {
    return destinationAddress;
  }

  public short getSourceAddress() {
    return sourceAddress;
  }

  public DF1Command getCommand() {
    return command;
  }

  public short getMessageEnd() {
    return MESSAGEEND;
  }

  public short getEndTransaction() {
    return ENDTRANSACTION;
  }

  @Override
  protected void serializeDF1SymbolChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DF1SymbolMessageFrame");

    // Simple Field (destinationAddress)
    writeSimpleField(
        "destinationAddress",
        destinationAddress,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (sourceAddress)
    writeSimpleField(
        "sourceAddress",
        sourceAddress,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (command)
    writeSimpleField(
        "command",
        command,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (messageEnd)
    writeConstField(
        "messageEnd",
        MESSAGEEND,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (endTransaction)
    writeConstField(
        "endTransaction",
        ENDTRANSACTION,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Checksum Field (checksum) (Calculated)
    writeChecksumField(
        "crc",
        (int)
            (org.apache.plc4x.java.df1.readwrite.utils.StaticHelper.crcCheck(
                destinationAddress, sourceAddress, command)),
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("DF1SymbolMessageFrame");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DF1SymbolMessageFrame _value = this;

    // Simple field (destinationAddress)
    lengthInBits += 8;

    // Simple field (sourceAddress)
    lengthInBits += 8;

    // Simple field (command)
    lengthInBits += command.getLengthInBits();

    // Const Field (messageEnd)
    lengthInBits += 8;

    // Const Field (endTransaction)
    lengthInBits += 8;

    // Checksum Field (checksum)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static DF1SymbolBuilder staticParseDF1SymbolBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("DF1SymbolMessageFrame");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short destinationAddress =
        readSimpleField(
            "destinationAddress",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short sourceAddress =
        readSimpleField(
            "sourceAddress",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    DF1Command command =
        readSimpleField(
            "command",
            new DataReaderComplexDefault<>(() -> DF1Command.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short messageEnd =
        readConstField(
            "messageEnd",
            readUnsignedShort(readBuffer, 8),
            DF1SymbolMessageFrame.MESSAGEEND,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short endTransaction =
        readConstField(
            "endTransaction",
            readUnsignedShort(readBuffer, 8),
            DF1SymbolMessageFrame.ENDTRANSACTION,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int crc =
        readChecksumField(
            "crc",
            readUnsignedInt(readBuffer, 16),
            (int)
                (org.apache.plc4x.java.df1.readwrite.utils.StaticHelper.crcCheck(
                    destinationAddress, sourceAddress, command)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("DF1SymbolMessageFrame");
    // Create the instance
    return new DF1SymbolMessageFrameBuilderImpl(destinationAddress, sourceAddress, command);
  }

  public static class DF1SymbolMessageFrameBuilderImpl implements DF1Symbol.DF1SymbolBuilder {
    private final short destinationAddress;
    private final short sourceAddress;
    private final DF1Command command;

    public DF1SymbolMessageFrameBuilderImpl(
        short destinationAddress, short sourceAddress, DF1Command command) {
      this.destinationAddress = destinationAddress;
      this.sourceAddress = sourceAddress;
      this.command = command;
    }

    public DF1SymbolMessageFrame build() {
      DF1SymbolMessageFrame dF1SymbolMessageFrame =
          new DF1SymbolMessageFrame(destinationAddress, sourceAddress, command);
      return dF1SymbolMessageFrame;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DF1SymbolMessageFrame)) {
      return false;
    }
    DF1SymbolMessageFrame that = (DF1SymbolMessageFrame) o;
    return (getDestinationAddress() == that.getDestinationAddress())
        && (getSourceAddress() == that.getSourceAddress())
        && (getCommand() == that.getCommand())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getDestinationAddress(), getSourceAddress(), getCommand());
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
