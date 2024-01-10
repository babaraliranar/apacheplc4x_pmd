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

public class UserManagementDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "24283";
  }

  // Properties.
  protected final PascalString userName;
  protected final UserConfigurationMask userConfiguration;
  protected final PascalString description;

  public UserManagementDataType(
      PascalString userName, UserConfigurationMask userConfiguration, PascalString description) {
    super();
    this.userName = userName;
    this.userConfiguration = userConfiguration;
    this.description = description;
  }

  public PascalString getUserName() {
    return userName;
  }

  public UserConfigurationMask getUserConfiguration() {
    return userConfiguration;
  }

  public PascalString getDescription() {
    return description;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("UserManagementDataType");

    // Simple Field (userName)
    writeSimpleField("userName", userName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (userConfiguration)
    writeSimpleEnumField(
        "userConfiguration",
        "UserConfigurationMask",
        userConfiguration,
        new DataWriterEnumDefault<>(
            UserConfigurationMask::getValue,
            UserConfigurationMask::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (description)
    writeSimpleField("description", description, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("UserManagementDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    UserManagementDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (userName)
    lengthInBits += userName.getLengthInBits();

    // Simple field (userConfiguration)
    lengthInBits += 32;

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("UserManagementDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString userName =
        readSimpleField(
            "userName",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    UserConfigurationMask userConfiguration =
        readEnumField(
            "userConfiguration",
            "UserConfigurationMask",
            new DataReaderEnumDefault<>(
                UserConfigurationMask::enumForValue, readUnsignedLong(readBuffer, 32)));

    PascalString description =
        readSimpleField(
            "description",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("UserManagementDataType");
    // Create the instance
    return new UserManagementDataTypeBuilderImpl(userName, userConfiguration, description);
  }

  public static class UserManagementDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString userName;
    private final UserConfigurationMask userConfiguration;
    private final PascalString description;

    public UserManagementDataTypeBuilderImpl(
        PascalString userName, UserConfigurationMask userConfiguration, PascalString description) {
      this.userName = userName;
      this.userConfiguration = userConfiguration;
      this.description = description;
    }

    public UserManagementDataType build() {
      UserManagementDataType userManagementDataType =
          new UserManagementDataType(userName, userConfiguration, description);
      return userManagementDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof UserManagementDataType)) {
      return false;
    }
    UserManagementDataType that = (UserManagementDataType) o;
    return (getUserName() == that.getUserName())
        && (getUserConfiguration() == that.getUserConfiguration())
        && (getDescription() == that.getDescription())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUserName(), getUserConfiguration(), getDescription());
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
