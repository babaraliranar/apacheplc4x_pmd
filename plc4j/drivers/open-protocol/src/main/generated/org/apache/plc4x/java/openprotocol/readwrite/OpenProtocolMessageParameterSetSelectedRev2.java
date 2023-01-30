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
package org.apache.plc4x.java.openprotocol.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class OpenProtocolMessageParameterSetSelectedRev2
    extends OpenProtocolMessageParameterSetSelected implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 2;
  }

  // Constant values.
  public static final Integer BLOCKIDPARAMETERSETID = 1;
  public static final Integer BLOCKIDPARAMETERSETNAME = 2;
  public static final Integer BLOCKIDDATEOFLASTCHANGEINPARAMETERSETSETTING = 3;
  public static final Integer BLOCKIDROTATIONDIRECTION = 4;
  public static final Integer BLOCKIDBATCHSIZE = 5;
  public static final Integer BLOCKIDTORQUEMIN = 6;
  public static final Integer BLOCKIDTORQUEMAX = 7;
  public static final Integer BLOCKIDTORQUEFINALTARGET = 8;
  public static final Integer BLOCKIDANGLEMIN = 9;
  public static final Integer BLOCKIDANGLEMAX = 10;
  public static final Integer BLOCKIDFINALANGLETARGET = 11;
  public static final Integer BLOCKIDFIRSTTARGET = 12;
  public static final Integer BLOCKIDSTARTFINALANGLE = 13;

  // Properties.
  protected final long parameterSetId;
  protected final String parameterSetName;
  protected final String dateOfLastChangeInParameterSetSetting;
  protected final RotationDirection rotationDirection;
  protected final int batchSize;
  protected final BigInteger torqueMin;
  protected final BigInteger torqueMax;
  protected final BigInteger torqueFinalTarget;
  protected final BigInteger angleMin;
  protected final BigInteger angleMax;
  protected final BigInteger finalAngleTarget;
  protected final BigInteger firstTarget;
  protected final BigInteger startFinalAngle;

  public OpenProtocolMessageParameterSetSelectedRev2(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      long parameterSetId,
      String parameterSetName,
      String dateOfLastChangeInParameterSetSetting,
      RotationDirection rotationDirection,
      int batchSize,
      BigInteger torqueMin,
      BigInteger torqueMax,
      BigInteger torqueFinalTarget,
      BigInteger angleMin,
      BigInteger angleMax,
      BigInteger finalAngleTarget,
      BigInteger firstTarget,
      BigInteger startFinalAngle) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.parameterSetId = parameterSetId;
    this.parameterSetName = parameterSetName;
    this.dateOfLastChangeInParameterSetSetting = dateOfLastChangeInParameterSetSetting;
    this.rotationDirection = rotationDirection;
    this.batchSize = batchSize;
    this.torqueMin = torqueMin;
    this.torqueMax = torqueMax;
    this.torqueFinalTarget = torqueFinalTarget;
    this.angleMin = angleMin;
    this.angleMax = angleMax;
    this.finalAngleTarget = finalAngleTarget;
    this.firstTarget = firstTarget;
    this.startFinalAngle = startFinalAngle;
  }

  public long getParameterSetId() {
    return parameterSetId;
  }

  public String getParameterSetName() {
    return parameterSetName;
  }

  public String getDateOfLastChangeInParameterSetSetting() {
    return dateOfLastChangeInParameterSetSetting;
  }

  public RotationDirection getRotationDirection() {
    return rotationDirection;
  }

  public int getBatchSize() {
    return batchSize;
  }

  public BigInteger getTorqueMin() {
    return torqueMin;
  }

  public BigInteger getTorqueMax() {
    return torqueMax;
  }

  public BigInteger getTorqueFinalTarget() {
    return torqueFinalTarget;
  }

  public BigInteger getAngleMin() {
    return angleMin;
  }

  public BigInteger getAngleMax() {
    return angleMax;
  }

  public BigInteger getFinalAngleTarget() {
    return finalAngleTarget;
  }

  public BigInteger getFirstTarget() {
    return firstTarget;
  }

  public BigInteger getStartFinalAngle() {
    return startFinalAngle;
  }

  public int getBlockIdParameterSetId() {
    return BLOCKIDPARAMETERSETID;
  }

  public int getBlockIdParameterSetName() {
    return BLOCKIDPARAMETERSETNAME;
  }

  public int getBlockIdDateOfLastChangeInParameterSetSetting() {
    return BLOCKIDDATEOFLASTCHANGEINPARAMETERSETSETTING;
  }

  public int getBlockIdRotationDirection() {
    return BLOCKIDROTATIONDIRECTION;
  }

  public int getBlockIdBatchSize() {
    return BLOCKIDBATCHSIZE;
  }

  public int getBlockIdTorqueMin() {
    return BLOCKIDTORQUEMIN;
  }

  public int getBlockIdTorqueMax() {
    return BLOCKIDTORQUEMAX;
  }

  public int getBlockIdTorqueFinalTarget() {
    return BLOCKIDTORQUEFINALTARGET;
  }

  public int getBlockIdAngleMin() {
    return BLOCKIDANGLEMIN;
  }

  public int getBlockIdAngleMax() {
    return BLOCKIDANGLEMAX;
  }

  public int getBlockIdFinalAngleTarget() {
    return BLOCKIDFINALANGLETARGET;
  }

  public int getBlockIdFirstTarget() {
    return BLOCKIDFIRSTTARGET;
  }

  public int getBlockIdStartFinalAngle() {
    return BLOCKIDSTARTFINALANGLE;
  }

  @Override
  protected void serializeOpenProtocolMessageParameterSetSelectedChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageParameterSetSelectedRev2");

    // Const Field (blockIdParameterSetId)
    writeConstField(
        "blockIdParameterSetId",
        BLOCKIDPARAMETERSETID,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (parameterSetId)
    writeSimpleField(
        "parameterSetId",
        parameterSetId,
        writeUnsignedLong(writeBuffer, 24),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdParameterSetName)
    writeConstField(
        "blockIdParameterSetName",
        BLOCKIDPARAMETERSETNAME,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (parameterSetName)
    writeSimpleField(
        "parameterSetName",
        parameterSetName,
        writeString(writeBuffer, 200),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdDateOfLastChangeInParameterSetSetting)
    writeConstField(
        "blockIdDateOfLastChangeInParameterSetSetting",
        BLOCKIDDATEOFLASTCHANGEINPARAMETERSETSETTING,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (dateOfLastChangeInParameterSetSetting)
    writeSimpleField(
        "dateOfLastChangeInParameterSetSetting",
        dateOfLastChangeInParameterSetSetting,
        writeString(writeBuffer, 152),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdRotationDirection)
    writeConstField(
        "blockIdRotationDirection",
        BLOCKIDROTATIONDIRECTION,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (rotationDirection)
    writeSimpleEnumField(
        "rotationDirection",
        "RotationDirection",
        rotationDirection,
        new DataWriterEnumDefault<>(
            RotationDirection::getValue,
            RotationDirection::name,
            writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdBatchSize)
    writeConstField(
        "blockIdBatchSize",
        BLOCKIDBATCHSIZE,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (batchSize)
    writeSimpleField(
        "batchSize",
        batchSize,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdTorqueMin)
    writeConstField(
        "blockIdTorqueMin",
        BLOCKIDTORQUEMIN,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (torqueMin)
    writeSimpleField(
        "torqueMin",
        torqueMin,
        writeUnsignedBigInteger(writeBuffer, 48),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdTorqueMax)
    writeConstField(
        "blockIdTorqueMax",
        BLOCKIDTORQUEMAX,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (torqueMax)
    writeSimpleField(
        "torqueMax",
        torqueMax,
        writeUnsignedBigInteger(writeBuffer, 48),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdTorqueFinalTarget)
    writeConstField(
        "blockIdTorqueFinalTarget",
        BLOCKIDTORQUEFINALTARGET,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (torqueFinalTarget)
    writeSimpleField(
        "torqueFinalTarget",
        torqueFinalTarget,
        writeUnsignedBigInteger(writeBuffer, 48),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdAngleMin)
    writeConstField(
        "blockIdAngleMin",
        BLOCKIDANGLEMIN,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (angleMin)
    writeSimpleField(
        "angleMin",
        angleMin,
        writeUnsignedBigInteger(writeBuffer, 40),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdAngleMax)
    writeConstField(
        "blockIdAngleMax",
        BLOCKIDANGLEMAX,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (angleMax)
    writeSimpleField(
        "angleMax",
        angleMax,
        writeUnsignedBigInteger(writeBuffer, 40),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdFinalAngleTarget)
    writeConstField(
        "blockIdFinalAngleTarget",
        BLOCKIDFINALANGLETARGET,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (finalAngleTarget)
    writeSimpleField(
        "finalAngleTarget",
        finalAngleTarget,
        writeUnsignedBigInteger(writeBuffer, 40),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdFirstTarget)
    writeConstField(
        "blockIdFirstTarget",
        BLOCKIDFIRSTTARGET,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (firstTarget)
    writeSimpleField(
        "firstTarget",
        firstTarget,
        writeUnsignedBigInteger(writeBuffer, 48),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdStartFinalAngle)
    writeConstField(
        "blockIdStartFinalAngle",
        BLOCKIDSTARTFINALANGLE,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (startFinalAngle)
    writeSimpleField(
        "startFinalAngle",
        startFinalAngle,
        writeUnsignedBigInteger(writeBuffer, 48),
        WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageParameterSetSelectedRev2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageParameterSetSelectedRev2 _value = this;

    // Const Field (blockIdParameterSetId)
    lengthInBits += 16;

    // Simple field (parameterSetId)
    lengthInBits += 24;

    // Const Field (blockIdParameterSetName)
    lengthInBits += 16;

    // Simple field (parameterSetName)
    lengthInBits += 200;

    // Const Field (blockIdDateOfLastChangeInParameterSetSetting)
    lengthInBits += 16;

    // Simple field (dateOfLastChangeInParameterSetSetting)
    lengthInBits += 152;

    // Const Field (blockIdRotationDirection)
    lengthInBits += 16;

    // Simple field (rotationDirection)
    lengthInBits += 8;

    // Const Field (blockIdBatchSize)
    lengthInBits += 16;

    // Simple field (batchSize)
    lengthInBits += 16;

    // Const Field (blockIdTorqueMin)
    lengthInBits += 16;

    // Simple field (torqueMin)
    lengthInBits += 48;

    // Const Field (blockIdTorqueMax)
    lengthInBits += 16;

    // Simple field (torqueMax)
    lengthInBits += 48;

    // Const Field (blockIdTorqueFinalTarget)
    lengthInBits += 16;

    // Simple field (torqueFinalTarget)
    lengthInBits += 48;

    // Const Field (blockIdAngleMin)
    lengthInBits += 16;

    // Simple field (angleMin)
    lengthInBits += 40;

    // Const Field (blockIdAngleMax)
    lengthInBits += 16;

    // Simple field (angleMax)
    lengthInBits += 40;

    // Const Field (blockIdFinalAngleTarget)
    lengthInBits += 16;

    // Simple field (finalAngleTarget)
    lengthInBits += 40;

    // Const Field (blockIdFirstTarget)
    lengthInBits += 16;

    // Simple field (firstTarget)
    lengthInBits += 48;

    // Const Field (blockIdStartFinalAngle)
    lengthInBits += 16;

    // Simple field (startFinalAngle)
    lengthInBits += 48;

    return lengthInBits;
  }

  public static OpenProtocolMessageParameterSetSelectedBuilder
      staticParseOpenProtocolMessageParameterSetSelectedBuilder(
          ReadBuffer readBuffer, Long revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageParameterSetSelectedRev2");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int blockIdParameterSetId =
        readConstField(
            "blockIdParameterSetId",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDPARAMETERSETID,
            WithOption.WithEncoding("ASCII"));

    long parameterSetId =
        readSimpleField(
            "parameterSetId", readUnsignedLong(readBuffer, 24), WithOption.WithEncoding("ASCII"));

    int blockIdParameterSetName =
        readConstField(
            "blockIdParameterSetName",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDPARAMETERSETNAME,
            WithOption.WithEncoding("ASCII"));

    String parameterSetName =
        readSimpleField(
            "parameterSetName", readString(readBuffer, 200), WithOption.WithEncoding("ASCII"));

    int blockIdDateOfLastChangeInParameterSetSetting =
        readConstField(
            "blockIdDateOfLastChangeInParameterSetSetting",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2
                .BLOCKIDDATEOFLASTCHANGEINPARAMETERSETSETTING,
            WithOption.WithEncoding("ASCII"));

    String dateOfLastChangeInParameterSetSetting =
        readSimpleField(
            "dateOfLastChangeInParameterSetSetting",
            readString(readBuffer, 152),
            WithOption.WithEncoding("ASCII"));

    int blockIdRotationDirection =
        readConstField(
            "blockIdRotationDirection",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDROTATIONDIRECTION,
            WithOption.WithEncoding("ASCII"));

    RotationDirection rotationDirection =
        readEnumField(
            "rotationDirection",
            "RotationDirection",
            new DataReaderEnumDefault<>(
                RotationDirection::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithEncoding("ASCII"));

    int blockIdBatchSize =
        readConstField(
            "blockIdBatchSize",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDBATCHSIZE,
            WithOption.WithEncoding("ASCII"));

    int batchSize =
        readSimpleField(
            "batchSize", readUnsignedInt(readBuffer, 16), WithOption.WithEncoding("ASCII"));

    int blockIdTorqueMin =
        readConstField(
            "blockIdTorqueMin",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDTORQUEMIN,
            WithOption.WithEncoding("ASCII"));

    BigInteger torqueMin =
        readSimpleField(
            "torqueMin", readUnsignedBigInteger(readBuffer, 48), WithOption.WithEncoding("ASCII"));

    int blockIdTorqueMax =
        readConstField(
            "blockIdTorqueMax",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDTORQUEMAX,
            WithOption.WithEncoding("ASCII"));

    BigInteger torqueMax =
        readSimpleField(
            "torqueMax", readUnsignedBigInteger(readBuffer, 48), WithOption.WithEncoding("ASCII"));

    int blockIdTorqueFinalTarget =
        readConstField(
            "blockIdTorqueFinalTarget",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDTORQUEFINALTARGET,
            WithOption.WithEncoding("ASCII"));

    BigInteger torqueFinalTarget =
        readSimpleField(
            "torqueFinalTarget",
            readUnsignedBigInteger(readBuffer, 48),
            WithOption.WithEncoding("ASCII"));

    int blockIdAngleMin =
        readConstField(
            "blockIdAngleMin",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDANGLEMIN,
            WithOption.WithEncoding("ASCII"));

    BigInteger angleMin =
        readSimpleField(
            "angleMin", readUnsignedBigInteger(readBuffer, 40), WithOption.WithEncoding("ASCII"));

    int blockIdAngleMax =
        readConstField(
            "blockIdAngleMax",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDANGLEMAX,
            WithOption.WithEncoding("ASCII"));

    BigInteger angleMax =
        readSimpleField(
            "angleMax", readUnsignedBigInteger(readBuffer, 40), WithOption.WithEncoding("ASCII"));

    int blockIdFinalAngleTarget =
        readConstField(
            "blockIdFinalAngleTarget",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDFINALANGLETARGET,
            WithOption.WithEncoding("ASCII"));

    BigInteger finalAngleTarget =
        readSimpleField(
            "finalAngleTarget",
            readUnsignedBigInteger(readBuffer, 40),
            WithOption.WithEncoding("ASCII"));

    int blockIdFirstTarget =
        readConstField(
            "blockIdFirstTarget",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDFIRSTTARGET,
            WithOption.WithEncoding("ASCII"));

    BigInteger firstTarget =
        readSimpleField(
            "firstTarget",
            readUnsignedBigInteger(readBuffer, 48),
            WithOption.WithEncoding("ASCII"));

    int blockIdStartFinalAngle =
        readConstField(
            "blockIdStartFinalAngle",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageParameterSetSelectedRev2.BLOCKIDSTARTFINALANGLE,
            WithOption.WithEncoding("ASCII"));

    BigInteger startFinalAngle =
        readSimpleField(
            "startFinalAngle",
            readUnsignedBigInteger(readBuffer, 48),
            WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageParameterSetSelectedRev2");
    // Create the instance
    return new OpenProtocolMessageParameterSetSelectedRev2BuilderImpl(
        parameterSetId,
        parameterSetName,
        dateOfLastChangeInParameterSetSetting,
        rotationDirection,
        batchSize,
        torqueMin,
        torqueMax,
        torqueFinalTarget,
        angleMin,
        angleMax,
        finalAngleTarget,
        firstTarget,
        startFinalAngle);
  }

  public static class OpenProtocolMessageParameterSetSelectedRev2BuilderImpl
      implements OpenProtocolMessageParameterSetSelected
          .OpenProtocolMessageParameterSetSelectedBuilder {
    private final long parameterSetId;
    private final String parameterSetName;
    private final String dateOfLastChangeInParameterSetSetting;
    private final RotationDirection rotationDirection;
    private final int batchSize;
    private final BigInteger torqueMin;
    private final BigInteger torqueMax;
    private final BigInteger torqueFinalTarget;
    private final BigInteger angleMin;
    private final BigInteger angleMax;
    private final BigInteger finalAngleTarget;
    private final BigInteger firstTarget;
    private final BigInteger startFinalAngle;

    public OpenProtocolMessageParameterSetSelectedRev2BuilderImpl(
        long parameterSetId,
        String parameterSetName,
        String dateOfLastChangeInParameterSetSetting,
        RotationDirection rotationDirection,
        int batchSize,
        BigInteger torqueMin,
        BigInteger torqueMax,
        BigInteger torqueFinalTarget,
        BigInteger angleMin,
        BigInteger angleMax,
        BigInteger finalAngleTarget,
        BigInteger firstTarget,
        BigInteger startFinalAngle) {
      this.parameterSetId = parameterSetId;
      this.parameterSetName = parameterSetName;
      this.dateOfLastChangeInParameterSetSetting = dateOfLastChangeInParameterSetSetting;
      this.rotationDirection = rotationDirection;
      this.batchSize = batchSize;
      this.torqueMin = torqueMin;
      this.torqueMax = torqueMax;
      this.torqueFinalTarget = torqueFinalTarget;
      this.angleMin = angleMin;
      this.angleMax = angleMax;
      this.finalAngleTarget = finalAngleTarget;
      this.firstTarget = firstTarget;
      this.startFinalAngle = startFinalAngle;
    }

    public OpenProtocolMessageParameterSetSelectedRev2 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageParameterSetSelectedRev2 openProtocolMessageParameterSetSelectedRev2 =
          new OpenProtocolMessageParameterSetSelectedRev2(
              midRevision,
              noAckFlag,
              targetStationId,
              targetSpindleId,
              sequenceNumber,
              numberOfMessageParts,
              messagePartNumber,
              parameterSetId,
              parameterSetName,
              dateOfLastChangeInParameterSetSetting,
              rotationDirection,
              batchSize,
              torqueMin,
              torqueMax,
              torqueFinalTarget,
              angleMin,
              angleMax,
              finalAngleTarget,
              firstTarget,
              startFinalAngle);
      return openProtocolMessageParameterSetSelectedRev2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageParameterSetSelectedRev2)) {
      return false;
    }
    OpenProtocolMessageParameterSetSelectedRev2 that =
        (OpenProtocolMessageParameterSetSelectedRev2) o;
    return (getParameterSetId() == that.getParameterSetId())
        && (getParameterSetName() == that.getParameterSetName())
        && (getDateOfLastChangeInParameterSetSetting()
            == that.getDateOfLastChangeInParameterSetSetting())
        && (getRotationDirection() == that.getRotationDirection())
        && (getBatchSize() == that.getBatchSize())
        && (getTorqueMin() == that.getTorqueMin())
        && (getTorqueMax() == that.getTorqueMax())
        && (getTorqueFinalTarget() == that.getTorqueFinalTarget())
        && (getAngleMin() == that.getAngleMin())
        && (getAngleMax() == that.getAngleMax())
        && (getFinalAngleTarget() == that.getFinalAngleTarget())
        && (getFirstTarget() == that.getFirstTarget())
        && (getStartFinalAngle() == that.getStartFinalAngle())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getParameterSetId(),
        getParameterSetName(),
        getDateOfLastChangeInParameterSetSetting(),
        getRotationDirection(),
        getBatchSize(),
        getTorqueMin(),
        getTorqueMax(),
        getTorqueFinalTarget(),
        getAngleMin(),
        getAngleMax(),
        getFinalAngleTarget(),
        getFirstTarget(),
        getStartFinalAngle());
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
