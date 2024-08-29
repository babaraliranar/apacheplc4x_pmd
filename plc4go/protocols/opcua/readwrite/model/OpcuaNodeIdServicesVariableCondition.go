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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableCondition is an enum
type OpcuaNodeIdServicesVariableCondition int32

type IOpcuaNodeIdServicesVariableCondition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableCondition_ConditionType_Retain                               OpcuaNodeIdServicesVariableCondition = 3874
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh_InputArguments      OpcuaNodeIdServicesVariableCondition = 3876
	OpcuaNodeIdServicesVariableCondition_ConditionVariableType_SourceTimestamp              OpcuaNodeIdServicesVariableCondition = 9003
	OpcuaNodeIdServicesVariableCondition_ConditionRefreshMethodType_InputArguments          OpcuaNodeIdServicesVariableCondition = 9008
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionName                        OpcuaNodeIdServicesVariableCondition = 9009
	OpcuaNodeIdServicesVariableCondition_ConditionType_BranchId                             OpcuaNodeIdServicesVariableCondition = 9010
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState                         OpcuaNodeIdServicesVariableCondition = 9011
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Id                      OpcuaNodeIdServicesVariableCondition = 9012
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Name                    OpcuaNodeIdServicesVariableCondition = 9013
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Number                  OpcuaNodeIdServicesVariableCondition = 9014
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveDisplayName    OpcuaNodeIdServicesVariableCondition = 9015
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TransitionTime          OpcuaNodeIdServicesVariableCondition = 9016
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveTransitionTime OpcuaNodeIdServicesVariableCondition = 9017
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TrueState               OpcuaNodeIdServicesVariableCondition = 9018
	OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_FalseState              OpcuaNodeIdServicesVariableCondition = 9019
	OpcuaNodeIdServicesVariableCondition_ConditionType_Quality                              OpcuaNodeIdServicesVariableCondition = 9020
	OpcuaNodeIdServicesVariableCondition_ConditionType_Quality_SourceTimestamp              OpcuaNodeIdServicesVariableCondition = 9021
	OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity                         OpcuaNodeIdServicesVariableCondition = 9022
	OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity_SourceTimestamp         OpcuaNodeIdServicesVariableCondition = 9023
	OpcuaNodeIdServicesVariableCondition_ConditionType_Comment                              OpcuaNodeIdServicesVariableCondition = 9024
	OpcuaNodeIdServicesVariableCondition_ConditionType_Comment_SourceTimestamp              OpcuaNodeIdServicesVariableCondition = 9025
	OpcuaNodeIdServicesVariableCondition_ConditionType_ClientUserId                         OpcuaNodeIdServicesVariableCondition = 9026
	OpcuaNodeIdServicesVariableCondition_ConditionType_AddComment_InputArguments            OpcuaNodeIdServicesVariableCondition = 9030
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassId                     OpcuaNodeIdServicesVariableCondition = 11112
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassName                   OpcuaNodeIdServicesVariableCondition = 11113
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh2_InputArguments     OpcuaNodeIdServicesVariableCondition = 12913
	OpcuaNodeIdServicesVariableCondition_ConditionRefresh2MethodType_InputArguments         OpcuaNodeIdServicesVariableCondition = 12915
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassId                  OpcuaNodeIdServicesVariableCondition = 16363
	OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassName                OpcuaNodeIdServicesVariableCondition = 16364
	OpcuaNodeIdServicesVariableCondition_ConditionType_SupportsFilteredRetain               OpcuaNodeIdServicesVariableCondition = 32060
)

var OpcuaNodeIdServicesVariableConditionValues []OpcuaNodeIdServicesVariableCondition

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableConditionValues = []OpcuaNodeIdServicesVariableCondition{
		OpcuaNodeIdServicesVariableCondition_ConditionType_Retain,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh_InputArguments,
		OpcuaNodeIdServicesVariableCondition_ConditionVariableType_SourceTimestamp,
		OpcuaNodeIdServicesVariableCondition_ConditionRefreshMethodType_InputArguments,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionName,
		OpcuaNodeIdServicesVariableCondition_ConditionType_BranchId,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Id,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Name,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Number,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TransitionTime,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TrueState,
		OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_FalseState,
		OpcuaNodeIdServicesVariableCondition_ConditionType_Quality,
		OpcuaNodeIdServicesVariableCondition_ConditionType_Quality_SourceTimestamp,
		OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity,
		OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity_SourceTimestamp,
		OpcuaNodeIdServicesVariableCondition_ConditionType_Comment,
		OpcuaNodeIdServicesVariableCondition_ConditionType_Comment_SourceTimestamp,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ClientUserId,
		OpcuaNodeIdServicesVariableCondition_ConditionType_AddComment_InputArguments,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassId,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassName,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh2_InputArguments,
		OpcuaNodeIdServicesVariableCondition_ConditionRefresh2MethodType_InputArguments,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassName,
		OpcuaNodeIdServicesVariableCondition_ConditionType_SupportsFilteredRetain,
	}
}

func OpcuaNodeIdServicesVariableConditionByValue(value int32) (enum OpcuaNodeIdServicesVariableCondition, ok bool) {
	switch value {
	case 11112:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassId, true
	case 11113:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassName, true
	case 12913:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh2_InputArguments, true
	case 12915:
		return OpcuaNodeIdServicesVariableCondition_ConditionRefresh2MethodType_InputArguments, true
	case 16363:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassId, true
	case 16364:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassName, true
	case 32060:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_SupportsFilteredRetain, true
	case 3874:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Retain, true
	case 3876:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh_InputArguments, true
	case 9003:
		return OpcuaNodeIdServicesVariableCondition_ConditionVariableType_SourceTimestamp, true
	case 9008:
		return OpcuaNodeIdServicesVariableCondition_ConditionRefreshMethodType_InputArguments, true
	case 9009:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionName, true
	case 9010:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_BranchId, true
	case 9011:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState, true
	case 9012:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Id, true
	case 9013:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Name, true
	case 9014:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Number, true
	case 9015:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveDisplayName, true
	case 9016:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TransitionTime, true
	case 9017:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveTransitionTime, true
	case 9018:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TrueState, true
	case 9019:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_FalseState, true
	case 9020:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Quality, true
	case 9021:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Quality_SourceTimestamp, true
	case 9022:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity, true
	case 9023:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity_SourceTimestamp, true
	case 9024:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Comment, true
	case 9025:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Comment_SourceTimestamp, true
	case 9026:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ClientUserId, true
	case 9030:
		return OpcuaNodeIdServicesVariableCondition_ConditionType_AddComment_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableConditionByName(value string) (enum OpcuaNodeIdServicesVariableCondition, ok bool) {
	switch value {
	case "ConditionType_ConditionClassId":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassId, true
	case "ConditionType_ConditionClassName":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassName, true
	case "ConditionType_ConditionRefresh2_InputArguments":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh2_InputArguments, true
	case "ConditionRefresh2MethodType_InputArguments":
		return OpcuaNodeIdServicesVariableCondition_ConditionRefresh2MethodType_InputArguments, true
	case "ConditionType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassId, true
	case "ConditionType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassName, true
	case "ConditionType_SupportsFilteredRetain":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_SupportsFilteredRetain, true
	case "ConditionType_Retain":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Retain, true
	case "ConditionType_ConditionRefresh_InputArguments":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh_InputArguments, true
	case "ConditionVariableType_SourceTimestamp":
		return OpcuaNodeIdServicesVariableCondition_ConditionVariableType_SourceTimestamp, true
	case "ConditionRefreshMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableCondition_ConditionRefreshMethodType_InputArguments, true
	case "ConditionType_ConditionName":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionName, true
	case "ConditionType_BranchId":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_BranchId, true
	case "ConditionType_EnabledState":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState, true
	case "ConditionType_EnabledState_Id":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Id, true
	case "ConditionType_EnabledState_Name":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Name, true
	case "ConditionType_EnabledState_Number":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Number, true
	case "ConditionType_EnabledState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveDisplayName, true
	case "ConditionType_EnabledState_TransitionTime":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TransitionTime, true
	case "ConditionType_EnabledState_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveTransitionTime, true
	case "ConditionType_EnabledState_TrueState":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TrueState, true
	case "ConditionType_EnabledState_FalseState":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_FalseState, true
	case "ConditionType_Quality":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Quality, true
	case "ConditionType_Quality_SourceTimestamp":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Quality_SourceTimestamp, true
	case "ConditionType_LastSeverity":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity, true
	case "ConditionType_LastSeverity_SourceTimestamp":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity_SourceTimestamp, true
	case "ConditionType_Comment":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Comment, true
	case "ConditionType_Comment_SourceTimestamp":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_Comment_SourceTimestamp, true
	case "ConditionType_ClientUserId":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_ClientUserId, true
	case "ConditionType_AddComment_InputArguments":
		return OpcuaNodeIdServicesVariableCondition_ConditionType_AddComment_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableConditionKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableConditionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableCondition(structType any) OpcuaNodeIdServicesVariableCondition {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableCondition {
		if sOpcuaNodeIdServicesVariableCondition, ok := typ.(OpcuaNodeIdServicesVariableCondition); ok {
			return sOpcuaNodeIdServicesVariableCondition
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableCondition) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableCondition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableConditionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableCondition, error) {
	return OpcuaNodeIdServicesVariableConditionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableConditionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableCondition, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableCondition", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableCondition")
	}
	if enum, ok := OpcuaNodeIdServicesVariableConditionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableCondition")
		return OpcuaNodeIdServicesVariableCondition(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableCondition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableCondition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableCondition", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableCondition) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassId:
		return "ConditionType_ConditionClassId"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionClassName:
		return "ConditionType_ConditionClassName"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh2_InputArguments:
		return "ConditionType_ConditionRefresh2_InputArguments"
	case OpcuaNodeIdServicesVariableCondition_ConditionRefresh2MethodType_InputArguments:
		return "ConditionRefresh2MethodType_InputArguments"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassId:
		return "ConditionType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionSubClassName:
		return "ConditionType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_SupportsFilteredRetain:
		return "ConditionType_SupportsFilteredRetain"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_Retain:
		return "ConditionType_Retain"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionRefresh_InputArguments:
		return "ConditionType_ConditionRefresh_InputArguments"
	case OpcuaNodeIdServicesVariableCondition_ConditionVariableType_SourceTimestamp:
		return "ConditionVariableType_SourceTimestamp"
	case OpcuaNodeIdServicesVariableCondition_ConditionRefreshMethodType_InputArguments:
		return "ConditionRefreshMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ConditionName:
		return "ConditionType_ConditionName"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_BranchId:
		return "ConditionType_BranchId"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState:
		return "ConditionType_EnabledState"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Id:
		return "ConditionType_EnabledState_Id"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Name:
		return "ConditionType_EnabledState_Name"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_Number:
		return "ConditionType_EnabledState_Number"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveDisplayName:
		return "ConditionType_EnabledState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TransitionTime:
		return "ConditionType_EnabledState_TransitionTime"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_EffectiveTransitionTime:
		return "ConditionType_EnabledState_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_TrueState:
		return "ConditionType_EnabledState_TrueState"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_EnabledState_FalseState:
		return "ConditionType_EnabledState_FalseState"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_Quality:
		return "ConditionType_Quality"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_Quality_SourceTimestamp:
		return "ConditionType_Quality_SourceTimestamp"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity:
		return "ConditionType_LastSeverity"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_LastSeverity_SourceTimestamp:
		return "ConditionType_LastSeverity_SourceTimestamp"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_Comment:
		return "ConditionType_Comment"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_Comment_SourceTimestamp:
		return "ConditionType_Comment_SourceTimestamp"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_ClientUserId:
		return "ConditionType_ClientUserId"
	case OpcuaNodeIdServicesVariableCondition_ConditionType_AddComment_InputArguments:
		return "ConditionType_AddComment_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableCondition) String() string {
	return e.PLC4XEnumName()
}
