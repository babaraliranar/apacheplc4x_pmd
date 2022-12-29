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

public enum TsnFailureCode {
  tsnFailureCodeNoFailure((long) 0L),
  tsnFailureCodeInsufficientBandwidth((long) 1L),
  tsnFailureCodeInsufficientResources((long) 2L),
  tsnFailureCodeInsufficientTrafficClassBandwidth((long) 3L),
  tsnFailureCodeStreamIdInUse((long) 4L),
  tsnFailureCodeStreamDestinationAddressInUse((long) 5L),
  tsnFailureCodeStreamPreemptedByHigherRank((long) 6L),
  tsnFailureCodeLatencyHasChanged((long) 7L),
  tsnFailureCodeEgressPortNotAvbCapable((long) 8L),
  tsnFailureCodeUseDifferentDestinationAddress((long) 9L),
  tsnFailureCodeOutOfMsrpResources((long) 10L),
  tsnFailureCodeOutOfMmrpResources((long) 11L),
  tsnFailureCodeCannotStoreDestinationAddress((long) 12L),
  tsnFailureCodePriorityIsNotAnSrcClass((long) 13L),
  tsnFailureCodeMaxFrameSizeTooLarge((long) 14L),
  tsnFailureCodeMaxFanInPortsLimitReached((long) 15L),
  tsnFailureCodeFirstValueChangedForStreamId((long) 16L),
  tsnFailureCodeVlanBlockedOnEgress((long) 17L),
  tsnFailureCodeVlanTaggingDisabledOnEgress((long) 18L),
  tsnFailureCodeSrClassPriorityMismatch((long) 19L),
  tsnFailureCodeFeatureNotPropagated((long) 20L),
  tsnFailureCodeMaxLatencyExceeded((long) 21L),
  tsnFailureCodeBridgeDoesNotProvideNetworkId((long) 22L),
  tsnFailureCodeStreamTransformNotSupported((long) 23L),
  tsnFailureCodeStreamIdTypeNotSupported((long) 24L),
  tsnFailureCodeFeatureNotSupported((long) 25L);
  private static final Map<Long, TsnFailureCode> map;

  static {
    map = new HashMap<>();
    for (TsnFailureCode value : TsnFailureCode.values()) {
      map.put((long) value.getValue(), value);
    }
  }

  private long value;

  TsnFailureCode(long value) {
    this.value = value;
  }

  public long getValue() {
    return value;
  }

  public static TsnFailureCode enumForValue(long value) {
    return map.get(value);
  }

  public static Boolean isDefined(long value) {
    return map.containsKey(value);
  }
}
