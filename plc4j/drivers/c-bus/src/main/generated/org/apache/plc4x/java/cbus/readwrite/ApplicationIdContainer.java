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

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum ApplicationIdContainer {
  RESERVED_00((short) 0x00, LightingCompatible.NA, ApplicationId.RESERVED),
  FREE_USAGE_01((short) 0x01, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_02((short) 0x02, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_03((short) 0x03, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_04((short) 0x04, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_05((short) 0x05, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_06((short) 0x06, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_07((short) 0x07, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_08((short) 0x08, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_09((short) 0x09, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_0A((short) 0x0A, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_0B((short) 0x0B, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_0C((short) 0x0C, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_0D((short) 0x0D, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_0E((short) 0x0E, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  FREE_USAGE_0F((short) 0x0F, LightingCompatible.NA, ApplicationId.FREE_USAGE),
  RESERVED_10((short) 0x10, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_11((short) 0x11, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_12((short) 0x12, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_13((short) 0x13, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_14((short) 0x14, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_15((short) 0x15, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_16((short) 0x16, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_17((short) 0x17, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_18((short) 0x18, LightingCompatible.NA, ApplicationId.RESERVED),
  TEMPERATURE_BROADCAST_19(
      (short) 0x19, LightingCompatible.NO, ApplicationId.TEMPERATURE_BROADCAST),
  RESERVED_1A((short) 0x1A, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_1B((short) 0x1B, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_1C((short) 0x1C, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_1D((short) 0x1D, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_1E((short) 0x1E, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_1F((short) 0x1F, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_20((short) 0x20, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_21((short) 0x21, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_22((short) 0x22, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_23((short) 0x23, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_24((short) 0x24, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_25((short) 0x25, LightingCompatible.NA, ApplicationId.RESERVED),
  ROOM_CONTROL_SYSTEM_26((short) 0x26, LightingCompatible.YES, ApplicationId.ROOM_CONTROL_SYSTEM),
  RESERVED_27((short) 0x27, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_28((short) 0x28, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_29((short) 0x29, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_2A((short) 0x2A, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_2B((short) 0x2B, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_2C((short) 0x2C, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_2D((short) 0x2D, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_2E((short) 0x2E, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_2F((short) 0x2F, LightingCompatible.NA, ApplicationId.RESERVED),
  LIGHTING_30((short) 0x30, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_31((short) 0x31, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_32((short) 0x32, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_33((short) 0x33, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_34((short) 0x34, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_35((short) 0x35, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_36((short) 0x36, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_37((short) 0x37, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_38((short) 0x38, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_39((short) 0x39, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_3A((short) 0x3A, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_3B((short) 0x3B, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_3C((short) 0x3C, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_3D((short) 0x3D, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_3E((short) 0x3E, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_3F((short) 0x3F, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_40((short) 0x40, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_41((short) 0x41, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_42((short) 0x42, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_43((short) 0x43, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_44((short) 0x44, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_45((short) 0x45, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_46((short) 0x46, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_47((short) 0x47, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_48((short) 0x48, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_49((short) 0x49, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_4A((short) 0x4A, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_4B((short) 0x4B, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_4C((short) 0x4C, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_4D((short) 0x4D, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_4E((short) 0x4E, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_4F((short) 0x4F, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_50((short) 0x50, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_51((short) 0x51, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_52((short) 0x52, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_53((short) 0x53, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_54((short) 0x54, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_55((short) 0x55, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_56((short) 0x56, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_57((short) 0x57, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_58((short) 0x58, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_59((short) 0x59, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_5A((short) 0x5A, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_5B((short) 0x5B, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_5C((short) 0x5C, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_5D((short) 0x5D, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_5E((short) 0x5E, LightingCompatible.YES, ApplicationId.LIGHTING),
  LIGHTING_5F((short) 0x5F, LightingCompatible.YES, ApplicationId.LIGHTING),
  RESERVED_60((short) 0x60, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_61((short) 0x61, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_62((short) 0x62, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_63((short) 0x63, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_64((short) 0x64, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_65((short) 0x65, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_66((short) 0x66, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_67((short) 0x67, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_68((short) 0x68, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_69((short) 0x69, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_6A((short) 0x6A, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_6B((short) 0x6B, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_6C((short) 0x6C, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_6D((short) 0x6D, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_6E((short) 0x6E, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_6F((short) 0x6F, LightingCompatible.NA, ApplicationId.RESERVED),
  VENTILATION_70((short) 0x70, LightingCompatible.YES, ApplicationId.VENTILATION),
  IRRIGATION_CONTROL_71((short) 0x71, LightingCompatible.YES, ApplicationId.IRRIGATION_CONTROL),
  POOLS_SPAS_PONDS_FOUNTAINS_CONTROL_72(
      (short) 0x72, LightingCompatible.YES, ApplicationId.POOLS_SPAS_PONDS_FOUNTAINS_CONTROL),
  HVAC_ACTUATOR_73((short) 0x73, LightingCompatible.NA, ApplicationId.HVAC_ACTUATOR),
  HVAC_ACTUATOR_74((short) 0x74, LightingCompatible.NA, ApplicationId.HVAC_ACTUATOR),
  RESERVED_75((short) 0x75, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_76((short) 0x76, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_77((short) 0x77, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_78((short) 0x78, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_79((short) 0x79, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_7A((short) 0x7A, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_7B((short) 0x7B, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_7C((short) 0x7C, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_7D((short) 0x7D, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_7E((short) 0x7E, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_7F((short) 0x7F, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_80((short) 0x80, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_81((short) 0x81, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_82((short) 0x82, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_83((short) 0x83, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_84((short) 0x84, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_85((short) 0x85, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_86((short) 0x86, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_87((short) 0x87, LightingCompatible.NA, ApplicationId.RESERVED),
  HEATING_88((short) 0x88, LightingCompatible.YES, ApplicationId.HEATING),
  RESERVED_89((short) 0x89, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_8A((short) 0x8A, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_8B((short) 0x8B, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_8C((short) 0x8C, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_8D((short) 0x8D, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_8E((short) 0x8E, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_8F((short) 0x8F, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_90((short) 0x90, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_91((short) 0x91, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_92((short) 0x92, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_93((short) 0x93, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_94((short) 0x94, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_95((short) 0x95, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_96((short) 0x96, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_97((short) 0x97, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_98((short) 0x98, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_99((short) 0x99, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_9A((short) 0x9A, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_9B((short) 0x9B, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_9C((short) 0x9C, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_9D((short) 0x9D, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_9E((short) 0x9E, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_9F((short) 0x9F, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A0((short) 0xA0, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A1((short) 0xA1, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A2((short) 0xA2, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A3((short) 0xA3, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A4((short) 0xA4, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A5((short) 0xA5, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A6((short) 0xA6, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A7((short) 0xA7, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A8((short) 0xA8, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_A9((short) 0xA9, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_AA((short) 0xAA, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_AB((short) 0xAB, LightingCompatible.NA, ApplicationId.RESERVED),
  AIR_CONDITIONING_AC((short) 0xAC, LightingCompatible.NO, ApplicationId.AIR_CONDITIONING),
  INFO_MESSAGES((short) 0xAD, LightingCompatible.NA, ApplicationId.INFO_MESSAGES),
  RESERVED_AE((short) 0xAE, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_AF((short) 0xAF, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B0((short) 0xB0, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B1((short) 0xB1, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B2((short) 0xB2, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B3((short) 0xB3, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B4((short) 0xB4, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B5((short) 0xB5, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B6((short) 0xB6, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B7((short) 0xB7, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B8((short) 0xB8, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_B9((short) 0xB9, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_BA((short) 0xBA, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_BB((short) 0xBB, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_BC((short) 0xBC, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_BD((short) 0xBD, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_BE((short) 0xBE, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_BF((short) 0xBF, LightingCompatible.NA, ApplicationId.RESERVED),
  MEDIA_TRANSPORT_CONTROL_C0(
      (short) 0xC0, LightingCompatible.NA, ApplicationId.MEDIA_TRANSPORT_CONTROL),
  RESERVED_C1((short) 0xC1, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C2((short) 0xC2, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C3((short) 0xC3, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C4((short) 0xC4, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C5((short) 0xC5, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C6((short) 0xC6, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C7((short) 0xC7, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C8((short) 0xC8, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_C9((short) 0xC9, LightingCompatible.NA, ApplicationId.RESERVED),
  TRIGGER_CONTROL_CA(
      (short) 0xCA, LightingCompatible.YES_BUT_RESTRICTIONS, ApplicationId.TRIGGER_CONTROL),
  ENABLE_CONTROL_CB(
      (short) 0xCB, LightingCompatible.YES_BUT_RESTRICTIONS, ApplicationId.ENABLE_CONTROL),
  I_HAVE_NO_IDEA_CC((short) 0xCC, LightingCompatible.NA, ApplicationId.RESERVED),
  AUDIO_AND_VIDEO_CD(
      (short) 0xCD, LightingCompatible.YES_BUT_RESTRICTIONS, ApplicationId.AUDIO_AND_VIDEO),
  ERROR_REPORTING_CE((short) 0xCE, LightingCompatible.NA, ApplicationId.ERROR_REPORTING),
  RESERVED_CF((short) 0xCF, LightingCompatible.NA, ApplicationId.RESERVED),
  SECURITY_D0((short) 0xD0, LightingCompatible.NO, ApplicationId.SECURITY),
  METERING_D1((short) 0xD1, LightingCompatible.NO, ApplicationId.METERING),
  RESERVED_D2((short) 0xD2, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_D3((short) 0xD3, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_D4((short) 0xD4, LightingCompatible.NA, ApplicationId.RESERVED),
  ACCESS_CONTROL_D5((short) 0xD5, LightingCompatible.NO, ApplicationId.ACCESS_CONTROL),
  RESERVED_D6((short) 0xD6, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_D7((short) 0xD7, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_D8((short) 0xD8, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_D9((short) 0xD9, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_DA((short) 0xDA, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_DB((short) 0xDB, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_DC((short) 0xDC, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_DD((short) 0xDD, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_DE((short) 0xDE, LightingCompatible.NA, ApplicationId.RESERVED),
  CLOCK_AND_TIMEKEEPING_DF(
      (short) 0xDF, LightingCompatible.NO, ApplicationId.CLOCK_AND_TIMEKEEPING),
  TELEPHONY_STATUS_AND_CONTROL_E0(
      (short) 0xE0, LightingCompatible.NO, ApplicationId.TELEPHONY_STATUS_AND_CONTROL),
  RESERVED_E1((short) 0xE1, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_E2((short) 0xE2, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_E3((short) 0xE3, LightingCompatible.NA, ApplicationId.RESERVED),
  MEASUREMENT_E4((short) 0xE4, LightingCompatible.NO, ApplicationId.MEASUREMENT),
  RESERVED_E5((short) 0xE5, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_E6((short) 0xE6, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_E7((short) 0xE7, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_E8((short) 0xE8, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_E9((short) 0xE9, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_EA((short) 0xEA, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_EB((short) 0xEB, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_EC((short) 0xEC, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_ED((short) 0xED, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_EE((short) 0xEE, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_EF((short) 0xEF, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F0((short) 0xF0, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F1((short) 0xF1, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F2((short) 0xF2, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F3((short) 0xF3, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F4((short) 0xF4, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F5((short) 0xF5, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F6((short) 0xF6, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F7((short) 0xF7, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F8((short) 0xF8, LightingCompatible.NA, ApplicationId.RESERVED),
  RESERVED_F9((short) 0xF9, LightingCompatible.NA, ApplicationId.RESERVED),
  TESTING_FA((short) 0xFA, LightingCompatible.NA, ApplicationId.TESTING),
  RESERVED_FB((short) 0xFB, LightingCompatible.NO, ApplicationId.RESERVED),
  RESERVED_FC((short) 0xFC, LightingCompatible.NO, ApplicationId.RESERVED),
  RESERVED_FD((short) 0xFD, LightingCompatible.NO, ApplicationId.RESERVED),
  RESERVED_FE((short) 0xFE, LightingCompatible.NO, ApplicationId.RESERVED),
  NETWORK_CONTROL((short) 0xFF, LightingCompatible.NO, ApplicationId.NETWORK_CONTROL);
  private static final Map<Short, ApplicationIdContainer> map;

  static {
    map = new HashMap<>();
    for (ApplicationIdContainer value : ApplicationIdContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;
  private LightingCompatible lightingCompatible;
  private ApplicationId applicationId;

  ApplicationIdContainer(
      short value, LightingCompatible lightingCompatible, ApplicationId applicationId) {
    this.value = value;
    this.lightingCompatible = lightingCompatible;
    this.applicationId = applicationId;
  }

  public short getValue() {
    return value;
  }

  public LightingCompatible getLightingCompatible() {
    return lightingCompatible;
  }

  public static ApplicationIdContainer firstEnumForFieldLightingCompatible(
      LightingCompatible fieldValue) {
    for (ApplicationIdContainer _val : ApplicationIdContainer.values()) {
      if (_val.getLightingCompatible() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<ApplicationIdContainer> enumsForFieldLightingCompatible(
      LightingCompatible fieldValue) {
    List<ApplicationIdContainer> _values = new ArrayList();
    for (ApplicationIdContainer _val : ApplicationIdContainer.values()) {
      if (_val.getLightingCompatible() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public ApplicationId getApplicationId() {
    return applicationId;
  }

  public static ApplicationIdContainer firstEnumForFieldApplicationId(ApplicationId fieldValue) {
    for (ApplicationIdContainer _val : ApplicationIdContainer.values()) {
      if (_val.getApplicationId() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<ApplicationIdContainer> enumsForFieldApplicationId(ApplicationId fieldValue) {
    List<ApplicationIdContainer> _values = new ArrayList();
    for (ApplicationIdContainer _val : ApplicationIdContainer.values()) {
      if (_val.getApplicationId() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static ApplicationIdContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
