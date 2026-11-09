// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubCredentials - Supported GitHub connection credentials.
type GithubCredentials struct {
	GithubApp                 *GithubApp
	GithubPersonalAccessToken *GithubPersonalAccessToken

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GithubAppAsGithubCredentials is a convenience function that returns GithubApp wrapped in GithubCredentials.
func GithubAppAsGithubCredentials(v *GithubApp) GithubCredentials {
	return GithubCredentials{GithubApp: v}
}

// GithubPersonalAccessTokenAsGithubCredentials is a convenience function that returns GithubPersonalAccessToken wrapped in GithubCredentials.
func GithubPersonalAccessTokenAsGithubCredentials(v *GithubPersonalAccessToken) GithubCredentials {
	return GithubCredentials{GithubPersonalAccessToken: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GithubCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GithubApp
	err = datadog.Unmarshal(data, &obj.GithubApp)
	if err == nil {
		if obj.GithubApp != nil && obj.GithubApp.UnparsedObject == nil {
			jsonGithubApp, _ := datadog.Marshal(obj.GithubApp)
			if string(jsonGithubApp) == "{}" { // empty struct
				obj.GithubApp = nil
			} else {
				match++
			}
		} else {
			obj.GithubApp = nil
		}
	} else {
		obj.GithubApp = nil
	}

	// try to unmarshal data into GithubPersonalAccessToken
	err = datadog.Unmarshal(data, &obj.GithubPersonalAccessToken)
	if err == nil {
		if obj.GithubPersonalAccessToken != nil && obj.GithubPersonalAccessToken.UnparsedObject == nil {
			jsonGithubPersonalAccessToken, _ := datadog.Marshal(obj.GithubPersonalAccessToken)
			if string(jsonGithubPersonalAccessToken) == "{}" { // empty struct
				obj.GithubPersonalAccessToken = nil
			} else {
				match++
			}
		} else {
			obj.GithubPersonalAccessToken = nil
		}
	} else {
		obj.GithubPersonalAccessToken = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GithubApp = nil
		obj.GithubPersonalAccessToken = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GithubCredentials) MarshalJSON() ([]byte, error) {
	if obj.GithubApp != nil {
		return datadog.Marshal(&obj.GithubApp)
	}

	if obj.GithubPersonalAccessToken != nil {
		return datadog.Marshal(&obj.GithubPersonalAccessToken)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GithubCredentials) GetActualInstance() interface{} {
	if obj.GithubApp != nil {
		return obj.GithubApp
	}

	if obj.GithubPersonalAccessToken != nil {
		return obj.GithubPersonalAccessToken
	}

	// all schemas are nil
	return nil
}
