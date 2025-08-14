// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPMtlsAuthType The definition of the `HTTPMtlsAuth` object.
type HTTPMtlsAuthType string

// List of HTTPMtlsAuthType.
const (
	HTTPMTLSAUTHTYPE_HTTPMTLSAUTH HTTPMtlsAuthType = "HTTPMtlsAuth"
)

var allowedHTTPMtlsAuthTypeEnumValues = []HTTPMtlsAuthType{
	HTTPMTLSAUTHTYPE_HTTPMTLSAUTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HTTPMtlsAuthType) GetAllowedValues() []HTTPMtlsAuthType {
	return allowedHTTPMtlsAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HTTPMtlsAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HTTPMtlsAuthType(value)
	return nil
}

// NewHTTPMtlsAuthTypeFromValue returns a pointer to a valid HTTPMtlsAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHTTPMtlsAuthTypeFromValue(v string) (*HTTPMtlsAuthType, error) {
	ev := HTTPMtlsAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HTTPMtlsAuthType: valid values are %v", v, allowedHTTPMtlsAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HTTPMtlsAuthType) IsValid() bool {
	for _, existing := range allowedHTTPMtlsAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HTTPMtlsAuthType value.
func (v HTTPMtlsAuthType) Ptr() *HTTPMtlsAuthType {
	return &v
}
