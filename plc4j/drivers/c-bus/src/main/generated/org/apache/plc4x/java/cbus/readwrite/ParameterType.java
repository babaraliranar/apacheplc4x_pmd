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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum ParameterType {
  UNKNOWN((short) 0),
  APPLICATION_ADDRESS_1((short) 1),
  APPLICATION_ADDRESS_2((short) 2),
  INTERFACE_OPTIONS_1((short) 3),
  INTERFACE_OPTIONS_2((short) 4),
  INTERFACE_OPTIONS_3((short) 5),
  BAUD_RATE_SELECTOR((short) 6),
  INTERFACE_OPTIONS_1_POWER_UP_SETTINGS((short) 7),
  CUSTOM_MANUFACTURER((short) 8),
  SERIAL_NUMBER((short) 9),
  CUSTOM_TYPE((short) 10);
  private static final Map<Short, ParameterType> map;

  static {
    map = new HashMap<>();
    for (ParameterType value : ParameterType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  ParameterType(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static ParameterType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
