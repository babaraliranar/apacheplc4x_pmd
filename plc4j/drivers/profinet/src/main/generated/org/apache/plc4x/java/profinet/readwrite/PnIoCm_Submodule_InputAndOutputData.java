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

public class PnIoCm_Submodule_InputAndOutputData extends PnIoCm_Submodule implements Message {

  // Accessors for discriminator values.
  public PnIoCm_SubmoduleType getSubmoduleType() {
    return PnIoCm_SubmoduleType.INPUT_AND_OUTPUT_DATA;
  }

  // Constant values.
  public static final Integer INPUTDATADESCRIPTION = 0x0001;
  public static final Integer OUTPUTDATADESCRIPTION = 0x0002;

  // Properties.
  protected final int inputSubmoduleDataLength;
  protected final short inputLengthIoCs;
  protected final short inputLengthIoPs;
  protected final int outputSubmoduleDataLength;
  protected final short outputLengthIoCs;
  protected final short outputLengthIoPs;

  public PnIoCm_Submodule_InputAndOutputData(
      int slotNumber,
      long submoduleIdentNumber,
      boolean discardIoxs,
      boolean reduceOutputModuleDataLength,
      boolean reduceInputModuleDataLength,
      boolean sharedInput,
      int inputSubmoduleDataLength,
      short inputLengthIoCs,
      short inputLengthIoPs,
      int outputSubmoduleDataLength,
      short outputLengthIoCs,
      short outputLengthIoPs) {
    super(
        slotNumber,
        submoduleIdentNumber,
        discardIoxs,
        reduceOutputModuleDataLength,
        reduceInputModuleDataLength,
        sharedInput);
    this.inputSubmoduleDataLength = inputSubmoduleDataLength;
    this.inputLengthIoCs = inputLengthIoCs;
    this.inputLengthIoPs = inputLengthIoPs;
    this.outputSubmoduleDataLength = outputSubmoduleDataLength;
    this.outputLengthIoCs = outputLengthIoCs;
    this.outputLengthIoPs = outputLengthIoPs;
  }

  public int getInputSubmoduleDataLength() {
    return inputSubmoduleDataLength;
  }

  public short getInputLengthIoCs() {
    return inputLengthIoCs;
  }

  public short getInputLengthIoPs() {
    return inputLengthIoPs;
  }

  public int getOutputSubmoduleDataLength() {
    return outputSubmoduleDataLength;
  }

  public short getOutputLengthIoCs() {
    return outputLengthIoCs;
  }

  public short getOutputLengthIoPs() {
    return outputLengthIoPs;
  }

  public int getInputDataDescription() {
    return INPUTDATADESCRIPTION;
  }

  public int getOutputDataDescription() {
    return OUTPUTDATADESCRIPTION;
  }

  @Override
  protected void serializePnIoCm_SubmoduleChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Submodule_InputAndOutputData");

    // Const Field (inputDataDescription)
    writeConstField(
        "inputDataDescription", INPUTDATADESCRIPTION, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (inputSubmoduleDataLength)
    writeSimpleField(
        "inputSubmoduleDataLength", inputSubmoduleDataLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (inputLengthIoCs)
    writeSimpleField("inputLengthIoCs", inputLengthIoCs, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (inputLengthIoPs)
    writeSimpleField("inputLengthIoPs", inputLengthIoPs, writeUnsignedShort(writeBuffer, 8));

    // Const Field (outputDataDescription)
    writeConstField(
        "outputDataDescription", OUTPUTDATADESCRIPTION, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (outputSubmoduleDataLength)
    writeSimpleField(
        "outputSubmoduleDataLength", outputSubmoduleDataLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (outputLengthIoCs)
    writeSimpleField("outputLengthIoCs", outputLengthIoCs, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (outputLengthIoPs)
    writeSimpleField("outputLengthIoPs", outputLengthIoPs, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("PnIoCm_Submodule_InputAndOutputData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Submodule_InputAndOutputData _value = this;

    // Const Field (inputDataDescription)
    lengthInBits += 16;

    // Simple field (inputSubmoduleDataLength)
    lengthInBits += 16;

    // Simple field (inputLengthIoCs)
    lengthInBits += 8;

    // Simple field (inputLengthIoPs)
    lengthInBits += 8;

    // Const Field (outputDataDescription)
    lengthInBits += 16;

    // Simple field (outputSubmoduleDataLength)
    lengthInBits += 16;

    // Simple field (outputLengthIoCs)
    lengthInBits += 8;

    // Simple field (outputLengthIoPs)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static PnIoCm_SubmoduleBuilder staticParsePnIoCm_SubmoduleBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Submodule_InputAndOutputData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int inputDataDescription =
        readConstField(
            "inputDataDescription",
            readUnsignedInt(readBuffer, 16),
            PnIoCm_Submodule_InputAndOutputData.INPUTDATADESCRIPTION);

    int inputSubmoduleDataLength =
        readSimpleField("inputSubmoduleDataLength", readUnsignedInt(readBuffer, 16));

    short inputLengthIoCs = readSimpleField("inputLengthIoCs", readUnsignedShort(readBuffer, 8));

    short inputLengthIoPs = readSimpleField("inputLengthIoPs", readUnsignedShort(readBuffer, 8));

    int outputDataDescription =
        readConstField(
            "outputDataDescription",
            readUnsignedInt(readBuffer, 16),
            PnIoCm_Submodule_InputAndOutputData.OUTPUTDATADESCRIPTION);

    int outputSubmoduleDataLength =
        readSimpleField("outputSubmoduleDataLength", readUnsignedInt(readBuffer, 16));

    short outputLengthIoCs = readSimpleField("outputLengthIoCs", readUnsignedShort(readBuffer, 8));

    short outputLengthIoPs = readSimpleField("outputLengthIoPs", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("PnIoCm_Submodule_InputAndOutputData");
    // Create the instance
    return new PnIoCm_Submodule_InputAndOutputDataBuilderImpl(
        inputSubmoduleDataLength,
        inputLengthIoCs,
        inputLengthIoPs,
        outputSubmoduleDataLength,
        outputLengthIoCs,
        outputLengthIoPs);
  }

  public static class PnIoCm_Submodule_InputAndOutputDataBuilderImpl
      implements PnIoCm_Submodule.PnIoCm_SubmoduleBuilder {
    private final int inputSubmoduleDataLength;
    private final short inputLengthIoCs;
    private final short inputLengthIoPs;
    private final int outputSubmoduleDataLength;
    private final short outputLengthIoCs;
    private final short outputLengthIoPs;

    public PnIoCm_Submodule_InputAndOutputDataBuilderImpl(
        int inputSubmoduleDataLength,
        short inputLengthIoCs,
        short inputLengthIoPs,
        int outputSubmoduleDataLength,
        short outputLengthIoCs,
        short outputLengthIoPs) {
      this.inputSubmoduleDataLength = inputSubmoduleDataLength;
      this.inputLengthIoCs = inputLengthIoCs;
      this.inputLengthIoPs = inputLengthIoPs;
      this.outputSubmoduleDataLength = outputSubmoduleDataLength;
      this.outputLengthIoCs = outputLengthIoCs;
      this.outputLengthIoPs = outputLengthIoPs;
    }

    public PnIoCm_Submodule_InputAndOutputData build(
        int slotNumber,
        long submoduleIdentNumber,
        boolean discardIoxs,
        boolean reduceOutputModuleDataLength,
        boolean reduceInputModuleDataLength,
        boolean sharedInput) {
      PnIoCm_Submodule_InputAndOutputData pnIoCm_Submodule_InputAndOutputData =
          new PnIoCm_Submodule_InputAndOutputData(
              slotNumber,
              submoduleIdentNumber,
              discardIoxs,
              reduceOutputModuleDataLength,
              reduceInputModuleDataLength,
              sharedInput,
              inputSubmoduleDataLength,
              inputLengthIoCs,
              inputLengthIoPs,
              outputSubmoduleDataLength,
              outputLengthIoCs,
              outputLengthIoPs);
      return pnIoCm_Submodule_InputAndOutputData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Submodule_InputAndOutputData)) {
      return false;
    }
    PnIoCm_Submodule_InputAndOutputData that = (PnIoCm_Submodule_InputAndOutputData) o;
    return (getInputSubmoduleDataLength() == that.getInputSubmoduleDataLength())
        && (getInputLengthIoCs() == that.getInputLengthIoCs())
        && (getInputLengthIoPs() == that.getInputLengthIoPs())
        && (getOutputSubmoduleDataLength() == that.getOutputSubmoduleDataLength())
        && (getOutputLengthIoCs() == that.getOutputLengthIoCs())
        && (getOutputLengthIoPs() == that.getOutputLengthIoPs())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInputSubmoduleDataLength(),
        getInputLengthIoCs(),
        getInputLengthIoPs(),
        getOutputSubmoduleDataLength(),
        getOutputLengthIoCs(),
        getOutputLengthIoPs());
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
