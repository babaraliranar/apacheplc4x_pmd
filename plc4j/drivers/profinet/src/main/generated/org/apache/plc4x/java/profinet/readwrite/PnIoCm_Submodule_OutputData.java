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

public class PnIoCm_Submodule_OutputData extends PnIoCm_Submodule implements Message {

  // Accessors for discriminator values.
  public PnIoCm_SubmoduleType getSubmoduleType() {
    return PnIoCm_SubmoduleType.OUTPUT_DATA;
  }

  // Constant values.
  public static final Integer INPUTDATADESCRIPTION = 0x0002;

  // Properties.
  protected final int inputSubmoduleDataLength;
  protected final short inputLengthIoCs;
  protected final short inputLengthIoPs;

  public PnIoCm_Submodule_OutputData(
      int slotNumber,
      long submoduleIdentNumber,
      boolean discardIoxs,
      boolean reduceOutputModuleDataLength,
      boolean reduceInputModuleDataLength,
      boolean sharedInput,
      int inputSubmoduleDataLength,
      short inputLengthIoCs,
      short inputLengthIoPs) {
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

  public int getInputDataDescription() {
    return INPUTDATADESCRIPTION;
  }

  @Override
  protected void serializePnIoCm_SubmoduleChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Submodule_OutputData");

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

    writeBuffer.popContext("PnIoCm_Submodule_OutputData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Submodule_OutputData _value = this;

    // Const Field (inputDataDescription)
    lengthInBits += 16;

    // Simple field (inputSubmoduleDataLength)
    lengthInBits += 16;

    // Simple field (inputLengthIoCs)
    lengthInBits += 8;

    // Simple field (inputLengthIoPs)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static PnIoCm_Submodule_OutputDataBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Submodule_OutputData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int inputDataDescription =
        readConstField(
            "inputDataDescription",
            readUnsignedInt(readBuffer, 16),
            PnIoCm_Submodule_OutputData.INPUTDATADESCRIPTION);

    int inputSubmoduleDataLength =
        readSimpleField("inputSubmoduleDataLength", readUnsignedInt(readBuffer, 16));

    short inputLengthIoCs = readSimpleField("inputLengthIoCs", readUnsignedShort(readBuffer, 8));

    short inputLengthIoPs = readSimpleField("inputLengthIoPs", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("PnIoCm_Submodule_OutputData");
    // Create the instance
    return new PnIoCm_Submodule_OutputDataBuilder(
        inputSubmoduleDataLength, inputLengthIoCs, inputLengthIoPs);
  }

  public static class PnIoCm_Submodule_OutputDataBuilder
      implements PnIoCm_Submodule.PnIoCm_SubmoduleBuilder {
    private final int inputSubmoduleDataLength;
    private final short inputLengthIoCs;
    private final short inputLengthIoPs;

    public PnIoCm_Submodule_OutputDataBuilder(
        int inputSubmoduleDataLength, short inputLengthIoCs, short inputLengthIoPs) {

      this.inputSubmoduleDataLength = inputSubmoduleDataLength;
      this.inputLengthIoCs = inputLengthIoCs;
      this.inputLengthIoPs = inputLengthIoPs;
    }

    public PnIoCm_Submodule_OutputData build(
        int slotNumber,
        long submoduleIdentNumber,
        boolean discardIoxs,
        boolean reduceOutputModuleDataLength,
        boolean reduceInputModuleDataLength,
        boolean sharedInput) {
      PnIoCm_Submodule_OutputData pnIoCm_Submodule_OutputData =
          new PnIoCm_Submodule_OutputData(
              slotNumber,
              submoduleIdentNumber,
              discardIoxs,
              reduceOutputModuleDataLength,
              reduceInputModuleDataLength,
              sharedInput,
              inputSubmoduleDataLength,
              inputLengthIoCs,
              inputLengthIoPs);
      return pnIoCm_Submodule_OutputData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Submodule_OutputData)) {
      return false;
    }
    PnIoCm_Submodule_OutputData that = (PnIoCm_Submodule_OutputData) o;
    return (getInputSubmoduleDataLength() == that.getInputSubmoduleDataLength())
        && (getInputLengthIoCs() == that.getInputLengthIoCs())
        && (getInputLengthIoPs() == that.getInputLengthIoPs())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInputSubmoduleDataLength(),
        getInputLengthIoCs(),
        getInputLengthIoPs());
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
