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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetConfirmedServiceRequestConfirmedTextMessage extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.CONFIRMED_TEXT_MESSAGE;
  }

  // Properties.
  protected final BACnetContextTagObjectIdentifier textMessageSourceDevice;
  protected final BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass messageClass;
  protected final BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
      messagePriority;
  protected final BACnetContextTagCharacterString message;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestConfirmedTextMessage(
      BACnetContextTagObjectIdentifier textMessageSourceDevice,
      BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass messageClass,
      BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged messagePriority,
      BACnetContextTagCharacterString message,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.textMessageSourceDevice = textMessageSourceDevice;
    this.messageClass = messageClass;
    this.messagePriority = messagePriority;
    this.message = message;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagObjectIdentifier getTextMessageSourceDevice() {
    return textMessageSourceDevice;
  }

  public BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass getMessageClass() {
    return messageClass;
  }

  public BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
      getMessagePriority() {
    return messagePriority;
  }

  public BACnetContextTagCharacterString getMessage() {
    return message;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestConfirmedTextMessage");

    // Simple Field (textMessageSourceDevice)
    writeSimpleField(
        "textMessageSourceDevice",
        textMessageSourceDevice,
        new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (messageClass) (Can be skipped, if the value is null)
    writeOptionalField("messageClass", messageClass, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (messagePriority)
    writeSimpleField(
        "messagePriority", messagePriority, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (message)
    writeSimpleField("message", message, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestConfirmedTextMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestConfirmedTextMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (textMessageSourceDevice)
    lengthInBits += textMessageSourceDevice.getLengthInBits();

    // Optional Field (messageClass)
    if (messageClass != null) {
      lengthInBits += messageClass.getLengthInBits();
    }

    // Simple field (messagePriority)
    lengthInBits += messagePriority.getLengthInBits();

    // Simple field (message)
    lengthInBits += message.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestConfirmedTextMessage");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier textMessageSourceDevice =
        readSimpleField(
            "textMessageSourceDevice",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass messageClass =
        readOptionalField(
            "messageClass",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.staticParse(
                        readBuffer, (short) (1)),
                readBuffer));

    BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged messagePriority =
        readSimpleField(
            "messagePriority",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
                        .staticParse(
                            readBuffer, (short) (2), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagCharacterString message =
        readSimpleField(
            "message",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestConfirmedTextMessage");
    // Create the instance
    return new BACnetConfirmedServiceRequestConfirmedTextMessageBuilderImpl(
        textMessageSourceDevice, messageClass, messagePriority, message, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestConfirmedTextMessageBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagObjectIdentifier textMessageSourceDevice;
    private final BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass messageClass;
    private final BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
        messagePriority;
    private final BACnetContextTagCharacterString message;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestConfirmedTextMessageBuilderImpl(
        BACnetContextTagObjectIdentifier textMessageSourceDevice,
        BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass messageClass,
        BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged messagePriority,
        BACnetContextTagCharacterString message,
        Long serviceRequestLength) {
      this.textMessageSourceDevice = textMessageSourceDevice;
      this.messageClass = messageClass;
      this.messagePriority = messagePriority;
      this.message = message;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestConfirmedTextMessage build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestConfirmedTextMessage
          bACnetConfirmedServiceRequestConfirmedTextMessage =
              new BACnetConfirmedServiceRequestConfirmedTextMessage(
                  textMessageSourceDevice,
                  messageClass,
                  messagePriority,
                  message,
                  serviceRequestLength);
      return bACnetConfirmedServiceRequestConfirmedTextMessage;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestConfirmedTextMessage)) {
      return false;
    }
    BACnetConfirmedServiceRequestConfirmedTextMessage that =
        (BACnetConfirmedServiceRequestConfirmedTextMessage) o;
    return (getTextMessageSourceDevice() == that.getTextMessageSourceDevice())
        && (getMessageClass() == that.getMessageClass())
        && (getMessagePriority() == that.getMessagePriority())
        && (getMessage() == that.getMessage())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getTextMessageSourceDevice(),
        getMessageClass(),
        getMessagePriority(),
        getMessage());
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
