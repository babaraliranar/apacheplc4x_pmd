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
package org.apache.plc4x.java.profinet.readwrite;

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

public class DceRpc_Packet implements Message {

  // Constant values.
  public static final Short VERSION = 0x04;
  public static final Boolean BROADCAST = false;
  public static final Boolean MAYBE = false;
  public static final Boolean FRAGMENT = false;
  public static final Boolean CANCELWASPENDING = false;
  public static final Short SERIALHIGH = 0x00;
  public static final Long INTERFACEVER = 0x00000001L;
  public static final Integer INTERFACEHINT = 0xFFFF;
  public static final Integer ACTIVITYHINT = 0xFFFF;
  public static final Integer FRAGMENTNUM = 0x0000;
  public static final Short AUTHPROTO = 0x00;
  public static final Short SERIALLOW = 0x00;

  // Properties.
  protected final DceRpc_PacketType packetType;
  protected final boolean idempotent;
  protected final boolean noFragmentAcknowledgeRequested;
  protected final boolean lastFragment;
  protected final IntegerEncoding integerEncoding;
  protected final CharacterEncoding characterEncoding;
  protected final FloatingPointEncoding floatingPointEncoding;
  protected final DceRpc_ObjectUuid objectUuid;
  protected final DceRpc_InterfaceUuid interfaceUuid;
  protected final DceRpc_ActivityUuid activityUuid;
  protected final long serverBootTime;
  protected final long sequenceNumber;
  protected final DceRpc_Operation operation;
  protected final PnIoCm_Packet payload;

  // Reserved Fields
  private Boolean reservedField0;
  private Boolean reservedField1;
  private Short reservedField2;
  private Boolean reservedField3;
  private Short reservedField4;

  public DceRpc_Packet(
      DceRpc_PacketType packetType,
      boolean idempotent,
      boolean noFragmentAcknowledgeRequested,
      boolean lastFragment,
      IntegerEncoding integerEncoding,
      CharacterEncoding characterEncoding,
      FloatingPointEncoding floatingPointEncoding,
      DceRpc_ObjectUuid objectUuid,
      DceRpc_InterfaceUuid interfaceUuid,
      DceRpc_ActivityUuid activityUuid,
      long serverBootTime,
      long sequenceNumber,
      DceRpc_Operation operation,
      PnIoCm_Packet payload) {
    super();
    this.packetType = packetType;
    this.idempotent = idempotent;
    this.noFragmentAcknowledgeRequested = noFragmentAcknowledgeRequested;
    this.lastFragment = lastFragment;
    this.integerEncoding = integerEncoding;
    this.characterEncoding = characterEncoding;
    this.floatingPointEncoding = floatingPointEncoding;
    this.objectUuid = objectUuid;
    this.interfaceUuid = interfaceUuid;
    this.activityUuid = activityUuid;
    this.serverBootTime = serverBootTime;
    this.sequenceNumber = sequenceNumber;
    this.operation = operation;
    this.payload = payload;
  }

  public DceRpc_PacketType getPacketType() {
    return packetType;
  }

  public boolean getIdempotent() {
    return idempotent;
  }

  public boolean getNoFragmentAcknowledgeRequested() {
    return noFragmentAcknowledgeRequested;
  }

  public boolean getLastFragment() {
    return lastFragment;
  }

  public IntegerEncoding getIntegerEncoding() {
    return integerEncoding;
  }

  public CharacterEncoding getCharacterEncoding() {
    return characterEncoding;
  }

  public FloatingPointEncoding getFloatingPointEncoding() {
    return floatingPointEncoding;
  }

  public DceRpc_ObjectUuid getObjectUuid() {
    return objectUuid;
  }

  public DceRpc_InterfaceUuid getInterfaceUuid() {
    return interfaceUuid;
  }

  public DceRpc_ActivityUuid getActivityUuid() {
    return activityUuid;
  }

  public long getServerBootTime() {
    return serverBootTime;
  }

  public long getSequenceNumber() {
    return sequenceNumber;
  }

  public DceRpc_Operation getOperation() {
    return operation;
  }

  public PnIoCm_Packet getPayload() {
    return payload;
  }

  public short getVersion() {
    return VERSION;
  }

  public boolean getBroadcast() {
    return BROADCAST;
  }

  public boolean getMaybe() {
    return MAYBE;
  }

  public boolean getFragment() {
    return FRAGMENT;
  }

  public boolean getCancelWasPending() {
    return CANCELWASPENDING;
  }

  public short getSerialHigh() {
    return SERIALHIGH;
  }

  public long getInterfaceVer() {
    return INTERFACEVER;
  }

  public int getInterfaceHint() {
    return INTERFACEHINT;
  }

  public int getActivityHint() {
    return ACTIVITYHINT;
  }

  public int getFragmentNum() {
    return FRAGMENTNUM;
  }

  public short getAuthProto() {
    return AUTHPROTO;
  }

  public short getSerialLow() {
    return SERIALLOW;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DceRpc_Packet");

    // Const Field (version)
    writeConstField(
        "version",
        VERSION,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (packetType)
    writeSimpleEnumField(
        "packetType",
        "DceRpc_PacketType",
        packetType,
        new DataWriterEnumDefault<>(
            DceRpc_PacketType::getValue,
            DceRpc_PacketType::name,
            writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (boolean) false,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (broadcast)
    writeConstField(
        "broadcast",
        BROADCAST,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (idempotent)
    writeSimpleField(
        "idempotent",
        idempotent,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (maybe)
    writeConstField(
        "maybe", MAYBE, writeBoolean(writeBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (noFragmentAcknowledgeRequested)
    writeSimpleField(
        "noFragmentAcknowledgeRequested",
        noFragmentAcknowledgeRequested,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (fragment)
    writeConstField(
        "fragment",
        FRAGMENT,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (lastFragment)
    writeSimpleField(
        "lastFragment",
        lastFragment,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (boolean) false,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField2 != null ? reservedField2 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 6),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (cancelWasPending)
    writeConstField(
        "cancelWasPending",
        CANCELWASPENDING,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField3 != null ? reservedField3 : (boolean) false,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (integerEncoding)
    writeSimpleEnumField(
        "integerEncoding",
        "IntegerEncoding",
        integerEncoding,
        new DataWriterEnumDefault<>(
            IntegerEncoding::getValue, IntegerEncoding::name, writeUnsignedByte(writeBuffer, 4)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (characterEncoding)
    writeSimpleEnumField(
        "characterEncoding",
        "CharacterEncoding",
        characterEncoding,
        new DataWriterEnumDefault<>(
            CharacterEncoding::getValue,
            CharacterEncoding::name,
            writeUnsignedByte(writeBuffer, 4)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (floatingPointEncoding)
    writeSimpleEnumField(
        "floatingPointEncoding",
        "FloatingPointEncoding",
        floatingPointEncoding,
        new DataWriterEnumDefault<>(
            FloatingPointEncoding::getValue,
            FloatingPointEncoding::name,
            writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField4 != null ? reservedField4 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (serialHigh)
    writeConstField(
        "serialHigh",
        SERIALHIGH,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (objectUuid)
    writeSimpleField(
        "objectUuid",
        objectUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (interfaceUuid)
    writeSimpleField(
        "interfaceUuid",
        interfaceUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (activityUuid)
    writeSimpleField(
        "activityUuid",
        activityUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (serverBootTime)
    writeSimpleField(
        "serverBootTime",
        serverBootTime,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (interfaceVer)
    writeConstField(
        "interfaceVer",
        INTERFACEVER,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (sequenceNumber)
    writeSimpleField(
        "sequenceNumber",
        sequenceNumber,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (operation)
    writeSimpleEnumField(
        "operation",
        "DceRpc_Operation",
        operation,
        new DataWriterEnumDefault<>(
            DceRpc_Operation::getValue, DceRpc_Operation::name, writeUnsignedInt(writeBuffer, 16)),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (interfaceHint)
    writeConstField(
        "interfaceHint",
        INTERFACEHINT,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (activityHint)
    writeConstField(
        "activityHint",
        ACTIVITYHINT,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Implicit Field (lengthOfBody) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int lengthOfBody = (int) (getPayload().getLengthInBytes());
    writeImplicitField(
        "lengthOfBody",
        lengthOfBody,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (fragmentNum)
    writeConstField(
        "fragmentNum",
        FRAGMENTNUM,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (authProto)
    writeConstField(
        "authProto",
        AUTHPROTO,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (serialLow)
    writeConstField(
        "serialLow",
        SERIALLOW,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (payload)
    writeSimpleField(
        "payload",
        payload,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    writeBuffer.popContext("DceRpc_Packet");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DceRpc_Packet _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (version)
    lengthInBits += 8;

    // Simple field (packetType)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Const Field (broadcast)
    lengthInBits += 1;

    // Simple field (idempotent)
    lengthInBits += 1;

    // Const Field (maybe)
    lengthInBits += 1;

    // Simple field (noFragmentAcknowledgeRequested)
    lengthInBits += 1;

    // Const Field (fragment)
    lengthInBits += 1;

    // Simple field (lastFragment)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 6;

    // Const Field (cancelWasPending)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (integerEncoding)
    lengthInBits += 4;

    // Simple field (characterEncoding)
    lengthInBits += 4;

    // Simple field (floatingPointEncoding)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Const Field (serialHigh)
    lengthInBits += 8;

    // Simple field (objectUuid)
    lengthInBits += objectUuid.getLengthInBits();

    // Simple field (interfaceUuid)
    lengthInBits += interfaceUuid.getLengthInBits();

    // Simple field (activityUuid)
    lengthInBits += activityUuid.getLengthInBits();

    // Simple field (serverBootTime)
    lengthInBits += 32;

    // Const Field (interfaceVer)
    lengthInBits += 32;

    // Simple field (sequenceNumber)
    lengthInBits += 32;

    // Simple field (operation)
    lengthInBits += 16;

    // Const Field (interfaceHint)
    lengthInBits += 16;

    // Const Field (activityHint)
    lengthInBits += 16;

    // Implicit Field (lengthOfBody)
    lengthInBits += 16;

    // Const Field (fragmentNum)
    lengthInBits += 16;

    // Const Field (authProto)
    lengthInBits += 8;

    // Const Field (serialLow)
    lengthInBits += 8;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static DceRpc_Packet staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DceRpc_Packet staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DceRpc_Packet");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short version =
        readConstField(
            "version",
            readUnsignedShort(readBuffer, 8),
            DceRpc_Packet.VERSION,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    DceRpc_PacketType packetType =
        readEnumField(
            "packetType",
            "DceRpc_PacketType",
            new DataReaderEnumDefault<>(
                DceRpc_PacketType::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Boolean reservedField0 =
        readReservedField(
            "reserved",
            readBoolean(readBuffer),
            (boolean) false,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean broadcast =
        readConstField(
            "broadcast",
            readBoolean(readBuffer),
            DceRpc_Packet.BROADCAST,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean idempotent =
        readSimpleField(
            "idempotent", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean maybe =
        readConstField(
            "maybe",
            readBoolean(readBuffer),
            DceRpc_Packet.MAYBE,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean noFragmentAcknowledgeRequested =
        readSimpleField(
            "noFragmentAcknowledgeRequested",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean fragment =
        readConstField(
            "fragment",
            readBoolean(readBuffer),
            DceRpc_Packet.FRAGMENT,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean lastFragment =
        readSimpleField(
            "lastFragment",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Boolean reservedField1 =
        readReservedField(
            "reserved",
            readBoolean(readBuffer),
            (boolean) false,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Short reservedField2 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 6),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean cancelWasPending =
        readConstField(
            "cancelWasPending",
            readBoolean(readBuffer),
            DceRpc_Packet.CANCELWASPENDING,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Boolean reservedField3 =
        readReservedField(
            "reserved",
            readBoolean(readBuffer),
            (boolean) false,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    IntegerEncoding integerEncoding =
        readEnumField(
            "integerEncoding",
            "IntegerEncoding",
            new DataReaderEnumDefault<>(
                IntegerEncoding::enumForValue, readUnsignedByte(readBuffer, 4)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    CharacterEncoding characterEncoding =
        readEnumField(
            "characterEncoding",
            "CharacterEncoding",
            new DataReaderEnumDefault<>(
                CharacterEncoding::enumForValue, readUnsignedByte(readBuffer, 4)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    FloatingPointEncoding floatingPointEncoding =
        readEnumField(
            "floatingPointEncoding",
            "FloatingPointEncoding",
            new DataReaderEnumDefault<>(
                FloatingPointEncoding::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Short reservedField4 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 8),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short serialHigh =
        readConstField(
            "serialHigh",
            readUnsignedShort(readBuffer, 8),
            DceRpc_Packet.SERIALHIGH,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    DceRpc_ObjectUuid objectUuid =
        readSimpleField(
            "objectUuid",
            new DataReaderComplexDefault<>(
                () -> DceRpc_ObjectUuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    DceRpc_InterfaceUuid interfaceUuid =
        readSimpleField(
            "interfaceUuid",
            new DataReaderComplexDefault<>(
                () -> DceRpc_InterfaceUuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    DceRpc_ActivityUuid activityUuid =
        readSimpleField(
            "activityUuid",
            new DataReaderComplexDefault<>(
                () -> DceRpc_ActivityUuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    long serverBootTime =
        readSimpleField(
            "serverBootTime",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    long interfaceVer =
        readConstField(
            "interfaceVer",
            readUnsignedLong(readBuffer, 32),
            DceRpc_Packet.INTERFACEVER,
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    long sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    DceRpc_Operation operation =
        readEnumField(
            "operation",
            "DceRpc_Operation",
            new DataReaderEnumDefault<>(
                DceRpc_Operation::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int interfaceHint =
        readConstField(
            "interfaceHint",
            readUnsignedInt(readBuffer, 16),
            DceRpc_Packet.INTERFACEHINT,
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int activityHint =
        readConstField(
            "activityHint",
            readUnsignedInt(readBuffer, 16),
            DceRpc_Packet.ACTIVITYHINT,
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int lengthOfBody =
        readImplicitField(
            "lengthOfBody",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int fragmentNum =
        readConstField(
            "fragmentNum",
            readUnsignedInt(readBuffer, 16),
            DceRpc_Packet.FRAGMENTNUM,
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    short authProto =
        readConstField(
            "authProto",
            readUnsignedShort(readBuffer, 8),
            DceRpc_Packet.AUTHPROTO,
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    short serialLow =
        readConstField(
            "serialLow",
            readUnsignedShort(readBuffer, 8),
            DceRpc_Packet.SERIALLOW,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    PnIoCm_Packet payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () -> PnIoCm_Packet.staticParse(readBuffer, (DceRpc_PacketType) (packetType)),
                readBuffer),
            WithOption.WithByteOrder(
                (((integerEncoding) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    readBuffer.closeContext("DceRpc_Packet");
    // Create the instance
    DceRpc_Packet _dceRpc_Packet;
    _dceRpc_Packet =
        new DceRpc_Packet(
            packetType,
            idempotent,
            noFragmentAcknowledgeRequested,
            lastFragment,
            integerEncoding,
            characterEncoding,
            floatingPointEncoding,
            objectUuid,
            interfaceUuid,
            activityUuid,
            serverBootTime,
            sequenceNumber,
            operation,
            payload);
    _dceRpc_Packet.reservedField0 = reservedField0;
    _dceRpc_Packet.reservedField1 = reservedField1;
    _dceRpc_Packet.reservedField2 = reservedField2;
    _dceRpc_Packet.reservedField3 = reservedField3;
    _dceRpc_Packet.reservedField4 = reservedField4;
    return _dceRpc_Packet;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DceRpc_Packet)) {
      return false;
    }
    DceRpc_Packet that = (DceRpc_Packet) o;
    return (getPacketType() == that.getPacketType())
        && (getIdempotent() == that.getIdempotent())
        && (getNoFragmentAcknowledgeRequested() == that.getNoFragmentAcknowledgeRequested())
        && (getLastFragment() == that.getLastFragment())
        && (getIntegerEncoding() == that.getIntegerEncoding())
        && (getCharacterEncoding() == that.getCharacterEncoding())
        && (getFloatingPointEncoding() == that.getFloatingPointEncoding())
        && (getObjectUuid() == that.getObjectUuid())
        && (getInterfaceUuid() == that.getInterfaceUuid())
        && (getActivityUuid() == that.getActivityUuid())
        && (getServerBootTime() == that.getServerBootTime())
        && (getSequenceNumber() == that.getSequenceNumber())
        && (getOperation() == that.getOperation())
        && (getPayload() == that.getPayload())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getPacketType(),
        getIdempotent(),
        getNoFragmentAcknowledgeRequested(),
        getLastFragment(),
        getIntegerEncoding(),
        getCharacterEncoding(),
        getFloatingPointEncoding(),
        getObjectUuid(),
        getInterfaceUuid(),
        getActivityUuid(),
        getServerBootTime(),
        getSequenceNumber(),
        getOperation(),
        getPayload());
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
