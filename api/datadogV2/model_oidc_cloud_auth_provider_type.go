// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthProviderType Resource type.
type OIDCCloudAuthProviderType string

// List of OIDCCloudAuthProviderType.
const (
	OIDCCLOUDAUTHPROVIDERTYPE_CUSTOM_OIDC_PROVIDER_CONFIG OIDCCloudAuthProviderType = "custom_oidc_provider_config"
)

var allowedOIDCCloudAuthProviderTypeEnumValues = []OIDCCloudAuthProviderType{
	OIDCCLOUDAUTHPROVIDERTYPE_CUSTOM_OIDC_PROVIDER_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OIDCCloudAuthProviderType) GetAllowedValues() []OIDCCloudAuthProviderType {
	return allowedOIDCCloudAuthProviderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OIDCCloudAuthProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OIDCCloudAuthProviderType(value)
	return nil
}

// NewOIDCCloudAuthProviderTypeFromValue returns a pointer to a valid OIDCCloudAuthProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOIDCCloudAuthProviderTypeFromValue(v string) (*OIDCCloudAuthProviderType, error) {
	ev := OIDCCloudAuthProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OIDCCloudAuthProviderType: valid values are %v", v, allowedOIDCCloudAuthProviderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OIDCCloudAuthProviderType) IsValid() bool {
	for _, existing := range allowedOIDCCloudAuthProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OIDCCloudAuthProviderType value.
func (v OIDCCloudAuthProviderType) Ptr() *OIDCCloudAuthProviderType {
	return &v
}
