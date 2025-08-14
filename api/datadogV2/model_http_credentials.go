// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPCredentials - The definition of the `HTTPCredentials` object.
type HTTPCredentials struct {
	HTTPTokenAuth *HTTPTokenAuth
	HTTPBasicAuth *HTTPBasicAuth
	HTTPMtlsAuth  *HTTPMtlsAuth

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HTTPTokenAuthAsHTTPCredentials is a convenience function that returns HTTPTokenAuth wrapped in HTTPCredentials.
func HTTPTokenAuthAsHTTPCredentials(v *HTTPTokenAuth) HTTPCredentials {
	return HTTPCredentials{HTTPTokenAuth: v}
}

// HTTPBasicAuthAsHTTPCredentials is a convenience function that returns HTTPBasicAuth wrapped in HTTPCredentials.
func HTTPBasicAuthAsHTTPCredentials(v *HTTPBasicAuth) HTTPCredentials {
	return HTTPCredentials{HTTPBasicAuth: v}
}

// HTTPMtlsAuthAsHTTPCredentials is a convenience function that returns HTTPMtlsAuth wrapped in HTTPCredentials.
func HTTPMtlsAuthAsHTTPCredentials(v *HTTPMtlsAuth) HTTPCredentials {
	return HTTPCredentials{HTTPMtlsAuth: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HTTPCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HTTPTokenAuth
	err = datadog.Unmarshal(data, &obj.HTTPTokenAuth)
	if err == nil {
		if obj.HTTPTokenAuth != nil && obj.HTTPTokenAuth.UnparsedObject == nil {
			jsonHTTPTokenAuth, _ := datadog.Marshal(obj.HTTPTokenAuth)
			if string(jsonHTTPTokenAuth) == "{}" { // empty struct
				obj.HTTPTokenAuth = nil
			} else {
				match++
			}
		} else {
			obj.HTTPTokenAuth = nil
		}
	} else {
		obj.HTTPTokenAuth = nil
	}

	// try to unmarshal data into HTTPBasicAuth
	err = datadog.Unmarshal(data, &obj.HTTPBasicAuth)
	if err == nil {
		if obj.HTTPBasicAuth != nil && obj.HTTPBasicAuth.UnparsedObject == nil {
			jsonHTTPBasicAuth, _ := datadog.Marshal(obj.HTTPBasicAuth)
			if string(jsonHTTPBasicAuth) == "{}" { // empty struct
				obj.HTTPBasicAuth = nil
			} else {
				match++
			}
		} else {
			obj.HTTPBasicAuth = nil
		}
	} else {
		obj.HTTPBasicAuth = nil
	}

	// try to unmarshal data into HTTPMtlsAuth
	err = datadog.Unmarshal(data, &obj.HTTPMtlsAuth)
	if err == nil {
		if obj.HTTPMtlsAuth != nil && obj.HTTPMtlsAuth.UnparsedObject == nil {
			jsonHTTPMtlsAuth, _ := datadog.Marshal(obj.HTTPMtlsAuth)
			if string(jsonHTTPMtlsAuth) == "{}" { // empty struct
				obj.HTTPMtlsAuth = nil
			} else {
				match++
			}
		} else {
			obj.HTTPMtlsAuth = nil
		}
	} else {
		obj.HTTPMtlsAuth = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HTTPTokenAuth = nil
		obj.HTTPBasicAuth = nil
		obj.HTTPMtlsAuth = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HTTPCredentials) MarshalJSON() ([]byte, error) {
	if obj.HTTPTokenAuth != nil {
		return datadog.Marshal(&obj.HTTPTokenAuth)
	}

	if obj.HTTPBasicAuth != nil {
		return datadog.Marshal(&obj.HTTPBasicAuth)
	}

	if obj.HTTPMtlsAuth != nil {
		return datadog.Marshal(&obj.HTTPMtlsAuth)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HTTPCredentials) GetActualInstance() interface{} {
	if obj.HTTPTokenAuth != nil {
		return obj.HTTPTokenAuth
	}

	if obj.HTTPBasicAuth != nil {
		return obj.HTTPBasicAuth
	}

	if obj.HTTPMtlsAuth != nil {
		return obj.HTTPMtlsAuth
	}

	// all schemas are nil
	return nil
}
