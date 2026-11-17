// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountAPITokenAuthType The authentication method type.
type CoreWeaveCloudIntegrationAccountAPITokenAuthType string

// List of CoreWeaveCloudIntegrationAccountAPITokenAuthType.
const (
	COREWEAVECLOUDINTEGRATIONACCOUNTAPITOKENAUTHTYPE_COREWEAVE_API_TOKEN CoreWeaveCloudIntegrationAccountAPITokenAuthType = "coreweave_api_token"
)

var allowedCoreWeaveCloudIntegrationAccountAPITokenAuthTypeEnumValues = []CoreWeaveCloudIntegrationAccountAPITokenAuthType{
	COREWEAVECLOUDINTEGRATIONACCOUNTAPITOKENAUTHTYPE_COREWEAVE_API_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CoreWeaveCloudIntegrationAccountAPITokenAuthType) GetAllowedValues() []CoreWeaveCloudIntegrationAccountAPITokenAuthType {
	return allowedCoreWeaveCloudIntegrationAccountAPITokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CoreWeaveCloudIntegrationAccountAPITokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CoreWeaveCloudIntegrationAccountAPITokenAuthType(value)
	return nil
}

// NewCoreWeaveCloudIntegrationAccountAPITokenAuthTypeFromValue returns a pointer to a valid CoreWeaveCloudIntegrationAccountAPITokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCoreWeaveCloudIntegrationAccountAPITokenAuthTypeFromValue(v string) (*CoreWeaveCloudIntegrationAccountAPITokenAuthType, error) {
	ev := CoreWeaveCloudIntegrationAccountAPITokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CoreWeaveCloudIntegrationAccountAPITokenAuthType: valid values are %v", v, allowedCoreWeaveCloudIntegrationAccountAPITokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CoreWeaveCloudIntegrationAccountAPITokenAuthType) IsValid() bool {
	for _, existing := range allowedCoreWeaveCloudIntegrationAccountAPITokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CoreWeaveCloudIntegrationAccountAPITokenAuthType value.
func (v CoreWeaveCloudIntegrationAccountAPITokenAuthType) Ptr() *CoreWeaveCloudIntegrationAccountAPITokenAuthType {
	return &v
}
