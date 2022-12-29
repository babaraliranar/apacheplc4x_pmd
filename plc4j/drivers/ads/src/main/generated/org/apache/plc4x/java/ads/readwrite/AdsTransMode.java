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
package org.apache.plc4x.java.ads.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum AdsTransMode {
  NONE((long) 0L),
  CLIENT_CYCLE((long) 1L),
  CLIENT_ON_CHANGE((long) 2L),
  CYCLIC((long) 3L),
  ON_CHANGE((long) 4L),
  CYCLIC_IN_CONTEXT((long) 5L),
  ON_CHANGE_IN_CONTEXT((long) 6L);
  private static final Map<Long, AdsTransMode> map;

  static {
    map = new HashMap<>();
    for (AdsTransMode value : AdsTransMode.values()) {
      map.put((long) value.getValue(), value);
    }
  }

  private long value;

  AdsTransMode(long value) {
    this.value = value;
  }

  public long getValue() {
    return value;
  }

  public static AdsTransMode enumForValue(long value) {
    return map.get(value);
  }

  public static Boolean isDefined(long value) {
    return map.containsKey(value);
  }
}
