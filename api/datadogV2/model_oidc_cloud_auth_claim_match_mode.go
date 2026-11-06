// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthClaimMatchMode Case-sensitive matching mode. Exact compares strings. Glob matches the entire value: * matches any number of characters, including /, and ? matches one character. Other characters are literal; escaping is not supported.
type OIDCCloudAuthClaimMatchMode string

// List of OIDCCloudAuthClaimMatchMode.
const (
	OIDCCLOUDAUTHCLAIMMATCHMODE_EXACT OIDCCloudAuthClaimMatchMode = "exact"
	OIDCCLOUDAUTHCLAIMMATCHMODE_GLOB  OIDCCloudAuthClaimMatchMode = "glob"
)

var allowedOIDCCloudAuthClaimMatchModeEnumValues = []OIDCCloudAuthClaimMatchMode{
	OIDCCLOUDAUTHCLAIMMATCHMODE_EXACT,
	OIDCCLOUDAUTHCLAIMMATCHMODE_GLOB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OIDCCloudAuthClaimMatchMode) GetAllowedValues() []OIDCCloudAuthClaimMatchMode {
	return allowedOIDCCloudAuthClaimMatchModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OIDCCloudAuthClaimMatchMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OIDCCloudAuthClaimMatchMode(value)
	return nil
}

// NewOIDCCloudAuthClaimMatchModeFromValue returns a pointer to a valid OIDCCloudAuthClaimMatchMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOIDCCloudAuthClaimMatchModeFromValue(v string) (*OIDCCloudAuthClaimMatchMode, error) {
	ev := OIDCCloudAuthClaimMatchMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OIDCCloudAuthClaimMatchMode: valid values are %v", v, allowedOIDCCloudAuthClaimMatchModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OIDCCloudAuthClaimMatchMode) IsValid() bool {
	for _, existing := range allowedOIDCCloudAuthClaimMatchModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OIDCCloudAuthClaimMatchMode value.
func (v OIDCCloudAuthClaimMatchMode) Ptr() *OIDCCloudAuthClaimMatchMode {
	return &v
}
