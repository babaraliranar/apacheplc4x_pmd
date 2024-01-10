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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableNode {
  NodeVersion((int) 3068L),
  NodeClass_EnumValues((int) 11878L),
  NodeAttributesMask_EnumValues((int) 11881L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableNode> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableNode value : OpcuaNodeIdServicesVariableNode.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableNode(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableNode enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
