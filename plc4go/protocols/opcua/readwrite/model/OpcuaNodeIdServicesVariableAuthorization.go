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

// OpcuaNodeIdServicesVariableAuthorization is an enum
type OpcuaNodeIdServicesVariableAuthorization int32

type IOpcuaNodeIdServicesVariableAuthorization interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceCertificate                                OpcuaNodeIdServicesVariableAuthorization = 17860
	OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceUri                                        OpcuaNodeIdServicesVariableAuthorization = 18072
	OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_IssuerEndpointUrl                                 OpcuaNodeIdServicesVariableAuthorization = 18073
	OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri         OpcuaNodeIdServicesVariableAuthorization = 23558
	OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate OpcuaNodeIdServicesVariableAuthorization = 23559
	OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl  OpcuaNodeIdServicesVariableAuthorization = 23560
)

var OpcuaNodeIdServicesVariableAuthorizationValues []OpcuaNodeIdServicesVariableAuthorization

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAuthorizationValues = []OpcuaNodeIdServicesVariableAuthorization{
		OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceCertificate,
		OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceUri,
		OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_IssuerEndpointUrl,
		OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri,
		OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate,
		OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl,
	}
}

func OpcuaNodeIdServicesVariableAuthorizationByValue(value int32) (enum OpcuaNodeIdServicesVariableAuthorization, ok bool) {
	switch value {
	case 17860:
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceCertificate, true
	case 18072:
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceUri, true
	case 18073:
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_IssuerEndpointUrl, true
	case 23558:
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri, true
	case 23559:
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate, true
	case 23560:
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAuthorizationByName(value string) (enum OpcuaNodeIdServicesVariableAuthorization, ok bool) {
	switch value {
	case "AuthorizationServiceConfigurationType_ServiceCertificate":
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceCertificate, true
	case "AuthorizationServiceConfigurationType_ServiceUri":
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceUri, true
	case "AuthorizationServiceConfigurationType_IssuerEndpointUrl":
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_IssuerEndpointUrl, true
	case "AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri":
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri, true
	case "AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate":
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate, true
	case "AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl":
		return OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAuthorizationKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAuthorizationValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableAuthorization(structType any) OpcuaNodeIdServicesVariableAuthorization {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAuthorization {
		if sOpcuaNodeIdServicesVariableAuthorization, ok := typ.(OpcuaNodeIdServicesVariableAuthorization); ok {
			return sOpcuaNodeIdServicesVariableAuthorization
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAuthorization) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAuthorization) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAuthorizationParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAuthorization, error) {
	return OpcuaNodeIdServicesVariableAuthorizationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAuthorizationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAuthorization, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAuthorization", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAuthorization")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAuthorizationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAuthorization")
		return OpcuaNodeIdServicesVariableAuthorization(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAuthorization) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAuthorization) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAuthorization", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAuthorization) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceCertificate:
		return "AuthorizationServiceConfigurationType_ServiceCertificate"
	case OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_ServiceUri:
		return "AuthorizationServiceConfigurationType_ServiceUri"
	case OpcuaNodeIdServicesVariableAuthorization_AuthorizationServiceConfigurationType_IssuerEndpointUrl:
		return "AuthorizationServiceConfigurationType_IssuerEndpointUrl"
	case OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri:
		return "AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceUri"
	case OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate:
		return "AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_ServiceCertificate"
	case OpcuaNodeIdServicesVariableAuthorization_AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl:
		return "AuthorizationServicesConfigurationFolderType_ServiceName_Placeholder_IssuerEndpointUrl"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAuthorization) String() string {
	return e.PLC4XEnumName()
}
