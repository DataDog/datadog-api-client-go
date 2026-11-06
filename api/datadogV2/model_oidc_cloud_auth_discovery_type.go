// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthDiscoveryType How the signing keys are discovered.
type OIDCCloudAuthDiscoveryType string

// List of OIDCCloudAuthDiscoveryType.
const (
	OIDCCLOUDAUTHDISCOVERYTYPE_OIDC_DISCOVERY OIDCCloudAuthDiscoveryType = "oidc_discovery"
	OIDCCLOUDAUTHDISCOVERYTYPE_DIRECT_JWKS    OIDCCloudAuthDiscoveryType = "direct_jwks"
)

var allowedOIDCCloudAuthDiscoveryTypeEnumValues = []OIDCCloudAuthDiscoveryType{
	OIDCCLOUDAUTHDISCOVERYTYPE_OIDC_DISCOVERY,
	OIDCCLOUDAUTHDISCOVERYTYPE_DIRECT_JWKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OIDCCloudAuthDiscoveryType) GetAllowedValues() []OIDCCloudAuthDiscoveryType {
	return allowedOIDCCloudAuthDiscoveryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OIDCCloudAuthDiscoveryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OIDCCloudAuthDiscoveryType(value)
	return nil
}

// NewOIDCCloudAuthDiscoveryTypeFromValue returns a pointer to a valid OIDCCloudAuthDiscoveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOIDCCloudAuthDiscoveryTypeFromValue(v string) (*OIDCCloudAuthDiscoveryType, error) {
	ev := OIDCCloudAuthDiscoveryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OIDCCloudAuthDiscoveryType: valid values are %v", v, allowedOIDCCloudAuthDiscoveryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OIDCCloudAuthDiscoveryType) IsValid() bool {
	for _, existing := range allowedOIDCCloudAuthDiscoveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OIDCCloudAuthDiscoveryType value.
func (v OIDCCloudAuthDiscoveryType) Ptr() *OIDCCloudAuthDiscoveryType {
	return &v
}
