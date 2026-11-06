// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthPersonaMappingType Resource type returned by the existing Custom OIDC persona mapping API.
type OIDCCloudAuthPersonaMappingType string

// List of OIDCCloudAuthPersonaMappingType.
const (
	OIDCCLOUDAUTHPERSONAMAPPINGTYPE_CUSTOM_OIDC_IDENTITY_MAPPING OIDCCloudAuthPersonaMappingType = "custom_oidc_identity_mapping"
)

var allowedOIDCCloudAuthPersonaMappingTypeEnumValues = []OIDCCloudAuthPersonaMappingType{
	OIDCCLOUDAUTHPERSONAMAPPINGTYPE_CUSTOM_OIDC_IDENTITY_MAPPING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OIDCCloudAuthPersonaMappingType) GetAllowedValues() []OIDCCloudAuthPersonaMappingType {
	return allowedOIDCCloudAuthPersonaMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OIDCCloudAuthPersonaMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OIDCCloudAuthPersonaMappingType(value)
	return nil
}

// NewOIDCCloudAuthPersonaMappingTypeFromValue returns a pointer to a valid OIDCCloudAuthPersonaMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOIDCCloudAuthPersonaMappingTypeFromValue(v string) (*OIDCCloudAuthPersonaMappingType, error) {
	ev := OIDCCloudAuthPersonaMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OIDCCloudAuthPersonaMappingType: valid values are %v", v, allowedOIDCCloudAuthPersonaMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OIDCCloudAuthPersonaMappingType) IsValid() bool {
	for _, existing := range allowedOIDCCloudAuthPersonaMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OIDCCloudAuthPersonaMappingType value.
func (v OIDCCloudAuthPersonaMappingType) Ptr() *OIDCCloudAuthPersonaMappingType {
	return &v
}
