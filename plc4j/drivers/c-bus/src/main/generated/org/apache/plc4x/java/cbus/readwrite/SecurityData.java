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
package org.apache.plc4x.java.cbus.readwrite;

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

public abstract class SecurityData implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final SecurityCommandTypeContainer commandTypeContainer;
  protected final byte argument;

  public SecurityData(SecurityCommandTypeContainer commandTypeContainer, byte argument) {
    super();
    this.commandTypeContainer = commandTypeContainer;
    this.argument = argument;
  }

  public SecurityCommandTypeContainer getCommandTypeContainer() {
    return commandTypeContainer;
  }

  public byte getArgument() {
    return argument;
  }

  public SecurityCommandType getCommandType() {
    return (SecurityCommandType) (getCommandTypeContainer().getCommandType());
  }

  protected abstract void serializeSecurityDataChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SecurityData");

    // Simple Field (commandTypeContainer)
    writeSimpleEnumField(
        "commandTypeContainer",
        "SecurityCommandTypeContainer",
        commandTypeContainer,
        new DataWriterEnumDefault<>(
            SecurityCommandTypeContainer::getValue,
            SecurityCommandTypeContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    SecurityCommandType commandType = getCommandType();
    writeBuffer.writeVirtual("commandType", commandType);

    // Simple Field (argument)
    writeSimpleField("argument", argument, writeByte(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeSecurityDataChild(writeBuffer);

    writeBuffer.popContext("SecurityData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SecurityData _value = this;

    // Simple field (commandTypeContainer)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Simple field (argument)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static SecurityData staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static SecurityData staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("SecurityData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    // Validation
    if (!(org.apache.plc4x.java.cbus.readwrite.utils.StaticHelper.knowsSecurityCommandTypeContainer(
        readBuffer))) {
      throw new ParseAssertException("no command type could be found");
    }

    SecurityCommandTypeContainer commandTypeContainer =
        readEnumField(
            "commandTypeContainer",
            "SecurityCommandTypeContainer",
            new DataReaderEnumDefault<>(
                SecurityCommandTypeContainer::enumForValue, readUnsignedShort(readBuffer, 8)));
    SecurityCommandType commandType =
        readVirtualField(
            "commandType", SecurityCommandType.class, commandTypeContainer.getCommandType());

    byte argument = readSimpleField("argument", readByte(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    SecurityDataBuilder builder = null;
    if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x80)) {
      builder = SecurityDataSystemArmedDisarmed.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x80)) {
      builder = SecurityDataSystemDisarmed.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x81)) {
      builder = SecurityDataExitDelayStarted.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x82)) {
      builder = SecurityDataEntryDelayStarted.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x83)) {
      builder = SecurityDataAlarmOn.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x83)) {
      builder = SecurityDataAlarmOff.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x84)) {
      builder = SecurityDataTamperOn.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x84)) {
      builder = SecurityDataTamperOff.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x85)) {
      builder = SecurityDataPanicActivated.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x85)) {
      builder = SecurityDataPanicCleared.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x86)) {
      builder = SecurityDataZoneUnsealed.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x87)) {
      builder = SecurityDataZoneSealed.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x88)) {
      builder = SecurityDataZoneOpen.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x89)) {
      builder = SecurityDataZoneShort.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x89)) {
      builder = SecurityDataZoneIsolated.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x8B)) {
      builder = SecurityDataLowBatteryDetected.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x8B)) {
      builder = SecurityDataLowBatteryCorrected.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x8C)) {
      builder = SecurityDataLowBatteryCharging.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x8D)) {
      builder = SecurityDataZoneName.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x8E)) {
      builder = SecurityDataStatusReport1.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x8F)) {
      builder = SecurityDataStatusReport2.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x90)) {
      builder = SecurityDataPasswordEntryStatus.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x91)) {
      builder = SecurityDataMainsFailure.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x91)) {
      builder = SecurityDataMainsRestoredOrApplied.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x92)) {
      builder = SecurityDataArmReadyNotReady.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0x93)) {
      builder = SecurityDataCurrentAlarmType.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x94)) {
      builder = SecurityDataLineCutAlarmRaised.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x94)) {
      builder = SecurityDataLineCutAlarmCleared.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x95)) {
      builder = SecurityDataArmFailedRaised.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x95)) {
      builder = SecurityDataArmFailedCleared.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x96)) {
      builder = SecurityDataFireAlarmRaised.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x96)) {
      builder = SecurityDataFireAlarmCleared.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x97)) {
      builder = SecurityDataGasAlarmRaised.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x97)) {
      builder = SecurityDataGasAlarmCleared.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0x98)) {
      builder = SecurityDataOtherAlarmRaised.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0x98)) {
      builder = SecurityDataOtherAlarmCleared.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0xA0)) {
      builder = SecurityDataStatus1Request.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0xA1)) {
      builder = SecurityDataStatus2Request.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0xA2)) {
      builder = SecurityDataArmSystem.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0xA3)) {
      builder = SecurityDataRaiseTamper.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)
        && EvaluationHelper.equals(argument, (byte) 0xA3)) {
      builder = SecurityDataDropTamper.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0xA4)) {
      builder = SecurityDataRaiseAlarm.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0xA5)) {
      builder = SecurityDataEmulatedKeypad.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)
        && EvaluationHelper.equals(argument, (byte) 0xA6)) {
      builder =
          SecurityDataDisplayMessage.staticParseSecurityDataBuilder(
              readBuffer, commandTypeContainer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)
        && EvaluationHelper.equals(argument, (byte) 0xA7)) {
      builder = SecurityDataRequestZoneName.staticParseSecurityDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.OFF)) {
      builder = SecurityDataOff.staticParseSecurityDataBuilder(readBuffer, commandTypeContainer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.ON)) {
      builder = SecurityDataOn.staticParseSecurityDataBuilder(readBuffer, commandTypeContainer);
    } else if (EvaluationHelper.equals(commandType, SecurityCommandType.EVENT)) {
      builder = SecurityDataEvent.staticParseSecurityDataBuilder(readBuffer, commandTypeContainer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandType="
              + commandType
              + " "
              + "argument="
              + argument
              + "]");
    }

    readBuffer.closeContext("SecurityData");
    // Create the instance
    SecurityData _securityData = builder.build(commandTypeContainer, argument);
    return _securityData;
  }

  public interface SecurityDataBuilder {
    SecurityData build(SecurityCommandTypeContainer commandTypeContainer, byte argument);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityData)) {
      return false;
    }
    SecurityData that = (SecurityData) o;
    return (getCommandTypeContainer() == that.getCommandTypeContainer())
        && (getArgument() == that.getArgument())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCommandTypeContainer(), getArgument());
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
