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

public class UserTokenPolicy extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "306";
  }

  // Properties.
  protected final PascalString policyId;
  protected final UserTokenType tokenType;
  protected final PascalString issuedTokenType;
  protected final PascalString issuerEndpointUrl;
  protected final PascalString securityPolicyUri;

  public UserTokenPolicy(
      PascalString policyId,
      UserTokenType tokenType,
      PascalString issuedTokenType,
      PascalString issuerEndpointUrl,
      PascalString securityPolicyUri) {
    super();
    this.policyId = policyId;
    this.tokenType = tokenType;
    this.issuedTokenType = issuedTokenType;
    this.issuerEndpointUrl = issuerEndpointUrl;
    this.securityPolicyUri = securityPolicyUri;
  }

  public PascalString getPolicyId() {
    return policyId;
  }

  public UserTokenType getTokenType() {
    return tokenType;
  }

  public PascalString getIssuedTokenType() {
    return issuedTokenType;
  }

  public PascalString getIssuerEndpointUrl() {
    return issuerEndpointUrl;
  }

  public PascalString getSecurityPolicyUri() {
    return securityPolicyUri;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("UserTokenPolicy");

    // Simple Field (policyId)
    writeSimpleField("policyId", policyId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (tokenType)
    writeSimpleEnumField(
        "tokenType",
        "UserTokenType",
        tokenType,
        new DataWriterEnumDefault<>(
            UserTokenType::getValue, UserTokenType::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (issuedTokenType)
    writeSimpleField(
        "issuedTokenType", issuedTokenType, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (issuerEndpointUrl)
    writeSimpleField(
        "issuerEndpointUrl", issuerEndpointUrl, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (securityPolicyUri)
    writeSimpleField(
        "securityPolicyUri", securityPolicyUri, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("UserTokenPolicy");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    UserTokenPolicy _value = this;

    // Simple field (policyId)
    lengthInBits += policyId.getLengthInBits();

    // Simple field (tokenType)
    lengthInBits += 32;

    // Simple field (issuedTokenType)
    lengthInBits += issuedTokenType.getLengthInBits();

    // Simple field (issuerEndpointUrl)
    lengthInBits += issuerEndpointUrl.getLengthInBits();

    // Simple field (securityPolicyUri)
    lengthInBits += securityPolicyUri.getLengthInBits();

    return lengthInBits;
  }

  public static UserTokenPolicyBuilder staticParseBuilder(ReadBuffer readBuffer, String identifier)
      throws ParseException {
    readBuffer.pullContext("UserTokenPolicy");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    PascalString policyId =
        readSimpleField(
            "policyId",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    UserTokenType tokenType =
        readEnumField(
            "tokenType",
            "UserTokenType",
            new DataReaderEnumDefault<>(
                UserTokenType::enumForValue, readUnsignedLong(readBuffer, 32)));

    PascalString issuedTokenType =
        readSimpleField(
            "issuedTokenType",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString issuerEndpointUrl =
        readSimpleField(
            "issuerEndpointUrl",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString securityPolicyUri =
        readSimpleField(
            "securityPolicyUri",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("UserTokenPolicy");
    // Create the instance
    return new UserTokenPolicyBuilder(
        policyId, tokenType, issuedTokenType, issuerEndpointUrl, securityPolicyUri);
  }

  public static class UserTokenPolicyBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString policyId;
    private final UserTokenType tokenType;
    private final PascalString issuedTokenType;
    private final PascalString issuerEndpointUrl;
    private final PascalString securityPolicyUri;

    public UserTokenPolicyBuilder(
        PascalString policyId,
        UserTokenType tokenType,
        PascalString issuedTokenType,
        PascalString issuerEndpointUrl,
        PascalString securityPolicyUri) {

      this.policyId = policyId;
      this.tokenType = tokenType;
      this.issuedTokenType = issuedTokenType;
      this.issuerEndpointUrl = issuerEndpointUrl;
      this.securityPolicyUri = securityPolicyUri;
    }

    public UserTokenPolicy build() {
      UserTokenPolicy userTokenPolicy =
          new UserTokenPolicy(
              policyId, tokenType, issuedTokenType, issuerEndpointUrl, securityPolicyUri);
      return userTokenPolicy;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof UserTokenPolicy)) {
      return false;
    }
    UserTokenPolicy that = (UserTokenPolicy) o;
    return (getPolicyId() == that.getPolicyId())
        && (getTokenType() == that.getTokenType())
        && (getIssuedTokenType() == that.getIssuedTokenType())
        && (getIssuerEndpointUrl() == that.getIssuerEndpointUrl())
        && (getSecurityPolicyUri() == that.getSecurityPolicyUri())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getPolicyId(),
        getTokenType(),
        getIssuedTokenType(),
        getIssuerEndpointUrl(),
        getSecurityPolicyUri());
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
