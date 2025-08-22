// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuePlatform Platform associated with the issue
type IssuePlatform string

// List of IssuePlatform.
const (
	ISSUEPLATFORM_BROWSER      IssuePlatform = "browser"
	ISSUEPLATFORM_ANDROID      IssuePlatform = "android"
	ISSUEPLATFORM_BACKEND      IssuePlatform = "backend"
	ISSUEPLATFORM_IOS          IssuePlatform = "ios"
	ISSUEPLATFORM_REACT_NATIVE IssuePlatform = "react_native"
	ISSUEPLATFORM_FLUTTER      IssuePlatform = "flutter"
	ISSUEPLATFORM_ROKU         IssuePlatform = "roku"
)

var allowedIssuePlatformEnumValues = []IssuePlatform{
	ISSUEPLATFORM_BROWSER,
	ISSUEPLATFORM_ANDROID,
	ISSUEPLATFORM_BACKEND,
	ISSUEPLATFORM_IOS,
	ISSUEPLATFORM_REACT_NATIVE,
	ISSUEPLATFORM_FLUTTER,
	ISSUEPLATFORM_ROKU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssuePlatform) GetAllowedValues() []IssuePlatform {
	return allowedIssuePlatformEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssuePlatform) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssuePlatform(value)
	return nil
}

// NewIssuePlatformFromValue returns a pointer to a valid IssuePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssuePlatformFromValue(v string) (*IssuePlatform, error) {
	ev := IssuePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssuePlatform: valid values are %v", v, allowedIssuePlatformEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssuePlatform) IsValid() bool {
	for _, existing := range allowedIssuePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuePlatform value.
func (v IssuePlatform) Ptr() *IssuePlatform {
	return &v
}
