// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthOrgUUIDSource Source of the organization UUID in the signed token. The `aud` source accepts a datadog/{org_uuid} audience.
type OIDCCloudAuthOrgUUIDSource string

// List of OIDCCloudAuthOrgUUIDSource.
const (
	OIDCCLOUDAUTHORGUUIDSOURCE_DD_ORG_UUID OIDCCloudAuthOrgUUIDSource = "dd_org_uuid"
	OIDCCLOUDAUTHORGUUIDSOURCE_AUD         OIDCCloudAuthOrgUUIDSource = "aud"
)

var allowedOIDCCloudAuthOrgUUIDSourceEnumValues = []OIDCCloudAuthOrgUUIDSource{
	OIDCCLOUDAUTHORGUUIDSOURCE_DD_ORG_UUID,
	OIDCCLOUDAUTHORGUUIDSOURCE_AUD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OIDCCloudAuthOrgUUIDSource) GetAllowedValues() []OIDCCloudAuthOrgUUIDSource {
	return allowedOIDCCloudAuthOrgUUIDSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OIDCCloudAuthOrgUUIDSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OIDCCloudAuthOrgUUIDSource(value)
	return nil
}

// NewOIDCCloudAuthOrgUUIDSourceFromValue returns a pointer to a valid OIDCCloudAuthOrgUUIDSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOIDCCloudAuthOrgUUIDSourceFromValue(v string) (*OIDCCloudAuthOrgUUIDSource, error) {
	ev := OIDCCloudAuthOrgUUIDSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OIDCCloudAuthOrgUUIDSource: valid values are %v", v, allowedOIDCCloudAuthOrgUUIDSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OIDCCloudAuthOrgUUIDSource) IsValid() bool {
	for _, existing := range allowedOIDCCloudAuthOrgUUIDSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OIDCCloudAuthOrgUUIDSource value.
func (v OIDCCloudAuthOrgUUIDSource) Ptr() *OIDCCloudAuthOrgUUIDSource {
	return &v
}
